/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// ControlLoopLifeCyclePhase the model 'ControlLoopLifeCyclePhase'
type ControlLoopLifeCyclePhase string

// List of ControlLoopLifeCyclePhase
const (
	PREPARATION ControlLoopLifeCyclePhase = "PREPARATION"
	COMMISSIONING ControlLoopLifeCyclePhase = "COMMISSIONING"
	OPERATION ControlLoopLifeCyclePhase = "OPERATION"
	DECOMMISSIONING ControlLoopLifeCyclePhase = "DECOMMISSIONING"
)

// All allowed values of ControlLoopLifeCyclePhase enum
var AllowedControlLoopLifeCyclePhaseEnumValues = []ControlLoopLifeCyclePhase{
	"PREPARATION",
	"COMMISSIONING",
	"OPERATION",
	"DECOMMISSIONING",
}

func (v *ControlLoopLifeCyclePhase) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ControlLoopLifeCyclePhase(value)
	for _, existing := range AllowedControlLoopLifeCyclePhaseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ControlLoopLifeCyclePhase", value)
}

// NewControlLoopLifeCyclePhaseFromValue returns a pointer to a valid ControlLoopLifeCyclePhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewControlLoopLifeCyclePhaseFromValue(v string) (*ControlLoopLifeCyclePhase, error) {
	ev := ControlLoopLifeCyclePhase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ControlLoopLifeCyclePhase: valid values are %v", v, AllowedControlLoopLifeCyclePhaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ControlLoopLifeCyclePhase) IsValid() bool {
	for _, existing := range AllowedControlLoopLifeCyclePhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ControlLoopLifeCyclePhase value
func (v ControlLoopLifeCyclePhase) Ptr() *ControlLoopLifeCyclePhase {
	return &v
}

type NullableControlLoopLifeCyclePhase struct {
	value *ControlLoopLifeCyclePhase
	isSet bool
}

func (v NullableControlLoopLifeCyclePhase) Get() *ControlLoopLifeCyclePhase {
	return v.value
}

func (v *NullableControlLoopLifeCyclePhase) Set(val *ControlLoopLifeCyclePhase) {
	v.value = val
	v.isSet = true
}

func (v NullableControlLoopLifeCyclePhase) IsSet() bool {
	return v.isSet
}

func (v *NullableControlLoopLifeCyclePhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlLoopLifeCyclePhase(val *ControlLoopLifeCyclePhase) *NullableControlLoopLifeCyclePhase {
	return &NullableControlLoopLifeCyclePhase{value: val, isSet: true}
}

func (v NullableControlLoopLifeCyclePhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlLoopLifeCyclePhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

