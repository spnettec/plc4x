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

// Duplex is an enum
type Duplex uint32

type IDuplex interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	Duplex_duplexFull Duplex = 0
	Duplex_duplexHalf Duplex = 1
	Duplex_duplexUnknown Duplex = 2
)

var DuplexValues []Duplex

func init() {
	_ = errors.New
	DuplexValues = []Duplex {
		Duplex_duplexFull,
		Duplex_duplexHalf,
		Duplex_duplexUnknown,
	}
}

func DuplexByValue(value uint32) (enum Duplex, ok bool) {
	switch value {
		case 0:
			return Duplex_duplexFull, true
		case 1:
			return Duplex_duplexHalf, true
		case 2:
			return Duplex_duplexUnknown, true
	}
	return 0, false
}

func DuplexByName(value string) (enum Duplex, ok bool) {
	switch value {
	case "duplexFull":
		return Duplex_duplexFull, true
	case "duplexHalf":
		return Duplex_duplexHalf, true
	case "duplexUnknown":
		return Duplex_duplexUnknown, true
	}
	return 0, false
}

func DuplexKnows(value uint32)  bool {
	for _, typeValue := range DuplexValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastDuplex(structType any) Duplex {
	castFunc := func(typ any) Duplex {
		if sDuplex, ok := typ.(Duplex); ok {
			return sDuplex
		}
		return 0
	}
	return castFunc(structType)
}

func (m Duplex) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m Duplex) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DuplexParse(ctx context.Context, theBytes []byte) (Duplex, error) {
	return DuplexParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DuplexParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Duplex, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("Duplex", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading Duplex")
	}
	if enum, ok := DuplexByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return Duplex(val), nil
	} else {
		return enum, nil
	}
}

func (e Duplex) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e Duplex) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("Duplex", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e Duplex) PLC4XEnumName() string {
	switch e {
	case Duplex_duplexFull:
		return "duplexFull"
	case Duplex_duplexHalf:
		return "duplexHalf"
	case Duplex_duplexUnknown:
		return "duplexUnknown"
	}
	return ""
}

func (e Duplex) String() string {
	return e.PLC4XEnumName()
}

