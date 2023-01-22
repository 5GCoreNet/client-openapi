/*
Unified Data Repository Service API file for policy data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Policy_Data

import (
	"encoding/json"
	"fmt"
)

// SlicePolicyDataPatch - Contains the modified network slice specific policy control information.
type SlicePolicyDataPatch struct {
	Interface{} *interface{}
}

// interface{}AsSlicePolicyDataPatch is a convenience function that returns interface{} wrapped in SlicePolicyDataPatch
func Interface{}AsSlicePolicyDataPatch(v *interface{}) SlicePolicyDataPatch {
	return SlicePolicyDataPatch{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SlicePolicyDataPatch) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(SlicePolicyDataPatch)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SlicePolicyDataPatch)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SlicePolicyDataPatch) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SlicePolicyDataPatch) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableSlicePolicyDataPatch struct {
	value *SlicePolicyDataPatch
	isSet bool
}

func (v NullableSlicePolicyDataPatch) Get() *SlicePolicyDataPatch {
	return v.value
}

func (v *NullableSlicePolicyDataPatch) Set(val *SlicePolicyDataPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableSlicePolicyDataPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableSlicePolicyDataPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlicePolicyDataPatch(val *SlicePolicyDataPatch) *NullableSlicePolicyDataPatch {
	return &NullableSlicePolicyDataPatch{value: val, isSet: true}
}

func (v NullableSlicePolicyDataPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlicePolicyDataPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


