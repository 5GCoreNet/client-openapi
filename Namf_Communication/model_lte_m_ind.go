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

// LteMInd LTE-M Indication.
type LteMInd struct {
	LteCatMInd bool `json:"lteCatMInd"`
}

// NewLteMInd instantiates a new LteMInd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLteMInd(lteCatMInd bool) *LteMInd {
	this := LteMInd{}
	this.LteCatMInd = lteCatMInd
	return &this
}

// NewLteMIndWithDefaults instantiates a new LteMInd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLteMIndWithDefaults() *LteMInd {
	this := LteMInd{}
	return &this
}

// GetLteCatMInd returns the LteCatMInd field value
func (o *LteMInd) GetLteCatMInd() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LteCatMInd
}

// GetLteCatMIndOk returns a tuple with the LteCatMInd field value
// and a boolean to check if the value has been set.
func (o *LteMInd) GetLteCatMIndOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LteCatMInd, true
}

// SetLteCatMInd sets field value
func (o *LteMInd) SetLteCatMInd(v bool) {
	o.LteCatMInd = v
}

func (o LteMInd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["lteCatMInd"] = o.LteCatMInd
	}
	return json.Marshal(toSerialize)
}

type NullableLteMInd struct {
	value *LteMInd
	isSet bool
}

func (v NullableLteMInd) Get() *LteMInd {
	return v.value
}

func (v *NullableLteMInd) Set(val *LteMInd) {
	v.value = val
	v.isSet = true
}

func (v NullableLteMInd) IsSet() bool {
	return v.isSet
}

func (v *NullableLteMInd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLteMInd(val *LteMInd) *NullableLteMInd {
	return &NullableLteMInd{value: val, isSet: true}
}

func (v NullableLteMInd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLteMInd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


