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
package org.apache.plc4x.java.bacnetip;

import org.junit.jupiter.api.Test;


import static org.apache.plc4x.java.bacnetip.Utils.PAYLOAD_START_INDEX;
import static org.apache.plc4x.java.bacnetip.Utils.tryParseBytes;

public class TrickyPackagesTest {

    // from plugfest-tridium-1.pcap
    @Test
    public void testTridium255() throws Exception {
        int[] rawBytesAsInts = new int[]{
            /*0000*/   0x00, 0x1d, 0x09, 0xbc, 0x43, 0x5f, 0x00, 0x01, 0xf0, 0x8c, 0x11, 0x50, 0x08, 0x00, 0x45, 0x00,
            /*0010*/   0x01, 0xfc, 0x7e, 0x47, 0x00, 0x00, 0x40, 0x11, 0x27, 0x92, 0xac, 0x10, 0x24, 0xcd, 0xac, 0x10,
            /*0020*/   0x56, 0x2a, 0xba, 0xc0, 0xba, 0xc0, 0x01, 0xe8, 0x24, 0xd9, 0x81, 0x0a, 0x01, 0xe0, 0x01, 0x00,
            /*0030*/   0x30, 0x6e, 0x0e, 0x0c, 0x04, 0x40, 0x00, 0x01, 0x1e, 0x29, 0x4b, 0x4e, 0xc4, 0x04, 0x40, 0x00,
            /*0040*/   0x01, 0x4f, 0x29, 0x4d, 0x4e, 0x75, 0x13, 0x00, 0x53, 0x63, 0x68, 0x65, 0x64, 0x2e, 0x45, 0x6e,
            /*0050*/   0x75, 0x6d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4f, 0x29, 0x4f, 0x4e, 0x91, 0x11,
            /*0060*/   0x4f, 0x29, 0x55, 0x4e, 0x21, 0x07, 0x4f, 0x29, 0x20, 0x4e, 0xa4, 0xff, 0xff, 0xff, 0xff, 0xa4,
            /*0070*/   0xff, 0xff, 0xff, 0xff, 0x4f, 0x29, 0xae, 0x4e, 0x21, 0x01, 0x4f, 0x29, 0x36, 0x4e, 0x0c, 0x04,
            /*0080*/   0xc0, 0x00, 0x66, 0x19, 0x57, 0x29, 0x0a, 0x0c, 0x04, 0xc0, 0x00, 0x65, 0x19, 0x57, 0x29, 0x0a,
            /*0090*/   0x4f, 0x29, 0x58, 0x4e, 0x21, 0x10, 0x4f, 0x29, 0x6f, 0x4e, 0x82, 0x04, 0x00, 0x4f, 0x29, 0x67,
            /*00a0*/   0x4e, 0x91, 0x00, 0x4f, 0x29, 0x51, 0x4e, 0x10, 0x4f, 0x29, 0x7b, 0x4e, 0x0e, 0xb4, 0x06, 0x00,
            /*00b0*/   0x00, 0x00, 0x21, 0x01, 0xb4, 0x09, 0x00, 0x00, 0x00, 0x21, 0x02, 0xb4, 0x0c, 0x00, 0x00, 0x00,
            /*00c0*/   0x21, 0x03, 0xb4, 0x0f, 0x00, 0x00, 0x00, 0x21, 0x04, 0xb4, 0x12, 0x00, 0x00, 0x00, 0x00, 0x0f,
            /*00d0*/   0x0e, 0xb4, 0x06, 0x00, 0x00, 0x00, 0x21, 0x01, 0xb4, 0x09, 0x00, 0x00, 0x00, 0x21, 0x02, 0xb4,
            /*00e0*/   0x0c, 0x00, 0x00, 0x00, 0x21, 0x03, 0xb4, 0x0f, 0x00, 0x00, 0x00, 0x21, 0x04, 0xb4, 0x12, 0x00,
            /*00f0*/   0x00, 0x00, 0x00, 0x0f, 0x0e, 0xb4, 0x06, 0x00, 0x00, 0x00, 0x21, 0x01, 0xb4, 0x09, 0x00, 0x00,
            /*0100*/   0x00, 0x21, 0x02, 0xb4, 0x0c, 0x00, 0x00, 0x00, 0x21, 0x03, 0xb4, 0x0f, 0x00, 0x00, 0x00, 0x21,
            /*0110*/   0x04, 0xb4, 0x12, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x0e, 0xb4, 0x06, 0x00, 0x00, 0x00, 0x21, 0x01,
            /*0120*/   0xb4, 0x09, 0x00, 0x00, 0x00, 0x21, 0x02, 0xb4, 0x0c, 0x00, 0x00, 0x00, 0x21, 0x03, 0xb4, 0x0f,
            /*0130*/   0x00, 0x00, 0x00, 0x21, 0x04, 0xb4, 0x12, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x0e, 0xb4, 0x06, 0x00,
            /*0140*/   0x00, 0x00, 0x21, 0x01, 0xb4, 0x09, 0x00, 0x00, 0x00, 0x21, 0x02, 0xb4, 0x0c, 0x00, 0x00, 0x00,
            /*0150*/   0x21, 0x03, 0xb4, 0x0f, 0x00, 0x00, 0x00, 0x21, 0x04, 0xb4, 0x12, 0x00, 0x00, 0x00, 0x00, 0x0f,
            /*0160*/   0x0e, 0xb4, 0x06, 0x00, 0x00, 0x00, 0x21, 0x05, 0xb4, 0x0c, 0x00, 0x00, 0x00, 0x21, 0x06, 0xb4,
            /*0170*/   0x12, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x0e, 0xb4, 0x06, 0x00, 0x00, 0x00, 0x21, 0x05, 0xb4, 0x0c,
            /*0180*/   0x00, 0x00, 0x00, 0x21, 0x06, 0xb4, 0x12, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x4f, 0x29, 0x26, 0x4e,
            /*0190*/   0x0e, 0x0c, 0xff, 0xff, 0xff, 0x03, 0x0f, 0x2e, 0x2f, 0x39, 0x10, 0x1c, 0x01, 0x80, 0x00, 0x01,
            /*01a0*/   0x2e, 0xb4, 0x09, 0x00, 0x00, 0x00, 0x21, 0x08, 0xb4, 0x0c, 0x00, 0x00, 0x00, 0x00, 0xb4, 0x0f,
            /*01b0*/   0x1e, 0x00, 0x00, 0x21, 0x08, 0xb4, 0x12, 0x00, 0x00, 0x00, 0x00, 0x2f, 0x39, 0x10, 0x1c, 0x01,
            /*01c0*/   0x80, 0x00, 0x00, 0x2e, 0xb4, 0x03, 0x00, 0x00, 0x00, 0x21, 0x07, 0xb4, 0x06, 0x00, 0x00, 0x00,
            /*01d0*/   0x00, 0xb4, 0x07, 0x1e, 0x00, 0x00, 0x21, 0x07, 0xb4, 0x09, 0x1e, 0x00, 0x00, 0x00, 0xb4, 0x0a,
            /*01e0*/   0x1e, 0x00, 0x00, 0x21, 0x07, 0xb4, 0x0d, 0x00, 0x00, 0x00, 0x00, 0xb4, 0x0e, 0x1e, 0x00, 0x00,
            /*01f0*/   0x21, 0x07, 0xb4, 0x11, 0x1e, 0x00, 0x00, 0x00, 0xb4, 0x13, 0x00, 0x00, 0x00, 0x21, 0x07, 0xb4,
            /*0200*/   0x14, 0x1e, 0x00, 0x00, 0x00, 0x2f, 0x39, 0x10, 0x4f, 0x1f,
        };
        tryParseBytes(rawBytesAsInts);
    }

    // from plugfest-delta-2b.cap
    @Test
    public void testDelta1204() throws Exception {
        int[] rawBytesAsInts = new int[]{
            /*0000*/   0x00, 0x06, 0x5b, 0xd6, 0x9f, 0xfe, 0x00, 0x40, 0xae, 0x00, 0x45, 0xd7, 0x08, 0x00, 0x45, 0x00,
            /*0010*/   0x01, 0x5e, 0x28, 0x64, 0x00, 0x00, 0x3c, 0x11, 0xec, 0xc9, 0xac, 0x10, 0x08, 0x0a, 0xac, 0x10,
            /*0020*/   0x08, 0x37, 0xba, 0xc0, 0xba, 0xc0, 0x01, 0x4a, 0x00, 0x00, 0x81, 0x0a, 0x01, 0x42, 0x01, 0x00,
            /*0030*/   0x30, 0xef, 0x0e, 0x0c, 0x45, 0x80, 0x00, 0x01, 0x1e, 0x29, 0x1c, 0x4e, 0x71, 0x00, 0x4f, 0x2a,
            /*0040*/   0x04, 0x0a, 0x4e, 0x91, 0x00, 0x4f, 0x2a, 0x04, 0x0c, 0x4e, 0x0e, 0x09, 0x01, 0x19, 0x04, 0x2c,
            /*0050*/   0x01, 0xfe, 0xed, 0xd8, 0x3d, 0x06, 0x03, 0xff, 0x3b, 0xe8, 0x3e, 0xf1, 0x4a, 0x04, 0x57, 0x5a,
            /*0060*/   0x01, 0xe0, 0x69, 0x64, 0x79, 0x0a, 0x89, 0x2a, 0x99, 0x32, 0xaa, 0x13, 0x88, 0xba, 0x0b, 0xb8,
            /*0070*/   0xca, 0x1b, 0x58, 0xd9, 0x03, 0xea, 0x13, 0x88, 0xfa, 0x0f, 0x01, 0xe0, 0xf9, 0x10, 0x00, 0x0f,
            /*0080*/   0x4f, 0x2a, 0x04, 0x27, 0x4e, 0x0a, 0x02, 0x51, 0x19, 0x00, 0x29, 0x00, 0x39, 0x00, 0x4a, 0x06,
            /*0090*/   0x0b, 0x59, 0x00, 0x69, 0x00, 0x79, 0x00, 0x89, 0x00, 0x99, 0x00, 0xa9, 0x00, 0x4f, 0x2a, 0x04,
            /*00a0*/   0x3a, 0x39, 0x00, 0x4e, 0x21, 0x05, 0x4f, 0x2a, 0x04, 0x3a, 0x4e, 0x09, 0x00, 0x19, 0x00, 0x29,
            /*00b0*/   0x01, 0x39, 0x01, 0x4c, 0x45, 0x80, 0x00, 0x01, 0x59, 0x01, 0x69, 0x06, 0x79, 0x01, 0x8e, 0x0b,
            /*00c0*/   0x01, 0xea, 0xc9, 0x19, 0x00, 0x29, 0x00, 0x39, 0x00, 0x49, 0x00, 0x8f, 0x9c, 0x00, 0x00, 0x00,
            /*00d0*/   0x00, 0x09, 0x00, 0x19, 0x00, 0x29, 0x01, 0x39, 0x01, 0x4c, 0x45, 0x80, 0x00, 0x01, 0x59, 0x02,
            /*00e0*/   0x69, 0x06, 0x79, 0x02, 0x8e, 0x0a, 0x79, 0xef, 0x19, 0x00, 0x2b, 0x01, 0x80, 0x7e, 0x3a, 0x05,
            /*00f0*/   0xe7, 0x49, 0x32, 0x8f, 0x9c, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x19, 0x00, 0x29, 0x01, 0x39,
            /*0100*/   0x01, 0x4c, 0x45, 0x80, 0x00, 0x01, 0x59, 0x01, 0x69, 0x01, 0x79, 0x03, 0x8e, 0x0a, 0x01, 0x01,
            /*0110*/   0x19, 0x00, 0x29, 0x00, 0x39, 0x00, 0x49, 0x00, 0x8f, 0x9c, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00,
            /*0120*/   0x19, 0x00, 0x29, 0x01, 0x39, 0x01, 0x4c, 0x45, 0x80, 0x00, 0x01, 0x59, 0x04, 0x69, 0x01, 0x79,
            /*0130*/   0x05, 0x8e, 0x0a, 0x31, 0x20, 0x19, 0x00, 0x2a, 0x2f, 0x37, 0x39, 0xc3, 0x49, 0x00, 0x8f, 0x9c,
            /*0140*/   0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x19, 0x00, 0x29, 0x01, 0x39, 0x01, 0x4c, 0x45, 0x80, 0x00,
            /*0150*/   0x01, 0x59, 0x05, 0x69, 0x01, 0x79, 0x07, 0x8e, 0x0a, 0x26, 0xe1, 0x19, 0x00, 0x2a, 0x25, 0xb9,
            /*0160*/   0x39, 0x00, 0x49, 0x00, 0x8f, 0x9c, 0x00, 0x00, 0x00, 0x00, 0x4f, 0x1f,
        };
        tryParseBytes(rawBytesAsInts, PAYLOAD_START_INDEX);
    }

    // from TrendLogMultipleReadRangeSimple.pcap
    @Test
    public void testMultiReadProperty() throws Exception {
        int[] rawBytesAsInts = new int[]{
            /*00000000*/  0x00, 0x24, 0xe8, 0xf9, 0x21, 0xa0, 0x00, 0x22, 0x5f, 0x07, 0xcc, 0x90, 0x08, 0x00, 0x45, 0x00,  //|.$..!.."_.....E.|
            /*00000010*/  0x03, 0x3b, 0xe6, 0x0a, 0x00, 0x00, 0x80, 0x11, 0xcd, 0x87, 0xc0, 0xa8, 0x01, 0x67, 0xc0, 0xa8,  //|.;...........g..|
            /*00000020*/  0x01, 0x68, 0xba, 0xc0, 0xba, 0xc0, 0x03, 0x27, 0xe7, 0x1d, 0x81, 0x0a, 0x03, 0x1f, 0x01, 0x00,  //|.h.....'........|
            /*00000030*/  0x30, 0x14, 0x0e, 0x0c, 0x02, 0x00, 0x0f, 0xa0, 0x1e, 0x29, 0x4d, 0x4e, 0x75, 0x07, 0x00, 0x44,  //|0........)MNu..D|
            /*00000040*/  0x45, 0x34, 0x30, 0x30, 0x30, 0x4f, 0x1f, 0x0c, 0x05, 0xc0, 0x00, 0x01, 0x1e, 0x29, 0x4d, 0x4e,  //|E4000O.......)MN|
            /*00000050*/  0x74, 0x00, 0x41, 0x43, 0x31, 0x4f, 0x1f, 0x0c, 0x00, 0x00, 0x00, 0x01, 0x1e, 0x29, 0x4d, 0x4e,  //|t.AC1O.......)MN|
            /*00000060*/  0x74, 0x00, 0x41, 0x49, 0x31, 0x4f, 0x1f, 0x0c, 0x00, 0x40, 0x00, 0x01, 0x1e, 0x29, 0x4d, 0x4e,  //|t.AI1O...@...)MN|
            /*00000070*/  0x74, 0x00, 0x41, 0x4f, 0x31, 0x4f, 0x1f, 0x0c, 0x00, 0x80, 0x00, 0x01, 0x1e, 0x29, 0x4d, 0x4e,  //|t.AO1O.......)MN|
            /*00000080*/  0x74, 0x00, 0x41, 0x56, 0x31, 0x4f, 0x1f, 0x0c, 0x00, 0x80, 0x00, 0x02, 0x1e, 0x29, 0x4d, 0x4e,  //|t.AV1O.......)MN|
            /*00000090*/  0x74, 0x00, 0x41, 0x56, 0x32, 0x4f, 0x1f, 0x0c, 0x04, 0x80, 0x00, 0x01, 0x1e, 0x29, 0x4d, 0x4e,  //|t.AV2O.......)MN|
            /*000000a0*/  0x74, 0x00, 0x41, 0x47, 0x31, 0x4f, 0x1f, 0x0c, 0x00, 0xc0, 0x00, 0x01, 0x1e, 0x29, 0x4d, 0x4e,  //|t.AG1O.......)MN|
            /*000000b0*/  0x74, 0x00, 0x42, 0x49, 0x31, 0x4f, 0x1f, 0x0c, 0x01, 0x00, 0x00, 0x01, 0x1e, 0x29, 0x4d, 0x4e,  //|t.BI1O.......)MN|
            /*000000c0*/  0x74, 0x00, 0x42, 0x4f, 0x31, 0x4f, 0x1f, 0x0c, 0x01, 0x40, 0x00, 0x01, 0x1e, 0x29, 0x4d, 0x4e,  //|t.BO1O...@...)MN|
            /*000000d0*/  0x74, 0x00, 0x42, 0x56, 0x31, 0x4f, 0x1f, 0x0c, 0x01, 0x80, 0x00, 0x01, 0x1e, 0x29, 0x4d, 0x4e,  //|t.BV1O.......)MN|
            /*000000e0*/  0x74, 0x00, 0x43, 0x41, 0x31, 0x4f, 0x1f, 0x0c, 0x01, 0xc0, 0x00, 0x01, 0x1e, 0x29, 0x4d, 0x4e,  //|t.CA1O.......)MN|
            /*000000f0*/  0x74, 0x00, 0x43, 0x4f, 0x31, 0x4f, 0x1f, 0x0c, 0x02, 0x40, 0x00, 0x01, 0x1e, 0x29, 0x4d, 0x4e,  //|t.CO1O...@...)MN|
            /*00000100*/  0x74, 0x00, 0x45, 0x45, 0x31, 0x4f, 0x1f, 0x0c, 0x02, 0x40, 0x00, 0x02, 0x1e, 0x29, 0x4d, 0x4e,  //|t.EE1O...@...)MN|
            /*00000110*/  0x74, 0x00, 0x45, 0x45, 0x32, 0x4f, 0x1f, 0x0c, 0x02, 0x40, 0x00, 0x03, 0x1e, 0x29, 0x4d, 0x4e,  //|t.EE2O...@...)MN|
            /*00000120*/  0x74, 0x00, 0x45, 0x45, 0x33, 0x4f, 0x1f, 0x0c, 0x02, 0x40, 0x00, 0x04, 0x1e, 0x29, 0x4d, 0x4e,  //|t.EE3O...@...)MN|
            /*00000130*/  0x74, 0x00, 0x45, 0x45, 0x34, 0x4f, 0x1f, 0x0c, 0x02, 0x40, 0x00, 0x05, 0x1e, 0x29, 0x4d, 0x4e,  //|t.EE4O...@...)MN|
            /*00000140*/  0x74, 0x00, 0x45, 0x45, 0x35, 0x4f, 0x1f, 0x0c, 0x02, 0x40, 0x00, 0x06, 0x1e, 0x29, 0x4d, 0x4e,  //|t.EE5O...@...)MN|
            /*00000150*/  0x74, 0x00, 0x45, 0x45, 0x36, 0x4f, 0x1f, 0x0c, 0x02, 0x40, 0x00, 0x07, 0x1e, 0x29, 0x4d, 0x4e,  //|t.EE6O...@...)MN|
            /*00000160*/  0x74, 0x00, 0x45, 0x45, 0x37, 0x4f, 0x1f, 0x0c, 0x02, 0x40, 0x00, 0x08, 0x1e, 0x29, 0x4d, 0x4e,  //|t.EE7O...@...)MN|
            /*00000170*/  0x74, 0x00, 0x45, 0x45, 0x38, 0x4f, 0x1f, 0x0c, 0x02, 0x40, 0x00, 0x09, 0x1e, 0x29, 0x4d, 0x4e,  //|t.EE8O...@...)MN|
            /*00000180*/  0x74, 0x00, 0x45, 0x45, 0x39, 0x4f, 0x1f, 0x0c, 0x02, 0x40, 0x00, 0x0a, 0x1e, 0x29, 0x4d, 0x4e,  //|t.EE9O...@...)MN|
            /*00000190*/  0x75, 0x05, 0x00, 0x45, 0x45, 0x31, 0x30, 0x4f, 0x1f, 0x0c, 0x02, 0x80, 0x00, 0x01, 0x1e, 0x29,  //|u..EE10O.......)|
            /*000001a0*/  0x4d, 0x4e, 0x74, 0x00, 0x46, 0x49, 0x31, 0x4f, 0x1f, 0x0c, 0x02, 0x80, 0x00, 0x02, 0x1e, 0x29,  //|MNt.FI1O.......)|
            /*000001b0*/  0x4d, 0x4e, 0x74, 0x00, 0x46, 0x49, 0x32, 0x4f, 0x1f, 0x0c, 0x02, 0xc0, 0x00, 0x01, 0x1e, 0x29,  //|MNt.FI2O.......)|
            /*000001c0*/  0x4d, 0x4e, 0x74, 0x00, 0x47, 0x52, 0x31, 0x4f, 0x1f, 0x0c, 0x05, 0x40, 0x00, 0x01, 0x1e, 0x29,  //|MNt.GR1O...@...)|
            /*000001d0*/  0x4d, 0x4e, 0x74, 0x00, 0x4c, 0x50, 0x31, 0x4f, 0x1f, 0x0c, 0x05, 0x80, 0x00, 0x01, 0x1e, 0x29,  //|MNt.LP1O.......)|
            /*000001e0*/  0x4d, 0x4e, 0x74, 0x00, 0x4c, 0x5a, 0x31, 0x4f, 0x1f, 0x0c, 0x03, 0x00, 0x00, 0x01, 0x1e, 0x29,  //|MNt.LZ1O.......)|
            /*000001f0*/  0x4d, 0x4e, 0x74, 0x00, 0x4c, 0x4f, 0x31, 0x4f, 0x1f, 0x0c, 0x03, 0x40, 0x00, 0x01, 0x1e, 0x29,  //|MNt.LO1O...@...)|
            /*00000200*/  0x4d, 0x4e, 0x74, 0x00, 0x4d, 0x49, 0x31, 0x4f, 0x1f, 0x0c, 0x03, 0x80, 0x00, 0x01, 0x1e, 0x29,  //|MNt.MI1O.......)|
            /*00000210*/  0x4d, 0x4e, 0x74, 0x00, 0x4d, 0x4f, 0x31, 0x4f, 0x1f, 0x0c, 0x04, 0xc0, 0x00, 0x01, 0x1e, 0x29,  //|MNt.MO1O.......)|
            /*00000220*/  0x4d, 0x4e, 0x74, 0x00, 0x4d, 0x56, 0x31, 0x4f, 0x1f, 0x0c, 0x03, 0xc0, 0x00, 0x01, 0x1e, 0x29,  //|MNt.MV1O.......)|
            /*00000230*/  0x4d, 0x4e, 0x74, 0x00, 0x4e, 0x43, 0x31, 0x4f, 0x1f, 0x0c, 0x04, 0x00, 0x00, 0x01, 0x1e, 0x29,  //|MNt.NC1O.......)|
            /*00000240*/  0x4d, 0x4e, 0x74, 0x00, 0x50, 0x52, 0x31, 0x4f, 0x1f, 0x0c, 0x04, 0x00, 0x00, 0x02, 0x1e, 0x29,  //|MNt.PR1O.......)|
            /*00000250*/  0x4d, 0x4e, 0x74, 0x00, 0x50, 0x52, 0x32, 0x4f, 0x1f, 0x0c, 0x04, 0x00, 0x00, 0x03, 0x1e, 0x29,  //|MNt.PR2O.......)|
            /*00000260*/  0x4d, 0x4e, 0x74, 0x00, 0x50, 0x52, 0x33, 0x4f, 0x1f, 0x0c, 0x04, 0x00, 0x00, 0x04, 0x1e, 0x29,  //|MNt.PR3O.......)|
            /*00000270*/  0x4d, 0x4e, 0x74, 0x00, 0x50, 0x52, 0x34, 0x4f, 0x1f, 0x0c, 0x04, 0x00, 0x00, 0x05, 0x1e, 0x29,  //|MNt.PR4O.......)|
            /*00000280*/  0x4d, 0x4e, 0x74, 0x00, 0x50, 0x52, 0x35, 0x4f, 0x1f, 0x0c, 0x04, 0x00, 0x00, 0x06, 0x1e, 0x29,  //|MNt.PR5O.......)|
            /*00000290*/  0x4d, 0x4e, 0x74, 0x00, 0x50, 0x52, 0x36, 0x4f, 0x1f, 0x0c, 0x04, 0x00, 0x00, 0x07, 0x1e, 0x29,  //|MNt.PR6O.......)|
            /*000002a0*/  0x4d, 0x4e, 0x74, 0x00, 0x50, 0x52, 0x37, 0x4f, 0x1f, 0x0c, 0x04, 0x00, 0x00, 0x08, 0x1e, 0x29,  //|MNt.PR7O.......)|
            /*000002b0*/  0x4d, 0x4e, 0x74, 0x00, 0x50, 0x52, 0x38, 0x4f, 0x1f, 0x0c, 0x06, 0x00, 0x00, 0x01, 0x1e, 0x29,  //|MNt.PR8O.......)|
            /*000002c0*/  0x4d, 0x4e, 0x74, 0x00, 0x50, 0x43, 0x31, 0x4f, 0x1f, 0x0c, 0x06, 0x00, 0x00, 0x02, 0x1e, 0x29,  //|MNt.PC1O.......)|
            /*000002d0*/  0x4d, 0x4e, 0x74, 0x00, 0x50, 0x43, 0x32, 0x4f, 0x1f, 0x0c, 0x04, 0x40, 0x00, 0x01, 0x1e, 0x29,  //|MNt.PC2O...@...)|
            /*000002e0*/  0x4d, 0x4e, 0x74, 0x00, 0x53, 0x43, 0x31, 0x4f, 0x1f, 0x0c, 0x05, 0x00, 0x00, 0x01, 0x1e, 0x29,  //|MNt.SC1O.......)|
            /*000002f0*/  0x4d, 0x4e, 0x74, 0x00, 0x54, 0x4c, 0x31, 0x4f, 0x1f, 0x0c, 0x07, 0x80, 0x00, 0x01, 0x1e, 0x29,  //|MNt.TL1O.......)|
            /*00000300*/  0x4d, 0x4e, 0x74, 0x00, 0x41, 0x44, 0x31, 0x4f, 0x1f, 0x0c, 0x06, 0x40, 0x00, 0x01, 0x1e, 0x29,  //|MNt.AD1O...@...)|
            /*00000310*/  0x4d, 0x4e, 0x74, 0x00, 0x45, 0x4c, 0x31, 0x4f, 0x1f, 0x0c, 0x07, 0x00, 0x00, 0x01, 0x1e, 0x29,  //|MNt.EL1O.......)|
            /*00000320*/  0x4d, 0x4e, 0x74, 0x00, 0x4c, 0x43, 0x31, 0x4f, 0x1f, 0x0c, 0x07, 0x40, 0x00, 0x01, 0x1e, 0x29,  //|MNt.LC1O...@...)|
            /*00000330*/  0x4d, 0x4e, 0x74, 0x00, 0x53, 0x56, 0x31, 0x4f, 0x1f, 0x0c, 0x06, 0xc0, 0x00, 0x01, 0x1e, 0x29,  //|MNt.SV1O.......)|
            /*00000340*/  0x4d, 0x4e, 0x74, 0x00, 0x54, 0x4d, 0x31, 0x4f, 0x1f                                             //|MNt.TM1O.|
        };
        tryParseBytes(rawBytesAsInts);
    }

    // from TrendLogMultipleReadRangeSimple.pcap
    @Test
    public void testMultiReadProperty2() throws Exception {
        int[] rawBytesAsInts = new int[]{
            /*00000000*/  0x00, 0x1d, 0x09, 0xbc, 0x43, 0x5f, 0x00, 0x12, 0xea, 0x03, 0x11, 0x96, 0x08, 0x00, 0x45, 0x00,  //|....C_........E.|
            /*00000010*/  0x01, 0xa9, 0x00, 0x00, 0x40, 0x00, 0x40, 0x11, 0x88, 0xef, 0xac, 0x10, 0x02, 0x0a, 0xac, 0x10,  //|....@.@.........|
            /*00000020*/  0x56, 0x2a, 0xba, 0xc0, 0xba, 0xc0, 0x01, 0x95, 0xeb, 0x4c, 0x81, 0x0a, 0x01, 0x8d, 0x01, 0x00,  //|V*.......L......|
            /*00000030*/  0x30, 0x17, 0x0e, 0x0c, 0x00, 0x00, 0x00, 0x01, 0x1e, 0x29, 0x55, 0x4e, 0x44, 0x00, 0x00, 0x00,  //|0........)UND...|
            /*00000040*/  0x00, 0x4f, 0x29, 0x24, 0x4e, 0x91, 0x00, 0x4f, 0x29, 0x6f, 0x4e, 0x82, 0x04, 0x10, 0x4f, 0x29,  //|.O)$N..O)oN...O)|
            /*00000050*/  0x67, 0x4e, 0x91, 0x00, 0x4f, 0x29, 0x00, 0x4e, 0x82, 0x05, 0xe0, 0x4f, 0x29, 0x82, 0x4e, 0x2e,  //|gN..O).N...O).N.|
            /*00000060*/  0xa4, 0xff, 0xff, 0xff, 0xff, 0xb4, 0xff, 0xff, 0xff, 0xff, 0x2f, 0x2e, 0xa4, 0xff, 0xff, 0xff,  //|........../.....|
            /*00000070*/  0xff, 0xb4, 0xff, 0xff, 0xff, 0xff, 0x2f, 0x2e, 0xa4, 0xff, 0xff, 0xff, 0xff, 0xb4, 0xff, 0xff,  //|....../.........|
            /*00000080*/  0xff, 0xff, 0x2f, 0x4f, 0x29, 0x51, 0x4e, 0x11, 0x4f, 0x29, 0x4d, 0x4e, 0x75, 0x21, 0x00, 0x46,  //|../O)QN.O)MNu!.F|
            /*00000090*/  0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x20, 0x4f, 0x75, 0x74, 0x64, 0x6f, 0x6f, 0x72, 0x20,  //|acility Outdoor |
            /*000000a0*/  0x41, 0x69, 0x72, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4f,  //|Air TemperatureO|
            /*000000b0*/  0x29, 0x1c, 0x4e, 0x71, 0x00, 0x4f, 0x29, 0x11, 0x4e, 0x21, 0x00, 0x4f, 0x29, 0x48, 0x4e, 0x91,  //|).Nq.O).N!.O)HN.|
            /*000000c0*/  0x01, 0x4f, 0x29, 0x71, 0x4e, 0x21, 0x00, 0x4f, 0x29, 0x23, 0x4e, 0x82, 0x05, 0x00, 0x4f, 0x29,  //|.O)qN!.O)#N...O)|
            /*000000d0*/  0x45, 0x4e, 0x44, 0xc2, 0x20, 0x02, 0x46, 0x4f, 0x29, 0x41, 0x4e, 0x44, 0x42, 0x6f, 0xfd, 0xb6,  //|END. .FO)ANDBo..|
            /*000000e0*/  0x4f, 0x29, 0x16, 0x4e, 0x44, 0x00, 0x00, 0x00, 0x00, 0x4f, 0x29, 0x3b, 0x4e, 0x44, 0xfe, 0xff,  //|O).ND....O);ND..|
            /*000000f0*/  0xc9, 0x9e, 0x4f, 0x29, 0x2d, 0x4e, 0x44, 0x7e, 0xff, 0xc9, 0x9e, 0x4f, 0x29, 0x19, 0x4e, 0x44,  //|..O)-ND~...O).ND|
            /*00000100*/  0x3d, 0xcc, 0xcc, 0xcd, 0x4f, 0x29, 0x34, 0x4e, 0x82, 0x06, 0x00, 0x4f, 0x2a, 0x52, 0x08, 0x4e,  //|=...O)4N...O*R.N|
            /*00000110*/  0x71, 0x00, 0x4f, 0x2a, 0x52, 0x09, 0x4e, 0x71, 0x00, 0x4f, 0x2a, 0x52, 0x10, 0x4e, 0x91, 0x01,  //|q.O*R.Nq.O*R.N..|
            /*00000120*/  0x4f, 0x2a, 0x52, 0x0a, 0x4e, 0x82, 0x04, 0xf0, 0x4f, 0x2a, 0x52, 0x0c, 0x4e, 0x71, 0x00, 0x4f,  //|O*R.N...O*R.Nq.O|
            /*00000130*/  0x2a, 0x52, 0x0d, 0x4e, 0x21, 0x1e, 0x4f, 0x2a, 0x52, 0x0e, 0x4e, 0x44, 0x3f, 0x80, 0x00, 0x00,  //|*R.N!.O*R.ND?...|
            /*00000140*/  0x4f, 0x2a, 0x52, 0x0f, 0x4e, 0x44, 0x00, 0x00, 0x00, 0x00, 0x4f, 0x2a, 0x52, 0x20, 0x4e, 0x92,  //|O*R.ND....O*R N.|
            /*00000150*/  0x52, 0x08, 0x92, 0x52, 0x09, 0x92, 0x52, 0x10, 0x92, 0x52, 0x0a, 0x92, 0x52, 0x0c, 0x92, 0x52,  //|R..R..R..R..R..R|
            /*00000160*/  0x0d, 0x92, 0x52, 0x0e, 0x92, 0x52, 0x0f, 0x92, 0x52, 0x20, 0x92, 0x52, 0x21, 0x4f, 0x2a, 0x52,  //|..R..R..R .R!O*R|
            /*00000170*/  0x21, 0x4e, 0x91, 0x55, 0x91, 0x24, 0x91, 0x6f, 0x91, 0x67, 0x91, 0x00, 0x91, 0x82, 0x91, 0x51,  //|!N.U.$.o.g.....Q|
            /*00000180*/  0x91, 0x4d, 0x91, 0x1c, 0x91, 0x11, 0x91, 0x48, 0x91, 0x71, 0x91, 0x23, 0x91, 0x45, 0x91, 0x41,  //|.M.....H.q.#.E.A|
            /*00000190*/  0x91, 0x16, 0x91, 0x3b, 0x91, 0x2d, 0x91, 0x19, 0x91, 0x34, 0x91, 0x4b, 0x91, 0x4f, 0x91, 0x75,  //|...;.-...4.K.O.u|
            /*000001a0*/  0x4f, 0x29, 0x4b, 0x4e, 0xc4, 0x00, 0x00, 0x00, 0x01, 0x4f, 0x29, 0x4f, 0x4e, 0x91, 0x00, 0x4f,  //|O)KN.....O)ON..O|
            /*000001b0*/  0x29, 0x75, 0x4e, 0x91, 0x3e, 0x4f, 0x1f,                                                        //|)uN.>O.|
        };
        tryParseBytes(rawBytesAsInts);
    }

    // from plugfest-2011-siemens-1.pcap
    @Test
    public void subscribeCOV() throws Exception {
        int[] rawBytesAsInts = new int[]{
            /*00000000*/  0x00, 0x1a, 0x23, 0x03, 0x00, 0x61, 0x00, 0x26, 0xb9, 0xdd, 0x2a, 0x20, 0x08, 0x00, 0x45, 0x00,  //|..#..a.&..* ..E.|
            /*00000010*/  0x00, 0x3c, 0x0d, 0x16, 0x00, 0x00, 0x80, 0x11, 0x77, 0xb0, 0xac, 0x10, 0x07, 0xc9, 0xac, 0x10,  //|.<......w.......|
            /*00000020*/  0x56, 0x01, 0xba, 0xc0, 0xba, 0xc0, 0x00, 0x28, 0x6e, 0x39, 0x81, 0x0a, 0x00, 0x20, 0x01, 0x2c,  //|V......(n9... .,|
            /*00000030*/  0x21, 0x99, 0x01, 0x3f, 0x02, 0xbf, 0x06, 0xac, 0x10, 0x07, 0xc9, 0xba, 0xc0, 0xfe, 0x02, 0x55,  //|!..?...........U|
            /*00000040*/  0xb7, 0x05, 0x0a, 0x02, 0x58, 0x1c, 0x00, 0xc0, 0x00, 0x65,                                      //|....X....e|
        };
        tryParseBytes(rawBytesAsInts);
    }

    // from plugfest-tridium-1.pcap
    @Test
    public void readPropertyMultiple() throws Exception {
        int[] rawBytesAsInts = new int[]{
            /*00000000*/  0x00, 0x1d, 0x09, 0xbc, 0x43, 0x5f, 0x00, 0x01, 0xf0, 0x8c, 0x11, 0x50, 0x08, 0x00, 0x45, 0x00,  //|....C_.....P..E.|
            /*00000010*/  0x01, 0x13, 0x7e, 0x1f, 0x00, 0x00, 0x40, 0x11, 0x28, 0xa3, 0xac, 0x10, 0x24, 0xcd, 0xac, 0x10,  //|..~...@.(...$...|
            /*00000020*/  0x56, 0x2a, 0xba, 0xc0, 0xba, 0xc0, 0x00, 0xff, 0xa3, 0xe8, 0x81, 0x0a, 0x00, 0xf7, 0x01, 0x00,  //|V*..............|
            /*00000030*/  0x30, 0x49, 0x0e, 0x0c, 0x03, 0x00, 0x00, 0x00, 0x1e, 0x29, 0x4b, 0x4e, 0xc4, 0x03, 0x00, 0x00,  //|0I.......)KN....|
            /*00000040*/  0x00, 0x4f, 0x29, 0x4d, 0x4e, 0x75, 0x16, 0x00, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x4c,  //|.O)MNu..Points.L|
            /*00000050*/  0x6f, 0x6f, 0x70, 0x2e, 0x4c, 0x6f, 0x6f, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x29, 0x4f,  //|oop.LoopPointO)O|
            /*00000060*/  0x4e, 0x91, 0x0c, 0x4f, 0x29, 0x55, 0x4e, 0x44, 0x00, 0x00, 0x00, 0x00, 0x4f, 0x29, 0x6f, 0x4e,  //|N..O)UND....O)oN|
            /*00000070*/  0x82, 0x04, 0x00, 0x4f, 0x29, 0x24, 0x4e, 0x91, 0x00, 0x4f, 0x29, 0x51, 0x4e, 0x10, 0x4f, 0x29,  //|...O)$N..O)QN.O)|
            /*00000080*/  0x52, 0x4e, 0x91, 0x5f, 0x4f, 0x29, 0x3c, 0x4e, 0x0c, 0x00, 0x3f, 0xff, 0xff, 0x19, 0x55, 0x4f,  //|RN._O)<N..?...UO|
            /*00000090*/  0x29, 0x13, 0x4e, 0x0c, 0x00, 0x3f, 0xff, 0xff, 0x19, 0x55, 0x4f, 0x29, 0x15, 0x4e, 0x44, 0x00,  //|).N..?...UO).ND.|
            /*000000a0*/  0x00, 0x00, 0x00, 0x4f, 0x29, 0x14, 0x4e, 0x91, 0x5f, 0x4f, 0x29, 0x6d, 0x4e, 0x4f, 0x29, 0x6c,  //|...O).N._O)mNO)l|
            /*000000b0*/  0x4e, 0x44, 0x00, 0x00, 0x00, 0x00, 0x4f, 0x29, 0x02, 0x4e, 0x91, 0x00, 0x4f, 0x29, 0x58, 0x4e,  //|ND....O).N..O)XN|
            /*000000c0*/  0x21, 0x10, 0x4f, 0x29, 0x1c, 0x4e, 0x71, 0x00, 0x4f, 0x29, 0x67, 0x4e, 0x91, 0x00, 0x4f, 0x29,  //|!.O).Nq.O)gN..O)|
            /*000000d0*/  0x5d, 0x4e, 0x44, 0x00, 0x00, 0x00, 0x00, 0x4f, 0x29, 0x5e, 0x4e, 0x91, 0x5f, 0x4f, 0x29, 0x31,  //|]ND....O)^N._O)1|
            /*000000e0*/  0x4e, 0x44, 0x00, 0x00, 0x00, 0x00, 0x4f, 0x29, 0x32, 0x4e, 0x91, 0x64, 0x4f, 0x29, 0x1a, 0x4e,  //|ND....O)2N.dO).N|
            /*000000f0*/  0x44, 0x00, 0x00, 0x00, 0x00, 0x4f, 0x29, 0x1b, 0x4e, 0x91, 0x49, 0x4f, 0x29, 0x0e, 0x4e, 0x44,  //|D....O).N.IO).ND|
            /*00000100*/  0x00, 0x00, 0x00, 0x00, 0x4f, 0x29, 0x3d, 0x4e, 0x44, 0x42, 0xc8, 0x00, 0x00, 0x4f, 0x29, 0x44,  //|....O)=NDB...O)D|
            /*00000110*/  0x4e, 0x44, 0x00, 0x00, 0x00, 0x00, 0x4f, 0x29, 0x16, 0x4e, 0x44, 0x00, 0x00, 0x00, 0x00, 0x4f,  //|ND....O).ND....O|
            /*00000120*/  0x1f,                                                                                            //|.|
        };
        tryParseBytes(rawBytesAsInts);
    }

}
