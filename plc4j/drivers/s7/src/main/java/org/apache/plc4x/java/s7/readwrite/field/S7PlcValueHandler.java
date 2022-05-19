/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.s7.readwrite.field;

import org.apache.commons.lang3.ArrayUtils;
import org.apache.plc4x.java.api.exceptions.PlcUnsupportedDataTypeException;
import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.api.value.PlcValueHandler;
import org.apache.plc4x.java.spi.values.*;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.time.Duration;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.util.Arrays;
import java.util.Objects;
import java.util.stream.Collectors;

public class S7PlcValueHandler implements PlcValueHandler {
    @Override
    public PlcValue newPlcValue(Object value) {
        return of(new Object[]{value});
    }

    @Override
    public PlcValue newPlcValue(Object[] values) {
        return of(values);
    }

    @Override
    public PlcValue newPlcValue(PlcField field, Object value) {
        return of(field, new Object[]{value});
    }

    @Override
    public PlcValue newPlcValue(PlcField field, Object[] values) {
        return of(field, values);
    }

    public static PlcValue of(Object value) {
        return of(new Object[]{value});
    }

    public static PlcValue of(Object[] values) {
        if (values.length == 1) {
            Object value = values[0];
            if (value instanceof Boolean) {
                return PlcBOOL.of(value);
            } else if (value instanceof Byte) {
                return PlcSINT.of(value);
            } else if (value instanceof Short) {
                return PlcINT.of(value);
            } else if (value instanceof Integer) {
                return PlcDINT.of(value);
            } else if (value instanceof Long) {
                return PlcLINT.of(value);
            } else if (value instanceof BigInteger) {
                return new PlcBigInteger((BigInteger) value);
            } else if (value instanceof Float) {
                return PlcREAL.of(value);
            } else if (value instanceof Double) {
                return PlcLREAL.of(value);
            } else if (value instanceof BigDecimal) {
                return new PlcBigDecimal((BigDecimal) value);
            } else if (value instanceof Duration) {
                return new PlcTIME((Duration) value);
            } else if (value instanceof LocalTime) {
                return new PlcTIME_OF_DAY((LocalTime) value);
            } else if (value instanceof LocalDate) {
                return new PlcDATE((LocalDate) value);
            } else if (value instanceof LocalDateTime) {
                return new PlcDATE_AND_TIME((LocalDateTime) value);
            } else if (value instanceof String) {
                return new PlcSTRING((String) value);
            } else if (value instanceof PlcValue) {
                return (PlcValue) value;
            } else if (value instanceof byte[]) {
                return of(ArrayUtils.toObject((byte[])value));
            } else {
                throw new PlcUnsupportedDataTypeException("Data Type " + value.getClass()
                    + " Is not supported");
            }
        } else {
            Object vo = ((Object[]) values)[0];
            if (vo instanceof PlcCHAR) {
                String v = Arrays.stream(values).map(Objects::toString).collect(Collectors.joining());
                return PlcSTRING.of(v);
            }
            if (vo instanceof PlcWCHAR) {
                String v = Arrays.stream(values).map(Objects::toString).collect(Collectors.joining());
                return PlcSTRING.of(v);
            }
            PlcList list = new PlcList();
            for (Object value : values) {
                list.add(of(new Object[]{value}));
            }
            return list;
        }
    }


    public static PlcValue of(PlcField field, Object value) {
        return of(field, new Object[]{value});
    }


    public static PlcValue of(PlcField field, Object[] values) {
        if (values.length == 1) {
            Object value = values[0];
            switch (field.getPlcDataType().toUpperCase()) {
                case "BOOL":
                case "BIT":
                    return PlcBOOL.of(value);
                case "BYTE":
                case "BITARR8":
                    if(field.getNumberOfElements()>1){
                        if(value instanceof byte[]) {
                            return of(field, ArrayUtils.toObject((byte[]) value));
                        } else if (value instanceof String && ((String)value).contains(",")){
                            return of(field, stringToByteArray((String) value));
                        }
                    }
                    return PlcBYTE.of(value);
                case "SINT":
                case "INT8":
                    return PlcSINT.of(value);
                case "USINT":
                case "UINT8":
                case "BIT8":
                    return PlcUSINT.of(value);
                case "INT":
                case "INT16":
                    return PlcINT.of(value);
                case "UINT":
                case "UINT16":
                    return PlcUINT.of(value);
                case "WORD":
                case "BITARR16":
                    return PlcWORD.of(value);
                case "DINT":
                case "INT32":
                    return PlcDINT.of(value);
                case "UDINT":
                case "UINT32":
                    return PlcUDINT.of(value);
                case "DWORD":
                case "BITARR32":
                    return PlcDWORD.of(value);
                case "LINT":
                case "INT64":
                    return PlcLINT.of(value);
                case "ULINT":
                case "UINT64":
                    return PlcULINT.of(value);
                case "LWORD":
                case "BITARR64":
                    return PlcLWORD.of(value);
                case "REAL":
                case "FLOAT":
                    return PlcREAL.of(value);
                case "LREAL":
                case "DOUBLE":
                    return PlcLREAL.of(value);
                case "CHAR": {
                    if (field.getNumberOfElements() > 1) {
                        PlcList list = new PlcList();
                        char[] charValues = value.toString().toCharArray();
                        for (Object v : charValues) {
                            list.add(PlcCHAR.of(v.toString()));
                        }
                        return list;
                    }
                    return PlcCHAR.of(value);
                }
                case "WCHAR":
                    if (field.getNumberOfElements() > 1) {
                        PlcList list = new PlcList();
                        char[] charValues = value.toString().toCharArray();
                        for (Object v : charValues) {
                            list.add(PlcWCHAR.of(v.toString()));
                        }
                        return list;
                    }
                    return PlcWCHAR.of(value);
                case "STRING":
                    return PlcSTRING.of(value);
                case "WSTRING":
                case "STRING16":
                    return PlcSTRING.of(value);
                case "TIME":
                    return PlcTIME.of(value);
                case "LTIME":
                    return PlcLTIME.of(value);
                case "DATE":
                    return PlcDATE.of(value);
                case "TIME_OF_DAY":
                    return PlcTIME_OF_DAY.of(value);
                case "DATE_AND_TIME":
                    return PlcDATE_AND_TIME.of(value);
                default:
                    return customDataType(field, new Object[]{value});
            }
        } else {
            PlcList list = new PlcList();
            for (Object value : values) {
                list.add(of(field, new Object[]{value}));
            }
            return list;
        }
    }

    public static PlcValue customDataType(PlcField field, Object[] values) {
        return of(values);
    }

    private static byte[] stringToByteArray(String stringBytes){
        String[] byteValues = stringBytes.substring(1, stringBytes.length() - 1).split(",");
        byte[] bytes = new byte[byteValues.length];

        for (int i=0, len=bytes.length; i<len; i++) {
            int intvalue = Integer.parseInt(byteValues[i].trim());
            bytes[i] = (byte)(intvalue & 0xff);
        }
        return bytes;
    }
}
