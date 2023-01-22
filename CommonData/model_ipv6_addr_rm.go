/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CommonData

import (
	"encoding/json"
)

// Ipv6AddrRm String identifying an IPv6 address formatted according to clause 4 of RFC5952 with the OpenAPI 'nullable: true' property. The mixed IPv4 IPv6 notation according to clause 5 of RFC5952 shall not be used. 
type Ipv6AddrRm struct {
}

// NewIpv6AddrRm instantiates a new Ipv6AddrRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6AddrRm() *Ipv6AddrRm {
	this := Ipv6AddrRm{}
	return &this
}

// NewIpv6AddrRmWithDefaults instantiates a new Ipv6AddrRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6AddrRmWithDefaults() *Ipv6AddrRm {
	this := Ipv6AddrRm{}
	return &this
}

func (o Ipv6AddrRm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableIpv6AddrRm struct {
	value *Ipv6AddrRm
	isSet bool
}

func (v NullableIpv6AddrRm) Get() *Ipv6AddrRm {
	return v.value
}

func (v *NullableIpv6AddrRm) Set(val *Ipv6AddrRm) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6AddrRm) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6AddrRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6AddrRm(val *Ipv6AddrRm) *NullableIpv6AddrRm {
	return &NullableIpv6AddrRm{value: val, isSet: true}
}

func (v NullableIpv6AddrRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6AddrRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


