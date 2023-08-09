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


// NLMUpdateKeyUpdate is the corresponding interface of NLMUpdateKeyUpdate
type NLMUpdateKeyUpdate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetControlFlags returns ControlFlags (property field)
	GetControlFlags() NLMUpdateKeyUpdateControlFlags
	// GetSet1KeyRevision returns Set1KeyRevision (property field)
	GetSet1KeyRevision() *byte
	// GetSet1ActivationTime returns Set1ActivationTime (property field)
	GetSet1ActivationTime() *uint32
	// GetSet1ExpirationTime returns Set1ExpirationTime (property field)
	GetSet1ExpirationTime() *uint32
	// GetSet1KeyCount returns Set1KeyCount (property field)
	GetSet1KeyCount() *uint8
	// GetSet1Keys returns Set1Keys (property field)
	GetSet1Keys() []NLMUpdateKeyUpdateKeyEntry
	// GetSet2KeyRevision returns Set2KeyRevision (property field)
	GetSet2KeyRevision() *byte
	// GetSet2ActivationTime returns Set2ActivationTime (property field)
	GetSet2ActivationTime() *uint32
	// GetSet2ExpirationTime returns Set2ExpirationTime (property field)
	GetSet2ExpirationTime() *uint32
	// GetSet2KeyCount returns Set2KeyCount (property field)
	GetSet2KeyCount() *uint8
	// GetSet2Keys returns Set2Keys (property field)
	GetSet2Keys() []NLMUpdateKeyUpdateKeyEntry
}

// NLMUpdateKeyUpdateExactly can be used when we want exactly this type and not a type which fulfills NLMUpdateKeyUpdate.
// This is useful for switch cases.
type NLMUpdateKeyUpdateExactly interface {
	NLMUpdateKeyUpdate
	isNLMUpdateKeyUpdate() bool
}

// _NLMUpdateKeyUpdate is the data-structure of this message
type _NLMUpdateKeyUpdate struct {
	*_NLM
        ControlFlags NLMUpdateKeyUpdateControlFlags
        Set1KeyRevision *byte
        Set1ActivationTime *uint32
        Set1ExpirationTime *uint32
        Set1KeyCount *uint8
        Set1Keys []NLMUpdateKeyUpdateKeyEntry
        Set2KeyRevision *byte
        Set2ActivationTime *uint32
        Set2ExpirationTime *uint32
        Set2KeyCount *uint8
        Set2Keys []NLMUpdateKeyUpdateKeyEntry
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMUpdateKeyUpdate)  GetMessageType() uint8 {
return 0x0E}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMUpdateKeyUpdate) InitializeParent(parent NLM ) {}

func (m *_NLMUpdateKeyUpdate)  GetParent() NLM {
	return m._NLM
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMUpdateKeyUpdate) GetControlFlags() NLMUpdateKeyUpdateControlFlags {
	return m.ControlFlags
}

func (m *_NLMUpdateKeyUpdate) GetSet1KeyRevision() *byte {
	return m.Set1KeyRevision
}

func (m *_NLMUpdateKeyUpdate) GetSet1ActivationTime() *uint32 {
	return m.Set1ActivationTime
}

func (m *_NLMUpdateKeyUpdate) GetSet1ExpirationTime() *uint32 {
	return m.Set1ExpirationTime
}

func (m *_NLMUpdateKeyUpdate) GetSet1KeyCount() *uint8 {
	return m.Set1KeyCount
}

func (m *_NLMUpdateKeyUpdate) GetSet1Keys() []NLMUpdateKeyUpdateKeyEntry {
	return m.Set1Keys
}

func (m *_NLMUpdateKeyUpdate) GetSet2KeyRevision() *byte {
	return m.Set2KeyRevision
}

func (m *_NLMUpdateKeyUpdate) GetSet2ActivationTime() *uint32 {
	return m.Set2ActivationTime
}

func (m *_NLMUpdateKeyUpdate) GetSet2ExpirationTime() *uint32 {
	return m.Set2ExpirationTime
}

func (m *_NLMUpdateKeyUpdate) GetSet2KeyCount() *uint8 {
	return m.Set2KeyCount
}

func (m *_NLMUpdateKeyUpdate) GetSet2Keys() []NLMUpdateKeyUpdateKeyEntry {
	return m.Set2Keys
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewNLMUpdateKeyUpdate factory function for _NLMUpdateKeyUpdate
func NewNLMUpdateKeyUpdate( controlFlags NLMUpdateKeyUpdateControlFlags , set1KeyRevision *byte , set1ActivationTime *uint32 , set1ExpirationTime *uint32 , set1KeyCount *uint8 , set1Keys []NLMUpdateKeyUpdateKeyEntry , set2KeyRevision *byte , set2ActivationTime *uint32 , set2ExpirationTime *uint32 , set2KeyCount *uint8 , set2Keys []NLMUpdateKeyUpdateKeyEntry , apduLength uint16 ) *_NLMUpdateKeyUpdate {
	_result := &_NLMUpdateKeyUpdate{
		ControlFlags: controlFlags,
		Set1KeyRevision: set1KeyRevision,
		Set1ActivationTime: set1ActivationTime,
		Set1ExpirationTime: set1ExpirationTime,
		Set1KeyCount: set1KeyCount,
		Set1Keys: set1Keys,
		Set2KeyRevision: set2KeyRevision,
		Set2ActivationTime: set2ActivationTime,
		Set2ExpirationTime: set2ExpirationTime,
		Set2KeyCount: set2KeyCount,
		Set2Keys: set2Keys,
    	_NLM: NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMUpdateKeyUpdate(structType any) NLMUpdateKeyUpdate {
    if casted, ok := structType.(NLMUpdateKeyUpdate); ok {
		return casted
	}
	if casted, ok := structType.(*NLMUpdateKeyUpdate); ok {
		return *casted
	}
	return nil
}

func (m *_NLMUpdateKeyUpdate) GetTypeName() string {
	return "NLMUpdateKeyUpdate"
}

func (m *_NLMUpdateKeyUpdate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (controlFlags)
	lengthInBits += m.ControlFlags.GetLengthInBits(ctx)

	// Optional Field (set1KeyRevision)
	if m.Set1KeyRevision != nil {
		lengthInBits += 8
	}

	// Optional Field (set1ActivationTime)
	if m.Set1ActivationTime != nil {
		lengthInBits += 32
	}

	// Optional Field (set1ExpirationTime)
	if m.Set1ExpirationTime != nil {
		lengthInBits += 32
	}

	// Optional Field (set1KeyCount)
	if m.Set1KeyCount != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.Set1Keys) > 0 {
		for _curItem, element := range m.Set1Keys {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Set1Keys), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	// Optional Field (set2KeyRevision)
	if m.Set2KeyRevision != nil {
		lengthInBits += 8
	}

	// Optional Field (set2ActivationTime)
	if m.Set2ActivationTime != nil {
		lengthInBits += 32
	}

	// Optional Field (set2ExpirationTime)
	if m.Set2ExpirationTime != nil {
		lengthInBits += 32
	}

	// Optional Field (set2KeyCount)
	if m.Set2KeyCount != nil {
		lengthInBits += 8
	}

	// Array field
	if len(m.Set2Keys) > 0 {
		for _curItem, element := range m.Set2Keys {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Set2Keys), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}


func (m *_NLMUpdateKeyUpdate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMUpdateKeyUpdateParse(ctx context.Context, theBytes []byte, apduLength uint16) (NLMUpdateKeyUpdate, error) {
	return NLMUpdateKeyUpdateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMUpdateKeyUpdateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLMUpdateKeyUpdate, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NLMUpdateKeyUpdate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMUpdateKeyUpdate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (controlFlags)
	if pullErr := readBuffer.PullContext("controlFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for controlFlags")
	}
_controlFlags, _controlFlagsErr := NLMUpdateKeyUpdateControlFlagsParseWithBuffer(ctx, readBuffer)
	if _controlFlagsErr != nil {
		return nil, errors.Wrap(_controlFlagsErr, "Error parsing 'controlFlags' field of NLMUpdateKeyUpdate")
	}
	controlFlags := _controlFlags.(NLMUpdateKeyUpdateControlFlags)
	if closeErr := readBuffer.CloseContext("controlFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for controlFlags")
	}

	// Optional Field (set1KeyRevision) (Can be skipped, if a given expression evaluates to false)
	var set1KeyRevision *byte = nil
	if controlFlags.GetSet1KeyRevisionActivationTimeExpirationTimePresent() {
		_val, _err := readBuffer.ReadByte("set1KeyRevision")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'set1KeyRevision' field of NLMUpdateKeyUpdate")
		}
		set1KeyRevision = &_val
	}

	// Optional Field (set1ActivationTime) (Can be skipped, if a given expression evaluates to false)
	var set1ActivationTime *uint32 = nil
	if controlFlags.GetSet1KeyRevisionActivationTimeExpirationTimePresent() {
		_val, _err := readBuffer.ReadUint32("set1ActivationTime", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'set1ActivationTime' field of NLMUpdateKeyUpdate")
		}
		set1ActivationTime = &_val
	}

	// Optional Field (set1ExpirationTime) (Can be skipped, if a given expression evaluates to false)
	var set1ExpirationTime *uint32 = nil
	if controlFlags.GetSet1KeyCountKeyParametersPresent() {
		_val, _err := readBuffer.ReadUint32("set1ExpirationTime", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'set1ExpirationTime' field of NLMUpdateKeyUpdate")
		}
		set1ExpirationTime = &_val
	}

	// Optional Field (set1KeyCount) (Can be skipped, if a given expression evaluates to false)
	var set1KeyCount *uint8 = nil
	if controlFlags.GetSet1KeyCountKeyParametersPresent() {
		_val, _err := readBuffer.ReadUint8("set1KeyCount", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'set1KeyCount' field of NLMUpdateKeyUpdate")
		}
		set1KeyCount = &_val
	}

	// Array field (set1Keys)
	if pullErr := readBuffer.PullContext("set1Keys", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for set1Keys")
	}
	// Count array
	set1Keys := make([]NLMUpdateKeyUpdateKeyEntry, utils.Max(utils.InlineIf(bool(((set1KeyCount)) != (nil)), func() any {return uint16((*set1KeyCount))}, func() any {return uint16(uint16(0))}).(uint16), 0))
	// This happens when the size is set conditional to 0
	if len(set1Keys) == 0 {
		set1Keys = nil
	}
	{
		_numItems := uint16(utils.Max(utils.InlineIf(bool(((set1KeyCount)) != (nil)), func() any {return uint16((*set1KeyCount))}, func() any {return uint16(uint16(0))}).(uint16), 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := NLMUpdateKeyUpdateKeyEntryParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'set1Keys' field of NLMUpdateKeyUpdate")
			}
			set1Keys[_curItem] = _item.(NLMUpdateKeyUpdateKeyEntry)
		}
	}
	if closeErr := readBuffer.CloseContext("set1Keys", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for set1Keys")
	}

	// Optional Field (set2KeyRevision) (Can be skipped, if a given expression evaluates to false)
	var set2KeyRevision *byte = nil
	if controlFlags.GetSet1KeyRevisionActivationTimeExpirationTimePresent() {
		_val, _err := readBuffer.ReadByte("set2KeyRevision")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'set2KeyRevision' field of NLMUpdateKeyUpdate")
		}
		set2KeyRevision = &_val
	}

	// Optional Field (set2ActivationTime) (Can be skipped, if a given expression evaluates to false)
	var set2ActivationTime *uint32 = nil
	if controlFlags.GetSet1KeyRevisionActivationTimeExpirationTimePresent() {
		_val, _err := readBuffer.ReadUint32("set2ActivationTime", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'set2ActivationTime' field of NLMUpdateKeyUpdate")
		}
		set2ActivationTime = &_val
	}

	// Optional Field (set2ExpirationTime) (Can be skipped, if a given expression evaluates to false)
	var set2ExpirationTime *uint32 = nil
	if controlFlags.GetSet1KeyCountKeyParametersPresent() {
		_val, _err := readBuffer.ReadUint32("set2ExpirationTime", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'set2ExpirationTime' field of NLMUpdateKeyUpdate")
		}
		set2ExpirationTime = &_val
	}

	// Optional Field (set2KeyCount) (Can be skipped, if a given expression evaluates to false)
	var set2KeyCount *uint8 = nil
	if controlFlags.GetSet1KeyCountKeyParametersPresent() {
		_val, _err := readBuffer.ReadUint8("set2KeyCount", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'set2KeyCount' field of NLMUpdateKeyUpdate")
		}
		set2KeyCount = &_val
	}

	// Array field (set2Keys)
	if pullErr := readBuffer.PullContext("set2Keys", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for set2Keys")
	}
	// Count array
	set2Keys := make([]NLMUpdateKeyUpdateKeyEntry, utils.Max(utils.InlineIf(bool(((set1KeyCount)) != (nil)), func() any {return uint16((*set1KeyCount))}, func() any {return uint16(uint16(0))}).(uint16), 0))
	// This happens when the size is set conditional to 0
	if len(set2Keys) == 0 {
		set2Keys = nil
	}
	{
		_numItems := uint16(utils.Max(utils.InlineIf(bool(((set1KeyCount)) != (nil)), func() any {return uint16((*set1KeyCount))}, func() any {return uint16(uint16(0))}).(uint16), 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := NLMUpdateKeyUpdateKeyEntryParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'set2Keys' field of NLMUpdateKeyUpdate")
			}
			set2Keys[_curItem] = _item.(NLMUpdateKeyUpdateKeyEntry)
		}
	}
	if closeErr := readBuffer.CloseContext("set2Keys", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for set2Keys")
	}

	if closeErr := readBuffer.CloseContext("NLMUpdateKeyUpdate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMUpdateKeyUpdate")
	}

	// Create a partially initialized instance
	_child := &_NLMUpdateKeyUpdate{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		ControlFlags: controlFlags,
		Set1KeyRevision: set1KeyRevision,
		Set1ActivationTime: set1ActivationTime,
		Set1ExpirationTime: set1ExpirationTime,
		Set1KeyCount: set1KeyCount,
		Set1Keys: set1Keys,
		Set2KeyRevision: set2KeyRevision,
		Set2ActivationTime: set2ActivationTime,
		Set2ExpirationTime: set2ExpirationTime,
		Set2KeyCount: set2KeyCount,
		Set2Keys: set2Keys,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMUpdateKeyUpdate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMUpdateKeyUpdate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMUpdateKeyUpdate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMUpdateKeyUpdate")
		}

	// Simple Field (controlFlags)
	if pushErr := writeBuffer.PushContext("controlFlags"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for controlFlags")
	}
	_controlFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetControlFlags())
	if popErr := writeBuffer.PopContext("controlFlags"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for controlFlags")
	}
	if _controlFlagsErr != nil {
		return errors.Wrap(_controlFlagsErr, "Error serializing 'controlFlags' field")
	}

	// Optional Field (set1KeyRevision) (Can be skipped, if the value is null)
	var set1KeyRevision *byte = nil
	if m.GetSet1KeyRevision() != nil {
		set1KeyRevision = m.GetSet1KeyRevision()
		_set1KeyRevisionErr := writeBuffer.WriteByte("set1KeyRevision", *(set1KeyRevision))
		if _set1KeyRevisionErr != nil {
			return errors.Wrap(_set1KeyRevisionErr, "Error serializing 'set1KeyRevision' field")
		}
	}

	// Optional Field (set1ActivationTime) (Can be skipped, if the value is null)
	var set1ActivationTime *uint32 = nil
	if m.GetSet1ActivationTime() != nil {
		set1ActivationTime = m.GetSet1ActivationTime()
		_set1ActivationTimeErr := writeBuffer.WriteUint32("set1ActivationTime", 32, *(set1ActivationTime))
		if _set1ActivationTimeErr != nil {
			return errors.Wrap(_set1ActivationTimeErr, "Error serializing 'set1ActivationTime' field")
		}
	}

	// Optional Field (set1ExpirationTime) (Can be skipped, if the value is null)
	var set1ExpirationTime *uint32 = nil
	if m.GetSet1ExpirationTime() != nil {
		set1ExpirationTime = m.GetSet1ExpirationTime()
		_set1ExpirationTimeErr := writeBuffer.WriteUint32("set1ExpirationTime", 32, *(set1ExpirationTime))
		if _set1ExpirationTimeErr != nil {
			return errors.Wrap(_set1ExpirationTimeErr, "Error serializing 'set1ExpirationTime' field")
		}
	}

	// Optional Field (set1KeyCount) (Can be skipped, if the value is null)
	var set1KeyCount *uint8 = nil
	if m.GetSet1KeyCount() != nil {
		set1KeyCount = m.GetSet1KeyCount()
		_set1KeyCountErr := writeBuffer.WriteUint8("set1KeyCount", 8, *(set1KeyCount))
		if _set1KeyCountErr != nil {
			return errors.Wrap(_set1KeyCountErr, "Error serializing 'set1KeyCount' field")
		}
	}

	// Array Field (set1Keys)
	if pushErr := writeBuffer.PushContext("set1Keys", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for set1Keys")
	}
	for _curItem, _element := range m.GetSet1Keys() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetSet1Keys()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'set1Keys' field")
		}
	}
	if popErr := writeBuffer.PopContext("set1Keys", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for set1Keys")
	}

	// Optional Field (set2KeyRevision) (Can be skipped, if the value is null)
	var set2KeyRevision *byte = nil
	if m.GetSet2KeyRevision() != nil {
		set2KeyRevision = m.GetSet2KeyRevision()
		_set2KeyRevisionErr := writeBuffer.WriteByte("set2KeyRevision", *(set2KeyRevision))
		if _set2KeyRevisionErr != nil {
			return errors.Wrap(_set2KeyRevisionErr, "Error serializing 'set2KeyRevision' field")
		}
	}

	// Optional Field (set2ActivationTime) (Can be skipped, if the value is null)
	var set2ActivationTime *uint32 = nil
	if m.GetSet2ActivationTime() != nil {
		set2ActivationTime = m.GetSet2ActivationTime()
		_set2ActivationTimeErr := writeBuffer.WriteUint32("set2ActivationTime", 32, *(set2ActivationTime))
		if _set2ActivationTimeErr != nil {
			return errors.Wrap(_set2ActivationTimeErr, "Error serializing 'set2ActivationTime' field")
		}
	}

	// Optional Field (set2ExpirationTime) (Can be skipped, if the value is null)
	var set2ExpirationTime *uint32 = nil
	if m.GetSet2ExpirationTime() != nil {
		set2ExpirationTime = m.GetSet2ExpirationTime()
		_set2ExpirationTimeErr := writeBuffer.WriteUint32("set2ExpirationTime", 32, *(set2ExpirationTime))
		if _set2ExpirationTimeErr != nil {
			return errors.Wrap(_set2ExpirationTimeErr, "Error serializing 'set2ExpirationTime' field")
		}
	}

	// Optional Field (set2KeyCount) (Can be skipped, if the value is null)
	var set2KeyCount *uint8 = nil
	if m.GetSet2KeyCount() != nil {
		set2KeyCount = m.GetSet2KeyCount()
		_set2KeyCountErr := writeBuffer.WriteUint8("set2KeyCount", 8, *(set2KeyCount))
		if _set2KeyCountErr != nil {
			return errors.Wrap(_set2KeyCountErr, "Error serializing 'set2KeyCount' field")
		}
	}

	// Array Field (set2Keys)
	if pushErr := writeBuffer.PushContext("set2Keys", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for set2Keys")
	}
	for _curItem, _element := range m.GetSet2Keys() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetSet2Keys()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'set2Keys' field")
		}
	}
	if popErr := writeBuffer.PopContext("set2Keys", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for set2Keys")
	}

		if popErr := writeBuffer.PopContext("NLMUpdateKeyUpdate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMUpdateKeyUpdate")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_NLMUpdateKeyUpdate) isNLMUpdateKeyUpdate() bool {
	return true
}

func (m *_NLMUpdateKeyUpdate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



