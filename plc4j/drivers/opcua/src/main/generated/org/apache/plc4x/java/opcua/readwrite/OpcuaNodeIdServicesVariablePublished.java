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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpcuaNodeIdServicesVariablePublished {
  PublishedDataSetType_ConfigurationVersion((int) 14519L),
  PublishedDataItemsType_PublishedData((int) 14548L),
  PublishedDataItemsType_AddVariables_InputArguments((int) 14556L),
  PublishedDataItemsType_AddVariables_OutputArguments((int) 14557L),
  PublishedDataItemsType_RemoveVariables_InputArguments((int) 14559L),
  PublishedDataItemsType_RemoveVariables_OutputArguments((int) 14560L),
  PublishedDataItemsAddVariablesMethodType_InputArguments((int) 14565L),
  PublishedDataItemsAddVariablesMethodType_OutputArguments((int) 14566L),
  PublishedDataItemsRemoveVariablesMethodType_InputArguments((int) 14568L),
  PublishedDataItemsRemoveVariablesMethodType_OutputArguments((int) 14569L),
  PublishedEventsType_PubSubEventNotifier((int) 14586L),
  PublishedEventsType_SelectedFields((int) 14587L),
  PublishedEventsType_Filter((int) 14588L),
  PublishedEventsType_ModifyFieldSelection_InputArguments((int) 15053L),
  PublishedEventsTypeModifyFieldSelectionMethodType_InputArguments((int) 15055L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Status_State((int) 15224L),
  PublishedDataSetType_DataSetMetaData((int) 15229L),
  PublishedDataSetType_ExtensionFields_AddExtensionField_InputArguments((int) 15483L),
  PublishedDataSetType_ExtensionFields_AddExtensionField_OutputArguments((int) 15484L),
  PublishedDataSetType_ExtensionFields_RemoveExtensionField_InputArguments((int) 15486L),
  PublishedEventsType_ModifyFieldSelection_OutputArguments((int) 15517L),
  PublishedEventsTypeModifyFieldSelectionMethodType_OutputArguments((int) 15518L),
  PublishedDataSetType_DataSetWriterName_Placeholder_DataSetWriterId((int) 16720L),
  PublishedDataSetType_DataSetWriterName_Placeholder_DataSetFieldContentMask((int) 16721L),
  PublishedDataSetType_DataSetWriterName_Placeholder_KeyFrameCount((int) 16731L),
  PublishedDataSetType_DataSetClassId((int) 16759L),
  PublishedDataSetType_DataSetWriterName_Placeholder_DataSetWriterProperties((int) 17482L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_DiagnosticsLevel((int) 18872L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_TotalInformation((int) 18873L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_TotalInformation_Active(
      (int) 18874L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_TotalInformation_Classification(
      (int) 18875L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_TotalInformation_DiagnosticsLevel(
      (int) 18876L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_TotalInformation_TimeFirstChange(
      (int) 18877L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_TotalError((int) 18878L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_TotalError_Active((int) 18879L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_TotalError_Classification(
      (int) 18880L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_TotalError_DiagnosticsLevel(
      (int) 18881L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_TotalError_TimeFirstChange(
      (int) 18882L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_SubError((int) 18884L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateError((int) 18886L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateError_Active(
      (int) 18887L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateError_Classification(
      (int) 18888L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateError_DiagnosticsLevel(
      (int) 18889L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateError_TimeFirstChange(
      (int) 18890L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalByMethod(
      (int) 18891L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalByMethod_Active(
      (int) 18892L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalByMethod_Classification(
      (int) 18893L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalByMethod_DiagnosticsLevel(
      (int) 18894L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalByMethod_TimeFirstChange(
      (int) 18895L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalByParent(
      (int) 18896L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalByParent_Active(
      (int) 18897L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalByParent_Classification(
      (int) 18898L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalByParent_DiagnosticsLevel(
      (int) 18899L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalByParent_TimeFirstChange(
      (int) 18900L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalFromError(
      (int) 18901L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalFromError_Active(
      (int) 18902L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalFromError_Classification(
      (int) 18903L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalFromError_DiagnosticsLevel(
      (int) 18904L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateOperationalFromError_TimeFirstChange(
      (int) 18905L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StatePausedByParent(
      (int) 18906L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StatePausedByParent_Active(
      (int) 18907L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StatePausedByParent_Classification(
      (int) 18908L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StatePausedByParent_DiagnosticsLevel(
      (int) 18909L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StatePausedByParent_TimeFirstChange(
      (int) 18910L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateDisabledByMethod(
      (int) 18911L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateDisabledByMethod_Active(
      (int) 18912L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateDisabledByMethod_Classification(
      (int) 18913L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateDisabledByMethod_DiagnosticsLevel(
      (int) 18914L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_StateDisabledByMethod_TimeFirstChange(
      (int) 18915L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_FailedDataSetMessages(
      (int) 18917L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_FailedDataSetMessages_Active(
      (int) 18918L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_FailedDataSetMessages_Classification(
      (int) 18919L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_FailedDataSetMessages_DiagnosticsLevel(
      (int) 18920L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_Counters_FailedDataSetMessages_TimeFirstChange(
      (int) 18921L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_LiveValues_MessageSequenceNumber(
      (int) 18922L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_LiveValues_MessageSequenceNumber_DiagnosticsLevel(
      (int) 18923L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_LiveValues_StatusCode(
      (int) 18924L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_LiveValues_StatusCode_DiagnosticsLevel(
      (int) 18925L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_LiveValues_MajorVersion(
      (int) 18926L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_LiveValues_MajorVersion_DiagnosticsLevel(
      (int) 18927L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_LiveValues_MinorVersion(
      (int) 18928L),
  PublishedDataSetType_DataSetWriterName_Placeholder_Diagnostics_LiveValues_MinorVersion_DiagnosticsLevel(
      (int) 18929L),
  PublishedDataSetType_CyclicDataSet((int) 25521L);
  private static final Map<Integer, OpcuaNodeIdServicesVariablePublished> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariablePublished value :
        OpcuaNodeIdServicesVariablePublished.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariablePublished(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariablePublished enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
