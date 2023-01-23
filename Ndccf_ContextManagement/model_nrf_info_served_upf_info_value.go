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

// NrfInfoServedUpfInfoValue struct for NrfInfoServedUpfInfoValue
type NrfInfoServedUpfInfoValue struct {
	UpfInfo *UpfInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedUpfInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UpfInfo
	err = json.Unmarshal(data, &dst.UpfInfo);
	if err == nil {
		jsonUpfInfo, _ := json.Marshal(dst.UpfInfo)
		if string(jsonUpfInfo) == "{}" { // empty struct
			dst.UpfInfo = nil
		} else {
			return nil // data stored in dst.UpfInfo, return on the first match
		}
	} else {
		dst.UpfInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedUpfInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedUpfInfoValue) MarshalJSON() ([]byte, error) {
	if src.UpfInfo != nil {
		return json.Marshal(&src.UpfInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedUpfInfoValue struct {
	value *NrfInfoServedUpfInfoValue
	isSet bool
}

func (v NullableNrfInfoServedUpfInfoValue) Get() *NrfInfoServedUpfInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedUpfInfoValue) Set(val *NrfInfoServedUpfInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedUpfInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedUpfInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedUpfInfoValue(val *NrfInfoServedUpfInfoValue) *NullableNrfInfoServedUpfInfoValue {
	return &NullableNrfInfoServedUpfInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedUpfInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedUpfInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


