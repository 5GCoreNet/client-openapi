/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// ScpCapabilityAnyOf the model 'ScpCapabilityAnyOf'
type ScpCapabilityAnyOf string

// List of ScpCapability_anyOf
const (
	INDIRECT_COM_WITH_DELEG_DISC ScpCapabilityAnyOf = "INDIRECT_COM_WITH_DELEG_DISC"
)

// All allowed values of ScpCapabilityAnyOf enum
var AllowedScpCapabilityAnyOfEnumValues = []ScpCapabilityAnyOf{
	"INDIRECT_COM_WITH_DELEG_DISC",
}

func (v *ScpCapabilityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScpCapabilityAnyOf(value)
	for _, existing := range AllowedScpCapabilityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScpCapabilityAnyOf", value)
}

// NewScpCapabilityAnyOfFromValue returns a pointer to a valid ScpCapabilityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScpCapabilityAnyOfFromValue(v string) (*ScpCapabilityAnyOf, error) {
	ev := ScpCapabilityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScpCapabilityAnyOf: valid values are %v", v, AllowedScpCapabilityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScpCapabilityAnyOf) IsValid() bool {
	for _, existing := range AllowedScpCapabilityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScpCapability_anyOf value
func (v ScpCapabilityAnyOf) Ptr() *ScpCapabilityAnyOf {
	return &v
}

type NullableScpCapabilityAnyOf struct {
	value *ScpCapabilityAnyOf
	isSet bool
}

func (v NullableScpCapabilityAnyOf) Get() *ScpCapabilityAnyOf {
	return v.value
}

func (v *NullableScpCapabilityAnyOf) Set(val *ScpCapabilityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableScpCapabilityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableScpCapabilityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScpCapabilityAnyOf(val *ScpCapabilityAnyOf) *NullableScpCapabilityAnyOf {
	return &NullableScpCapabilityAnyOf{value: val, isSet: true}
}

func (v NullableScpCapabilityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScpCapabilityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

