/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// ExceptionIdAnyOf the model 'ExceptionIdAnyOf'
type ExceptionIdAnyOf string

// List of ExceptionId_anyOf
const (
	UNEXPECTED_UE_LOCATION ExceptionIdAnyOf = "UNEXPECTED_UE_LOCATION"
	UNEXPECTED_LONG_LIVE_FLOW ExceptionIdAnyOf = "UNEXPECTED_LONG_LIVE_FLOW"
	UNEXPECTED_LARGE_RATE_FLOW ExceptionIdAnyOf = "UNEXPECTED_LARGE_RATE_FLOW"
	UNEXPECTED_WAKEUP ExceptionIdAnyOf = "UNEXPECTED_WAKEUP"
	SUSPICION_OF_DDOS_ATTACK ExceptionIdAnyOf = "SUSPICION_OF_DDOS_ATTACK"
	WRONG_DESTINATION_ADDRESS ExceptionIdAnyOf = "WRONG_DESTINATION_ADDRESS"
	TOO_FREQUENT_SERVICE_ACCESS ExceptionIdAnyOf = "TOO_FREQUENT_SERVICE_ACCESS"
	UNEXPECTED_RADIO_LINK_FAILURES ExceptionIdAnyOf = "UNEXPECTED_RADIO_LINK_FAILURES"
	PING_PONG_ACROSS_CELLS ExceptionIdAnyOf = "PING_PONG_ACROSS_CELLS"
)

// All allowed values of ExceptionIdAnyOf enum
var AllowedExceptionIdAnyOfEnumValues = []ExceptionIdAnyOf{
	"UNEXPECTED_UE_LOCATION",
	"UNEXPECTED_LONG_LIVE_FLOW",
	"UNEXPECTED_LARGE_RATE_FLOW",
	"UNEXPECTED_WAKEUP",
	"SUSPICION_OF_DDOS_ATTACK",
	"WRONG_DESTINATION_ADDRESS",
	"TOO_FREQUENT_SERVICE_ACCESS",
	"UNEXPECTED_RADIO_LINK_FAILURES",
	"PING_PONG_ACROSS_CELLS",
}

func (v *ExceptionIdAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExceptionIdAnyOf(value)
	for _, existing := range AllowedExceptionIdAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExceptionIdAnyOf", value)
}

// NewExceptionIdAnyOfFromValue returns a pointer to a valid ExceptionIdAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExceptionIdAnyOfFromValue(v string) (*ExceptionIdAnyOf, error) {
	ev := ExceptionIdAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExceptionIdAnyOf: valid values are %v", v, AllowedExceptionIdAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExceptionIdAnyOf) IsValid() bool {
	for _, existing := range AllowedExceptionIdAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExceptionId_anyOf value
func (v ExceptionIdAnyOf) Ptr() *ExceptionIdAnyOf {
	return &v
}

type NullableExceptionIdAnyOf struct {
	value *ExceptionIdAnyOf
	isSet bool
}

func (v NullableExceptionIdAnyOf) Get() *ExceptionIdAnyOf {
	return v.value
}

func (v *NullableExceptionIdAnyOf) Set(val *ExceptionIdAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionIdAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionIdAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionIdAnyOf(val *ExceptionIdAnyOf) *NullableExceptionIdAnyOf {
	return &NullableExceptionIdAnyOf{value: val, isSet: true}
}

func (v NullableExceptionIdAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionIdAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

