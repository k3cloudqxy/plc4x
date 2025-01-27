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

// BACnetNotificationParametersChangeOfCharacterString is the data-structure of this message
type BACnetNotificationParametersChangeOfCharacterString struct {
	*BACnetNotificationParameters
	InnerOpeningTag *BACnetOpeningTag
	ChangedValue    *BACnetContextTagCharacterString
	StatusFlags     *BACnetStatusFlags
	AlarmValue      *BACnetContextTagCharacterString
	InnerClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber  uint8
	ObjectType BACnetObjectType
}

// IBACnetNotificationParametersChangeOfCharacterString is the corresponding interface of BACnetNotificationParametersChangeOfCharacterString
type IBACnetNotificationParametersChangeOfCharacterString interface {
	IBACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() *BACnetOpeningTag
	// GetChangedValue returns ChangedValue (property field)
	GetChangedValue() *BACnetContextTagCharacterString
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() *BACnetStatusFlags
	// GetAlarmValue returns AlarmValue (property field)
	GetAlarmValue() *BACnetContextTagCharacterString
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

func (m *BACnetNotificationParametersChangeOfCharacterString) InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParameters.OpeningTag = openingTag
	m.BACnetNotificationParameters.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParameters.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersChangeOfCharacterString) GetParent() *BACnetNotificationParameters {
	return m.BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfCharacterString) GetInnerOpeningTag() *BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *BACnetNotificationParametersChangeOfCharacterString) GetChangedValue() *BACnetContextTagCharacterString {
	return m.ChangedValue
}

func (m *BACnetNotificationParametersChangeOfCharacterString) GetStatusFlags() *BACnetStatusFlags {
	return m.StatusFlags
}

func (m *BACnetNotificationParametersChangeOfCharacterString) GetAlarmValue() *BACnetContextTagCharacterString {
	return m.AlarmValue
}

func (m *BACnetNotificationParametersChangeOfCharacterString) GetInnerClosingTag() *BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfCharacterString factory function for BACnetNotificationParametersChangeOfCharacterString
func NewBACnetNotificationParametersChangeOfCharacterString(innerOpeningTag *BACnetOpeningTag, changedValue *BACnetContextTagCharacterString, statusFlags *BACnetStatusFlags, alarmValue *BACnetContextTagCharacterString, innerClosingTag *BACnetClosingTag, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, objectType BACnetObjectType) *BACnetNotificationParametersChangeOfCharacterString {
	_result := &BACnetNotificationParametersChangeOfCharacterString{
		InnerOpeningTag:              innerOpeningTag,
		ChangedValue:                 changedValue,
		StatusFlags:                  statusFlags,
		AlarmValue:                   alarmValue,
		InnerClosingTag:              innerClosingTag,
		BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectType),
	}
	_result.Child = _result
	return _result
}

func CastBACnetNotificationParametersChangeOfCharacterString(structType interface{}) *BACnetNotificationParametersChangeOfCharacterString {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfCharacterString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersChangeOfCharacterString(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersChangeOfCharacterString(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfCharacterString) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfCharacterString"
}

func (m *BACnetNotificationParametersChangeOfCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (changedValue)
	lengthInBits += m.ChangedValue.GetLengthInBits()

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits()

	// Simple field (alarmValue)
	lengthInBits += m.AlarmValue.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfCharacterStringParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, peekedTagNumber uint8) (*BACnetNotificationParametersChangeOfCharacterString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfCharacterString"); pullErr != nil {
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

	// Simple Field (changedValue)
	if pullErr := readBuffer.PullContext("changedValue"); pullErr != nil {
		return nil, pullErr
	}
	_changedValue, _changedValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _changedValueErr != nil {
		return nil, errors.Wrap(_changedValueErr, "Error parsing 'changedValue' field")
	}
	changedValue := CastBACnetContextTagCharacterString(_changedValue)
	if closeErr := readBuffer.CloseContext("changedValue"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, pullErr
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsParse(readBuffer, uint8(uint8(1)))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field")
	}
	statusFlags := CastBACnetStatusFlags(_statusFlags)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (alarmValue)
	if pullErr := readBuffer.PullContext("alarmValue"); pullErr != nil {
		return nil, pullErr
	}
	_alarmValue, _alarmValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _alarmValueErr != nil {
		return nil, errors.Wrap(_alarmValueErr, "Error parsing 'alarmValue' field")
	}
	alarmValue := CastBACnetContextTagCharacterString(_alarmValue)
	if closeErr := readBuffer.CloseContext("alarmValue"); closeErr != nil {
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

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfCharacterString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfCharacterString{
		InnerOpeningTag:              CastBACnetOpeningTag(innerOpeningTag),
		ChangedValue:                 CastBACnetContextTagCharacterString(changedValue),
		StatusFlags:                  CastBACnetStatusFlags(statusFlags),
		AlarmValue:                   CastBACnetContextTagCharacterString(alarmValue),
		InnerClosingTag:              CastBACnetClosingTag(innerClosingTag),
		BACnetNotificationParameters: &BACnetNotificationParameters{},
	}
	_child.BACnetNotificationParameters.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersChangeOfCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfCharacterString"); pushErr != nil {
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

		// Simple Field (changedValue)
		if pushErr := writeBuffer.PushContext("changedValue"); pushErr != nil {
			return pushErr
		}
		_changedValueErr := m.ChangedValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("changedValue"); popErr != nil {
			return popErr
		}
		if _changedValueErr != nil {
			return errors.Wrap(_changedValueErr, "Error serializing 'changedValue' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return pushErr
		}
		_statusFlagsErr := m.StatusFlags.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return popErr
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (alarmValue)
		if pushErr := writeBuffer.PushContext("alarmValue"); pushErr != nil {
			return pushErr
		}
		_alarmValueErr := m.AlarmValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("alarmValue"); popErr != nil {
			return popErr
		}
		if _alarmValueErr != nil {
			return errors.Wrap(_alarmValueErr, "Error serializing 'alarmValue' field")
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

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfCharacterString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
