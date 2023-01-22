/*
Nnwdaf_MLModelProvision

Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_MLModelProvision

import (
	"encoding/json"
	"fmt"
)

// FailureCodeAnyOf the model 'FailureCodeAnyOf'
type FailureCodeAnyOf string

// List of FailureCode_anyOf
const (
	UNAVAILABLE_ML_MODEL FailureCodeAnyOf = "UNAVAILABLE_ML_MODEL"
)

// All allowed values of FailureCodeAnyOf enum
var AllowedFailureCodeAnyOfEnumValues = []FailureCodeAnyOf{
	"UNAVAILABLE_ML_MODEL",
}

func (v *FailureCodeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FailureCodeAnyOf(value)
	for _, existing := range AllowedFailureCodeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FailureCodeAnyOf", value)
}

// NewFailureCodeAnyOfFromValue returns a pointer to a valid FailureCodeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFailureCodeAnyOfFromValue(v string) (*FailureCodeAnyOf, error) {
	ev := FailureCodeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FailureCodeAnyOf: valid values are %v", v, AllowedFailureCodeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FailureCodeAnyOf) IsValid() bool {
	for _, existing := range AllowedFailureCodeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FailureCode_anyOf value
func (v FailureCodeAnyOf) Ptr() *FailureCodeAnyOf {
	return &v
}

type NullableFailureCodeAnyOf struct {
	value *FailureCodeAnyOf
	isSet bool
}

func (v NullableFailureCodeAnyOf) Get() *FailureCodeAnyOf {
	return v.value
}

func (v *NullableFailureCodeAnyOf) Set(val *FailureCodeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureCodeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureCodeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureCodeAnyOf(val *FailureCodeAnyOf) *NullableFailureCodeAnyOf {
	return &NullableFailureCodeAnyOf{value: val, isSet: true}
}

func (v NullableFailureCodeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureCodeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

