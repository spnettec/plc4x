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

public class BACnetChannelValueBoolean extends BACnetChannelValue implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagBoolean booleanValue;

  public BACnetChannelValueBoolean(
      BACnetTagHeader peekedTagHeader, BACnetApplicationTagBoolean booleanValue) {
    super(peekedTagHeader);
    this.booleanValue = booleanValue;
  }

  public BACnetApplicationTagBoolean getBooleanValue() {
    return booleanValue;
  }

  @Override
  protected void serializeBACnetChannelValueChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetChannelValueBoolean");

    // Simple Field (booleanValue)
    writeSimpleField("booleanValue", booleanValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetChannelValueBoolean");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetChannelValueBoolean _value = this;

    // Simple field (booleanValue)
    lengthInBits += booleanValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetChannelValueBooleanBuilder staticParseBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("BACnetChannelValueBoolean");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetApplicationTagBoolean booleanValue =
        readSimpleField(
            "booleanValue",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagBoolean) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetChannelValueBoolean");
    // Create the instance
    return new BACnetChannelValueBooleanBuilder(booleanValue);
  }

  public static class BACnetChannelValueBooleanBuilder
      implements BACnetChannelValue.BACnetChannelValueBuilder {
    private final BACnetApplicationTagBoolean booleanValue;

    public BACnetChannelValueBooleanBuilder(BACnetApplicationTagBoolean booleanValue) {

      this.booleanValue = booleanValue;
    }

    public BACnetChannelValueBoolean build(BACnetTagHeader peekedTagHeader) {
      BACnetChannelValueBoolean bACnetChannelValueBoolean =
          new BACnetChannelValueBoolean(peekedTagHeader, booleanValue);
      return bACnetChannelValueBoolean;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetChannelValueBoolean)) {
      return false;
    }
    BACnetChannelValueBoolean that = (BACnetChannelValueBoolean) o;
    return (getBooleanValue() == that.getBooleanValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getBooleanValue());
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
