/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// EcRestrictionDataWb struct for EcRestrictionDataWb
type EcRestrictionDataWb struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EcRestrictionDataWb) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(EcRestrictionDataWb)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EcRestrictionDataWb) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEcRestrictionDataWb struct {
	value *EcRestrictionDataWb
	isSet bool
}

func (v NullableEcRestrictionDataWb) Get() *EcRestrictionDataWb {
	return v.value
}

func (v *NullableEcRestrictionDataWb) Set(val *EcRestrictionDataWb) {
	v.value = val
	v.isSet = true
}

func (v NullableEcRestrictionDataWb) IsSet() bool {
	return v.isSet
}

func (v *NullableEcRestrictionDataWb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcRestrictionDataWb(val *EcRestrictionDataWb) *NullableEcRestrictionDataWb {
	return &NullableEcRestrictionDataWb{value: val, isSet: true}
}

func (v NullableEcRestrictionDataWb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcRestrictionDataWb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


