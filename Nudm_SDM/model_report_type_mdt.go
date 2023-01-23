/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// ReportTypeMdt The enumeration ReportTypeMdt defines Report Type for logged MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.4-1. 
type ReportTypeMdt struct {
	ReportTypeMdtAnyOf *ReportTypeMdtAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportTypeMdt) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ReportTypeMdtAnyOf
	err = json.Unmarshal(data, &dst.ReportTypeMdtAnyOf);
	if err == nil {
		jsonReportTypeMdtAnyOf, _ := json.Marshal(dst.ReportTypeMdtAnyOf)
		if string(jsonReportTypeMdtAnyOf) == "{}" { // empty struct
			dst.ReportTypeMdtAnyOf = nil
		} else {
			return nil // data stored in dst.ReportTypeMdtAnyOf, return on the first match
		}
	} else {
		dst.ReportTypeMdtAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReportTypeMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReportTypeMdt) MarshalJSON() ([]byte, error) {
	if src.ReportTypeMdtAnyOf != nil {
		return json.Marshal(&src.ReportTypeMdtAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReportTypeMdt struct {
	value *ReportTypeMdt
	isSet bool
}

func (v NullableReportTypeMdt) Get() *ReportTypeMdt {
	return v.value
}

func (v *NullableReportTypeMdt) Set(val *ReportTypeMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableReportTypeMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableReportTypeMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportTypeMdt(val *ReportTypeMdt) *NullableReportTypeMdt {
	return &NullableReportTypeMdt{value: val, isSet: true}
}

func (v NullableReportTypeMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportTypeMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


