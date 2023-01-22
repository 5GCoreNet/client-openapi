/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndcaf_DataReporting

import (
	"encoding/json"
	"fmt"
)

// ReportingConditionTypeAnyOf the model 'ReportingConditionTypeAnyOf'
type ReportingConditionTypeAnyOf string

// List of ReportingConditionType_anyOf
const (
	INTERVAL ReportingConditionTypeAnyOf = "INTERVAL"
	THRESHOLD ReportingConditionTypeAnyOf = "THRESHOLD"
	EVENT ReportingConditionTypeAnyOf = "EVENT"
)

// All allowed values of ReportingConditionTypeAnyOf enum
var AllowedReportingConditionTypeAnyOfEnumValues = []ReportingConditionTypeAnyOf{
	"INTERVAL",
	"THRESHOLD",
	"EVENT",
}

func (v *ReportingConditionTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportingConditionTypeAnyOf(value)
	for _, existing := range AllowedReportingConditionTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportingConditionTypeAnyOf", value)
}

// NewReportingConditionTypeAnyOfFromValue returns a pointer to a valid ReportingConditionTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportingConditionTypeAnyOfFromValue(v string) (*ReportingConditionTypeAnyOf, error) {
	ev := ReportingConditionTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportingConditionTypeAnyOf: valid values are %v", v, AllowedReportingConditionTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportingConditionTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedReportingConditionTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportingConditionType_anyOf value
func (v ReportingConditionTypeAnyOf) Ptr() *ReportingConditionTypeAnyOf {
	return &v
}

type NullableReportingConditionTypeAnyOf struct {
	value *ReportingConditionTypeAnyOf
	isSet bool
}

func (v NullableReportingConditionTypeAnyOf) Get() *ReportingConditionTypeAnyOf {
	return v.value
}

func (v *NullableReportingConditionTypeAnyOf) Set(val *ReportingConditionTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingConditionTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingConditionTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingConditionTypeAnyOf(val *ReportingConditionTypeAnyOf) *NullableReportingConditionTypeAnyOf {
	return &NullableReportingConditionTypeAnyOf{value: val, isSet: true}
}

func (v NullableReportingConditionTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingConditionTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

