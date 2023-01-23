/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// RemoteAddress struct for RemoteAddress
type RemoteAddress struct {
	Ipv4Address *string `json:"ipv4Address,omitempty"`
	Ipv6Address *Ipv6Addr `json:"ipv6Address,omitempty"`
}

// NewRemoteAddress instantiates a new RemoteAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteAddress() *RemoteAddress {
	this := RemoteAddress{}
	return &this
}

// NewRemoteAddressWithDefaults instantiates a new RemoteAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteAddressWithDefaults() *RemoteAddress {
	this := RemoteAddress{}
	return &this
}

// GetIpv4Address returns the Ipv4Address field value if set, zero value otherwise.
func (o *RemoteAddress) GetIpv4Address() string {
	if o == nil || isNil(o.Ipv4Address) {
		var ret string
		return ret
	}
	return *o.Ipv4Address
}

// GetIpv4AddressOk returns a tuple with the Ipv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAddress) GetIpv4AddressOk() (*string, bool) {
	if o == nil || isNil(o.Ipv4Address) {
    return nil, false
	}
	return o.Ipv4Address, true
}

// HasIpv4Address returns a boolean if a field has been set.
func (o *RemoteAddress) HasIpv4Address() bool {
	if o != nil && !isNil(o.Ipv4Address) {
		return true
	}

	return false
}

// SetIpv4Address gets a reference to the given string and assigns it to the Ipv4Address field.
func (o *RemoteAddress) SetIpv4Address(v string) {
	o.Ipv4Address = &v
}

// GetIpv6Address returns the Ipv6Address field value if set, zero value otherwise.
func (o *RemoteAddress) GetIpv6Address() Ipv6Addr {
	if o == nil || isNil(o.Ipv6Address) {
		var ret Ipv6Addr
		return ret
	}
	return *o.Ipv6Address
}

// GetIpv6AddressOk returns a tuple with the Ipv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteAddress) GetIpv6AddressOk() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.Ipv6Address) {
    return nil, false
	}
	return o.Ipv6Address, true
}

// HasIpv6Address returns a boolean if a field has been set.
func (o *RemoteAddress) HasIpv6Address() bool {
	if o != nil && !isNil(o.Ipv6Address) {
		return true
	}

	return false
}

// SetIpv6Address gets a reference to the given Ipv6Addr and assigns it to the Ipv6Address field.
func (o *RemoteAddress) SetIpv6Address(v Ipv6Addr) {
	o.Ipv6Address = &v
}

func (o RemoteAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ipv4Address) {
		toSerialize["ipv4Address"] = o.Ipv4Address
	}
	if !isNil(o.Ipv6Address) {
		toSerialize["ipv6Address"] = o.Ipv6Address
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteAddress struct {
	value *RemoteAddress
	isSet bool
}

func (v NullableRemoteAddress) Get() *RemoteAddress {
	return v.value
}

func (v *NullableRemoteAddress) Set(val *RemoteAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteAddress(val *RemoteAddress) *NullableRemoteAddress {
	return &NullableRemoteAddress{value: val, isSet: true}
}

func (v NullableRemoteAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


