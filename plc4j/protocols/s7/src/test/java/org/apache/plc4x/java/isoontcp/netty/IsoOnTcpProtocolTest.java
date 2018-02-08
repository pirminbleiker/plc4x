/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.isoontcp.netty;

import static org.assertj.core.api.Assertions.*;
import static org.mockito.Mockito.*;

import ch.qos.logback.classic.Level;
import ch.qos.logback.classic.Logger;
import ch.qos.logback.classic.spi.LoggingEvent;
import ch.qos.logback.core.Appender;
import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import io.netty.channel.embedded.EmbeddedChannel;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.isoontcp.netty.model.IsoOnTcpMessage;
import org.apache.plc4x.java.netty.NettyTestBase;
import org.apache.plc4x.test.FastTests;
import org.junit.Test;
import org.junit.experimental.categories.Category;
import org.slf4j.LoggerFactory;


public class IsoOnTcpProtocolTest extends NettyTestBase {

    @Test
    @Category(FastTests.class)
    public void encode() {
        IsoOnTcpMessage isoOnTcpMessage = new IsoOnTcpMessage(
            Unpooled.wrappedBuffer(new byte[]{(byte)0x01,(byte)0x02,(byte)0x03}));
        EmbeddedChannel channel = new EmbeddedChannel(new IsoOnTcpProtocol());
        channel.writeOutbound(isoOnTcpMessage);
        channel.checkException();
        Object obj = channel.readOutbound();
        assertThat(obj).isInstanceOf(ByteBuf.class);
        ByteBuf byteBuf = (ByteBuf) obj;
        assertThat(byteBuf.readableBytes()).isEqualTo(4 + 3).withFailMessage("The TCP on ISO Header should add 4 bytes to the data sent");
        assertThat(byteBuf.getByte(0)).isEqualTo(IsoOnTcpProtocol.ISO_ON_TCP_MAGIC_NUMBER);
        assertThat(byteBuf.getShort(2)).isEqualTo((short) (4 + 3)).withFailMessage("The length value in the packet should reflect the size of the entire data being sent");
    }

    /**
     * Happy path test.
     */
    @Test
    @Category(FastTests.class)
    public void decode() {
        EmbeddedChannel channel = new EmbeddedChannel(new IsoOnTcpProtocol());
        channel.writeInbound(Unpooled.wrappedBuffer(new byte[]{IsoOnTcpProtocol.ISO_ON_TCP_MAGIC_NUMBER,
            (byte)0x00,(byte)0x00,(byte)0x0D,
            (byte)0x01,(byte)0x02,(byte)0x03,(byte)0x04,(byte)0x05,(byte)0x06,(byte)0x07,(byte)0x08,(byte)0x09}));
        channel.checkException();
        Object obj = channel.readInbound();
        assertThat(obj).isInstanceOf(IsoOnTcpMessage.class);
        IsoOnTcpMessage isoOnTcpMessage = (IsoOnTcpMessage) obj;
        assertThat(isoOnTcpMessage.getUserData()).isNotNull();
        assertThat(isoOnTcpMessage.getUserData().readableBytes()).isEqualTo(9);
    }

    /**
     * If the packet doesn't start with the ISO on TCP magic byte 0x03
     * an exception should be thrown.
     */
    @Test
    @Category(FastTests.class)
    public void decodeWrongMagicByte() {
        EmbeddedChannel channel = new EmbeddedChannel(new IsoOnTcpProtocol());
        // TODO: Check this is done the same way as in the rest of the project
        Throwable throwable = catchThrowable(() -> channel.writeInbound(Unpooled.wrappedBuffer(new byte[]{0x12,
            (byte)0x00,(byte)0x00,(byte)0x0D,
            (byte)0x01,(byte)0x02,(byte)0x03,(byte)0x04,(byte)0x05,(byte)0x06,(byte)0x07,(byte)0x08,(byte)0x09})));
        assertThat(throwable).isInstanceOf(PlcProtocolException.class);
        assertThat(throwable.getMessage()).contains("ISO on TCP magic number");
    }

    /**
     * If the available amount of data is so small we can't even find out how big
     * the entire package should be, nothing should be read.
     */
    @Test
    @Category(FastTests.class)
    public void decodeWayTooShort() {
        EmbeddedChannel channel = new EmbeddedChannel(new IsoOnTcpProtocol());
        channel.writeInbound(Unpooled.wrappedBuffer(new byte[]{IsoOnTcpProtocol.ISO_ON_TCP_MAGIC_NUMBER,
            (byte)0x00,(byte)0x00,(byte)0x0D}));
        channel.checkException();
        Object obj = channel.readInbound();
        assertThat(obj).isNull();
    }

    /**
     * If the available amount of data is smaller than what the packet size says
     * it should be, nothing should be read.
     */
    @Test
    @Category(FastTests.class)
    public void decodeTooShort() {
        EmbeddedChannel channel = new EmbeddedChannel(new IsoOnTcpProtocol());
        channel.writeInbound(Unpooled.wrappedBuffer(new byte[]{IsoOnTcpProtocol.ISO_ON_TCP_MAGIC_NUMBER,
            (byte)0x00,(byte)0x00,(byte)0x0D,
            (byte)0x01,(byte)0x02,(byte)0x03,(byte)0x04,(byte)0x05,(byte)0x06,(byte)0x07,(byte)0x08}));
        channel.checkException();
        Object obj = channel.readInbound();
        assertThat(obj).isNull();
    }

    /**
     * If logging is set to `DEBUG` then a hexdump of the entire captured packet
     * should be logged
     */
    @Test
    @Category(FastTests.class)
    public void decodeLogPacketIfTraceLogging() {
        // Setup the mock logger.
        Logger root = (Logger) LoggerFactory.getLogger(Logger.ROOT_LOGGER_NAME);
        Appender mockAppender = mock(Appender.class);
        when(mockAppender.getName()).thenReturn("MOCK");
        root.addAppender(mockAppender);
        // Save the current default logging level
        Level defaultLevel = root.getLevel();
        try {
            // Change the logging to TRACE.
            root.setLevel(Level.TRACE);

            // Do some deserialization
            EmbeddedChannel channel = new EmbeddedChannel(new IsoOnTcpProtocol());
            channel.writeInbound(Unpooled.wrappedBuffer(new byte[]{IsoOnTcpProtocol.ISO_ON_TCP_MAGIC_NUMBER,
                (byte) 0x00, (byte) 0x00, (byte) 0x0D,
                (byte) 0x01, (byte) 0x02, (byte) 0x03, (byte) 0x04, (byte) 0x05, (byte) 0x06, (byte) 0x07, (byte) 0x08, (byte) 0x09}));
            channel.checkException();
            Object obj = channel.readInbound();
            assertThat(obj).isNotNull();

            // Check that the packet dump was logged.
            verify(mockAppender).doAppend(argThat(argument ->
                ((LoggingEvent) argument).getFormattedMessage().contains("Got Data: 0300000d010203040506070809")));
        } finally {
            // Reset the log level to the default.
            root.setLevel(defaultLevel);
        }
    }


}
