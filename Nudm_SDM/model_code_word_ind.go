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

// CodeWordInd struct for CodeWordInd
type CodeWordInd struct {
	CodeWordIndAnyOf *CodeWordIndAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CodeWordInd) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CodeWordIndAnyOf
	err = json.Unmarshal(data, &dst.CodeWordIndAnyOf);
	if err == nil {
		jsonCodeWordIndAnyOf, _ := json.Marshal(dst.CodeWordIndAnyOf)
		if string(jsonCodeWordIndAnyOf) == "{}" { // empty struct
			dst.CodeWordIndAnyOf = nil
		} else {
			return nil // data stored in dst.CodeWordIndAnyOf, return on the first match
		}
	} else {
		dst.CodeWordIndAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CodeWordInd)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CodeWordInd) MarshalJSON() ([]byte, error) {
	if src.CodeWordIndAnyOf != nil {
		return json.Marshal(&src.CodeWordIndAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCodeWordInd struct {
	value *CodeWordInd
	isSet bool
}

func (v NullableCodeWordInd) Get() *CodeWordInd {
	return v.value
}

func (v *NullableCodeWordInd) Set(val *CodeWordInd) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeWordInd) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeWordInd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeWordInd(val *CodeWordInd) *NullableCodeWordInd {
	return &NullableCodeWordInd{value: val, isSet: true}
}

func (v NullableCodeWordInd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeWordInd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


