/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// Tmgi1 Temporary Mobile Group Identity
type Tmgi1 struct {
	// MBS Service ID
	MbsServiceId string `json:"mbsServiceId"`
	PlmnId PlmnId1 `json:"plmnId"`
}

// NewTmgi1 instantiates a new Tmgi1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTmgi1(mbsServiceId string, plmnId PlmnId1) *Tmgi1 {
	this := Tmgi1{}
	this.MbsServiceId = mbsServiceId
	this.PlmnId = plmnId
	return &this
}

// NewTmgi1WithDefaults instantiates a new Tmgi1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTmgi1WithDefaults() *Tmgi1 {
	this := Tmgi1{}
	return &this
}

// GetMbsServiceId returns the MbsServiceId field value
func (o *Tmgi1) GetMbsServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MbsServiceId
}

// GetMbsServiceIdOk returns a tuple with the MbsServiceId field value
// and a boolean to check if the value has been set.
func (o *Tmgi1) GetMbsServiceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MbsServiceId, true
}

// SetMbsServiceId sets field value
func (o *Tmgi1) SetMbsServiceId(v string) {
	o.MbsServiceId = v
}

// GetPlmnId returns the PlmnId field value
func (o *Tmgi1) GetPlmnId() PlmnId1 {
	if o == nil {
		var ret PlmnId1
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *Tmgi1) GetPlmnIdOk() (*PlmnId1, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *Tmgi1) SetPlmnId(v PlmnId1) {
	o.PlmnId = v
}

func (o Tmgi1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mbsServiceId"] = o.MbsServiceId
	}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	return json.Marshal(toSerialize)
}

type NullableTmgi1 struct {
	value *Tmgi1
	isSet bool
}

func (v NullableTmgi1) Get() *Tmgi1 {
	return v.value
}

func (v *NullableTmgi1) Set(val *Tmgi1) {
	v.value = val
	v.isSet = true
}

func (v NullableTmgi1) IsSet() bool {
	return v.isSet
}

func (v *NullableTmgi1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTmgi1(val *Tmgi1) *NullableTmgi1 {
	return &NullableTmgi1{value: val, isSet: true}
}

func (v NullableTmgi1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTmgi1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


