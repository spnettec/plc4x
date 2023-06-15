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


// BACnetConfirmedServiceRequestVTOpen is the corresponding interface of BACnetConfirmedServiceRequestVTOpen
type BACnetConfirmedServiceRequestVTOpen interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetVtClass returns VtClass (property field)
	GetVtClass() BACnetVTClassTagged
	// GetLocalVtSessionIdentifier returns LocalVtSessionIdentifier (property field)
	GetLocalVtSessionIdentifier() BACnetApplicationTagUnsignedInteger
}

// BACnetConfirmedServiceRequestVTOpenExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestVTOpen.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestVTOpenExactly interface {
	BACnetConfirmedServiceRequestVTOpen
	isBACnetConfirmedServiceRequestVTOpen() bool
}

// _BACnetConfirmedServiceRequestVTOpen is the data-structure of this message
type _BACnetConfirmedServiceRequestVTOpen struct {
	*_BACnetConfirmedServiceRequest
        VtClass BACnetVTClassTagged
        LocalVtSessionIdentifier BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTOpen)  GetServiceChoice() BACnetConfirmedServiceChoice {
return BACnetConfirmedServiceChoice_VT_OPEN}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestVTOpen) InitializeParent(parent BACnetConfirmedServiceRequest ) {}

func (m *_BACnetConfirmedServiceRequestVTOpen)  GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTOpen) GetVtClass() BACnetVTClassTagged {
	return m.VtClass
}

func (m *_BACnetConfirmedServiceRequestVTOpen) GetLocalVtSessionIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.LocalVtSessionIdentifier
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConfirmedServiceRequestVTOpen factory function for _BACnetConfirmedServiceRequestVTOpen
func NewBACnetConfirmedServiceRequestVTOpen( vtClass BACnetVTClassTagged , localVtSessionIdentifier BACnetApplicationTagUnsignedInteger , serviceRequestLength uint32 ) *_BACnetConfirmedServiceRequestVTOpen {
	_result := &_BACnetConfirmedServiceRequestVTOpen{
		VtClass: vtClass,
		LocalVtSessionIdentifier: localVtSessionIdentifier,
    	_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestVTOpen(structType any) BACnetConfirmedServiceRequestVTOpen {
    if casted, ok := structType.(BACnetConfirmedServiceRequestVTOpen); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestVTOpen); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestVTOpen) GetTypeName() string {
	return "BACnetConfirmedServiceRequestVTOpen"
}

func (m *_BACnetConfirmedServiceRequestVTOpen) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (vtClass)
	lengthInBits += m.VtClass.GetLengthInBits(ctx)

	// Simple field (localVtSessionIdentifier)
	lengthInBits += m.LocalVtSessionIdentifier.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetConfirmedServiceRequestVTOpen) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestVTOpenParse(ctx context.Context, theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestVTOpen, error) {
	return BACnetConfirmedServiceRequestVTOpenParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestVTOpenParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestVTOpen, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestVTOpen"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestVTOpen")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (vtClass)
	if pullErr := readBuffer.PullContext("vtClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vtClass")
	}
_vtClass, _vtClassErr := BACnetVTClassTaggedParseWithBuffer(ctx, readBuffer , uint8( uint8(0) ) , TagClass( TagClass_APPLICATION_TAGS ) )
	if _vtClassErr != nil {
		return nil, errors.Wrap(_vtClassErr, "Error parsing 'vtClass' field of BACnetConfirmedServiceRequestVTOpen")
	}
	vtClass := _vtClass.(BACnetVTClassTagged)
	if closeErr := readBuffer.CloseContext("vtClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vtClass")
	}

	// Simple Field (localVtSessionIdentifier)
	if pullErr := readBuffer.PullContext("localVtSessionIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for localVtSessionIdentifier")
	}
_localVtSessionIdentifier, _localVtSessionIdentifierErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _localVtSessionIdentifierErr != nil {
		return nil, errors.Wrap(_localVtSessionIdentifierErr, "Error parsing 'localVtSessionIdentifier' field of BACnetConfirmedServiceRequestVTOpen")
	}
	localVtSessionIdentifier := _localVtSessionIdentifier.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("localVtSessionIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for localVtSessionIdentifier")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestVTOpen"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestVTOpen")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestVTOpen{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		VtClass: vtClass,
		LocalVtSessionIdentifier: localVtSessionIdentifier,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestVTOpen) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestVTOpen) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestVTOpen"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestVTOpen")
		}

	// Simple Field (vtClass)
	if pushErr := writeBuffer.PushContext("vtClass"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for vtClass")
	}
	_vtClassErr := writeBuffer.WriteSerializable(ctx, m.GetVtClass())
	if popErr := writeBuffer.PopContext("vtClass"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for vtClass")
	}
	if _vtClassErr != nil {
		return errors.Wrap(_vtClassErr, "Error serializing 'vtClass' field")
	}

	// Simple Field (localVtSessionIdentifier)
	if pushErr := writeBuffer.PushContext("localVtSessionIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for localVtSessionIdentifier")
	}
	_localVtSessionIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetLocalVtSessionIdentifier())
	if popErr := writeBuffer.PopContext("localVtSessionIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for localVtSessionIdentifier")
	}
	if _localVtSessionIdentifierErr != nil {
		return errors.Wrap(_localVtSessionIdentifierErr, "Error serializing 'localVtSessionIdentifier' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestVTOpen"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestVTOpen")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConfirmedServiceRequestVTOpen) isBACnetConfirmedServiceRequestVTOpen() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestVTOpen) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



