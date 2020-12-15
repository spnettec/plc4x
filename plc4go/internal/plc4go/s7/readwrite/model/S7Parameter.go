//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "encoding/xml"
    "errors"
    "io"
    "github.com/apache/plc4x/plc4go/internal/plc4go/utils"
    "reflect"
    "strings"
)

// The data-structure of this message
type S7Parameter struct {
    Child IS7ParameterChild
    IS7Parameter
    IS7ParameterParent
}

// The corresponding interface
type IS7Parameter interface {
    MessageType() uint8
    ParameterType() uint8
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

type IS7ParameterParent interface {
    SerializeParent(io utils.WriteBuffer, child IS7Parameter, serializeChildFunction func() error) error
    GetTypeName() string
}

type IS7ParameterChild interface {
    Serialize(io utils.WriteBuffer) error
    InitializeParent(parent *S7Parameter)
    GetTypeName() string
    IS7Parameter
}

func NewS7Parameter() *S7Parameter {
    return &S7Parameter{}
}

func CastS7Parameter(structType interface{}) *S7Parameter {
    castFunc := func(typ interface{}) *S7Parameter {
        if casted, ok := typ.(S7Parameter); ok {
            return &casted
        }
        if casted, ok := typ.(*S7Parameter); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *S7Parameter) GetTypeName() string {
    return "S7Parameter"
}

func (m *S7Parameter) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Discriminator Field (parameterType)
    lengthInBits += 8

    // Length of sub-type elements will be added by sub-type...
    lengthInBits += m.Child.LengthInBits()

    return lengthInBits
}

func (m *S7Parameter) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func S7ParameterParse(io *utils.ReadBuffer, messageType uint8) (*S7Parameter, error) {

    // Discriminator Field (parameterType) (Used as input to a switch field)
    parameterType, _parameterTypeErr := io.ReadUint8(8)
    if _parameterTypeErr != nil {
        return nil, errors.New("Error parsing 'parameterType' field " + _parameterTypeErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var _parent *S7Parameter
    var typeSwitchError error
    switch {
    case parameterType == 0xF0:
        _parent, typeSwitchError = S7ParameterSetupCommunicationParse(io)
    case parameterType == 0x04 && messageType == 0x01:
        _parent, typeSwitchError = S7ParameterReadVarRequestParse(io)
    case parameterType == 0x04 && messageType == 0x03:
        _parent, typeSwitchError = S7ParameterReadVarResponseParse(io)
    case parameterType == 0x05 && messageType == 0x01:
        _parent, typeSwitchError = S7ParameterWriteVarRequestParse(io)
    case parameterType == 0x05 && messageType == 0x03:
        _parent, typeSwitchError = S7ParameterWriteVarResponseParse(io)
    case parameterType == 0x00 && messageType == 0x07:
        _parent, typeSwitchError = S7ParameterUserDataParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Finish initializing
    _parent.Child.InitializeParent(_parent)
    return _parent, nil
}

func (m *S7Parameter) Serialize(io utils.WriteBuffer) error {
    return m.Child.Serialize(io)
}

func (m *S7Parameter) SerializeParent(io utils.WriteBuffer, child IS7Parameter, serializeChildFunction func() error) error {

    // Discriminator Field (parameterType) (Used as input to a switch field)
    parameterType := uint8(child.ParameterType())
    _parameterTypeErr := io.WriteUint8(8, (parameterType))
    if _parameterTypeErr != nil {
        return errors.New("Error serializing 'parameterType' field " + _parameterTypeErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := serializeChildFunction()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *S7Parameter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    for {
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
                default:
                    switch start.Attr[0].Value {
                        case "org.apache.plc4x.java.s7.readwrite.S7ParameterSetupCommunication":
                            var dt *S7ParameterSetupCommunication
                            if m.Child != nil {
                                dt = m.Child.(*S7ParameterSetupCommunication)
                            }
                            if err := d.DecodeElement(&dt, &tok); err != nil {
                                return err
                            }
                            if m.Child == nil {
                                dt.Parent = m
                                m.Child = dt
                            }
                        case "org.apache.plc4x.java.s7.readwrite.S7ParameterReadVarRequest":
                            var dt *S7ParameterReadVarRequest
                            if m.Child != nil {
                                dt = m.Child.(*S7ParameterReadVarRequest)
                            }
                            if err := d.DecodeElement(&dt, &tok); err != nil {
                                return err
                            }
                            if m.Child == nil {
                                dt.Parent = m
                                m.Child = dt
                            }
                        case "org.apache.plc4x.java.s7.readwrite.S7ParameterReadVarResponse":
                            var dt *S7ParameterReadVarResponse
                            if m.Child != nil {
                                dt = m.Child.(*S7ParameterReadVarResponse)
                            }
                            if err := d.DecodeElement(&dt, &tok); err != nil {
                                return err
                            }
                            if m.Child == nil {
                                dt.Parent = m
                                m.Child = dt
                            }
                        case "org.apache.plc4x.java.s7.readwrite.S7ParameterWriteVarRequest":
                            var dt *S7ParameterWriteVarRequest
                            if m.Child != nil {
                                dt = m.Child.(*S7ParameterWriteVarRequest)
                            }
                            if err := d.DecodeElement(&dt, &tok); err != nil {
                                return err
                            }
                            if m.Child == nil {
                                dt.Parent = m
                                m.Child = dt
                            }
                        case "org.apache.plc4x.java.s7.readwrite.S7ParameterWriteVarResponse":
                            var dt *S7ParameterWriteVarResponse
                            if m.Child != nil {
                                dt = m.Child.(*S7ParameterWriteVarResponse)
                            }
                            if err := d.DecodeElement(&dt, &tok); err != nil {
                                return err
                            }
                            if m.Child == nil {
                                dt.Parent = m
                                m.Child = dt
                            }
                        case "org.apache.plc4x.java.s7.readwrite.S7ParameterUserData":
                            var dt *S7ParameterUserData
                            if m.Child != nil {
                                dt = m.Child.(*S7ParameterUserData)
                            }
                            if err := d.DecodeElement(&dt, &tok); err != nil {
                                return err
                            }
                            if m.Child == nil {
                                dt.Parent = m
                                m.Child = dt
                            }
                    }
            }
        }
    }
}

func (m *S7Parameter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := reflect.TypeOf(m.Child).String()
    className = "org.apache.plc4x.java.s7.readwrite." + className[strings.LastIndex(className, ".") + 1:]
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    marshaller, ok := m.Child.(xml.Marshaler)
    if !ok {
        return errors.New("child is not castable to Marshaler")
    }
    marshaller.MarshalXML(e, start)
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

