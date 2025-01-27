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

package simulated

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/apache/plc4x/plc4go/pkg/plc4go/values"
	"github.com/apache/plc4x/plc4go/protocols/simulated/readwrite/model"
	"github.com/rs/zerolog/log"
	"math/rand"
)

type Device struct {
	Name  string
	State map[SimulatedField]*values.PlcValue
}

func NewDevice(name string) *Device {
	return &Device{
		Name:  name,
		State: make(map[SimulatedField]*values.PlcValue),
	}
}

func (t *Device) Get(field SimulatedField) *values.PlcValue {
	switch field.FieldType {
	case FieldState:
		return t.State[field]
	case FieldRandom:
		return t.getRandomValue(field)
	}
	return nil
}

func (t *Device) Set(field SimulatedField, value *values.PlcValue) {
	switch field.FieldType {
	case FieldState:
		t.State[field] = value
		break
	case FieldRandom:
		// TODO: Doesn't really make any sense to write a random
		break
	case FieldStdOut:
		log.Debug().Msgf("TEST PLC STDOUT [%s]: %s", field.Name, (*value).GetString())
		break
	}
}

func (t *Device) getRandomValue(field SimulatedField) *values.PlcValue {
	size := field.GetDataTypeSize().DataTypeSize()
	data := make([]byte, uint16(size)*field.Quantity)
	rand.Read(data)
	readBuffer := utils.NewReadBufferByteBased(data)
	plcValue, err := model.DataItemParse(readBuffer, field.DataTypeSize.String(), field.Quantity)
	if err != nil {
		panic("Unable to parse random bytes")
	}
	return &plcValue
}
