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
package org.apache.plc4x.java.utils.connectionpool;

import org.assertj.core.api.WithAssertions;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;

/**
 * See Cahnges due to PLC4X-223 and PLC4X-224
 */
@Disabled
class AbstractPoolKeyFactoryTest implements WithAssertions {

    private PoolKeyFactory SUT = new PoolKeyFactory();

    @Nested
    class Generic {
        @Test
        void getPoolKey() throws Exception {
            AbstractPoolKey poolKey = SUT.getPoolKey("randomProtocol://randomHost/1/1?someOptions", PooledPlcConnectionManager.noPlcAuthentication);
            assertThat(poolKey.getUrl()).isEqualTo("randomProtocol://randomHost/1/1?someOptions");
            assertThat(poolKey.getPlcAuthentication()).isEqualTo(PooledPlcConnectionManager.noPlcAuthentication);
            assertThat(poolKey.getPoolableKey()).isEqualTo("randomProtocol://randomHost/1/1?someOptions");
        }
    }

    @Nested
    class S7 {
        @Test
        void getPoolKey() throws Exception {
            AbstractPoolKey poolKey = SUT.getPoolKey("s7://localhost?randomOption=true", PooledPlcConnectionManager.noPlcAuthentication);
            assertThat(poolKey.getUrl()).isEqualTo("s7://localhost?randomOption=true");
            assertThat(poolKey.getPlcAuthentication()).isEqualTo(PooledPlcConnectionManager.noPlcAuthentication);
            assertThat(poolKey.getPoolableKey()).isEqualTo("s7://localhost");
        }
    }

    @Nested
    class ADS {
        @Test
        void getPoolKey_TCP() throws Exception {
            AbstractPoolKey poolKey = SUT.getPoolKey("ads:tcp://10.10.64.40/10.10.64.40.1.1:851/10.10.56.23.1.1:30000", PooledPlcConnectionManager.noPlcAuthentication);
            assertThat(poolKey.getUrl()).isEqualTo("ads:tcp://10.10.64.40/10.10.64.40.1.1:851/10.10.56.23.1.1:30000");
            assertThat(poolKey.getPlcAuthentication()).isEqualTo(PooledPlcConnectionManager.noPlcAuthentication);
            assertThat(poolKey.getPoolableKey()).isEqualTo("ads:tcp://10.10.64.40");
        }

        @Test
        void getPoolKey_SERIAL() throws Exception {
            AbstractPoolKey poolKey = SUT.getPoolKey("ads:serial:///dev/ttys003/10.10.64.40.1.1:851/10.10.56.23.1.1:30000", PooledPlcConnectionManager.noPlcAuthentication);
            assertThat(poolKey.getUrl()).isEqualTo("ads:serial:///dev/ttys003/10.10.64.40.1.1:851/10.10.56.23.1.1:30000");
            assertThat(poolKey.getPlcAuthentication()).isEqualTo(PooledPlcConnectionManager.noPlcAuthentication);
            assertThat(poolKey.getPoolableKey()).isEqualTo("ads:serial:///dev/ttys003");
        }
    }

    @Nested
    class Modbus {
        @Test
        void getPoolKey_TCP() throws Exception {
            AbstractPoolKey poolKey = SUT.getPoolKey("modbus-tcp://10.10.64.40?someRandomOption=true", PooledPlcConnectionManager.noPlcAuthentication);
            assertThat(poolKey.getUrl()).isEqualTo("modbus-tcp://10.10.64.40?someRandomOption=true");
            assertThat(poolKey.getPlcAuthentication()).isEqualTo(PooledPlcConnectionManager.noPlcAuthentication);
            assertThat(poolKey.getPoolableKey()).isEqualTo("modbus-tcp://10.10.64.40");
        }

        @Disabled("Modbus serial pooling doesn't work right now as intended")
        @Test
        void getPoolKey_SERIAL() throws Exception {
            AbstractPoolKey poolKey = SUT.getPoolKey("modbus-adu:///dev/ttys003?someRandomOption=true", PooledPlcConnectionManager.noPlcAuthentication);
            assertThat(poolKey.getUrl()).isEqualTo("modbus-adu:///dev/ttys003?someRandomOption=true");
            assertThat(poolKey.getPlcAuthentication()).isEqualTo(PooledPlcConnectionManager.noPlcAuthentication);
            assertThat(poolKey.getPoolableKey()).isEqualTo("modbus-adu:///dev/ttys003");
        }
    }


}