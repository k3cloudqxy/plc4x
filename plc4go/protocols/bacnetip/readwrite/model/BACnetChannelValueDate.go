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

// BACnetChannelValueDate is the data-structure of this message
type BACnetChannelValueDate struct {
	*BACnetChannelValue
	DateValue *BACnetApplicationTagDate
}

// IBACnetChannelValueDate is the corresponding interface of BACnetChannelValueDate
type IBACnetChannelValueDate interface {
	IBACnetChannelValue
	// GetDateValue returns DateValue (property field)
	GetDateValue() *BACnetApplicationTagDate
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

func (m *BACnetChannelValueDate) InitializeParent(parent *BACnetChannelValue, peekedTagHeader *BACnetTagHeader) {
	m.BACnetChannelValue.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetChannelValueDate) GetParent() *BACnetChannelValue {
	return m.BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetChannelValueDate) GetDateValue() *BACnetApplicationTagDate {
	return m.DateValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueDate factory function for BACnetChannelValueDate
func NewBACnetChannelValueDate(dateValue *BACnetApplicationTagDate, peekedTagHeader *BACnetTagHeader) *BACnetChannelValueDate {
	_result := &BACnetChannelValueDate{
		DateValue:          dateValue,
		BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetChannelValueDate(structType interface{}) *BACnetChannelValueDate {
	if casted, ok := structType.(BACnetChannelValueDate); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetChannelValueDate); ok {
		return casted
	}
	if casted, ok := structType.(BACnetChannelValue); ok {
		return CastBACnetChannelValueDate(casted.Child)
	}
	if casted, ok := structType.(*BACnetChannelValue); ok {
		return CastBACnetChannelValueDate(casted.Child)
	}
	return nil
}

func (m *BACnetChannelValueDate) GetTypeName() string {
	return "BACnetChannelValueDate"
}

func (m *BACnetChannelValueDate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetChannelValueDate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetChannelValueDate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueDateParse(readBuffer utils.ReadBuffer) (*BACnetChannelValueDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueDate"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateValue)
	if pullErr := readBuffer.PullContext("dateValue"); pullErr != nil {
		return nil, pullErr
	}
	_dateValue, _dateValueErr := BACnetApplicationTagParse(readBuffer)
	if _dateValueErr != nil {
		return nil, errors.Wrap(_dateValueErr, "Error parsing 'dateValue' field")
	}
	dateValue := CastBACnetApplicationTagDate(_dateValue)
	if closeErr := readBuffer.CloseContext("dateValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueDate"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetChannelValueDate{
		DateValue:          CastBACnetApplicationTagDate(dateValue),
		BACnetChannelValue: &BACnetChannelValue{},
	}
	_child.BACnetChannelValue.Child = _child
	return _child, nil
}

func (m *BACnetChannelValueDate) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueDate"); pushErr != nil {
			return pushErr
		}

		// Simple Field (dateValue)
		if pushErr := writeBuffer.PushContext("dateValue"); pushErr != nil {
			return pushErr
		}
		_dateValueErr := m.DateValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dateValue"); popErr != nil {
			return popErr
		}
		if _dateValueErr != nil {
			return errors.Wrap(_dateValueErr, "Error serializing 'dateValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueDate"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetChannelValueDate) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
