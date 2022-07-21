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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// RequestObsolete is the corresponding interface of RequestObsolete
type RequestObsolete interface {
	utils.LengthAware
	utils.Serializable
	Request
	// GetCalData returns CalData (property field)
	GetCalData() CALData
	// GetAlpha returns Alpha (property field)
	GetAlpha() Alpha
	// GetObsoletePayloadLength returns ObsoletePayloadLength (virtual field)
	GetObsoletePayloadLength() uint16
}

// RequestObsoleteExactly can be used when we want exactly this type and not a type which fulfills RequestObsolete.
// This is useful for switch cases.
type RequestObsoleteExactly interface {
	RequestObsolete
	isRequestObsolete() bool
}

// _RequestObsolete is the data-structure of this message
type _RequestObsolete struct {
	*_Request
	CalData CALData
	Alpha   Alpha

	// Arguments.
	PayloadLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestObsolete) InitializeParent(parent Request, peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination) {
	m.PeekedByte = peekedByte
	m.StartingCR = startingCR
	m.ResetMode = resetMode
	m.SecondPeek = secondPeek
	m.Termination = termination
}

func (m *_RequestObsolete) GetParent() Request {
	return m._Request
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestObsolete) GetCalData() CALData {
	return m.CalData
}

func (m *_RequestObsolete) GetAlpha() Alpha {
	return m.Alpha
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_RequestObsolete) GetObsoletePayloadLength() uint16 {
	alpha := m.Alpha
	_ = alpha
	return uint16(uint16(m.PayloadLength) + uint16(uint16(1)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestObsolete factory function for _RequestObsolete
func NewRequestObsolete(calData CALData, alpha Alpha, peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination, cBusOptions CBusOptions, messageLength uint16, payloadLength uint16) *_RequestObsolete {
	_result := &_RequestObsolete{
		CalData:  calData,
		Alpha:    alpha,
		_Request: NewRequest(peekedByte, startingCR, resetMode, secondPeek, termination, cBusOptions, messageLength),
	}
	_result._Request._RequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRequestObsolete(structType interface{}) RequestObsolete {
	if casted, ok := structType.(RequestObsolete); ok {
		return casted
	}
	if casted, ok := structType.(*RequestObsolete); ok {
		return *casted
	}
	return nil
}

func (m *_RequestObsolete) GetTypeName() string {
	return "RequestObsolete"
}

func (m *_RequestObsolete) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_RequestObsolete) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Manual Field (calData)
	lengthInBits += uint16(int32(int32(int32(m.GetLengthInBytes())*int32(int32(2)))) * int32(int32(8)))

	// Optional Field (alpha)
	if m.Alpha != nil {
		lengthInBits += m.Alpha.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_RequestObsolete) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RequestObsoleteParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions, messageLength uint16, payloadLength uint16) (RequestObsolete, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestObsolete"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestObsolete")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_obsoletePayloadLength := uint16(payloadLength) + uint16(uint16(1))
	obsoletePayloadLength := uint16(_obsoletePayloadLength)
	_ = obsoletePayloadLength

	// Manual Field (calData)
	_calData, _calDataErr := ReadCALData(readBuffer, obsoletePayloadLength)
	if _calDataErr != nil {
		return nil, errors.Wrap(_calDataErr, "Error parsing 'calData' field of RequestObsolete")
	}
	calData := _calData.(CALData)

	// Optional Field (alpha) (Can be skipped, if a given expression evaluates to false)
	var alpha Alpha = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("alpha"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for alpha")
		}
		_val, _err := AlphaParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'alpha' field of RequestObsolete")
		default:
			alpha = _val.(Alpha)
			if closeErr := readBuffer.CloseContext("alpha"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for alpha")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("RequestObsolete"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestObsolete")
	}

	// Create a partially initialized instance
	_child := &_RequestObsolete{
		CalData: calData,
		Alpha:   alpha,
		_Request: &_Request{
			CBusOptions:   cBusOptions,
			MessageLength: messageLength,
		},
	}
	_child._Request._RequestChildRequirements = _child
	return _child, nil
}

func (m *_RequestObsolete) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestObsolete"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestObsolete")
		}
		// Virtual field
		if _obsoletePayloadLengthErr := writeBuffer.WriteVirtual("obsoletePayloadLength", m.GetObsoletePayloadLength()); _obsoletePayloadLengthErr != nil {
			return errors.Wrap(_obsoletePayloadLengthErr, "Error serializing 'obsoletePayloadLength' field")
		}

		// Manual Field (calData)
		_calDataErr := WriteCALData(writeBuffer, m.GetCalData())
		if _calDataErr != nil {
			return errors.Wrap(_calDataErr, "Error serializing 'calData' field")
		}

		// Optional Field (alpha) (Can be skipped, if the value is null)
		var alpha Alpha = nil
		if m.GetAlpha() != nil {
			if pushErr := writeBuffer.PushContext("alpha"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for alpha")
			}
			alpha = m.GetAlpha()
			_alphaErr := writeBuffer.WriteSerializable(alpha)
			if popErr := writeBuffer.PopContext("alpha"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for alpha")
			}
			if _alphaErr != nil {
				return errors.Wrap(_alphaErr, "Error serializing 'alpha' field")
			}
		}

		if popErr := writeBuffer.PopContext("RequestObsolete"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestObsolete")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_RequestObsolete) isRequestObsolete() bool {
	return true
}

func (m *_RequestObsolete) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
