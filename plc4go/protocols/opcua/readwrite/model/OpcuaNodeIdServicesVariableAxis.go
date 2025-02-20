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

// OpcuaNodeIdServicesVariableAxis is an enum
type OpcuaNodeIdServicesVariableAxis int32

type IOpcuaNodeIdServicesVariableAxis interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableAxis_AxisScaleEnumeration_EnumStrings OpcuaNodeIdServicesVariableAxis = 12078
)

var OpcuaNodeIdServicesVariableAxisValues []OpcuaNodeIdServicesVariableAxis

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableAxisValues = []OpcuaNodeIdServicesVariableAxis{
		OpcuaNodeIdServicesVariableAxis_AxisScaleEnumeration_EnumStrings,
	}
}

func OpcuaNodeIdServicesVariableAxisByValue(value int32) (enum OpcuaNodeIdServicesVariableAxis, ok bool) {
	switch value {
	case 12078:
		return OpcuaNodeIdServicesVariableAxis_AxisScaleEnumeration_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAxisByName(value string) (enum OpcuaNodeIdServicesVariableAxis, ok bool) {
	switch value {
	case "AxisScaleEnumeration_EnumStrings":
		return OpcuaNodeIdServicesVariableAxis_AxisScaleEnumeration_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAxisKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableAxisValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableAxis(structType any) OpcuaNodeIdServicesVariableAxis {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableAxis {
		if sOpcuaNodeIdServicesVariableAxis, ok := typ.(OpcuaNodeIdServicesVariableAxis); ok {
			return sOpcuaNodeIdServicesVariableAxis
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableAxis) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableAxis) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableAxisParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableAxis, error) {
	return OpcuaNodeIdServicesVariableAxisParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableAxisParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableAxis, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableAxis", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableAxis")
	}
	if enum, ok := OpcuaNodeIdServicesVariableAxisByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableAxis")
		return OpcuaNodeIdServicesVariableAxis(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableAxis) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableAxis) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableAxis", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableAxis) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableAxis_AxisScaleEnumeration_EnumStrings:
		return "AxisScaleEnumeration_EnumStrings"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableAxis) String() string {
	return e.PLC4XEnumName()
}
