/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// AfRequestedDataAnyOf the model 'AfRequestedDataAnyOf'
type AfRequestedDataAnyOf string

// List of AfRequestedData_anyOf
const (
	UE_IDENTITY AfRequestedDataAnyOf = "UE_IDENTITY"
)

// All allowed values of AfRequestedDataAnyOf enum
var AllowedAfRequestedDataAnyOfEnumValues = []AfRequestedDataAnyOf{
	"UE_IDENTITY",
}

func (v *AfRequestedDataAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AfRequestedDataAnyOf(value)
	for _, existing := range AllowedAfRequestedDataAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AfRequestedDataAnyOf", value)
}

// NewAfRequestedDataAnyOfFromValue returns a pointer to a valid AfRequestedDataAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAfRequestedDataAnyOfFromValue(v string) (*AfRequestedDataAnyOf, error) {
	ev := AfRequestedDataAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AfRequestedDataAnyOf: valid values are %v", v, AllowedAfRequestedDataAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AfRequestedDataAnyOf) IsValid() bool {
	for _, existing := range AllowedAfRequestedDataAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AfRequestedData_anyOf value
func (v AfRequestedDataAnyOf) Ptr() *AfRequestedDataAnyOf {
	return &v
}

type NullableAfRequestedDataAnyOf struct {
	value *AfRequestedDataAnyOf
	isSet bool
}

func (v NullableAfRequestedDataAnyOf) Get() *AfRequestedDataAnyOf {
	return v.value
}

func (v *NullableAfRequestedDataAnyOf) Set(val *AfRequestedDataAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAfRequestedDataAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAfRequestedDataAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfRequestedDataAnyOf(val *AfRequestedDataAnyOf) *NullableAfRequestedDataAnyOf {
	return &NullableAfRequestedDataAnyOf{value: val, isSet: true}
}

func (v NullableAfRequestedDataAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfRequestedDataAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

