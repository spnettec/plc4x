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
package org.apache.plc4x.java.eip.readwrite;

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

public class CipWriteResponse extends CipService implements Message {

  // Accessors for discriminator values.
  public Short getService() {
    return (short) 0x4D;
  }

  public Boolean getResponse() {
    return (boolean) true;
  }

  public Boolean getConnected() {
    return false;
  }

  // Properties.
  protected final short status;
  protected final short extStatus;

  public CipWriteResponse(short status, short extStatus) {
    super();
    this.status = status;
    this.extStatus = extStatus;
  }

  public short getStatus() {
    return status;
  }

  public short getExtStatus() {
    return extStatus;
  }

  @Override
  protected void serializeCipServiceChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CipWriteResponse");

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (status)
    writeSimpleField("status", status, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (extStatus)
    writeSimpleField("extStatus", extStatus, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("CipWriteResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CipWriteResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (status)
    lengthInBits += 8;

    // Simple field (extStatus)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static CipServiceBuilder staticParseCipServiceBuilder(
      ReadBuffer readBuffer, Boolean connected, Integer serviceLen) throws ParseException {
    readBuffer.pullContext("CipWriteResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 8), (short) 0x00);

    short status = readSimpleField("status", readUnsignedShort(readBuffer, 8));

    short extStatus = readSimpleField("extStatus", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("CipWriteResponse");
    // Create the instance
    return new CipWriteResponseBuilderImpl(status, extStatus);
  }

  public static class CipWriteResponseBuilderImpl implements CipService.CipServiceBuilder {
    private final short status;
    private final short extStatus;

    public CipWriteResponseBuilderImpl(short status, short extStatus) {
      this.status = status;
      this.extStatus = extStatus;
    }

    public CipWriteResponse build() {
      CipWriteResponse cipWriteResponse = new CipWriteResponse(status, extStatus);
      return cipWriteResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CipWriteResponse)) {
      return false;
    }
    CipWriteResponse that = (CipWriteResponse) o;
    return (getStatus() == that.getStatus())
        && (getExtStatus() == that.getExtStatus())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getStatus(), getExtStatus());
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
