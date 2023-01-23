/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// ReportAmountMdt The enumeration ReportAmountMdt defines Report Amount for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.10-1. 
type ReportAmountMdt struct {
	ReportAmountMdtAnyOf *ReportAmountMdtAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportAmountMdt) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ReportAmountMdtAnyOf
	err = json.Unmarshal(data, &dst.ReportAmountMdtAnyOf);
	if err == nil {
		jsonReportAmountMdtAnyOf, _ := json.Marshal(dst.ReportAmountMdtAnyOf)
		if string(jsonReportAmountMdtAnyOf) == "{}" { // empty struct
			dst.ReportAmountMdtAnyOf = nil
		} else {
			return nil // data stored in dst.ReportAmountMdtAnyOf, return on the first match
		}
	} else {
		dst.ReportAmountMdtAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReportAmountMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReportAmountMdt) MarshalJSON() ([]byte, error) {
	if src.ReportAmountMdtAnyOf != nil {
		return json.Marshal(&src.ReportAmountMdtAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReportAmountMdt struct {
	value *ReportAmountMdt
	isSet bool
}

func (v NullableReportAmountMdt) Get() *ReportAmountMdt {
	return v.value
}

func (v *NullableReportAmountMdt) Set(val *ReportAmountMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableReportAmountMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableReportAmountMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportAmountMdt(val *ReportAmountMdt) *NullableReportAmountMdt {
	return &NullableReportAmountMdt{value: val, isSet: true}
}

func (v NullableReportAmountMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportAmountMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


