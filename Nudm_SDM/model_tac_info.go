/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_SDM

import (
	"encoding/json"
)

// TacInfo contains tracking area information (tracking area codes).
type TacInfo struct {
	TacList []string `json:"tacList"`
}

// NewTacInfo instantiates a new TacInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTacInfo(tacList []string) *TacInfo {
	this := TacInfo{}
	this.TacList = tacList
	return &this
}

// NewTacInfoWithDefaults instantiates a new TacInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTacInfoWithDefaults() *TacInfo {
	this := TacInfo{}
	return &this
}

// GetTacList returns the TacList field value
func (o *TacInfo) GetTacList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TacList
}

// GetTacListOk returns a tuple with the TacList field value
// and a boolean to check if the value has been set.
func (o *TacInfo) GetTacListOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TacList, true
}

// SetTacList sets field value
func (o *TacInfo) SetTacList(v []string) {
	o.TacList = v
}

func (o TacInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tacList"] = o.TacList
	}
	return json.Marshal(toSerialize)
}

type NullableTacInfo struct {
	value *TacInfo
	isSet bool
}

func (v NullableTacInfo) Get() *TacInfo {
	return v.value
}

func (v *NullableTacInfo) Set(val *TacInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTacInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTacInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTacInfo(val *TacInfo) *NullableTacInfo {
	return &NullableTacInfo{value: val, isSet: true}
}

func (v NullableTacInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTacInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


