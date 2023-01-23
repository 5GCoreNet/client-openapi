/*
Npcf_EventExposure

PCF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_EventExposure

import (
	"encoding/json"
)

// SnssaiDnnCombination Represents a combination of S-NSSAI and DNN(s).
type SnssaiDnnCombination struct {
	Snssai *Snssai `json:"snssai,omitempty"`
	Dnns []string `json:"dnns,omitempty"`
}

// NewSnssaiDnnCombination instantiates a new SnssaiDnnCombination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiDnnCombination() *SnssaiDnnCombination {
	this := SnssaiDnnCombination{}
	return &this
}

// NewSnssaiDnnCombinationWithDefaults instantiates a new SnssaiDnnCombination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiDnnCombinationWithDefaults() *SnssaiDnnCombination {
	this := SnssaiDnnCombination{}
	return &this
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *SnssaiDnnCombination) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnssaiDnnCombination) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
    return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *SnssaiDnnCombination) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *SnssaiDnnCombination) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetDnns returns the Dnns field value if set, zero value otherwise.
func (o *SnssaiDnnCombination) GetDnns() []string {
	if o == nil || isNil(o.Dnns) {
		var ret []string
		return ret
	}
	return o.Dnns
}

// GetDnnsOk returns a tuple with the Dnns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnssaiDnnCombination) GetDnnsOk() ([]string, bool) {
	if o == nil || isNil(o.Dnns) {
    return nil, false
	}
	return o.Dnns, true
}

// HasDnns returns a boolean if a field has been set.
func (o *SnssaiDnnCombination) HasDnns() bool {
	if o != nil && !isNil(o.Dnns) {
		return true
	}

	return false
}

// SetDnns gets a reference to the given []string and assigns it to the Dnns field.
func (o *SnssaiDnnCombination) SetDnns(v []string) {
	o.Dnns = v
}

func (o SnssaiDnnCombination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !isNil(o.Dnns) {
		toSerialize["dnns"] = o.Dnns
	}
	return json.Marshal(toSerialize)
}

type NullableSnssaiDnnCombination struct {
	value *SnssaiDnnCombination
	isSet bool
}

func (v NullableSnssaiDnnCombination) Get() *SnssaiDnnCombination {
	return v.value
}

func (v *NullableSnssaiDnnCombination) Set(val *SnssaiDnnCombination) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiDnnCombination) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiDnnCombination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiDnnCombination(val *SnssaiDnnCombination) *NullableSnssaiDnnCombination {
	return &NullableSnssaiDnnCombination{value: val, isSet: true}
}

func (v NullableSnssaiDnnCombination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiDnnCombination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


