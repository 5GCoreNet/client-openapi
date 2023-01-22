/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// IdentityTypeAnyOf the model 'IdentityTypeAnyOf'
type IdentityTypeAnyOf string

// List of IdentityType_anyOf
const (
	DISTINCT_IMPU IdentityTypeAnyOf = "DISTINCT_IMPU"
	DISTINCT_PSI IdentityTypeAnyOf = "DISTINCT_PSI"
	WILDCARDED_IMPU IdentityTypeAnyOf = "WILDCARDED_IMPU"
	WILDCARDED_PSI IdentityTypeAnyOf = "WILDCARDED_PSI"
)

// All allowed values of IdentityTypeAnyOf enum
var AllowedIdentityTypeAnyOfEnumValues = []IdentityTypeAnyOf{
	"DISTINCT_IMPU",
	"DISTINCT_PSI",
	"WILDCARDED_IMPU",
	"WILDCARDED_PSI",
}

func (v *IdentityTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdentityTypeAnyOf(value)
	for _, existing := range AllowedIdentityTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IdentityTypeAnyOf", value)
}

// NewIdentityTypeAnyOfFromValue returns a pointer to a valid IdentityTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIdentityTypeAnyOfFromValue(v string) (*IdentityTypeAnyOf, error) {
	ev := IdentityTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IdentityTypeAnyOf: valid values are %v", v, AllowedIdentityTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdentityTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedIdentityTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IdentityType_anyOf value
func (v IdentityTypeAnyOf) Ptr() *IdentityTypeAnyOf {
	return &v
}

type NullableIdentityTypeAnyOf struct {
	value *IdentityTypeAnyOf
	isSet bool
}

func (v NullableIdentityTypeAnyOf) Get() *IdentityTypeAnyOf {
	return v.value
}

func (v *NullableIdentityTypeAnyOf) Set(val *IdentityTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityTypeAnyOf(val *IdentityTypeAnyOf) *NullableIdentityTypeAnyOf {
	return &NullableIdentityTypeAnyOf{value: val, isSet: true}
}

func (v NullableIdentityTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

