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

#include <string.h>
#include <plc4c/driver_s7_static.h>

#include "alarm_type.h"

// Code generated by code-generation. DO NOT EDIT.


// Create an empty NULL-struct
static const plc4c_s7_read_write_alarm_type plc4c_s7_read_write_alarm_type_null_const;

plc4c_s7_read_write_alarm_type plc4c_s7_read_write_alarm_type_null() {
  return plc4c_s7_read_write_alarm_type_null_const;
}

// Parse function.
plc4c_return_code plc4c_s7_read_write_alarm_type_parse(plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_alarm_type* _message) {
    plc4c_return_code _res = OK;

    uint8_t value;
    _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &value);
    *_message = plc4c_s7_read_write_alarm_type_for_value(value);

    return _res;
}

plc4c_return_code plc4c_s7_read_write_alarm_type_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_alarm_type* _message) {
    plc4c_return_code _res = OK;

    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, *_message);

    return _res;
}

plc4c_s7_read_write_alarm_type plc4c_s7_read_write_alarm_type_for_value(uint8_t value) {
    for(int i = 0; i < plc4c_s7_read_write_alarm_type_num_values(); i++) {
        if(plc4c_s7_read_write_alarm_type_value_for_index(i) == value) {
            return plc4c_s7_read_write_alarm_type_value_for_index(i);
        }
    }
    return -1;
}

plc4c_s7_read_write_alarm_type plc4c_s7_read_write_alarm_type_value_of(char* value_string) {
    if(strcmp(value_string, "SCAN") == 0) {
        return plc4c_s7_read_write_alarm_type_SCAN;
    }
    if(strcmp(value_string, "ALARM_8") == 0) {
        return plc4c_s7_read_write_alarm_type_ALARM_8;
    }
    if(strcmp(value_string, "ALARM_S") == 0) {
        return plc4c_s7_read_write_alarm_type_ALARM_S;
    }
    return -1;
}

int plc4c_s7_read_write_alarm_type_num_values() {
  return 3;
}

plc4c_s7_read_write_alarm_type plc4c_s7_read_write_alarm_type_value_for_index(int index) {
    switch(index) {
      case 0: {
        return plc4c_s7_read_write_alarm_type_SCAN;
      }
      case 1: {
        return plc4c_s7_read_write_alarm_type_ALARM_8;
      }
      case 2: {
        return plc4c_s7_read_write_alarm_type_ALARM_S;
      }
      default: {
        return -1;
      }
    }
}

uint16_t plc4c_s7_read_write_alarm_type_length_in_bytes(plc4c_s7_read_write_alarm_type* _message) {
    return plc4c_s7_read_write_alarm_type_length_in_bits(_message) / 8;
}

uint16_t plc4c_s7_read_write_alarm_type_length_in_bits(plc4c_s7_read_write_alarm_type* _message) {
    return 8;
}
