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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestGetEnrollmentSummary is the corresponding interface of BACnetConfirmedServiceRequestGetEnrollmentSummary
type BACnetConfirmedServiceRequestGetEnrollmentSummary interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetAcknowledgmentFilter returns AcknowledgmentFilter (property field)
	GetAcknowledgmentFilter() BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged
	// GetEnrollmentFilter returns EnrollmentFilter (property field)
	GetEnrollmentFilter() BACnetRecipientProcessEnclosed
	// GetEventStateFilter returns EventStateFilter (property field)
	GetEventStateFilter() BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged
	// GetEventTypeFilter returns EventTypeFilter (property field)
	GetEventTypeFilter() BACnetEventTypeTagged
	// GetPriorityFilter returns PriorityFilter (property field)
	GetPriorityFilter() BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
	// GetNotificationClassFilter returns NotificationClassFilter (property field)
	GetNotificationClassFilter() BACnetContextTagUnsignedInteger
}

// BACnetConfirmedServiceRequestGetEnrollmentSummaryExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestGetEnrollmentSummary.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestGetEnrollmentSummaryExactly interface {
	BACnetConfirmedServiceRequestGetEnrollmentSummary
	isBACnetConfirmedServiceRequestGetEnrollmentSummary() bool
}

// _BACnetConfirmedServiceRequestGetEnrollmentSummary is the data-structure of this message
type _BACnetConfirmedServiceRequestGetEnrollmentSummary struct {
	*_BACnetConfirmedServiceRequest
	AcknowledgmentFilter    BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged
	EnrollmentFilter        BACnetRecipientProcessEnclosed
	EventStateFilter        BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged
	EventTypeFilter         BACnetEventTypeTagged
	PriorityFilter          BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter
	NotificationClassFilter BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetAcknowledgmentFilter() BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged {
	return m.AcknowledgmentFilter
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetEnrollmentFilter() BACnetRecipientProcessEnclosed {
	return m.EnrollmentFilter
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetEventStateFilter() BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged {
	return m.EventStateFilter
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetEventTypeFilter() BACnetEventTypeTagged {
	return m.EventTypeFilter
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetPriorityFilter() BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter {
	return m.PriorityFilter
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetNotificationClassFilter() BACnetContextTagUnsignedInteger {
	return m.NotificationClassFilter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestGetEnrollmentSummary factory function for _BACnetConfirmedServiceRequestGetEnrollmentSummary
func NewBACnetConfirmedServiceRequestGetEnrollmentSummary(acknowledgmentFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged, enrollmentFilter BACnetRecipientProcessEnclosed, eventStateFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged, eventTypeFilter BACnetEventTypeTagged, priorityFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter, notificationClassFilter BACnetContextTagUnsignedInteger, serviceRequestLength uint16) *_BACnetConfirmedServiceRequestGetEnrollmentSummary {
	_result := &_BACnetConfirmedServiceRequestGetEnrollmentSummary{
		AcknowledgmentFilter:           acknowledgmentFilter,
		EnrollmentFilter:               enrollmentFilter,
		EventStateFilter:               eventStateFilter,
		EventTypeFilter:                eventTypeFilter,
		PriorityFilter:                 priorityFilter,
		NotificationClassFilter:        notificationClassFilter,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestGetEnrollmentSummary(structType interface{}) BACnetConfirmedServiceRequestGetEnrollmentSummary {
	if casted, ok := structType.(BACnetConfirmedServiceRequestGetEnrollmentSummary); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestGetEnrollmentSummary); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetTypeName() string {
	return "BACnetConfirmedServiceRequestGetEnrollmentSummary"
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (acknowledgmentFilter)
	lengthInBits += m.AcknowledgmentFilter.GetLengthInBits()

	// Optional Field (enrollmentFilter)
	if m.EnrollmentFilter != nil {
		lengthInBits += m.EnrollmentFilter.GetLengthInBits()
	}

	// Optional Field (eventStateFilter)
	if m.EventStateFilter != nil {
		lengthInBits += m.EventStateFilter.GetLengthInBits()
	}

	// Optional Field (eventTypeFilter)
	if m.EventTypeFilter != nil {
		lengthInBits += m.EventTypeFilter.GetLengthInBits()
	}

	// Optional Field (priorityFilter)
	if m.PriorityFilter != nil {
		lengthInBits += m.PriorityFilter.GetLengthInBits()
	}

	// Optional Field (notificationClassFilter)
	if m.NotificationClassFilter != nil {
		lengthInBits += m.NotificationClassFilter.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetConfirmedServiceRequestGetEnrollmentSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestGetEnrollmentSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestGetEnrollmentSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (acknowledgmentFilter)
	if pullErr := readBuffer.PullContext("acknowledgmentFilter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for acknowledgmentFilter")
	}
	_acknowledgmentFilter, _acknowledgmentFilterErr := BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _acknowledgmentFilterErr != nil {
		return nil, errors.Wrap(_acknowledgmentFilterErr, "Error parsing 'acknowledgmentFilter' field of BACnetConfirmedServiceRequestGetEnrollmentSummary")
	}
	acknowledgmentFilter := _acknowledgmentFilter.(BACnetConfirmedServiceRequestGetEnrollmentSummaryAcknowledgementFilterTagged)
	if closeErr := readBuffer.CloseContext("acknowledgmentFilter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for acknowledgmentFilter")
	}

	// Optional Field (enrollmentFilter) (Can be skipped, if a given expression evaluates to false)
	var enrollmentFilter BACnetRecipientProcessEnclosed = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("enrollmentFilter"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for enrollmentFilter")
		}
		_val, _err := BACnetRecipientProcessEnclosedParse(readBuffer, uint8(1))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'enrollmentFilter' field of BACnetConfirmedServiceRequestGetEnrollmentSummary")
		default:
			enrollmentFilter = _val.(BACnetRecipientProcessEnclosed)
			if closeErr := readBuffer.CloseContext("enrollmentFilter"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for enrollmentFilter")
			}
		}
	}

	// Optional Field (eventStateFilter) (Can be skipped, if a given expression evaluates to false)
	var eventStateFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("eventStateFilter"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for eventStateFilter")
		}
		_val, _err := BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTaggedParse(readBuffer, uint8(2), TagClass_CONTEXT_SPECIFIC_TAGS)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'eventStateFilter' field of BACnetConfirmedServiceRequestGetEnrollmentSummary")
		default:
			eventStateFilter = _val.(BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged)
			if closeErr := readBuffer.CloseContext("eventStateFilter"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for eventStateFilter")
			}
		}
	}

	// Optional Field (eventTypeFilter) (Can be skipped, if a given expression evaluates to false)
	var eventTypeFilter BACnetEventTypeTagged = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("eventTypeFilter"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for eventTypeFilter")
		}
		_val, _err := BACnetEventTypeTaggedParse(readBuffer, uint8(3), TagClass_CONTEXT_SPECIFIC_TAGS)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'eventTypeFilter' field of BACnetConfirmedServiceRequestGetEnrollmentSummary")
		default:
			eventTypeFilter = _val.(BACnetEventTypeTagged)
			if closeErr := readBuffer.CloseContext("eventTypeFilter"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for eventTypeFilter")
			}
		}
	}

	// Optional Field (priorityFilter) (Can be skipped, if a given expression evaluates to false)
	var priorityFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("priorityFilter"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for priorityFilter")
		}
		_val, _err := BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilterParse(readBuffer, uint8(4))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'priorityFilter' field of BACnetConfirmedServiceRequestGetEnrollmentSummary")
		default:
			priorityFilter = _val.(BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter)
			if closeErr := readBuffer.CloseContext("priorityFilter"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for priorityFilter")
			}
		}
	}

	// Optional Field (notificationClassFilter) (Can be skipped, if a given expression evaluates to false)
	var notificationClassFilter BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("notificationClassFilter"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for notificationClassFilter")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(5), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'notificationClassFilter' field of BACnetConfirmedServiceRequestGetEnrollmentSummary")
		default:
			notificationClassFilter = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("notificationClassFilter"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for notificationClassFilter")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestGetEnrollmentSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestGetEnrollmentSummary")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestGetEnrollmentSummary{
		AcknowledgmentFilter:    acknowledgmentFilter,
		EnrollmentFilter:        enrollmentFilter,
		EventStateFilter:        eventStateFilter,
		EventTypeFilter:         eventTypeFilter,
		PriorityFilter:          priorityFilter,
		NotificationClassFilter: notificationClassFilter,
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestGetEnrollmentSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestGetEnrollmentSummary")
		}

		// Simple Field (acknowledgmentFilter)
		if pushErr := writeBuffer.PushContext("acknowledgmentFilter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for acknowledgmentFilter")
		}
		_acknowledgmentFilterErr := writeBuffer.WriteSerializable(m.GetAcknowledgmentFilter())
		if popErr := writeBuffer.PopContext("acknowledgmentFilter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for acknowledgmentFilter")
		}
		if _acknowledgmentFilterErr != nil {
			return errors.Wrap(_acknowledgmentFilterErr, "Error serializing 'acknowledgmentFilter' field")
		}

		// Optional Field (enrollmentFilter) (Can be skipped, if the value is null)
		var enrollmentFilter BACnetRecipientProcessEnclosed = nil
		if m.GetEnrollmentFilter() != nil {
			if pushErr := writeBuffer.PushContext("enrollmentFilter"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for enrollmentFilter")
			}
			enrollmentFilter = m.GetEnrollmentFilter()
			_enrollmentFilterErr := writeBuffer.WriteSerializable(enrollmentFilter)
			if popErr := writeBuffer.PopContext("enrollmentFilter"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for enrollmentFilter")
			}
			if _enrollmentFilterErr != nil {
				return errors.Wrap(_enrollmentFilterErr, "Error serializing 'enrollmentFilter' field")
			}
		}

		// Optional Field (eventStateFilter) (Can be skipped, if the value is null)
		var eventStateFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged = nil
		if m.GetEventStateFilter() != nil {
			if pushErr := writeBuffer.PushContext("eventStateFilter"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for eventStateFilter")
			}
			eventStateFilter = m.GetEventStateFilter()
			_eventStateFilterErr := writeBuffer.WriteSerializable(eventStateFilter)
			if popErr := writeBuffer.PopContext("eventStateFilter"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for eventStateFilter")
			}
			if _eventStateFilterErr != nil {
				return errors.Wrap(_eventStateFilterErr, "Error serializing 'eventStateFilter' field")
			}
		}

		// Optional Field (eventTypeFilter) (Can be skipped, if the value is null)
		var eventTypeFilter BACnetEventTypeTagged = nil
		if m.GetEventTypeFilter() != nil {
			if pushErr := writeBuffer.PushContext("eventTypeFilter"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for eventTypeFilter")
			}
			eventTypeFilter = m.GetEventTypeFilter()
			_eventTypeFilterErr := writeBuffer.WriteSerializable(eventTypeFilter)
			if popErr := writeBuffer.PopContext("eventTypeFilter"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for eventTypeFilter")
			}
			if _eventTypeFilterErr != nil {
				return errors.Wrap(_eventTypeFilterErr, "Error serializing 'eventTypeFilter' field")
			}
		}

		// Optional Field (priorityFilter) (Can be skipped, if the value is null)
		var priorityFilter BACnetConfirmedServiceRequestGetEnrollmentSummaryPriorityFilter = nil
		if m.GetPriorityFilter() != nil {
			if pushErr := writeBuffer.PushContext("priorityFilter"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for priorityFilter")
			}
			priorityFilter = m.GetPriorityFilter()
			_priorityFilterErr := writeBuffer.WriteSerializable(priorityFilter)
			if popErr := writeBuffer.PopContext("priorityFilter"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for priorityFilter")
			}
			if _priorityFilterErr != nil {
				return errors.Wrap(_priorityFilterErr, "Error serializing 'priorityFilter' field")
			}
		}

		// Optional Field (notificationClassFilter) (Can be skipped, if the value is null)
		var notificationClassFilter BACnetContextTagUnsignedInteger = nil
		if m.GetNotificationClassFilter() != nil {
			if pushErr := writeBuffer.PushContext("notificationClassFilter"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for notificationClassFilter")
			}
			notificationClassFilter = m.GetNotificationClassFilter()
			_notificationClassFilterErr := writeBuffer.WriteSerializable(notificationClassFilter)
			if popErr := writeBuffer.PopContext("notificationClassFilter"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for notificationClassFilter")
			}
			if _notificationClassFilterErr != nil {
				return errors.Wrap(_notificationClassFilterErr, "Error serializing 'notificationClassFilter' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestGetEnrollmentSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestGetEnrollmentSummary")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) isBACnetConfirmedServiceRequestGetEnrollmentSummary() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
