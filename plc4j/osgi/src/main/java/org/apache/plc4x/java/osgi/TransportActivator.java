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
package org.apache.plc4x.java.osgi;

import org.osgi.framework.BundleActivator;
import org.osgi.framework.BundleContext;
import org.osgi.framework.ServiceRegistration;
import org.osgi.framework.wiring.BundleWiring;
import org.apache.plc4x.java.spi.transport.Transport;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.ArrayList;
import java.util.Hashtable;
import java.util.List;
import java.util.ServiceLoader;

public class TransportActivator implements BundleActivator {
    private static final Logger LOGGER = LoggerFactory.getLogger(TransportActivator.class);
    private List<ServiceRegistration<Transport>> regs = new ArrayList<>();
    private final String TRANSPORT_CODE ="org.apache.plc4x.transport.code";
    private final String TRANSPORT_NAME ="org.apache.plc4x.transport.name";


    @Override
    public void start(BundleContext context) {
        ServiceLoader<Transport> transports = ServiceLoader.load(Transport.class, context.getBundle().adapt(BundleWiring.class).getClassLoader());
        for (Transport transport : transports) {
            Hashtable<String, String> props = new Hashtable<String, String>();
            props.put(TRANSPORT_CODE, transport.getTransportCode());
            props.put(TRANSPORT_NAME, transport.getTransportName());
            regs.add(context.registerService(Transport.class, transport, props));
            LOGGER.info("register transport {}",transport.getTransportName());
        }
        LOGGER.info("register {}  transports done",regs.size());
    }

    @Override
    public void stop(BundleContext context) {
        regs.forEach(ServiceRegistration::unregister);
    }
}


