/*
Npcf_AMPolicyControl

Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_AMPolicyControl

import (
	"encoding/json"
)

// UeSliceMbr Contains a UE-Slice-MBR and the related information.
type UeSliceMbr struct {
	// Contains the MBR for uplink and the MBR for downlink.
	SliceMbr map[string]SliceMbr `json:"sliceMbr"`
	ServingSnssai Snssai `json:"servingSnssai"`
	MappedHomeSnssai *Snssai `json:"mappedHomeSnssai,omitempty"`
}

// NewUeSliceMbr instantiates a new UeSliceMbr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeSliceMbr(sliceMbr map[string]SliceMbr, servingSnssai Snssai) *UeSliceMbr {
	this := UeSliceMbr{}
	this.SliceMbr = sliceMbr
	this.ServingSnssai = servingSnssai
	return &this
}

// NewUeSliceMbrWithDefaults instantiates a new UeSliceMbr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeSliceMbrWithDefaults() *UeSliceMbr {
	this := UeSliceMbr{}
	return &this
}

// GetSliceMbr returns the SliceMbr field value
func (o *UeSliceMbr) GetSliceMbr() map[string]SliceMbr {
	if o == nil {
		var ret map[string]SliceMbr
		return ret
	}

	return o.SliceMbr
}

// GetSliceMbrOk returns a tuple with the SliceMbr field value
// and a boolean to check if the value has been set.
func (o *UeSliceMbr) GetSliceMbrOk() (*map[string]SliceMbr, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SliceMbr, true
}

// SetSliceMbr sets field value
func (o *UeSliceMbr) SetSliceMbr(v map[string]SliceMbr) {
	o.SliceMbr = v
}

// GetServingSnssai returns the ServingSnssai field value
func (o *UeSliceMbr) GetServingSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.ServingSnssai
}

// GetServingSnssaiOk returns a tuple with the ServingSnssai field value
// and a boolean to check if the value has been set.
func (o *UeSliceMbr) GetServingSnssaiOk() (*Snssai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServingSnssai, true
}

// SetServingSnssai sets field value
func (o *UeSliceMbr) SetServingSnssai(v Snssai) {
	o.ServingSnssai = v
}

// GetMappedHomeSnssai returns the MappedHomeSnssai field value if set, zero value otherwise.
func (o *UeSliceMbr) GetMappedHomeSnssai() Snssai {
	if o == nil || isNil(o.MappedHomeSnssai) {
		var ret Snssai
		return ret
	}
	return *o.MappedHomeSnssai
}

// GetMappedHomeSnssaiOk returns a tuple with the MappedHomeSnssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeSliceMbr) GetMappedHomeSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.MappedHomeSnssai) {
    return nil, false
	}
	return o.MappedHomeSnssai, true
}

// HasMappedHomeSnssai returns a boolean if a field has been set.
func (o *UeSliceMbr) HasMappedHomeSnssai() bool {
	if o != nil && !isNil(o.MappedHomeSnssai) {
		return true
	}

	return false
}

// SetMappedHomeSnssai gets a reference to the given Snssai and assigns it to the MappedHomeSnssai field.
func (o *UeSliceMbr) SetMappedHomeSnssai(v Snssai) {
	o.MappedHomeSnssai = &v
}

func (o UeSliceMbr) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sliceMbr"] = o.SliceMbr
	}
	if true {
		toSerialize["servingSnssai"] = o.ServingSnssai
	}
	if !isNil(o.MappedHomeSnssai) {
		toSerialize["mappedHomeSnssai"] = o.MappedHomeSnssai
	}
	return json.Marshal(toSerialize)
}

type NullableUeSliceMbr struct {
	value *UeSliceMbr
	isSet bool
}

func (v NullableUeSliceMbr) Get() *UeSliceMbr {
	return v.value
}

func (v *NullableUeSliceMbr) Set(val *UeSliceMbr) {
	v.value = val
	v.isSet = true
}

func (v NullableUeSliceMbr) IsSet() bool {
	return v.isSet
}

func (v *NullableUeSliceMbr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeSliceMbr(val *UeSliceMbr) *NullableUeSliceMbr {
	return &NullableUeSliceMbr{value: val, isSet: true}
}

func (v NullableUeSliceMbr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeSliceMbr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


