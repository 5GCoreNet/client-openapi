/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// UeAuthAnyOf the model 'UeAuthAnyOf'
type UeAuthAnyOf string

// List of UeAuth_anyOf
const (
	AUTHORIZED UeAuthAnyOf = "AUTHORIZED"
	NOT_AUTHORIZED UeAuthAnyOf = "NOT_AUTHORIZED"
)

// All allowed values of UeAuthAnyOf enum
var AllowedUeAuthAnyOfEnumValues = []UeAuthAnyOf{
	"AUTHORIZED",
	"NOT_AUTHORIZED",
}

func (v *UeAuthAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UeAuthAnyOf(value)
	for _, existing := range AllowedUeAuthAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UeAuthAnyOf", value)
}

// NewUeAuthAnyOfFromValue returns a pointer to a valid UeAuthAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUeAuthAnyOfFromValue(v string) (*UeAuthAnyOf, error) {
	ev := UeAuthAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UeAuthAnyOf: valid values are %v", v, AllowedUeAuthAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UeAuthAnyOf) IsValid() bool {
	for _, existing := range AllowedUeAuthAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UeAuth_anyOf value
func (v UeAuthAnyOf) Ptr() *UeAuthAnyOf {
	return &v
}

type NullableUeAuthAnyOf struct {
	value *UeAuthAnyOf
	isSet bool
}

func (v NullableUeAuthAnyOf) Get() *UeAuthAnyOf {
	return v.value
}

func (v *NullableUeAuthAnyOf) Set(val *UeAuthAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUeAuthAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUeAuthAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeAuthAnyOf(val *UeAuthAnyOf) *NullableUeAuthAnyOf {
	return &NullableUeAuthAnyOf{value: val, isSet: true}
}

func (v NullableUeAuthAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeAuthAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

