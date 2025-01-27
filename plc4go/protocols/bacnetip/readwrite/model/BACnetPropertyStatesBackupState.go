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

// BACnetPropertyStatesBackupState is the data-structure of this message
type BACnetPropertyStatesBackupState struct {
	*BACnetPropertyStates
	BackupState *BACnetBackupStateTagged
}

// IBACnetPropertyStatesBackupState is the corresponding interface of BACnetPropertyStatesBackupState
type IBACnetPropertyStatesBackupState interface {
	IBACnetPropertyStates
	// GetBackupState returns BackupState (property field)
	GetBackupState() *BACnetBackupStateTagged
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

func (m *BACnetPropertyStatesBackupState) InitializeParent(parent *BACnetPropertyStates, peekedTagHeader *BACnetTagHeader) {
	m.BACnetPropertyStates.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetPropertyStatesBackupState) GetParent() *BACnetPropertyStates {
	return m.BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetPropertyStatesBackupState) GetBackupState() *BACnetBackupStateTagged {
	return m.BackupState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesBackupState factory function for BACnetPropertyStatesBackupState
func NewBACnetPropertyStatesBackupState(backupState *BACnetBackupStateTagged, peekedTagHeader *BACnetTagHeader) *BACnetPropertyStatesBackupState {
	_result := &BACnetPropertyStatesBackupState{
		BackupState:          backupState,
		BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetPropertyStatesBackupState(structType interface{}) *BACnetPropertyStatesBackupState {
	if casted, ok := structType.(BACnetPropertyStatesBackupState); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBackupState); ok {
		return casted
	}
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesBackupState(casted.Child)
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return CastBACnetPropertyStatesBackupState(casted.Child)
	}
	return nil
}

func (m *BACnetPropertyStatesBackupState) GetTypeName() string {
	return "BACnetPropertyStatesBackupState"
}

func (m *BACnetPropertyStatesBackupState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesBackupState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (backupState)
	lengthInBits += m.BackupState.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetPropertyStatesBackupState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesBackupStateParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (*BACnetPropertyStatesBackupState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBackupState"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (backupState)
	if pullErr := readBuffer.PullContext("backupState"); pullErr != nil {
		return nil, pullErr
	}
	_backupState, _backupStateErr := BACnetBackupStateTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _backupStateErr != nil {
		return nil, errors.Wrap(_backupStateErr, "Error parsing 'backupState' field")
	}
	backupState := CastBACnetBackupStateTagged(_backupState)
	if closeErr := readBuffer.CloseContext("backupState"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBackupState"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesBackupState{
		BackupState:          CastBACnetBackupStateTagged(backupState),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child, nil
}

func (m *BACnetPropertyStatesBackupState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBackupState"); pushErr != nil {
			return pushErr
		}

		// Simple Field (backupState)
		if pushErr := writeBuffer.PushContext("backupState"); pushErr != nil {
			return pushErr
		}
		_backupStateErr := m.BackupState.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("backupState"); popErr != nil {
			return popErr
		}
		if _backupStateErr != nil {
			return errors.Wrap(_backupStateErr, "Error serializing 'backupState' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBackupState"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesBackupState) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
