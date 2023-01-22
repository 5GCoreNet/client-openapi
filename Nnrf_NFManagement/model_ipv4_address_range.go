/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFManagement

import (
	"encoding/json"
)

// Ipv4AddressRange Range of IPv4 addresses
type Ipv4AddressRange struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Start *string `json:"start,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	End *string `json:"end,omitempty"`
}

// NewIpv4AddressRange instantiates a new Ipv4AddressRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv4AddressRange() *Ipv4AddressRange {
	this := Ipv4AddressRange{}
	return &this
}

// NewIpv4AddressRangeWithDefaults instantiates a new Ipv4AddressRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv4AddressRangeWithDefaults() *Ipv4AddressRange {
	this := Ipv4AddressRange{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Ipv4AddressRange) GetStart() string {
	if o == nil || isNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv4AddressRange) GetStartOk() (*string, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Ipv4AddressRange) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *Ipv4AddressRange) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Ipv4AddressRange) GetEnd() string {
	if o == nil || isNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv4AddressRange) GetEndOk() (*string, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Ipv4AddressRange) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *Ipv4AddressRange) SetEnd(v string) {
	o.End = &v
}

func (o Ipv4AddressRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableIpv4AddressRange struct {
	value *Ipv4AddressRange
	isSet bool
}

func (v NullableIpv4AddressRange) Get() *Ipv4AddressRange {
	return v.value
}

func (v *NullableIpv4AddressRange) Set(val *Ipv4AddressRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv4AddressRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv4AddressRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv4AddressRange(val *Ipv4AddressRange) *NullableIpv4AddressRange {
	return &NullableIpv4AddressRange{value: val, isSet: true}
}

func (v NullableIpv4AddressRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv4AddressRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


