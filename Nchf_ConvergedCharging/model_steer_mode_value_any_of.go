/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// SteerModeValueAnyOf the model 'SteerModeValueAnyOf'
type SteerModeValueAnyOf string

// List of SteerModeValue_anyOf
const (
	ACTIVE_STANDBY SteerModeValueAnyOf = "ACTIVE_STANDBY"
	LOAD_BALANCING SteerModeValueAnyOf = "LOAD_BALANCING"
	SMALLEST_DELAY SteerModeValueAnyOf = "SMALLEST_DELAY"
	PRIORITY_BASED SteerModeValueAnyOf = "PRIORITY_BASED"
)

// All allowed values of SteerModeValueAnyOf enum
var AllowedSteerModeValueAnyOfEnumValues = []SteerModeValueAnyOf{
	"ACTIVE_STANDBY",
	"LOAD_BALANCING",
	"SMALLEST_DELAY",
	"PRIORITY_BASED",
}

func (v *SteerModeValueAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SteerModeValueAnyOf(value)
	for _, existing := range AllowedSteerModeValueAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SteerModeValueAnyOf", value)
}

// NewSteerModeValueAnyOfFromValue returns a pointer to a valid SteerModeValueAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSteerModeValueAnyOfFromValue(v string) (*SteerModeValueAnyOf, error) {
	ev := SteerModeValueAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SteerModeValueAnyOf: valid values are %v", v, AllowedSteerModeValueAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SteerModeValueAnyOf) IsValid() bool {
	for _, existing := range AllowedSteerModeValueAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SteerModeValue_anyOf value
func (v SteerModeValueAnyOf) Ptr() *SteerModeValueAnyOf {
	return &v
}

type NullableSteerModeValueAnyOf struct {
	value *SteerModeValueAnyOf
	isSet bool
}

func (v NullableSteerModeValueAnyOf) Get() *SteerModeValueAnyOf {
	return v.value
}

func (v *NullableSteerModeValueAnyOf) Set(val *SteerModeValueAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSteerModeValueAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSteerModeValueAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSteerModeValueAnyOf(val *SteerModeValueAnyOf) *NullableSteerModeValueAnyOf {
	return &NullableSteerModeValueAnyOf{value: val, isSet: true}
}

func (v NullableSteerModeValueAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSteerModeValueAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

