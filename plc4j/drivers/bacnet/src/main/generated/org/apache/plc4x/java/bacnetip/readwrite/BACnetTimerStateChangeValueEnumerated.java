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

public class BACnetTimerStateChangeValueEnumerated extends BACnetTimerStateChangeValue
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagEnumerated enumeratedValue;

  // Arguments.
  protected final BACnetObjectType objectTypeArgument;

  public BACnetTimerStateChangeValueEnumerated(
      BACnetTagHeader peekedTagHeader,
      BACnetApplicationTagEnumerated enumeratedValue,
      BACnetObjectType objectTypeArgument) {
    super(peekedTagHeader, objectTypeArgument);
    this.enumeratedValue = enumeratedValue;
    this.objectTypeArgument = objectTypeArgument;
  }

  public BACnetApplicationTagEnumerated getEnumeratedValue() {
    return enumeratedValue;
  }

  @Override
  protected void serializeBACnetTimerStateChangeValueChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetTimerStateChangeValueEnumerated");

    // Simple Field (enumeratedValue)
    writeSimpleField(
        "enumeratedValue", enumeratedValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetTimerStateChangeValueEnumerated");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetTimerStateChangeValueEnumerated _value = this;

    // Simple field (enumeratedValue)
    lengthInBits += enumeratedValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetTimerStateChangeValueBuilder staticParseBACnetTimerStateChangeValueBuilder(
      ReadBuffer readBuffer, BACnetObjectType objectTypeArgument) throws ParseException {
    readBuffer.pullContext("BACnetTimerStateChangeValueEnumerated");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetApplicationTagEnumerated enumeratedValue =
        readSimpleField(
            "enumeratedValue",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagEnumerated) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetTimerStateChangeValueEnumerated");
    // Create the instance
    return new BACnetTimerStateChangeValueEnumeratedBuilderImpl(
        enumeratedValue, objectTypeArgument);
  }

  public static class BACnetTimerStateChangeValueEnumeratedBuilderImpl
      implements BACnetTimerStateChangeValue.BACnetTimerStateChangeValueBuilder {
    private final BACnetApplicationTagEnumerated enumeratedValue;
    private final BACnetObjectType objectTypeArgument;

    public BACnetTimerStateChangeValueEnumeratedBuilderImpl(
        BACnetApplicationTagEnumerated enumeratedValue, BACnetObjectType objectTypeArgument) {
      this.enumeratedValue = enumeratedValue;
      this.objectTypeArgument = objectTypeArgument;
    }

    public BACnetTimerStateChangeValueEnumerated build(
        BACnetTagHeader peekedTagHeader, BACnetObjectType objectTypeArgument) {
      BACnetTimerStateChangeValueEnumerated bACnetTimerStateChangeValueEnumerated =
          new BACnetTimerStateChangeValueEnumerated(
              peekedTagHeader, enumeratedValue, objectTypeArgument);
      return bACnetTimerStateChangeValueEnumerated;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetTimerStateChangeValueEnumerated)) {
      return false;
    }
    BACnetTimerStateChangeValueEnumerated that = (BACnetTimerStateChangeValueEnumerated) o;
    return (getEnumeratedValue() == that.getEnumeratedValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getEnumeratedValue());
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
