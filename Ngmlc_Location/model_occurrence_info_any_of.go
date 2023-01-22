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

// OccurrenceInfoAnyOf the model 'OccurrenceInfoAnyOf'
type OccurrenceInfoAnyOf string

// List of OccurrenceInfo_anyOf
const (
	ONE_TIME_EVENT OccurrenceInfoAnyOf = "ONE_TIME_EVENT"
	MULTIPLE_TIME_EVENT OccurrenceInfoAnyOf = "MULTIPLE_TIME_EVENT"
)

// All allowed values of OccurrenceInfoAnyOf enum
var AllowedOccurrenceInfoAnyOfEnumValues = []OccurrenceInfoAnyOf{
	"ONE_TIME_EVENT",
	"MULTIPLE_TIME_EVENT",
}

func (v *OccurrenceInfoAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OccurrenceInfoAnyOf(value)
	for _, existing := range AllowedOccurrenceInfoAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OccurrenceInfoAnyOf", value)
}

// NewOccurrenceInfoAnyOfFromValue returns a pointer to a valid OccurrenceInfoAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOccurrenceInfoAnyOfFromValue(v string) (*OccurrenceInfoAnyOf, error) {
	ev := OccurrenceInfoAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OccurrenceInfoAnyOf: valid values are %v", v, AllowedOccurrenceInfoAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OccurrenceInfoAnyOf) IsValid() bool {
	for _, existing := range AllowedOccurrenceInfoAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OccurrenceInfo_anyOf value
func (v OccurrenceInfoAnyOf) Ptr() *OccurrenceInfoAnyOf {
	return &v
}

type NullableOccurrenceInfoAnyOf struct {
	value *OccurrenceInfoAnyOf
	isSet bool
}

func (v NullableOccurrenceInfoAnyOf) Get() *OccurrenceInfoAnyOf {
	return v.value
}

func (v *NullableOccurrenceInfoAnyOf) Set(val *OccurrenceInfoAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOccurrenceInfoAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOccurrenceInfoAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOccurrenceInfoAnyOf(val *OccurrenceInfoAnyOf) *NullableOccurrenceInfoAnyOf {
	return &NullableOccurrenceInfoAnyOf{value: val, isSet: true}
}

func (v NullableOccurrenceInfoAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOccurrenceInfoAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
