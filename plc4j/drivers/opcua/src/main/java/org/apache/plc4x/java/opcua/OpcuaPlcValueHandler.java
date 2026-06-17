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
package org.apache.plc4x.java.opcua;

import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcValueType;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.values.DefaultPlcValueHandler;
import org.apache.plc4x.java.spi.values.PlcList;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;

/**
 * OPC UA-specific value handler that infers the PLC value type from the
 * actual Java value when the tag type is {@link PlcValueType#NULL}.
 *
 * <p>The new SPI's {@link DefaultPlcValueHandler} creates {@code PlcNull}
 * for NULL-typed tags, discarding the actual value.  The pre-merge (old SPI)
 * handler inferred the type from the Java object, which is what OPC UA needs
 * because many tag addresses don't carry a type suffix.</p>
 */
public class OpcuaPlcValueHandler extends DefaultPlcValueHandler {

    @Override
    public PlcValue newPlcValue(PlcTag tag, Object value) {
        if (tag.getPlcValueType() == PlcValueType.NULL && value != null) {
            PlcValueType inferred = inferType(value);
            if (inferred != PlcValueType.NULL) {
                tag.setPlcValueType(inferred);
            }
        }
        return super.newPlcValue(tag, value);
    }

    @Override
    public PlcValue newPlcValue(PlcTag tag, Object[] values) {
        if (values != null && values.length > 0 && values[0] != null) {
            if (tag.getPlcValueType() == PlcValueType.NULL) {
                PlcValueType inferred = inferType(values[0]);
                if (inferred != PlcValueType.NULL) {
                    tag.setPlcValueType(inferred);
                }
            }
            // OPC UA array tags often have no [n] dimension in the address,
            // so tag.getArrayInfo() is empty.  The default handler rejects
            // multiple values for non-array tags.  We handle this by creating
            // a PlcList from the individual elements — matching the pre-merge
            // (old SPI) behavior where PlcList was created transparently.
            if (values.length > 1 && tag.getArrayInfo().isEmpty()) {
                PlcValueType elementType = tag.getPlcValueType();
                List<PlcValue> elements = new ArrayList<>(values.length);
                for (Object v : values) {
                    elements.add(super.newPlcValue(tag, v));
                }
                return new PlcList(elements);
            }
        }
        return super.newPlcValue(tag, values);
    }

    /**
     * Infer the {@link PlcValueType} from a Java object, replicating the
     * pre-merge behavior of the old SPI's value handler.
     */
    private static PlcValueType inferType(Object value) {
        if (value instanceof Boolean)    return PlcValueType.BOOL;
        if (value instanceof Byte)       return PlcValueType.SINT;
        if (value instanceof Short)      return PlcValueType.INT;
        if (value instanceof Integer)    return PlcValueType.DINT;
        if (value instanceof Long)       return PlcValueType.LINT;
        if (value instanceof Float)      return PlcValueType.REAL;
        if (value instanceof Double)     return PlcValueType.LREAL;
        if (value instanceof String)     return PlcValueType.STRING;
        if (value instanceof BigInteger) return PlcValueType.ULINT;
        if (value instanceof boolean[])  return PlcValueType.BOOL;
        return PlcValueType.NULL;
    }
}
