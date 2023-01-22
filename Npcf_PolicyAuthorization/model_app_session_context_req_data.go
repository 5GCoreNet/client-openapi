/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// AppSessionContextReqData - Identifies the service requirements of an Individual Application Session Context.
type AppSessionContextReqData struct {
	Interface{} *interface{}
}

// interface{}AsAppSessionContextReqData is a convenience function that returns interface{} wrapped in AppSessionContextReqData
func Interface{}AsAppSessionContextReqData(v *interface{}) AppSessionContextReqData {
	return AppSessionContextReqData{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppSessionContextReqData) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(AppSessionContextReqData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppSessionContextReqData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppSessionContextReqData) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppSessionContextReqData) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableAppSessionContextReqData struct {
	value *AppSessionContextReqData
	isSet bool
}

func (v NullableAppSessionContextReqData) Get() *AppSessionContextReqData {
	return v.value
}

func (v *NullableAppSessionContextReqData) Set(val *AppSessionContextReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSessionContextReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSessionContextReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSessionContextReqData(val *AppSessionContextReqData) *NullableAppSessionContextReqData {
	return &NullableAppSessionContextReqData{value: val, isSet: true}
}

func (v NullableAppSessionContextReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSessionContextReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


