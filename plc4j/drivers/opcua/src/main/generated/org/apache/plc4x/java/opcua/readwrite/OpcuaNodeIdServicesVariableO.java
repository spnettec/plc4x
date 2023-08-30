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

public enum OpcuaNodeIdServicesVariableO {
  OPCUANamespaceMetadata_NamespaceUri((int) 15958L),
  OPCUANamespaceMetadata_NamespaceVersion((int) 15959L),
  OPCUANamespaceMetadata_NamespacePublicationDate((int) 15960L),
  OPCUANamespaceMetadata_IsNamespaceSubset((int) 15961L),
  OPCUANamespaceMetadata_StaticNodeIdTypes((int) 15962L),
  OPCUANamespaceMetadata_StaticNumericNodeIdRange((int) 15963L),
  OPCUANamespaceMetadata_StaticStringNodeIdPattern((int) 15964L),
  OPCUANamespaceMetadata_NamespaceFile_Size((int) 15966L),
  OPCUANamespaceMetadata_NamespaceFile_Writable((int) 15967L),
  OPCUANamespaceMetadata_NamespaceFile_UserWritable((int) 15968L),
  OPCUANamespaceMetadata_NamespaceFile_OpenCount((int) 15969L),
  OPCUANamespaceMetadata_NamespaceFile_MimeType((int) 15970L),
  OPCUANamespaceMetadata_NamespaceFile_Open_InputArguments((int) 15972L),
  OPCUANamespaceMetadata_NamespaceFile_Open_OutputArguments((int) 15973L),
  OPCUANamespaceMetadata_NamespaceFile_Close_InputArguments((int) 15975L),
  OPCUANamespaceMetadata_NamespaceFile_Read_InputArguments((int) 15977L),
  OPCUANamespaceMetadata_NamespaceFile_Read_OutputArguments((int) 15978L),
  OPCUANamespaceMetadata_NamespaceFile_Write_InputArguments((int) 15980L),
  OPCUANamespaceMetadata_NamespaceFile_GetPosition_InputArguments((int) 15982L),
  OPCUANamespaceMetadata_NamespaceFile_GetPosition_OutputArguments((int) 15983L),
  OPCUANamespaceMetadata_NamespaceFile_SetPosition_InputArguments((int) 15985L),
  OPCUANamespaceMetadata_DefaultRolePermissions((int) 16134L),
  OPCUANamespaceMetadata_DefaultUserRolePermissions((int) 16135L),
  OPCUANamespaceMetadata_DefaultAccessRestrictions((int) 16136L),
  OPCUANamespaceMetadata_NamespaceFile_MaxByteStringLength((int) 24243L),
  OPCUANamespaceMetadata_NamespaceFile_LastModifiedTime((int) 25199L),
  OPCUANamespaceMetadata_ConfigurationVersion((int) 25266L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableO> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableO value : OpcuaNodeIdServicesVariableO.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableO(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableO enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
