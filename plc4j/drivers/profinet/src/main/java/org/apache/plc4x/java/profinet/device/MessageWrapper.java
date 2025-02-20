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

package org.apache.plc4x.java.profinet.device;

import org.apache.plc4x.java.profinet.context.ProfinetDeviceContext;
import org.apache.plc4x.java.profinet.readwrite.DceRpc_Packet;
import org.apache.plc4x.java.profinet.readwrite.Ethernet_Frame;

public interface MessageWrapper {
    void sendUdpMessage(ProfinetCallable<DceRpc_Packet> callable, ProfinetDeviceContext context) throws RuntimeException;
    void sendPnioMessage(ProfinetCallable<Ethernet_Frame> callable, ProfinetDeviceContext context) throws RuntimeException;
}
