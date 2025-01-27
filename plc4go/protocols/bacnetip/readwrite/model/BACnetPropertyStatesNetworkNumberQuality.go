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

// BACnetPropertyStatesNetworkNumberQuality is the data-structure of this message
type BACnetPropertyStatesNetworkNumberQuality struct {
	*BACnetPropertyStates
	NetworkNumberQuality *BACnetNetworkNumberQualityTagged
}

// IBACnetPropertyStatesNetworkNumberQuality is the corresponding interface of BACnetPropertyStatesNetworkNumberQuality
type IBACnetPropertyStatesNetworkNumberQuality interface {
	IBACnetPropertyStates
	// GetNetworkNumberQuality returns NetworkNumberQuality (property field)
	GetNetworkNumberQuality() *BACnetNetworkNumberQualityTagged
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

func (m *BACnetPropertyStatesNetworkNumberQuality) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesNetworkNumberQuality) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesNetworkNumberQuality) GetNetworkNumberQuality() *BACnetNetworkNumberQualityTagged {
	return m.NetworkNumberQuality
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesNetworkNumberQuality factory function for BACnetPropertyStatesNetworkNumberQuality
func NewBACnetPropertyStatesNetworkNumberQuality(networkNumberQuality *BACnetNetworkNumberQualityTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesNetworkNumberQuality {
	_result := &BACnetPropertyStatesNetworkNumberQuality{
		NetworkNumberQuality: networkNumberQuality,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesNetworkNumberQuality(structType interface{}) *BACnetPropertyStatesNetworkNumberQuality {
	if casted, ok := structType.(BACnetPropertyStatesNetworkNumberQuality); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesNetworkNumberQuality); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesNetworkNumberQuality(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesNetworkNumberQuality(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesNetworkNumberQuality) GetTypeName() string {
	return "BACnetPropertyStatesNetworkNumberQuality"
}

func (m *BACnetPropertyStatesNetworkNumberQuality) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesNetworkNumberQuality) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (networkNumberQuality)
	lengthInBits += m.NetworkNumberQuality.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesNetworkNumberQuality) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesNetworkNumberQualityParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesNetworkNumberQuality, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesNetworkNumberQuality"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkNumberQuality)
	if pullErr := readBuffer.PullContext("networkNumberQuality"); pullErr != nil {
		return nil, pullErr
	}
	_networkNumberQuality, _networkNumberQualityErr := BACnetNetworkNumberQualityTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _networkNumberQualityErr != nil {
		return nil, errors.Wrap(_networkNumberQualityErr, "Error parsing 'networkNumberQuality' field")
	}
	networkNumberQuality := CastBACnetNetworkNumberQualityTagged(_networkNumberQuality)
	if closeErr := readBuffer.CloseContext("networkNumberQuality"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesNetworkNumberQuality"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesNetworkNumberQuality{
		NetworkNumberQuality: CastBACnetNetworkNumberQualityTagged(networkNumberQuality),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesNetworkNumberQuality) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesNetworkNumberQuality"); pushErr != nil {
			return pushErr
		}

		// Simple Field (networkNumberQuality)
		if pushErr := writeBuffer.PushContext("networkNumberQuality"); pushErr != nil {
			return pushErr
		}
		_networkNumberQualityErr := m.NetworkNumberQuality.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("networkNumberQuality"); popErr != nil {
			return popErr
		}
		if _networkNumberQualityErr != nil {
			return errors.Wrap(_networkNumberQualityErr, "Error serializing 'networkNumberQuality' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesNetworkNumberQuality"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesNetworkNumberQuality) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
