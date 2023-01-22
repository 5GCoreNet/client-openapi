/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// UpdatedItem Identifies a fragment of a resource.
type UpdatedItem struct {
	// Identifies a fragment (subset of resource data) of a given resource.
	Item string `json:"item"`
	Value interface{} `json:"value"`
}

// NewUpdatedItem instantiates a new UpdatedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatedItem(item string, value interface{}) *UpdatedItem {
	this := UpdatedItem{}
	this.Item = item
	this.Value = value
	return &this
}

// NewUpdatedItemWithDefaults instantiates a new UpdatedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatedItemWithDefaults() *UpdatedItem {
	this := UpdatedItem{}
	return &this
}

// GetItem returns the Item field value
func (o *UpdatedItem) GetItem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *UpdatedItem) GetItemOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *UpdatedItem) SetItem(v string) {
	o.Item = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *UpdatedItem) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatedItem) GetValueOk() (*interface{}, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UpdatedItem) SetValue(v interface{}) {
	o.Value = v
}

func (o UpdatedItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatedItem struct {
	value *UpdatedItem
	isSet bool
}

func (v NullableUpdatedItem) Get() *UpdatedItem {
	return v.value
}

func (v *NullableUpdatedItem) Set(val *UpdatedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatedItem(val *UpdatedItem) *NullableUpdatedItem {
	return &NullableUpdatedItem{value: val, isSet: true}
}

func (v NullableUpdatedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


