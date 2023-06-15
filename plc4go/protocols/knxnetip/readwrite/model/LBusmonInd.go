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


// LBusmonInd is the corresponding interface of LBusmonInd
type LBusmonInd interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
	// GetAdditionalInformationLength returns AdditionalInformationLength (property field)
	GetAdditionalInformationLength() uint8
	// GetAdditionalInformation returns AdditionalInformation (property field)
	GetAdditionalInformation() []CEMIAdditionalInformation
	// GetDataFrame returns DataFrame (property field)
	GetDataFrame() LDataFrame
	// GetCrc returns Crc (property field)
	GetCrc() *uint8
}

// LBusmonIndExactly can be used when we want exactly this type and not a type which fulfills LBusmonInd.
// This is useful for switch cases.
type LBusmonIndExactly interface {
	LBusmonInd
	isLBusmonInd() bool
}

// _LBusmonInd is the data-structure of this message
type _LBusmonInd struct {
	*_CEMI
        AdditionalInformationLength uint8
        AdditionalInformation []CEMIAdditionalInformation
        DataFrame LDataFrame
        Crc *uint8
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LBusmonInd)  GetMessageCode() uint8 {
return 0x2B}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LBusmonInd) InitializeParent(parent CEMI ) {}

func (m *_LBusmonInd)  GetParent() CEMI {
	return m._CEMI
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LBusmonInd) GetAdditionalInformationLength() uint8 {
	return m.AdditionalInformationLength
}

func (m *_LBusmonInd) GetAdditionalInformation() []CEMIAdditionalInformation {
	return m.AdditionalInformation
}

func (m *_LBusmonInd) GetDataFrame() LDataFrame {
	return m.DataFrame
}

func (m *_LBusmonInd) GetCrc() *uint8 {
	return m.Crc
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewLBusmonInd factory function for _LBusmonInd
func NewLBusmonInd( additionalInformationLength uint8 , additionalInformation []CEMIAdditionalInformation , dataFrame LDataFrame , crc *uint8 , size uint16 ) *_LBusmonInd {
	_result := &_LBusmonInd{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation: additionalInformation,
		DataFrame: dataFrame,
		Crc: crc,
    	_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLBusmonInd(structType any) LBusmonInd {
    if casted, ok := structType.(LBusmonInd); ok {
		return casted
	}
	if casted, ok := structType.(*LBusmonInd); ok {
		return *casted
	}
	return nil
}

func (m *_LBusmonInd) GetTypeName() string {
	return "LBusmonInd"
}

func (m *_LBusmonInd) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (additionalInformationLength)
	lengthInBits += 8;

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.GetLengthInBits(ctx)

	// Optional Field (crc)
	if m.Crc != nil {
		lengthInBits += 8
	}

	return lengthInBits
}


func (m *_LBusmonInd) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LBusmonIndParse(ctx context.Context, theBytes []byte, size uint16) (LBusmonInd, error) {
	return LBusmonIndParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), size)
}

func LBusmonIndParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (LBusmonInd, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("LBusmonInd"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LBusmonInd")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (additionalInformationLength)
_additionalInformationLength, _additionalInformationLengthErr := readBuffer.ReadUint8("additionalInformationLength", 8)
	if _additionalInformationLengthErr != nil {
		return nil, errors.Wrap(_additionalInformationLengthErr, "Error parsing 'additionalInformationLength' field of LBusmonInd")
	}
	additionalInformationLength := _additionalInformationLength

	// Array field (additionalInformation)
	if pullErr := readBuffer.PullContext("additionalInformation", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for additionalInformation")
	}
	// Length array
	var additionalInformation []CEMIAdditionalInformation
	{
		_additionalInformationLength := additionalInformationLength
		_additionalInformationEndPos := positionAware.GetPos() + uint16(_additionalInformationLength)
		for ;positionAware.GetPos() < _additionalInformationEndPos; {
_item, _err := CEMIAdditionalInformationParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'additionalInformation' field of LBusmonInd")
			}
			additionalInformation = append(additionalInformation, _item.(CEMIAdditionalInformation))
		}
	}
	if closeErr := readBuffer.CloseContext("additionalInformation", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for additionalInformation")
	}

	// Simple Field (dataFrame)
	if pullErr := readBuffer.PullContext("dataFrame"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataFrame")
	}
_dataFrame, _dataFrameErr := LDataFrameParseWithBuffer(ctx, readBuffer)
	if _dataFrameErr != nil {
		return nil, errors.Wrap(_dataFrameErr, "Error parsing 'dataFrame' field of LBusmonInd")
	}
	dataFrame := _dataFrame.(LDataFrame)
	if closeErr := readBuffer.CloseContext("dataFrame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataFrame")
	}

	// Optional Field (crc) (Can be skipped, if a given expression evaluates to false)
	var crc *uint8 = nil
	if CastLDataFrame(dataFrame).GetNotAckFrame() {
		_val, _err := readBuffer.ReadUint8("crc", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'crc' field of LBusmonInd")
		}
		crc = &_val
	}

	if closeErr := readBuffer.CloseContext("LBusmonInd"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LBusmonInd")
	}

	// Create a partially initialized instance
	_child := &_LBusmonInd{
		_CEMI: &_CEMI{
			Size: size,
		},
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation: additionalInformation,
		DataFrame: dataFrame,
		Crc: crc,
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_LBusmonInd) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LBusmonInd) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LBusmonInd"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LBusmonInd")
		}

	// Simple Field (additionalInformationLength)
	additionalInformationLength := uint8(m.GetAdditionalInformationLength())
	_additionalInformationLengthErr := writeBuffer.WriteUint8("additionalInformationLength", 8, (additionalInformationLength))
	if _additionalInformationLengthErr != nil {
		return errors.Wrap(_additionalInformationLengthErr, "Error serializing 'additionalInformationLength' field")
	}

	// Array Field (additionalInformation)
	if pushErr := writeBuffer.PushContext("additionalInformation", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for additionalInformation")
	}
	for _curItem, _element := range m.GetAdditionalInformation() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetAdditionalInformation()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'additionalInformation' field")
		}
	}
	if popErr := writeBuffer.PopContext("additionalInformation", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for additionalInformation")
	}

	// Simple Field (dataFrame)
	if pushErr := writeBuffer.PushContext("dataFrame"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for dataFrame")
	}
	_dataFrameErr := writeBuffer.WriteSerializable(ctx, m.GetDataFrame())
	if popErr := writeBuffer.PopContext("dataFrame"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for dataFrame")
	}
	if _dataFrameErr != nil {
		return errors.Wrap(_dataFrameErr, "Error serializing 'dataFrame' field")
	}

	// Optional Field (crc) (Can be skipped, if the value is null)
	var crc *uint8 = nil
	if m.GetCrc() != nil {
		crc = m.GetCrc()
		_crcErr := writeBuffer.WriteUint8("crc", 8, *(crc))
		if _crcErr != nil {
			return errors.Wrap(_crcErr, "Error serializing 'crc' field")
		}
	}

		if popErr := writeBuffer.PopContext("LBusmonInd"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LBusmonInd")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_LBusmonInd) isLBusmonInd() bool {
	return true
}

func (m *_LBusmonInd) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



