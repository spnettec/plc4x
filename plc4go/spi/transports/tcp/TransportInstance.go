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

package tcp

import (
	"bufio"
	"context"
	"fmt"
	"net"

	"github.com/apache/plc4x/plc4go/spi/options"
	transportUtils "github.com/apache/plc4x/plc4go/spi/transports/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type TransportInstance struct {
	transportUtils.DefaultBufferedTransportInstance
	RemoteAddress  *net.TCPAddr
	LocalAddress   *net.TCPAddr
	ConnectTimeout uint32
	transport      *Transport
	tcpConn        net.Conn
	reader         *bufio.Reader

	log zerolog.Logger
}

func NewTcpTransportInstance(remoteAddress *net.TCPAddr, connectTimeout uint32, transport *Transport, _options ...options.WithOption) *TransportInstance {
	transportInstance := &TransportInstance{
		RemoteAddress:  remoteAddress,
		ConnectTimeout: connectTimeout,
		transport:      transport,

		log: options.ExtractCustomLogger(_options...),
	}
	transportInstance.DefaultBufferedTransportInstance = transportUtils.NewDefaultBufferedTransportInstance(transportInstance, _options...)
	return transportInstance
}

func (m *TransportInstance) Connect() error {
	return m.ConnectWithContext(context.Background())
}

func (m *TransportInstance) ConnectWithContext(ctx context.Context) error {
	if m.RemoteAddress == nil {
		return errors.New("Required remote address missing")
	}
	var err error
	var d net.Dialer
	m.tcpConn, err = d.DialContext(ctx, "tcp", m.RemoteAddress.String())
	if err != nil {
		return errors.Wrap(err, "error connecting to remote address")
	}

	m.LocalAddress = m.tcpConn.LocalAddr().(*net.TCPAddr)

	m.reader = bufio.NewReaderSize(m.tcpConn, 100000)

	return nil
}

func (m *TransportInstance) Close() error {
	if m.tcpConn == nil {
		return nil
	}
	err := m.tcpConn.Close()
	if err != nil {
		return errors.Wrap(err, "error closing connection")
	}
	m.tcpConn = nil
	return nil
}

func (m *TransportInstance) IsConnected() bool {
	return m.tcpConn != nil
}

func (m *TransportInstance) Write(data []byte) error {
	if m.tcpConn == nil {
		return errors.New("error writing to transport. No writer available")
	}
	num, err := m.tcpConn.Write(data)
	if err != nil {
		return errors.Wrap(err, "error writing")
	}
	if num != len(data) {
		return errors.New("error writing: not all bytes written")
	}
	return nil
}

func (m *TransportInstance) GetReader() *bufio.Reader {
	return m.reader
}

func (m *TransportInstance) String() string {
	localAddress := ""
	if m.LocalAddress != nil {
		localAddress = m.LocalAddress.String() + "->"
	}
	return fmt.Sprintf("tcp:%s%s", localAddress, m.RemoteAddress)
}
