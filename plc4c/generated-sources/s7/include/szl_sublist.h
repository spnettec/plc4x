/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

#ifndef PLC4C_S7_READ_WRITE_SZL_SUBLIST_H_
#define PLC4C_S7_READ_WRITE_SZL_SUBLIST_H_

#include <stdbool.h>
#include <stdint.h>

// Code generated by code-generation. DO NOT EDIT.


#ifdef __cplusplus
extern "C" {
#endif

enum plc4c_s7_read_write_szl_sublist {
  plc4c_s7_read_write_szl_sublist_MODULE_IDENTIFICATION = 0x11,
  plc4c_s7_read_write_szl_sublist_CPU_FEATURES = 0x12,
  plc4c_s7_read_write_szl_sublist_USER_MEMORY_AREA = 0x13,
  plc4c_s7_read_write_szl_sublist_SYSTEM_AREAS = 0x14,
  plc4c_s7_read_write_szl_sublist_BLOCK_TYPES = 0x15,
  plc4c_s7_read_write_szl_sublist_STATUS_MODULE_LEDS = 0x19,
  plc4c_s7_read_write_szl_sublist_COMPONENT_IDENTIFICATION = 0x1C,
  plc4c_s7_read_write_szl_sublist_INTERRUPT_STATUS = 0x22,
  plc4c_s7_read_write_szl_sublist_ASSIGNMENT_BETWEEN_PROCESS_IMAGE_PARTITIONS_AND_OBS = 0x25,
  plc4c_s7_read_write_szl_sublist_COMMUNICATION_STATUS_DATA = 0x32,
  plc4c_s7_read_write_szl_sublist_STATUS_SINGLE_MODULE_LED = 0x74,
  plc4c_s7_read_write_szl_sublist_DP_MASTER_SYSTEM_INFORMATION = 0x90,
  plc4c_s7_read_write_szl_sublist_MODULE_STATUS_INFORMATION = 0x91,
  plc4c_s7_read_write_szl_sublist_RACK_OR_STATION_STATUS_INFORMATION = 0x92,
  plc4c_s7_read_write_szl_sublist_RACK_OR_STATION_STATUS_INFORMATION_2 = 0x94,
  plc4c_s7_read_write_szl_sublist_ADDITIONAL_DP_MASTER_SYSTEM_OR_PROFINET_IO_SYSTEM_INFORMATION = 0x95,
  plc4c_s7_read_write_szl_sublist_MODULE_STATUS_INFORMATION_PROFINET_IO_AND_PROFIBUS_DP = 0x96,
  plc4c_s7_read_write_szl_sublist_DIAGNOSTIC_BUFFER = 0xA0,
  plc4c_s7_read_write_szl_sublist_MODULE_DIAGNOSTIC_DATA = 0xB1
};
typedef enum plc4c_s7_read_write_szl_sublist plc4c_s7_read_write_szl_sublist;

// Get an empty NULL-struct
plc4c_s7_read_write_szl_sublist plc4c_s7_read_write_szl_sublist_null();

plc4c_s7_read_write_szl_sublist plc4c_s7_read_write_szl_sublist_value_of(char* value_string);

int plc4c_s7_read_write_szl_sublist_num_values();

plc4c_s7_read_write_szl_sublist plc4c_s7_read_write_szl_sublist_value_for_index(int index);

#ifdef __cplusplus
}
#endif

#endif  // PLC4C_S7_READ_WRITE_SZL_SUBLIST_H_
