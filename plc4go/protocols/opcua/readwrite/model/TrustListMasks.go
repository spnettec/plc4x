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

// TrustListMasks is an enum
type TrustListMasks uint32

type ITrustListMasks interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	TrustListMasks_trustListMasksNone TrustListMasks = 0
	TrustListMasks_trustListMasksTrustedCertificates TrustListMasks = 1
	TrustListMasks_trustListMasksTrustedCrls TrustListMasks = 2
	TrustListMasks_trustListMasksIssuerCertificates TrustListMasks = 4
	TrustListMasks_trustListMasksIssuerCrls TrustListMasks = 8
	TrustListMasks_trustListMasksAll TrustListMasks = 15
)

var TrustListMasksValues []TrustListMasks

func init() {
	_ = errors.New
	TrustListMasksValues = []TrustListMasks {
		TrustListMasks_trustListMasksNone,
		TrustListMasks_trustListMasksTrustedCertificates,
		TrustListMasks_trustListMasksTrustedCrls,
		TrustListMasks_trustListMasksIssuerCertificates,
		TrustListMasks_trustListMasksIssuerCrls,
		TrustListMasks_trustListMasksAll,
	}
}

func TrustListMasksByValue(value uint32) (enum TrustListMasks, ok bool) {
	switch value {
		case 0:
			return TrustListMasks_trustListMasksNone, true
		case 1:
			return TrustListMasks_trustListMasksTrustedCertificates, true
		case 15:
			return TrustListMasks_trustListMasksAll, true
		case 2:
			return TrustListMasks_trustListMasksTrustedCrls, true
		case 4:
			return TrustListMasks_trustListMasksIssuerCertificates, true
		case 8:
			return TrustListMasks_trustListMasksIssuerCrls, true
	}
	return 0, false
}

func TrustListMasksByName(value string) (enum TrustListMasks, ok bool) {
	switch value {
	case "trustListMasksNone":
		return TrustListMasks_trustListMasksNone, true
	case "trustListMasksTrustedCertificates":
		return TrustListMasks_trustListMasksTrustedCertificates, true
	case "trustListMasksAll":
		return TrustListMasks_trustListMasksAll, true
	case "trustListMasksTrustedCrls":
		return TrustListMasks_trustListMasksTrustedCrls, true
	case "trustListMasksIssuerCertificates":
		return TrustListMasks_trustListMasksIssuerCertificates, true
	case "trustListMasksIssuerCrls":
		return TrustListMasks_trustListMasksIssuerCrls, true
	}
	return 0, false
}

func TrustListMasksKnows(value uint32)  bool {
	for _, typeValue := range TrustListMasksValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastTrustListMasks(structType any) TrustListMasks {
	castFunc := func(typ any) TrustListMasks {
		if sTrustListMasks, ok := typ.(TrustListMasks); ok {
			return sTrustListMasks
		}
		return 0
	}
	return castFunc(structType)
}

func (m TrustListMasks) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m TrustListMasks) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TrustListMasksParse(ctx context.Context, theBytes []byte) (TrustListMasks, error) {
	return TrustListMasksParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TrustListMasksParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TrustListMasks, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("TrustListMasks", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TrustListMasks")
	}
	if enum, ok := TrustListMasksByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return TrustListMasks(val), nil
	} else {
		return enum, nil
	}
}

func (e TrustListMasks) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TrustListMasks) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("TrustListMasks", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TrustListMasks) PLC4XEnumName() string {
	switch e {
	case TrustListMasks_trustListMasksNone:
		return "trustListMasksNone"
	case TrustListMasks_trustListMasksTrustedCertificates:
		return "trustListMasksTrustedCertificates"
	case TrustListMasks_trustListMasksAll:
		return "trustListMasksAll"
	case TrustListMasks_trustListMasksTrustedCrls:
		return "trustListMasksTrustedCrls"
	case TrustListMasks_trustListMasksIssuerCertificates:
		return "trustListMasksIssuerCertificates"
	case TrustListMasks_trustListMasksIssuerCrls:
		return "trustListMasksIssuerCrls"
	}
	return ""
}

func (e TrustListMasks) String() string {
	return e.PLC4XEnumName()
}

