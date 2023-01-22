/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
	"fmt"
)

// NotFulfilledState the model 'NotFulfilledState'
type NotFulfilledState string

// List of NotFulfilledState
const (
	ACKNOWLEDGED NotFulfilledState = "ACKNOWLEDGED"
	COMPLIANT NotFulfilledState = "COMPLIANT"
	DEGRADED NotFulfilledState = "DEGRADED"
	SUSPENDED NotFulfilledState = "SUSPENDED"
	TERMINATED NotFulfilledState = "TERMINATED"
	FULFILMENTFAILED NotFulfilledState = "FULFILMENTFAILED"
)

// All allowed values of NotFulfilledState enum
var AllowedNotFulfilledStateEnumValues = []NotFulfilledState{
	"ACKNOWLEDGED",
	"COMPLIANT",
	"DEGRADED",
	"SUSPENDED",
	"TERMINATED",
	"FULFILMENTFAILED",
}

func (v *NotFulfilledState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotFulfilledState(value)
	for _, existing := range AllowedNotFulfilledStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotFulfilledState", value)
}

// NewNotFulfilledStateFromValue returns a pointer to a valid NotFulfilledState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotFulfilledStateFromValue(v string) (*NotFulfilledState, error) {
	ev := NotFulfilledState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotFulfilledState: valid values are %v", v, AllowedNotFulfilledStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotFulfilledState) IsValid() bool {
	for _, existing := range AllowedNotFulfilledStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotFulfilledState value
func (v NotFulfilledState) Ptr() *NotFulfilledState {
	return &v
}

type NullableNotFulfilledState struct {
	value *NotFulfilledState
	isSet bool
}

func (v NullableNotFulfilledState) Get() *NotFulfilledState {
	return v.value
}

func (v *NullableNotFulfilledState) Set(val *NotFulfilledState) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFulfilledState) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFulfilledState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFulfilledState(val *NotFulfilledState) *NullableNotFulfilledState {
	return &NullableNotFulfilledState{value: val, isSet: true}
}

func (v NullableNotFulfilledState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotFulfilledState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

