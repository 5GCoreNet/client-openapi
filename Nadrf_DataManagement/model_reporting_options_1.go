/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// ReportingOptions1 - Represents reporting options for processed notifications.
type ReportingOptions1 struct {
	Interface{} *interface{}
}

// interface{}AsReportingOptions1 is a convenience function that returns interface{} wrapped in ReportingOptions1
func Interface{}AsReportingOptions1(v *interface{}) ReportingOptions1 {
	return ReportingOptions1{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ReportingOptions1) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ReportingOptions1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ReportingOptions1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ReportingOptions1) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ReportingOptions1) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableReportingOptions1 struct {
	value *ReportingOptions1
	isSet bool
}

func (v NullableReportingOptions1) Get() *ReportingOptions1 {
	return v.value
}

func (v *NullableReportingOptions1) Set(val *ReportingOptions1) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingOptions1) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingOptions1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingOptions1(val *ReportingOptions1) *NullableReportingOptions1 {
	return &NullableReportingOptions1{value: val, isSet: true}
}

func (v NullableReportingOptions1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingOptions1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


