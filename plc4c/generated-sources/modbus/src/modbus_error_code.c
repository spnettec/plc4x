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

#include "modbus_error_code.h"
#include <string.h>

// Code generated by code-generation. DO NOT EDIT.


// Create an empty NULL-struct
static const plc4c_modbus_read_write_modbus_error_code plc4c_modbus_read_write_modbus_error_code_null_const;

plc4c_modbus_read_write_modbus_error_code plc4c_modbus_read_write_modbus_error_code_null() {
  return plc4c_modbus_read_write_modbus_error_code_null_const;
}

// Parse function.
plc4c_return_code plc4c_modbus_read_write_modbus_error_code_parse(plc4c_spi_read_buffer* readBuffer, plc4c_modbus_read_write_modbus_error_code** _message) {
    plc4c_return_code _res = OK;

    // Allocate enough memory to contain this data structure.
    (*_message) = malloc(sizeof(plc4c_modbus_read_write_modbus_error_code));
    if(*_message == NULL) {
        return NO_MEMORY;
    }

    _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) *_message);

    return _res;
}

plc4c_return_code plc4c_modbus_read_write_modbus_error_code_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_modbus_read_write_modbus_error_code* _message) {
    plc4c_return_code _res = OK;

    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, *_message);

    return _res;
}

plc4c_modbus_read_write_modbus_error_code plc4c_modbus_read_write_modbus_error_code_value_of(char* value_string) {
    if(strcmp(value_string, "ILLEGAL_FUNCTION") == 0) {
        return plc4c_modbus_read_write_modbus_error_code_ILLEGAL_FUNCTION;
    }
    if(strcmp(value_string, "ILLEGAL_DATA_ADDRESS") == 0) {
        return plc4c_modbus_read_write_modbus_error_code_ILLEGAL_DATA_ADDRESS;
    }
    if(strcmp(value_string, "ILLEGAL_DATA_VALUE") == 0) {
        return plc4c_modbus_read_write_modbus_error_code_ILLEGAL_DATA_VALUE;
    }
    if(strcmp(value_string, "SLAVE_DEVICE_FAILURE") == 0) {
        return plc4c_modbus_read_write_modbus_error_code_SLAVE_DEVICE_FAILURE;
    }
    if(strcmp(value_string, "ACKNOWLEDGE") == 0) {
        return plc4c_modbus_read_write_modbus_error_code_ACKNOWLEDGE;
    }
    if(strcmp(value_string, "SLAVE_DEVICE_BUSY") == 0) {
        return plc4c_modbus_read_write_modbus_error_code_SLAVE_DEVICE_BUSY;
    }
    if(strcmp(value_string, "NEGATIVE_ACKNOWLEDGE") == 0) {
        return plc4c_modbus_read_write_modbus_error_code_NEGATIVE_ACKNOWLEDGE;
    }
    if(strcmp(value_string, "MEMORY_PARITY_ERROR") == 0) {
        return plc4c_modbus_read_write_modbus_error_code_MEMORY_PARITY_ERROR;
    }
    if(strcmp(value_string, "GATEWAY_PATH_UNAVAILABLE") == 0) {
        return plc4c_modbus_read_write_modbus_error_code_GATEWAY_PATH_UNAVAILABLE;
    }
    if(strcmp(value_string, "GATEWAY_TARGET_DEVICE_FAILED_TO_RESPOND") == 0) {
        return plc4c_modbus_read_write_modbus_error_code_GATEWAY_TARGET_DEVICE_FAILED_TO_RESPOND;
    }
    return -1;
}

int plc4c_modbus_read_write_modbus_error_code_num_values() {
  return 10;
}

plc4c_modbus_read_write_modbus_error_code plc4c_modbus_read_write_modbus_error_code_value_for_index(int index) {
    switch(index) {
      case 0: {
        return plc4c_modbus_read_write_modbus_error_code_ILLEGAL_FUNCTION;
      }
      case 1: {
        return plc4c_modbus_read_write_modbus_error_code_ILLEGAL_DATA_ADDRESS;
      }
      case 2: {
        return plc4c_modbus_read_write_modbus_error_code_ILLEGAL_DATA_VALUE;
      }
      case 3: {
        return plc4c_modbus_read_write_modbus_error_code_SLAVE_DEVICE_FAILURE;
      }
      case 4: {
        return plc4c_modbus_read_write_modbus_error_code_ACKNOWLEDGE;
      }
      case 5: {
        return plc4c_modbus_read_write_modbus_error_code_SLAVE_DEVICE_BUSY;
      }
      case 6: {
        return plc4c_modbus_read_write_modbus_error_code_NEGATIVE_ACKNOWLEDGE;
      }
      case 7: {
        return plc4c_modbus_read_write_modbus_error_code_MEMORY_PARITY_ERROR;
      }
      case 8: {
        return plc4c_modbus_read_write_modbus_error_code_GATEWAY_PATH_UNAVAILABLE;
      }
      case 9: {
        return plc4c_modbus_read_write_modbus_error_code_GATEWAY_TARGET_DEVICE_FAILED_TO_RESPOND;
      }
      default: {
        return -1;
      }
    }
}

uint16_t plc4c_modbus_read_write_modbus_error_code_length_in_bytes(plc4c_modbus_read_write_modbus_error_code* _message) {
    return plc4c_modbus_read_write_modbus_error_code_length_in_bits(_message) / 8;
}

uint16_t plc4c_modbus_read_write_modbus_error_code_length_in_bits(plc4c_modbus_read_write_modbus_error_code* _message) {
    return 8;
}
