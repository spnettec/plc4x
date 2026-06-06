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

package org.apache.plc4x.java.utils.cache;

import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.messages.PlcWriteResponse;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcValue;

import java.util.LinkedHashSet;
import java.util.List;
import java.util.concurrent.CompletableFuture;
import java.util.function.Consumer;

class LeasedPlcWriteRequest implements PlcWriteRequest {

    private final PlcWriteRequest inner;
    private final Consumer<Boolean> onError;

    LeasedPlcWriteRequest(PlcWriteRequest inner, Consumer<Boolean> onError) {
        this.inner = inner;
        this.onError = onError;
    }

    @Override
    public CompletableFuture<? extends PlcWriteResponse> execute() {
        CompletableFuture<? extends PlcWriteResponse> future = inner.execute();
        if (future != null) {
            future.whenComplete((r, t) -> { if (t != null) onError.accept(true); });
        }
        return future;
    }

    @Override public int getNumberOfValues(String name) { return inner.getNumberOfValues(name); }
    @Override public PlcValue getPlcValue(String name) { return inner.getPlcValue(name); }
    @Override public int getNumberOfTags() { return inner.getNumberOfTags(); }
    @Override public LinkedHashSet<String> getTagNames() { return inner.getTagNames(); }
    @Override public PlcResponseCode getTagResponseCode(String name) { return inner.getTagResponseCode(name); }
    @Override public PlcTag getTag(String name) { return inner.getTag(name); }
    @Override public List<PlcTag> getTags() { return inner.getTags(); }
}
