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

// BACnetChannelValueOctetString is the data-structure of this message
type BACnetChannelValueOctetString struct {
	*BACnetChannelValue
	OctetStringValue *BACnetApplicationTagOctetString
}

// IBACnetChannelValueOctetString is the corresponding interface of BACnetChannelValueOctetString
type IBACnetChannelValueOctetString interface {
	IBACnetChannelValue
	// GetOctetStringValue returns OctetStringValue (property field)
	GetOctetStringValue() *BACnetApplicationTagOctetString
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

func (m *BACnetChannelValueOctetString) InitializeParent(parent *BACnetChannelValue, peekedTagHeader *BACnetTagHeader) {
	m.BACnetChannelValue.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetChannelValueOctetString) GetParent() *BACnetChannelValue {
	return m.BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetChannelValueOctetString) GetOctetStringValue() *BACnetApplicationTagOctetString {
	return m.OctetStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueOctetString factory function for BACnetChannelValueOctetString
func NewBACnetChannelValueOctetString(octetStringValue *BACnetApplicationTagOctetString, peekedTagHeader *BACnetTagHeader) *BACnetChannelValueOctetString {
	_result := &BACnetChannelValueOctetString{
		OctetStringValue:   octetStringValue,
		BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetChannelValueOctetString(structType interface{}) *BACnetChannelValueOctetString {
	if casted, ok := structType.(BACnetChannelValueOctetString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetChannelValueOctetString); ok {
		return casted
	}
	if casted, ok := structType.(BACnetChannelValue); ok {
		return CastBACnetChannelValueOctetString(casted.Child)
	}
	if casted, ok := structType.(*BACnetChannelValue); ok {
		return CastBACnetChannelValueOctetString(casted.Child)
	}
	return nil
}

func (m *BACnetChannelValueOctetString) GetTypeName() string {
	return "BACnetChannelValueOctetString"
}

func (m *BACnetChannelValueOctetString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetChannelValueOctetString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (octetStringValue)
	lengthInBits += m.OctetStringValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetChannelValueOctetString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueOctetStringParse(readBuffer utils.ReadBuffer) (*BACnetChannelValueOctetString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueOctetString"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (octetStringValue)
	if pullErr := readBuffer.PullContext("octetStringValue"); pullErr != nil {
		return nil, pullErr
	}
	_octetStringValue, _octetStringValueErr := BACnetApplicationTagParse(readBuffer)
	if _octetStringValueErr != nil {
		return nil, errors.Wrap(_octetStringValueErr, "Error parsing 'octetStringValue' field")
	}
	octetStringValue := CastBACnetApplicationTagOctetString(_octetStringValue)
	if closeErr := readBuffer.CloseContext("octetStringValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueOctetString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetChannelValueOctetString{
		OctetStringValue:   CastBACnetApplicationTagOctetString(octetStringValue),
		BACnetChannelValue: &BACnetChannelValue{},
	}
	_child.BACnetChannelValue.Child = _child
	return _child, nil
}

func (m *BACnetChannelValueOctetString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueOctetString"); pushErr != nil {
			return pushErr
		}

		// Simple Field (octetStringValue)
		if pushErr := writeBuffer.PushContext("octetStringValue"); pushErr != nil {
			return pushErr
		}
		_octetStringValueErr := m.OctetStringValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("octetStringValue"); popErr != nil {
			return popErr
		}
		if _octetStringValueErr != nil {
			return errors.Wrap(_octetStringValueErr, "Error serializing 'octetStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueOctetString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetChannelValueOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
