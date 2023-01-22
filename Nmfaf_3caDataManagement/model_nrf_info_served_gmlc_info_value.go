/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// NrfInfoServedGmlcInfoValue struct for NrfInfoServedGmlcInfoValue
type NrfInfoServedGmlcInfoValue struct {
	GmlcInfo *GmlcInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedGmlcInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into GmlcInfo
	err = json.Unmarshal(data, &dst.GmlcInfo);
	if err == nil {
		jsonGmlcInfo, _ := json.Marshal(dst.GmlcInfo)
		if string(jsonGmlcInfo) == "{}" { // empty struct
			dst.GmlcInfo = nil
		} else {
			return nil // data stored in dst.GmlcInfo, return on the first match
		}
	} else {
		dst.GmlcInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedGmlcInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedGmlcInfoValue) MarshalJSON() ([]byte, error) {
	if src.GmlcInfo != nil {
		return json.Marshal(&src.GmlcInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedGmlcInfoValue struct {
	value *NrfInfoServedGmlcInfoValue
	isSet bool
}

func (v NullableNrfInfoServedGmlcInfoValue) Get() *NrfInfoServedGmlcInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedGmlcInfoValue) Set(val *NrfInfoServedGmlcInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedGmlcInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedGmlcInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedGmlcInfoValue(val *NrfInfoServedGmlcInfoValue) *NullableNrfInfoServedGmlcInfoValue {
	return &NullableNrfInfoServedGmlcInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedGmlcInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedGmlcInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


