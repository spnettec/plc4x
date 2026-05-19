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
package org.apache.plc4x.java.transport.serial;

import io.netty.channel.Channel;
import io.netty.channel.nio.AbstractNioChannel.NioUnsafe;
import io.netty.util.concurrent.DefaultEventExecutor;
import io.netty.util.concurrent.DefaultPromise;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.nio.channels.SelectionKey;
import java.nio.channels.Selector;
import java.nio.channels.spi.AbstractSelectableChannel;
import java.nio.channels.spi.AbstractSelector;
import java.nio.channels.spi.SelectorProvider;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;
import java.util.stream.Collectors;

class SerialPollingSelector extends AbstractSelector {

    private static final Logger logger = LoggerFactory.getLogger(SerialPollingSelector.class);

    private final List<SelectionKey> registeredChannels;
    private final Set<SelectorEvent> events = ConcurrentHashMap.newKeySet();

    // Use a Netty Promise
    private final DefaultEventExecutor executor = new DefaultEventExecutor();
    private DefaultPromise<Void> selectPromise;

    public static class SelectorEvent {

        private final SelectionKey key;

        private final int event;
        public SelectorEvent(SelectionKey key, int event) {
            this.key = key;
            this.event = event;
        }

        public SelectionKey getKey() {
            return this.key;
        }

        public int getEvent() {
            return event;
        }

    }
    public SerialPollingSelector(SelectorProvider selectorProvider) {
        super(selectorProvider);
        registeredChannels = new ArrayList<>();
    }

    @Override
    public Set<SelectionKey> keys() {
        return new HashSet<>(registeredChannels);
    }

    /**
     * Always returns an empty set.
     *
     * Netty 4.2's {@code NioIoHandler.processSelectedKey()} unconditionally casts the key's
     * attachment to its package-private {@code DefaultNioRegistration} and crashes on anything
     * else. Since we register the SerialChannel itself as the attachment (4.1 contract), we have
     * to make sure the NioIoHandler never iterates our selected keys. Dispatch happens directly
     * from {@link #addEvent} via {@code eventLoop.execute(unsafe::read)} on the channel's event
     * loop instead.
     */
    @Override
    public Set<SelectionKey> selectedKeys() {
        return Collections.emptySet();
    }

    @Override
    public int selectNow() {
        logger.debug("selectNow()");
        // Reads are dispatched directly from addEvent(); never report selected keys to NioIoHandler.
        return 0;
    }

    @Override
    public int select(long timeout) {
        this.selectPromise = new DefaultPromise<>(executor);
        try {
            if(timeout == 0) {
                selectPromise.await();
            } else {
                selectPromise.await(timeout);
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            throw new RuntimeException("Was interrupted", e);
        }
        logger.debug("returning from select");
        // Reads are dispatched directly via addEvent(); never report selected keys to NioIoHandler.
        // Returning 0 also lets the event loop fall through to its queued-task processing where our
        // direct unsafe.read() lives.
        return 0;
    }

    @Override
    public int select() {
        return select(0);
    }

    @Override
    public Selector wakeup() {
        logger.debug("being asked to wake up from select");
        // throw new NotImplementedException("Not implemented for this selector, should not be needed.");
        if ((selectPromise != null) && !selectPromise.isDone()) {
            selectPromise.setSuccess(null);
        }
        return this;
    }

    public void addEvent(SelectorEvent event) {
        logger.debug("Adding Event to Selector, dispatching directly...");
        // Add the OP to the SelectionKey for any code that still polls interestOps via the key.
        ((SerialSelectionKey) event.key).addReadyOp(event.event);

        // Netty 4.2: NioIoHandler.processSelectedKey() requires SelectionKey.attachment() to be a
        // DefaultNioRegistration (package-private final). Since we register SerialChannel itself as
        // the attachment, we can't go through the standard NioIoHandler dispatch path without
        // ClassCastException. Instead, push the read directly onto the channel's event loop here.
        if ((event.event & SelectionKey.OP_READ) != 0) {
            Object attachment = event.key.attachment();
            if (attachment instanceof Channel) {
                Channel channel = (Channel) attachment;
                try {
                    channel.eventLoop().execute(() -> {
                        try {
                            ((NioUnsafe) channel.unsafe()).read();
                        } catch (Throwable t) {
                            logger.warn("SerialChannel read dispatch failed", t);
                        }
                    });
                } catch (Throwable t) {
                    logger.warn("Failed to schedule SerialChannel read on event loop", t);
                }
            }
        }

        // Wake any pending select() so the NioEventLoop returns to its task loop (where our
        // queued read() task lives). Also keep the historical events bookkeeping for now in case
        // anything else inspects it.
        this.events.add(event);
        if (selectPromise != null && !selectPromise.isDone()) {
            selectPromise.setSuccess(null);
        }
    }

    public void removeEvent(SerialSelectionKey serialSelectionKey) {
        events.removeIf(event -> event.key.equals(serialSelectionKey));
    }

    @Override
    protected void implCloseSelector() {
        // TODO should we do something here?
    }

    @Override
    protected SelectionKey register(AbstractSelectableChannel ch, int ops, Object att) {
        logger.debug("Registering Channel for selector {} with operations {}", ch, ops);
        if (!(ch instanceof SerialSocketChannel)) {
            throw new IllegalArgumentException("Given channel has to be of type " + SerialSocketChannel.class);
        }
        final SerialSelectionKey key = new SerialSelectionKey(ch, this, ops);
        // Attach attr
        key.attach(att);
        synchronized (this) {
            // TODO is this always the case??
            final int index = registeredChannels.size();
            registeredChannels.add(key);
            key.setIndex(index);
        }
        return key;
    }

}
