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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean is the data-structure of this message
type BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean struct {
	*BACnetNotificationParametersChangeOfDiscreteValueNewValue
	BooleanValue *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber uint8
}

// IBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
type IBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean interface {
	IBACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() *BACnetApplicationTagBoolean
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

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) InitializeParent(parent *BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.OpeningTag = openingTag
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetParent() *BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetBooleanValue() *BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean factory function for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(booleanValue *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean {
	_result := &BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean{
		BooleanValue: booleanValue,
		BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(structType interface{}) *BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return CastBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return CastBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (booleanValue)
	if pullErr := readBuffer.PullContext("booleanValue"); pullErr != nil {
		return nil, pullErr
	}
	_booleanValue, _booleanValueErr := BACnetApplicationTagParse(readBuffer)
	if _booleanValueErr != nil {
		return nil, errors.Wrap(_booleanValueErr, "Error parsing 'booleanValue' field")
	}
	booleanValue := CastBACnetApplicationTagBoolean(_booleanValue)
	if closeErr := readBuffer.CloseContext("booleanValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean{
		BooleanValue: CastBACnetApplicationTagBoolean(booleanValue),
		BACnetNotificationParametersChangeOfDiscreteValueNewValue: &BACnetNotificationParametersChangeOfDiscreteValueNewValue{},
	}
	_child.BACnetNotificationParametersChangeOfDiscreteValueNewValue.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); pushErr != nil {
			return pushErr
		}

		// Simple Field (booleanValue)
		if pushErr := writeBuffer.PushContext("booleanValue"); pushErr != nil {
			return pushErr
		}
		_booleanValueErr := m.BooleanValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("booleanValue"); popErr != nil {
			return popErr
		}
		if _booleanValueErr != nil {
			return errors.Wrap(_booleanValueErr, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
