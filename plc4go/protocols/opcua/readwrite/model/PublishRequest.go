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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// PublishRequest is the corresponding interface of PublishRequest
type PublishRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfSubscriptionAcknowledgements returns NoOfSubscriptionAcknowledgements (property field)
	GetNoOfSubscriptionAcknowledgements() int32
	// GetSubscriptionAcknowledgements returns SubscriptionAcknowledgements (property field)
	GetSubscriptionAcknowledgements() []ExtensionObjectDefinition
}

// PublishRequestExactly can be used when we want exactly this type and not a type which fulfills PublishRequest.
// This is useful for switch cases.
type PublishRequestExactly interface {
	PublishRequest
	isPublishRequest() bool
}

// _PublishRequest is the data-structure of this message
type _PublishRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader                    ExtensionObjectDefinition
	NoOfSubscriptionAcknowledgements int32
	SubscriptionAcknowledgements     []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PublishRequest) GetIdentifier() string {
	return "826"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PublishRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_PublishRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PublishRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_PublishRequest) GetNoOfSubscriptionAcknowledgements() int32 {
	return m.NoOfSubscriptionAcknowledgements
}

func (m *_PublishRequest) GetSubscriptionAcknowledgements() []ExtensionObjectDefinition {
	return m.SubscriptionAcknowledgements
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPublishRequest factory function for _PublishRequest
func NewPublishRequest(requestHeader ExtensionObjectDefinition, noOfSubscriptionAcknowledgements int32, subscriptionAcknowledgements []ExtensionObjectDefinition) *_PublishRequest {
	_result := &_PublishRequest{
		RequestHeader:                    requestHeader,
		NoOfSubscriptionAcknowledgements: noOfSubscriptionAcknowledgements,
		SubscriptionAcknowledgements:     subscriptionAcknowledgements,
		_ExtensionObjectDefinition:       NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPublishRequest(structType any) PublishRequest {
	if casted, ok := structType.(PublishRequest); ok {
		return casted
	}
	if casted, ok := structType.(*PublishRequest); ok {
		return *casted
	}
	return nil
}

func (m *_PublishRequest) GetTypeName() string {
	return "PublishRequest"
}

func (m *_PublishRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfSubscriptionAcknowledgements)
	lengthInBits += 32

	// Array field
	if len(m.SubscriptionAcknowledgements) > 0 {
		for _curItem, element := range m.SubscriptionAcknowledgements {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.SubscriptionAcknowledgements), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_PublishRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PublishRequestParse(ctx context.Context, theBytes []byte, identifier string) (PublishRequest, error) {
	return PublishRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func PublishRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (PublishRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PublishRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PublishRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of PublishRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (noOfSubscriptionAcknowledgements)
	_noOfSubscriptionAcknowledgements, _noOfSubscriptionAcknowledgementsErr := readBuffer.ReadInt32("noOfSubscriptionAcknowledgements", 32)
	if _noOfSubscriptionAcknowledgementsErr != nil {
		return nil, errors.Wrap(_noOfSubscriptionAcknowledgementsErr, "Error parsing 'noOfSubscriptionAcknowledgements' field of PublishRequest")
	}
	noOfSubscriptionAcknowledgements := _noOfSubscriptionAcknowledgements

	// Array field (subscriptionAcknowledgements)
	if pullErr := readBuffer.PullContext("subscriptionAcknowledgements", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for subscriptionAcknowledgements")
	}
	// Count array
	subscriptionAcknowledgements := make([]ExtensionObjectDefinition, utils.Max(noOfSubscriptionAcknowledgements, 0))
	// This happens when the size is set conditional to 0
	if len(subscriptionAcknowledgements) == 0 {
		subscriptionAcknowledgements = nil
	}
	{
		_numItems := uint16(utils.Max(noOfSubscriptionAcknowledgements, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "823")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'subscriptionAcknowledgements' field of PublishRequest")
			}
			subscriptionAcknowledgements[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("subscriptionAcknowledgements", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for subscriptionAcknowledgements")
	}

	if closeErr := readBuffer.CloseContext("PublishRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PublishRequest")
	}

	// Create a partially initialized instance
	_child := &_PublishRequest{
		_ExtensionObjectDefinition:       &_ExtensionObjectDefinition{},
		RequestHeader:                    requestHeader,
		NoOfSubscriptionAcknowledgements: noOfSubscriptionAcknowledgements,
		SubscriptionAcknowledgements:     subscriptionAcknowledgements,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_PublishRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PublishRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PublishRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PublishRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (noOfSubscriptionAcknowledgements)
		noOfSubscriptionAcknowledgements := int32(m.GetNoOfSubscriptionAcknowledgements())
		_noOfSubscriptionAcknowledgementsErr := writeBuffer.WriteInt32("noOfSubscriptionAcknowledgements", 32, (noOfSubscriptionAcknowledgements))
		if _noOfSubscriptionAcknowledgementsErr != nil {
			return errors.Wrap(_noOfSubscriptionAcknowledgementsErr, "Error serializing 'noOfSubscriptionAcknowledgements' field")
		}

		// Array Field (subscriptionAcknowledgements)
		if pushErr := writeBuffer.PushContext("subscriptionAcknowledgements", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for subscriptionAcknowledgements")
		}
		for _curItem, _element := range m.GetSubscriptionAcknowledgements() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetSubscriptionAcknowledgements()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'subscriptionAcknowledgements' field")
			}
		}
		if popErr := writeBuffer.PopContext("subscriptionAcknowledgements", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for subscriptionAcknowledgements")
		}

		if popErr := writeBuffer.PopContext("PublishRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PublishRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PublishRequest) isPublishRequest() bool {
	return true
}

func (m *_PublishRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
