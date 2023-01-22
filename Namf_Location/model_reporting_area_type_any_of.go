/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Location

import (
	"encoding/json"
	"fmt"
)

// ReportingAreaTypeAnyOf the model 'ReportingAreaTypeAnyOf'
type ReportingAreaTypeAnyOf string

// List of ReportingAreaType_anyOf
const (
	EPS_TRACKING_AREA_IDENTITY ReportingAreaTypeAnyOf = "EPS_TRACKING_AREA_IDENTITY"
	E_UTRAN_CELL_GLOBAL_IDENTIFICATION ReportingAreaTypeAnyOf = "E-UTRAN_CELL_GLOBAL_IDENTIFICATION"
	_5_GS_TRACKING_AREA_IDENTITY ReportingAreaTypeAnyOf = "5GS_TRACKING_AREA_IDENTITY"
	NR_CELL_GLOBAL_IDENTITY ReportingAreaTypeAnyOf = "NR_CELL_GLOBAL_IDENTITY"
)

// All allowed values of ReportingAreaTypeAnyOf enum
var AllowedReportingAreaTypeAnyOfEnumValues = []ReportingAreaTypeAnyOf{
	"EPS_TRACKING_AREA_IDENTITY",
	"E-UTRAN_CELL_GLOBAL_IDENTIFICATION",
	"5GS_TRACKING_AREA_IDENTITY",
	"NR_CELL_GLOBAL_IDENTITY",
}

func (v *ReportingAreaTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportingAreaTypeAnyOf(value)
	for _, existing := range AllowedReportingAreaTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportingAreaTypeAnyOf", value)
}

// NewReportingAreaTypeAnyOfFromValue returns a pointer to a valid ReportingAreaTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportingAreaTypeAnyOfFromValue(v string) (*ReportingAreaTypeAnyOf, error) {
	ev := ReportingAreaTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportingAreaTypeAnyOf: valid values are %v", v, AllowedReportingAreaTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportingAreaTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedReportingAreaTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportingAreaType_anyOf value
func (v ReportingAreaTypeAnyOf) Ptr() *ReportingAreaTypeAnyOf {
	return &v
}

type NullableReportingAreaTypeAnyOf struct {
	value *ReportingAreaTypeAnyOf
	isSet bool
}

func (v NullableReportingAreaTypeAnyOf) Get() *ReportingAreaTypeAnyOf {
	return v.value
}

func (v *NullableReportingAreaTypeAnyOf) Set(val *ReportingAreaTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingAreaTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingAreaTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingAreaTypeAnyOf(val *ReportingAreaTypeAnyOf) *NullableReportingAreaTypeAnyOf {
	return &NullableReportingAreaTypeAnyOf{value: val, isSet: true}
}

func (v NullableReportingAreaTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingAreaTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

