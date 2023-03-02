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
	"encoding/binary"
	"fmt"

	"github.com/apache/plc4x/plc4go/pkg/api/model"
	"github.com/apache/plc4x/plc4go/pkg/api/values"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

type DefaultPlcBrowseItem struct {
	Tag          model.PlcTag
	Name         string
	DataTypeName string
	Readable     bool
	Writable     bool
	Subscribable bool
	Children     map[string]model.PlcBrowseItem
	Options      map[string]values.PlcValue
}

func (d *DefaultPlcBrowseItem) GetTag() model.PlcTag {
	return d.Tag
}

func (d *DefaultPlcBrowseItem) GetName() string {
	return d.Name
}

func (d *DefaultPlcBrowseItem) GetDataTypeName() string {
	return d.DataTypeName
}

func (d *DefaultPlcBrowseItem) IsReadable() bool {
	return d.Readable
}

func (d *DefaultPlcBrowseItem) IsWritable() bool {
	return d.Writable
}

func (d *DefaultPlcBrowseItem) IsSubscribable() bool {
	return d.Subscribable
}

func (d *DefaultPlcBrowseItem) GetChildren() map[string]model.PlcBrowseItem {
	return d.Children
}

func (d *DefaultPlcBrowseItem) GetOptions() map[string]values.PlcValue {
	return d.Options
}

func (d *DefaultPlcBrowseItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *DefaultPlcBrowseItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("PlcBrowseItem"); err != nil {
		return err
	}

	if d.Tag != nil {
		if serializableField, ok := d.Tag.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("tag"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("tag"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.Tag)
			if err := writeBuffer.WriteString("tag", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteString("name", uint32(len(d.Name)*8), "UTF-8", d.Name); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("dataTypeName", uint32(len(d.DataTypeName)*8), "UTF-8", d.DataTypeName); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("readable", d.Readable); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("writable", d.Writable); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("subscribable", d.Subscribable); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("children", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for name, elem := range d.Children {

		var elem interface{} = elem
		if serializable, ok := elem.(utils.Serializable); ok {
			if err := writeBuffer.PushContext(name); err != nil {
				return err
			}
			if err := serializable.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext(name); err != nil {
				return err
			}
		} else {
			elemAsString := fmt.Sprintf("%v", elem)
			if err := writeBuffer.WriteString(name, uint32(len(elemAsString)*8), "UTF-8", elemAsString); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("children", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("options", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for name, elem := range d.Options {

		var elem interface{} = elem
		if serializable, ok := elem.(utils.Serializable); ok {
			if err := writeBuffer.PushContext(name); err != nil {
				return err
			}
			if err := serializable.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext(name); err != nil {
				return err
			}
		} else {
			elemAsString := fmt.Sprintf("%v", elem)
			if err := writeBuffer.WriteString(name, uint32(len(elemAsString)*8), "UTF-8", elemAsString); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("options", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("PlcBrowseItem"); err != nil {
		return err
	}
	return nil
}

func (d *DefaultPlcBrowseItem) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
