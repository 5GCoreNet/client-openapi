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

// NrfInfoServedSmfInfoValue struct for NrfInfoServedSmfInfoValue
type NrfInfoServedSmfInfoValue struct {
	SmfInfo *SmfInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedSmfInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SmfInfo
	err = json.Unmarshal(data, &dst.SmfInfo);
	if err == nil {
		jsonSmfInfo, _ := json.Marshal(dst.SmfInfo)
		if string(jsonSmfInfo) == "{}" { // empty struct
			dst.SmfInfo = nil
		} else {
			return nil // data stored in dst.SmfInfo, return on the first match
		}
	} else {
		dst.SmfInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedSmfInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedSmfInfoValue) MarshalJSON() ([]byte, error) {
	if src.SmfInfo != nil {
		return json.Marshal(&src.SmfInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedSmfInfoValue struct {
	value *NrfInfoServedSmfInfoValue
	isSet bool
}

func (v NullableNrfInfoServedSmfInfoValue) Get() *NrfInfoServedSmfInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedSmfInfoValue) Set(val *NrfInfoServedSmfInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedSmfInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedSmfInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedSmfInfoValue(val *NrfInfoServedSmfInfoValue) *NullableNrfInfoServedSmfInfoValue {
	return &NullableNrfInfoServedSmfInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedSmfInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedSmfInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


