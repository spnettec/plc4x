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

#include "data_transport_error_code.h"

// Code generated by code-generation. DO NOT EDIT.


// Create an empty NULL-struct
static const plc4c_s7_read_write_data_transport_error_code plc4c_s7_read_write_data_transport_error_code_null_const;

plc4c_s7_read_write_data_transport_error_code plc4c_s7_read_write_data_transport_error_code_null() {
  return plc4c_s7_read_write_data_transport_error_code_null_const;
}

// Parse function.
plc4c_return_code plc4c_s7_read_write_data_transport_error_code_parse(plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_data_transport_error_code* _message) {
    plc4c_return_code _res = OK;

    uint8_t value;
    _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &value);
    *_message = plc4c_s7_read_write_data_transport_error_code_for_value(value);

    return _res;
}

plc4c_return_code plc4c_s7_read_write_data_transport_error_code_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_data_transport_error_code* _message) {
    plc4c_return_code _res = OK;

    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, *_message);

    return _res;
}

plc4c_s7_read_write_data_transport_error_code plc4c_s7_read_write_data_transport_error_code_for_value(uint8_t value) {
    for(int i = 0; i < plc4c_s7_read_write_data_transport_error_code_num_values(); i++) {
        if(plc4c_s7_read_write_data_transport_error_code_value_for_index(i) == value) {
            return plc4c_s7_read_write_data_transport_error_code_value_for_index(i);
        }
    }
    return -1;
}

plc4c_s7_read_write_data_transport_error_code plc4c_s7_read_write_data_transport_error_code_value_of(char* value_string) {
    if(strcmp(value_string, "RESERVED") == 0) {
        return plc4c_s7_read_write_data_transport_error_code_RESERVED;
    }
    if(strcmp(value_string, "OK") == 0) {
        return plc4c_s7_read_write_data_transport_error_code_OK;
    }
    if(strcmp(value_string, "ACCESS_DENIED") == 0) {
        return plc4c_s7_read_write_data_transport_error_code_ACCESS_DENIED;
    }
    if(strcmp(value_string, "INVALID_ADDRESS") == 0) {
        return plc4c_s7_read_write_data_transport_error_code_INVALID_ADDRESS;
    }
    if(strcmp(value_string, "DATA_TYPE_NOT_SUPPORTED") == 0) {
        return plc4c_s7_read_write_data_transport_error_code_DATA_TYPE_NOT_SUPPORTED;
    }
    if(strcmp(value_string, "NOT_FOUND") == 0) {
        return plc4c_s7_read_write_data_transport_error_code_NOT_FOUND;
    }
    return -1;
}

int plc4c_s7_read_write_data_transport_error_code_num_values() {
  return 6;
}

plc4c_s7_read_write_data_transport_error_code plc4c_s7_read_write_data_transport_error_code_value_for_index(int index) {
    switch(index) {
      case 0: {
        return plc4c_s7_read_write_data_transport_error_code_RESERVED;
      }
      case 1: {
        return plc4c_s7_read_write_data_transport_error_code_OK;
      }
      case 2: {
        return plc4c_s7_read_write_data_transport_error_code_ACCESS_DENIED;
      }
      case 3: {
        return plc4c_s7_read_write_data_transport_error_code_INVALID_ADDRESS;
      }
      case 4: {
        return plc4c_s7_read_write_data_transport_error_code_DATA_TYPE_NOT_SUPPORTED;
      }
      case 5: {
        return plc4c_s7_read_write_data_transport_error_code_NOT_FOUND;
      }
      default: {
        return -1;
      }
    }
}

uint16_t plc4c_s7_read_write_data_transport_error_code_length_in_bytes(plc4c_s7_read_write_data_transport_error_code* _message) {
    return plc4c_s7_read_write_data_transport_error_code_length_in_bits(_message) / 8;
}

uint16_t plc4c_s7_read_write_data_transport_error_code_length_in_bits(plc4c_s7_read_write_data_transport_error_code* _message) {
    return 8;
}
