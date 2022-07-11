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

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// RequestContext is the corresponding interface of RequestContext
type RequestContext interface {
	utils.LengthAware
	utils.Serializable
	// GetSendCalCommandBefore returns SendCalCommandBefore (property field)
	GetSendCalCommandBefore() bool
	// GetSendSALStatusRequestBefore returns SendSALStatusRequestBefore (property field)
	GetSendSALStatusRequestBefore() bool
	// GetSendIdentifyRequestBefore returns SendIdentifyRequestBefore (property field)
	GetSendIdentifyRequestBefore() bool
}

// RequestContextExactly can be used when we want exactly this type and not a type which fulfills RequestContext.
// This is useful for switch cases.
type RequestContextExactly interface {
	RequestContext
	isRequestContext() bool
}

// _RequestContext is the data-structure of this message
type _RequestContext struct {
	SendCalCommandBefore       bool
	SendSALStatusRequestBefore bool
	SendIdentifyRequestBefore  bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestContext) GetSendCalCommandBefore() bool {
	return m.SendCalCommandBefore
}

func (m *_RequestContext) GetSendSALStatusRequestBefore() bool {
	return m.SendSALStatusRequestBefore
}

func (m *_RequestContext) GetSendIdentifyRequestBefore() bool {
	return m.SendIdentifyRequestBefore
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestContext factory function for _RequestContext
func NewRequestContext(sendCalCommandBefore bool, sendSALStatusRequestBefore bool, sendIdentifyRequestBefore bool) *_RequestContext {
	return &_RequestContext{SendCalCommandBefore: sendCalCommandBefore, SendSALStatusRequestBefore: sendSALStatusRequestBefore, SendIdentifyRequestBefore: sendIdentifyRequestBefore}
}

// Deprecated: use the interface for direct cast
func CastRequestContext(structType interface{}) RequestContext {
	if casted, ok := structType.(RequestContext); ok {
		return casted
	}
	if casted, ok := structType.(*RequestContext); ok {
		return *casted
	}
	return nil
}

func (m *_RequestContext) GetTypeName() string {
	return "RequestContext"
}

func (m *_RequestContext) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_RequestContext) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (sendCalCommandBefore)
	lengthInBits += 1

	// Simple field (sendSALStatusRequestBefore)
	lengthInBits += 1

	// Simple field (sendIdentifyRequestBefore)
	lengthInBits += 1

	return lengthInBits
}

func (m *_RequestContext) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RequestContextParse(readBuffer utils.ReadBuffer) (RequestContext, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestContext"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestContext")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (sendCalCommandBefore)
	_sendCalCommandBefore, _sendCalCommandBeforeErr := readBuffer.ReadBit("sendCalCommandBefore")
	if _sendCalCommandBeforeErr != nil {
		return nil, errors.Wrap(_sendCalCommandBeforeErr, "Error parsing 'sendCalCommandBefore' field of RequestContext")
	}
	sendCalCommandBefore := _sendCalCommandBefore

	// Simple Field (sendSALStatusRequestBefore)
	_sendSALStatusRequestBefore, _sendSALStatusRequestBeforeErr := readBuffer.ReadBit("sendSALStatusRequestBefore")
	if _sendSALStatusRequestBeforeErr != nil {
		return nil, errors.Wrap(_sendSALStatusRequestBeforeErr, "Error parsing 'sendSALStatusRequestBefore' field of RequestContext")
	}
	sendSALStatusRequestBefore := _sendSALStatusRequestBefore

	// Simple Field (sendIdentifyRequestBefore)
	_sendIdentifyRequestBefore, _sendIdentifyRequestBeforeErr := readBuffer.ReadBit("sendIdentifyRequestBefore")
	if _sendIdentifyRequestBeforeErr != nil {
		return nil, errors.Wrap(_sendIdentifyRequestBeforeErr, "Error parsing 'sendIdentifyRequestBefore' field of RequestContext")
	}
	sendIdentifyRequestBefore := _sendIdentifyRequestBefore

	if closeErr := readBuffer.CloseContext("RequestContext"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestContext")
	}

	// Create the instance
	return NewRequestContext(sendCalCommandBefore, sendSALStatusRequestBefore, sendIdentifyRequestBefore), nil
}

func (m *_RequestContext) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("RequestContext"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for RequestContext")
	}

	// Simple Field (sendCalCommandBefore)
	sendCalCommandBefore := bool(m.GetSendCalCommandBefore())
	_sendCalCommandBeforeErr := writeBuffer.WriteBit("sendCalCommandBefore", (sendCalCommandBefore))
	if _sendCalCommandBeforeErr != nil {
		return errors.Wrap(_sendCalCommandBeforeErr, "Error serializing 'sendCalCommandBefore' field")
	}

	// Simple Field (sendSALStatusRequestBefore)
	sendSALStatusRequestBefore := bool(m.GetSendSALStatusRequestBefore())
	_sendSALStatusRequestBeforeErr := writeBuffer.WriteBit("sendSALStatusRequestBefore", (sendSALStatusRequestBefore))
	if _sendSALStatusRequestBeforeErr != nil {
		return errors.Wrap(_sendSALStatusRequestBeforeErr, "Error serializing 'sendSALStatusRequestBefore' field")
	}

	// Simple Field (sendIdentifyRequestBefore)
	sendIdentifyRequestBefore := bool(m.GetSendIdentifyRequestBefore())
	_sendIdentifyRequestBeforeErr := writeBuffer.WriteBit("sendIdentifyRequestBefore", (sendIdentifyRequestBefore))
	if _sendIdentifyRequestBeforeErr != nil {
		return errors.Wrap(_sendIdentifyRequestBeforeErr, "Error serializing 'sendIdentifyRequestBefore' field")
	}

	if popErr := writeBuffer.PopContext("RequestContext"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for RequestContext")
	}
	return nil
}

func (m *_RequestContext) isRequestContext() bool {
	return true
}

func (m *_RequestContext) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
