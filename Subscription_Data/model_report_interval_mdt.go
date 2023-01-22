/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// ReportIntervalMdt The enumeration ReportIntervalMdt defines Report Interval for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.9-1. 
type ReportIntervalMdt struct {
	ReportIntervalMdtAnyOf *ReportIntervalMdtAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportIntervalMdt) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ReportIntervalMdtAnyOf
	err = json.Unmarshal(data, &dst.ReportIntervalMdtAnyOf);
	if err == nil {
		jsonReportIntervalMdtAnyOf, _ := json.Marshal(dst.ReportIntervalMdtAnyOf)
		if string(jsonReportIntervalMdtAnyOf) == "{}" { // empty struct
			dst.ReportIntervalMdtAnyOf = nil
		} else {
			return nil // data stored in dst.ReportIntervalMdtAnyOf, return on the first match
		}
	} else {
		dst.ReportIntervalMdtAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReportIntervalMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReportIntervalMdt) MarshalJSON() ([]byte, error) {
	if src.ReportIntervalMdtAnyOf != nil {
		return json.Marshal(&src.ReportIntervalMdtAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReportIntervalMdt struct {
	value *ReportIntervalMdt
	isSet bool
}

func (v NullableReportIntervalMdt) Get() *ReportIntervalMdt {
	return v.value
}

func (v *NullableReportIntervalMdt) Set(val *ReportIntervalMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableReportIntervalMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableReportIntervalMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportIntervalMdt(val *ReportIntervalMdt) *NullableReportIntervalMdt {
	return &NullableReportIntervalMdt{value: val, isSet: true}
}

func (v NullableReportIntervalMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportIntervalMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


