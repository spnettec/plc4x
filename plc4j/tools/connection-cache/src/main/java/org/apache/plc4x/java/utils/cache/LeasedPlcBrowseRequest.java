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

import org.apache.plc4x.java.api.messages.PlcBrowseRequest;
import org.apache.plc4x.java.api.messages.PlcBrowseRequestInterceptor;
import org.apache.plc4x.java.api.messages.PlcBrowseResponse;
import org.apache.plc4x.java.api.model.PlcQuery;

import java.util.LinkedHashSet;
import java.util.concurrent.CompletableFuture;
import java.util.function.Consumer;

class LeasedPlcBrowseRequest implements PlcBrowseRequest {

    private final PlcBrowseRequest inner;
    private final Consumer<Boolean> onError;

    LeasedPlcBrowseRequest(PlcBrowseRequest inner, Consumer<Boolean> onError) {
        this.inner = inner;
        this.onError = onError;
    }

    @Override
    public CompletableFuture<? extends PlcBrowseResponse> execute() {
        CompletableFuture<? extends PlcBrowseResponse> future = inner.execute();
        if (future != null) {
            future.whenComplete((r, t) -> { if (t != null) onError.accept(true); });
        }
        return future;
    }

    @Override
    public CompletableFuture<? extends PlcBrowseResponse> executeWithInterceptor(PlcBrowseRequestInterceptor interceptor) {
        CompletableFuture<? extends PlcBrowseResponse> future = inner.executeWithInterceptor(interceptor);
        if (future != null) {
            future.whenComplete((r, t) -> { if (t != null) onError.accept(true); });
        }
        return future;
    }

    @Override public LinkedHashSet<String> getQueryNames() { return inner.getQueryNames(); }
    @Override public PlcQuery getQuery(String name) { return inner.getQuery(name); }
}
