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
package org.apache.plc4x.java.s7.readwrite.protocol;

import io.netty.buffer.ByteBuf;
import io.netty.buffer.ByteBufUtil;
import io.netty.buffer.Unpooled;
import org.apache.plc4x.java.api.exceptions.PlcInvalidTagException;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.s7.events.*;
import org.apache.plc4x.java.s7.readwrite.*;
import org.apache.plc4x.java.s7.readwrite.configuration.S7Configuration;
import org.apache.plc4x.java.s7.readwrite.context.S7DriverContext;
import org.apache.plc4x.java.s7.readwrite.optimizer.LargeTagPlcReadRequest;
import org.apache.plc4x.java.s7.readwrite.optimizer.LargeTagPlcWriteRequest;
import org.apache.plc4x.java.s7.readwrite.tag.*;
import org.apache.plc4x.java.s7.utils.S7ParamErrorCode;
import org.apache.plc4x.java.spi.ConversationContext;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.spi.configuration.HasConfiguration;
import org.apache.plc4x.java.spi.connection.PlcTagHandler;
import org.apache.plc4x.java.spi.context.DriverContext;
import org.apache.plc4x.java.spi.generation.*;
import org.apache.plc4x.java.spi.messages.*;
import org.apache.plc4x.java.spi.messages.utils.DefaultPlcResponseItem;
import org.apache.plc4x.java.spi.messages.utils.PlcResponseItem;
import org.apache.plc4x.java.spi.transaction.RequestTransactionManager;
import org.apache.plc4x.java.spi.transaction.TransactionErrorCallback;
import org.apache.plc4x.java.spi.transaction.TransactionTimeOutCallback;
import org.apache.plc4x.java.spi.values.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.nio.charset.Charset;
import java.time.Duration;
import java.time.LocalDateTime;
import java.util.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import static org.apache.plc4x.java.s7.readwrite.optimizer.S7Optimizer.*;
import static org.apache.plc4x.java.spi.connection.AbstractPlcConnection.IS_CONNECTED;

/**
 * The S7 Protocol states that there can not be more then {min(maxAmqCaller, maxAmqCallee} "ongoing" requests. So we
 * need to limit those. Thus, each request goes to a Work Queue and this Queue ensures, that only 3 are open at the same
 * time.
 */
public class S7NonHProtocolLogic extends Plc4xProtocolBase<TPKTPacket> implements HasConfiguration<S7Configuration> {

    private static final Logger logger = LoggerFactory.getLogger(S7NonHProtocolLogic.class);

    private final AtomicInteger tpduGenerator = new AtomicInteger(1);

    private S7Configuration configuration;

    private S7DriverContext s7DriverContext;
    private RequestTransactionManager tm;

    @Override
    public void setDriverContext(DriverContext driverContext) {
        super.setDriverContext(driverContext);
        this.s7DriverContext = (S7DriverContext) driverContext;

        // Initialize Transaction Manager.
        // Until the number of concurrent requests is successfully negotiated we set it to a
        // maximum of only one request being able to be sent at a time. During the login process
        // No concurrent requests can be sent anyway. It will be updated when receiving the
        // S7ParameterSetupCommunication response.
        this.tm = new RequestTransactionManager(1, "S7NonHProtocolLogic");
    }

    @Override
    public void close(ConversationContext<TPKTPacket> context) {
        // TODO: Find out how to close this prior to Java 19
        //clientExecutorService.close();
        tm.shutdown();
    }

    @Override
    public void setConfiguration(S7Configuration configuration) {
        this.configuration = configuration;
    }

    @Override
    public void onConnect(ConversationContext<TPKTPacket> context) {
        if (context.isPassive()) {
            logger.info("S7 Driver running in PASSIVE mode.");
            s7DriverContext.setPassiveMode(true);
            // No login required, just confirm that we're connected.
            context.fireConnected();
            return;
        }

        //Set feature for all handlers in the pipeline from
        //the driver configuration.
        setChannelFeatures();

        // Only the TCP transport supports login.
        logger.info("S7 Driver running in ACTIVE mode.");
        logger.debug("Sending COTP Connection Request");
        // Open the session on ISO Transport Protocol first.
        TPKTPacket packet = new TPKTPacket(
                createCOTPConnectionRequest(s7DriverContext.getCalledTsapId(), s7DriverContext.getCallingTsapId(),
                        s7DriverContext.getCotpTpduSize()));

        context.getChannel().pipeline().names().forEach(s -> {
            logger.debug("Nombre tuberias: " + s);
        });

        context.sendRequest(packet).onTimeout(e -> {
                    logger.info("Timeout during Connection establishing, closing channel...");
                    // TODO: We're saying that we're closing the channel, but not closing the channel ... sure, this is what we want?
                    //context.getChannel().close();
                }).expectResponse(TPKTPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
                .unwrap(TPKTPacket::getPayload).only(COTPPacketConnectionResponse.class)
                .handle(cotpPacketConnectionResponse -> {
                    logger.debug("Got COTP Connection Response");
                    logger.debug("Sending S7 Connection Request");
                    context.sendRequest(createS7ConnectionRequest(cotpPacketConnectionResponse)).onTimeout(e -> {
                                logger.warn("Timeout during Connection establishing, closing channel...");
                                context.getChannel().close();
                            }).expectResponse(TPKTPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
                            .unwrap(TPKTPacket::getPayload).only(COTPPacketData.class).unwrap(COTPPacket::getPayload)
                            .only(S7MessageResponseData.class).unwrap(S7Message::getParameter)
                            .only(S7ParameterSetupCommunication.class).handle(setupCommunication -> {
                                logger.debug("Got S7 Connection Response");
                                // Save some data from the response.
                                s7DriverContext.setMaxAmqCaller(setupCommunication.getMaxAmqCaller());
                                s7DriverContext.setMaxAmqCallee(setupCommunication.getMaxAmqCallee());
                                s7DriverContext.setPduSize(setupCommunication.getPduLength());

                                // Update the number of concurrent requests to the negotiated number.
                                // I have never seen anything else than equal values for caller and
                                // callee, but if they were different, we're only limiting the outgoing
                                // requests.
                                tm.setNumberOfConcurrentRequests(s7DriverContext.getMaxAmqCallee());

                                // If the controller type is explicitly set, were finished with the login
                                // process. If it's set to ANY, we have to query the serial number information
                                // in order to detect the type of PLC.
                                if (s7DriverContext.getControllerType() != ControllerType.ANY) {
                                    // Send an event that connection setup is complete.
                                    context.fireConnected();
                                    return;
                                }

                                // Prepare a message to request the remote to identify itself.
                                logger.debug("Sending S7 Identification Request");
                                TPKTPacket tpktPacket = createIdentifyRemoteMessage();
                                context.sendRequest(tpktPacket).onTimeout(e -> {
                                            logger.warn("Timeout during Connection establishing, closing channel...");
                                            context.getChannel().close();
                                        }).expectResponse(TPKTPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
                                        .unwrap(TPKTPacket::getPayload).only(COTPPacketData.class)
                                        .unwrap(COTPPacketData::getPayload).only(S7MessageUserData.class)
                                        .unwrap(S7MessageUserData::getPayload).only(S7PayloadUserData.class)
                                        .handle(payloadUserData -> {
                                            logger.debug("Got S7 Identification Response");
                                            extractControllerTypeAndFireConnected(context, payloadUserData);
                                        });
                            });
                });
    }

    /*
     * It performs the sequential and safe shutdown of the driver.
     * Completion of pending requests, executors and associated tasks.
     */
    @Override
    public void onDisconnect(ConversationContext<TPKTPacket> context) {
        tm.shutdown();
    }

    @Override
    public PlcTagHandler getTagHandler() {
        return new S7PlcTagHandler();
    }

    @Override
    public CompletableFuture<PlcReadResponse> read(PlcReadRequest readRequest) {
        // If we're not connected, just abort with an error.
        if (!isConnected()) {
            CompletableFuture<PlcReadResponse> future = new CompletableFuture<>();
            future.completeExceptionally(new PlcRuntimeException("Disconnected"));
            return future;
        }
        if (readRequest instanceof LargeTagPlcReadRequest) {
            final LargeTagPlcReadRequest largeTagPlcReadRequest = (LargeTagPlcReadRequest) readRequest;
            return toLargePlcReadResponse(largeTagPlcReadRequest);
        }
        DefaultPlcReadRequest request = (DefaultPlcReadRequest) readRequest;
        CompletableFuture<S7Message> responseFuture;
        if (request.getTagNames().stream().anyMatch(t -> request.getTag(t) instanceof S7SzlTag)) {
            // TODO: Is it correct, that there can only be one szl tag?
            S7SzlTag szlTag = (S7SzlTag) request.getTags().get(0);
            S7Message s7Message = new S7MessageUserData(getTpduId(), new S7ParameterUserData(
                    List.of(new S7ParameterUserDataItemCPUFunctions((short) 0x11, (byte) 0x4, (byte) 0x4, (short) 0x01,
                            (short) 0x00, null, null, null))), new S7PayloadUserData(
                    List.of(new S7PayloadUserDataItemCpuFunctionReadSzlRequest(DataTransportErrorCode.OK,
                            DataTransportSize.OCTET_STRING, 0x04,
                            new SzlId(SzlModuleTypeClass.enumForValue((byte) ((szlTag.getSzlId() & 0xf000) >> 12)),
                                    (byte) ((szlTag.getSzlId() & 0x0f00) >> 8),
                                    SzlSublist.enumForValue((short) (szlTag.getSzlId() & 0x00ff))),
                            szlTag.getIndex()))));
            responseFuture = sendInternal(s7Message);
        } else if (request.getTagNames().stream().anyMatch(t -> request.getTag(t) instanceof S7ClkTag)) {
            responseFuture = performClkRequest(request);
        }

        // If the request contains at least one var-length string field, we need to get the real length first.
        //else if (request.getTagNames().stream().anyMatch(t -> request.getTag(t) instanceof S7StringVarLengthTag)) {
        //	responseFuture = performVarLengthStringReadRequest(request);
        //}

        // This is a "normal" read request.
        else {
            responseFuture = performOrdinaryReadRequest(request);
        }

        // Just send a single response and chain it as Response
        return toPlcReadResponse(readRequest, responseFuture);
    }

    public <T> CompletableFuture<List<T>> allOf(List<CompletableFuture<T>> futuresList) {
        CompletableFuture<Void> allFuturesResult = CompletableFuture.allOf(
                futuresList.toArray(CompletableFuture[]::new));
        return allFuturesResult.thenApply(
                v -> futuresList.stream().map(CompletableFuture::join).collect(Collectors.<T> toList()));
    }

    private CompletableFuture<PlcReadResponse> toLargePlcReadResponse(LargeTagPlcReadRequest readRequest) {
        List<CompletableFuture<S7Message>> response = readLargeInternal(readRequest);
        return allOf(response).thenApply(value -> {
            try {
                return (PlcReadResponse) decodeLargeReadResponse(value, readRequest);
            } catch (PlcProtocolException e) {
                throw new RuntimeException(e);
            }
        });
    }

    private CompletableFuture<PlcWriteResponse> toLargePlcWriteResponse(LargeTagPlcWriteRequest writeRequest) {
        List<CompletableFuture<S7Message>> response = writeLargeInternal(writeRequest);
        return allOf(response).thenApply(value -> {
            try {
                return (PlcWriteResponse) decodeLargeWriteResponse(value, writeRequest);
            } catch (PlcProtocolException e) {
                throw new RuntimeException(e);
            }
        });
    }

    /**
     * Maps the S7ReadResponse of a PlcReadRequest to a PlcReadResponse
     */
    private CompletableFuture<PlcReadResponse> toPlcReadResponse(PlcReadRequest readRequest,
            CompletableFuture<S7Message> responseFuture) {
        CompletableFuture<PlcReadResponse> clientFuture = new CompletableFuture<>();

        responseFuture.whenComplete((s7Message, throwable) -> {
            if (throwable != null) {
                clientFuture.completeExceptionally(new PlcProtocolException("Error reading", throwable));
            } else {
                try {
                    PlcReadResponse response = (PlcReadResponse) decodeReadResponse(s7Message, readRequest);
                    clientFuture.complete(response);
                } catch (Exception ex) {
                    logger.info(ex.toString());
                }
            }
        });

        return clientFuture;
    }

    @Override
    public CompletableFuture<PlcWriteResponse> write(PlcWriteRequest writeRequest) {
        // If we're not connected, just abort with an error.
        if (!isConnected()) {
            CompletableFuture<PlcWriteResponse> future = new CompletableFuture<>();
            future.completeExceptionally(new PlcRuntimeException("Disconnected"));
            return future;
        }

        CompletableFuture<S7Message> responseFuture;
        // TODO: Write one or two lines on what happens here ... to me it looks as if there's at least on S7ClkTag, then all is handled by the writeClk method, but what happens if a request would contain mixed tag types?
        if (writeRequest.getTagNames().stream().anyMatch(t -> writeRequest.getTag(t) instanceof S7ClkTag)) {
            responseFuture = performClkSetRequest(writeRequest);
        } else {
            if (writeRequest instanceof LargeTagPlcWriteRequest) {
                final LargeTagPlcWriteRequest largeTagPlcWriteRequest = (LargeTagPlcWriteRequest) writeRequest;
                return toLargePlcWriteResponse(largeTagPlcWriteRequest);
            } else {
                responseFuture = performOrdinaryWriteRequest(writeRequest);
            }
        }

        return toPlcWriteResponse(writeRequest, responseFuture);
    }

    private CompletableFuture<PlcWriteResponse> toPlcWriteResponse(PlcWriteRequest writeRequest,
            CompletableFuture<S7Message> responseFuture) {
        CompletableFuture<PlcWriteResponse> clientFuture = new CompletableFuture<>();

        responseFuture.whenComplete((s7Message, throwable) -> {
            if (throwable != null) {
                Map<String, PlcResponseCode> responses = new HashMap<>();
                for (String tagName : writeRequest.getTagNames()) {
                    responses.put(tagName, PlcResponseCode.INTERNAL_ERROR);
                }
                clientFuture.complete(new DefaultPlcWriteResponse(writeRequest, responses));
                //clientFuture.completeExceptionally(new PlcProtocolException("Error writing", throwable));
            } else {
                try {
                    PlcWriteResponse response = (PlcWriteResponse) decodeWriteResponse(s7Message, writeRequest);
                    clientFuture.complete(response);
                } catch (Exception ex) {
                    logger.info(ex.toString());
                }
            }
        });

        return clientFuture;
    }


    private List<CompletableFuture<S7Message>> readLargeInternal(LargeTagPlcReadRequest largeTagPlcReadRequest) {
        List<CompletableFuture<S7Message>> futures = new ArrayList<>();
        if (!isConnected()) {
            CompletableFuture<S7Message> future = new CompletableFuture<>();
            futures.add(future);
            future.completeExceptionally(new PlcRuntimeException("Disconnected"));
            return futures;
        }

        List<S7Message> messages = splitLargeTagReadVarParameter(largeTagPlcReadRequest,
                this.s7DriverContext.getPduSize());
        messages.forEach(message -> futures.add(sendInternal(message)));
        return futures;
    }

    private List<CompletableFuture<S7Message>> writeLargeInternal(LargeTagPlcWriteRequest largeTagPlcWriteRequest) {
        List<CompletableFuture<S7Message>> futures = new ArrayList<>();
        if (!isConnected()) {
            CompletableFuture<S7Message> future = new CompletableFuture<>();
            futures.add(future);
            future.completeExceptionally(new PlcRuntimeException("Disconnected"));
            return futures;
        }

        List<S7Message> messages = splitLargeTagWriteVarParameter(largeTagPlcWriteRequest,
                this.s7DriverContext.getPduSize());
        messages.forEach(message -> futures.add(sendInternal(message)));
        return futures;
    }

    private List<S7Message> splitLargeTagReadVarParameter(LargeTagPlcReadRequest largeTagPlcReadRequest, int pduSize) {
        List<S7Message> result = new LinkedList<>();
        int maxResponseSize = pduSize - EMPTY_READ_RESPONSE_SIZE - 4;
        final S7AddressAny address = (S7AddressAny) encodeS7Address(largeTagPlcReadRequest.getTag());
        // Calculate the sizes in the request and response adding this item to the current request would add.
        int readRequestItemSize = S7_ADDRESS_ANY_SIZE;
        // Constant size of the parameter item in the response (0 bytes) + Constant size of the payload item +
        // payload data size.
        int readResponseItemSize = 4 + (address.getNumberOfElements() * address.getTransportSize().getSizeInBytes());
        // If it's an odd number of bytes, add one to make it even
        if (readResponseItemSize % 2 == 1) {
            readResponseItemSize++;
        }

        // If the item would not fit into a separate message, we have to split it.
        if (((EMPTY_READ_REQUEST_SIZE + readRequestItemSize) > pduSize) || (
                EMPTY_READ_RESPONSE_SIZE + readResponseItemSize > pduSize)) {
            // Create a new sub message.
            S7Message subMessage;
            // Calculate the maximum number of items that would fit in a single request.
            int maxNumElements = (int) Math.floor(
                    (double) maxResponseSize / (double) address.getTransportSize().getSizeInBytes());
            int sizeMaxNumElementInBytes = maxNumElements * address.getTransportSize().getSizeInBytes();

            // Initialize the loop with the total number of elements and the original address.
            int remainingNumElements = address.getNumberOfElements();
            int curByteAddress = address.getByteAddress();

            // Keep on adding chunks of the original address until all have been added.
            while (remainingNumElements > 0) {
                int numCurElements = Math.min(remainingNumElements, maxNumElements);
                S7VarRequestParameterItemAddress subVarParameterItem = new S7VarRequestParameterItemAddress(
                        new S7AddressAny(address.getTransportSize(), numCurElements, address.getDbNumber(),
                                address.getArea(), curByteAddress, address.getBitAddress()));

                subMessage = new S7MessageRequest(getTpduId(),
                        new S7ParameterReadVarRequest(Collections.singletonList(subVarParameterItem)), null);

                result.add(subMessage);

                remainingNumElements -= maxNumElements;
                curByteAddress += sizeMaxNumElementInBytes;
            }
        } else {
            logger.warn("Read Var Request size is too small, this is not large enough");

            S7Message subMessage = new S7MessageRequest(getTpduId(), new S7ParameterReadVarRequest(
                    Collections.singletonList(new S7VarRequestParameterItemAddress(address))), null);

            result.add(subMessage);
        }

        return result;
    }

    private List<S7Message> splitLargeTagWriteVarParameter(LargeTagPlcWriteRequest largeTagPlcWriteRequest,
            int pduSize) {
        List<S7Message> result = new LinkedList<>();
        int maxRequestSize = pduSize - S7_ADDRESS_ANY_SIZE - EMPTY_WRITE_REQUEST_SIZE - 4;
        final S7AddressAny address = (S7AddressAny) encodeS7Address(largeTagPlcWriteRequest.getTag());
        // Calculate the sizes in the request and response adding this item to the current request would add.
        //int writeRequestItemSize = address.getLengthInBytes();
        // Constant size of the parameter item in the response (0 bytes) + Constant size of the payload item +
        // payload data size.
        int writeRequestItemSize = S7_ADDRESS_ANY_SIZE + 4 + (address.getNumberOfElements() * address.getTransportSize().getSizeInBytes());
        // If it's an odd number of bytes, add one to make it even
        if (writeRequestItemSize % 2 == 1) {
            writeRequestItemSize++;
        }
        int writeResponseItemSize = 4;
        // If the item would not fit into a separate message, we have to split it.
        if (((EMPTY_WRITE_REQUEST_SIZE + writeRequestItemSize) > pduSize) || (
                EMPTY_WRITE_RESPONSE_SIZE + writeResponseItemSize > pduSize)) {
            // Create a new sub message.
            S7Message subMessage;
            // Calculate the maximum number of items that would fit in a single request.
            int maxNumElements = (int) Math.floor(
                    (double) maxRequestSize / (double) address.getTransportSize().getSizeInBytes());

            // Initialize the loop with the total number of elements and the original address.
            int remainingNumElements = address.getNumberOfElements();
            int curByteAddress = address.getByteAddress();
            PlcValue plcValue = largeTagPlcWriteRequest.getPlcValue();
            WriteBufferByteBased writeBufferByteBased = serializePlcValueToWriteBuffer(largeTagPlcWriteRequest.getTag(), plcValue);
            if (writeBufferByteBased == null) {
                throw new PlcRuntimeException("writeBufferByteBased is null");
            }
            // Keep on adding chunks of the original address until all have been added.
            while (remainingNumElements > 0) {
                int numCurElements = Math.min(remainingNumElements, maxNumElements);
                S7VarRequestParameterItemAddress subVarParameterItem = new S7VarRequestParameterItemAddress(
                        new S7AddressAny(address.getTransportSize(), numCurElements, address.getDbNumber(),
                                address.getArea(), curByteAddress, address.getBitAddress()));
                int start = curByteAddress - address.getByteAddress();
                byte[] bytes = writeBufferByteBased.getBytes(start, start + numCurElements);
                subMessage = new S7MessageRequest(getTpduId(),
                        new S7ParameterWriteVarRequest(Collections.singletonList(subVarParameterItem)),
                        new S7PayloadWriteVarRequest(Collections.singletonList(
                                new S7VarPayloadDataItem(DataTransportErrorCode.OK,
                                        address.getTransportSize().getDataTransportSize(), bytes))));

                result.add(subMessage);

                remainingNumElements -= numCurElements;
                curByteAddress += numCurElements;
            }
        } else {
            logger.warn("Write Var Request size is too small, this is not large enough");

            S7Message subMessage = new S7MessageRequest(getTpduId(), new S7ParameterWriteVarRequest(
                    Collections.singletonList(new S7VarRequestParameterItemAddress(address))),
                    new S7PayloadWriteVarRequest(Collections.singletonList(
                            serializePlcValue(largeTagPlcWriteRequest.getTag(),largeTagPlcWriteRequest.getPlcValue()))));

            result.add(subMessage);
        }

        return result;
    }

    private CompletableFuture<S7Message> performClkRequest(DefaultPlcReadRequest request) {
        List<S7ParameterUserDataItem> parameterItems = new ArrayList<>(request.getNumberOfTags());
        List<S7PayloadUserDataItem> payloadItems = new ArrayList<>(request.getNumberOfTags());

        final S7ClkTag tag = (S7ClkTag) request.getTags().get(0);
        int subFunction = tag.getAddressString().equals("CLK") ? 1 : 3;

        S7ParameterUserDataItemCPUFunctions parameter = new S7ParameterUserDataItemCPUFunctions((short) 0x11,   //Method
                (byte) 0x04,    //FunctionType
                (byte) 0x07,    //FunctionGroup
                (short) subFunction,   //SubFunction
                (short) 0x00,   //SequenceNumber
                null,   //DataUnitReferenceNumber
                null,   //LastDataUnit
                null         //errorCode
        );
        parameterItems.add(parameter);

        S7PayloadUserDataItemClkRequest payload;
        payload = new S7PayloadUserDataItemClkRequest(DataTransportErrorCode.NOT_FOUND, DataTransportSize.NULL, 0x00);
        payloadItems.add(payload);

        return sendInternal(new S7MessageUserData(getTpduId(), new S7ParameterUserData(parameterItems),
                new S7PayloadUserData(payloadItems)));
    }

    /*
     *
     */
    private CompletableFuture<S7Message> performClkSetRequest(PlcWriteRequest request) {
        List<S7ParameterUserDataItem> parameterItems = new ArrayList<>(request.getNumberOfTags());
        List<S7PayloadUserDataItem> payloadItems = new ArrayList<>(request.getNumberOfTags());

        S7ParameterUserDataItemCPUFunctions parameter = new S7ParameterUserDataItemCPUFunctions((short) 0x11,   //Method
                (byte) 0x04,    //FunctionType
                (byte) 0x07,    //FunctionGroup
                (short) 0x04,   //SubFunction
                (short) 0x00,   //SequenceNumber
                null,   //DataUnitReferenceNumber
                null,   //LastDataUnit
                null         //errorCode
        );
        parameterItems.add(parameter);

        S7ClkTag tag = (S7ClkTag) request.getTags().get(0);

        S7PayloadUserDataItemClkSetRequest payload;
        payload = new S7PayloadUserDataItemClkSetRequest(DataTransportErrorCode.OK, DataTransportSize.OCTET_STRING,
                0x0A, tag.getDateAndTime());
        payloadItems.add(payload);

        return sendInternal(new S7MessageUserData(getTpduId(), new S7ParameterUserData(parameterItems),
                new S7PayloadUserData(payloadItems)));
    }

    private CompletableFuture<S7Message> performOrdinaryReadRequest(DefaultPlcReadRequest request) {
        // Convert each tag in the request into a corresponding item used in the S7 protocol.
        List<S7VarRequestParameterItem> requestItems = new ArrayList<>(request.getNumberOfTags());
        for (PlcTag tag : request.getTags()) {
            requestItems.add(new S7VarRequestParameterItemAddress(encodeS7Address(tag)));
        }

        // Create a read request template.
        // tpuId will be inserted before sending in #readInternal, so we insert -1 as dummy here
        S7Message requestMessage = new S7MessageRequest(getTpduId(), new S7ParameterReadVarRequest(requestItems), null);

        return sendInternal(requestMessage);
    }

    private CompletableFuture<S7Message> performOrdinaryWriteRequest(PlcWriteRequest request) {
        List<S7VarRequestParameterItem> parameterItems = new ArrayList<>(request.getNumberOfTags());
        List<S7VarPayloadDataItem> payloadItems = new ArrayList<>(request.getNumberOfTags());

        for (String tagName : request.getTagNames()) {
            final S7Tag tag = (S7Tag) request.getTag(tagName);
            final PlcValue plcValue = request.getPlcValue(tagName);
            parameterItems.add(new S7VarRequestParameterItemAddress(encodeS7Address(tag)));
            payloadItems.add(serializePlcValue(tag, plcValue));
        }

        return sendInternal(new S7MessageRequest(getTpduId(), new S7ParameterWriteVarRequest(parameterItems),
                new S7PayloadWriteVarRequest(payloadItems)));
    }

    /**
     * Sends one Read over the Wire and internally returns the Response Do sending of normally sized single-message
     * request.
     * <p>
     * Assumes that the {@link S7MessageRequest} and its expected {@link S7MessageResponseData} and does not further
     * check that!
     */
    private CompletableFuture<S7Message> sendInternal(S7Message request) {
        CompletableFuture<S7Message> future = new CompletableFuture<>();

        // Get the tpduId from the S7 message.
        int tpduId = request.getTpduReference();

        TPKTPacket tpktPacket = new TPKTPacket(new COTPPacketData(null, request, true, (byte) tpduId));

        // Start a new request-transaction (Is ended in the response-handler)
        RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
        // Send the request.
        transaction.submit(() -> conversationContext.sendRequest(tpktPacket)
                .onTimeout(new TransactionTimeOutCallback<>(future, transaction, conversationContext.getChannel()))
                .onError(new TransactionErrorCallback<>(future, transaction, conversationContext.getChannel()))
                .expectResponse(TPKTPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
                .unwrap(TPKTPacket::getPayload).only(COTPPacketData.class).check(p -> p.getPayload() != null)
                .unwrap(COTPPacket::getPayload).check(p -> p.getTpduReference() == tpduId).handle(p -> {
                    future.complete(p);
                    // Finish the request-transaction.
                    transaction.endRequest();
                }));

        return future;
    }

    /**
     * DECODE: This method is called when there is no handler for the message. By default it must correspond to
     * asynchronous events, which if so, must be transferred to the event queue.
     *
     * The event's own information is encapsulated in the parameters and payload field. From this it is abstracted to
     * the corresponding event model.
     *
     * 01. S7ModeEvent: 02. S7UserEvent: 03. S7SysEvent: 04. S7AlarmEvent 05. S7CyclicEvent: 06. S7CyclicEvent:
     *
     * TODO: Use mspec to generate types that allow better interpretation of
     * the code using "instanceof".
     */
    @Override
    protected void decode(ConversationContext<TPKTPacket> context, TPKTPacket msg) throws Exception {

        final S7Message s7msg = msg.getPayload().getPayload();
        final S7Parameter parameter = s7msg.getParameter();
        final S7PayloadUserData payload = (S7PayloadUserData) s7msg.getPayload();

        if (parameter instanceof S7ParameterModeTransition) {  //(01)

            S7ModeEvent modeEvent = new S7ModeEvent((S7ParameterModeTransition) parameter);
            //eventQueue.add(modeEvent);

        } else if (parameter instanceof S7ParameterUserData) {

            S7ParameterUserData parameterUD = (S7ParameterUserData) parameter;
            List<S7ParameterUserDataItem> parameterUDItems = parameterUD.getItems();

            for (S7ParameterUserDataItem parameterUDItem : parameterUDItems) {

                if (parameterUDItem instanceof S7ParameterUserDataItemCPUFunctions) {

                    S7ParameterUserDataItemCPUFunctions myParameter = (S7ParameterUserDataItemCPUFunctions) parameterUDItem;

                    if ((myParameter.getCpuFunctionType() == 0x00) && (myParameter.getCpuSubfunction()
                            == 0x03)) { //(02)

                        payload.getItems().forEach(item -> {
                            if (item instanceof S7PayloadDiagnosticMessage) {
                                final S7PayloadDiagnosticMessage pload = (S7PayloadDiagnosticMessage) item;
                                if ((pload.getEventId() >= 0x0A000) & (pload.getEventId() <= 0x0BFFF)) {
                                    S7UserEvent userEvent = S7UserEvent.of(pload);
                                    //eventQueue.add(userEvent);
                                } else {
                                    S7SysEvent sysEvent = S7SysEvent.of(pload);
                                    //eventQueue.add(sysEvent);
                                }
                            }
                        });

                    } else if ((myParameter.getCpuFunctionType() == 0x00) && ((myParameter.getCpuSubfunction() == 0x05)
                            || (myParameter.getCpuSubfunction() == 0x06) || (myParameter.getCpuSubfunction() == 0x0c)
                            || (myParameter.getCpuSubfunction() == 0x11) || (myParameter.getCpuSubfunction() == 0x12)
                            || (myParameter.getCpuSubfunction() == 0x13) || (myParameter.getCpuSubfunction()
                            == 0x16))) { //(04)

                        payload.getItems().forEach(item -> {
                            S7AlarmEvent alrmEvent = S7AlarmEvent.of(item);
                            //eventQueue.add(alrmEvent);
                        });

                    } else if ((myParameter.getCpuFunctionType() == 0x00) && (myParameter.getCpuSubfunction()
                            == 0x13)) {
                        //TODO: Requires reverse engineering.
                    } else if ((myParameter.getCpuFunctionGroup() == 0x02) && (myParameter.getCpuFunctionType() == 0x00)
                            && (myParameter.getCpuSubfunction() == 0x01)) { //(05)

                        S7ParameterUserDataItemCPUFunctions parameterItem = (S7ParameterUserDataItemCPUFunctions) ((S7ParameterUserData) parameter).getItems()
                                .get(0);

                        S7PayloadUserDataItemCyclicServicesPush payloadItem = (S7PayloadUserDataItemCyclicServicesPush) payload.getItems()
                                .get(0);

                        //S7CyclicEvent cycEvent = new S7CyclicEvent(cycRequests.get(parameterItem.getSequenceNumber()),
                        //		parameterItem.getSequenceNumber(),
                        //		payloadItem);
                        //eventQueue.add(cycEvent);

                    } else if ((myParameter.getCpuFunctionGroup() == 0x02) && (myParameter.getCpuFunctionType() == 0x00)
                            && (myParameter.getCpuSubfunction() == 0x05)) { //(06)

                        S7ParameterUserDataItemCPUFunctions parameterItem = (S7ParameterUserDataItemCPUFunctions) ((S7ParameterUserData) parameter).getItems()
                                .get(0);

                        S7PayloadUserDataItemCyclicServicesChangeDrivenPush payloadItem = (S7PayloadUserDataItemCyclicServicesChangeDrivenPush) payload.getItems()
                                .get(0);

                        S7CyclicEvent cycEvent = new S7CyclicEvent(null, parameterItem.getSequenceNumber(),
                                payloadItem);
                        //eventQueue.add(cycEvent);

                    } else if ((myParameter.getCpuFunctionType() == 0x08) && (myParameter.getCpuSubfunction()
                            == 0x01)) {
                        //TODO: Requires reverse engineering.
                    } else if ((myParameter.getCpuFunctionType() == 0x08) && (myParameter.getCpuSubfunction()
                            == 0x04)) {
                        //TODO: Requires reverse engineering.
                    }
                }
            }
        }
    }

    private void extractControllerTypeAndFireConnected(ConversationContext<TPKTPacket> context,
            S7PayloadUserData payloadUserData) {
        for (S7PayloadUserDataItem item : payloadUserData.getItems()) {
            if (!(item instanceof S7PayloadUserDataItemCpuFunctionReadSzlResponse)) {
                continue;
            }
            S7PayloadUserDataItemCpuFunctionReadSzlResponse readSzlResponseItem = (S7PayloadUserDataItemCpuFunctionReadSzlResponse) item;

            //            for (SzlDataTreeItem readSzlResponseItemItem : readSzlResponseItem.getItems()) {
            //                if (readSzlResponseItemItem.getItemIndex() != 0x0001) {
            //                    continue;
            //                }
            //                final String articleNumber = new String(readSzlResponseItemItem.getMlfb());
            //                s7DriverContext.setControllerType(decodeControllerType(articleNumber));
            //
            //                // Send an event that connection setup is complete.
            //                context.fireConnected();
            //            }
            ByteBuf szlItem = Unpooled.wrappedBuffer(readSzlResponseItem.getItems());
            String articleNumber = szlItem.toString(2, 20, Charset.defaultCharset());
            s7DriverContext.setControllerType(decodeControllerType(articleNumber));
            context.fireConnected();
        }
    }

    private TPKTPacket createIdentifyRemoteMessage() {
        S7MessageUserData identifyRemoteMessage = new S7MessageUserData(1, new S7ParameterUserData(
                Collections.singletonList(
                        new S7ParameterUserDataItemCPUFunctions((short) 0x11, (byte) 0x4, (byte) 0x4, (short) 0x01,
                                (short) 0x00, null, null, null))), new S7PayloadUserData(Collections.singletonList(
                new S7PayloadUserDataItemCpuFunctionReadSzlRequest(DataTransportErrorCode.OK,
                        DataTransportSize.OCTET_STRING, 0x0C,
                        new SzlId(SzlModuleTypeClass.CPU, (byte) 0x00, SzlSublist.MODULE_IDENTIFICATION), 0x0000))));
        COTPPacketData cotpPacketData = new COTPPacketData(null, identifyRemoteMessage, true, (byte) 2);
        return new TPKTPacket(cotpPacketData);
    }

    private TPKTPacket createS7ConnectionRequest(COTPPacketConnectionResponse cotpPacketConnectionResponse) {
        for (COTPParameter parameter : cotpPacketConnectionResponse.getParameters()) {
            if (parameter instanceof COTPParameterCalledTsap) {
                COTPParameterCalledTsap cotpParameterCalledTsap = (COTPParameterCalledTsap) parameter;
                s7DriverContext.setCalledTsapId(cotpParameterCalledTsap.getTsapId());
            } else if (parameter instanceof COTPParameterCallingTsap) {
                COTPParameterCallingTsap cotpParameterCallingTsap = (COTPParameterCallingTsap) parameter;
                if (cotpParameterCallingTsap.getTsapId() != s7DriverContext.getCallingTsapId()) {
                    s7DriverContext.setCallingTsapId(cotpParameterCallingTsap.getTsapId());
                    logger.warn("Switching calling TSAP id to '{}'", s7DriverContext.getCallingTsapId());
                }
            } else if (parameter instanceof COTPParameterTpduSize) {
                COTPParameterTpduSize cotpParameterTpduSize = (COTPParameterTpduSize) parameter;
                s7DriverContext.setCotpTpduSize(cotpParameterTpduSize.getTpduSize());
            } else {
                logger.warn("Got unknown parameter type '{}'", parameter.getClass().getName());
            }
        }

        // Send an S7 login message.
        S7ParameterSetupCommunication s7ParameterSetupCommunication = new S7ParameterSetupCommunication(
                s7DriverContext.getMaxAmqCaller(), s7DriverContext.getMaxAmqCallee(), s7DriverContext.getPduSize());
        S7Message s7Message = new S7MessageRequest(0, s7ParameterSetupCommunication, null);
        int tpduId = 1;
        if (this.s7DriverContext.getControllerType() == ControllerType.S7_200) {
            tpduId = 0;
        }
        COTPPacketData cotpPacketData = new COTPPacketData(null, s7Message, true, (byte) tpduId);
        return new TPKTPacket(cotpPacketData);
    }

    private COTPPacketConnectionRequest createCOTPConnectionRequest(int calledTsapId, int callingTsapId,
            COTPTpduSize cotpTpduSize) {
        return new COTPPacketConnectionRequest(
                Arrays.asList(new COTPParameterCallingTsap(callingTsapId), new COTPParameterCalledTsap(calledTsapId),
                        new COTPParameterTpduSize(cotpTpduSize)), null, (short) 0x0000, (short) 0x000F,
                COTPProtocolClass.CLASS_0);
    }

    private PlcResponse decodeLargeReadResponse(List<S7Message> responseMessages, LargeTagPlcReadRequest plcReadRequest)
            throws PlcProtocolException {
        Map<String, PlcResponseItem<PlcValue>> values = new HashMap<>();
        short errorClass;
        short errorCode;
        S7ParameterUserDataItemCPUFunctions parameteritem;
        String tagName = plcReadRequest.getTagName();
        S7Tag tag = (S7Tag) plcReadRequest.getTag();
        ByteBuf data = Unpooled.buffer();
        try {
            PlcResponseCode responseCode = null;
            for (S7Message responseMessage : responseMessages) {
                if (responseMessage instanceof S7MessageResponseData) {
                    S7MessageResponseData messageResponseData = (S7MessageResponseData) responseMessage;
                    errorClass = messageResponseData.getErrorClass();
                    errorCode = messageResponseData.getErrorCode();
                } else if (responseMessage instanceof S7MessageResponse) {
                    S7MessageResponse messageResponse = (S7MessageResponse) responseMessage;
                    errorClass = messageResponse.getErrorClass();
                    errorCode = messageResponse.getErrorCode();
                } else if (responseMessage instanceof S7MessageUserData) {
                    S7MessageUserData messageResponse = (S7MessageUserData) responseMessage;
                    S7ParameterUserData parameters = (S7ParameterUserData) messageResponse.getParameter();
                    parameteritem = (S7ParameterUserDataItemCPUFunctions) parameters.getItems().get(0);
                    errorClass = 0;
                    errorCode = parameteritem.getErrorCode().shortValue();
                } else {
                    throw new PlcProtocolException("Unsupported message type " + responseMessage.getClass().getName());
                }
                if ((errorClass != 0) || (errorCode != 0)) {
                    if ((errorClass == 129) && (errorCode == 4)) {
                        logger.warn("Got an error response from the PLC. This particular response code usually indicates "
                                + "that PUT/GET is not enabled on the PLC.");
                        PlcResponseItem<PlcValue> result = new DefaultPlcResponseItem<>(PlcResponseCode.ACCESS_DENIED,
                                new PlcNull());
                        values.put(tagName, result);
                        return new DefaultPlcReadResponse(plcReadRequest, values);
                    } else {
                        logger.warn("Got an unknown error response from the PLC. Error Class: {}, Error Code {}. "
                                + "We probably need to implement explicit handling for this, so please file a bug-report "
                                + "on https://issues.apache.org/jira/projects/PLC4X and ideally attach a WireShark dump "
                                + "containing a capture of the communication.", errorClass, errorCode);
                        PlcResponseItem<PlcValue> result = new DefaultPlcResponseItem<>(PlcResponseCode.INTERNAL_ERROR,
                                new PlcNull());
                        values.put(tagName, result);
                        return new DefaultPlcReadResponse(plcReadRequest, values);
                    }
                }
                S7PayloadReadVarResponse payload = (S7PayloadReadVarResponse) responseMessage.getPayload();

                S7VarPayloadDataItem payloadItem = payload.getItems().get(0);

                responseCode = decodeResponseCode(payloadItem.getReturnCode());
                data.writeBytes(payloadItem.getData());
                if (responseCode != PlcResponseCode.OK) {
                    break;
                }
            }
            PlcValue plcValue;
            try {
                plcValue = parsePlcValue(tag, data.array());
            } catch (Exception e) {
                throw new PlcProtocolException("Error decoding PlcValue", e);
            }
            PlcResponseItem<PlcValue> result = new DefaultPlcResponseItem<>(responseCode, plcValue);
            values.put(tagName, result);
            return new DefaultPlcReadResponse(plcReadRequest, values);
        } finally {
            data.release();
        }
    }

    private PlcResponse decodeLargeWriteResponse(List<S7Message> responseMessages,
            LargeTagPlcWriteRequest plcWriteRequest) throws PlcProtocolException {
        Map<String, PlcResponseCode> values = new HashMap<>();
        short errorClass;
        short errorCode;
        S7ParameterUserDataItemCPUFunctions parameteritem;
        String tagName = plcWriteRequest.getTagName();
        PlcResponseCode responseCode = null;
        for (S7Message responseMessage : responseMessages) {
            if (responseMessage instanceof S7MessageResponseData) {
                S7MessageResponseData messageResponseData = (S7MessageResponseData) responseMessage;
                errorClass = messageResponseData.getErrorClass();
                errorCode = messageResponseData.getErrorCode();
            } else if (responseMessage instanceof S7MessageResponse) {
                S7MessageResponse messageResponse = (S7MessageResponse) responseMessage;
                errorClass = messageResponse.getErrorClass();
                errorCode = messageResponse.getErrorCode();
            } else if (responseMessage instanceof S7MessageUserData) {
                S7MessageUserData messageResponse = (S7MessageUserData) responseMessage;
                S7ParameterUserData parameters = (S7ParameterUserData) messageResponse.getParameter();
                parameteritem = (S7ParameterUserDataItemCPUFunctions) parameters.getItems().get(0);
                errorClass = 0;
                errorCode = parameteritem.getErrorCode().shortValue();
            } else {
                throw new PlcProtocolException("Unsupported message type " + responseMessage.getClass().getName());
            }
            if ((errorClass != 0) || (errorCode != 0)) {
                if ((errorClass == 129) && (errorCode == 4)) {
                    logger.warn("Got an error response from the PLC. This particular response code usually indicates "
                            + "that PUT/GET is not enabled on the PLC.");
                    values.put(tagName, PlcResponseCode.ACCESS_DENIED);
                    return new DefaultPlcWriteResponse(plcWriteRequest, values);
                } else {
                    logger.warn("Got an unknown error response from the PLC. Error Class: {}, Error Code {}. "
                            + "We probably need to implement explicit handling for this, so please file a bug-report "
                            + "on https://issues.apache.org/jira/projects/PLC4X and ideally attach a WireShark dump "
                            + "containing a capture of the communication.", errorClass, errorCode);
                    values.put(tagName, PlcResponseCode.INTERNAL_ERROR);
                    return new DefaultPlcWriteResponse(plcWriteRequest, values);
                }
            }
            S7PayloadWriteVarResponse payload = (S7PayloadWriteVarResponse) responseMessage.getPayload();

            S7VarPayloadStatusItem payloadItem = payload.getItems().get(0);

            responseCode = decodeResponseCode(payloadItem.getReturnCode());
            if (responseCode != PlcResponseCode.OK) {
                break;
            }
        }
        values.put(tagName, responseCode);
        return new DefaultPlcWriteResponse(plcWriteRequest, values);
    }

    private PlcResponse decodeReadResponse(S7Message responseMessage, PlcReadRequest plcReadRequest)
            throws PlcProtocolException {
        Map<String, PlcResponseItem<PlcValue>> values = new HashMap<>();
        short errorClass;
        short errorCode;

        S7ParameterUserDataItemCPUFunctions parameteritem = null;
        if (responseMessage instanceof S7MessageResponseData) {
            S7MessageResponseData messageResponseData = (S7MessageResponseData) responseMessage;
            errorClass = messageResponseData.getErrorClass();
            errorCode = messageResponseData.getErrorCode();
        } else if (responseMessage instanceof S7MessageResponse) {
            S7MessageResponse messageResponse = (S7MessageResponse) responseMessage;
            errorClass = messageResponse.getErrorClass();
            errorCode = messageResponse.getErrorCode();
        } else if (responseMessage instanceof S7MessageUserData) {
            S7MessageUserData messageResponse = (S7MessageUserData) responseMessage;
            S7ParameterUserData parameters = (S7ParameterUserData) messageResponse.getParameter();
            parameteritem = (S7ParameterUserDataItemCPUFunctions) parameters.getItems().get(0);
            errorClass = 0;
            errorCode = parameteritem.getErrorCode().shortValue();
        } else {
            throw new PlcProtocolException("Unsupported message type " + responseMessage.getClass().getName());
        }
        // If the result contains any form of non-null error code, handle this instead.
        if ((errorClass != 0) || (errorCode != 0)) {
            // This is usually the case if PUT/GET wasn't enabled on the PLC
            if ((errorClass == 129) && (errorCode == 4)) {
                logger.warn("Got an error response from the PLC. This particular response code usually indicates "
                        + "that PUT/GET is not enabled on the PLC.");
                for (String tagName : plcReadRequest.getTagNames()) {
                    PlcResponseItem<PlcValue> result = new DefaultPlcResponseItem<>(PlcResponseCode.ACCESS_DENIED,
                            new PlcNull());
                    values.put(tagName, result);
                }
                return new DefaultPlcReadResponse(plcReadRequest, values);
            } else {
                logger.warn("Got an unknown error response from the PLC. Error Class: {}, Error Code {}. "
                        + "We probably need to implement explicit handling for this, so please file a bug-report "
                        + "on https://github.com/apache/plc4x/issues and ideally attach a WireShark dump "
                        + "containing a capture of the communication.", errorClass, errorCode);
                for (String tagName : plcReadRequest.getTagNames()) {
                    PlcResponseItem<PlcValue> result = new DefaultPlcResponseItem<>(PlcResponseCode.INTERNAL_ERROR,
                            new PlcNull());
                    values.put(tagName, result);
                }
                return new DefaultPlcReadResponse(plcReadRequest, values);
            }
        }

        //TODO: Reassembling message.
        if (responseMessage instanceof S7MessageResponseData) {
            for (String tagName : plcReadRequest.getTagNames()) {
                if (plcReadRequest.getTag(tagName) instanceof S7StringTag) {
                    PlcValue plcValue = null;
                    PlcResponseCode responseCode = PlcResponseCode.INTERNAL_ERROR;
                    PlcResponseItem<PlcValue> result = new DefaultPlcResponseItem<>(responseCode, plcValue);
                    values.put(tagName, result);
                }
            }
        } else if (responseMessage instanceof S7MessageUserData) {

            S7PayloadUserData payload = (S7PayloadUserData) responseMessage.getPayload();
            if (plcReadRequest.getNumberOfTags() != payload.getItems().size()) {
                throw new PlcProtocolException(
                        "The number of requested items doesn't match the number of returned items");
            }

            List<S7PayloadUserDataItem> payloadItems = payload.getItems();

            PlcResponseCode responseCode = PlcResponseCode.INTERNAL_ERROR;
            PlcValue plcValue = null;
            int index = 0;
            for (String tagName : plcReadRequest.getTagNames()) {

                if (plcReadRequest.getTag(tagName) instanceof S7SzlTag) {

                    S7PayloadUserDataItemCpuFunctionReadSzlResponse payloadItem = (S7PayloadUserDataItemCpuFunctionReadSzlResponse) payloadItems.get(
                            index);
                    responseCode = decodeResponseCode(payloadItem.getReturnCode());

                    if (responseCode == PlcResponseCode.OK) {
                        try {
                            List<PlcValue> plcValues;
                            byte[] data = payloadItem.getItems();

                            plcValues = new LinkedList<>();
                            for (byte b : data) {
                                plcValues.add(new PlcSINT(b));
                            }

                            if (parameteritem.getLastDataUnit() == 1) {
                                CompletableFuture<S7MessageUserData> nextFuture;
                                S7ParameterUserData nextParameter;
                                S7PayloadUserData nextPayload;
                                S7PayloadUserDataItemCpuFunctionReadSzlResponse nextPayloadItem;

                                while (parameteritem.getLastDataUnit() == 1) {
                                    //TODO: Just wait for one answer!. Pending for other packages for rearm.
                                    nextFuture = reassembledMessage(parameteritem.getSequenceNumber(), plcValues);

                                    S7MessageUserData msg;

                                    msg = nextFuture.get();
                                    if (msg != null) {
                                        nextParameter = (S7ParameterUserData) msg.getParameter();
                                        parameteritem = (S7ParameterUserDataItemCPUFunctions) nextParameter.getItems()
                                                .get(0);
                                        nextPayload = (S7PayloadUserData) msg.getPayload();
                                        nextPayloadItem = (S7PayloadUserDataItemCpuFunctionReadSzlResponse) nextPayload.getItems()
                                                .get(0);
                                        for (byte b : nextPayloadItem.getItems()) {
                                            plcValues.add(new PlcSINT(b));
                                        }
                                    }

                                    plcValue = new PlcList(plcValues);
                                }
                            } else {
                                plcValue = new PlcList(plcValues);
                            }
                        } catch (Exception e) {
                            throw new PlcProtocolException("Error decoding PlcValue", e);
                        }

                    }

                }
                if (plcReadRequest.getTag(tagName) instanceof S7AckTag) {
                    S7PayloadUserDataItemCpuFunctionAlarmAckResponse payloadItem = (S7PayloadUserDataItemCpuFunctionAlarmAckResponse) payloadItems.get(
                            index);
                    responseCode = decodeResponseCode(payloadItem.getReturnCode());
                    List<Short> data = payloadItem.getMessageObjects();
                    List<PlcValue> plcValues = new LinkedList<>();
                    for (short b : data) {
                        plcValues.add(new PlcSINT((byte) b));
                    }
                    plcValue = new PlcList(plcValues);
                }
                if (plcReadRequest.getTag(tagName) instanceof S7ClkTag) {
                    DateAndTime dt;
                    if (payloadItems.get(index) instanceof S7PayloadUserDataItemClkResponse) {
                        final S7PayloadUserDataItemClkResponse payloadItem = (S7PayloadUserDataItemClkResponse) payloadItems.get(
                                index);
                        responseCode = decodeResponseCode(payloadItem.getReturnCode());
                        dt = payloadItem.getTimeStamp();
                    } else if (payloadItems.get(index) instanceof S7PayloadUserDataItemClkFResponse) {
                        final S7PayloadUserDataItemClkFResponse payloadItem = (S7PayloadUserDataItemClkFResponse) payloadItems.get(
                                index);
                        responseCode = decodeResponseCode(payloadItem.getReturnCode());
                        dt = payloadItem.getTimeStamp();
                    } else {
                        throw new PlcRuntimeException("unknown date-time type.");
                    }

                    List<PlcValue> plcValues = new LinkedList<>();
                    plcValues.add(PlcDATE_AND_LTIME.of(
                            LocalDateTime.of(dt.getYear() + 2000, dt.getMonth(), dt.getDay(), dt.getHour(),
                                    dt.getMinutes(), dt.getSeconds(), dt.getMsec() * 1000000)));
                    plcValue = new PlcList(plcValues);
                }

                PlcResponseItem<PlcValue> result = new DefaultPlcResponseItem<>(responseCode, plcValue);
                values.put(tagName, result);
                index++;
            }

            return new DefaultPlcReadResponse(plcReadRequest, values);
        }

        // In all other cases all went well.
        S7PayloadReadVarResponse payload = (S7PayloadReadVarResponse) responseMessage.getPayload();

        // If the numbers of items don't match, we're in big trouble as the only
        // way to know how to interpret the responses is by aligning them with the
        // items from the request as this information is not returned by the PLC.
        if (plcReadRequest.getNumberOfTags() != payload.getItems().size()) {
            throw new PlcProtocolException("The number of requested items doesn't match the number of returned items");
        }

        List<S7VarPayloadDataItem> payloadItems = payload.getItems();
        int index = 0;
        PlcResponseCode responseCode;
        PlcValue plcValue;
        for (String tagName : plcReadRequest.getTagNames()) {
            S7Tag tag = (S7Tag) plcReadRequest.getTag(tagName);
            S7VarPayloadDataItem payloadItem = payloadItems.get(index);

            responseCode = decodeResponseCode(payloadItem.getReturnCode());
            plcValue = null;

            if (responseCode == PlcResponseCode.OK) {
                try {
                    plcValue = parsePlcValue(tag, payloadItem.getData());
                } catch (Exception e) {
                    throw new PlcProtocolException("Error decoding PlcValue", e);
                }
            }

            PlcResponseItem<PlcValue> result = new DefaultPlcResponseItem<>(responseCode, plcValue);
            values.put(tagName, result);
            index++;
        }

        return new DefaultPlcReadResponse(plcReadRequest, values);
    }

    private PlcResponse decodeWriteResponse(S7Message responseMessage, PlcWriteRequest plcWriteRequest)
            throws PlcProtocolException {
        Map<String, PlcResponseCode> responses = new HashMap<>();
        short errorClass;
        short errorCode;

        if (responseMessage instanceof S7MessageResponseData) {
            S7MessageResponseData messageResponseData = (S7MessageResponseData) responseMessage;
            errorClass = messageResponseData.getErrorClass();
            errorCode = messageResponseData.getErrorCode();
        } else if (responseMessage instanceof S7MessageResponse) {
            S7MessageResponse messageResponse = (S7MessageResponse) responseMessage;
            errorClass = messageResponse.getErrorClass();
            errorCode = messageResponse.getErrorCode();
        } else if (responseMessage instanceof S7MessageUserData) {
            String tagName = (String) plcWriteRequest.getTagNames().toArray()[0];
            responses.put(tagName, PlcResponseCode.OK);
            return new DefaultPlcWriteResponse(plcWriteRequest, responses);
        } else {
            throw new PlcProtocolException("Unsupported message type " + responseMessage.getClass().getName());
        }
        // If the result contains any form of non-null error code, handle this instead.
        if ((errorClass != 0) || (errorCode != 0)) {
            // This is usually the case if PUT/GET wasn't enabled on the PLC
            if ((errorClass == 129) && (errorCode == 4)) {
                logger.warn("Got an error response from the PLC. This particular response code usually indicates "
                        + "that PUT/GET is not enabled on the PLC.");
                for (String tagName : plcWriteRequest.getTagNames()) {
                    responses.put(tagName, PlcResponseCode.ACCESS_DENIED);
                }
                return new DefaultPlcWriteResponse(plcWriteRequest, responses);
            } else {
                logger.warn("Got an unknown error response from the PLC. Error Class: {}, Error Code {}. "
                                + "We probably need to implement explicit handling for this, so please file a bug-report "
                                + "on https://issues.apache.org/jira/projects/PLC4X and ideally attach a WireShark dump "
                                + "containing a capture of the communication.tags:{}", errorClass, errorCode,
                        plcWriteRequest.getTagNames());
                for (String tagName : plcWriteRequest.getTagNames()) {
                    responses.put(tagName, PlcResponseCode.INTERNAL_ERROR);
                }
                return new DefaultPlcWriteResponse(plcWriteRequest, responses);
            }
        }

        // In all other cases all went well.
        S7PayloadWriteVarResponse payload = (S7PayloadWriteVarResponse) responseMessage.getPayload();
        if (payload == null) {
            throw new PlcProtocolException("response null");
        }
        // If the numbers of items don't match, we're in big trouble as the only
        // way to know how to interpret the responses is by aligning them with the
        // items from the request as this information is not returned by the PLC.
        if (plcWriteRequest.getNumberOfTags() != payload.getItems().size()) {
            throw new PlcProtocolException("The number of requested items doesn't match the number of returned items");
        }

        List<S7VarPayloadStatusItem> payloadItems = payload.getItems();
        int index = 0;
        for (String tagName : plcWriteRequest.getTagNames()) {
            S7VarPayloadStatusItem payloadItem = payloadItems.get(index);

            PlcResponseCode responseCode = decodeResponseCode(payloadItem.getReturnCode());
            responses.put(tagName, responseCode);
            index++;
        }

        return new DefaultPlcWriteResponse(plcWriteRequest, responses);
    }

    private S7VarPayloadDataItem serializePlcValue(S7Tag tag, PlcValue plcValue) {

        DataTransportSize transportSize = tag.getDataType().getDataTransportSize();

        final WriteBufferByteBased writeBuffer = serializePlcValueToWriteBuffer(tag, plcValue);
        if (writeBuffer == null) {
            return null;
        }
        byte[] data = writeBuffer.getBytes();
        if (data != null && data.length > 0) {
            return new S7VarPayloadDataItem(DataTransportErrorCode.OK, transportSize, data);
        }

        return null;
    }

    private WriteBufferByteBased serializePlcValueToWriteBuffer(S7Tag tag, PlcValue plcValue) {
        try {
            int stringLength = (tag instanceof S7StringTag) ? ((S7StringTag) tag).getStringLength() : 254;
            int lengthInBits = DataItem.getLengthInBits(plcValue.getIndex(0), tag.getDataType().getDataProtocolId(),
                    s7DriverContext.getControllerType(), stringLength, tag.getStringEncoding());
            // Cap the length of the string with the maximum allowed size.
            if (tag.getDataType() == TransportSize.STRING) {
                lengthInBits = Math.min(lengthInBits, (stringLength * 8) + 16);
            } else if (tag.getDataType() == TransportSize.WSTRING) {
                lengthInBits = Math.min(lengthInBits, (stringLength * 16) + 32);
            } else if (tag.getDataType() == TransportSize.S5TIME) {
                lengthInBits = lengthInBits * 8;
            }
            lengthInBits = lengthInBits * tag.getNumberOfElements();
            final WriteBufferByteBased writeBuffer = new WriteBufferByteBased(
                    (int) Math.ceil(((float) lengthInBits) / 8.0f));
            for (int i = 0; i < tag.getNumberOfElements(); i++) {
                DataItem.staticSerialize(writeBuffer, plcValue.getIndex(i), tag.getDataType().getDataProtocolId(),
                        s7DriverContext.getControllerType(), stringLength, tag.getStringEncoding());
            }
            return writeBuffer;
        } catch (SerializationException e) {
            logger.warn("Error serializing tag item of type: '{}'", tag.getDataType().name(), e);
        }
        return null;
    }

    private PlcValue parsePlcValue(S7Tag tag, byte[] data) {
        ReadBuffer readBuffer = new ReadBufferByteBased(data);
        try {
            int stringLength = (tag instanceof S7StringTag) ? ((S7StringTag) tag).getStringLength() : 254;
            if (tag.getNumberOfElements() == 1) {
                return DataItem.staticParse(readBuffer, tag.getDataType().getDataProtocolId(),
                        s7DriverContext.getControllerType(), stringLength, tag.getStringEncoding());
            } else {
                // In case of reading an array of bytes, make use of our simpler PlcRawByteArray as the user is
                // probably expecting to process the read raw data.
                if (tag.getDataType() == TransportSize.BYTE) {
                    return new PlcRawByteArray(data);
                } else {
                    // Fetch all
                    final PlcValue[] resultItems = IntStream.range(0, tag.getNumberOfElements()).mapToObj(i -> {
                        try {
                            return DataItem.staticParse(readBuffer, tag.getDataType().getDataProtocolId(),
                                    s7DriverContext.getControllerType(), stringLength, tag.getStringEncoding());
                        } catch (ParseException e) {
                            logger.warn("Error parsing tag item of type: '{}' (at position {}})",
                                    tag.getDataType().name(), i, e);
                        }
                        return null;
                    }).toArray(PlcValue[]::new);
                    return DefaultPlcValueHandler.of(resultItems);
                }
            }
        } catch (ParseException e) {
            logger.warn("Error parsing tag item of type: '{}'", tag.getDataType().name(), e);
        }
        return null;
    }

    /**
     * Helper to convert the return codes returned from the S7 into one of our standard PLC4X return codes
     *
     * @param dataTransportErrorCode
     *         S7 return code
     * @return PLC4X return code.
     */
    private PlcResponseCode decodeResponseCode(DataTransportErrorCode dataTransportErrorCode) {
        if (dataTransportErrorCode == null) {
            return PlcResponseCode.INTERNAL_ERROR;
        }
        switch (dataTransportErrorCode) {
        case OK:
            return PlcResponseCode.OK;
        case NOT_FOUND:
            return PlcResponseCode.NOT_FOUND;
        case INVALID_ADDRESS:
            return PlcResponseCode.INVALID_ADDRESS;
        case DATA_TYPE_NOT_SUPPORTED:
            return PlcResponseCode.INVALID_DATATYPE;
        default:
            return PlcResponseCode.INTERNAL_ERROR;
        }
    }

    /**
     * Little helper method to parse Siemens article numbers and extract the type of controller.
     *
     * @param articleNumber
     *         article number string.
     * @return type of controller.
     */
    private ControllerType decodeControllerType(String articleNumber) {
        if (!articleNumber.startsWith("6ES7 ")) {
            return ControllerType.ANY;
        }
        String model = articleNumber.substring(articleNumber.indexOf(' ') + 1, articleNumber.indexOf(' ') + 2);
        switch (model) {
        case "2":
            return ControllerType.S7_1200;
        case "5":
            return ControllerType.S7_1500;
        case "3":
            return ControllerType.S7_300;
        case "4":
            return ControllerType.S7_400;
        default:
            if (logger.isInfoEnabled()) {
                logger.info("Looking up unknown article number {}", articleNumber);
            }
            return ControllerType.ANY;
        }
    }

    /**
     * Currently we only support the S7 Any type of addresses. This helper simply converts the S7Tag from PLC4X into
     * S7Address objects.
     *
     * @param tag
     *         S7Tag instance we need to convert into an S7Address
     * @return the S7Address
     */
    protected S7Address encodeS7Address(PlcTag tag) {
        if (!(tag instanceof S7Tag)) {
            throw new PlcRuntimeException("Unsupported address type " + tag.getClass().getName());
        }
        S7Tag s7Tag = (S7Tag) tag;
        TransportSize transportSize = s7Tag.getDataType();
        int numElements = s7Tag.getNumberOfElements();
        // For these date-types we have to convert the requests to simple byte-array requests
        // As otherwise the S7 will deny them with "Data type not supported" replies.
        if (transportSize == TransportSize.STRING) {
            transportSize = TransportSize.CHAR;
            int stringLength = (s7Tag instanceof S7StringTag) ? ((S7StringTag) s7Tag).getStringLength() : 254;
            numElements = numElements * (stringLength + 2);
        } else if (transportSize == TransportSize.WSTRING) {
            transportSize = TransportSize.CHAR;
            int stringLength = (s7Tag instanceof S7StringTag) ? ((S7StringTag) s7Tag).getStringLength() : 254;
            numElements = numElements * (stringLength + 2) * 2;
        } else if (transportSize == TransportSize.BOOL && s7Tag.getNumberOfElements() > 1) {
            //numElements = (int) Math.ceil((double) numElements / 8);
            transportSize = TransportSize.BYTE;
        } else if (transportSize == TransportSize.BIT && s7Tag.getNumberOfElements() > 1) {
            //numElements = (int) Math.ceil((double) numElements / 8);
            transportSize = TransportSize.BYTE;
        }
        if (transportSize.getCode() == 0x00) {
            numElements = numElements * transportSize.getSizeInBytes();
            transportSize = TransportSize.BYTE;
        }
        return new S7AddressAny(transportSize, numElements, s7Tag.getBlockNumber(), s7Tag.getMemoryArea(),
                s7Tag.getByteOffset(), s7Tag.getBitOffset());
    }

    private boolean isConnected() {
        return conversationContext.getChannel().attr(IS_CONNECTED).get();
        //return true;
    }

    private void setChannelFeatures() {
        conversationContext.getChannel().attr(S7HMuxImpl.READ_TIME_OUT).set(s7DriverContext.getReadTimeout());
        conversationContext.getChannel().attr(S7HMuxImpl.IS_PING_ACTIVE).set(s7DriverContext.getPing());
        conversationContext.getChannel().attr(S7HMuxImpl.PING_TIME).set(s7DriverContext.getPingTime());
        conversationContext.getChannel().attr(S7HMuxImpl.RETRY_TIME).set(s7DriverContext.getRetryTime());
    }

    private CompletableFuture<S7MessageUserData> reassembledMessage(short sequenceNumber, List<PlcValue> plcValues) {

        CompletableFuture<S7MessageUserData> future = new CompletableFuture<>();

        //TODO: We need to verify that the returned tpdu id is the same in the response.
        int tpduId = getTpduId();

        TPKTPacket request = createSzlReassembledRequest(tpduId, sequenceNumber);

        conversationContext.sendRequest(request).onTimeout(e -> {
                    logger.warn("Timeout during Connection establishing, closing channel...");
                    //context.getChannel().close();
                }).expectResponse(TPKTPacket.class, Duration.ofMillis(1000)).unwrap(TPKTPacket::getPayload)
                .only(COTPPacketData.class).unwrap(COTPPacketData::getPayload).only(S7MessageUserData.class)
                .check(p -> p.getPayload() instanceof S7PayloadUserData).handle(future::complete);

        return future;
    }

    /*
     *
     */
    private TPKTPacket createSzlReassembledRequest(int tpduId, short sequenceNumber) {
        S7MessageUserData identifyRemoteMessage = new S7MessageUserData(tpduId, new S7ParameterUserData(
                List.of(new S7ParameterUserDataItemCPUFunctions((short) 0x12, (byte) 0x4, (byte) 0x4, (short) 0x01,
                        sequenceNumber, (short) 0x00, (short) 0x00, 0))), new S7PayloadUserData(
                List.of(new S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest(DataTransportErrorCode.NOT_FOUND,
                        DataTransportSize.NULL, 0x00))));
        COTPPacketData cotpPacketData = new COTPPacketData(null, identifyRemoteMessage, true, (byte) 2);
        return new TPKTPacket(cotpPacketData);
    }

    private CompletableFuture<S7MessageUserData> reassembledAlarmEvents(short sequenceNumber) {
        CompletableFuture<S7MessageUserData> future = new CompletableFuture<>();

        //TODO: We need to verify that the returned tpdu id is the same in the response.
        int tpduId = getTpduId();

        TPKTPacket request = createAlarmQueryReassembledRequest(tpduId, sequenceNumber);

        conversationContext.sendRequest(request).onTimeout(e -> {
                    logger.warn("Timeout during Connection establishing, closing channel...");
                    //context.getChannel().close();
                }).expectResponse(TPKTPacket.class, Duration.ofMillis(1000)).unwrap(TPKTPacket::getPayload)
                .only(COTPPacketData.class).unwrap(COTPPacketData::getPayload).only(S7MessageUserData.class)
                .check(p -> p.getPayload() instanceof S7PayloadUserData).handle(future::complete);

        return future;
    }

    //TODO: S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest to S7PayloadUserDataItemCpuFunctionAlarmQueryNoDataRequest
    private TPKTPacket createAlarmQueryReassembledRequest(int tpduId, short sequenceNumber) {
        S7MessageUserData identifyRemoteMessage = new S7MessageUserData(tpduId, new S7ParameterUserData(
                List.of(new S7ParameterUserDataItemCPUFunctions((short) 0x12, (byte) 0x4, (byte) 0x4, (short) 0x13,
                        sequenceNumber, (short) 0x00, (short) 0x00, 0))), new S7PayloadUserData(
                List.of(new S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest(DataTransportErrorCode.NOT_FOUND,
                        DataTransportSize.NULL, 0x00))));
        COTPPacketData cotpPacketData = new COTPPacketData(null, identifyRemoteMessage, true, (byte) 2);
        return new TPKTPacket(cotpPacketData);
    }

    private int getTpduId() {
        int tpduId = 0;
        if (this.s7DriverContext.getControllerType() != ControllerType.S7_200) {
            tpduId = tpduGenerator.getAndIncrement();
            tpduGenerator.compareAndExchange(0xFFFF, 1);
        }
        return tpduId;
    }

    @Override
    public void channelInactive(ConversationContext<TPKTPacket> context) {
        tm.shutdown();
    }
}
