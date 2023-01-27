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
package org.apache.plc4x.java.opcua.readwrite;

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

public class StructureField extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "103";
  }

  // Properties.
  protected final PascalString name;
  protected final LocalizedText description;
  protected final NodeId dataType;
  protected final int valueRank;
  protected final int noOfArrayDimensions;
  protected final List<Long> arrayDimensions;
  protected final long maxStringLength;
  protected final boolean isOptional;

  public StructureField(
      PascalString name,
      LocalizedText description,
      NodeId dataType,
      int valueRank,
      int noOfArrayDimensions,
      List<Long> arrayDimensions,
      long maxStringLength,
      boolean isOptional) {
    super();
    this.name = name;
    this.description = description;
    this.dataType = dataType;
    this.valueRank = valueRank;
    this.noOfArrayDimensions = noOfArrayDimensions;
    this.arrayDimensions = arrayDimensions;
    this.maxStringLength = maxStringLength;
    this.isOptional = isOptional;
  }

  public PascalString getName() {
    return name;
  }

  public LocalizedText getDescription() {
    return description;
  }

  public NodeId getDataType() {
    return dataType;
  }

  public int getValueRank() {
    return valueRank;
  }

  public int getNoOfArrayDimensions() {
    return noOfArrayDimensions;
  }

  public List<Long> getArrayDimensions() {
    return arrayDimensions;
  }

  public long getMaxStringLength() {
    return maxStringLength;
  }

  public boolean getIsOptional() {
    return isOptional;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("StructureField");

    // Simple Field (name)
    writeSimpleField("name", name, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (dataType)
    writeSimpleField("dataType", dataType, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (valueRank)
    writeSimpleField("valueRank", valueRank, writeSignedInt(writeBuffer, 32));

    // Simple Field (noOfArrayDimensions)
    writeSimpleField("noOfArrayDimensions", noOfArrayDimensions, writeSignedInt(writeBuffer, 32));

    // Array Field (arrayDimensions)
    writeSimpleTypeArrayField(
        "arrayDimensions", arrayDimensions, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (maxStringLength)
    writeSimpleField("maxStringLength", maxStringLength, writeUnsignedLong(writeBuffer, 32));

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 7));

    // Simple Field (isOptional)
    writeSimpleField("isOptional", isOptional, writeBoolean(writeBuffer));

    writeBuffer.popContext("StructureField");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    StructureField _value = this;

    // Simple field (name)
    lengthInBits += name.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Simple field (dataType)
    lengthInBits += dataType.getLengthInBits();

    // Simple field (valueRank)
    lengthInBits += 32;

    // Simple field (noOfArrayDimensions)
    lengthInBits += 32;

    // Array field
    if (arrayDimensions != null) {
      lengthInBits += 32 * arrayDimensions.size();
    }

    // Simple field (maxStringLength)
    lengthInBits += 32;

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (isOptional)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("StructureField");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    PascalString name =
        readSimpleField(
            "name",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    LocalizedText description =
        readSimpleField(
            "description",
            new DataReaderComplexDefault<>(
                () -> LocalizedText.staticParse(readBuffer), readBuffer));

    NodeId dataType =
        readSimpleField(
            "dataType",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    int valueRank = readSimpleField("valueRank", readSignedInt(readBuffer, 32));

    int noOfArrayDimensions = readSimpleField("noOfArrayDimensions", readSignedInt(readBuffer, 32));

    List<Long> arrayDimensions =
        readCountArrayField(
            "arrayDimensions", readUnsignedLong(readBuffer, 32), noOfArrayDimensions);

    long maxStringLength = readSimpleField("maxStringLength", readUnsignedLong(readBuffer, 32));

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 7), (short) 0x00);

    boolean isOptional = readSimpleField("isOptional", readBoolean(readBuffer));

    readBuffer.closeContext("StructureField");
    // Create the instance
    return new StructureFieldBuilderImpl(
        name,
        description,
        dataType,
        valueRank,
        noOfArrayDimensions,
        arrayDimensions,
        maxStringLength,
        isOptional);
  }

  public static class StructureFieldBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString name;
    private final LocalizedText description;
    private final NodeId dataType;
    private final int valueRank;
    private final int noOfArrayDimensions;
    private final List<Long> arrayDimensions;
    private final long maxStringLength;
    private final boolean isOptional;

    public StructureFieldBuilderImpl(
        PascalString name,
        LocalizedText description,
        NodeId dataType,
        int valueRank,
        int noOfArrayDimensions,
        List<Long> arrayDimensions,
        long maxStringLength,
        boolean isOptional) {
      this.name = name;
      this.description = description;
      this.dataType = dataType;
      this.valueRank = valueRank;
      this.noOfArrayDimensions = noOfArrayDimensions;
      this.arrayDimensions = arrayDimensions;
      this.maxStringLength = maxStringLength;
      this.isOptional = isOptional;
    }

    public StructureField build() {
      StructureField structureField =
          new StructureField(
              name,
              description,
              dataType,
              valueRank,
              noOfArrayDimensions,
              arrayDimensions,
              maxStringLength,
              isOptional);
      return structureField;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof StructureField)) {
      return false;
    }
    StructureField that = (StructureField) o;
    return (getName() == that.getName())
        && (getDescription() == that.getDescription())
        && (getDataType() == that.getDataType())
        && (getValueRank() == that.getValueRank())
        && (getNoOfArrayDimensions() == that.getNoOfArrayDimensions())
        && (getArrayDimensions() == that.getArrayDimensions())
        && (getMaxStringLength() == that.getMaxStringLength())
        && (getIsOptional() == that.getIsOptional())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getName(),
        getDescription(),
        getDataType(),
        getValueRank(),
        getNoOfArrayDimensions(),
        getArrayDimensions(),
        getMaxStringLength(),
        getIsOptional());
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
