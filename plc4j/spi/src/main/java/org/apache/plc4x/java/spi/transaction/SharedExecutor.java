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
package org.apache.plc4x.java.spi.transaction;

import org.apache.commons.lang3.concurrent.BasicThreadFactory;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

/**
 * JVM-level shared thread pools for plc4j drivers.
 *
 * Per-connection ThreadPoolExecutors accumulate when connections are cached and
 * never fully torn down. Two static pools replace them: one for the transaction
 * manager (all drivers) and one for the S7 application layer. Nothing changes
 * for single-connection use; cached-connection scenarios no longer leak threads.
 *
 * Thread counts are set to 10 (vs. the original per-connection 4) to keep
 * throughput when many connections share the same pool. Increase if you run
 * a large number of concurrent PLC connections.
 */
public class SharedExecutor {

    /** Used by RequestTransactionManager (all drivers). */
    private static final ExecutorService TM_EXECUTOR = Executors.newFixedThreadPool(
            10,
            new BasicThreadFactory.Builder()
                    .namingPattern("plc4x-tm-thread-%d")
                    .daemon(true)
                    .priority(Thread.MAX_PRIORITY)
                    .build()
    );

    /** Used by S7ProtocolLogic application layer. */
    private static final ExecutorService APP_EXECUTOR = Executors.newFixedThreadPool(
            10,
            new BasicThreadFactory.Builder()
                    .namingPattern("plc4x-app-thread-%d")
                    .daemon(true)
                    .priority(Thread.MAX_PRIORITY)
                    .build()
    );

    private SharedExecutor() {}

    public static ExecutorService getTmExecutor() {
        return TM_EXECUTOR;
    }

    public static ExecutorService getAppExecutor() {
        return APP_EXECUTOR;
    }

    /** Hard shutdown — call only on JVM exit or full driver manager teardown. */
    public static void shutdown() {
        TM_EXECUTOR.shutdownNow();
        APP_EXECUTOR.shutdownNow();
    }
}
