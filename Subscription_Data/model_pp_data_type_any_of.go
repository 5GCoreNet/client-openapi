/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// PpDataTypeAnyOf the model 'PpDataTypeAnyOf'
type PpDataTypeAnyOf string

// List of PpDataType_anyOf
const (
	COMMUNICATION_CHARACTERISTICS PpDataTypeAnyOf = "COMMUNICATION_CHARACTERISTICS"
	EXPECTED_UE_BEHAVIOUR PpDataTypeAnyOf = "EXPECTED_UE_BEHAVIOUR"
	EC_RESTRICTION PpDataTypeAnyOf = "EC_RESTRICTION"
	ACS_INFO PpDataTypeAnyOf = "ACS_INFO"
	TRACE PpDataTypeAnyOf = "TRACE"
	STN_SR PpDataTypeAnyOf = "STN_SR"
	LCS_PRIVACY PpDataTypeAnyOf = "LCS_PRIVACY"
	SOR_INFO PpDataTypeAnyOf = "SOR_INFO"
)

// All allowed values of PpDataTypeAnyOf enum
var AllowedPpDataTypeAnyOfEnumValues = []PpDataTypeAnyOf{
	"COMMUNICATION_CHARACTERISTICS",
	"EXPECTED_UE_BEHAVIOUR",
	"EC_RESTRICTION",
	"ACS_INFO",
	"TRACE",
	"STN_SR",
	"LCS_PRIVACY",
	"SOR_INFO",
}

func (v *PpDataTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PpDataTypeAnyOf(value)
	for _, existing := range AllowedPpDataTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PpDataTypeAnyOf", value)
}

// NewPpDataTypeAnyOfFromValue returns a pointer to a valid PpDataTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPpDataTypeAnyOfFromValue(v string) (*PpDataTypeAnyOf, error) {
	ev := PpDataTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PpDataTypeAnyOf: valid values are %v", v, AllowedPpDataTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PpDataTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedPpDataTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PpDataType_anyOf value
func (v PpDataTypeAnyOf) Ptr() *PpDataTypeAnyOf {
	return &v
}

type NullablePpDataTypeAnyOf struct {
	value *PpDataTypeAnyOf
	isSet bool
}

func (v NullablePpDataTypeAnyOf) Get() *PpDataTypeAnyOf {
	return v.value
}

func (v *NullablePpDataTypeAnyOf) Set(val *PpDataTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePpDataTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePpDataTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePpDataTypeAnyOf(val *PpDataTypeAnyOf) *NullablePpDataTypeAnyOf {
	return &NullablePpDataTypeAnyOf{value: val, isSet: true}
}

func (v NullablePpDataTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePpDataTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

