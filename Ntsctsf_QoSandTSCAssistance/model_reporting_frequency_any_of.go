/*
Ntsctsf_QoSandTSCAssistance Service API

TSCTSF QoS and TSC Assistance Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_QoSandTSCAssistance

import (
	"encoding/json"
	"fmt"
)

// ReportingFrequencyAnyOf the model 'ReportingFrequencyAnyOf'
type ReportingFrequencyAnyOf string

// List of ReportingFrequency_anyOf
const (
	EVENT_TRIGGERED ReportingFrequencyAnyOf = "EVENT_TRIGGERED"
	PERIODIC ReportingFrequencyAnyOf = "PERIODIC"
	SESSION_RELEASE ReportingFrequencyAnyOf = "SESSION_RELEASE"
)

// All allowed values of ReportingFrequencyAnyOf enum
var AllowedReportingFrequencyAnyOfEnumValues = []ReportingFrequencyAnyOf{
	"EVENT_TRIGGERED",
	"PERIODIC",
	"SESSION_RELEASE",
}

func (v *ReportingFrequencyAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportingFrequencyAnyOf(value)
	for _, existing := range AllowedReportingFrequencyAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportingFrequencyAnyOf", value)
}

// NewReportingFrequencyAnyOfFromValue returns a pointer to a valid ReportingFrequencyAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportingFrequencyAnyOfFromValue(v string) (*ReportingFrequencyAnyOf, error) {
	ev := ReportingFrequencyAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportingFrequencyAnyOf: valid values are %v", v, AllowedReportingFrequencyAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportingFrequencyAnyOf) IsValid() bool {
	for _, existing := range AllowedReportingFrequencyAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportingFrequency_anyOf value
func (v ReportingFrequencyAnyOf) Ptr() *ReportingFrequencyAnyOf {
	return &v
}

type NullableReportingFrequencyAnyOf struct {
	value *ReportingFrequencyAnyOf
	isSet bool
}

func (v NullableReportingFrequencyAnyOf) Get() *ReportingFrequencyAnyOf {
	return v.value
}

func (v *NullableReportingFrequencyAnyOf) Set(val *ReportingFrequencyAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingFrequencyAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingFrequencyAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingFrequencyAnyOf(val *ReportingFrequencyAnyOf) *NullableReportingFrequencyAnyOf {
	return &NullableReportingFrequencyAnyOf{value: val, isSet: true}
}

func (v NullableReportingFrequencyAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingFrequencyAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

