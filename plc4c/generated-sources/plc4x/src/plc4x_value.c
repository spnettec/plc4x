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

#include <stdio.h>
#include <string.h>
#include <time.h>
#include <plc4c/data.h>
#include <plc4c/utils/list.h>
#include <plc4c/spi/evaluation_helper.h>
#include <plc4c/spi/types_private.h>
#include <plc4c/driver_plc4x_static.h>

#include "plc4x_value.h"

// Code generated by code-generation. DO NOT EDIT.

// Parse function.
plc4c_return_code plc4c_plc4x_read_write_plc4x_value_parse(plc4x_spi_context ctx, plc4c_spi_read_buffer* readBuffer, plc4c_plc4x_read_write_plc4x_value_type valueType, plc4c_data** data_item) {
    uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
    uint16_t curPos;
    plc4c_return_code _res = OK;

        if(valueType == plc4c_plc4x_read_write_plc4x_value_type_BOOL) { /* BOOL */

                // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
                {
                    uint8_t _reserved = 0;
                    _res = plc4c_spi_read_unsigned_byte(readBuffer, 7, (uint8_t*) &_reserved);
                    if(_res != OK) {
                        return _res;
                    }
                    if(_reserved != 0x00) {
                      printf("Expected constant value '%d' but got '%d' for reserved field.", 0x00, _reserved);
                    }
                }

                // Simple Field (value)
                bool value = false;
                _res = plc4c_spi_read_bit(readBuffer, (bool*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_bool_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_BYTE) { /* BYTE */

                // Simple Field (value)
                uint8_t value = 0;
                _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_byte_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_WORD) { /* WORD */

                // Simple Field (value)
                uint16_t value = 0;
                _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_word_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DWORD) { /* DWORD */

                // Simple Field (value)
                uint32_t value = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_dword_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LWORD) { /* LWORD */

                // Simple Field (value)
                uint64_t value = 0;
                _res = plc4c_spi_read_unsigned_long(readBuffer, 64, (uint64_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_lword_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_USINT) { /* USINT */

                // Simple Field (value)
                uint8_t value = 0;
                _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_usint_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_UINT) { /* UINT */

                // Simple Field (value)
                uint16_t value = 0;
                _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_UDINT) { /* UDINT */

                // Simple Field (value)
                uint32_t value = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_udint_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_ULINT) { /* ULINT */

                // Simple Field (value)
                uint64_t value = 0;
                _res = plc4c_spi_read_unsigned_long(readBuffer, 64, (uint64_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_ulint_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_SINT) { /* SINT */

                // Simple Field (value)
                int8_t value = 0;
                _res = plc4c_spi_read_signed_byte(readBuffer, 8, (int8_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_sint_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_INT) { /* INT */

                // Simple Field (value)
                int16_t value = 0;
                _res = plc4c_spi_read_signed_short(readBuffer, 16, (int16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_int_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DINT) { /* DINT */

                // Simple Field (value)
                int32_t value = 0;
                _res = plc4c_spi_read_signed_int(readBuffer, 32, (int32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_dint_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LINT) { /* LINT */

                // Simple Field (value)
                int64_t value = 0;
                _res = plc4c_spi_read_signed_long(readBuffer, 64, (int64_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_lint_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_REAL) { /* REAL */

                // Simple Field (value)
                float value = 0.0f;
                _res = plc4c_spi_read_float(readBuffer, 32, (float*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_real_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LREAL) { /* LREAL */

                // Simple Field (value)
                double value = 0.0f;
                _res = plc4c_spi_read_double(readBuffer, 64, (double*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_lreal_data(value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_CHAR) { /* STRING */

                // Simple Field (value)
                char* value = "";
                _res = plc4c_spi_read_string(readBuffer, 8, "UTF-8", (char**) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_string_data(8, value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_WCHAR) { /* STRING */

                // Simple Field (value)
                char* value = "";
                _res = plc4c_spi_read_string(readBuffer, 16, "UTF-16", (char**) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_string_data(16, value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_STRING) { /* STRING */

                // Manual Field (value)
                char* value = (char*) (plc4c_plc4x_read_write_parse_string(readBuffer, "UTF-8"));

                *data_item = plc4c_data_create_string_data(strlen(value), value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_WSTRING) { /* STRING */

                // Manual Field (value)
                char* value = (char*) (plc4c_plc4x_read_write_parse_string(readBuffer, "UTF-16"));

                *data_item = plc4c_data_create_string_data(strlen(value), value);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_TIME) { /* TIME */

                // Simple Field (milliseconds)
                uint32_t milliseconds = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &milliseconds);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_time_data(milliseconds);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LTIME) { /* LTIME */

                // Simple Field (nanoseconds)
                uint64_t nanoseconds = 0;
                _res = plc4c_spi_read_unsigned_long(readBuffer, 64, (uint64_t*) &nanoseconds);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_ltime_data(nanoseconds);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DATE) { /* DATE */

                // Simple Field (secondsSinceEpoch)
                uint32_t secondsSinceEpoch = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &secondsSinceEpoch);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_date_data(secondsSinceEpoch);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LDATE) { /* LDATE */

                // Simple Field (nanosecondsSinceEpoch)
                uint64_t nanosecondsSinceEpoch = 0;
                _res = plc4c_spi_read_unsigned_long(readBuffer, 64, (uint64_t*) &nanosecondsSinceEpoch);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_ldate_data(nanosecondsSinceEpoch);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_TIME_OF_DAY) { /* TIME_OF_DAY */

                // Simple Field (millisecondsSinceMidnight)
                uint32_t millisecondsSinceMidnight = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &millisecondsSinceMidnight);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_time_of_day_data(millisecondsSinceMidnight);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LTIME_OF_DAY) { /* LTIME_OF_DAY */

                // Simple Field (nanosecondsSinceMidnight)
                uint64_t nanosecondsSinceMidnight = 0;
                _res = plc4c_spi_read_unsigned_long(readBuffer, 64, (uint64_t*) &nanosecondsSinceMidnight);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_ltime_of_day_data(nanosecondsSinceMidnight);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DATE_AND_TIME) { /* DATE_AND_TIME */

                // Simple Field (secondsSinceEpoch)
                uint32_t secondsSinceEpoch = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &secondsSinceEpoch);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_date_and_time_data(secondsSinceEpoch);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DATE_AND_LTIME) { /* LDATE_AND_TIME */

                // Simple Field (nanosecondsSinceEpoch)
                uint64_t nanosecondsSinceEpoch = 0;
                _res = plc4c_spi_read_unsigned_long(readBuffer, 64, (uint64_t*) &nanosecondsSinceEpoch);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_ldate_and_time_data(nanosecondsSinceEpoch);

    } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LDATE_AND_TIME) { /* LDATE_AND_TIME */

                // Simple Field (nanosecondsSinceEpoch)
                uint64_t nanosecondsSinceEpoch = 0;
                _res = plc4c_spi_read_unsigned_long(readBuffer, 64, (uint64_t*) &nanosecondsSinceEpoch);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_ldate_and_time_data(nanosecondsSinceEpoch);

    }

  return OK;
}

plc4c_return_code plc4c_plc4x_read_write_plc4x_value_serialize(plc4x_spi_context ctx, plc4c_spi_write_buffer* writeBuffer, plc4c_plc4x_read_write_plc4x_value_type valueType, plc4c_data** data_item) {
  plc4c_return_code _res = OK;
        if(valueType == plc4c_plc4x_read_write_plc4x_value_type_BOOL) { /* BOOL */

                    // Reserved Field (reserved)

                    // Simple field (value)
                    _res = plc4c_spi_write_bit(writeBuffer, (*data_item)->data.bool_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_BYTE) { /* BYTE */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, (*data_item)->data.byte_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_WORD) { /* WORD */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, (*data_item)->data.word_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DWORD) { /* DWORD */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, (*data_item)->data.dword_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LWORD) { /* LWORD */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_long(writeBuffer, 64, (*data_item)->data.lword_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_USINT) { /* USINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, (*data_item)->data.usint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_UINT) { /* UINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, (*data_item)->data.uint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_UDINT) { /* UDINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, (*data_item)->data.udint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_ULINT) { /* ULINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_long(writeBuffer, 64, (*data_item)->data.ulint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_SINT) { /* SINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_signed_byte(writeBuffer, 8, (*data_item)->data.sint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_INT) { /* INT */

                    // Simple field (value)
                    _res = plc4c_spi_write_signed_short(writeBuffer, 16, (*data_item)->data.int_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DINT) { /* DINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_signed_int(writeBuffer, 32, (*data_item)->data.dint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LINT) { /* LINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_signed_long(writeBuffer, 64, (*data_item)->data.lint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_REAL) { /* REAL */

                    // Simple field (value)
                    _res = plc4c_spi_write_float(writeBuffer, 32, (*data_item)->data.real_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LREAL) { /* LREAL */

                    // Simple field (value)
                    _res = plc4c_spi_write_double(writeBuffer, 64, (*data_item)->data.lreal_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_CHAR) { /* STRING */

                    // Simple field (value)
                    _res = plc4c_spi_write_string(writeBuffer, 8, "UTF-8", (char*) (*data_item)->data.string_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_WCHAR) { /* STRING */

                    // Simple field (value)
                    _res = plc4c_spi_write_string(writeBuffer, 16, "UTF-16", (char*) (*data_item)->data.string_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_STRING) { /* STRING */

                    // Manual Field (value)
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_WSTRING) { /* STRING */

                    // Manual Field (value)
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_TIME) { /* TIME */

                    // Simple field (milliseconds)
                    _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, (*data_item)->data.time_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LTIME) { /* LTIME */

                    // Simple field (nanoseconds)
                    _res = plc4c_spi_write_unsigned_long(writeBuffer, 64, (*data_item)->data.ltime_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DATE) { /* DATE */

                    // Simple field (secondsSinceEpoch)
                    _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, (*data_item)->data.date_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LDATE) { /* LDATE */

                    // Simple field (nanosecondsSinceEpoch)
                    _res = plc4c_spi_write_unsigned_long(writeBuffer, 64, (*data_item)->data.ldate_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_TIME_OF_DAY) { /* TIME_OF_DAY */

                    // Simple field (millisecondsSinceMidnight)
                    _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, (*data_item)->data.time_of_day_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LTIME_OF_DAY) { /* LTIME_OF_DAY */

                    // Simple field (nanosecondsSinceMidnight)
                    _res = plc4c_spi_write_unsigned_long(writeBuffer, 64, (*data_item)->data.ltime_of_day_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DATE_AND_TIME) { /* DATE_AND_TIME */

                    // Simple field (secondsSinceEpoch)
                    _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, (*data_item)->data.date_and_time_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DATE_AND_LTIME) { /* LDATE_AND_TIME */

                    // Simple field (nanosecondsSinceEpoch)
                    _res = plc4c_spi_write_unsigned_long(writeBuffer, 64, (*data_item)->data.ldate_and_time_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LDATE_AND_TIME) { /* LDATE_AND_TIME */

                    // Simple field (nanosecondsSinceEpoch)
                    _res = plc4c_spi_write_unsigned_long(writeBuffer, 64, (*data_item)->data.ldate_and_time_value);
                    if(_res != OK) {
                        return _res;
                    }
        }
  return OK;
}

uint16_t plc4c_plc4x_read_write_plc4x_value_length_in_bytes(plc4x_spi_context ctx, plc4c_data* data_item, plc4c_plc4x_read_write_plc4x_value_type value_type) {
  return plc4c_plc4x_read_write_plc4x_value_length_in_bits(ctx, data_item, value_type) / 8;
}

uint16_t plc4c_plc4x_read_write_plc4x_value_length_in_bits(plc4x_spi_context ctx, plc4c_data* data_item, plc4c_plc4x_read_write_plc4x_value_type valueType) {
  uint16_t lengthInBits = 0;
    if(valueType == plc4c_plc4x_read_write_plc4x_value_type_BOOL) { /* BOOL */

        // Reserved Field (reserved)
        lengthInBits += 7;

        // Simple field (value)
        lengthInBits += 1;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_BYTE) { /* BYTE */

        // Simple field (value)
        lengthInBits += 8;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_WORD) { /* WORD */

        // Simple field (value)
        lengthInBits += 16;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DWORD) { /* DWORD */

        // Simple field (value)
        lengthInBits += 32;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LWORD) { /* LWORD */

        // Simple field (value)
        lengthInBits += 64;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_USINT) { /* USINT */

        // Simple field (value)
        lengthInBits += 8;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_UINT) { /* UINT */

        // Simple field (value)
        lengthInBits += 16;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_UDINT) { /* UDINT */

        // Simple field (value)
        lengthInBits += 32;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_ULINT) { /* ULINT */

        // Simple field (value)
        lengthInBits += 64;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_SINT) { /* SINT */

        // Simple field (value)
        lengthInBits += 8;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_INT) { /* INT */

        // Simple field (value)
        lengthInBits += 16;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DINT) { /* DINT */

        // Simple field (value)
        lengthInBits += 32;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LINT) { /* LINT */

        // Simple field (value)
        lengthInBits += 64;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_REAL) { /* REAL */

        // Simple field (value)
        lengthInBits += 32;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LREAL) { /* LREAL */

        // Simple field (value)
        lengthInBits += 64;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_CHAR) { /* STRING */

        // Simple field (value)
        lengthInBits += 8;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_WCHAR) { /* STRING */

        // Simple field (value)
        lengthInBits += 16;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_STRING) { /* STRING */

        // Manual Field (value)
        {
            char* _value = data_item->data.string_value;
            lengthInBits += (((plc4c_spi_evaluation_helper_str_len(_value)) + (1))) * (8);
        }
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_WSTRING) { /* STRING */

        // Manual Field (value)
        {
            char* _value = data_item->data.string_value;
            lengthInBits += (((plc4c_spi_evaluation_helper_str_len(_value)) + (1))) * (16);
        }
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_TIME) { /* TIME */

        // Simple field (milliseconds)
        lengthInBits += 32;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LTIME) { /* LTIME */

        // Simple field (nanoseconds)
        lengthInBits += 64;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DATE) { /* DATE */

        // Simple field (secondsSinceEpoch)
        lengthInBits += 32;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LDATE) { /* LDATE */

        // Simple field (nanosecondsSinceEpoch)
        lengthInBits += 64;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_TIME_OF_DAY) { /* TIME_OF_DAY */

        // Simple field (millisecondsSinceMidnight)
        lengthInBits += 32;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LTIME_OF_DAY) { /* LTIME_OF_DAY */

        // Simple field (nanosecondsSinceMidnight)
        lengthInBits += 64;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DATE_AND_TIME) { /* DATE_AND_TIME */

        // Simple field (secondsSinceEpoch)
        lengthInBits += 32;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_DATE_AND_LTIME) { /* LDATE_AND_TIME */

        // Simple field (nanosecondsSinceEpoch)
        lengthInBits += 64;
    } else     if(valueType == plc4c_plc4x_read_write_plc4x_value_type_LDATE_AND_TIME) { /* LDATE_AND_TIME */

        // Simple field (nanosecondsSinceEpoch)
        lengthInBits += 64;
    }
  return lengthInBits;
}

