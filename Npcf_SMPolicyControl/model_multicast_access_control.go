/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// MulticastAccessControl Indicates whether the service data flow, corresponding to the service data flow template, is allowed or not allowed.
type MulticastAccessControl struct {
	MulticastAccessControlAnyOf *MulticastAccessControlAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MulticastAccessControl) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MulticastAccessControlAnyOf
	err = json.Unmarshal(data, &dst.MulticastAccessControlAnyOf);
	if err == nil {
		jsonMulticastAccessControlAnyOf, _ := json.Marshal(dst.MulticastAccessControlAnyOf)
		if string(jsonMulticastAccessControlAnyOf) == "{}" { // empty struct
			dst.MulticastAccessControlAnyOf = nil
		} else {
			return nil // data stored in dst.MulticastAccessControlAnyOf, return on the first match
		}
	} else {
		dst.MulticastAccessControlAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MulticastAccessControl)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MulticastAccessControl) MarshalJSON() ([]byte, error) {
	if src.MulticastAccessControlAnyOf != nil {
		return json.Marshal(&src.MulticastAccessControlAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMulticastAccessControl struct {
	value *MulticastAccessControl
	isSet bool
}

func (v NullableMulticastAccessControl) Get() *MulticastAccessControl {
	return v.value
}

func (v *NullableMulticastAccessControl) Set(val *MulticastAccessControl) {
	v.value = val
	v.isSet = true
}

func (v NullableMulticastAccessControl) IsSet() bool {
	return v.isSet
}

func (v *NullableMulticastAccessControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMulticastAccessControl(val *MulticastAccessControl) *NullableMulticastAccessControl {
	return &NullableMulticastAccessControl{value: val, isSet: true}
}

func (v NullableMulticastAccessControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMulticastAccessControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


