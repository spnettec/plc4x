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

import org.apache.plc4x.java.api.messages.PlcUnsubscriptionRequest;
import org.apache.plc4x.java.api.messages.PlcUnsubscriptionResponse;
import org.apache.plc4x.java.api.model.PlcSubscriptionHandle;

import java.util.List;
import java.util.concurrent.CompletableFuture;
import java.util.function.Consumer;

class LeasedPlcUnsubscriptionRequest implements PlcUnsubscriptionRequest {

    private final PlcUnsubscriptionRequest inner;
    private final Consumer<Boolean> onError;

    LeasedPlcUnsubscriptionRequest(PlcUnsubscriptionRequest inner, Consumer<Boolean> onError) {
        this.inner = inner;
        this.onError = onError;
    }

    @Override
    public CompletableFuture<PlcUnsubscriptionResponse> execute() {
        CompletableFuture<? extends PlcUnsubscriptionResponse> future = inner.execute();
        if (future != null) {
            future.whenComplete((r, t) -> { if (t != null) onError.accept(true); });
        }
        @SuppressWarnings("unchecked")
        CompletableFuture<PlcUnsubscriptionResponse> result = (CompletableFuture<PlcUnsubscriptionResponse>) future;
        return result;
    }

    @Override
    public List<PlcSubscriptionHandle> getSubscriptionHandles() { return inner.getSubscriptionHandles(); }
}
