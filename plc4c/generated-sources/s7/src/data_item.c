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
#include <string.h>
#include <time.h>
#include <plc4c/data.h>
#include <plc4c/utils/list.h>
#include <plc4c/spi/evaluation_helper.h>
#include <plc4c/spi/types_private.h>
#include <plc4c/driver_s7_static.h>

#include "data_item.h"

// Code generated by code-generation. DO NOT EDIT.

// Parse function.
plc4c_return_code plc4c_s7_read_write_data_item_parse(plc4c_spi_read_buffer* readBuffer, char* dataProtocolId, int32_t stringLength, char* stringEncoding, plc4c_data** data_item) {
    uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
    uint16_t curPos;
    plc4c_return_code _res = OK;

        if(strcmp(dataProtocolId, "IEC61131_BOOL") == 0) { /* BOOL */

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

    } else         if(strcmp(dataProtocolId, "IEC61131_BYTE") == 0) { /* BYTE */

                // Simple Field (value)
                uint8_t value = 0;
                _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_byte_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_WORD") == 0) { /* WORD */

                // Simple Field (value)
                uint16_t value = 0;
                _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_word_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_DWORD") == 0) { /* DWORD */

                // Simple Field (value)
                uint32_t value = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_dword_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_LWORD") == 0) { /* LWORD */

                // Simple Field (value)
                uint64_t value = 0;
                _res = plc4c_spi_read_unsigned_long(readBuffer, 64, (uint64_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_lword_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_SINT") == 0) { /* SINT */

                // Simple Field (value)
                int8_t value = 0;
                _res = plc4c_spi_read_signed_byte(readBuffer, 8, (int8_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_sint_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_USINT") == 0) { /* USINT */

                // Simple Field (value)
                uint8_t value = 0;
                _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_usint_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_INT") == 0) { /* INT */

                // Simple Field (value)
                int16_t value = 0;
                _res = plc4c_spi_read_signed_short(readBuffer, 16, (int16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_int_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_UINT") == 0) { /* UINT */

                // Simple Field (value)
                uint16_t value = 0;
                _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_uint_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_DINT") == 0) { /* DINT */

                // Simple Field (value)
                int32_t value = 0;
                _res = plc4c_spi_read_signed_int(readBuffer, 32, (int32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_dint_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_UDINT") == 0) { /* UDINT */

                // Simple Field (value)
                uint32_t value = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_udint_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_LINT") == 0) { /* LINT */

                // Simple Field (value)
                int64_t value = 0;
                _res = plc4c_spi_read_signed_long(readBuffer, 64, (int64_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_lint_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_ULINT") == 0) { /* ULINT */

                // Simple Field (value)
                uint64_t value = 0;
                _res = plc4c_spi_read_unsigned_long(readBuffer, 64, (uint64_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_ulint_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_REAL") == 0) { /* REAL */

                // Simple Field (value)
                float value = 0.0f;
                _res = plc4c_spi_read_float(readBuffer, 32, (float*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_real_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_LREAL") == 0) { /* LREAL */

                // Simple Field (value)
                double value = 0.0f;
                _res = plc4c_spi_read_double(readBuffer, 64, (double*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_lreal_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_CHAR") == 0) { /* CHAR */

                // Manual Field (value)
                char* value = (char*) (plc4c_s7_read_write_parse_s7_char(readBuffer, "UTF-8", stringEncoding));

                *data_item = plc4c_data_create_char_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_WCHAR") == 0) { /* CHAR */

                // Manual Field (value)
                char* value = (char*) (plc4c_s7_read_write_parse_s7_char(readBuffer, "UTF-16", stringEncoding));

                *data_item = plc4c_data_create_char_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_STRING") == 0) { /* STRING */

                // Manual Field (value)
                char* value = (char*) (plc4c_s7_read_write_parse_s7_string(readBuffer, stringLength, "UTF-8", stringEncoding));

                *data_item = plc4c_data_create_string_data(stringLength, value);

    } else         if(strcmp(dataProtocolId, "IEC61131_WSTRING") == 0) { /* STRING */

                // Manual Field (value)
                char* value = (char*) (plc4c_s7_read_write_parse_s7_string(readBuffer, stringLength, "UTF-16", stringEncoding));

                *data_item = plc4c_data_create_string_data(stringLength, value);

    } else         if(strcmp(dataProtocolId, "IEC61131_TIME") == 0) { /* TIME */

                // Simple Field (value)
                uint32_t value = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_time_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_LTIME") == 0) { /* LTIME */

                // Simple Field (value)
                uint64_t value = 0;
                _res = plc4c_spi_read_unsigned_long(readBuffer, 64, (uint64_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_ltime_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_DATE") == 0) { /* DATE */

                // Simple Field (value)
                uint16_t value = 0;
                _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_date_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_TIME_OF_DAY") == 0) { /* TIME_OF_DAY */

                // Simple Field (value)
                uint32_t value = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &value);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_time_of_day_data(value);

    } else         if(strcmp(dataProtocolId, "IEC61131_DATE_AND_TIME") == 0) { /* DATE_AND_TIME */

                // Simple Field (year)
                uint16_t year = 0;
                _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &year);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_date_and_time_data(year);


                // Simple Field (month)
                uint8_t month = 0;
                _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &month);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_date_and_time_data(month);


                // Simple Field (day)
                uint8_t day = 0;
                _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &day);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_date_and_time_data(day);


                // Simple Field (dayOfWeek)
                uint8_t dayOfWeek = 0;
                _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &dayOfWeek);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_date_and_time_data(dayOfWeek);


                // Simple Field (hour)
                uint8_t hour = 0;
                _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &hour);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_date_and_time_data(hour);


                // Simple Field (minutes)
                uint8_t minutes = 0;
                _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &minutes);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_date_and_time_data(minutes);


                // Simple Field (seconds)
                uint8_t seconds = 0;
                _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &seconds);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_date_and_time_data(seconds);


                // Simple Field (nanos)
                uint32_t nanos = 0;
                _res = plc4c_spi_read_unsigned_int(readBuffer, 32, (uint32_t*) &nanos);
                if(_res != OK) {
                    return _res;
                }

                *data_item = plc4c_data_create_date_and_time_data(nanos);

    }

  return OK;
}

plc4c_return_code plc4c_s7_read_write_data_item_serialize(plc4c_spi_write_buffer* writeBuffer, char* dataProtocolId, int32_t stringLength, char* stringEncoding, plc4c_data** data_item) {
  plc4c_return_code _res = OK;
        if(strcmp(dataProtocolId, "IEC61131_BOOL") == 0) { /* BOOL */

                    // Reserved Field (reserved)

                    // Simple field (value)
                    _res = plc4c_spi_write_bit(writeBuffer, (*data_item)->data.bool_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_BYTE") == 0) { /* BYTE */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, (*data_item)->data.byte_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_WORD") == 0) { /* WORD */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, (*data_item)->data.word_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_DWORD") == 0) { /* DWORD */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, (*data_item)->data.dword_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_LWORD") == 0) { /* LWORD */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_long(writeBuffer, 64, (*data_item)->data.lword_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_SINT") == 0) { /* SINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_signed_byte(writeBuffer, 8, (*data_item)->data.sint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_USINT") == 0) { /* USINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, (*data_item)->data.usint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_INT") == 0) { /* INT */

                    // Simple field (value)
                    _res = plc4c_spi_write_signed_short(writeBuffer, 16, (*data_item)->data.int_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_UINT") == 0) { /* UINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, (*data_item)->data.uint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_DINT") == 0) { /* DINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_signed_int(writeBuffer, 32, (*data_item)->data.dint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_UDINT") == 0) { /* UDINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, (*data_item)->data.udint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_LINT") == 0) { /* LINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_signed_long(writeBuffer, 64, (*data_item)->data.lint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_ULINT") == 0) { /* ULINT */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_long(writeBuffer, 64, (*data_item)->data.ulint_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_REAL") == 0) { /* REAL */

                    // Simple field (value)
                    _res = plc4c_spi_write_float(writeBuffer, 32, (*data_item)->data.real_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_LREAL") == 0) { /* LREAL */

                    // Simple field (value)
                    _res = plc4c_spi_write_double(writeBuffer, 64, (*data_item)->data.lreal_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_CHAR") == 0) { /* CHAR */

                    // Manual Field (value)
        } else         if(strcmp(dataProtocolId, "IEC61131_WCHAR") == 0) { /* CHAR */

                    // Manual Field (value)
        } else         if(strcmp(dataProtocolId, "IEC61131_STRING") == 0) { /* STRING */

                    // Manual Field (value)
        } else         if(strcmp(dataProtocolId, "IEC61131_WSTRING") == 0) { /* STRING */

                    // Manual Field (value)
        } else         if(strcmp(dataProtocolId, "IEC61131_TIME") == 0) { /* TIME */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, (*data_item)->data.time_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_LTIME") == 0) { /* LTIME */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_long(writeBuffer, 64, (*data_item)->data.ltime_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_DATE") == 0) { /* DATE */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, (*data_item)->data.date_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_TIME_OF_DAY") == 0) { /* TIME_OF_DAY */

                    // Simple field (value)
                    _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, (*data_item)->data.time_of_day_value);
                    if(_res != OK) {
                        return _res;
                    }
        } else         if(strcmp(dataProtocolId, "IEC61131_DATE_AND_TIME") == 0) { /* DATE_AND_TIME */

                    // Simple field (year)
                    _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, (*data_item)->data.date_and_time_value);
                    if(_res != OK) {
                        return _res;
                    }

                    // Simple field (month)
                    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, (*data_item)->data.date_and_time_value);
                    if(_res != OK) {
                        return _res;
                    }

                    // Simple field (day)
                    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, (*data_item)->data.date_and_time_value);
                    if(_res != OK) {
                        return _res;
                    }

                    // Simple field (dayOfWeek)
                    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, (*data_item)->data.date_and_time_value);
                    if(_res != OK) {
                        return _res;
                    }

                    // Simple field (hour)
                    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, (*data_item)->data.date_and_time_value);
                    if(_res != OK) {
                        return _res;
                    }

                    // Simple field (minutes)
                    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, (*data_item)->data.date_and_time_value);
                    if(_res != OK) {
                        return _res;
                    }

                    // Simple field (seconds)
                    _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, (*data_item)->data.date_and_time_value);
                    if(_res != OK) {
                        return _res;
                    }

                    // Simple field (nanos)
                    _res = plc4c_spi_write_unsigned_int(writeBuffer, 32, (*data_item)->data.date_and_time_value);
                    if(_res != OK) {
                        return _res;
                    }
        }
  return OK;
}

uint16_t plc4c_s7_read_write_data_item_length_in_bytes(plc4c_data* data_item, char* data_protocol_id, int32_t string_length, char* string_encoding) {
  return plc4c_s7_read_write_data_item_length_in_bits(data_item, data_protocol_id, string_length, string_encoding) / 8;
}

uint16_t plc4c_s7_read_write_data_item_length_in_bits(plc4c_data* data_item, char* dataProtocolId, int32_t stringLength, char* stringEncoding) {
  uint16_t lengthInBits = 0;
    if(strcmp(dataProtocolId, "IEC61131_BOOL") == 0) { /* BOOL */

        // Reserved Field (reserved)
        lengthInBits += 7;

        // Simple field (value)
        lengthInBits += 1;
    } else     if(strcmp(dataProtocolId, "IEC61131_BYTE") == 0) { /* BYTE */

        // Simple field (value)
        lengthInBits += 8;
    } else     if(strcmp(dataProtocolId, "IEC61131_WORD") == 0) { /* WORD */

        // Simple field (value)
        lengthInBits += 16;
    } else     if(strcmp(dataProtocolId, "IEC61131_DWORD") == 0) { /* DWORD */

        // Simple field (value)
        lengthInBits += 32;
    } else     if(strcmp(dataProtocolId, "IEC61131_LWORD") == 0) { /* LWORD */

        // Simple field (value)
        lengthInBits += 64;
    } else     if(strcmp(dataProtocolId, "IEC61131_SINT") == 0) { /* SINT */

        // Simple field (value)
        lengthInBits += 8;
    } else     if(strcmp(dataProtocolId, "IEC61131_USINT") == 0) { /* USINT */

        // Simple field (value)
        lengthInBits += 8;
    } else     if(strcmp(dataProtocolId, "IEC61131_INT") == 0) { /* INT */

        // Simple field (value)
        lengthInBits += 16;
    } else     if(strcmp(dataProtocolId, "IEC61131_UINT") == 0) { /* UINT */

        // Simple field (value)
        lengthInBits += 16;
    } else     if(strcmp(dataProtocolId, "IEC61131_DINT") == 0) { /* DINT */

        // Simple field (value)
        lengthInBits += 32;
    } else     if(strcmp(dataProtocolId, "IEC61131_UDINT") == 0) { /* UDINT */

        // Simple field (value)
        lengthInBits += 32;
    } else     if(strcmp(dataProtocolId, "IEC61131_LINT") == 0) { /* LINT */

        // Simple field (value)
        lengthInBits += 64;
    } else     if(strcmp(dataProtocolId, "IEC61131_ULINT") == 0) { /* ULINT */

        // Simple field (value)
        lengthInBits += 64;
    } else     if(strcmp(dataProtocolId, "IEC61131_REAL") == 0) { /* REAL */

        // Simple field (value)
        lengthInBits += 32;
    } else     if(strcmp(dataProtocolId, "IEC61131_LREAL") == 0) { /* LREAL */

        // Simple field (value)
        lengthInBits += 64;
    } else     if(strcmp(dataProtocolId, "IEC61131_CHAR") == 0) { /* CHAR */

        // Manual Field (value)
        lengthInBits += 8;
    } else     if(strcmp(dataProtocolId, "IEC61131_WCHAR") == 0) { /* CHAR */

        // Manual Field (value)
        lengthInBits += 16;
    } else     if(strcmp(dataProtocolId, "IEC61131_STRING") == 0) { /* STRING */

        // Manual Field (value)
        lengthInBits += (((stringLength) + (2))) * (8);
    } else     if(strcmp(dataProtocolId, "IEC61131_WSTRING") == 0) { /* STRING */

        // Manual Field (value)
        lengthInBits += (((stringLength) + (2))) * (16);
    } else     if(strcmp(dataProtocolId, "IEC61131_TIME") == 0) { /* TIME */

        // Simple field (value)
        lengthInBits += 32;
    } else     if(strcmp(dataProtocolId, "IEC61131_LTIME") == 0) { /* LTIME */

        // Simple field (value)
        lengthInBits += 64;
    } else     if(strcmp(dataProtocolId, "IEC61131_DATE") == 0) { /* DATE */

        // Simple field (value)
        lengthInBits += 16;
    } else     if(strcmp(dataProtocolId, "IEC61131_TIME_OF_DAY") == 0) { /* TIME_OF_DAY */

        // Simple field (value)
        lengthInBits += 32;
    } else     if(strcmp(dataProtocolId, "IEC61131_DATE_AND_TIME") == 0) { /* DATE_AND_TIME */

        // Simple field (year)
        lengthInBits += 16;

        // Simple field (month)
        lengthInBits += 8;

        // Simple field (day)
        lengthInBits += 8;

        // Simple field (dayOfWeek)
        lengthInBits += 8;

        // Simple field (hour)
        lengthInBits += 8;

        // Simple field (minutes)
        lengthInBits += 8;

        // Simple field (seconds)
        lengthInBits += 8;

        // Simple field (nanos)
        lengthInBits += 32;
    }
  return lengthInBits;
}

