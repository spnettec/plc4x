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
package org.apache.plc4x.java.profinet.readwrite;

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

public class TlvChassisId extends LldpUnit implements Message {

  // Accessors for discriminator values.
  public TlvType getTlvId() {
    return TlvType.CHASSIS_ID;
  }

  // Properties.
  protected final short chassisIdSubType;
  protected final String chassisId;

  public TlvChassisId(int tlvIdLength, short chassisIdSubType, String chassisId) {
    super(tlvIdLength);
    this.chassisIdSubType = chassisIdSubType;
    this.chassisId = chassisId;
  }

  public short getChassisIdSubType() {
    return chassisIdSubType;
  }

  public String getChassisId() {
    return chassisId;
  }

  @Override
  protected void serializeLldpUnitChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("TlvChassisId");

    // Simple Field (chassisIdSubType)
    writeSimpleField("chassisIdSubType", chassisIdSubType, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (chassisId)
    writeSimpleField(
        "chassisId", chassisId, writeString(writeBuffer, (((tlvIdLength) - (1))) * (8)));

    writeBuffer.popContext("TlvChassisId");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TlvChassisId _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (chassisIdSubType)
    lengthInBits += 8;

    // Simple field (chassisId)
    lengthInBits += (((tlvIdLength) - (1))) * (8);

    return lengthInBits;
  }

  public static LldpUnitBuilder staticParseLldpUnitBuilder(
      ReadBuffer readBuffer, Integer tlvIdLength) throws ParseException {
    readBuffer.pullContext("TlvChassisId");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short chassisIdSubType = readSimpleField("chassisIdSubType", readUnsignedShort(readBuffer, 8));

    String chassisId =
        readSimpleField("chassisId", readString(readBuffer, (((tlvIdLength) - (1))) * (8)));

    readBuffer.closeContext("TlvChassisId");
    // Create the instance
    return new TlvChassisIdBuilderImpl(chassisIdSubType, chassisId);
  }

  public static class TlvChassisIdBuilderImpl implements LldpUnit.LldpUnitBuilder {
    private final short chassisIdSubType;
    private final String chassisId;

    public TlvChassisIdBuilderImpl(short chassisIdSubType, String chassisId) {
      this.chassisIdSubType = chassisIdSubType;
      this.chassisId = chassisId;
    }

    public TlvChassisId build(int tlvIdLength) {
      TlvChassisId tlvChassisId = new TlvChassisId(tlvIdLength, chassisIdSubType, chassisId);
      return tlvChassisId;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TlvChassisId)) {
      return false;
    }
    TlvChassisId that = (TlvChassisId) o;
    return (getChassisIdSubType() == that.getChassisIdSubType())
        && (getChassisId() == that.getChassisId())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getChassisIdSubType(), getChassisId());
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
