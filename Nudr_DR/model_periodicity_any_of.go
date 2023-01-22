/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// PeriodicityAnyOf the model 'PeriodicityAnyOf'
type PeriodicityAnyOf string

// List of Periodicity_anyOf
const (
	YEARLY PeriodicityAnyOf = "YEARLY"
	MONTHLY PeriodicityAnyOf = "MONTHLY"
	WEEKLY PeriodicityAnyOf = "WEEKLY"
	DAILY PeriodicityAnyOf = "DAILY"
	HOURLY PeriodicityAnyOf = "HOURLY"
)

// All allowed values of PeriodicityAnyOf enum
var AllowedPeriodicityAnyOfEnumValues = []PeriodicityAnyOf{
	"YEARLY",
	"MONTHLY",
	"WEEKLY",
	"DAILY",
	"HOURLY",
}

func (v *PeriodicityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PeriodicityAnyOf(value)
	for _, existing := range AllowedPeriodicityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PeriodicityAnyOf", value)
}

// NewPeriodicityAnyOfFromValue returns a pointer to a valid PeriodicityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPeriodicityAnyOfFromValue(v string) (*PeriodicityAnyOf, error) {
	ev := PeriodicityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PeriodicityAnyOf: valid values are %v", v, AllowedPeriodicityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PeriodicityAnyOf) IsValid() bool {
	for _, existing := range AllowedPeriodicityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Periodicity_anyOf value
func (v PeriodicityAnyOf) Ptr() *PeriodicityAnyOf {
	return &v
}

type NullablePeriodicityAnyOf struct {
	value *PeriodicityAnyOf
	isSet bool
}

func (v NullablePeriodicityAnyOf) Get() *PeriodicityAnyOf {
	return v.value
}

func (v *NullablePeriodicityAnyOf) Set(val *PeriodicityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodicityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodicityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodicityAnyOf(val *PeriodicityAnyOf) *NullablePeriodicityAnyOf {
	return &NullablePeriodicityAnyOf{value: val, isSet: true}
}

func (v NullablePeriodicityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodicityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

