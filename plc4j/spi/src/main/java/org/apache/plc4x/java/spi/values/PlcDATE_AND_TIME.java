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
import java.time.*;
import java.time.format.DateTimeFormatter;
import java.time.format.DateTimeParseException;

public class PlcDATE_AND_TIME extends PlcSimpleValue<LocalDateTime> {

    public static PlcDATE_AND_TIME of(Object value) {
        if (value instanceof LocalDateTime) {
            return new PlcDATE_AND_TIME((LocalDateTime) value);
        } else if (value instanceof Long) {
            return new PlcDATE_AND_TIME(LocalDateTime.ofInstant(
                Instant.ofEpochSecond((long) value), ZoneId.systemDefault()));
        }
        if (value instanceof Instant) {
            return new PlcDATE_AND_TIME(((Instant) value).toEpochMilli());
        }
        if (value instanceof LocalDate) {
            return new PlcDATE_AND_TIME(((LocalDate) value).atStartOfDay(ZoneId.of("UTC")).toLocalDateTime());
        } else if (value instanceof String) {
            String strValue = (String) value;
            LocalDateTime date;
            try {
                date = LocalDateTime.parse(strValue);
            } catch (DateTimeParseException e) {
                try {
                    date = LocalDateTime.parse(strValue, DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
                } catch (DateTimeParseException e1) {
                    try {
                        date = LocalDateTime.parse(strValue, DateTimeFormatter.ofPattern("yyyyMMdd HH:mm:ss"));
                    } catch (DateTimeParseException e2) {
                        try {
                            date = LocalDateTime.parse(strValue, DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
                        } catch (DateTimeParseException e3) {
                            try {
                                date = LocalDateTime.parse(strValue, DateTimeFormatter.ofPattern("yyyy-MM-dd"));
                            } catch (DateTimeParseException e4) {
                                date = LocalDateTime.parse(strValue, DateTimeFormatter.ofPattern("yyyyMMdd"));
                            }
                        }
                    }
                }
            }
            return new PlcDATE_AND_TIME(date);
        }
        throw new PlcRuntimeException("Invalid value type");
    }

    public static PlcDATE_AND_TIME ofSecondsSinceEpoch(long secondsSinceEpoch) {
        return new PlcDATE_AND_TIME(LocalDateTime.ofEpochSecond(secondsSinceEpoch, 0,
            ZoneOffset.UTC));
    }

    public static PlcDATE_AND_TIME ofSegments(int year, int month, int day, int hour, int minutes, int seconds, int nanoseconds) {
        return new PlcDATE_AND_TIME(LocalDateTime.of(year, month, day, hour, minutes, seconds, nanoseconds));
    }

    public PlcDATE_AND_TIME(LocalDateTime value) {
        super(value, true);
    }

    public PlcDATE_AND_TIME(long secondsSinceEpoch) {
        super(LocalDateTime.ofEpochSecond(secondsSinceEpoch, 0,
            ZoneOffset.UTC), true);
    }

    public PlcDATE_AND_TIME(int year, int month, int day, int hour, int minutes, int seconds, int nanoseconds) {
        super(LocalDateTime.of(year, month, day, hour, minutes, seconds, nanoseconds), true);
    }

    @Override
    public PlcValueType getPlcValueType() {
        return PlcValueType.DATE_AND_TIME;
    }

    public long getSecondsSinceEpoch() {
        Instant instant = getDateTime().toInstant(ZoneOffset.of(ZoneOffset.UTC.getId()));
        return instant.getEpochSecond();
    }

    public int getYear() {
        return value.getYear();
    }

    public int getMonth() {
        return value.getMonthValue();
    }

    public int getDay() {
        return value.getDayOfMonth();
    }

    public int getDayOfWeek() {
        return value.getDayOfWeek().getValue();
    }

    public int getHour() {
        return value.getHour();
    }

    public int getMinutes() {
        return value.getMinute();
    }

    public int getSeconds() {
        return value.getSecond();
    }

    public int getNanoseconds() {
        return value.getNano();
    }

    @Override
    public boolean isLong() {
        return true;
    }

    @Override
    public long getLong() {
        Instant instant = value.atZone(ZoneId.systemDefault()).toInstant();
        return instant.getEpochSecond();
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
    public boolean isTime() {
        return true;
    }

    @Override
    public LocalTime getTime() {
        return value.toLocalTime();
    }

    @Override
    public boolean isDate() {
        return true;
    }

    @Override
    public LocalDate getDate() {
        return value.toLocalDate();
    }

    @Override
    public boolean isDateTime() {
        return true;
    }

    @Override
    public LocalDateTime getDateTime() {
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
            valueString.getBytes(StandardCharsets.UTF_8).length*8,
            valueString, WithOption.WithEncoding(StandardCharsets.UTF_8.name()));
    }
    @Override
    public Object getPropertyByName(String property){
        switch (property){
        case "year":
            return this.getYear();
        case "month":
            return (short)this.getMonth();
        case "day":
            return (short)this.getDay();
        case "dayOfWeek":
            return (short)this.getDayOfWeek();
        case "hour":
            return (short)this.getHour();
        case "minutes":
            return (short)this.getMinutes();
        case "seconds":
            return (short)this.getSeconds();
        case "nanoseconds":
            return (long)this.getNanoseconds();
        case "secondsSinceEpoch":
            return this.getSecondsSinceEpoch();
        default:
            throw new PlcRuntimeException("Invalid property name");
        }
    }
}
