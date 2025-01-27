
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

// Code generated by code-generation. DO NOT EDIT.

#include <unity.h>

#include "plc4c/spi/read_buffer.h"
#include "plc4c/spi/write_buffer.h"

#include "plc4x_message.h"


void internal_assert_arrays_equal(uint8_t* expected_array, uint8_t* actual_array, uint8_t num_bytes);

        
void parser_serializer_test_plc4x_read_write_connection_request() {
            
    uint8_t payload[] = {
        0x01, 0x00, 0x1c, 0x00, 0x01, 0x01, 0x15, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x3a, 0x2f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_plc4x_read_write_plc4x_message* message = NULL;
    return_code = plc4c_plc4x_read_write_plc4x_message_parse(read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error error parsing tpkt packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_plc4x_read_write_plc4x_message_serialize(write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_plc4x_read_write_connection_response() {
            
    uint8_t payload[] = {
        0x01, 0x00, 0x09, 0x00, 0x01, 0x02, 0x00, 0x01, 0x01
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_plc4x_read_write_plc4x_message* message = NULL;
    return_code = plc4c_plc4x_read_write_plc4x_message_parse(read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error error parsing tpkt packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_plc4x_read_write_plc4x_message_serialize(write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_plc4x_read_write_read_request_single_item() {
            
    uint8_t payload[] = {
        0x01, 0x00, 0x1f, 0x00, 0x02, 0x05, 0x00, 0x01, 0x01, 0x06, 0x6c, 0x61, 0x6c, 0x61, 0x6c, 0x61, 0x0e, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x49, 0x4e, 0x54
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_plc4x_read_write_plc4x_message* message = NULL;
    return_code = plc4c_plc4x_read_write_plc4x_message_parse(read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error error parsing tpkt packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_plc4x_read_write_plc4x_message_serialize(write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_plc4x_read_write_read_response_single_item() {
            
    uint8_t payload[] = {
        0x01, 0x00, 0x24, 0x00, 0x02, 0x06, 0x00, 0x01, 0x01, 0x01, 0x06, 0x6c, 0x61, 0x6c, 0x61, 0x6c, 0x61, 0x0e, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x49, 0x4e, 0x54, 0x01, 0x22, 0x3c, 0xeb
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_plc4x_read_write_plc4x_message* message = NULL;
    return_code = plc4c_plc4x_read_write_plc4x_message_parse(read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error error parsing tpkt packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_plc4x_read_write_plc4x_message_serialize(write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_plc4x_read_write_read_request_multiple_items() {
            
    uint8_t payload[] = {
        0x01, 0x01, 0xb3, 0x00, 0x02, 0x05, 0x00, 0x01, 0x10, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x42, 0x4f, 0x4f, 0x4c, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x42, 0x4f, 0x4f, 0x4c, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x42, 0x59, 0x54, 0x45, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x42, 0x59, 0x54, 0x45, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x57, 0x4f, 0x52, 0x44, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x57, 0x4f, 0x52, 0x44, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x44, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x44, 0x57, 0x4f, 0x52, 0x44, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x55, 0x53, 0x49, 0x4e, 0x54, 0x10, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x55, 0x53, 0x49, 0x4e, 0x54, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x55, 0x49, 0x4e, 0x54, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x55, 0x49, 0x4e, 0x54, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x55, 0x44, 0x49, 0x4e, 0x54, 0x10, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x55, 0x44, 0x49, 0x4e, 0x54, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x55, 0x4c, 0x49, 0x4e, 0x54, 0x10, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x55, 0x4c, 0x49, 0x4e, 0x54, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x53, 0x49, 0x4e, 0x54, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x53, 0x49, 0x4e, 0x54, 0x08, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x49, 0x4e, 0x54, 0x0e, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x49, 0x4e, 0x54, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x44, 0x49, 0x4e, 0x54, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x44, 0x49, 0x4e, 0x54, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x4c, 0x49, 0x4e, 0x54, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x4c, 0x49, 0x4e, 0x54, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x52, 0x45, 0x41, 0x4c, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x52, 0x45, 0x41, 0x4c, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x4c, 0x52, 0x45, 0x41, 0x4c, 0x10, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x4c, 0x52, 0x45, 0x41, 0x4c, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x43, 0x48, 0x41, 0x52, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x43, 0x48, 0x41, 0x52, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x57, 0x43, 0x48, 0x41, 0x52, 0x10, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x57, 0x43, 0x48, 0x41, 0x52
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_plc4x_read_write_plc4x_message* message = NULL;
    return_code = plc4c_plc4x_read_write_plc4x_message_parse(read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error error parsing tpkt packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_plc4x_read_write_plc4x_message_serialize(write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_plc4x_read_write_read_response_multiple_items() {
            
    uint8_t payload[] = {
        0x01, 0x02, 0x09, 0x00, 0x02, 0x06, 0x00, 0x01, 0x01, 0x10, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x42, 0x4f, 0x4f, 0x4c, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x42, 0x4f, 0x4f, 0x4c, 0x01, 0x01, 0x00, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x42, 0x59, 0x54, 0x45, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x42, 0x59, 0x54, 0x45, 0x01, 0x02, 0xbf, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x57, 0x4f, 0x52, 0x44, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x57, 0x4f, 0x52, 0x44, 0x01, 0x03, 0x29, 0x1f, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x44, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x44, 0x57, 0x4f, 0x52, 0x44, 0x01, 0x04, 0x78, 0x23, 0x99, 0x39, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x55, 0x53, 0x49, 0x4e, 0x54, 0x10, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x55, 0x53, 0x49, 0x4e, 0x54, 0x01, 0x11, 0xaf, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x55, 0x49, 0x4e, 0x54, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x55, 0x49, 0x4e, 0x54, 0x01, 0x12, 0xaa, 0xe8, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x55, 0x44, 0x49, 0x4e, 0x54, 0x10, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x55, 0x44, 0x49, 0x4e, 0x54, 0x01, 0x13, 0xd1, 0x8a, 0x89, 0xd4, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x55, 0x4c, 0x49, 0x4e, 0x54, 0x10, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x55, 0x4c, 0x49, 0x4e, 0x54, 0x01, 0x14, 0x97, 0x3d, 0xdb, 0x08, 0xe6, 0xb6, 0x5f, 0xad, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x53, 0x49, 0x4e, 0x54, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x53, 0x49, 0x4e, 0x54, 0x01, 0x21, 0x60, 0x08, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x49, 0x4e, 0x54, 0x0e, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x49, 0x4e, 0x54, 0x01, 0x22, 0xe6, 0x80, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x44, 0x49, 0x4e, 0x54, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x44, 0x49, 0x4e, 0x54, 0x01, 0x23, 0x38, 0x82, 0x36, 0xae, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x4c, 0x49, 0x4e, 0x54, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x4c, 0x49, 0x4e, 0x54, 0x01, 0x24, 0xe3, 0xa2, 0xe5, 0x6f, 0xd4, 0xab, 0x3a, 0xcf, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x52, 0x45, 0x41, 0x4c, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x52, 0x45, 0x41, 0x4c, 0x01, 0x31, 0xf3, 0x80, 0xff, 0x42, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x4c, 0x52, 0x45, 0x41, 0x4c, 0x10, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x4c, 0x52, 0x45, 0x41, 0x4c, 0x01, 0x32, 0x17, 0x51, 0x0d, 0x30, 0x33, 0x21, 0x9c, 0x25, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x43, 0x48, 0x41, 0x52, 0x0f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x43, 0x48, 0x41, 0x52, 0x01, 0x41, 0x50, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x57, 0x43, 0x48, 0x41, 0x52, 0x10, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x57, 0x43, 0x48, 0x41, 0x52, 0x01, 0x42, 0x03, 0xe0
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_plc4x_read_write_plc4x_message* message = NULL;
    return_code = plc4c_plc4x_read_write_plc4x_message_parse(read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error error parsing tpkt packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_plc4x_read_write_plc4x_message_serialize(write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_plc4x_read_write_write_request() {
            
    uint8_t payload[] = {
        0x01, 0x00, 0x22, 0x00, 0x02, 0x07, 0x00, 0x01, 0x01, 0x06, 0x6c, 0x61, 0x6c, 0x61, 0x6c, 0x61, 0x0e, 0x53, 0x54, 0x44, 0x4f, 0x55, 0x54, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x49, 0x4e, 0x54, 0x22, 0x00, 0x17
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_plc4x_read_write_plc4x_message* message = NULL;
    return_code = plc4c_plc4x_read_write_plc4x_message_parse(read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error error parsing tpkt packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_plc4x_read_write_plc4x_message_serialize(write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        
void parser_serializer_test_plc4x_read_write_write_response() {
            
    uint8_t payload[] = {
        0x01, 0x00, 0x21, 0x00, 0x02, 0x08, 0x00, 0x01, 0x01, 0x01, 0x06, 0x6c, 0x61, 0x6c, 0x61, 0x6c, 0x61, 0x0e, 0x53, 0x54, 0x44, 0x4f, 0x55, 0x54, 0x2f, 0x66, 0x6f, 0x6f, 0x3a, 0x49, 0x4e, 0x54, 0x01
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, &read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_plc4x_read_write_plc4x_message* message = NULL;
    return_code = plc4c_plc4x_read_write_plc4x_message_parse(read_buffer, &message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error error parsing tpkt packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, &write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_plc4x_read_write_plc4x_message_serialize(write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer->data, payload_size);

    printf("Success");
}
        


void parser_serializer_test_plc4x_read_write(void) {
        
    RUN_TEST(
        parser_serializer_test_plc4x_read_write_connection_request);
        
    RUN_TEST(
        parser_serializer_test_plc4x_read_write_connection_response);
        
    RUN_TEST(
        parser_serializer_test_plc4x_read_write_read_request_single_item);
        
    RUN_TEST(
        parser_serializer_test_plc4x_read_write_read_response_single_item);
        
    RUN_TEST(
        parser_serializer_test_plc4x_read_write_read_request_multiple_items);
        
    RUN_TEST(
        parser_serializer_test_plc4x_read_write_read_response_multiple_items);
        
    RUN_TEST(
        parser_serializer_test_plc4x_read_write_write_request);
        
    RUN_TEST(
        parser_serializer_test_plc4x_read_write_write_response);
        
}
    