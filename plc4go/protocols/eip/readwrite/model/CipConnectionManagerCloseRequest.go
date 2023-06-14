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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CipConnectionManagerCloseRequest is the corresponding interface of CipConnectionManagerCloseRequest
type CipConnectionManagerCloseRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetRequestPathSize returns RequestPathSize (property field)
	GetRequestPathSize() uint8
	// GetClassSegment returns ClassSegment (property field)
	GetClassSegment() PathSegment
	// GetInstanceSegment returns InstanceSegment (property field)
	GetInstanceSegment() PathSegment
	// GetPriority returns Priority (property field)
	GetPriority() uint8
	// GetTickTime returns TickTime (property field)
	GetTickTime() uint8
	// GetTimeoutTicks returns TimeoutTicks (property field)
	GetTimeoutTicks() uint8
	// GetConnectionSerialNumber returns ConnectionSerialNumber (property field)
	GetConnectionSerialNumber() uint16
	// GetOriginatorVendorId returns OriginatorVendorId (property field)
	GetOriginatorVendorId() uint16
	// GetOriginatorSerialNumber returns OriginatorSerialNumber (property field)
	GetOriginatorSerialNumber() uint32
	// GetConnectionPathSize returns ConnectionPathSize (property field)
	GetConnectionPathSize() uint8
	// GetConnectionPaths returns ConnectionPaths (property field)
	GetConnectionPaths() []PathSegment
}

// CipConnectionManagerCloseRequestExactly can be used when we want exactly this type and not a type which fulfills CipConnectionManagerCloseRequest.
// This is useful for switch cases.
type CipConnectionManagerCloseRequestExactly interface {
	CipConnectionManagerCloseRequest
	isCipConnectionManagerCloseRequest() bool
}

// _CipConnectionManagerCloseRequest is the data-structure of this message
type _CipConnectionManagerCloseRequest struct {
	*_CipService
	RequestPathSize        uint8
	ClassSegment           PathSegment
	InstanceSegment        PathSegment
	Priority               uint8
	TickTime               uint8
	TimeoutTicks           uint8
	ConnectionSerialNumber uint16
	OriginatorVendorId     uint16
	OriginatorSerialNumber uint32
	ConnectionPathSize     uint8
	ConnectionPaths        []PathSegment
	// Reserved Fields
	reservedField0 *byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectionManagerCloseRequest) GetService() uint8 {
	return 0x4E
}

func (m *_CipConnectionManagerCloseRequest) GetResponse() bool {
	return bool(false)
}

func (m *_CipConnectionManagerCloseRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectionManagerCloseRequest) InitializeParent(parent CipService) {}

func (m *_CipConnectionManagerCloseRequest) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectionManagerCloseRequest) GetRequestPathSize() uint8 {
	return m.RequestPathSize
}

func (m *_CipConnectionManagerCloseRequest) GetClassSegment() PathSegment {
	return m.ClassSegment
}

func (m *_CipConnectionManagerCloseRequest) GetInstanceSegment() PathSegment {
	return m.InstanceSegment
}

func (m *_CipConnectionManagerCloseRequest) GetPriority() uint8 {
	return m.Priority
}

func (m *_CipConnectionManagerCloseRequest) GetTickTime() uint8 {
	return m.TickTime
}

func (m *_CipConnectionManagerCloseRequest) GetTimeoutTicks() uint8 {
	return m.TimeoutTicks
}

func (m *_CipConnectionManagerCloseRequest) GetConnectionSerialNumber() uint16 {
	return m.ConnectionSerialNumber
}

func (m *_CipConnectionManagerCloseRequest) GetOriginatorVendorId() uint16 {
	return m.OriginatorVendorId
}

func (m *_CipConnectionManagerCloseRequest) GetOriginatorSerialNumber() uint32 {
	return m.OriginatorSerialNumber
}

func (m *_CipConnectionManagerCloseRequest) GetConnectionPathSize() uint8 {
	return m.ConnectionPathSize
}

func (m *_CipConnectionManagerCloseRequest) GetConnectionPaths() []PathSegment {
	return m.ConnectionPaths
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipConnectionManagerCloseRequest factory function for _CipConnectionManagerCloseRequest
func NewCipConnectionManagerCloseRequest(requestPathSize uint8, classSegment PathSegment, instanceSegment PathSegment, priority uint8, tickTime uint8, timeoutTicks uint8, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, connectionPathSize uint8, connectionPaths []PathSegment, serviceLen uint16) *_CipConnectionManagerCloseRequest {
	_result := &_CipConnectionManagerCloseRequest{
		RequestPathSize:        requestPathSize,
		ClassSegment:           classSegment,
		InstanceSegment:        instanceSegment,
		Priority:               priority,
		TickTime:               tickTime,
		TimeoutTicks:           timeoutTicks,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		ConnectionPathSize:     connectionPathSize,
		ConnectionPaths:        connectionPaths,
		_CipService:            NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipConnectionManagerCloseRequest(structType any) CipConnectionManagerCloseRequest {
	if casted, ok := structType.(CipConnectionManagerCloseRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectionManagerCloseRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectionManagerCloseRequest) GetTypeName() string {
	return "CipConnectionManagerCloseRequest"
}

func (m *_CipConnectionManagerCloseRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestPathSize)
	lengthInBits += 8

	// Simple field (classSegment)
	lengthInBits += m.ClassSegment.GetLengthInBits(ctx)

	// Simple field (instanceSegment)
	lengthInBits += m.InstanceSegment.GetLengthInBits(ctx)

	// Simple field (priority)
	lengthInBits += 4

	// Simple field (tickTime)
	lengthInBits += 4

	// Simple field (timeoutTicks)
	lengthInBits += 8

	// Simple field (connectionSerialNumber)
	lengthInBits += 16

	// Simple field (originatorVendorId)
	lengthInBits += 16

	// Simple field (originatorSerialNumber)
	lengthInBits += 32

	// Simple field (connectionPathSize)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Array field
	if len(m.ConnectionPaths) > 0 {
		for _, element := range m.ConnectionPaths {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_CipConnectionManagerCloseRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CipConnectionManagerCloseRequestParse(ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (CipConnectionManagerCloseRequest, error) {
	return CipConnectionManagerCloseRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func CipConnectionManagerCloseRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (CipConnectionManagerCloseRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipConnectionManagerCloseRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectionManagerCloseRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestPathSize)
	_requestPathSize, _requestPathSizeErr := readBuffer.ReadUint8("requestPathSize", 8)
	if _requestPathSizeErr != nil {
		return nil, errors.Wrap(_requestPathSizeErr, "Error parsing 'requestPathSize' field of CipConnectionManagerCloseRequest")
	}
	requestPathSize := _requestPathSize

	// Simple Field (classSegment)
	if pullErr := readBuffer.PullContext("classSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for classSegment")
	}
	_classSegment, _classSegmentErr := PathSegmentParseWithBuffer(ctx, readBuffer)
	if _classSegmentErr != nil {
		return nil, errors.Wrap(_classSegmentErr, "Error parsing 'classSegment' field of CipConnectionManagerCloseRequest")
	}
	classSegment := _classSegment.(PathSegment)
	if closeErr := readBuffer.CloseContext("classSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for classSegment")
	}

	// Simple Field (instanceSegment)
	if pullErr := readBuffer.PullContext("instanceSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for instanceSegment")
	}
	_instanceSegment, _instanceSegmentErr := PathSegmentParseWithBuffer(ctx, readBuffer)
	if _instanceSegmentErr != nil {
		return nil, errors.Wrap(_instanceSegmentErr, "Error parsing 'instanceSegment' field of CipConnectionManagerCloseRequest")
	}
	instanceSegment := _instanceSegment.(PathSegment)
	if closeErr := readBuffer.CloseContext("instanceSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for instanceSegment")
	}

	// Simple Field (priority)
	_priority, _priorityErr := readBuffer.ReadUint8("priority", 4)
	if _priorityErr != nil {
		return nil, errors.Wrap(_priorityErr, "Error parsing 'priority' field of CipConnectionManagerCloseRequest")
	}
	priority := _priority

	// Simple Field (tickTime)
	_tickTime, _tickTimeErr := readBuffer.ReadUint8("tickTime", 4)
	if _tickTimeErr != nil {
		return nil, errors.Wrap(_tickTimeErr, "Error parsing 'tickTime' field of CipConnectionManagerCloseRequest")
	}
	tickTime := _tickTime

	// Simple Field (timeoutTicks)
	_timeoutTicks, _timeoutTicksErr := readBuffer.ReadUint8("timeoutTicks", 8)
	if _timeoutTicksErr != nil {
		return nil, errors.Wrap(_timeoutTicksErr, "Error parsing 'timeoutTicks' field of CipConnectionManagerCloseRequest")
	}
	timeoutTicks := _timeoutTicks

	// Simple Field (connectionSerialNumber)
	_connectionSerialNumber, _connectionSerialNumberErr := readBuffer.ReadUint16("connectionSerialNumber", 16)
	if _connectionSerialNumberErr != nil {
		return nil, errors.Wrap(_connectionSerialNumberErr, "Error parsing 'connectionSerialNumber' field of CipConnectionManagerCloseRequest")
	}
	connectionSerialNumber := _connectionSerialNumber

	// Simple Field (originatorVendorId)
	_originatorVendorId, _originatorVendorIdErr := readBuffer.ReadUint16("originatorVendorId", 16)
	if _originatorVendorIdErr != nil {
		return nil, errors.Wrap(_originatorVendorIdErr, "Error parsing 'originatorVendorId' field of CipConnectionManagerCloseRequest")
	}
	originatorVendorId := _originatorVendorId

	// Simple Field (originatorSerialNumber)
	_originatorSerialNumber, _originatorSerialNumberErr := readBuffer.ReadUint32("originatorSerialNumber", 32)
	if _originatorSerialNumberErr != nil {
		return nil, errors.Wrap(_originatorSerialNumberErr, "Error parsing 'originatorSerialNumber' field of CipConnectionManagerCloseRequest")
	}
	originatorSerialNumber := _originatorSerialNumber

	// Simple Field (connectionPathSize)
	_connectionPathSize, _connectionPathSizeErr := readBuffer.ReadUint8("connectionPathSize", 8)
	if _connectionPathSizeErr != nil {
		return nil, errors.Wrap(_connectionPathSizeErr, "Error parsing 'connectionPathSize' field of CipConnectionManagerCloseRequest")
	}
	connectionPathSize := _connectionPathSize

	var reservedField0 *byte
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CipConnectionManagerCloseRequest")
		}
		if reserved != byte(0x00) {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": byte(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Array field (connectionPaths)
	if pullErr := readBuffer.PullContext("connectionPaths", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for connectionPaths")
	}
	// Terminated array
	var connectionPaths []PathSegment
	{
		for !bool(NoMorePathSegments(ctx, readBuffer)) {
			_item, _err := PathSegmentParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'connectionPaths' field of CipConnectionManagerCloseRequest")
			}
			connectionPaths = append(connectionPaths, _item.(PathSegment))
		}
	}
	if closeErr := readBuffer.CloseContext("connectionPaths", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for connectionPaths")
	}

	if closeErr := readBuffer.CloseContext("CipConnectionManagerCloseRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectionManagerCloseRequest")
	}

	// Create a partially initialized instance
	_child := &_CipConnectionManagerCloseRequest{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		RequestPathSize:        requestPathSize,
		ClassSegment:           classSegment,
		InstanceSegment:        instanceSegment,
		Priority:               priority,
		TickTime:               tickTime,
		TimeoutTicks:           timeoutTicks,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		ConnectionPathSize:     connectionPathSize,
		ConnectionPaths:        connectionPaths,
		reservedField0:         reservedField0,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_CipConnectionManagerCloseRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectionManagerCloseRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectionManagerCloseRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectionManagerCloseRequest")
		}

		// Simple Field (requestPathSize)
		requestPathSize := uint8(m.GetRequestPathSize())
		_requestPathSizeErr := writeBuffer.WriteUint8("requestPathSize", 8, (requestPathSize))
		if _requestPathSizeErr != nil {
			return errors.Wrap(_requestPathSizeErr, "Error serializing 'requestPathSize' field")
		}

		// Simple Field (classSegment)
		if pushErr := writeBuffer.PushContext("classSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for classSegment")
		}
		_classSegmentErr := writeBuffer.WriteSerializable(ctx, m.GetClassSegment())
		if popErr := writeBuffer.PopContext("classSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for classSegment")
		}
		if _classSegmentErr != nil {
			return errors.Wrap(_classSegmentErr, "Error serializing 'classSegment' field")
		}

		// Simple Field (instanceSegment)
		if pushErr := writeBuffer.PushContext("instanceSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for instanceSegment")
		}
		_instanceSegmentErr := writeBuffer.WriteSerializable(ctx, m.GetInstanceSegment())
		if popErr := writeBuffer.PopContext("instanceSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for instanceSegment")
		}
		if _instanceSegmentErr != nil {
			return errors.Wrap(_instanceSegmentErr, "Error serializing 'instanceSegment' field")
		}

		// Simple Field (priority)
		priority := uint8(m.GetPriority())
		_priorityErr := writeBuffer.WriteUint8("priority", 4, (priority))
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}

		// Simple Field (tickTime)
		tickTime := uint8(m.GetTickTime())
		_tickTimeErr := writeBuffer.WriteUint8("tickTime", 4, (tickTime))
		if _tickTimeErr != nil {
			return errors.Wrap(_tickTimeErr, "Error serializing 'tickTime' field")
		}

		// Simple Field (timeoutTicks)
		timeoutTicks := uint8(m.GetTimeoutTicks())
		_timeoutTicksErr := writeBuffer.WriteUint8("timeoutTicks", 8, (timeoutTicks))
		if _timeoutTicksErr != nil {
			return errors.Wrap(_timeoutTicksErr, "Error serializing 'timeoutTicks' field")
		}

		// Simple Field (connectionSerialNumber)
		connectionSerialNumber := uint16(m.GetConnectionSerialNumber())
		_connectionSerialNumberErr := writeBuffer.WriteUint16("connectionSerialNumber", 16, (connectionSerialNumber))
		if _connectionSerialNumberErr != nil {
			return errors.Wrap(_connectionSerialNumberErr, "Error serializing 'connectionSerialNumber' field")
		}

		// Simple Field (originatorVendorId)
		originatorVendorId := uint16(m.GetOriginatorVendorId())
		_originatorVendorIdErr := writeBuffer.WriteUint16("originatorVendorId", 16, (originatorVendorId))
		if _originatorVendorIdErr != nil {
			return errors.Wrap(_originatorVendorIdErr, "Error serializing 'originatorVendorId' field")
		}

		// Simple Field (originatorSerialNumber)
		originatorSerialNumber := uint32(m.GetOriginatorSerialNumber())
		_originatorSerialNumberErr := writeBuffer.WriteUint32("originatorSerialNumber", 32, (originatorSerialNumber))
		if _originatorSerialNumberErr != nil {
			return errors.Wrap(_originatorSerialNumberErr, "Error serializing 'originatorSerialNumber' field")
		}

		// Simple Field (connectionPathSize)
		connectionPathSize := uint8(m.GetConnectionPathSize())
		_connectionPathSizeErr := writeBuffer.WriteUint8("connectionPathSize", 8, (connectionPathSize))
		if _connectionPathSizeErr != nil {
			return errors.Wrap(_connectionPathSizeErr, "Error serializing 'connectionPathSize' field")
		}

		// Reserved Field (reserved)
		{
			var reserved byte = byte(0x00)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]any{
					"expected value": byte(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteByte("reserved", reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Array Field (connectionPaths)
		if pushErr := writeBuffer.PushContext("connectionPaths", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for connectionPaths")
		}
		for _curItem, _element := range m.GetConnectionPaths() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetConnectionPaths()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'connectionPaths' field")
			}
		}
		if popErr := writeBuffer.PopContext("connectionPaths", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for connectionPaths")
		}

		if popErr := writeBuffer.PopContext("CipConnectionManagerCloseRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectionManagerCloseRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectionManagerCloseRequest) isCipConnectionManagerCloseRequest() bool {
	return true
}

func (m *_CipConnectionManagerCloseRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
