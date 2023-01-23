/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// MatchingDirectionAnyOf the model 'MatchingDirectionAnyOf'
type MatchingDirectionAnyOf string

// List of MatchingDirection_anyOf
const (
	ASCENDING MatchingDirectionAnyOf = "ASCENDING"
	DESCENDING MatchingDirectionAnyOf = "DESCENDING"
	CROSSED MatchingDirectionAnyOf = "CROSSED"
)

// All allowed values of MatchingDirectionAnyOf enum
var AllowedMatchingDirectionAnyOfEnumValues = []MatchingDirectionAnyOf{
	"ASCENDING",
	"DESCENDING",
	"CROSSED",
}

func (v *MatchingDirectionAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MatchingDirectionAnyOf(value)
	for _, existing := range AllowedMatchingDirectionAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MatchingDirectionAnyOf", value)
}

// NewMatchingDirectionAnyOfFromValue returns a pointer to a valid MatchingDirectionAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMatchingDirectionAnyOfFromValue(v string) (*MatchingDirectionAnyOf, error) {
	ev := MatchingDirectionAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MatchingDirectionAnyOf: valid values are %v", v, AllowedMatchingDirectionAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MatchingDirectionAnyOf) IsValid() bool {
	for _, existing := range AllowedMatchingDirectionAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MatchingDirection_anyOf value
func (v MatchingDirectionAnyOf) Ptr() *MatchingDirectionAnyOf {
	return &v
}

type NullableMatchingDirectionAnyOf struct {
	value *MatchingDirectionAnyOf
	isSet bool
}

func (v NullableMatchingDirectionAnyOf) Get() *MatchingDirectionAnyOf {
	return v.value
}

func (v *NullableMatchingDirectionAnyOf) Set(val *MatchingDirectionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchingDirectionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchingDirectionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchingDirectionAnyOf(val *MatchingDirectionAnyOf) *NullableMatchingDirectionAnyOf {
	return &NullableMatchingDirectionAnyOf{value: val, isSet: true}
}

func (v NullableMatchingDirectionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchingDirectionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

