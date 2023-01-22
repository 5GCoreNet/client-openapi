/*
Ntsctsf_QoSandTSCAssistance Service API

TSCTSF QoS and TSC Assistance Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ntsctsf_QoSandTSCAssistance

import (
	"encoding/json"
	"fmt"
)

// TscEventAnyOf the model 'TscEventAnyOf'
type TscEventAnyOf string

// List of TscEvent_anyOf
const (
	FAILED_RESOURCES_ALLOCATION TscEventAnyOf = "FAILED_RESOURCES_ALLOCATION"
	QOS_MONITORING TscEventAnyOf = "QOS_MONITORING"
	QOS_GUARANTEED TscEventAnyOf = "QOS_GUARANTEED"
	QOS_NOT_GUARANTEED TscEventAnyOf = "QOS_NOT_GUARANTEED"
	SUCCESSFUL_RESOURCES_ALLOCATION TscEventAnyOf = "SUCCESSFUL_RESOURCES_ALLOCATION"
	USAGE_REPORT TscEventAnyOf = "USAGE_REPORT"
)

// All allowed values of TscEventAnyOf enum
var AllowedTscEventAnyOfEnumValues = []TscEventAnyOf{
	"FAILED_RESOURCES_ALLOCATION",
	"QOS_MONITORING",
	"QOS_GUARANTEED",
	"QOS_NOT_GUARANTEED",
	"SUCCESSFUL_RESOURCES_ALLOCATION",
	"USAGE_REPORT",
}

func (v *TscEventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TscEventAnyOf(value)
	for _, existing := range AllowedTscEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TscEventAnyOf", value)
}

// NewTscEventAnyOfFromValue returns a pointer to a valid TscEventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTscEventAnyOfFromValue(v string) (*TscEventAnyOf, error) {
	ev := TscEventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TscEventAnyOf: valid values are %v", v, AllowedTscEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TscEventAnyOf) IsValid() bool {
	for _, existing := range AllowedTscEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TscEvent_anyOf value
func (v TscEventAnyOf) Ptr() *TscEventAnyOf {
	return &v
}

type NullableTscEventAnyOf struct {
	value *TscEventAnyOf
	isSet bool
}

func (v NullableTscEventAnyOf) Get() *TscEventAnyOf {
	return v.value
}

func (v *NullableTscEventAnyOf) Set(val *TscEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTscEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTscEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTscEventAnyOf(val *TscEventAnyOf) *NullableTscEventAnyOf {
	return &NullableTscEventAnyOf{value: val, isSet: true}
}

func (v NullableTscEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTscEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
