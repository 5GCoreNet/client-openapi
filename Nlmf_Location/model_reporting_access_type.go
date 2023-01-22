/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nlmf_Location

import (
	"encoding/json"
	"fmt"
)

// ReportingAccessType Specifies access types of event reporting.
type ReportingAccessType struct {
	ReportingAccessTypeAnyOf *ReportingAccessTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportingAccessType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ReportingAccessTypeAnyOf
	err = json.Unmarshal(data, &dst.ReportingAccessTypeAnyOf);
	if err == nil {
		jsonReportingAccessTypeAnyOf, _ := json.Marshal(dst.ReportingAccessTypeAnyOf)
		if string(jsonReportingAccessTypeAnyOf) == "{}" { // empty struct
			dst.ReportingAccessTypeAnyOf = nil
		} else {
			return nil // data stored in dst.ReportingAccessTypeAnyOf, return on the first match
		}
	} else {
		dst.ReportingAccessTypeAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ReportingAccessType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReportingAccessType) MarshalJSON() ([]byte, error) {
	if src.ReportingAccessTypeAnyOf != nil {
		return json.Marshal(&src.ReportingAccessTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReportingAccessType struct {
	value *ReportingAccessType
	isSet bool
}

func (v NullableReportingAccessType) Get() *ReportingAccessType {
	return v.value
}

func (v *NullableReportingAccessType) Set(val *ReportingAccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingAccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingAccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingAccessType(val *ReportingAccessType) *NullableReportingAccessType {
	return &NullableReportingAccessType{value: val, isSet: true}
}

func (v NullableReportingAccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingAccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


