/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// DispersionCollection1 - Contains the dispersion information collected for an AF.
type DispersionCollection1 struct {
	Interface{} *interface{}
}

// interface{}AsDispersionCollection1 is a convenience function that returns interface{} wrapped in DispersionCollection1
func Interface{}AsDispersionCollection1(v *interface{}) DispersionCollection1 {
	return DispersionCollection1{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DispersionCollection1) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DispersionCollection1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DispersionCollection1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DispersionCollection1) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DispersionCollection1) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableDispersionCollection1 struct {
	value *DispersionCollection1
	isSet bool
}

func (v NullableDispersionCollection1) Get() *DispersionCollection1 {
	return v.value
}

func (v *NullableDispersionCollection1) Set(val *DispersionCollection1) {
	v.value = val
	v.isSet = true
}

func (v NullableDispersionCollection1) IsSet() bool {
	return v.isSet
}

func (v *NullableDispersionCollection1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispersionCollection1(val *DispersionCollection1) *NullableDispersionCollection1 {
	return &NullableDispersionCollection1{value: val, isSet: true}
}

func (v NullableDispersionCollection1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispersionCollection1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


