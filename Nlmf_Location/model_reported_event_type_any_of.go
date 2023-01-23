/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Location

import (
	"encoding/json"
	"fmt"
)

// ReportedEventTypeAnyOf the model 'ReportedEventTypeAnyOf'
type ReportedEventTypeAnyOf string

// List of ReportedEventType_anyOf
const (
	PERIODIC_EVENT ReportedEventTypeAnyOf = "PERIODIC_EVENT"
	ENTERING_AREA_EVENT ReportedEventTypeAnyOf = "ENTERING_AREA_EVENT"
	LEAVING_AREA_EVENT ReportedEventTypeAnyOf = "LEAVING_AREA_EVENT"
	BEING_INSIDE_AREA_EVENT ReportedEventTypeAnyOf = "BEING_INSIDE_AREA_EVENT"
	MOTION_EVENT ReportedEventTypeAnyOf = "MOTION_EVENT"
	MAXIMUM_INTERVAL_EXPIRATION_EVENT ReportedEventTypeAnyOf = "MAXIMUM_INTERVAL_EXPIRATION_EVENT"
	LOCATION_CANCELLATION_EVENT ReportedEventTypeAnyOf = "LOCATION_CANCELLATION_EVENT"
)

// All allowed values of ReportedEventTypeAnyOf enum
var AllowedReportedEventTypeAnyOfEnumValues = []ReportedEventTypeAnyOf{
	"PERIODIC_EVENT",
	"ENTERING_AREA_EVENT",
	"LEAVING_AREA_EVENT",
	"BEING_INSIDE_AREA_EVENT",
	"MOTION_EVENT",
	"MAXIMUM_INTERVAL_EXPIRATION_EVENT",
	"LOCATION_CANCELLATION_EVENT",
}

func (v *ReportedEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportedEventTypeAnyOf(value)
	for _, existing := range AllowedReportedEventTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportedEventTypeAnyOf", value)
}

// NewReportedEventTypeAnyOfFromValue returns a pointer to a valid ReportedEventTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportedEventTypeAnyOfFromValue(v string) (*ReportedEventTypeAnyOf, error) {
	ev := ReportedEventTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportedEventTypeAnyOf: valid values are %v", v, AllowedReportedEventTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportedEventTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedReportedEventTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportedEventType_anyOf value
func (v ReportedEventTypeAnyOf) Ptr() *ReportedEventTypeAnyOf {
	return &v
}

type NullableReportedEventTypeAnyOf struct {
	value *ReportedEventTypeAnyOf
	isSet bool
}

func (v NullableReportedEventTypeAnyOf) Get() *ReportedEventTypeAnyOf {
	return v.value
}

func (v *NullableReportedEventTypeAnyOf) Set(val *ReportedEventTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportedEventTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportedEventTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportedEventTypeAnyOf(val *ReportedEventTypeAnyOf) *NullableReportedEventTypeAnyOf {
	return &NullableReportedEventTypeAnyOf{value: val, isSet: true}
}

func (v NullableReportedEventTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportedEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

