/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// LocalAddress struct for LocalAddress
type LocalAddress struct {
	AddressWithVlan *AddressWithVlan `json:"addressWithVlan,omitempty"`
	Port *int32 `json:"port,omitempty"`
}

// NewLocalAddress instantiates a new LocalAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalAddress() *LocalAddress {
	this := LocalAddress{}
	return &this
}

// NewLocalAddressWithDefaults instantiates a new LocalAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalAddressWithDefaults() *LocalAddress {
	this := LocalAddress{}
	return &this
}

// GetAddressWithVlan returns the AddressWithVlan field value if set, zero value otherwise.
func (o *LocalAddress) GetAddressWithVlan() AddressWithVlan {
	if o == nil || isNil(o.AddressWithVlan) {
		var ret AddressWithVlan
		return ret
	}
	return *o.AddressWithVlan
}

// GetAddressWithVlanOk returns a tuple with the AddressWithVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalAddress) GetAddressWithVlanOk() (*AddressWithVlan, bool) {
	if o == nil || isNil(o.AddressWithVlan) {
    return nil, false
	}
	return o.AddressWithVlan, true
}

// HasAddressWithVlan returns a boolean if a field has been set.
func (o *LocalAddress) HasAddressWithVlan() bool {
	if o != nil && !isNil(o.AddressWithVlan) {
		return true
	}

	return false
}

// SetAddressWithVlan gets a reference to the given AddressWithVlan and assigns it to the AddressWithVlan field.
func (o *LocalAddress) SetAddressWithVlan(v AddressWithVlan) {
	o.AddressWithVlan = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LocalAddress) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalAddress) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *LocalAddress) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *LocalAddress) SetPort(v int32) {
	o.Port = &v
}

func (o LocalAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AddressWithVlan) {
		toSerialize["addressWithVlan"] = o.AddressWithVlan
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableLocalAddress struct {
	value *LocalAddress
	isSet bool
}

func (v NullableLocalAddress) Get() *LocalAddress {
	return v.value
}

func (v *NullableLocalAddress) Set(val *LocalAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalAddress(val *LocalAddress) *NullableLocalAddress {
	return &NullableLocalAddress{value: val, isSet: true}
}

func (v NullableLocalAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


