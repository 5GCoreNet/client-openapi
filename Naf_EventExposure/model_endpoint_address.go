/*
Naf_EventExposure

AF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Naf_EventExposure

import (
	"encoding/json"
)

// EndpointAddress struct for EndpointAddress
type EndpointAddress struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr *string `json:"ipv4Addr,omitempty"`
	Ipv6Addr *Ipv6Addr `json:"ipv6Addr,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	PortNumber int32 `json:"portNumber"`
}

// NewEndpointAddress instantiates a new EndpointAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointAddress(portNumber int32) *EndpointAddress {
	this := EndpointAddress{}
	this.PortNumber = portNumber
	return &this
}

// NewEndpointAddressWithDefaults instantiates a new EndpointAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointAddressWithDefaults() *EndpointAddress {
	this := EndpointAddress{}
	return &this
}

// GetIpv4Addr returns the Ipv4Addr field value if set, zero value otherwise.
func (o *EndpointAddress) GetIpv4Addr() string {
	if o == nil || isNil(o.Ipv4Addr) {
		var ret string
		return ret
	}
	return *o.Ipv4Addr
}

// GetIpv4AddrOk returns a tuple with the Ipv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAddress) GetIpv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.Ipv4Addr) {
    return nil, false
	}
	return o.Ipv4Addr, true
}

// HasIpv4Addr returns a boolean if a field has been set.
func (o *EndpointAddress) HasIpv4Addr() bool {
	if o != nil && !isNil(o.Ipv4Addr) {
		return true
	}

	return false
}

// SetIpv4Addr gets a reference to the given string and assigns it to the Ipv4Addr field.
func (o *EndpointAddress) SetIpv4Addr(v string) {
	o.Ipv4Addr = &v
}

// GetIpv6Addr returns the Ipv6Addr field value if set, zero value otherwise.
func (o *EndpointAddress) GetIpv6Addr() Ipv6Addr {
	if o == nil || isNil(o.Ipv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.Ipv6Addr
}

// GetIpv6AddrOk returns a tuple with the Ipv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAddress) GetIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.Ipv6Addr) {
    return nil, false
	}
	return o.Ipv6Addr, true
}

// HasIpv6Addr returns a boolean if a field has been set.
func (o *EndpointAddress) HasIpv6Addr() bool {
	if o != nil && !isNil(o.Ipv6Addr) {
		return true
	}

	return false
}

// SetIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the Ipv6Addr field.
func (o *EndpointAddress) SetIpv6Addr(v Ipv6Addr) {
	o.Ipv6Addr = &v
}

// GetPortNumber returns the PortNumber field value
func (o *EndpointAddress) GetPortNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PortNumber
}

// GetPortNumberOk returns a tuple with the PortNumber field value
// and a boolean to check if the value has been set.
func (o *EndpointAddress) GetPortNumberOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PortNumber, true
}

// SetPortNumber sets field value
func (o *EndpointAddress) SetPortNumber(v int32) {
	o.PortNumber = v
}

func (o EndpointAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ipv4Addr) {
		toSerialize["ipv4Addr"] = o.Ipv4Addr
	}
	if !isNil(o.Ipv6Addr) {
		toSerialize["ipv6Addr"] = o.Ipv6Addr
	}
	if true {
		toSerialize["portNumber"] = o.PortNumber
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointAddress struct {
	value *EndpointAddress
	isSet bool
}

func (v NullableEndpointAddress) Get() *EndpointAddress {
	return v.value
}

func (v *NullableEndpointAddress) Set(val *EndpointAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointAddress(val *EndpointAddress) *NullableEndpointAddress {
	return &NullableEndpointAddress{value: val, isSet: true}
}

func (v NullableEndpointAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


