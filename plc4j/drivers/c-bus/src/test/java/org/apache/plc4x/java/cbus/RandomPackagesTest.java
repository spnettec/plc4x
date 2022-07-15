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
package org.apache.plc4x.java.cbus;

import org.apache.commons.codec.binary.Hex;
import org.apache.plc4x.java.cbus.readwrite.*;
import org.apache.plc4x.java.spi.generation.ReadBufferByteBased;
import org.apache.plc4x.java.spi.generation.WriteBufferByteBased;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;

import java.nio.charset.StandardCharsets;

import static org.assertj.core.api.Assertions.assertThat;

public class RandomPackagesTest {

    RequestContext requestContext;
    CBusOptions cBusOptions;

    @BeforeEach
    void setUp() {
        requestContext = new RequestContext(false, false, false);
        cBusOptions = new CBusOptions(false, false, false, false, false, false, false, false, false);
    }

    // from: https://updates.clipsal.com/ClipsalSoftwareDownload/DL/downloads/OpenCBus/Serial%20Interface%20User%20Guide.pdf
    @Nested
    class ReferenceDocumentationTest {

        // 3.4
        @Nested
        class Header {
            @Test
            void Point_point_multipoint_lowest_priority_class() throws Exception {
                byte[] bytes = new byte[]{0x03};
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusHeader msg = CBusHeader.staticParse(readBufferByteBased);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Test
            void Point_multipoint_lowest_priority_class() throws Exception {
                byte[] bytes = new byte[]{0x05};
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusHeader msg = CBusHeader.staticParse(readBufferByteBased);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Test
            void Point_point_lowest_priority_class() throws Exception {
                byte[] bytes = new byte[]{0x06};
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusHeader msg = CBusHeader.staticParse(readBufferByteBased);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

        }

        // 4
        @Nested
        class SerialInterface {

            // 4.2.3
            @Test
            void reset() throws Exception {
                byte[] bytes = "~\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Disabled("not implemented yet")
            // 4.2.4
            @Test
            void cancel() throws Exception {
                byte[] bytes = "AB0123?9876\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            // 4.2.5
            @Test
            void smartConnectShortcut() throws Exception {
                byte[] bytes = "\r|\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Disabled("not implemented yet")
            // 4.2.4
            @Test
            void confirmation() throws Exception {
                // If you follow the spec a confirmation can occur at any place... seems strange
                byte[] bytes = "AB0123n9876\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            // 4.2.7
            @Test
            void directCommandAccess1() throws Exception {
                byte[] bytes = "@2102\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                CALData calData = ((CALDataOrSetParameterValue) ((RequestDirectCommandAccess) ((CBusMessageToServer) msg).getRequest()).getCalDataOrSetParameter()).getCalData();
                System.out.println(calData);
            }

            @Disabled("not implemented yet")
            // 4.2.7
            @Test
            void directCommandAccess2() throws Exception {
                // TODO: this should be the same as the @above but that is not yet implemented
                byte[] bytes = "~2102\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                CALData calData = ((CALDataOrSetParameterValue) ((RequestDirectCommandAccess) ((CBusMessageToServer) msg).getRequest()).getCalDataOrSetParameter()).getCalData();
                System.out.println(calData);
            }

        }


        // 4.2.9.1
        @Nested
        class PointToPointCommands {
            @Test
            void pointToPointCommandDirect() throws Exception {
                byte[] bytes = "\\0603002102D4\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                CBusCommand cbusCommand = ((RequestCommand) ((CBusMessageToServer) msg).getRequest()).getCbusCommand();
                System.out.println(cbusCommand);
            }

            @Test
            void pointToPointCommandBridged() throws Exception {
                byte[] bytes = "\\06420903210289\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                CBusCommand cbusCommand = ((RequestCommand) ((CBusMessageToServer) msg).getRequest()).getCbusCommand();
                System.out.println(cbusCommand);
            }
        }


        // 4.2.9.2
        @Nested
        class PointToMultiPointCommands {
            @Test
            void pointToMultiPointCommandDirect() throws Exception {
                byte[] bytes = "\\0538000108BA\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                CBusCommand cbusCommand = ((RequestCommand) ((CBusMessageToServer) msg).getRequest()).getCbusCommand();
                System.out.println(cbusCommand);
            }

            @Test
            void pointToMultiPointCommandBridged() throws Exception {
                byte[] bytes = "\\05FF007A38004A\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                CBusCommand cbusCommand = ((RequestCommand) ((CBusMessageToServer) msg).getRequest()).getCbusCommand();
                System.out.println(cbusCommand);
            }
        }

        // 4.2.9.3
        @Nested
        class PointToPointToMultoPointCommands {
            @Test
            void pointToPointToMultiPointCommand2() throws Exception {
                byte[] bytes = "\\03420938010871\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                CBusCommand cbusCommand = ((RequestCommand) ((CBusMessageToServer) msg).getRequest()).getCbusCommand();
                System.out.println(cbusCommand);
            }
        }

        // 4.3.3.1
        @Nested
        class _CALReply {
            @Test
            void calRequest() throws Exception {
                byte[] bytes = "\\0605002102\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Test
            void calReplyNormal() throws Exception {
                byte[] bytes = Hex.decodeHex("8902312E322E363620200A");
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CALReply msg = CALReplyShort.staticParse(readBufferByteBased, cBusOptions, requestContext);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Test
            void calReplySmart() throws Exception {
                byte[] bytes = Hex.decodeHex("860593008902312E322E363620207F");
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CALReply msg = CALReplyLong.staticParse(readBufferByteBased, cBusOptions, requestContext);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }
        }

        // 4.3.3.2
        @Nested
        class _MonitoredSAL {
            @Test
            void monitoredSal() throws Exception {
                byte[] bytes = "0503380079083F\r\n".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                ReplyOrConfirmation msg = ReplyOrConfirmation.staticParse(readBufferByteBased, cBusOptions, bytes.length, requestContext);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                MonitoredSAL monitoredSAL = ((MonitoredSALReply) ((ReplyOrConfirmationReply) msg).getReply()).getMonitoredSAL();
                System.out.println(monitoredSAL);
            }
        }

        // 4.3.3.3
        @Nested
        class Confirmation {
            @Test
            void successful() throws Exception {
                byte[] bytes = "g.".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                ReplyOrConfirmation msg = ReplyOrConfirmation.staticParse(readBufferByteBased, cBusOptions, bytes.length, requestContext);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Test
            void toManyRetransmissions() throws Exception {
                byte[] bytes = "g#".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                ReplyOrConfirmation msg = ReplyOrConfirmation.staticParse(readBufferByteBased, cBusOptions, bytes.length, requestContext);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Test
            void corruption() throws Exception {
                byte[] bytes = "g$".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                ReplyOrConfirmation msg = ReplyOrConfirmation.staticParse(readBufferByteBased, cBusOptions, bytes.length, requestContext);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Test
            void desync() throws Exception {
                byte[] bytes = "g%".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                ReplyOrConfirmation msg = ReplyOrConfirmation.staticParse(readBufferByteBased, cBusOptions, bytes.length, requestContext);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Test
            void tooLong() throws Exception {
                byte[] bytes = "g'".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                ReplyOrConfirmation msg = ReplyOrConfirmation.staticParse(readBufferByteBased, cBusOptions, bytes.length, requestContext);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }
        }

        // 7.3
        @Test
        void StandardFormatStatusReply1() throws Exception {
            byte[] bytes = Hex.decodeHex("D8380068AA0140550550001000000014000000000000000000CF");
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            StandardFormatStatusReply msg = StandardFormatStatusReply.staticParse(readBufferByteBased, false);
            assertThat(msg).isNotNull();
            System.out.println(msg);
        }

        @Test
        void StandardFormatStatusReply2() throws Exception {
            byte[] bytes = Hex.decodeHex("D838580000000000000000000000000000000000000000000098");
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            StandardFormatStatusReply msg = StandardFormatStatusReply.staticParse(readBufferByteBased, false);
            assertThat(msg).isNotNull();
            System.out.println(msg);
        }

        @Test
        void StandardFormatStatusReply3() throws Exception {
            byte[] bytes = Hex.decodeHex("D638B000000000FF00000000000000000000000000000043");
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            StandardFormatStatusReply msg = StandardFormatStatusReply.staticParse(readBufferByteBased, false);
            assertThat(msg).isNotNull();
            System.out.println(msg);
        }

        // 7.4
        @Test
        void ExtendedFormatStatusReply1() throws Exception {
            byte[] bytes = Hex.decodeHex("F9073800AAAA000095990000000005555000000000005555555548");
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            ExtendedFormatStatusReply msg = ExtendedFormatStatusReply.staticParse(readBufferByteBased, false);
            assertThat(msg).isNotNull();
            System.out.println(msg);
        }

        @Test
        void ExtendedFormatStatusReply2() throws Exception {
            byte[] bytes = Hex.decodeHex("F907380B0000000000005555000000000000000000000000000013");
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            ExtendedFormatStatusReply msg = ExtendedFormatStatusReply.staticParse(readBufferByteBased, false);
            assertThat(msg).isNotNull();
            System.out.println(msg);
        }

        @Test
        void ExtendedFormatStatusReply3() throws Exception {
            byte[] bytes = Hex.decodeHex("f70738160000000000000000000000000000000000000000B4");
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            ExtendedFormatStatusReply msg = ExtendedFormatStatusReply.staticParse(readBufferByteBased, false);
            assertThat(msg).isNotNull();
            System.out.println(msg);
        }

        // 9.1
        @Nested
        class PointToMultiPointCommandsIntoLocalCBusNetwork {
            @Test
            void LightningOff() throws Exception {
                // TODO: the section describes that on non smart mode the message doesn't have the last CR
                byte[] bytes = "\\0538000114AE\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Test
            void LightningStatus() throws Exception {
                // TODO: the section describes that on non smart mode the message doesn't have the last CR
                byte[] bytes = "\\05FF007A38004A\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Test
            void LightningStatusReply1() throws Exception {
                // TODO: the section describes that on non smart mode the message doesn't have the last CR
                byte[] bytes = "D83800A8AA02000000000000000000000000000000000000009C\r\n".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                requestContext = new RequestContext(false, true, false);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                StandardFormatStatusReply reply = ((ReplyStandardFormatStatusReply) ((ReplyOrConfirmationReply) ((CBusMessageToClient) msg).getReply()).getReply()).getReply();
                System.out.println(reply);
            }

            @Test
            void LightningStatusReply2() throws Exception {
                // TODO: the section describes that on non smart mode the message doesn't have the last CR
                byte[] bytes = "D838580000000000000000000000000000000000000000000098\r\n".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                requestContext = new RequestContext(false, true, false);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                StandardFormatStatusReply reply = ((ReplyStandardFormatStatusReply) ((ReplyOrConfirmationReply) ((CBusMessageToClient) msg).getReply()).getReply()).getReply();
                System.out.println(reply);
            }

            @Test
            void LightningStatusReply3() throws Exception {
                // TODO: the section describes that on non smart mode the message doesn't have the last CR
                byte[] bytes = "D638B0000000000000000000000000000000000000000042\r\n".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                requestContext = new RequestContext(false, true, false);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                StandardFormatStatusReply reply = ((ReplyStandardFormatStatusReply) ((ReplyOrConfirmationReply) ((CBusMessageToClient) msg).getReply()).getReply()).getReply();
                System.out.println(reply);
            }

            @Test
            void LightningStatusReply4() throws Exception {
                // TODO: the section describes that on non smart mode the message doesn't have the last CR
                byte[] bytes = "86999900F8003800A8AA0200000000000000000000000000000000000000C4\r\n".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                requestContext = new RequestContext(false, true, false);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                StandardFormatStatusReply reply = ((ReplyStandardFormatStatusReply) ((ReplyOrConfirmationReply) ((CBusMessageToClient) msg).getReply()).getReply()).getReply();
                System.out.println(reply);
            }


            @Test
            void LightningStatusReply5() throws Exception {
                // TODO: the section describes that on non smart mode the message doesn't have the last CR
                byte[] bytes = "86999900F800385800000000000000000000000000000000000000000000C0\r\n".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                requestContext = new RequestContext(false, true, false);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                StandardFormatStatusReply reply = ((ReplyStandardFormatStatusReply) ((ReplyOrConfirmationReply) ((CBusMessageToClient) msg).getReply()).getReply()).getReply();
                System.out.println(reply);
            }

            @Test
            void LightningStatusReply6() throws Exception {
                // TODO: the section describes that on non smart mode the message doesn't have the last CR
                byte[] bytes = "86999900F60038B000000000000000000000000000000000000000008F\r\n".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                requestContext = new RequestContext(false, true, false);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                StandardFormatStatusReply reply = ((ReplyStandardFormatStatusReply) ((ReplyOrConfirmationReply) ((CBusMessageToClient) msg).getReply()).getReply()).getReply();
                System.out.println(reply);
            }
        }

        // 9.2
        @Nested
        class PointToPointCommandsIntoLocalCBusNetwork {
            @Test
            void RecallCurrentValueOfParameter0x30onUnit0x04() throws Exception {
                // TODO: the section describes that on non smart mode the message doesn't have the last CR
                byte[] bytes = "\\0604001A3001AB\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

            @Test
            void Reply() throws Exception {
                byte[] bytes = "8604990082300328\r\n".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                requestContext = new RequestContext(true, false, false);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

        }

        // 9.3
        @Nested
        class PointToMultiPointCommandsIntoaRemoteCBusNetwork {
            @Test
            void IssueLightningOf() throws Exception {
                // TODO: the section describes that on non smart mode the message doesn't have the last CR
                byte[] bytes = "\\03421B53643801149C\r".getBytes(StandardCharsets.UTF_8);
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
                assertThat(msg).isNotNull();
                System.out.println(msg);
                CBusCommand cbusCommand = ((RequestCommand) ((CBusMessageToServer) msg).getRequest()).getCbusCommand();
                System.out.println(cbusCommand);
            }

            @Disabled("the transformation from point to point to multipoint message is not yet implemented as described in 8.4... not sure if we will ever implement that")
            @Test
            void Reply() throws Exception {
                byte[] bytes = Hex.decodeHex("0565380354432101148E");
                ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
                CBusCommand msg = CBusCommand.staticParse(readBufferByteBased, cBusOptions);
                assertThat(msg).isNotNull();
                System.out.println(msg);
            }

        }

        // 9.4
        @Test
        void SwitchMode() throws Exception {
            byte[] bytes = "~@A3300019\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
        }

        // 9.5
        @Test
        void MultipleCommands() throws Exception {
            byte[] bytes = "\\05380001210122012301240A25010A2601D4\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
        }

        // 10.2.1
        @Test
        void testParameterSet() throws Exception {
            byte[] bytes = "@A3470011\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
        }

        // 10.2.1
        @Test
        void testParameterSetObsolete() throws Exception {
            byte[] bytes = "A3470011\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);

            CALDataOrSetParameterSetParameter calDataOrSetParameter = (CALDataOrSetParameterSetParameter) ((RequestObsolete) ((CBusMessageToServer) msg).getRequest()).getCalDataOrSetParameter();
            System.out.println(calDataOrSetParameter);
        }
    }

    @Nested
    class CBusQuickStartGuideTest {

        // 4.2.9.1
        @Test
        void pointToPointCommandDirect() throws Exception {
            byte[] bytes = "\\0538007902D4\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
        }

    }

    @Nested
    class OwnCaptures {

        @Disabled
        @Test
        void whatEverThisIs() throws Exception {
            byte[] bytes = "\\3436303230303231303167\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
        }

        @Test
        void deviceManagementInstruction() throws Exception {
            byte[] bytes = "@1A2001\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
        }

        @Test
        void setLight() throws Exception {
            byte[] bytes = "\\0538000100g\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            CBusMessageToServer msgToServer = (CBusMessageToServer) msg;
            RequestCommand requestCommand = (RequestCommand) msgToServer.getRequest();
            CBusCommand cbusCommand = requestCommand.getCbusCommand();
            System.out.println(cbusCommand);
        }

        @Test
        void identifyResponse() throws Exception {
            byte[] bytes = "g.890150435F434E49454421\r\n".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            // We know we send an identify command so we set the cal flag
            requestContext = new RequestContext(true, false, false);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            CBusMessageToClient messageToClient = (CBusMessageToClient) msg;
            ReplyOrConfirmationConfirmation confirmationReply = (ReplyOrConfirmationConfirmation) messageToClient.getReply();
            ReplyOrConfirmationReply normalReply = (ReplyOrConfirmationReply) confirmationReply.getEmbeddedReply();
            ReplyCALReply calReplyReply = (ReplyCALReply) normalReply.getReply();
            System.out.println(calReplyReply.getCalReply());
        }

        @Test
        void someResponse() throws Exception {
            byte[] bytes = "nl.8220025C\r\n".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            requestContext = new RequestContext(true, false, false);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            CBusMessageToClient messageToClient = (CBusMessageToClient) msg;
            ReplyOrConfirmationConfirmation confirmationReply = (ReplyOrConfirmationConfirmation) messageToClient.getReply();
            ReplyOrConfirmationReply normalReply = (ReplyOrConfirmationReply) confirmationReply.getEmbeddedReply();
            ReplyCALReply calReplyReply = (ReplyCALReply) normalReply.getReply();
            System.out.println(calReplyReply.getCalReply());
        }

        @Test
        void someOtherResponse() throws Exception {
            byte[] bytes = "\\0538000100g\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            CBusMessageToServer messageToServer = (CBusMessageToServer) msg;
            RequestCommand requestCommand = (RequestCommand) messageToServer.getRequest();
            System.out.println(requestCommand.getCbusCommand());
        }


        @Test
        void identifyRequest2() throws Exception {
            byte[] bytes = "21021A2102i\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            CBusMessageToServer messageToServer = (CBusMessageToServer) msg;
            RequestObsolete requestObsolete = (RequestObsolete) messageToServer.getRequest();
            CALData calData = ((CALDataOrSetParameterValue) requestObsolete.getCalDataOrSetParameter()).getCalData();
            System.out.println(calData);
        }

        @Test
        void identifyResponse2() throws Exception {
            byte[] bytes = "i.8902352E342E3030202010\r\n".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            // We know we send an identify command so we set the cal flag
            requestContext = new RequestContext(true, false, true);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            CBusMessageToClient messageToClient = (CBusMessageToClient) msg;
            ReplyOrConfirmationConfirmation confirmationReply = (ReplyOrConfirmationConfirmation) messageToClient.getReply();
            ReplyOrConfirmationReply normalReply = (ReplyOrConfirmationReply) confirmationReply.getEmbeddedReply();
            System.out.println(((ReplyCALReply) normalReply.getReply()).getCalReply());
        }

        @Test
        void recall() throws Exception {
            byte[] bytes = "@1A2001\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            CBusMessageToServer messageToServer = (CBusMessageToServer) msg;
            RequestDirectCommandAccess requestDirectCommandAccess = (RequestDirectCommandAccess) messageToServer.getRequest();
            CALDataOrSetParameter calDataOrSetParameter = requestDirectCommandAccess.getCalDataOrSetParameter();
            System.out.println(calDataOrSetParameter);

            WriteBufferByteBased writeBuffer = new WriteBufferByteBased(bytes.length);
            msg.serialize(writeBuffer);
            assertThat(writeBuffer.getBytes()).isEqualTo(bytes);
        }

        @Test
        void identifyTypeReply() throws Exception {
            byte[] bytes = "h.890150435F434E49454421\r\n".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            requestContext = new RequestContext(true, false, true);
            cBusOptions = new CBusOptions(false, false, false, false, false, false, false, false, true);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            CBusMessageToClient messageToClient = (CBusMessageToClient) msg;
            ReplyOrConfirmationConfirmation confirmationReply = (ReplyOrConfirmationConfirmation) messageToClient.getReply();
            ReplyOrConfirmationReply normalReply = (ReplyOrConfirmationReply) confirmationReply.getEmbeddedReply();
            ReplyCALReply calReplyReply = (ReplyCALReply) normalReply.getReply();
            System.out.println(calReplyReply.getCalReply());

            WriteBufferByteBased writeBuffer = new WriteBufferByteBased(bytes.length);
            msg.serialize(writeBuffer);
            byte[] bytes1 = writeBuffer.getBytes();
            assertThat(bytes1).isEqualTo(bytes);
        }

        @Disabled
        @Test
        void strangeNotYetParsableCommand() throws Exception {
            byte[] bytes = "A3309755s\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            cBusOptions = new CBusOptions(false, false, false, false, false, false, false, false, true);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            CBusMessageToClient messageToClient = (CBusMessageToClient) msg;
            ReplyOrConfirmationConfirmation confirmationReply = (ReplyOrConfirmationConfirmation) messageToClient.getReply();
            ReplyOrConfirmationReply normalReply = (ReplyOrConfirmationReply) confirmationReply.getEmbeddedReply();
            ReplyCALReply calReplyReply = (ReplyCALReply) normalReply.getReply();
            System.out.println(calReplyReply.getCalReply());

            WriteBufferByteBased writeBuffer = new WriteBufferByteBased(bytes.length);
            msg.serialize(writeBuffer);
            byte[] bytes1 = writeBuffer.getBytes();
            assertThat(bytes1).isEqualTo(bytes);
        }

        @Test
        void strangeNotYetParsableCommandResponse() throws Exception {
            byte[] bytes = "s.860202003230977D\r\n".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            requestContext = new RequestContext(true, false, false);
            cBusOptions = new CBusOptions(false, false, false, false, false, false, false, false, true);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            CBusMessageToClient messageToClient = (CBusMessageToClient) msg;
            ReplyOrConfirmationConfirmation confirmationReply = (ReplyOrConfirmationConfirmation) messageToClient.getReply();
            ReplyOrConfirmationReply normalReply = (ReplyOrConfirmationReply) confirmationReply.getEmbeddedReply();
            ReplyCALReply calReplyReply = (ReplyCALReply) normalReply.getReply();
            System.out.println(calReplyReply.getCalReply());

            WriteBufferByteBased writeBuffer = new WriteBufferByteBased(bytes.length);
            msg.serialize(writeBuffer);
            byte[] bytes1 = writeBuffer.getBytes();
            assertThat(bytes1).isEqualTo(bytes);
        }

        @Test
        void statusRequestBinaryState() throws Exception {
            byte[] bytes = "\\05FF00FAFF00v\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            cBusOptions = new CBusOptions(false, false, false, false, false, false, false, false, true);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            CBusMessageToServer messageToServer = (CBusMessageToServer) msg;
            RequestCommand requestCommand = (RequestCommand) messageToServer.getRequest();
            CBusCommand cbusCommand = requestCommand.getCbusCommand();
            System.out.println(cbusCommand);
        }

        @Test
        void wat() throws Exception {
            byte[] bytes = "D8FF0024000002000000000000000008000000000000000000\r\n".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            cBusOptions = new CBusOptions(false, false, false, false, false, false, false, false, true);
            requestContext = new RequestContext(true, false, false);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, true, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            CBusMessageToClient messageToClient = (CBusMessageToClient) msg;
            ReplyOrConfirmationReply reply = (ReplyOrConfirmationReply) messageToClient.getReply();
            /*
            MonitoredSALReply monitoredSALReply = (MonitoredSALReply) reply.getReply();
            System.out.println(monitoredSALReply.getMonitoredSAL());
             */
            ReplyCALReply replyReply = (ReplyCALReply) reply.getReply();
            System.out.println(replyReply.getCalReply());
        }

        @Test
        void WriteCommand() throws Exception {
            byte[] bytes = "\\46310900A400410600r\r".getBytes(StandardCharsets.UTF_8);
            ReadBufferByteBased readBufferByteBased = new ReadBufferByteBased(bytes);
            cBusOptions = new CBusOptions(false, false, false, false, false, false, false, false, true);
            CBusMessage msg = CBusMessage.staticParse(readBufferByteBased, false, requestContext, cBusOptions, bytes.length);
            assertThat(msg).isNotNull();
            System.out.println(msg);
            System.out.println(((RequestCommand) ((CBusMessageToServer) msg).getRequest()).getCbusCommand());
        }
    }
}
