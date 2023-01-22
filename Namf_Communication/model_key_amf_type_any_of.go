/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// KeyAmfTypeAnyOf the model 'KeyAmfTypeAnyOf'
type KeyAmfTypeAnyOf string

// List of KeyAmfType_anyOf
const (
	KAMF KeyAmfTypeAnyOf = "KAMF"
	KPRIMEAMF KeyAmfTypeAnyOf = "KPRIMEAMF"
)

// All allowed values of KeyAmfTypeAnyOf enum
var AllowedKeyAmfTypeAnyOfEnumValues = []KeyAmfTypeAnyOf{
	"KAMF",
	"KPRIMEAMF",
}

func (v *KeyAmfTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KeyAmfTypeAnyOf(value)
	for _, existing := range AllowedKeyAmfTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KeyAmfTypeAnyOf", value)
}

// NewKeyAmfTypeAnyOfFromValue returns a pointer to a valid KeyAmfTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeyAmfTypeAnyOfFromValue(v string) (*KeyAmfTypeAnyOf, error) {
	ev := KeyAmfTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeyAmfTypeAnyOf: valid values are %v", v, AllowedKeyAmfTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeyAmfTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedKeyAmfTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KeyAmfType_anyOf value
func (v KeyAmfTypeAnyOf) Ptr() *KeyAmfTypeAnyOf {
	return &v
}

type NullableKeyAmfTypeAnyOf struct {
	value *KeyAmfTypeAnyOf
	isSet bool
}

func (v NullableKeyAmfTypeAnyOf) Get() *KeyAmfTypeAnyOf {
	return v.value
}

func (v *NullableKeyAmfTypeAnyOf) Set(val *KeyAmfTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyAmfTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyAmfTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyAmfTypeAnyOf(val *KeyAmfTypeAnyOf) *NullableKeyAmfTypeAnyOf {
	return &NullableKeyAmfTypeAnyOf{value: val, isSet: true}
}

func (v NullableKeyAmfTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyAmfTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

