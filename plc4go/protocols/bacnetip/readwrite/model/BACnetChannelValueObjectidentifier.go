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

// BACnetChannelValueObjectidentifier is the data-structure of this message
type BACnetChannelValueObjectidentifier struct {
	*BACnetChannelValue
	ObjectidentifierValue *BACnetApplicationTagObjectIdentifier
}

// IBACnetChannelValueObjectidentifier is the corresponding interface of BACnetChannelValueObjectidentifier
type IBACnetChannelValueObjectidentifier interface {
	IBACnetChannelValue
	// GetObjectidentifierValue returns ObjectidentifierValue (property field)
	GetObjectidentifierValue() *BACnetApplicationTagObjectIdentifier
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

func (m *BACnetChannelValueObjectidentifier) InitializeParent(parent *BACnetChannelValue, peekedTagHeader *BACnetTagHeader) {
	m.BACnetChannelValue.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetChannelValueObjectidentifier) GetParent() *BACnetChannelValue {
	return m.BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetChannelValueObjectidentifier) GetObjectidentifierValue() *BACnetApplicationTagObjectIdentifier {
	return m.ObjectidentifierValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueObjectidentifier factory function for BACnetChannelValueObjectidentifier
func NewBACnetChannelValueObjectidentifier(objectidentifierValue *BACnetApplicationTagObjectIdentifier, peekedTagHeader *BACnetTagHeader) *BACnetChannelValueObjectidentifier {
	_result := &BACnetChannelValueObjectidentifier{
		ObjectidentifierValue: objectidentifierValue,
		BACnetChannelValue:    NewBACnetChannelValue(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetChannelValueObjectidentifier(structType interface{}) *BACnetChannelValueObjectidentifier {
	if casted, ok := structType.(BACnetChannelValueObjectidentifier); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetChannelValueObjectidentifier); ok {
		return casted
	}
	if casted, ok := structType.(BACnetChannelValue); ok {
		return CastBACnetChannelValueObjectidentifier(casted.Child)
	}
	if casted, ok := structType.(*BACnetChannelValue); ok {
		return CastBACnetChannelValueObjectidentifier(casted.Child)
	}
	return nil
}

func (m *BACnetChannelValueObjectidentifier) GetTypeName() string {
	return "BACnetChannelValueObjectidentifier"
}

func (m *BACnetChannelValueObjectidentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetChannelValueObjectidentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectidentifierValue)
	lengthInBits += m.ObjectidentifierValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetChannelValueObjectidentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueObjectidentifierParse(readBuffer utils.ReadBuffer) (*BACnetChannelValueObjectidentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueObjectidentifier"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectidentifierValue)
	if pullErr := readBuffer.PullContext("objectidentifierValue"); pullErr != nil {
		return nil, pullErr
	}
	_objectidentifierValue, _objectidentifierValueErr := BACnetApplicationTagParse(readBuffer)
	if _objectidentifierValueErr != nil {
		return nil, errors.Wrap(_objectidentifierValueErr, "Error parsing 'objectidentifierValue' field")
	}
	objectidentifierValue := CastBACnetApplicationTagObjectIdentifier(_objectidentifierValue)
	if closeErr := readBuffer.CloseContext("objectidentifierValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueObjectidentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetChannelValueObjectidentifier{
		ObjectidentifierValue: CastBACnetApplicationTagObjectIdentifier(objectidentifierValue),
		BACnetChannelValue:    &BACnetChannelValue{},
	}
	_child.BACnetChannelValue.Child = _child
	return _child, nil
}

func (m *BACnetChannelValueObjectidentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueObjectidentifier"); pushErr != nil {
			return pushErr
		}

		// Simple Field (objectidentifierValue)
		if pushErr := writeBuffer.PushContext("objectidentifierValue"); pushErr != nil {
			return pushErr
		}
		_objectidentifierValueErr := m.ObjectidentifierValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("objectidentifierValue"); popErr != nil {
			return popErr
		}
		if _objectidentifierValueErr != nil {
			return errors.Wrap(_objectidentifierValueErr, "Error serializing 'objectidentifierValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueObjectidentifier"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetChannelValueObjectidentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
