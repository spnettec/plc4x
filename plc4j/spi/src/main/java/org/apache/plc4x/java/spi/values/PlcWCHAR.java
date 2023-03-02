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

import org.apache.plc4x.java.api.exceptions.PlcInvalidTagException;
import org.apache.plc4x.java.api.types.PlcValueType;
import org.apache.plc4x.java.spi.codegen.WithOption;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WriteBuffer;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.nio.charset.StandardCharsets;

public class PlcWCHAR extends PlcIECValue<Integer> {

    private static final String VALUE_OUT_OF_RANGE = "Value of type %s is out of range %d - %d for a %s Value";
    static Integer minValue = 0;
    static Integer maxValue = Short.MAX_VALUE * 2 + 1;

    public static PlcWCHAR of(Object value) {
        if (value instanceof Boolean) {
            return new PlcWCHAR((Boolean) value);
        } else if (value instanceof Byte) {
            return new PlcWCHAR((Byte) value);
        } else if (value instanceof Short) {
            return new PlcWCHAR((Short) value);
        } else if (value instanceof Integer) {
            return new PlcWCHAR((Integer) value);
        } else if (value instanceof Long) {
            return new PlcWCHAR((Long) value);
        } else if (value instanceof Float) {
            return new PlcWCHAR((Float) value);
        } else if (value instanceof Double) {
            return new PlcWCHAR((Double) value);
        } else if (value instanceof BigInteger) {
            return new PlcWCHAR((BigInteger) value);
        } else if (value instanceof BigDecimal) {
            return new PlcWCHAR((BigDecimal) value);
        } else if (value instanceof Number) {
            return new PlcWCHAR(((Number) value).intValue());
        } else {
            return new PlcWCHAR((String) value);
        }
    }

    public PlcWCHAR(Boolean value) {
        this.value = value ? (Integer) 1 : (Integer) 0;
        this.isNullable = false;
    }

    public PlcWCHAR(Byte value) {
        this.value = value.intValue();
        this.isNullable = false;
    }

    public PlcWCHAR(Character value) {
        int val = (int) value;
        if ((val < minValue) || (val > maxValue)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = val;
        this.isNullable = false;
    }

    public PlcWCHAR(Short value) {
        if ((value < minValue) || (value > maxValue)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value.intValue();
        this.isNullable = false;
    }

    public PlcWCHAR(Integer value) {
        if ((value < minValue) || (value > maxValue)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value;
        this.isNullable = false;
    }

    public PlcWCHAR(Long value) {
        if ((value < minValue) || (value > maxValue)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        } else {
            this.value = value.intValue();
            this.isNullable = false;
        }
    }

    public PlcWCHAR(Float value) {
        if ((value < minValue) || (value > maxValue) || (value % 1 != 0)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value.intValue();
        this.isNullable = false;
    }

    public PlcWCHAR(Double value) {
        if ((value < minValue) || (value > maxValue) || (value % 1 != 0)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value.intValue();
        this.isNullable = false;
    }

    public PlcWCHAR(BigInteger value) {
        if ((value.compareTo(BigInteger.valueOf(minValue)) < 0) || (value.compareTo(BigInteger.valueOf(maxValue)) > 0)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value.intValue();
        this.isNullable = true;
    }

    public PlcWCHAR(BigDecimal value) {
        if ((value.compareTo(BigDecimal.valueOf(minValue)) < 0) || (value.compareTo(BigDecimal.valueOf(maxValue)) > 0) || (value.scale() > 0)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value.intValue();
        this.isNullable = true;
    }

    public PlcWCHAR(String value) {
        try {
            //If there is a extra space around the character trim it, unless you are actually sending a space
            String s = value.trim();
            if (s.length() == 0) {
                s = " ";
            }
            int val = (int) s.charAt(0);
            if ((val < minValue) || (val > maxValue)) {
                throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
            }
            this.value = val;
            this.isNullable = false;
        } catch (Exception e) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()), e);
        }
    }

    public PlcWCHAR(int value) {
        if ((value < minValue) || (value > maxValue)) {
            throw new PlcInvalidTagException(String.format(VALUE_OUT_OF_RANGE, value, minValue, maxValue, this.getClass().getSimpleName()));
        }
        this.value = value;
        this.isNullable = false;
    }

    @Override
    public PlcValueType getPlcValueType() {
        return PlcValueType.WCHAR;
    }

    @Override
    public boolean isBoolean() {
        return true;
    }

    @Override
    public boolean getBoolean() {
        return (value != null) && !value.equals(0);
    }

    @Override
    public boolean isByte() {
        return (value != null) && (value <= Byte.MAX_VALUE) && (value >= Byte.MIN_VALUE);
    }

    @Override
    public byte getByte() {
        return value.byteValue();
    }

    @Override
    public boolean isShort() {
        return (value != null) && (value <= Short.MAX_VALUE) && (value >= Short.MIN_VALUE);
    }

    @Override
    public short getShort() {
        return value.shortValue();
    }

    @Override
    public boolean isInteger() {
        return true;
    }

    @Override
    public int getInteger() {
        return value;
    }

    @Override
    public boolean isLong() {
        return true;
    }

    @Override
    public long getLong() {
        return value.longValue();
    }

    @Override
    public boolean isBigInteger() {
        return true;
    }

    @Override
    public BigInteger getBigInteger() {
        return BigInteger.valueOf(getLong());
    }

    @Override
    public boolean isFloat() {
        return true;
    }

    @Override
    public float getFloat() {
        return value.floatValue();
    }

    @Override
    public boolean isDouble() {
        return true;
    }

    @Override
    public double getDouble() {
        return value.doubleValue();
    }

    @Override
    public boolean isBigDecimal() {
        return true;
    }

    @Override
    public BigDecimal getBigDecimal() {
        return BigDecimal.valueOf(getFloat());
    }

    @Override
    public boolean isString() {
        return true;
    }

    @Override
    public String getString() {
        return toString();
    }

    @Override
    public Object getObject() {
        return Character.toString((char) ((int) value));
    }

    @Override
    public String toString() {
        return Character.toString((char) ((int) value));
    }

    public byte[] getBytes() {
        return new byte[]{
            (byte) ((value >> 8) & 0xff),
            (byte) (value & 0xff)
        };
    }

    @Override
    public void serialize(WriteBuffer writeBuffer) throws SerializationException {
        String valueString = value.toString();
        writeBuffer.writeString(getClass().getSimpleName(),
            16,
            valueString, WithOption.WithEncoding(StandardCharsets.UTF_8.name()));
    }

}
