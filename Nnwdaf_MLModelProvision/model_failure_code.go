/*
Nnwdaf_MLModelProvision

Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_MLModelProvision

import (
	"encoding/json"
	"fmt"
)

// FailureCode Possible values are - UNAVAILABLE_ML_MODEL: Indicates the requested ML model for the event is unavailable. 
type FailureCode struct {
	FailureCodeAnyOf *FailureCodeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FailureCode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FailureCodeAnyOf
	err = json.Unmarshal(data, &dst.FailureCodeAnyOf);
	if err == nil {
		jsonFailureCodeAnyOf, _ := json.Marshal(dst.FailureCodeAnyOf)
		if string(jsonFailureCodeAnyOf) == "{}" { // empty struct
			dst.FailureCodeAnyOf = nil
		} else {
			return nil // data stored in dst.FailureCodeAnyOf, return on the first match
		}
	} else {
		dst.FailureCodeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(FailureCode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FailureCode) MarshalJSON() ([]byte, error) {
	if src.FailureCodeAnyOf != nil {
		return json.Marshal(&src.FailureCodeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFailureCode struct {
	value *FailureCode
	isSet bool
}

func (v NullableFailureCode) Get() *FailureCode {
	return v.value
}

func (v *NullableFailureCode) Set(val *FailureCode) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureCode) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureCode(val *FailureCode) *NullableFailureCode {
	return &NullableFailureCode{value: val, isSet: true}
}

func (v NullableFailureCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


