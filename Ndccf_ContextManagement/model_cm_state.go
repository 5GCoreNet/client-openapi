/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// CmState Describes the connection management state of a UE
type CmState struct {
	CmStateAnyOf *CmStateAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CmState) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CmStateAnyOf
	err = json.Unmarshal(data, &dst.CmStateAnyOf);
	if err == nil {
		jsonCmStateAnyOf, _ := json.Marshal(dst.CmStateAnyOf)
		if string(jsonCmStateAnyOf) == "{}" { // empty struct
			dst.CmStateAnyOf = nil
		} else {
			return nil // data stored in dst.CmStateAnyOf, return on the first match
		}
	} else {
		dst.CmStateAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CmState)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CmState) MarshalJSON() ([]byte, error) {
	if src.CmStateAnyOf != nil {
		return json.Marshal(&src.CmStateAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCmState struct {
	value *CmState
	isSet bool
}

func (v NullableCmState) Get() *CmState {
	return v.value
}

func (v *NullableCmState) Set(val *CmState) {
	v.value = val
	v.isSet = true
}

func (v NullableCmState) IsSet() bool {
	return v.isSet
}

func (v *NullableCmState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmState(val *CmState) *NullableCmState {
	return &NullableCmState{value: val, isSet: true}
}

func (v NullableCmState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


