/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// ReportAmountMdtAnyOf the model 'ReportAmountMdtAnyOf'
type ReportAmountMdtAnyOf string

// List of ReportAmountMdt_anyOf
const (
	_1 ReportAmountMdtAnyOf = "1"
	_2 ReportAmountMdtAnyOf = "2"
	_4 ReportAmountMdtAnyOf = "4"
	_8 ReportAmountMdtAnyOf = "8"
	_16 ReportAmountMdtAnyOf = "16"
	_32 ReportAmountMdtAnyOf = "32"
	_64 ReportAmountMdtAnyOf = "64"
	INFINITY ReportAmountMdtAnyOf = "infinity"
)

// All allowed values of ReportAmountMdtAnyOf enum
var AllowedReportAmountMdtAnyOfEnumValues = []ReportAmountMdtAnyOf{
	"1",
	"2",
	"4",
	"8",
	"16",
	"32",
	"64",
	"infinity",
}

func (v *ReportAmountMdtAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportAmountMdtAnyOf(value)
	for _, existing := range AllowedReportAmountMdtAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportAmountMdtAnyOf", value)
}

// NewReportAmountMdtAnyOfFromValue returns a pointer to a valid ReportAmountMdtAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAmountMdtAnyOfFromValue(v string) (*ReportAmountMdtAnyOf, error) {
	ev := ReportAmountMdtAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAmountMdtAnyOf: valid values are %v", v, AllowedReportAmountMdtAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAmountMdtAnyOf) IsValid() bool {
	for _, existing := range AllowedReportAmountMdtAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportAmountMdt_anyOf value
func (v ReportAmountMdtAnyOf) Ptr() *ReportAmountMdtAnyOf {
	return &v
}

type NullableReportAmountMdtAnyOf struct {
	value *ReportAmountMdtAnyOf
	isSet bool
}

func (v NullableReportAmountMdtAnyOf) Get() *ReportAmountMdtAnyOf {
	return v.value
}

func (v *NullableReportAmountMdtAnyOf) Set(val *ReportAmountMdtAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportAmountMdtAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportAmountMdtAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportAmountMdtAnyOf(val *ReportAmountMdtAnyOf) *NullableReportAmountMdtAnyOf {
	return &NullableReportAmountMdtAnyOf{value: val, isSet: true}
}

func (v NullableReportAmountMdtAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportAmountMdtAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

