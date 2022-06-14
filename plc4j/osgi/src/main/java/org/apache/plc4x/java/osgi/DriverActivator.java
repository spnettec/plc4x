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
package org.apache.plc4x.java.osgi;

import org.apache.plc4x.java.api.PlcDriver;
import org.osgi.framework.BundleActivator;
import org.osgi.framework.BundleContext;
import org.osgi.framework.ServiceRegistration;
import org.osgi.framework.wiring.BundleWiring;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.ArrayList;
import java.util.Hashtable;
import java.util.List;
import java.util.ServiceLoader;

public class DriverActivator implements BundleActivator {
    private static final Logger LOGGER = LoggerFactory.getLogger(DriverActivator.class);
    private List<ServiceRegistration<PlcDriver>> regs = new ArrayList<>();
    public static final String PROTOCOL_NAME = "org.apache.plc4x.driver.name";
    public static final String PROTOCOL_CODE = "org.apache.plc4x.driver.code";


    @Override
    public void start(BundleContext context) {
        ServiceLoader<PlcDriver> drivers = ServiceLoader.load(PlcDriver.class, context.getBundle().adapt(BundleWiring.class).getClassLoader());
        for (PlcDriver driver : drivers) {
            Hashtable<String, String> props = new Hashtable<>();
            props.put(PROTOCOL_CODE, driver.getProtocolCode());
            props.put(PROTOCOL_NAME, driver.getProtocolName());
            regs.add(context.registerService(PlcDriver.class, driver, props));
            LOGGER.info("register driver {}",driver.getProtocolName());
        }
        LOGGER.info("register {}  drivers done",regs.size());
    }

    @Override
    public void stop(BundleContext context) {
        regs.forEach(ServiceRegistration::unregister);
    }
}

