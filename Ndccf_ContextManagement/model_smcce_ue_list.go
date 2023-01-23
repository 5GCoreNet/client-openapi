/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// SmcceUeList Represents the List of UEs classified based on experience level of Session Management  congestion control. 
type SmcceUeList struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SmcceUeList) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SmcceUeList)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SmcceUeList) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSmcceUeList struct {
	value *SmcceUeList
	isSet bool
}

func (v NullableSmcceUeList) Get() *SmcceUeList {
	return v.value
}

func (v *NullableSmcceUeList) Set(val *SmcceUeList) {
	v.value = val
	v.isSet = true
}

func (v NullableSmcceUeList) IsSet() bool {
	return v.isSet
}

func (v *NullableSmcceUeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmcceUeList(val *SmcceUeList) *NullableSmcceUeList {
	return &NullableSmcceUeList{value: val, isSet: true}
}

func (v NullableSmcceUeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmcceUeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


