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

// SuccessType Success Type to indicate full or partial success
type SuccessType struct {
	SuccessTypeAnyOf *SuccessTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SuccessType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SuccessTypeAnyOf
	err = json.Unmarshal(data, &dst.SuccessTypeAnyOf);
	if err == nil {
		jsonSuccessTypeAnyOf, _ := json.Marshal(dst.SuccessTypeAnyOf)
		if string(jsonSuccessTypeAnyOf) == "{}" { // empty struct
			dst.SuccessTypeAnyOf = nil
		} else {
			return nil // data stored in dst.SuccessTypeAnyOf, return on the first match
		}
	} else {
		dst.SuccessTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SuccessType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SuccessType) MarshalJSON() ([]byte, error) {
	if src.SuccessTypeAnyOf != nil {
		return json.Marshal(&src.SuccessTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSuccessType struct {
	value *SuccessType
	isSet bool
}

func (v NullableSuccessType) Get() *SuccessType {
	return v.value
}

func (v *NullableSuccessType) Set(val *SuccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessType(val *SuccessType) *NullableSuccessType {
	return &NullableSuccessType{value: val, isSet: true}
}

func (v NullableSuccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


