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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned is the data-structure of this message
type BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned struct {
	*BACnetNotificationParametersChangeOfDiscreteValueNewValue
	UnsignedValue *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned
type IBACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned interface {
	IBACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) InitializeParent(parent *BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.OpeningTag = openingTag
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) GetParent() *BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) GetUnsignedValue() *BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned factory function for BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned(unsignedValue *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned {
	_result := &BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned{
		UnsignedValue: unsignedValue,
		BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned(structType interface{}) *BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return CastBACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return CastBACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned"
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsignedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unsignedValue)
	if pullErr := readBuffer.PullContext("unsignedValue"); pullErr != nil {
		return nil, pullErr
	}
	_unsignedValue, _unsignedValueErr := BACnetApplicationTagParse(readBuffer)
	if _unsignedValueErr != nil {
		return nil, errors.Wrap(_unsignedValueErr, "Error parsing 'unsignedValue' field")
	}
	unsignedValue := CastBACnetApplicationTagUnsignedInteger(_unsignedValue)
	if closeErr := readBuffer.CloseContext("unsignedValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned{
		UnsignedValue: CastBACnetApplicationTagUnsignedInteger(unsignedValue),
		BACnetNotificationParametersChangeOfDiscreteValueNewValue: &BACnetNotificationParametersChangeOfDiscreteValueNewValue{},
	}
	_child.BACnetNotificationParametersChangeOfDiscreteValueNewValue.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned"); pushErr != nil {
			return pushErr
		}

		// Simple Field (unsignedValue)
		if pushErr := writeBuffer.PushContext("unsignedValue"); pushErr != nil {
			return pushErr
		}
		_unsignedValueErr := m.UnsignedValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("unsignedValue"); popErr != nil {
			return popErr
		}
		if _unsignedValueErr != nil {
			return errors.Wrap(_unsignedValueErr, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
