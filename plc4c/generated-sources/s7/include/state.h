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

#ifndef PLC4C_S7_READ_WRITE_STATE_H_
#define PLC4C_S7_READ_WRITE_STATE_H_

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>
#include <plc4c/utils/list.h>

// Code generated by code-generation. DO NOT EDIT.


struct plc4c_s7_read_write_state {
  /* Properties */
  bool si_g_8 : 1;
  bool si_g_7 : 1;
  bool si_g_6 : 1;
  bool si_g_5 : 1;
  bool si_g_4 : 1;
  bool si_g_3 : 1;
  bool si_g_2 : 1;
  bool si_g_1 : 1;
};
typedef struct plc4c_s7_read_write_state plc4c_s7_read_write_state;

// Create an empty NULL-struct
plc4c_s7_read_write_state plc4c_s7_read_write_state_null();

plc4c_return_code plc4c_s7_read_write_state_parse(plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_state** message);

plc4c_return_code plc4c_s7_read_write_state_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_state* message);

uint16_t plc4c_s7_read_write_state_length_in_bytes(plc4c_s7_read_write_state* message);

uint16_t plc4c_s7_read_write_state_length_in_bits(plc4c_s7_read_write_state* message);

#endif  // PLC4C_S7_READ_WRITE_STATE_H_
