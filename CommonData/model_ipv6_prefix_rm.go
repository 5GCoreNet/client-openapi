/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
)

// Ipv6PrefixRm String identifying an IPv6 address prefix formatted according to clause 4 of RFC 5952 with the OpenAPI 'nullable: true' property. IPv6Prefix data type may contain an individual /128 IPv6 address. 
type Ipv6PrefixRm struct {
}

// NewIpv6PrefixRm instantiates a new Ipv6PrefixRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6PrefixRm() *Ipv6PrefixRm {
	this := Ipv6PrefixRm{}
	return &this
}

// NewIpv6PrefixRmWithDefaults instantiates a new Ipv6PrefixRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6PrefixRmWithDefaults() *Ipv6PrefixRm {
	this := Ipv6PrefixRm{}
	return &this
}

func (o Ipv6PrefixRm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableIpv6PrefixRm struct {
	value *Ipv6PrefixRm
	isSet bool
}

func (v NullableIpv6PrefixRm) Get() *Ipv6PrefixRm {
	return v.value
}

func (v *NullableIpv6PrefixRm) Set(val *Ipv6PrefixRm) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6PrefixRm) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6PrefixRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6PrefixRm(val *Ipv6PrefixRm) *NullableIpv6PrefixRm {
	return &NullableIpv6PrefixRm{value: val, isSet: true}
}

func (v NullableIpv6PrefixRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6PrefixRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


