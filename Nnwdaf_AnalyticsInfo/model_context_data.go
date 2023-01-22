/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// ContextData Contains context information related to analytics subscriptions corresponding with one or  more context identifiers. 
type ContextData struct {
	// List of items that contain context information corresponding with a context identifier. 
	ContextElems []ContextElement `json:"contextElems"`
}

// NewContextData instantiates a new ContextData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextData(contextElems []ContextElement) *ContextData {
	this := ContextData{}
	this.ContextElems = contextElems
	return &this
}

// NewContextDataWithDefaults instantiates a new ContextData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextDataWithDefaults() *ContextData {
	this := ContextData{}
	return &this
}

// GetContextElems returns the ContextElems field value
func (o *ContextData) GetContextElems() []ContextElement {
	if o == nil {
		var ret []ContextElement
		return ret
	}

	return o.ContextElems
}

// GetContextElemsOk returns a tuple with the ContextElems field value
// and a boolean to check if the value has been set.
func (o *ContextData) GetContextElemsOk() ([]ContextElement, bool) {
	if o == nil {
    return nil, false
	}
	return o.ContextElems, true
}

// SetContextElems sets field value
func (o *ContextData) SetContextElems(v []ContextElement) {
	o.ContextElems = v
}

func (o ContextData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contextElems"] = o.ContextElems
	}
	return json.Marshal(toSerialize)
}

type NullableContextData struct {
	value *ContextData
	isSet bool
}

func (v NullableContextData) Get() *ContextData {
	return v.value
}

func (v *NullableContextData) Set(val *ContextData) {
	v.value = val
	v.isSet = true
}

func (v NullableContextData) IsSet() bool {
	return v.isSet
}

func (v *NullableContextData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextData(val *ContextData) *NullableContextData {
	return &NullableContextData{value: val, isSet: true}
}

func (v NullableContextData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

