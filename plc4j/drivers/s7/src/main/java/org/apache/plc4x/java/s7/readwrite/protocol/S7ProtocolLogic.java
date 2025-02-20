/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.s7.readwrite.protocol;

import io.netty.buffer.ByteBuf;
import io.netty.buffer.ByteBufUtil;
import io.netty.buffer.Unpooled;
import org.apache.commons.lang3.concurrent.BasicThreadFactory;
import org.apache.commons.lang3.tuple.MutablePair;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.model.PlcSubscriptionHandle;
import org.apache.plc4x.java.api.model.PlcSubscriptionTag;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.s7.events.S7CyclicEvent;
import org.apache.plc4x.java.s7.readwrite.*;
import org.apache.plc4x.java.s7.readwrite.connection.S7HMuxImpl;
import org.apache.plc4x.java.s7.readwrite.context.S7DriverContext;
import org.apache.plc4x.java.s7.readwrite.tag.*;
import org.apache.plc4x.java.s7.readwrite.types.S7ControllerType;
import org.apache.plc4x.java.s7.readwrite.types.S7SubscriptionType;
import org.apache.plc4x.java.s7.readwrite.utils.S7PlcSubscriptionHandle;
import org.apache.plc4x.java.s7.utils.S7ParamErrorCode;
import org.apache.plc4x.java.spi.ConversationContext;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.spi.context.DriverContext;
import org.apache.plc4x.java.spi.generation.*;
import org.apache.plc4x.java.spi.messages.*;
import org.apache.plc4x.java.spi.messages.utils.ResponseItem;
import org.apache.plc4x.java.spi.model.DefaultPlcSubscriptionTag;
import org.apache.plc4x.java.spi.transaction.RequestTransactionManager;
import org.apache.plc4x.java.spi.values.PlcList;
import org.apache.plc4x.java.spi.values.PlcNull;
import org.apache.plc4x.java.spi.values.PlcSINT;
import org.apache.plc4x.java.spi.values.PlcValueHandler;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.nio.ByteBuffer;
import java.nio.charset.Charset;
import java.time.Duration;
import java.util.*;
import java.util.concurrent.*;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.BiConsumer;
import java.util.function.Consumer;
import java.util.stream.IntStream;

/**
 * The S7 Protocol states that there can not be more then {min(maxAmqCaller, maxAmqCallee} "ongoing" requests.
 * So we need to limit those.
 * Thus, each request goes to a Work Queue and this Queue ensures, that only 3 are open at the same time.
 */
public class S7ProtocolLogic extends Plc4xProtocolBase<TPKTPacket> {

    private final Logger logger = LoggerFactory.getLogger(S7ProtocolLogic.class);

    public static final Duration REQUEST_TIMEOUT = Duration.ofMillis(10000);
    private final AtomicInteger tpduGenerator = new AtomicInteger(10);

    /*
     * Task group for managing connection redundancy.
     */
    private final ExecutorService clientExecutorService = Executors.newFixedThreadPool(4, new BasicThreadFactory.Builder()
        .namingPattern("plc4x-app-thread-%d")
        .daemon(true)
        .priority(Thread.MAX_PRIORITY)
        .build());

    /*
     * Take into account that the size of this buffer depends on the final device.
     * S7-300 goes from 20 to 300 and for S7-400 it goes from 300 to 10000.
     * Depending on the configuration of the alarm system, a large number of
     * them should be expected when starting the connection.
     * (Examples of this are PCS7 and Braumat).
     * Alarm filtering, ack, etc. must be performed by the client application.
     */
    private final BlockingQueue<Object> eventQueue = new ArrayBlockingQueue<>(1024);
    private final S7ProtocolEventLogic EventLogic = new S7ProtocolEventLogic(eventQueue);
    private final S7PlcSubscriptionHandle modeHandle = new S7PlcSubscriptionHandle(EventType.MODE, EventLogic);
    private final S7PlcSubscriptionHandle sysHandle = new S7PlcSubscriptionHandle(EventType.SYS, EventLogic);
    private final S7PlcSubscriptionHandle usrHandle = new S7PlcSubscriptionHandle(EventType.USR, EventLogic);
    private final S7PlcSubscriptionHandle almHandle = new S7PlcSubscriptionHandle(EventType.ALM, EventLogic);
    private final S7PlcSubscriptionHandle cycHandle = new S7PlcSubscriptionHandle(EventType.CYC, EventLogic);

    /*
     * For the reconnection functionality by a "TimeOut" of the connection,
     * you must keep track of open transactions. In general, an S7 device
     * supports a couple of simultaneous requests.
     * The rhythm of execution must be determined by the TransactionManager.
     * So far it is the way to indicate to the user that he must redo
     * his request.
     */
    private final HashMap<Object, MutablePair<RequestTransactionManager.RequestTransaction, Object>> activeRequests = new HashMap<>();


    private S7DriverContext s7DriverContext;
    private RequestTransactionManager tm;

    @Override
    public void setDriverContext(DriverContext driverContext) {
        super.setDriverContext(driverContext);
        this.s7DriverContext = (S7DriverContext) driverContext;

        // Initialize Transaction Manager.
        // Until the number of concurrent requests is successfully negotiated we set it to a
        // maximum of only one request being able to be sent at a time. During the login process
        // No concurrent requests can be sent anyway. It will be updated when receiving the
        // S7ParameterSetupCommunication response.
        this.tm = new RequestTransactionManager(1);
        EventLogic.start();
    }

    @Override
    public void onConnect(ConversationContext<TPKTPacket> context) {
        if (context.isPassive()) {
            logger.info("S7 Driver running in PASSIVE mode.");
            s7DriverContext.setPassiveMode(true);
            // No login required, just confirm that we're connected.
            context.fireConnected();
            return;
        }

        //Set feature for all handlers in the pipeline from the driver configuration.
        setChannelFeatures();

        // Only the TCP transport supports login.
        logger.info("S7 Driver running in ACTIVE mode.");
        logger.debug("Sending COTP Connection Request");
        // Open the session on ISO Transport Protocol first.
        TPKTPacket packet = new TPKTPacket(createCOTPConnectionRequest(
            s7DriverContext.getCalledTsapId(), s7DriverContext.getCallingTsapId(), s7DriverContext.getCotpTpduSize()));

        context.sendRequest(packet)
            .onTimeout(e -> {
                logger.warn("Timeout during Connection establishing, closing channel...");
                context.getChannel().close();
            })
            .expectResponse(TPKTPacket.class, REQUEST_TIMEOUT)
            .check(p -> p.getPayload() instanceof COTPPacketConnectionResponse)
            .unwrap(p -> (COTPPacketConnectionResponse) p.getPayload())
            .handle(cotpPacketConnectionResponse -> {
                logger.debug("Got COTP Connection Response");
                logger.debug("Sending S7 Connection Request");
                context.sendRequest(createS7ConnectionRequest(cotpPacketConnectionResponse))
                    .onTimeout(e -> {
                        logger.warn("Timeout during Connection establishing, closing channel...");
                        context.getChannel().close();
                    })
                    .expectResponse(TPKTPacket.class, REQUEST_TIMEOUT)
                    .unwrap(TPKTPacket::getPayload)
                    .only(COTPPacketData.class)
                    .unwrap(COTPPacket::getPayload)
                    .only(S7MessageResponseData.class)
                    .unwrap(S7Message::getParameter)
                    .only(S7ParameterSetupCommunication.class)
                    .handle(setupCommunication -> {
                        logger.debug("Got S7 Connection Response");
                        // Save some data from the response.
                        s7DriverContext.setMaxAmqCaller(setupCommunication.getMaxAmqCaller());
                        s7DriverContext.setMaxAmqCallee(setupCommunication.getMaxAmqCallee());
                        s7DriverContext.setPduSize(setupCommunication.getPduLength());

                        // Update the number of concurrent requests to the negotiated number.
                        // I have never seen anything else than equal values for caller and
                        // callee, but if they were different, we're only limiting the outgoing
                        // requests.
                        tm.setNumberOfConcurrentRequests(s7DriverContext.getMaxAmqCallee());

                        // If the controller type is explicitly set, were finished with the login
                        // process. If it's set to ANY, we have to query the serial number information
                        // in order to detect the type of PLC.
                        if (s7DriverContext.getControllerType() != S7ControllerType.ANY) {
                            // Send an event that connection setup is complete.
                            context.fireConnected();
                            return;
                        }

                        // Prepare a message to request the remote to identify itself.
                        logger.debug("Sending S7 Identification Request");
                        TPKTPacket tpktPacket = createIdentifyRemoteMessage();
                        context.sendRequest(tpktPacket)
                            .onTimeout(e -> {
                                logger.warn("Timeout during Connection establishing, closing channel...");
                                context.getChannel().close();
                            })
                            .expectResponse(TPKTPacket.class, REQUEST_TIMEOUT)
                            .check(p -> p.getPayload() instanceof COTPPacketData)
                            .unwrap(p -> ((COTPPacketData) p.getPayload()))
                            .check(p -> p.getPayload() instanceof S7MessageUserData)
                            .unwrap(p -> ((S7MessageUserData) p.getPayload()))
                            .check(p -> p.getPayload() instanceof S7PayloadUserData)
                            .handle(messageUserData -> {
                                logger.debug("Got S7 Identification Response");
                                S7PayloadUserData payloadUserData = (S7PayloadUserData) messageUserData.getPayload();
                                extractControllerTypeAndFireConnected(context, payloadUserData);
                                cleanFutures();
                            });
                    });
            });
    }


    /*
     * It performs the sequential and safe shutdown of the driver.
     * Completion of pending requests, executors and associated tasks.
     */
    @Override
    public void onDisconnect(ConversationContext<TPKTPacket> context) {
        //1. Clear all pending requests and their associated transaction          
        cleanFutures();
        //2. Here we shut down the local task executor.
        clientExecutorService.shutdown();
        //3. Performs the shutdown of the transaction executor.
        tm.shutdown();
        //4. Finish the execution of the tasks for the handling of Events. 
        EventLogic.stop();
        //5. Executes the closing of the main channel.
        context.getChannel().close();
        //6. Here is the stop of any task or state machine that is added.        
    }


    @Override
    public CompletableFuture<PlcReadResponse> read(PlcReadRequest readRequest) {
        if (!isConnected()) {
            CompletableFuture<PlcReadResponse> future = new CompletableFuture<>();
            future.completeExceptionally(new PlcRuntimeException("Disconnected"));
            return future;
        }
        DefaultPlcReadRequest request = (DefaultPlcReadRequest) readRequest;
        List<S7VarRequestParameterItem> requestItems = new ArrayList<>(request.getNumberOfTags());

        if (request.getTags().get(0) instanceof S7SzlTag) {
            S7SzlTag szltag = (S7SzlTag) request.getTags().get(0);

            final S7MessageUserData s7SzlMessageRequest = new S7MessageUserData(1, new S7ParameterUserData(List.of(
                new S7ParameterUserDataItemCPUFunctions((short) 0x11, (byte) 0x4, (byte) 0x4, (short) 0x01, (short) 0x00, null, null, null)
            )), new S7PayloadUserData(List.of(
                new S7PayloadUserDataItemCpuFunctionReadSzlRequest(DataTransportErrorCode.OK,
                    DataTransportSize.OCTET_STRING,
                    0x04,
                    new SzlId(SzlModuleTypeClass.enumForValue((byte) ((szltag.getSzlId() & 0xf000) >> 12)),
                        (byte) ((szltag.getSzlId() & 0x0f00) >> 8),
                        SzlSublist.enumForValue((short) (szltag.getSzlId() & 0x00ff))),
                    szltag.getIndex())
            )));

            return toPlcReadResponse(readRequest, readInternal(s7SzlMessageRequest));

        } else if (request.getTags().get(0) instanceof S7AckTag) {

            List<S7ParameterUserDataItem> parameterItems = new ArrayList<>(request.getNumberOfTags());
            List<S7PayloadUserDataItem> payloadItems = new ArrayList<>(request.getNumberOfTags());
            encodeAlarmAckRequest(request, parameterItems, payloadItems);
            final S7MessageUserData s7MessageRequest = new S7MessageUserData(-1,
                new S7ParameterUserData(parameterItems),
                new S7PayloadUserData(payloadItems));
            return toPlcReadResponse(readRequest, readInternal(s7MessageRequest));

        }


        for (PlcTag tag : request.getTags()) {
            requestItems.add(new S7VarRequestParameterItemAddress(encodeS7Address(tag)));
        }

        // Create a read request template.
        // tpuId will be inserted before sending in #readInternal so we insert -1 as dummy here
        final S7MessageRequest s7MessageRequest = new S7MessageRequest(-1,
            new S7ParameterReadVarRequest(requestItems),
            null);

        // Just send a single response and chain it as Response
        return toPlcReadResponse(readRequest, readInternal(s7MessageRequest));
    }

    /**
     * Maps the S7ReadResponse of a PlcReadRequest to a PlcReadResponse
     */
    private CompletableFuture<PlcReadResponse> toPlcReadResponse(PlcReadRequest readRequest, CompletableFuture<S7Message> response) {
        CompletableFuture<PlcReadResponse> clientFuture = new CompletableFuture<>();
        activeRequests.get(response).setRight(clientFuture);

        try {
            clientExecutorService.execute(() -> {
                try {
                    PlcReadResponse plcItems = (PlcReadResponse) decodeReadResponse(response.get(), readRequest);
                    clientFuture.complete(plcItems);
                } catch (Exception e) {
                    logger.info("uh", e);
                }
            });
        } catch (Exception e) {
            logger.info("uh", e);
        }

        return clientFuture;

// TODO: whoever out commented this describe why it is out commented and describe what the above does different
//        return response
//            .thenApply(p -> {
//                try {
//                    return ((PlcReadResponse) decodeReadResponse(p, readRequest));
//                } catch (PlcProtocolException e) {
//                    throw new PlcRuntimeException("Unable to decode Response", e);
//                }
//            });
    }

    /**
     * Sends one Read over the Wire and internally returns the Response
     * Do sending of normally sized single-message request.
     * <p>
     * Assumes that the {@link S7MessageRequest} and its expected {@link S7MessageResponseData}
     * and does not further check that!
     */
    private CompletableFuture<S7Message> readInternal(S7Message request) {
        CompletableFuture<S7Message> future = new CompletableFuture<>();
        int tpduId = tpduGenerator.getAndIncrement();
        // If we've reached the max value for a 16 bit transaction identifier, reset back to 1
        tpduGenerator.compareAndExchange(0xFFFF, 1);

        S7Message message = (request instanceof S7MessageUserData) ?
            new S7MessageUserData(tpduId, request.getParameter(), request.getPayload()) :
            new S7MessageRequest(tpduId, request.getParameter(), request.getPayload());

        TPKTPacket tpktPacket = new TPKTPacket(new COTPPacketData(null, message, true, (byte) tpduId));

        // Start a new request-transaction (Is ended in the response-handler)
        RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
        transaction.submit(() -> context.sendRequest(tpktPacket)
            .onTimeout(new TransactionErrorCallback<>(future, transaction))
            .onError(new TransactionErrorCallback<>(future, transaction))
            .expectResponse(TPKTPacket.class, REQUEST_TIMEOUT)
            .check(p -> p.getPayload() instanceof COTPPacketData)
            .unwrap(p -> (COTPPacketData) p.getPayload())
            .check(p -> p.getPayload() != null)
            .unwrap(COTPPacket::getPayload)
            .check(p -> p.getTpduReference() == tpduId)
            .handle(p -> {
                activeRequests.remove(future);
                future.complete(p);
                // Finish the request-transaction.
                transaction.endRequest();
            }));
        activeRequests.put(future, new MutablePair<>(transaction, null));

        return future;
    }


    @Override
    public CompletableFuture<PlcWriteResponse> write(PlcWriteRequest writeRequest) {
        CompletableFuture<PlcWriteResponse> future = new CompletableFuture<>();
        DefaultPlcWriteRequest request = (DefaultPlcWriteRequest) writeRequest;
        List<S7VarRequestParameterItem> parameterItems = new ArrayList<>(request.getNumberOfTags());
        List<S7VarPayloadDataItem> payloadItems = new ArrayList<>(request.getNumberOfTags());

        Iterator<String> iter = request.getTagNames().iterator();

        String tagName = null;
        while (iter.hasNext()) {
            tagName = iter.next();
            final S7Tag tag = (S7Tag) request.getTag(tagName);
            final PlcValue plcValue = request.getPlcValue(tagName);
            parameterItems.add(new S7VarRequestParameterItemAddress(encodeS7Address(tag)));
            payloadItems.add(serializePlcValue(tag, plcValue, iter.hasNext()));
        }


//        for (String tagName : request.getTagNames()) {
//            final S7Tag tag = (S7Tag) request.getTag(tagName);
//            final PlcValue plcValue = request.getPlcValue(tagName);
//            parameterItems.add(new S7VarRequestParameterItemAddress(encodeS7Address(tag)));
//            payloadItems.add(serializePlcValue(tag, plcValue));
//
//        }


        final int tpduId = tpduGenerator.getAndIncrement();
        // If we've reached the max value for a 16 bit transaction identifier, reset back to 1
        tpduGenerator.compareAndExchange(0xFFFF, 1);

        TPKTPacket tpktPacket = new TPKTPacket(
            new COTPPacketData(
                null,
                new S7MessageRequest(tpduId,
                    new S7ParameterWriteVarRequest(parameterItems),
                    new S7PayloadWriteVarRequest(payloadItems)
                ),
                true,
                (byte) tpduId
            )
        );

        // Start a new request-transaction (Is ended in the response-handler)
        RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
        transaction.submit(() -> context.sendRequest(tpktPacket)
            .onTimeout(new TransactionErrorCallback<>(future, transaction))
            .onError(new TransactionErrorCallback<>(future, transaction))
            .expectResponse(TPKTPacket.class, REQUEST_TIMEOUT)
            .check(p -> p.getPayload() instanceof COTPPacketData)
            .unwrap(p -> ((COTPPacketData) p.getPayload()))
            .unwrap(COTPPacket::getPayload)
            .check(p -> p.getTpduReference() == tpduId)
            .handle(p -> {
                try {
                    future.complete(((PlcWriteResponse) decodeWriteResponse(p, writeRequest)));
                } catch (PlcProtocolException e) {
                    logger.warn("Error sending 'write' message: '{}'", e.getMessage(), e);
                }
                // Finish the request-transaction.
                transaction.endRequest();
            }));
        return future;
    }

    @Override
    public CompletableFuture<PlcSubscriptionResponse> subscribe(PlcSubscriptionRequest subscriptionRequest) {
        if (!isConnected()) {
            CompletableFuture<PlcSubscriptionResponse> future = new CompletableFuture<>();
            future.completeExceptionally(new PlcRuntimeException("Disconnected"));
            return future;
        }
        if (!isSubscriptionSupported()) {
            CompletableFuture<PlcSubscriptionResponse> future = new CompletableFuture<>();
            future.completeExceptionally(new PlcRuntimeException("Not Supported"));
            return future;
        }

        CompletableFuture<PlcSubscriptionResponse> future = new CompletableFuture<>();
        HashMap<String, PlcSubscriptionResponse> valuesResponse = new HashMap<>();
        HashMap<String, CompletableFuture<S7Message>> futures = new HashMap<>();

        //Initialize multiple requests.
        CompletableFuture<PlcSubscriptionResponse> response = new CompletableFuture<>();
        subscriptionRequest.getTagNames().forEach(fieldname -> futures.put(fieldname, new CompletableFuture<>()));


        DefaultPlcSubscriptionRequest request = (DefaultPlcSubscriptionRequest) subscriptionRequest;

        List<S7ParameterUserDataItem> parameterItems = new ArrayList<>(request.getNumberOfTags());
        List<S7PayloadUserDataItem> payloadItems = new ArrayList<>(request.getNumberOfTags());

        //The main task that runs the subscriptions.
        CompletableFuture<Void> mainTask;
        mainTask = CompletableFuture.
            allOf(futures.values().toArray(new CompletableFuture[0]));

        Thread t1 = new Thread(() -> {

            for (String tagName : request.getTagNames()) {
                final DefaultPlcSubscriptionTag sf = (DefaultPlcSubscriptionTag) request.getTag(tagName);
                final S7SubscriptionTag tag = (S7SubscriptionTag) sf.getTag();

                switch (tag.getTagType()) {
                    case EVENT_SUBSCRIPTION:
                        encodeEventSubscriptionRequest(request, parameterItems, payloadItems);
                        break;
                    case EVENT_UNSUBSCRIPTION:
                        //encodeEventUnSubscriptionRequest(request, parameterItems, payloadItems);
                        break;
                    case ALARM_ACK:
                        //encodeAlarmAckRequest(request, parameterItems, payloadItems);
                        break;
                    case ALARM_QUERY:
                        encodeAlarmQueryRequest(request, parameterItems, payloadItems);
                        break;
                    case CYCLIC_SUBSCRIPTION:
                        encodeCycledS7ANYSubscriptionRequest(request.getTag(tagName), parameterItems, payloadItems);
                        break;
                    case CYCLIC_DB_SUBSCRIPTION:
                        encodeCycledDBREADSubscriptionRequest(request.getTag(tagName), parameterItems, payloadItems);
                        break;
                    case CYCLIC_UNSUBSCRIPTION:
                        //encodeCycledUnSubscriptionRequest(request, parameterItems, payloadItems);
                        break;
                    default:
                }
                //final PlcValue plcValue = request.getPlcValue(tagName);
                //parameterItems.add(new S7VarRequestParameterItemAddress(encodeS7Address(tag)));
                //payloadItems.add(serializePlcValue(tag, plcValue));
                final int tpduId = tpduGenerator.getAndIncrement();
                // If we've reached the max value for a 16 bit transaction identifier, reset back to 1
                tpduGenerator.compareAndExchange(0xFFFF, 1);

                TPKTPacket tpktPacket = new TPKTPacket(new COTPPacketData(null,
                    new S7MessageUserData(tpduId,
                        new S7ParameterUserData(parameterItems),
                        new S7PayloadUserData(payloadItems)),
                    true, (byte) tpduId));

                // Start a new request-transaction (Is ended in the response-handler)
                RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
                transaction.submit(() -> context.sendRequest(tpktPacket)
                    .onTimeout(new TransactionErrorCallback<>(future, transaction))
                    .onError(new TransactionErrorCallback<>(future, transaction))
                    .expectResponse(TPKTPacket.class, REQUEST_TIMEOUT)
                    .check(p -> p.getPayload() instanceof COTPPacketData)
                    .unwrap(p -> ((COTPPacketData) p.getPayload()))
                    .unwrap(COTPPacket::getPayload)
                    .check(p -> p.getTpduReference() == tpduId)
                    .handle(p -> {
                        try {
                            //future.complete(decodeEventSubscriptionRequest(tagName, p, subscriptionRequest));
                            futures.get(tagName).complete(p);
                        } catch (Exception e) {
                            logger.warn("Error sending 'write' message: '{}'", e.getMessage(), e);
                        }
                        // Finish the request-transaction.
                        transaction.endRequest();
                    }));

                try {
                    valuesResponse.put(tagName, decodeEventSubscriptionRequest(tagName, subscriptionRequest, futures.get(tagName).get()));
                } catch (Exception ex) {
                    logger.warn(ex.toString());
                }

            } //Next

            try {
                mainTask.get();
                HashMap<String, ResponseItem<PlcSubscriptionHandle>> values = new HashMap<>();

                valuesResponse.forEach((s, p) -> {
                    if (p != null)
                        values.putAll(((DefaultPlcSubscriptionResponse) p).getValues());
                });

                response.complete(new DefaultPlcSubscriptionResponse(subscriptionRequest, values));

            } catch (Exception e) {
                logger.warn("uh", e);
            }

        });

        t1.start();

        return response;
    }


    @Override
    public CompletableFuture<PlcUnsubscriptionResponse> unsubscribe(PlcUnsubscriptionRequest unsubscriptionRequest) {
        if (!isConnected()) {
            CompletableFuture<PlcUnsubscriptionResponse> future = new CompletableFuture<>();
            future.completeExceptionally(new PlcRuntimeException("Disconnected"));
            return future;
        }
        if (!isSubscriptionSupported()) {
            CompletableFuture<PlcUnsubscriptionResponse> future = new CompletableFuture<>();
            future.completeExceptionally(new PlcRuntimeException("Not Supported"));
            return future;
        }
        /*CompletableFuture<PlcUnsubscriptionResponse> future =*/ return new CompletableFuture<>();
        // TODO: It seems the unsubscription hasn't been implemented yet.
        //DefaultPlcUnsubscriptionRequest request = (DefaultPlcUnsubscriptionRequest) unsubscriptionRequest;
    }

    private void encodeEventSubscriptionRequest(DefaultPlcSubscriptionRequest request,
                                                List<S7ParameterUserDataItem> parameterItems,
                                                List<S7PayloadUserDataItem> payloadItems) {
        byte subsevent = 0x00;
        for (String tagName : request.getTagNames()) {
            if (request.getTag(tagName) instanceof DefaultPlcSubscriptionTag) {
                PlcTag event = ((DefaultPlcSubscriptionTag) request.getTag(tagName)).getTag();
                if (event instanceof S7SubscriptionTag) {
                    if (((S7SubscriptionTag) event).getTagType() == S7SubscriptionType.EVENT_SUBSCRIPTION)
                        subsevent = (byte) (subsevent | ((S7SubscriptionTag) event).getEventType().getValue());
                }
            }
        }

        S7ParameterUserDataItemCPUFunctions parameter = new S7ParameterUserDataItemCPUFunctions(
            (short) 0x11,   //Method
            (byte) 0x04,    //FunctionType
            (byte) 0x04,    //FunctionGroup
            (short) 0x02,   //SubFunction
            (short) 0x00,   //SequenceNumber
            null,   //DataUnitReferenceNumber
            null,   //LastDataUnit
            null         //errorCode
        );
        parameterItems.clear();
        parameterItems.add(parameter);


        S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest payload;
        if (subsevent > 0) {
            payload = new S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest(
                DataTransportErrorCode.OK,
                DataTransportSize.OCTET_STRING,
                0x0a,
                subsevent,
                "HmiRtm  ",
                null,
                null);
        } else {
            //TODO: Check for ALARM_S (S7300) and ALARM_8 (S7400), maybe we need verify the CPU
            AlarmStateType alarmtype;
            if (s7DriverContext.getControllerType() == S7ControllerType.S7_400) {
                alarmtype = AlarmStateType.ALARM_INITIATE;
            } else {
                alarmtype = AlarmStateType.ALARM_S_INITIATE;
            }

            short auxSubsEvent = (short) (subsevent & 0xFF);

            payload = new S7PayloadUserDataItemCpuFunctionMsgSubscriptionRequest(
                DataTransportErrorCode.OK,
                DataTransportSize.OCTET_STRING,
                (short) 0x0c,
                auxSubsEvent,
                "HmiRtm  ",
                alarmtype,
                (short) 0x00);
        }

        payloadItems.clear();
        payloadItems.add(payload);

    }

    private PlcSubscriptionResponse decodeEventSubscriptionRequest(String strTagName,
                                                                   PlcSubscriptionRequest plcSubscriptionRequest,
                                                                   S7Message responseMessage)
        throws PlcProtocolException {

        Map<String, ResponseItem<PlcSubscriptionHandle>> values = new HashMap<>();
        short errorClass = 0;
        short errorCode = 0;
        // TODO: It seems this hasn't been implemented yet.
        /*if (responseMessage instanceof S7MessageUserData) {
            S7MessageUserData messageUserData = (S7MessageUserData) responseMessage;
            S7PayloadUserData payload = (S7PayloadUserData) messageUserData.getPayload();
            // errorClass = payload.getItems()[0].
            // errorCode = messageUserData.getParameter().

        } else*/ if (responseMessage instanceof S7MessageResponse) {
            S7MessageResponse messageResponse = (S7MessageResponse) responseMessage;
            errorClass = messageResponse.getErrorClass();
            errorCode = messageResponse.getErrorCode();
        } else {
            throw new PlcProtocolException("Unsupported message type " + responseMessage.getClass().getName());
        }

        // If the result contains any form of non-null error code, handle this instead.
        if ((errorClass != 0) || (errorCode != 0)) {
            // This is usually the case if PUT/GET wasn't enabled on the PLC
            if ((errorClass == 129) && (errorCode == 4)) {
                logger.warn("Got an error response from the PLC. This particular response code usually indicates " +
                    "that PUT/GET is not enabled on the PLC.");
                for (String tagName : plcSubscriptionRequest.getTagNames()) {
                    values.put(tagName, null);
                }
                return new DefaultPlcSubscriptionResponse(plcSubscriptionRequest, values);
            } else {
                logger.warn("Got an unknown error response from the PLC. Error Class: {}, Error Code {}. " +
                        "We probably need to implement explicit handling for this, so please file a bug-report " +
                        "on https://issues.apache.org/jira/projects/PLC4X and ideally attach a WireShark dump " +
                        "containing a capture of the communication.",
                    errorClass, errorCode);
                for (String tagName : plcSubscriptionRequest.getTagNames()) {
                    values.put(tagName, null);
                }
                return new DefaultPlcSubscriptionResponse(plcSubscriptionRequest, values);
            }
        }

        S7ParameterUserData parameter = (S7ParameterUserData) responseMessage.getParameter();
        List<S7ParameterUserDataItem> parameters = parameter.getItems();
        S7ParameterUserDataItemCPUFunctions itemparameter = (S7ParameterUserDataItemCPUFunctions) parameters.get(0);
        errorCode = itemparameter.getErrorCode().shortValue();

        // In all other cases all went well.
        S7PayloadUserData payload = (S7PayloadUserData) responseMessage.getPayload();

        List<S7PayloadUserDataItem> payloadItems = payload.getItems();

        //Only one item for any number of subscription (4)
        if (payloadItems.size() == 0) {
            throw new PlcProtocolException(
                "The number of requested items doesn't match the number of returned items");
        }

        boolean responseOk = false;
        if (payloadItems.get(0) instanceof S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) {
            S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse item =
                (S7PayloadUserDataItemCpuFunctionMsgSubscriptionResponse) payloadItems.get(0);
            if ((item.getReturnCode() == DataTransportErrorCode.OK) &&
                (item.getTransportSize() == DataTransportSize.OCTET_STRING)) {
                responseOk = true;
            }
        } else if (payloadItems.get(0) instanceof S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) {
            S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse item =
                (S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) payloadItems.get(0);
            if ((item.getReturnCode() == DataTransportErrorCode.OK) &&
                (item.getTransportSize() == DataTransportSize.OCTET_STRING)) {
                responseOk = true;
            }
        } else if (payloadItems.get(0) instanceof S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) {
            S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse item =
                (S7PayloadUserDataItemCpuFunctionMsgSubscriptionAlarmResponse) payloadItems.get(0);
            if ((item.getReturnCode() == DataTransportErrorCode.OK) &&
                (item.getTransportSize() == DataTransportSize.OCTET_STRING)) {
                responseOk = true;
            }
        } else if (payloadItems.get(0) instanceof S7PayloadUserDataItemCpuFunctionAlarmAckResponse) {
            S7PayloadUserDataItemCpuFunctionAlarmAckResponse items =
                (S7PayloadUserDataItemCpuFunctionAlarmAckResponse)
                    payloadItems.get(0);
            //String tagName = (String) plcSubscriptionRequest.getTagNames().toArray()[0];
            //TODO: Chequear si tagName es el correcto           
            logger.info("strTagName: {}", strTagName);
            values.put(strTagName, new ResponseItem<>(PlcResponseCode.OK, null));
            for (short s : items.getMessageObjects()) {

                if (s == 0x0000) {
                    values.put(Integer.toHexString(s), new ResponseItem<>(PlcResponseCode.OK, null));
                } else if (s == 0x000a) {
                    values.put(Integer.toHexString(s), new ResponseItem<>(PlcResponseCode.NOT_FOUND, null));
                }
            }

            return new DefaultPlcSubscriptionResponse(plcSubscriptionRequest, values);

        } else if (payloadItems.get(0) instanceof S7PayloadUserDataItemCpuFunctionAlarmAckErrorResponse) {

            /*S7PayloadUserDataItemCpuFunctionAlarmAckResponse items =
                (S7PayloadUserDataItemCpuFunctionAlarmAckResponse)
                    payloadItems.get(0);*/
            //String fieldName = (String) S7PayloadUserDataItemCyclicServicesPush .getFieldNames().toArray()[0];
            logger.warn("Request field: {}: {} {}", strTagName, S7ParamErrorCode.valueOf(errorCode), S7ParamErrorCode.valueOf(errorCode).getEvent());
            values.put(strTagName, new ResponseItem<>(PlcResponseCode.NOT_FOUND, null));
            return new DefaultPlcSubscriptionResponse(plcSubscriptionRequest, values);

        } else if (payloadItems.get(0) instanceof S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) {

            S7PayloadUserDataItemCpuFunctionAlarmQueryResponse items =
                (S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) payloadItems.get(0);

            ByteBuf buffer = Unpooled.directBuffer(items.getItems().length * 2);
            ByteBuf rxBuffer = Unpooled.directBuffer(items.getItems().length * 2);
            buffer.writeBytes(items.getItems());


            //int numberOfItems = 1;

            if (itemparameter.getLastDataUnit() == 1) {

                short loop = 0xff;
                CompletableFuture<S7MessageUserData> loopfuture;
                S7MessageUserData msg;
                S7ParameterUserDataItemCPUFunctions loopparameter;
                S7PayloadUserDataItemCpuFunctionAlarmQueryResponse looppayload = null;

                do {
                    loopfuture = reassembledAlarmEvents(itemparameter.getSequenceNumber());

                    try {

                        msg = loopfuture.get();
                        if (msg != null) {
                            loopparameter = (S7ParameterUserDataItemCPUFunctions) ((S7ParameterUserData) msg.getParameter()).getItems().get(0);
                            looppayload = (S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) ((S7PayloadUserData) msg.getPayload()).getItems().get(0);

                            buffer.writeBytes(looppayload.getItems());

                            loop = loopparameter.getLastDataUnit();

                        } else loop = 0x00;

                    } catch (Exception ex) {
                        logger.warn(ex.toString());
                    }
                } while (loop > 0x00);

                rxBuffer.writeByte(looppayload.getReturnCode().getValue());
                rxBuffer.writeByte(looppayload.getTransportSize().getValue());
                rxBuffer.writeShort(looppayload.getDataLength());
                rxBuffer.writeBytes(buffer);

            } else {
                rxBuffer.writeByte(payloadItems.get(0).getReturnCode().getValue());
                rxBuffer.writeByte(payloadItems.get(0).getTransportSize().getValue());
                rxBuffer.writeShort(payloadItems.get(0).getDataLength());
                rxBuffer.writeBytes(buffer);
            }

            ReadBuffer readBuffer = new ReadBufferByteBased(ByteBufUtil.getBytes(rxBuffer));

            try {

                short cpuSubfunction = 0x00;

                if (s7DriverContext.getControllerType() == S7ControllerType.S7_300) {
                    cpuSubfunction = 0x13;
                } else cpuSubfunction = 0xf0;

                S7PayloadUserDataItem payloadItem =
                    S7PayloadUserDataItem.staticParse(readBuffer,
                        (byte) 0x04,
                        (byte) 0x00,
                        cpuSubfunction);

                eventQueue.add(payloadItem);
            } catch (Exception ex) {
                logger.info(ex.toString());
            }

            PlcResponseCode resCode = (items.getReturnCode() == DataTransportErrorCode.OK) ? PlcResponseCode.OK : PlcResponseCode.INTERNAL_ERROR;
            values.put(strTagName, new ResponseItem<>(resCode, null));
            return new DefaultPlcSubscriptionResponse(plcSubscriptionRequest, values);

        } else if (payloadItems.get(0) instanceof S7PayloadUserDataItemCyclicServicesSubscribeResponse) {
            //S7ParameterUserData parameter = (S7ParameterUserData) responseMessage.getParameter();  
            //logger.info("Aqui debe responder a Cyclic transfer");
            S7ParameterUserDataItemCPUFunctions msgParameter = (S7ParameterUserDataItemCPUFunctions)
                parameter.getItems().get(0);

            S7CyclicEvent cycevent = new S7CyclicEvent(plcSubscriptionRequest,
                msgParameter.getSequenceNumber(),
                (S7PayloadUserDataItemCyclicServicesSubscribeResponse) payloadItems.get(0));

            eventQueue.add(cycevent);
            values.put(strTagName, new ResponseItem<>(PlcResponseCode.OK, cycHandle));
            return new DefaultPlcSubscriptionResponse(plcSubscriptionRequest, values);

        } else if (payloadItems.get(0) instanceof S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) {
            //logger.info("Aqui debe responder a Cyclic transfer Change Driven");
            S7ParameterUserDataItemCPUFunctions msgParameter = (S7ParameterUserDataItemCPUFunctions)
                parameter.getItems().get(0);

            S7CyclicEvent cycEvent = new S7CyclicEvent(plcSubscriptionRequest,
                msgParameter.getSequenceNumber(),
                (S7PayloadUserDataItemCyclicServicesChangeDrivenSubscribeResponse) payloadItems.get(0));

            eventQueue.add(cycEvent);
            values.put(strTagName, new ResponseItem<>(PlcResponseCode.OK, cycHandle));
            return new DefaultPlcSubscriptionResponse(plcSubscriptionRequest, values);

        } else if (payloadItems.get(0) instanceof S7PayloadUserDataItemCyclicServicesErrorResponse) {
            // TODO: It seems this isn't fully implemented.
            //S7ParameterUserData parameter = (S7ParameterUserData) responseMessage.getParameter();
            //S7ParameterUserDataItem[] parameters = parameter.getItems();
            //S7ParameterUserDataItemCPUFunctions itemparameter = (S7ParameterUserDataItemCPUFunctions) parameters[0];
            //errorCode = itemparameter.getErrorCode().shortValue();
            logger.warn("Request field: " + strTagName + ": " + S7ParamErrorCode.valueOf(errorCode) + " " + S7ParamErrorCode.valueOf(errorCode).getEvent());
            if (errorCode == 0x8104) {
                values.put(strTagName, new ResponseItem<>(PlcResponseCode.UNSUPPORTED, null));
            } else
                values.put(strTagName, new ResponseItem<>(PlcResponseCode.INTERNAL_ERROR, null));
            return new DefaultPlcSubscriptionResponse(plcSubscriptionRequest, values);
        } else if (payloadItems.get(0) instanceof S7PayloadUserDataItemCyclicServicesUnsubscribeResponse) {
            values.put(strTagName, new ResponseItem<>(PlcResponseCode.OK, null));
            return new DefaultPlcSubscriptionResponse(plcSubscriptionRequest, values);
        }

        if (responseOk) {
            for (String tagName : plcSubscriptionRequest.getTagNames()) {
                DefaultPlcSubscriptionTag dTag = (DefaultPlcSubscriptionTag) plcSubscriptionRequest.getTag(tagName);
                S7SubscriptionTag tag = (S7SubscriptionTag) dTag.getTag();
                switch (tag.getEventType()) {
                    case MODE:
                        values.put(tagName, new ResponseItem<>(PlcResponseCode.OK, modeHandle));
                        break;
                    case SYS:
                        values.put(tagName, new ResponseItem<>(PlcResponseCode.OK, sysHandle));
                        break;
                    case USR:
                        values.put(tagName, new ResponseItem<>(PlcResponseCode.OK, usrHandle));
                        break;
                    case ALM:
                        values.put(tagName, new ResponseItem<>(PlcResponseCode.OK, almHandle));
                        break;
                }

            }
            return new DefaultPlcSubscriptionResponse(plcSubscriptionRequest, values);
        }

        return null;
    }

    private void encodeEventUnSubscriptionRequest(DefaultPlcSubscriptionRequest request,
                                                  List<S7VarRequestParameterItem> parameterItems,
                                                  List<S7VarPayloadDataItem> payloadItems) {

    }

    private void encodeAlarmAckRequest(DefaultPlcReadRequest request,
                                       List<S7ParameterUserDataItem> parameterItems,
                                       List<S7PayloadUserDataItem> payloadItems) {

        S7ParameterUserDataItemCPUFunctions parameter = new S7ParameterUserDataItemCPUFunctions(
            (short) 0x11,   //Method
            (byte) 0x04,    //FunctionType
            (byte) 0x04,    //FunctionGroup
            (short) 0x0b,   //SubFunction
            (short) 0x00,   //SequenceNumber
            null,   //DataUnitReferenceNumber
            null,   //LastDataUnit
            null         //errorCode
        );
        parameterItems.clear();
        parameterItems.add(parameter);

        ArrayList<AlarmMessageObjectAckType> messageObjects = null;
        BitSet bs;
        for (String fieldName : request.getTagNames()) {
            if (request.getTag(fieldName) instanceof S7AckTag) {
                PlcTag field = request.getTag(fieldName);
                if (field instanceof S7AckTag) {
                    ArrayList<Integer> arrAlarmIds = ((S7AckTag) field).getAlarmIds();
                    ArrayList<Integer> arrAlarmSigs = ((S7AckTag) field).getAlarmSigs();
                    //messageObjects = new AlarmMessageObjectAckType[arrAlarmIds.size()];
                    messageObjects = new ArrayList<>();
                    int j = 0;
                    for (int i = 0; i < arrAlarmIds.size(); i++) {
                        bs = BitSet.valueOf(new byte[]{arrAlarmSigs.get(i).byteValue()});

                        AlarmMessageObjectAckType messageObject = new AlarmMessageObjectAckType(
                            SyntaxIdType.ALARM_ACKSET,
                            (short) 0,
                            arrAlarmIds.get(i),
                            new State(bs.get(7), bs.get(6), bs.get(5), bs.get(4), bs.get(3), bs.get(2), bs.get(1), bs.get(0)),
                            new State(bs.get(7), bs.get(6), bs.get(5), bs.get(4), bs.get(3), bs.get(2), bs.get(1), bs.get(0))
                        );
                        //messageObjects[i] = messageObject;
                        messageObjects.add(messageObject);
                    }
                }
            }
        }


        S7PayloadUserDataItemCpuFunctionAlarmAckRequest payload =
            new S7PayloadUserDataItemCpuFunctionAlarmAckRequest(
                DataTransportErrorCode.OK,
                DataTransportSize.OCTET_STRING,
                0x0c,
                messageObjects);
//
        payloadItems.clear();
        payloadItems.add(payload);

    }

    private void encodeAlarmQueryRequest(DefaultPlcSubscriptionRequest request,
                                         List<S7ParameterUserDataItem> parameterItems,
                                         List<S7PayloadUserDataItem> payloadItems) {

        S7ParameterUserDataItemCPUFunctions parameter = new S7ParameterUserDataItemCPUFunctions(
            (short) 0x11,   //Method
            (byte) 0x04,    //FunctionType
            (byte) 0x04,    //FunctionGroup
            (short) 0x13,   //SubFunction
            (short) 0x00,   //SequenceNumber
            null,   //DataUnitReferenceNumber
            null,   //LastDataUnit
            null         //errorCode
        );

        parameterItems.clear();
        parameterItems.add(parameter);

        //TODO: Chequear el tipo dfe larma.
        S7PayloadUserDataItemCpuFunctionAlarmQueryRequest payload =
            new S7PayloadUserDataItemCpuFunctionAlarmQueryRequest(
                DataTransportErrorCode.OK,
                DataTransportSize.OCTET_STRING,
                0x0c,
                SyntaxIdType.ALARM_QUERYREQSET,
                QueryType.ALARM_8P,
                AlarmType.ALARM_8);

        payloadItems.clear();
        payloadItems.add(payload);
    }

    private void encodeCycledSubscriptionRequest(DefaultPlcSubscriptionRequest request,
                                                 List<S7ParameterUserDataItem> parameterItems,
                                                 List<S7PayloadUserDataItem> payloadItem) {

    }

    private void encodeCycledS7ANYSubscriptionRequest(PlcSubscriptionTag plctag,
                                                      List<S7ParameterUserDataItem> parameterItems,
                                                      List<S7PayloadUserDataItem> payloadItems) {
        S7ParameterUserDataItemCPUFunctions parameter = new S7ParameterUserDataItemCPUFunctions(
            (short) 0x11,   //Method
            (byte) 0x04,    //FunctionType
            (byte) 0x02,    //FunctionGroup
            (short) 0x01,   //SubFunction
            (short) 0x00,   //SequenceNumber
            null,   //DataUnitReferenceNumber
            null,   //LastDataUnit
            null         //errorCode
        );

        parameterItems.clear();
        parameterItems.add(parameter);

        //TODO: Chequear la asignacionde tipo

        PlcTag tag = ((DefaultPlcSubscriptionTag) plctag).getTag();
        S7SubscriptionTag s7tag = (S7SubscriptionTag) tag;
//
        List<CycServiceItemType> items = new ArrayList<>();
//        
        int i = 0;
        for (S7Tag userfield : s7tag.getS7Tags()) {
            items.add(new CycServiceItemAnyType(
                (short) 0x0a,
                (short) 0x10, //S7ANY
                userfield.getDataType(),
                userfield.getNumberOfElements(),
                userfield.getBlockNumber(),
                userfield.getMemoryArea(),
                (((userfield.getByteOffset() << 3) | (userfield.getBitOffset() & 0x0007)))
            ));
            i++;
        }
//        
//        //Length in bytes
//        //4+12*items.legth
        int lengthInBytes = 4 + items.size() * 12;
//
        S7PayloadUserDataItemCyclicServicesSubscribeRequest payload =
            new S7PayloadUserDataItemCyclicServicesSubscribeRequest(
                DataTransportErrorCode.OK,
                DataTransportSize.OCTET_STRING,
                lengthInBytes,
                items.size(),
                s7tag.getTimeBase(),
                s7tag.getMultiplier(),
                items
            );

        payloadItems.clear();
        payloadItems.add(payload);
    }


    private void encodeCycledDBREADSubscriptionRequest(PlcSubscriptionTag plctag,
                                                       List<S7ParameterUserDataItem> parameterItems,
                                                       List<S7PayloadUserDataItem> payloadItems) {
        S7ParameterUserDataItemCPUFunctions parameter = new S7ParameterUserDataItemCPUFunctions(
            (short) 0x11,   //Method
            (byte) 0x04,    //FunctionType
            (byte) 0x02,    //FunctionGroup
            (short) 0x01,   //SubFunction
            (short) 0x00,   //SequenceNumber
            null,   //DataUnitReferenceNumber
            null,   //LastDataUnit
            null         //errorCode
        );

        parameterItems.clear();
        parameterItems.add(parameter);

        //TODO: Chequear la asignacionde tipo

        PlcTag tag = ((DefaultPlcSubscriptionTag) plctag).getTag();
        S7SubscriptionTag s7tag = (S7SubscriptionTag) tag;


        double cycitemcount = Math.ceil(((double) s7tag.getS7Tags().length / (double) 50));

        //CycServiceItemDbReadType[] cycitems = new CycServiceItemDbReadType[(int) cycitemcount];
        List<CycServiceItemType> cycitems = new ArrayList();
        ArrayList<SubItem> subitems = new ArrayList();
        for (S7Tag userfield : s7tag.getS7Tags()) {
            subitems.add(new SubItem((short) userfield.getNumberOfElements(),
                userfield.getBlockNumber(),
                userfield.getByteOffset()));
        }

        int initpos = 0;
        int endpos = (subitems.size() < 50) ? subitems.size() : 49;
        int j = 0;
        int lengthInBytes = 4;
        do {
            List<SubItem> arraysubitems = subitems.subList(initpos, endpos);

            cycitems.add(j, new CycServiceItemDbReadType(
                (short) (arraysubitems.size() * 5 + 2),
                (short) 0xb0,
                (short) arraysubitems.size(),
                arraysubitems));

            lengthInBytes += 4 + arraysubitems.size() * 5;
            initpos = endpos++;
            endpos = ((initpos + 49) < subitems.size()) ? (initpos + 49) : subitems.size();
            j++;
        } while (j < cycitems.size());

        S7PayloadUserDataItemCyclicServicesSubscribeRequest payload =
            new S7PayloadUserDataItemCyclicServicesSubscribeRequest(
                DataTransportErrorCode.OK,
                DataTransportSize.OCTET_STRING,
                lengthInBytes,
                cycitems.size(),
                s7tag.getTimeBase(),
                s7tag.getMultiplier(),
                cycitems
            );

        payloadItems.clear();
        payloadItems.add(payload);
    }


    private void encodeCycledUnSubscriptionRequest(DefaultPlcSubscriptionRequest request,
                                                   List<S7ParameterUserDataItem> parameterItems,
                                                   List<S7PayloadUserDataItem> payloadItems) {

        S7ParameterUserDataItemCPUFunctions parameter = new S7ParameterUserDataItemCPUFunctions(
            (short) 0x11,   //Method
            (byte) 0x04,    //FunctionType
            (byte) 0x02,    //FunctionGroup
            (short) 0x04,   //SubFunction
            (short) 0x00,   //SequenceNumber
            null,   //DataUnitReferenceNumber
            null,   //LastDataUnit
            null         //errorCode
        );

        parameterItems.clear();
        parameterItems.add(parameter);

        PlcSubscriptionTag tag = request.getTags().get(0);
        S7SubscriptionTag s7tag = (S7SubscriptionTag) tag;

        //PlcField field = ((DefaultPlcSubscriptionField) plcfield).getPlcField(); 
        //S7SubscriptionField s7field = (S7SubscriptionField) field;

        //TODO:Check CPU type
        S7PayloadUserDataItemCyclicServicesUnsubscribeRequest payload =
            new S7PayloadUserDataItemCyclicServicesUnsubscribeRequest(
                DataTransportErrorCode.OK,
                DataTransportSize.OCTET_STRING,
                0x02,
                (short) 0x05,
                s7tag.getAckAlarms().get(0).byteValue()
            );

        payloadItems.clear();
        payloadItems.add(payload);

    }

    /**
     * This method is only called when there is no Response Handler.
     */
    @Override
    protected void decode(ConversationContext<TPKTPacket> context, TPKTPacket msg) throws Exception {
        S7Message s7msg = msg.getPayload().getPayload();
        S7Parameter parameter = s7msg.getParameter();
        if (parameter instanceof S7ParameterModeTransition) {
            eventQueue.add(parameter);
        } else if (parameter instanceof S7ParameterUserData) {
            S7ParameterUserData parameterud = (S7ParameterUserData) parameter;
            List<S7ParameterUserDataItem> parameterudis = parameterud.getItems();
            for (S7ParameterUserDataItem parameterudi : parameterudis) {
                if (parameterudi instanceof S7ParameterUserDataItemCPUFunctions) {
                    S7ParameterUserDataItemCPUFunctions myparameter = (S7ParameterUserDataItemCPUFunctions) parameterudi;
                    //TODO: Check from mspec. We can try using "instanceof"
                    if ((myparameter.getCpuFunctionType() == 0x00) && (myparameter.getCpuSubfunction() == 0x03)) {
                        S7PayloadUserData payload = (S7PayloadUserData) s7msg.getPayload();
                        List<S7PayloadUserDataItem> items = payload.getItems();
                        for (S7PayloadUserDataItem item : items) {
                            if (item instanceof S7PayloadDiagnosticMessage) {
                                eventQueue.add(item);
                            }
                        }
                    } else if ((myparameter.getCpuFunctionType() == 0x00) &&
                        ((myparameter.getCpuSubfunction() == 0x05) ||
                            (myparameter.getCpuSubfunction() == 0x06) ||
                            (myparameter.getCpuSubfunction() == 0x0c) ||
                            (myparameter.getCpuSubfunction() == 0x11) ||
                            (myparameter.getCpuSubfunction() == 0x12) ||
                            (myparameter.getCpuSubfunction() == 0x13) ||
                            (myparameter.getCpuSubfunction() == 0x16))) {
                        S7PayloadUserData payload = (S7PayloadUserData) s7msg.getPayload();
                        List<S7PayloadUserDataItem> items = payload.getItems();
                        eventQueue.addAll(items);
                    } else if ((myparameter.getCpuFunctionType() == 0x00) && (myparameter.getCpuSubfunction() == 0x13)) {

                    } else if ((myparameter.getCpuFunctionGroup() == 0x02) && (myparameter.getCpuFunctionType() == 0x00) && (myparameter.getCpuSubfunction() == 0x01)) {

                        S7ParameterUserDataItemCPUFunctions parameteritem =
                            (S7ParameterUserDataItemCPUFunctions)
                                ((S7ParameterUserData) parameter).getItems().get(0);

                        S7PayloadUserData payload = (S7PayloadUserData) s7msg.getPayload();

                        S7PayloadUserDataItemCyclicServicesPush payloaditem =
                            (S7PayloadUserDataItemCyclicServicesPush)
                                payload.getItems().get(0);

                        S7CyclicEvent cycevent = new S7CyclicEvent(null,
                            parameteritem.getSequenceNumber(),
                            payloaditem);
                        eventQueue.add(cycevent);

                    } else if ((myparameter.getCpuFunctionGroup() == 0x02) && (myparameter.getCpuFunctionType() == 0x00) && (myparameter.getCpuSubfunction() == 0x05)) {
                        S7ParameterUserDataItemCPUFunctions parameterItem =
                            (S7ParameterUserDataItemCPUFunctions)
                                ((S7ParameterUserData) parameter).getItems().get(0);

                        S7PayloadUserData payload = (S7PayloadUserData) s7msg.getPayload();

                        S7PayloadUserDataItemCyclicServicesChangeDrivenPush payloadItem =
                            (S7PayloadUserDataItemCyclicServicesChangeDrivenPush)
                                payload.getItems().get(0);

                        S7CyclicEvent cycEvent = new S7CyclicEvent(null,
                            parameterItem.getSequenceNumber(),
                            payloadItem);
                        eventQueue.add(cycEvent);

                    } else if ((myparameter.getCpuFunctionType() == 0x08) && (myparameter.getCpuSubfunction() == 0x01)) {

                    } else if ((myparameter.getCpuFunctionType() == 0x08) && (myparameter.getCpuSubfunction() == 0x04)) {

                    }
                }
            }
        }
    }

    @Override
    public void close(ConversationContext<TPKTPacket> context) {
        // TODO Implement Closing on Protocol Level
        EventLogic.stop();
    }

    private void extractControllerTypeAndFireConnected(ConversationContext<TPKTPacket> context, S7PayloadUserData payloadUserData) {
        for (S7PayloadUserDataItem item : payloadUserData.getItems()) {
            if (!(item instanceof S7PayloadUserDataItemCpuFunctionReadSzlResponse)) {
                continue;
            }
            S7PayloadUserDataItemCpuFunctionReadSzlResponse readSzlResponseItem =
                (S7PayloadUserDataItemCpuFunctionReadSzlResponse) item;

            // TODO: @carcia: why did you do that this way? it looks like you skip the two bytes you removed with the mspec changes
//            for (SzlDataTreeItem readSzlResponseItemItem : readSzlResponseItem.getItems()) {
//                if (readSzlResponseItemItem.getItemIndex() != 0x0001) {
//                    continue;
//                }
//                final String articleNumber = new String(readSzlResponseItemItem.getMlfb());
//                s7DriverContext.setControllerType(decodeControllerType(articleNumber));
//
//                // Send an event that connection setup is complete.
//                context.fireConnected();
//            }
            ByteBuf szlitem = Unpooled.wrappedBuffer(readSzlResponseItem.getItems());
            String articleNumber = szlitem.toString(2, 20, Charset.defaultCharset());
            s7DriverContext.setControllerType(decodeControllerType(articleNumber));
            context.fireConnected();
        }
    }

    private TPKTPacket createIdentifyRemoteMessage() {
        S7MessageUserData identifyRemoteMessage = new S7MessageUserData(1, new S7ParameterUserData(Collections.singletonList(
            new S7ParameterUserDataItemCPUFunctions((short) 0x11, (byte) 0x4, (byte) 0x4, (short) 0x01, (short) 0x00, null, null, null)
        )), new S7PayloadUserData(Collections.singletonList(
            new S7PayloadUserDataItemCpuFunctionReadSzlRequest(DataTransportErrorCode.OK,
                DataTransportSize.OCTET_STRING,
                0x0C,
                new SzlId(SzlModuleTypeClass.CPU,
                    (byte) 0x00,
                    SzlSublist.MODULE_IDENTIFICATION),
                0x0000)
        )));
        COTPPacketData cotpPacketData = new COTPPacketData(null, identifyRemoteMessage, true, (byte) 2);
        return new TPKTPacket(cotpPacketData);
    }

    private TPKTPacket createS7ConnectionRequest(COTPPacketConnectionResponse cotpPacketConnectionResponse) {
        for (COTPParameter parameter : cotpPacketConnectionResponse.getParameters()) {
            if (parameter instanceof COTPParameterCalledTsap) {
                COTPParameterCalledTsap cotpParameterCalledTsap = (COTPParameterCalledTsap) parameter;
                s7DriverContext.setCalledTsapId(cotpParameterCalledTsap.getTsapId());
            } else if (parameter instanceof COTPParameterCallingTsap) {
                COTPParameterCallingTsap cotpParameterCallingTsap = (COTPParameterCallingTsap) parameter;
                if (cotpParameterCallingTsap.getTsapId() != s7DriverContext.getCallingTsapId()) {
                    s7DriverContext.setCallingTsapId(cotpParameterCallingTsap.getTsapId());
                    logger.warn("Switching calling TSAP id to '{}'", s7DriverContext.getCallingTsapId());
                }
            } else if (parameter instanceof COTPParameterTpduSize) {
                COTPParameterTpduSize cotpParameterTpduSize = (COTPParameterTpduSize) parameter;
                s7DriverContext.setCotpTpduSize(cotpParameterTpduSize.getTpduSize());
            } else {
                logger.warn("Got unknown parameter type '{}'", parameter.getClass().getName());
            }
        }

        // Send an S7 login message.
        S7ParameterSetupCommunication s7ParameterSetupCommunication =
            new S7ParameterSetupCommunication(
                s7DriverContext.getMaxAmqCaller(), s7DriverContext.getMaxAmqCallee(), s7DriverContext.getPduSize());
        S7Message s7Message = new S7MessageRequest(0, s7ParameterSetupCommunication,
            null);
        COTPPacketData cotpPacketData = new COTPPacketData(null, s7Message, true, (byte) 1);
        return new TPKTPacket(cotpPacketData);
    }

    private COTPPacketConnectionRequest createCOTPConnectionRequest(int calledTsapId, int callingTsapId, COTPTpduSize cotpTpduSize) {
        return new COTPPacketConnectionRequest(
            Arrays.asList(
                new COTPParameterCallingTsap(callingTsapId),
                new COTPParameterCalledTsap(calledTsapId),
                new COTPParameterTpduSize(cotpTpduSize)
            ), null, (short) 0x0000, (short) 0x000F, COTPProtocolClass.CLASS_0);
    }

    private PlcResponse decodeReadResponse(S7Message responseMessage, PlcReadRequest plcReadRequest) throws PlcProtocolException {
        Map<String, ResponseItem<PlcValue>> values = new HashMap<>();
        short errorClass;
        short errorCode;

        S7ParameterUserDataItemCPUFunctions parameteritem = null;
        if (responseMessage instanceof S7MessageResponseData) {
            S7MessageResponseData messageResponseData = (S7MessageResponseData) responseMessage;
            errorClass = messageResponseData.getErrorClass();
            errorCode = messageResponseData.getErrorCode();
        } else if (responseMessage instanceof S7MessageResponse) {
            S7MessageResponse messageResponse = (S7MessageResponse) responseMessage;
            errorClass = messageResponse.getErrorClass();
            errorCode = messageResponse.getErrorCode();
        } else if (responseMessage instanceof S7MessageUserData) {
            S7MessageUserData messageResponse = (S7MessageUserData) responseMessage;
            S7ParameterUserData parameters = (S7ParameterUserData) messageResponse.getParameter();
            parameteritem = (S7ParameterUserDataItemCPUFunctions) parameters.getItems().get(0);
            errorClass = 0;
            errorCode = parameteritem.getErrorCode().shortValue();
        } else {
            throw new PlcProtocolException("Unsupported message type " + responseMessage.getClass().getName());
        }
        // If the result contains any form of non-null error code, handle this instead.
        if ((errorClass != 0) || (errorCode != 0)) {
            // This is usually the case if PUT/GET wasn't enabled on the PLC
            if ((errorClass == 129) && (errorCode == 4)) {
                logger.warn("Got an error response from the PLC. This particular response code usually indicates " +
                    "that PUT/GET is not enabled on the PLC.");
                for (String tagName : plcReadRequest.getTagNames()) {
                    ResponseItem<PlcValue> result = new ResponseItem<>(PlcResponseCode.ACCESS_DENIED, new PlcNull());
                    values.put(tagName, result);
                }
                return new DefaultPlcReadResponse(plcReadRequest, values);
            } else {
                logger.warn("Got an unknown error response from the PLC. Error Class: {}, Error Code {}. " +
                        "We probably need to implement explicit handling for this, so please file a bug-report " +
                        "on https://issues.apache.org/jira/projects/PLC4X and ideally attach a WireShark dump " +
                        "containing a capture of the communication.",
                    errorClass, errorCode);
                for (String tagName : plcReadRequest.getTagNames()) {
                    ResponseItem<PlcValue> result = new ResponseItem<>(PlcResponseCode.INTERNAL_ERROR, new PlcNull());
                    values.put(tagName, result);
                }
                return new DefaultPlcReadResponse(plcReadRequest, values);
            }
        }

        //TODO: Reemsambling message.
        if (responseMessage instanceof S7MessageUserData) {

            S7PayloadUserData payload = (S7PayloadUserData) responseMessage.getPayload();
            if (plcReadRequest.getNumberOfTags() != payload.getItems().size()) {
                throw new PlcProtocolException(
                    "The number of requested items doesn't match the number of returned items");
            }

            List<S7PayloadUserDataItem> payloadItems = payload.getItems();

            PlcResponseCode responseCode = PlcResponseCode.INTERNAL_ERROR;
            PlcValue plcValue = null;
            int index = 0;
            for (String tagName : plcReadRequest.getTagNames()) {

                if (plcReadRequest.getTag(tagName) instanceof S7SzlTag) {
                    S7PayloadUserDataItemCpuFunctionReadSzlResponse payloadItem = (S7PayloadUserDataItemCpuFunctionReadSzlResponse) payloadItems.get(index);
                    responseCode = decodeResponseCode(payloadItem.getReturnCode());

                    if (responseCode == PlcResponseCode.OK) {
                        try {
                            List<PlcValue> plcvalues = new LinkedList<>();
                            byte[] data = payloadItem.getItems();

                            for (byte b : data) plcvalues.add(new PlcSINT(b));

                            if (parameteritem.getLastDataUnit() == 1) {
                                CompletableFuture<S7MessageUserData> nextFuture;
                                S7ParameterUserData next_parameter;
                                S7PayloadUserData next_payload;
                                S7PayloadUserDataItemCpuFunctionReadSzlResponse next_payloadItem;

                                while (parameteritem.getLastDataUnit() == 1) {
                                    //TODO: Just wait for one answer!. Pending for other packages for rearm.
                                    nextFuture = reassembledMessage(parameteritem.getSequenceNumber());
                                    S7MessageUserData msg = nextFuture.get();
                                    if (msg != null) {
                                        next_parameter = (S7ParameterUserData) msg.getParameter();
                                        parameteritem = (S7ParameterUserDataItemCPUFunctions) next_parameter.getItems().get(0);
                                        next_payload = (S7PayloadUserData) msg.getPayload();
                                        next_payloadItem = (S7PayloadUserDataItemCpuFunctionReadSzlResponse) next_payload.getItems().get(0);
                                        for (byte b : next_payloadItem.getItems()) plcvalues.add(new PlcSINT(b));
                                    }

                                    plcValue = new PlcList(plcvalues);
                                }

                            } else {
                                plcValue = new PlcList(plcvalues);
                            }
                        } catch (Exception e) {
                            throw new PlcProtocolException("Error decoding PlcValue", e);
                        }

                    }

                }
                if (plcReadRequest.getTag(tagName) instanceof S7AckTag) {
                    S7AckTag field = (S7AckTag) plcReadRequest.getTag(tagName);
                    S7PayloadUserDataItemCpuFunctionAlarmAckResponse payloadItem =
                        (S7PayloadUserDataItemCpuFunctionAlarmAckResponse) payloadItems.get(index);
                    responseCode = decodeResponseCode(payloadItem.getReturnCode());
                    List<Short> data = payloadItem.getMessageObjects();
                    LinkedList<PlcValue> plcValues = new LinkedList<>();
                    for (short b : data) plcValues.add(new PlcSINT((byte) b));
                    plcValue = new PlcList(plcValues);
                }

                ResponseItem<PlcValue> result = new ResponseItem<>(responseCode, plcValue);
                values.put(tagName, result);
                index++;
            }

            return new DefaultPlcReadResponse(plcReadRequest, values);

        }

        // In all other cases all went well.
        S7PayloadReadVarResponse payload = (S7PayloadReadVarResponse) responseMessage.getPayload();

        // If the numbers of items don't match, we're in big trouble as the only
        // way to know how to interpret the responses is by aligning them with the
        // items from the request as this information is not returned by the PLC.
        if (plcReadRequest.getNumberOfTags() != payload.getItems().size()) {
            throw new PlcProtocolException(
                "The number of requested items doesn't match the number of returned items");
        }

        List<S7VarPayloadDataItem> payloadItems = payload.getItems();
        int index = 0;
        for (String tagName : plcReadRequest.getTagNames()) {
            S7Tag tag = (S7Tag) plcReadRequest.getTag(tagName);
            S7VarPayloadDataItem payloadItem = payloadItems.get(index);

            PlcResponseCode responseCode = decodeResponseCode(payloadItem.getReturnCode());
            PlcValue plcValue = null;
            ByteBuf data = Unpooled.wrappedBuffer(payloadItem.getData());
            if (responseCode == PlcResponseCode.OK) {
                try {
                    plcValue = parsePlcValue(tag, data);
                } catch (Exception e) {
                    throw new PlcProtocolException("Error decoding PlcValue", e);
                }
            }
            ResponseItem<PlcValue> result = new ResponseItem<>(responseCode, plcValue);
            values.put(tagName, result);
            index++;
        }

        return new DefaultPlcReadResponse(plcReadRequest, values);
    }

    private PlcResponse decodeWriteResponse(S7Message responseMessage, PlcWriteRequest plcWriteRequest) throws PlcProtocolException {
        Map<String, PlcResponseCode> responses = new HashMap<>();
        short errorClass;
        short errorCode;
        if (responseMessage instanceof S7MessageResponseData) {
            S7MessageResponseData messageResponseData = (S7MessageResponseData) responseMessage;
            errorClass = messageResponseData.getErrorClass();
            errorCode = messageResponseData.getErrorCode();
        } else if (responseMessage instanceof S7MessageResponse) {
            S7MessageResponse messageResponse = (S7MessageResponse) responseMessage;
            errorClass = messageResponse.getErrorClass();
            errorCode = messageResponse.getErrorCode();
        } else {
            throw new PlcProtocolException("Unsupported message type " + responseMessage.getClass().getName());
        }
        // If the result contains any form of non-null error code, handle this instead.
        if ((errorClass != 0) || (errorCode != 0)) {
            // This is usually the case if PUT/GET wasn't enabled on the PLC
            if ((errorClass == 129) && (errorCode == 4)) {
                logger.warn("Got an error response from the PLC. This particular response code usually indicates " +
                    "that PUT/GET is not enabled on the PLC.");
                for (String tagName : plcWriteRequest.getTagNames()) {
                    responses.put(tagName, PlcResponseCode.ACCESS_DENIED);
                }
                return new DefaultPlcWriteResponse(plcWriteRequest, responses);
            } else {
                logger.warn("Got an unknown error response from the PLC. Error Class: {}, Error Code {}. " +
                        "We probably need to implement explicit handling for this, so please file a bug-report " +
                        "on https://issues.apache.org/jira/projects/PLC4X and ideally attach a WireShark dump " +
                        "containing a capture of the communication.",
                    errorClass, errorCode);
                for (String tagName : plcWriteRequest.getTagNames()) {
                    responses.put(tagName, PlcResponseCode.INTERNAL_ERROR);
                }
                return new DefaultPlcWriteResponse(plcWriteRequest, responses);
            }
        }

        // In all other cases all went well.
        S7PayloadWriteVarResponse payload = (S7PayloadWriteVarResponse) responseMessage.getPayload();

        // If the numbers of items don't match, we're in big trouble as the only
        // way to know how to interpret the responses is by aligning them with the
        // items from the request as this information is not returned by the PLC.
        if (plcWriteRequest.getNumberOfTags() != payload.getItems().size()) {
            throw new PlcProtocolException(
                "The number of requested items doesn't match the number of returned items");
        }

        List<S7VarPayloadStatusItem> payloadItems = payload.getItems();
        int index = 0;
        for (String tagName : plcWriteRequest.getTagNames()) {
            S7VarPayloadStatusItem payloadItem = payloadItems.get(index);

            PlcResponseCode responseCode = decodeResponseCode(payloadItem.getReturnCode());
            responses.put(tagName, responseCode);
            index++;
        }

        return new DefaultPlcWriteResponse(plcWriteRequest, responses);
    }

    private S7VarPayloadDataItem serializePlcValue(S7Tag tag, PlcValue plcValue, Boolean hasNext) {
        try {
            DataTransportSize transportSize = tag.getDataType().getDataTransportSize();
            int stringLength = (tag instanceof S7StringTag) ? ((S7StringTag) tag).getStringLength() : 254;
            ByteBuffer byteBuffer = null;
            for (int i = 0; i < tag.getNumberOfElements(); i++) {
                final int lengthInBits = DataItem.getLengthInBits(plcValue.getIndex(i), tag.getDataType().getDataProtocolId(), stringLength);
                final WriteBufferByteBased writeBuffer = new WriteBufferByteBased((int) Math.ceil(((float) lengthInBits) / 8.0f));
                DataItem.staticSerialize(writeBuffer, plcValue.getIndex(i), tag.getDataType().getDataProtocolId(), stringLength);
                // Allocate enough space for all items.
                if (byteBuffer == null) {
                    byteBuffer = ByteBuffer.allocate(writeBuffer.getBytes().length * tag.getNumberOfElements());
                }
                byteBuffer.put(writeBuffer.getBytes());
            }
            if (byteBuffer != null) {
                byte[] data = byteBuffer.array();
                return new S7VarPayloadDataItem(DataTransportErrorCode.OK, transportSize, data/*, hasNext*/);
            }
        } catch (SerializationException e) {
            logger.warn("Error serializing tag item of type: '{}'", tag.getDataType().name(), e);
        }
        return null;
    }

    private PlcValue parsePlcValue(S7Tag tag, ByteBuf data) {
        ReadBuffer readBuffer = new ReadBufferByteBased(data.array());
        try {
            int stringLength = (tag instanceof S7StringTag) ? ((S7StringTag) tag).getStringLength() : 254;
            if (tag.getNumberOfElements() == 1) {
                return DataItem.staticParse(readBuffer, tag.getDataType().getDataProtocolId(),
                    stringLength);
            } else {
                // Fetch all
                final PlcValue[] resultItems = IntStream.range(0, tag.getNumberOfElements()).mapToObj(i -> {
                    try {
                        return DataItem.staticParse(readBuffer, tag.getDataType().getDataProtocolId(),
                            stringLength);
                    } catch (ParseException e) {
                        logger.warn("Error parsing tag item of type: '{}' (at position {}})", tag.getDataType().name(), i, e);
                    }
                    return null;
                }).toArray(PlcValue[]::new);
                return PlcValueHandler.of(resultItems);
            }
        } catch (ParseException e) {
            logger.warn("Error parsing tag item of type: '{}'", tag.getDataType().name(), e);
        }
        return null;
    }

    /**
     * Helper to convert the return codes returned from the S7 into one of our standard
     * PLC4X return codes
     *
     * @param dataTransportErrorCode S7 return code
     * @return PLC4X return code.
     */
    private PlcResponseCode decodeResponseCode(DataTransportErrorCode dataTransportErrorCode) {
        if (dataTransportErrorCode == null) {
            return PlcResponseCode.INTERNAL_ERROR;
        }
        switch (dataTransportErrorCode) {
            case OK:
                return PlcResponseCode.OK;
            case NOT_FOUND:
                return PlcResponseCode.NOT_FOUND;
            case INVALID_ADDRESS:
                return PlcResponseCode.INVALID_ADDRESS;
            case DATA_TYPE_NOT_SUPPORTED:
                return PlcResponseCode.INVALID_DATATYPE;
            default:
                return PlcResponseCode.INTERNAL_ERROR;
        }
    }

    /**
     * Little helper method to parse Siemens article numbers and extract the type of controller.
     *
     * @param articleNumber article number string.
     * @return type of controller.
     */
    private S7ControllerType decodeControllerType(String articleNumber) {
        if (!articleNumber.startsWith("6ES7 ")) {
            return S7ControllerType.ANY;
        }
        String model = articleNumber.substring(articleNumber.indexOf(' ') + 1, articleNumber.indexOf(' ') + 2);
        switch (model) {
            case "2":
                return S7ControllerType.S7_1200;
            case "5":
                return S7ControllerType.S7_1500;
            case "3":
                return S7ControllerType.S7_300;
            case "4":
                return S7ControllerType.S7_400;
            default:
                if (logger.isInfoEnabled()) {
                    logger.info("Looking up unknown article number {}", articleNumber);
                }
                return S7ControllerType.ANY;
        }
    }

    /**
     * Currently we only support the S7 Any type of addresses. This helper simply converts the S7Tag
     * from PLC4X into S7Address objects.
     *
     * @param tag S7Tag instance we need to convert into an S7Address
     * @return the S7Address
     */
    protected S7Address encodeS7Address(PlcTag tag) {
        if (!(tag instanceof S7Tag)) {
            throw new PlcRuntimeException("Unsupported address type " + tag.getClass().getName());
        }
        S7Tag s7Tag = (S7Tag) tag;
        TransportSize transportSize = s7Tag.getDataType();
        int numElements = s7Tag.getNumberOfElements();
        // For these date-types we have to convert the requests to simple byte-array requests
        // As otherwise the S7 will deny them with "Data type not supported" replies.
        if ((transportSize == TransportSize.TIME) /*|| (transportSize == TransportSize.S7_S5TIME)*/ ||
            (transportSize == TransportSize.LTIME) || (transportSize == TransportSize.DATE) ||
            (transportSize == TransportSize.TIME_OF_DAY) || (transportSize == TransportSize.DATE_AND_TIME)) {
            numElements = numElements * transportSize.getSizeInBytes();
            transportSize = TransportSize.BYTE;
        }
        if (transportSize == TransportSize.STRING) {
            transportSize = TransportSize.CHAR;
            int stringLength = (s7Tag instanceof S7StringTag) ? ((S7StringTag) s7Tag).getStringLength() : 254;
            numElements = numElements * (stringLength + 2);
        } else if (transportSize == TransportSize.WSTRING) {
            transportSize = TransportSize.CHAR;
            int stringLength = (s7Tag instanceof S7StringTag) ? ((S7StringTag) s7Tag).getStringLength() : 254;
            numElements = numElements * (stringLength + 2) * 2;
        }
        return new S7AddressAny(transportSize, numElements, s7Tag.getBlockNumber(),
            s7Tag.getMemoryArea(), s7Tag.getByteOffset(), s7Tag.getBitOffset());
    }


    private void cleanFutures() {
        //TODO: Debe ser ejecutado si la conexion esta levanta.
        activeRequests.forEach((f, p) -> {
            CompletableFuture<Object> cf = (CompletableFuture<Object>) f;
            try {
                if (!cf.isDone()) {
                    logger.info("CF");
                    cf.cancel(true);
                    logger.info("ClientCF");
                    ((CompletableFuture<Object>) p.getRight()).completeExceptionally(new PlcRuntimeException("Disconnected"));
                    logger.info("TM");
                    p.getLeft().endRequest();
                }
            } catch (Exception ex) {
                logger.info(ex.toString());
            }
        });
        activeRequests.clear();

    }

    private boolean isConnected() {
        return context.getChannel().attr(S7HMuxImpl.IS_CONNECTED).get();
        //return true;
    }

    private boolean isPrimaryChannel() {
        return context.getChannel().attr(S7HMuxImpl.IS_PRIMARY).get() == null || context.getChannel().attr(S7HMuxImpl.IS_PRIMARY).get();
    }

    private void setChannelFeatures() {
        context.getChannel().attr(S7HMuxImpl.READ_TIME_OUT).set(s7DriverContext.getReadTimeout());
        context.getChannel().attr(S7HMuxImpl.IS_PING_ACTIVE).set(s7DriverContext.getPing());
        context.getChannel().attr(S7HMuxImpl.PING_TIME).set(s7DriverContext.getPingTime());
        context.getChannel().attr(S7HMuxImpl.RETRY_TIME).set(s7DriverContext.getRetryTime());
    }


    private boolean isSubscriptionSupported() {
        return (s7DriverContext.getControllerType() == S7ControllerType.S7_300) ||
            (s7DriverContext.getControllerType() == S7ControllerType.S7_400);
    }

    private CompletableFuture<S7MessageUserData> reassembledMessage(short sequenceNumber) {

        CompletableFuture<S7MessageUserData> future = new CompletableFuture<>();

        //TODO: PDU id is the same, we need check.
        int tpduId = tpduGenerator.getAndIncrement();
        // If we've reached the max value for a 16 bit transaction identifier, reset back to 1
        tpduGenerator.compareAndExchange(0xFFFF, 1);

        TPKTPacket request = createSzlReassembledRequest(tpduId, sequenceNumber);

        context.sendRequest(request)
            .onTimeout(e -> {
                logger.warn("Timeout during Connection establishing, closing channel...");
                //context.getChannel().close();
            })
            .expectResponse(TPKTPacket.class, Duration.ofMillis(1000))
            .check(p -> p.getPayload() instanceof COTPPacketData)
            .unwrap(p -> ((COTPPacketData) p.getPayload()))
            .check(p -> p.getPayload() instanceof S7MessageUserData)
            .unwrap(p -> ((S7MessageUserData) p.getPayload()))
            .check(p -> p.getPayload() instanceof S7PayloadUserData)
            .handle(future::complete);

        return future;
    }

    /*
     *
     */
    private TPKTPacket createSzlReassembledRequest(int tpduId, short sequenceNumber) {
        S7MessageUserData identifyRemoteMessage = new S7MessageUserData(tpduId, new S7ParameterUserData(Arrays.asList(
            new S7ParameterUserDataItemCPUFunctions((short) 0x12, (byte) 0x4, (byte) 0x4, (short) 0x01, sequenceNumber, (short) 0x00, (short) 0x00, 0)
        )), new S7PayloadUserData(Arrays.asList(
            new S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest(
                DataTransportErrorCode.NOT_FOUND,
                DataTransportSize.NULL,
                0x00)
        )));
        COTPPacketData cotpPacketData = new COTPPacketData(null, identifyRemoteMessage, true, (byte) 2);
        return new TPKTPacket(cotpPacketData);
    }

    private CompletableFuture<S7MessageUserData> reassembledAlarmEvents(short sequenceNumber) {
        CompletableFuture<S7MessageUserData> future = new CompletableFuture<>();

        //TODO: PDU id is the same, we need check.
        int tpduId = tpduGenerator.getAndIncrement();
        // If we've reached the max value for a 16 bit transaction identifier, reset back to 1
        tpduGenerator.compareAndExchange(0xFFFF, 1);

        TPKTPacket request = createAlarmQueryReassembledRequest(tpduId, sequenceNumber);

        context.sendRequest(request)
            .onTimeout(e -> {
                logger.warn("Timeout during Connection establishing, closing channel...");
                //context.getChannel().close();
            })
            .expectResponse(TPKTPacket.class, Duration.ofMillis(1000))
            .check(p -> p.getPayload() instanceof COTPPacketData)
            .unwrap(p -> ((COTPPacketData) p.getPayload()))
            .check(p -> p.getPayload() instanceof S7MessageUserData)
            .unwrap(p -> ((S7MessageUserData) p.getPayload()))
            .check(p -> p.getPayload() instanceof S7PayloadUserData)
            .handle(future::complete);

        return future;
    }

    //TODO: S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest to S7PayloadUserDataItemCpuFunctionAlarmQueryNoDataRequest 
    private TPKTPacket createAlarmQueryReassembledRequest(int tpduId, short sequenceNumber) {
        S7MessageUserData identifyRemoteMessage = new S7MessageUserData(tpduId, new S7ParameterUserData(Arrays.asList(
            new S7ParameterUserDataItemCPUFunctions((short) 0x12, (byte) 0x4, (byte) 0x4, (short) 0x13, sequenceNumber, (short) 0x00, (short) 0x00, 0)
        )), new S7PayloadUserData(Arrays.asList(
            new S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest(
                DataTransportErrorCode.NOT_FOUND,
                DataTransportSize.NULL,
                0x00)
        )));
        COTPPacketData cotpPacketData = new COTPPacketData(null, identifyRemoteMessage, true, (byte) 2);
        return new TPKTPacket(cotpPacketData);
    }


    /**
     * A generic purpose error handler which terminates transaction and calls back given future with error message.
     */
    static class TransactionErrorCallback<T, E extends Throwable> implements Consumer<TimeoutException>, BiConsumer<TPKTPacket, E> {

        private final CompletableFuture<T> future;
        private final RequestTransactionManager.RequestTransaction transaction;

        TransactionErrorCallback(CompletableFuture<T> future, RequestTransactionManager.RequestTransaction transaction) {
            this.future = future;
            this.transaction = transaction;
        }

        @Override
        public void accept(TimeoutException e) {
            transaction.endRequest();
            future.completeExceptionally(e);
        }

        @Override
        public void accept(TPKTPacket tpktPacket, E e) {
            transaction.endRequest();
            future.completeExceptionally(e);
        }
    }

}
