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

package org.apache.plc4x.java.base.connection;

import com.fazecast.jSerialComm.SerialPort;
import com.fazecast.jSerialComm.SerialPortDataListener;
import com.fazecast.jSerialComm.SerialPortEvent;
import io.netty.buffer.ByteBuf;
import io.netty.buffer.ByteBufAllocator;
import io.netty.channel.AbstractChannel;
import io.netty.channel.Channel;
import io.netty.channel.ChannelConfig;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelOutboundBuffer;
import io.netty.channel.ChannelPipeline;
import io.netty.channel.ChannelPromise;
import io.netty.channel.DefaultChannelConfig;
import io.netty.channel.DefaultChannelPipeline;
import io.netty.channel.EventLoop;
import io.netty.channel.FileRegion;
import io.netty.channel.RecvByteBufAllocator;
import io.netty.channel.VoidChannelPromise;
import io.netty.channel.jsc.JSerialCommDeviceAddress;
import io.netty.channel.nio.AbstractNioByteChannel;
import io.netty.channel.nio.AbstractNioChannel;
import io.netty.channel.nio.NioEventLoop;
import io.netty.channel.socket.DuplexChannel;
import org.apache.commons.lang3.NotImplementedException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.lang.reflect.Field;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.net.SocketAddress;
import java.nio.channels.ClosedChannelException;
import java.nio.channels.SelectableChannel;
import java.util.concurrent.RejectedExecutionException;

/**
 * TODO write comment
 *
 * @author julian
 * Created by julian on 2019-08-10
 */
public class SerialChannel extends AbstractNioByteChannel implements DuplexChannel {

    private static final Logger logger = LoggerFactory.getLogger(SerialChannel.class);
    private final ChannelConfig config;

    private final VoidChannelPromise unsafeVoidPromise = new VoidChannelPromise(this, false);
    private boolean readPending = false; // Did we receive an EOF?
    private SocketAddress remoteAddress;
    private boolean active = false;
    private SerialSelectionKey selectionKey;
    private SerialChannelHandler comPort;


    public SerialChannel() {
        this(null, new SerialSocketChannel(new SerialSelectorProvider()));
        ((SerialSocketChannel) javaChannel()).setChild(this);
    }

    /**
     * Create a new instance
     *
     * @param parent the parent {@link Channel} by which this instance was created. May be {@code null}
     * @param ch     the underlying {@link SelectableChannel} on which it operates
     */
    protected SerialChannel(Channel parent, SelectableChannel ch) {
        super(parent, ch);
        config = new DefaultChannelConfig(this);
    }

    @Override
    public NioUnsafe unsafe() {
        return new SerialNioUnsafe();
    }

    @Override
    public boolean isInputShutdown() {
        throw new NotImplementedException("");
    }

    @Override
    public ChannelFuture shutdownInput() {
        throw new NotImplementedException("");
    }

    @Override
    public ChannelFuture shutdownInput(ChannelPromise promise) {
        throw new NotImplementedException("");
    }

    @Override
    public boolean isOutputShutdown() {
        throw new NotImplementedException("");
    }

    @Override
    public ChannelFuture shutdownOutput() {
        throw new NotImplementedException("");
    }

    @Override
    public ChannelFuture shutdownOutput(ChannelPromise promise) {
        throw new NotImplementedException("");
    }

    @Override
    public boolean isShutdown() {
        throw new NotImplementedException("");
    }

    @Override
    public ChannelFuture shutdown() {
        throw new NotImplementedException("");
    }

    @Override
    public ChannelFuture shutdown(ChannelPromise promise) {
        throw new NotImplementedException("");
    }

    @Override
    protected long doWriteFileRegion(FileRegion region) throws Exception {
        throw new NotImplementedException("");
    }

    @Override
    protected int doReadBytes(ByteBuf buf) throws Exception {
        if (!active) {
            return 0;
        }
        // TODO Here we really read the bytes
        logger.debug("Trying to read bytes from wire...");
        buf.writeByte(0x01);
        return 1;
    }

    @Override
    protected int doWriteBytes(ByteBuf buf) throws Exception {
        throw new NotImplementedException("");
    }

    @Override
    protected boolean doConnect(SocketAddress remoteAddress, SocketAddress localAddress) throws Exception {
        this.remoteAddress = remoteAddress;
        if (!(remoteAddress instanceof JSerialCommDeviceAddress)) {
            throw new IllegalArgumentException("Socket Address has to be of type " + JSerialCommDeviceAddress.class);
        }
        logger.debug("Connecting to Socket Address '{}'", ((JSerialCommDeviceAddress) remoteAddress).value());

        try {
            // TODO this should take port from remote Adress
            comPort = SerialChannelHandler.DummyHandler.INSTANCE;
            logger.debug("Using Com Port {}, trying to open port", comPort.getIdentifier());
            if (comPort.open()) {
                logger.debug("Opened port successful to {}", comPort.getIdentifier());
                comPort.registerSelectionKey(selectionKey);

                this.active = true;
                return true;
            } else {
                logger.debug("Unable to open port {}", comPort.getIdentifier());
                return false;
            }
        } catch (Exception e) {
            e.printStackTrace();
            this.active = false;
            return false;
        }
    }

    @Override
    protected void doFinishConnect() throws Exception {
        throw new NotImplementedException("");
    }

    @Override
    protected SocketAddress localAddress0() {
        return null;
    }

    @Override
    protected SocketAddress remoteAddress0() {
        return null;
    }

    @Override
    protected void doBind(SocketAddress localAddress) throws Exception {
        throw new NotImplementedException("");
    }

    @Override
    protected void doDisconnect() throws Exception {
        throw new NotImplementedException("");
    }

    @Override
    public ChannelConfig config() {
        return this.config;
    }

    @Override
    public boolean isActive() {
        return active;
    }

    private class SerialNioUnsafe implements NioUnsafe {

        private RecvByteBufAllocator.Handle recvHandle;

        @Override
        public SelectableChannel ch() {
            throw new NotImplementedException("");
        }

        @Override
        public void finishConnect() {
            throw new NotImplementedException("");
        }

        // See NioByteUnsafe#read()
        @Override
        public void read() {
            logger.debug("Reading...");
            // TODO we should read something here, okay?!
            final ChannelConfig config = config();
            final ChannelPipeline pipeline = pipeline();
            final ByteBufAllocator allocator = config.getAllocator();
            final RecvByteBufAllocator.Handle allocHandle = recvBufAllocHandle();
            allocHandle.reset(config);

            ByteBuf byteBuf = null;
            boolean close = false;
            try {
                do {
                    byteBuf = allocHandle.allocate(allocator);
                    allocHandle.lastBytesRead(doReadBytes(byteBuf));
                    if (allocHandle.lastBytesRead() <= 0) {
                        // nothing was read. release the buffer.
                        byteBuf.release();
                        byteBuf = null;
                        close = allocHandle.lastBytesRead() < 0;
                        if (close) {
                            // There is nothing left to read as we received an EOF.
                            readPending = false;
                        }
                        break;
                    }

                    allocHandle.incMessagesRead(1);
                    readPending = false;
                    pipeline.fireChannelRead(byteBuf);
                    byteBuf = null;
                } while (allocHandle.continueReading());

                allocHandle.readComplete();
                pipeline.fireChannelReadComplete();

                if (close) {
                    // TODO
                    //closeOnRead(pipeline);
                }
            } catch (Throwable t) {
                // TODO
                // handleReadException(pipeline, byteBuf, t, close, allocHandle);
            } finally {
                // Check if there is a readPending which was not processed yet.
                // This could be for two reasons:
                // * The user called Channel.read() or ChannelHandlerContext.read() in channelRead(...) method
                // * The user called Channel.read() or ChannelHandlerContext.read() in channelReadComplete(...) method
                //
                // See https://github.com/netty/netty/issues/2254
                if (!readPending && !config.isAutoRead()) {
                    // TODO
                }
            }
        }

        @Override
        public void forceFlush() {
            throw new NotImplementedException("");
        }

        @Override
        public RecvByteBufAllocator.Handle recvBufAllocHandle() {
            if (recvHandle == null) {
                recvHandle = config().getRecvByteBufAllocator().newHandle();
            }
            return recvHandle;
        }

        @Override
        public SocketAddress localAddress() {
            throw new NotImplementedException("");
        }

        @Override
        public SocketAddress remoteAddress() {
            return null;
        }

        @Override
        public void register(EventLoop eventLoop, ChannelPromise promise) {
            // Register
            if (!(eventLoop instanceof NioEventLoop)) {
                throw new IllegalArgumentException("Only valid for NioEventLoop!");
            }
            if (!(promise.channel() instanceof SerialChannel)) {
                throw new IllegalArgumentException("Only valid for " + SerialChannel.class + " but is " + promise.channel().getClass());
            }
            // Register channel to event loop
            // We have to use reflection here, I fear
            try {
                Method method = NioEventLoop.class.getDeclaredMethod("unwrappedSelector");
                method.setAccessible(true);
                SerialPollingSelector selector = (SerialPollingSelector) method.invoke(eventLoop);

                // Register the channel
                selectionKey = (SerialSelectionKey) ((SerialChannel) promise.channel()).javaChannel().register(selector, 0, SerialChannel.this);

                // Set selection key
                final Field selectionKeyField = AbstractNioChannel.class.getDeclaredField("selectionKey");
                selectionKeyField.setAccessible(true);
                selectionKeyField.set(SerialChannel.this, selectionKey);

                // Set event loop (again, via reflection)
                final Field loop = AbstractChannel.class.getDeclaredField("eventLoop");
                loop.setAccessible(true);
                loop.set(SerialChannel.this, eventLoop);

                // Register Pipeline, if necessary
                // Ensure we call handlerAdded(...) before we actually notify the promise. This is needed as the
                // user may already fire events through the pipeline in the ChannelFutureListener.
                if (!(pipeline() instanceof DefaultChannelPipeline)) {
                    throw new IllegalStateException("Pipeline should be of Type " + DefaultChannelPipeline.class);
                }
                // Again reflection, but has to be done in an event loop
                eventLoop().execute(() -> {
                    try {
                        final Method invokeHandlerAddedIfNeeded = DefaultChannelPipeline.class.getDeclaredMethod("invokeHandlerAddedIfNeeded");
                        invokeHandlerAddedIfNeeded.setAccessible(true);

                        invokeHandlerAddedIfNeeded.invoke(pipeline());

                        pipeline().fireChannelRegistered();
                    } catch (IllegalAccessException | InvocationTargetException | NoSuchMethodException e) {
                        e.printStackTrace();
                    }
                });

                // Return promise
                promise.setSuccess();
            } catch (NoSuchMethodException | IllegalAccessException | InvocationTargetException | ClosedChannelException | NoSuchFieldException e) {
                e.printStackTrace();
                throw new NotImplementedException("Should register channel to event loop!!!");
            }
        }

        @Override
        public void bind(SocketAddress localAddress, ChannelPromise promise) {
            throw new NotImplementedException("");
        }

        @Override
        public void connect(SocketAddress remoteAddress, SocketAddress localAddress, ChannelPromise promise) {
            eventLoop().execute(() -> {
                try {
                    final boolean sucess = doConnect(remoteAddress, localAddress);
                    if (sucess) {
                        // Send a message to the pipeline
                        pipeline().fireChannelActive();
                        // Finally, close the promise
                        promise.setSuccess();
                    } else {
                        promise.setFailure(new RuntimeException("Unable to open the com port " + remoteAddress.toString()));
                    }
                } catch (Exception e) {
                    promise.setFailure(e);
                }
            });
        }

        @Override
        public void disconnect(ChannelPromise promise) {
            throw new NotImplementedException("");
        }

        @Override
        public void close(ChannelPromise promise) {
            logger.debug("Closing the Serial Port '{}'", this.remoteAddress());
            // TODO this should close the Serial Port
            throw new NotImplementedException("");
        }

        @Override
        public void closeForcibly() {
            throw new NotImplementedException("");
        }

        @Override
        public void deregister(ChannelPromise promise) {
            throw new NotImplementedException("");
        }

        @Override
        public final void beginRead() {
            assert eventLoop().inEventLoop();

            if (!isActive()) {
                return;
            }

            try {
                doBeginRead();
            } catch (final Exception e) {
                invokeLater(new Runnable() {
                    @Override
                    public void run() {
                        pipeline().fireExceptionCaught(e);
                    }
                });
                close(voidPromise());
            }
        }

        private void invokeLater(Runnable task) {
            try {
                // This method is used by outbound operation implementations to trigger an inbound event later.
                // They do not trigger an inbound event immediately because an outbound operation might have been
                // triggered by another inbound event handler method.  If fired immediately, the call stack
                // will look like this for example:
                //
                //   handlerA.inboundBufferUpdated() - (1) an inbound handler method closes a connection.
                //   -> handlerA.ctx.close()
                //      -> channel.unsafe.close()
                //         -> handlerA.channelInactive() - (2) another inbound handler method called while in (1) yet
                //
                // which means the execution of two inbound handler methods of the same handler overlap undesirably.
                eventLoop().execute(task);
            } catch (RejectedExecutionException e) {
                logger.warn("Can't invoke task later as EventLoop rejected it", e);
            }
        }

        @Override
        public void write(Object msg, ChannelPromise promise) {
            throw new NotImplementedException("");
        }

        @Override
        public void flush() {
            throw new NotImplementedException("");
        }

        @Override
        public ChannelPromise voidPromise() {
            return unsafeVoidPromise;
        }

        @Override
        public ChannelOutboundBuffer outboundBuffer() {
            throw new NotImplementedException("");
        }
    }
}
