/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// RedTransExpOrderingCriterionAnyOf the model 'RedTransExpOrderingCriterionAnyOf'
type RedTransExpOrderingCriterionAnyOf string

// List of RedTransExpOrderingCriterion_anyOf
const (
	TIME_SLOT_START RedTransExpOrderingCriterionAnyOf = "TIME_SLOT_START"
	RED_TRANS_EXP RedTransExpOrderingCriterionAnyOf = "RED_TRANS_EXP"
)

// All allowed values of RedTransExpOrderingCriterionAnyOf enum
var AllowedRedTransExpOrderingCriterionAnyOfEnumValues = []RedTransExpOrderingCriterionAnyOf{
	"TIME_SLOT_START",
	"RED_TRANS_EXP",
}

func (v *RedTransExpOrderingCriterionAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RedTransExpOrderingCriterionAnyOf(value)
	for _, existing := range AllowedRedTransExpOrderingCriterionAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RedTransExpOrderingCriterionAnyOf", value)
}

// NewRedTransExpOrderingCriterionAnyOfFromValue returns a pointer to a valid RedTransExpOrderingCriterionAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRedTransExpOrderingCriterionAnyOfFromValue(v string) (*RedTransExpOrderingCriterionAnyOf, error) {
	ev := RedTransExpOrderingCriterionAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RedTransExpOrderingCriterionAnyOf: valid values are %v", v, AllowedRedTransExpOrderingCriterionAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RedTransExpOrderingCriterionAnyOf) IsValid() bool {
	for _, existing := range AllowedRedTransExpOrderingCriterionAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RedTransExpOrderingCriterion_anyOf value
func (v RedTransExpOrderingCriterionAnyOf) Ptr() *RedTransExpOrderingCriterionAnyOf {
	return &v
}

type NullableRedTransExpOrderingCriterionAnyOf struct {
	value *RedTransExpOrderingCriterionAnyOf
	isSet bool
}

func (v NullableRedTransExpOrderingCriterionAnyOf) Get() *RedTransExpOrderingCriterionAnyOf {
	return v.value
}

func (v *NullableRedTransExpOrderingCriterionAnyOf) Set(val *RedTransExpOrderingCriterionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRedTransExpOrderingCriterionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRedTransExpOrderingCriterionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedTransExpOrderingCriterionAnyOf(val *RedTransExpOrderingCriterionAnyOf) *NullableRedTransExpOrderingCriterionAnyOf {
	return &NullableRedTransExpOrderingCriterionAnyOf{value: val, isSet: true}
}

func (v NullableRedTransExpOrderingCriterionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedTransExpOrderingCriterionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

