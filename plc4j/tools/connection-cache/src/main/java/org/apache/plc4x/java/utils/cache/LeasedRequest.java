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

import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.model.*;
import org.apache.plc4x.java.api.value.PlcValue;

import java.util.function.Consumer;

/**
 * Shared proxy logic for leased request wrappers that marks the parent
 * {@link LeasedPlcConnection} as invalid when an I/O error occurs during
 * {@code execute()}.
 */
final class LeasedRequest {

    private LeasedRequest() {}

    static PlcReadRequest.Builder read(PlcReadRequest.Builder inner, Consumer<Boolean> onError) {
        return new PlcReadRequest.Builder() {
            @Override public PlcReadRequest build() { return new LeasedPlcReadRequest(inner.build(), onError); }
            @Override public PlcReadRequest.Builder addTagAddress(String n, String a) { inner.addTagAddress(n, a); return this; }
            @Override public PlcReadRequest.Builder addTag(String n, PlcTag t) { inner.addTag(n, t); return this; }
        };
    }

    static PlcWriteRequest.Builder write(PlcWriteRequest.Builder inner, Consumer<Boolean> onError) {
        return new PlcWriteRequest.Builder() {
            @Override public PlcWriteRequest build() { return new LeasedPlcWriteRequest(inner.build(), onError); }
            @Override public PlcWriteRequest.Builder addTagAddress(String n, String a, Object... v) { inner.addTagAddress(n, a, v); return this; }
            @Override public PlcWriteRequest.Builder addTag(String n, PlcTag t, Object... v) { inner.addTag(n, t, v); return this; }
        };
    }

    static PlcSubscriptionRequest.Builder subscription(PlcSubscriptionRequest.Builder inner, Consumer<Boolean> onError) {
        return new PlcSubscriptionRequest.Builder() {
            @Override public PlcSubscriptionRequest build() { return new LeasedPlcSubscriptionRequest(inner.build(), onError); }
            @Override public PlcSubscriptionRequest.Builder setConsumer(java.util.function.Consumer<PlcSubscriptionEvent> c) { inner.setConsumer(c); return this; }
            @Override public PlcSubscriptionRequest.Builder addCyclicTagAddress(String n, String a, java.time.Duration d) { inner.addCyclicTagAddress(n, a, d); return this; }
            @Override public PlcSubscriptionRequest.Builder addCyclicTagAddress(String n, String a, java.time.Duration d, java.util.function.Consumer<PlcSubscriptionEvent> c) { inner.addCyclicTagAddress(n, a, d, c); return this; }
            @Override public PlcSubscriptionRequest.Builder addCyclicTag(String n, PlcTag t, java.time.Duration d) { inner.addCyclicTag(n, t, d); return this; }
            @Override public PlcSubscriptionRequest.Builder addChangeOfStateTagAddress(String n, String a, java.time.Duration d) { inner.addChangeOfStateTagAddress(n, a, d); return this; }
            @Override public PlcSubscriptionRequest.Builder addCyclicTag(String n, PlcTag t, java.time.Duration d, java.util.function.Consumer<PlcSubscriptionEvent> c) { inner.addCyclicTag(n, t, d, c); return this; }
            @Override public PlcSubscriptionRequest.Builder addChangeOfStateTagAddress(String n, String a) { inner.addChangeOfStateTagAddress(n, a); return this; }
            @Override public PlcSubscriptionRequest.Builder addChangeOfStateTagAddress(String n, String a, java.util.function.Consumer<PlcSubscriptionEvent> c) { inner.addChangeOfStateTagAddress(n, a, c); return this; }
            @Override public PlcSubscriptionRequest.Builder addChangeOfStateTagAddress(String n, String a, java.util.function.Consumer<PlcSubscriptionEvent> c, java.time.Duration d) { inner.addChangeOfStateTagAddress(n, a, c, d); return this; }
            @Override public PlcSubscriptionRequest.Builder addChangeOfStateTag(String n, PlcTag t) { inner.addChangeOfStateTag(n, t); return this; }
            @Override public PlcSubscriptionRequest.Builder addChangeOfStateTag(String n, PlcTag t, java.util.function.Consumer<PlcSubscriptionEvent> c) { inner.addChangeOfStateTag(n, t, c); return this; }
            @Override public PlcSubscriptionRequest.Builder addChangeOfStateTag(String n, PlcTag t, java.time.Duration d) { inner.addChangeOfStateTag(n, t, d); return this; }
            @Override public PlcSubscriptionRequest.Builder addChangeOfStateTag(String n, PlcTag t, java.util.function.Consumer<PlcSubscriptionEvent> c, java.time.Duration d) { inner.addChangeOfStateTag(n, t, c, d); return this; }
            @Override public PlcSubscriptionRequest.Builder addEventTagAddress(String n, String a) { inner.addEventTagAddress(n, a); return this; }
            @Override public PlcSubscriptionRequest.Builder addEventTagAddress(String n, String a, java.util.function.Consumer<PlcSubscriptionEvent> c) { inner.addEventTagAddress(n, a, c); return this; }
            @Override public PlcSubscriptionRequest.Builder addEventTag(String n, PlcTag t) { inner.addEventTag(n, t); return this; }
            @Override public PlcSubscriptionRequest.Builder addEventTag(String n, PlcTag t, java.util.function.Consumer<PlcSubscriptionEvent> c) { inner.addEventTag(n, t, c); return this; }
        };
    }

    static PlcUnsubscriptionRequest.Builder unsubscription(PlcUnsubscriptionRequest.Builder inner, Consumer<Boolean> onError) {
        return new PlcUnsubscriptionRequest.Builder() {
            @Override public PlcUnsubscriptionRequest build() { return new LeasedPlcUnsubscriptionRequest(inner.build(), onError); }
            @Override public PlcUnsubscriptionRequest.Builder addHandles(PlcSubscriptionHandle h) { inner.addHandles(h); return this; }
            @Override public PlcUnsubscriptionRequest.Builder addHandles(PlcSubscriptionHandle h1, PlcSubscriptionHandle... h) { inner.addHandles(h1, h); return this; }
            @Override public PlcUnsubscriptionRequest.Builder addHandles(java.util.Collection<PlcSubscriptionHandle> h) { inner.addHandles(h); return this; }
        };
    }

    static PlcBrowseRequest.Builder browse(PlcBrowseRequest.Builder inner, Consumer<Boolean> onError) {
        return new PlcBrowseRequest.Builder() {
            @Override public PlcBrowseRequest build() { return new LeasedPlcBrowseRequest(inner.build(), onError); }
            @Override public PlcBrowseRequest.Builder addQuery(String n, String q) { inner.addQuery(n, q); return this; }
        };
    }
}
