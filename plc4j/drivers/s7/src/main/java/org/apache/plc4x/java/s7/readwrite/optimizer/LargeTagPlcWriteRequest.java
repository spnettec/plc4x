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

package org.apache.plc4x.java.s7.readwrite.optimizer;

import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.messages.PlcWriteResponse;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.messages.utils.PlcTagItem;

import java.util.*;
import java.util.concurrent.CompletableFuture;

public class LargeTagPlcWriteRequest implements PlcWriteRequest {

    private final PlcValue plcValue;
    private final PlcTagItem<PlcTag> tagItem;
    public LargeTagPlcWriteRequest(String tagName, PlcValue plcValue, PlcTagItem<PlcTag> tagItem) {
        this.tagName = tagName;
        this.plcValue = plcValue;
        this.tagItem = tagItem;
    }

    private final String tagName;

    public String getTagName() {
        return tagName;
    }

    public PlcValue getPlcValue() {
        return plcValue;
    }

    @Override
    public int getNumberOfTags() {
        return 1;
    }

    @Override
    public LinkedHashSet<String> getTagNames() {
        return new LinkedHashSet<>(Collections.singletonList(tagName));
    }

    @Override
    public PlcResponseCode getTagResponseCode(String tagName) {
        return tagItem.getResponseCode();
    }

    @Override
    public PlcTag getTag(String name) {
        return tagItem.getTag();
    }

    @Override
    public List<PlcTag> getTags() {
        return Collections.singletonList(tagItem.getTag());
    }

    @Override
    public CompletableFuture<? extends PlcWriteResponse> execute() {
        throw new UnsupportedOperationException();
    }

    @Override
    public int getNumberOfValues(String name) {
        return 1;
    }

    @Override
    public PlcValue getPlcValue(String name) {
        return plcValue;
    }
}
