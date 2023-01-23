/*
GBA BSF Nbsp_GBA Service

GBA BSF Nbsp_GBA Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsp_GBA

import (
	"encoding/json"
)

// FlagsItem Data item in a Flags array list
type FlagsItem struct {
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	Flag int32 `json:"flag"`
}

// NewFlagsItem instantiates a new FlagsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlagsItem(flag int32) *FlagsItem {
	this := FlagsItem{}
	this.Flag = flag
	return &this
}

// NewFlagsItemWithDefaults instantiates a new FlagsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlagsItemWithDefaults() *FlagsItem {
	this := FlagsItem{}
	return &this
}

// GetFlag returns the Flag field value
func (o *FlagsItem) GetFlag() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flag
}

// GetFlagOk returns a tuple with the Flag field value
// and a boolean to check if the value has been set.
func (o *FlagsItem) GetFlagOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Flag, true
}

// SetFlag sets field value
func (o *FlagsItem) SetFlag(v int32) {
	o.Flag = v
}

func (o FlagsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["flag"] = o.Flag
	}
	return json.Marshal(toSerialize)
}

type NullableFlagsItem struct {
	value *FlagsItem
	isSet bool
}

func (v NullableFlagsItem) Get() *FlagsItem {
	return v.value
}

func (v *NullableFlagsItem) Set(val *FlagsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFlagsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFlagsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlagsItem(val *FlagsItem) *NullableFlagsItem {
	return &NullableFlagsItem{value: val, isSet: true}
}

func (v NullableFlagsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlagsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


