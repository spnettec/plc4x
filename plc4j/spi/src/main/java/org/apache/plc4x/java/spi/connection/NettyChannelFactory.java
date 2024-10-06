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
package org.apache.plc4x.java.spi.connection;

import io.netty.bootstrap.Bootstrap;
import io.netty.channel.Channel;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelHandler;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.epoll.Epoll;
import io.netty.channel.epoll.EpollEventLoopGroup;
import io.netty.channel.epoll.EpollSocketChannel;
import io.netty.channel.kqueue.KQueue;
import io.netty.channel.kqueue.KQueueEventLoopGroup;
import io.netty.channel.kqueue.KQueueSocketChannel;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioSocketChannel;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.spi.utils.ClassUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.SocketAddress;

/**
 * Adapter with sensible defaults for a Netty Based Channel Factory.
 * <p>
 * By Default Nettys {@link NioEventLoopGroup} is used.
 * Transports which have to use a different EventLoopGroup have to override {@link #getEventLoopGroup()}.
 */
public abstract class NettyChannelFactory implements ChannelFactory {

    private static final Logger logger = LoggerFactory.getLogger(NettyChannelFactory.class);

    private final SocketAddress localAddress;
    private final SocketAddress remoteAddress;
    private EventLoopGroup workerGroup;

    protected NettyChannelFactory(SocketAddress remoteAddress) {
        this(null, remoteAddress);
    }

    protected NettyChannelFactory(SocketAddress localAddress, SocketAddress remoteAddress) {
        this.localAddress = localAddress;
        this.remoteAddress = remoteAddress;
    }

    /**
     * Channel to Use, e.g. NiO, EiO
     */
    public Class<? extends Channel> getChannel(){
        if (ClassUtils.classIsPresent("io.netty.channel.epoll.Epoll") && Epoll.isAvailable()) {
            return EpollSocketChannel.class;
        } else if(ClassUtils.classIsPresent("io.netty.channel.kqueue.KQueue") && KQueue.isAvailable()) {
            return KQueueSocketChannel.class;
        } else {
            return NioSocketChannel.class;
        }
    }

    /**
     * We need to be able to override this in the TestChanelFactory
     *
     * @return the Bootstrap instance we will be using to initialize the channel.
     */
    protected Bootstrap createBootstrap() {
        return new Bootstrap();
    }

    /**
     * This Method is used to modify the Bootstrap Element of Netty, if one wishes to do so.
     * E.g. for Protocol Specific extension.
     * For TCP e.g.
     * <code>
     * bootstrap.option(ChannelOption.SO_KEEPALIVE, true);
     * bootstrap.option(ChannelOption.TCP_NODELAY, true);
     * </code>
     */
    public abstract void configureBootstrap(Bootstrap bootstrap);

    /**
     * Event Loop Group to use.
     * Has to be in accordance with {@link #getChannel()}
     * otherwise a Runtime Exception will be produced by Netty
     * <p>
     * By Default Netty's {@link NioEventLoopGroup} is used.
     * Transports which have to use a different EventLoopGroup have to override {#getEventLoopGroup()}.
     */
    public EventLoopGroup getEventLoopGroup() {
        if (ClassUtils.classIsPresent("io.netty.channel.epoll.Epoll") && Epoll.isAvailable()) {
            return  new EpollEventLoopGroup();
        } else if(ClassUtils.classIsPresent("io.netty.channel.kqueue.KQueue") && KQueue.isAvailable()) {
            return new KQueueEventLoopGroup();
        } else {
            return new NioEventLoopGroup();
        }
    }

    @Override
    public Channel createChannel(ChannelHandler channelHandler) throws PlcConnectionException {
        try {
            Bootstrap bootstrap = createBootstrap();

            workerGroup = getEventLoopGroup();
            if (workerGroup != null) {
                bootstrap.group(workerGroup);
            }

            bootstrap.channel(getChannel());
            // Callback to allow subclasses to modify the Bootstrap

            configureBootstrap(bootstrap);
            bootstrap.handler(channelHandler);
            // Start the client.
            final ChannelFuture f = (localAddress == null) ?
                bootstrap.connect(remoteAddress) : bootstrap.connect(remoteAddress, localAddress);
            f.addListener(future -> {
                if (!future.isSuccess()) {
                    logger.info("Unable to connect, shutting down worker thread.");
                    if (workerGroup != null) {
                        workerGroup.shutdownGracefully();
                    }
                }
            });

            final Channel channel = f.channel();

            if (workerGroup != null) {
                // Shut down the workerGroup when channel closing to avoid open too many files
                channel.closeFuture().addListener(future -> workerGroup.shutdownGracefully());
            }

            // It seems the embedded channel operates differently.
            // Intentionally using the class name as we don't want to require a
            // hard dependency on the test-channel.
            if (!"Plc4xEmbeddedChannel".equals(channel.getClass().getSimpleName())) {
                // Wait for sync
                f.sync();
                // Wait till the session is finished initializing.
                f.awaitUninterruptibly(); // jf: unsure if we need that
            }

            return channel;
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            throw new PlcConnectionException("Error creating channel.", e);
        } catch (Throwable t) {
            throw new PlcConnectionException("Error creating channel.", t);
        }
    }

    @Override
    public void closeEventLoopForChannel(Channel channel) {
        if(workerGroup == null) {
            return;
        }
        if(!(workerGroup.isShuttingDown() || workerGroup.isTerminated())) {
            workerGroup.shutdownGracefully().awaitUninterruptibly(2000);
            logger.info("Worker Group was closed successfully!");
        } else {
            logger.warn("Worker Group is ShuttingDown or isTerminated");
        }
    }

}
