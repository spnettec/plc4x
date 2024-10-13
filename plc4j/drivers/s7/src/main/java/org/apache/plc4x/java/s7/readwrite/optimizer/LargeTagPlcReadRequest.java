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

import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.spi.messages.utils.PlcTagItem;

import java.util.Collections;
import java.util.LinkedHashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.concurrent.CompletableFuture;

public class LargeTagPlcReadRequest implements PlcReadRequest {
    private final String tagName;
    private final PlcTagItem<PlcTag> tagItem;
    public LargeTagPlcReadRequest(String tagName, PlcTagItem<PlcTag> tagItem) {
        this.tagName = tagName;
        this.tagItem = tagItem;

    }
    public String getTagName() {
        return tagName;
    }
    public PlcTag getTag() {
        return tagItem.getTag();
    }

    @Override
    public CompletableFuture<? extends PlcReadResponse> execute() {
        throw new UnsupportedOperationException();
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
        throw new UnsupportedOperationException();
    }

    @Override
    public List<PlcTag> getTags() {
        return new LinkedList<>(Collections.singletonList(tagItem.getTag()));
    }
}
