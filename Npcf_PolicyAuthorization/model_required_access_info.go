/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// RequiredAccessInfo Indicates the access network information required for an AF session.
type RequiredAccessInfo struct {
	RequiredAccessInfoAnyOf *RequiredAccessInfoAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RequiredAccessInfo) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RequiredAccessInfoAnyOf
	err = json.Unmarshal(data, &dst.RequiredAccessInfoAnyOf);
	if err == nil {
		jsonRequiredAccessInfoAnyOf, _ := json.Marshal(dst.RequiredAccessInfoAnyOf)
		if string(jsonRequiredAccessInfoAnyOf) == "{}" { // empty struct
			dst.RequiredAccessInfoAnyOf = nil
		} else {
			return nil // data stored in dst.RequiredAccessInfoAnyOf, return on the first match
		}
	} else {
		dst.RequiredAccessInfoAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RequiredAccessInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RequiredAccessInfo) MarshalJSON() ([]byte, error) {
	if src.RequiredAccessInfoAnyOf != nil {
		return json.Marshal(&src.RequiredAccessInfoAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRequiredAccessInfo struct {
	value *RequiredAccessInfo
	isSet bool
}

func (v NullableRequiredAccessInfo) Get() *RequiredAccessInfo {
	return v.value
}

func (v *NullableRequiredAccessInfo) Set(val *RequiredAccessInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRequiredAccessInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRequiredAccessInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequiredAccessInfo(val *RequiredAccessInfo) *NullableRequiredAccessInfo {
	return &NullableRequiredAccessInfo{value: val, isSet: true}
}

func (v NullableRequiredAccessInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequiredAccessInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


