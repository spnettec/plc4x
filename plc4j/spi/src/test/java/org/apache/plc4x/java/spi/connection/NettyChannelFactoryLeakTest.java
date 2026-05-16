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
import io.netty.channel.ChannelInboundHandlerAdapter;
import io.netty.channel.ChannelOption;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioSocketChannel;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import static java.lang.System.out;

import java.io.IOException;
import java.net.InetAddress;
import java.net.InetSocketAddress;
import java.net.ServerSocket;
import java.net.Socket;
import java.net.SocketAddress;
import java.time.Duration;
import java.time.Instant;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.concurrent.atomic.AtomicInteger;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

/**
 * Tests {@link NettyChannelFactory} lifecycle around EventLoopGroup
 * allocation/cleanup, with special focus on the "PLC drops then recovers"
 * scenario which historically left groups (and sockets) leaked.
 *
 * <p>Three scenarios are exercised:
 * <ol>
 *   <li><b>singleFailedConnect</b> — one failed connect must result in the
 *       group being terminated (no map-based leak).</li>
 *   <li><b>stormScenario</b> — many rapid failed connects (the 100ms-tick
 *       wire-driven retry pattern) must all terminate within bounded time;
 *       captures peak alive for baseline vs. fix comparison.</li>
 *   <li><b>recoveryCycle</b> — connect-success → server-down (refused) →
 *       server-back-up must still produce a working channel and leave no
 *       EventLoopGroup behind.</li>
 * </ol>
 */
class NettyChannelFactoryLeakTest {

    private static final Logger LOGGER = LoggerFactory.getLogger(NettyChannelFactoryLeakTest.class);

    /** 127.0.0.1:1 — kernel returns ECONNREFUSED immediately. */
    private static final SocketAddress UNREACHABLE = new InetSocketAddress("127.0.0.1", 1);

    private static final int STORM_ATTEMPTS = 20;
    private static final Duration TERMINATION_DEADLINE = Duration.ofSeconds(25);

    private ToggleServer server;

    @AfterEach
    void stopServer() throws IOException {
        if (server != null) {
            server.close();
            server = null;
        }
    }

    @Test
    void singleFailedConnect_groupTerminates() throws InterruptedException {
        InstrumentedFactory factory = new InstrumentedFactory(UNREACHABLE);

        assertThrows(PlcConnectionException.class,
            () -> factory.createChannel(new ChannelInboundHandlerAdapter()));

        waitForTermination(factory, 1);

        out.printf("[METRIC] single attempt: created=%d terminated=%d peakAlive=%d%n",
            factory.created.get(), factory.terminated.get(), factory.peakAlive.get());
        assertEquals(1, factory.terminated.get(), "group should terminate after failed connect");
    }

    @Test
    void stormScenario_allGroupsTerminate() throws InterruptedException {
        InstrumentedFactory factory = new InstrumentedFactory(UNREACHABLE);

        Instant startCreate = Instant.now();
        for (int i = 0; i < STORM_ATTEMPTS; i++) {
            try {
                factory.createChannel(new ChannelInboundHandlerAdapter());
            } catch (PlcConnectionException expected) {
                // simulated upstream caller continues retrying on next tick
            }
        }
        Duration createPhase = Duration.between(startCreate, Instant.now());

        Instant startCleanup = Instant.now();
        waitForTermination(factory, STORM_ATTEMPTS);
        Duration cleanupPhase = Duration.between(startCleanup, Instant.now());

        out.printf("[METRIC] storm: attempts=%d createPhase=%dms peakAlive=%d cleanupPhase=%dms%n",
            STORM_ATTEMPTS, createPhase.toMillis(), factory.peakAlive.get(), cleanupPhase.toMillis());

        assertEquals(STORM_ATTEMPTS, factory.terminated.get(),
            "all " + STORM_ATTEMPTS + " groups should terminate within deadline");
    }

    /**
     * The historical motivation for the YOFC fork's NettyChannelFactory
     * change. Connect once successfully, drop the "PLC", retry while down,
     * then bring it back. After all of that, every group should have been
     * shut down and a fresh connect must succeed.
     */
    @Test
    void recoveryCycle_connectAfterFailuresStillWorks() throws Exception {
        server = ToggleServer.startOnFreePort();
        InstrumentedFactory factory = new InstrumentedFactory(server.address());

        // Phase 1 — PLC reachable. Connect succeeds, then we close the channel.
        Channel channel = factory.createChannel(new ChannelInboundHandlerAdapter());
        assertNotNull(channel, "channel from successful connect");
        assertTrue(channel.isActive(), "channel should be active after successful connect");
        channel.close().syncUninterruptibly();
        int afterPhase1 = factory.created.get();
        out.printf("[METRIC] phase1 up:   created=%d alive=%d%n", afterPhase1, factory.alive.get());

        // Phase 2 — PLC down. Several failed retries (mirroring wire 100ms tick).
        server.stop();
        for (int i = 0; i < 5; i++) {
            assertThrows(PlcConnectionException.class,
                () -> factory.createChannel(new ChannelInboundHandlerAdapter()));
        }
        int afterPhase2 = factory.created.get();
        out.printf("[METRIC] phase2 down: created=%d peakAlive=%d%n", afterPhase2, factory.peakAlive.get());

        // Phase 3 — PLC reachable again on the SAME port. Recovery must succeed.
        server.restart();
        Channel recovered = factory.createChannel(new ChannelInboundHandlerAdapter());
        assertNotNull(recovered, "channel after recovery");
        assertTrue(recovered.isActive(), "channel should be active after recovery");
        recovered.close().syncUninterruptibly();

        // All groups (success + failures + recovery) must eventually terminate.
        waitForTermination(factory, factory.created.get());

        out.printf("[METRIC] cycle done: created=%d terminated=%d peakAlive=%d%n",
            factory.created.get(), factory.terminated.get(), factory.peakAlive.get());
        assertEquals(factory.created.get(), factory.terminated.get(),
            "every allocated group should terminate by the end of the cycle");
    }

    private static void waitForTermination(InstrumentedFactory factory, int expected) throws InterruptedException {
        long deadlineNanos = System.nanoTime() + TERMINATION_DEADLINE.toNanos();
        while (factory.terminated.get() < expected && System.nanoTime() < deadlineNanos) {
            Thread.sleep(100);
        }
    }

    /**
     * Concrete factory mimicking a TCP transport (NIO + connect timeout),
     * with hooks to count every EventLoopGroup's lifecycle.
     */
    private static final class InstrumentedFactory extends NettyChannelFactory {

        final AtomicInteger created = new AtomicInteger();
        final AtomicInteger terminated = new AtomicInteger();
        final AtomicInteger alive = new AtomicInteger();
        final AtomicInteger peakAlive = new AtomicInteger();

        InstrumentedFactory(SocketAddress remoteAddress) {
            super(remoteAddress);
        }

        @Override
        public Class<? extends Channel> getChannel() {
            return NioSocketChannel.class;
        }

        @Override
        public boolean isPassive() {
            return false;
        }

        @Override
        public EventLoopGroup getEventLoopGroup() {
            NioEventLoopGroup group = new NioEventLoopGroup(1);
            int currentAlive = alive.incrementAndGet();
            created.incrementAndGet();
            peakAlive.updateAndGet(prev -> Math.max(prev, currentAlive));
            group.terminationFuture().addListener(future -> {
                alive.decrementAndGet();
                terminated.incrementAndGet();
            });
            return group;
        }

        @Override
        public void configureBootstrap(Bootstrap bootstrap) {
            bootstrap.option(ChannelOption.CONNECT_TIMEOUT_MILLIS, 500);
        }
    }

    /**
     * Trivially accepts inbound TCP connections on a fixed loopback port.
     * Supports stop/restart on the same port so the recovery scenario can be
     * exercised without changing the factory's target address.
     */
    private static final class ToggleServer implements AutoCloseable {

        private final int port;
        private ServerSocket serverSocket;
        private Thread acceptor;
        private final List<Socket> accepted = Collections.synchronizedList(new ArrayList<>());
        private volatile boolean running;

        private ToggleServer(int port) {
            this.port = port;
        }

        static ToggleServer startOnFreePort() throws IOException {
            int port;
            try (ServerSocket probe = new ServerSocket(0, 0, InetAddress.getLoopbackAddress())) {
                probe.setReuseAddress(true);
                port = probe.getLocalPort();
            }
            ToggleServer s = new ToggleServer(port);
            s.start();
            return s;
        }

        InetSocketAddress address() {
            return new InetSocketAddress(InetAddress.getLoopbackAddress(), port);
        }

        private void start() throws IOException {
            serverSocket = new ServerSocket();
            serverSocket.setReuseAddress(true);
            serverSocket.bind(new InetSocketAddress(InetAddress.getLoopbackAddress(), port));
            running = true;
            acceptor = new Thread(this::acceptLoop, "ToggleServer-" + port);
            acceptor.setDaemon(true);
            acceptor.start();
        }

        void stop() throws IOException {
            running = false;
            if (serverSocket != null) {
                serverSocket.close();
                serverSocket = null;
            }
            for (Socket s : accepted) {
                try { s.close(); } catch (IOException ignored) { /* */ }
            }
            accepted.clear();
            if (acceptor != null) {
                try { acceptor.join(2000); } catch (InterruptedException ignored) { Thread.currentThread().interrupt(); }
                acceptor = null;
            }
        }

        void restart() throws IOException {
            start();
        }

        private void acceptLoop() {
            while (running) {
                try {
                    Socket client = serverSocket.accept();
                    accepted.add(client);
                } catch (IOException e) {
                    if (running) {
                        LOGGER.warn("ToggleServer accept failed", e);
                    }
                    return;
                }
            }
        }

        @Override
        public void close() throws IOException {
            stop();
        }
    }
}
