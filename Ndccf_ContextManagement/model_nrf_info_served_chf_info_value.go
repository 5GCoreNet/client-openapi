/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// NrfInfoServedChfInfoValue struct for NrfInfoServedChfInfoValue
type NrfInfoServedChfInfoValue struct {
	ChfInfo *ChfInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedChfInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ChfInfo
	err = json.Unmarshal(data, &dst.ChfInfo);
	if err == nil {
		jsonChfInfo, _ := json.Marshal(dst.ChfInfo)
		if string(jsonChfInfo) == "{}" { // empty struct
			dst.ChfInfo = nil
		} else {
			return nil // data stored in dst.ChfInfo, return on the first match
		}
	} else {
		dst.ChfInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedChfInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedChfInfoValue) MarshalJSON() ([]byte, error) {
	if src.ChfInfo != nil {
		return json.Marshal(&src.ChfInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedChfInfoValue struct {
	value *NrfInfoServedChfInfoValue
	isSet bool
}

func (v NullableNrfInfoServedChfInfoValue) Get() *NrfInfoServedChfInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedChfInfoValue) Set(val *NrfInfoServedChfInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedChfInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedChfInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedChfInfoValue(val *NrfInfoServedChfInfoValue) *NullableNrfInfoServedChfInfoValue {
	return &NullableNrfInfoServedChfInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedChfInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedChfInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


