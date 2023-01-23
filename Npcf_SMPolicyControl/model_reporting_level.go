/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// ReportingLevel Possible values are - SER_ID_LEVEL: Indicates that the usage shall be reported on service id and rating group combination level. - RAT_GR_LEVEL: Indicates that the usage shall be reported on rating group level. - SPON_CON_LEVEL: Indicates that the usage shall be reported on sponsor identity and rating group combination level. 
type ReportingLevel struct {
	NullValue *NullValue
	ReportingLevelAnyOf *ReportingLevelAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportingLevel) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	// try to unmarshal JSON data into ReportingLevelAnyOf
	err = json.Unmarshal(data, &dst.ReportingLevelAnyOf);
	if err == nil {
		jsonReportingLevelAnyOf, _ := json.Marshal(dst.ReportingLevelAnyOf)
		if string(jsonReportingLevelAnyOf) == "{}" { // empty struct
			dst.ReportingLevelAnyOf = nil
		} else {
			return nil // data stored in dst.ReportingLevelAnyOf, return on the first match
		}
	} else {
		dst.ReportingLevelAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReportingLevel)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReportingLevel) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.ReportingLevelAnyOf != nil {
		return json.Marshal(&src.ReportingLevelAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReportingLevel struct {
	value *ReportingLevel
	isSet bool
}

func (v NullableReportingLevel) Get() *ReportingLevel {
	return v.value
}

func (v *NullableReportingLevel) Set(val *ReportingLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingLevel(val *ReportingLevel) *NullableReportingLevel {
	return &NullableReportingLevel{value: val, isSet: true}
}

func (v NullableReportingLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


