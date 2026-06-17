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
package org.apache.plc4x.java.spi.utils;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class ClassUtilsTest {

    @Test
    void testClassIsPresent_existingClass() {
        assertTrue(ClassUtils.classIsPresent("java.lang.String"));
    }

    @Test
    void testClassIsPresent_anotherExistingClass() {
        assertTrue(ClassUtils.classIsPresent("java.util.List"));
    }

    @Test
    void testClassIsPresent_nonExistingClass() {
        assertFalse(ClassUtils.classIsPresent("com.nonexistent.FakeClass"));
    }

    @Test
    void testClassIsPresent_emptyString() {
        assertFalse(ClassUtils.classIsPresent(""));
    }

    @Test
    void testClassIsPresent_ownClass() {
        assertTrue(ClassUtils.classIsPresent("org.apache.plc4x.java.spi.utils.ClassUtils"));
    }
}
