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

import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.types.PlcValueType;
import org.apache.plc4x.java.spi.codegen.WithOption;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WriteBuffer;

import java.nio.charset.StandardCharsets;
import java.time.Instant;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.ZoneId;

public class PlcDATE extends PlcSimpleValue<LocalDate> {

    public static PlcDATE of(Object value) {
        if (value instanceof LocalDate) {
            return new PlcDATE((LocalDate) value);
        } else if (value instanceof Long) {
            return new PlcDATE(LocalDateTime.ofInstant(
                Instant.ofEpochSecond((long) value), ZoneId.systemDefault()).toLocalDate());
        }
        throw new PlcRuntimeException("Invalid value type");
    }

    public static PlcDATE ofSecondsSinceEpoch(long secondsSinceEpoch) {
        return new PlcDATE(LocalDateTime.ofInstant(
            Instant.ofEpochSecond(secondsSinceEpoch), ZoneId.systemDefault()).toLocalDate());
    }

    public static PlcDATE ofDaysSinceEpoch(int daysSinceEpoch) {
        // 86400 = 24 hours x 60 Minutes x 60 Seconds
        return new PlcDATE(LocalDateTime.ofInstant(
            Instant.ofEpochSecond(((long) daysSinceEpoch) * 86400), ZoneId.systemDefault()).toLocalDate());
    }

    public static PlcDATE ofDaysSinceSiemensEpoch(int daysSinceSiemensEpoch) {
        return ofDaysSinceEpoch(daysSinceSiemensEpoch + 7305);
    }

    public PlcDATE(LocalDate value) {
        super(value, true);
    }

    public PlcDATE(int daysSinceEpoch) {
        // REMARK: Yes, I'm using LocalDataTime.ofInstant as LocalDate.ofInstant is marked "JDK 1.9"
        super(LocalDateTime.ofInstant(
            Instant.ofEpochSecond(((long) daysSinceEpoch) * 86400), ZoneId.systemDefault()).toLocalDate(), true);
    }

    public PlcDATE(long secondsSinceEpoch) {
        // REMARK: Yes, I'm using LocalDataTime.ofInstant as LocalDate.ofInstant is marked "JDK 1.9"
        super(LocalDateTime.ofInstant(
            Instant.ofEpochSecond(secondsSinceEpoch), ZoneId.systemDefault()).toLocalDate(), true);
    }

    @Override
    public PlcValueType getPlcValueType() {
        return PlcValueType.DATE;
    }

    public long getSecondsSinceEpoch() {
        return value.atStartOfDay(ZoneId.systemDefault()).toEpochSecond();
    }

    public int getDaysSinceEpoch() {
        return (int) (value.atStartOfDay(ZoneId.systemDefault()).toEpochSecond() / 86400);
    }

    public int getDaysSinceSiemensEpoch() {
        return ((int) (value.atStartOfDay(ZoneId.systemDefault()).toEpochSecond() / 86400)) - 7305;
    }

    @Override
    public boolean isLong() {
        return true;
    }

    @Override
    public long getLong() {
        Instant instant = value.atStartOfDay(ZoneId.systemDefault()).toInstant();
        return (instant.toEpochMilli() / 1000);
    }

    @Override
    public boolean isString() {
        return true;
    }

    @Override
    public String getString() {
        return value.toString();
    }

    @Override
    public boolean isDate() {
        return true;
    }

    @Override
    public LocalDate getDate() {
        return value;
    }

    @Override
    public String toString() {
        return String.valueOf(value);
    }

    @Override
    public void serialize(WriteBuffer writeBuffer) throws SerializationException {
        String valueString = value.toString();
        writeBuffer.writeString(getClass().getSimpleName(),
            valueString.getBytes(StandardCharsets.UTF_8).length * 8,
            valueString, WithOption.WithEncoding(StandardCharsets.UTF_8.name()));
    }
    @Override
    public Object getPropertyByName(String property){
        switch (property){
        case "daysSinceSiemensEpoch":
            return this.getDaysSinceSiemensEpoch();
        case "secondsSinceEpoch":
            return this.getSecondsSinceEpoch();
        default:
            throw new PlcRuntimeException("Invalid property name");
        }
    }
}
