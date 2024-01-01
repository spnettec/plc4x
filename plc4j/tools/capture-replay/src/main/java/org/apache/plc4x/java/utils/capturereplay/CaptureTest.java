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
import org.pcap4j.core.PcapNetworkInterface;
import org.pcap4j.core.Pcaps;
import org.pcap4j.packet.IpV4Packet;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Arrays;

public class CaptureTest {
    private static final Logger logger = LoggerFactory.getLogger(CaptureTest.class);
    public static void main(String[] args) throws Exception {
        PacketCaptureThread.Builder builder = new PacketCaptureThread.Builder();
        PcapNetworkInterface device = Pcaps.getDevByName("en0");
        if (device == null) {
            logger.debug("");
            throw new IllegalArgumentException("Could not open output device eth0");
        }
        PacketCaptureListener listener = new PacketCaptureListener();
        listener.setCallback((packet,rawPacket) -> {

            if(packet==null){
                return;
            }
            if(isS7Packet(packet.getRawData())) {
                IpV4Packet p = rawPacket.get(IpV4Packet.class);
                if (p != null) {
                    logger.info("src ip:{}", p.getHeader().getSrcAddr());
                    logger.info("desc ip:{}", p.getHeader().getDstAddr());
                }
                logger.info("S7 packet:{}", packet);
            } else {
                //logger.info("unknown packet:{}", packet);
            }
        });
        PacketCaptureThread pThead = builder.withDevice(device)
                .withTimeout(20)
                .withFilterExpression("dst port 102")
                .withListener(listener).build();
        pThead.setDaemon(true);
        pThead.start();
        Thread.sleep(300000);
        pThead.breakLoop();
    }
    private static boolean isS7Packet(byte[] bytes){

        if(bytes.length<10 && bytes[0]!=0x03){
            return false;
        }
        ReadBufferByteBased readBuffer = new ReadBufferByteBased(bytes, ByteOrder.BIG_ENDIAN);
        TPKTPacket tPKTPacket = null;
        try {
            tPKTPacket = TPKTPacket.staticParse(readBuffer);
            logger.info(tPKTPacket.toString());
        } catch (ParseException e) {
            logger.error("error:{}", Arrays.toString(bytes));
        }
        return tPKTPacket!=null;
    }
}
