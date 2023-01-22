/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_DataManagement

import (
	"encoding/json"
)

// SliceLoadLevelInformation Contains load level information applicable for one or several slices.
type SliceLoadLevelInformation struct {
	// Load level information of the network slice and the optionally associated network slice  instance. 
	LoadLevelInformation int32 `json:"loadLevelInformation"`
	// Identification(s) of network slice to which the subscription applies.
	Snssais []Snssai `json:"snssais"`
}

// NewSliceLoadLevelInformation instantiates a new SliceLoadLevelInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSliceLoadLevelInformation(loadLevelInformation int32, snssais []Snssai) *SliceLoadLevelInformation {
	this := SliceLoadLevelInformation{}
	this.LoadLevelInformation = loadLevelInformation
	this.Snssais = snssais
	return &this
}

// NewSliceLoadLevelInformationWithDefaults instantiates a new SliceLoadLevelInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliceLoadLevelInformationWithDefaults() *SliceLoadLevelInformation {
	this := SliceLoadLevelInformation{}
	return &this
}

// GetLoadLevelInformation returns the LoadLevelInformation field value
func (o *SliceLoadLevelInformation) GetLoadLevelInformation() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoadLevelInformation
}

// GetLoadLevelInformationOk returns a tuple with the LoadLevelInformation field value
// and a boolean to check if the value has been set.
func (o *SliceLoadLevelInformation) GetLoadLevelInformationOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LoadLevelInformation, true
}

// SetLoadLevelInformation sets field value
func (o *SliceLoadLevelInformation) SetLoadLevelInformation(v int32) {
	o.LoadLevelInformation = v
}

// GetSnssais returns the Snssais field value
func (o *SliceLoadLevelInformation) GetSnssais() []Snssai {
	if o == nil {
		var ret []Snssai
		return ret
	}

	return o.Snssais
}

// GetSnssaisOk returns a tuple with the Snssais field value
// and a boolean to check if the value has been set.
func (o *SliceLoadLevelInformation) GetSnssaisOk() ([]Snssai, bool) {
	if o == nil {
    return nil, false
	}
	return o.Snssais, true
}

// SetSnssais sets field value
func (o *SliceLoadLevelInformation) SetSnssais(v []Snssai) {
	o.Snssais = v
}

func (o SliceLoadLevelInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["loadLevelInformation"] = o.LoadLevelInformation
	}
	if true {
		toSerialize["snssais"] = o.Snssais
	}
	return json.Marshal(toSerialize)
}

type NullableSliceLoadLevelInformation struct {
	value *SliceLoadLevelInformation
	isSet bool
}

func (v NullableSliceLoadLevelInformation) Get() *SliceLoadLevelInformation {
	return v.value
}

func (v *NullableSliceLoadLevelInformation) Set(val *SliceLoadLevelInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableSliceLoadLevelInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableSliceLoadLevelInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliceLoadLevelInformation(val *SliceLoadLevelInformation) *NullableSliceLoadLevelInformation {
	return &NullableSliceLoadLevelInformation{value: val, isSet: true}
}

func (v NullableSliceLoadLevelInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliceLoadLevelInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


