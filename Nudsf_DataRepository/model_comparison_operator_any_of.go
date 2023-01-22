/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudsf_DataRepository

import (
	"encoding/json"
	"fmt"
)

// ComparisonOperatorAnyOf the model 'ComparisonOperatorAnyOf'
type ComparisonOperatorAnyOf string

// List of ComparisonOperator_anyOf
const (
	EQ ComparisonOperatorAnyOf = "EQ"
	NEQ ComparisonOperatorAnyOf = "NEQ"
	GT ComparisonOperatorAnyOf = "GT"
	GTE ComparisonOperatorAnyOf = "GTE"
	LT ComparisonOperatorAnyOf = "LT"
	LTE ComparisonOperatorAnyOf = "LTE"
)

// All allowed values of ComparisonOperatorAnyOf enum
var AllowedComparisonOperatorAnyOfEnumValues = []ComparisonOperatorAnyOf{
	"EQ",
	"NEQ",
	"GT",
	"GTE",
	"LT",
	"LTE",
}

func (v *ComparisonOperatorAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ComparisonOperatorAnyOf(value)
	for _, existing := range AllowedComparisonOperatorAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ComparisonOperatorAnyOf", value)
}

// NewComparisonOperatorAnyOfFromValue returns a pointer to a valid ComparisonOperatorAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComparisonOperatorAnyOfFromValue(v string) (*ComparisonOperatorAnyOf, error) {
	ev := ComparisonOperatorAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ComparisonOperatorAnyOf: valid values are %v", v, AllowedComparisonOperatorAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComparisonOperatorAnyOf) IsValid() bool {
	for _, existing := range AllowedComparisonOperatorAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ComparisonOperator_anyOf value
func (v ComparisonOperatorAnyOf) Ptr() *ComparisonOperatorAnyOf {
	return &v
}

type NullableComparisonOperatorAnyOf struct {
	value *ComparisonOperatorAnyOf
	isSet bool
}

func (v NullableComparisonOperatorAnyOf) Get() *ComparisonOperatorAnyOf {
	return v.value
}

func (v *NullableComparisonOperatorAnyOf) Set(val *ComparisonOperatorAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComparisonOperatorAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComparisonOperatorAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComparisonOperatorAnyOf(val *ComparisonOperatorAnyOf) *NullableComparisonOperatorAnyOf {
	return &NullableComparisonOperatorAnyOf{value: val, isSet: true}
}

func (v NullableComparisonOperatorAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComparisonOperatorAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

