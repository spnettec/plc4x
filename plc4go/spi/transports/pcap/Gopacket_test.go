//go:build cgo || windows

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

package pcap

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"testing"

	"github.com/gopacket/gopacket"
	"github.com/gopacket/gopacket/layers"
)

func TestGopacketTruncatedGTPv2(t *testing.T) {
	for _, data := range [][]byte{
		{0x40, 1, 0, 3, 0, 0, 1},
		{0x48, 1, 0, 7, 0, 0, 0, 1, 0, 0, 1},
	} {
		t.Run(fmt.Sprintf("header_bytes_%d", len(data)), func(t *testing.T) {
			// Both headers end after the sequence number, before the spare byte.
			// Calling the decoder directly ensures panic recovery cannot hide a regression.
			var gtp layers.GTPv2
			if err := gtp.DecodeFromBytes(data, gopacket.NilDecodeFeedback); err == nil {
				t.Fatal("expected a truncated GTPv2 decode error")
			}

			udp := &layers.UDP{SrcPort: 12345, DstPort: 2123}
			frame := gopacketTestFrame(t, layers.IPProtocolUDP, udp, data)
			// Match TransportInstance.Connect: decode the entire Ethernet frame with Default.
			packet := gopacket.NewPacket(frame, layers.LayerTypeEthernet, gopacket.Default)
			failure := packet.ErrorLayer()
			if failure == nil || !strings.Contains(failure.Error().Error(), "too small for SequenceNumber") {
				t.Fatalf("expected explicit GTPv2 boundary error, got %v", failure)
			}
			if got := packet.Layer(layers.LayerTypeUDP); got == nil || !bytes.Equal(got.LayerPayload(), data) {
				t.Fatal("UDP payload was not preserved after GTPv2 decode error")
			}
		})
	}
}

func TestGopacketReplayPayload(t *testing.T) {
	payload := []byte("plc4x replay payload")
	for _, tcp := range []bool{false, true} {
		var transport gopacket.SerializableLayer = &layers.UDP{SrcPort: 30000, DstPort: 30001}
		protocol, layerType := layers.IPProtocolUDP, layers.LayerTypeUDP
		if tcp {
			transport = &layers.TCP{SrcPort: 30000, DstPort: 30001}
			protocol, layerType = layers.IPProtocolTCP, layers.LayerTypeTCP
		}
		packet := gopacket.NewPacket(gopacketTestFrame(t, protocol, transport, payload), layers.LayerTypeEthernet, gopacket.Default)
		if failure := packet.ErrorLayer(); failure != nil {
			t.Fatalf("unexpected decode error: %v", failure)
		}
		if got := packet.Layer(layerType); got == nil || !bytes.Equal(got.LayerPayload(), payload) {
			t.Fatalf("%v payload was not preserved", layerType)
		}
	}
}

func gopacketTestFrame(t *testing.T, protocol layers.IPProtocol, transport gopacket.SerializableLayer, payload []byte) []byte {
	t.Helper()
	buffer := gopacket.NewSerializeBuffer()
	err := gopacket.SerializeLayers(buffer, gopacket.SerializeOptions{FixLengths: true},
		&layers.Ethernet{SrcMAC: net.HardwareAddr{0, 1, 2, 3, 4, 5}, DstMAC: net.HardwareAddr{6, 7, 8, 9, 10, 11}, EthernetType: layers.EthernetTypeIPv4},
		&layers.IPv4{Version: 4, TTL: 64, Protocol: protocol, SrcIP: net.IPv4(192, 0, 2, 1), DstIP: net.IPv4(192, 0, 2, 2)},
		transport, gopacket.Payload(payload))
	if err != nil {
		t.Fatal(err)
	}
	return buffer.Bytes()
}
