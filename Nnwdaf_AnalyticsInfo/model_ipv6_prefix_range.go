/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// Ipv6PrefixRange Range of IPv6 prefixes
type Ipv6PrefixRange struct {
	Start *Ipv6Prefix `json:"start,omitempty"`
	End *Ipv6Prefix `json:"end,omitempty"`
}

// NewIpv6PrefixRange instantiates a new Ipv6PrefixRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6PrefixRange() *Ipv6PrefixRange {
	this := Ipv6PrefixRange{}
	return &this
}

// NewIpv6PrefixRangeWithDefaults instantiates a new Ipv6PrefixRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6PrefixRangeWithDefaults() *Ipv6PrefixRange {
	this := Ipv6PrefixRange{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Ipv6PrefixRange) GetStart() Ipv6Prefix {
	if o == nil || isNil(o.Start) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6PrefixRange) GetStartOk() (*Ipv6Prefix, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Ipv6PrefixRange) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given Ipv6Prefix and assigns it to the Start field.
func (o *Ipv6PrefixRange) SetStart(v Ipv6Prefix) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Ipv6PrefixRange) GetEnd() Ipv6Prefix {
	if o == nil || isNil(o.End) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6PrefixRange) GetEndOk() (*Ipv6Prefix, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Ipv6PrefixRange) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given Ipv6Prefix and assigns it to the End field.
func (o *Ipv6PrefixRange) SetEnd(v Ipv6Prefix) {
	o.End = &v
}

func (o Ipv6PrefixRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableIpv6PrefixRange struct {
	value *Ipv6PrefixRange
	isSet bool
}

func (v NullableIpv6PrefixRange) Get() *Ipv6PrefixRange {
	return v.value
}

func (v *NullableIpv6PrefixRange) Set(val *Ipv6PrefixRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6PrefixRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6PrefixRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6PrefixRange(val *Ipv6PrefixRange) *NullableIpv6PrefixRange {
	return &NullableIpv6PrefixRange{value: val, isSet: true}
}

func (v NullableIpv6PrefixRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6PrefixRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


