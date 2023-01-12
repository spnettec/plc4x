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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// NLMUpdateKeyUpdateKeyEntry is the corresponding interface of NLMUpdateKeyUpdateKeyEntry
type NLMUpdateKeyUpdateKeyEntry interface {
	utils.LengthAware
	utils.Serializable
	// GetKeyIdentifier returns KeyIdentifier (property field)
	GetKeyIdentifier() uint16
	// GetKeySize returns KeySize (property field)
	GetKeySize() uint8
	// GetKey returns Key (property field)
	GetKey() []byte
}

// NLMUpdateKeyUpdateKeyEntryExactly can be used when we want exactly this type and not a type which fulfills NLMUpdateKeyUpdateKeyEntry.
// This is useful for switch cases.
type NLMUpdateKeyUpdateKeyEntryExactly interface {
	NLMUpdateKeyUpdateKeyEntry
	isNLMUpdateKeyUpdateKeyEntry() bool
}

// _NLMUpdateKeyUpdateKeyEntry is the data-structure of this message
type _NLMUpdateKeyUpdateKeyEntry struct {
        KeyIdentifier uint16
        KeySize uint8
        Key []byte
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMUpdateKeyUpdateKeyEntry) GetKeyIdentifier() uint16 {
	return m.KeyIdentifier
}

func (m *_NLMUpdateKeyUpdateKeyEntry) GetKeySize() uint8 {
	return m.KeySize
}

func (m *_NLMUpdateKeyUpdateKeyEntry) GetKey() []byte {
	return m.Key
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewNLMUpdateKeyUpdateKeyEntry factory function for _NLMUpdateKeyUpdateKeyEntry
func NewNLMUpdateKeyUpdateKeyEntry( keyIdentifier uint16 , keySize uint8 , key []byte ) *_NLMUpdateKeyUpdateKeyEntry {
return &_NLMUpdateKeyUpdateKeyEntry{ KeyIdentifier: keyIdentifier , KeySize: keySize , Key: key }
}

// Deprecated: use the interface for direct cast
func CastNLMUpdateKeyUpdateKeyEntry(structType interface{}) NLMUpdateKeyUpdateKeyEntry {
    if casted, ok := structType.(NLMUpdateKeyUpdateKeyEntry); ok {
		return casted
	}
	if casted, ok := structType.(*NLMUpdateKeyUpdateKeyEntry); ok {
		return *casted
	}
	return nil
}

func (m *_NLMUpdateKeyUpdateKeyEntry) GetTypeName() string {
	return "NLMUpdateKeyUpdateKeyEntry"
}

func (m *_NLMUpdateKeyUpdateKeyEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_NLMUpdateKeyUpdateKeyEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (keyIdentifier)
	lengthInBits += 16;

	// Simple field (keySize)
	lengthInBits += 8;

	// Array field
	if len(m.Key) > 0 {
		lengthInBits += 8 * uint16(len(m.Key))
	}

	return lengthInBits
}


func (m *_NLMUpdateKeyUpdateKeyEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMUpdateKeyUpdateKeyEntryParse(theBytes []byte) (NLMUpdateKeyUpdateKeyEntry, error) {
	return NLMUpdateKeyUpdateKeyEntryParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func NLMUpdateKeyUpdateKeyEntryParseWithBuffer(readBuffer utils.ReadBuffer) (NLMUpdateKeyUpdateKeyEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMUpdateKeyUpdateKeyEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMUpdateKeyUpdateKeyEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (keyIdentifier)
_keyIdentifier, _keyIdentifierErr := readBuffer.ReadUint16("keyIdentifier", 16)
	if _keyIdentifierErr != nil {
		return nil, errors.Wrap(_keyIdentifierErr, "Error parsing 'keyIdentifier' field of NLMUpdateKeyUpdateKeyEntry")
	}
	keyIdentifier := _keyIdentifier

	// Simple Field (keySize)
_keySize, _keySizeErr := readBuffer.ReadUint8("keySize", 8)
	if _keySizeErr != nil {
		return nil, errors.Wrap(_keySizeErr, "Error parsing 'keySize' field of NLMUpdateKeyUpdateKeyEntry")
	}
	keySize := _keySize
	// Byte Array field (key)
	numberOfByteskey := int(keySize)
	key, _readArrayErr := readBuffer.ReadByteArray("key", numberOfByteskey)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'key' field of NLMUpdateKeyUpdateKeyEntry")
	}

	if closeErr := readBuffer.CloseContext("NLMUpdateKeyUpdateKeyEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMUpdateKeyUpdateKeyEntry")
	}

	// Create the instance
	return &_NLMUpdateKeyUpdateKeyEntry{
			KeyIdentifier: keyIdentifier,
			KeySize: keySize,
			Key: key,
		}, nil
}

func (m *_NLMUpdateKeyUpdateKeyEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMUpdateKeyUpdateKeyEntry) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("NLMUpdateKeyUpdateKeyEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NLMUpdateKeyUpdateKeyEntry")
	}

	// Simple Field (keyIdentifier)
	keyIdentifier := uint16(m.GetKeyIdentifier())
	_keyIdentifierErr := writeBuffer.WriteUint16("keyIdentifier", 16, (keyIdentifier))
	if _keyIdentifierErr != nil {
		return errors.Wrap(_keyIdentifierErr, "Error serializing 'keyIdentifier' field")
	}

	// Simple Field (keySize)
	keySize := uint8(m.GetKeySize())
	_keySizeErr := writeBuffer.WriteUint8("keySize", 8, (keySize))
	if _keySizeErr != nil {
		return errors.Wrap(_keySizeErr, "Error serializing 'keySize' field")
	}

	// Array Field (key)
	// Byte Array field (key)
	if err := writeBuffer.WriteByteArray("key", m.GetKey()); err != nil {
		return errors.Wrap(err, "Error serializing 'key' field")
	}

	if popErr := writeBuffer.PopContext("NLMUpdateKeyUpdateKeyEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NLMUpdateKeyUpdateKeyEntry")
	}
	return nil
}


func (m *_NLMUpdateKeyUpdateKeyEntry) isNLMUpdateKeyUpdateKeyEntry() bool {
	return true
}

func (m *_NLMUpdateKeyUpdateKeyEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



