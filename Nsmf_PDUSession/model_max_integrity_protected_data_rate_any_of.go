/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// MaxIntegrityProtectedDataRateAnyOf the model 'MaxIntegrityProtectedDataRateAnyOf'
type MaxIntegrityProtectedDataRateAnyOf string

// List of MaxIntegrityProtectedDataRate_anyOf
const (
	_64_KBPS MaxIntegrityProtectedDataRateAnyOf = "64_KBPS"
	MAX_UE_RATE MaxIntegrityProtectedDataRateAnyOf = "MAX_UE_RATE"
)

// All allowed values of MaxIntegrityProtectedDataRateAnyOf enum
var AllowedMaxIntegrityProtectedDataRateAnyOfEnumValues = []MaxIntegrityProtectedDataRateAnyOf{
	"64_KBPS",
	"MAX_UE_RATE",
}

func (v *MaxIntegrityProtectedDataRateAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MaxIntegrityProtectedDataRateAnyOf(value)
	for _, existing := range AllowedMaxIntegrityProtectedDataRateAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MaxIntegrityProtectedDataRateAnyOf", value)
}

// NewMaxIntegrityProtectedDataRateAnyOfFromValue returns a pointer to a valid MaxIntegrityProtectedDataRateAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMaxIntegrityProtectedDataRateAnyOfFromValue(v string) (*MaxIntegrityProtectedDataRateAnyOf, error) {
	ev := MaxIntegrityProtectedDataRateAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MaxIntegrityProtectedDataRateAnyOf: valid values are %v", v, AllowedMaxIntegrityProtectedDataRateAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MaxIntegrityProtectedDataRateAnyOf) IsValid() bool {
	for _, existing := range AllowedMaxIntegrityProtectedDataRateAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MaxIntegrityProtectedDataRate_anyOf value
func (v MaxIntegrityProtectedDataRateAnyOf) Ptr() *MaxIntegrityProtectedDataRateAnyOf {
	return &v
}

type NullableMaxIntegrityProtectedDataRateAnyOf struct {
	value *MaxIntegrityProtectedDataRateAnyOf
	isSet bool
}

func (v NullableMaxIntegrityProtectedDataRateAnyOf) Get() *MaxIntegrityProtectedDataRateAnyOf {
	return v.value
}

func (v *NullableMaxIntegrityProtectedDataRateAnyOf) Set(val *MaxIntegrityProtectedDataRateAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxIntegrityProtectedDataRateAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxIntegrityProtectedDataRateAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxIntegrityProtectedDataRateAnyOf(val *MaxIntegrityProtectedDataRateAnyOf) *NullableMaxIntegrityProtectedDataRateAnyOf {
	return &NullableMaxIntegrityProtectedDataRateAnyOf{value: val, isSet: true}
}

func (v NullableMaxIntegrityProtectedDataRateAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxIntegrityProtectedDataRateAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

