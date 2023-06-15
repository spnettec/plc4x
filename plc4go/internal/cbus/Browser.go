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

package cbus

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	plc4go "github.com/apache/plc4x/plc4go/pkg/api"
	apiModel "github.com/apache/plc4x/plc4go/pkg/api/model"
	"github.com/apache/plc4x/plc4go/pkg/api/values"
	readWriteModel "github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"
	_default "github.com/apache/plc4x/plc4go/spi/default"
	spiModel "github.com/apache/plc4x/plc4go/spi/model"
	"github.com/apache/plc4x/plc4go/spi/options"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Browser struct {
	_default.DefaultBrowser
	connection      plc4go.PlcConnection
	sequenceCounter uint8

	log zerolog.Logger
}

func NewBrowser(connection plc4go.PlcConnection, _options ...options.WithOption) *Browser {
	browser := Browser{
		connection:      connection,
		sequenceCounter: 0,

		log: options.ExtractCustomLogger(_options...),
	}
	browser.DefaultBrowser = _default.NewDefaultBrowser(browser, _options...)
	return &browser
}

func (m Browser) BrowseQuery(ctx context.Context, interceptor func(result apiModel.PlcBrowseItem) bool, queryName string, query apiModel.PlcQuery) (apiModel.PlcResponseCode, []apiModel.PlcBrowseItem) {
	var queryResults []apiModel.PlcBrowseItem
	switch query := query.(type) {
	case *unitInfoQuery:
		m.log.Trace().Msg("extract units")
		units, allUnits, err := m.extractUnits(ctx, query, m.getInstalledUnitAddressBytes)
		if err != nil {
			m.log.Error().Err(err).Msg("Error extracting units")
			return apiModel.PlcResponseCode_INTERNAL_ERROR, nil
		}
		attributes, allAttributes := m.extractAttributes(query)

		if allUnits {
			m.log.Info().Msg("Querying all (available) units")
		}
	unitLoop:
		for _, unit := range units {
			m.log.Trace().Msgf("checking unit:\n%s", unit)
			if err := ctx.Err(); err != nil {
				m.log.Info().Err(err).Msgf("Aborting scan at unit %s", unit)
				return apiModel.PlcResponseCode_INVALID_ADDRESS, nil
			}
			unitAddress := unit.GetAddress()
			if !allUnits && allAttributes {
				m.log.Info().Msgf("Querying all attributes of unit %d", unitAddress)
			}
			event := m.log.Info()
			if allUnits {
				event = m.log.Debug()
			}
			event.Msgf("Query unit  %d", unitAddress)
			for _, attribute := range attributes {
				if err := ctx.Err(); err != nil {
					m.log.Info().Err(err).Msgf("Aborting scan at unit %s", unit)
					return apiModel.PlcResponseCode_INVALID_ADDRESS, nil
				}
				if !allUnits && !allAttributes {
					m.log.Info().Msgf("Querying attribute %s of unit %d", attribute, unitAddress)
				} else {
					event.Msgf("unit %d: Query %s", unitAddress, attribute)
				}
				m.log.Trace().Msg("Building request")
				readTagName := fmt.Sprintf("%s/%d/%s", queryName, unitAddress, attribute)
				readRequest, _ := m.connection.ReadRequestBuilder().
					AddTag(readTagName, NewCALIdentifyTag(unit, nil /*TODO: add bridge support*/, attribute, 1)).
					Build()
				timeoutCtx, timeoutCancel := context.WithTimeout(ctx, time.Second*2)
				m.log.Trace().Msgf("Executing readRequest\n%s\nwith timeout %s", readRequest, timeoutCtx)
				requestResult := <-readRequest.ExecuteWithContext(timeoutCtx)
				m.log.Trace().Msg("got a response")
				timeoutCancel()
				if err := requestResult.GetErr(); err != nil {
					if allUnits || allAttributes {
						event = m.log.Trace()
					}
					event.Err(err).Msgf("unit %d: Can't read attribute %s", unitAddress, attribute)
					continue unitLoop
				}
				response := requestResult.GetResponse()
				if code := response.GetResponseCode(readTagName); code != apiModel.PlcResponseCode_OK {
					event.Msgf("unit %d: error reading tag %s. Code %s", unitAddress, attribute, code)
					continue unitLoop
				}
				queryResult := spiModel.NewDefaultPlcBrowseItem(
					NewCALIdentifyTag(unit, nil /*TODO: add bridge support*/, attribute, 1),
					queryName,
					"",
					true,
					false,
					false,
					nil,
					map[string]values.PlcValue{
						"CurrentValue": response.GetValue(readTagName),
					},
				)
				if interceptor != nil {
					m.log.Trace().Msg("forwarding query result to interceptor")
					interceptor(queryResult)
				}
				queryResults = append(queryResults, queryResult)
			}
		}
	default:
		m.log.Warn().Msgf("unsupported query type supplied %T", query)
		return apiModel.PlcResponseCode_INVALID_ADDRESS, nil
	}
	m.log.Trace().Msgf("Browse done with \n%s", queryResults)
	return apiModel.PlcResponseCode_OK, queryResults
}

func (m Browser) extractUnits(ctx context.Context, query *unitInfoQuery, getInstalledUnitAddressBytes func(ctx context.Context) (map[byte]any, error)) ([]readWriteModel.UnitAddress, bool, error) {
	if unitAddress := query.unitAddress; unitAddress != nil {
		return []readWriteModel.UnitAddress{unitAddress}, false, nil
	} else {
		// TODO: check if we still want the option to brute force all addresses
		installedUnitAddressBytes, err := getInstalledUnitAddressBytes(ctx)
		if err != nil {
			return nil, false, errors.New("Unable to get installed uints")
		}

		var units []readWriteModel.UnitAddress
		for i := 0; i <= 0xFF; i++ {
			unitAddressByte := byte(i)
			if _, ok := installedUnitAddressBytes[unitAddressByte]; ok {
				units = append(units, readWriteModel.NewUnitAddress(unitAddressByte))
			}
		}
		return units, true, nil
	}
}

func (m Browser) extractAttributes(query *unitInfoQuery) ([]readWriteModel.Attribute, bool) {
	if attribute := query.attribute; attribute != nil {
		return []readWriteModel.Attribute{*attribute}, false
	} else {
		var attributes []readWriteModel.Attribute
		for _, attribute := range readWriteModel.AttributeValues {
			attributes = append(attributes, attribute)
		}
		return attributes, true
	}
}

func (m Browser) getInstalledUnitAddressBytes(ctx context.Context) (map[byte]any, error) {
	start := time.Now()
	defer func() {
		m.log.Debug().Msgf("Ending unit address acquiring after %s", time.Since(start))
	}()
	// We need to pre-subscribe to catch the 2 followup responses
	subscriptionRequest, err := m.connection.SubscriptionRequestBuilder().
		AddEventTagAddress("installationMMIMonitor", "mmimonitor/*/NETWORK_CONTROL").
		Build()
	if err != nil {
		return nil, errors.Wrap(err, "Error subscribing to the installation MMI")
	}
	subCtx, subCtxCancel := context.WithTimeout(ctx, time.Second*2)
	defer subCtxCancel()
	subscriptionResult := <-subscriptionRequest.ExecuteWithContext(subCtx)
	if err := subscriptionResult.GetErr(); err != nil {
		return nil, errors.Wrap(err, "Error subscribing to the mmi")
	}
	response := subscriptionResult.GetResponse()
	if responseCode := response.GetResponseCode("installationMMIMonitor"); responseCode != apiModel.PlcResponseCode_OK {
		return nil, errors.Errorf("Got %s", responseCode)
	}
	subscriptionHandle, err := response.GetSubscriptionHandle("installationMMIMonitor")
	if err != nil {
		return nil, errors.Wrap(err, "Error getting the subscription handle")
	}

	blockOffset0Received := false
	blockOffset0ReceivedChan := make(chan any, 100) // We only expect one, but we make it a bit bigger to no clog up
	blockOffset88Received := false
	blockOffset88ReceivedChan := make(chan any, 100) // We only expect one, but we make it a bit bigger to no clog up
	blockOffset176Received := false
	blockOffset176ReceivedChan := make(chan any, 100) // We only expect one, but we make it a bit bigger to no clog up
	result := make(map[byte]any)
	plcConsumerRegistration := subscriptionHandle.Register(func(event apiModel.PlcSubscriptionEvent) {
		m.log.Trace().Msgf("handling event:\n%s", event)
		if responseCode := event.GetResponseCode("installationMMIMonitor"); responseCode != apiModel.PlcResponseCode_OK {
			m.log.Warn().Msgf("Ignoring %v", event)
			return
		}
		rootValue := event.GetValue("installationMMIMonitor")
		if !rootValue.IsStruct() {
			m.log.Warn().Msgf("Ignoring %v should be a struct", rootValue)
			return
		}
		rootStruct := rootValue.GetStruct()
		if applicationValue := rootStruct["application"]; applicationValue == nil || !applicationValue.IsString() || applicationValue.GetString() != "NETWORK_CONTROL" {
			m.log.Warn().Msgf("Ignoring %v should contain a application tag of type string with value NETWORK_CONTROL", rootStruct)
			return
		}
		var blockStart int
		if blockStartValue := rootStruct["blockStart"]; blockStartValue == nil || !blockStartValue.IsByte() {
			m.log.Warn().Msgf("Ignoring %v should contain a blockStart tag of type byte", rootStruct)
			return
		} else {
			blockStart = int(blockStartValue.GetByte())
		}

		if plcListValue := rootStruct["values"]; plcListValue == nil || !plcListValue.IsList() {
			m.log.Warn().Msgf("Ignoring %v should contain a values tag of type list", rootStruct)
			return
		} else {
			for unitByteAddress, plcValue := range plcListValue.GetList() {
				unitByteAddress = blockStart + unitByteAddress
				if !plcValue.IsString() {
					m.log.Warn().Msgf("Ignoring %v at %d should be a string", plcValue, unitByteAddress)
					return
				}
				switch plcValue.GetString() {
				case readWriteModel.GAVState_ON.PLC4XEnumName(), readWriteModel.GAVState_OFF.PLC4XEnumName():
					m.log.Debug().Msgf("unit %d does exists", unitByteAddress)
					result[byte(unitByteAddress)] = true
				case readWriteModel.GAVState_DOES_NOT_EXIST.PLC4XEnumName():
					m.log.Debug().Msgf("unit %d does not exists", unitByteAddress)
				case readWriteModel.GAVState_ERROR.PLC4XEnumName():
					m.log.Warn().Msgf("unit %d is in error state", unitByteAddress)
				}
			}
		}
		// We notify here so we don't exit to early
		switch blockStart {
		case 0:
			select {
			case blockOffset0ReceivedChan <- true:
				m.log.Trace().Msg("0 notified")
			default:
				m.log.Warn().Msg("0 blocked")
			}
		case 88:
			select {
			case blockOffset88ReceivedChan <- true:
				m.log.Trace().Msg("88 notified")
			default:
				m.log.Warn().Msg("88 blocked")
			}
		case 176:
			select {
			case blockOffset176ReceivedChan <- true:
				m.log.Trace().Msg("176 notified")
			default:
				m.log.Warn().Msg("176 blocked")
			}
		}
	})
	defer plcConsumerRegistration.Unregister()

	readRequest, err := m.connection.ReadRequestBuilder().
		AddTagAddress("installationMMI", "status/binary/0xFF").
		Build()
	if err != nil {
		return nil, errors.Wrap(err, "Error getting the installation MMI")
	}
	readCtx, readCtxCancel := context.WithTimeout(ctx, time.Second*2)
	defer readCtxCancel()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				m.log.Error().Msgf("panic-ed %v. Stack:\n%s", err, debug.Stack())
			}
		}()
		defer readCtxCancel()
		readRequestResult := <-readRequest.ExecuteWithContext(readCtx)
		if err := readRequestResult.GetErr(); err != nil {
			m.log.Warn().Err(err).Msg("Error reading the mmi")
			return
		}
		response := readRequestResult.GetResponse()
		if responseCode := response.GetResponseCode("installationMMI"); responseCode == apiModel.PlcResponseCode_OK {
			rootValue := response.GetValue("installationMMI")
			if !rootValue.IsStruct() {
				m.log.Warn().Err(err).Msgf("%v should be a struct", rootValue)
				return
			}
			rootStruct := rootValue.GetStruct()
			if applicationValue := rootStruct["application"]; applicationValue == nil || !applicationValue.IsString() || applicationValue.GetString() != "NETWORK_CONTROL" {
				m.log.Warn().Err(err).Msgf("%v should contain a application tag of type string with value NETWORK_CONTROL", rootStruct)
				return
			}
			var blockStart int
			if blockStartValue := rootStruct["blockStart"]; blockStartValue == nil || !blockStartValue.IsByte() || blockStartValue.GetByte() != 0 {
				m.log.Warn().Err(err).Msgf("%v should contain a blockStart tag of type byte with value 0", rootStruct)
				return
			} else {
				blockStart = int(blockStartValue.GetByte())
			}
			m.log.Debug().Msgf("Read MMI with block start %d", blockStart)

			if plcListValue := rootStruct["values"]; plcListValue == nil || !plcListValue.IsList() {
				m.log.Warn().Err(err).Msgf("%v should contain a values tag of type list", rootStruct)
				return
			} else {
				for unitByteAddress, plcValue := range plcListValue.GetList() {
					unitByteAddress = blockStart + unitByteAddress
					if !plcValue.IsString() {
						m.log.Warn().Err(err).Msgf("%v at %d should be a string", plcValue, unitByteAddress)
						return
					}
					switch plcValue.GetString() {
					case readWriteModel.GAVState_ON.PLC4XEnumName(), readWriteModel.GAVState_OFF.PLC4XEnumName():
						m.log.Debug().Msgf("unit %d does exists", unitByteAddress)
						result[byte(unitByteAddress)] = true
					case readWriteModel.GAVState_DOES_NOT_EXIST.PLC4XEnumName():
						m.log.Debug().Msgf("unit %d does not exists", unitByteAddress)
					case readWriteModel.GAVState_ERROR.PLC4XEnumName():
						m.log.Warn().Msgf("unit %d is in error state", unitByteAddress)
					}
				}
			}
			switch blockStart {
			case 0:
				blockOffset0Received = true
				m.log.Trace().Msg("block 0 read by read")
			case 88:
				blockOffset88Received = true
				m.log.Trace().Msg("block 88 read by read")
			case 176:
				blockOffset176Received = true
				m.log.Trace().Msg("block 176 read by read")
			}

		} else {
			m.log.Warn().Msgf("We got %s as response code for installation mmi so we rely on getting it via subscription", responseCode)
		}
	}()

	syncCtx, syncCtxCancel := context.WithTimeout(ctx, time.Second*6)
	defer syncCtxCancel()
	for !blockOffset0Received || !blockOffset88Received || !blockOffset176Received {
		select {
		case <-blockOffset0ReceivedChan:
			m.log.Trace().Msg("Offset 0 received")
			blockOffset0Received = true
		case <-blockOffset88ReceivedChan:
			m.log.Trace().Msg("Offset 88 received")
			blockOffset88Received = true
		case <-blockOffset176ReceivedChan:
			m.log.Trace().Msg("Offset 176 received")
			blockOffset176Received = true
		case <-syncCtx.Done():
			err = syncCtx.Err()
			m.log.Trace().Err(err).Msg("Ending prematurely")
			return nil, errors.Wrap(err, "error waiting for other offsets")
		}
	}
	return result, nil
}
