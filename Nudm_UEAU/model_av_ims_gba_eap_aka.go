/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UEAU

import (
	"encoding/json"
)

// AvImsGbaEapAka struct for AvImsGbaEapAka
type AvImsGbaEapAka struct {
	AvType HssAvType `json:"avType"`
	Rand string `json:"rand"`
	Xres string `json:"xres"`
	Autn string `json:"autn"`
	Ck string `json:"ck"`
	Ik string `json:"ik"`
}

// NewAvImsGbaEapAka instantiates a new AvImsGbaEapAka object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvImsGbaEapAka(avType HssAvType, rand string, xres string, autn string, ck string, ik string) *AvImsGbaEapAka {
	this := AvImsGbaEapAka{}
	this.AvType = avType
	this.Rand = rand
	this.Xres = xres
	this.Autn = autn
	this.Ck = ck
	this.Ik = ik
	return &this
}

// NewAvImsGbaEapAkaWithDefaults instantiates a new AvImsGbaEapAka object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvImsGbaEapAkaWithDefaults() *AvImsGbaEapAka {
	this := AvImsGbaEapAka{}
	return &this
}

// GetAvType returns the AvType field value
func (o *AvImsGbaEapAka) GetAvType() HssAvType {
	if o == nil {
		var ret HssAvType
		return ret
	}

	return o.AvType
}

// GetAvTypeOk returns a tuple with the AvType field value
// and a boolean to check if the value has been set.
func (o *AvImsGbaEapAka) GetAvTypeOk() (*HssAvType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AvType, true
}

// SetAvType sets field value
func (o *AvImsGbaEapAka) SetAvType(v HssAvType) {
	o.AvType = v
}

// GetRand returns the Rand field value
func (o *AvImsGbaEapAka) GetRand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rand
}

// GetRandOk returns a tuple with the Rand field value
// and a boolean to check if the value has been set.
func (o *AvImsGbaEapAka) GetRandOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Rand, true
}

// SetRand sets field value
func (o *AvImsGbaEapAka) SetRand(v string) {
	o.Rand = v
}

// GetXres returns the Xres field value
func (o *AvImsGbaEapAka) GetXres() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Xres
}

// GetXresOk returns a tuple with the Xres field value
// and a boolean to check if the value has been set.
func (o *AvImsGbaEapAka) GetXresOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Xres, true
}

// SetXres sets field value
func (o *AvImsGbaEapAka) SetXres(v string) {
	o.Xres = v
}

// GetAutn returns the Autn field value
func (o *AvImsGbaEapAka) GetAutn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Autn
}

// GetAutnOk returns a tuple with the Autn field value
// and a boolean to check if the value has been set.
func (o *AvImsGbaEapAka) GetAutnOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Autn, true
}

// SetAutn sets field value
func (o *AvImsGbaEapAka) SetAutn(v string) {
	o.Autn = v
}

// GetCk returns the Ck field value
func (o *AvImsGbaEapAka) GetCk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ck
}

// GetCkOk returns a tuple with the Ck field value
// and a boolean to check if the value has been set.
func (o *AvImsGbaEapAka) GetCkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ck, true
}

// SetCk sets field value
func (o *AvImsGbaEapAka) SetCk(v string) {
	o.Ck = v
}

// GetIk returns the Ik field value
func (o *AvImsGbaEapAka) GetIk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ik
}

// GetIkOk returns a tuple with the Ik field value
// and a boolean to check if the value has been set.
func (o *AvImsGbaEapAka) GetIkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ik, true
}

// SetIk sets field value
func (o *AvImsGbaEapAka) SetIk(v string) {
	o.Ik = v
}

func (o AvImsGbaEapAka) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["avType"] = o.AvType
	}
	if true {
		toSerialize["rand"] = o.Rand
	}
	if true {
		toSerialize["xres"] = o.Xres
	}
	if true {
		toSerialize["autn"] = o.Autn
	}
	if true {
		toSerialize["ck"] = o.Ck
	}
	if true {
		toSerialize["ik"] = o.Ik
	}
	return json.Marshal(toSerialize)
}

type NullableAvImsGbaEapAka struct {
	value *AvImsGbaEapAka
	isSet bool
}

func (v NullableAvImsGbaEapAka) Get() *AvImsGbaEapAka {
	return v.value
}

func (v *NullableAvImsGbaEapAka) Set(val *AvImsGbaEapAka) {
	v.value = val
	v.isSet = true
}

func (v NullableAvImsGbaEapAka) IsSet() bool {
	return v.isSet
}

func (v *NullableAvImsGbaEapAka) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvImsGbaEapAka(val *AvImsGbaEapAka) *NullableAvImsGbaEapAka {
	return &NullableAvImsGbaEapAka{value: val, isSet: true}
}

func (v NullableAvImsGbaEapAka) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvImsGbaEapAka) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


