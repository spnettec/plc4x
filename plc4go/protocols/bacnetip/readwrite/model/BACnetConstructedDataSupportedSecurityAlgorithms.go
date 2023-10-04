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

// BACnetConstructedDataSupportedSecurityAlgorithms is the corresponding interface of BACnetConstructedDataSupportedSecurityAlgorithms
type BACnetConstructedDataSupportedSecurityAlgorithms interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSupportedSecurityAlgorithms returns SupportedSecurityAlgorithms (property field)
	GetSupportedSecurityAlgorithms() []BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataSupportedSecurityAlgorithmsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSupportedSecurityAlgorithms.
// This is useful for switch cases.
type BACnetConstructedDataSupportedSecurityAlgorithmsExactly interface {
	BACnetConstructedDataSupportedSecurityAlgorithms
	isBACnetConstructedDataSupportedSecurityAlgorithms() bool
}

// _BACnetConstructedDataSupportedSecurityAlgorithms is the data-structure of this message
type _BACnetConstructedDataSupportedSecurityAlgorithms struct {
	*_BACnetConstructedData
	SupportedSecurityAlgorithms []BACnetApplicationTagUnsignedInteger
}

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

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) GetSupportedSecurityAlgorithms() []BACnetApplicationTagUnsignedInteger {
	return m.SupportedSecurityAlgorithms
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSupportedSecurityAlgorithms factory function for _BACnetConstructedDataSupportedSecurityAlgorithms
func NewBACnetConstructedDataSupportedSecurityAlgorithms(supportedSecurityAlgorithms []BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSupportedSecurityAlgorithms {
	_result := &_BACnetConstructedDataSupportedSecurityAlgorithms{
		SupportedSecurityAlgorithms: supportedSecurityAlgorithms,
		_BACnetConstructedData:      NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
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
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

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

func BACnetConstructedDataSupportedSecurityAlgorithmsParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSupportedSecurityAlgorithms, error) {
	return BACnetConstructedDataSupportedSecurityAlgorithmsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataSupportedSecurityAlgorithmsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSupportedSecurityAlgorithms, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSupportedSecurityAlgorithms"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSupportedSecurityAlgorithms")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (supportedSecurityAlgorithms)
	if pullErr := readBuffer.PullContext("supportedSecurityAlgorithms", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for supportedSecurityAlgorithms")
	}
	// Terminated array
	var supportedSecurityAlgorithms []BACnetApplicationTagUnsignedInteger
	{
		for !bool(IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber)) {
			_item, _err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'supportedSecurityAlgorithms' field of BACnetConstructedDataSupportedSecurityAlgorithms")
			}
			supportedSecurityAlgorithms = append(supportedSecurityAlgorithms, _item.(BACnetApplicationTagUnsignedInteger))
		}
	}
	if closeErr := readBuffer.CloseContext("supportedSecurityAlgorithms", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for supportedSecurityAlgorithms")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSupportedSecurityAlgorithms"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSupportedSecurityAlgorithms")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSupportedSecurityAlgorithms{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		SupportedSecurityAlgorithms: supportedSecurityAlgorithms,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
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

		// Array Field (supportedSecurityAlgorithms)
		if pushErr := writeBuffer.PushContext("supportedSecurityAlgorithms", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for supportedSecurityAlgorithms")
		}
		for _curItem, _element := range m.GetSupportedSecurityAlgorithms() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetSupportedSecurityAlgorithms()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'supportedSecurityAlgorithms' field")
			}
		}
		if popErr := writeBuffer.PopContext("supportedSecurityAlgorithms", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for supportedSecurityAlgorithms")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSupportedSecurityAlgorithms"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSupportedSecurityAlgorithms")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSupportedSecurityAlgorithms) isBACnetConstructedDataSupportedSecurityAlgorithms() bool {
	return true
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
