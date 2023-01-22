/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// NrfInfoServedLmfInfoValue struct for NrfInfoServedLmfInfoValue
type NrfInfoServedLmfInfoValue struct {
	LmfInfo *LmfInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedLmfInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LmfInfo
	err = json.Unmarshal(data, &dst.LmfInfo);
	if err == nil {
		jsonLmfInfo, _ := json.Marshal(dst.LmfInfo)
		if string(jsonLmfInfo) == "{}" { // empty struct
			dst.LmfInfo = nil
		} else {
			return nil // data stored in dst.LmfInfo, return on the first match
		}
	} else {
		dst.LmfInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedLmfInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedLmfInfoValue) MarshalJSON() ([]byte, error) {
	if src.LmfInfo != nil {
		return json.Marshal(&src.LmfInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedLmfInfoValue struct {
	value *NrfInfoServedLmfInfoValue
	isSet bool
}

func (v NullableNrfInfoServedLmfInfoValue) Get() *NrfInfoServedLmfInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedLmfInfoValue) Set(val *NrfInfoServedLmfInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedLmfInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedLmfInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedLmfInfoValue(val *NrfInfoServedLmfInfoValue) *NullableNrfInfoServedLmfInfoValue {
	return &NullableNrfInfoServedLmfInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedLmfInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedLmfInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


