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

// ImsVoPs struct for ImsVoPs
type ImsVoPs struct {
	ImsVoPsAnyOf *ImsVoPsAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ImsVoPs) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ImsVoPsAnyOf
	err = json.Unmarshal(data, &dst.ImsVoPsAnyOf);
	if err == nil {
		jsonImsVoPsAnyOf, _ := json.Marshal(dst.ImsVoPsAnyOf)
		if string(jsonImsVoPsAnyOf) == "{}" { // empty struct
			dst.ImsVoPsAnyOf = nil
		} else {
			return nil // data stored in dst.ImsVoPsAnyOf, return on the first match
		}
	} else {
		dst.ImsVoPsAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ImsVoPs)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ImsVoPs) MarshalJSON() ([]byte, error) {
	if src.ImsVoPsAnyOf != nil {
		return json.Marshal(&src.ImsVoPsAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableImsVoPs struct {
	value *ImsVoPs
	isSet bool
}

func (v NullableImsVoPs) Get() *ImsVoPs {
	return v.value
}

func (v *NullableImsVoPs) Set(val *ImsVoPs) {
	v.value = val
	v.isSet = true
}

func (v NullableImsVoPs) IsSet() bool {
	return v.isSet
}

func (v *NullableImsVoPs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImsVoPs(val *ImsVoPs) *NullableImsVoPs {
	return &NullableImsVoPs{value: val, isSet: true}
}

func (v NullableImsVoPs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImsVoPs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


