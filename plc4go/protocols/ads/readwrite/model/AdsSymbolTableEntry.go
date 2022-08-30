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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

	// Code generated by code-generation. DO NOT EDIT.


// Constant values.
const AdsSymbolTableEntry_NAMETERMINATOR uint8 = 0x00
const AdsSymbolTableEntry_DATATYPENAMETERMINATOR uint8 = 0x00
const AdsSymbolTableEntry_COMMENTTERMINATOR uint8 = 0x00

// AdsSymbolTableEntry is the corresponding interface of AdsSymbolTableEntry
type AdsSymbolTableEntry interface {
	utils.LengthAware
	utils.Serializable
	// GetEntryLength returns EntryLength (property field)
	GetEntryLength() uint32
	// GetGroup returns Group (property field)
	GetGroup() uint32
	// GetOffset returns Offset (property field)
	GetOffset() uint32
	// GetSize returns Size (property field)
	GetSize() uint32
	// GetDataType returns DataType (property field)
	GetDataType() uint32
	// GetFlagMethodDeref returns FlagMethodDeref (property field)
	GetFlagMethodDeref() bool
	// GetFlagItfMethodAccess returns FlagItfMethodAccess (property field)
	GetFlagItfMethodAccess() bool
	// GetFlagReadOnly returns FlagReadOnly (property field)
	GetFlagReadOnly() bool
	// GetFlagTComInterfacePointer returns FlagTComInterfacePointer (property field)
	GetFlagTComInterfacePointer() bool
	// GetFlagTypeGuid returns FlagTypeGuid (property field)
	GetFlagTypeGuid() bool
	// GetFlagReferenceTo returns FlagReferenceTo (property field)
	GetFlagReferenceTo() bool
	// GetFlagBitValue returns FlagBitValue (property field)
	GetFlagBitValue() bool
	// GetFlagPersistent returns FlagPersistent (property field)
	GetFlagPersistent() bool
	// GetFlagExtendedFlags returns FlagExtendedFlags (property field)
	GetFlagExtendedFlags() bool
	// GetFlagInitOnReset returns FlagInitOnReset (property field)
	GetFlagInitOnReset() bool
	// GetFlagStatic returns FlagStatic (property field)
	GetFlagStatic() bool
	// GetFlagAttributes returns FlagAttributes (property field)
	GetFlagAttributes() bool
	// GetFlagContextMask returns FlagContextMask (property field)
	GetFlagContextMask() bool
	// GetName returns Name (property field)
	GetName() string
	// GetDataTypeName returns DataTypeName (property field)
	GetDataTypeName() string
	// GetComment returns Comment (property field)
	GetComment() string
	// GetRest returns Rest (property field)
	GetRest() []byte
}

// AdsSymbolTableEntryExactly can be used when we want exactly this type and not a type which fulfills AdsSymbolTableEntry.
// This is useful for switch cases.
type AdsSymbolTableEntryExactly interface {
	AdsSymbolTableEntry
	isAdsSymbolTableEntry() bool
}

// _AdsSymbolTableEntry is the data-structure of this message
type _AdsSymbolTableEntry struct {
        EntryLength uint32
        Group uint32
        Offset uint32
        Size uint32
        DataType uint32
        FlagMethodDeref bool
        FlagItfMethodAccess bool
        FlagReadOnly bool
        FlagTComInterfacePointer bool
        FlagTypeGuid bool
        FlagReferenceTo bool
        FlagBitValue bool
        FlagPersistent bool
        FlagExtendedFlags bool
        FlagInitOnReset bool
        FlagStatic bool
        FlagAttributes bool
        FlagContextMask bool
        Name string
        DataTypeName string
        Comment string
        Rest []byte
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint16
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsSymbolTableEntry) GetEntryLength() uint32 {
	return m.EntryLength
}

func (m *_AdsSymbolTableEntry) GetGroup() uint32 {
	return m.Group
}

func (m *_AdsSymbolTableEntry) GetOffset() uint32 {
	return m.Offset
}

func (m *_AdsSymbolTableEntry) GetSize() uint32 {
	return m.Size
}

func (m *_AdsSymbolTableEntry) GetDataType() uint32 {
	return m.DataType
}

func (m *_AdsSymbolTableEntry) GetFlagMethodDeref() bool {
	return m.FlagMethodDeref
}

func (m *_AdsSymbolTableEntry) GetFlagItfMethodAccess() bool {
	return m.FlagItfMethodAccess
}

func (m *_AdsSymbolTableEntry) GetFlagReadOnly() bool {
	return m.FlagReadOnly
}

func (m *_AdsSymbolTableEntry) GetFlagTComInterfacePointer() bool {
	return m.FlagTComInterfacePointer
}

func (m *_AdsSymbolTableEntry) GetFlagTypeGuid() bool {
	return m.FlagTypeGuid
}

func (m *_AdsSymbolTableEntry) GetFlagReferenceTo() bool {
	return m.FlagReferenceTo
}

func (m *_AdsSymbolTableEntry) GetFlagBitValue() bool {
	return m.FlagBitValue
}

func (m *_AdsSymbolTableEntry) GetFlagPersistent() bool {
	return m.FlagPersistent
}

func (m *_AdsSymbolTableEntry) GetFlagExtendedFlags() bool {
	return m.FlagExtendedFlags
}

func (m *_AdsSymbolTableEntry) GetFlagInitOnReset() bool {
	return m.FlagInitOnReset
}

func (m *_AdsSymbolTableEntry) GetFlagStatic() bool {
	return m.FlagStatic
}

func (m *_AdsSymbolTableEntry) GetFlagAttributes() bool {
	return m.FlagAttributes
}

func (m *_AdsSymbolTableEntry) GetFlagContextMask() bool {
	return m.FlagContextMask
}

func (m *_AdsSymbolTableEntry) GetName() string {
	return m.Name
}

func (m *_AdsSymbolTableEntry) GetDataTypeName() string {
	return m.DataTypeName
}

func (m *_AdsSymbolTableEntry) GetComment() string {
	return m.Comment
}

func (m *_AdsSymbolTableEntry) GetRest() []byte {
	return m.Rest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AdsSymbolTableEntry) GetNameTerminator() uint8 {
	return AdsSymbolTableEntry_NAMETERMINATOR
}

func (m *_AdsSymbolTableEntry) GetDataTypeNameTerminator() uint8 {
	return AdsSymbolTableEntry_DATATYPENAMETERMINATOR
}

func (m *_AdsSymbolTableEntry) GetCommentTerminator() uint8 {
	return AdsSymbolTableEntry_COMMENTTERMINATOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewAdsSymbolTableEntry factory function for _AdsSymbolTableEntry
func NewAdsSymbolTableEntry( entryLength uint32 , group uint32 , offset uint32 , size uint32 , dataType uint32 , flagMethodDeref bool , flagItfMethodAccess bool , flagReadOnly bool , flagTComInterfacePointer bool , flagTypeGuid bool , flagReferenceTo bool , flagBitValue bool , flagPersistent bool , flagExtendedFlags bool , flagInitOnReset bool , flagStatic bool , flagAttributes bool , flagContextMask bool , name string , dataTypeName string , comment string , rest []byte ) *_AdsSymbolTableEntry {
return &_AdsSymbolTableEntry{ EntryLength: entryLength , Group: group , Offset: offset , Size: size , DataType: dataType , FlagMethodDeref: flagMethodDeref , FlagItfMethodAccess: flagItfMethodAccess , FlagReadOnly: flagReadOnly , FlagTComInterfacePointer: flagTComInterfacePointer , FlagTypeGuid: flagTypeGuid , FlagReferenceTo: flagReferenceTo , FlagBitValue: flagBitValue , FlagPersistent: flagPersistent , FlagExtendedFlags: flagExtendedFlags , FlagInitOnReset: flagInitOnReset , FlagStatic: flagStatic , FlagAttributes: flagAttributes , FlagContextMask: flagContextMask , Name: name , DataTypeName: dataTypeName , Comment: comment , Rest: rest }
}

// Deprecated: use the interface for direct cast
func CastAdsSymbolTableEntry(structType interface{}) AdsSymbolTableEntry {
    if casted, ok := structType.(AdsSymbolTableEntry); ok {
		return casted
	}
	if casted, ok := structType.(*AdsSymbolTableEntry); ok {
		return *casted
	}
	return nil
}

func (m *_AdsSymbolTableEntry) GetTypeName() string {
	return "AdsSymbolTableEntry"
}

func (m *_AdsSymbolTableEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsSymbolTableEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (entryLength)
	lengthInBits += 32;

	// Simple field (group)
	lengthInBits += 32;

	// Simple field (offset)
	lengthInBits += 32;

	// Simple field (size)
	lengthInBits += 32;

	// Simple field (dataType)
	lengthInBits += 32;

	// Simple field (flagMethodDeref)
	lengthInBits += 1;

	// Simple field (flagItfMethodAccess)
	lengthInBits += 1;

	// Simple field (flagReadOnly)
	lengthInBits += 1;

	// Simple field (flagTComInterfacePointer)
	lengthInBits += 1;

	// Simple field (flagTypeGuid)
	lengthInBits += 1;

	// Simple field (flagReferenceTo)
	lengthInBits += 1;

	// Simple field (flagBitValue)
	lengthInBits += 1;

	// Simple field (flagPersistent)
	lengthInBits += 1;

	// Reserved Field (reserved)
	lengthInBits += 3

	// Simple field (flagExtendedFlags)
	lengthInBits += 1;

	// Simple field (flagInitOnReset)
	lengthInBits += 1;

	// Simple field (flagStatic)
	lengthInBits += 1;

	// Simple field (flagAttributes)
	lengthInBits += 1;

	// Simple field (flagContextMask)
	lengthInBits += 1;

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (nameLength)
	lengthInBits += 16

	// Implicit Field (dataTypeNameLength)
	lengthInBits += 16

	// Implicit Field (commentLength)
	lengthInBits += 16

	// Simple field (name)
	lengthInBits += uint16(int32(uint16(len(m.GetName()))) * int32(int32(8)))

	// Const Field (nameTerminator)
	lengthInBits += 8

	// Simple field (dataTypeName)
	lengthInBits += uint16(int32(uint16(len(m.GetDataTypeName()))) * int32(int32(8)))

	// Const Field (dataTypeNameTerminator)
	lengthInBits += 8

	// Simple field (comment)
	lengthInBits += uint16(int32(uint16(len(m.GetComment()))) * int32(int32(8)))

	// Const Field (commentTerminator)
	lengthInBits += 8

	// Array field
	if len(m.Rest) > 0 {
		lengthInBits += 8 * uint16(len(m.Rest))
	}

	return lengthInBits
}


func (m *_AdsSymbolTableEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsSymbolTableEntryParse(readBuffer utils.ReadBuffer) (AdsSymbolTableEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsSymbolTableEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsSymbolTableEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	var startPos = positionAware.GetPos()
	_ = startPos
	var curPos uint16

	// Simple Field (entryLength)
_entryLength, _entryLengthErr := readBuffer.ReadUint32("entryLength", 32)
	if _entryLengthErr != nil {
		return nil, errors.Wrap(_entryLengthErr, "Error parsing 'entryLength' field of AdsSymbolTableEntry")
	}
	entryLength := _entryLength

	// Simple Field (group)
_group, _groupErr := readBuffer.ReadUint32("group", 32)
	if _groupErr != nil {
		return nil, errors.Wrap(_groupErr, "Error parsing 'group' field of AdsSymbolTableEntry")
	}
	group := _group

	// Simple Field (offset)
_offset, _offsetErr := readBuffer.ReadUint32("offset", 32)
	if _offsetErr != nil {
		return nil, errors.Wrap(_offsetErr, "Error parsing 'offset' field of AdsSymbolTableEntry")
	}
	offset := _offset

	// Simple Field (size)
_size, _sizeErr := readBuffer.ReadUint32("size", 32)
	if _sizeErr != nil {
		return nil, errors.Wrap(_sizeErr, "Error parsing 'size' field of AdsSymbolTableEntry")
	}
	size := _size

	// Simple Field (dataType)
_dataType, _dataTypeErr := readBuffer.ReadUint32("dataType", 32)
	if _dataTypeErr != nil {
		return nil, errors.Wrap(_dataTypeErr, "Error parsing 'dataType' field of AdsSymbolTableEntry")
	}
	dataType := _dataType

	// Simple Field (flagMethodDeref)
_flagMethodDeref, _flagMethodDerefErr := readBuffer.ReadBit("flagMethodDeref")
	if _flagMethodDerefErr != nil {
		return nil, errors.Wrap(_flagMethodDerefErr, "Error parsing 'flagMethodDeref' field of AdsSymbolTableEntry")
	}
	flagMethodDeref := _flagMethodDeref

	// Simple Field (flagItfMethodAccess)
_flagItfMethodAccess, _flagItfMethodAccessErr := readBuffer.ReadBit("flagItfMethodAccess")
	if _flagItfMethodAccessErr != nil {
		return nil, errors.Wrap(_flagItfMethodAccessErr, "Error parsing 'flagItfMethodAccess' field of AdsSymbolTableEntry")
	}
	flagItfMethodAccess := _flagItfMethodAccess

	// Simple Field (flagReadOnly)
_flagReadOnly, _flagReadOnlyErr := readBuffer.ReadBit("flagReadOnly")
	if _flagReadOnlyErr != nil {
		return nil, errors.Wrap(_flagReadOnlyErr, "Error parsing 'flagReadOnly' field of AdsSymbolTableEntry")
	}
	flagReadOnly := _flagReadOnly

	// Simple Field (flagTComInterfacePointer)
_flagTComInterfacePointer, _flagTComInterfacePointerErr := readBuffer.ReadBit("flagTComInterfacePointer")
	if _flagTComInterfacePointerErr != nil {
		return nil, errors.Wrap(_flagTComInterfacePointerErr, "Error parsing 'flagTComInterfacePointer' field of AdsSymbolTableEntry")
	}
	flagTComInterfacePointer := _flagTComInterfacePointer

	// Simple Field (flagTypeGuid)
_flagTypeGuid, _flagTypeGuidErr := readBuffer.ReadBit("flagTypeGuid")
	if _flagTypeGuidErr != nil {
		return nil, errors.Wrap(_flagTypeGuidErr, "Error parsing 'flagTypeGuid' field of AdsSymbolTableEntry")
	}
	flagTypeGuid := _flagTypeGuid

	// Simple Field (flagReferenceTo)
_flagReferenceTo, _flagReferenceToErr := readBuffer.ReadBit("flagReferenceTo")
	if _flagReferenceToErr != nil {
		return nil, errors.Wrap(_flagReferenceToErr, "Error parsing 'flagReferenceTo' field of AdsSymbolTableEntry")
	}
	flagReferenceTo := _flagReferenceTo

	// Simple Field (flagBitValue)
_flagBitValue, _flagBitValueErr := readBuffer.ReadBit("flagBitValue")
	if _flagBitValueErr != nil {
		return nil, errors.Wrap(_flagBitValueErr, "Error parsing 'flagBitValue' field of AdsSymbolTableEntry")
	}
	flagBitValue := _flagBitValue

	// Simple Field (flagPersistent)
_flagPersistent, _flagPersistentErr := readBuffer.ReadBit("flagPersistent")
	if _flagPersistentErr != nil {
		return nil, errors.Wrap(_flagPersistentErr, "Error parsing 'flagPersistent' field of AdsSymbolTableEntry")
	}
	flagPersistent := _flagPersistent

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 3)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of AdsSymbolTableEntry")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (flagExtendedFlags)
_flagExtendedFlags, _flagExtendedFlagsErr := readBuffer.ReadBit("flagExtendedFlags")
	if _flagExtendedFlagsErr != nil {
		return nil, errors.Wrap(_flagExtendedFlagsErr, "Error parsing 'flagExtendedFlags' field of AdsSymbolTableEntry")
	}
	flagExtendedFlags := _flagExtendedFlags

	// Simple Field (flagInitOnReset)
_flagInitOnReset, _flagInitOnResetErr := readBuffer.ReadBit("flagInitOnReset")
	if _flagInitOnResetErr != nil {
		return nil, errors.Wrap(_flagInitOnResetErr, "Error parsing 'flagInitOnReset' field of AdsSymbolTableEntry")
	}
	flagInitOnReset := _flagInitOnReset

	// Simple Field (flagStatic)
_flagStatic, _flagStaticErr := readBuffer.ReadBit("flagStatic")
	if _flagStaticErr != nil {
		return nil, errors.Wrap(_flagStaticErr, "Error parsing 'flagStatic' field of AdsSymbolTableEntry")
	}
	flagStatic := _flagStatic

	// Simple Field (flagAttributes)
_flagAttributes, _flagAttributesErr := readBuffer.ReadBit("flagAttributes")
	if _flagAttributesErr != nil {
		return nil, errors.Wrap(_flagAttributesErr, "Error parsing 'flagAttributes' field of AdsSymbolTableEntry")
	}
	flagAttributes := _flagAttributes

	// Simple Field (flagContextMask)
_flagContextMask, _flagContextMaskErr := readBuffer.ReadBit("flagContextMask")
	if _flagContextMaskErr != nil {
		return nil, errors.Wrap(_flagContextMaskErr, "Error parsing 'flagContextMask' field of AdsSymbolTableEntry")
	}
	flagContextMask := _flagContextMask

	var reservedField1 *uint16
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of AdsSymbolTableEntry")
		}
		if reserved != uint16(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	// Implicit Field (nameLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	nameLength, _nameLengthErr := readBuffer.ReadUint16("nameLength", 16)
	_ = nameLength
	if _nameLengthErr != nil {
		return nil, errors.Wrap(_nameLengthErr, "Error parsing 'nameLength' field of AdsSymbolTableEntry")
	}

	// Implicit Field (dataTypeNameLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	dataTypeNameLength, _dataTypeNameLengthErr := readBuffer.ReadUint16("dataTypeNameLength", 16)
	_ = dataTypeNameLength
	if _dataTypeNameLengthErr != nil {
		return nil, errors.Wrap(_dataTypeNameLengthErr, "Error parsing 'dataTypeNameLength' field of AdsSymbolTableEntry")
	}

	// Implicit Field (commentLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	commentLength, _commentLengthErr := readBuffer.ReadUint16("commentLength", 16)
	_ = commentLength
	if _commentLengthErr != nil {
		return nil, errors.Wrap(_commentLengthErr, "Error parsing 'commentLength' field of AdsSymbolTableEntry")
	}

	// Simple Field (name)
_name, _nameErr := readBuffer.ReadString("name", uint32((nameLength) * ((8))))
	if _nameErr != nil {
		return nil, errors.Wrap(_nameErr, "Error parsing 'name' field of AdsSymbolTableEntry")
	}
	name := _name

	// Const Field (nameTerminator)
	nameTerminator, _nameTerminatorErr := readBuffer.ReadUint8("nameTerminator", 8)
	if _nameTerminatorErr != nil {
		return nil, errors.Wrap(_nameTerminatorErr, "Error parsing 'nameTerminator' field of AdsSymbolTableEntry")
	}
	if nameTerminator != AdsSymbolTableEntry_NAMETERMINATOR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AdsSymbolTableEntry_NAMETERMINATOR) + " but got " + fmt.Sprintf("%d", nameTerminator))
	}

	// Simple Field (dataTypeName)
_dataTypeName, _dataTypeNameErr := readBuffer.ReadString("dataTypeName", uint32((dataTypeNameLength) * ((8))))
	if _dataTypeNameErr != nil {
		return nil, errors.Wrap(_dataTypeNameErr, "Error parsing 'dataTypeName' field of AdsSymbolTableEntry")
	}
	dataTypeName := _dataTypeName

	// Const Field (dataTypeNameTerminator)
	dataTypeNameTerminator, _dataTypeNameTerminatorErr := readBuffer.ReadUint8("dataTypeNameTerminator", 8)
	if _dataTypeNameTerminatorErr != nil {
		return nil, errors.Wrap(_dataTypeNameTerminatorErr, "Error parsing 'dataTypeNameTerminator' field of AdsSymbolTableEntry")
	}
	if dataTypeNameTerminator != AdsSymbolTableEntry_DATATYPENAMETERMINATOR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AdsSymbolTableEntry_DATATYPENAMETERMINATOR) + " but got " + fmt.Sprintf("%d", dataTypeNameTerminator))
	}

	// Simple Field (comment)
_comment, _commentErr := readBuffer.ReadString("comment", uint32((commentLength) * ((8))))
	if _commentErr != nil {
		return nil, errors.Wrap(_commentErr, "Error parsing 'comment' field of AdsSymbolTableEntry")
	}
	comment := _comment

	// Const Field (commentTerminator)
	commentTerminator, _commentTerminatorErr := readBuffer.ReadUint8("commentTerminator", 8)
	if _commentTerminatorErr != nil {
		return nil, errors.Wrap(_commentTerminatorErr, "Error parsing 'commentTerminator' field of AdsSymbolTableEntry")
	}
	if commentTerminator != AdsSymbolTableEntry_COMMENTTERMINATOR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AdsSymbolTableEntry_COMMENTTERMINATOR) + " but got " + fmt.Sprintf("%d", commentTerminator))
	}
	// Byte Array field (rest)
	numberOfBytesrest := int(uint16(entryLength) - uint16(curPos))
	rest, _readArrayErr := readBuffer.ReadByteArray("rest", numberOfBytesrest)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'rest' field of AdsSymbolTableEntry")
	}

	if closeErr := readBuffer.CloseContext("AdsSymbolTableEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsSymbolTableEntry")
	}

	// Create the instance
	return &_AdsSymbolTableEntry{
			EntryLength: entryLength,
			Group: group,
			Offset: offset,
			Size: size,
			DataType: dataType,
			FlagMethodDeref: flagMethodDeref,
			FlagItfMethodAccess: flagItfMethodAccess,
			FlagReadOnly: flagReadOnly,
			FlagTComInterfacePointer: flagTComInterfacePointer,
			FlagTypeGuid: flagTypeGuid,
			FlagReferenceTo: flagReferenceTo,
			FlagBitValue: flagBitValue,
			FlagPersistent: flagPersistent,
			FlagExtendedFlags: flagExtendedFlags,
			FlagInitOnReset: flagInitOnReset,
			FlagStatic: flagStatic,
			FlagAttributes: flagAttributes,
			FlagContextMask: flagContextMask,
			Name: name,
			DataTypeName: dataTypeName,
			Comment: comment,
			Rest: rest,
			reservedField0: reservedField0,
			reservedField1: reservedField1,
		}, nil
}

func (m *_AdsSymbolTableEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("AdsSymbolTableEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AdsSymbolTableEntry")
	}

	// Simple Field (entryLength)
	entryLength := uint32(m.GetEntryLength())
	_entryLengthErr := writeBuffer.WriteUint32("entryLength", 32, (entryLength))
	if _entryLengthErr != nil {
		return errors.Wrap(_entryLengthErr, "Error serializing 'entryLength' field")
	}

	// Simple Field (group)
	group := uint32(m.GetGroup())
	_groupErr := writeBuffer.WriteUint32("group", 32, (group))
	if _groupErr != nil {
		return errors.Wrap(_groupErr, "Error serializing 'group' field")
	}

	// Simple Field (offset)
	offset := uint32(m.GetOffset())
	_offsetErr := writeBuffer.WriteUint32("offset", 32, (offset))
	if _offsetErr != nil {
		return errors.Wrap(_offsetErr, "Error serializing 'offset' field")
	}

	// Simple Field (size)
	size := uint32(m.GetSize())
	_sizeErr := writeBuffer.WriteUint32("size", 32, (size))
	if _sizeErr != nil {
		return errors.Wrap(_sizeErr, "Error serializing 'size' field")
	}

	// Simple Field (dataType)
	dataType := uint32(m.GetDataType())
	_dataTypeErr := writeBuffer.WriteUint32("dataType", 32, (dataType))
	if _dataTypeErr != nil {
		return errors.Wrap(_dataTypeErr, "Error serializing 'dataType' field")
	}

	// Simple Field (flagMethodDeref)
	flagMethodDeref := bool(m.GetFlagMethodDeref())
	_flagMethodDerefErr := writeBuffer.WriteBit("flagMethodDeref", (flagMethodDeref))
	if _flagMethodDerefErr != nil {
		return errors.Wrap(_flagMethodDerefErr, "Error serializing 'flagMethodDeref' field")
	}

	// Simple Field (flagItfMethodAccess)
	flagItfMethodAccess := bool(m.GetFlagItfMethodAccess())
	_flagItfMethodAccessErr := writeBuffer.WriteBit("flagItfMethodAccess", (flagItfMethodAccess))
	if _flagItfMethodAccessErr != nil {
		return errors.Wrap(_flagItfMethodAccessErr, "Error serializing 'flagItfMethodAccess' field")
	}

	// Simple Field (flagReadOnly)
	flagReadOnly := bool(m.GetFlagReadOnly())
	_flagReadOnlyErr := writeBuffer.WriteBit("flagReadOnly", (flagReadOnly))
	if _flagReadOnlyErr != nil {
		return errors.Wrap(_flagReadOnlyErr, "Error serializing 'flagReadOnly' field")
	}

	// Simple Field (flagTComInterfacePointer)
	flagTComInterfacePointer := bool(m.GetFlagTComInterfacePointer())
	_flagTComInterfacePointerErr := writeBuffer.WriteBit("flagTComInterfacePointer", (flagTComInterfacePointer))
	if _flagTComInterfacePointerErr != nil {
		return errors.Wrap(_flagTComInterfacePointerErr, "Error serializing 'flagTComInterfacePointer' field")
	}

	// Simple Field (flagTypeGuid)
	flagTypeGuid := bool(m.GetFlagTypeGuid())
	_flagTypeGuidErr := writeBuffer.WriteBit("flagTypeGuid", (flagTypeGuid))
	if _flagTypeGuidErr != nil {
		return errors.Wrap(_flagTypeGuidErr, "Error serializing 'flagTypeGuid' field")
	}

	// Simple Field (flagReferenceTo)
	flagReferenceTo := bool(m.GetFlagReferenceTo())
	_flagReferenceToErr := writeBuffer.WriteBit("flagReferenceTo", (flagReferenceTo))
	if _flagReferenceToErr != nil {
		return errors.Wrap(_flagReferenceToErr, "Error serializing 'flagReferenceTo' field")
	}

	// Simple Field (flagBitValue)
	flagBitValue := bool(m.GetFlagBitValue())
	_flagBitValueErr := writeBuffer.WriteBit("flagBitValue", (flagBitValue))
	if _flagBitValueErr != nil {
		return errors.Wrap(_flagBitValueErr, "Error serializing 'flagBitValue' field")
	}

	// Simple Field (flagPersistent)
	flagPersistent := bool(m.GetFlagPersistent())
	_flagPersistentErr := writeBuffer.WriteBit("flagPersistent", (flagPersistent))
	if _flagPersistentErr != nil {
		return errors.Wrap(_flagPersistentErr, "Error serializing 'flagPersistent' field")
	}

	// Reserved Field (reserved)
	{
		var reserved uint8 = uint8(0x00)
		if m.reservedField0 != nil {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteUint8("reserved", 3, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (flagExtendedFlags)
	flagExtendedFlags := bool(m.GetFlagExtendedFlags())
	_flagExtendedFlagsErr := writeBuffer.WriteBit("flagExtendedFlags", (flagExtendedFlags))
	if _flagExtendedFlagsErr != nil {
		return errors.Wrap(_flagExtendedFlagsErr, "Error serializing 'flagExtendedFlags' field")
	}

	// Simple Field (flagInitOnReset)
	flagInitOnReset := bool(m.GetFlagInitOnReset())
	_flagInitOnResetErr := writeBuffer.WriteBit("flagInitOnReset", (flagInitOnReset))
	if _flagInitOnResetErr != nil {
		return errors.Wrap(_flagInitOnResetErr, "Error serializing 'flagInitOnReset' field")
	}

	// Simple Field (flagStatic)
	flagStatic := bool(m.GetFlagStatic())
	_flagStaticErr := writeBuffer.WriteBit("flagStatic", (flagStatic))
	if _flagStaticErr != nil {
		return errors.Wrap(_flagStaticErr, "Error serializing 'flagStatic' field")
	}

	// Simple Field (flagAttributes)
	flagAttributes := bool(m.GetFlagAttributes())
	_flagAttributesErr := writeBuffer.WriteBit("flagAttributes", (flagAttributes))
	if _flagAttributesErr != nil {
		return errors.Wrap(_flagAttributesErr, "Error serializing 'flagAttributes' field")
	}

	// Simple Field (flagContextMask)
	flagContextMask := bool(m.GetFlagContextMask())
	_flagContextMaskErr := writeBuffer.WriteBit("flagContextMask", (flagContextMask))
	if _flagContextMaskErr != nil {
		return errors.Wrap(_flagContextMaskErr, "Error serializing 'flagContextMask' field")
	}

	// Reserved Field (reserved)
	{
		var reserved uint16 = uint16(0x0000)
		if m.reservedField1 != nil {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField1
		}
		_err := writeBuffer.WriteUint16("reserved", 16, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Implicit Field (nameLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	nameLength := uint16(uint16(len(m.GetName())))
	_nameLengthErr := writeBuffer.WriteUint16("nameLength", 16, (nameLength))
	if _nameLengthErr != nil {
		return errors.Wrap(_nameLengthErr, "Error serializing 'nameLength' field")
	}

	// Implicit Field (dataTypeNameLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataTypeNameLength := uint16(uint16(len(m.GetDataTypeName())))
	_dataTypeNameLengthErr := writeBuffer.WriteUint16("dataTypeNameLength", 16, (dataTypeNameLength))
	if _dataTypeNameLengthErr != nil {
		return errors.Wrap(_dataTypeNameLengthErr, "Error serializing 'dataTypeNameLength' field")
	}

	// Implicit Field (commentLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	commentLength := uint16(uint16(len(m.GetComment())))
	_commentLengthErr := writeBuffer.WriteUint16("commentLength", 16, (commentLength))
	if _commentLengthErr != nil {
		return errors.Wrap(_commentLengthErr, "Error serializing 'commentLength' field")
	}

	// Simple Field (name)
	name := string(m.GetName())
	_nameErr := writeBuffer.WriteString("name", uint32((uint16(len(m.GetName()))) * ((8))), "UTF-8", (name))
	if _nameErr != nil {
		return errors.Wrap(_nameErr, "Error serializing 'name' field")
	}

	// Const Field (nameTerminator)
	_nameTerminatorErr := writeBuffer.WriteUint8("nameTerminator", 8, 0x00)
	if _nameTerminatorErr != nil {
		return errors.Wrap(_nameTerminatorErr, "Error serializing 'nameTerminator' field")
	}

	// Simple Field (dataTypeName)
	dataTypeName := string(m.GetDataTypeName())
	_dataTypeNameErr := writeBuffer.WriteString("dataTypeName", uint32((uint16(len(m.GetDataTypeName()))) * ((8))), "UTF-8", (dataTypeName))
	if _dataTypeNameErr != nil {
		return errors.Wrap(_dataTypeNameErr, "Error serializing 'dataTypeName' field")
	}

	// Const Field (dataTypeNameTerminator)
	_dataTypeNameTerminatorErr := writeBuffer.WriteUint8("dataTypeNameTerminator", 8, 0x00)
	if _dataTypeNameTerminatorErr != nil {
		return errors.Wrap(_dataTypeNameTerminatorErr, "Error serializing 'dataTypeNameTerminator' field")
	}

	// Simple Field (comment)
	comment := string(m.GetComment())
	_commentErr := writeBuffer.WriteString("comment", uint32((uint16(len(m.GetComment()))) * ((8))), "UTF-8", (comment))
	if _commentErr != nil {
		return errors.Wrap(_commentErr, "Error serializing 'comment' field")
	}

	// Const Field (commentTerminator)
	_commentTerminatorErr := writeBuffer.WriteUint8("commentTerminator", 8, 0x00)
	if _commentTerminatorErr != nil {
		return errors.Wrap(_commentTerminatorErr, "Error serializing 'commentTerminator' field")
	}

	// Array Field (rest)
	// Byte Array field (rest)
	if err := writeBuffer.WriteByteArray("rest", m.GetRest()); err != nil {
		return errors.Wrap(err, "Error serializing 'rest' field")
	}

	if popErr := writeBuffer.PopContext("AdsSymbolTableEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AdsSymbolTableEntry")
	}
	return nil
}


func (m *_AdsSymbolTableEntry) isAdsSymbolTableEntry() bool {
	return true
}

func (m *_AdsSymbolTableEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



