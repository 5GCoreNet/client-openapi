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

// GmCapableAnyOf the model 'GmCapableAnyOf'
type GmCapableAnyOf string

// List of GmCapable_anyOf
const (
	GPTP GmCapableAnyOf = "GPTP"
	PTP GmCapableAnyOf = "PTP"
)

// All allowed values of GmCapableAnyOf enum
var AllowedGmCapableAnyOfEnumValues = []GmCapableAnyOf{
	"GPTP",
	"PTP",
}

func (v *GmCapableAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GmCapableAnyOf(value)
	for _, existing := range AllowedGmCapableAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GmCapableAnyOf", value)
}

// NewGmCapableAnyOfFromValue returns a pointer to a valid GmCapableAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGmCapableAnyOfFromValue(v string) (*GmCapableAnyOf, error) {
	ev := GmCapableAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GmCapableAnyOf: valid values are %v", v, AllowedGmCapableAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GmCapableAnyOf) IsValid() bool {
	for _, existing := range AllowedGmCapableAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GmCapable_anyOf value
func (v GmCapableAnyOf) Ptr() *GmCapableAnyOf {
	return &v
}

type NullableGmCapableAnyOf struct {
	value *GmCapableAnyOf
	isSet bool
}

func (v NullableGmCapableAnyOf) Get() *GmCapableAnyOf {
	return v.value
}

func (v *NullableGmCapableAnyOf) Set(val *GmCapableAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGmCapableAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGmCapableAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGmCapableAnyOf(val *GmCapableAnyOf) *NullableGmCapableAnyOf {
	return &NullableGmCapableAnyOf{value: val, isSet: true}
}

func (v NullableGmCapableAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGmCapableAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

