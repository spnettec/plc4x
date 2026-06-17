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
import org.junit.jupiter.api.Test;

import java.math.BigDecimal;
import java.math.BigInteger;

import static org.junit.jupiter.api.Assertions.*;

class PlcBITTest {

    @Test
    void testTrueValue() {
        PlcBIT value = new PlcBIT(true);
        assertTrue(value.getBoolean());
        assertEquals(PlcValueType.BOOL, value.getPlcValueType());
    }

    @Test
    void testFalseValue() {
        PlcBIT value = new PlcBIT(false);
        assertFalse(value.getBoolean());
    }

    @Test
    void testNullBooleanConstructor() {
        PlcBIT value = new PlcBIT((Boolean) null);
        assertTrue(value.isNullable());
    }

    @Test
    void testByteConstructorZero() {
        PlcBIT value = new PlcBIT((byte) 0);
        assertFalse(value.getBoolean());
    }

    @Test
    void testByteConstructorNonZero() {
        PlcBIT value = new PlcBIT((byte) 1);
        assertTrue(value.getBoolean());
    }

    @Test
    void testByteConstructorNull() {
        assertThrows(PlcInvalidTagException.class, () -> new PlcBIT((Byte) null));
    }

    @Test
    void testShortConstructor() {
        assertFalse(new PlcBIT((short) 0).getBoolean());
        assertTrue(new PlcBIT((short) 1).getBoolean());
    }

    @Test
    void testShortConstructorNull() {
        assertThrows(PlcInvalidTagException.class, () -> new PlcBIT((Short) null));
    }

    @Test
    void testIntegerConstructor() {
        assertFalse(new PlcBIT(0).getBoolean());
        assertTrue(new PlcBIT(1).getBoolean());
    }

    @Test
    void testIntegerConstructorNull() {
        assertThrows(PlcInvalidTagException.class, () -> new PlcBIT((Integer) null));
    }

    @Test
    void testLongConstructor() {
        assertFalse(new PlcBIT(0L).getBoolean());
        assertTrue(new PlcBIT(1L).getBoolean());
    }

    @Test
    void testLongConstructorNull() {
        assertThrows(PlcInvalidTagException.class, () -> new PlcBIT((Long) null));
    }

    @Test
    void testFloatConstructor() {
        assertFalse(new PlcBIT(0.0f).getBoolean());
        assertTrue(new PlcBIT(0.1f).getBoolean());
    }

    @Test
    void testFloatConstructorNull() {
        assertThrows(PlcInvalidTagException.class, () -> new PlcBIT((Float) null));
    }

    @Test
    void testDoubleConstructor() {
        assertFalse(new PlcBIT(0.0).getBoolean());
        assertTrue(new PlcBIT(0.1).getBoolean());
    }

    @Test
    void testDoubleConstructorNull() {
        assertThrows(PlcInvalidTagException.class, () -> new PlcBIT((Double) null));
    }

    @Test
    void testBigIntegerConstructor() {
        assertFalse(new PlcBIT(BigInteger.ZERO).getBoolean());
        assertTrue(new PlcBIT(BigInteger.ONE).getBoolean());
    }

    @Test
    void testBigIntegerConstructorNull() {
        assertThrows(PlcInvalidTagException.class, () -> new PlcBIT((BigInteger) null));
    }

    @Test
    void testBigDecimalConstructor() {
        assertFalse(new PlcBIT(BigDecimal.ZERO).getBoolean());
        assertTrue(new PlcBIT(BigDecimal.ONE).getBoolean());
    }

    @Test
    void testBigDecimalConstructorNull() {
        assertThrows(PlcInvalidTagException.class, () -> new PlcBIT((BigDecimal) null));
    }

    @Test
    void testStringConstructor() {
        assertTrue(new PlcBIT("true").getBoolean());
        assertTrue(new PlcBIT("TRUE").getBoolean());
        assertFalse(new PlcBIT("false").getBoolean());
        assertFalse(new PlcBIT("0").getBoolean());
        assertTrue(new PlcBIT("1").getBoolean());
    }

    @Test
    void testOfMethodWithPlcBIT() {
        PlcBIT original = new PlcBIT(true);
        PlcBIT copy = PlcBIT.of(original);
        assertSame(original, copy);
    }

    @Test
    void testOfMethodWithNull() {
        PlcBIT value = PlcBIT.of(null);
        assertNotNull(value);
    }

    @Test
    void testOfMethodWithBoolean() {
        assertTrue(PlcBIT.of(true).getBoolean());
    }

    @Test
    void testOfMethodWithInteger() {
        assertTrue(PlcBIT.of(1).getBoolean());
        assertFalse(PlcBIT.of(0).getBoolean());
    }

    @Test
    void testOfMethodWithString() {
        assertTrue(PlcBIT.of("true").getBoolean());
    }

    @Test
    void testOfMethodWithNumber() {
        // Generic Number type (not Byte/Short/Integer/Long/Float/Double)
        Number num = new Number() {
            @Override public int intValue() { return 1; }
            @Override public long longValue() { return 1L; }
            @Override public float floatValue() { return 1.0f; }
            @Override public double doubleValue() { return 1.0; }
        };
        assertTrue(PlcBIT.of(num).getBoolean());
    }

    @Test
    void testGetters() {
        PlcBIT trueVal = new PlcBIT(true);
        assertEquals(1, trueVal.getByte());
        assertEquals(1, trueVal.getShort());
        assertEquals(1, trueVal.getInteger());
        assertEquals(1L, trueVal.getLong());
        assertEquals(BigInteger.ONE, trueVal.getBigInteger());
        assertEquals("true", trueVal.getString());

        PlcBIT falseVal = new PlcBIT(false);
        assertEquals(0, falseVal.getByte());
        assertEquals(0, falseVal.getInteger());
        assertEquals("false", falseVal.getString());
    }

    @Test
    void testIsTypes() {
        PlcBIT value = new PlcBIT(true);
        assertTrue(value.isBoolean());
        assertTrue(value.isByte());
        assertTrue(value.isShort());
        assertTrue(value.isInteger());
        assertTrue(value.isLong());
        assertTrue(value.isBigInteger());
        assertTrue(value.isString());
    }

    @Test
    void testGetBytes() {
        PlcBIT trueVal = new PlcBIT(true);
        byte[] bytes = trueVal.getBytes();
        assertEquals(1, bytes.length);
        assertEquals(1, bytes[0]);

        PlcBIT falseVal = new PlcBIT(false);
        assertEquals(0, falseVal.getBytes()[0]);
    }

    @Test
    void testToString() {
        assertEquals("true", new PlcBIT(true).toString());
        assertEquals("false", new PlcBIT(false).toString());
    }
}
