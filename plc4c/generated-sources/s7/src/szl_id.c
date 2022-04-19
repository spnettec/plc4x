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

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include <plc4c/driver_s7_static.h>

#include "szl_id.h"

// Code generated by code-generation. DO NOT EDIT.


// Parse function.
plc4c_return_code plc4c_s7_read_write_szl_id_parse(plc4c_spi_read_buffer* readBuffer, plc4c_s7_read_write_szl_id** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_s7_read_write_szl_id));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Simple Field (typeClass)
  plc4c_s7_read_write_szl_module_type_class typeClass;
  _res = plc4c_s7_read_write_szl_module_type_class_parse(readBuffer, (void*) &typeClass);
  if(_res != OK) {
    return _res;
  }
  (*_message)->type_class = typeClass;

  // Simple Field (sublistExtract)
  uint8_t sublistExtract = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 4, (uint8_t*) &sublistExtract);
  if(_res != OK) {
    return _res;
  }
  (*_message)->sublist_extract = sublistExtract;

  // Simple Field (sublistList)
  plc4c_s7_read_write_szl_sublist sublistList;
  _res = plc4c_s7_read_write_szl_sublist_parse(readBuffer, (void*) &sublistList);
  if(_res != OK) {
    return _res;
  }
  (*_message)->sublist_list = sublistList;

  return OK;
}

plc4c_return_code plc4c_s7_read_write_szl_id_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_s7_read_write_szl_id* _message) {
  plc4c_return_code _res = OK;

  // Simple Field (typeClass)
  _res = plc4c_s7_read_write_szl_module_type_class_serialize(writeBuffer, &_message->type_class);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (sublistExtract)
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 4, _message->sublist_extract);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (sublistList)
  _res = plc4c_s7_read_write_szl_sublist_serialize(writeBuffer, &_message->sublist_list);
  if(_res != OK) {
    return _res;
  }

  return OK;
}

uint16_t plc4c_s7_read_write_szl_id_length_in_bytes(plc4c_s7_read_write_szl_id* _message) {
  return plc4c_s7_read_write_szl_id_length_in_bits(_message) / 8;
}

uint16_t plc4c_s7_read_write_szl_id_length_in_bits(plc4c_s7_read_write_szl_id* _message) {
  uint16_t lengthInBits = 0;

  // Simple field (typeClass)
  lengthInBits += plc4c_s7_read_write_szl_module_type_class_length_in_bits(&_message->type_class);

  // Simple field (sublistExtract)
  lengthInBits += 4;

  // Simple field (sublistList)
  lengthInBits += plc4c_s7_read_write_szl_sublist_length_in_bits(&_message->sublist_list);

  return lengthInBits;
}

