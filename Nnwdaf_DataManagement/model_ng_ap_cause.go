/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// NgApCause Represents the NGAP cause.
type NgApCause struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Group int32 `json:"group"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Value int32 `json:"value"`
}

// NewNgApCause instantiates a new NgApCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgApCause(group int32, value int32) *NgApCause {
	this := NgApCause{}
	this.Group = group
	this.Value = value
	return &this
}

// NewNgApCauseWithDefaults instantiates a new NgApCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgApCauseWithDefaults() *NgApCause {
	this := NgApCause{}
	return &this
}

// GetGroup returns the Group field value
func (o *NgApCause) GetGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *NgApCause) GetGroupOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *NgApCause) SetGroup(v int32) {
	o.Group = v
}

// GetValue returns the Value field value
func (o *NgApCause) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *NgApCause) GetValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *NgApCause) SetValue(v int32) {
	o.Value = v
}

func (o NgApCause) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableNgApCause struct {
	value *NgApCause
	isSet bool
}

func (v NullableNgApCause) Get() *NgApCause {
	return v.value
}

func (v *NullableNgApCause) Set(val *NgApCause) {
	v.value = val
	v.isSet = true
}

func (v NullableNgApCause) IsSet() bool {
	return v.isSet
}

func (v *NullableNgApCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgApCause(val *NgApCause) *NullableNgApCause {
	return &NullableNgApCause{value: val, isSet: true}
}

func (v NullableNgApCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgApCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


