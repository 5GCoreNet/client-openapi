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

// NrfInfoServedAanfInfoListValueValue struct for NrfInfoServedAanfInfoListValueValue
type NrfInfoServedAanfInfoListValueValue struct {
	AanfInfo *AanfInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedAanfInfoListValueValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AanfInfo
	err = json.Unmarshal(data, &dst.AanfInfo);
	if err == nil {
		jsonAanfInfo, _ := json.Marshal(dst.AanfInfo)
		if string(jsonAanfInfo) == "{}" { // empty struct
			dst.AanfInfo = nil
		} else {
			return nil // data stored in dst.AanfInfo, return on the first match
		}
	} else {
		dst.AanfInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedAanfInfoListValueValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedAanfInfoListValueValue) MarshalJSON() ([]byte, error) {
	if src.AanfInfo != nil {
		return json.Marshal(&src.AanfInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedAanfInfoListValueValue struct {
	value *NrfInfoServedAanfInfoListValueValue
	isSet bool
}

func (v NullableNrfInfoServedAanfInfoListValueValue) Get() *NrfInfoServedAanfInfoListValueValue {
	return v.value
}

func (v *NullableNrfInfoServedAanfInfoListValueValue) Set(val *NrfInfoServedAanfInfoListValueValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedAanfInfoListValueValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedAanfInfoListValueValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedAanfInfoListValueValue(val *NrfInfoServedAanfInfoListValueValue) *NullableNrfInfoServedAanfInfoListValueValue {
	return &NullableNrfInfoServedAanfInfoListValueValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedAanfInfoListValueValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedAanfInfoListValueValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


