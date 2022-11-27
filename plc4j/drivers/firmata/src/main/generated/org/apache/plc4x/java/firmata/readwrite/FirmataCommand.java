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
package org.apache.plc4x.java.firmata.readwrite;

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

public abstract class FirmataCommand implements Message {

  // Abstract accessors for discriminator values.
  public abstract Byte getCommandCode();

  // Arguments.
  protected final Boolean response;

  public FirmataCommand(Boolean response) {
    super();
    this.response = response;
  }

  protected abstract void serializeFirmataCommandChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("FirmataCommand");

    // Discriminator Field (commandCode) (Used as input to a switch field)
    writeDiscriminatorField("commandCode", getCommandCode(), writeUnsignedByte(writeBuffer, 4));

    // Switch field (Serialize the sub-type)
    serializeFirmataCommandChild(writeBuffer);

    writeBuffer.popContext("FirmataCommand");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    FirmataCommand _value = this;

    // Discriminator Field (commandCode)
    lengthInBits += 4;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static FirmataCommand staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Boolean response;
    if (args[0] instanceof Boolean) {
      response = (Boolean) args[0];
    } else if (args[0] instanceof String) {
      response = Boolean.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Boolean or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, response);
  }

  public static FirmataCommand staticParse(ReadBuffer readBuffer, Boolean response)
      throws ParseException {
    readBuffer.pullContext("FirmataCommand");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    byte commandCode = readDiscriminatorField("commandCode", readUnsignedByte(readBuffer, 4));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    FirmataCommandBuilder builder = null;
    if (EvaluationHelper.equals(commandCode, (byte) 0x0)) {
      builder = FirmataCommandSysex.staticParseBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(commandCode, (byte) 0x4)) {
      builder = FirmataCommandSetPinMode.staticParseBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(commandCode, (byte) 0x5)) {
      builder = FirmataCommandSetDigitalPinValue.staticParseBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(commandCode, (byte) 0x9)) {
      builder = FirmataCommandProtocolVersion.staticParseBuilder(readBuffer, response);
    } else if (EvaluationHelper.equals(commandCode, (byte) 0xF)) {
      builder = FirmataCommandSystemReset.staticParseBuilder(readBuffer, response);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "commandCode="
              + commandCode
              + "]");
    }

    readBuffer.closeContext("FirmataCommand");
    // Create the instance
    FirmataCommand _firmataCommand = builder.build(response);

    return _firmataCommand;
  }

  public static interface FirmataCommandBuilder {
    FirmataCommand build(Boolean response);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof FirmataCommand)) {
      return false;
    }
    FirmataCommand that = (FirmataCommand) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
