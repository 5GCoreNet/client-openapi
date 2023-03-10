/*
Nchf_SpendingLimitControl

Nchf Spending Limit Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_SpendingLimitControl

import (
	"encoding/json"
	"fmt"
)

// TerminationCauseAnyOf the model 'TerminationCauseAnyOf'
type TerminationCauseAnyOf string

// List of TerminationCause_anyOf
const (
	REMOVED_SUBSCRIBER TerminationCauseAnyOf = "REMOVED_SUBSCRIBER"
)

// All allowed values of TerminationCauseAnyOf enum
var AllowedTerminationCauseAnyOfEnumValues = []TerminationCauseAnyOf{
	"REMOVED_SUBSCRIBER",
}

func (v *TerminationCauseAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TerminationCauseAnyOf(value)
	for _, existing := range AllowedTerminationCauseAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TerminationCauseAnyOf", value)
}

// NewTerminationCauseAnyOfFromValue returns a pointer to a valid TerminationCauseAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTerminationCauseAnyOfFromValue(v string) (*TerminationCauseAnyOf, error) {
	ev := TerminationCauseAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TerminationCauseAnyOf: valid values are %v", v, AllowedTerminationCauseAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TerminationCauseAnyOf) IsValid() bool {
	for _, existing := range AllowedTerminationCauseAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TerminationCause_anyOf value
func (v TerminationCauseAnyOf) Ptr() *TerminationCauseAnyOf {
	return &v
}

type NullableTerminationCauseAnyOf struct {
	value *TerminationCauseAnyOf
	isSet bool
}

func (v NullableTerminationCauseAnyOf) Get() *TerminationCauseAnyOf {
	return v.value
}

func (v *NullableTerminationCauseAnyOf) Set(val *TerminationCauseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminationCauseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminationCauseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminationCauseAnyOf(val *TerminationCauseAnyOf) *NullableTerminationCauseAnyOf {
	return &NullableTerminationCauseAnyOf{value: val, isSet: true}
}

func (v NullableTerminationCauseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminationCauseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

