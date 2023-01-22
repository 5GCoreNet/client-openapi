/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
	"fmt"
)

// JobTypeAnyOf the model 'JobTypeAnyOf'
type JobTypeAnyOf string

// List of JobType_anyOf
const (
	IMMEDIATE_MDT_ONLY JobTypeAnyOf = "IMMEDIATE_MDT_ONLY"
	LOGGED_MDT_ONLY JobTypeAnyOf = "LOGGED_MDT_ONLY"
	TRACE_ONLY JobTypeAnyOf = "TRACE_ONLY"
	IMMEDIATE_MDT_AND_TRACE JobTypeAnyOf = "IMMEDIATE_MDT_AND_TRACE"
	RLF_REPORTS_ONLY JobTypeAnyOf = "RLF_REPORTS_ONLY"
	RCEF_REPORTS_ONLY JobTypeAnyOf = "RCEF_REPORTS_ONLY"
	LOGGED_MBSFN_MDT JobTypeAnyOf = "LOGGED_MBSFN_MDT"
)

// All allowed values of JobTypeAnyOf enum
var AllowedJobTypeAnyOfEnumValues = []JobTypeAnyOf{
	"IMMEDIATE_MDT_ONLY",
	"LOGGED_MDT_ONLY",
	"TRACE_ONLY",
	"IMMEDIATE_MDT_AND_TRACE",
	"RLF_REPORTS_ONLY",
	"RCEF_REPORTS_ONLY",
	"LOGGED_MBSFN_MDT",
}

func (v *JobTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobTypeAnyOf(value)
	for _, existing := range AllowedJobTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobTypeAnyOf", value)
}

// NewJobTypeAnyOfFromValue returns a pointer to a valid JobTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobTypeAnyOfFromValue(v string) (*JobTypeAnyOf, error) {
	ev := JobTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobTypeAnyOf: valid values are %v", v, AllowedJobTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedJobTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JobType_anyOf value
func (v JobTypeAnyOf) Ptr() *JobTypeAnyOf {
	return &v
}

type NullableJobTypeAnyOf struct {
	value *JobTypeAnyOf
	isSet bool
}

func (v NullableJobTypeAnyOf) Get() *JobTypeAnyOf {
	return v.value
}

func (v *NullableJobTypeAnyOf) Set(val *JobTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableJobTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableJobTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobTypeAnyOf(val *JobTypeAnyOf) *NullableJobTypeAnyOf {
	return &NullableJobTypeAnyOf{value: val, isSet: true}
}

func (v NullableJobTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

