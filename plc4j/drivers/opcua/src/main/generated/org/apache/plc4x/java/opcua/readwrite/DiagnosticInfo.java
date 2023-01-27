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

public class DiagnosticInfo implements Message {

  // Properties.
  protected final boolean innerDiagnosticInfoSpecified;
  protected final boolean innerStatusCodeSpecified;
  protected final boolean additionalInfoSpecified;
  protected final boolean localeSpecified;
  protected final boolean localizedTextSpecified;
  protected final boolean namespaceURISpecified;
  protected final boolean symbolicIdSpecified;
  protected final Integer symbolicId;
  protected final Integer namespaceURI;
  protected final Integer locale;
  protected final Integer localizedText;
  protected final PascalString additionalInfo;
  protected final StatusCode innerStatusCode;
  protected final DiagnosticInfo innerDiagnosticInfo;

  public DiagnosticInfo(
      boolean innerDiagnosticInfoSpecified,
      boolean innerStatusCodeSpecified,
      boolean additionalInfoSpecified,
      boolean localeSpecified,
      boolean localizedTextSpecified,
      boolean namespaceURISpecified,
      boolean symbolicIdSpecified,
      Integer symbolicId,
      Integer namespaceURI,
      Integer locale,
      Integer localizedText,
      PascalString additionalInfo,
      StatusCode innerStatusCode,
      DiagnosticInfo innerDiagnosticInfo) {
    super();
    this.innerDiagnosticInfoSpecified = innerDiagnosticInfoSpecified;
    this.innerStatusCodeSpecified = innerStatusCodeSpecified;
    this.additionalInfoSpecified = additionalInfoSpecified;
    this.localeSpecified = localeSpecified;
    this.localizedTextSpecified = localizedTextSpecified;
    this.namespaceURISpecified = namespaceURISpecified;
    this.symbolicIdSpecified = symbolicIdSpecified;
    this.symbolicId = symbolicId;
    this.namespaceURI = namespaceURI;
    this.locale = locale;
    this.localizedText = localizedText;
    this.additionalInfo = additionalInfo;
    this.innerStatusCode = innerStatusCode;
    this.innerDiagnosticInfo = innerDiagnosticInfo;
  }

  public boolean getInnerDiagnosticInfoSpecified() {
    return innerDiagnosticInfoSpecified;
  }

  public boolean getInnerStatusCodeSpecified() {
    return innerStatusCodeSpecified;
  }

  public boolean getAdditionalInfoSpecified() {
    return additionalInfoSpecified;
  }

  public boolean getLocaleSpecified() {
    return localeSpecified;
  }

  public boolean getLocalizedTextSpecified() {
    return localizedTextSpecified;
  }

  public boolean getNamespaceURISpecified() {
    return namespaceURISpecified;
  }

  public boolean getSymbolicIdSpecified() {
    return symbolicIdSpecified;
  }

  public Integer getSymbolicId() {
    return symbolicId;
  }

  public Integer getNamespaceURI() {
    return namespaceURI;
  }

  public Integer getLocale() {
    return locale;
  }

  public Integer getLocalizedText() {
    return localizedText;
  }

  public PascalString getAdditionalInfo() {
    return additionalInfo;
  }

  public StatusCode getInnerStatusCode() {
    return innerStatusCode;
  }

  public DiagnosticInfo getInnerDiagnosticInfo() {
    return innerDiagnosticInfo;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("DiagnosticInfo");

    // Reserved Field (reserved)
    writeReservedField("reserved", (boolean) false, writeBoolean(writeBuffer));

    // Simple Field (innerDiagnosticInfoSpecified)
    writeSimpleField(
        "innerDiagnosticInfoSpecified", innerDiagnosticInfoSpecified, writeBoolean(writeBuffer));

    // Simple Field (innerStatusCodeSpecified)
    writeSimpleField(
        "innerStatusCodeSpecified", innerStatusCodeSpecified, writeBoolean(writeBuffer));

    // Simple Field (additionalInfoSpecified)
    writeSimpleField("additionalInfoSpecified", additionalInfoSpecified, writeBoolean(writeBuffer));

    // Simple Field (localeSpecified)
    writeSimpleField("localeSpecified", localeSpecified, writeBoolean(writeBuffer));

    // Simple Field (localizedTextSpecified)
    writeSimpleField("localizedTextSpecified", localizedTextSpecified, writeBoolean(writeBuffer));

    // Simple Field (namespaceURISpecified)
    writeSimpleField("namespaceURISpecified", namespaceURISpecified, writeBoolean(writeBuffer));

    // Simple Field (symbolicIdSpecified)
    writeSimpleField("symbolicIdSpecified", symbolicIdSpecified, writeBoolean(writeBuffer));

    // Optional Field (symbolicId) (Can be skipped, if the value is null)
    writeOptionalField("symbolicId", symbolicId, writeSignedInt(writeBuffer, 32));

    // Optional Field (namespaceURI) (Can be skipped, if the value is null)
    writeOptionalField("namespaceURI", namespaceURI, writeSignedInt(writeBuffer, 32));

    // Optional Field (locale) (Can be skipped, if the value is null)
    writeOptionalField("locale", locale, writeSignedInt(writeBuffer, 32));

    // Optional Field (localizedText) (Can be skipped, if the value is null)
    writeOptionalField("localizedText", localizedText, writeSignedInt(writeBuffer, 32));

    // Optional Field (additionalInfo) (Can be skipped, if the value is null)
    writeOptionalField(
        "additionalInfo", additionalInfo, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (innerStatusCode) (Can be skipped, if the value is null)
    writeOptionalField(
        "innerStatusCode", innerStatusCode, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (innerDiagnosticInfo) (Can be skipped, if the value is null)
    writeOptionalField(
        "innerDiagnosticInfo", innerDiagnosticInfo, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("DiagnosticInfo");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    DiagnosticInfo _value = this;

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Simple field (innerDiagnosticInfoSpecified)
    lengthInBits += 1;

    // Simple field (innerStatusCodeSpecified)
    lengthInBits += 1;

    // Simple field (additionalInfoSpecified)
    lengthInBits += 1;

    // Simple field (localeSpecified)
    lengthInBits += 1;

    // Simple field (localizedTextSpecified)
    lengthInBits += 1;

    // Simple field (namespaceURISpecified)
    lengthInBits += 1;

    // Simple field (symbolicIdSpecified)
    lengthInBits += 1;

    // Optional Field (symbolicId)
    if (symbolicId != null) {
      lengthInBits += 32;
    }

    // Optional Field (namespaceURI)
    if (namespaceURI != null) {
      lengthInBits += 32;
    }

    // Optional Field (locale)
    if (locale != null) {
      lengthInBits += 32;
    }

    // Optional Field (localizedText)
    if (localizedText != null) {
      lengthInBits += 32;
    }

    // Optional Field (additionalInfo)
    if (additionalInfo != null) {
      lengthInBits += additionalInfo.getLengthInBits();
    }

    // Optional Field (innerStatusCode)
    if (innerStatusCode != null) {
      lengthInBits += innerStatusCode.getLengthInBits();
    }

    // Optional Field (innerDiagnosticInfo)
    if (innerDiagnosticInfo != null) {
      lengthInBits += innerDiagnosticInfo.getLengthInBits();
    }

    return lengthInBits;
  }

  public static DiagnosticInfo staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static DiagnosticInfo staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("DiagnosticInfo");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    Boolean reservedField0 =
        readReservedField("reserved", readBoolean(readBuffer), (boolean) false);

    boolean innerDiagnosticInfoSpecified =
        readSimpleField("innerDiagnosticInfoSpecified", readBoolean(readBuffer));

    boolean innerStatusCodeSpecified =
        readSimpleField("innerStatusCodeSpecified", readBoolean(readBuffer));

    boolean additionalInfoSpecified =
        readSimpleField("additionalInfoSpecified", readBoolean(readBuffer));

    boolean localeSpecified = readSimpleField("localeSpecified", readBoolean(readBuffer));

    boolean localizedTextSpecified =
        readSimpleField("localizedTextSpecified", readBoolean(readBuffer));

    boolean namespaceURISpecified =
        readSimpleField("namespaceURISpecified", readBoolean(readBuffer));

    boolean symbolicIdSpecified = readSimpleField("symbolicIdSpecified", readBoolean(readBuffer));

    Integer symbolicId =
        readOptionalField("symbolicId", readSignedInt(readBuffer, 32), symbolicIdSpecified);

    Integer namespaceURI =
        readOptionalField("namespaceURI", readSignedInt(readBuffer, 32), namespaceURISpecified);

    Integer locale = readOptionalField("locale", readSignedInt(readBuffer, 32), localeSpecified);

    Integer localizedText =
        readOptionalField("localizedText", readSignedInt(readBuffer, 32), localizedTextSpecified);

    PascalString additionalInfo =
        readOptionalField(
            "additionalInfo",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer),
            additionalInfoSpecified);

    StatusCode innerStatusCode =
        readOptionalField(
            "innerStatusCode",
            new DataReaderComplexDefault<>(() -> StatusCode.staticParse(readBuffer), readBuffer),
            innerStatusCodeSpecified);

    DiagnosticInfo innerDiagnosticInfo =
        readOptionalField(
            "innerDiagnosticInfo",
            new DataReaderComplexDefault<>(
                () -> DiagnosticInfo.staticParse(readBuffer), readBuffer),
            innerDiagnosticInfoSpecified);

    readBuffer.closeContext("DiagnosticInfo");
    // Create the instance
    DiagnosticInfo _diagnosticInfo;
    _diagnosticInfo =
        new DiagnosticInfo(
            innerDiagnosticInfoSpecified,
            innerStatusCodeSpecified,
            additionalInfoSpecified,
            localeSpecified,
            localizedTextSpecified,
            namespaceURISpecified,
            symbolicIdSpecified,
            symbolicId,
            namespaceURI,
            locale,
            localizedText,
            additionalInfo,
            innerStatusCode,
            innerDiagnosticInfo);
    return _diagnosticInfo;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DiagnosticInfo)) {
      return false;
    }
    DiagnosticInfo that = (DiagnosticInfo) o;
    return (getInnerDiagnosticInfoSpecified() == that.getInnerDiagnosticInfoSpecified())
        && (getInnerStatusCodeSpecified() == that.getInnerStatusCodeSpecified())
        && (getAdditionalInfoSpecified() == that.getAdditionalInfoSpecified())
        && (getLocaleSpecified() == that.getLocaleSpecified())
        && (getLocalizedTextSpecified() == that.getLocalizedTextSpecified())
        && (getNamespaceURISpecified() == that.getNamespaceURISpecified())
        && (getSymbolicIdSpecified() == that.getSymbolicIdSpecified())
        && (getSymbolicId() == that.getSymbolicId())
        && (getNamespaceURI() == that.getNamespaceURI())
        && (getLocale() == that.getLocale())
        && (getLocalizedText() == that.getLocalizedText())
        && (getAdditionalInfo() == that.getAdditionalInfo())
        && (getInnerStatusCode() == that.getInnerStatusCode())
        && (getInnerDiagnosticInfo() == that.getInnerDiagnosticInfo())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getInnerDiagnosticInfoSpecified(),
        getInnerStatusCodeSpecified(),
        getAdditionalInfoSpecified(),
        getLocaleSpecified(),
        getLocalizedTextSpecified(),
        getNamespaceURISpecified(),
        getSymbolicIdSpecified(),
        getSymbolicId(),
        getNamespaceURI(),
        getLocale(),
        getLocalizedText(),
        getAdditionalInfo(),
        getInnerStatusCode(),
        getInnerDiagnosticInfo());
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
