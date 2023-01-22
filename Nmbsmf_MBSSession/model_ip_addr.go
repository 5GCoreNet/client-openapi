/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsmf_MBSSession

import (
	"encoding/json"
	"fmt"
)

// IpAddr - Contains an IP adresse.
type IpAddr struct {
	Interface{} *interface{}
}

// interface{}AsIpAddr is a convenience function that returns interface{} wrapped in IpAddr
func Interface{}AsIpAddr(v *interface{}) IpAddr {
	return IpAddr{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IpAddr) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IpAddr)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IpAddr)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IpAddr) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IpAddr) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableIpAddr struct {
	value *IpAddr
	isSet bool
}

func (v NullableIpAddr) Get() *IpAddr {
	return v.value
}

func (v *NullableIpAddr) Set(val *IpAddr) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAddr) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAddr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAddr(val *IpAddr) *NullableIpAddr {
	return &NullableIpAddr{value: val, isSet: true}
}

func (v NullableIpAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAddr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


