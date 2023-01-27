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

public class DataChangeNotification extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "811";
  }

  // Properties.
  protected final int noOfMonitoredItems;
  protected final List<ExtensionObjectDefinition> monitoredItems;
  protected final int noOfDiagnosticInfos;
  protected final List<DiagnosticInfo> diagnosticInfos;

  public DataChangeNotification(
      int noOfMonitoredItems,
      List<ExtensionObjectDefinition> monitoredItems,
      int noOfDiagnosticInfos,
      List<DiagnosticInfo> diagnosticInfos) {
    super();
    this.noOfMonitoredItems = noOfMonitoredItems;
    this.monitoredItems = monitoredItems;
    this.noOfDiagnosticInfos = noOfDiagnosticInfos;
    this.diagnosticInfos = diagnosticInfos;
  }

  public int getNoOfMonitoredItems() {
    return noOfMonitoredItems;
  }

  public List<ExtensionObjectDefinition> getMonitoredItems() {
    return monitoredItems;
  }

  public int getNoOfDiagnosticInfos() {
    return noOfDiagnosticInfos;
  }

  public List<DiagnosticInfo> getDiagnosticInfos() {
    return diagnosticInfos;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("DataChangeNotification");

    // Implicit Field (notificationLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int notificationLength = (int) (getLengthInBytes());
    writeImplicitField("notificationLength", notificationLength, writeSignedInt(writeBuffer, 32));

    // Simple Field (noOfMonitoredItems)
    writeSimpleField("noOfMonitoredItems", noOfMonitoredItems, writeSignedInt(writeBuffer, 32));

    // Array Field (monitoredItems)
    writeComplexTypeArrayField("monitoredItems", monitoredItems, writeBuffer);

    // Simple Field (noOfDiagnosticInfos)
    writeSimpleField("noOfDiagnosticInfos", noOfDiagnosticInfos, writeSignedInt(writeBuffer, 32));

    // Array Field (diagnosticInfos)
    writeComplexTypeArrayField("diagnosticInfos", diagnosticInfos, writeBuffer);

    writeBuffer.popContext("DataChangeNotification");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    DataChangeNotification _value = this;

    // Implicit Field (notificationLength)
    lengthInBits += 32;

    // Simple field (noOfMonitoredItems)
    lengthInBits += 32;

    // Array field
    if (monitoredItems != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : monitoredItems) {
        boolean last = ++i >= monitoredItems.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (noOfDiagnosticInfos)
    lengthInBits += 32;

    // Array field
    if (diagnosticInfos != null) {
      int i = 0;
      for (DiagnosticInfo element : diagnosticInfos) {
        boolean last = ++i >= diagnosticInfos.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("DataChangeNotification");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    int notificationLength = readImplicitField("notificationLength", readSignedInt(readBuffer, 32));

    int noOfMonitoredItems = readSimpleField("noOfMonitoredItems", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> monitoredItems =
        readCountArrayField(
            "monitoredItems",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("808")),
                readBuffer),
            noOfMonitoredItems);

    int noOfDiagnosticInfos = readSimpleField("noOfDiagnosticInfos", readSignedInt(readBuffer, 32));

    List<DiagnosticInfo> diagnosticInfos =
        readCountArrayField(
            "diagnosticInfos",
            new DataReaderComplexDefault<>(
                () -> DiagnosticInfo.staticParse(readBuffer), readBuffer),
            noOfDiagnosticInfos);

    readBuffer.closeContext("DataChangeNotification");
    // Create the instance
    return new DataChangeNotificationBuilderImpl(
        noOfMonitoredItems, monitoredItems, noOfDiagnosticInfos, diagnosticInfos);
  }

  public static class DataChangeNotificationBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final int noOfMonitoredItems;
    private final List<ExtensionObjectDefinition> monitoredItems;
    private final int noOfDiagnosticInfos;
    private final List<DiagnosticInfo> diagnosticInfos;

    public DataChangeNotificationBuilderImpl(
        int noOfMonitoredItems,
        List<ExtensionObjectDefinition> monitoredItems,
        int noOfDiagnosticInfos,
        List<DiagnosticInfo> diagnosticInfos) {
      this.noOfMonitoredItems = noOfMonitoredItems;
      this.monitoredItems = monitoredItems;
      this.noOfDiagnosticInfos = noOfDiagnosticInfos;
      this.diagnosticInfos = diagnosticInfos;
    }

    public DataChangeNotification build() {
      DataChangeNotification dataChangeNotification =
          new DataChangeNotification(
              noOfMonitoredItems, monitoredItems, noOfDiagnosticInfos, diagnosticInfos);
      return dataChangeNotification;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof DataChangeNotification)) {
      return false;
    }
    DataChangeNotification that = (DataChangeNotification) o;
    return (getNoOfMonitoredItems() == that.getNoOfMonitoredItems())
        && (getMonitoredItems() == that.getMonitoredItems())
        && (getNoOfDiagnosticInfos() == that.getNoOfDiagnosticInfos())
        && (getDiagnosticInfos() == that.getDiagnosticInfos())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getNoOfMonitoredItems(),
        getMonitoredItems(),
        getNoOfDiagnosticInfos(),
        getDiagnosticInfos());
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
