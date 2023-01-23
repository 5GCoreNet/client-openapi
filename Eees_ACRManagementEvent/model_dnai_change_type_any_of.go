/*
EES ACR Management Event_API

API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRManagementEvent

import (
	"encoding/json"
	"fmt"
)

// DnaiChangeTypeAnyOf the model 'DnaiChangeTypeAnyOf'
type DnaiChangeTypeAnyOf string

// List of DnaiChangeType_anyOf
const (
	EARLY DnaiChangeTypeAnyOf = "EARLY"
	EARLY_LATE DnaiChangeTypeAnyOf = "EARLY_LATE"
	LATE DnaiChangeTypeAnyOf = "LATE"
)

// All allowed values of DnaiChangeTypeAnyOf enum
var AllowedDnaiChangeTypeAnyOfEnumValues = []DnaiChangeTypeAnyOf{
	"EARLY",
	"EARLY_LATE",
	"LATE",
}

func (v *DnaiChangeTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DnaiChangeTypeAnyOf(value)
	for _, existing := range AllowedDnaiChangeTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DnaiChangeTypeAnyOf", value)
}

// NewDnaiChangeTypeAnyOfFromValue returns a pointer to a valid DnaiChangeTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDnaiChangeTypeAnyOfFromValue(v string) (*DnaiChangeTypeAnyOf, error) {
	ev := DnaiChangeTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DnaiChangeTypeAnyOf: valid values are %v", v, AllowedDnaiChangeTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DnaiChangeTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedDnaiChangeTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DnaiChangeType_anyOf value
func (v DnaiChangeTypeAnyOf) Ptr() *DnaiChangeTypeAnyOf {
	return &v
}

type NullableDnaiChangeTypeAnyOf struct {
	value *DnaiChangeTypeAnyOf
	isSet bool
}

func (v NullableDnaiChangeTypeAnyOf) Get() *DnaiChangeTypeAnyOf {
	return v.value
}

func (v *NullableDnaiChangeTypeAnyOf) Set(val *DnaiChangeTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaiChangeTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaiChangeTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaiChangeTypeAnyOf(val *DnaiChangeTypeAnyOf) *NullableDnaiChangeTypeAnyOf {
	return &NullableDnaiChangeTypeAnyOf{value: val, isSet: true}
}

func (v NullableDnaiChangeTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaiChangeTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

