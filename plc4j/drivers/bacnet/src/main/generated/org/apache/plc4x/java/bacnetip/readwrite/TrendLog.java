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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class TrendLog implements Message {

  // Properties.
  protected final ReadableProperty objectIdentifier;
  protected final ReadableProperty objectName;
  protected final ReadableProperty objectType;
  protected final OptionalProperty description;
  protected final WritableProperty enable;
  protected final OptionalProperty startTime;
  protected final OptionalProperty stopTime;
  protected final OptionalProperty logDeviceObjectProperty;
  protected final OptionalProperty logInterval;
  protected final OptionalProperty cOVResubscriptionInterval;
  protected final OptionalProperty clientCOVIncrement;
  protected final ReadableProperty stopWhenFull;
  protected final ReadableProperty bufferSize;
  protected final ReadableProperty logBuffer;
  protected final WritableProperty recordCount;
  protected final ReadableProperty totalRecordCount;
  protected final ReadableProperty loggingType;
  protected final OptionalProperty alignIntervals;
  protected final OptionalProperty intervalOffset;
  protected final OptionalProperty trigger;
  protected final ReadableProperty statusFlags;
  protected final OptionalProperty reliability;
  protected final OptionalProperty notificationThreshold;
  protected final OptionalProperty recordsSinceNotification;
  protected final OptionalProperty lastNotifyRecord;
  protected final ReadableProperty eventState;
  protected final OptionalProperty notificationClass;
  protected final OptionalProperty eventEnable;
  protected final OptionalProperty ackedTransitions;
  protected final OptionalProperty notifyType;
  protected final OptionalProperty eventTimeStamps;
  protected final OptionalProperty eventMessageTexts;
  protected final OptionalProperty eventMessageTextsConfig;
  protected final OptionalProperty eventDetectionEnable;
  protected final OptionalProperty eventAlgorithmInhibitRef;
  protected final OptionalProperty eventAlgorithmInhibit;
  protected final OptionalProperty reliabilityEvaluationInhibit;
  protected final ReadableProperty propertyList;
  protected final OptionalProperty tags;
  protected final OptionalProperty profileLocation;
  protected final OptionalProperty profileName;

  public TrendLog(
      ReadableProperty objectIdentifier,
      ReadableProperty objectName,
      ReadableProperty objectType,
      OptionalProperty description,
      WritableProperty enable,
      OptionalProperty startTime,
      OptionalProperty stopTime,
      OptionalProperty logDeviceObjectProperty,
      OptionalProperty logInterval,
      OptionalProperty cOVResubscriptionInterval,
      OptionalProperty clientCOVIncrement,
      ReadableProperty stopWhenFull,
      ReadableProperty bufferSize,
      ReadableProperty logBuffer,
      WritableProperty recordCount,
      ReadableProperty totalRecordCount,
      ReadableProperty loggingType,
      OptionalProperty alignIntervals,
      OptionalProperty intervalOffset,
      OptionalProperty trigger,
      ReadableProperty statusFlags,
      OptionalProperty reliability,
      OptionalProperty notificationThreshold,
      OptionalProperty recordsSinceNotification,
      OptionalProperty lastNotifyRecord,
      ReadableProperty eventState,
      OptionalProperty notificationClass,
      OptionalProperty eventEnable,
      OptionalProperty ackedTransitions,
      OptionalProperty notifyType,
      OptionalProperty eventTimeStamps,
      OptionalProperty eventMessageTexts,
      OptionalProperty eventMessageTextsConfig,
      OptionalProperty eventDetectionEnable,
      OptionalProperty eventAlgorithmInhibitRef,
      OptionalProperty eventAlgorithmInhibit,
      OptionalProperty reliabilityEvaluationInhibit,
      ReadableProperty propertyList,
      OptionalProperty tags,
      OptionalProperty profileLocation,
      OptionalProperty profileName) {
    super();
    this.objectIdentifier = objectIdentifier;
    this.objectName = objectName;
    this.objectType = objectType;
    this.description = description;
    this.enable = enable;
    this.startTime = startTime;
    this.stopTime = stopTime;
    this.logDeviceObjectProperty = logDeviceObjectProperty;
    this.logInterval = logInterval;
    this.cOVResubscriptionInterval = cOVResubscriptionInterval;
    this.clientCOVIncrement = clientCOVIncrement;
    this.stopWhenFull = stopWhenFull;
    this.bufferSize = bufferSize;
    this.logBuffer = logBuffer;
    this.recordCount = recordCount;
    this.totalRecordCount = totalRecordCount;
    this.loggingType = loggingType;
    this.alignIntervals = alignIntervals;
    this.intervalOffset = intervalOffset;
    this.trigger = trigger;
    this.statusFlags = statusFlags;
    this.reliability = reliability;
    this.notificationThreshold = notificationThreshold;
    this.recordsSinceNotification = recordsSinceNotification;
    this.lastNotifyRecord = lastNotifyRecord;
    this.eventState = eventState;
    this.notificationClass = notificationClass;
    this.eventEnable = eventEnable;
    this.ackedTransitions = ackedTransitions;
    this.notifyType = notifyType;
    this.eventTimeStamps = eventTimeStamps;
    this.eventMessageTexts = eventMessageTexts;
    this.eventMessageTextsConfig = eventMessageTextsConfig;
    this.eventDetectionEnable = eventDetectionEnable;
    this.eventAlgorithmInhibitRef = eventAlgorithmInhibitRef;
    this.eventAlgorithmInhibit = eventAlgorithmInhibit;
    this.reliabilityEvaluationInhibit = reliabilityEvaluationInhibit;
    this.propertyList = propertyList;
    this.tags = tags;
    this.profileLocation = profileLocation;
    this.profileName = profileName;
  }

  public ReadableProperty getObjectIdentifier() {
    return objectIdentifier;
  }

  public ReadableProperty getObjectName() {
    return objectName;
  }

  public ReadableProperty getObjectType() {
    return objectType;
  }

  public OptionalProperty getDescription() {
    return description;
  }

  public WritableProperty getEnable() {
    return enable;
  }

  public OptionalProperty getStartTime() {
    return startTime;
  }

  public OptionalProperty getStopTime() {
    return stopTime;
  }

  public OptionalProperty getLogDeviceObjectProperty() {
    return logDeviceObjectProperty;
  }

  public OptionalProperty getLogInterval() {
    return logInterval;
  }

  public OptionalProperty getCOVResubscriptionInterval() {
    return cOVResubscriptionInterval;
  }

  public OptionalProperty getClientCOVIncrement() {
    return clientCOVIncrement;
  }

  public ReadableProperty getStopWhenFull() {
    return stopWhenFull;
  }

  public ReadableProperty getBufferSize() {
    return bufferSize;
  }

  public ReadableProperty getLogBuffer() {
    return logBuffer;
  }

  public WritableProperty getRecordCount() {
    return recordCount;
  }

  public ReadableProperty getTotalRecordCount() {
    return totalRecordCount;
  }

  public ReadableProperty getLoggingType() {
    return loggingType;
  }

  public OptionalProperty getAlignIntervals() {
    return alignIntervals;
  }

  public OptionalProperty getIntervalOffset() {
    return intervalOffset;
  }

  public OptionalProperty getTrigger() {
    return trigger;
  }

  public ReadableProperty getStatusFlags() {
    return statusFlags;
  }

  public OptionalProperty getReliability() {
    return reliability;
  }

  public OptionalProperty getNotificationThreshold() {
    return notificationThreshold;
  }

  public OptionalProperty getRecordsSinceNotification() {
    return recordsSinceNotification;
  }

  public OptionalProperty getLastNotifyRecord() {
    return lastNotifyRecord;
  }

  public ReadableProperty getEventState() {
    return eventState;
  }

  public OptionalProperty getNotificationClass() {
    return notificationClass;
  }

  public OptionalProperty getEventEnable() {
    return eventEnable;
  }

  public OptionalProperty getAckedTransitions() {
    return ackedTransitions;
  }

  public OptionalProperty getNotifyType() {
    return notifyType;
  }

  public OptionalProperty getEventTimeStamps() {
    return eventTimeStamps;
  }

  public OptionalProperty getEventMessageTexts() {
    return eventMessageTexts;
  }

  public OptionalProperty getEventMessageTextsConfig() {
    return eventMessageTextsConfig;
  }

  public OptionalProperty getEventDetectionEnable() {
    return eventDetectionEnable;
  }

  public OptionalProperty getEventAlgorithmInhibitRef() {
    return eventAlgorithmInhibitRef;
  }

  public OptionalProperty getEventAlgorithmInhibit() {
    return eventAlgorithmInhibit;
  }

  public OptionalProperty getReliabilityEvaluationInhibit() {
    return reliabilityEvaluationInhibit;
  }

  public ReadableProperty getPropertyList() {
    return propertyList;
  }

  public OptionalProperty getTags() {
    return tags;
  }

  public OptionalProperty getProfileLocation() {
    return profileLocation;
  }

  public OptionalProperty getProfileName() {
    return profileName;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("TrendLog");

    // Simple Field (objectIdentifier)
    writeSimpleField("objectIdentifier", objectIdentifier, writeComplex(writeBuffer));

    // Simple Field (objectName)
    writeSimpleField("objectName", objectName, writeComplex(writeBuffer));

    // Simple Field (objectType)
    writeSimpleField("objectType", objectType, writeComplex(writeBuffer));

    // Simple Field (description)
    writeSimpleField("description", description, writeComplex(writeBuffer));

    // Simple Field (enable)
    writeSimpleField("enable", enable, writeComplex(writeBuffer));

    // Simple Field (startTime)
    writeSimpleField("startTime", startTime, writeComplex(writeBuffer));

    // Simple Field (stopTime)
    writeSimpleField("stopTime", stopTime, writeComplex(writeBuffer));

    // Simple Field (logDeviceObjectProperty)
    writeSimpleField("logDeviceObjectProperty", logDeviceObjectProperty, writeComplex(writeBuffer));

    // Simple Field (logInterval)
    writeSimpleField("logInterval", logInterval, writeComplex(writeBuffer));

    // Simple Field (cOVResubscriptionInterval)
    writeSimpleField(
        "cOVResubscriptionInterval", cOVResubscriptionInterval, writeComplex(writeBuffer));

    // Simple Field (clientCOVIncrement)
    writeSimpleField("clientCOVIncrement", clientCOVIncrement, writeComplex(writeBuffer));

    // Simple Field (stopWhenFull)
    writeSimpleField("stopWhenFull", stopWhenFull, writeComplex(writeBuffer));

    // Simple Field (bufferSize)
    writeSimpleField("bufferSize", bufferSize, writeComplex(writeBuffer));

    // Simple Field (logBuffer)
    writeSimpleField("logBuffer", logBuffer, writeComplex(writeBuffer));

    // Simple Field (recordCount)
    writeSimpleField("recordCount", recordCount, writeComplex(writeBuffer));

    // Simple Field (totalRecordCount)
    writeSimpleField("totalRecordCount", totalRecordCount, writeComplex(writeBuffer));

    // Simple Field (loggingType)
    writeSimpleField("loggingType", loggingType, writeComplex(writeBuffer));

    // Simple Field (alignIntervals)
    writeSimpleField("alignIntervals", alignIntervals, writeComplex(writeBuffer));

    // Simple Field (intervalOffset)
    writeSimpleField("intervalOffset", intervalOffset, writeComplex(writeBuffer));

    // Simple Field (trigger)
    writeSimpleField("trigger", trigger, writeComplex(writeBuffer));

    // Simple Field (statusFlags)
    writeSimpleField("statusFlags", statusFlags, writeComplex(writeBuffer));

    // Simple Field (reliability)
    writeSimpleField("reliability", reliability, writeComplex(writeBuffer));

    // Simple Field (notificationThreshold)
    writeSimpleField("notificationThreshold", notificationThreshold, writeComplex(writeBuffer));

    // Simple Field (recordsSinceNotification)
    writeSimpleField(
        "recordsSinceNotification", recordsSinceNotification, writeComplex(writeBuffer));

    // Simple Field (lastNotifyRecord)
    writeSimpleField("lastNotifyRecord", lastNotifyRecord, writeComplex(writeBuffer));

    // Simple Field (eventState)
    writeSimpleField("eventState", eventState, writeComplex(writeBuffer));

    // Simple Field (notificationClass)
    writeSimpleField("notificationClass", notificationClass, writeComplex(writeBuffer));

    // Simple Field (eventEnable)
    writeSimpleField("eventEnable", eventEnable, writeComplex(writeBuffer));

    // Simple Field (ackedTransitions)
    writeSimpleField("ackedTransitions", ackedTransitions, writeComplex(writeBuffer));

    // Simple Field (notifyType)
    writeSimpleField("notifyType", notifyType, writeComplex(writeBuffer));

    // Simple Field (eventTimeStamps)
    writeSimpleField("eventTimeStamps", eventTimeStamps, writeComplex(writeBuffer));

    // Simple Field (eventMessageTexts)
    writeSimpleField("eventMessageTexts", eventMessageTexts, writeComplex(writeBuffer));

    // Simple Field (eventMessageTextsConfig)
    writeSimpleField("eventMessageTextsConfig", eventMessageTextsConfig, writeComplex(writeBuffer));

    // Simple Field (eventDetectionEnable)
    writeSimpleField("eventDetectionEnable", eventDetectionEnable, writeComplex(writeBuffer));

    // Simple Field (eventAlgorithmInhibitRef)
    writeSimpleField(
        "eventAlgorithmInhibitRef", eventAlgorithmInhibitRef, writeComplex(writeBuffer));

    // Simple Field (eventAlgorithmInhibit)
    writeSimpleField("eventAlgorithmInhibit", eventAlgorithmInhibit, writeComplex(writeBuffer));

    // Simple Field (reliabilityEvaluationInhibit)
    writeSimpleField(
        "reliabilityEvaluationInhibit", reliabilityEvaluationInhibit, writeComplex(writeBuffer));

    // Simple Field (propertyList)
    writeSimpleField("propertyList", propertyList, writeComplex(writeBuffer));

    // Simple Field (tags)
    writeSimpleField("tags", tags, writeComplex(writeBuffer));

    // Simple Field (profileLocation)
    writeSimpleField("profileLocation", profileLocation, writeComplex(writeBuffer));

    // Simple Field (profileName)
    writeSimpleField("profileName", profileName, writeComplex(writeBuffer));

    writeBuffer.popContext("TrendLog");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    TrendLog _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (objectIdentifier)
    lengthInBits += objectIdentifier.getLengthInBits();

    // Simple field (objectName)
    lengthInBits += objectName.getLengthInBits();

    // Simple field (objectType)
    lengthInBits += objectType.getLengthInBits();

    // Simple field (description)
    lengthInBits += description.getLengthInBits();

    // Simple field (enable)
    lengthInBits += enable.getLengthInBits();

    // Simple field (startTime)
    lengthInBits += startTime.getLengthInBits();

    // Simple field (stopTime)
    lengthInBits += stopTime.getLengthInBits();

    // Simple field (logDeviceObjectProperty)
    lengthInBits += logDeviceObjectProperty.getLengthInBits();

    // Simple field (logInterval)
    lengthInBits += logInterval.getLengthInBits();

    // Simple field (cOVResubscriptionInterval)
    lengthInBits += cOVResubscriptionInterval.getLengthInBits();

    // Simple field (clientCOVIncrement)
    lengthInBits += clientCOVIncrement.getLengthInBits();

    // Simple field (stopWhenFull)
    lengthInBits += stopWhenFull.getLengthInBits();

    // Simple field (bufferSize)
    lengthInBits += bufferSize.getLengthInBits();

    // Simple field (logBuffer)
    lengthInBits += logBuffer.getLengthInBits();

    // Simple field (recordCount)
    lengthInBits += recordCount.getLengthInBits();

    // Simple field (totalRecordCount)
    lengthInBits += totalRecordCount.getLengthInBits();

    // Simple field (loggingType)
    lengthInBits += loggingType.getLengthInBits();

    // Simple field (alignIntervals)
    lengthInBits += alignIntervals.getLengthInBits();

    // Simple field (intervalOffset)
    lengthInBits += intervalOffset.getLengthInBits();

    // Simple field (trigger)
    lengthInBits += trigger.getLengthInBits();

    // Simple field (statusFlags)
    lengthInBits += statusFlags.getLengthInBits();

    // Simple field (reliability)
    lengthInBits += reliability.getLengthInBits();

    // Simple field (notificationThreshold)
    lengthInBits += notificationThreshold.getLengthInBits();

    // Simple field (recordsSinceNotification)
    lengthInBits += recordsSinceNotification.getLengthInBits();

    // Simple field (lastNotifyRecord)
    lengthInBits += lastNotifyRecord.getLengthInBits();

    // Simple field (eventState)
    lengthInBits += eventState.getLengthInBits();

    // Simple field (notificationClass)
    lengthInBits += notificationClass.getLengthInBits();

    // Simple field (eventEnable)
    lengthInBits += eventEnable.getLengthInBits();

    // Simple field (ackedTransitions)
    lengthInBits += ackedTransitions.getLengthInBits();

    // Simple field (notifyType)
    lengthInBits += notifyType.getLengthInBits();

    // Simple field (eventTimeStamps)
    lengthInBits += eventTimeStamps.getLengthInBits();

    // Simple field (eventMessageTexts)
    lengthInBits += eventMessageTexts.getLengthInBits();

    // Simple field (eventMessageTextsConfig)
    lengthInBits += eventMessageTextsConfig.getLengthInBits();

    // Simple field (eventDetectionEnable)
    lengthInBits += eventDetectionEnable.getLengthInBits();

    // Simple field (eventAlgorithmInhibitRef)
    lengthInBits += eventAlgorithmInhibitRef.getLengthInBits();

    // Simple field (eventAlgorithmInhibit)
    lengthInBits += eventAlgorithmInhibit.getLengthInBits();

    // Simple field (reliabilityEvaluationInhibit)
    lengthInBits += reliabilityEvaluationInhibit.getLengthInBits();

    // Simple field (propertyList)
    lengthInBits += propertyList.getLengthInBits();

    // Simple field (tags)
    lengthInBits += tags.getLengthInBits();

    // Simple field (profileLocation)
    lengthInBits += profileLocation.getLengthInBits();

    // Simple field (profileName)
    lengthInBits += profileName.getLengthInBits();

    return lengthInBits;
  }

  public static TrendLog staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("TrendLog");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ReadableProperty objectIdentifier =
        readSimpleField(
            "objectIdentifier",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetObjectIdentifier")),
                readBuffer));

    ReadableProperty objectName =
        readSimpleField(
            "objectName",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    ReadableProperty objectType =
        readSimpleField(
            "objectType",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetObjectType")),
                readBuffer));

    OptionalProperty description =
        readSimpleField(
            "description",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    WritableProperty enable =
        readSimpleField(
            "enable",
            readComplex(
                () -> WritableProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    OptionalProperty startTime =
        readSimpleField(
            "startTime",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetDateTime")),
                readBuffer));

    OptionalProperty stopTime =
        readSimpleField(
            "stopTime",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetDateTime")),
                readBuffer));

    OptionalProperty logDeviceObjectProperty =
        readSimpleField(
            "logDeviceObjectProperty",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetDeviceObjectPropertyReference")),
                readBuffer));

    OptionalProperty logInterval =
        readSimpleField(
            "logInterval",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty cOVResubscriptionInterval =
        readSimpleField(
            "cOVResubscriptionInterval",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty clientCOVIncrement =
        readSimpleField(
            "clientCOVIncrement",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetClientCOV")),
                readBuffer));

    ReadableProperty stopWhenFull =
        readSimpleField(
            "stopWhenFull",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    ReadableProperty bufferSize =
        readSimpleField(
            "bufferSize",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("Unsigned32")),
                readBuffer));

    ReadableProperty logBuffer =
        readSimpleField(
            "logBuffer",
            readComplex(
                () ->
                    ReadableProperty.staticParse(
                        readBuffer, (String) ("BACnetLIST of BACnetLogRecord")),
                readBuffer));

    WritableProperty recordCount =
        readSimpleField(
            "recordCount",
            readComplex(
                () -> WritableProperty.staticParse(readBuffer, (String) ("Unsigned32")),
                readBuffer));

    ReadableProperty totalRecordCount =
        readSimpleField(
            "totalRecordCount",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("Unsigned32")),
                readBuffer));

    ReadableProperty loggingType =
        readSimpleField(
            "loggingType",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetLoggingType")),
                readBuffer));

    OptionalProperty alignIntervals =
        readSimpleField(
            "alignIntervals",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    OptionalProperty intervalOffset =
        readSimpleField(
            "intervalOffset",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty trigger =
        readSimpleField(
            "trigger",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    ReadableProperty statusFlags =
        readSimpleField(
            "statusFlags",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetStatusFlags")),
                readBuffer));

    OptionalProperty reliability =
        readSimpleField(
            "reliability",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetReliability")),
                readBuffer));

    OptionalProperty notificationThreshold =
        readSimpleField(
            "notificationThreshold",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned32")),
                readBuffer));

    OptionalProperty recordsSinceNotification =
        readSimpleField(
            "recordsSinceNotification",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned32")),
                readBuffer));

    OptionalProperty lastNotifyRecord =
        readSimpleField(
            "lastNotifyRecord",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned32")),
                readBuffer));

    ReadableProperty eventState =
        readSimpleField(
            "eventState",
            readComplex(
                () -> ReadableProperty.staticParse(readBuffer, (String) ("BACnetEventState")),
                readBuffer));

    OptionalProperty notificationClass =
        readSimpleField(
            "notificationClass",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("Unsigned")), readBuffer));

    OptionalProperty eventEnable =
        readSimpleField(
            "eventEnable",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetEventTransitionBits")),
                readBuffer));

    OptionalProperty ackedTransitions =
        readSimpleField(
            "ackedTransitions",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetEventTransitionBits")),
                readBuffer));

    OptionalProperty notifyType =
        readSimpleField(
            "notifyType",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BACnetNotifyType")),
                readBuffer));

    OptionalProperty eventTimeStamps =
        readSimpleField(
            "eventTimeStamps",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[3] of BACnetTimeStamp")),
                readBuffer));

    OptionalProperty eventMessageTexts =
        readSimpleField(
            "eventMessageTexts",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[3] of CharacterString")),
                readBuffer));

    OptionalProperty eventMessageTextsConfig =
        readSimpleField(
            "eventMessageTextsConfig",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[3] of CharacterString")),
                readBuffer));

    OptionalProperty eventDetectionEnable =
        readSimpleField(
            "eventDetectionEnable",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    OptionalProperty eventAlgorithmInhibitRef =
        readSimpleField(
            "eventAlgorithmInhibitRef",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetObjectPropertyReference")),
                readBuffer));

    OptionalProperty eventAlgorithmInhibit =
        readSimpleField(
            "eventAlgorithmInhibit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    OptionalProperty reliabilityEvaluationInhibit =
        readSimpleField(
            "reliabilityEvaluationInhibit",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("BOOLEAN")), readBuffer));

    ReadableProperty propertyList =
        readSimpleField(
            "propertyList",
            readComplex(
                () ->
                    ReadableProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[N] of BACnetPropertyIdentifier")),
                readBuffer));

    OptionalProperty tags =
        readSimpleField(
            "tags",
            readComplex(
                () ->
                    OptionalProperty.staticParse(
                        readBuffer, (String) ("BACnetARRAY[N] of BACnetNameValue")),
                readBuffer));

    OptionalProperty profileLocation =
        readSimpleField(
            "profileLocation",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    OptionalProperty profileName =
        readSimpleField(
            "profileName",
            readComplex(
                () -> OptionalProperty.staticParse(readBuffer, (String) ("CharacterString")),
                readBuffer));

    readBuffer.closeContext("TrendLog");
    // Create the instance
    TrendLog _trendLog;
    _trendLog =
        new TrendLog(
            objectIdentifier,
            objectName,
            objectType,
            description,
            enable,
            startTime,
            stopTime,
            logDeviceObjectProperty,
            logInterval,
            cOVResubscriptionInterval,
            clientCOVIncrement,
            stopWhenFull,
            bufferSize,
            logBuffer,
            recordCount,
            totalRecordCount,
            loggingType,
            alignIntervals,
            intervalOffset,
            trigger,
            statusFlags,
            reliability,
            notificationThreshold,
            recordsSinceNotification,
            lastNotifyRecord,
            eventState,
            notificationClass,
            eventEnable,
            ackedTransitions,
            notifyType,
            eventTimeStamps,
            eventMessageTexts,
            eventMessageTextsConfig,
            eventDetectionEnable,
            eventAlgorithmInhibitRef,
            eventAlgorithmInhibit,
            reliabilityEvaluationInhibit,
            propertyList,
            tags,
            profileLocation,
            profileName);
    return _trendLog;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TrendLog)) {
      return false;
    }
    TrendLog that = (TrendLog) o;
    return (getObjectIdentifier() == that.getObjectIdentifier())
        && (getObjectName() == that.getObjectName())
        && (getObjectType() == that.getObjectType())
        && (getDescription() == that.getDescription())
        && (getEnable() == that.getEnable())
        && (getStartTime() == that.getStartTime())
        && (getStopTime() == that.getStopTime())
        && (getLogDeviceObjectProperty() == that.getLogDeviceObjectProperty())
        && (getLogInterval() == that.getLogInterval())
        && (getCOVResubscriptionInterval() == that.getCOVResubscriptionInterval())
        && (getClientCOVIncrement() == that.getClientCOVIncrement())
        && (getStopWhenFull() == that.getStopWhenFull())
        && (getBufferSize() == that.getBufferSize())
        && (getLogBuffer() == that.getLogBuffer())
        && (getRecordCount() == that.getRecordCount())
        && (getTotalRecordCount() == that.getTotalRecordCount())
        && (getLoggingType() == that.getLoggingType())
        && (getAlignIntervals() == that.getAlignIntervals())
        && (getIntervalOffset() == that.getIntervalOffset())
        && (getTrigger() == that.getTrigger())
        && (getStatusFlags() == that.getStatusFlags())
        && (getReliability() == that.getReliability())
        && (getNotificationThreshold() == that.getNotificationThreshold())
        && (getRecordsSinceNotification() == that.getRecordsSinceNotification())
        && (getLastNotifyRecord() == that.getLastNotifyRecord())
        && (getEventState() == that.getEventState())
        && (getNotificationClass() == that.getNotificationClass())
        && (getEventEnable() == that.getEventEnable())
        && (getAckedTransitions() == that.getAckedTransitions())
        && (getNotifyType() == that.getNotifyType())
        && (getEventTimeStamps() == that.getEventTimeStamps())
        && (getEventMessageTexts() == that.getEventMessageTexts())
        && (getEventMessageTextsConfig() == that.getEventMessageTextsConfig())
        && (getEventDetectionEnable() == that.getEventDetectionEnable())
        && (getEventAlgorithmInhibitRef() == that.getEventAlgorithmInhibitRef())
        && (getEventAlgorithmInhibit() == that.getEventAlgorithmInhibit())
        && (getReliabilityEvaluationInhibit() == that.getReliabilityEvaluationInhibit())
        && (getPropertyList() == that.getPropertyList())
        && (getTags() == that.getTags())
        && (getProfileLocation() == that.getProfileLocation())
        && (getProfileName() == that.getProfileName())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getObjectIdentifier(),
        getObjectName(),
        getObjectType(),
        getDescription(),
        getEnable(),
        getStartTime(),
        getStopTime(),
        getLogDeviceObjectProperty(),
        getLogInterval(),
        getCOVResubscriptionInterval(),
        getClientCOVIncrement(),
        getStopWhenFull(),
        getBufferSize(),
        getLogBuffer(),
        getRecordCount(),
        getTotalRecordCount(),
        getLoggingType(),
        getAlignIntervals(),
        getIntervalOffset(),
        getTrigger(),
        getStatusFlags(),
        getReliability(),
        getNotificationThreshold(),
        getRecordsSinceNotification(),
        getLastNotifyRecord(),
        getEventState(),
        getNotificationClass(),
        getEventEnable(),
        getAckedTransitions(),
        getNotifyType(),
        getEventTimeStamps(),
        getEventMessageTexts(),
        getEventMessageTextsConfig(),
        getEventDetectionEnable(),
        getEventAlgorithmInhibitRef(),
        getEventAlgorithmInhibit(),
        getReliabilityEvaluationInhibit(),
        getPropertyList(),
        getTags(),
        getProfileLocation(),
        getProfileName());
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
