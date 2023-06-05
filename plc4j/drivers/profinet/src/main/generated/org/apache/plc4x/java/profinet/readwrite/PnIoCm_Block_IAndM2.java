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

public class PnIoCm_Block_IAndM2 extends PnIoCm_Block implements Message {

  // Accessors for discriminator values.
  public PnIoCm_BlockType getBlockType() {
    return PnIoCm_BlockType.I_AND_M_2;
  }

  // Properties.
  protected final short blockVersionHigh;
  protected final short blockVersionLow;
  protected final String installationDate;

  public PnIoCm_Block_IAndM2(
      short blockVersionHigh, short blockVersionLow, String installationDate) {
    super();
    this.blockVersionHigh = blockVersionHigh;
    this.blockVersionLow = blockVersionLow;
    this.installationDate = installationDate;
  }

  public short getBlockVersionHigh() {
    return blockVersionHigh;
  }

  public short getBlockVersionLow() {
    return blockVersionLow;
  }

  public String getInstallationDate() {
    return installationDate;
  }

  @Override
  protected void serializePnIoCm_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PnIoCm_Block_IAndM2");

    // Implicit Field (blockLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int blockLength = (int) ((getLengthInBytes()) - (4));
    writeImplicitField(
        "blockLength",
        blockLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionHigh)
    writeSimpleField(
        "blockVersionHigh",
        blockVersionHigh,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionLow)
    writeSimpleField(
        "blockVersionLow",
        blockVersionLow,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (installationDate)
    writeSimpleField(
        "installationDate",
        installationDate,
        writeString(writeBuffer, 128),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PnIoCm_Block_IAndM2");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnIoCm_Block_IAndM2 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (blockLength)
    lengthInBits += 16;

    // Simple field (blockVersionHigh)
    lengthInBits += 8;

    // Simple field (blockVersionLow)
    lengthInBits += 8;

    // Simple field (installationDate)
    lengthInBits += 128;

    return lengthInBits;
  }

  public static PnIoCm_BlockBuilder staticParsePnIoCm_BlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnIoCm_Block_IAndM2");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int blockLength =
        readImplicitField(
            "blockLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionHigh =
        readSimpleField(
            "blockVersionHigh",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionLow =
        readSimpleField(
            "blockVersionLow",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    String installationDate =
        readSimpleField(
            "installationDate",
            readString(readBuffer, 128),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PnIoCm_Block_IAndM2");
    // Create the instance
    return new PnIoCm_Block_IAndM2BuilderImpl(blockVersionHigh, blockVersionLow, installationDate);
  }

  public static class PnIoCm_Block_IAndM2BuilderImpl implements PnIoCm_Block.PnIoCm_BlockBuilder {
    private final short blockVersionHigh;
    private final short blockVersionLow;
    private final String installationDate;

    public PnIoCm_Block_IAndM2BuilderImpl(
        short blockVersionHigh, short blockVersionLow, String installationDate) {
      this.blockVersionHigh = blockVersionHigh;
      this.blockVersionLow = blockVersionLow;
      this.installationDate = installationDate;
    }

    public PnIoCm_Block_IAndM2 build() {
      PnIoCm_Block_IAndM2 pnIoCm_Block_IAndM2 =
          new PnIoCm_Block_IAndM2(blockVersionHigh, blockVersionLow, installationDate);
      return pnIoCm_Block_IAndM2;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Block_IAndM2)) {
      return false;
    }
    PnIoCm_Block_IAndM2 that = (PnIoCm_Block_IAndM2) o;
    return (getBlockVersionHigh() == that.getBlockVersionHigh())
        && (getBlockVersionLow() == that.getBlockVersionLow())
        && (getInstallationDate() == that.getInstallationDate())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getBlockVersionHigh(), getBlockVersionLow(), getInstallationDate());
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
