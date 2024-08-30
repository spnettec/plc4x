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
package org.apache.plc4x.java.profinet.readwrite;

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

public abstract class LldpUnit implements Message {

  // Abstract accessors for discriminator values.
  public abstract TlvType getTlvId();

  // Properties.
  protected final short tlvIdLength;

  public LldpUnit(short tlvIdLength) {
    super();
    this.tlvIdLength = tlvIdLength;
  }

  public short getTlvIdLength() {
    return tlvIdLength;
  }

  protected abstract void serializeLldpUnitChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("LldpUnit");

    // Discriminator Field (tlvId) (Used as input to a switch field)
    writeDiscriminatorEnumField(
        "tlvId",
        "TlvType",
        getTlvId(),
        writeEnum(TlvType::getValue, TlvType::name, writeUnsignedByte(writeBuffer, 7)));

    // Simple Field (tlvIdLength)
    writeSimpleField("tlvIdLength", tlvIdLength, writeUnsignedShort(writeBuffer, 9));

    // Switch field (Serialize the sub-type)
    serializeLldpUnitChild(writeBuffer);

    writeBuffer.popContext("LldpUnit");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    LldpUnit _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (tlvId)
    lengthInBits += 7;

    // Simple field (tlvIdLength)
    lengthInBits += 9;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static LldpUnit staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("LldpUnit");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    TlvType tlvId =
        readDiscriminatorEnumField(
            "tlvId", "TlvType", readEnum(TlvType::enumForValue, readUnsignedByte(readBuffer, 7)));

    short tlvIdLength = readSimpleField("tlvIdLength", readUnsignedShort(readBuffer, 9));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    LldpUnitBuilder builder = null;
    if (EvaluationHelper.equals(tlvId, TlvType.END_OF_LLDP)) {
      builder = EndOfLldp.staticParseLldpUnitBuilder(readBuffer);
    } else if (EvaluationHelper.equals(tlvId, TlvType.CHASSIS_ID)) {
      builder = TlvChassisId.staticParseLldpUnitBuilder(readBuffer, tlvIdLength);
    } else if (EvaluationHelper.equals(tlvId, TlvType.PORT_ID)) {
      builder = TlvPortId.staticParseLldpUnitBuilder(readBuffer, tlvIdLength);
    } else if (EvaluationHelper.equals(tlvId, TlvType.TIME_TO_LIVE)) {
      builder = TlvTimeToLive.staticParseLldpUnitBuilder(readBuffer);
    } else if (EvaluationHelper.equals(tlvId, TlvType.PORT_DESCRIPTION)) {
      builder = TlvPortDescription.staticParseLldpUnitBuilder(readBuffer, tlvIdLength);
    } else if (EvaluationHelper.equals(tlvId, TlvType.SYSTEM_NAME)) {
      builder = TlvSystemName.staticParseLldpUnitBuilder(readBuffer, tlvIdLength);
    } else if (EvaluationHelper.equals(tlvId, TlvType.SYSTEM_DESCRIPTION)) {
      builder = TlvSystemDescription.staticParseLldpUnitBuilder(readBuffer, tlvIdLength);
    } else if (EvaluationHelper.equals(tlvId, TlvType.SYSTEM_CAPABILITIES)) {
      builder = TlvSystemCapabilities.staticParseLldpUnitBuilder(readBuffer);
    } else if (EvaluationHelper.equals(tlvId, TlvType.MANAGEMENT_ADDRESS)) {
      builder = TlvManagementAddress.staticParseLldpUnitBuilder(readBuffer);
    } else if (EvaluationHelper.equals(tlvId, TlvType.ORGANIZATION_SPECIFIC)) {
      builder = TlvOrganizationSpecific.staticParseLldpUnitBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type" + " parameters [" + "tlvId=" + tlvId + "]");
    }

    readBuffer.closeContext("LldpUnit");
    // Create the instance
    LldpUnit _lldpUnit = builder.build(tlvIdLength);
    return _lldpUnit;
  }

  public interface LldpUnitBuilder {
    LldpUnit build(short tlvIdLength);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof LldpUnit)) {
      return false;
    }
    LldpUnit that = (LldpUnit) o;
    return (getTlvIdLength() == that.getTlvIdLength()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getTlvIdLength());
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
