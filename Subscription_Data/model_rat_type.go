/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// RatType Indicates the radio access used.
type RatType struct {
	RatTypeAnyOf *RatTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RatType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RatTypeAnyOf
	err = json.Unmarshal(data, &dst.RatTypeAnyOf);
	if err == nil {
		jsonRatTypeAnyOf, _ := json.Marshal(dst.RatTypeAnyOf)
		if string(jsonRatTypeAnyOf) == "{}" { // empty struct
			dst.RatTypeAnyOf = nil
		} else {
			return nil // data stored in dst.RatTypeAnyOf, return on the first match
		}
	} else {
		dst.RatTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RatType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RatType) MarshalJSON() ([]byte, error) {
	if src.RatTypeAnyOf != nil {
		return json.Marshal(&src.RatTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRatType struct {
	value *RatType
	isSet bool
}

func (v NullableRatType) Get() *RatType {
	return v.value
}

func (v *NullableRatType) Set(val *RatType) {
	v.value = val
	v.isSet = true
}

func (v NullableRatType) IsSet() bool {
	return v.isSet
}

func (v *NullableRatType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatType(val *RatType) *NullableRatType {
	return &NullableRatType{value: val, isSet: true}
}

func (v NullableRatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


