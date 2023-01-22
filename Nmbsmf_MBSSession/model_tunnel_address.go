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

// TunnelAddress Tunnel address
type TunnelAddress struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TunnelAddress) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(TunnelAddress)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TunnelAddress) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTunnelAddress struct {
	value *TunnelAddress
	isSet bool
}

func (v NullableTunnelAddress) Get() *TunnelAddress {
	return v.value
}

func (v *NullableTunnelAddress) Set(val *TunnelAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelAddress(val *TunnelAddress) *NullableTunnelAddress {
	return &NullableTunnelAddress{value: val, isSet: true}
}

func (v NullableTunnelAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


