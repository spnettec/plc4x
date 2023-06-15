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
	"io"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple is the corresponding interface of BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
type BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() BACnetContextTagUnsignedInteger
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStampEnclosed
	// GetListOfCovNotifications returns ListOfCovNotifications (property field)
	GetListOfCovNotifications() ListOfCovNotificationsList
}

// BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleExactly interface {
	BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
	isBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple() bool
}

// _BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple is the data-structure of this message
type _BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple struct {
	*_BACnetUnconfirmedServiceRequest
        SubscriberProcessIdentifier BACnetContextTagUnsignedInteger
        InitiatingDeviceIdentifier BACnetContextTagObjectIdentifier
        TimeRemaining BACnetContextTagUnsignedInteger
        Timestamp BACnetTimeStampEnclosed
        ListOfCovNotifications ListOfCovNotificationsList
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple)  GetServiceChoice() BACnetUnconfirmedServiceChoice {
return BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) InitializeParent(parent BACnetUnconfirmedServiceRequest ) {}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple)  GetParent() BACnetUnconfirmedServiceRequest {
	return m._BACnetUnconfirmedServiceRequest
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetTimeRemaining() BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetTimestamp() BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetListOfCovNotifications() ListOfCovNotificationsList {
	return m.ListOfCovNotifications
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple factory function for _BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
func NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple( subscriberProcessIdentifier BACnetContextTagUnsignedInteger , initiatingDeviceIdentifier BACnetContextTagObjectIdentifier , timeRemaining BACnetContextTagUnsignedInteger , timestamp BACnetTimeStampEnclosed , listOfCovNotifications ListOfCovNotificationsList , serviceRequestLength uint16 ) *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple {
	_result := &_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple{
		SubscriberProcessIdentifier: subscriberProcessIdentifier,
		InitiatingDeviceIdentifier: initiatingDeviceIdentifier,
		TimeRemaining: timeRemaining,
		Timestamp: timestamp,
		ListOfCovNotifications: listOfCovNotifications,
    	_BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(structType any) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple {
    if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits(ctx)

	// Optional Field (timestamp)
	if m.Timestamp != nil {
		lengthInBits += m.Timestamp.GetLengthInBits(ctx)
	}

	// Simple field (listOfCovNotifications)
	lengthInBits += m.ListOfCovNotifications.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleParse(ctx context.Context, theBytes []byte, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple, error) {
	return BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (subscriberProcessIdentifier)
	if pullErr := readBuffer.PullContext("subscriberProcessIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for subscriberProcessIdentifier")
	}
_subscriberProcessIdentifier, _subscriberProcessIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer , uint8( uint8(0) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field of BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
	}
	subscriberProcessIdentifier := _subscriberProcessIdentifier.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("subscriberProcessIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for subscriberProcessIdentifier")
	}

	// Simple Field (initiatingDeviceIdentifier)
	if pullErr := readBuffer.PullContext("initiatingDeviceIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for initiatingDeviceIdentifier")
	}
_initiatingDeviceIdentifier, _initiatingDeviceIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer , uint8( uint8(1) ) , BACnetDataType( BACnetDataType_BACNET_OBJECT_IDENTIFIER ) )
	if _initiatingDeviceIdentifierErr != nil {
		return nil, errors.Wrap(_initiatingDeviceIdentifierErr, "Error parsing 'initiatingDeviceIdentifier' field of BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
	}
	initiatingDeviceIdentifier := _initiatingDeviceIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("initiatingDeviceIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for initiatingDeviceIdentifier")
	}

	// Simple Field (timeRemaining)
	if pullErr := readBuffer.PullContext("timeRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeRemaining")
	}
_timeRemaining, _timeRemainingErr := BACnetContextTagParseWithBuffer(ctx, readBuffer , uint8( uint8(2) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _timeRemainingErr != nil {
		return nil, errors.Wrap(_timeRemainingErr, "Error parsing 'timeRemaining' field of BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
	}
	timeRemaining := _timeRemaining.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeRemaining")
	}

	// Optional Field (timestamp) (Can be skipped, if a given expression evaluates to false)
	var timestamp BACnetTimeStampEnclosed = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("timestamp"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for timestamp")
		}
_val, _err := BACnetTimeStampEnclosedParseWithBuffer(ctx, readBuffer , uint8(3) )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'timestamp' field of BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
		default:
			timestamp = _val.(BACnetTimeStampEnclosed)
			if closeErr := readBuffer.CloseContext("timestamp"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for timestamp")
			}
		}
	}

	// Simple Field (listOfCovNotifications)
	if pullErr := readBuffer.PullContext("listOfCovNotifications"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfCovNotifications")
	}
_listOfCovNotifications, _listOfCovNotificationsErr := ListOfCovNotificationsListParseWithBuffer(ctx, readBuffer , uint8( uint8(4) ) )
	if _listOfCovNotificationsErr != nil {
		return nil, errors.Wrap(_listOfCovNotificationsErr, "Error parsing 'listOfCovNotifications' field of BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
	}
	listOfCovNotifications := _listOfCovNotifications.(ListOfCovNotificationsList)
	if closeErr := readBuffer.CloseContext("listOfCovNotifications"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfCovNotifications")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple{
		_BACnetUnconfirmedServiceRequest: &_BACnetUnconfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		SubscriberProcessIdentifier: subscriberProcessIdentifier,
		InitiatingDeviceIdentifier: initiatingDeviceIdentifier,
		TimeRemaining: timeRemaining,
		Timestamp: timestamp,
		ListOfCovNotifications: listOfCovNotifications,
	}
	_child._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
		}

	// Simple Field (subscriberProcessIdentifier)
	if pushErr := writeBuffer.PushContext("subscriberProcessIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for subscriberProcessIdentifier")
	}
	_subscriberProcessIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetSubscriberProcessIdentifier())
	if popErr := writeBuffer.PopContext("subscriberProcessIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for subscriberProcessIdentifier")
	}
	if _subscriberProcessIdentifierErr != nil {
		return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
	}

	// Simple Field (initiatingDeviceIdentifier)
	if pushErr := writeBuffer.PushContext("initiatingDeviceIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for initiatingDeviceIdentifier")
	}
	_initiatingDeviceIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetInitiatingDeviceIdentifier())
	if popErr := writeBuffer.PopContext("initiatingDeviceIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for initiatingDeviceIdentifier")
	}
	if _initiatingDeviceIdentifierErr != nil {
		return errors.Wrap(_initiatingDeviceIdentifierErr, "Error serializing 'initiatingDeviceIdentifier' field")
	}

	// Simple Field (timeRemaining)
	if pushErr := writeBuffer.PushContext("timeRemaining"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeRemaining")
	}
	_timeRemainingErr := writeBuffer.WriteSerializable(ctx, m.GetTimeRemaining())
	if popErr := writeBuffer.PopContext("timeRemaining"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeRemaining")
	}
	if _timeRemainingErr != nil {
		return errors.Wrap(_timeRemainingErr, "Error serializing 'timeRemaining' field")
	}

	// Optional Field (timestamp) (Can be skipped, if the value is null)
	var timestamp BACnetTimeStampEnclosed = nil
	if m.GetTimestamp() != nil {
		if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timestamp")
		}
		timestamp = m.GetTimestamp()
		_timestampErr := writeBuffer.WriteSerializable(ctx, timestamp)
		if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timestamp")
		}
		if _timestampErr != nil {
			return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
		}
	}

	// Simple Field (listOfCovNotifications)
	if pushErr := writeBuffer.PushContext("listOfCovNotifications"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfCovNotifications")
	}
	_listOfCovNotificationsErr := writeBuffer.WriteSerializable(ctx, m.GetListOfCovNotifications())
	if popErr := writeBuffer.PopContext("listOfCovNotifications"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfCovNotifications")
	}
	if _listOfCovNotificationsErr != nil {
		return errors.Wrap(_listOfCovNotificationsErr, "Error serializing 'listOfCovNotifications' field")
	}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) isBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



