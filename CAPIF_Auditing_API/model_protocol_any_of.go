/*
CAPIF_Auditing_API

API for auditing.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Auditing_API

import (
	"encoding/json"
	"fmt"
)

// ProtocolAnyOf the model 'ProtocolAnyOf'
type ProtocolAnyOf string

// List of Protocol_anyOf
const (
	_1_1 ProtocolAnyOf = "HTTP_1_1"
	_2 ProtocolAnyOf = "HTTP_2"
)

// All allowed values of ProtocolAnyOf enum
var AllowedProtocolAnyOfEnumValues = []ProtocolAnyOf{
	"HTTP_1_1",
	"HTTP_2",
}

func (v *ProtocolAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProtocolAnyOf(value)
	for _, existing := range AllowedProtocolAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProtocolAnyOf", value)
}

// NewProtocolAnyOfFromValue returns a pointer to a valid ProtocolAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProtocolAnyOfFromValue(v string) (*ProtocolAnyOf, error) {
	ev := ProtocolAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProtocolAnyOf: valid values are %v", v, AllowedProtocolAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProtocolAnyOf) IsValid() bool {
	for _, existing := range AllowedProtocolAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Protocol_anyOf value
func (v ProtocolAnyOf) Ptr() *ProtocolAnyOf {
	return &v
}

type NullableProtocolAnyOf struct {
	value *ProtocolAnyOf
	isSet bool
}

func (v NullableProtocolAnyOf) Get() *ProtocolAnyOf {
	return v.value
}

func (v *NullableProtocolAnyOf) Set(val *ProtocolAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolAnyOf(val *ProtocolAnyOf) *NullableProtocolAnyOf {
	return &NullableProtocolAnyOf{value: val, isSet: true}
}

func (v NullableProtocolAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

