/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// PrivateIdentities A list of IMS Private Identities
type PrivateIdentities struct {
	PrivateIdentities []PrivateIdentity `json:"privateIdentities"`
}

// NewPrivateIdentities instantiates a new PrivateIdentities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateIdentities(privateIdentities []PrivateIdentity) *PrivateIdentities {
	this := PrivateIdentities{}
	this.PrivateIdentities = privateIdentities
	return &this
}

// NewPrivateIdentitiesWithDefaults instantiates a new PrivateIdentities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateIdentitiesWithDefaults() *PrivateIdentities {
	this := PrivateIdentities{}
	return &this
}

// GetPrivateIdentities returns the PrivateIdentities field value
func (o *PrivateIdentities) GetPrivateIdentities() []PrivateIdentity {
	if o == nil {
		var ret []PrivateIdentity
		return ret
	}

	return o.PrivateIdentities
}

// GetPrivateIdentitiesOk returns a tuple with the PrivateIdentities field value
// and a boolean to check if the value has been set.
func (o *PrivateIdentities) GetPrivateIdentitiesOk() ([]PrivateIdentity, bool) {
	if o == nil {
    return nil, false
	}
	return o.PrivateIdentities, true
}

// SetPrivateIdentities sets field value
func (o *PrivateIdentities) SetPrivateIdentities(v []PrivateIdentity) {
	o.PrivateIdentities = v
}

func (o PrivateIdentities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["privateIdentities"] = o.PrivateIdentities
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateIdentities struct {
	value *PrivateIdentities
	isSet bool
}

func (v NullablePrivateIdentities) Get() *PrivateIdentities {
	return v.value
}

func (v *NullablePrivateIdentities) Set(val *PrivateIdentities) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateIdentities) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateIdentities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateIdentities(val *PrivateIdentities) *NullablePrivateIdentities {
	return &NullablePrivateIdentities{value: val, isSet: true}
}

func (v NullablePrivateIdentities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateIdentities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


