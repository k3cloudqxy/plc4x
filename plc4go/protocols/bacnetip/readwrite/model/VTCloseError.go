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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// VTCloseError is the data-structure of this message
type VTCloseError struct {
	*BACnetError
	ErrorType                  *ErrorEnclosed
	ListOfVtSessionIdentifiers *VTCloseErrorListOfVTSessionIdentifiers
}

// IVTCloseError is the corresponding interface of VTCloseError
type IVTCloseError interface {
	IBACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() *ErrorEnclosed
	// GetListOfVtSessionIdentifiers returns ListOfVtSessionIdentifiers (property field)
	GetListOfVtSessionIdentifiers() *VTCloseErrorListOfVTSessionIdentifiers
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

func (m *VTCloseError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_CLOSE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *VTCloseError) InitializeParent(parent *BACnetError) {}

func (m *VTCloseError) GetParent() *BACnetError {
	return m.BACnetError
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *VTCloseError) GetErrorType() *ErrorEnclosed {
	return m.ErrorType
}

func (m *VTCloseError) GetListOfVtSessionIdentifiers() *VTCloseErrorListOfVTSessionIdentifiers {
	return m.ListOfVtSessionIdentifiers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewVTCloseError factory function for VTCloseError
func NewVTCloseError(errorType *ErrorEnclosed, listOfVtSessionIdentifiers *VTCloseErrorListOfVTSessionIdentifiers) *VTCloseError {
	_result := &VTCloseError{
		ErrorType:                  errorType,
		ListOfVtSessionIdentifiers: listOfVtSessionIdentifiers,
		BACnetError:                NewBACnetError(),
	}
	_result.Child = _result
	return _result
}

func CastVTCloseError(structType interface{}) *VTCloseError {
	if casted, ok := structType.(VTCloseError); ok {
		return &casted
	}
	if casted, ok := structType.(*VTCloseError); ok {
		return casted
	}
	if casted, ok := structType.(BACnetError); ok {
		return CastVTCloseError(casted.Child)
	}
	if casted, ok := structType.(*BACnetError); ok {
		return CastVTCloseError(casted.Child)
	}
	return nil
}

func (m *VTCloseError) GetTypeName() string {
	return "VTCloseError"
}

func (m *VTCloseError) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *VTCloseError) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits()

	// Optional Field (listOfVtSessionIdentifiers)
	if m.ListOfVtSessionIdentifiers != nil {
		lengthInBits += (*m.ListOfVtSessionIdentifiers).GetLengthInBits()
	}

	return lengthInBits
}

func (m *VTCloseError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func VTCloseErrorParse(readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (*VTCloseError, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VTCloseError"); pullErr != nil {
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

	// Optional Field (listOfVtSessionIdentifiers) (Can be skipped, if a given expression evaluates to false)
	var listOfVtSessionIdentifiers *VTCloseErrorListOfVTSessionIdentifiers = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("listOfVtSessionIdentifiers"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := VTCloseErrorListOfVTSessionIdentifiersParse(readBuffer, uint8(1))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'listOfVtSessionIdentifiers' field")
		default:
			listOfVtSessionIdentifiers = CastVTCloseErrorListOfVTSessionIdentifiers(_val)
			if closeErr := readBuffer.CloseContext("listOfVtSessionIdentifiers"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("VTCloseError"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &VTCloseError{
		ErrorType:                  CastErrorEnclosed(errorType),
		ListOfVtSessionIdentifiers: CastVTCloseErrorListOfVTSessionIdentifiers(listOfVtSessionIdentifiers),
		BACnetError:                &BACnetError{},
	}
	_child.BACnetError.Child = _child
	return _child, nil
}

func (m *VTCloseError) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VTCloseError"); pushErr != nil {
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

		// Optional Field (listOfVtSessionIdentifiers) (Can be skipped, if the value is null)
		var listOfVtSessionIdentifiers *VTCloseErrorListOfVTSessionIdentifiers = nil
		if m.ListOfVtSessionIdentifiers != nil {
			if pushErr := writeBuffer.PushContext("listOfVtSessionIdentifiers"); pushErr != nil {
				return pushErr
			}
			listOfVtSessionIdentifiers = m.ListOfVtSessionIdentifiers
			_listOfVtSessionIdentifiersErr := listOfVtSessionIdentifiers.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("listOfVtSessionIdentifiers"); popErr != nil {
				return popErr
			}
			if _listOfVtSessionIdentifiersErr != nil {
				return errors.Wrap(_listOfVtSessionIdentifiersErr, "Error serializing 'listOfVtSessionIdentifiers' field")
			}
		}

		if popErr := writeBuffer.PopContext("VTCloseError"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *VTCloseError) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
