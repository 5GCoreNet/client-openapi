/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// PseudonymIndicatorAnyOf the model 'PseudonymIndicatorAnyOf'
type PseudonymIndicatorAnyOf string

// List of PseudonymIndicator_anyOf
const (
	REQUESTED PseudonymIndicatorAnyOf = "PSEUDONYM_REQUESTED"
	NOT_REQUESTED PseudonymIndicatorAnyOf = "PSEUDONYM_NOT_REQUESTED"
)

// All allowed values of PseudonymIndicatorAnyOf enum
var AllowedPseudonymIndicatorAnyOfEnumValues = []PseudonymIndicatorAnyOf{
	"PSEUDONYM_REQUESTED",
	"PSEUDONYM_NOT_REQUESTED",
}

func (v *PseudonymIndicatorAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PseudonymIndicatorAnyOf(value)
	for _, existing := range AllowedPseudonymIndicatorAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PseudonymIndicatorAnyOf", value)
}

// NewPseudonymIndicatorAnyOfFromValue returns a pointer to a valid PseudonymIndicatorAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPseudonymIndicatorAnyOfFromValue(v string) (*PseudonymIndicatorAnyOf, error) {
	ev := PseudonymIndicatorAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PseudonymIndicatorAnyOf: valid values are %v", v, AllowedPseudonymIndicatorAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PseudonymIndicatorAnyOf) IsValid() bool {
	for _, existing := range AllowedPseudonymIndicatorAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PseudonymIndicator_anyOf value
func (v PseudonymIndicatorAnyOf) Ptr() *PseudonymIndicatorAnyOf {
	return &v
}

type NullablePseudonymIndicatorAnyOf struct {
	value *PseudonymIndicatorAnyOf
	isSet bool
}

func (v NullablePseudonymIndicatorAnyOf) Get() *PseudonymIndicatorAnyOf {
	return v.value
}

func (v *NullablePseudonymIndicatorAnyOf) Set(val *PseudonymIndicatorAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePseudonymIndicatorAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePseudonymIndicatorAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePseudonymIndicatorAnyOf(val *PseudonymIndicatorAnyOf) *NullablePseudonymIndicatorAnyOf {
	return &NullablePseudonymIndicatorAnyOf{value: val, isSet: true}
}

func (v NullablePseudonymIndicatorAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePseudonymIndicatorAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
