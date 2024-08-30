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

public class AdsReadDeviceInfoResponse extends AmsPacket implements Message {

  // Accessors for discriminator values.
  public CommandId getCommandId() {
    return CommandId.ADS_READ_DEVICE_INFO;
  }

  public Boolean getResponse() {
    return (boolean) true;
  }

  // Properties.
  protected final ReturnCode result;
  protected final short majorVersion;
  protected final short minorVersion;
  protected final int version;
  protected final byte[] device;

  public AdsReadDeviceInfoResponse(
      AmsNetId targetAmsNetId,
      int targetAmsPort,
      AmsNetId sourceAmsNetId,
      int sourceAmsPort,
      long errorCode,
      long invokeId,
      ReturnCode result,
      short majorVersion,
      short minorVersion,
      int version,
      byte[] device) {
    super(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId);
    this.result = result;
    this.majorVersion = majorVersion;
    this.minorVersion = minorVersion;
    this.version = version;
    this.device = device;
  }

  public ReturnCode getResult() {
    return result;
  }

  public short getMajorVersion() {
    return majorVersion;
  }

  public short getMinorVersion() {
    return minorVersion;
  }

  public int getVersion() {
    return version;
  }

  public byte[] getDevice() {
    return device;
  }

  @Override
  protected void serializeAmsPacketChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AdsReadDeviceInfoResponse");

    // Simple Field (result)
    writeSimpleEnumField(
        "result",
        "ReturnCode",
        result,
        writeEnum(ReturnCode::getValue, ReturnCode::name, writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (majorVersion)
    writeSimpleField("majorVersion", majorVersion, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (minorVersion)
    writeSimpleField("minorVersion", minorVersion, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (version)
    writeSimpleField("version", version, writeUnsignedInt(writeBuffer, 16));

    // Array Field (device)
    writeByteArrayField("device", device, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("AdsReadDeviceInfoResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AdsReadDeviceInfoResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (result)
    lengthInBits += 32;

    // Simple field (majorVersion)
    lengthInBits += 8;

    // Simple field (minorVersion)
    lengthInBits += 8;

    // Simple field (version)
    lengthInBits += 16;

    // Array field
    if (device != null) {
      lengthInBits += 8 * device.length;
    }

    return lengthInBits;
  }

  public static AmsPacketBuilder staticParseAmsPacketBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("AdsReadDeviceInfoResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ReturnCode result =
        readEnumField(
            "result",
            "ReturnCode",
            readEnum(ReturnCode::enumForValue, readUnsignedLong(readBuffer, 32)));

    short majorVersion = readSimpleField("majorVersion", readUnsignedShort(readBuffer, 8));

    short minorVersion = readSimpleField("minorVersion", readUnsignedShort(readBuffer, 8));

    int version = readSimpleField("version", readUnsignedInt(readBuffer, 16));

    byte[] device = readBuffer.readByteArray("device", Math.toIntExact(16));

    readBuffer.closeContext("AdsReadDeviceInfoResponse");
    // Create the instance
    return new AdsReadDeviceInfoResponseBuilderImpl(
        result, majorVersion, minorVersion, version, device);
  }

  public static class AdsReadDeviceInfoResponseBuilderImpl implements AmsPacket.AmsPacketBuilder {
    private final ReturnCode result;
    private final short majorVersion;
    private final short minorVersion;
    private final int version;
    private final byte[] device;

    public AdsReadDeviceInfoResponseBuilderImpl(
        ReturnCode result, short majorVersion, short minorVersion, int version, byte[] device) {
      this.result = result;
      this.majorVersion = majorVersion;
      this.minorVersion = minorVersion;
      this.version = version;
      this.device = device;
    }

    public AdsReadDeviceInfoResponse build(
        AmsNetId targetAmsNetId,
        int targetAmsPort,
        AmsNetId sourceAmsNetId,
        int sourceAmsPort,
        long errorCode,
        long invokeId) {
      AdsReadDeviceInfoResponse adsReadDeviceInfoResponse =
          new AdsReadDeviceInfoResponse(
              targetAmsNetId,
              targetAmsPort,
              sourceAmsNetId,
              sourceAmsPort,
              errorCode,
              invokeId,
              result,
              majorVersion,
              minorVersion,
              version,
              device);
      return adsReadDeviceInfoResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AdsReadDeviceInfoResponse)) {
      return false;
    }
    AdsReadDeviceInfoResponse that = (AdsReadDeviceInfoResponse) o;
    return (getResult() == that.getResult())
        && (getMajorVersion() == that.getMajorVersion())
        && (getMinorVersion() == that.getMinorVersion())
        && (getVersion() == that.getVersion())
        && (getDevice() == that.getDevice())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getResult(),
        getMajorVersion(),
        getMinorVersion(),
        getVersion(),
        getDevice());
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
