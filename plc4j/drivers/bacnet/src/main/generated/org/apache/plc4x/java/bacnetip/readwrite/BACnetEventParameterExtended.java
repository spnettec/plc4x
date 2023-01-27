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

public class BACnetEventParameterExtended extends BACnetEventParameter implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final BACnetVendorIdTagged vendorId;
  protected final BACnetContextTagUnsignedInteger extendedEventType;
  protected final BACnetEventParameterExtendedParameters parameters;
  protected final BACnetClosingTag closingTag;

  public BACnetEventParameterExtended(
      BACnetTagHeader peekedTagHeader,
      BACnetOpeningTag openingTag,
      BACnetVendorIdTagged vendorId,
      BACnetContextTagUnsignedInteger extendedEventType,
      BACnetEventParameterExtendedParameters parameters,
      BACnetClosingTag closingTag) {
    super(peekedTagHeader);
    this.openingTag = openingTag;
    this.vendorId = vendorId;
    this.extendedEventType = extendedEventType;
    this.parameters = parameters;
    this.closingTag = closingTag;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public BACnetVendorIdTagged getVendorId() {
    return vendorId;
  }

  public BACnetContextTagUnsignedInteger getExtendedEventType() {
    return extendedEventType;
  }

  public BACnetEventParameterExtendedParameters getParameters() {
    return parameters;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  @Override
  protected void serializeBACnetEventParameterChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetEventParameterExtended");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (vendorId)
    writeSimpleField("vendorId", vendorId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (extendedEventType)
    writeSimpleField(
        "extendedEventType", extendedEventType, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (parameters)
    writeSimpleField("parameters", parameters, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetEventParameterExtended");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetEventParameterExtended _value = this;

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Simple field (vendorId)
    lengthInBits += vendorId.getLengthInBits();

    // Simple field (extendedEventType)
    lengthInBits += extendedEventType.getLengthInBits();

    // Simple field (parameters)
    lengthInBits += parameters.getLengthInBits();

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetEventParameterBuilder staticParseBACnetEventParameterBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetEventParameterExtended");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (9)), readBuffer));

    BACnetVendorIdTagged vendorId =
        readSimpleField(
            "vendorId",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetVendorIdTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagUnsignedInteger extendedEventType =
        readSimpleField(
            "extendedEventType",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetEventParameterExtendedParameters parameters =
        readSimpleField(
            "parameters",
            new DataReaderComplexDefault<>(
                () -> BACnetEventParameterExtendedParameters.staticParse(readBuffer, (short) (2)),
                readBuffer));

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (9)), readBuffer));

    readBuffer.closeContext("BACnetEventParameterExtended");
    // Create the instance
    return new BACnetEventParameterExtendedBuilderImpl(
        openingTag, vendorId, extendedEventType, parameters, closingTag);
  }

  public static class BACnetEventParameterExtendedBuilderImpl
      implements BACnetEventParameter.BACnetEventParameterBuilder {
    private final BACnetOpeningTag openingTag;
    private final BACnetVendorIdTagged vendorId;
    private final BACnetContextTagUnsignedInteger extendedEventType;
    private final BACnetEventParameterExtendedParameters parameters;
    private final BACnetClosingTag closingTag;

    public BACnetEventParameterExtendedBuilderImpl(
        BACnetOpeningTag openingTag,
        BACnetVendorIdTagged vendorId,
        BACnetContextTagUnsignedInteger extendedEventType,
        BACnetEventParameterExtendedParameters parameters,
        BACnetClosingTag closingTag) {
      this.openingTag = openingTag;
      this.vendorId = vendorId;
      this.extendedEventType = extendedEventType;
      this.parameters = parameters;
      this.closingTag = closingTag;
    }

    public BACnetEventParameterExtended build(BACnetTagHeader peekedTagHeader) {
      BACnetEventParameterExtended bACnetEventParameterExtended =
          new BACnetEventParameterExtended(
              peekedTagHeader, openingTag, vendorId, extendedEventType, parameters, closingTag);
      return bACnetEventParameterExtended;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetEventParameterExtended)) {
      return false;
    }
    BACnetEventParameterExtended that = (BACnetEventParameterExtended) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getVendorId() == that.getVendorId())
        && (getExtendedEventType() == that.getExtendedEventType())
        && (getParameters() == that.getParameters())
        && (getClosingTag() == that.getClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getOpeningTag(),
        getVendorId(),
        getExtendedEventType(),
        getParameters(),
        getClosingTag());
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
