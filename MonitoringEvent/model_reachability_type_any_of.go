/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
	"fmt"
)

// ReachabilityTypeAnyOf the model 'ReachabilityTypeAnyOf'
type ReachabilityTypeAnyOf string

// List of ReachabilityType_anyOf
const (
	SMS ReachabilityTypeAnyOf = "SMS"
	DATA ReachabilityTypeAnyOf = "DATA"
)

// All allowed values of ReachabilityTypeAnyOf enum
var AllowedReachabilityTypeAnyOfEnumValues = []ReachabilityTypeAnyOf{
	"SMS",
	"DATA",
}

func (v *ReachabilityTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReachabilityTypeAnyOf(value)
	for _, existing := range AllowedReachabilityTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReachabilityTypeAnyOf", value)
}

// NewReachabilityTypeAnyOfFromValue returns a pointer to a valid ReachabilityTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReachabilityTypeAnyOfFromValue(v string) (*ReachabilityTypeAnyOf, error) {
	ev := ReachabilityTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReachabilityTypeAnyOf: valid values are %v", v, AllowedReachabilityTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReachabilityTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedReachabilityTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReachabilityType_anyOf value
func (v ReachabilityTypeAnyOf) Ptr() *ReachabilityTypeAnyOf {
	return &v
}

type NullableReachabilityTypeAnyOf struct {
	value *ReachabilityTypeAnyOf
	isSet bool
}

func (v NullableReachabilityTypeAnyOf) Get() *ReachabilityTypeAnyOf {
	return v.value
}

func (v *NullableReachabilityTypeAnyOf) Set(val *ReachabilityTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityTypeAnyOf(val *ReachabilityTypeAnyOf) *NullableReachabilityTypeAnyOf {
	return &NullableReachabilityTypeAnyOf{value: val, isSet: true}
}

func (v NullableReachabilityTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

