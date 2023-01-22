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

// ReportIntervalMdtAnyOf the model 'ReportIntervalMdtAnyOf'
type ReportIntervalMdtAnyOf string

// List of ReportIntervalMdt_anyOf
const (
	_120 ReportIntervalMdtAnyOf = "120"
	_240 ReportIntervalMdtAnyOf = "240"
	_480 ReportIntervalMdtAnyOf = "480"
	_640 ReportIntervalMdtAnyOf = "640"
	_1024 ReportIntervalMdtAnyOf = "1024"
	_2048 ReportIntervalMdtAnyOf = "2048"
	_5120 ReportIntervalMdtAnyOf = "5120"
	_10240 ReportIntervalMdtAnyOf = "10240"
	_60000 ReportIntervalMdtAnyOf = "60000"
	_360000 ReportIntervalMdtAnyOf = "360000"
	_720000 ReportIntervalMdtAnyOf = "720000"
	_1800000 ReportIntervalMdtAnyOf = "1800000"
	_3600000 ReportIntervalMdtAnyOf = "3600000"
)

// All allowed values of ReportIntervalMdtAnyOf enum
var AllowedReportIntervalMdtAnyOfEnumValues = []ReportIntervalMdtAnyOf{
	"120",
	"240",
	"480",
	"640",
	"1024",
	"2048",
	"5120",
	"10240",
	"60000",
	"360000",
	"720000",
	"1800000",
	"3600000",
}

func (v *ReportIntervalMdtAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportIntervalMdtAnyOf(value)
	for _, existing := range AllowedReportIntervalMdtAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportIntervalMdtAnyOf", value)
}

// NewReportIntervalMdtAnyOfFromValue returns a pointer to a valid ReportIntervalMdtAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportIntervalMdtAnyOfFromValue(v string) (*ReportIntervalMdtAnyOf, error) {
	ev := ReportIntervalMdtAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportIntervalMdtAnyOf: valid values are %v", v, AllowedReportIntervalMdtAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportIntervalMdtAnyOf) IsValid() bool {
	for _, existing := range AllowedReportIntervalMdtAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportIntervalMdt_anyOf value
func (v ReportIntervalMdtAnyOf) Ptr() *ReportIntervalMdtAnyOf {
	return &v
}

type NullableReportIntervalMdtAnyOf struct {
	value *ReportIntervalMdtAnyOf
	isSet bool
}

func (v NullableReportIntervalMdtAnyOf) Get() *ReportIntervalMdtAnyOf {
	return v.value
}

func (v *NullableReportIntervalMdtAnyOf) Set(val *ReportIntervalMdtAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportIntervalMdtAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportIntervalMdtAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportIntervalMdtAnyOf(val *ReportIntervalMdtAnyOf) *NullableReportIntervalMdtAnyOf {
	return &NullableReportIntervalMdtAnyOf{value: val, isSet: true}
}

func (v NullableReportIntervalMdtAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportIntervalMdtAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

