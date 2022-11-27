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
package org.apache.plc4x.java.ads.readwrite;

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

public class AmsSerialAcknowledgeFrame implements Message {

  // Properties.
  protected final int magicCookie;
  protected final byte transmitterAddress;
  protected final byte receiverAddress;
  protected final byte fragmentNumber;
  protected final byte length;
  protected final int crc;

  public AmsSerialAcknowledgeFrame(
      int magicCookie,
      byte transmitterAddress,
      byte receiverAddress,
      byte fragmentNumber,
      byte length,
      int crc) {
    super();
    this.magicCookie = magicCookie;
    this.transmitterAddress = transmitterAddress;
    this.receiverAddress = receiverAddress;
    this.fragmentNumber = fragmentNumber;
    this.length = length;
    this.crc = crc;
  }

  public int getMagicCookie() {
    return magicCookie;
  }

  public byte getTransmitterAddress() {
    return transmitterAddress;
  }

  public byte getReceiverAddress() {
    return receiverAddress;
  }

  public byte getFragmentNumber() {
    return fragmentNumber;
  }

  public byte getLength() {
    return length;
  }

  public int getCrc() {
    return crc;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AmsSerialAcknowledgeFrame");

    // Simple Field (magicCookie)
    writeSimpleField("magicCookie", magicCookie, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (transmitterAddress)
    writeSimpleField("transmitterAddress", transmitterAddress, writeSignedByte(writeBuffer, 8));

    // Simple Field (receiverAddress)
    writeSimpleField("receiverAddress", receiverAddress, writeSignedByte(writeBuffer, 8));

    // Simple Field (fragmentNumber)
    writeSimpleField("fragmentNumber", fragmentNumber, writeSignedByte(writeBuffer, 8));

    // Simple Field (length)
    writeSimpleField("length", length, writeSignedByte(writeBuffer, 8));

    // Simple Field (crc)
    writeSimpleField("crc", crc, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("AmsSerialAcknowledgeFrame");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    AmsSerialAcknowledgeFrame _value = this;

    // Simple field (magicCookie)
    lengthInBits += 16;

    // Simple field (transmitterAddress)
    lengthInBits += 8;

    // Simple field (receiverAddress)
    lengthInBits += 8;

    // Simple field (fragmentNumber)
    lengthInBits += 8;

    // Simple field (length)
    lengthInBits += 8;

    // Simple field (crc)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static AmsSerialAcknowledgeFrame staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static AmsSerialAcknowledgeFrame staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("AmsSerialAcknowledgeFrame");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    int magicCookie = readSimpleField("magicCookie", readUnsignedInt(readBuffer, 16));

    byte transmitterAddress = readSimpleField("transmitterAddress", readSignedByte(readBuffer, 8));

    byte receiverAddress = readSimpleField("receiverAddress", readSignedByte(readBuffer, 8));

    byte fragmentNumber = readSimpleField("fragmentNumber", readSignedByte(readBuffer, 8));

    byte length = readSimpleField("length", readSignedByte(readBuffer, 8));

    int crc = readSimpleField("crc", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("AmsSerialAcknowledgeFrame");
    // Create the instance
    AmsSerialAcknowledgeFrame _amsSerialAcknowledgeFrame;
    _amsSerialAcknowledgeFrame =
        new AmsSerialAcknowledgeFrame(
            magicCookie, transmitterAddress, receiverAddress, fragmentNumber, length, crc);
    return _amsSerialAcknowledgeFrame;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AmsSerialAcknowledgeFrame)) {
      return false;
    }
    AmsSerialAcknowledgeFrame that = (AmsSerialAcknowledgeFrame) o;
    return (getMagicCookie() == that.getMagicCookie())
        && (getTransmitterAddress() == that.getTransmitterAddress())
        && (getReceiverAddress() == that.getReceiverAddress())
        && (getFragmentNumber() == that.getFragmentNumber())
        && (getLength() == that.getLength())
        && (getCrc() == that.getCrc())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getMagicCookie(),
        getTransmitterAddress(),
        getReceiverAddress(),
        getFragmentNumber(),
        getLength(),
        getCrc());
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
