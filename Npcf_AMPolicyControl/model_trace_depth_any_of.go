/*
Npcf_AMPolicyControl

Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_AMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// TraceDepthAnyOf the model 'TraceDepthAnyOf'
type TraceDepthAnyOf string

// List of TraceDepth_anyOf
const (
	MINIMUM TraceDepthAnyOf = "MINIMUM"
	MEDIUM TraceDepthAnyOf = "MEDIUM"
	MAXIMUM TraceDepthAnyOf = "MAXIMUM"
	MINIMUM_WO_VENDOR_EXTENSION TraceDepthAnyOf = "MINIMUM_WO_VENDOR_EXTENSION"
	MEDIUM_WO_VENDOR_EXTENSION TraceDepthAnyOf = "MEDIUM_WO_VENDOR_EXTENSION"
	MAXIMUM_WO_VENDOR_EXTENSION TraceDepthAnyOf = "MAXIMUM_WO_VENDOR_EXTENSION"
)

// All allowed values of TraceDepthAnyOf enum
var AllowedTraceDepthAnyOfEnumValues = []TraceDepthAnyOf{
	"MINIMUM",
	"MEDIUM",
	"MAXIMUM",
	"MINIMUM_WO_VENDOR_EXTENSION",
	"MEDIUM_WO_VENDOR_EXTENSION",
	"MAXIMUM_WO_VENDOR_EXTENSION",
}

func (v *TraceDepthAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TraceDepthAnyOf(value)
	for _, existing := range AllowedTraceDepthAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TraceDepthAnyOf", value)
}

// NewTraceDepthAnyOfFromValue returns a pointer to a valid TraceDepthAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTraceDepthAnyOfFromValue(v string) (*TraceDepthAnyOf, error) {
	ev := TraceDepthAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TraceDepthAnyOf: valid values are %v", v, AllowedTraceDepthAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TraceDepthAnyOf) IsValid() bool {
	for _, existing := range AllowedTraceDepthAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TraceDepth_anyOf value
func (v TraceDepthAnyOf) Ptr() *TraceDepthAnyOf {
	return &v
}

type NullableTraceDepthAnyOf struct {
	value *TraceDepthAnyOf
	isSet bool
}

func (v NullableTraceDepthAnyOf) Get() *TraceDepthAnyOf {
	return v.value
}

func (v *NullableTraceDepthAnyOf) Set(val *TraceDepthAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceDepthAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceDepthAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceDepthAnyOf(val *TraceDepthAnyOf) *NullableTraceDepthAnyOf {
	return &NullableTraceDepthAnyOf{value: val, isSet: true}
}

func (v NullableTraceDepthAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceDepthAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

