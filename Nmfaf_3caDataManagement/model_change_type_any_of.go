/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// ChangeTypeAnyOf the model 'ChangeTypeAnyOf'
type ChangeTypeAnyOf string

// List of ChangeType_anyOf
const (
	ADD ChangeTypeAnyOf = "ADD"
	MOVE ChangeTypeAnyOf = "MOVE"
	REMOVE ChangeTypeAnyOf = "REMOVE"
	REPLACE ChangeTypeAnyOf = "REPLACE"
)

// All allowed values of ChangeTypeAnyOf enum
var AllowedChangeTypeAnyOfEnumValues = []ChangeTypeAnyOf{
	"ADD",
	"MOVE",
	"REMOVE",
	"REPLACE",
}

func (v *ChangeTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChangeTypeAnyOf(value)
	for _, existing := range AllowedChangeTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChangeTypeAnyOf", value)
}

// NewChangeTypeAnyOfFromValue returns a pointer to a valid ChangeTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChangeTypeAnyOfFromValue(v string) (*ChangeTypeAnyOf, error) {
	ev := ChangeTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChangeTypeAnyOf: valid values are %v", v, AllowedChangeTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChangeTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedChangeTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeType_anyOf value
func (v ChangeTypeAnyOf) Ptr() *ChangeTypeAnyOf {
	return &v
}

type NullableChangeTypeAnyOf struct {
	value *ChangeTypeAnyOf
	isSet bool
}

func (v NullableChangeTypeAnyOf) Get() *ChangeTypeAnyOf {
	return v.value
}

func (v *NullableChangeTypeAnyOf) Set(val *ChangeTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeTypeAnyOf(val *ChangeTypeAnyOf) *NullableChangeTypeAnyOf {
	return &NullableChangeTypeAnyOf{value: val, isSet: true}
}

func (v NullableChangeTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

