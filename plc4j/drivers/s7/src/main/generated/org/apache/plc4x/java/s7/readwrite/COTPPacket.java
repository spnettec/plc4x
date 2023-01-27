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
package org.apache.plc4x.java.s7.readwrite;

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

public abstract class COTPPacket implements Message {

  // Abstract accessors for discriminator values.
  public abstract Short getTpduCode();

  // Properties.
  protected final List<COTPParameter> parameters;
  protected final S7Message payload;

  public COTPPacket(List<COTPParameter> parameters, S7Message payload) {
    super();
    this.parameters = parameters;
    this.payload = payload;
  }

  public List<COTPParameter> getParameters() {
    return parameters;
  }

  public S7Message getPayload() {
    return payload;
  }

  protected abstract void serializeCOTPPacketChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("COTPPacket");

    // Implicit Field (headerLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    short headerLength =
        (short)
            ((getLengthInBytes())
                - ((((((((getPayload()) != (null))) ? getPayload().getLengthInBytes() : 0)))
                    + (1))));
    writeImplicitField("headerLength", headerLength, writeUnsignedShort(writeBuffer, 8));

    // Discriminator Field (tpduCode) (Used as input to a switch field)
    writeDiscriminatorField("tpduCode", getTpduCode(), writeUnsignedShort(writeBuffer, 8));

    // Switch field (Serialize the sub-type)
    serializeCOTPPacketChild(writeBuffer);

    // Array Field (parameters)
    writeComplexTypeArrayField("parameters", parameters, writeBuffer);

    // Optional Field (payload) (Can be skipped, if the value is null)
    writeOptionalField("payload", payload, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("COTPPacket");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    COTPPacket _value = this;

    // Implicit Field (headerLength)
    lengthInBits += 8;

    // Discriminator Field (tpduCode)
    lengthInBits += 8;

    // Length of sub-type elements will be added by sub-type...

    // Array field
    if (parameters != null) {
      for (Message element : parameters) {
        lengthInBits += element.getLengthInBits();
      }
    }

    // Optional Field (payload)
    if (payload != null) {
      lengthInBits += payload.getLengthInBits();
    }

    return lengthInBits;
  }

  public static COTPPacket staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Integer cotpLen;
    if (args[0] instanceof Integer) {
      cotpLen = (Integer) args[0];
    } else if (args[0] instanceof String) {
      cotpLen = Integer.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Integer or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, cotpLen);
  }

  public static COTPPacket staticParse(ReadBuffer readBuffer, Integer cotpLen)
      throws ParseException {
    readBuffer.pullContext("COTPPacket");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    short headerLength = readImplicitField("headerLength", readUnsignedShort(readBuffer, 8));

    short tpduCode = readDiscriminatorField("tpduCode", readUnsignedShort(readBuffer, 8));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    COTPPacketBuilder builder = null;
    if (EvaluationHelper.equals(tpduCode, (short) 0xF0)) {
      builder = COTPPacketData.staticParseCOTPPacketBuilder(readBuffer, cotpLen);
    } else if (EvaluationHelper.equals(tpduCode, (short) 0xE0)) {
      builder = COTPPacketConnectionRequest.staticParseCOTPPacketBuilder(readBuffer, cotpLen);
    } else if (EvaluationHelper.equals(tpduCode, (short) 0xD0)) {
      builder = COTPPacketConnectionResponse.staticParseCOTPPacketBuilder(readBuffer, cotpLen);
    } else if (EvaluationHelper.equals(tpduCode, (short) 0x80)) {
      builder = COTPPacketDisconnectRequest.staticParseCOTPPacketBuilder(readBuffer, cotpLen);
    } else if (EvaluationHelper.equals(tpduCode, (short) 0xC0)) {
      builder = COTPPacketDisconnectResponse.staticParseCOTPPacketBuilder(readBuffer, cotpLen);
    } else if (EvaluationHelper.equals(tpduCode, (short) 0x70)) {
      builder = COTPPacketTpduError.staticParseCOTPPacketBuilder(readBuffer, cotpLen);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "tpduCode="
              + tpduCode
              + "]");
    }

    List<COTPParameter> parameters =
        readLengthArrayField(
            "parameters",
            new DataReaderComplexDefault<>(
                () ->
                    COTPParameter.staticParse(
                        readBuffer,
                        (short) ((((headerLength) + (1))) - ((positionAware.getPos() - startPos)))),
                readBuffer),
            (((headerLength) + (1))) - ((positionAware.getPos() - startPos)));

    S7Message payload =
        readOptionalField(
            "payload",
            new DataReaderComplexDefault<>(() -> S7Message.staticParse(readBuffer), readBuffer),
            ((positionAware.getPos() - startPos)) < (cotpLen));

    readBuffer.closeContext("COTPPacket");
    // Create the instance
    COTPPacket _cOTPPacket = builder.build(parameters, payload);
    return _cOTPPacket;
  }

  public interface COTPPacketBuilder {
    COTPPacket build(List<COTPParameter> parameters, S7Message payload);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof COTPPacket)) {
      return false;
    }
    COTPPacket that = (COTPPacket) o;
    return (getParameters() == that.getParameters()) && (getPayload() == that.getPayload()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getParameters(), getPayload());
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
