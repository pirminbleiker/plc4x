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
package org.apache.plc4x.java.transport.tcp;

import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.spi.configuration.HasConfiguration;
import org.apache.plc4x.java.spi.connection.ChannelFactory;
import org.apache.plc4x.java.spi.transport.Transport;
import org.apache.plc4x.java.spi.transport.TransportConfiguration;

import java.net.InetSocketAddress;
import java.net.SocketAddress;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class TcpTransport implements Transport, HasConfiguration<TcpTransportConfiguration> {

    private static final Pattern TRANSPORT_TCP_PATTERN = Pattern.compile(
        "^((?<ip>[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3})|(?<hostname>[a-zA-Z0-9.\\-]+))(:(?<port>[0-9]{1,5}))?.*");

    private TcpTransportConfiguration configuration;

    @Override
    public String getTransportCode() {
        return "tcp";
    }

    @Override
    public String getTransportName() {
        return "IP/TCP Transport";
    }

    @Override
    public void setConfiguration(TcpTransportConfiguration configuration) {
        this.configuration = configuration;
    }

    @Override
    public ChannelFactory createChannelFactory(String transportConfig) {
        final Matcher matcher = TRANSPORT_TCP_PATTERN.matcher(transportConfig);
        if (!matcher.matches()) {
            throw new PlcRuntimeException("Invalid url for TCP transport: " + transportConfig);
        }
        String ip = matcher.group("ip");
        String hostname = matcher.group("hostname");
        String portString = matcher.group("port");

        // If the port wasn't specified, try to get a default port from the configuration.
        int port;
        if (portString != null) {
            port = Integer.parseInt(portString);
        } else if ((configuration != null) &&
            (configuration.getDefaultPort() != TcpTransportConfiguration.NO_DEFAULT_PORT)) {
            port = configuration.getDefaultPort();
        } else {
            throw new PlcRuntimeException("No port defined");
        }

        // Create the fully qualified remote socket address which we should connect to.
        SocketAddress address = new InetSocketAddress((ip == null) ? hostname : ip, port);

        // Initialize the channel factory with the default socket address we want to connect to.
        TcpChannelFactory tcpChannelFactory = new TcpChannelFactory(address);
        if(configuration != null) {
            tcpChannelFactory.setConfiguration(configuration);
        }
        return tcpChannelFactory;
    }

    @Override
    public Class<? extends TransportConfiguration> getTransportConfigType() {
        return DefaultTcpTransportConfiguration.class;
    }

}
