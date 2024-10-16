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
package org.apache.plc4x.java.transport.udp;

import io.netty.bootstrap.Bootstrap;
import io.netty.channel.Channel;
import io.netty.channel.ChannelPipeline;
import io.netty.channel.epoll.Epoll;
import io.netty.channel.epoll.EpollDatagramChannel;
import io.netty.channel.epoll.EpollEventLoopGroup;
import io.netty.channel.kqueue.KQueue;
import io.netty.channel.kqueue.KQueueDatagramChannel;
import io.netty.channel.kqueue.KQueueEventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioDatagramChannel;
import org.apache.plc4x.java.spi.configuration.HasConfiguration;
import org.apache.plc4x.java.spi.connection.NettyChannelFactory;
import org.apache.plc4x.java.spi.utils.ClassUtils;
import org.apache.plc4x.java.transport.udp.protocol.DatagramUnpackingHandler;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.SocketAddress;

public class UdpChannelFactory extends NettyChannelFactory implements HasConfiguration<UdpTransportConfiguration> {

    private static final Logger logger = LoggerFactory.getLogger(UdpChannelFactory.class);

    private UdpTransportConfiguration configuration;

    public UdpChannelFactory(SocketAddress remoteAddress) {
        super(remoteAddress);
    }

    public UdpChannelFactory(SocketAddress localAddress, SocketAddress remoteAddress) {
        super(localAddress, remoteAddress);
    }

    @Override
    public void setConfiguration(UdpTransportConfiguration configuration) {
        this.configuration = configuration;
    }

    @Override
    public Class<? extends Channel> getChannel() {
        if (ClassUtils.classIsPresent("io.netty.channel.epoll.Epoll") && Epoll.isAvailable()) {
            return  EpollDatagramChannel.class;
        } else if(ClassUtils.classIsPresent("io.netty.channel.kqueue.KQueue") && KQueue.isAvailable()) {
            return KQueueDatagramChannel.class;
        } else {
            return NioDatagramChannel.class;
        }
    }

    @Override
    public boolean isPassive() {
        return false;
    }

    @Override
    public void configureBootstrap(Bootstrap bootstrap) {
        if(configuration != null) {
            logger.info("Configuring Bootstrap with {}", configuration);
        }
    }

    @Override
    public void initializePipeline(ChannelPipeline pipeline) {
        pipeline.addLast(new DatagramUnpackingHandler());
    }

}
