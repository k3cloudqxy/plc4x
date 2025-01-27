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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger is the data-structure of this message
type BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger struct {
	*BACnetNotificationParametersChangeOfDiscreteValueNewValue
	IntegerValue *BACnetApplicationTagSignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetNotificationParametersChangeOfDiscreteValueNewValueInteger is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger
type IBACnetNotificationParametersChangeOfDiscreteValueNewValueInteger interface {
	IBACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() *BACnetApplicationTagSignedInteger
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

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger) InitializeParent(parent *BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.OpeningTag = openingTag
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger) GetParent() *BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger) GetIntegerValue() *BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueInteger factory function for BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueInteger(integerValue *BACnetApplicationTagSignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger {
	_result := &BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger{
		IntegerValue: integerValue,
		BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueInteger(structType interface{}) *BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return CastBACnetNotificationParametersChangeOfDiscreteValueNewValueInteger(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return CastBACnetNotificationParametersChangeOfDiscreteValueNewValueInteger(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger"
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueIntegerParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integerValue)
	if pullErr := readBuffer.PullContext("integerValue"); pullErr != nil {
		return nil, pullErr
	}
	_integerValue, _integerValueErr := BACnetApplicationTagParse(readBuffer)
	if _integerValueErr != nil {
		return nil, errors.Wrap(_integerValueErr, "Error parsing 'integerValue' field")
	}
	integerValue := CastBACnetApplicationTagSignedInteger(_integerValue)
	if closeErr := readBuffer.CloseContext("integerValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger{
		IntegerValue: CastBACnetApplicationTagSignedInteger(integerValue),
		BACnetNotificationParametersChangeOfDiscreteValueNewValue: &BACnetNotificationParametersChangeOfDiscreteValueNewValue{},
	}
	_child.BACnetNotificationParametersChangeOfDiscreteValueNewValue.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger"); pushErr != nil {
			return pushErr
		}

		// Simple Field (integerValue)
		if pushErr := writeBuffer.PushContext("integerValue"); pushErr != nil {
			return pushErr
		}
		_integerValueErr := m.IntegerValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("integerValue"); popErr != nil {
			return popErr
		}
		if _integerValueErr != nil {
			return errors.Wrap(_integerValueErr, "Error serializing 'integerValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
