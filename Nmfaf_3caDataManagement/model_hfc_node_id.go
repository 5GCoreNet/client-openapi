/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// HfcNodeId REpresents the HFC Node Identifer received over NGAP.
type HfcNodeId struct {
	// This IE represents the identifier of the HFC node Id as specified in CableLabs WR-TR-5WWC-ARCH. It is provisioned by the wireline operator as part of wireline operations and may contain up to six characters. 
	HfcNId string `json:"hfcNId"`
}

// NewHfcNodeId instantiates a new HfcNodeId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHfcNodeId(hfcNId string) *HfcNodeId {
	this := HfcNodeId{}
	this.HfcNId = hfcNId
	return &this
}

// NewHfcNodeIdWithDefaults instantiates a new HfcNodeId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHfcNodeIdWithDefaults() *HfcNodeId {
	this := HfcNodeId{}
	return &this
}

// GetHfcNId returns the HfcNId field value
func (o *HfcNodeId) GetHfcNId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HfcNId
}

// GetHfcNIdOk returns a tuple with the HfcNId field value
// and a boolean to check if the value has been set.
func (o *HfcNodeId) GetHfcNIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HfcNId, true
}

// SetHfcNId sets field value
func (o *HfcNodeId) SetHfcNId(v string) {
	o.HfcNId = v
}

func (o HfcNodeId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hfcNId"] = o.HfcNId
	}
	return json.Marshal(toSerialize)
}

type NullableHfcNodeId struct {
	value *HfcNodeId
	isSet bool
}

func (v NullableHfcNodeId) Get() *HfcNodeId {
	return v.value
}

func (v *NullableHfcNodeId) Set(val *HfcNodeId) {
	v.value = val
	v.isSet = true
}

func (v NullableHfcNodeId) IsSet() bool {
	return v.isSet
}

func (v *NullableHfcNodeId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHfcNodeId(val *HfcNodeId) *NullableHfcNodeId {
	return &NullableHfcNodeId{value: val, isSet: true}
}

func (v NullableHfcNodeId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHfcNodeId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


