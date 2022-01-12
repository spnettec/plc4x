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
package org.apache.plc4x.java.bacnetip;

import com.vdurmont.semver4j.Semver;
import org.apache.commons.io.FileUtils;
import org.apache.commons.io.IOUtils;
import org.apache.commons.lang3.StringUtils;
import org.apache.commons.lang3.SystemUtils;
import org.apache.plc4x.java.bacnetip.readwrite.*;
import org.apache.plc4x.java.bacnetip.readwrite.io.BVLCIO;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBufferByteBased;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WriteBufferBoxBased;
import org.apache.plc4x.java.spi.utils.Serializable;
import org.apache.plc4x.java.spi.utils.hex.Hex;
import org.junit.jupiter.api.*;
import org.pcap4j.core.NotOpenException;
import org.pcap4j.core.PcapHandle;
import org.pcap4j.core.PcapNativeException;
import org.pcap4j.core.Pcaps;
import org.pcap4j.packet.Packet;
import org.pcap4j.packet.UdpPacket;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.Closeable;
import java.io.File;
import java.io.IOException;
import java.math.BigInteger;
import java.net.URL;
import java.nio.file.FileSystems;
import java.util.*;
import java.util.concurrent.ConcurrentLinkedDeque;
import java.util.stream.IntStream;

import static org.junit.jupiter.api.Assertions.*;
import static org.junit.jupiter.api.Assumptions.assumeTrue;

// Tests from http://kargs.net/captures
public class RandomPackagesTest {

    private static final Logger LOGGER = LoggerFactory.getLogger(RandomPackagesTest.class);

    @BeforeAll
    static void setUp() {
        // TODO: for mac only don't commit
        //System.getProperties().setProperty("jna.library.path", "/usr/local/Cellar/libpcap//1.10.1/lib");
        assumeTrue(() -> {
            try {
                String version = Pcaps.libVersion();
                LOGGER.info("Pcap version: " + version);
                String libpcap_version_string = StringUtils.removeStart(version, "libpcap version ");
                // Remove any trailing extra info
                libpcap_version_string = StringUtils.split(libpcap_version_string, " ")[0];
                Semver libpcap_version = new Semver(libpcap_version_string);
                if (SystemUtils.IS_OS_MAC) {
                    Semver minimumVersion = new Semver("1.10.1");

                    if (libpcap_version.isLowerThan(minimumVersion)) {
                        LOGGER.info("pcap with at least " + minimumVersion + " required.");
                        return false;
                    }
                }
            } catch (Exception | Error e) {
                e.printStackTrace();
                return false;
            }
            return true;
        }, "no pcap version on system");
    }

    Queue<Closeable> toBeClosed = new ConcurrentLinkedDeque<>();

    @AfterEach
    void closeStuff() {
        for (Closeable closeable = toBeClosed.poll(); closeable != null; closeable = toBeClosed.poll()) {
            LOGGER.info("Closing closeable " + closeable);
            IOUtils.closeQuietly(closeable);
        }
    }

    @TestFactory
    @DisplayName("BACnet-BBMD-on-same-subnet")
    Collection<DynamicNode> BACnet_BBMD_on_same_subnet() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("BACnet-BBMD-on-same-subnet.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("BACnet Virtual Link Control BVLC Function Register-Foreign-Device",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed whoIs",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("BACnet Virtual Link Control BVLC Function BVLC-Results",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Is",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,123",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,123",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,18",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,18",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,2401",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,2401",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,86114",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,86114",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,884456",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,884456",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("BACnet-MSTP-SNAP-Mixed")
    Collection<DynamicNode> BACnet_MSTP_SNAP_Mixed() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("BACnet-MSTP-SNAP-Mixed.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("BACnetARRAY-element-0")
    Collection<DynamicNode> BACnetARRAY_element_0() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("BACnetARRAY-element-0.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("BACnetARRAY-elements")
    Collection<DynamicNode> BACnetARRAY_elements() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("BACnetARRAY-elements.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("BACnetDeviceObjectReference")
    Collection<DynamicNode> BACnetDeviceObjectReference() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("BACnetDeviceObjectReference.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[  0] life-safety-zone,1 zone-members",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.LIFE_SAFETY_ZONE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(1, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.ZONE_MEMBERS, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[  0] life-safety-zone,1 zone-members life-safety-zone,3 life-safety-zone,4 life-safety-zone,5 life-safety-zone,6 life-safety-zone,7 life-safety-zone,8 life-safety-zone,9 life-safety-zone,16 life-safety-zone,494 life-safety-zone,255 life-safety-zone,231 life-safety-zone,4193620 life-safety-zone,222 life-safety-zone,300 life-safety-zone,166",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.LIFE_SAFETY_ZONE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(1, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.ZONE_MEMBERS, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    // TODO: assert object identifiers
                })
        );
    }

    @Disabled("mostly llc so we don't use that here for now")
    @TestFactory
    @DisplayName("BACnetIP-MSTP-Mix")
    Collection<DynamicNode> BACnet_MSTP_Mix() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("BACnetIP-MSTP-Mix.cap");
        return null;
    }

    @Disabled("mostly llc so we don't use that here for now")
    @TestFactory
    @DisplayName("BBMD_Results")
    Collection<DynamicNode> BBMD_Results() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("BBMD_Results.pcap");
        return null;
    }

    @TestFactory
    @DisplayName("BBMD_readproperty")
    Collection<DynamicNode> BBMD_readProperty() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("BBMD_readproperty.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Is 12345 12345",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCDistributeBroadcastToNetwork bvlcDistributeBroadcastToNetwork = (BVLCDistributeBroadcastToNetwork) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcDistributeBroadcastToNetwork.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequest serviceRequest = apduUnconfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetUnconfirmedServiceRequestWhoIs baCnetUnconfirmedServiceRequestWhoIs = (BACnetUnconfirmedServiceRequestWhoIs) serviceRequest;
                    assertEquals(12345L, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeLowLimit().getActualValue());
                    assertEquals(12345L, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeHighLimit().getActualValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Is 12345 12345",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequest serviceRequest = apduUnconfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetUnconfirmedServiceRequestWhoIs baCnetUnconfirmedServiceRequestWhoIs = (BACnetUnconfirmedServiceRequestWhoIs) serviceRequest;
                    assertEquals(12345L, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeLowLimit().getActualValue());
                    assertEquals(12345L, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeHighLimit().getActualValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,12345",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequest serviceRequest = apduUnconfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetUnconfirmedServiceRequestIAm baCnetUnconfirmedServiceRequestIAm = (BACnetUnconfirmedServiceRequestIAm) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIAm.getDeviceIdentifier().getObjectType());
                    assertEquals(12345, baCnetUnconfirmedServiceRequestIAm.getDeviceIdentifier().getInstanceNumber());
                    assertEquals(480, baCnetUnconfirmedServiceRequestIAm.getMaximumApduLengthAcceptedLength().getActualValue());
                    // TODO: change to enum
                    assertEquals(List.of((byte) 0x03), baCnetUnconfirmedServiceRequestIAm.getSegmentationSupported().getData());
                    assertEquals(260L, baCnetUnconfirmedServiceRequestIAm.getVendorId().getActualValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Is 12345 12345",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCForwardedNPDU bvlcForwardedNPDU = (BVLCForwardedNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcForwardedNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequest serviceRequest = apduUnconfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetUnconfirmedServiceRequestWhoIs baCnetUnconfirmedServiceRequestWhoIs = (BACnetUnconfirmedServiceRequestWhoIs) serviceRequest;
                    assertEquals(12345L, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeLowLimit().getActualValue());
                    assertEquals(12345L, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeHighLimit().getActualValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Is 12345 12345",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCForwardedNPDU bvlcForwardedNPDU = (BVLCForwardedNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcForwardedNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequest serviceRequest = apduUnconfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetUnconfirmedServiceRequestWhoIs baCnetUnconfirmedServiceRequestWhoIs = (BACnetUnconfirmedServiceRequestWhoIs) serviceRequest;
                    assertEquals(12345L, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeLowLimit().getActualValue());
                    assertEquals(12345L, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeHighLimit().getActualValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,12345",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCForwardedNPDU bvlcForwardedNPDU = (BVLCForwardedNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcForwardedNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequest serviceRequest = apduUnconfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetUnconfirmedServiceRequestIAm baCnetUnconfirmedServiceRequestIAm = (BACnetUnconfirmedServiceRequestIAm) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIAm.getDeviceIdentifier().getObjectType());
                    assertEquals(12345, baCnetUnconfirmedServiceRequestIAm.getDeviceIdentifier().getInstanceNumber());
                    assertEquals(480, baCnetUnconfirmedServiceRequestIAm.getMaximumApduLengthAcceptedLength().getActualValue());
                    // TODO: change to enum
                    assertEquals(List.of((byte) 0x03), baCnetUnconfirmedServiceRequestIAm.getSegmentationSupported().getData());
                    assertEquals(260L, baCnetUnconfirmedServiceRequestIAm.getVendorId().getActualValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ readProperty[ 1] analog-output,0 priority-array",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.ANALOG_OUTPUT, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(0, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PRIORITY_ARRAY, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK readProperty[ 1] analog-output,0 priority-array",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.ANALOG_OUTPUT, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(0, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PRIORITY_ARRAY, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    /* FIXME: we get now a bunch of tags here
                    BACnetPropertyValuePriorityValue baCnetPropertyValuePriorityValue = (BACnetPropertyValuePriorityValue) baCnetServiceAckReadProperty.getValues().getData();
                    assertArrayEquals(new byte[]{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, baCnetPropertyValuePriorityValue.getValues());
                     */
                }),
            DynamicTest.dynamicTest("BACnet Virtual Link Control BVLC Function Register-Foreign-Device",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCRegisterForeignDevice bvlcRegisterForeignDevice = (BVLCRegisterForeignDevice) bvlc;
                    assertEquals(60000, bvlcRegisterForeignDevice.getTtl());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Is 12345 12345",
                () -> {
                    // this is a repeat from the package above
                    pcapEvaluator.skipPackages(1);
                }),
            DynamicTest.dynamicTest("BACnet Virtual Link Control BVLC Function BVLC-Result",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCResult bvlcResult = (BVLCResult) bvlc;
                    assertEquals(BVLCResultCode.SUCCESSFUL_COMPLETION, bvlcResult.getCode());
                }),
            DynamicTest.dynamicTest("Skip Unconfirmed-REQ who-Is/I-Am",
                () -> {
                    // this is a repeat from the package above
                    pcapEvaluator.skipPackages(5);
                }),
            DynamicTest.dynamicTest("Confirmed-REQ readProperty[ 1] analog-output,0 present-value",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.ANALOG_OUTPUT, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(0, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PRESENT_VALUE, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK readProperty[ 1] analog-output,0 present-value",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.ANALOG_OUTPUT, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(0, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PRESENT_VALUE, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTagReal baCnetApplicationTagReal = (BACnetApplicationTagReal) baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    assertEquals(0, baCnetApplicationTagReal.getValue());
                }),
            DynamicTest.dynamicTest("Skip Misc packages",
                () -> {
                    // this is a repeat from the package above
                    pcapEvaluator.skipPackages(8);
                }),
            DynamicTest.dynamicTest("Confirmed-REQ readProperty[ 1] analog-output,0 relinquish-default",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.ANALOG_OUTPUT, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(0, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.RELINQUISH_DEFAULT, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK readProperty[ 1] analog-output,0 relinquish-default",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.ANALOG_OUTPUT, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(0, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.RELINQUISH_DEFAULT, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    /* FIXME: wrong data here too
                    BACnetApplicationTagReal baCnetApplicationTagReal = (BACnetApplicationTagReal) baCnetServiceAckReadProperty.getValues().getData().get(0);
                    assertEquals(0f, baCnetApplicationTagReal);
                     */
                }),
            DynamicTest.dynamicTest("Skip Misc packages",
                () -> {
                    // this is a repeat from the package above
                    pcapEvaluator.skipPackages(48);
                }),
            DynamicTest.dynamicTest("Confirmed-REQ writeProperty[ 1] analog-output,0 priority-array",
                () -> {
                    // This package is broken as from the spec it requires 16 values // TODO: validate that
                    pcapEvaluator.skipPackages(1);
                }),
            DynamicTest.dynamicTest("Error writeProperty[ 1]",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUError apduError = (APDUError) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetErrorWriteProperty baCnetErrorWriteProperty = (BACnetErrorWriteProperty) apduError.getError();
                    // TODO: change to enum
                    assertEquals(List.of((byte) 0x02), baCnetErrorWriteProperty.getErrorClass().getData());
                    // TODO: change to enum
                    assertEquals(List.of((byte) 0x28), baCnetErrorWriteProperty.getErrorCode().getData());
                }),
            DynamicTest.dynamicTest("Skip Misc 8 packages",
                () -> {
                    // this is a repeat from the package above
                    pcapEvaluator.skipPackages(8);
                }),
            DynamicTest.dynamicTest("Confirmed-REQ writeProperty[ 1] analog-output,0 present-value",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestWriteProperty baCnetConfirmedServiceRequestWriteProperty = (BACnetConfirmedServiceRequestWriteProperty) serviceRequest;
                    assertEquals(BACnetObjectType.ANALOG_OUTPUT, baCnetConfirmedServiceRequestWriteProperty.getObjectIdentifier().getObjectType());
                    assertEquals(0, baCnetConfirmedServiceRequestWriteProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PRESENT_VALUE, baCnetConfirmedServiceRequestWriteProperty.getPropertyIdentifier().getValue());

                    BACnetApplicationTagReal baCnetApplicationTagReal = (BACnetApplicationTagReal) baCnetConfirmedServiceRequestWriteProperty.getPropertyValue().getData().get(0).getApplicationTag();
                    assertEquals(123.449997f, baCnetApplicationTagReal.getValue());
                    BACnetContextTagUnsignedInteger priority = baCnetConfirmedServiceRequestWriteProperty.getPriority();
                    assertEquals(10, priority.getActualValue());
                }),
            DynamicTest.dynamicTest("Error writeProperty[ 1]",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUError apduError = (APDUError) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetErrorWriteProperty baCnetErrorWriteProperty = (BACnetErrorWriteProperty) apduError.getError();
                    // TODO: change to enum
                    assertEquals(List.of((byte) 0x02), baCnetErrorWriteProperty.getErrorClass().getData());
                    // TODO: change to enum
                    assertEquals(List.of((byte) 0x25), baCnetErrorWriteProperty.getErrorCode().getData());
                }),
            DynamicTest.dynamicTest("Skip Misc packages",
                () -> {
                    // this is a repeat from the package above
                    pcapEvaluator.skipTo(143);
                }),
            DynamicTest.dynamicTest("Confirmed-REQ writeProperty[ 1] analog-output,0 present-value",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestWriteProperty baCnetConfirmedServiceRequestWriteProperty = (BACnetConfirmedServiceRequestWriteProperty) serviceRequest;
                    assertEquals(BACnetObjectType.ANALOG_OUTPUT, baCnetConfirmedServiceRequestWriteProperty.getObjectIdentifier().getObjectType());
                    assertEquals(0, baCnetConfirmedServiceRequestWriteProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PRESENT_VALUE, baCnetConfirmedServiceRequestWriteProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTagNull baCnetApplicationTagNull = (BACnetApplicationTagNull) baCnetConfirmedServiceRequestWriteProperty.getPropertyValue().getData().get(0).getApplicationTag();
                    assertNotNull(baCnetApplicationTagNull);
                    BACnetContextTagUnsignedInteger priority = baCnetConfirmedServiceRequestWriteProperty.getPriority();
                    assertEquals(1, priority.getActualValue());
                }),
            DynamicTest.dynamicTest("Simple-ACK writeProperty[ 1]", () -> {
                BVLC bvlc = pcapEvaluator.nextBVLC();
                dump(bvlc);
                BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                APDUSimpleAck apduSimpleAck = (APDUSimpleAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                assertEquals(15, apduSimpleAck.getServiceChoice());
            }),
            DynamicTest.dynamicTest("Skip Misc packages",
                () -> {
                    // this is a repeat from the package above
                    pcapEvaluator.skipTo(201);
                }),
            DynamicTest.dynamicTest("Confirmed-REQ readProperty[  1] device,12345 object-identifier",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(12345, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_IDENTIFIER, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK   readProperty[  1] device,12345 object-identifier device,12345",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(12345, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_IDENTIFIER, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTagObjectIdentifier objectIdentifier = (BACnetApplicationTagObjectIdentifier) baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    assertEquals(BACnetObjectType.DEVICE, objectIdentifier.getObjectType());
                    assertEquals(12345, objectIdentifier.getInstanceNumber());
                }),
            parseEmAll(pcapEvaluator, 202, 337)
        );
    }

    @TestFactory
    @DisplayName("CEN_9_11")
    Collection<DynamicNode> CEN_9_11() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("CEN_9_11.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ   confirmedEventNotification[119] event-enrollment,11 analog-input,1",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestConfirmedEventNotification baCnetConfirmedServiceRequestConfirmedEventNotification = (BACnetConfirmedServiceRequestConfirmedEventNotification) serviceRequest;
                    assertEquals((short) 111, baCnetConfirmedServiceRequestConfirmedEventNotification.getProcessIdentifier().getValueUint8());
                    assertEquals(BACnetObjectType.EVENT_ENROLLMENT, baCnetConfirmedServiceRequestConfirmedEventNotification.getInitiatingDeviceIdentifier().getObjectType());
                    assertEquals(11, baCnetConfirmedServiceRequestConfirmedEventNotification.getInitiatingDeviceIdentifier().getInstanceNumber());
                    assertEquals(BACnetObjectType.ANALOG_INPUT, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventObjectIdentifier().getObjectType());
                    assertEquals(1, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventObjectIdentifier().getInstanceNumber());
                    {
                        BACnetTimeStampSequence timestamp = (BACnetTimeStampSequence) baCnetConfirmedServiceRequestConfirmedEventNotification.getTimestamp();
                        assertEquals(2, timestamp.getSequenceNumber().getActualValue());
                    }
                    {
                        assertEquals(1, baCnetConfirmedServiceRequestConfirmedEventNotification.getNotificationClass().getActualValue());
                    }
                    {
                        assertEquals(1, baCnetConfirmedServiceRequestConfirmedEventNotification.getPriority().getActualValue());
                    }
                    {
                        assertEquals(BACnetEventType.UNSIGNED_RANGE, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventType().getValue());
                    }
                    {
                        assertEquals("My Message", baCnetConfirmedServiceRequestConfirmedEventNotification.getMessageText().getValue());
                    }
                    {
                        assertEquals(BACnetNotifyType.EVENT, baCnetConfirmedServiceRequestConfirmedEventNotification.getNotifyType().getValue());
                    }
                    {
                        assertTrue(baCnetConfirmedServiceRequestConfirmedEventNotification.getAckRequired().getIsFalse());
                    }
                    {
                        assertEquals(BACnetEventState.NORMAL, baCnetConfirmedServiceRequestConfirmedEventNotification.getFromState().getValue());
                    }
                    {
                        assertEquals(BACnetEventState.OFFNORMAL, baCnetConfirmedServiceRequestConfirmedEventNotification.getToState().getValue());
                    }
                    {
                        BACnetNotificationParametersUnsignedRange baCnetNotificationParametersUnsignedRange = (BACnetNotificationParametersUnsignedRange) baCnetConfirmedServiceRequestConfirmedEventNotification.getEventValues();
                        assertEquals(50, baCnetNotificationParametersUnsignedRange.getSequenceNumber().getActualValue());
                        assertTrue(baCnetNotificationParametersUnsignedRange.getStatusFlags().getInAlarm());
                        assertFalse(baCnetNotificationParametersUnsignedRange.getStatusFlags().getFault());
                        assertFalse(baCnetNotificationParametersUnsignedRange.getStatusFlags().getOverriden());
                        assertFalse(baCnetNotificationParametersUnsignedRange.getStatusFlags().getOutOfService());
                        assertEquals(40, baCnetNotificationParametersUnsignedRange.getExceededLimit().getActualValue());
                    }
                }),
            DynamicTest.dynamicTest("Simple-ACK      confirmedEventNotification[119]",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUSimpleAck apduSimpleAck = (APDUSimpleAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    assertEquals(119, apduSimpleAck.getOriginalInvokeId());
                    assertEquals(2, apduSimpleAck.getServiceChoice());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   confirmedEventNotification[120] event-enrollment,11 analog-input,1",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestConfirmedEventNotification baCnetConfirmedServiceRequestConfirmedEventNotification = (BACnetConfirmedServiceRequestConfirmedEventNotification) serviceRequest;
                    assertEquals((short) 111, baCnetConfirmedServiceRequestConfirmedEventNotification.getProcessIdentifier().getValueUint8());
                    assertEquals(BACnetObjectType.EVENT_ENROLLMENT, baCnetConfirmedServiceRequestConfirmedEventNotification.getInitiatingDeviceIdentifier().getObjectType());
                    assertEquals(11, baCnetConfirmedServiceRequestConfirmedEventNotification.getInitiatingDeviceIdentifier().getInstanceNumber());
                    assertEquals(BACnetObjectType.ANALOG_INPUT, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventObjectIdentifier().getObjectType());
                    assertEquals(1, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventObjectIdentifier().getInstanceNumber());
                    {
                        BACnetTimeStampSequence timestamp = (BACnetTimeStampSequence) baCnetConfirmedServiceRequestConfirmedEventNotification.getTimestamp();
                        assertEquals(2, timestamp.getSequenceNumber().getActualValue());
                    }
                    {
                        assertEquals(1, baCnetConfirmedServiceRequestConfirmedEventNotification.getNotificationClass().getActualValue());
                    }
                    {
                        assertEquals(1, baCnetConfirmedServiceRequestConfirmedEventNotification.getPriority().getActualValue());
                    }
                    {
                        assertEquals(BACnetEventType.EXTENDED, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventType().getValue());
                    }
                    {
                        assertEquals("My Message", baCnetConfirmedServiceRequestConfirmedEventNotification.getMessageText().getValue());
                    }
                    {
                        assertEquals(BACnetNotifyType.EVENT, baCnetConfirmedServiceRequestConfirmedEventNotification.getNotifyType().getValue());
                    }
                    {
                        assertTrue(baCnetConfirmedServiceRequestConfirmedEventNotification.getAckRequired().getIsFalse());
                    }
                    {
                        assertEquals(BACnetEventState.NORMAL, baCnetConfirmedServiceRequestConfirmedEventNotification.getFromState().getValue());
                    }
                    {
                        assertEquals(BACnetEventState.OFFNORMAL, baCnetConfirmedServiceRequestConfirmedEventNotification.getToState().getValue());
                    }
                    {
                        BACnetNotificationParametersExtended baCnetNotificationParametersExtended = (BACnetNotificationParametersExtended) baCnetConfirmedServiceRequestConfirmedEventNotification.getEventValues();
                        assertEquals(24, baCnetNotificationParametersExtended.getVendorId().getActualValue());
                        assertEquals(33, baCnetNotificationParametersExtended.getExtendedEventType().getActualValue());
                    }
                }),
            DynamicTest.dynamicTest("Simple-ACK      confirmedEventNotification[120]",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUSimpleAck apduSimpleAck = (APDUSimpleAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    assertEquals(120, apduSimpleAck.getOriginalInvokeId());
                    assertEquals(2, apduSimpleAck.getServiceChoice());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Is 140 140",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    NPDU npdu = ((BVLCOriginalBroadcastNPDU) bvlc).getNpdu();
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) npdu.getApdu();
                    BACnetUnconfirmedServiceRequestWhoIs baCnetUnconfirmedServiceRequestWhoIs = (BACnetUnconfirmedServiceRequestWhoIs) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(140, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeLowLimit().getActualValue());
                    assertEquals(140, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeHighLimit().getActualValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Is 140 140",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    NPDU npdu = ((BVLCOriginalBroadcastNPDU) bvlc).getNpdu();
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) npdu.getApdu();
                    BACnetUnconfirmedServiceRequestWhoIs baCnetUnconfirmedServiceRequestWhoIs = (BACnetUnconfirmedServiceRequestWhoIs) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(140, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeLowLimit().getActualValue());
                    assertEquals(140, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeHighLimit().getActualValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Is 871 871",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    NPDU npdu = ((BVLCOriginalBroadcastNPDU) bvlc).getNpdu();
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) npdu.getApdu();
                    BACnetUnconfirmedServiceRequestWhoIs baCnetUnconfirmedServiceRequestWhoIs = (BACnetUnconfirmedServiceRequestWhoIs) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(871, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeLowLimit().getActualValue());
                    assertEquals(871, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeHighLimit().getActualValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Is 871 871",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    NPDU npdu = ((BVLCOriginalBroadcastNPDU) bvlc).getNpdu();
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) npdu.getApdu();
                    BACnetUnconfirmedServiceRequestWhoIs baCnetUnconfirmedServiceRequestWhoIs = (BACnetUnconfirmedServiceRequestWhoIs) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(871, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeLowLimit().getActualValue());
                    assertEquals(871, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeHighLimit().getActualValue());
                })
        );
    }

    @TestFactory
    @DisplayName("CEN_10")
    Collection<DynamicNode> CEN_10() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("CEN_10.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ   confirmedEventNotification[  7] device,151 trend-log,1 trend-log,1 log-buffer device,151",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestConfirmedEventNotification baCnetConfirmedServiceRequestConfirmedEventNotification = (BACnetConfirmedServiceRequestConfirmedEventNotification) serviceRequest;
                    assertEquals((short) 0, baCnetConfirmedServiceRequestConfirmedEventNotification.getProcessIdentifier().getValueUint8());
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestConfirmedEventNotification.getInitiatingDeviceIdentifier().getObjectType());
                    assertEquals(151, baCnetConfirmedServiceRequestConfirmedEventNotification.getInitiatingDeviceIdentifier().getInstanceNumber());
                    assertEquals(BACnetObjectType.TREND_LOG, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventObjectIdentifier().getObjectType());
                    assertEquals(1, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventObjectIdentifier().getInstanceNumber());
                    {
                        BACnetTimeStampDateTime timestamp = (BACnetTimeStampDateTime) baCnetConfirmedServiceRequestConfirmedEventNotification.getTimestamp();
                        assertEquals(2008, timestamp.getDateTimeValue().getDateValue().getYear());
                        assertEquals(5, timestamp.getDateTimeValue().getDateValue().getMonth());
                        assertEquals(2, timestamp.getDateTimeValue().getDateValue().getDayOfMonth());
                        assertEquals(5, timestamp.getDateTimeValue().getDateValue().getDayOfWeek());
                        assertEquals(11, timestamp.getDateTimeValue().getTimeValue().getHour());
                        assertEquals(11, timestamp.getDateTimeValue().getTimeValue().getMinute());
                        assertEquals(30, timestamp.getDateTimeValue().getTimeValue().getSecond());
                        assertEquals(0, timestamp.getDateTimeValue().getTimeValue().getFractional());
                    }
                    {
                        assertEquals(1, baCnetConfirmedServiceRequestConfirmedEventNotification.getNotificationClass().getActualValue());
                    }
                    {
                        assertEquals(15, baCnetConfirmedServiceRequestConfirmedEventNotification.getPriority().getActualValue());
                    }
                    {
                        assertEquals(BACnetEventType.BUFFER_READY, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventType().getValue());
                    }
                    {
                        assertEquals(BACnetNotifyType.EVENT, baCnetConfirmedServiceRequestConfirmedEventNotification.getNotifyType().getValue());
                    }
                    {
                        assertTrue(baCnetConfirmedServiceRequestConfirmedEventNotification.getAckRequired().getIsTrue());
                    }
                    {
                        assertEquals(BACnetEventState.NORMAL, baCnetConfirmedServiceRequestConfirmedEventNotification.getFromState().getValue());
                    }
                    {
                        assertEquals(BACnetEventState.NORMAL, baCnetConfirmedServiceRequestConfirmedEventNotification.getToState().getValue());
                    }
                    {
                        BACnetNotificationParametersBufferReady baCnetNotificationParametersBufferReady = (BACnetNotificationParametersBufferReady) baCnetConfirmedServiceRequestConfirmedEventNotification.getEventValues();
                        assertEquals(BACnetObjectType.TREND_LOG, baCnetNotificationParametersBufferReady.getBufferProperty().getObjectIdentifier().getObjectType());
                        assertEquals(BACnetPropertyIdentifier.LOG_BUFFER, baCnetNotificationParametersBufferReady.getBufferProperty().getPropertyIdentifier().getValue());
                        assertEquals(BACnetObjectType.DEVICE, baCnetNotificationParametersBufferReady.getBufferProperty().getDeviceIdentifier().getObjectType());
                        assertEquals(1640, baCnetNotificationParametersBufferReady.getPreviousNotification().getActualValue());
                        assertEquals(1653, baCnetNotificationParametersBufferReady.getCurrentNotification().getActualValue());
                    }
                })
        );
    }

    @TestFactory
    @DisplayName("COV_AWF_ARF")
    Collection<DynamicNode> COV_AWF_ARF() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("COV_AWF_ARF.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ   subscribeCOV[ 10] binary-input,0",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestSubscribeCOV baCnetConfirmedServiceRequestSubscribeCOV = (BACnetConfirmedServiceRequestSubscribeCOV) serviceRequest;
                    assertEquals((short) 123, baCnetConfirmedServiceRequestSubscribeCOV.getSubscriberProcessIdentifier().getValueUint8());
                    assertEquals(BACnetObjectType.BINARY_INPUT, baCnetConfirmedServiceRequestSubscribeCOV.getMonitoredObjectIdentifier().getObjectType());
                    assertEquals(0, baCnetConfirmedServiceRequestSubscribeCOV.getMonitoredObjectIdentifier().getInstanceNumber());
                    assertTrue(baCnetConfirmedServiceRequestSubscribeCOV.getIssueConfirmed().getIsFalse());
                    assertEquals(10, baCnetConfirmedServiceRequestSubscribeCOV.getLifetimeInSeconds().getActualValue() / 60);
                }),
            DynamicTest.dynamicTest("Simple-ACK      subscribeCOV[ 10]",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUSimpleAck apduSimpleAck = (APDUSimpleAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    assertEquals((short) 5, apduSimpleAck.getServiceChoice());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ unconfirmedCOVNotification device,12345 binary-input,0 present-value status-flags",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequest serviceRequest = apduUnconfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification baCnetUnconfirmedServiceRequestUnconfirmedCOVNotification = (BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification) serviceRequest;
                    assertEquals((short) 123, baCnetUnconfirmedServiceRequestUnconfirmedCOVNotification.getSubscriberProcessIdentifier().getValueUint8());
                    assertEquals(BACnetObjectType.BINARY_INPUT, baCnetUnconfirmedServiceRequestUnconfirmedCOVNotification.getMonitoredObjectIdentifier().getObjectType());
                    assertEquals(0, baCnetUnconfirmedServiceRequestUnconfirmedCOVNotification.getMonitoredObjectIdentifier().getInstanceNumber());
                    assertEquals(9, baCnetUnconfirmedServiceRequestUnconfirmedCOVNotification.getLifetimeInSeconds().getActualValue() / 60);
                    {
                        BACnetContextTagPropertyIdentifier baCnetContextTagPropertyIdentifier = baCnetUnconfirmedServiceRequestUnconfirmedCOVNotification.getListOfValues().getData().get(0).getPropertyIdentifier();
                        assertEquals(BACnetPropertyIdentifier.PRESENT_VALUE, baCnetContextTagPropertyIdentifier.getValue());
                    }
                    {
                        BACnetApplicationTagEnumerated baCnetApplicationTagEnumerated = (BACnetApplicationTagEnumerated) baCnetUnconfirmedServiceRequestUnconfirmedCOVNotification.getListOfValues().getData().get(0).getPropertyValue().getConstructedData().getData().get(0).getApplicationTag();
                        assertEquals(List.of((byte) 0x0), baCnetApplicationTagEnumerated.getData());
                    }
                    {
                        BACnetContextTagPropertyIdentifier baCnetContextTagPropertyIdentifier = baCnetUnconfirmedServiceRequestUnconfirmedCOVNotification.getListOfValues().getData().get(1).getPropertyIdentifier();
                        assertEquals(BACnetPropertyIdentifier.STATUS_FLAGS, baCnetContextTagPropertyIdentifier.getValue());
                    }
                    {
                        BACnetApplicationTagBitString baCnetApplicationTagBitString = (BACnetApplicationTagBitString) baCnetUnconfirmedServiceRequestUnconfirmedCOVNotification.getListOfValues().getData().get(1).getPropertyValue().getConstructedData().getData().get(0).getApplicationTag();
                        assertEquals(Arrays.asList(false, false, false, false), baCnetApplicationTagBitString.getData());
                    }
                }),
            parseEmAll(pcapEvaluator, 3, 1351)
        );
    }

    private DynamicContainer parseEmAll(PCAPEvaluator pcapEvaluator, int startInclusive, int endExclusive) {
        return DynamicContainer.dynamicContainer("Parse em all (from " + startInclusive + " to " + endExclusive + ")", () -> IntStream.range(startInclusive, endExclusive).mapToObj((i) -> DynamicTest.dynamicTest("test n." + i, () -> {
            assumeTrue(i > pcapEvaluator.getCurrentPackageNumber(), "package nr." + i + " already parsed");
            BVLC bvlc = pcapEvaluator.nextBVLC();
            assumeTrue(bvlc != null, "No more package left");
            dump(bvlc);
        })).map(DynamicNode.class::cast).iterator());
    }

    @TestFactory
    @DisplayName("ContextTagAbove14Sample_1")
    Collection<DynamicNode> ContextTagAbove14Sample_1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("ContextTagAbove14Sample_1.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("skipLLC",
                () -> {
                    pcapEvaluator.skipPackages(1);
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   confirmedEventNotification[138] device,1 event-enrollment,1",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestConfirmedEventNotification baCnetConfirmedServiceRequestConfirmedEventNotification = (BACnetConfirmedServiceRequestConfirmedEventNotification) serviceRequest;
                    assertEquals((short) 1, baCnetConfirmedServiceRequestConfirmedEventNotification.getProcessIdentifier().getValueUint8());
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestConfirmedEventNotification.getInitiatingDeviceIdentifier().getObjectType());
                    assertEquals(1, baCnetConfirmedServiceRequestConfirmedEventNotification.getInitiatingDeviceIdentifier().getInstanceNumber());
                    assertEquals(BACnetObjectType.EVENT_ENROLLMENT, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventObjectIdentifier().getObjectType());
                    assertEquals(1, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventObjectIdentifier().getInstanceNumber());
                    {
                        BACnetTimeStampSequence timestamp = (BACnetTimeStampSequence) baCnetConfirmedServiceRequestConfirmedEventNotification.getTimestamp();
                        assertEquals(1, timestamp.getSequenceNumber().getActualValue());
                    }
                    {
                        assertEquals(1, baCnetConfirmedServiceRequestConfirmedEventNotification.getNotificationClass().getActualValue());
                    }
                    {
                        assertEquals(111, baCnetConfirmedServiceRequestConfirmedEventNotification.getPriority().getActualValue());
                    }
                    {
                        assertEquals(BACnetEventType.CHANGE_OF_STATE, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventType().getValue());
                    }
                    {
                        assertEquals(BACnetNotifyType.EVENT, baCnetConfirmedServiceRequestConfirmedEventNotification.getNotifyType().getValue());
                    }
                    {
                        assertTrue(baCnetConfirmedServiceRequestConfirmedEventNotification.getAckRequired().getIsFalse());
                    }
                    {
                        assertEquals(BACnetEventState.NORMAL, baCnetConfirmedServiceRequestConfirmedEventNotification.getFromState().getValue());
                    }
                    {
                        assertEquals(BACnetEventState.NORMAL, baCnetConfirmedServiceRequestConfirmedEventNotification.getToState().getValue());
                    }
                    {
                        BACnetNotificationParametersChangeOfState baCnetNotificationParametersChangeOfState = (BACnetNotificationParametersChangeOfState) baCnetConfirmedServiceRequestConfirmedEventNotification.getEventValues();
                        assertEquals(true, baCnetNotificationParametersChangeOfState.getStatusFlags().getInAlarm());
                        assertEquals(false, baCnetNotificationParametersChangeOfState.getStatusFlags().getFault());
                        assertEquals(false, baCnetNotificationParametersChangeOfState.getStatusFlags().getOverriden());
                        assertEquals(false, baCnetNotificationParametersChangeOfState.getStatusFlags().getOutOfService());
                    }
                })
        );
    }

    @TestFactory
    @DisplayName("CriticalRoom55-1")
    Collection<DynamicNode> CriticalRoom55_1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("CriticalRoom55-1.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("CriticalRoom55-2")
    Collection<DynamicNode> CriticalRoom55_2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("CriticalRoom55-2.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ writeProperty[113] analog-value,1 present-value",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestWriteProperty baCnetConfirmedServiceRequestWriteProperty = (BACnetConfirmedServiceRequestWriteProperty) serviceRequest;
                    BACnetApplicationTagReal baCnetApplicationTagReal = (BACnetApplicationTagReal) baCnetConfirmedServiceRequestWriteProperty.getPropertyValue().getData().get(0).getApplicationTag();
                    assertEquals(123.0f, baCnetApplicationTagReal.getValue());
                }),
            DynamicTest.dynamicTest("Abort",
                () -> {
                    // TODO: package is malformed
                    assumeTrue(false);
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("DRI%20CAVE%20log%20udp-0168-20081216-1117-03")
    Collection<DynamicNode> DRI_CAVE_log_udp_0168_20081216_1117_03() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("DRI%20CAVE%20log%20udp-0168-20081216-1117-03.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("I-Am-Router-To-Network")
    Collection<DynamicNode> I_Am_Router_To_Network() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("I-Am-Router-To-Network.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("Ethereal-Misinterpreted-Packet")
    Collection<DynamicNode> Ethereal_Misinterpreted_Packet() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("Ethereal-Misinterpreted-Packet.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ   confirmedEventNotification[ 10] device,1041000 analog-input,3000016 (2200) Vendor Proprietary Value object-name (2201) Vendor Proprietary Value (2202) Vendor Proprietary Value reliability (661) VendorProprietary Value units (1659) Vendor Proprietary Value (2203) Vendor Proprietary Value vendor-identifier",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestConfirmedEventNotification baCnetConfirmedServiceRequestConfirmedEventNotification = (BACnetConfirmedServiceRequestConfirmedEventNotification) serviceRequest;
                    assertEquals((short) 0, baCnetConfirmedServiceRequestConfirmedEventNotification.getProcessIdentifier().getValueUint8());
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestConfirmedEventNotification.getInitiatingDeviceIdentifier().getObjectType());
                    assertEquals(1041000, baCnetConfirmedServiceRequestConfirmedEventNotification.getInitiatingDeviceIdentifier().getInstanceNumber());
                    assertEquals(BACnetObjectType.ANALOG_INPUT, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventObjectIdentifier().getObjectType());
                    assertEquals(3000016, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventObjectIdentifier().getInstanceNumber());
                    {
                        BACnetTimeStampDateTime timestamp = (BACnetTimeStampDateTime) baCnetConfirmedServiceRequestConfirmedEventNotification.getTimestamp();
                        assertEquals(2005, timestamp.getDateTimeValue().getDateValue().getYear());
                        assertEquals(12, timestamp.getDateTimeValue().getDateValue().getMonth());
                        assertEquals(8, timestamp.getDateTimeValue().getDateValue().getDayOfMonth());
                        assertEquals(4, timestamp.getDateTimeValue().getDateValue().getDayOfWeek());
                        assertEquals(14, timestamp.getDateTimeValue().getTimeValue().getHour());
                        assertEquals(12, timestamp.getDateTimeValue().getTimeValue().getMinute());
                        assertEquals(49, timestamp.getDateTimeValue().getTimeValue().getSecond());
                        assertEquals(0, timestamp.getDateTimeValue().getTimeValue().getFractional());
                    }
                    {
                        assertEquals(1, baCnetConfirmedServiceRequestConfirmedEventNotification.getNotificationClass().getActualValue());
                    }
                    {
                        assertEquals(200, baCnetConfirmedServiceRequestConfirmedEventNotification.getPriority().getActualValue());
                    }
                    {
                        assertEquals(BACnetEventType.OUT_OF_RANGE, baCnetConfirmedServiceRequestConfirmedEventNotification.getEventType().getValue());
                    }
                    {
                        assertEquals(BACnetNotifyType.ALARM, baCnetConfirmedServiceRequestConfirmedEventNotification.getNotifyType().getValue());
                    }
                    {
                        assertTrue(baCnetConfirmedServiceRequestConfirmedEventNotification.getAckRequired().getIsTrue());
                    }
                    {
                        assertEquals(BACnetEventState.HIGH_LIMIT, baCnetConfirmedServiceRequestConfirmedEventNotification.getFromState().getValue());
                    }
                    {
                        assertEquals(BACnetEventState.NORMAL, baCnetConfirmedServiceRequestConfirmedEventNotification.getToState().getValue());
                    }
                    {
                        BACnetNotificationParametersComplexEventType baCnetNotificationParametersComplexEventType = (BACnetNotificationParametersComplexEventType) baCnetConfirmedServiceRequestConfirmedEventNotification.getEventValues();
                        {
                            BACnetPropertyValue baCnetPropertyValue = baCnetNotificationParametersComplexEventType.getListOfValues().getData().get(0);
                            assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetPropertyValue.getPropertyIdentifier().getValue());
                            assertEquals(2200, baCnetPropertyValue.getPropertyIdentifier().getProprietaryValue());
                            BACnetApplicationTagCharacterString baCnetApplicationTagCharacterString = (BACnetApplicationTagCharacterString) baCnetPropertyValue.getPropertyValue().getConstructedData().getData().get(0).getApplicationTag();
                            assertEquals("StockingNAE", baCnetApplicationTagCharacterString.getValue());
                        }
                        {
                            BACnetPropertyValue baCnetPropertyValue = baCnetNotificationParametersComplexEventType.getListOfValues().getData().get(1);
                            assertEquals(BACnetPropertyIdentifier.OBJECT_NAME, baCnetPropertyValue.getPropertyIdentifier().getValue());
                            BACnetApplicationTagCharacterString baCnetApplicationTagCharacterString = (BACnetApplicationTagCharacterString) baCnetPropertyValue.getPropertyValue().getConstructedData().getData().get(0).getApplicationTag();
                            assertEquals("StockingNAE/N2-1.NAE4-N2A-DX1.OA-T", baCnetApplicationTagCharacterString.getValue());
                        }
                        {
                            BACnetPropertyValue baCnetPropertyValue = baCnetNotificationParametersComplexEventType.getListOfValues().getData().get(2);
                            assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetPropertyValue.getPropertyIdentifier().getValue());
                            assertEquals(2201, baCnetPropertyValue.getPropertyIdentifier().getProprietaryValue());
                            BACnetApplicationTagEnumerated baCnetApplicationTagEnumerated = (BACnetApplicationTagEnumerated) baCnetPropertyValue.getPropertyValue().getConstructedData().getData().get(0).getApplicationTag();
                            assertEquals(List.of((byte) 85), baCnetApplicationTagEnumerated.getData());
                        }
                        {
                            BACnetPropertyValue baCnetPropertyValue = baCnetNotificationParametersComplexEventType.getListOfValues().getData().get(3);
                            assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetPropertyValue.getPropertyIdentifier().getValue());
                            assertEquals(2202, baCnetPropertyValue.getPropertyIdentifier().getProprietaryValue());
                            BACnetApplicationTagReal baCnetApplicationTagReal = (BACnetApplicationTagReal) baCnetPropertyValue.getPropertyValue().getConstructedData().getData().get(0).getApplicationTag();
                            assertEquals(35.093750, baCnetApplicationTagReal.getValue());
                        }
                        {
                            BACnetPropertyValue baCnetPropertyValue = baCnetNotificationParametersComplexEventType.getListOfValues().getData().get(4);
                            assertEquals(BACnetPropertyIdentifier.RELIABILITY, baCnetPropertyValue.getPropertyIdentifier().getValue());
                            BACnetApplicationTagEnumerated baCnetApplicationTagEnumerated = (BACnetApplicationTagEnumerated) baCnetPropertyValue.getPropertyValue().getConstructedData().getData().get(0).getApplicationTag();
                            assertEquals(List.of((byte) 0), baCnetApplicationTagEnumerated.getData());
                        }
                        {
                            BACnetPropertyValue baCnetPropertyValue = baCnetNotificationParametersComplexEventType.getListOfValues().getData().get(5);
                            assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetPropertyValue.getPropertyIdentifier().getValue());
                            assertEquals(661, baCnetPropertyValue.getPropertyIdentifier().getProprietaryValue());
                            BACnetApplicationTagEnumerated baCnetApplicationTagEnumerated = (BACnetApplicationTagEnumerated) baCnetPropertyValue.getPropertyValue().getConstructedData().getData().get(0).getApplicationTag();
                            assertEquals(List.of((byte) 5), baCnetApplicationTagEnumerated.getData());
                        }
                        {
                            BACnetPropertyValue baCnetPropertyValue = baCnetNotificationParametersComplexEventType.getListOfValues().getData().get(6);
                            assertEquals(BACnetPropertyIdentifier.UNITS, baCnetPropertyValue.getPropertyIdentifier().getValue());
                            BACnetApplicationTagEnumerated baCnetApplicationTagEnumerated = (BACnetApplicationTagEnumerated) baCnetPropertyValue.getPropertyValue().getConstructedData().getData().get(0).getApplicationTag();
                            assertEquals(List.of((byte) 64), baCnetApplicationTagEnumerated.getData());
                        }
                        {
                            BACnetPropertyValue baCnetPropertyValue = baCnetNotificationParametersComplexEventType.getListOfValues().getData().get(7);
                            assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetPropertyValue.getPropertyIdentifier().getValue());
                            assertEquals(1659, baCnetPropertyValue.getPropertyIdentifier().getProprietaryValue());
                            BACnetApplicationTagEnumerated baCnetApplicationTagEnumerated = (BACnetApplicationTagEnumerated) baCnetPropertyValue.getPropertyValue().getConstructedData().getData().get(0).getApplicationTag();
                            assertEquals(List.of((byte) 0), baCnetApplicationTagEnumerated.getData());
                        }
                        {
                            BACnetPropertyValue baCnetPropertyValue = baCnetNotificationParametersComplexEventType.getListOfValues().getData().get(8);
                            assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetPropertyValue.getPropertyIdentifier().getValue());
                            assertEquals(2203, baCnetPropertyValue.getPropertyIdentifier().getProprietaryValue());
                            BACnetApplicationTagEnumerated baCnetApplicationTagEnumerated = (BACnetApplicationTagEnumerated) baCnetPropertyValue.getPropertyValue().getConstructedData().getData().get(0).getApplicationTag();
                            assertEquals(List.of((byte) 0), baCnetApplicationTagEnumerated.getData());
                        }
                        {
                            BACnetPropertyValue baCnetPropertyValue = baCnetNotificationParametersComplexEventType.getListOfValues().getData().get(9);
                            assertEquals(BACnetPropertyIdentifier.VENDOR_IDENTIFIER, baCnetPropertyValue.getPropertyIdentifier().getValue());
                            BACnetApplicationTagUnsignedInteger baCnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) baCnetPropertyValue.getPropertyValue().getConstructedData().getData().get(0).getApplicationTag();
                            assertEquals(5, baCnetApplicationTagUnsignedInteger.getActualValue());
                        }
                    }
                })
        );
    }

    @TestFactory
    @DisplayName("MSTP_Malformed_Packets")
    Collection<DynamicNode> MSTP_Malformed_Packets() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("MSTP_Malformed_Packets.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("NPDU")
    Collection<DynamicNode> NPDU() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("NPDU.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("PrivateTransferError-octetstring")
    Collection<DynamicNode> PrivateTransferError_octetstring() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("PrivateTransferError-octetstring.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("PrivateTransferError")
    Collection<DynamicNode> PrivateTransferError() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("PrivateTransferError.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("RPM_ALL_Allobjecttypes1")
    Collection<DynamicNode> RPM_ALL_Allobjecttypes1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("RPM_ALL_Allobjecttypes1.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("ReadPropertyMultiple")
    Collection<DynamicNode> ReadPropertyMultiple() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("ReadPropertyMultiple.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("ReadPropertyMultipleDeviceAll")
    Collection<DynamicNode> ReadPropertyMultipleDeviceAll() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("ReadPropertyMultipleDeviceAll.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("Subordinate List")
    Collection<DynamicNode> Subordinate_List() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("Subordinate%20List.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ readProperty[152] structured-view,1 subordinate-list",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[152] structured-view,1 subordinate-list device,128 analog-input,1 device,128 analog-input,3 device,128 analog-output,1 device,128 analog-output,3",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("Subordinate List2")
    Collection<DynamicNode> Subordinate_List2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("Subordinate%20List2.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[143] structured-view,1 subordinate-list",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[143] structured-view,1 subordinate-list device,4000 analog-input,1 analog-value,1 binary-input,1 binary-value,1",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[144] structured-view,1 subordinate-annotations",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[144] structured-view,1 subordinate-annotations",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("SubordinateListDecodeSample")
    Collection<DynamicNode> SubordinateListDecodeSample() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("SubordinateListDecodeSample.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("SynergyBlinkWarn")
    Collection<DynamicNode> SynergyBlinkWarn() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("SynergyBlinkWarn.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("SynergyReadProperties")
    Collection<DynamicNode> SynergyReadProperties() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("SynergyReadProperties.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("SynergyWriteProperty")
    Collection<DynamicNode> SynergyWriteProperty() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("SynergyWriteProperty.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("Sysco-1")
    Collection<DynamicNode> Sysco_1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("Sysco-1.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("Sysco-2")
    Collection<DynamicNode> Sysco_2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("Sysco-2.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("Sysco-3")
    Collection<DynamicNode> Sysco_3() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("Sysco-3.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TA02 MST")
    Collection<DynamicNode> TA02_MST() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TA02%20MST.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TAO2 TES.39A")
    Collection<DynamicNode> TAO2_TES_39_A() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TAO2%20TES.39A.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TC51103_BTL-9.21.1.X3_bool_ext_3")
    Collection<DynamicNode> TC51103_BTL_9_21_1_X3_bool_ext_3() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TC51103_BTL-9.21.1.X3_bool_ext_3.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TC51103_BTL-9.21.1.X3_int_ext_1")
    Collection<DynamicNode> TC51103_BTL_9_21_1_X3_int_ext_1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TC51103_BTL-9.21.1.X3_int_ext_1.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TestRun4 - Internal side of Router")
    Collection<DynamicNode> TestRun4___Internal_side_of_Router() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TestRun4%20-%20Internal%20side%20of%20Router.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TestRun4 - Outside of Router")
    Collection<DynamicNode> TestRun4___Outside_of_Router() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TestRun4%20-%20Outside%20of%20Router.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TestRun5 - Internal side of Router")
    Collection<DynamicNode> TestRun5___Internal_side_of_Router() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TestRun5%20-%20Internal%20side%20of%20Router.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TestRun5 - Outside of Router")
    Collection<DynamicNode> TestRun5___Outside_of_Router() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TestRun5%20-%20Outside%20of%20Router.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TestRun8 - Internal side of Router")
    Collection<DynamicNode> TestRun8___Internal_side_of_Router() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TestRun8%20-%20Internal%20side%20of%20Router.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TestRun8 - Outside of Router")
    Collection<DynamicNode> TestRun8___Outside_of_Router() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TestRun8%20-%20Outside%20of%20Router.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TimeSync_Decode_Noon")
    Collection<DynamicNode> TimeSync_Decode_Noon() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TimeSync_Decode_Noon.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("Tower333 lighting 5min IP")
    Collection<DynamicNode> Tower333_lighting_5min_IP() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("Tower333%20lighting%205min%20IP.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TrendLogMultipleReadRange")
    Collection<DynamicNode> TrendLogMultipleReadRange() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TrendLogMultipleReadRange.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TrendLogMultipleReadRange2")
    Collection<DynamicNode> TrendLogMultipleReadRange2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TrendLogMultipleReadRange2.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TrendLogMultipleReadRange3")
    Collection<DynamicNode> TrendLogMultipleReadRange3() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TrendLogMultipleReadRange3.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TrendLogMultipleReadRangeSimple")
    Collection<DynamicNode> TrendLogMultipleReadRangeSimple() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TrendLogMultipleReadRangeSimple.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("TrendLogMultipleUsage")
    Collection<DynamicNode> TrendLogMultipleUsage() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("TrendLogMultipleUsage.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("WhoIsRouterToNetwork-test")
    Collection<DynamicNode> WhoIsRouterToNetwork_test() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("WhoIsRouterToNetwork-test.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("WhoIsRouterToNetwork")
    Collection<DynamicNode> WhoIsRouterToNetwork() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("WhoIsRouterToNetwork.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("WhoIs_I-Am_Epics")
    Collection<DynamicNode> WhoIs_I_Am_Epics() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("WhoIs_I-Am_Epics.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("WireSharkError_ArrayIndex")
    Collection<DynamicNode> WireSharkError_ArrayIndex() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("WireSharkError_ArrayIndex.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 74] schedule,1 exception-schedule",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 74] schedule,1 exception-schedule",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("WireSharkError_BufferReadyNotification")
    Collection<DynamicNode> WireSharkError_BufferReadyNotification() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("WireSharkError_BufferReadyNotification.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("WireSharkOfNewObjects")
    Collection<DynamicNode> WireSharkOfNewObjects() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("WireSharkOfNewObjects.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("WriteProperty_BinaryOutput")
    Collection<DynamicNode> WriteProperty_BinaryOutput() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("WriteProperty_BinaryOutput.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("action-list")
    Collection<DynamicNode> action_list() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("action-list.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[107] command,1 action",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[107] command,1 action binary-value,0 present-value device,0 analog-value,1 present-value binary-value,2 present-value device,0 analog-value,3 present-value binary-value,4 present-value device,0 analog-value,5 present-value binary-value,6 present-value device,0 analog-value,7 present-value binary-value,8 present-value device,0 analog-value,9 present-value", () -> {
                BVLC bvlc;
                bvlc = pcapEvaluator.nextBVLC();
                dump(bvlc);
                // TODO:
                assumeTrue(false, "not properly implemented. Check manually and add asserts");
            })
        );
    }

    @TestFactory
    @DisplayName("aha_220_to_20_lost_b")
    Collection<DynamicNode> aha_220_to_20_lost_b() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("aha_220_to_20_lost_b.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("alerton-plugfest-2")
    Collection<DynamicNode> alerton_plugfest_2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("alerton-plugfest-2.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Complex-ACK readProperty[155] device,42222 protocol-version",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);

                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAckReadProperty);
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_VERSION, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Complex-ACK readProperty[155] device,42222 protocol-conformance-class",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("alerton-plugfest-3")
    Collection<DynamicNode> alerton_plugfest_3() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("alerton-plugfest-3.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("alerton-plugfest")
    Collection<DynamicNode> alerton_plugfest() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("alerton-plugfest.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("arf-empty-file")
    Collection<DynamicNode> arf_empty_file() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("arf-empty-file.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic-empty-file")
    Collection<DynamicNode> atomic_empty_file() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic-empty-file.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic-read-file-50")
    Collection<DynamicNode> atomic_read_file_50() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic-read-file-50.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic-read-file-50x1500k")
    Collection<DynamicNode> atomic_read_file_50x1500k() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic-read-file-50x1500k.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic-read-file-480")
    Collection<DynamicNode> atomic_read_file_480() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic-read-file-480.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic-read-file-1470")
    Collection<DynamicNode> atomic_read_file_1470() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic-read-file-1470.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic-read-file")
    Collection<DynamicNode> atomic_read_file() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic-read-file.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic-write-file-2")
    Collection<DynamicNode> atomic_write_file_2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic-write-file-2.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic-write-file-3")
    Collection<DynamicNode> atomic_write_file_3() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic-write-file-3.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic-write-file-50x1k")
    Collection<DynamicNode> atomic_write_file_50x1k() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic-write-file-50x1k.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic-write-file-480")
    Collection<DynamicNode> atomic_write_file_480() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic-write-file-480.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic-write-file-seg")
    Collection<DynamicNode> atomic_write_file_seg() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic-write-file-seg.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic-write-file")
    Collection<DynamicNode> atomic_write_file() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic-write-file.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("atomic_write_file_bad_ack")
    Collection<DynamicNode> atomic_write_file_bad_ack() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("atomic_write_file_bad_ack.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bacapp-malform")
    Collection<DynamicNode> bacapp_malform() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bacapp-malform.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bacnet-arcnet")
    Collection<DynamicNode> bacnet_arcnet() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bacnet-arcnet.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bacnet-ethernet-device")
    Collection<DynamicNode> bacnet_ethernet_device() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bacnet-ethernet-device.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bacnet-ethernet")
    Collection<DynamicNode> bacnet_ethernet() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bacnet-ethernet.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bacnet-ip")
    Collection<DynamicNode> bacnet_ip() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bacnet-ip.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bacnet-properties")
    Collection<DynamicNode> bacnet_properties() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bacnet-properties.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bacnet-services")
    Collection<DynamicNode> bacnet_services() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bacnet-services.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bacnet-stack-services")
    Collection<DynamicNode> bacnet_stack_services() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bacnet-stack-services.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bacrpm-test")
    Collection<DynamicNode> bacrpm_test() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bacrpm-test.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bad_hub_restart")
    Collection<DynamicNode> bad_hub_restart() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bad_hub_restart.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bip-discover")
    Collection<DynamicNode> bip_discover() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bip-discover.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bip-readprop-2")
    Collection<DynamicNode> bip_readprop_2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bip-readprop-2.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bip-readprop")
    Collection<DynamicNode> bip_readprop() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bip-readprop.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bip-readwrite-test")
    Collection<DynamicNode> bip_readwrite_test() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bip-readwrite-test.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bo_command_failure")
    Collection<DynamicNode> bo_command_failure() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bo_command_failure.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bo_command_failure_original")
    Collection<DynamicNode> bo_command_failure_original() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bo_command_failure_original.pcap");
        // TODO: we should set a filter for bacnet
        // Pcap starts with 20 non bacnet packages
        pcapEvaluator.skipPackages(20);
        return Arrays.asList(
            DynamicTest.dynamicTest("Unconfirmed-REQ unconfirmedCOVNotification device,1 accumulator,21 present-value status-flags",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ unconfirmedCOVNotification device,1 accumulator,22 present-value status-flags",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ unconfirmedCOVNotification device,1 binary-input,217 present-value status-flags",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ unconfirmedCOVNotification device,1 accumulator,21 present-value status-flags",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ unconfirmedCOVNotification device,1 binary-input,217 present-value status-flags",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ unconfirmedCOVNotification device,1 binary-output,1 present-value status-flags",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ unconfirmedEventNotification device,1 binary-output,1",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ unconfirmedCOVNotification device,1 accumulator,22 present-value status-flags",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("btl-plugfest")
    Collection<DynamicNode> btl_plugfest() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("btl-plugfest.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bvlc-bac4-rp")
    Collection<DynamicNode> bvlc_bac4_rp() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bvlc-bac4-rp.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bvlc-bac4")
    Collection<DynamicNode> bvlc_bac4() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bvlc-bac4.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bvlc-fdreg-readprop-47809")
    Collection<DynamicNode> bvlc_fdreg_readprop_47809() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bvlc-fdreg-readprop-47809.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bvlc-loop")
    Collection<DynamicNode> bvlc_loop() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bvlc-loop.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("bvlc")
    Collection<DynamicNode>
    bvlc() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("bvlc.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("cimetrics_mstp")
    Collection<DynamicNode> cimetrics_mstp() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("cimetrics_mstp.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("config-tool-discover")
    Collection<DynamicNode> config_tool_discover() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("config-tool-discover.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("confirmedEventNotification")
    Collection<DynamicNode> confirmedEventNotification() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("confirmedEventNotification.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ confirmedEventNotification[103] device,1041000 analog-input,3000016 present-value",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("cov-testing-1")
    Collection<DynamicNode> cov_testing_1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("cov-testing-1.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("cov-testing-2")
    Collection<DynamicNode> cov_testing_2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("cov-testing-2.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("cov-testing-3")
    Collection<DynamicNode> cov_testing_3() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("cov-testing-3.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("device-address-binding")
    Collection<DynamicNode> device_address_binding() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("device-address-binding.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("epics-1")
    Collection<DynamicNode> epics_1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("epics-1.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("epics-2")
    Collection<DynamicNode> epics_2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("epics-2.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("eventLog_ReadRange")
    Collection<DynamicNode> eventLog_ReadRange() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("eventLog_ReadRange.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("eventLog_rpm")
    Collection<DynamicNode> eventLog_rpm() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("eventLog_rpm.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("eventTimeStamp_rp")
    Collection<DynamicNode> eventTimeStamp_rp() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("eventTimeStamp_rp.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[148] load-control,1 event-time-stamp",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[148] load-control,1 event-time-stamp",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("eventTimeStamp_rpm")
    Collection<DynamicNode> eventTimeStamp_rpm() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("eventTimeStamp_rpm.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @Disabled("Needs filtering")
    @TestFactory
    @DisplayName("foreign-device-npdu")
    Collection<DynamicNode> foreign_device_npdu() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("foreign-device-npdu.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("getEventInformation")
    Collection<DynamicNode> getEventInformation() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("getEventInformation.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("i-am-vendor-id-over-255")
    Collection<DynamicNode> i_am_vendor_id_over_255() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("i-am-vendor-id-over-255.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("lmbc-300-bootload")
    Collection<DynamicNode> lmbc_300_bootload() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("lmbc-300-bootload.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("load-control-properties")
    Collection<DynamicNode> load_control_properties() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("load-control-properties.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("load-control")
    Collection<DynamicNode> load_control() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("load-control.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("log-buffer_readRange")
    Collection<DynamicNode> log_buffer_readRange() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("log-buffer_readRange.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("loop2")
    Collection<DynamicNode> loop2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("loop2.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp-cimetrics")
    Collection<DynamicNode> mstp_cimetrics() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp-cimetrics.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp-test-4")
    Collection<DynamicNode> mstp_test_4() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp-test-4.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp-whois-basrt-mix")
    Collection<DynamicNode> mstp_whois_basrt_mix() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp-whois-basrt-mix.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp-whois-basrt-mix2")
    Collection<DynamicNode> mstp_whois_basrt_mix2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp-whois-basrt-mix2.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp-whois-iam")
    Collection<DynamicNode> mstp_whois_iam() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp-whois-iam.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20090227094623")
    Collection<DynamicNode> mstp_20090227094623() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20090227094623.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20090304105820")
    Collection<DynamicNode> mstp_20090304105820() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20090304105820.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20090304110410")
    Collection<DynamicNode> mstp_20090304110410() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20090304110410.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20090807145500")
    Collection<DynamicNode> mstp_20090807145500() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20090807145500.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20091013121352")
    Collection<DynamicNode> mstp_20091013121352() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20091013121352.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20091013121410")
    Collection<DynamicNode> mstp_20091013121410() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20091013121410.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20091013122053")
    Collection<DynamicNode> mstp_20091013122053() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20091013122053.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20091013122451")
    Collection<DynamicNode> mstp_20091013122451() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20091013122451.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20091013123108")
    Collection<DynamicNode> mstp_20091013123108() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20091013123108.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20091013124218")
    Collection<DynamicNode> mstp_20091013124218() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20091013124218.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20091013130259")
    Collection<DynamicNode> mstp_20091013130259() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20091013130259.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20091013162906")
    Collection<DynamicNode> mstp_20091013162906() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20091013162906.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20091014093519")
    Collection<DynamicNode> mstp_20091014093519() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20091014093519.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20091014112427")
    Collection<DynamicNode> mstp_20091014112427() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20091014112427.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_20140225214217")
    Collection<DynamicNode> mstp_20140225214217() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_20140225214217.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_mix_basrt_V124")
    Collection<DynamicNode> mstp_mix_basrt_V124() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_mix_basrt_V124.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_mix_basrt_V124_bad")
    Collection<DynamicNode> mstp_mix_basrt_V124_bad() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_mix_basrt_V124_bad.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("mstp_wtap")
    Collection<DynamicNode> mstp_wtap() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("mstp_wtap.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("nb-binary-output")
    Collection<DynamicNode> nb_binary_output() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("nb-binary-output.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("plugfest-2011-delta-1")
    Collection<DynamicNode> plugfest_2011_delta_1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("plugfest-2011-delta-1.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("plugfest-2011-delta-2")
    Collection<DynamicNode> plugfest_2011_delta_2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("plugfest-2011-delta-2.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("plugfest-2011-delta-3")
    Collection<DynamicNode> plugfest_2011_delta_3() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("plugfest-2011-delta-3.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("plugfest-2011-mstp-roundtable")
    Collection<DynamicNode> plugfest_2011_mstp_roundtable() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("plugfest-2011-mstp-roundtable.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("plugfest-2011-sauter-1")
    Collection<DynamicNode> plugfest_2011_sauter_1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("plugfest-2011-sauter-1.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("plugfest-2011-siemens-1")
    Collection<DynamicNode> plugfest_2011_siemens_1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("plugfest-2011-siemens-1.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("plugfest-2011-trane-1")
    Collection<DynamicNode> plugfest_2011_trane_1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("plugfest-2011-trane-1.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Unconfirmed REQ who-Is",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequest serviceRequest = apduUnconfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    assertTrue(serviceRequest instanceof BACnetUnconfirmedServiceRequestWhoIs);
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("plugfest-delta-2")
    Collection<DynamicNode> plugfest_delta_2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("plugfest-delta-2.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("plugfest-delta-2b")
    Collection<DynamicNode> plugfest_delta_2b() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("plugfest-delta-2b.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("plugfest-tridium-1")
    Collection<DynamicNode> plugfest_tridium_1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("plugfest-tridium-1.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("plugfest-tridium-2")
    Collection<DynamicNode> plugfest_tridium_2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("plugfest-tridium-2.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("polarsoft-free-range-router-init-routing-table")
    Collection<DynamicNode> polarsoft_free_range_router_init_routing_table() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("polarsoft-free-range-router-init-routing-table.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("polarsoft-free-range-router")
    Collection<DynamicNode> polarsoft_free_range_router() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("polarsoft-free-range-router.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("properties")
    Collection<DynamicNode> properties() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("properties.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("read-file")
    Collection<DynamicNode> read_file() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("read-file.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("read-properties")
    Collection<DynamicNode> read_properties() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("read-properties.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,111",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestIAm baCnetUnconfirmedServiceRequestIAm = (BACnetUnconfirmedServiceRequestIAm) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIAm.getDeviceIdentifier().getObjectType());
                    assertEquals(111, baCnetUnconfirmedServiceRequestIAm.getDeviceIdentifier().getInstanceNumber());
                    assertEquals(50, baCnetUnconfirmedServiceRequestIAm.getMaximumApduLengthAcceptedLength().getActualValue());
                    assertEquals(List.of((byte) 0x03), baCnetUnconfirmedServiceRequestIAm.getSegmentationSupported().getData());
                    assertEquals(42, baCnetUnconfirmedServiceRequestIAm.getVendorId().getActualValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Is",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestWhoIs baCnetUnconfirmedServiceRequestWhoIs = (BACnetUnconfirmedServiceRequestWhoIs) apduUnconfirmedRequest.getServiceRequest();
                }),
            DynamicTest.dynamicTest("skip 5 packages",
                () -> {
                    pcapEvaluator.skipPackages(5);
                }
            ),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestIAm baCnetUnconfirmedServiceRequestIAm = (BACnetUnconfirmedServiceRequestIAm) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIAm.getDeviceIdentifier().getObjectType());
                    assertEquals(201, baCnetUnconfirmedServiceRequestIAm.getDeviceIdentifier().getInstanceNumber());
                    assertEquals(1476, baCnetUnconfirmedServiceRequestIAm.getMaximumApduLengthAcceptedLength().getActualValue());
                    assertEquals(List.of((byte) 0x00), baCnetUnconfirmedServiceRequestIAm.getSegmentationSupported().getData());
                    assertEquals(18, baCnetUnconfirmedServiceRequestIAm.getVendorId().getActualValue());
                }),
            DynamicTest.dynamicTest("skip 1 packages",
                () -> {
                    pcapEvaluator.skipPackages(1);
                }
            ),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Am device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestIAm baCnetUnconfirmedServiceRequestIAm = (BACnetUnconfirmedServiceRequestIAm) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIAm.getDeviceIdentifier().getObjectType());
                    assertEquals(61, baCnetUnconfirmedServiceRequestIAm.getDeviceIdentifier().getInstanceNumber());
                    assertEquals(480, baCnetUnconfirmedServiceRequestIAm.getMaximumApduLengthAcceptedLength().getActualValue());
                    assertEquals(List.of((byte) 0x00), baCnetUnconfirmedServiceRequestIAm.getSegmentationSupported().getData());
                    assertEquals(42, baCnetUnconfirmedServiceRequestIAm.getVendorId().getActualValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 29] device,201 object-identifier",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_IDENTIFIER, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 29] device,201 object-identifier device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_IDENTIFIER, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 30] device,201 object-name",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_NAME, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 30] device,201 object-name device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_NAME, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagCharacterString baCnetApplicationTagCharacterString = (BACnetApplicationTagCharacterString) value;
                    assertEquals("Lithonia Router", baCnetApplicationTagCharacterString.getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 31] device,201 object-type",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_TYPE, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 31] device,201 object-type device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_TYPE, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 32] device,201 system-status",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.SYSTEM_STATUS, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 32] device,201 system-status device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.SYSTEM_STATUS, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagEnumerated baCnetApplicationTagEnumerated = (BACnetApplicationTagEnumerated) value;
                    assertEquals(List.of((byte) 0x0), baCnetApplicationTagEnumerated.getData());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 33] device,201 vendor-name",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_NAME, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 33] device,201 vendor-name device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_NAME, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagCharacterString BACnetApplicationTagCharacterString = (BACnetApplicationTagCharacterString) value;
                    assertEquals("Alerton Technologies, Inc.", BACnetApplicationTagCharacterString.getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 34] device,201 vendor-identifier",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_IDENTIFIER, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 34] device,201 vendor-identifier device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_IDENTIFIER, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 35] device,201 model-name",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MODEL_NAME, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 35] device,201 model-name device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MODEL_NAME, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagCharacterString BACnetApplicationTagCharacterString = (BACnetApplicationTagCharacterString) value;
                    assertEquals("LSi Controller", BACnetApplicationTagCharacterString.getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 36] device,201 model-name",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.FIRMWARE_REVISION, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 36] device,201 model-name device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.FIRMWARE_REVISION, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagCharacterString BACnetApplicationTagCharacterString = (BACnetApplicationTagCharacterString) value;
                    assertEquals("BACtalk LSi   v3.10 A         ", BACnetApplicationTagCharacterString.getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 37] device,201 application-software-version",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.APPLICATION_SOFTWARE_VERSION, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 37] device,201 application-software-version device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.APPLICATION_SOFTWARE_VERSION, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagCharacterString BACnetApplicationTagCharacterString = (BACnetApplicationTagCharacterString) value;
                    assertEquals("LSi Controller v3.11 E", BACnetApplicationTagCharacterString.getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 38] device,201 protocol-version",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_VERSION, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 38] device,201 protocol-version device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_VERSION, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals(1, BACnetApplicationTagUnsignedInteger.getActualValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 39] device,201 protocol-conformance-class",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_CONFORMANCE_CLASS, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 39] device,201 protocol-conformance-class device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_CONFORMANCE_CLASS, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals(3, BACnetApplicationTagUnsignedInteger.getActualValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 40] device,201 protocol-services-supported",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_SERVICES_SUPPORTED, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 40] device,201 protocol-services-supported device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_SERVICES_SUPPORTED, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagBitString BACnetApplicationTagBitString = (BACnetApplicationTagBitString) value;
                    assertEquals(Arrays.asList(true, false, true, true, false, true, true, true, false, false, true, true, true, false, true, true, true, true, true, false, true, false, false, false, false, false, true, true, false, false, true, false, true, true, true), BACnetApplicationTagBitString.getData());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 41] device,201 protocol-object-types-supported",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_OBJECT_TYPES_SUPPORTED, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 41] device,201 protocol-object-types-supported device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_OBJECT_TYPES_SUPPORTED, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagBitString BACnetApplicationTagBitString = (BACnetApplicationTagBitString) value;
                    assertEquals(Arrays.asList(false, false, true, false, false, true, true, false, true, true, true, false, false, false, false, true, true, true), BACnetApplicationTagBitString.getData());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 42] device,201 max-apdu-length-accepted",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MAX_APDU_LENGTH_ACCEPTED, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 42] device,201 max-apdu-length-accepted device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MAX_APDU_LENGTH_ACCEPTED, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals(1476, BACnetApplicationTagUnsignedInteger.getActualValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 43] device,201 segmentation-supported",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.SEGMENTATION_SUPPORTED, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 43] device,201 segmentation-supported device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.SEGMENTATION_SUPPORTED, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagEnumerated BACnetApplicationTagEnumerated = (BACnetApplicationTagEnumerated) value;
                    assertEquals(List.of((byte) 0), BACnetApplicationTagEnumerated.getData());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 44] device,201 local-time",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.LOCAL_TIME, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 44] device,201 local-time device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.LOCAL_TIME, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagTime BACnetApplicationTagTime = (BACnetApplicationTagTime) value;
                    assertEquals(15, BACnetApplicationTagTime.getHour());
                    assertEquals(28, BACnetApplicationTagTime.getMinute());
                    assertEquals(41, BACnetApplicationTagTime.getSecond());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 45] device,201 local-date",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.LOCAL_DATE, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 45] device,201 local-date device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.LOCAL_DATE, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagDate BACnetApplicationTagDate = (BACnetApplicationTagDate) value;
                    assertEquals(2005, BACnetApplicationTagDate.getYear());
                    assertEquals(9, BACnetApplicationTagDate.getMonth());
                    assertEquals(1, BACnetApplicationTagDate.getDayOfMonth());
                    assertEquals(4, BACnetApplicationTagDate.getDayOfWeek());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 46] device,201 utc-offset",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.UTC_OFFSET, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 46] device,201 utc-offset device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.UTC_OFFSET, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagSignedInteger BACnetApplicationTagSignedInteger = (BACnetApplicationTagSignedInteger) value;
                    assertEquals(BigInteger.ZERO, BACnetApplicationTagSignedInteger.getActualValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 47] device,201 daylights-savings-status",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.DAYLIGHT_SAVINGS_STATUS, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 47] device,201 daylights-savings-status device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.DAYLIGHT_SAVINGS_STATUS, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagBoolean BACnetApplicationTagBoolean = (BACnetApplicationTagBoolean) value;
                    assertTrue(BACnetApplicationTagBoolean.getIsFalse());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 48] device,201 apdu-segment-timeout",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.APDU_SEGMENT_TIMEOUT, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 48] device,201 apdu-segment-timeout device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.APDU_SEGMENT_TIMEOUT, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals(6000, BACnetApplicationTagUnsignedInteger.getValueUint16());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 49] device,201 apdu-timeout",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.APDU_TIMEOUT, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 49] device,201 apdu-timeout device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.APDU_TIMEOUT, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals(6000, BACnetApplicationTagUnsignedInteger.getValueUint16());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 50] device,201 number-of-APDU-retries",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.NUMBER_OF_APDU_RETRIES, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 50] device,201 number-of-APDU-retries device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.NUMBER_OF_APDU_RETRIES, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals((short) 3, BACnetApplicationTagUnsignedInteger.getValueUint8());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 51] device,201 time-synchronization-recipients",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.TIME_SYNCHRONIZATION_RECIPIENTS, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("ERROR           readProperty[ 51] device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUError apduError = (APDUError) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetErrorReadProperty baCnetErrorReadProperty = (BACnetErrorReadProperty) apduError.getError();
                    // TODO: use proper enums
                    assertEquals(List.of((byte) 32), baCnetErrorReadProperty.getErrorCode().getData());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 52] device,201 max-master",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MAX_MASTER, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 52] device,201 max-master device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MAX_MASTER, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals((short) 127, BACnetApplicationTagUnsignedInteger.getValueUint8());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 53] device,201 max-info-frames",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MAX_INFO_FRAMES, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 53] device,201 max-info-frames device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MAX_INFO_FRAMES, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals((short) 40, BACnetApplicationTagUnsignedInteger.getValueUint8());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 54] device,201 device-address-binding",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.DEVICE_ADDRESS_BINDING, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 54] device,201 device-address-binding device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.DEVICE_ADDRESS_BINDING, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 55] device,201 (514) Vendor Proprietary Value",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 55] device,201 (514) Vendor Proprietary Value device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagBoolean BACnetApplicationTagBoolean = (BACnetApplicationTagBoolean) value;
                    assertTrue(BACnetApplicationTagBoolean.getIsFalse());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 56] device,201 (515) Vendor Proprietary Value device,201",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(201, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 56] device,201 Error",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUError apduError = (APDUError) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetErrorReadProperty baCnetErrorReadProperty = (BACnetErrorReadProperty) apduError.getError();
                    // TODO: use proper enums
                    assertEquals(List.of((byte) 32), baCnetErrorReadProperty.getErrorCode().getData());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 57] device,61 object-identifier",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_IDENTIFIER, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 57] object-identifier device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_IDENTIFIER, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagObjectIdentifier BACnetApplicationTagObjectIdentifier = (BACnetApplicationTagObjectIdentifier) value;
                    assertEquals(BACnetObjectType.DEVICE, BACnetApplicationTagObjectIdentifier.getObjectType());
                    assertEquals(61, BACnetApplicationTagObjectIdentifier.getInstanceNumber());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 58] device,61 object-name",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_NAME, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 58] object-name device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_NAME, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagCharacterString BACnetApplicationTagObjectIdentifier = (BACnetApplicationTagCharacterString) value;
                    assertEquals("SYNERGY", BACnetApplicationTagObjectIdentifier.getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 59] device,61 object-type",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_TYPE, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 59] object-type device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.OBJECT_TYPE, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagEnumerated baCnetApplicationTagEnumerated = (BACnetApplicationTagEnumerated) value;
                    assertEquals(List.of((byte) 8), baCnetApplicationTagEnumerated.getData());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 60] device,61 system-status",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.SYSTEM_STATUS, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 60] device,61 system-status device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.SYSTEM_STATUS, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagEnumerated baCnetApplicationTagEnumerated = (BACnetApplicationTagEnumerated) value;
                    assertEquals(List.of((byte) 0x0), baCnetApplicationTagEnumerated.getData());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 61] device,61 vendor-name",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_NAME, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 61] device,61 vendor-name device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_NAME, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagCharacterString BACnetApplicationTagCharacterString = (BACnetApplicationTagCharacterString) value;
                    assertEquals("Lithonia Lighting, Inc.", BACnetApplicationTagCharacterString.getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 62] device,61 vendor-identifier",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_IDENTIFIER, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 62] device,61 vendor-identifier device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_IDENTIFIER, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 63] device,61 model-name",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MODEL_NAME, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 63] device,61 model-name device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MODEL_NAME, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagCharacterString BACnetApplicationTagCharacterString = (BACnetApplicationTagCharacterString) value;
                    assertEquals("SYSC MLX", BACnetApplicationTagCharacterString.getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 64] device,61 model-name",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.FIRMWARE_REVISION, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 64] device,61 model-name device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.FIRMWARE_REVISION, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagCharacterString BACnetApplicationTagCharacterString = (BACnetApplicationTagCharacterString) value;
                    assertEquals("2x62i", BACnetApplicationTagCharacterString.getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 65] device,61 application-software-version",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.APPLICATION_SOFTWARE_VERSION, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 65] device,61 application-software-version device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.APPLICATION_SOFTWARE_VERSION, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagCharacterString BACnetApplicationTagCharacterString = (BACnetApplicationTagCharacterString) value;
                    assertEquals("10-Nov-2004", BACnetApplicationTagCharacterString.getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 66] device,61 protocol-version",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_VERSION, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 66] device,61 protocol-version device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_VERSION, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals(1, BACnetApplicationTagUnsignedInteger.getActualValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 67] device,61 protocol-conformance-class",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_CONFORMANCE_CLASS, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 67] device,61 protocol-conformance-class device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_CONFORMANCE_CLASS, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals(2, BACnetApplicationTagUnsignedInteger.getActualValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 68] device,61 protocol-services-supported",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_SERVICES_SUPPORTED, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 68] device,61 protocol-services-supported device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_SERVICES_SUPPORTED, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagBitString BACnetApplicationTagBitString = (BACnetApplicationTagBitString) value;
                    assertEquals(Arrays.asList(false, false, false, false, false, false, true, true, false, false, false, false, true, false, true, true, false, true, false, false, true, false, false, false, false, false, true, true, false, false, false, false, true, true, true), BACnetApplicationTagBitString.getData());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 69] device,61 protocol-object-types-supported",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_OBJECT_TYPES_SUPPORTED, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 69] device,61 protocol-object-types-supported device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.PROTOCOL_OBJECT_TYPES_SUPPORTED, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagBitString BACnetApplicationTagBitString = (BACnetApplicationTagBitString) value;
                    assertEquals(Arrays.asList(true, true, true, true, true, true, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false), BACnetApplicationTagBitString.getData());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 70] device,61 max-apdu-length-accepted",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MAX_APDU_LENGTH_ACCEPTED, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 70] device,61 max-apdu-length-accepted device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MAX_APDU_LENGTH_ACCEPTED, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals(480, BACnetApplicationTagUnsignedInteger.getActualValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 71] device,61 segmentation-supported",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.SEGMENTATION_SUPPORTED, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 71] device,61 segmentation-supported device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.SEGMENTATION_SUPPORTED, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagEnumerated BACnetApplicationTagEnumerated = (BACnetApplicationTagEnumerated) value;
                    assertEquals(List.of((byte) 0), BACnetApplicationTagEnumerated.getData());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 72] device,61 local-time",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.LOCAL_TIME, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 72] device,61 local-time device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.LOCAL_TIME, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagTime BACnetApplicationTagTime = (BACnetApplicationTagTime) value;
                    assertEquals(15, BACnetApplicationTagTime.getHour());
                    assertEquals(26, BACnetApplicationTagTime.getMinute());
                    assertEquals(40, BACnetApplicationTagTime.getSecond());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 73] device,61 local-date",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.LOCAL_DATE, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 73] device,61 local-date device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.LOCAL_DATE, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagDate BACnetApplicationTagDate = (BACnetApplicationTagDate) value;
                    assertEquals(2005, BACnetApplicationTagDate.getYear());
                    assertEquals(9, BACnetApplicationTagDate.getMonth());
                    assertEquals(1, BACnetApplicationTagDate.getDayOfMonth());
                    assertEquals(-1, BACnetApplicationTagDate.getDayOfWeek());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 74] device,61 utc-offset",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.UTC_OFFSET, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 74] device,61 utc-offset device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.UTC_OFFSET, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagSignedInteger BACnetApplicationTagSignedInteger = (BACnetApplicationTagSignedInteger) value;
                    assertEquals(BigInteger.valueOf(-300), BACnetApplicationTagSignedInteger.getActualValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 75] device,61 daylights-savings-status",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.DAYLIGHT_SAVINGS_STATUS, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 75] device,61 daylights-savings-status device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.DAYLIGHT_SAVINGS_STATUS, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagBoolean BACnetApplicationTagBoolean = (BACnetApplicationTagBoolean) value;
                    assertTrue(BACnetApplicationTagBoolean.getIsTrue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 76] device,61 apdu-segment-timeout",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.APDU_SEGMENT_TIMEOUT, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 76] device,61 apdu-segment-timeout device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.APDU_SEGMENT_TIMEOUT, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals(8000, BACnetApplicationTagUnsignedInteger.getValueUint16());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 77] device,61 apdu-timeout",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.APDU_TIMEOUT, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 77] device,61 apdu-timeout device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.APDU_TIMEOUT, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals(8000, BACnetApplicationTagUnsignedInteger.getValueUint16());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 78] device,61 number-of-APDU-retries",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.NUMBER_OF_APDU_RETRIES, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 78] device,61 number-of-APDU-retries device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.NUMBER_OF_APDU_RETRIES, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals((short) 3, BACnetApplicationTagUnsignedInteger.getValueUint8());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 79] device,61 time-synchronization-recipients",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.TIME_SYNCHRONIZATION_RECIPIENTS, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("ERROR           readProperty[ 79] device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUError apduError = (APDUError) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetErrorReadProperty baCnetErrorReadProperty = (BACnetErrorReadProperty) apduError.getError();
                    // TODO: use proper enums
                    assertEquals(List.of((byte) 32), baCnetErrorReadProperty.getErrorCode().getData());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 80] device,61 max-master",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MAX_MASTER, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 80] device,61 max-master device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MAX_MASTER, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals((short) 127, BACnetApplicationTagUnsignedInteger.getValueUint8());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 81] device,61 max-info-frames",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MAX_INFO_FRAMES, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 81] device,61 max-info-frames device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.MAX_INFO_FRAMES, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagUnsignedInteger BACnetApplicationTagUnsignedInteger = (BACnetApplicationTagUnsignedInteger) value;
                    assertEquals((short) 1, BACnetApplicationTagUnsignedInteger.getValueUint8());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 82] device,61 device-address-binding",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.DEVICE_ADDRESS_BINDING, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 82] device,61 device-address-binding device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.DEVICE_ADDRESS_BINDING, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 83] device,61 (514) Vendor Proprietary Value",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 83] device,61 (514) Vendor Proprietary Value device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagTime baCnetApplicationTagTime = (BACnetApplicationTagTime) value;
                    assertEquals(7, baCnetApplicationTagTime.getHour());
                    assertEquals(11, baCnetApplicationTagTime.getMinute());
                    assertEquals(38, baCnetApplicationTagTime.getSecond());
                }),
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[ 84] device,61 (515) Vendor Proprietary Value device,61",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetConfirmedServiceRequest serviceRequest = apduConfirmedRequest.getServiceRequest();
                    assertNotNull(serviceRequest);
                    BACnetConfirmedServiceRequestReadProperty baCnetConfirmedServiceRequestReadProperty = (BACnetConfirmedServiceRequestReadProperty) serviceRequest;
                    assertEquals(BACnetObjectType.DEVICE, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetConfirmedServiceRequestReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetConfirmedServiceRequestReadProperty.getPropertyIdentifier().getValue());
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[ 84] device,61 Error",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetServiceAck baCnetServiceAck = apduComplexAck.getServiceAck();
                    assertNotNull(baCnetServiceAck);
                    BACnetServiceAckReadProperty baCnetServiceAckReadProperty = (BACnetServiceAckReadProperty) baCnetServiceAck;
                    assertEquals(BACnetObjectType.DEVICE, baCnetServiceAckReadProperty.getObjectIdentifier().getObjectType());
                    assertEquals(61, baCnetServiceAckReadProperty.getObjectIdentifier().getInstanceNumber());
                    assertEquals(BACnetPropertyIdentifier.VENDOR_PROPRIETARY_VALUE, baCnetServiceAckReadProperty.getPropertyIdentifier().getValue());
                    BACnetApplicationTag value = baCnetServiceAckReadProperty.getValues().getData().get(0).getApplicationTag();
                    BACnetApplicationTagTime baCnetApplicationTagTime = (BACnetApplicationTagTime) value;
                    assertEquals(20, baCnetApplicationTagTime.getHour());
                    assertEquals(3, baCnetApplicationTagTime.getMinute());
                    assertEquals(18, baCnetApplicationTagTime.getSecond());
                })
        );
    }

    @TestFactory
    @DisplayName("read-property-bad")
    Collection<DynamicNode> read_property_bad() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("read-property-bad.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("read-property-epics")
    Collection<DynamicNode> read_property_epics() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("read-property-epics.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("read-property-synergy")
    Collection<DynamicNode> read_property_synergy() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("read-property-synergy.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("read-property")
    Collection<DynamicNode> read_property() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("read-property.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("readfile")
    Collection<DynamicNode> readfile() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("readfile.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("readrange_malformed")
    Collection<DynamicNode> readrange_malformed() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("readrange_malformed.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("reinit-device")
    Collection<DynamicNode> reinit_device() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("reinit-device.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("router")
    Collection<DynamicNode> router() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("router.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("routers")
    Collection<DynamicNode> routers() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("routers.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("rp-device")
    Collection<DynamicNode> rp_device() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("rp-device.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("rp-shed-level")
    Collection<DynamicNode> rp_shed_level() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("rp-shed-level.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Confirmed-REQ   readProperty[  1] load-control,0 expected-shed-level",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("Complex-ACK     readProperty[  1] load-control,0 expected-shed-level",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("rp")
    Collection<DynamicNode> rp() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("rp.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("rpm-error")
    Collection<DynamicNode> rpm_error() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("rpm-error.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("rpm")
    Collection<DynamicNode> rpm() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("rpm.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("rpm_multiple_scheduler_bug")
    Collection<DynamicNode> rpm_multiple_scheduler_bug() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("rpm_multiple_scheduler_bug.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("schedule-rpm-foreign")
    Collection<DynamicNode> schedule_rpm_foreign() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("schedule-rpm-foreign.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("signed_value_negative")
    Collection<DynamicNode> signed_value_negative() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("signed_value_negative.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("single-RPM")
    Collection<DynamicNode> single_RPM() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("single-RPM.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("softdel-BTL")
    Collection<DynamicNode> softdel_BTL() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("softdel-BTL.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @Disabled("Needs filtering")
    @TestFactory
    @DisplayName("special-events")
    Collection<DynamicNode> special_events() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("special-events.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("startup-exchange")
    Collection<DynamicNode> startup_exchange() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("startup-exchange.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("state_text")
    Collection<DynamicNode> state_text() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("state_text.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("state_text_good")
    Collection<DynamicNode> state_text_good() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("state_text_good.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("subordinatelist_rpm")
    Collection<DynamicNode> subordinatelist_rpm() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("subordinatelist_rpm.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("synergy-binding-2x63y")
    Collection<DynamicNode> synergy_binding_2x63y() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("synergy-binding-2x63y.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("synergy-broken-rpm")
    Collection<DynamicNode> synergy_broken_rpm() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("synergy-broken-rpm.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("synergy-device")
    Collection<DynamicNode> synergy_device() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("synergy-device.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("time-sync")
    Collection<DynamicNode> time_sync() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("time-sync.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @Disabled("Needs filtering")
    @TestFactory
    @DisplayName("tridium jace2")
    Collection<DynamicNode> tridium_jace2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("tridium%20jace2.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("u+4_MSTP")
    Collection<DynamicNode> u_4_MSTP() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("u+4_MSTP.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("weekend")
    Collection<DynamicNode> weekend() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("weekend.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("who-has-I-have")
    Collection<DynamicNode> who_has_I_have() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("who-has-I-have.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("who-has")
    Collection<DynamicNode> who_has() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("who-has.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Is 133 133",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestWhoIs baCnetUnconfirmedServiceRequestWhoIs = (BACnetUnconfirmedServiceRequestWhoIs) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(133, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeLowLimit().getActualValue());
                    assertEquals(133, baCnetUnconfirmedServiceRequestWhoIs.getDeviceInstanceRangeLowLimit().getActualValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Has device,133",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestWhoHas baCnetUnconfirmedServiceRequestWhoHas = (BACnetUnconfirmedServiceRequestWhoHas) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(133, baCnetUnconfirmedServiceRequestWhoHas.getDeviceInstanceRangeLowLimit().getActualValue());
                    assertEquals(133, baCnetUnconfirmedServiceRequestWhoHas.getDeviceInstanceRangeLowLimit().getActualValue());
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestWhoHas.getObjectIdentifier().getObjectType());
                    assertEquals(133, baCnetUnconfirmedServiceRequestWhoHas.getObjectIdentifier().getInstanceNumber());
                }),
            DynamicTest.dynamicTest("skip 2 LLC packages",
                () -> {
                    pcapEvaluator.skipPackages(2);
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ I-Am 133 133",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestIAm baCnetUnconfirmedServiceRequestIAm = (BACnetUnconfirmedServiceRequestIAm) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIAm.getDeviceIdentifier().getObjectType());
                    assertEquals(133, baCnetUnconfirmedServiceRequestIAm.getDeviceIdentifier().getInstanceNumber());
                    assertEquals(480, baCnetUnconfirmedServiceRequestIAm.getMaximumApduLengthAcceptedLength().getActualValue());
                    // TODO: we should use a enum here
                    assertEquals(List.of((byte) (byte) 0x0), baCnetUnconfirmedServiceRequestIAm.getSegmentationSupported().getData());
                    assertEquals(42, baCnetUnconfirmedServiceRequestIAm.getVendorId().getActualValue());
                }),
            DynamicContainer.dynamicContainer("Confirmed-REQ atomicWriteFile 1-30", () -> {
                Collection<DynamicNode> nodes = new LinkedList<>();
                IntStream.range(1, 31).forEach(i -> {
                    nodes.add(DynamicTest.dynamicTest("Confirmed-REQ atomicWriteFile [" + i + "] file,0", () -> {
                        BVLC bvlc = pcapEvaluator.nextBVLC();
                        dump(bvlc);
                        BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                        APDUConfirmedRequest apduConfirmedRequest = (APDUConfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                        BACnetConfirmedServiceRequestAtomicWriteFile baCnetConfirmedServiceRequestAtomicWriteFile = (BACnetConfirmedServiceRequestAtomicWriteFile) apduConfirmedRequest.getServiceRequest();
                        assertEquals(BACnetObjectType.FILE, baCnetConfirmedServiceRequestAtomicWriteFile.getDeviceIdentifier().getObjectType());
                        assertNotNull(baCnetConfirmedServiceRequestAtomicWriteFile.getFileStartPosition());
                        assertNotNull(baCnetConfirmedServiceRequestAtomicWriteFile.getFileData());
                    }));
                    nodes.add(DynamicTest.dynamicTest("Confirmed-Ack     atomicWriteFile [" + i + "]", () -> {
                        BVLC bvlc = pcapEvaluator.nextBVLC();
                        dump(bvlc);
                        BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                        APDUComplexAck apduComplexAck = (APDUComplexAck) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                        BACnetServiceAckAtomicWriteFile baCnetServiceAckAtomicWriteFile = (BACnetServiceAckAtomicWriteFile) apduComplexAck.getServiceAck();
                        assertNotNull(baCnetServiceAckAtomicWriteFile.getFileStartPosition());
                    }));
                });
                return nodes.iterator();
            }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Has device,133",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestWhoHas baCnetUnconfirmedServiceRequestWhoHas = (BACnetUnconfirmedServiceRequestWhoHas) apduUnconfirmedRequest.getServiceRequest();
                    assertNull(baCnetUnconfirmedServiceRequestWhoHas.getDeviceInstanceRangeLowLimit());
                    assertNull(baCnetUnconfirmedServiceRequestWhoHas.getDeviceInstanceRangeLowLimit());
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestWhoHas.getObjectIdentifier().getObjectType());
                    assertEquals(133, baCnetUnconfirmedServiceRequestWhoHas.getObjectIdentifier().getInstanceNumber());
                }),
            DynamicTest.dynamicTest("skip 1 LLC packages",
                () -> pcapEvaluator.skipPackages(1)),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Have device,4194303 device,133",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestIHave baCnetUnconfirmedServiceRequestIHave = (BACnetUnconfirmedServiceRequestIHave) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIHave.getDeviceIdentifier().getObjectType());
                    assertEquals(4194303, baCnetUnconfirmedServiceRequestIHave.getDeviceIdentifier().getInstanceNumber());
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIHave.getObjectIdentifier().getObjectType());
                    assertEquals(133, baCnetUnconfirmedServiceRequestIHave.getObjectIdentifier().getInstanceNumber());
                    assertEquals("Unknown", baCnetUnconfirmedServiceRequestIHave.getObjectName().getValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Have device,133 device,133",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestIHave baCnetUnconfirmedServiceRequestIHave = (BACnetUnconfirmedServiceRequestIHave) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIHave.getDeviceIdentifier().getObjectType());
                    assertEquals(133, baCnetUnconfirmedServiceRequestIHave.getDeviceIdentifier().getInstanceNumber());
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIHave.getObjectIdentifier().getObjectType());
                    assertEquals(133, baCnetUnconfirmedServiceRequestIHave.getObjectIdentifier().getInstanceNumber());
                    assertEquals("SYNERGY", baCnetUnconfirmedServiceRequestIHave.getObjectName().getValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ who-Has device,133",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestWhoHas baCnetUnconfirmedServiceRequestWhoHas = (BACnetUnconfirmedServiceRequestWhoHas) apduUnconfirmedRequest.getServiceRequest();
                    assertNull(baCnetUnconfirmedServiceRequestWhoHas.getDeviceInstanceRangeLowLimit());
                    assertNull(baCnetUnconfirmedServiceRequestWhoHas.getDeviceInstanceRangeLowLimit());
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestWhoHas.getObjectIdentifier().getObjectType());
                    assertEquals(133, baCnetUnconfirmedServiceRequestWhoHas.getObjectIdentifier().getInstanceNumber());
                }),
            DynamicTest.dynamicTest("skip 1 LLC packages",
                () -> pcapEvaluator.skipPackages(1)),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Have device,4194303 device,133",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalBroadcastNPDU bvlcOriginalBroadcastNPDU = (BVLCOriginalBroadcastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalBroadcastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestIHave baCnetUnconfirmedServiceRequestIHave = (BACnetUnconfirmedServiceRequestIHave) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIHave.getDeviceIdentifier().getObjectType());
                    assertEquals(4194303, baCnetUnconfirmedServiceRequestIHave.getDeviceIdentifier().getInstanceNumber());
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIHave.getObjectIdentifier().getObjectType());
                    assertEquals(133, baCnetUnconfirmedServiceRequestIHave.getObjectIdentifier().getInstanceNumber());
                    assertEquals("Unknown", baCnetUnconfirmedServiceRequestIHave.getObjectName().getValue());
                }),
            DynamicTest.dynamicTest("Unconfirmed-REQ i-Have device,133 device,133",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
                    APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) bvlcOriginalUnicastNPDU.getNpdu().getApdu();
                    BACnetUnconfirmedServiceRequestIHave baCnetUnconfirmedServiceRequestIHave = (BACnetUnconfirmedServiceRequestIHave) apduUnconfirmedRequest.getServiceRequest();
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIHave.getDeviceIdentifier().getObjectType());
                    assertEquals(133, baCnetUnconfirmedServiceRequestIHave.getDeviceIdentifier().getInstanceNumber());
                    assertEquals(BACnetObjectType.DEVICE, baCnetUnconfirmedServiceRequestIHave.getObjectIdentifier().getObjectType());
                    assertEquals(133, baCnetUnconfirmedServiceRequestIHave.getObjectIdentifier().getInstanceNumber());
                    assertEquals("SYNERGY", baCnetUnconfirmedServiceRequestIHave.getObjectName().getValue());
                })
        );
    }

    @TestFactory
    @DisplayName("who-is-i-am")
    Collection<DynamicNode> who_is_i_am() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("who-is-i-am.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("whois-basrtp-b-1")
    Collection<DynamicNode> whois_basrtp_b_1() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("whois-basrtp-b-1.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("whois-basrtp-b-2")
    Collection<DynamicNode> whois_basrtp_b_2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("whois-basrtp-b-2.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("whois-iam")
    Collection<DynamicNode> whois_iam() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("whois-iam.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("wireshark_BBMDError")
    Collection<DynamicNode> wireshark_BBMDError() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("wireshark_BBMDError.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("wireshark_CEN_9_11")
    Collection<DynamicNode> wireshark_CEN_9_11() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("wireshark_CEN_9_11.pcap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("wp-rp-index")
    Collection<DynamicNode> wp_rp_index() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("wp-rp-index.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("wp-weekly-schedule-index")
    Collection<DynamicNode> wp_weekly_schedule_index() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("wp-weekly-schedule-index.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("wp-weekly-schedule-test")
    Collection<DynamicNode> wp_weekly_schedule_test() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("wp-weekly-schedule-test.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("wp_weekly_schedule")
    Collection<DynamicNode> wp_weekly_schedule() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("wp_weekly_schedule.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    assumeTrue(false, "BACnetConfirmedServiceRequestWriteProperty wrongly implemented");
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    assumeTrue(false, "BACnetConfirmedServiceRequestWriteProperty wrongly implemented");
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("write-property-array")
    Collection<DynamicNode> write_property_array() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("write-property-array.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("write-property-multiple")
    Collection<DynamicNode> write_property_multiple() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("write-property-multiple.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("write-property-wattstopper-panel")
    Collection<DynamicNode> write_property_wattstopper_panel() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("write-property-wattstopper-panel.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("write-property")
    Collection<DynamicNode> write_property() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("write-property.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }

    @TestFactory
    @DisplayName("write-property2")
    Collection<DynamicNode> write_property2() throws Exception {
        PCAPEvaluator pcapEvaluator = pcapEvaluator("write-property2.cap");
        return Arrays.asList(
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                }),
            DynamicTest.dynamicTest("TODO",
                () -> {
                    BVLC bvlc = pcapEvaluator.nextBVLC();
                    assertNotNull(bvlc);
                    dump(bvlc);
                    // TODO:
                    assumeTrue(false, "not properly implemented. Check manually and add asserts");
                })
        );
    }


    private void dump(Serializable serializable) throws SerializationException {
        WriteBufferBoxBased writeBuffer = new WriteBufferBoxBased(true, true);
        serializable.serialize(writeBuffer);
        LOGGER.info("{}\n{}", serializable.getClass().getName(), writeBuffer.getBox());
    }

    private PCAPEvaluator pcapEvaluator(String pcapFile) throws IOException, PcapNativeException {
        PCAPEvaluator pcapEvaluator = new PCAPEvaluator(pcapFile);
        toBeClosed.offer(pcapEvaluator);
        return pcapEvaluator;
    }

    private static class PCAPEvaluator implements Closeable {
        private int currentPackage = 0;
        private final String pcapFile;
        private final PcapHandle pcapHandle;

        public PCAPEvaluator(String pcapFile) throws IOException, PcapNativeException {
            this.pcapFile = pcapFile;
            String toParse = DownloadAndCache(pcapFile);
            LOGGER.info("Reading " + toParse);
            pcapHandle = getHandle(toParse);
        }

        public void skipTo(int packageNumber) {
            if (packageNumber <= currentPackage) {
                throw new IllegalArgumentException("Package number must be bigger than " + currentPackage);
            }
            LOGGER.info("Skipping to package number {}", packageNumber);
            skipPackages(packageNumber - currentPackage - 1);
        }

        public void skipPackages(int numberOfPackages) {
            IntStream.rangeClosed(1, numberOfPackages).forEach(i -> {
                LOGGER.info("Skipping package " + (currentPackage + i));
                try {
                    pcapHandle.getNextPacket();
                } catch (NotOpenException e) {
                    e.printStackTrace();
                }
            });
            currentPackage += numberOfPackages;
        }

        private int getCurrentPackageNumber() {
            return currentPackage;
        }

        private BVLC nextBVLC() throws NotOpenException, ParseException {
            currentPackage += 1;
            Packet packet = pcapHandle.getNextPacket();
            LOGGER.info("({})Next packet:\n{}", currentPackage, packet);
            if (packet == null) {
                return null;
            }

            UdpPacket udpPacket = packet.get(UdpPacket.class);
            assumeTrue(udpPacket != null, "nextBVLC assumes a UDP Packet. If non is there it might by LLC");
            LOGGER.info("Handling UDP\n{}", udpPacket);
            byte[] rawData = udpPacket.getPayload().getRawData();
            LOGGER.info("Reading BVLC from:\n{}", Hex.dump(rawData));
            return BVLCIO.staticParse(new ReadBufferByteBased(rawData));
        }

        private PcapHandle getHandle(String file) throws PcapNativeException {
            return Pcaps.openOffline(file, PcapHandle.TimestampPrecision.NANO);
        }

        private String DownloadAndCache(String file) throws IOException {
            String tempDirectory = FileUtils.getTempDirectoryPath();
            File pcapFile = FileSystems.getDefault().getPath(tempDirectory, RandomPackagesTest.class.getSimpleName(), file).toFile();
            FileUtils.createParentDirectories(pcapFile);
            if (!pcapFile.exists()) {
                URL source = new URL("http://kargs.net/captures/" + file);
                LOGGER.info("Downloading {}", source);
                FileUtils.copyURLToFile(source, pcapFile);
            }
            return pcapFile.getAbsolutePath();
        }

        @Override
        public void close() throws IOException {
            pcapHandle.close();
        }

        @Override
        public String toString() {
            return "PCAPEvaluator{" +
                "pcapFile='" + pcapFile + '\'' +
                ", pcapHandle=" + pcapHandle +
                '}';
        }
    }
}
