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
package org.apache.plc4x.java.scraper.config;

import tools.jackson.core.JacksonException;
import tools.jackson.core.json.JsonFactory;
import tools.jackson.databind.ObjectMapper;
import tools.jackson.databind.exc.MismatchedInputException;
import tools.jackson.dataformat.yaml.YAMLFactory;
import org.apache.plc4x.java.scraper.ScrapeJob;
import org.apache.plc4x.java.scraper.exception.ScraperConfigurationException;
import org.apache.plc4x.java.scraper.exception.ScraperException;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.util.List;
import java.util.Map;

/**
 * interface for basic configuration of scraper
 */
public interface ScraperConfiguration {
    static <T>T fromYaml(String yaml, Class<T> clazz) {
        ObjectMapper mapper = new ObjectMapper(new YAMLFactory());
        try {
            return mapper.readValue(yaml, clazz);
        } catch (JacksonException e) {
            throw new ScraperConfigurationException("Unable to parse given yaml configuration!", e);
        }
    }

    static <T>T fromJson(String json, Class<T> clazz) {
        ObjectMapper mapper = new ObjectMapper(new JsonFactory());
        try {
            return mapper.readValue(json, clazz);
        } catch (JacksonException e) {
            throw new ScraperConfigurationException("Unable to parse given json configuration!", e);
        }
    }

    static <T>T fromFile(String path, Class<T> clazz) throws IOException {
        ObjectMapper mapper;
        if (path.endsWith("json")) {
            mapper = new ObjectMapper(new JsonFactory());
        } else if (path.endsWith("yml") || path.endsWith("yaml")) {
            mapper = new ObjectMapper(new YAMLFactory());
        } else {
            throw new ScraperConfigurationException("Only files with extensions json, yml or yaml can be read");
        }
        File file = new File(path);
        if (!file.exists()) {
            throw new ScraperConfigurationException("Unable to find configuration given configuration file at '" + path + "'");
        }
        try {
            return mapper.readValue(file, clazz);
        } catch (MismatchedInputException e) {
            throw new ScraperConfigurationException("Given configuration is in wrong format!", e);
        }
    }

    Map<String, String> getSources();

    List<JobConfigurationImpl> getJobConfigurations();

    @com.fasterxml.jackson.annotation.JsonIgnore
    List<ScrapeJob> getJobs() throws ScraperException;
}
