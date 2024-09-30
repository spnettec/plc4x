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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ReplyOrConfirmationReply is the corresponding interface of ReplyOrConfirmationReply
type ReplyOrConfirmationReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ReplyOrConfirmation
	// GetReply returns Reply (property field)
	GetReply() Reply
	// GetTermination returns Termination (property field)
	GetTermination() ResponseTermination
	// IsReplyOrConfirmationReply is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReplyOrConfirmationReply()
	// CreateBuilder creates a ReplyOrConfirmationReplyBuilder
	CreateReplyOrConfirmationReplyBuilder() ReplyOrConfirmationReplyBuilder
}

// _ReplyOrConfirmationReply is the data-structure of this message
type _ReplyOrConfirmationReply struct {
	ReplyOrConfirmationContract
	Reply       Reply
	Termination ResponseTermination
}

var _ ReplyOrConfirmationReply = (*_ReplyOrConfirmationReply)(nil)
var _ ReplyOrConfirmationRequirements = (*_ReplyOrConfirmationReply)(nil)

// NewReplyOrConfirmationReply factory function for _ReplyOrConfirmationReply
func NewReplyOrConfirmationReply(peekedByte byte, reply Reply, termination ResponseTermination, cBusOptions CBusOptions, requestContext RequestContext) *_ReplyOrConfirmationReply {
	if reply == nil {
		panic("reply of type Reply for ReplyOrConfirmationReply must not be nil")
	}
	if termination == nil {
		panic("termination of type ResponseTermination for ReplyOrConfirmationReply must not be nil")
	}
	_result := &_ReplyOrConfirmationReply{
		ReplyOrConfirmationContract: NewReplyOrConfirmation(peekedByte, cBusOptions, requestContext),
		Reply:                       reply,
		Termination:                 termination,
	}
	_result.ReplyOrConfirmationContract.(*_ReplyOrConfirmation)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ReplyOrConfirmationReplyBuilder is a builder for ReplyOrConfirmationReply
type ReplyOrConfirmationReplyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(reply Reply, termination ResponseTermination) ReplyOrConfirmationReplyBuilder
	// WithReply adds Reply (property field)
	WithReply(Reply) ReplyOrConfirmationReplyBuilder
	// WithReplyBuilder adds Reply (property field) which is build by the builder
	WithReplyBuilder(func(ReplyBuilder) ReplyBuilder) ReplyOrConfirmationReplyBuilder
	// WithTermination adds Termination (property field)
	WithTermination(ResponseTermination) ReplyOrConfirmationReplyBuilder
	// WithTerminationBuilder adds Termination (property field) which is build by the builder
	WithTerminationBuilder(func(ResponseTerminationBuilder) ResponseTerminationBuilder) ReplyOrConfirmationReplyBuilder
	// Build builds the ReplyOrConfirmationReply or returns an error if something is wrong
	Build() (ReplyOrConfirmationReply, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ReplyOrConfirmationReply
}

// NewReplyOrConfirmationReplyBuilder() creates a ReplyOrConfirmationReplyBuilder
func NewReplyOrConfirmationReplyBuilder() ReplyOrConfirmationReplyBuilder {
	return &_ReplyOrConfirmationReplyBuilder{_ReplyOrConfirmationReply: new(_ReplyOrConfirmationReply)}
}

type _ReplyOrConfirmationReplyBuilder struct {
	*_ReplyOrConfirmationReply

	parentBuilder *_ReplyOrConfirmationBuilder

	err *utils.MultiError
}

var _ (ReplyOrConfirmationReplyBuilder) = (*_ReplyOrConfirmationReplyBuilder)(nil)

func (b *_ReplyOrConfirmationReplyBuilder) setParent(contract ReplyOrConfirmationContract) {
	b.ReplyOrConfirmationContract = contract
}

func (b *_ReplyOrConfirmationReplyBuilder) WithMandatoryFields(reply Reply, termination ResponseTermination) ReplyOrConfirmationReplyBuilder {
	return b.WithReply(reply).WithTermination(termination)
}

func (b *_ReplyOrConfirmationReplyBuilder) WithReply(reply Reply) ReplyOrConfirmationReplyBuilder {
	b.Reply = reply
	return b
}

func (b *_ReplyOrConfirmationReplyBuilder) WithReplyBuilder(builderSupplier func(ReplyBuilder) ReplyBuilder) ReplyOrConfirmationReplyBuilder {
	builder := builderSupplier(b.Reply.CreateReplyBuilder())
	var err error
	b.Reply, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ReplyBuilder failed"))
	}
	return b
}

func (b *_ReplyOrConfirmationReplyBuilder) WithTermination(termination ResponseTermination) ReplyOrConfirmationReplyBuilder {
	b.Termination = termination
	return b
}

func (b *_ReplyOrConfirmationReplyBuilder) WithTerminationBuilder(builderSupplier func(ResponseTerminationBuilder) ResponseTerminationBuilder) ReplyOrConfirmationReplyBuilder {
	builder := builderSupplier(b.Termination.CreateResponseTerminationBuilder())
	var err error
	b.Termination, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ResponseTerminationBuilder failed"))
	}
	return b
}

func (b *_ReplyOrConfirmationReplyBuilder) Build() (ReplyOrConfirmationReply, error) {
	if b.Reply == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'reply' not set"))
	}
	if b.Termination == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'termination' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ReplyOrConfirmationReply.deepCopy(), nil
}

func (b *_ReplyOrConfirmationReplyBuilder) MustBuild() ReplyOrConfirmationReply {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ReplyOrConfirmationReplyBuilder) Done() ReplyOrConfirmationBuilder {
	return b.parentBuilder
}

func (b *_ReplyOrConfirmationReplyBuilder) buildForReplyOrConfirmation() (ReplyOrConfirmation, error) {
	return b.Build()
}

func (b *_ReplyOrConfirmationReplyBuilder) DeepCopy() any {
	_copy := b.CreateReplyOrConfirmationReplyBuilder().(*_ReplyOrConfirmationReplyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateReplyOrConfirmationReplyBuilder creates a ReplyOrConfirmationReplyBuilder
func (b *_ReplyOrConfirmationReply) CreateReplyOrConfirmationReplyBuilder() ReplyOrConfirmationReplyBuilder {
	if b == nil {
		return NewReplyOrConfirmationReplyBuilder()
	}
	return &_ReplyOrConfirmationReplyBuilder{_ReplyOrConfirmationReply: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReplyOrConfirmationReply) GetParent() ReplyOrConfirmationContract {
	return m.ReplyOrConfirmationContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyOrConfirmationReply) GetReply() Reply {
	return m.Reply
}

func (m *_ReplyOrConfirmationReply) GetTermination() ResponseTermination {
	return m.Termination
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastReplyOrConfirmationReply(structType any) ReplyOrConfirmationReply {
	if casted, ok := structType.(ReplyOrConfirmationReply); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyOrConfirmationReply); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyOrConfirmationReply) GetTypeName() string {
	return "ReplyOrConfirmationReply"
}

func (m *_ReplyOrConfirmationReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation).getLengthInBits(ctx))

	// Simple field (reply)
	lengthInBits += m.Reply.GetLengthInBits(ctx)

	// Simple field (termination)
	lengthInBits += m.Termination.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ReplyOrConfirmationReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReplyOrConfirmationReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ReplyOrConfirmation, cBusOptions CBusOptions, requestContext RequestContext) (__replyOrConfirmationReply ReplyOrConfirmationReply, err error) {
	m.ReplyOrConfirmationContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReplyOrConfirmationReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyOrConfirmationReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reply, err := ReadSimpleField[Reply](ctx, "reply", ReadComplex[Reply](ReplyParseWithBufferProducer[Reply]((CBusOptions)(cBusOptions), (RequestContext)(requestContext)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reply' field"))
	}
	m.Reply = reply

	termination, err := ReadSimpleField[ResponseTermination](ctx, "termination", ReadComplex[ResponseTermination](ResponseTerminationParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'termination' field"))
	}
	m.Termination = termination

	if closeErr := readBuffer.CloseContext("ReplyOrConfirmationReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyOrConfirmationReply")
	}

	return m, nil
}

func (m *_ReplyOrConfirmationReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReplyOrConfirmationReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReplyOrConfirmationReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReplyOrConfirmationReply")
		}

		if err := WriteSimpleField[Reply](ctx, "reply", m.GetReply(), WriteComplex[Reply](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'reply' field")
		}

		if err := WriteSimpleField[ResponseTermination](ctx, "termination", m.GetTermination(), WriteComplex[ResponseTermination](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'termination' field")
		}

		if popErr := writeBuffer.PopContext("ReplyOrConfirmationReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReplyOrConfirmationReply")
		}
		return nil
	}
	return m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReplyOrConfirmationReply) IsReplyOrConfirmationReply() {}

func (m *_ReplyOrConfirmationReply) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ReplyOrConfirmationReply) deepCopy() *_ReplyOrConfirmationReply {
	if m == nil {
		return nil
	}
	_ReplyOrConfirmationReplyCopy := &_ReplyOrConfirmationReply{
		m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation).deepCopy(),
		m.Reply.DeepCopy().(Reply),
		m.Termination.DeepCopy().(ResponseTermination),
	}
	m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation)._SubType = m
	return _ReplyOrConfirmationReplyCopy
}

func (m *_ReplyOrConfirmationReply) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
