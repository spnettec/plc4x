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

public class BACnetServiceAckReadProperty extends BACnetServiceAck implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.READ_PROPERTY;
  }

  // Properties.
  protected final BACnetContextTagObjectIdentifier objectIdentifier;
  protected final BACnetPropertyIdentifierTagged propertyIdentifier;
  protected final BACnetContextTagUnsignedInteger arrayIndex;
  protected final BACnetConstructedData values;

  // Arguments.
  protected final Long serviceAckLength;

  public BACnetServiceAckReadProperty(
      BACnetContextTagObjectIdentifier objectIdentifier,
      BACnetPropertyIdentifierTagged propertyIdentifier,
      BACnetContextTagUnsignedInteger arrayIndex,
      BACnetConstructedData values,
      Long serviceAckLength) {
    super(serviceAckLength);
    this.objectIdentifier = objectIdentifier;
    this.propertyIdentifier = propertyIdentifier;
    this.arrayIndex = arrayIndex;
    this.values = values;
    this.serviceAckLength = serviceAckLength;
  }

  public BACnetContextTagObjectIdentifier getObjectIdentifier() {
    return objectIdentifier;
  }

  public BACnetPropertyIdentifierTagged getPropertyIdentifier() {
    return propertyIdentifier;
  }

  public BACnetContextTagUnsignedInteger getArrayIndex() {
    return arrayIndex;
  }

  public BACnetConstructedData getValues() {
    return values;
  }

  @Override
  protected void serializeBACnetServiceAckChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetServiceAckReadProperty");

    // Simple Field (objectIdentifier)
    writeSimpleField(
        "objectIdentifier", objectIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (propertyIdentifier)
    writeSimpleField(
        "propertyIdentifier", propertyIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (arrayIndex) (Can be skipped, if the value is null)
    writeOptionalField("arrayIndex", arrayIndex, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (values) (Can be skipped, if the value is null)
    writeOptionalField("values", values, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetServiceAckReadProperty");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetServiceAckReadProperty _value = this;

    // Simple field (objectIdentifier)
    lengthInBits += objectIdentifier.getLengthInBits();

    // Simple field (propertyIdentifier)
    lengthInBits += propertyIdentifier.getLengthInBits();

    // Optional Field (arrayIndex)
    if (arrayIndex != null) {
      lengthInBits += arrayIndex.getLengthInBits();
    }

    // Optional Field (values)
    if (values != null) {
      lengthInBits += values.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetServiceAckReadPropertyBuilder staticParseBuilder(
      ReadBuffer readBuffer, Long serviceAckLength) throws ParseException {
    readBuffer.pullContext("BACnetServiceAckReadProperty");
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

    BACnetContextTagUnsignedInteger arrayIndex =
        readOptionalField(
            "arrayIndex",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetConstructedData values =
        readOptionalField(
            "values",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetConstructedData.staticParse(
                        readBuffer,
                        (short) (3),
                        (BACnetObjectType) (objectIdentifier.getObjectType()),
                        (BACnetPropertyIdentifier) (propertyIdentifier.getValue()),
                        (BACnetTagPayloadUnsignedInteger)
                            (((((arrayIndex) != (null)) ? arrayIndex.getPayload() : null)))),
                readBuffer));

    readBuffer.closeContext("BACnetServiceAckReadProperty");
    // Create the instance
    return new BACnetServiceAckReadPropertyBuilder(
        objectIdentifier, propertyIdentifier, arrayIndex, values, serviceAckLength);
  }

  public static class BACnetServiceAckReadPropertyBuilder
      implements BACnetServiceAck.BACnetServiceAckBuilder {
    private final BACnetContextTagObjectIdentifier objectIdentifier;
    private final BACnetPropertyIdentifierTagged propertyIdentifier;
    private final BACnetContextTagUnsignedInteger arrayIndex;
    private final BACnetConstructedData values;
    private final Long serviceAckLength;

    public BACnetServiceAckReadPropertyBuilder(
        BACnetContextTagObjectIdentifier objectIdentifier,
        BACnetPropertyIdentifierTagged propertyIdentifier,
        BACnetContextTagUnsignedInteger arrayIndex,
        BACnetConstructedData values,
        Long serviceAckLength) {

      this.objectIdentifier = objectIdentifier;
      this.propertyIdentifier = propertyIdentifier;
      this.arrayIndex = arrayIndex;
      this.values = values;
      this.serviceAckLength = serviceAckLength;
    }

    public BACnetServiceAckReadProperty build(Long serviceAckLength) {

      BACnetServiceAckReadProperty bACnetServiceAckReadProperty =
          new BACnetServiceAckReadProperty(
              objectIdentifier, propertyIdentifier, arrayIndex, values, serviceAckLength);
      return bACnetServiceAckReadProperty;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetServiceAckReadProperty)) {
      return false;
    }
    BACnetServiceAckReadProperty that = (BACnetServiceAckReadProperty) o;
    return (getObjectIdentifier() == that.getObjectIdentifier())
        && (getPropertyIdentifier() == that.getPropertyIdentifier())
        && (getArrayIndex() == that.getArrayIndex())
        && (getValues() == that.getValues())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getObjectIdentifier(),
        getPropertyIdentifier(),
        getArrayIndex(),
        getValues());
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
