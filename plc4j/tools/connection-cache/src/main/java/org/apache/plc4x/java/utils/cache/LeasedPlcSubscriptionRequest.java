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

import org.apache.plc4x.java.api.messages.PlcSubscriptionEvent;
import org.apache.plc4x.java.api.messages.PlcSubscriptionRequest;
import org.apache.plc4x.java.api.messages.PlcSubscriptionResponse;
import org.apache.plc4x.java.api.model.PlcSubscriptionTag;
import org.apache.plc4x.java.api.types.PlcResponseCode;

import java.util.LinkedHashSet;
import java.util.List;
import java.util.concurrent.CompletableFuture;
import java.util.function.Consumer;

class LeasedPlcSubscriptionRequest implements PlcSubscriptionRequest {

    private final PlcSubscriptionRequest inner;
    private final Consumer<Boolean> onError;

    LeasedPlcSubscriptionRequest(PlcSubscriptionRequest inner, Consumer<Boolean> onError) {
        this.inner = inner;
        this.onError = onError;
    }

    @Override
    public CompletableFuture<? extends PlcSubscriptionResponse> execute() {
        CompletableFuture<? extends PlcSubscriptionResponse> future = inner.execute();
        if (future != null) {
            future.whenComplete((r, t) -> { if (t != null) onError.accept(true); });
        }
        return future;
    }

    @Override public int getNumberOfTags() { return inner.getNumberOfTags(); }
    @Override public LinkedHashSet<String> getTagNames() { return inner.getTagNames(); }
    @Override public PlcSubscriptionTag getTag(String name) { return inner.getTag(name); }
    @Override public PlcResponseCode getTagResponseCode(String name) { return inner.getTagResponseCode(name); }
    @Override public List<PlcSubscriptionTag> getTags() { return inner.getTags(); }
    @Override public Consumer<PlcSubscriptionEvent> getConsumer() { return inner.getConsumer(); }
    @Override public Consumer<PlcSubscriptionEvent> getTagConsumer(String name) { return inner.getTagConsumer(name); }
}
