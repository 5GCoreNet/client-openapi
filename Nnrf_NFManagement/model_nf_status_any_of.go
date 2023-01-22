/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFManagement

import (
	"encoding/json"
	"fmt"
)

// NFStatusAnyOf the model 'NFStatusAnyOf'
type NFStatusAnyOf string

// List of NFStatus_anyOf
const (
	REGISTERED NFStatusAnyOf = "REGISTERED"
	SUSPENDED NFStatusAnyOf = "SUSPENDED"
	UNDISCOVERABLE NFStatusAnyOf = "UNDISCOVERABLE"
)

// All allowed values of NFStatusAnyOf enum
var AllowedNFStatusAnyOfEnumValues = []NFStatusAnyOf{
	"REGISTERED",
	"SUSPENDED",
	"UNDISCOVERABLE",
}

func (v *NFStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NFStatusAnyOf(value)
	for _, existing := range AllowedNFStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NFStatusAnyOf", value)
}

// NewNFStatusAnyOfFromValue returns a pointer to a valid NFStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNFStatusAnyOfFromValue(v string) (*NFStatusAnyOf, error) {
	ev := NFStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NFStatusAnyOf: valid values are %v", v, AllowedNFStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NFStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedNFStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NFStatus_anyOf value
func (v NFStatusAnyOf) Ptr() *NFStatusAnyOf {
	return &v
}

type NullableNFStatusAnyOf struct {
	value *NFStatusAnyOf
	isSet bool
}

func (v NullableNFStatusAnyOf) Get() *NFStatusAnyOf {
	return v.value
}

func (v *NullableNFStatusAnyOf) Set(val *NFStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNFStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNFStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFStatusAnyOf(val *NFStatusAnyOf) *NullableNFStatusAnyOf {
	return &NullableNFStatusAnyOf{value: val, isSet: true}
}

func (v NullableNFStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

