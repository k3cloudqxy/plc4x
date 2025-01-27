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

// BACnetNotificationParametersChangeOfStatusFlags is the data-structure of this message
type BACnetNotificationParametersChangeOfStatusFlags struct {
	*BACnetNotificationParameters
	InnerOpeningTag *BACnetOpeningTag
	PresentValue    *BACnetConstructedData
	ReferencedFlags *BACnetStatusFlags
	InnerClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber  uint8
	ObjectType BACnetObjectType
}

// IBACnetNotificationParametersChangeOfStatusFlags is the corresponding interface of BACnetNotificationParametersChangeOfStatusFlags
type IBACnetNotificationParametersChangeOfStatusFlags interface {
	IBACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() *BACnetOpeningTag
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() *BACnetConstructedData
	// GetReferencedFlags returns ReferencedFlags (property field)
	GetReferencedFlags() *BACnetStatusFlags
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() *BACnetClosingTag
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

func (m *BACnetNotificationParametersChangeOfStatusFlags) InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParameters.OpeningTag = openingTag
	m.BACnetNotificationParameters.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParameters.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersChangeOfStatusFlags) GetParent() *BACnetNotificationParameters {
	return m.BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfStatusFlags) GetInnerOpeningTag() *BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *BACnetNotificationParametersChangeOfStatusFlags) GetPresentValue() *BACnetConstructedData {
	return m.PresentValue
}

func (m *BACnetNotificationParametersChangeOfStatusFlags) GetReferencedFlags() *BACnetStatusFlags {
	return m.ReferencedFlags
}

func (m *BACnetNotificationParametersChangeOfStatusFlags) GetInnerClosingTag() *BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfStatusFlags factory function for BACnetNotificationParametersChangeOfStatusFlags
func NewBACnetNotificationParametersChangeOfStatusFlags(innerOpeningTag *BACnetOpeningTag, presentValue *BACnetConstructedData, referencedFlags *BACnetStatusFlags, innerClosingTag *BACnetClosingTag, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, objectType BACnetObjectType) *BACnetNotificationParametersChangeOfStatusFlags {
	_result := &BACnetNotificationParametersChangeOfStatusFlags{
		InnerOpeningTag:              innerOpeningTag,
		PresentValue:                 presentValue,
		ReferencedFlags:              referencedFlags,
		InnerClosingTag:              innerClosingTag,
		BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectType),
	}
	_result.Child = _result
	return _result
}

func CastBACnetNotificationParametersChangeOfStatusFlags(structType interface{}) *BACnetNotificationParametersChangeOfStatusFlags {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfStatusFlags); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfStatusFlags); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersChangeOfStatusFlags(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersChangeOfStatusFlags(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfStatusFlags) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfStatusFlags"
}

func (m *BACnetNotificationParametersChangeOfStatusFlags) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfStatusFlags) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits()

	// Simple field (referencedFlags)
	lengthInBits += m.ReferencedFlags.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfStatusFlags) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfStatusFlagsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, peekedTagNumber uint8) (*BACnetNotificationParametersChangeOfStatusFlags, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfStatusFlags"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := CastBACnetOpeningTag(_innerOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (presentValue)
	if pullErr := readBuffer.PullContext("presentValue"); pullErr != nil {
		return nil, pullErr
	}
	_presentValue, _presentValueErr := BACnetConstructedDataParse(readBuffer, uint8(uint8(0)), BACnetObjectType(objectType), BACnetPropertyIdentifier(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE))
	if _presentValueErr != nil {
		return nil, errors.Wrap(_presentValueErr, "Error parsing 'presentValue' field")
	}
	presentValue := CastBACnetConstructedData(_presentValue)
	if closeErr := readBuffer.CloseContext("presentValue"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (referencedFlags)
	if pullErr := readBuffer.PullContext("referencedFlags"); pullErr != nil {
		return nil, pullErr
	}
	_referencedFlags, _referencedFlagsErr := BACnetStatusFlagsParse(readBuffer, uint8(uint8(1)))
	if _referencedFlagsErr != nil {
		return nil, errors.Wrap(_referencedFlagsErr, "Error parsing 'referencedFlags' field")
	}
	referencedFlags := CastBACnetStatusFlags(_referencedFlags)
	if closeErr := readBuffer.CloseContext("referencedFlags"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := CastBACnetClosingTag(_innerClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfStatusFlags"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfStatusFlags{
		InnerOpeningTag:              CastBACnetOpeningTag(innerOpeningTag),
		PresentValue:                 CastBACnetConstructedData(presentValue),
		ReferencedFlags:              CastBACnetStatusFlags(referencedFlags),
		InnerClosingTag:              CastBACnetClosingTag(innerClosingTag),
		BACnetNotificationParameters: &BACnetNotificationParameters{},
	}
	_child.BACnetNotificationParameters.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersChangeOfStatusFlags) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfStatusFlags"); pushErr != nil {
			return pushErr
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return pushErr
		}
		_innerOpeningTagErr := m.InnerOpeningTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return popErr
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (presentValue)
		if pushErr := writeBuffer.PushContext("presentValue"); pushErr != nil {
			return pushErr
		}
		_presentValueErr := m.PresentValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("presentValue"); popErr != nil {
			return popErr
		}
		if _presentValueErr != nil {
			return errors.Wrap(_presentValueErr, "Error serializing 'presentValue' field")
		}

		// Simple Field (referencedFlags)
		if pushErr := writeBuffer.PushContext("referencedFlags"); pushErr != nil {
			return pushErr
		}
		_referencedFlagsErr := m.ReferencedFlags.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("referencedFlags"); popErr != nil {
			return popErr
		}
		if _referencedFlagsErr != nil {
			return errors.Wrap(_referencedFlagsErr, "Error serializing 'referencedFlags' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return pushErr
		}
		_innerClosingTagErr := m.InnerClosingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return popErr
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfStatusFlags"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfStatusFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
