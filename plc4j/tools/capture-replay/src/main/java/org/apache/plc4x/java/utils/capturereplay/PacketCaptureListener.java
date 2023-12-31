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

package org.apache.plc4x.java.utils.capturereplay;

import org.apache.plc4x.java.s7.readwrite.TPKTPacket;
import org.apache.plc4x.java.spi.generation.ByteOrder;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBufferByteBased;
import org.pcap4j.core.PacketListener;
import org.pcap4j.packet.Packet;
import org.pcap4j.packet.UnknownPacket;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Arrays;

public class PacketCaptureListener implements PacketListener {

    public void setCallback(Callback callback) {
        this.callback = callback;
    }

    private Callback callback;
    private static final Logger logger = LoggerFactory.getLogger(PacketCaptureListener.class);

    @Override
    public void gotPacket(Packet packet) {
        packet = packet.get(UnknownPacket.class);

        if (packet != null) {
            if (callback != null) {
                callback.call(packet);
            } else {
                byte[] bytes = packet.getRawData();
                logger.info(Arrays.toString(bytes));

                ReadBufferByteBased readBuffer = new ReadBufferByteBased(bytes, ByteOrder.BIG_ENDIAN);
                try {
                    TPKTPacket tPKTPacket = TPKTPacket.staticParse(readBuffer);
                    logger.info(tPKTPacket.toString());
                } catch (ParseException e) {
                    logger.error("error", e);
                }
            }
        }
    }

    public interface Callback {

        void call(Packet packet);
    }
}
