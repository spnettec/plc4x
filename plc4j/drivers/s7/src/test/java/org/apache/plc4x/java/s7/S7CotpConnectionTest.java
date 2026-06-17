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
package org.apache.plc4x.java.s7;

import org.apache.plc4x.java.s7.configuration.S7Configuration;
import org.apache.plc4x.java.s7.readwrite.ControllerType;
import org.apache.plc4x.java.spi.transports.api.TransportInstance;
import org.apache.plc4x.java.utils.auditlog.api.AuditLog;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;

import java.lang.reflect.Method;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotEquals;

class S7CotpConnectionTest {

    @Test
    void s7200UsesZeroTpduReference() throws Exception {
        S7CotpConnection connection = newConnection(ControllerType.S7_200);

        assertEquals(0, nextTpduId(connection));
        assertEquals(0, nextTpduId(connection));
    }

    @Test
    void s7300UsesNonZeroIncrementingTpduReference() throws Exception {
        S7CotpConnection connection = newConnection(ControllerType.S7_300);

        int first = nextTpduId(connection);
        int second = nextTpduId(connection);
        assertNotEquals(0, first);
        assertEquals(first + 1, second);
    }

    private static S7CotpConnection newConnection(ControllerType controllerType) {
        S7Configuration configuration = new S7Configuration();
        configuration.setControllerType(controllerType);
        configuration.setMaxAmqCallee(4);
        TransportInstance<?> transport = Mockito.mock(TransportInstance.class);
        AuditLog auditLog = Mockito.mock(AuditLog.class);
        return new S7CotpConnection(configuration, transport, auditLog);
    }

    private static int nextTpduId(S7CotpConnection connection) throws Exception {
        Method method = S7CotpConnection.class.getDeclaredMethod("getTpduId");
        method.setAccessible(true);
        return (int) method.invoke(connection);
    }

}
