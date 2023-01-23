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

// DispersionClassOneOf the model 'DispersionClassOneOf'
type DispersionClassOneOf string

// List of DispersionClass_oneOf
const (
	FIXED DispersionClassOneOf = "FIXED"
	CAMPER DispersionClassOneOf = "CAMPER"
	TRAVELLER DispersionClassOneOf = "TRAVELLER"
	TOP_HEAVY DispersionClassOneOf = "TOP_HEAVY"
)

// All allowed values of DispersionClassOneOf enum
var AllowedDispersionClassOneOfEnumValues = []DispersionClassOneOf{
	"FIXED",
	"CAMPER",
	"TRAVELLER",
	"TOP_HEAVY",
}

func (v *DispersionClassOneOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DispersionClassOneOf(value)
	for _, existing := range AllowedDispersionClassOneOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DispersionClassOneOf", value)
}

// NewDispersionClassOneOfFromValue returns a pointer to a valid DispersionClassOneOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDispersionClassOneOfFromValue(v string) (*DispersionClassOneOf, error) {
	ev := DispersionClassOneOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DispersionClassOneOf: valid values are %v", v, AllowedDispersionClassOneOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DispersionClassOneOf) IsValid() bool {
	for _, existing := range AllowedDispersionClassOneOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DispersionClass_oneOf value
func (v DispersionClassOneOf) Ptr() *DispersionClassOneOf {
	return &v
}

type NullableDispersionClassOneOf struct {
	value *DispersionClassOneOf
	isSet bool
}

func (v NullableDispersionClassOneOf) Get() *DispersionClassOneOf {
	return v.value
}

func (v *NullableDispersionClassOneOf) Set(val *DispersionClassOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDispersionClassOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDispersionClassOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispersionClassOneOf(val *DispersionClassOneOf) *NullableDispersionClassOneOf {
	return &NullableDispersionClassOneOf{value: val, isSet: true}
}

func (v NullableDispersionClassOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispersionClassOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

