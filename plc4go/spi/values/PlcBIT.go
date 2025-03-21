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
	"context"
	"encoding/binary"
	"fmt"

	apiValues "github.com/apache/plc4x/plc4go/pkg/api/values"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

type PlcBIT struct {
	PlcSimpleValueAdapter
	value bool
}

func NewPlcBIT(value bool) PlcBIT {
	return PlcBIT{
		value: value,
	}
}

func (m PlcBIT) IsRaw() bool {
	return true
}

func (m PlcBIT) GetRaw() []byte {
	if m.value {
		return []byte{0x01}
	}
	return []byte{0x00}
}

func (m PlcBIT) IsBool() bool {
	return true
}

func (m PlcBIT) GetBoolLength() uint32 {
	return 1
}

func (m PlcBIT) GetBool() bool {
	return m.value
}

func (m PlcBIT) GetBoolAt(index uint32) bool {
	if index == 0 {
		return m.value
	}
	return false
}

func (m PlcBIT) GetBoolArray() []bool {
	return []bool{m.value}
}

func (m PlcBIT) IsString() bool {
	return true
}

func (m PlcBIT) GetString() string {
	if m.GetBool() {
		return "true"
	} else {
		return "false"
	}
}

func (m PlcBIT) GetPlcValueType() apiValues.PlcValueType {
	return apiValues.BOOL
}

func (m PlcBIT) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m PlcBIT) SerializeWithWriteBuffer(_ context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteBit("PlcBOOL", m.value)
}

func (m PlcBIT) String() string {
	return fmt.Sprintf("%s(%dbit):%v", m.GetPlcValueType(), 1, m.value)
}
