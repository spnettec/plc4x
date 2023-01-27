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

public class BACnetOptionalREALValue extends BACnetOptionalREAL implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagReal realValue;

  public BACnetOptionalREALValue(
      BACnetTagHeader peekedTagHeader, BACnetApplicationTagReal realValue) {
    super(peekedTagHeader);
    this.realValue = realValue;
  }

  public BACnetApplicationTagReal getRealValue() {
    return realValue;
  }

  @Override
  protected void serializeBACnetOptionalREALChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetOptionalREALValue");

    // Simple Field (realValue)
    writeSimpleField("realValue", realValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetOptionalREALValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetOptionalREALValue _value = this;

    // Simple field (realValue)
    lengthInBits += realValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetOptionalREALBuilder staticParseBACnetOptionalREALBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetOptionalREALValue");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetApplicationTagReal realValue =
        readSimpleField(
            "realValue",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagReal) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetOptionalREALValue");
    // Create the instance
    return new BACnetOptionalREALValueBuilderImpl(realValue);
  }

  public static class BACnetOptionalREALValueBuilderImpl
      implements BACnetOptionalREAL.BACnetOptionalREALBuilder {
    private final BACnetApplicationTagReal realValue;

    public BACnetOptionalREALValueBuilderImpl(BACnetApplicationTagReal realValue) {
      this.realValue = realValue;
    }

    public BACnetOptionalREALValue build(BACnetTagHeader peekedTagHeader) {
      BACnetOptionalREALValue bACnetOptionalREALValue =
          new BACnetOptionalREALValue(peekedTagHeader, realValue);
      return bACnetOptionalREALValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetOptionalREALValue)) {
      return false;
    }
    BACnetOptionalREALValue that = (BACnetOptionalREALValue) o;
    return (getRealValue() == that.getRealValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRealValue());
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
