/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesEventType is the data-structure of this message
type BACnetPropertyStatesEventType struct {
	*BACnetPropertyStates
	EventType *BACnetEventTypeTagged
}

// IBACnetPropertyStatesEventType is the corresponding interface of BACnetPropertyStatesEventType
type IBACnetPropertyStatesEventType interface {
	IBACnetPropertyStates
	// GetEventType returns EventType (property field)
	GetEventType() *BACnetEventTypeTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetPropertyStatesEventType) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesEventType) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesEventType) GetEventType() *BACnetEventTypeTagged {
	return m.EventType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesEventType factory function for BACnetPropertyStatesEventType
func NewBACnetPropertyStatesEventType(eventType *BACnetEventTypeTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesEventType {
	_result := &BACnetPropertyStatesEventType{
		EventType:            eventType,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesEventType(structType interface{}) *BACnetPropertyStatesEventType {
	if casted, ok := structType.(BACnetPropertyStatesEventType); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesEventType); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesEventType(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesEventType(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesEventType) GetTypeName() string {
	return "BACnetPropertyStatesEventType"
}

func (m *BACnetPropertyStatesEventType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesEventType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (eventType)
	lengthInBits += m.EventType.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesEventType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesEventTypeParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesEventType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesEventType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (eventType)
	if pullErr := readBuffer.PullContext("eventType"); pullErr != nil {
		return nil, pullErr
	}
	_eventType, _eventTypeErr := BACnetEventTypeTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _eventTypeErr != nil {
		return nil, errors.Wrap(_eventTypeErr, "Error parsing 'eventType' field")
	}
	eventType := CastBACnetEventTypeTagged(_eventType)
	if closeErr := readBuffer.CloseContext("eventType"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesEventType"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesEventType{
		EventType:            CastBACnetEventTypeTagged(eventType),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesEventType) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesEventType"); pushErr != nil {
			return pushErr
		}

		// Simple Field (eventType)
		if pushErr := writeBuffer.PushContext("eventType"); pushErr != nil {
			return pushErr
		}
		_eventTypeErr := m.EventType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("eventType"); popErr != nil {
			return popErr
		}
		if _eventTypeErr != nil {
			return errors.Wrap(_eventTypeErr, "Error serializing 'eventType' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesEventType"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesEventType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
