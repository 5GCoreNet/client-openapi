/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// OutputStrategyAnyOf the model 'OutputStrategyAnyOf'
type OutputStrategyAnyOf string

// List of OutputStrategy_anyOf
const (
	BINARY OutputStrategyAnyOf = "BINARY"
	GRADIENT OutputStrategyAnyOf = "GRADIENT"
)

// All allowed values of OutputStrategyAnyOf enum
var AllowedOutputStrategyAnyOfEnumValues = []OutputStrategyAnyOf{
	"BINARY",
	"GRADIENT",
}

func (v *OutputStrategyAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OutputStrategyAnyOf(value)
	for _, existing := range AllowedOutputStrategyAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OutputStrategyAnyOf", value)
}

// NewOutputStrategyAnyOfFromValue returns a pointer to a valid OutputStrategyAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOutputStrategyAnyOfFromValue(v string) (*OutputStrategyAnyOf, error) {
	ev := OutputStrategyAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OutputStrategyAnyOf: valid values are %v", v, AllowedOutputStrategyAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OutputStrategyAnyOf) IsValid() bool {
	for _, existing := range AllowedOutputStrategyAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OutputStrategy_anyOf value
func (v OutputStrategyAnyOf) Ptr() *OutputStrategyAnyOf {
	return &v
}

type NullableOutputStrategyAnyOf struct {
	value *OutputStrategyAnyOf
	isSet bool
}

func (v NullableOutputStrategyAnyOf) Get() *OutputStrategyAnyOf {
	return v.value
}

func (v *NullableOutputStrategyAnyOf) Set(val *OutputStrategyAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputStrategyAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputStrategyAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputStrategyAnyOf(val *OutputStrategyAnyOf) *NullableOutputStrategyAnyOf {
	return &NullableOutputStrategyAnyOf{value: val, isSet: true}
}

func (v NullableOutputStrategyAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputStrategyAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

