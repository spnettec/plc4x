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

import org.apache.commons.codec.binary.Hex;
import org.apache.commons.io.HexDump;
import org.apache.commons.lang3.ArrayUtils;
import org.apache.plc4x.java.bacnetip.readwrite.BVLC;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.ReadBufferByteBased;

import java.util.stream.IntStream;

import static org.apache.plc4x.java.bacnetip.Utils.tryParseBytes;
import static org.junit.jupiter.api.Assertions.assertNotNull;

public class ManualBacNetDecoder {

    public static void main(String[] args) throws Exception {
        int[] rawBytesAsInts = new int[]{
            /*00000000*/ 0x81, 0x0A, 0x01, 0x68, 0x01, 0x00, 0x30, 0x41, 0x0E, 0x0C, 0x05, 0x40, 0x00, 0x01, 0x1E, 0x29, //...h..0A...@...)
            /*00000010*/ 0x4B, 0x4E, 0xC4, 0x05, 0x40, 0x00, 0x01, 0x4F, 0x29, 0x4D, 0x4E, 0x74, 0x00, 0x4C, 0x50, 0x31, //KN..@..O)MNt.LP1
            /*00000020*/ 0x4F, 0x29, 0x4F, 0x4E, 0x91, 0x15, 0x4F, 0x29, 0x55, 0x4E, 0x91, 0x00, 0x4F, 0x29, 0x6F, 0x4E, //O)ON..O)UN..O)oN
            /*00000030*/ 0x82, 0x04, 0x00, 0x4F, 0x29, 0x24, 0x4E, 0x91, 0x00, 0x4F, 0x29, 0x67, 0x4E, 0x91, 0x00, 0x4F, //...O)$N..O)gN..O
            /*00000040*/ 0x29, 0x51, 0x4E, 0x10, 0x4F, 0x29, 0xA0, 0x4E, 0x91, 0x01, 0x4F, 0x29, 0xAF, 0x4E, 0x91, 0x00, //)QN.O).N..O).N..
            /*00000050*/ 0x91, 0x01, 0x91, 0x02, 0x91, 0x03, 0x91, 0x04, 0x91, 0x05, 0x91, 0x06, 0x91, 0x07, 0x91, 0x08, //................
            /*00000060*/ 0x91, 0x09, 0x91, 0x0A, 0x91, 0x0B, 0x91, 0x0C, 0x91, 0x0D, 0x91, 0x0E, 0x4F, 0x29, 0xA3, 0x4E, //............O).N
            /*00000070*/ 0x91, 0x00, 0x4F, 0x29, 0xA1, 0x4E, 0x91, 0x00, 0x4F, 0x29, 0x1C, 0x4E, 0x75, 0x19, 0x00, 0x54, //..O).N..O).Nu..T
            /*00000080*/ 0x65, 0x73, 0x74, 0x20, 0x4C, 0x69, 0x66, 0x65, 0x20, 0x53, 0x61, 0x66, 0x65, 0x74, 0x79, 0x20, //est Life Safety
            /*00000090*/ 0x50, 0x6F, 0x69, 0x6E, 0x74, 0x20, 0x31, 0x4F, 0x29, 0x1F, 0x4E, 0x75, 0x1B, 0x00, 0x53, 0x69, //Point 1O).Nu..Si
            /*000000A0*/ 0x6D, 0x75, 0x6C, 0x61, 0x74, 0x65, 0x64, 0x20, 0x53, 0x65, 0x6E, 0x73, 0x6F, 0x72, 0x20, 0x75, //mulated Sensor u
            /*000000B0*/ 0x73, 0x69, 0x6E, 0x67, 0x20, 0x50, 0x47, 0x37, 0x4F, 0x29, 0xA4, 0x4E, 0x91, 0x00, 0x4F, 0x29, //sing PG7O).N..O)
            /*000000C0*/ 0xA8, 0x4E, 0x75, 0x19, 0x00, 0x30, 0x2D, 0x4F, 0x62, 0x6A, 0x65, 0x63, 0x74, 0x2D, 0x4C, 0x69, //.Nu..0-Object-Li
            /*000000D0*/ 0x66, 0x65, 0x53, 0x61, 0x66, 0x65, 0x74, 0x79, 0x50, 0x6F, 0x69, 0x6E, 0x74, 0x4F, 0x29, 0x9E, //feSafetyPointO).
            /*000000E0*/ 0x4E, 0x91, 0x00, 0x4F, 0x29, 0xA2, 0x4E, 0x21, 0x00, 0x4F, 0x29, 0x9C, 0x4E, 0x44, 0x00, 0x00, //N..O).N!.O).ND..
            /*000000F0*/ 0x00, 0x00, 0x4F, 0x29, 0x75, 0x4E, 0x91, 0x5F, 0x4F, 0x29, 0x9F, 0x4E, 0x0C, 0x02, 0x00, 0x0F, //..O)uN._O).N....
            /*00000100*/ 0xA0, 0x1C, 0x05, 0x80, 0x00, 0x01, 0x1C, 0x05, 0x80, 0x00, 0x01, 0x4F, 0x29, 0x71, 0x4E, 0x21, //...........O)qN!
            /*00000110*/ 0x05, 0x4F, 0x29, 0x11, 0x4E, 0x21, 0x01, 0x4F, 0x29, 0x07, 0x4E, 0x91, 0x02, 0x91, 0x05, 0x91, //.O).N!.O).N.....
            /*00000120*/ 0x09, 0x91, 0x0C, 0x91, 0x0F, 0x91, 0x14, 0x91, 0x15, 0x4F, 0x29, 0x27, 0x4E, 0x91, 0x03, 0x91, //.........O)'N...
            /*00000130*/ 0x0B, 0x4F, 0x29, 0xA6, 0x4E, 0x91, 0x08, 0x91, 0x0D, 0x91, 0x0E, 0x4F, 0x29, 0x23, 0x4E, 0x82, //.O).N......O)#N.
            /*00000140*/ 0x05, 0x00, 0x4F, 0x29, 0x00, 0x4E, 0x82, 0x05, 0xE0, 0x4F, 0x29, 0x48, 0x4E, 0x91, 0x00, 0x4F, //..O).N...O)HN..O
            /*00000150*/ 0x29, 0x82, 0x4E, 0x0C, 0xFF, 0xFF, 0xFF, 0xFF, 0x19, 0x00, 0x2E, 0xA4, 0xFF, 0xFF, 0xFF, 0xFF, //).N.............
            /*00000160*/ 0xB4, 0xFF, 0xFF, 0xFF, 0xFF, 0x2F, 0x4F, 0x1F,                                                 //...../O.
        };
        tryParseBytes(rawBytesAsInts,0, true);
    }

}
