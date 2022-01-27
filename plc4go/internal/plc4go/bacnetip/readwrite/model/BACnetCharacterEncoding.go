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

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type BACnetCharacterEncoding byte

type IBACnetCharacterEncoding interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetCharacterEncoding_ISO_10646          BACnetCharacterEncoding = 0x0
	BACnetCharacterEncoding_IBM_Microsoft_DBCS BACnetCharacterEncoding = 0x1
	BACnetCharacterEncoding_JIS_X_0208         BACnetCharacterEncoding = 0x2
	BACnetCharacterEncoding_ISO_10646_4        BACnetCharacterEncoding = 0x3
	BACnetCharacterEncoding_ISO_10646_2        BACnetCharacterEncoding = 0x4
	BACnetCharacterEncoding_ISO_8859_1         BACnetCharacterEncoding = 0x5
)

var BACnetCharacterEncodingValues []BACnetCharacterEncoding

func init() {
	_ = errors.New
	BACnetCharacterEncodingValues = []BACnetCharacterEncoding{
		BACnetCharacterEncoding_ISO_10646,
		BACnetCharacterEncoding_IBM_Microsoft_DBCS,
		BACnetCharacterEncoding_JIS_X_0208,
		BACnetCharacterEncoding_ISO_10646_4,
		BACnetCharacterEncoding_ISO_10646_2,
		BACnetCharacterEncoding_ISO_8859_1,
	}
}

func BACnetCharacterEncodingByValue(value byte) BACnetCharacterEncoding {
	switch value {
	case 0x0:
		return BACnetCharacterEncoding_ISO_10646
	case 0x1:
		return BACnetCharacterEncoding_IBM_Microsoft_DBCS
	case 0x2:
		return BACnetCharacterEncoding_JIS_X_0208
	case 0x3:
		return BACnetCharacterEncoding_ISO_10646_4
	case 0x4:
		return BACnetCharacterEncoding_ISO_10646_2
	case 0x5:
		return BACnetCharacterEncoding_ISO_8859_1
	}
	return 0
}

func BACnetCharacterEncodingByName(value string) BACnetCharacterEncoding {
	switch value {
	case "ISO_10646":
		return BACnetCharacterEncoding_ISO_10646
	case "IBM_Microsoft_DBCS":
		return BACnetCharacterEncoding_IBM_Microsoft_DBCS
	case "JIS_X_0208":
		return BACnetCharacterEncoding_JIS_X_0208
	case "ISO_10646_4":
		return BACnetCharacterEncoding_ISO_10646_4
	case "ISO_10646_2":
		return BACnetCharacterEncoding_ISO_10646_2
	case "ISO_8859_1":
		return BACnetCharacterEncoding_ISO_8859_1
	}
	return 0
}

func BACnetCharacterEncodingKnows(value byte) bool {
	for _, typeValue := range BACnetCharacterEncodingValues {
		if byte(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetCharacterEncoding(structType interface{}) BACnetCharacterEncoding {
	castFunc := func(typ interface{}) BACnetCharacterEncoding {
		if sBACnetCharacterEncoding, ok := typ.(BACnetCharacterEncoding); ok {
			return sBACnetCharacterEncoding
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetCharacterEncoding) LengthInBits() uint16 {
	return 8
}

func (m BACnetCharacterEncoding) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetCharacterEncodingParse(readBuffer utils.ReadBuffer) (BACnetCharacterEncoding, error) {
	val, err := readBuffer.ReadByte("BACnetCharacterEncoding")
	if err != nil {
		return 0, nil
	}
	return BACnetCharacterEncodingByValue(val), nil
}

func (e BACnetCharacterEncoding) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteByte("BACnetCharacterEncoding", byte(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BACnetCharacterEncoding) name() string {
	switch e {
	case BACnetCharacterEncoding_ISO_10646:
		return "ISO_10646"
	case BACnetCharacterEncoding_IBM_Microsoft_DBCS:
		return "IBM_Microsoft_DBCS"
	case BACnetCharacterEncoding_JIS_X_0208:
		return "JIS_X_0208"
	case BACnetCharacterEncoding_ISO_10646_4:
		return "ISO_10646_4"
	case BACnetCharacterEncoding_ISO_10646_2:
		return "ISO_10646_2"
	case BACnetCharacterEncoding_ISO_8859_1:
		return "ISO_8859_1"
	}
	return ""
}

func (e BACnetCharacterEncoding) String() string {
	return e.name()
}
