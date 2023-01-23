/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// NgRanTargetId Indicates a NG RAN as target of the handover
type NgRanTargetId struct {
	RanNodeId GlobalRanNodeId `json:"ranNodeId"`
	Tai Tai `json:"tai"`
}

// NewNgRanTargetId instantiates a new NgRanTargetId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgRanTargetId(ranNodeId GlobalRanNodeId, tai Tai) *NgRanTargetId {
	this := NgRanTargetId{}
	this.RanNodeId = ranNodeId
	this.Tai = tai
	return &this
}

// NewNgRanTargetIdWithDefaults instantiates a new NgRanTargetId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgRanTargetIdWithDefaults() *NgRanTargetId {
	this := NgRanTargetId{}
	return &this
}

// GetRanNodeId returns the RanNodeId field value
func (o *NgRanTargetId) GetRanNodeId() GlobalRanNodeId {
	if o == nil {
		var ret GlobalRanNodeId
		return ret
	}

	return o.RanNodeId
}

// GetRanNodeIdOk returns a tuple with the RanNodeId field value
// and a boolean to check if the value has been set.
func (o *NgRanTargetId) GetRanNodeIdOk() (*GlobalRanNodeId, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RanNodeId, true
}

// SetRanNodeId sets field value
func (o *NgRanTargetId) SetRanNodeId(v GlobalRanNodeId) {
	o.RanNodeId = v
}

// GetTai returns the Tai field value
func (o *NgRanTargetId) GetTai() Tai {
	if o == nil {
		var ret Tai
		return ret
	}

	return o.Tai
}

// GetTaiOk returns a tuple with the Tai field value
// and a boolean to check if the value has been set.
func (o *NgRanTargetId) GetTaiOk() (*Tai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tai, true
}

// SetTai sets field value
func (o *NgRanTargetId) SetTai(v Tai) {
	o.Tai = v
}

func (o NgRanTargetId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ranNodeId"] = o.RanNodeId
	}
	if true {
		toSerialize["tai"] = o.Tai
	}
	return json.Marshal(toSerialize)
}

type NullableNgRanTargetId struct {
	value *NgRanTargetId
	isSet bool
}

func (v NullableNgRanTargetId) Get() *NgRanTargetId {
	return v.value
}

func (v *NullableNgRanTargetId) Set(val *NgRanTargetId) {
	v.value = val
	v.isSet = true
}

func (v NullableNgRanTargetId) IsSet() bool {
	return v.isSet
}

func (v *NullableNgRanTargetId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgRanTargetId(val *NgRanTargetId) *NullableNgRanTargetId {
	return &NullableNgRanTargetId{value: val, isSet: true}
}

func (v NullableNgRanTargetId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgRanTargetId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


