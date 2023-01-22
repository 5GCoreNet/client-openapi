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

// ConditionEventTypeAnyOf the model 'ConditionEventTypeAnyOf'
type ConditionEventTypeAnyOf string

// List of ConditionEventType_anyOf
const (
	ADDED ConditionEventTypeAnyOf = "NF_ADDED"
	REMOVED ConditionEventTypeAnyOf = "NF_REMOVED"
)

// All allowed values of ConditionEventTypeAnyOf enum
var AllowedConditionEventTypeAnyOfEnumValues = []ConditionEventTypeAnyOf{
	"NF_ADDED",
	"NF_REMOVED",
}

func (v *ConditionEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConditionEventTypeAnyOf(value)
	for _, existing := range AllowedConditionEventTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConditionEventTypeAnyOf", value)
}

// NewConditionEventTypeAnyOfFromValue returns a pointer to a valid ConditionEventTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConditionEventTypeAnyOfFromValue(v string) (*ConditionEventTypeAnyOf, error) {
	ev := ConditionEventTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConditionEventTypeAnyOf: valid values are %v", v, AllowedConditionEventTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConditionEventTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedConditionEventTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConditionEventType_anyOf value
func (v ConditionEventTypeAnyOf) Ptr() *ConditionEventTypeAnyOf {
	return &v
}

type NullableConditionEventTypeAnyOf struct {
	value *ConditionEventTypeAnyOf
	isSet bool
}

func (v NullableConditionEventTypeAnyOf) Get() *ConditionEventTypeAnyOf {
	return v.value
}

func (v *NullableConditionEventTypeAnyOf) Set(val *ConditionEventTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionEventTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionEventTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionEventTypeAnyOf(val *ConditionEventTypeAnyOf) *NullableConditionEventTypeAnyOf {
	return &NullableConditionEventTypeAnyOf{value: val, isSet: true}
}

func (v NullableConditionEventTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

