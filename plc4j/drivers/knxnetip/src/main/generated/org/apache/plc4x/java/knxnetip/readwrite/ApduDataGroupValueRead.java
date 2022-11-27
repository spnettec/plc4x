/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.knxnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class ApduDataGroupValueRead extends ApduData implements Message {

  // Accessors for discriminator values.
  public Byte getApciType() {
    return (byte) 0x0;
  }

  // Arguments.
  protected final Short dataLength;
  // Reserved Fields
  private Short reservedField0;

  public ApduDataGroupValueRead(Short dataLength) {
    super(dataLength);
    this.dataLength = dataLength;
  }

  @Override
  protected void serializeApduDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ApduDataGroupValueRead");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (short) 0x00,
        writeUnsignedShort(writeBuffer, 6));

    writeBuffer.popContext("ApduDataGroupValueRead");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ApduDataGroupValueRead _value = this;

    // Reserved Field (reserved)
    lengthInBits += 6;

    return lengthInBits;
  }

  public static ApduDataGroupValueReadBuilder staticParseBuilder(
      ReadBuffer readBuffer, Short dataLength) throws ParseException {
    readBuffer.pullContext("ApduDataGroupValueRead");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 6), (short) 0x00);

    readBuffer.closeContext("ApduDataGroupValueRead");
    // Create the instance
    return new ApduDataGroupValueReadBuilder(dataLength, reservedField0);
  }

  public static class ApduDataGroupValueReadBuilder implements ApduData.ApduDataBuilder {
    private final Short dataLength;
    private final Short reservedField0;

    public ApduDataGroupValueReadBuilder(Short dataLength, Short reservedField0) {
      this.dataLength = dataLength;
      this.reservedField0 = reservedField0;
    }

    public ApduDataGroupValueRead build(Short dataLength) {

      ApduDataGroupValueRead apduDataGroupValueRead = new ApduDataGroupValueRead(dataLength);

      apduDataGroupValueRead.reservedField0 = reservedField0;
      return apduDataGroupValueRead;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ApduDataGroupValueRead)) {
      return false;
    }
    ApduDataGroupValueRead that = (ApduDataGroupValueRead) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
