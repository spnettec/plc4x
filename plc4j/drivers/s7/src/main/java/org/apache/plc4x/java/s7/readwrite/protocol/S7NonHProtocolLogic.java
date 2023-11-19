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
import org.apache.commons.lang3.concurrent.BasicThreadFactory;
import org.apache.commons.lang3.tuple.MutablePair;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.model.PlcSubscriptionHandle;
import org.apache.plc4x.java.api.model.PlcSubscriptionTag;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.s7.events.S7CyclicEvent;
import org.apache.plc4x.java.s7.events.S7Event;
import org.apache.plc4x.java.s7.readwrite.*;
import org.apache.plc4x.java.s7.readwrite.configuration.S7Configuration;
import org.apache.plc4x.java.s7.readwrite.context.S7DriverContext;
import org.apache.plc4x.java.s7.readwrite.tag.*;
import org.apache.plc4x.java.s7.readwrite.types.S7ControllerType;
import org.apache.plc4x.java.s7.readwrite.types.S7SubscriptionType;
import org.apache.plc4x.java.s7.readwrite.utils.S7PlcSubscriptionHandle;
import org.apache.plc4x.java.s7.utils.S7ParamErrorCode;
import org.apache.plc4x.java.spi.ConversationContext;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.spi.configuration.HasConfiguration;
import org.apache.plc4x.java.spi.context.DriverContext;
import org.apache.plc4x.java.spi.generation.*;
import org.apache.plc4x.java.spi.messages.*;
import org.apache.plc4x.java.spi.messages.utils.ResponseItem;
import org.apache.plc4x.java.spi.model.DefaultPlcSubscriptionTag;
import org.apache.plc4x.java.spi.transaction.RequestTransactionManager;
import org.apache.plc4x.java.spi.values.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.nio.ByteBuffer;
import java.nio.charset.Charset;
import java.time.Duration;
import java.time.LocalDateTime;
import java.util.*;
import java.util.concurrent.*;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.BiConsumer;
import java.util.function.Consumer;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import static org.apache.plc4x.java.spi.connection.AbstractPlcConnection.IS_CONNECTED;

/**
 * The S7 Protocol states that there can not be more then {min(maxAmqCaller,
 * maxAmqCallee} "ongoing" requests. So we need to limit those. Thus, each
 * request goes to a Work Queue and this Queue ensures, that only 3 are open at
 * the same time.
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
		this.tm = new RequestTransactionManager(1);
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
			context.fireConnected();
			return;
		}

		// Only the TCP transport supports login.
		logger.info("S7 Driver running in ACTIVE mode.");
		logger.debug("Sending COTP Connection Request");
		// Open the session on ISO Transport Protocol first.
		TPKTPacket packet = new TPKTPacket(createCOTPConnectionRequest(s7DriverContext.getCalledTsapId(),
				s7DriverContext.getCallingTsapId(), s7DriverContext.getCotpTpduSize()));

		context.getChannel().pipeline().names().forEach(s -> {
			logger.debug("Nombre tuberias: " + s);
		});

		context.sendRequest(packet).onTimeout(e -> {
			logger.info("Timeout during Connection establishing, closing channel...");
			// TODO: We're saying that we're closing the channel, but not closing the
			// channel ... sure, this is what we want?
			// context.getChannel().close();
		}).expectResponse(TPKTPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
				.check(p -> p.getPayload() instanceof COTPPacketConnectionResponse)
				.unwrap(p -> (COTPPacketConnectionResponse) p.getPayload()).handle(cotpPacketConnectionResponse -> {
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
								if (s7DriverContext.getControllerType() != S7ControllerType.ANY) {
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
								}).expectResponse(TPKTPacket.class,
										Duration.ofMillis(configuration.getTimeoutRequest()))
										.check(p -> p.getPayload() instanceof COTPPacketData)
										.unwrap(p -> ((COTPPacketData) p.getPayload()))
										.check(p -> p.getPayload() instanceof S7MessageUserData)
										.unwrap(p -> ((S7MessageUserData) p.getPayload()))
										.check(p -> p.getPayload() instanceof S7PayloadUserData)
										.handle(messageUserData -> {
											logger.debug("Got S7 Identification Response");
											S7PayloadUserData payloadUserData = (S7PayloadUserData) messageUserData
													.getPayload();
											extractControllerTypeAndFireConnected(context, payloadUserData);
										});
							});
				});
	}

	/*
	 * It performs the sequential and safe shutdown of the driver. Completion of
	 * pending requests, executors and associated tasks.
	 */
	@Override
	public void onDisconnect(ConversationContext<TPKTPacket> context) {
		// 1. Clear all pending requests and their associated transaction
		// 3. Performs the shutdown of the transaction executor.
		tm.shutdown();
		// 5. Executes the closing of the main channel.
		context.fireDisconnected();
		// 6. Here is the stop of any task or state machine that is added.
	}

	@Override
	public CompletableFuture<PlcReadResponse> read(PlcReadRequest readRequest) {
		if (!isConnected()) {
			CompletableFuture<PlcReadResponse> future = new CompletableFuture<>();
			future.completeExceptionally(new PlcRuntimeException("Disconnected"));
			return future;
		}
		DefaultPlcReadRequest request = (DefaultPlcReadRequest) readRequest;
		List<S7VarRequestParameterItem> requestItems = new ArrayList<>(request.getNumberOfTags());

		if (request.getTags().get(0) instanceof S7SzlTag) {
			S7SzlTag szltag = (S7SzlTag) request.getTags().get(0);

			final S7MessageUserData s7SzlMessageRequest = new S7MessageUserData(1,
					new S7ParameterUserData(List.of(new S7ParameterUserDataItemCPUFunctions((short) 0x11, (byte) 0x4,
							(byte) 0x4, (short) 0x01, (short) 0x00, null, null, null))),
					new S7PayloadUserData(List.of(new S7PayloadUserDataItemCpuFunctionReadSzlRequest(
							DataTransportErrorCode.OK, DataTransportSize.OCTET_STRING, 0x04,
							new SzlId(SzlModuleTypeClass.enumForValue((byte) ((szltag.getSzlId() & 0xf000) >> 12)),
									(byte) ((szltag.getSzlId() & 0x0f00) >> 8),
									SzlSublist.enumForValue((short) (szltag.getSzlId() & 0x00ff))),
							szltag.getIndex()))));

			return toPlcReadResponse(readRequest, readInternal(s7SzlMessageRequest));

		}

		for (PlcTag tag : request.getTags()) {
			requestItems.add(new S7VarRequestParameterItemAddress(encodeS7Address(tag)));
		}

		// Create a read request template.
		// tpuId will be inserted before sending in #readInternal so we insert -1 as
		// dummy here
		final S7MessageRequest s7MessageRequest = new S7MessageRequest(-1, new S7ParameterReadVarRequest(requestItems),
				null);

		// Just send a single response and chain it as Response
		return toPlcReadResponse(readRequest, readInternal(s7MessageRequest));
	}

	/**
	 * Maps the S7ReadResponse of a PlcReadRequest to a PlcReadResponse
	 */
	private CompletableFuture<PlcReadResponse> toPlcReadResponse(PlcReadRequest readRequest,
			CompletableFuture<S7Message> response) {
		return response.thenApply(value -> {
			try {
				return (PlcReadResponse) decodeReadResponse(value, readRequest);
			} catch (PlcProtocolException e) {
				throw new RuntimeException(e);
			}
		});
	}

	/**
	 * Sends one Read over the Wire and internally returns the Response Do sending
	 * of normally sized single-message request.
	 * <p>
	 * Assumes that the {@link S7MessageRequest} and its expected
	 * {@link S7MessageResponseData} and does not further check that!
	 */
	private CompletableFuture<S7Message> readInternal(S7Message request) {
		CompletableFuture<S7Message> future = new CompletableFuture<>();
		int thisTpduId = 0;
		if (this.s7DriverContext.getControllerType() != S7ControllerType.S7_200) {
			thisTpduId = tpduGenerator.getAndIncrement();
		}
		final int tpduId = thisTpduId;
		// If we've reached the max value for a 16 bit transaction identifier, reset
		// back to 1
		tpduGenerator.compareAndExchange(0xFFFF, 1);

		S7Message message = (request instanceof S7MessageUserData)
				? new S7MessageUserData(tpduId, request.getParameter(), request.getPayload())
				: new S7MessageRequest(tpduId, request.getParameter(), request.getPayload());

		TPKTPacket tpktPacket = new TPKTPacket(new COTPPacketData(null, message, true, (byte) tpduId));

		// Start a new request-transaction (Is ended in the response-handler)

		RequestTransactionManager.RequestTransaction transaction = tm.startRequest();

		transaction.submit(
				() -> context.sendRequest(tpktPacket).onTimeout(new TransactionErrorCallback<>(future, transaction))
						.onError(new TransactionErrorCallback<>(future, transaction))
						.expectResponse(TPKTPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
						.check(p -> p.getPayload() instanceof COTPPacketData)
						.unwrap(p -> (COTPPacketData) p.getPayload()).check(p -> p.getPayload() != null)
						.unwrap(COTPPacket::getPayload).check(p -> p.getTpduReference() == tpduId).handle(p -> {
							future.complete(p);
							// Finish the request-transaction.
							transaction.endRequest();
						}));

		return future;
	}

	// TODO: Clean code
	@Override
	public CompletableFuture<PlcWriteResponse> write(PlcWriteRequest writeRequest) {

		CompletableFuture<PlcWriteResponse> future = new CompletableFuture<>();
		DefaultPlcWriteRequest request = (DefaultPlcWriteRequest) writeRequest;

		List<String> clkTags = request.getTagNames().stream().filter(t -> request.getTag(t) instanceof S7ClkTag)
				.collect(Collectors.toList());

		List<S7VarRequestParameterItem> parameterItems = new ArrayList<>(request.getNumberOfTags());
		List<S7VarPayloadDataItem> payloadItems = new ArrayList<>(request.getNumberOfTags());

		Iterator<String> iter = request.getTagNames().iterator();

		String tagName;
		while (iter.hasNext()) {
			tagName = iter.next();
			final S7Tag tag = (S7Tag) request.getTag(tagName);
			final PlcValue plcValue = request.getPlcValue(tagName);
			parameterItems.add(new S7VarRequestParameterItemAddress(encodeS7Address(tag)));
			payloadItems.add(serializePlcValue(tag, plcValue, iter.hasNext()));
		}

//        for (String tagName : request.getTagNames()) {
//            final S7Tag tag = (S7Tag) request.getTag(tagName);
//            final PlcValue plcValue = request.getPlcValue(tagName);
//            parameterItems.add(new S7VarRequestParameterItemAddress(encodeS7Address(tag)));
//            payloadItems.add(serializePlcValue(tag, plcValue));
//
//        }

		final int tpduId = tpduGenerator.getAndIncrement();
		// If we've reached the max value for a 16 bit transaction identifier, reset
		// back to 1
		tpduGenerator.compareAndExchange(0xFFFF, 1);

		TPKTPacket tpktPacket = new TPKTPacket(
				new COTPPacketData(null, new S7MessageRequest(tpduId, new S7ParameterWriteVarRequest(parameterItems),
						new S7PayloadWriteVarRequest(payloadItems)), true, (byte) tpduId));

		// Start a new request-transaction (Is ended in the response-handler)
		RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
		transaction.submit(() -> context.sendRequest(tpktPacket)
				.onTimeout(new TransactionErrorCallback<>(future, transaction))
				.onError(new TransactionErrorCallback<>(future, transaction))
				.expectResponse(TPKTPacket.class, Duration.ofMillis(configuration.getTimeoutRequest()))
				.check(p -> p.getPayload() instanceof COTPPacketData).unwrap(p -> ((COTPPacketData) p.getPayload()))
				.unwrap(COTPPacket::getPayload).check(p -> p.getTpduReference() == tpduId).handle(p -> {
					try {
						future.complete(((PlcWriteResponse) decodeWriteResponse(p, writeRequest)));
					} catch (PlcProtocolException e) {
						logger.warn("Error sending 'write' message: '{}'", e.getMessage(), e);
					}
					// Finish the request-transaction.
					transaction.endRequest();
				}));
		return future;
	}

	@Override
	public CompletableFuture<PlcSubscriptionResponse> subscribe(PlcSubscriptionRequest subscriptionRequest) {
		throw new UnsupportedOperationException("subscribe, Not supported.");
	}

	@Override
	public CompletableFuture<PlcUnsubscriptionResponse> unsubscribe(PlcUnsubscriptionRequest unsubscriptionRequest) {
		throw new UnsupportedOperationException("unsubscribe, Not supported.");
	}

	@Override
	public void close(ConversationContext<TPKTPacket> context) {

	}

	/**
	 * This method is only called when there is no Response Handler.
	 */
	@Override
	protected void decode(ConversationContext<TPKTPacket> context, TPKTPacket msg) throws Exception {

		S7Message s7msg = msg.getPayload().getPayload();
		S7Parameter parameter = s7msg.getParameter();
		if (parameter instanceof S7ParameterModeTransition) {

		} else if (parameter instanceof S7ParameterUserData) {
			S7ParameterUserData parameterUD = (S7ParameterUserData) parameter;
			List<S7ParameterUserDataItem> parameterUDItems = parameterUD.getItems();
			for (S7ParameterUserDataItem parameterUDItem : parameterUDItems) {
				if (parameterUDItem instanceof S7ParameterUserDataItemCPUFunctions) {
					S7ParameterUserDataItemCPUFunctions myParameter = (S7ParameterUserDataItemCPUFunctions) parameterUDItem;
					// TODO: Check from mspec. We can try using "instanceof"
					if ((myParameter.getCpuFunctionType() == 0x00) && (myParameter.getCpuSubfunction() == 0x03)) {
						S7PayloadUserData payload = (S7PayloadUserData) s7msg.getPayload();
						List<S7PayloadUserDataItem> items = payload.getItems();
						for (S7PayloadUserDataItem item : items) {
							if (item instanceof S7PayloadDiagnosticMessage) {

							}
						}
					} else if ((myParameter.getCpuFunctionType() == 0x00) && ((myParameter.getCpuSubfunction() == 0x05)
							|| (myParameter.getCpuSubfunction() == 0x06) || (myParameter.getCpuSubfunction() == 0x0c)
							|| (myParameter.getCpuSubfunction() == 0x11) || (myParameter.getCpuSubfunction() == 0x12)
							|| (myParameter.getCpuSubfunction() == 0x13)
							|| (myParameter.getCpuSubfunction() == 0x16))) {

					} else if ((myParameter.getCpuFunctionType() == 0x00)
							&& (myParameter.getCpuSubfunction() == 0x13)) {

					} else if ((myParameter.getCpuFunctionGroup() == 0x02) && (myParameter.getCpuFunctionType() == 0x00)
							&& (myParameter.getCpuSubfunction() == 0x01)) {

					} else if ((myParameter.getCpuFunctionGroup() == 0x02) && (myParameter.getCpuFunctionType() == 0x00)
							&& (myParameter.getCpuSubfunction() == 0x05)) {

					} else if ((myParameter.getCpuFunctionType() == 0x08)
							&& (myParameter.getCpuSubfunction() == 0x01)) {

					} else if ((myParameter.getCpuFunctionType() == 0x08)
							&& (myParameter.getCpuSubfunction() == 0x04)) {

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
			ByteBuf szlItem = Unpooled.wrappedBuffer(((S7PayloadUserDataItemCpuFunctionReadSzlResponse)item).getItems());
			String articleNumber = szlItem.toString(2, 20, Charset.defaultCharset());
			s7DriverContext.setControllerType(decodeControllerType(articleNumber));
			context.fireConnected();
            break;
		}
	}

	private TPKTPacket createIdentifyRemoteMessage() {
		S7MessageUserData identifyRemoteMessage = new S7MessageUserData(1,
				new S7ParameterUserData(Collections.singletonList(new S7ParameterUserDataItemCPUFunctions((short) 0x11,
						(byte) 0x4, (byte) 0x4, (short) 0x01, (short) 0x00, null, null, null))),
				new S7PayloadUserData(Collections.singletonList(new S7PayloadUserDataItemCpuFunctionReadSzlRequest(
						DataTransportErrorCode.OK, DataTransportSize.OCTET_STRING, 0x0C,
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
		if (this.s7DriverContext.getControllerType() == S7ControllerType.S7_200) {
			tpduId = 0;
		}
		COTPPacketData cotpPacketData = new COTPPacketData(null, s7Message, true, (byte) tpduId);
		return new TPKTPacket(cotpPacketData);
	}

	private COTPPacketConnectionRequest createCOTPConnectionRequest(int calledTsapId, int callingTsapId,
			COTPTpduSize cotpTpduSize) {
		return new COTPPacketConnectionRequest(
				Arrays.asList(new COTPParameterCallingTsap(callingTsapId), new COTPParameterCalledTsap(calledTsapId),
						new COTPParameterTpduSize(cotpTpduSize)),
				null, (short) 0x0000, (short) 0x000F, COTPProtocolClass.CLASS_0);
	}

	private PlcResponse decodeReadResponse(S7Message responseMessage, PlcReadRequest plcReadRequest)
			throws PlcProtocolException {
		Map<String, ResponseItem<PlcValue>> values = new HashMap<>();
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
					ResponseItem<PlcValue> result = new ResponseItem<>(PlcResponseCode.ACCESS_DENIED, new PlcNull());
					values.put(tagName, result);
				}
				return new DefaultPlcReadResponse(plcReadRequest, values);
			} else {
				logger.warn("Got an unknown error response from the PLC. Error Class: {}, Error Code {}. "
						+ "We probably need to implement explicit handling for this, so please file a bug-report "
						+ "on https://issues.apache.org/jira/projects/PLC4X and ideally attach a WireShark dump "
						+ "containing a capture of the communication.", errorClass, errorCode);
				for (String tagName : plcReadRequest.getTagNames()) {
					ResponseItem<PlcValue> result = new ResponseItem<>(PlcResponseCode.INTERNAL_ERROR, new PlcNull());
					values.put(tagName, result);
				}
				return new DefaultPlcReadResponse(plcReadRequest, values);
			}
		}

		// TODO: Reemsambling message.
		if (responseMessage instanceof S7MessageUserData) {

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

					S7PayloadUserDataItemCpuFunctionReadSzlResponse payloadItem = (S7PayloadUserDataItemCpuFunctionReadSzlResponse) payloadItems
							.get(index);
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
									// TODO: Just wait for one answer!. Pending for other packages for rearm.
									nextFuture = reassembledMessage(parameteritem.getSequenceNumber(), plcValues);

									S7MessageUserData msg;

									msg = nextFuture.get();
									if (msg != null) {
										nextParameter = (S7ParameterUserData) msg.getParameter();
										parameteritem = (S7ParameterUserDataItemCPUFunctions) nextParameter.getItems()
												.get(0);
										nextPayload = (S7PayloadUserData) msg.getPayload();
										nextPayloadItem = (S7PayloadUserDataItemCpuFunctionReadSzlResponse) nextPayload
												.getItems().get(0);
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
					S7PayloadUserDataItemCpuFunctionAlarmAckResponse payloadItem = (S7PayloadUserDataItemCpuFunctionAlarmAckResponse) payloadItems
							.get(index);
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
						final S7PayloadUserDataItemClkResponse payloadItem = (S7PayloadUserDataItemClkResponse) payloadItems
								.get(index);
						responseCode = decodeResponseCode(payloadItem.getReturnCode());
						dt = payloadItem.getTimeStamp();
					} else if (payloadItems.get(index) instanceof S7PayloadUserDataItemClkFResponse) {
						final S7PayloadUserDataItemClkFResponse payloadItem = (S7PayloadUserDataItemClkFResponse) payloadItems
								.get(index);
						responseCode = decodeResponseCode(payloadItem.getReturnCode());
						dt = payloadItem.getTimeStamp();
					} else {
						throw new PlcRuntimeException("unknown date-time type.");
					}

					List<PlcValue> plcValues = new LinkedList<>();
					plcValues.add(PlcLDATE_AND_TIME.of(LocalDateTime.of(dt.getYear() + 2000, dt.getMonth(), dt.getDay(),
							dt.getHour(), dt.getMinutes(), dt.getSeconds(), dt.getMsec() * 1000000)));
					plcValue = new PlcList(plcValues);
				}

				ResponseItem<PlcValue> result = new ResponseItem<>(responseCode, plcValue);
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
		for (String tagName : plcReadRequest.getTagNames()) {
			S7Tag tag = (S7Tag) plcReadRequest.getTag(tagName);
			S7VarPayloadDataItem payloadItem = payloadItems.get(index);

			PlcResponseCode responseCode = decodeResponseCode(payloadItem.getReturnCode());
			PlcValue plcValue = null;
			ByteBuf data = Unpooled.wrappedBuffer(payloadItem.getData());
			if (responseCode == PlcResponseCode.OK) {
				try {
					plcValue = parsePlcValue(tag, data);
				} catch (Exception e) {
					throw new PlcProtocolException("Error decoding PlcValue", e);
				}
			}
			ResponseItem<PlcValue> result = new ResponseItem<>(responseCode, plcValue);
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
						+ "containing a capture of the communication.", errorClass, errorCode);
				for (String tagName : plcWriteRequest.getTagNames()) {
					responses.put(tagName, PlcResponseCode.INTERNAL_ERROR);
				}
				return new DefaultPlcWriteResponse(plcWriteRequest, responses);
			}
		}

		// In all other cases all went well.
		S7PayloadWriteVarResponse payload = (S7PayloadWriteVarResponse) responseMessage.getPayload();

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

	private S7VarPayloadDataItem serializePlcValue(S7Tag tag, PlcValue plcValue, Boolean hasNext) {
		try {
			DataTransportSize transportSize = tag.getDataType().getDataTransportSize();
			int stringLength = (tag instanceof S7StringTag) ? ((S7StringTag) tag).getStringLength() : 254;
			ByteBuffer byteBuffer = null;
			for (int i = 0; i < tag.getNumberOfElements(); i++) {
				final int lengthInBytes = DataItem.getLengthInBytes(plcValue.getIndex(i),
						tag.getDataType().getDataProtocolId(), stringLength, tag.getStringEncoding());
				final WriteBufferByteBased writeBuffer = new WriteBufferByteBased(lengthInBytes);
				if (tag.getDataType() == TransportSize.BOOL && tag.getNumberOfElements() == 1) {
					writeBuffer.writeUnsignedShort("", 7, ((Number) (short) 0x00).shortValue());
				}
				DataItem.staticSerialize(writeBuffer, plcValue.getIndex(i), tag.getDataType().getDataProtocolId(),
						stringLength, tag.getStringEncoding());
				// Allocate enough space for all items.
				if (byteBuffer == null) {
					byteBuffer = ByteBuffer.allocate(lengthInBytes * tag.getNumberOfElements());
				}
				byteBuffer.put(writeBuffer.getBytes());
			}
			if (byteBuffer != null) {
				byte[] data = byteBuffer.array();
				return new S7VarPayloadDataItem(DataTransportErrorCode.OK, transportSize, data/* , hasNext */);
			}
		} catch (SerializationException e) {
			logger.warn("Error serializing tag item of type: '{}'", tag.getDataType().name(), e);
		}
		return null;
	}

	private PlcValue parsePlcValue(S7Tag tag, ByteBuf data) {
		ReadBuffer readBuffer = new ReadBufferByteBased(data.array());
		try {
			int stringLength = (tag instanceof S7StringTag) ? ((S7StringTag) tag).getStringLength() : 254;
			if (tag.getNumberOfElements() == 1) {
				if (tag.getDataType() == TransportSize.BOOL) {
					short reserved = readBuffer.readUnsignedShort("", 7);
					if (reserved != (short) 0x00) {
						logger.info(
								"Expected constant value " + 0x00 + " but got " + reserved + " for reserved field.");
					}
				}
				return DataItem.staticParse(readBuffer, tag.getDataType().getDataProtocolId(), stringLength,
						tag.getStringEncoding());
			} else {
				// Fetch all
				final PlcValue[] resultItems = IntStream.range(0, tag.getNumberOfElements()).mapToObj(i -> {
					try {
						return DataItem.staticParse(readBuffer, tag.getDataType().getDataProtocolId(), stringLength,
								tag.getStringEncoding());
					} catch (ParseException e) {
						logger.warn("Error parsing tag item of type: '{}' (at position {}})", tag.getDataType().name(),
								i, e);
					}
					return null;
				}).toArray(PlcValue[]::new);
				return PlcValueHandler.of(resultItems);
			}
		} catch (ParseException e) {
			logger.warn("Error parsing tag item of type: '{}'", tag.getDataType().name(), e);
		}
		return null;
	}

	/**
	 * Helper to convert the return codes returned from the S7 into one of our
	 * standard PLC4X return codes
	 *
	 * @param dataTransportErrorCode S7 return code
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
	 * Little helper method to parse Siemens article numbers and extract the type of
	 * controller.
	 *
	 * @param articleNumber article number string.
	 * @return type of controller.
	 */
	private S7ControllerType decodeControllerType(String articleNumber) {
		if (!articleNumber.startsWith("6ES7 ")) {
			return S7ControllerType.ANY;
		}
		String model = articleNumber.substring(articleNumber.indexOf(' ') + 1, articleNumber.indexOf(' ') + 2);
		switch (model) {
		case "2":
			return S7ControllerType.S7_1200;
		case "5":
			return S7ControllerType.S7_1500;
		case "3":
			return S7ControllerType.S7_300;
		case "4":
			return S7ControllerType.S7_400;
		default:
			if (logger.isInfoEnabled()) {
				logger.info("Looking up unknown article number {}", articleNumber);
			}
			return S7ControllerType.ANY;
		}
	}

	/**
	 * Currently we only support the S7 Any type of addresses. This helper simply
	 * converts the S7Tag from PLC4X into S7Address objects.
	 *
	 * @param tag S7Tag instance we need to convert into an S7Address
	 * @return the S7Address
	 */
	protected S7Address encodeS7Address(PlcTag tag) {
		if (!(tag instanceof S7Tag)) {
			throw new PlcRuntimeException("Unsupported address type " + tag.getClass().getName());
		}
		S7Tag s7Tag = (S7Tag) tag;
		TransportSize transportSize = s7Tag.getDataType();
		int numElements = s7Tag.getNumberOfElements();
		// For these date-types we have to convert the requests to simple byte-array
		// requests
		// As otherwise the S7 will deny them with "Data type not supported" replies.
		if ((transportSize == TransportSize.TIME)
				/* || (transportSize == TransportSize.S7_S5TIME) */ || (transportSize == TransportSize.LINT)
				|| (transportSize == TransportSize.ULINT) || (transportSize == TransportSize.LWORD)
				|| (transportSize == TransportSize.LREAL) || (transportSize == TransportSize.REAL)
				|| (transportSize == TransportSize.LTIME) || (transportSize == TransportSize.DATE)
				|| (transportSize == TransportSize.TIME_OF_DAY) || (transportSize == TransportSize.DATE_AND_TIME)) {
			numElements = numElements * transportSize.getSizeInBytes();
			// ((S7Field) field).setDataType(transportSize);
			transportSize = TransportSize.BYTE;
		}
		if (transportSize == TransportSize.BOOL) {
			if (numElements > 1) {
				transportSize = TransportSize.BYTE;
			}
			numElements = numElements * transportSize.getSizeInBytes();
		}
		if (transportSize == TransportSize.CHAR) {
			transportSize = TransportSize.BYTE;
			numElements = numElements * transportSize.getSizeInBytes();
		}
		if (transportSize == TransportSize.WCHAR) {
			transportSize = TransportSize.BYTE;
			numElements = numElements * transportSize.getSizeInBytes() * 2;
		}
		if (transportSize == TransportSize.STRING) {
			transportSize = TransportSize.BYTE;
			int stringLength = (s7Tag instanceof S7StringTag) ? ((S7StringTag) s7Tag).getStringLength() : 254;
			numElements = numElements * (stringLength + 2);
		} else if (transportSize == TransportSize.WSTRING) {
			transportSize = TransportSize.BYTE;
			int stringLength = (s7Tag instanceof S7StringTag) ? ((S7StringTag) s7Tag).getStringLength() : 254;
			numElements = numElements * (stringLength + 2) * 2;
		}
		return new S7AddressAny(transportSize, numElements, s7Tag.getBlockNumber(), s7Tag.getMemoryArea(),
				s7Tag.getByteOffset(), s7Tag.getBitOffset());
	}

	/*
	 * In the case of a reconnection, there may be requests waiting, for which the
	 * operation must be terminated by exception and canceled in the transaction
	 * manager. If this does not happen, the driver operation can be frozen.
	 */
	private boolean isConnected() {
		return context.getChannel().attr(IS_CONNECTED).get();
		// return true;
	}

	private CompletableFuture<S7MessageUserData> reassembledMessage(short sequenceNumber, List<PlcValue> plcValues) {

		CompletableFuture<S7MessageUserData> future = new CompletableFuture<>();

		// TODO: PDU id is the same, we need check.
		int tpduId = tpduGenerator.getAndIncrement();
		// If we've reached the max value for a 16 bit transaction identifier, reset
		// back to 1
		tpduGenerator.compareAndExchange(0xFFFF, 1);

		TPKTPacket request = createSzlReassembledRequest(tpduId, sequenceNumber);

		context.sendRequest(request).onTimeout(e -> {
			logger.warn("Timeout during Connection establishing, closing channel...");
			// context.getChannel().close();
		}).expectResponse(TPKTPacket.class, Duration.ofMillis(1000))
				.check(p -> p.getPayload() instanceof COTPPacketData).unwrap(p -> ((COTPPacketData) p.getPayload()))
				.check(p -> p.getPayload() instanceof S7MessageUserData)
				.unwrap(p -> ((S7MessageUserData) p.getPayload()))
				.check(p -> p.getPayload() instanceof S7PayloadUserData).handle(future::complete);

		return future;
	}

	/*
	 *
	 */
	private TPKTPacket createSzlReassembledRequest(int tpduId, short sequenceNumber) {
		S7MessageUserData identifyRemoteMessage = new S7MessageUserData(tpduId,
				new S7ParameterUserData(List.of(new S7ParameterUserDataItemCPUFunctions((short) 0x12, (byte) 0x4,
						(byte) 0x4, (short) 0x01, sequenceNumber, (short) 0x00, (short) 0x00, 0))),
				new S7PayloadUserData(List.of(new S7PayloadUserDataItemCpuFunctionReadSzlNoDataRequest(
						DataTransportErrorCode.NOT_FOUND, DataTransportSize.NULL, 0x00))));
		COTPPacketData cotpPacketData = new COTPPacketData(null, identifyRemoteMessage, true, (byte) 2);
		return new TPKTPacket(cotpPacketData);
	}

	/**
	 * A generic purpose error handler which terminates transaction and calls back
	 * given future with error message.
	 */
	static class TransactionErrorCallback<T, E extends Throwable>
			implements Consumer<TimeoutException>, BiConsumer<TPKTPacket, E> {

		private final CompletableFuture<T> future;
		private final RequestTransactionManager.RequestTransaction transaction;

		TransactionErrorCallback(CompletableFuture<T> future,
				RequestTransactionManager.RequestTransaction transaction) {
			this.future = future;
			this.transaction = transaction;
		}

		@Override
		public void accept(TimeoutException e) {
			try {
				transaction.endRequest();
			} catch (Exception ex) {
				logger.info(ex.getMessage());
			}
			future.completeExceptionally(e);
		}

		@Override
		public void accept(TPKTPacket tpktPacket, E e) {
			try {
				transaction.endRequest();
			} catch (Exception ex) {
				logger.info(ex.getMessage());
			}
			future.completeExceptionally(e);
		}
	}

}
