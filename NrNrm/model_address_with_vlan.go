/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
)

// AddressWithVlan struct for AddressWithVlan
type AddressWithVlan struct {
	Ipv4Address *string `json:"ipv4Address,omitempty"`
	Ipv6Address *Ipv6Addr `json:"ipv6Address,omitempty"`
	VlanId *int32 `json:"vlanId,omitempty"`
}

// NewAddressWithVlan instantiates a new AddressWithVlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressWithVlan() *AddressWithVlan {
	this := AddressWithVlan{}
	return &this
}

// NewAddressWithVlanWithDefaults instantiates a new AddressWithVlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithVlanWithDefaults() *AddressWithVlan {
	this := AddressWithVlan{}
	return &this
}

// GetIpv4Address returns the Ipv4Address field value if set, zero value otherwise.
func (o *AddressWithVlan) GetIpv4Address() string {
	if o == nil || isNil(o.Ipv4Address) {
		var ret string
		return ret
	}
	return *o.Ipv4Address
}

// GetIpv4AddressOk returns a tuple with the Ipv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressWithVlan) GetIpv4AddressOk() (*string, bool) {
	if o == nil || isNil(o.Ipv4Address) {
    return nil, false
	}
	return o.Ipv4Address, true
}

// HasIpv4Address returns a boolean if a field has been set.
func (o *AddressWithVlan) HasIpv4Address() bool {
	if o != nil && !isNil(o.Ipv4Address) {
		return true
	}

	return false
}

// SetIpv4Address gets a reference to the given string and assigns it to the Ipv4Address field.
func (o *AddressWithVlan) SetIpv4Address(v string) {
	o.Ipv4Address = &v
}

// GetIpv6Address returns the Ipv6Address field value if set, zero value otherwise.
func (o *AddressWithVlan) GetIpv6Address() Ipv6Addr {
	if o == nil || isNil(o.Ipv6Address) {
		var ret Ipv6Addr
		return ret
	}
	return *o.Ipv6Address
}

// GetIpv6AddressOk returns a tuple with the Ipv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressWithVlan) GetIpv6AddressOk() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.Ipv6Address) {
    return nil, false
	}
	return o.Ipv6Address, true
}

// HasIpv6Address returns a boolean if a field has been set.
func (o *AddressWithVlan) HasIpv6Address() bool {
	if o != nil && !isNil(o.Ipv6Address) {
		return true
	}

	return false
}

// SetIpv6Address gets a reference to the given Ipv6Addr and assigns it to the Ipv6Address field.
func (o *AddressWithVlan) SetIpv6Address(v Ipv6Addr) {
	o.Ipv6Address = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *AddressWithVlan) GetVlanId() int32 {
	if o == nil || isNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressWithVlan) GetVlanIdOk() (*int32, bool) {
	if o == nil || isNil(o.VlanId) {
    return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *AddressWithVlan) HasVlanId() bool {
	if o != nil && !isNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *AddressWithVlan) SetVlanId(v int32) {
	o.VlanId = &v
}

func (o AddressWithVlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ipv4Address) {
		toSerialize["ipv4Address"] = o.Ipv4Address
	}
	if !isNil(o.Ipv6Address) {
		toSerialize["ipv6Address"] = o.Ipv6Address
	}
	if !isNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	return json.Marshal(toSerialize)
}

type NullableAddressWithVlan struct {
	value *AddressWithVlan
	isSet bool
}

func (v NullableAddressWithVlan) Get() *AddressWithVlan {
	return v.value
}

func (v *NullableAddressWithVlan) Set(val *AddressWithVlan) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressWithVlan) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressWithVlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressWithVlan(val *AddressWithVlan) *NullableAddressWithVlan {
	return &NullableAddressWithVlan{value: val, isSet: true}
}

func (v NullableAddressWithVlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressWithVlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


