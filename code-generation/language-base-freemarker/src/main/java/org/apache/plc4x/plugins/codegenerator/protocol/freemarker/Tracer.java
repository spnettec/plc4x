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
package org.apache.plc4x.plugins.codegenerator.protocol.freemarker;

/**
 * Can be used to annotate generated code with traces out of the generator
 * <p>
 * uses c-style comments to inject traces. To customize sub this class and override separator(), prefix(), suffix()
 */
public class Tracer {

    protected static boolean ENABLED = System.getenv().containsKey("PLC4X_TRACE_CODE_GEN");

    protected final String currentTrace;

    protected Tracer(String currentTrace) {
        this.currentTrace = currentTrace;
    }

    /**
     * use this method to start a trace
     *
     * @param base usually the method name
     * @return a new trace containing the method name as base
     */
    public static Tracer start(String base) {
        return new Tracer(base);
    }

    /**
     * Returns a new Tracer with the appended sub
     *
     * @param sub usually a logical if name
     * @return a new trace with current trace + sub trace
     */
    public Tracer dive(String sub) {
        return new Tracer(currentTrace + separator() + sub);
    }

    protected String separator() {
        return "/";
    }

    protected String prefix() {
        return "/*";
    }

    protected String suffix() {
        return "*/";
    }

    protected boolean isEnabled() {
        return ENABLED;
    }

    @Override
    public String toString() {
        if (!isEnabled()) {
            return "";
        }
        return prefix() + currentTrace + suffix();
    }
}
