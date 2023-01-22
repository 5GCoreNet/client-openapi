/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// UpCnxState User Plane Connection State. Possible values are - ACTIVATED - DEACTIVATED - ACTIVATING - SUSPENDED 
type UpCnxState struct {
	UpCnxStateAnyOf *UpCnxStateAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UpCnxState) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UpCnxStateAnyOf
	err = json.Unmarshal(data, &dst.UpCnxStateAnyOf);
	if err == nil {
		jsonUpCnxStateAnyOf, _ := json.Marshal(dst.UpCnxStateAnyOf)
		if string(jsonUpCnxStateAnyOf) == "{}" { // empty struct
			dst.UpCnxStateAnyOf = nil
		} else {
			return nil // data stored in dst.UpCnxStateAnyOf, return on the first match
		}
	} else {
		dst.UpCnxStateAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(UpCnxState)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UpCnxState) MarshalJSON() ([]byte, error) {
	if src.UpCnxStateAnyOf != nil {
		return json.Marshal(&src.UpCnxStateAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUpCnxState struct {
	value *UpCnxState
	isSet bool
}

func (v NullableUpCnxState) Get() *UpCnxState {
	return v.value
}

func (v *NullableUpCnxState) Set(val *UpCnxState) {
	v.value = val
	v.isSet = true
}

func (v NullableUpCnxState) IsSet() bool {
	return v.isSet
}

func (v *NullableUpCnxState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpCnxState(val *UpCnxState) *NullableUpCnxState {
	return &NullableUpCnxState{value: val, isSet: true}
}

func (v NullableUpCnxState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpCnxState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


