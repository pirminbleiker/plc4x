/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.profinet.config;

import org.apache.plc4x.java.spi.configuration.annotations.ConfigurationParameter;
import org.apache.plc4x.java.spi.configuration.annotations.defaults.BooleanDefaultValue;
import org.apache.plc4x.java.transport.rawsocket.DefaultRawSocketTransportConfiguration;
import org.apache.plc4x.java.utils.pcap.netty.handlers.PacketHandler;

public class ProfinetRawSocketTransportConfiguration extends DefaultRawSocketTransportConfiguration {

    @ConfigurationParameter("resolve-mac-address")
    @BooleanDefaultValue(true)
    private boolean resolveMacAddress;

    @Override
    public boolean isResolveMacAccess() {
        return resolveMacAddress;
    }

    public void setResolveMacAddress(boolean resolveMacAddress) {
        this.resolveMacAddress = resolveMacAddress;
    }

    @Override
    public int getDefaultPort() {
        return 34964;
    }

    @Override
    public PacketHandler getPcapPacketHandler() {
        return null;
    }

}
