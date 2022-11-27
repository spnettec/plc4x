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

public class BACnetConfirmedServiceRequestReadRange extends BACnetConfirmedServiceRequest
    implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.READ_RANGE;
  }

  // Properties.
  protected final BACnetContextTagObjectIdentifier objectIdentifier;
  protected final BACnetPropertyIdentifierTagged propertyIdentifier;
  protected final BACnetContextTagUnsignedInteger propertyArrayIndex;
  protected final BACnetConfirmedServiceRequestReadRangeRange readRange;

  // Arguments.
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestReadRange(
      BACnetContextTagObjectIdentifier objectIdentifier,
      BACnetPropertyIdentifierTagged propertyIdentifier,
      BACnetContextTagUnsignedInteger propertyArrayIndex,
      BACnetConfirmedServiceRequestReadRangeRange readRange,
      Long serviceRequestLength) {
    super(serviceRequestLength);
    this.objectIdentifier = objectIdentifier;
    this.propertyIdentifier = propertyIdentifier;
    this.propertyArrayIndex = propertyArrayIndex;
    this.readRange = readRange;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetContextTagObjectIdentifier getObjectIdentifier() {
    return objectIdentifier;
  }

  public BACnetPropertyIdentifierTagged getPropertyIdentifier() {
    return propertyIdentifier;
  }

  public BACnetContextTagUnsignedInteger getPropertyArrayIndex() {
    return propertyArrayIndex;
  }

  public BACnetConfirmedServiceRequestReadRangeRange getReadRange() {
    return readRange;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestReadRange");

    // Simple Field (objectIdentifier)
    writeSimpleField(
        "objectIdentifier", objectIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (propertyIdentifier)
    writeSimpleField(
        "propertyIdentifier", propertyIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (propertyArrayIndex) (Can be skipped, if the value is null)
    writeOptionalField(
        "propertyArrayIndex", propertyArrayIndex, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (readRange) (Can be skipped, if the value is null)
    writeOptionalField("readRange", readRange, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestReadRange");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestReadRange _value = this;

    // Simple field (objectIdentifier)
    lengthInBits += objectIdentifier.getLengthInBits();

    // Simple field (propertyIdentifier)
    lengthInBits += propertyIdentifier.getLengthInBits();

    // Optional Field (propertyArrayIndex)
    if (propertyArrayIndex != null) {
      lengthInBits += propertyArrayIndex.getLengthInBits();
    }

    // Optional Field (readRange)
    if (readRange != null) {
      lengthInBits += readRange.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestReadRangeBuilder staticParseBuilder(
      ReadBuffer readBuffer, Long serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestReadRange");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetContextTagObjectIdentifier objectIdentifier =
        readSimpleField(
            "objectIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    BACnetPropertyIdentifierTagged propertyIdentifier =
        readSimpleField(
            "propertyIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetPropertyIdentifierTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagUnsignedInteger propertyArrayIndex =
        readOptionalField(
            "propertyArrayIndex",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetConfirmedServiceRequestReadRangeRange readRange =
        readOptionalField(
            "readRange",
            new DataReaderComplexDefault<>(
                () -> BACnetConfirmedServiceRequestReadRangeRange.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestReadRange");
    // Create the instance
    return new BACnetConfirmedServiceRequestReadRangeBuilder(
        objectIdentifier, propertyIdentifier, propertyArrayIndex, readRange, serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestReadRangeBuilder
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final BACnetContextTagObjectIdentifier objectIdentifier;
    private final BACnetPropertyIdentifierTagged propertyIdentifier;
    private final BACnetContextTagUnsignedInteger propertyArrayIndex;
    private final BACnetConfirmedServiceRequestReadRangeRange readRange;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestReadRangeBuilder(
        BACnetContextTagObjectIdentifier objectIdentifier,
        BACnetPropertyIdentifierTagged propertyIdentifier,
        BACnetContextTagUnsignedInteger propertyArrayIndex,
        BACnetConfirmedServiceRequestReadRangeRange readRange,
        Long serviceRequestLength) {

      this.objectIdentifier = objectIdentifier;
      this.propertyIdentifier = propertyIdentifier;
      this.propertyArrayIndex = propertyArrayIndex;
      this.readRange = readRange;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestReadRange build(Long serviceRequestLength) {

      BACnetConfirmedServiceRequestReadRange bACnetConfirmedServiceRequestReadRange =
          new BACnetConfirmedServiceRequestReadRange(
              objectIdentifier,
              propertyIdentifier,
              propertyArrayIndex,
              readRange,
              serviceRequestLength);
      return bACnetConfirmedServiceRequestReadRange;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestReadRange)) {
      return false;
    }
    BACnetConfirmedServiceRequestReadRange that = (BACnetConfirmedServiceRequestReadRange) o;
    return (getObjectIdentifier() == that.getObjectIdentifier())
        && (getPropertyIdentifier() == that.getPropertyIdentifier())
        && (getPropertyArrayIndex() == that.getPropertyArrayIndex())
        && (getReadRange() == that.getReadRange())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getObjectIdentifier(),
        getPropertyIdentifier(),
        getPropertyArrayIndex(),
        getReadRange());
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
