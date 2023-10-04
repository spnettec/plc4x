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

// BACnetFaultParameterFaultExtended is the corresponding interface of BACnetFaultParameterFaultExtended
type BACnetFaultParameterFaultExtended interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetExtendedFaultType returns ExtendedFaultType (property field)
	GetExtendedFaultType() BACnetContextTagUnsignedInteger
	// GetParameters returns Parameters (property field)
	GetParameters() BACnetFaultParameterFaultExtendedParameters
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetFaultParameterFaultExtendedExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtended.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedExactly interface {
	BACnetFaultParameterFaultExtended
	isBACnetFaultParameterFaultExtended() bool
}

// _BACnetFaultParameterFaultExtended is the data-structure of this message
type _BACnetFaultParameterFaultExtended struct {
	*_BACnetFaultParameter
	OpeningTag        BACnetOpeningTag
	VendorId          BACnetVendorIdTagged
	ExtendedFaultType BACnetContextTagUnsignedInteger
	Parameters        BACnetFaultParameterFaultExtendedParameters
	ClosingTag        BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultExtended) InitializeParent(parent BACnetFaultParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterFaultExtended) GetParent() BACnetFaultParameter {
	return m._BACnetFaultParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtended) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultExtended) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetFaultParameterFaultExtended) GetExtendedFaultType() BACnetContextTagUnsignedInteger {
	return m.ExtendedFaultType
}

func (m *_BACnetFaultParameterFaultExtended) GetParameters() BACnetFaultParameterFaultExtendedParameters {
	return m.Parameters
}

func (m *_BACnetFaultParameterFaultExtended) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtended factory function for _BACnetFaultParameterFaultExtended
func NewBACnetFaultParameterFaultExtended(openingTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedFaultType BACnetContextTagUnsignedInteger, parameters BACnetFaultParameterFaultExtendedParameters, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultExtended {
	_result := &_BACnetFaultParameterFaultExtended{
		OpeningTag:            openingTag,
		VendorId:              vendorId,
		ExtendedFaultType:     extendedFaultType,
		Parameters:            parameters,
		ClosingTag:            closingTag,
		_BACnetFaultParameter: NewBACnetFaultParameter(peekedTagHeader),
	}
	_result._BACnetFaultParameter._BACnetFaultParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtended(structType any) BACnetFaultParameterFaultExtended {
	if casted, ok := structType.(BACnetFaultParameterFaultExtended); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtended); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtended) GetTypeName() string {
	return "BACnetFaultParameterFaultExtended"
}

func (m *_BACnetFaultParameterFaultExtended) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	// Simple field (extendedFaultType)
	lengthInBits += m.ExtendedFaultType.GetLengthInBits(ctx)

	// Simple field (parameters)
	lengthInBits += m.Parameters.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtended) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultExtendedParse(ctx context.Context, theBytes []byte) (BACnetFaultParameterFaultExtended, error) {
	return BACnetFaultParameterFaultExtendedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetFaultParameterFaultExtendedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultExtended, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetFaultParameterFaultExtended")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (vendorId)
	if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vendorId")
	}
	_vendorId, _vendorIdErr := BACnetVendorIdTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field of BACnetFaultParameterFaultExtended")
	}
	vendorId := _vendorId.(BACnetVendorIdTagged)
	if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vendorId")
	}

	// Simple Field (extendedFaultType)
	if pullErr := readBuffer.PullContext("extendedFaultType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for extendedFaultType")
	}
	_extendedFaultType, _extendedFaultTypeErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _extendedFaultTypeErr != nil {
		return nil, errors.Wrap(_extendedFaultTypeErr, "Error parsing 'extendedFaultType' field of BACnetFaultParameterFaultExtended")
	}
	extendedFaultType := _extendedFaultType.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("extendedFaultType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for extendedFaultType")
	}

	// Simple Field (parameters)
	if pullErr := readBuffer.PullContext("parameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for parameters")
	}
	_parameters, _parametersErr := BACnetFaultParameterFaultExtendedParametersParseWithBuffer(ctx, readBuffer, uint8(uint8(2)))
	if _parametersErr != nil {
		return nil, errors.Wrap(_parametersErr, "Error parsing 'parameters' field of BACnetFaultParameterFaultExtended")
	}
	parameters := _parameters.(BACnetFaultParameterFaultExtendedParameters)
	if closeErr := readBuffer.CloseContext("parameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for parameters")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetFaultParameterFaultExtended")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtended")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultExtended{
		_BACnetFaultParameter: &_BACnetFaultParameter{},
		OpeningTag:            openingTag,
		VendorId:              vendorId,
		ExtendedFaultType:     extendedFaultType,
		Parameters:            parameters,
		ClosingTag:            closingTag,
	}
	_child._BACnetFaultParameter._BACnetFaultParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtended) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtended")
		}

		// Simple Field (openingTag)
		if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for openingTag")
		}
		_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
		if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for openingTag")
		}
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
		}

		// Simple Field (vendorId)
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorId")
		}
		_vendorIdErr := writeBuffer.WriteSerializable(ctx, m.GetVendorId())
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorId")
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Simple Field (extendedFaultType)
		if pushErr := writeBuffer.PushContext("extendedFaultType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for extendedFaultType")
		}
		_extendedFaultTypeErr := writeBuffer.WriteSerializable(ctx, m.GetExtendedFaultType())
		if popErr := writeBuffer.PopContext("extendedFaultType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for extendedFaultType")
		}
		if _extendedFaultTypeErr != nil {
			return errors.Wrap(_extendedFaultTypeErr, "Error serializing 'extendedFaultType' field")
		}

		// Simple Field (parameters)
		if pushErr := writeBuffer.PushContext("parameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for parameters")
		}
		_parametersErr := writeBuffer.WriteSerializable(ctx, m.GetParameters())
		if popErr := writeBuffer.PopContext("parameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for parameters")
		}
		if _parametersErr != nil {
			return errors.Wrap(_parametersErr, "Error serializing 'parameters' field")
		}

		// Simple Field (closingTag)
		if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for closingTag")
		}
		_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
		if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for closingTag")
		}
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtended")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtended) isBACnetFaultParameterFaultExtended() bool {
	return true
}

func (m *_BACnetFaultParameterFaultExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
