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
package org.apache.plc4x.java.cbus.readwrite;

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

public class IdentifyReplyCommandLogicalAssignment extends IdentifyReplyCommand implements Message {

  // Accessors for discriminator values.
  public Attribute getAttribute() {
    return Attribute.LogicalAssignment;
  }

  // Properties.
  protected final List<LogicAssignment> logicAssigment;

  // Arguments.
  protected final Short numBytes;

  public IdentifyReplyCommandLogicalAssignment(
      List<LogicAssignment> logicAssigment, Short numBytes) {
    super(numBytes);
    this.logicAssigment = logicAssigment;
    this.numBytes = numBytes;
  }

  public List<LogicAssignment> getLogicAssigment() {
    return logicAssigment;
  }

  @Override
  protected void serializeIdentifyReplyCommandChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("IdentifyReplyCommandLogicalAssignment");

    // Array Field (logicAssigment)
    writeComplexTypeArrayField("logicAssigment", logicAssigment, writeBuffer);

    writeBuffer.popContext("IdentifyReplyCommandLogicalAssignment");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    IdentifyReplyCommandLogicalAssignment _value = this;

    // Array field
    if (logicAssigment != null) {
      int i = 0;
      for (LogicAssignment element : logicAssigment) {
        boolean last = ++i >= logicAssigment.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static IdentifyReplyCommandLogicalAssignmentBuilder staticParseBuilder(
      ReadBuffer readBuffer, Attribute attribute, Short numBytes) throws ParseException {
    readBuffer.pullContext("IdentifyReplyCommandLogicalAssignment");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    List<LogicAssignment> logicAssigment =
        readCountArrayField(
            "logicAssigment",
            new DataReaderComplexDefault<>(
                () -> LogicAssignment.staticParse(readBuffer), readBuffer),
            numBytes);

    readBuffer.closeContext("IdentifyReplyCommandLogicalAssignment");
    // Create the instance
    return new IdentifyReplyCommandLogicalAssignmentBuilder(logicAssigment, numBytes);
  }

  public static class IdentifyReplyCommandLogicalAssignmentBuilder
      implements IdentifyReplyCommand.IdentifyReplyCommandBuilder {
    private final List<LogicAssignment> logicAssigment;
    private final Short numBytes;

    public IdentifyReplyCommandLogicalAssignmentBuilder(
        List<LogicAssignment> logicAssigment, Short numBytes) {

      this.logicAssigment = logicAssigment;
      this.numBytes = numBytes;
    }

    public IdentifyReplyCommandLogicalAssignment build(Short numBytes) {

      IdentifyReplyCommandLogicalAssignment identifyReplyCommandLogicalAssignment =
          new IdentifyReplyCommandLogicalAssignment(logicAssigment, numBytes);
      return identifyReplyCommandLogicalAssignment;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof IdentifyReplyCommandLogicalAssignment)) {
      return false;
    }
    IdentifyReplyCommandLogicalAssignment that = (IdentifyReplyCommandLogicalAssignment) o;
    return (getLogicAssigment() == that.getLogicAssigment()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getLogicAssigment());
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
