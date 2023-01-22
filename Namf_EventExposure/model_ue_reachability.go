/*
Namf_EventExposure

AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// UeReachability Describes the reachability of the UE
type UeReachability struct {
	UeReachabilityAnyOf *UeReachabilityAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UeReachability) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UeReachabilityAnyOf
	err = json.Unmarshal(data, &dst.UeReachabilityAnyOf);
	if err == nil {
		jsonUeReachabilityAnyOf, _ := json.Marshal(dst.UeReachabilityAnyOf)
		if string(jsonUeReachabilityAnyOf) == "{}" { // empty struct
			dst.UeReachabilityAnyOf = nil
		} else {
			return nil // data stored in dst.UeReachabilityAnyOf, return on the first match
		}
	} else {
		dst.UeReachabilityAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(UeReachability)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UeReachability) MarshalJSON() ([]byte, error) {
	if src.UeReachabilityAnyOf != nil {
		return json.Marshal(&src.UeReachabilityAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUeReachability struct {
	value *UeReachability
	isSet bool
}

func (v NullableUeReachability) Get() *UeReachability {
	return v.value
}

func (v *NullableUeReachability) Set(val *UeReachability) {
	v.value = val
	v.isSet = true
}

func (v NullableUeReachability) IsSet() bool {
	return v.isSet
}

func (v *NullableUeReachability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeReachability(val *UeReachability) *NullableUeReachability {
	return &NullableUeReachability{value: val, isSet: true}
}

func (v NullableUeReachability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeReachability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


