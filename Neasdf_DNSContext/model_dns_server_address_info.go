/*
Neasdf_DNSContext

EASDF DNS Context Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Neasdf_DNSContext

import (
	"encoding/json"
	"fmt"
)

// DnsServerAddressInfo DNS Server Address Information
type DnsServerAddressInfo struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DnsServerAddressInfo) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DnsServerAddressInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DnsServerAddressInfo) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDnsServerAddressInfo struct {
	value *DnsServerAddressInfo
	isSet bool
}

func (v NullableDnsServerAddressInfo) Get() *DnsServerAddressInfo {
	return v.value
}

func (v *NullableDnsServerAddressInfo) Set(val *DnsServerAddressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsServerAddressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsServerAddressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsServerAddressInfo(val *DnsServerAddressInfo) *NullableDnsServerAddressInfo {
	return &NullableDnsServerAddressInfo{value: val, isSet: true}
}

func (v NullableDnsServerAddressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsServerAddressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


