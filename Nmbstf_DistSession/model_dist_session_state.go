/*
Nmbstf-distsession

MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbstf_DistSession

import (
	"encoding/json"
	"fmt"
)

// DistSessionState Current State of MBS distribution session
type DistSessionState struct {
	DistSessionStateAnyOf *DistSessionStateAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DistSessionState) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DistSessionStateAnyOf
	err = json.Unmarshal(data, &dst.DistSessionStateAnyOf);
	if err == nil {
		jsonDistSessionStateAnyOf, _ := json.Marshal(dst.DistSessionStateAnyOf)
		if string(jsonDistSessionStateAnyOf) == "{}" { // empty struct
			dst.DistSessionStateAnyOf = nil
		} else {
			return nil // data stored in dst.DistSessionStateAnyOf, return on the first match
		}
	} else {
		dst.DistSessionStateAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DistSessionState)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DistSessionState) MarshalJSON() ([]byte, error) {
	if src.DistSessionStateAnyOf != nil {
		return json.Marshal(&src.DistSessionStateAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDistSessionState struct {
	value *DistSessionState
	isSet bool
}

func (v NullableDistSessionState) Get() *DistSessionState {
	return v.value
}

func (v *NullableDistSessionState) Set(val *DistSessionState) {
	v.value = val
	v.isSet = true
}

func (v NullableDistSessionState) IsSet() bool {
	return v.isSet
}

func (v *NullableDistSessionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistSessionState(val *DistSessionState) *NullableDistSessionState {
	return &NullableDistSessionState{value: val, isSet: true}
}

func (v NullableDistSessionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistSessionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


