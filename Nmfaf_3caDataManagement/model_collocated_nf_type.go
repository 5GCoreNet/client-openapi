/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// CollocatedNfType NF types for a collocated NF
type CollocatedNfType struct {
	CollocatedNfTypeAnyOf *CollocatedNfTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CollocatedNfType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CollocatedNfTypeAnyOf
	err = json.Unmarshal(data, &dst.CollocatedNfTypeAnyOf);
	if err == nil {
		jsonCollocatedNfTypeAnyOf, _ := json.Marshal(dst.CollocatedNfTypeAnyOf)
		if string(jsonCollocatedNfTypeAnyOf) == "{}" { // empty struct
			dst.CollocatedNfTypeAnyOf = nil
		} else {
			return nil // data stored in dst.CollocatedNfTypeAnyOf, return on the first match
		}
	} else {
		dst.CollocatedNfTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CollocatedNfType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CollocatedNfType) MarshalJSON() ([]byte, error) {
	if src.CollocatedNfTypeAnyOf != nil {
		return json.Marshal(&src.CollocatedNfTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCollocatedNfType struct {
	value *CollocatedNfType
	isSet bool
}

func (v NullableCollocatedNfType) Get() *CollocatedNfType {
	return v.value
}

func (v *NullableCollocatedNfType) Set(val *CollocatedNfType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollocatedNfType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollocatedNfType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollocatedNfType(val *CollocatedNfType) *NullableCollocatedNfType {
	return &NullableCollocatedNfType{value: val, isSet: true}
}

func (v NullableCollocatedNfType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollocatedNfType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


