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

// BACnetEventTransitionBitsApplication is the data-structure of this message
type BACnetEventTransitionBitsApplication struct {
	RawBits *BACnetApplicationTagBitString
}

// IBACnetEventTransitionBitsApplication is the corresponding interface of BACnetEventTransitionBitsApplication
type IBACnetEventTransitionBitsApplication interface {
	// GetRawBits returns RawBits (property field)
	GetRawBits() *BACnetApplicationTagBitString
	// GetToOffnormal returns ToOffnormal (virtual field)
	GetToOffnormal() bool
	// GetToFault returns ToFault (virtual field)
	GetToFault() bool
	// GetToNormal returns ToNormal (virtual field)
	GetToNormal() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventTransitionBitsApplication) GetRawBits() *BACnetApplicationTagBitString {
	return m.RawBits
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetEventTransitionBitsApplication) GetToOffnormal() bool {
	return bool(m.GetRawBits().GetPayload().GetData()[0])
}

func (m *BACnetEventTransitionBitsApplication) GetToFault() bool {
	return bool(m.GetRawBits().GetPayload().GetData()[1])
}

func (m *BACnetEventTransitionBitsApplication) GetToNormal() bool {
	return bool(m.GetRawBits().GetPayload().GetData()[2])
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventTransitionBitsApplication factory function for BACnetEventTransitionBitsApplication
func NewBACnetEventTransitionBitsApplication(rawBits *BACnetApplicationTagBitString) *BACnetEventTransitionBitsApplication {
	return &BACnetEventTransitionBitsApplication{RawBits: rawBits}
}

func CastBACnetEventTransitionBitsApplication(structType interface{}) *BACnetEventTransitionBitsApplication {
	if casted, ok := structType.(BACnetEventTransitionBitsApplication); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventTransitionBitsApplication); ok {
		return casted
	}
	return nil
}

func (m *BACnetEventTransitionBitsApplication) GetTypeName() string {
	return "BACnetEventTransitionBitsApplication"
}

func (m *BACnetEventTransitionBitsApplication) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventTransitionBitsApplication) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (rawBits)
	lengthInBits += m.RawBits.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetEventTransitionBitsApplication) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventTransitionBitsApplicationParse(readBuffer utils.ReadBuffer) (*BACnetEventTransitionBitsApplication, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventTransitionBitsApplication"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (rawBits)
	if pullErr := readBuffer.PullContext("rawBits"); pullErr != nil {
		return nil, pullErr
	}
	_rawBits, _rawBitsErr := BACnetApplicationTagParse(readBuffer)
	if _rawBitsErr != nil {
		return nil, errors.Wrap(_rawBitsErr, "Error parsing 'rawBits' field")
	}
	rawBits := CastBACnetApplicationTagBitString(_rawBits)
	if closeErr := readBuffer.CloseContext("rawBits"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_toOffnormal := rawBits.GetPayload().GetData()[0]
	toOffnormal := bool(_toOffnormal)
	_ = toOffnormal

	// Virtual field
	_toFault := rawBits.GetPayload().GetData()[1]
	toFault := bool(_toFault)
	_ = toFault

	// Virtual field
	_toNormal := rawBits.GetPayload().GetData()[2]
	toNormal := bool(_toNormal)
	_ = toNormal

	if closeErr := readBuffer.CloseContext("BACnetEventTransitionBitsApplication"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetEventTransitionBitsApplication(rawBits), nil
}

func (m *BACnetEventTransitionBitsApplication) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventTransitionBitsApplication"); pushErr != nil {
		return pushErr
	}

	// Simple Field (rawBits)
	if pushErr := writeBuffer.PushContext("rawBits"); pushErr != nil {
		return pushErr
	}
	_rawBitsErr := m.RawBits.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("rawBits"); popErr != nil {
		return popErr
	}
	if _rawBitsErr != nil {
		return errors.Wrap(_rawBitsErr, "Error serializing 'rawBits' field")
	}
	// Virtual field
	if _toOffnormalErr := writeBuffer.WriteVirtual("toOffnormal", m.GetToOffnormal()); _toOffnormalErr != nil {
		return errors.Wrap(_toOffnormalErr, "Error serializing 'toOffnormal' field")
	}
	// Virtual field
	if _toFaultErr := writeBuffer.WriteVirtual("toFault", m.GetToFault()); _toFaultErr != nil {
		return errors.Wrap(_toFaultErr, "Error serializing 'toFault' field")
	}
	// Virtual field
	if _toNormalErr := writeBuffer.WriteVirtual("toNormal", m.GetToNormal()); _toNormalErr != nil {
		return errors.Wrap(_toNormalErr, "Error serializing 'toNormal' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventTransitionBitsApplication"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetEventTransitionBitsApplication) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
