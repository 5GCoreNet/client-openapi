/*
GBA BSF Nbsp_GBA Service

GBA BSF Nbsp_GBA Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nbsp_GBA

import (
	"encoding/json"
)

// NafId NAF ID, containing the NAF FQDN and the Ua Security Protocol Identifier
type NafId struct {
	// Fully Qualified Domain Name
	NafFqdn string `json:"nafFqdn"`
	UaSecProtId string `json:"uaSecProtId"`
}

// NewNafId instantiates a new NafId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNafId(nafFqdn string, uaSecProtId string) *NafId {
	this := NafId{}
	this.NafFqdn = nafFqdn
	this.UaSecProtId = uaSecProtId
	return &this
}

// NewNafIdWithDefaults instantiates a new NafId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNafIdWithDefaults() *NafId {
	this := NafId{}
	return &this
}

// GetNafFqdn returns the NafFqdn field value
func (o *NafId) GetNafFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NafFqdn
}

// GetNafFqdnOk returns a tuple with the NafFqdn field value
// and a boolean to check if the value has been set.
func (o *NafId) GetNafFqdnOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NafFqdn, true
}

// SetNafFqdn sets field value
func (o *NafId) SetNafFqdn(v string) {
	o.NafFqdn = v
}

// GetUaSecProtId returns the UaSecProtId field value
func (o *NafId) GetUaSecProtId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UaSecProtId
}

// GetUaSecProtIdOk returns a tuple with the UaSecProtId field value
// and a boolean to check if the value has been set.
func (o *NafId) GetUaSecProtIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UaSecProtId, true
}

// SetUaSecProtId sets field value
func (o *NafId) SetUaSecProtId(v string) {
	o.UaSecProtId = v
}

func (o NafId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nafFqdn"] = o.NafFqdn
	}
	if true {
		toSerialize["uaSecProtId"] = o.UaSecProtId
	}
	return json.Marshal(toSerialize)
}

type NullableNafId struct {
	value *NafId
	isSet bool
}

func (v NullableNafId) Get() *NafId {
	return v.value
}

func (v *NullableNafId) Set(val *NafId) {
	v.value = val
	v.isSet = true
}

func (v NullableNafId) IsSet() bool {
	return v.isSet
}

func (v *NullableNafId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNafId(val *NafId) *NullableNafId {
	return &NullableNafId{value: val, isSet: true}
}

func (v NullableNafId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNafId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


