/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// FlowStatusAnyOf the model 'FlowStatusAnyOf'
type FlowStatusAnyOf string

// List of FlowStatus_anyOf
const (
	ENABLED_UPLINK FlowStatusAnyOf = "ENABLED-UPLINK"
	ENABLED_DOWNLINK FlowStatusAnyOf = "ENABLED-DOWNLINK"
	ENABLED FlowStatusAnyOf = "ENABLED"
	DISABLED FlowStatusAnyOf = "DISABLED"
	REMOVED FlowStatusAnyOf = "REMOVED"
)

// All allowed values of FlowStatusAnyOf enum
var AllowedFlowStatusAnyOfEnumValues = []FlowStatusAnyOf{
	"ENABLED-UPLINK",
	"ENABLED-DOWNLINK",
	"ENABLED",
	"DISABLED",
	"REMOVED",
}

func (v *FlowStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FlowStatusAnyOf(value)
	for _, existing := range AllowedFlowStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FlowStatusAnyOf", value)
}

// NewFlowStatusAnyOfFromValue returns a pointer to a valid FlowStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlowStatusAnyOfFromValue(v string) (*FlowStatusAnyOf, error) {
	ev := FlowStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FlowStatusAnyOf: valid values are %v", v, AllowedFlowStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlowStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedFlowStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlowStatus_anyOf value
func (v FlowStatusAnyOf) Ptr() *FlowStatusAnyOf {
	return &v
}

type NullableFlowStatusAnyOf struct {
	value *FlowStatusAnyOf
	isSet bool
}

func (v NullableFlowStatusAnyOf) Get() *FlowStatusAnyOf {
	return v.value
}

func (v *NullableFlowStatusAnyOf) Set(val *FlowStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowStatusAnyOf(val *FlowStatusAnyOf) *NullableFlowStatusAnyOf {
	return &NullableFlowStatusAnyOf{value: val, isSet: true}
}

func (v NullableFlowStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

