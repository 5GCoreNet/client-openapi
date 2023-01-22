/*
Nmbstf-distsession

MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbstf_DistSession

import (
	"encoding/json"
	"fmt"
)

// DistSessionEventType Status Event Type
type DistSessionEventType struct {
	DistSessionEventTypeAnyOf *DistSessionEventTypeAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DistSessionEventType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DistSessionEventTypeAnyOf
	err = json.Unmarshal(data, &dst.DistSessionEventTypeAnyOf);
	if err == nil {
		jsonDistSessionEventTypeAnyOf, _ := json.Marshal(dst.DistSessionEventTypeAnyOf)
		if string(jsonDistSessionEventTypeAnyOf) == "{}" { // empty struct
			dst.DistSessionEventTypeAnyOf = nil
		} else {
			return nil // data stored in dst.DistSessionEventTypeAnyOf, return on the first match
		}
	} else {
		dst.DistSessionEventTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DistSessionEventType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DistSessionEventType) MarshalJSON() ([]byte, error) {
	if src.DistSessionEventTypeAnyOf != nil {
		return json.Marshal(&src.DistSessionEventTypeAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDistSessionEventType struct {
	value *DistSessionEventType
	isSet bool
}

func (v NullableDistSessionEventType) Get() *DistSessionEventType {
	return v.value
}

func (v *NullableDistSessionEventType) Set(val *DistSessionEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableDistSessionEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableDistSessionEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistSessionEventType(val *DistSessionEventType) *NullableDistSessionEventType {
	return &NullableDistSessionEventType{value: val, isSet: true}
}

func (v NullableDistSessionEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistSessionEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

