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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataSupportedSecurityAlgorithms is the corresponding interface of BACnetConstructedDataSupportedSecurityAlgorithms
type BACnetConstructedDataSupportedSecurityAlgorithms interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSupportedSecurityAlgorithms returns SupportedSecurityAlgorithms (property field)
	GetSupportedSecurityAlgorithms() []BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataSupportedSecurityAlgorithms is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSupportedSecurityAlgorithms()
}

// _BACnetConstructedDataSupportedSecurityAlgorithms is the data-structure of this message
type _BACnetConstructedDataSupportedSecurityAlgorithms struct {
	BACnetConstructedDataContract
	SupportedSecurityAlgorithms []BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataSupportedSecurityAlgorithms = (*_BACnetConstructedDataSupportedSecurityAlgorithms)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSupportedSecurityAlgorithms)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUPPORTED_SECURITY_ALGORITHMS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetSupportedSecurityAlgorithms() []BACnetApplicationTagUnsignedInteger {
	return m.SupportedSecurityAlgorithms
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSupportedSecurityAlgorithms factory function for _BACnetConstructedDataSupportedSecurityAlgorithms
func NewBACnetConstructedDataSupportedSecurityAlgorithms(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, supportedSecurityAlgorithms []BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSupportedSecurityAlgorithms {
	_result := &_BACnetConstructedDataSupportedSecurityAlgorithms{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		SupportedSecurityAlgorithms:   supportedSecurityAlgorithms,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSupportedSecurityAlgorithms(structType any) BACnetConstructedDataSupportedSecurityAlgorithms {
	if casted, ok := structType.(BACnetConstructedDataSupportedSecurityAlgorithms); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSupportedSecurityAlgorithms); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetTypeName() string {
	return "BACnetConstructedDataSupportedSecurityAlgorithms"
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.SupportedSecurityAlgorithms) > 0 {
		for _, element := range m.SupportedSecurityAlgorithms {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSupportedSecurityAlgorithms BACnetConstructedDataSupportedSecurityAlgorithms, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSupportedSecurityAlgorithms"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSupportedSecurityAlgorithms")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	supportedSecurityAlgorithms, err := ReadTerminatedArrayField[BACnetApplicationTagUnsignedInteger](ctx, "supportedSecurityAlgorithms", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'supportedSecurityAlgorithms' field"))
	}
	m.SupportedSecurityAlgorithms = supportedSecurityAlgorithms

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSupportedSecurityAlgorithms"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSupportedSecurityAlgorithms")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSupportedSecurityAlgorithms"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSupportedSecurityAlgorithms")
		}

		if err := WriteComplexTypeArrayField(ctx, "supportedSecurityAlgorithms", m.GetSupportedSecurityAlgorithms(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'supportedSecurityAlgorithms' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSupportedSecurityAlgorithms"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSupportedSecurityAlgorithms")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) IsBACnetConstructedDataSupportedSecurityAlgorithms() {
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
