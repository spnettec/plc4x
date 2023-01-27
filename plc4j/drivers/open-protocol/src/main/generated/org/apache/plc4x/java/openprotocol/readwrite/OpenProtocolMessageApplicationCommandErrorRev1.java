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
package org.apache.plc4x.java.openprotocol.readwrite;

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

public class OpenProtocolMessageApplicationCommandErrorRev1
    extends OpenProtocolMessageApplicationCommandError implements Message {

  // Accessors for discriminator values.
  public Long getRevision() {
    return (long) 1;
  }

  // Properties.
  protected final Mid requestMid;
  protected final Error error;

  public OpenProtocolMessageApplicationCommandErrorRev1(
      Long midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber,
      Mid requestMid,
      Error error) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
    this.requestMid = requestMid;
    this.error = error;
  }

  public Mid getRequestMid() {
    return requestMid;
  }

  public Error getError() {
    return error;
  }

  @Override
  protected void serializeOpenProtocolMessageApplicationCommandErrorChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("OpenProtocolMessageApplicationCommandErrorRev1");

    // Simple Field (requestMid)
    writeSimpleEnumField(
        "requestMid",
        "Mid",
        requestMid,
        new DataWriterEnumDefault<>(Mid::getValue, Mid::name, writeUnsignedLong(writeBuffer, 32)),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (error)
    writeSimpleEnumField(
        "error",
        "Error",
        error,
        new DataWriterEnumDefault<>(
            Error::getValue, Error::name, writeUnsignedInt(writeBuffer, 16)),
        WithOption.WithEncoding("ASCII"));

    writeBuffer.popContext("OpenProtocolMessageApplicationCommandErrorRev1");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageApplicationCommandErrorRev1 _value = this;

    // Simple field (requestMid)
    lengthInBits += 32;

    // Simple field (error)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static OpenProtocolMessageApplicationCommandErrorBuilder
      staticParseOpenProtocolMessageApplicationCommandErrorBuilder(
          ReadBuffer readBuffer, Long revision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageApplicationCommandErrorRev1");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    Mid requestMid =
        readEnumField(
            "requestMid",
            "Mid",
            new DataReaderEnumDefault<>(Mid::enumForValue, readUnsignedLong(readBuffer, 32)),
            WithOption.WithEncoding("ASCII"));

    Error error =
        readEnumField(
            "error",
            "Error",
            new DataReaderEnumDefault<>(Error::enumForValue, readUnsignedInt(readBuffer, 16)),
            WithOption.WithEncoding("ASCII"));

    readBuffer.closeContext("OpenProtocolMessageApplicationCommandErrorRev1");
    // Create the instance
    return new OpenProtocolMessageApplicationCommandErrorRev1BuilderImpl(requestMid, error);
  }

  public static class OpenProtocolMessageApplicationCommandErrorRev1BuilderImpl
      implements OpenProtocolMessageApplicationCommandError
          .OpenProtocolMessageApplicationCommandErrorBuilder {
    private final Mid requestMid;
    private final Error error;

    public OpenProtocolMessageApplicationCommandErrorRev1BuilderImpl(Mid requestMid, Error error) {
      this.requestMid = requestMid;
      this.error = error;
    }

    public OpenProtocolMessageApplicationCommandErrorRev1 build(
        Long midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      OpenProtocolMessageApplicationCommandErrorRev1
          openProtocolMessageApplicationCommandErrorRev1 =
              new OpenProtocolMessageApplicationCommandErrorRev1(
                  midRevision,
                  noAckFlag,
                  targetStationId,
                  targetSpindleId,
                  sequenceNumber,
                  numberOfMessageParts,
                  messagePartNumber,
                  requestMid,
                  error);
      return openProtocolMessageApplicationCommandErrorRev1;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageApplicationCommandErrorRev1)) {
      return false;
    }
    OpenProtocolMessageApplicationCommandErrorRev1 that =
        (OpenProtocolMessageApplicationCommandErrorRev1) o;
    return (getRequestMid() == that.getRequestMid())
        && (getError() == that.getError())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRequestMid(), getError());
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
