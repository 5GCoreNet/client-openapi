/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// NfGroupCond Subscription to a set of NFs based on their Group Id
type NfGroupCond struct {
	NfType string `json:"nfType"`
	// Identifier of a group of NFs.
	NfGroupId string `json:"nfGroupId"`
}

// NewNfGroupCond instantiates a new NfGroupCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfGroupCond(nfType string, nfGroupId string) *NfGroupCond {
	this := NfGroupCond{}
	this.NfType = nfType
	this.NfGroupId = nfGroupId
	return &this
}

// NewNfGroupCondWithDefaults instantiates a new NfGroupCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfGroupCondWithDefaults() *NfGroupCond {
	this := NfGroupCond{}
	return &this
}

// GetNfType returns the NfType field value
func (o *NfGroupCond) GetNfType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfType
}

// GetNfTypeOk returns a tuple with the NfType field value
// and a boolean to check if the value has been set.
func (o *NfGroupCond) GetNfTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NfType, true
}

// SetNfType sets field value
func (o *NfGroupCond) SetNfType(v string) {
	o.NfType = v
}

// GetNfGroupId returns the NfGroupId field value
func (o *NfGroupCond) GetNfGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfGroupId
}

// GetNfGroupIdOk returns a tuple with the NfGroupId field value
// and a boolean to check if the value has been set.
func (o *NfGroupCond) GetNfGroupIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NfGroupId, true
}

// SetNfGroupId sets field value
func (o *NfGroupCond) SetNfGroupId(v string) {
	o.NfGroupId = v
}

func (o NfGroupCond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nfType"] = o.NfType
	}
	if true {
		toSerialize["nfGroupId"] = o.NfGroupId
	}
	return json.Marshal(toSerialize)
}

type NullableNfGroupCond struct {
	value *NfGroupCond
	isSet bool
}

func (v NullableNfGroupCond) Get() *NfGroupCond {
	return v.value
}

func (v *NullableNfGroupCond) Set(val *NfGroupCond) {
	v.value = val
	v.isSet = true
}

func (v NullableNfGroupCond) IsSet() bool {
	return v.isSet
}

func (v *NullableNfGroupCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfGroupCond(val *NfGroupCond) *NullableNfGroupCond {
	return &NullableNfGroupCond{value: val, isSet: true}
}

func (v NullableNfGroupCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfGroupCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


