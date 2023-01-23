/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Location

import (
	"encoding/json"
	"fmt"
)

// PositioningModeAnyOf the model 'PositioningModeAnyOf'
type PositioningModeAnyOf string

// List of PositioningMode_anyOf
const (
	UE_BASED PositioningModeAnyOf = "UE_BASED"
	UE_ASSISTED PositioningModeAnyOf = "UE_ASSISTED"
	CONVENTIONAL PositioningModeAnyOf = "CONVENTIONAL"
)

// All allowed values of PositioningModeAnyOf enum
var AllowedPositioningModeAnyOfEnumValues = []PositioningModeAnyOf{
	"UE_BASED",
	"UE_ASSISTED",
	"CONVENTIONAL",
}

func (v *PositioningModeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PositioningModeAnyOf(value)
	for _, existing := range AllowedPositioningModeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PositioningModeAnyOf", value)
}

// NewPositioningModeAnyOfFromValue returns a pointer to a valid PositioningModeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPositioningModeAnyOfFromValue(v string) (*PositioningModeAnyOf, error) {
	ev := PositioningModeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PositioningModeAnyOf: valid values are %v", v, AllowedPositioningModeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PositioningModeAnyOf) IsValid() bool {
	for _, existing := range AllowedPositioningModeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PositioningMode_anyOf value
func (v PositioningModeAnyOf) Ptr() *PositioningModeAnyOf {
	return &v
}

type NullablePositioningModeAnyOf struct {
	value *PositioningModeAnyOf
	isSet bool
}

func (v NullablePositioningModeAnyOf) Get() *PositioningModeAnyOf {
	return v.value
}

func (v *NullablePositioningModeAnyOf) Set(val *PositioningModeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePositioningModeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePositioningModeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositioningModeAnyOf(val *PositioningModeAnyOf) *NullablePositioningModeAnyOf {
	return &NullablePositioningModeAnyOf{value: val, isSet: true}
}

func (v NullablePositioningModeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositioningModeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

