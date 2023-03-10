/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// FlowUsageAnyOf the model 'FlowUsageAnyOf'
type FlowUsageAnyOf string

// List of FlowUsage_anyOf
const (
	NO_INFO FlowUsageAnyOf = "NO_INFO"
	RTCP FlowUsageAnyOf = "RTCP"
	AF_SIGNALLING FlowUsageAnyOf = "AF_SIGNALLING"
)

// All allowed values of FlowUsageAnyOf enum
var AllowedFlowUsageAnyOfEnumValues = []FlowUsageAnyOf{
	"NO_INFO",
	"RTCP",
	"AF_SIGNALLING",
}

func (v *FlowUsageAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FlowUsageAnyOf(value)
	for _, existing := range AllowedFlowUsageAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FlowUsageAnyOf", value)
}

// NewFlowUsageAnyOfFromValue returns a pointer to a valid FlowUsageAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlowUsageAnyOfFromValue(v string) (*FlowUsageAnyOf, error) {
	ev := FlowUsageAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FlowUsageAnyOf: valid values are %v", v, AllowedFlowUsageAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlowUsageAnyOf) IsValid() bool {
	for _, existing := range AllowedFlowUsageAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlowUsage_anyOf value
func (v FlowUsageAnyOf) Ptr() *FlowUsageAnyOf {
	return &v
}

type NullableFlowUsageAnyOf struct {
	value *FlowUsageAnyOf
	isSet bool
}

func (v NullableFlowUsageAnyOf) Get() *FlowUsageAnyOf {
	return v.value
}

func (v *NullableFlowUsageAnyOf) Set(val *FlowUsageAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowUsageAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowUsageAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowUsageAnyOf(val *FlowUsageAnyOf) *NullableFlowUsageAnyOf {
	return &NullableFlowUsageAnyOf{value: val, isSet: true}
}

func (v NullableFlowUsageAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowUsageAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

