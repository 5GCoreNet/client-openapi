/*
Nmfaf_3daDataManagement

MFAF 3GPP DCCF Adaptor (3DA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3daDataManagement

import (
	"encoding/json"
	"fmt"
)

// NefEventAnyOf the model 'NefEventAnyOf'
type NefEventAnyOf string

// List of NefEvent_anyOf
const (
	SVC_EXPERIENCE NefEventAnyOf = "SVC_EXPERIENCE"
	UE_MOBILITY NefEventAnyOf = "UE_MOBILITY"
	UE_COMM NefEventAnyOf = "UE_COMM"
	EXCEPTIONS NefEventAnyOf = "EXCEPTIONS"
	USER_DATA_CONGESTION NefEventAnyOf = "USER_DATA_CONGESTION"
	PERF_DATA NefEventAnyOf = "PERF_DATA"
	DISPERSION NefEventAnyOf = "DISPERSION"
	COLLECTIVE_BEHAVIOUR NefEventAnyOf = "COLLECTIVE_BEHAVIOUR"
	MS_QOE_METRICS NefEventAnyOf = "MS_QOE_METRICS"
	MS_CONSUMPTION NefEventAnyOf = "MS_CONSUMPTION"
	MS_NET_ASSIST_INVOCATION NefEventAnyOf = "MS_NET_ASSIST_INVOCATION"
	MS_DYN_POLICY_INVOCATION NefEventAnyOf = "MS_DYN_POLICY_INVOCATION"
	MS_ACCESS_ACTIVITY NefEventAnyOf = "MS_ACCESS_ACTIVITY"
)

// All allowed values of NefEventAnyOf enum
var AllowedNefEventAnyOfEnumValues = []NefEventAnyOf{
	"SVC_EXPERIENCE",
	"UE_MOBILITY",
	"UE_COMM",
	"EXCEPTIONS",
	"USER_DATA_CONGESTION",
	"PERF_DATA",
	"DISPERSION",
	"COLLECTIVE_BEHAVIOUR",
	"MS_QOE_METRICS",
	"MS_CONSUMPTION",
	"MS_NET_ASSIST_INVOCATION",
	"MS_DYN_POLICY_INVOCATION",
	"MS_ACCESS_ACTIVITY",
}

func (v *NefEventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NefEventAnyOf(value)
	for _, existing := range AllowedNefEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NefEventAnyOf", value)
}

// NewNefEventAnyOfFromValue returns a pointer to a valid NefEventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNefEventAnyOfFromValue(v string) (*NefEventAnyOf, error) {
	ev := NefEventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NefEventAnyOf: valid values are %v", v, AllowedNefEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NefEventAnyOf) IsValid() bool {
	for _, existing := range AllowedNefEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NefEvent_anyOf value
func (v NefEventAnyOf) Ptr() *NefEventAnyOf {
	return &v
}

type NullableNefEventAnyOf struct {
	value *NefEventAnyOf
	isSet bool
}

func (v NullableNefEventAnyOf) Get() *NefEventAnyOf {
	return v.value
}

func (v *NullableNefEventAnyOf) Set(val *NefEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNefEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNefEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNefEventAnyOf(val *NefEventAnyOf) *NullableNefEventAnyOf {
	return &NullableNefEventAnyOf{value: val, isSet: true}
}

func (v NullableNefEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNefEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

