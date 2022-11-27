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

public class EphemeralKeyType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "17550";
  }

  // Properties.
  protected final PascalByteString publicKey;
  protected final PascalByteString signature;

  public EphemeralKeyType(PascalByteString publicKey, PascalByteString signature) {
    super();
    this.publicKey = publicKey;
    this.signature = signature;
  }

  public PascalByteString getPublicKey() {
    return publicKey;
  }

  public PascalByteString getSignature() {
    return signature;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("EphemeralKeyType");

    // Simple Field (publicKey)
    writeSimpleField("publicKey", publicKey, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (signature)
    writeSimpleField("signature", signature, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("EphemeralKeyType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    EphemeralKeyType _value = this;

    // Simple field (publicKey)
    lengthInBits += publicKey.getLengthInBits();

    // Simple field (signature)
    lengthInBits += signature.getLengthInBits();

    return lengthInBits;
  }

  public static EphemeralKeyTypeBuilder staticParseBuilder(ReadBuffer readBuffer, String identifier)
      throws ParseException {
    readBuffer.pullContext("EphemeralKeyType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    PascalByteString publicKey =
        readSimpleField(
            "publicKey",
            new DataReaderComplexDefault<>(
                () -> PascalByteString.staticParse(readBuffer), readBuffer));

    PascalByteString signature =
        readSimpleField(
            "signature",
            new DataReaderComplexDefault<>(
                () -> PascalByteString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("EphemeralKeyType");
    // Create the instance
    return new EphemeralKeyTypeBuilder(publicKey, signature);
  }

  public static class EphemeralKeyTypeBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalByteString publicKey;
    private final PascalByteString signature;

    public EphemeralKeyTypeBuilder(PascalByteString publicKey, PascalByteString signature) {

      this.publicKey = publicKey;
      this.signature = signature;
    }

    public EphemeralKeyType build() {
      EphemeralKeyType ephemeralKeyType = new EphemeralKeyType(publicKey, signature);
      return ephemeralKeyType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof EphemeralKeyType)) {
      return false;
    }
    EphemeralKeyType that = (EphemeralKeyType) o;
    return (getPublicKey() == that.getPublicKey())
        && (getSignature() == that.getSignature())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPublicKey(), getSignature());
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
