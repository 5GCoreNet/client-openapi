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

// NfStatus Contains the percentage of time spent on various NF states.
type NfStatus struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NfStatus) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(NfStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NfStatus) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNfStatus struct {
	value *NfStatus
	isSet bool
}

func (v NullableNfStatus) Get() *NfStatus {
	return v.value
}

func (v *NullableNfStatus) Set(val *NfStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNfStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNfStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfStatus(val *NfStatus) *NullableNfStatus {
	return &NullableNfStatus{value: val, isSet: true}
}

func (v NullableNfStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


