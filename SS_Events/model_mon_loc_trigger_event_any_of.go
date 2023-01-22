/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SS_Events

import (
	"encoding/json"
	"fmt"
)

// MonLocTriggerEventAnyOf the model 'MonLocTriggerEventAnyOf'
type MonLocTriggerEventAnyOf string

// List of MonLocTriggerEvent_anyOf
const (
	DISTANCE_TRAVELLED MonLocTriggerEventAnyOf = "DISTANCE_TRAVELLED"
)

// All allowed values of MonLocTriggerEventAnyOf enum
var AllowedMonLocTriggerEventAnyOfEnumValues = []MonLocTriggerEventAnyOf{
	"DISTANCE_TRAVELLED",
}

func (v *MonLocTriggerEventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MonLocTriggerEventAnyOf(value)
	for _, existing := range AllowedMonLocTriggerEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MonLocTriggerEventAnyOf", value)
}

// NewMonLocTriggerEventAnyOfFromValue returns a pointer to a valid MonLocTriggerEventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonLocTriggerEventAnyOfFromValue(v string) (*MonLocTriggerEventAnyOf, error) {
	ev := MonLocTriggerEventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MonLocTriggerEventAnyOf: valid values are %v", v, AllowedMonLocTriggerEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonLocTriggerEventAnyOf) IsValid() bool {
	for _, existing := range AllowedMonLocTriggerEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonLocTriggerEvent_anyOf value
func (v MonLocTriggerEventAnyOf) Ptr() *MonLocTriggerEventAnyOf {
	return &v
}

type NullableMonLocTriggerEventAnyOf struct {
	value *MonLocTriggerEventAnyOf
	isSet bool
}

func (v NullableMonLocTriggerEventAnyOf) Get() *MonLocTriggerEventAnyOf {
	return v.value
}

func (v *NullableMonLocTriggerEventAnyOf) Set(val *MonLocTriggerEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMonLocTriggerEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMonLocTriggerEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonLocTriggerEventAnyOf(val *MonLocTriggerEventAnyOf) *NullableMonLocTriggerEventAnyOf {
	return &NullableMonLocTriggerEventAnyOf{value: val, isSet: true}
}

func (v NullableMonLocTriggerEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonLocTriggerEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

