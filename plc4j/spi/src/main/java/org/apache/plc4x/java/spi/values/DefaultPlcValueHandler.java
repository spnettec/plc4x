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
package org.apache.plc4x.java.spi.values;

import org.apache.commons.lang3.ArrayUtils;
import org.apache.commons.lang3.NotImplementedException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.exceptions.PlcUnsupportedDataTypeException;
import org.apache.plc4x.java.api.model.ArrayInfo;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcValueType;
import org.apache.plc4x.java.api.value.PlcValue;

import java.lang.reflect.Array;
import java.math.BigDecimal;
import java.math.BigInteger;
import java.time.Duration;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.util.*;
import java.util.stream.Collectors;

public class DefaultPlcValueHandler implements PlcValueHandler {

    public PlcValue newPlcValue(Object value) {
        return of(new Object[]{value});
    }

    public PlcValue newPlcValue(Object[] values) {
        return of(values);
    }

    @Override
    public PlcValue newPlcValue(PlcTag tag, Object value) {
        return of(tag, new Object[]{value});
    }

    @Override
    public PlcValue newPlcValue(PlcTag tag, Object[] values) {
        return of(tag, values);
    }

    public static PlcValue of(Object value) {
        return of(new Object[]{value});
    }

    public static PlcValue of(List<?> value) {
        return of(value.toArray());
    }
    public static PlcValue of(PlcValueType valueType, Object value) {
        if(value instanceof  PlcValue){
            PlcValue originValue = (PlcValue)value;
            if(originValue.getPlcValueType()==valueType)
            {
                return originValue;
            }
            value = originValue.getObject();
        }
        return typeValue(valueType,value);
    }
    public static PlcValue of(Object[] values) {
        if (values.length != 1) {
            Object vo = ((Object[]) values)[0];
            if (vo instanceof PlcCHAR) {
                String v = Arrays.stream(values).map(Objects::toString).collect(Collectors.joining());
                return PlcSTRING.of(v);
            }
            if (vo instanceof PlcWCHAR) {
                String v = Arrays.stream(values).map(Objects::toString).collect(Collectors.joining());
                return PlcSTRING.of(v);
            }
            if (vo instanceof Byte) {
                return PlcRawByteArray.of(getByteArray(values));
            }
            PlcList list = new PlcList();
            for (Object value : values) {
                list.add(of(new Object[]{value}));
            }
            return list;
        }
        Object value = values[0];
        if (value instanceof Boolean) {
            return PlcBOOL.of(value);
        }
        if (value instanceof Byte) {
            return PlcSINT.of(value);
        }
        if (value instanceof byte[]) {
            return PlcRawByteArray.of(value);
        }
        if (value instanceof Short) {
            return PlcINT.of(value);
        }
        if (value instanceof Integer) {
            return PlcDINT.of(value);
        }
        if (value instanceof Long) {
            return PlcLINT.of(value);
        }
        if (value instanceof BigInteger) {
            return PlcLINT.of(value);
        }
        if (value instanceof BigDecimal) {
            return new PlcLINT((BigDecimal) value);
        }
        if (value instanceof Float) {
            return PlcREAL.of(value);
        }
        if (value instanceof Double) {
            return PlcLREAL.of(value);
        }
        if (value instanceof Duration) {
            return new PlcTIME((Duration) value);
        }
        if (value instanceof LocalTime) {
            return new PlcTIME_OF_DAY((LocalTime) value);
        }
        if (value instanceof LocalDate) {
            return new PlcDATE((LocalDate) value);
        }
        if (value instanceof LocalDateTime) {
            return new PlcDATE_AND_TIME((LocalDateTime) value);
        }
        if (value instanceof String) {
            return new PlcSTRING((String) value);
        }
        if (value instanceof PlcValue) {
            return (PlcValue) value;
        }
        throw new PlcUnsupportedDataTypeException("Data Type " + value.getClass()
            + " Is not supported");
    }


    public static PlcValue of(PlcTag tag, Object value) {
        return of(tag, new Object[]{value});
    }


    public static PlcValue of(PlcTag tag, Object[] values) {
        if (values.length == 1) {
            Object value = values[0];
            if (tag.getPlcValueType() == null) {
                // TODO: This is a hacky shortcut ..
                if (value instanceof PlcValue) {
                    return (PlcValue) value;
                }
                return customDataType(getArray(value));
            }
            if (value instanceof PlcValue) {
                value = ((PlcValue) value).getObject();
            }
            if (tag.getNumberOfElements() > 1) {
                if(value==null){
                    value = new Object[tag.getNumberOfElements()];
                }
                int len = 1;
                try {
                    len = ArrayUtils.getLength(value);
                } catch (IllegalArgumentException ignored){}
                if (len>1) {
                    if(value instanceof byte[]){
                        return of(tag, ArrayUtils.toObject((byte[])value));
                    }
                    if(value instanceof int[]){
                        return of(tag, ArrayUtils.toObject((int[])value));
                    }
                    if(value instanceof float[]){
                        return of(tag, ArrayUtils.toObject((float[])value));
                    }
                    if(value instanceof double[]){
                        return of(tag, ArrayUtils.toObject((double[])value));
                    }
                    if(value instanceof long[]){
                        return of(tag, ArrayUtils.toObject((long[])value));
                    }
                    if(value instanceof short[]){
                        return of(tag, ArrayUtils.toObject((short[])value));
                    }
                    if(value instanceof char[]){
                        return of(tag, ArrayUtils.toObject((char[])value));
                    }
                    return of(tag, (Object[])value);
                } else if (value instanceof String && ((String) value).contains(",")) {
                    switch (tag.getPlcValueType()) {
                        case BYTE:
                            return of(tag, stringToByteArray((String) value));
                        default: return of(tag, stringToStringArray((String) value));
                    }
                }
            }
            if (value instanceof PlcValue) {
                PlcValue plcValue = (PlcValue) value;
                if (plcValue.getPlcValueType() == tag.getPlcValueType()) {
                    return (PlcValue) value;
                } else {
                    throw new PlcRuntimeException("Expected PlcValue of type " + tag.getPlcValueType().name() + " but got " + plcValue.getPlcValueType().name());
                }
            }
            return typeValue(tag.getPlcValueType(),value);
        } else {
            PlcList list = new PlcList();
            for (Object value : values) {
                list.add(of(tag.getPlcValueType(), value));
            }
            if(tag.getNumberOfElements() > 1) {
                int rest = tag.getNumberOfElements() - values.length;
                if (rest > 0) {
                    while (rest > 0) {
                        list.add(of(tag.getPlcValueType(), null));
                        rest--;
                    }
                }
            }
            return list;
        }
    }

    public static PlcValue customDataType(Object[] values) {
        return of(values);
    }

    private static byte[] stringToByteArray(String stringBytes) {
        String[] byteValues = stringBytes.substring(1, stringBytes.length() - 1).split(",");
        byte[] bytes = new byte[byteValues.length];

        for (int i = 0, len = bytes.length; i < len; i++) {
            int intvalue = Integer.parseInt(byteValues[i].trim());
            bytes[i] = (byte) (intvalue & 0xff);
        }
        return bytes;
    }
    private static String[] stringToStringArray(String strings) {
        return strings.substring(1, strings.length() - 1).split(",");
    }
    private static byte[] getByteArray(Object[] arrs){
        byte[] bytes = new byte[arrs.length];
        for (int i=0;i<arrs.length;i++)
        {
            bytes[i] = (byte) arrs[i];
        }
        return bytes;
    }
    private static Object[] getArray(Object val){
        if (val instanceof Object[]) {
            return (Object[])val;
        }
        if(!val.getClass().isArray())
        {
            return new Object[]{val};
        }
        int arrLength = Array.getLength(val);
        Object[] outputArray = new Object[arrLength];
        for(int i = 0; i < arrLength; ++i){
            outputArray[i] = Array.get(val, i);
        }
        return outputArray;
    }
    private static PlcValue typeValue(PlcValueType type,Object value) {
        switch (type) {
            case BOOL:
                return PlcBOOL.of(value);
            case BYTE:
                return PlcBYTE.of(value);
            case SINT:
                return PlcSINT.of(value);
            case USINT:
                return PlcUSINT.of(value);
            case INT:
                return PlcINT.of(value);
            case UINT:
                return PlcUINT.of(value);
            case WORD:
                return PlcWORD.of(value);
            case DINT:
                return PlcDINT.of(value);
            case UDINT:
                return PlcUDINT.of(value);
            case DWORD:
                return PlcDWORD.of(value);
            case LINT:
                return PlcLINT.of(value);
            case ULINT:
                return PlcULINT.of(value);
            case LWORD:
                return PlcLWORD.of(value);
            case REAL:
                return PlcREAL.of(value);
            case LREAL:
                return PlcLREAL.of(value);
            case CHAR:
                return PlcCHAR.of(value);
            case WCHAR:
                return PlcWCHAR.of(value);
            case STRING:
                return PlcSTRING.of(value);
            case WSTRING:
                return PlcWSTRING.of(value);
            case TIME:
                return PlcTIME.of(value);
            case DATE:
                return PlcDATE.of(value);
            case TIME_OF_DAY:
                return PlcTIME_OF_DAY.of(value);
            case DATE_AND_TIME:
                return PlcDATE_AND_TIME.of(value);
            case LDATE:
                return PlcLDATE.of(value);
            case LTIME:
                return PlcLTIME.of(value);
            case LDATE_AND_TIME:
                return PlcLDATE_AND_TIME.of(value);
            case DATE_AND_LTIME:
                return PlcDATE_AND_LTIME.of(value);
            case LTIME_OF_DAY:
                return PlcLTIME_OF_DAY.of(value);
            default:
                return customDataType(getArray(value));
        }
    }
}
