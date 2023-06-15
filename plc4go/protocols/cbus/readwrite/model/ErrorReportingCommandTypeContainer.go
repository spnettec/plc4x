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

// ErrorReportingCommandTypeContainer is an enum
type ErrorReportingCommandTypeContainer uint8

type IErrorReportingCommandTypeContainer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NumBytes() uint8
	CommandType() ErrorReportingCommandType
}

const(
	ErrorReportingCommandTypeContainer_ErrorReportingCommandDeprecated ErrorReportingCommandTypeContainer = 0x05
	ErrorReportingCommandTypeContainer_ErrorReportingCommandErrorReport ErrorReportingCommandTypeContainer = 0x15
	ErrorReportingCommandTypeContainer_ErrorReportingCommandAcknowledge ErrorReportingCommandTypeContainer = 0x25
	ErrorReportingCommandTypeContainer_ErrorReportingCommandClearMostSevere ErrorReportingCommandTypeContainer = 0x35
)

var ErrorReportingCommandTypeContainerValues []ErrorReportingCommandTypeContainer

func init() {
	_ = errors.New
	ErrorReportingCommandTypeContainerValues = []ErrorReportingCommandTypeContainer {
		ErrorReportingCommandTypeContainer_ErrorReportingCommandDeprecated,
		ErrorReportingCommandTypeContainer_ErrorReportingCommandErrorReport,
		ErrorReportingCommandTypeContainer_ErrorReportingCommandAcknowledge,
		ErrorReportingCommandTypeContainer_ErrorReportingCommandClearMostSevere,
	}
}


func (e ErrorReportingCommandTypeContainer) NumBytes() uint8 {
	switch e  {
		case 0x05: { /* '0x05' */
            return 5
		}
		case 0x15: { /* '0x15' */
            return 5
		}
		case 0x25: { /* '0x25' */
            return 5
		}
		case 0x35: { /* '0x35' */
            return 5
		}
		default: {
			return 0
		}
	}
}

func ErrorReportingCommandTypeContainerFirstEnumForFieldNumBytes(value uint8) (ErrorReportingCommandTypeContainer, error) {
	for _, sizeValue := range ErrorReportingCommandTypeContainerValues {
		if sizeValue.NumBytes() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumBytes not found", value)
}

func (e ErrorReportingCommandTypeContainer) CommandType() ErrorReportingCommandType {
	switch e  {
		case 0x05: { /* '0x05' */
			return ErrorReportingCommandType_DEPRECATED
		}
		case 0x15: { /* '0x15' */
			return ErrorReportingCommandType_ERROR_REPORT
		}
		case 0x25: { /* '0x25' */
			return ErrorReportingCommandType_ACKNOWLEDGE
		}
		case 0x35: { /* '0x35' */
			return ErrorReportingCommandType_CLEAR_MOST_SEVERE
		}
		default: {
			return 0
		}
	}
}

func ErrorReportingCommandTypeContainerFirstEnumForFieldCommandType(value ErrorReportingCommandType) (ErrorReportingCommandTypeContainer, error) {
	for _, sizeValue := range ErrorReportingCommandTypeContainerValues {
		if sizeValue.CommandType() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing CommandType not found", value)
}
func ErrorReportingCommandTypeContainerByValue(value uint8) (enum ErrorReportingCommandTypeContainer, ok bool) {
	switch value {
		case 0x05:
			return ErrorReportingCommandTypeContainer_ErrorReportingCommandDeprecated, true
		case 0x15:
			return ErrorReportingCommandTypeContainer_ErrorReportingCommandErrorReport, true
		case 0x25:
			return ErrorReportingCommandTypeContainer_ErrorReportingCommandAcknowledge, true
		case 0x35:
			return ErrorReportingCommandTypeContainer_ErrorReportingCommandClearMostSevere, true
	}
	return 0, false
}

func ErrorReportingCommandTypeContainerByName(value string) (enum ErrorReportingCommandTypeContainer, ok bool) {
	switch value {
	case "ErrorReportingCommandDeprecated":
		return ErrorReportingCommandTypeContainer_ErrorReportingCommandDeprecated, true
	case "ErrorReportingCommandErrorReport":
		return ErrorReportingCommandTypeContainer_ErrorReportingCommandErrorReport, true
	case "ErrorReportingCommandAcknowledge":
		return ErrorReportingCommandTypeContainer_ErrorReportingCommandAcknowledge, true
	case "ErrorReportingCommandClearMostSevere":
		return ErrorReportingCommandTypeContainer_ErrorReportingCommandClearMostSevere, true
	}
	return 0, false
}

func ErrorReportingCommandTypeContainerKnows(value uint8)  bool {
	for _, typeValue := range ErrorReportingCommandTypeContainerValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastErrorReportingCommandTypeContainer(structType any) ErrorReportingCommandTypeContainer {
	castFunc := func(typ any) ErrorReportingCommandTypeContainer {
		if sErrorReportingCommandTypeContainer, ok := typ.(ErrorReportingCommandTypeContainer); ok {
			return sErrorReportingCommandTypeContainer
		}
		return 0
	}
	return castFunc(structType)
}

func (m ErrorReportingCommandTypeContainer) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m ErrorReportingCommandTypeContainer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingCommandTypeContainerParse(ctx context.Context, theBytes []byte) (ErrorReportingCommandTypeContainer, error) {
	return ErrorReportingCommandTypeContainerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingCommandTypeContainerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingCommandTypeContainer, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("ErrorReportingCommandTypeContainer", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ErrorReportingCommandTypeContainer")
	}
	if enum, ok := ErrorReportingCommandTypeContainerByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return ErrorReportingCommandTypeContainer(val), nil
	} else {
		return enum, nil
	}
}

func (e ErrorReportingCommandTypeContainer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ErrorReportingCommandTypeContainer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("ErrorReportingCommandTypeContainer", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ErrorReportingCommandTypeContainer) PLC4XEnumName() string {
	switch e {
	case ErrorReportingCommandTypeContainer_ErrorReportingCommandDeprecated:
		return "ErrorReportingCommandDeprecated"
	case ErrorReportingCommandTypeContainer_ErrorReportingCommandErrorReport:
		return "ErrorReportingCommandErrorReport"
	case ErrorReportingCommandTypeContainer_ErrorReportingCommandAcknowledge:
		return "ErrorReportingCommandAcknowledge"
	case ErrorReportingCommandTypeContainer_ErrorReportingCommandClearMostSevere:
		return "ErrorReportingCommandClearMostSevere"
	}
	return ""
}

func (e ErrorReportingCommandTypeContainer) String() string {
	return e.PLC4XEnumName()
}

