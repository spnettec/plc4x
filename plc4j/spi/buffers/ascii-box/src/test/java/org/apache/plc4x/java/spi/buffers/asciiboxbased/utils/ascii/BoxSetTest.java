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
package org.apache.plc4x.java.spi.buffers.asciiboxbased.utils.ascii;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class BoxSetTest {

    private BoxSet createBoxSet(String ul, String ur, String h, String v, String ll, String lr) {
        return new BoxSet(ul, ur, h, v, ll, lr);
    }

    private AsciiBox createBox(BoxSet set, String content) {
        // AsciiBox(String) is protected, accessible from same package
        AsciiBox box = new AsciiBox(content);
        // compressedBoxSet is package-private
        box.compressedBoxSet = set.compressBoxSet();
        return box;
    }

    @Test
    void testCompressBoxSet() {
        BoxSet set = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        assertEquals("╔╗═║╚╝", set.compressBoxSet());
    }

    @Test
    void testCompressBoxSetAscii() {
        BoxSet set = createBoxSet("+", "+", "-", "|", "+", "+");
        assertEquals("++-|++", set.compressBoxSet());
    }

    @Test
    void testContributeToCompressedBoxSet_alreadyContained() {
        BoxSet set = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        AsciiBox box = createBox(set, "test content");

        String result = set.contributeToCompressedBoxSet(box);
        assertEquals(set.compressBoxSet(), result);
    }

    @Test
    void testContributeToCompressedBoxSet_notContained() {
        BoxSet set1 = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        BoxSet set2 = createBoxSet("+", "+", "-", "|", "+", "+");
        AsciiBox box = createBox(set1, "test content");

        String result = set2.contributeToCompressedBoxSet(box);
        assertTrue(result.contains(set1.compressBoxSet()));
        assertTrue(result.contains(set2.compressBoxSet()));
        assertTrue(result.contains(","));
    }

    @Test
    void testCombineCompressedBoxSets_differentSets() {
        BoxSet set1 = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        BoxSet set2 = createBoxSet("+", "+", "-", "|", "+", "+");
        AsciiBox box1 = createBox(set1, "content1");
        AsciiBox box2 = createBox(set2, "content2");

        String combined = BoxSet.combineCompressedBoxSets(box1, box2);
        assertTrue(combined.contains(set1.compressBoxSet()));
        assertTrue(combined.contains(set2.compressBoxSet()));
    }

    @Test
    void testCombineCompressedBoxSets_sameSet() {
        BoxSet set = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        AsciiBox box1 = createBox(set, "content1");
        AsciiBox box2 = createBox(set, "content2");

        String combined = BoxSet.combineCompressedBoxSets(box1, box2);
        assertFalse(combined.contains(","));
    }

    @Test
    void testEquals_sameObject() {
        BoxSet set = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        assertEquals(set, set);
    }

    @Test
    void testEquals_equalObjects() {
        BoxSet set1 = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        BoxSet set2 = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        assertEquals(set1, set2);
    }

    @Test
    void testEquals_differentObjects() {
        BoxSet set1 = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        BoxSet set2 = createBoxSet("+", "+", "-", "|", "+", "+");
        assertNotEquals(set1, set2);
    }

    @Test
    void testEquals_null() {
        BoxSet set = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        assertNotEquals(null, set);
    }

    @Test
    void testEquals_differentType() {
        BoxSet set = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        assertNotEquals("not a BoxSet", set);
    }

    @Test
    void testHashCode_equalObjects() {
        BoxSet set1 = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        BoxSet set2 = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        assertEquals(set1.hashCode(), set2.hashCode());
    }

    @Test
    void testHashCode_differentObjects() {
        BoxSet set1 = createBoxSet("╔", "╗", "═", "║", "╚", "╝");
        BoxSet set2 = createBoxSet("+", "+", "-", "|", "+", "+");
        assertNotEquals(set1.hashCode(), set2.hashCode());
    }

    @Test
    void testFieldAccess() {
        BoxSet set = createBoxSet("A", "B", "C", "D", "E", "F");
        assertEquals("A", set.upperLeftCorner);
        assertEquals("B", set.upperRightCorner);
        assertEquals("C", set.horizontalLine);
        assertEquals("D", set.verticalLine);
        assertEquals("E", set.lowerLeftCorner);
        assertEquals("F", set.lowerRightCorner);
    }
}
