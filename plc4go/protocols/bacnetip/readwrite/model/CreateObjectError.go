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

// CreateObjectError is the data-structure of this message
type CreateObjectError struct {
	*BACnetError
	ErrorType                *ErrorEnclosed
	FirstFailedElementNumber *BACnetContextTagUnsignedInteger
}

// ICreateObjectError is the corresponding interface of CreateObjectError
type ICreateObjectError interface {
	IBACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() *ErrorEnclosed
	// GetFirstFailedElementNumber returns FirstFailedElementNumber (property field)
	GetFirstFailedElementNumber() *BACnetContextTagUnsignedInteger
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

func (m *CreateObjectError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CREATE_OBJECT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *CreateObjectError) InitializeParent(parent *BACnetError) {}

func (m *CreateObjectError) GetParent() *BACnetError {
	return m.BACnetError
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CreateObjectError) GetErrorType() *ErrorEnclosed {
	return m.ErrorType
}

func (m *CreateObjectError) GetFirstFailedElementNumber() *BACnetContextTagUnsignedInteger {
	return m.FirstFailedElementNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCreateObjectError factory function for CreateObjectError
func NewCreateObjectError(errorType *ErrorEnclosed, firstFailedElementNumber *BACnetContextTagUnsignedInteger) *CreateObjectError {
	_result := &CreateObjectError{
		ErrorType:                errorType,
		FirstFailedElementNumber: firstFailedElementNumber,
		BACnetError:              NewBACnetError(),
	}
	_result.Child = _result
	return _result
}

func CastCreateObjectError(structType interface{}) *CreateObjectError {
	if casted, ok := structType.(CreateObjectError); ok {
		return &casted
	}
	if casted, ok := structType.(*CreateObjectError); ok {
		return casted
	}
	if casted, ok := structType.(BACnetError); ok {
		return CastCreateObjectError(casted.Child)
	}
	if casted, ok := structType.(*BACnetError); ok {
		return CastCreateObjectError(casted.Child)
	}
	return nil
}

func (m *CreateObjectError) GetTypeName() string {
	return "CreateObjectError"
}

func (m *CreateObjectError) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CreateObjectError) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits()

	// Simple field (firstFailedElementNumber)
	lengthInBits += m.FirstFailedElementNumber.GetLengthInBits()

	return lengthInBits
}

func (m *CreateObjectError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CreateObjectErrorParse(readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (*CreateObjectError, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CreateObjectError"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (errorType)
	if pullErr := readBuffer.PullContext("errorType"); pullErr != nil {
		return nil, pullErr
	}
	_errorType, _errorTypeErr := ErrorEnclosedParse(readBuffer, uint8(uint8(0)))
	if _errorTypeErr != nil {
		return nil, errors.Wrap(_errorTypeErr, "Error parsing 'errorType' field")
	}
	errorType := CastErrorEnclosed(_errorType)
	if closeErr := readBuffer.CloseContext("errorType"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (firstFailedElementNumber)
	if pullErr := readBuffer.PullContext("firstFailedElementNumber"); pullErr != nil {
		return nil, pullErr
	}
	_firstFailedElementNumber, _firstFailedElementNumberErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _firstFailedElementNumberErr != nil {
		return nil, errors.Wrap(_firstFailedElementNumberErr, "Error parsing 'firstFailedElementNumber' field")
	}
	firstFailedElementNumber := CastBACnetContextTagUnsignedInteger(_firstFailedElementNumber)
	if closeErr := readBuffer.CloseContext("firstFailedElementNumber"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("CreateObjectError"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CreateObjectError{
		ErrorType:                CastErrorEnclosed(errorType),
		FirstFailedElementNumber: CastBACnetContextTagUnsignedInteger(firstFailedElementNumber),
		BACnetError:              &BACnetError{},
	}
	_child.BACnetError.Child = _child
	return _child, nil
}

func (m *CreateObjectError) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateObjectError"); pushErr != nil {
			return pushErr
		}

		// Simple Field (errorType)
		if pushErr := writeBuffer.PushContext("errorType"); pushErr != nil {
			return pushErr
		}
		_errorTypeErr := m.ErrorType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("errorType"); popErr != nil {
			return popErr
		}
		if _errorTypeErr != nil {
			return errors.Wrap(_errorTypeErr, "Error serializing 'errorType' field")
		}

		// Simple Field (firstFailedElementNumber)
		if pushErr := writeBuffer.PushContext("firstFailedElementNumber"); pushErr != nil {
			return pushErr
		}
		_firstFailedElementNumberErr := m.FirstFailedElementNumber.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("firstFailedElementNumber"); popErr != nil {
			return popErr
		}
		if _firstFailedElementNumberErr != nil {
			return errors.Wrap(_firstFailedElementNumberErr, "Error serializing 'firstFailedElementNumber' field")
		}

		if popErr := writeBuffer.PopContext("CreateObjectError"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CreateObjectError) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
