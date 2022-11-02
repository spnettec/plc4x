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

package model

import (
	"encoding/binary"

	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Language is an enum
type Language uint8

type ILanguage interface {
	utils.Serializable
}

const(
	Language_NO_LANGUAGE Language = 0x00
	Language_ENGLISH Language = 0x01
	Language_ENGLISH_AUSTRALIA Language = 0x02
	Language_ENGLISH_BELIZE Language = 0x03
	Language_ENGLISH_CANADA Language = 0x04
	Language_ENGLISH_CARRIBEAN Language = 0x05
	Language_ENGLISH_IRELAND Language = 0x06
	Language_ENGLISH_JAMAICA Language = 0x07
	Language_ENGLISH_NEW_ZEALAND Language = 0x08
	Language_ENGLISH_PHILIPPINES Language = 0x09
	Language_ENGLISH_SOUTH_AFRICA Language = 0x0A
	Language_ENGLISH_TRINIDAD Language = 0x0B
	Language_ENGLISH_UK Language = 0x0C
	Language_ENGLISH_USA Language = 0x0D
	Language_ENGLISH_ZIMBABWE Language = 0x0E
	Language_AFRIKAANS Language = 0x40
	Language_BASQUE Language = 0x41
	Language_CATALAN Language = 0x42
	Language_DANISH Language = 0x43
	Language_DUTCH_BELGIUM Language = 0x44
	Language_DUTCH_NETHERLANDS Language = 0x45
	Language_FAEROESE Language = 0x46
	Language_FINNISH Language = 0x47
	Language_FRENCH_BELGIUM Language = 0x48
	Language_FRENCH_CANADA Language = 0x49
	Language_FRENCH Language = 0x4A
	Language_FRENCH_LUXEMBOURG Language = 0x4B
	Language_FRENCH_MONACO Language = 0x4C
	Language_FRENCH_SWITZERLAND Language = 0x4D
	Language_GALICIAN Language = 0x4E
	Language_GERMAN_AUSTRIA Language = 0x4F
	Language_GERMAN Language = 0x50
	Language_GERMAN_LIECHTENSTEIN Language = 0x51
	Language_GERMAN_LUXEMBOURG Language = 0x52
	Language_GERMAN_SWITZERLAND Language = 0x53
	Language_ICELANDIC Language = 0x54
	Language_INDONESIAN Language = 0x55
	Language_ITALIAN Language = 0x56
	Language_ITALIAN_SWITZERLAND Language = 0x57
	Language_MALAY_BRUNEI Language = 0x58
	Language_MALAY Language = 0x59
	Language_NORWEGIAN Language = 0x5A
	Language_NORWEGIAN_NYNORSK Language = 0x5B
	Language_PORTUGUESE_BRAZIL Language = 0x5C
	Language_PORTUGUESE Language = 0x5D
	Language_SPANISH_ARGENTINE Language = 0x5E
	Language_SPANISH_BOLIVIA Language = 0x5F
	Language_SPANISH_CHILE Language = 0x60
	Language_SPANISH_COLOMBIA Language = 0x61
	Language_SPANISH_COSTA_RICA Language = 0x62
	Language_SPANISH_DOMINICAN_REPUBLIC Language = 0x63
	Language_SPANISH_ECUADOR Language = 0x64
	Language_SPANISH_EL_SALVADOR Language = 0x65
	Language_SPANISH_GUATEMALA Language = 0x66
	Language_SPANISH_HONDURAS Language = 0x67
	Language_SPANISH Language = 0x68
	Language_SPANISH_MEXICO Language = 0x69
	Language_SPANISH_NICARAGUA Language = 0x6A
	Language_SPANISH_PANAMA Language = 0x6B
	Language_SPANISH_PARAGUAY Language = 0x6C
	Language_SPANISH_PERU Language = 0x6D
	Language_SPANISH_PERTO_RICO Language = 0x6E
	Language_SPANISH_TRADITIONAL Language = 0x6F
	Language_SPANISH_URUGUAY Language = 0x70
	Language_SPANISH_VENEZUELA Language = 0x71
	Language_SWAHILI Language = 0x72
	Language_SWEDISH Language = 0x73
	Language_SWEDISH_FINLAND Language = 0x74
	Language_CHINESE_CP936 Language = 0xCA
)

var LanguageValues []Language

func init() {
	_ = errors.New
	LanguageValues = []Language {
		Language_NO_LANGUAGE,
		Language_ENGLISH,
		Language_ENGLISH_AUSTRALIA,
		Language_ENGLISH_BELIZE,
		Language_ENGLISH_CANADA,
		Language_ENGLISH_CARRIBEAN,
		Language_ENGLISH_IRELAND,
		Language_ENGLISH_JAMAICA,
		Language_ENGLISH_NEW_ZEALAND,
		Language_ENGLISH_PHILIPPINES,
		Language_ENGLISH_SOUTH_AFRICA,
		Language_ENGLISH_TRINIDAD,
		Language_ENGLISH_UK,
		Language_ENGLISH_USA,
		Language_ENGLISH_ZIMBABWE,
		Language_AFRIKAANS,
		Language_BASQUE,
		Language_CATALAN,
		Language_DANISH,
		Language_DUTCH_BELGIUM,
		Language_DUTCH_NETHERLANDS,
		Language_FAEROESE,
		Language_FINNISH,
		Language_FRENCH_BELGIUM,
		Language_FRENCH_CANADA,
		Language_FRENCH,
		Language_FRENCH_LUXEMBOURG,
		Language_FRENCH_MONACO,
		Language_FRENCH_SWITZERLAND,
		Language_GALICIAN,
		Language_GERMAN_AUSTRIA,
		Language_GERMAN,
		Language_GERMAN_LIECHTENSTEIN,
		Language_GERMAN_LUXEMBOURG,
		Language_GERMAN_SWITZERLAND,
		Language_ICELANDIC,
		Language_INDONESIAN,
		Language_ITALIAN,
		Language_ITALIAN_SWITZERLAND,
		Language_MALAY_BRUNEI,
		Language_MALAY,
		Language_NORWEGIAN,
		Language_NORWEGIAN_NYNORSK,
		Language_PORTUGUESE_BRAZIL,
		Language_PORTUGUESE,
		Language_SPANISH_ARGENTINE,
		Language_SPANISH_BOLIVIA,
		Language_SPANISH_CHILE,
		Language_SPANISH_COLOMBIA,
		Language_SPANISH_COSTA_RICA,
		Language_SPANISH_DOMINICAN_REPUBLIC,
		Language_SPANISH_ECUADOR,
		Language_SPANISH_EL_SALVADOR,
		Language_SPANISH_GUATEMALA,
		Language_SPANISH_HONDURAS,
		Language_SPANISH,
		Language_SPANISH_MEXICO,
		Language_SPANISH_NICARAGUA,
		Language_SPANISH_PANAMA,
		Language_SPANISH_PARAGUAY,
		Language_SPANISH_PERU,
		Language_SPANISH_PERTO_RICO,
		Language_SPANISH_TRADITIONAL,
		Language_SPANISH_URUGUAY,
		Language_SPANISH_VENEZUELA,
		Language_SWAHILI,
		Language_SWEDISH,
		Language_SWEDISH_FINLAND,
		Language_CHINESE_CP936,
	}
}

func LanguageByValue(value uint8) (enum Language, ok bool) {
	switch value {
		case 0x00:
			return Language_NO_LANGUAGE, true
		case 0x01:
			return Language_ENGLISH, true
		case 0x02:
			return Language_ENGLISH_AUSTRALIA, true
		case 0x03:
			return Language_ENGLISH_BELIZE, true
		case 0x04:
			return Language_ENGLISH_CANADA, true
		case 0x05:
			return Language_ENGLISH_CARRIBEAN, true
		case 0x06:
			return Language_ENGLISH_IRELAND, true
		case 0x07:
			return Language_ENGLISH_JAMAICA, true
		case 0x08:
			return Language_ENGLISH_NEW_ZEALAND, true
		case 0x09:
			return Language_ENGLISH_PHILIPPINES, true
		case 0x0A:
			return Language_ENGLISH_SOUTH_AFRICA, true
		case 0x0B:
			return Language_ENGLISH_TRINIDAD, true
		case 0x0C:
			return Language_ENGLISH_UK, true
		case 0x0D:
			return Language_ENGLISH_USA, true
		case 0x0E:
			return Language_ENGLISH_ZIMBABWE, true
		case 0x40:
			return Language_AFRIKAANS, true
		case 0x41:
			return Language_BASQUE, true
		case 0x42:
			return Language_CATALAN, true
		case 0x43:
			return Language_DANISH, true
		case 0x44:
			return Language_DUTCH_BELGIUM, true
		case 0x45:
			return Language_DUTCH_NETHERLANDS, true
		case 0x46:
			return Language_FAEROESE, true
		case 0x47:
			return Language_FINNISH, true
		case 0x48:
			return Language_FRENCH_BELGIUM, true
		case 0x49:
			return Language_FRENCH_CANADA, true
		case 0x4A:
			return Language_FRENCH, true
		case 0x4B:
			return Language_FRENCH_LUXEMBOURG, true
		case 0x4C:
			return Language_FRENCH_MONACO, true
		case 0x4D:
			return Language_FRENCH_SWITZERLAND, true
		case 0x4E:
			return Language_GALICIAN, true
		case 0x4F:
			return Language_GERMAN_AUSTRIA, true
		case 0x50:
			return Language_GERMAN, true
		case 0x51:
			return Language_GERMAN_LIECHTENSTEIN, true
		case 0x52:
			return Language_GERMAN_LUXEMBOURG, true
		case 0x53:
			return Language_GERMAN_SWITZERLAND, true
		case 0x54:
			return Language_ICELANDIC, true
		case 0x55:
			return Language_INDONESIAN, true
		case 0x56:
			return Language_ITALIAN, true
		case 0x57:
			return Language_ITALIAN_SWITZERLAND, true
		case 0x58:
			return Language_MALAY_BRUNEI, true
		case 0x59:
			return Language_MALAY, true
		case 0x5A:
			return Language_NORWEGIAN, true
		case 0x5B:
			return Language_NORWEGIAN_NYNORSK, true
		case 0x5C:
			return Language_PORTUGUESE_BRAZIL, true
		case 0x5D:
			return Language_PORTUGUESE, true
		case 0x5E:
			return Language_SPANISH_ARGENTINE, true
		case 0x5F:
			return Language_SPANISH_BOLIVIA, true
		case 0x60:
			return Language_SPANISH_CHILE, true
		case 0x61:
			return Language_SPANISH_COLOMBIA, true
		case 0x62:
			return Language_SPANISH_COSTA_RICA, true
		case 0x63:
			return Language_SPANISH_DOMINICAN_REPUBLIC, true
		case 0x64:
			return Language_SPANISH_ECUADOR, true
		case 0x65:
			return Language_SPANISH_EL_SALVADOR, true
		case 0x66:
			return Language_SPANISH_GUATEMALA, true
		case 0x67:
			return Language_SPANISH_HONDURAS, true
		case 0x68:
			return Language_SPANISH, true
		case 0x69:
			return Language_SPANISH_MEXICO, true
		case 0x6A:
			return Language_SPANISH_NICARAGUA, true
		case 0x6B:
			return Language_SPANISH_PANAMA, true
		case 0x6C:
			return Language_SPANISH_PARAGUAY, true
		case 0x6D:
			return Language_SPANISH_PERU, true
		case 0x6E:
			return Language_SPANISH_PERTO_RICO, true
		case 0x6F:
			return Language_SPANISH_TRADITIONAL, true
		case 0x70:
			return Language_SPANISH_URUGUAY, true
		case 0x71:
			return Language_SPANISH_VENEZUELA, true
		case 0x72:
			return Language_SWAHILI, true
		case 0x73:
			return Language_SWEDISH, true
		case 0x74:
			return Language_SWEDISH_FINLAND, true
		case 0xCA:
			return Language_CHINESE_CP936, true
	}
	return 0, false
}

func LanguageByName(value string) (enum Language, ok bool) {
	switch value {
	case "NO_LANGUAGE":
		return Language_NO_LANGUAGE, true
	case "ENGLISH":
		return Language_ENGLISH, true
	case "ENGLISH_AUSTRALIA":
		return Language_ENGLISH_AUSTRALIA, true
	case "ENGLISH_BELIZE":
		return Language_ENGLISH_BELIZE, true
	case "ENGLISH_CANADA":
		return Language_ENGLISH_CANADA, true
	case "ENGLISH_CARRIBEAN":
		return Language_ENGLISH_CARRIBEAN, true
	case "ENGLISH_IRELAND":
		return Language_ENGLISH_IRELAND, true
	case "ENGLISH_JAMAICA":
		return Language_ENGLISH_JAMAICA, true
	case "ENGLISH_NEW_ZEALAND":
		return Language_ENGLISH_NEW_ZEALAND, true
	case "ENGLISH_PHILIPPINES":
		return Language_ENGLISH_PHILIPPINES, true
	case "ENGLISH_SOUTH_AFRICA":
		return Language_ENGLISH_SOUTH_AFRICA, true
	case "ENGLISH_TRINIDAD":
		return Language_ENGLISH_TRINIDAD, true
	case "ENGLISH_UK":
		return Language_ENGLISH_UK, true
	case "ENGLISH_USA":
		return Language_ENGLISH_USA, true
	case "ENGLISH_ZIMBABWE":
		return Language_ENGLISH_ZIMBABWE, true
	case "AFRIKAANS":
		return Language_AFRIKAANS, true
	case "BASQUE":
		return Language_BASQUE, true
	case "CATALAN":
		return Language_CATALAN, true
	case "DANISH":
		return Language_DANISH, true
	case "DUTCH_BELGIUM":
		return Language_DUTCH_BELGIUM, true
	case "DUTCH_NETHERLANDS":
		return Language_DUTCH_NETHERLANDS, true
	case "FAEROESE":
		return Language_FAEROESE, true
	case "FINNISH":
		return Language_FINNISH, true
	case "FRENCH_BELGIUM":
		return Language_FRENCH_BELGIUM, true
	case "FRENCH_CANADA":
		return Language_FRENCH_CANADA, true
	case "FRENCH":
		return Language_FRENCH, true
	case "FRENCH_LUXEMBOURG":
		return Language_FRENCH_LUXEMBOURG, true
	case "FRENCH_MONACO":
		return Language_FRENCH_MONACO, true
	case "FRENCH_SWITZERLAND":
		return Language_FRENCH_SWITZERLAND, true
	case "GALICIAN":
		return Language_GALICIAN, true
	case "GERMAN_AUSTRIA":
		return Language_GERMAN_AUSTRIA, true
	case "GERMAN":
		return Language_GERMAN, true
	case "GERMAN_LIECHTENSTEIN":
		return Language_GERMAN_LIECHTENSTEIN, true
	case "GERMAN_LUXEMBOURG":
		return Language_GERMAN_LUXEMBOURG, true
	case "GERMAN_SWITZERLAND":
		return Language_GERMAN_SWITZERLAND, true
	case "ICELANDIC":
		return Language_ICELANDIC, true
	case "INDONESIAN":
		return Language_INDONESIAN, true
	case "ITALIAN":
		return Language_ITALIAN, true
	case "ITALIAN_SWITZERLAND":
		return Language_ITALIAN_SWITZERLAND, true
	case "MALAY_BRUNEI":
		return Language_MALAY_BRUNEI, true
	case "MALAY":
		return Language_MALAY, true
	case "NORWEGIAN":
		return Language_NORWEGIAN, true
	case "NORWEGIAN_NYNORSK":
		return Language_NORWEGIAN_NYNORSK, true
	case "PORTUGUESE_BRAZIL":
		return Language_PORTUGUESE_BRAZIL, true
	case "PORTUGUESE":
		return Language_PORTUGUESE, true
	case "SPANISH_ARGENTINE":
		return Language_SPANISH_ARGENTINE, true
	case "SPANISH_BOLIVIA":
		return Language_SPANISH_BOLIVIA, true
	case "SPANISH_CHILE":
		return Language_SPANISH_CHILE, true
	case "SPANISH_COLOMBIA":
		return Language_SPANISH_COLOMBIA, true
	case "SPANISH_COSTA_RICA":
		return Language_SPANISH_COSTA_RICA, true
	case "SPANISH_DOMINICAN_REPUBLIC":
		return Language_SPANISH_DOMINICAN_REPUBLIC, true
	case "SPANISH_ECUADOR":
		return Language_SPANISH_ECUADOR, true
	case "SPANISH_EL_SALVADOR":
		return Language_SPANISH_EL_SALVADOR, true
	case "SPANISH_GUATEMALA":
		return Language_SPANISH_GUATEMALA, true
	case "SPANISH_HONDURAS":
		return Language_SPANISH_HONDURAS, true
	case "SPANISH":
		return Language_SPANISH, true
	case "SPANISH_MEXICO":
		return Language_SPANISH_MEXICO, true
	case "SPANISH_NICARAGUA":
		return Language_SPANISH_NICARAGUA, true
	case "SPANISH_PANAMA":
		return Language_SPANISH_PANAMA, true
	case "SPANISH_PARAGUAY":
		return Language_SPANISH_PARAGUAY, true
	case "SPANISH_PERU":
		return Language_SPANISH_PERU, true
	case "SPANISH_PERTO_RICO":
		return Language_SPANISH_PERTO_RICO, true
	case "SPANISH_TRADITIONAL":
		return Language_SPANISH_TRADITIONAL, true
	case "SPANISH_URUGUAY":
		return Language_SPANISH_URUGUAY, true
	case "SPANISH_VENEZUELA":
		return Language_SPANISH_VENEZUELA, true
	case "SWAHILI":
		return Language_SWAHILI, true
	case "SWEDISH":
		return Language_SWEDISH, true
	case "SWEDISH_FINLAND":
		return Language_SWEDISH_FINLAND, true
	case "CHINESE_CP936":
		return Language_CHINESE_CP936, true
	}
	return 0, false
}

func LanguageKnows(value uint8)  bool {
	for _, typeValue := range LanguageValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastLanguage(structType interface{}) Language {
	castFunc := func(typ interface{}) Language {
		if sLanguage, ok := typ.(Language); ok {
			return sLanguage
		}
		return 0
	}
	return castFunc(structType)
}

func (m Language) GetLengthInBits() uint16 {
	return 8
}

func (m Language) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LanguageParse(readBuffer utils.ReadBuffer) (Language, error) {
	val, err := readBuffer.ReadUint8("Language", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading Language")
	}
	if enum, ok := LanguageByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return Language(val), nil
	} else {
		return enum, nil
	}
}

func (e Language) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian)) // TODO: get endianness from mspec
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e Language) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("Language", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e Language) PLC4XEnumName() string {
	switch e {
	case Language_NO_LANGUAGE:
		return "NO_LANGUAGE"
	case Language_ENGLISH:
		return "ENGLISH"
	case Language_ENGLISH_AUSTRALIA:
		return "ENGLISH_AUSTRALIA"
	case Language_ENGLISH_BELIZE:
		return "ENGLISH_BELIZE"
	case Language_ENGLISH_CANADA:
		return "ENGLISH_CANADA"
	case Language_ENGLISH_CARRIBEAN:
		return "ENGLISH_CARRIBEAN"
	case Language_ENGLISH_IRELAND:
		return "ENGLISH_IRELAND"
	case Language_ENGLISH_JAMAICA:
		return "ENGLISH_JAMAICA"
	case Language_ENGLISH_NEW_ZEALAND:
		return "ENGLISH_NEW_ZEALAND"
	case Language_ENGLISH_PHILIPPINES:
		return "ENGLISH_PHILIPPINES"
	case Language_ENGLISH_SOUTH_AFRICA:
		return "ENGLISH_SOUTH_AFRICA"
	case Language_ENGLISH_TRINIDAD:
		return "ENGLISH_TRINIDAD"
	case Language_ENGLISH_UK:
		return "ENGLISH_UK"
	case Language_ENGLISH_USA:
		return "ENGLISH_USA"
	case Language_ENGLISH_ZIMBABWE:
		return "ENGLISH_ZIMBABWE"
	case Language_AFRIKAANS:
		return "AFRIKAANS"
	case Language_BASQUE:
		return "BASQUE"
	case Language_CATALAN:
		return "CATALAN"
	case Language_DANISH:
		return "DANISH"
	case Language_DUTCH_BELGIUM:
		return "DUTCH_BELGIUM"
	case Language_DUTCH_NETHERLANDS:
		return "DUTCH_NETHERLANDS"
	case Language_FAEROESE:
		return "FAEROESE"
	case Language_FINNISH:
		return "FINNISH"
	case Language_FRENCH_BELGIUM:
		return "FRENCH_BELGIUM"
	case Language_FRENCH_CANADA:
		return "FRENCH_CANADA"
	case Language_FRENCH:
		return "FRENCH"
	case Language_FRENCH_LUXEMBOURG:
		return "FRENCH_LUXEMBOURG"
	case Language_FRENCH_MONACO:
		return "FRENCH_MONACO"
	case Language_FRENCH_SWITZERLAND:
		return "FRENCH_SWITZERLAND"
	case Language_GALICIAN:
		return "GALICIAN"
	case Language_GERMAN_AUSTRIA:
		return "GERMAN_AUSTRIA"
	case Language_GERMAN:
		return "GERMAN"
	case Language_GERMAN_LIECHTENSTEIN:
		return "GERMAN_LIECHTENSTEIN"
	case Language_GERMAN_LUXEMBOURG:
		return "GERMAN_LUXEMBOURG"
	case Language_GERMAN_SWITZERLAND:
		return "GERMAN_SWITZERLAND"
	case Language_ICELANDIC:
		return "ICELANDIC"
	case Language_INDONESIAN:
		return "INDONESIAN"
	case Language_ITALIAN:
		return "ITALIAN"
	case Language_ITALIAN_SWITZERLAND:
		return "ITALIAN_SWITZERLAND"
	case Language_MALAY_BRUNEI:
		return "MALAY_BRUNEI"
	case Language_MALAY:
		return "MALAY"
	case Language_NORWEGIAN:
		return "NORWEGIAN"
	case Language_NORWEGIAN_NYNORSK:
		return "NORWEGIAN_NYNORSK"
	case Language_PORTUGUESE_BRAZIL:
		return "PORTUGUESE_BRAZIL"
	case Language_PORTUGUESE:
		return "PORTUGUESE"
	case Language_SPANISH_ARGENTINE:
		return "SPANISH_ARGENTINE"
	case Language_SPANISH_BOLIVIA:
		return "SPANISH_BOLIVIA"
	case Language_SPANISH_CHILE:
		return "SPANISH_CHILE"
	case Language_SPANISH_COLOMBIA:
		return "SPANISH_COLOMBIA"
	case Language_SPANISH_COSTA_RICA:
		return "SPANISH_COSTA_RICA"
	case Language_SPANISH_DOMINICAN_REPUBLIC:
		return "SPANISH_DOMINICAN_REPUBLIC"
	case Language_SPANISH_ECUADOR:
		return "SPANISH_ECUADOR"
	case Language_SPANISH_EL_SALVADOR:
		return "SPANISH_EL_SALVADOR"
	case Language_SPANISH_GUATEMALA:
		return "SPANISH_GUATEMALA"
	case Language_SPANISH_HONDURAS:
		return "SPANISH_HONDURAS"
	case Language_SPANISH:
		return "SPANISH"
	case Language_SPANISH_MEXICO:
		return "SPANISH_MEXICO"
	case Language_SPANISH_NICARAGUA:
		return "SPANISH_NICARAGUA"
	case Language_SPANISH_PANAMA:
		return "SPANISH_PANAMA"
	case Language_SPANISH_PARAGUAY:
		return "SPANISH_PARAGUAY"
	case Language_SPANISH_PERU:
		return "SPANISH_PERU"
	case Language_SPANISH_PERTO_RICO:
		return "SPANISH_PERTO_RICO"
	case Language_SPANISH_TRADITIONAL:
		return "SPANISH_TRADITIONAL"
	case Language_SPANISH_URUGUAY:
		return "SPANISH_URUGUAY"
	case Language_SPANISH_VENEZUELA:
		return "SPANISH_VENEZUELA"
	case Language_SWAHILI:
		return "SWAHILI"
	case Language_SWEDISH:
		return "SWEDISH"
	case Language_SWEDISH_FINLAND:
		return "SWEDISH_FINLAND"
	case Language_CHINESE_CP936:
		return "CHINESE_CP936"
	}
	return ""
}

func (e Language) String() string {
	return e.PLC4XEnumName()
}

