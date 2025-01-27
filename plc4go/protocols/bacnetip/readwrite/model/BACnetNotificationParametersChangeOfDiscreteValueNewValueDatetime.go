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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime is the data-structure of this message
type BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime struct {
	*BACnetNotificationParametersChangeOfDiscreteValueNewValue
	DateTimeValue *BACnetDateTimeEnclosed

	// Arguments.
	TagNumber uint8
}

// IBACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime
type IBACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime interface {
	IBACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetDateTimeValue returns DateTimeValue (property field)
	GetDateTimeValue() *BACnetDateTimeEnclosed
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

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) InitializeParent(parent *BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.OpeningTag = openingTag
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) GetParent() *BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) GetDateTimeValue() *BACnetDateTimeEnclosed {
	return m.DateTimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime factory function for BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime(dateTimeValue *BACnetDateTimeEnclosed, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime {
	_result := &BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime{
		DateTimeValue: dateTimeValue,
		BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime(structType interface{}) *BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return CastBACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return CastBACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime"
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dateTimeValue)
	lengthInBits += m.DateTimeValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetimeParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateTimeValue)
	if pullErr := readBuffer.PullContext("dateTimeValue"); pullErr != nil {
		return nil, pullErr
	}
	_dateTimeValue, _dateTimeValueErr := BACnetDateTimeEnclosedParse(readBuffer, uint8(uint8(0)))
	if _dateTimeValueErr != nil {
		return nil, errors.Wrap(_dateTimeValueErr, "Error parsing 'dateTimeValue' field")
	}
	dateTimeValue := CastBACnetDateTimeEnclosed(_dateTimeValue)
	if closeErr := readBuffer.CloseContext("dateTimeValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime{
		DateTimeValue: CastBACnetDateTimeEnclosed(dateTimeValue),
		BACnetNotificationParametersChangeOfDiscreteValueNewValue: &BACnetNotificationParametersChangeOfDiscreteValueNewValue{},
	}
	_child.BACnetNotificationParametersChangeOfDiscreteValueNewValue.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (dateTimeValue)
		if pushErr := writeBuffer.PushContext("dateTimeValue"); pushErr != nil {
			return pushErr
		}
		_dateTimeValueErr := m.DateTimeValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dateTimeValue"); popErr != nil {
			return popErr
		}
		if _dateTimeValueErr != nil {
			return errors.Wrap(_dateTimeValueErr, "Error serializing 'dateTimeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
