/*
CAPIF_Events_API

API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CAPIF_Events_API

import (
	"encoding/json"
)

// Ipv6AddressRange Represents IPv6 address range.
type Ipv6AddressRange struct {
	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	Start string `json:"start"`
	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	End string `json:"end"`
}

// NewIpv6AddressRange instantiates a new Ipv6AddressRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6AddressRange(start string, end string) *Ipv6AddressRange {
	this := Ipv6AddressRange{}
	this.Start = start
	this.End = end
	return &this
}

// NewIpv6AddressRangeWithDefaults instantiates a new Ipv6AddressRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6AddressRangeWithDefaults() *Ipv6AddressRange {
	this := Ipv6AddressRange{}
	return &this
}

// GetStart returns the Start field value
func (o *Ipv6AddressRange) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *Ipv6AddressRange) GetStartOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *Ipv6AddressRange) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *Ipv6AddressRange) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *Ipv6AddressRange) GetEndOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *Ipv6AddressRange) SetEnd(v string) {
	o.End = v
}

func (o Ipv6AddressRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableIpv6AddressRange struct {
	value *Ipv6AddressRange
	isSet bool
}

func (v NullableIpv6AddressRange) Get() *Ipv6AddressRange {
	return v.value
}

func (v *NullableIpv6AddressRange) Set(val *Ipv6AddressRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6AddressRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6AddressRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6AddressRange(val *Ipv6AddressRange) *NullableIpv6AddressRange {
	return &NullableIpv6AddressRange{value: val, isSet: true}
}

func (v NullableIpv6AddressRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6AddressRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


