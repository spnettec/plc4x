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
package org.apache.plc4x.java.plc4x.readwrite;

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

public class Plc4xReadResponse extends Plc4xMessage implements Message {

  // Accessors for discriminator values.
  public Plc4xRequestType getRequestType() {
    return Plc4xRequestType.READ_RESPONSE;
  }

  // Properties.
  protected final int connectionId;
  protected final Plc4xResponseCode responseCode;
  protected final List<Plc4xTagValueResponse> tags;

  public Plc4xReadResponse(
      int requestId,
      int connectionId,
      Plc4xResponseCode responseCode,
      List<Plc4xTagValueResponse> tags) {
    super(requestId);
    this.connectionId = connectionId;
    this.responseCode = responseCode;
    this.tags = tags;
  }

  public int getConnectionId() {
    return connectionId;
  }

  public Plc4xResponseCode getResponseCode() {
    return responseCode;
  }

  public List<Plc4xTagValueResponse> getTags() {
    return tags;
  }

  @Override
  protected void serializePlc4xMessageChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("Plc4xReadResponse");

    // Simple Field (connectionId)
    writeSimpleField(
        "connectionId",
        connectionId,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (responseCode)
    writeSimpleEnumField(
        "responseCode",
        "Plc4xResponseCode",
        responseCode,
        new DataWriterEnumDefault<>(
            Plc4xResponseCode::getValue,
            Plc4xResponseCode::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Implicit Field (numTags) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    short numTags = (short) (COUNT(getTags()));
    writeImplicitField("numTags", numTags, writeUnsignedShort(writeBuffer, 8));

    // Array Field (tags)
    writeComplexTypeArrayField("tags", tags, writeBuffer);

    writeBuffer.popContext("Plc4xReadResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    Plc4xReadResponse _value = this;

    // Simple field (connectionId)
    lengthInBits += 16;

    // Simple field (responseCode)
    lengthInBits += 8;

    // Implicit Field (numTags)
    lengthInBits += 8;

    // Array field
    if (tags != null) {
      int i = 0;
      for (Plc4xTagValueResponse element : tags) {
        boolean last = ++i >= tags.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static Plc4xReadResponseBuilder staticParseBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("Plc4xReadResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    int connectionId =
        readSimpleField(
            "connectionId",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Plc4xResponseCode responseCode =
        readEnumField(
            "responseCode",
            "Plc4xResponseCode",
            new DataReaderEnumDefault<>(
                Plc4xResponseCode::enumForValue, readUnsignedShort(readBuffer, 8)),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short numTags =
        readImplicitField(
            "numTags",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    List<Plc4xTagValueResponse> tags =
        readCountArrayField(
            "tags",
            new DataReaderComplexDefault<>(
                () -> Plc4xTagValueResponse.staticParse(readBuffer), readBuffer),
            numTags,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("Plc4xReadResponse");
    // Create the instance
    return new Plc4xReadResponseBuilder(connectionId, responseCode, tags);
  }

  public static class Plc4xReadResponseBuilder implements Plc4xMessage.Plc4xMessageBuilder {
    private final int connectionId;
    private final Plc4xResponseCode responseCode;
    private final List<Plc4xTagValueResponse> tags;

    public Plc4xReadResponseBuilder(
        int connectionId, Plc4xResponseCode responseCode, List<Plc4xTagValueResponse> tags) {

      this.connectionId = connectionId;
      this.responseCode = responseCode;
      this.tags = tags;
    }

    public Plc4xReadResponse build(int requestId) {
      Plc4xReadResponse plc4xReadResponse =
          new Plc4xReadResponse(requestId, connectionId, responseCode, tags);
      return plc4xReadResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Plc4xReadResponse)) {
      return false;
    }
    Plc4xReadResponse that = (Plc4xReadResponse) o;
    return (getConnectionId() == that.getConnectionId())
        && (getResponseCode() == that.getResponseCode())
        && (getTags() == that.getTags())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getConnectionId(), getResponseCode(), getTags());
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
