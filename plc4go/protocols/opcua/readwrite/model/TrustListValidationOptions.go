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

// TrustListValidationOptions is an enum
type TrustListValidationOptions uint32

type ITrustListValidationOptions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	TrustListValidationOptions_trustListValidationOptionsNone TrustListValidationOptions = 0
	TrustListValidationOptions_trustListValidationOptionsSuppressCertificateExpired TrustListValidationOptions = 1
	TrustListValidationOptions_trustListValidationOptionsSuppressHostNameInvalid TrustListValidationOptions = 2
	TrustListValidationOptions_trustListValidationOptionsSuppressRevocationStatusUnknown TrustListValidationOptions = 4
	TrustListValidationOptions_trustListValidationOptionsSuppressIssuerCertificateExpired TrustListValidationOptions = 8
	TrustListValidationOptions_trustListValidationOptionsSuppressIssuerRevocationStatusUnknown TrustListValidationOptions = 16
	TrustListValidationOptions_trustListValidationOptionsCheckRevocationStatusOnline TrustListValidationOptions = 32
	TrustListValidationOptions_trustListValidationOptionsCheckRevocationStatusOffline TrustListValidationOptions = 64
)

var TrustListValidationOptionsValues []TrustListValidationOptions

func init() {
	_ = errors.New
	TrustListValidationOptionsValues = []TrustListValidationOptions {
		TrustListValidationOptions_trustListValidationOptionsNone,
		TrustListValidationOptions_trustListValidationOptionsSuppressCertificateExpired,
		TrustListValidationOptions_trustListValidationOptionsSuppressHostNameInvalid,
		TrustListValidationOptions_trustListValidationOptionsSuppressRevocationStatusUnknown,
		TrustListValidationOptions_trustListValidationOptionsSuppressIssuerCertificateExpired,
		TrustListValidationOptions_trustListValidationOptionsSuppressIssuerRevocationStatusUnknown,
		TrustListValidationOptions_trustListValidationOptionsCheckRevocationStatusOnline,
		TrustListValidationOptions_trustListValidationOptionsCheckRevocationStatusOffline,
	}
}

func TrustListValidationOptionsByValue(value uint32) (enum TrustListValidationOptions, ok bool) {
	switch value {
		case 0:
			return TrustListValidationOptions_trustListValidationOptionsNone, true
		case 1:
			return TrustListValidationOptions_trustListValidationOptionsSuppressCertificateExpired, true
		case 16:
			return TrustListValidationOptions_trustListValidationOptionsSuppressIssuerRevocationStatusUnknown, true
		case 2:
			return TrustListValidationOptions_trustListValidationOptionsSuppressHostNameInvalid, true
		case 32:
			return TrustListValidationOptions_trustListValidationOptionsCheckRevocationStatusOnline, true
		case 4:
			return TrustListValidationOptions_trustListValidationOptionsSuppressRevocationStatusUnknown, true
		case 64:
			return TrustListValidationOptions_trustListValidationOptionsCheckRevocationStatusOffline, true
		case 8:
			return TrustListValidationOptions_trustListValidationOptionsSuppressIssuerCertificateExpired, true
	}
	return 0, false
}

func TrustListValidationOptionsByName(value string) (enum TrustListValidationOptions, ok bool) {
	switch value {
	case "trustListValidationOptionsNone":
		return TrustListValidationOptions_trustListValidationOptionsNone, true
	case "trustListValidationOptionsSuppressCertificateExpired":
		return TrustListValidationOptions_trustListValidationOptionsSuppressCertificateExpired, true
	case "trustListValidationOptionsSuppressIssuerRevocationStatusUnknown":
		return TrustListValidationOptions_trustListValidationOptionsSuppressIssuerRevocationStatusUnknown, true
	case "trustListValidationOptionsSuppressHostNameInvalid":
		return TrustListValidationOptions_trustListValidationOptionsSuppressHostNameInvalid, true
	case "trustListValidationOptionsCheckRevocationStatusOnline":
		return TrustListValidationOptions_trustListValidationOptionsCheckRevocationStatusOnline, true
	case "trustListValidationOptionsSuppressRevocationStatusUnknown":
		return TrustListValidationOptions_trustListValidationOptionsSuppressRevocationStatusUnknown, true
	case "trustListValidationOptionsCheckRevocationStatusOffline":
		return TrustListValidationOptions_trustListValidationOptionsCheckRevocationStatusOffline, true
	case "trustListValidationOptionsSuppressIssuerCertificateExpired":
		return TrustListValidationOptions_trustListValidationOptionsSuppressIssuerCertificateExpired, true
	}
	return 0, false
}

func TrustListValidationOptionsKnows(value uint32)  bool {
	for _, typeValue := range TrustListValidationOptionsValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastTrustListValidationOptions(structType any) TrustListValidationOptions {
	castFunc := func(typ any) TrustListValidationOptions {
		if sTrustListValidationOptions, ok := typ.(TrustListValidationOptions); ok {
			return sTrustListValidationOptions
		}
		return 0
	}
	return castFunc(structType)
}

func (m TrustListValidationOptions) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m TrustListValidationOptions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TrustListValidationOptionsParse(ctx context.Context, theBytes []byte) (TrustListValidationOptions, error) {
	return TrustListValidationOptionsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TrustListValidationOptionsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TrustListValidationOptions, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("TrustListValidationOptions", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TrustListValidationOptions")
	}
	if enum, ok := TrustListValidationOptionsByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for TrustListValidationOptions")
		return TrustListValidationOptions(val), nil
	} else {
		return enum, nil
	}
}

func (e TrustListValidationOptions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TrustListValidationOptions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("TrustListValidationOptions", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TrustListValidationOptions) PLC4XEnumName() string {
	switch e {
	case TrustListValidationOptions_trustListValidationOptionsNone:
		return "trustListValidationOptionsNone"
	case TrustListValidationOptions_trustListValidationOptionsSuppressCertificateExpired:
		return "trustListValidationOptionsSuppressCertificateExpired"
	case TrustListValidationOptions_trustListValidationOptionsSuppressIssuerRevocationStatusUnknown:
		return "trustListValidationOptionsSuppressIssuerRevocationStatusUnknown"
	case TrustListValidationOptions_trustListValidationOptionsSuppressHostNameInvalid:
		return "trustListValidationOptionsSuppressHostNameInvalid"
	case TrustListValidationOptions_trustListValidationOptionsCheckRevocationStatusOnline:
		return "trustListValidationOptionsCheckRevocationStatusOnline"
	case TrustListValidationOptions_trustListValidationOptionsSuppressRevocationStatusUnknown:
		return "trustListValidationOptionsSuppressRevocationStatusUnknown"
	case TrustListValidationOptions_trustListValidationOptionsCheckRevocationStatusOffline:
		return "trustListValidationOptionsCheckRevocationStatusOffline"
	case TrustListValidationOptions_trustListValidationOptionsSuppressIssuerCertificateExpired:
		return "trustListValidationOptionsSuppressIssuerCertificateExpired"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e TrustListValidationOptions) String() string {
	return e.PLC4XEnumName()
}

