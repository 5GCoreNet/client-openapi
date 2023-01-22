/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_OfflineOnlyCharging

import (
	"encoding/json"
	"fmt"
)

// SteeringFunctionalityAnyOf the model 'SteeringFunctionalityAnyOf'
type SteeringFunctionalityAnyOf string

// List of SteeringFunctionality_anyOf
const (
	MPTCP SteeringFunctionalityAnyOf = "MPTCP"
	ATSSS_LL SteeringFunctionalityAnyOf = "ATSSS_LL"
)

// All allowed values of SteeringFunctionalityAnyOf enum
var AllowedSteeringFunctionalityAnyOfEnumValues = []SteeringFunctionalityAnyOf{
	"MPTCP",
	"ATSSS_LL",
}

func (v *SteeringFunctionalityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SteeringFunctionalityAnyOf(value)
	for _, existing := range AllowedSteeringFunctionalityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SteeringFunctionalityAnyOf", value)
}

// NewSteeringFunctionalityAnyOfFromValue returns a pointer to a valid SteeringFunctionalityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSteeringFunctionalityAnyOfFromValue(v string) (*SteeringFunctionalityAnyOf, error) {
	ev := SteeringFunctionalityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SteeringFunctionalityAnyOf: valid values are %v", v, AllowedSteeringFunctionalityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SteeringFunctionalityAnyOf) IsValid() bool {
	for _, existing := range AllowedSteeringFunctionalityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SteeringFunctionality_anyOf value
func (v SteeringFunctionalityAnyOf) Ptr() *SteeringFunctionalityAnyOf {
	return &v
}

type NullableSteeringFunctionalityAnyOf struct {
	value *SteeringFunctionalityAnyOf
	isSet bool
}

func (v NullableSteeringFunctionalityAnyOf) Get() *SteeringFunctionalityAnyOf {
	return v.value
}

func (v *NullableSteeringFunctionalityAnyOf) Set(val *SteeringFunctionalityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSteeringFunctionalityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSteeringFunctionalityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSteeringFunctionalityAnyOf(val *SteeringFunctionalityAnyOf) *NullableSteeringFunctionalityAnyOf {
	return &NullableSteeringFunctionalityAnyOf{value: val, isSet: true}
}

func (v NullableSteeringFunctionalityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSteeringFunctionalityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

