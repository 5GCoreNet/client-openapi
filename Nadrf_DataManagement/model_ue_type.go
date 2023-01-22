/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// UeType Describes the type of UEs
type UeType struct {
	UeTypeAnyOf *UeTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UeType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UeTypeAnyOf
	err = json.Unmarshal(data, &dst.UeTypeAnyOf);
	if err == nil {
		jsonUeTypeAnyOf, _ := json.Marshal(dst.UeTypeAnyOf)
		if string(jsonUeTypeAnyOf) == "{}" { // empty struct
			dst.UeTypeAnyOf = nil
		} else {
			return nil // data stored in dst.UeTypeAnyOf, return on the first match
		}
	} else {
		dst.UeTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(UeType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UeType) MarshalJSON() ([]byte, error) {
	if src.UeTypeAnyOf != nil {
		return json.Marshal(&src.UeTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUeType struct {
	value *UeType
	isSet bool
}

func (v NullableUeType) Get() *UeType {
	return v.value
}

func (v *NullableUeType) Set(val *UeType) {
	v.value = val
	v.isSet = true
}

func (v NullableUeType) IsSet() bool {
	return v.isSet
}

func (v *NullableUeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeType(val *UeType) *NullableUeType {
	return &NullableUeType{value: val, isSet: true}
}

func (v NullableUeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


