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

// DispersionTypeOneOf the model 'DispersionTypeOneOf'
type DispersionTypeOneOf string

// List of DispersionType_oneOf
const (
	DVDA DispersionTypeOneOf = "DVDA"
	TDA DispersionTypeOneOf = "TDA"
	DVDA_AND_TDA DispersionTypeOneOf = "DVDA_AND_TDA"
)

// All allowed values of DispersionTypeOneOf enum
var AllowedDispersionTypeOneOfEnumValues = []DispersionTypeOneOf{
	"DVDA",
	"TDA",
	"DVDA_AND_TDA",
}

func (v *DispersionTypeOneOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DispersionTypeOneOf(value)
	for _, existing := range AllowedDispersionTypeOneOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DispersionTypeOneOf", value)
}

// NewDispersionTypeOneOfFromValue returns a pointer to a valid DispersionTypeOneOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDispersionTypeOneOfFromValue(v string) (*DispersionTypeOneOf, error) {
	ev := DispersionTypeOneOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DispersionTypeOneOf: valid values are %v", v, AllowedDispersionTypeOneOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DispersionTypeOneOf) IsValid() bool {
	for _, existing := range AllowedDispersionTypeOneOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DispersionType_oneOf value
func (v DispersionTypeOneOf) Ptr() *DispersionTypeOneOf {
	return &v
}

type NullableDispersionTypeOneOf struct {
	value *DispersionTypeOneOf
	isSet bool
}

func (v NullableDispersionTypeOneOf) Get() *DispersionTypeOneOf {
	return v.value
}

func (v *NullableDispersionTypeOneOf) Set(val *DispersionTypeOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDispersionTypeOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDispersionTypeOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispersionTypeOneOf(val *DispersionTypeOneOf) *NullableDispersionTypeOneOf {
	return &NullableDispersionTypeOneOf{value: val, isSet: true}
}

func (v NullableDispersionTypeOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispersionTypeOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

