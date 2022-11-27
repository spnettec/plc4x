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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class NLMVendorProprietaryMessage extends NLM implements Message {

  // Accessors for discriminator values.
  public Short getMessageType() {
    return 0;
  }

  // Properties.
  protected final BACnetVendorId vendorId;
  protected final byte[] proprietaryMessage;

  // Arguments.
  protected final Integer apduLength;

  public NLMVendorProprietaryMessage(
      BACnetVendorId vendorId, byte[] proprietaryMessage, Integer apduLength) {
    super(apduLength);
    this.vendorId = vendorId;
    this.proprietaryMessage = proprietaryMessage;
    this.apduLength = apduLength;
  }

  public BACnetVendorId getVendorId() {
    return vendorId;
  }

  public byte[] getProprietaryMessage() {
    return proprietaryMessage;
  }

  @Override
  protected void serializeNLMChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("NLMVendorProprietaryMessage");

    // Simple Field (vendorId)
    writeSimpleEnumField(
        "vendorId",
        "BACnetVendorId",
        vendorId,
        new DataWriterEnumDefault<>(
            BACnetVendorId::getValue, BACnetVendorId::name, writeUnsignedInt(writeBuffer, 16)));

    // Array Field (proprietaryMessage)
    writeByteArrayField("proprietaryMessage", proprietaryMessage, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("NLMVendorProprietaryMessage");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    NLMVendorProprietaryMessage _value = this;

    // Simple field (vendorId)
    lengthInBits += 16;

    // Array field
    if (proprietaryMessage != null) {
      lengthInBits += 8 * proprietaryMessage.length;
    }

    return lengthInBits;
  }

  public static NLMVendorProprietaryMessageBuilder staticParseBuilder(
      ReadBuffer readBuffer, Integer apduLength) throws ParseException {
    readBuffer.pullContext("NLMVendorProprietaryMessage");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetVendorId vendorId =
        readEnumField(
            "vendorId",
            "BACnetVendorId",
            new DataReaderEnumDefault<>(
                BACnetVendorId::enumForValue, readUnsignedInt(readBuffer, 16)));

    byte[] proprietaryMessage =
        readBuffer.readByteArray(
            "proprietaryMessage",
            Math.toIntExact(((((apduLength) > (0))) ? ((apduLength) - (3)) : 0)));

    readBuffer.closeContext("NLMVendorProprietaryMessage");
    // Create the instance
    return new NLMVendorProprietaryMessageBuilder(vendorId, proprietaryMessage, apduLength);
  }

  public static class NLMVendorProprietaryMessageBuilder implements NLM.NLMBuilder {
    private final BACnetVendorId vendorId;
    private final byte[] proprietaryMessage;
    private final Integer apduLength;

    public NLMVendorProprietaryMessageBuilder(
        BACnetVendorId vendorId, byte[] proprietaryMessage, Integer apduLength) {

      this.vendorId = vendorId;
      this.proprietaryMessage = proprietaryMessage;
      this.apduLength = apduLength;
    }

    public NLMVendorProprietaryMessage build(Integer apduLength) {

      NLMVendorProprietaryMessage nLMVendorProprietaryMessage =
          new NLMVendorProprietaryMessage(vendorId, proprietaryMessage, apduLength);
      return nLMVendorProprietaryMessage;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NLMVendorProprietaryMessage)) {
      return false;
    }
    NLMVendorProprietaryMessage that = (NLMVendorProprietaryMessage) o;
    return (getVendorId() == that.getVendorId())
        && (getProprietaryMessage() == that.getProprietaryMessage())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getVendorId(), getProprietaryMessage());
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
