/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
	"fmt"
)

// NrfInfoServedNwdafInfoValue struct for NrfInfoServedNwdafInfoValue
type NrfInfoServedNwdafInfoValue struct {
	NwdafInfo *NwdafInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedNwdafInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NwdafInfo
	err = json.Unmarshal(data, &dst.NwdafInfo);
	if err == nil {
		jsonNwdafInfo, _ := json.Marshal(dst.NwdafInfo)
		if string(jsonNwdafInfo) == "{}" { // empty struct
			dst.NwdafInfo = nil
		} else {
			return nil // data stored in dst.NwdafInfo, return on the first match
		}
	} else {
		dst.NwdafInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedNwdafInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedNwdafInfoValue) MarshalJSON() ([]byte, error) {
	if src.NwdafInfo != nil {
		return json.Marshal(&src.NwdafInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedNwdafInfoValue struct {
	value *NrfInfoServedNwdafInfoValue
	isSet bool
}

func (v NullableNrfInfoServedNwdafInfoValue) Get() *NrfInfoServedNwdafInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedNwdafInfoValue) Set(val *NrfInfoServedNwdafInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedNwdafInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedNwdafInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedNwdafInfoValue(val *NrfInfoServedNwdafInfoValue) *NullableNrfInfoServedNwdafInfoValue {
	return &NullableNrfInfoServedNwdafInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedNwdafInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedNwdafInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


