/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// SqnSchemeAnyOf the model 'SqnSchemeAnyOf'
type SqnSchemeAnyOf string

// List of SqnScheme_anyOf
const (
	GENERAL SqnSchemeAnyOf = "GENERAL"
	NON_TIME_BASED SqnSchemeAnyOf = "NON_TIME_BASED"
	TIME_BASED SqnSchemeAnyOf = "TIME_BASED"
)

// All allowed values of SqnSchemeAnyOf enum
var AllowedSqnSchemeAnyOfEnumValues = []SqnSchemeAnyOf{
	"GENERAL",
	"NON_TIME_BASED",
	"TIME_BASED",
}

func (v *SqnSchemeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SqnSchemeAnyOf(value)
	for _, existing := range AllowedSqnSchemeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SqnSchemeAnyOf", value)
}

// NewSqnSchemeAnyOfFromValue returns a pointer to a valid SqnSchemeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSqnSchemeAnyOfFromValue(v string) (*SqnSchemeAnyOf, error) {
	ev := SqnSchemeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SqnSchemeAnyOf: valid values are %v", v, AllowedSqnSchemeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SqnSchemeAnyOf) IsValid() bool {
	for _, existing := range AllowedSqnSchemeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SqnScheme_anyOf value
func (v SqnSchemeAnyOf) Ptr() *SqnSchemeAnyOf {
	return &v
}

type NullableSqnSchemeAnyOf struct {
	value *SqnSchemeAnyOf
	isSet bool
}

func (v NullableSqnSchemeAnyOf) Get() *SqnSchemeAnyOf {
	return v.value
}

func (v *NullableSqnSchemeAnyOf) Set(val *SqnSchemeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSqnSchemeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSqnSchemeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqnSchemeAnyOf(val *SqnSchemeAnyOf) *NullableSqnSchemeAnyOf {
	return &NullableSqnSchemeAnyOf{value: val, isSet: true}
}

func (v NullableSqnSchemeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSqnSchemeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

