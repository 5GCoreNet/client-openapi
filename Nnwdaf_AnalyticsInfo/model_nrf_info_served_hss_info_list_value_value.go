/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// NrfInfoServedHssInfoListValueValue struct for NrfInfoServedHssInfoListValueValue
type NrfInfoServedHssInfoListValueValue struct {
	HssInfo *HssInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedHssInfoListValueValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into HssInfo
	err = json.Unmarshal(data, &dst.HssInfo);
	if err == nil {
		jsonHssInfo, _ := json.Marshal(dst.HssInfo)
		if string(jsonHssInfo) == "{}" { // empty struct
			dst.HssInfo = nil
		} else {
			return nil // data stored in dst.HssInfo, return on the first match
		}
	} else {
		dst.HssInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedHssInfoListValueValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedHssInfoListValueValue) MarshalJSON() ([]byte, error) {
	if src.HssInfo != nil {
		return json.Marshal(&src.HssInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedHssInfoListValueValue struct {
	value *NrfInfoServedHssInfoListValueValue
	isSet bool
}

func (v NullableNrfInfoServedHssInfoListValueValue) Get() *NrfInfoServedHssInfoListValueValue {
	return v.value
}

func (v *NullableNrfInfoServedHssInfoListValueValue) Set(val *NrfInfoServedHssInfoListValueValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedHssInfoListValueValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedHssInfoListValueValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedHssInfoListValueValue(val *NrfInfoServedHssInfoListValueValue) *NullableNrfInfoServedHssInfoListValueValue {
	return &NullableNrfInfoServedHssInfoListValueValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedHssInfoListValueValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedHssInfoListValueValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


