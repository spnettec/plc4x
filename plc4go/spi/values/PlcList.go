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

package values

import (
	"github.com/apache/plc4x/plc4go/pkg/api/values"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

type PlcList struct {
	Values []values.PlcValue
	PlcValueAdapter
}

func NewPlcList(values []values.PlcValue) values.PlcValue {
	return PlcList{
		Values: values,
	}
}

func singleOrAdapter[R any](plcList PlcList, f func(values.PlcValue) R) R {
	if len(plcList.Values) == 1 {
		return f(plcList.Values[0])
	} else {
		return f(plcList.PlcValueAdapter)
	}
}

////
// Simple Types

func (m PlcList) IsSimple() bool   { return singleOrAdapter(m, values.PlcValue.IsSimple) }
func (m PlcList) IsNullable() bool { return singleOrAdapter(m, values.PlcValue.IsNullable) }
func (m PlcList) IsNull() bool     { return singleOrAdapter(m, values.PlcValue.IsNull) }

//
///

////
// Boolean

func (m PlcList) IsBool() bool          { return singleOrAdapter(m, values.PlcValue.IsBool) }
func (m PlcList) GetBoolLength() uint32 { return singleOrAdapter(m, values.PlcValue.GetBoolLength) }
func (m PlcList) GetBool() bool         { return singleOrAdapter(m, values.PlcValue.GetBool) }
func (m PlcList) GetBoolAt(index uint32) bool {
	return m.PlcValueAdapter.GetBoolAt(index)
}
func (m PlcList) GetBoolArray() []bool { return singleOrAdapter(m, values.PlcValue.GetBoolArray) }

//
///

////
// Byte

func (m PlcList) IsByte() bool  { return singleOrAdapter(m, values.PlcValue.IsByte) }
func (m PlcList) GetByte() byte { return singleOrAdapter(m, values.PlcValue.GetByte) }

//
///

////
// Integer

func (m PlcList) IsUint8() bool     { return singleOrAdapter(m, values.PlcValue.IsUint8) }
func (m PlcList) GetUint8() uint8   { return singleOrAdapter(m, values.PlcValue.GetUint8) }
func (m PlcList) IsUint16() bool    { return singleOrAdapter(m, values.PlcValue.IsUint16) }
func (m PlcList) GetUint16() uint16 { return singleOrAdapter(m, values.PlcValue.GetUint16) }
func (m PlcList) IsUint32() bool    { return singleOrAdapter(m, values.PlcValue.IsUint32) }
func (m PlcList) GetUint32() uint32 { return singleOrAdapter(m, values.PlcValue.GetUint32) }
func (m PlcList) IsUint64() bool    { return singleOrAdapter(m, values.PlcValue.IsUint64) }
func (m PlcList) GetUint64() uint64 { return singleOrAdapter(m, values.PlcValue.GetUint64) }
func (m PlcList) IsInt8() bool      { return singleOrAdapter(m, values.PlcValue.IsInt8) }
func (m PlcList) GetInt8() int8     { return singleOrAdapter(m, values.PlcValue.GetInt8) }
func (m PlcList) IsInt16() bool     { return singleOrAdapter(m, values.PlcValue.IsInt16) }
func (m PlcList) GetInt16() int16   { return singleOrAdapter(m, values.PlcValue.GetInt16) }
func (m PlcList) IsInt32() bool     { return singleOrAdapter(m, values.PlcValue.IsInt32) }
func (m PlcList) GetInt32() int32   { return singleOrAdapter(m, values.PlcValue.GetInt32) }
func (m PlcList) IsInt64() bool     { return singleOrAdapter(m, values.PlcValue.IsInt64) }
func (m PlcList) GetInt64() int64   { return singleOrAdapter(m, values.PlcValue.GetInt64) }

//
///

////
// Floating Point

func (m PlcList) IsFloat32() bool     { return singleOrAdapter(m, values.PlcValue.IsFloat32) }
func (m PlcList) GetFloat32() float32 { return singleOrAdapter(m, values.PlcValue.GetFloat32) }
func (m PlcList) IsFloat64() bool     { return singleOrAdapter(m, values.PlcValue.IsFloat64) }
func (m PlcList) GetFloat64() float64 { return singleOrAdapter(m, values.PlcValue.GetFloat64) }

//
///

////
// String

func (m PlcList) IsString() bool { return singleOrAdapter(m, values.PlcValue.IsString) }
func (m PlcList) GetString() string {
	stringValues := make([]string, len(m.Values))
	for i, v := range m.Values {
		stringValues[i] = v.GetString()
	}
	return strings.Join(stringValues, ", ")
}

//
///

////
// Time

func (m PlcList) IsTime() bool           { return singleOrAdapter(m, values.PlcValue.IsTime) }
func (m PlcList) GetTime() time.Time     { return singleOrAdapter(m, values.PlcValue.GetTime) }
func (m PlcList) IsDate() bool           { return singleOrAdapter(m, values.PlcValue.IsDate) }
func (m PlcList) GetDate() time.Time     { return singleOrAdapter(m, values.PlcValue.GetDate) }
func (m PlcList) IsDateTime() bool       { return singleOrAdapter(m, values.PlcValue.IsDateTime) }
func (m PlcList) GetDateTime() time.Time { return singleOrAdapter(m, values.PlcValue.GetDateTime) }

//
///

////
// Raw Access

func (m PlcList) GetRaw() []byte {
	buf := utils.NewWriteBufferByteBased()
	if err := m.Serialize(buf); err != nil {
		log.Error().Err(err).Msg("Error getting raw")
		return nil
	}
	return buf.GetBytes()
}

//
///

////
// List Methods

func (m PlcList) IsList() bool { return true }
func (m PlcList) GetLength() uint32 {
	return uint32(len(m.Values))
}
func (m PlcList) GetIndex(i uint32) values.PlcValue {
	return m.Values[i]
}
func (m PlcList) GetList() []values.PlcValue {
	return m.Values
}

//
///

////
// Struct Methods

func (m PlcList) IsStruct() bool    { return singleOrAdapter(m, values.PlcValue.IsStruct) }
func (m PlcList) GetKeys() []string { return singleOrAdapter(m, values.PlcValue.GetKeys) }
func (m PlcList) HasKey(key string) bool {
	if len(m.Values) == 1 {
		return m.Values[0].HasKey(key)
	} else {
		return m.PlcValueAdapter.HasKey(key)
	}

}
func (m PlcList) GetValue(key string) values.PlcValue {
	if len(m.Values) == 1 {
		return m.Values[0].GetValue(key)
	} else {
		return m.PlcValueAdapter.GetValue(key)
	}
}
func (m PlcList) GetStruct() map[string]values.PlcValue {
	return singleOrAdapter(m, values.PlcValue.GetStruct)
}

//
///

func (m PlcList) Serialize(writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("PlcList"); err != nil {
		return err
	}
	for _, listItem := range m.GetList() {
		if listItemSerializable, ok := listItem.(utils.Serializable); ok {
			if err := listItemSerializable.Serialize(writeBuffer); err != nil {
				return err
			}
		} else {
			return errors.New("Error serializing. List item doesn't implement Serializable")
		}
	}
	if err := writeBuffer.PopContext("PlcList"); err != nil {
		return err
	}
	return nil
}
