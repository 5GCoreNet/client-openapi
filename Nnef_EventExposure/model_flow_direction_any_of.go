/*
Nnef_EventExposure

NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnef_EventExposure

import (
	"encoding/json"
	"fmt"
)

// FlowDirectionAnyOf the model 'FlowDirectionAnyOf'
type FlowDirectionAnyOf string

// List of FlowDirection_anyOf
const (
	DOWNLINK FlowDirectionAnyOf = "DOWNLINK"
	UPLINK FlowDirectionAnyOf = "UPLINK"
	BIDIRECTIONAL FlowDirectionAnyOf = "BIDIRECTIONAL"
	UNSPECIFIED FlowDirectionAnyOf = "UNSPECIFIED"
)

// All allowed values of FlowDirectionAnyOf enum
var AllowedFlowDirectionAnyOfEnumValues = []FlowDirectionAnyOf{
	"DOWNLINK",
	"UPLINK",
	"BIDIRECTIONAL",
	"UNSPECIFIED",
}

func (v *FlowDirectionAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FlowDirectionAnyOf(value)
	for _, existing := range AllowedFlowDirectionAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FlowDirectionAnyOf", value)
}

// NewFlowDirectionAnyOfFromValue returns a pointer to a valid FlowDirectionAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlowDirectionAnyOfFromValue(v string) (*FlowDirectionAnyOf, error) {
	ev := FlowDirectionAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FlowDirectionAnyOf: valid values are %v", v, AllowedFlowDirectionAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlowDirectionAnyOf) IsValid() bool {
	for _, existing := range AllowedFlowDirectionAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlowDirection_anyOf value
func (v FlowDirectionAnyOf) Ptr() *FlowDirectionAnyOf {
	return &v
}

type NullableFlowDirectionAnyOf struct {
	value *FlowDirectionAnyOf
	isSet bool
}

func (v NullableFlowDirectionAnyOf) Get() *FlowDirectionAnyOf {
	return v.value
}

func (v *NullableFlowDirectionAnyOf) Set(val *FlowDirectionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowDirectionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowDirectionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowDirectionAnyOf(val *FlowDirectionAnyOf) *NullableFlowDirectionAnyOf {
	return &NullableFlowDirectionAnyOf{value: val, isSet: true}
}

func (v NullableFlowDirectionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowDirectionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

