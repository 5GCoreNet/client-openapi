/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_OfflineOnlyCharging

import (
	"encoding/json"
	"fmt"
)

// SteerModeIndicatorAnyOf the model 'SteerModeIndicatorAnyOf'
type SteerModeIndicatorAnyOf string

// List of SteerModeIndicator_anyOf
const (
	AUTO_LOAD_BALANCE SteerModeIndicatorAnyOf = "AUTO_LOAD_BALANCE"
	UE_ASSISTANCE SteerModeIndicatorAnyOf = "UE_ASSISTANCE"
)

// All allowed values of SteerModeIndicatorAnyOf enum
var AllowedSteerModeIndicatorAnyOfEnumValues = []SteerModeIndicatorAnyOf{
	"AUTO_LOAD_BALANCE",
	"UE_ASSISTANCE",
}

func (v *SteerModeIndicatorAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SteerModeIndicatorAnyOf(value)
	for _, existing := range AllowedSteerModeIndicatorAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SteerModeIndicatorAnyOf", value)
}

// NewSteerModeIndicatorAnyOfFromValue returns a pointer to a valid SteerModeIndicatorAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSteerModeIndicatorAnyOfFromValue(v string) (*SteerModeIndicatorAnyOf, error) {
	ev := SteerModeIndicatorAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SteerModeIndicatorAnyOf: valid values are %v", v, AllowedSteerModeIndicatorAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SteerModeIndicatorAnyOf) IsValid() bool {
	for _, existing := range AllowedSteerModeIndicatorAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SteerModeIndicator_anyOf value
func (v SteerModeIndicatorAnyOf) Ptr() *SteerModeIndicatorAnyOf {
	return &v
}

type NullableSteerModeIndicatorAnyOf struct {
	value *SteerModeIndicatorAnyOf
	isSet bool
}

func (v NullableSteerModeIndicatorAnyOf) Get() *SteerModeIndicatorAnyOf {
	return v.value
}

func (v *NullableSteerModeIndicatorAnyOf) Set(val *SteerModeIndicatorAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSteerModeIndicatorAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSteerModeIndicatorAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSteerModeIndicatorAnyOf(val *SteerModeIndicatorAnyOf) *NullableSteerModeIndicatorAnyOf {
	return &NullableSteerModeIndicatorAnyOf{value: val, isSet: true}
}

func (v NullableSteerModeIndicatorAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSteerModeIndicatorAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

