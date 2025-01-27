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

// BACnetClosingTag is the data-structure of this message
type BACnetClosingTag struct {
	Header *BACnetTagHeader

	// Arguments.
	TagNumberArgument uint8
}

// IBACnetClosingTag is the corresponding interface of BACnetClosingTag
type IBACnetClosingTag interface {
	// GetHeader returns Header (property field)
	GetHeader() *BACnetTagHeader
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

func (m *BACnetClosingTag) GetHeader() *BACnetTagHeader {
	return m.Header
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetClosingTag factory function for BACnetClosingTag
func NewBACnetClosingTag(header *BACnetTagHeader, tagNumberArgument uint8) *BACnetClosingTag {
	return &BACnetClosingTag{Header: header, TagNumberArgument: tagNumberArgument}
}

func CastBACnetClosingTag(structType interface{}) *BACnetClosingTag {
	if casted, ok := structType.(BACnetClosingTag); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetClosingTag); ok {
		return casted
	}
	return nil
}

func (m *BACnetClosingTag) GetTypeName() string {
	return "BACnetClosingTag"
}

func (m *BACnetClosingTag) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetClosingTag) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetClosingTag) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetClosingTagParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8) (*BACnetClosingTag, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetClosingTag"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, pullErr
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field")
	}
	header := CastBACnetTagHeader(_header)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, closeErr
	}

	// Validation
	if !(bool((header.GetActualTagNumber()) == (tagNumberArgument))) {
		return nil, utils.ParseAssertError{"tagnumber doesn't match"}
	}

	// Validation
	if !(bool((header.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) {
		return nil, utils.ParseValidationError{"should be a context tag"}
	}

	// Validation
	if !(bool((header.GetLengthValueType()) == (7))) {
		return nil, utils.ParseValidationError{"closing tag should have a value of 7"}
	}

	if closeErr := readBuffer.CloseContext("BACnetClosingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetClosingTag(header, tagNumberArgument), nil
}

func (m *BACnetClosingTag) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetClosingTag"); pushErr != nil {
		return pushErr
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return pushErr
	}
	_headerErr := m.Header.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return popErr
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	if popErr := writeBuffer.PopContext("BACnetClosingTag"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetClosingTag) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
