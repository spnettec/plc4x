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
package org.apache.plc4x.java.opcua.security;

import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import org.apache.plc4x.java.opcua.TestCertificateGenerator;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;

/**
 * Covers the {@link CertificateVerifier#checkCertificateChainTrusted(List)} default method:
 * every verifier that doesn't override it inherits this behavior.
 */
class CertificateVerifierTest {

    @Test
    void rejectsNullChain() {
        CertificateVerifier verifier = certificate -> { };
        assertThatThrownBy(() -> verifier.checkCertificateChainTrusted(null))
            .isInstanceOf(CertificateException.class)
            .hasMessageContaining("No certificate to check");
    }

    @Test
    void rejectsEmptyChain() {
        CertificateVerifier verifier = certificate -> { };
        assertThatThrownBy(() -> verifier.checkCertificateChainTrusted(Collections.emptyList()))
            .isInstanceOf(CertificateException.class)
            .hasMessageContaining("No certificate to check");
    }

    @Test
    void checksOnlyTheLeafOfTheChain() throws CertificateException {
        X509Certificate leaf = TestCertificateGenerator.generate(2048, "CN=leaf", 3600).getValue();
        X509Certificate intermediate = TestCertificateGenerator.generate(2048, "CN=intermediate", 3600).getValue();
        List<X509Certificate> checked = new ArrayList<>();
        CertificateVerifier verifier = checked::add;

        verifier.checkCertificateChainTrusted(List.of(leaf, intermediate));

        assertThat(checked).containsExactly(leaf);
    }

    @Test
    void propagatesLeafRejection() {
        X509Certificate leaf = TestCertificateGenerator.generate(2048, "CN=leaf", 3600).getValue();
        CertificateVerifier verifier = certificate -> {
            throw new CertificateException("untrusted leaf");
        };

        assertThatThrownBy(() -> verifier.checkCertificateChainTrusted(List.of(leaf)))
            .isInstanceOf(CertificateException.class)
            .hasMessageContaining("untrusted leaf");
    }
}
