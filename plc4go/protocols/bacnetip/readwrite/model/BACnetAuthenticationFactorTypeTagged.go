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

// BACnetAuthenticationFactorTypeTagged is the corresponding interface of BACnetAuthenticationFactorTypeTagged
type BACnetAuthenticationFactorTypeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAuthenticationFactorType
}

// BACnetAuthenticationFactorTypeTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetAuthenticationFactorTypeTagged.
// This is useful for switch cases.
type BACnetAuthenticationFactorTypeTaggedExactly interface {
	BACnetAuthenticationFactorTypeTagged
	isBACnetAuthenticationFactorTypeTagged() bool
}

// _BACnetAuthenticationFactorTypeTagged is the data-structure of this message
type _BACnetAuthenticationFactorTypeTagged struct {
	Header BACnetTagHeader
	Value  BACnetAuthenticationFactorType

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthenticationFactorTypeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAuthenticationFactorTypeTagged) GetValue() BACnetAuthenticationFactorType {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAuthenticationFactorTypeTagged factory function for _BACnetAuthenticationFactorTypeTagged
func NewBACnetAuthenticationFactorTypeTagged(header BACnetTagHeader, value BACnetAuthenticationFactorType, tagNumber uint8, tagClass TagClass) *_BACnetAuthenticationFactorTypeTagged {
	return &_BACnetAuthenticationFactorTypeTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetAuthenticationFactorTypeTagged(structType any) BACnetAuthenticationFactorTypeTagged {
	if casted, ok := structType.(BACnetAuthenticationFactorTypeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthenticationFactorTypeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthenticationFactorTypeTagged) GetTypeName() string {
	return "BACnetAuthenticationFactorTypeTagged"
}

func (m *_BACnetAuthenticationFactorTypeTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetAuthenticationFactorTypeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationFactorTypeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAuthenticationFactorTypeTagged, error) {
	return BACnetAuthenticationFactorTypeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAuthenticationFactorTypeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAuthenticationFactorTypeTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationFactorTypeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationFactorTypeTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetAuthenticationFactorTypeTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetAuthenticationFactorType_UNDEFINED)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetAuthenticationFactorTypeTagged")
	}
	var value BACnetAuthenticationFactorType
	if _value != nil {
		value = _value.(BACnetAuthenticationFactorType)
	}

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationFactorTypeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationFactorTypeTagged")
	}

	// Create the instance
	return &_BACnetAuthenticationFactorTypeTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetAuthenticationFactorTypeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthenticationFactorTypeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationFactorTypeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationFactorTypeTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(ctx, writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationFactorTypeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationFactorTypeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAuthenticationFactorTypeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAuthenticationFactorTypeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAuthenticationFactorTypeTagged) isBACnetAuthenticationFactorTypeTagged() bool {
	return true
}

func (m *_BACnetAuthenticationFactorTypeTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
