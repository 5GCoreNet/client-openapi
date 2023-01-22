/*
Ntsctsf_TimeSynchronization Service API

TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ntsctsf_TimeSynchronization

import (
	"encoding/json"
	"fmt"
)

// AsTimeResourceAnyOf the model 'AsTimeResourceAnyOf'
type AsTimeResourceAnyOf string

// List of AsTimeResource_anyOf
const (
	ATOMIC_CLOCK AsTimeResourceAnyOf = "ATOMIC_CLOCK"
	GNSS AsTimeResourceAnyOf = "GNSS"
	TERRESTRIAL_RADIO AsTimeResourceAnyOf = "TERRESTRIAL_RADIO"
	SERIAL_TIME_CODE AsTimeResourceAnyOf = "SERIAL_TIME_CODE"
	PTP AsTimeResourceAnyOf = "PTP"
	NTP AsTimeResourceAnyOf = "NTP"
	HAND_SET AsTimeResourceAnyOf = "HAND_SET"
	INTERNAL_OSCILLATOR AsTimeResourceAnyOf = "INTERNAL_OSCILLATOR"
	OTHER AsTimeResourceAnyOf = "OTHER"
)

// All allowed values of AsTimeResourceAnyOf enum
var AllowedAsTimeResourceAnyOfEnumValues = []AsTimeResourceAnyOf{
	"ATOMIC_CLOCK",
	"GNSS",
	"TERRESTRIAL_RADIO",
	"SERIAL_TIME_CODE",
	"PTP",
	"NTP",
	"HAND_SET",
	"INTERNAL_OSCILLATOR",
	"OTHER",
}

func (v *AsTimeResourceAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AsTimeResourceAnyOf(value)
	for _, existing := range AllowedAsTimeResourceAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AsTimeResourceAnyOf", value)
}

// NewAsTimeResourceAnyOfFromValue returns a pointer to a valid AsTimeResourceAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAsTimeResourceAnyOfFromValue(v string) (*AsTimeResourceAnyOf, error) {
	ev := AsTimeResourceAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AsTimeResourceAnyOf: valid values are %v", v, AllowedAsTimeResourceAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AsTimeResourceAnyOf) IsValid() bool {
	for _, existing := range AllowedAsTimeResourceAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AsTimeResource_anyOf value
func (v AsTimeResourceAnyOf) Ptr() *AsTimeResourceAnyOf {
	return &v
}

type NullableAsTimeResourceAnyOf struct {
	value *AsTimeResourceAnyOf
	isSet bool
}

func (v NullableAsTimeResourceAnyOf) Get() *AsTimeResourceAnyOf {
	return v.value
}

func (v *NullableAsTimeResourceAnyOf) Set(val *AsTimeResourceAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAsTimeResourceAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAsTimeResourceAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsTimeResourceAnyOf(val *AsTimeResourceAnyOf) *NullableAsTimeResourceAnyOf {
	return &NullableAsTimeResourceAnyOf{value: val, isSet: true}
}

func (v NullableAsTimeResourceAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsTimeResourceAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

