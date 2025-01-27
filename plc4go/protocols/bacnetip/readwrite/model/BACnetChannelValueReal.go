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

// BACnetChannelValueReal is the data-structure of this message
type BACnetChannelValueReal struct {
	*BACnetChannelValue
	RealValue *BACnetApplicationTagReal
}

// IBACnetChannelValueReal is the corresponding interface of BACnetChannelValueReal
type IBACnetChannelValueReal interface {
	IBACnetChannelValue
	// GetRealValue returns RealValue (property field)
	GetRealValue() *BACnetApplicationTagReal
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

func (m *BACnetChannelValueReal) InitializeParent(parent *BACnetChannelValue, peekedTagHeader *BACnetTagHeader) {
	m.BACnetChannelValue.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetChannelValueReal) GetParent() *BACnetChannelValue {
	return m.BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetChannelValueReal) GetRealValue() *BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueReal factory function for BACnetChannelValueReal
func NewBACnetChannelValueReal(realValue *BACnetApplicationTagReal, peekedTagHeader *BACnetTagHeader) *BACnetChannelValueReal {
	_result := &BACnetChannelValueReal{
		RealValue:          realValue,
		BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetChannelValueReal(structType interface{}) *BACnetChannelValueReal {
	if casted, ok := structType.(BACnetChannelValueReal); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetChannelValueReal); ok {
		return casted
	}
	if casted, ok := structType.(BACnetChannelValue); ok {
		return CastBACnetChannelValueReal(casted.Child)
	}
	if casted, ok := structType.(*BACnetChannelValue); ok {
		return CastBACnetChannelValueReal(casted.Child)
	}
	return nil
}

func (m *BACnetChannelValueReal) GetTypeName() string {
	return "BACnetChannelValueReal"
}

func (m *BACnetChannelValueReal) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetChannelValueReal) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetChannelValueReal) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueRealParse(readBuffer utils.ReadBuffer) (*BACnetChannelValueReal, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueReal"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (realValue)
	if pullErr := readBuffer.PullContext("realValue"); pullErr != nil {
		return nil, pullErr
	}
	_realValue, _realValueErr := BACnetApplicationTagParse(readBuffer)
	if _realValueErr != nil {
		return nil, errors.Wrap(_realValueErr, "Error parsing 'realValue' field")
	}
	realValue := CastBACnetApplicationTagReal(_realValue)
	if closeErr := readBuffer.CloseContext("realValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueReal"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetChannelValueReal{
		RealValue:          CastBACnetApplicationTagReal(realValue),
		BACnetChannelValue: &BACnetChannelValue{},
	}
	_child.BACnetChannelValue.Child = _child
	return _child, nil
}

func (m *BACnetChannelValueReal) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueReal"); pushErr != nil {
			return pushErr
		}

		// Simple Field (realValue)
		if pushErr := writeBuffer.PushContext("realValue"); pushErr != nil {
			return pushErr
		}
		_realValueErr := m.RealValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("realValue"); popErr != nil {
			return popErr
		}
		if _realValueErr != nil {
			return errors.Wrap(_realValueErr, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueReal"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetChannelValueReal) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
