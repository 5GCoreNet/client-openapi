/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
	"fmt"
)

// MbsSessionActivityStatusAnyOf the model 'MbsSessionActivityStatusAnyOf'
type MbsSessionActivityStatusAnyOf string

// List of MbsSessionActivityStatus_anyOf
const (
	ACTIVE MbsSessionActivityStatusAnyOf = "ACTIVE"
	INACTIVE MbsSessionActivityStatusAnyOf = "INACTIVE"
)

// All allowed values of MbsSessionActivityStatusAnyOf enum
var AllowedMbsSessionActivityStatusAnyOfEnumValues = []MbsSessionActivityStatusAnyOf{
	"ACTIVE",
	"INACTIVE",
}

func (v *MbsSessionActivityStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MbsSessionActivityStatusAnyOf(value)
	for _, existing := range AllowedMbsSessionActivityStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MbsSessionActivityStatusAnyOf", value)
}

// NewMbsSessionActivityStatusAnyOfFromValue returns a pointer to a valid MbsSessionActivityStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMbsSessionActivityStatusAnyOfFromValue(v string) (*MbsSessionActivityStatusAnyOf, error) {
	ev := MbsSessionActivityStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MbsSessionActivityStatusAnyOf: valid values are %v", v, AllowedMbsSessionActivityStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MbsSessionActivityStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedMbsSessionActivityStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MbsSessionActivityStatus_anyOf value
func (v MbsSessionActivityStatusAnyOf) Ptr() *MbsSessionActivityStatusAnyOf {
	return &v
}

type NullableMbsSessionActivityStatusAnyOf struct {
	value *MbsSessionActivityStatusAnyOf
	isSet bool
}

func (v NullableMbsSessionActivityStatusAnyOf) Get() *MbsSessionActivityStatusAnyOf {
	return v.value
}

func (v *NullableMbsSessionActivityStatusAnyOf) Set(val *MbsSessionActivityStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionActivityStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionActivityStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionActivityStatusAnyOf(val *MbsSessionActivityStatusAnyOf) *NullableMbsSessionActivityStatusAnyOf {
	return &NullableMbsSessionActivityStatusAnyOf{value: val, isSet: true}
}

func (v NullableMbsSessionActivityStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionActivityStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

