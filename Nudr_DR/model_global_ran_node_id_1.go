/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// GlobalRanNodeId1 - One of the six attributes n3IwfId, gNbIdm, ngeNbId, wagfId, tngfId, eNbId shall be present. 
type GlobalRanNodeId1 struct {
	Interface{} *interface{}
}

// interface{}AsGlobalRanNodeId1 is a convenience function that returns interface{} wrapped in GlobalRanNodeId1
func Interface{}AsGlobalRanNodeId1(v *interface{}) GlobalRanNodeId1 {
	return GlobalRanNodeId1{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GlobalRanNodeId1) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(GlobalRanNodeId1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GlobalRanNodeId1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GlobalRanNodeId1) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GlobalRanNodeId1) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableGlobalRanNodeId1 struct {
	value *GlobalRanNodeId1
	isSet bool
}

func (v NullableGlobalRanNodeId1) Get() *GlobalRanNodeId1 {
	return v.value
}

func (v *NullableGlobalRanNodeId1) Set(val *GlobalRanNodeId1) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalRanNodeId1) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalRanNodeId1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalRanNodeId1(val *GlobalRanNodeId1) *NullableGlobalRanNodeId1 {
	return &NullableGlobalRanNodeId1{value: val, isSet: true}
}

func (v NullableGlobalRanNodeId1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalRanNodeId1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


