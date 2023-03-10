/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// ReportingAreaType Indicates type of event reporting area.
type ReportingAreaType struct {
	ReportingAreaTypeAnyOf *ReportingAreaTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportingAreaType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ReportingAreaTypeAnyOf
	err = json.Unmarshal(data, &dst.ReportingAreaTypeAnyOf);
	if err == nil {
		jsonReportingAreaTypeAnyOf, _ := json.Marshal(dst.ReportingAreaTypeAnyOf)
		if string(jsonReportingAreaTypeAnyOf) == "{}" { // empty struct
			dst.ReportingAreaTypeAnyOf = nil
		} else {
			return nil // data stored in dst.ReportingAreaTypeAnyOf, return on the first match
		}
	} else {
		dst.ReportingAreaTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReportingAreaType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReportingAreaType) MarshalJSON() ([]byte, error) {
	if src.ReportingAreaTypeAnyOf != nil {
		return json.Marshal(&src.ReportingAreaTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReportingAreaType struct {
	value *ReportingAreaType
	isSet bool
}

func (v NullableReportingAreaType) Get() *ReportingAreaType {
	return v.value
}

func (v *NullableReportingAreaType) Set(val *ReportingAreaType) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingAreaType) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingAreaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingAreaType(val *ReportingAreaType) *NullableReportingAreaType {
	return &NullableReportingAreaType{value: val, isSet: true}
}

func (v NullableReportingAreaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingAreaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


