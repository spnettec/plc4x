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

package org.apache.plc4x.java.spi.transaction;

import io.netty.channel.Channel;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeoutException;
import java.util.function.Consumer;

public class TransactionTimeOutCallback<T>
        implements Consumer<TimeoutException>{
    private static final Logger logger = LoggerFactory.getLogger(TransactionTimeOutCallback.class);
    private final CompletableFuture<T> future;
    private final RequestTransactionManager.RequestTransaction transaction;
    private final Channel channel;

    private final boolean timeOutCloseChannel;

    public TransactionTimeOutCallback(CompletableFuture<T> future,
            RequestTransactionManager.RequestTransaction transaction,Channel channel) {
        this.future = future;
        this.transaction = transaction;
        this.channel = channel;
        this.timeOutCloseChannel = false;
    }
    public TransactionTimeOutCallback(CompletableFuture<T> future,
            RequestTransactionManager.RequestTransaction transaction,Channel channel,boolean timeOutCloseChannel) {
        this.future = future;
        this.transaction = transaction;
        this.channel = channel;
        this.timeOutCloseChannel = timeOutCloseChannel;
    }

    @Override
    public void accept(TimeoutException e) {
        try {
            if (transaction != null) {
                transaction.failRequest(e);
            }
            if(timeOutCloseChannel && channel != null && channel.isOpen()) {
                channel.close();
            }
        } catch (Exception ex) {
            logger.error(ex.getMessage());
        }
        if (future != null) {
            future.completeExceptionally(e);
        }
    }
}
