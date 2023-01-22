/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// ChargingInfo Diameter addresses of the charging function
type ChargingInfo struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ChargingInfo) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ChargingInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ChargingInfo) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableChargingInfo struct {
	value *ChargingInfo
	isSet bool
}

func (v NullableChargingInfo) Get() *ChargingInfo {
	return v.value
}

func (v *NullableChargingInfo) Set(val *ChargingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableChargingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableChargingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargingInfo(val *ChargingInfo) *NullableChargingInfo {
	return &NullableChargingInfo{value: val, isSet: true}
}

func (v NullableChargingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


