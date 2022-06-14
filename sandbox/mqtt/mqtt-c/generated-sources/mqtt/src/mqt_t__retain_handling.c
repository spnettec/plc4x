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

#include "mqt_t__retain_handling.h"
#include <string.h>

// Code generated by code-generation. DO NOT EDIT.


// Create an empty NULL-struct
static const plc4c_mqtt_read_write_mqt_t__retain_handling plc4c_mqtt_read_write_mqt_t__retain_handling_null_const;

plc4c_mqtt_read_write_mqt_t__retain_handling plc4c_mqtt_read_write_mqt_t__retain_handling_null() {
  return plc4c_mqtt_read_write_mqt_t__retain_handling_null_const;
}

// Parse function.
plc4c_return_code plc4c_mqtt_read_write_mqt_t__retain_handling_parse(plc4c_spi_read_buffer* readBuffer, plc4c_mqtt_read_write_mqt_t__retain_handling** _message) {
    plc4c_return_code _res = OK;

    // Allocate enough memory to contain this data structure.
    (*_message) = malloc(sizeof(plc4c_mqtt_read_write_mqt_t__retain_handling));
    if(*_message == NULL) {
        return NO_MEMORY;
    }

    _res = plc4c_spi_read_unsigned_byte(readBuffer, 2, (uint8_t*) *_message);

    return _res;
}

plc4c_return_code plc4c_mqtt_read_write_mqt_t__retain_handling_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_mqtt_read_write_mqt_t__retain_handling* _message) {
    plc4c_return_code _res = OK;

    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 2, *_message);

    return _res;
}

plc4c_mqtt_read_write_mqt_t__retain_handling plc4c_mqtt_read_write_mqt_t__retain_handling_value_of(char* value_string) {
    if(strcmp(value_string, "SEND_RETAINED_MESSAGES_AT_THE_TIME_OF_THE_SUBSCRIBE") == 0) {
        return plc4c_mqtt_read_write_mqt_t__retain_handling_SEND_RETAINED_MESSAGES_AT_THE_TIME_OF_THE_SUBSCRIBE;
    }
    if(strcmp(value_string, "SEND_RETAINED_MESSAGES_AT_SUBSCRIBE_ONLY_IF_THE_SUBSCRIPTION_DOES_NOT_CURRENTLY_EXIST") == 0) {
        return plc4c_mqtt_read_write_mqt_t__retain_handling_SEND_RETAINED_MESSAGES_AT_SUBSCRIBE_ONLY_IF_THE_SUBSCRIPTION_DOES_NOT_CURRENTLY_EXIST;
    }
    if(strcmp(value_string, "DO_NOT_SEND_RETAINED_MESSAGES_AT_THE_TIME_OF_SUBSCRIBE") == 0) {
        return plc4c_mqtt_read_write_mqt_t__retain_handling_DO_NOT_SEND_RETAINED_MESSAGES_AT_THE_TIME_OF_SUBSCRIBE;
    }
    return -1;
}

int plc4c_mqtt_read_write_mqt_t__retain_handling_num_values() {
  return 3;
}

plc4c_mqtt_read_write_mqt_t__retain_handling plc4c_mqtt_read_write_mqt_t__retain_handling_value_for_index(int index) {
    switch(index) {
      case 0: {
        return plc4c_mqtt_read_write_mqt_t__retain_handling_SEND_RETAINED_MESSAGES_AT_THE_TIME_OF_THE_SUBSCRIBE;
      }
      case 1: {
        return plc4c_mqtt_read_write_mqt_t__retain_handling_SEND_RETAINED_MESSAGES_AT_SUBSCRIBE_ONLY_IF_THE_SUBSCRIPTION_DOES_NOT_CURRENTLY_EXIST;
      }
      case 2: {
        return plc4c_mqtt_read_write_mqt_t__retain_handling_DO_NOT_SEND_RETAINED_MESSAGES_AT_THE_TIME_OF_SUBSCRIBE;
      }
      default: {
        return -1;
      }
    }
}

uint16_t plc4c_mqtt_read_write_mqt_t__retain_handling_length_in_bytes(plc4c_mqtt_read_write_mqt_t__retain_handling* _message) {
    return plc4c_mqtt_read_write_mqt_t__retain_handling_length_in_bits(_message) / 8;
}

uint16_t plc4c_mqtt_read_write_mqt_t__retain_handling_length_in_bits(plc4c_mqtt_read_write_mqt_t__retain_handling* _message) {
    return 2;
}
