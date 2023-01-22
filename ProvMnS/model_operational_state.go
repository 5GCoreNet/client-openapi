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

// OperationalState the model 'OperationalState'
type OperationalState string

// List of OperationalState
const (
	ENABLED OperationalState = "ENABLED"
	DISABLED OperationalState = "DISABLED"
)

// All allowed values of OperationalState enum
var AllowedOperationalStateEnumValues = []OperationalState{
	"ENABLED",
	"DISABLED",
}

func (v *OperationalState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperationalState(value)
	for _, existing := range AllowedOperationalStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperationalState", value)
}

// NewOperationalStateFromValue returns a pointer to a valid OperationalState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperationalStateFromValue(v string) (*OperationalState, error) {
	ev := OperationalState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperationalState: valid values are %v", v, AllowedOperationalStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperationalState) IsValid() bool {
	for _, existing := range AllowedOperationalStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperationalState value
func (v OperationalState) Ptr() *OperationalState {
	return &v
}

type NullableOperationalState struct {
	value *OperationalState
	isSet bool
}

func (v NullableOperationalState) Get() *OperationalState {
	return v.value
}

func (v *NullableOperationalState) Set(val *OperationalState) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationalState) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationalState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationalState(val *OperationalState) *NullableOperationalState {
	return &NullableOperationalState{value: val, isSet: true}
}

func (v NullableOperationalState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationalState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

