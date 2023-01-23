/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_EE

import (
	"encoding/json"
	"fmt"
)

// FailedCauseAnyOf the model 'FailedCauseAnyOf'
type FailedCauseAnyOf string

// List of FailedCause_anyOf
const (
	AF_NOT_ALLOWED FailedCauseAnyOf = "AF_NOT_ALLOWED"
	MTC_PROVIDER_NOT_ALLOWED FailedCauseAnyOf = "MTC_PROVIDER_NOT_ALLOWED"
	MONITORING_NOT_ALLOWED FailedCauseAnyOf = "MONITORING_NOT_ALLOWED"
	UNSUPPORTED_MONITORING_EVENT_TYPE FailedCauseAnyOf = "UNSUPPORTED_MONITORING_EVENT_TYPE"
	UNSUPPORTED_MONITORING_REPORT_OPTIONS FailedCauseAnyOf = "UNSUPPORTED_MONITORING_REPORT_OPTIONS"
	UNSPECIFIED FailedCauseAnyOf = "UNSPECIFIED"
)

// All allowed values of FailedCauseAnyOf enum
var AllowedFailedCauseAnyOfEnumValues = []FailedCauseAnyOf{
	"AF_NOT_ALLOWED",
	"MTC_PROVIDER_NOT_ALLOWED",
	"MONITORING_NOT_ALLOWED",
	"UNSUPPORTED_MONITORING_EVENT_TYPE",
	"UNSUPPORTED_MONITORING_REPORT_OPTIONS",
	"UNSPECIFIED",
}

func (v *FailedCauseAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FailedCauseAnyOf(value)
	for _, existing := range AllowedFailedCauseAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FailedCauseAnyOf", value)
}

// NewFailedCauseAnyOfFromValue returns a pointer to a valid FailedCauseAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFailedCauseAnyOfFromValue(v string) (*FailedCauseAnyOf, error) {
	ev := FailedCauseAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FailedCauseAnyOf: valid values are %v", v, AllowedFailedCauseAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FailedCauseAnyOf) IsValid() bool {
	for _, existing := range AllowedFailedCauseAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FailedCause_anyOf value
func (v FailedCauseAnyOf) Ptr() *FailedCauseAnyOf {
	return &v
}

type NullableFailedCauseAnyOf struct {
	value *FailedCauseAnyOf
	isSet bool
}

func (v NullableFailedCauseAnyOf) Get() *FailedCauseAnyOf {
	return v.value
}

func (v *NullableFailedCauseAnyOf) Set(val *FailedCauseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFailedCauseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFailedCauseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailedCauseAnyOf(val *FailedCauseAnyOf) *NullableFailedCauseAnyOf {
	return &NullableFailedCauseAnyOf{value: val, isSet: true}
}

func (v NullableFailedCauseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailedCauseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

