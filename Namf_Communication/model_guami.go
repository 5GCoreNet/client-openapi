/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
)

// Guami Globally Unique AMF Identifier constructed out of PLMN, Network and AMF identity.
type Guami struct {
	PlmnId PlmnIdNid `json:"plmnId"`
	// String identifying the AMF ID composed of AMF Region ID (8 bits), AMF Set ID (10 bits) and AMF  Pointer (6 bits) as specified in clause 2.10.1 of 3GPP TS 23.003. It is encoded as a string of  6 hexadecimal characters (i.e., 24 bits).  
	AmfId string `json:"amfId"`
}

// NewGuami instantiates a new Guami object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuami(plmnId PlmnIdNid, amfId string) *Guami {
	this := Guami{}
	this.PlmnId = plmnId
	this.AmfId = amfId
	return &this
}

// NewGuamiWithDefaults instantiates a new Guami object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuamiWithDefaults() *Guami {
	this := Guami{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *Guami) GetPlmnId() PlmnIdNid {
	if o == nil {
		var ret PlmnIdNid
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *Guami) GetPlmnIdOk() (*PlmnIdNid, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *Guami) SetPlmnId(v PlmnIdNid) {
	o.PlmnId = v
}

// GetAmfId returns the AmfId field value
func (o *Guami) GetAmfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmfId
}

// GetAmfIdOk returns a tuple with the AmfId field value
// and a boolean to check if the value has been set.
func (o *Guami) GetAmfIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AmfId, true
}

// SetAmfId sets field value
func (o *Guami) SetAmfId(v string) {
	o.AmfId = v
}

func (o Guami) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if true {
		toSerialize["amfId"] = o.AmfId
	}
	return json.Marshal(toSerialize)
}

type NullableGuami struct {
	value *Guami
	isSet bool
}

func (v NullableGuami) Get() *Guami {
	return v.value
}

func (v *NullableGuami) Set(val *Guami) {
	v.value = val
	v.isSet = true
}

func (v NullableGuami) IsSet() bool {
	return v.isSet
}

func (v *NullableGuami) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuami(val *Guami) *NullableGuami {
	return &NullableGuami{value: val, isSet: true}
}

func (v NullableGuami) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuami) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


