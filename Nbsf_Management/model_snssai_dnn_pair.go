/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsf_Management

import (
	"encoding/json"
)

// SnssaiDnnPair Contains a S-NSSAI and DNN combination.
type SnssaiDnnPair struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn"`
	Snssai Snssai `json:"snssai"`
}

// NewSnssaiDnnPair instantiates a new SnssaiDnnPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiDnnPair(dnn string, snssai Snssai) *SnssaiDnnPair {
	this := SnssaiDnnPair{}
	this.Dnn = dnn
	this.Snssai = snssai
	return &this
}

// NewSnssaiDnnPairWithDefaults instantiates a new SnssaiDnnPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiDnnPairWithDefaults() *SnssaiDnnPair {
	this := SnssaiDnnPair{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *SnssaiDnnPair) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *SnssaiDnnPair) GetDnnOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *SnssaiDnnPair) SetDnn(v string) {
	o.Dnn = v
}

// GetSnssai returns the Snssai field value
func (o *SnssaiDnnPair) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *SnssaiDnnPair) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *SnssaiDnnPair) SetSnssai(v Snssai) {
	o.Snssai = v
}

func (o SnssaiDnnPair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnn"] = o.Dnn
	}
	if true {
		toSerialize["snssai"] = o.Snssai
	}
	return json.Marshal(toSerialize)
}

type NullableSnssaiDnnPair struct {
	value *SnssaiDnnPair
	isSet bool
}

func (v NullableSnssaiDnnPair) Get() *SnssaiDnnPair {
	return v.value
}

func (v *NullableSnssaiDnnPair) Set(val *SnssaiDnnPair) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiDnnPair) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiDnnPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiDnnPair(val *SnssaiDnnPair) *NullableSnssaiDnnPair {
	return &NullableSnssaiDnnPair{value: val, isSet: true}
}

func (v NullableSnssaiDnnPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiDnnPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


