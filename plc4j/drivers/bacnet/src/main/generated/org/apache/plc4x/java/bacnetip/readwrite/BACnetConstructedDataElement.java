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

public class BACnetConstructedDataElement implements Message {

  // Properties.
  protected final BACnetTagHeader peekedTagHeader;
  protected final BACnetApplicationTag applicationTag;
  protected final BACnetContextTag contextTag;
  protected final BACnetConstructedData constructedData;

  // Arguments.
  protected final BACnetObjectType objectTypeArgument;
  protected final BACnetPropertyIdentifier propertyIdentifierArgument;
  protected final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

  public BACnetConstructedDataElement(
      BACnetTagHeader peekedTagHeader,
      BACnetApplicationTag applicationTag,
      BACnetContextTag contextTag,
      BACnetConstructedData constructedData,
      BACnetObjectType objectTypeArgument,
      BACnetPropertyIdentifier propertyIdentifierArgument,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
    super();
    this.peekedTagHeader = peekedTagHeader;
    this.applicationTag = applicationTag;
    this.contextTag = contextTag;
    this.constructedData = constructedData;
    this.objectTypeArgument = objectTypeArgument;
    this.propertyIdentifierArgument = propertyIdentifierArgument;
    this.arrayIndexArgument = arrayIndexArgument;
  }

  public BACnetTagHeader getPeekedTagHeader() {
    return peekedTagHeader;
  }

  public BACnetApplicationTag getApplicationTag() {
    return applicationTag;
  }

  public BACnetContextTag getContextTag() {
    return contextTag;
  }

  public BACnetConstructedData getConstructedData() {
    return constructedData;
  }

  public short getPeekedTagNumber() {
    return (short) (getPeekedTagHeader().getActualTagNumber());
  }

  public boolean getIsApplicationTag() {
    return (boolean) ((getPeekedTagHeader().getTagClass()) == (TagClass.APPLICATION_TAGS));
  }

  public boolean getIsConstructedData() {
    return (boolean)
        ((!(getIsApplicationTag())) && ((getPeekedTagHeader().getLengthValueType()) == (0x6)));
  }

  public boolean getIsContextTag() {
    return (boolean) ((!(getIsConstructedData())) && (!(getIsApplicationTag())));
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetConstructedDataElement");

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    short peekedTagNumber = getPeekedTagNumber();
    writeBuffer.writeVirtual("peekedTagNumber", peekedTagNumber);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isApplicationTag = getIsApplicationTag();
    writeBuffer.writeVirtual("isApplicationTag", isApplicationTag);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isConstructedData = getIsConstructedData();
    writeBuffer.writeVirtual("isConstructedData", isConstructedData);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isContextTag = getIsContextTag();
    writeBuffer.writeVirtual("isContextTag", isContextTag);

    // Optional Field (applicationTag) (Can be skipped, if the value is null)
    writeOptionalField(
        "applicationTag",
        applicationTag,
        new DataWriterComplexDefault<>(writeBuffer),
        getIsApplicationTag());

    // Optional Field (contextTag) (Can be skipped, if the value is null)
    writeOptionalField(
        "contextTag", contextTag, new DataWriterComplexDefault<>(writeBuffer), getIsContextTag());

    // Optional Field (constructedData) (Can be skipped, if the value is null)
    writeOptionalField(
        "constructedData",
        constructedData,
        new DataWriterComplexDefault<>(writeBuffer),
        getIsConstructedData());

    writeBuffer.popContext("BACnetConstructedDataElement");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetConstructedDataElement _value = this;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Optional Field (applicationTag)
    if (applicationTag != null) {
      lengthInBits += applicationTag.getLengthInBits();
    }

    // Optional Field (contextTag)
    if (contextTag != null) {
      lengthInBits += contextTag.getLengthInBits();
    }

    // Optional Field (constructedData)
    if (constructedData != null) {
      lengthInBits += constructedData.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetConstructedDataElement staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 3)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 3, but got " + args.length);
    }
    BACnetObjectType objectTypeArgument;
    if (args[0] instanceof BACnetObjectType) {
      objectTypeArgument = (BACnetObjectType) args[0];
    } else if (args[0] instanceof String) {
      objectTypeArgument = BACnetObjectType.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type BACnetObjectType or a string which is parseable but"
              + " was "
              + args[0].getClass().getName());
    }
    BACnetPropertyIdentifier propertyIdentifierArgument;
    if (args[1] instanceof BACnetPropertyIdentifier) {
      propertyIdentifierArgument = (BACnetPropertyIdentifier) args[1];
    } else if (args[1] instanceof String) {
      propertyIdentifierArgument = BACnetPropertyIdentifier.valueOf((String) args[1]);
    } else {
      throw new PlcRuntimeException(
          "Argument 1 expected to be of type BACnetPropertyIdentifier or a string which is"
              + " parseable but was "
              + args[1].getClass().getName());
    }
    BACnetTagPayloadUnsignedInteger arrayIndexArgument;
    if (args[2] instanceof BACnetTagPayloadUnsignedInteger) {
      arrayIndexArgument = (BACnetTagPayloadUnsignedInteger) args[2];
    } else {
      throw new PlcRuntimeException(
          "Argument 2 expected to be of type BACnetTagPayloadUnsignedInteger or a string which is"
              + " parseable but was "
              + args[2].getClass().getName());
    }
    return staticParse(
        readBuffer, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument);
  }

  public static BACnetConstructedDataElement staticParse(
      ReadBuffer readBuffer,
      BACnetObjectType objectTypeArgument,
      BACnetPropertyIdentifier propertyIdentifierArgument,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetConstructedDataElement");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    BACnetTagHeader peekedTagHeader =
        readPeekField(
            "peekedTagHeader",
            new DataReaderComplexDefault<>(
                () -> BACnetTagHeader.staticParse(readBuffer), readBuffer));
    short peekedTagNumber =
        readVirtualField("peekedTagNumber", short.class, peekedTagHeader.getActualTagNumber());
    boolean isApplicationTag =
        readVirtualField(
            "isApplicationTag",
            boolean.class,
            (peekedTagHeader.getTagClass()) == (TagClass.APPLICATION_TAGS));
    boolean isConstructedData =
        readVirtualField(
            "isConstructedData",
            boolean.class,
            (!(isApplicationTag)) && ((peekedTagHeader.getLengthValueType()) == (0x6)));
    boolean isContextTag =
        readVirtualField(
            "isContextTag", boolean.class, (!(isConstructedData)) && (!(isApplicationTag)));
    // Validation
    if (!((!(isContextTag))
        || (((isContextTag) && ((peekedTagHeader.getLengthValueType()) != (0x7)))))) {
      throw new ParseValidationException("unexpected closing tag");
    }

    BACnetApplicationTag applicationTag =
        readOptionalField(
            "applicationTag",
            new DataReaderComplexDefault<>(
                () -> BACnetApplicationTag.staticParse(readBuffer), readBuffer),
            isApplicationTag);

    BACnetContextTag contextTag =
        readOptionalField(
            "contextTag",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetContextTag.staticParse(
                        readBuffer,
                        (short) (peekedTagNumber),
                        (BACnetDataType) (BACnetDataType.UNKNOWN)),
                readBuffer),
            isContextTag);

    BACnetConstructedData constructedData =
        readOptionalField(
            "constructedData",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetConstructedData.staticParse(
                        readBuffer,
                        (short) (peekedTagNumber),
                        (BACnetObjectType) (objectTypeArgument),
                        (BACnetPropertyIdentifier) (propertyIdentifierArgument),
                        (BACnetTagPayloadUnsignedInteger) (arrayIndexArgument)),
                readBuffer),
            isConstructedData);
    // Validation
    if (!(((((isApplicationTag) && ((applicationTag) != (null))))
            || (((isContextTag) && ((contextTag) != (null)))))
        || (((isConstructedData) && ((constructedData) != (null)))))) {
      throw new ParseValidationException("BACnetConstructedDataElement could not parse anything");
    }

    readBuffer.closeContext("BACnetConstructedDataElement");
    // Create the instance
    BACnetConstructedDataElement _bACnetConstructedDataElement;
    _bACnetConstructedDataElement =
        new BACnetConstructedDataElement(
            peekedTagHeader,
            applicationTag,
            contextTag,
            constructedData,
            objectTypeArgument,
            propertyIdentifierArgument,
            arrayIndexArgument);
    return _bACnetConstructedDataElement;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConstructedDataElement)) {
      return false;
    }
    BACnetConstructedDataElement that = (BACnetConstructedDataElement) o;
    return (getPeekedTagHeader() == that.getPeekedTagHeader())
        && (getApplicationTag() == that.getApplicationTag())
        && (getContextTag() == that.getContextTag())
        && (getConstructedData() == that.getConstructedData())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getPeekedTagHeader(), getApplicationTag(), getContextTag(), getConstructedData());
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
