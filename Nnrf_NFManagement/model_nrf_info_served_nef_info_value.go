/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFManagement

import (
	"encoding/json"
	"fmt"
)

// NrfInfoServedNefInfoValue struct for NrfInfoServedNefInfoValue
type NrfInfoServedNefInfoValue struct {
	NefInfo *NefInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedNefInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NefInfo
	err = json.Unmarshal(data, &dst.NefInfo);
	if err == nil {
		jsonNefInfo, _ := json.Marshal(dst.NefInfo)
		if string(jsonNefInfo) == "{}" { // empty struct
			dst.NefInfo = nil
		} else {
			return nil // data stored in dst.NefInfo, return on the first match
		}
	} else {
		dst.NefInfo = nil
	}

	// try to unmarshal JSON data into map[string]interface{}
	err = json.Unmarshal(data, &dst.map[string]interface{});
	if err == nil {
		jsonmap[string]interface{}, _ := json.Marshal(dst.map[string]interface{})
		if string(jsonmap[string]interface{}) == "{}" { // empty struct
			dst.map[string]interface{} = nil
		} else {
			return nil // data stored in dst.map[string]interface{}, return on the first match
		}
	} else {
		dst.map[string]interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedNefInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedNefInfoValue) MarshalJSON() ([]byte, error) {
	if src.NefInfo != nil {
		return json.Marshal(&src.NefInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedNefInfoValue struct {
	value *NrfInfoServedNefInfoValue
	isSet bool
}

func (v NullableNrfInfoServedNefInfoValue) Get() *NrfInfoServedNefInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedNefInfoValue) Set(val *NrfInfoServedNefInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedNefInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedNefInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedNefInfoValue(val *NrfInfoServedNefInfoValue) *NullableNrfInfoServedNefInfoValue {
	return &NullableNrfInfoServedNefInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedNefInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedNefInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


