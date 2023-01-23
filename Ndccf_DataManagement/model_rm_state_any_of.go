/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// RmStateAnyOf the model 'RmStateAnyOf'
type RmStateAnyOf string

// List of RmState_anyOf
const (
	REGISTERED RmStateAnyOf = "REGISTERED"
	DEREGISTERED RmStateAnyOf = "DEREGISTERED"
)

// All allowed values of RmStateAnyOf enum
var AllowedRmStateAnyOfEnumValues = []RmStateAnyOf{
	"REGISTERED",
	"DEREGISTERED",
}

func (v *RmStateAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RmStateAnyOf(value)
	for _, existing := range AllowedRmStateAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RmStateAnyOf", value)
}

// NewRmStateAnyOfFromValue returns a pointer to a valid RmStateAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRmStateAnyOfFromValue(v string) (*RmStateAnyOf, error) {
	ev := RmStateAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RmStateAnyOf: valid values are %v", v, AllowedRmStateAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RmStateAnyOf) IsValid() bool {
	for _, existing := range AllowedRmStateAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RmState_anyOf value
func (v RmStateAnyOf) Ptr() *RmStateAnyOf {
	return &v
}

type NullableRmStateAnyOf struct {
	value *RmStateAnyOf
	isSet bool
}

func (v NullableRmStateAnyOf) Get() *RmStateAnyOf {
	return v.value
}

func (v *NullableRmStateAnyOf) Set(val *RmStateAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRmStateAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRmStateAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmStateAnyOf(val *RmStateAnyOf) *NullableRmStateAnyOf {
	return &NullableRmStateAnyOf{value: val, isSet: true}
}

func (v NullableRmStateAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmStateAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

