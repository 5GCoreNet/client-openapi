/*
AUSF API

AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nausf_UEAuthentication

import (
	"encoding/json"
)

// Av5gAka Contains Authentication Vector for method 5G AKA.
type Av5gAka struct {
	Rand string `json:"rand"`
	// Contains the HXRES*.
	HxresStar string `json:"hxresStar"`
	Autn string `json:"autn"`
}

// NewAv5gAka instantiates a new Av5gAka object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAv5gAka(rand string, hxresStar string, autn string) *Av5gAka {
	this := Av5gAka{}
	this.Rand = rand
	this.HxresStar = hxresStar
	this.Autn = autn
	return &this
}

// NewAv5gAkaWithDefaults instantiates a new Av5gAka object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAv5gAkaWithDefaults() *Av5gAka {
	this := Av5gAka{}
	return &this
}

// GetRand returns the Rand field value
func (o *Av5gAka) GetRand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rand
}

// GetRandOk returns a tuple with the Rand field value
// and a boolean to check if the value has been set.
func (o *Av5gAka) GetRandOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Rand, true
}

// SetRand sets field value
func (o *Av5gAka) SetRand(v string) {
	o.Rand = v
}

// GetHxresStar returns the HxresStar field value
func (o *Av5gAka) GetHxresStar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HxresStar
}

// GetHxresStarOk returns a tuple with the HxresStar field value
// and a boolean to check if the value has been set.
func (o *Av5gAka) GetHxresStarOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HxresStar, true
}

// SetHxresStar sets field value
func (o *Av5gAka) SetHxresStar(v string) {
	o.HxresStar = v
}

// GetAutn returns the Autn field value
func (o *Av5gAka) GetAutn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Autn
}

// GetAutnOk returns a tuple with the Autn field value
// and a boolean to check if the value has been set.
func (o *Av5gAka) GetAutnOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Autn, true
}

// SetAutn sets field value
func (o *Av5gAka) SetAutn(v string) {
	o.Autn = v
}

func (o Av5gAka) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rand"] = o.Rand
	}
	if true {
		toSerialize["hxresStar"] = o.HxresStar
	}
	if true {
		toSerialize["autn"] = o.Autn
	}
	return json.Marshal(toSerialize)
}

type NullableAv5gAka struct {
	value *Av5gAka
	isSet bool
}

func (v NullableAv5gAka) Get() *Av5gAka {
	return v.value
}

func (v *NullableAv5gAka) Set(val *Av5gAka) {
	v.value = val
	v.isSet = true
}

func (v NullableAv5gAka) IsSet() bool {
	return v.isSet
}

func (v *NullableAv5gAka) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAv5gAka(val *Av5gAka) *NullableAv5gAka {
	return &NullableAv5gAka{value: val, isSet: true}
}

func (v NullableAv5gAka) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAv5gAka) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


