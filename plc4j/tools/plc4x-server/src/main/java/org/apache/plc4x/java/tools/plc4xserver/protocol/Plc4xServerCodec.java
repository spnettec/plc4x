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
package org.apache.plc4x.java.tools.plc4xserver.protocol;

import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import io.netty.channel.ChannelDuplexHandler;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelPromise;
import org.apache.plc4x.java.plc4x.readwrite.Plc4xMessage;
import org.apache.plc4x.java.spi.buffers.api.WithOption;
import org.apache.plc4x.java.spi.buffers.bytebased.ReadBufferByteBased;
import org.apache.plc4x.java.spi.buffers.bytebased.WithByteBasedOption;
import org.apache.plc4x.java.spi.buffers.bytebased.WriteBufferByteBased;

/**
 * Netty codec for the PLC4X proxy protocol.
 * Handles framing (3-byte header: version + 2-byte length) and
 * serialization/deserialization using the new SPI buffer API.
 */
public class Plc4xServerCodec extends ChannelDuplexHandler {

    private static final int HEADER_SIZE = 3;
    private ByteBuf cumulation;

    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
        if (!(msg instanceof ByteBuf)) {
            ctx.fireChannelRead(msg);
            return;
        }
        ByteBuf in = (ByteBuf) msg;
        if (cumulation == null) {
            cumulation = in;
        } else {
            cumulation = Unpooled.wrappedBuffer(cumulation, in);
        }

        while (cumulation.readableBytes() >= HEADER_SIZE) {
            cumulation.markReaderIndex();
            cumulation.readByte(); // version
            int totalLength = cumulation.readUnsignedShort();
            cumulation.resetReaderIndex();

            if (cumulation.readableBytes() < totalLength) {
                break;
            }

            byte[] data = new byte[totalLength];
            cumulation.readBytes(data);
            ReadBufferByteBased readBuffer = new ReadBufferByteBased(data,
                WithOption.WithUnsignedIntegerEncoding("unsigned-binary"),
                WithOption.WithSignedIntegerEncoding("twos-complement"),
                WithOption.WithFloatEncoding("IEEE754"),
                WithOption.WithEncoding("UTF8"),
                WithOption.WithStringEncoding("UTF8"),
                WithByteBasedOption.WithByteOrder("BIG_ENDIAN"));
            Plc4xMessage message = Plc4xMessage.staticParse(readBuffer);
            ctx.fireChannelRead(message);
        }

        if (cumulation != null && !cumulation.isReadable()) {
            cumulation.release();
            cumulation = null;
        }
    }

    @Override
    public void write(ChannelHandlerContext ctx, Object msg, ChannelPromise promise) throws Exception {
        if (msg instanceof Plc4xMessage plc4xMessage) {
            int size = plc4xMessage.getLengthInBytes();
            WriteBufferByteBased writeBuffer = new WriteBufferByteBased(new byte[size],
                WithOption.WithUnsignedIntegerEncoding("unsigned-binary"),
                WithOption.WithSignedIntegerEncoding("twos-complement"),
                WithOption.WithFloatEncoding("IEEE754"),
                WithOption.WithEncoding("UTF8"),
                WithOption.WithStringEncoding("UTF8"),
                WithByteBasedOption.WithByteOrder("BIG_ENDIAN"));
            plc4xMessage.serialize(writeBuffer);
            ctx.write(Unpooled.wrappedBuffer(writeBuffer.getBytes()), promise);
        } else {
            ctx.write(msg, promise);
        }
    }

    @Override
    public void channelInactive(ChannelHandlerContext ctx) throws Exception {
        if (cumulation != null) {
            cumulation.release();
            cumulation = null;
        }
        super.channelInactive(ctx);
    }
}
