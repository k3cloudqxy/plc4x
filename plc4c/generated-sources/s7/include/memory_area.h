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

#ifndef PLC4C_S7_READ_WRITE_MEMORY_AREA_H_
#define PLC4C_S7_READ_WRITE_MEMORY_AREA_H_

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>

// Code generated by code-generation. DO NOT EDIT.


enum plc4c_s7_read_write_memory_area {
  plc4c_s7_read_write_memory_area_COUNTERS = 0x1C,
  plc4c_s7_read_write_memory_area_TIMERS = 0x1D,
  plc4c_s7_read_write_memory_area_DIRECT_PERIPHERAL_ACCESS = 0x80,
  plc4c_s7_read_write_memory_area_INPUTS = 0x81,
  plc4c_s7_read_write_memory_area_OUTPUTS = 0x82,
  plc4c_s7_read_write_memory_area_FLAGS_MARKERS = 0x83,
  plc4c_s7_read_write_memory_area_DATA_BLOCKS = 0x84,
  plc4c_s7_read_write_memory_area_INSTANCE_DATA_BLOCKS = 0x85,
  plc4c_s7_read_write_memory_area_LOCAL_DATA = 0x86
};
typedef enum plc4c_s7_read_write_memory_area plc4c_s7_read_write_memory_area;

// Get an empty NULL-struct
plc4c_s7_read_write_memory_area plc4c_s7_read_write_memory_area_null();

plc4c_return_code plc4c_s7_read_write_memory_area_parse(plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_memory_area* message);

plc4c_return_code plc4c_s7_read_write_memory_area_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_memory_area* message);

plc4c_s7_read_write_memory_area plc4c_s7_read_write_memory_area_value_of(char* value_string);

int plc4c_s7_read_write_memory_area_num_values();

plc4c_s7_read_write_memory_area plc4c_s7_read_write_memory_area_value_for_index(int index);

char* plc4c_s7_read_write_memory_area_get_short_name(plc4c_s7_read_write_memory_area value);
plc4c_s7_read_write_memory_area plc4c_s7_read_write_memory_area_get_first_enum_for_field_short_name(char* value);

uint16_t plc4c_s7_read_write_memory_area_length_in_bytes(plc4c_s7_read_write_memory_area* message);

uint16_t plc4c_s7_read_write_memory_area_length_in_bits(plc4c_s7_read_write_memory_area* message);


#endif  // PLC4C_S7_READ_WRITE_MEMORY_AREA_H_
