/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
	"fmt"
)

// AnalyticsSubsetAnyOf the model 'AnalyticsSubsetAnyOf'
type AnalyticsSubsetAnyOf string

// List of AnalyticsSubset_anyOf
const (
	NUM_OF_UE_REG AnalyticsSubsetAnyOf = "NUM_OF_UE_REG"
	NUM_OF_PDU_SESS_ESTBL AnalyticsSubsetAnyOf = "NUM_OF_PDU_SESS_ESTBL"
	RES_USAGE AnalyticsSubsetAnyOf = "RES_USAGE"
	NUM_OF_EXCEED_RES_USAGE_LOAD_LEVEL_THR AnalyticsSubsetAnyOf = "NUM_OF_EXCEED_RES_USAGE_LOAD_LEVEL_THR"
	PERIOD_OF_EXCEED_RES_USAGE_LOAD_LEVEL_THR AnalyticsSubsetAnyOf = "PERIOD_OF_EXCEED_RES_USAGE_LOAD_LEVEL_THR"
	EXCEED_LOAD_LEVEL_THR_IND AnalyticsSubsetAnyOf = "EXCEED_LOAD_LEVEL_THR_IND"
	LIST_OF_TOP_APP_UL AnalyticsSubsetAnyOf = "LIST_OF_TOP_APP_UL"
	LIST_OF_TOP_APP_DL AnalyticsSubsetAnyOf = "LIST_OF_TOP_APP_DL"
	NF_STATUS AnalyticsSubsetAnyOf = "NF_STATUS"
	NF_RESOURCE_USAGE AnalyticsSubsetAnyOf = "NF_RESOURCE_USAGE"
	NF_LOAD AnalyticsSubsetAnyOf = "NF_LOAD"
	NF_PEAK_LOAD AnalyticsSubsetAnyOf = "NF_PEAK_LOAD"
	NF_LOAD_AVG_IN_AOI AnalyticsSubsetAnyOf = "NF_LOAD_AVG_IN_AOI"
	DISPER_AMOUNT AnalyticsSubsetAnyOf = "DISPER_AMOUNT"
	DISPER_CLASS AnalyticsSubsetAnyOf = "DISPER_CLASS"
	RANKING AnalyticsSubsetAnyOf = "RANKING"
	PERCENTILE_RANKING AnalyticsSubsetAnyOf = "PERCENTILE_RANKING"
	RSSI AnalyticsSubsetAnyOf = "RSSI"
	RTT AnalyticsSubsetAnyOf = "RTT"
	TRAFFIC_INFO AnalyticsSubsetAnyOf = "TRAFFIC_INFO"
	NUMBER_OF_UES AnalyticsSubsetAnyOf = "NUMBER_OF_UES"
	APP_LIST_FOR_UE_COMM AnalyticsSubsetAnyOf = "APP_LIST_FOR_UE_COMM"
	N4_SESS_INACT_TIMER_FOR_UE_COMM AnalyticsSubsetAnyOf = "N4_SESS_INACT_TIMER_FOR_UE_COMM"
	AVG_TRAFFIC_RATE AnalyticsSubsetAnyOf = "AVG_TRAFFIC_RATE"
	MAX_TRAFFIC_RATE AnalyticsSubsetAnyOf = "MAX_TRAFFIC_RATE"
	AVG_PACKET_DELAY AnalyticsSubsetAnyOf = "AVG_PACKET_DELAY"
	MAX_PACKET_DELAY AnalyticsSubsetAnyOf = "MAX_PACKET_DELAY"
	AVG_PACKET_LOSS_RATE AnalyticsSubsetAnyOf = "AVG_PACKET_LOSS_RATE"
	UE_LOCATION AnalyticsSubsetAnyOf = "UE_LOCATION"
	LIST_OF_HIGH_EXP_UE AnalyticsSubsetAnyOf = "LIST_OF_HIGH_EXP_UE"
	LIST_OF_MEDIUM_EXP_UE AnalyticsSubsetAnyOf = "LIST_OF_MEDIUM_EXP_UE"
	LIST_OF_LOW_EXP_UE AnalyticsSubsetAnyOf = "LIST_OF_LOW_EXP_UE"
	AVG_UL_PKT_DROP_RATE AnalyticsSubsetAnyOf = "AVG_UL_PKT_DROP_RATE"
	VAR_UL_PKT_DROP_RATE AnalyticsSubsetAnyOf = "VAR_UL_PKT_DROP_RATE"
	AVG_DL_PKT_DROP_RATE AnalyticsSubsetAnyOf = "AVG_DL_PKT_DROP_RATE"
	VAR_DL_PKT_DROP_RATE AnalyticsSubsetAnyOf = "VAR_DL_PKT_DROP_RATE"
	AVG_UL_PKT_DELAY AnalyticsSubsetAnyOf = "AVG_UL_PKT_DELAY"
	VAR_UL_PKT_DELAY AnalyticsSubsetAnyOf = "VAR_UL_PKT_DELAY"
	AVG_DL_PKT_DELAY AnalyticsSubsetAnyOf = "AVG_DL_PKT_DELAY"
	VAR_DL_PKT_DELAY AnalyticsSubsetAnyOf = "VAR_DL_PKT_DELAY"
)

// All allowed values of AnalyticsSubsetAnyOf enum
var AllowedAnalyticsSubsetAnyOfEnumValues = []AnalyticsSubsetAnyOf{
	"NUM_OF_UE_REG",
	"NUM_OF_PDU_SESS_ESTBL",
	"RES_USAGE",
	"NUM_OF_EXCEED_RES_USAGE_LOAD_LEVEL_THR",
	"PERIOD_OF_EXCEED_RES_USAGE_LOAD_LEVEL_THR",
	"EXCEED_LOAD_LEVEL_THR_IND",
	"LIST_OF_TOP_APP_UL",
	"LIST_OF_TOP_APP_DL",
	"NF_STATUS",
	"NF_RESOURCE_USAGE",
	"NF_LOAD",
	"NF_PEAK_LOAD",
	"NF_LOAD_AVG_IN_AOI",
	"DISPER_AMOUNT",
	"DISPER_CLASS",
	"RANKING",
	"PERCENTILE_RANKING",
	"RSSI",
	"RTT",
	"TRAFFIC_INFO",
	"NUMBER_OF_UES",
	"APP_LIST_FOR_UE_COMM",
	"N4_SESS_INACT_TIMER_FOR_UE_COMM",
	"AVG_TRAFFIC_RATE",
	"MAX_TRAFFIC_RATE",
	"AVG_PACKET_DELAY",
	"MAX_PACKET_DELAY",
	"AVG_PACKET_LOSS_RATE",
	"UE_LOCATION",
	"LIST_OF_HIGH_EXP_UE",
	"LIST_OF_MEDIUM_EXP_UE",
	"LIST_OF_LOW_EXP_UE",
	"AVG_UL_PKT_DROP_RATE",
	"VAR_UL_PKT_DROP_RATE",
	"AVG_DL_PKT_DROP_RATE",
	"VAR_DL_PKT_DROP_RATE",
	"AVG_UL_PKT_DELAY",
	"VAR_UL_PKT_DELAY",
	"AVG_DL_PKT_DELAY",
	"VAR_DL_PKT_DELAY",
}

func (v *AnalyticsSubsetAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnalyticsSubsetAnyOf(value)
	for _, existing := range AllowedAnalyticsSubsetAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnalyticsSubsetAnyOf", value)
}

// NewAnalyticsSubsetAnyOfFromValue returns a pointer to a valid AnalyticsSubsetAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnalyticsSubsetAnyOfFromValue(v string) (*AnalyticsSubsetAnyOf, error) {
	ev := AnalyticsSubsetAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AnalyticsSubsetAnyOf: valid values are %v", v, AllowedAnalyticsSubsetAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AnalyticsSubsetAnyOf) IsValid() bool {
	for _, existing := range AllowedAnalyticsSubsetAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnalyticsSubset_anyOf value
func (v AnalyticsSubsetAnyOf) Ptr() *AnalyticsSubsetAnyOf {
	return &v
}

type NullableAnalyticsSubsetAnyOf struct {
	value *AnalyticsSubsetAnyOf
	isSet bool
}

func (v NullableAnalyticsSubsetAnyOf) Get() *AnalyticsSubsetAnyOf {
	return v.value
}

func (v *NullableAnalyticsSubsetAnyOf) Set(val *AnalyticsSubsetAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsSubsetAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsSubsetAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsSubsetAnyOf(val *AnalyticsSubsetAnyOf) *NullableAnalyticsSubsetAnyOf {
	return &NullableAnalyticsSubsetAnyOf{value: val, isSet: true}
}

func (v NullableAnalyticsSubsetAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsSubsetAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

