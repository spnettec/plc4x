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

package org.apache.plc4x.java.tools.eventpump.config;

import tools.jackson.core.JacksonException;
import tools.jackson.core.JsonParser;
import tools.jackson.databind.DeserializationContext;
import tools.jackson.databind.JsonNode;
import tools.jackson.databind.deser.std.StdDeserializer;

import java.util.LinkedHashMap;
import java.util.Map;

/**
 * Custom deserializer for tag maps that supports both simple and extended formats.
 * <p>
 * Handles:
 * - Simple format: "tagName": "address"
 * - Extended format: "tagName": { "address": "...", "transform": "..." }
 */
public class TagMapDeserializer extends StdDeserializer<Map<String, TagConfiguration>> {

    public TagMapDeserializer() {
        super(Map.class);
    }

    @Override
    public Map<String, TagConfiguration> deserialize(JsonParser p, DeserializationContext ctxt) throws JacksonException {
        Map<String, TagConfiguration> tags = new LinkedHashMap<>();
        JsonNode node = p.readValueAsTree();

        for (Map.Entry<String, JsonNode> entry : node.properties()) {
            String tagName = entry.getKey();
            JsonNode tagNode = entry.getValue();

            TagConfiguration tagConfig;
            if (tagNode.isTextual()) {
                // Simple format: just a string address
                tagConfig = new TagConfiguration(tagNode.asText());
            } else if (tagNode.isObject()) {
                // Extended format: object with address and optional transform
                tagConfig = ctxt.readTreeAsValue(tagNode, TagConfiguration.class);
            } else {
                throw new IllegalArgumentException("Invalid tag configuration for '" + tagName + "': expected string or object");
            }

            tags.put(tagName, tagConfig);
        }

        return tags;
    }
}
