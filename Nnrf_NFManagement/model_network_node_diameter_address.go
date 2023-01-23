/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
)

// NetworkNodeDiameterAddress struct for NetworkNodeDiameterAddress
type NetworkNodeDiameterAddress struct {
	// Fully Qualified Domain Name
	Name string `json:"name"`
	// Fully Qualified Domain Name
	Realm string `json:"realm"`
}

// NewNetworkNodeDiameterAddress instantiates a new NetworkNodeDiameterAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkNodeDiameterAddress(name string, realm string) *NetworkNodeDiameterAddress {
	this := NetworkNodeDiameterAddress{}
	this.Name = name
	this.Realm = realm
	return &this
}

// NewNetworkNodeDiameterAddressWithDefaults instantiates a new NetworkNodeDiameterAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkNodeDiameterAddressWithDefaults() *NetworkNodeDiameterAddress {
	this := NetworkNodeDiameterAddress{}
	return &this
}

// GetName returns the Name field value
func (o *NetworkNodeDiameterAddress) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkNodeDiameterAddress) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkNodeDiameterAddress) SetName(v string) {
	o.Name = v
}

// GetRealm returns the Realm field value
func (o *NetworkNodeDiameterAddress) GetRealm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Realm
}

// GetRealmOk returns a tuple with the Realm field value
// and a boolean to check if the value has been set.
func (o *NetworkNodeDiameterAddress) GetRealmOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Realm, true
}

// SetRealm sets field value
func (o *NetworkNodeDiameterAddress) SetRealm(v string) {
	o.Realm = v
}

func (o NetworkNodeDiameterAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["realm"] = o.Realm
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkNodeDiameterAddress struct {
	value *NetworkNodeDiameterAddress
	isSet bool
}

func (v NullableNetworkNodeDiameterAddress) Get() *NetworkNodeDiameterAddress {
	return v.value
}

func (v *NullableNetworkNodeDiameterAddress) Set(val *NetworkNodeDiameterAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkNodeDiameterAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkNodeDiameterAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkNodeDiameterAddress(val *NetworkNodeDiameterAddress) *NullableNetworkNodeDiameterAddress {
	return &NullableNetworkNodeDiameterAddress{value: val, isSet: true}
}

func (v NullableNetworkNodeDiameterAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkNodeDiameterAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


