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

// UeReachableInd UE Reachable Indication
type UeReachableInd struct {
	UeReachableIndAnyOf *UeReachableIndAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UeReachableInd) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UeReachableIndAnyOf
	err = json.Unmarshal(data, &dst.UeReachableIndAnyOf);
	if err == nil {
		jsonUeReachableIndAnyOf, _ := json.Marshal(dst.UeReachableIndAnyOf)
		if string(jsonUeReachableIndAnyOf) == "{}" { // empty struct
			dst.UeReachableIndAnyOf = nil
		} else {
			return nil // data stored in dst.UeReachableIndAnyOf, return on the first match
		}
	} else {
		dst.UeReachableIndAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(UeReachableInd)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UeReachableInd) MarshalJSON() ([]byte, error) {
	if src.UeReachableIndAnyOf != nil {
		return json.Marshal(&src.UeReachableIndAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUeReachableInd struct {
	value *UeReachableInd
	isSet bool
}

func (v NullableUeReachableInd) Get() *UeReachableInd {
	return v.value
}

func (v *NullableUeReachableInd) Set(val *UeReachableInd) {
	v.value = val
	v.isSet = true
}

func (v NullableUeReachableInd) IsSet() bool {
	return v.isSet
}

func (v *NullableUeReachableInd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeReachableInd(val *UeReachableInd) *NullableUeReachableInd {
	return &NullableUeReachableInd{value: val, isSet: true}
}

func (v NullableUeReachableInd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeReachableInd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


