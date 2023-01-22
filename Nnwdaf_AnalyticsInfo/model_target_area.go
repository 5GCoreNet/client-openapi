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

// TargetArea TA list or TAI range list or any TA
type TargetArea struct {
	TaList []Tai `json:"taList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	AnyTa *bool `json:"anyTa,omitempty"`
}

// NewTargetArea instantiates a new TargetArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetArea() *TargetArea {
	this := TargetArea{}
	var anyTa bool = false
	this.AnyTa = &anyTa
	return &this
}

// NewTargetAreaWithDefaults instantiates a new TargetArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetAreaWithDefaults() *TargetArea {
	this := TargetArea{}
	var anyTa bool = false
	this.AnyTa = &anyTa
	return &this
}

// GetTaList returns the TaList field value if set, zero value otherwise.
func (o *TargetArea) GetTaList() []Tai {
	if o == nil || isNil(o.TaList) {
		var ret []Tai
		return ret
	}
	return o.TaList
}

// GetTaListOk returns a tuple with the TaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetArea) GetTaListOk() ([]Tai, bool) {
	if o == nil || isNil(o.TaList) {
    return nil, false
	}
	return o.TaList, true
}

// HasTaList returns a boolean if a field has been set.
func (o *TargetArea) HasTaList() bool {
	if o != nil && !isNil(o.TaList) {
		return true
	}

	return false
}

// SetTaList gets a reference to the given []Tai and assigns it to the TaList field.
func (o *TargetArea) SetTaList(v []Tai) {
	o.TaList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *TargetArea) GetTaiRangeList() []TaiRange {
	if o == nil || isNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetArea) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || isNil(o.TaiRangeList) {
    return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *TargetArea) HasTaiRangeList() bool {
	if o != nil && !isNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *TargetArea) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetAnyTa returns the AnyTa field value if set, zero value otherwise.
func (o *TargetArea) GetAnyTa() bool {
	if o == nil || isNil(o.AnyTa) {
		var ret bool
		return ret
	}
	return *o.AnyTa
}

// GetAnyTaOk returns a tuple with the AnyTa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetArea) GetAnyTaOk() (*bool, bool) {
	if o == nil || isNil(o.AnyTa) {
    return nil, false
	}
	return o.AnyTa, true
}

// HasAnyTa returns a boolean if a field has been set.
func (o *TargetArea) HasAnyTa() bool {
	if o != nil && !isNil(o.AnyTa) {
		return true
	}

	return false
}

// SetAnyTa gets a reference to the given bool and assigns it to the AnyTa field.
func (o *TargetArea) SetAnyTa(v bool) {
	o.AnyTa = &v
}

func (o TargetArea) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TaList) {
		toSerialize["taList"] = o.TaList
	}
	if !isNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !isNil(o.AnyTa) {
		toSerialize["anyTa"] = o.AnyTa
	}
	return json.Marshal(toSerialize)
}

type NullableTargetArea struct {
	value *TargetArea
	isSet bool
}

func (v NullableTargetArea) Get() *TargetArea {
	return v.value
}

func (v *NullableTargetArea) Set(val *TargetArea) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetArea) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetArea(val *TargetArea) *NullableTargetArea {
	return &NullableTargetArea{value: val, isSet: true}
}

func (v NullableTargetArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


