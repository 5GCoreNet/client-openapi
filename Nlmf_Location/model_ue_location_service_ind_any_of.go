/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nlmf_Location

import (
	"encoding/json"
	"fmt"
)

// UeLocationServiceIndAnyOf the model 'UeLocationServiceIndAnyOf'
type UeLocationServiceIndAnyOf string

// List of UeLocationServiceInd_anyOf
const (
	ESTIMATE UeLocationServiceIndAnyOf = "LOCATION_ESTIMATE"
	ASSISTANCE_DATA UeLocationServiceIndAnyOf = "LOCATION_ASSISTANCE_DATA"
)

// All allowed values of UeLocationServiceIndAnyOf enum
var AllowedUeLocationServiceIndAnyOfEnumValues = []UeLocationServiceIndAnyOf{
	"LOCATION_ESTIMATE",
	"LOCATION_ASSISTANCE_DATA",
}

func (v *UeLocationServiceIndAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UeLocationServiceIndAnyOf(value)
	for _, existing := range AllowedUeLocationServiceIndAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UeLocationServiceIndAnyOf", value)
}

// NewUeLocationServiceIndAnyOfFromValue returns a pointer to a valid UeLocationServiceIndAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUeLocationServiceIndAnyOfFromValue(v string) (*UeLocationServiceIndAnyOf, error) {
	ev := UeLocationServiceIndAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UeLocationServiceIndAnyOf: valid values are %v", v, AllowedUeLocationServiceIndAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UeLocationServiceIndAnyOf) IsValid() bool {
	for _, existing := range AllowedUeLocationServiceIndAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UeLocationServiceInd_anyOf value
func (v UeLocationServiceIndAnyOf) Ptr() *UeLocationServiceIndAnyOf {
	return &v
}

type NullableUeLocationServiceIndAnyOf struct {
	value *UeLocationServiceIndAnyOf
	isSet bool
}

func (v NullableUeLocationServiceIndAnyOf) Get() *UeLocationServiceIndAnyOf {
	return v.value
}

func (v *NullableUeLocationServiceIndAnyOf) Set(val *UeLocationServiceIndAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUeLocationServiceIndAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUeLocationServiceIndAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeLocationServiceIndAnyOf(val *UeLocationServiceIndAnyOf) *NullableUeLocationServiceIndAnyOf {
	return &NullableUeLocationServiceIndAnyOf{value: val, isSet: true}
}

func (v NullableUeLocationServiceIndAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeLocationServiceIndAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

