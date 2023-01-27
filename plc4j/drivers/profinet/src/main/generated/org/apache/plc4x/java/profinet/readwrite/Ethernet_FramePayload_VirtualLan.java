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

public class Ethernet_FramePayload_VirtualLan extends Ethernet_FramePayload implements Message {

  // Accessors for discriminator values.
  public Integer getPacketType() {
    return (int) 0x8100;
  }

  // Properties.
  protected final VirtualLanPriority priority;
  protected final boolean ineligible;
  protected final int id;
  protected final Ethernet_FramePayload payload;

  public Ethernet_FramePayload_VirtualLan(
      VirtualLanPriority priority, boolean ineligible, int id, Ethernet_FramePayload payload) {
    super();
    this.priority = priority;
    this.ineligible = ineligible;
    this.id = id;
    this.payload = payload;
  }

  public VirtualLanPriority getPriority() {
    return priority;
  }

  public boolean getIneligible() {
    return ineligible;
  }

  public int getId() {
    return id;
  }

  public Ethernet_FramePayload getPayload() {
    return payload;
  }

  @Override
  protected void serializeEthernet_FramePayloadChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("Ethernet_FramePayload_VirtualLan");

    // Simple Field (priority)
    writeSimpleEnumField(
        "priority",
        "VirtualLanPriority",
        priority,
        new DataWriterEnumDefault<>(
            VirtualLanPriority::getValue,
            VirtualLanPriority::name,
            writeUnsignedByte(writeBuffer, 3)));

    // Simple Field (ineligible)
    writeSimpleField("ineligible", ineligible, writeBoolean(writeBuffer));

    // Simple Field (id)
    writeSimpleField("id", id, writeUnsignedInt(writeBuffer, 12));

    // Simple Field (payload)
    writeSimpleField("payload", payload, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("Ethernet_FramePayload_VirtualLan");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    Ethernet_FramePayload_VirtualLan _value = this;

    // Simple field (priority)
    lengthInBits += 3;

    // Simple field (ineligible)
    lengthInBits += 1;

    // Simple field (id)
    lengthInBits += 12;

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    return lengthInBits;
  }

  public static Ethernet_FramePayloadBuilder staticParseEthernet_FramePayloadBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("Ethernet_FramePayload_VirtualLan");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    VirtualLanPriority priority =
        readEnumField(
            "priority",
            "VirtualLanPriority",
            new DataReaderEnumDefault<>(
                VirtualLanPriority::enumForValue, readUnsignedByte(readBuffer, 3)));

    boolean ineligible = readSimpleField("ineligible", readBoolean(readBuffer));

    int id = readSimpleField("id", readUnsignedInt(readBuffer, 12));

    Ethernet_FramePayload payload =
        readSimpleField(
            "payload",
            new DataReaderComplexDefault<>(
                () -> Ethernet_FramePayload.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("Ethernet_FramePayload_VirtualLan");
    // Create the instance
    return new Ethernet_FramePayload_VirtualLanBuilderImpl(priority, ineligible, id, payload);
  }

  public static class Ethernet_FramePayload_VirtualLanBuilderImpl
      implements Ethernet_FramePayload.Ethernet_FramePayloadBuilder {
    private final VirtualLanPriority priority;
    private final boolean ineligible;
    private final int id;
    private final Ethernet_FramePayload payload;

    public Ethernet_FramePayload_VirtualLanBuilderImpl(
        VirtualLanPriority priority, boolean ineligible, int id, Ethernet_FramePayload payload) {
      this.priority = priority;
      this.ineligible = ineligible;
      this.id = id;
      this.payload = payload;
    }

    public Ethernet_FramePayload_VirtualLan build() {
      Ethernet_FramePayload_VirtualLan ethernet_FramePayload_VirtualLan =
          new Ethernet_FramePayload_VirtualLan(priority, ineligible, id, payload);
      return ethernet_FramePayload_VirtualLan;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Ethernet_FramePayload_VirtualLan)) {
      return false;
    }
    Ethernet_FramePayload_VirtualLan that = (Ethernet_FramePayload_VirtualLan) o;
    return (getPriority() == that.getPriority())
        && (getIneligible() == that.getIneligible())
        && (getId() == that.getId())
        && (getPayload() == that.getPayload())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPriority(), getIneligible(), getId(), getPayload());
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
