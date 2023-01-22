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

// NrfInfoServedPcscfInfoListValueValue struct for NrfInfoServedPcscfInfoListValueValue
type NrfInfoServedPcscfInfoListValueValue struct {
	PcscfInfo *PcscfInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedPcscfInfoListValueValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PcscfInfo
	err = json.Unmarshal(data, &dst.PcscfInfo);
	if err == nil {
		jsonPcscfInfo, _ := json.Marshal(dst.PcscfInfo)
		if string(jsonPcscfInfo) == "{}" { // empty struct
			dst.PcscfInfo = nil
		} else {
			return nil // data stored in dst.PcscfInfo, return on the first match
		}
	} else {
		dst.PcscfInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedPcscfInfoListValueValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedPcscfInfoListValueValue) MarshalJSON() ([]byte, error) {
	if src.PcscfInfo != nil {
		return json.Marshal(&src.PcscfInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedPcscfInfoListValueValue struct {
	value *NrfInfoServedPcscfInfoListValueValue
	isSet bool
}

func (v NullableNrfInfoServedPcscfInfoListValueValue) Get() *NrfInfoServedPcscfInfoListValueValue {
	return v.value
}

func (v *NullableNrfInfoServedPcscfInfoListValueValue) Set(val *NrfInfoServedPcscfInfoListValueValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedPcscfInfoListValueValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedPcscfInfoListValueValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedPcscfInfoListValueValue(val *NrfInfoServedPcscfInfoListValueValue) *NullableNrfInfoServedPcscfInfoListValueValue {
	return &NullableNrfInfoServedPcscfInfoListValueValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedPcscfInfoListValueValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedPcscfInfoListValueValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


