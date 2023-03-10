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

// ReportingAccessTypeAnyOf the model 'ReportingAccessTypeAnyOf'
type ReportingAccessTypeAnyOf string

// List of ReportingAccessType_anyOf
const (
	NR ReportingAccessTypeAnyOf = "NR"
	EUTRA_CONNECTED_TO_5_GC ReportingAccessTypeAnyOf = "EUTRA_CONNECTED_TO_5GC"
	NON_3_GPP_CONNECTED_TO_5_GC ReportingAccessTypeAnyOf = "NON_3GPP_CONNECTED_TO_5GC"
	NR_LEO ReportingAccessTypeAnyOf = "NR_LEO"
	NR_MEO ReportingAccessTypeAnyOf = "NR_MEO"
	NR_GEO ReportingAccessTypeAnyOf = "NR_GEO"
	NR_OTHER_SAT ReportingAccessTypeAnyOf = "NR_OTHER_SAT"
)

// All allowed values of ReportingAccessTypeAnyOf enum
var AllowedReportingAccessTypeAnyOfEnumValues = []ReportingAccessTypeAnyOf{
	"NR",
	"EUTRA_CONNECTED_TO_5GC",
	"NON_3GPP_CONNECTED_TO_5GC",
	"NR_LEO",
	"NR_MEO",
	"NR_GEO",
	"NR_OTHER_SAT",
}

func (v *ReportingAccessTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportingAccessTypeAnyOf(value)
	for _, existing := range AllowedReportingAccessTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportingAccessTypeAnyOf", value)
}

// NewReportingAccessTypeAnyOfFromValue returns a pointer to a valid ReportingAccessTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportingAccessTypeAnyOfFromValue(v string) (*ReportingAccessTypeAnyOf, error) {
	ev := ReportingAccessTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportingAccessTypeAnyOf: valid values are %v", v, AllowedReportingAccessTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportingAccessTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedReportingAccessTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportingAccessType_anyOf value
func (v ReportingAccessTypeAnyOf) Ptr() *ReportingAccessTypeAnyOf {
	return &v
}

type NullableReportingAccessTypeAnyOf struct {
	value *ReportingAccessTypeAnyOf
	isSet bool
}

func (v NullableReportingAccessTypeAnyOf) Get() *ReportingAccessTypeAnyOf {
	return v.value
}

func (v *NullableReportingAccessTypeAnyOf) Set(val *ReportingAccessTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingAccessTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingAccessTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingAccessTypeAnyOf(val *ReportingAccessTypeAnyOf) *NullableReportingAccessTypeAnyOf {
	return &NullableReportingAccessTypeAnyOf{value: val, isSet: true}
}

func (v NullableReportingAccessTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingAccessTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

