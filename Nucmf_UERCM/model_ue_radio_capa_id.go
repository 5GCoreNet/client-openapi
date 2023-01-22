/*
Nucmf_UECapabilityManagement

Nucmf_UECapabilityManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nucmf_UERCM

import (
	"encoding/json"
)

// UeRadioCapaId UE Radio Capability ID
type UeRadioCapaId struct {
	// string with format 'bytes' as defined in OpenAPI
	PlmnAssiUeRadioCapId *string `json:"plmnAssiUeRadioCapId,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	ManAssiUeRadioCapId *string `json:"manAssiUeRadioCapId,omitempty"`
}

// NewUeRadioCapaId instantiates a new UeRadioCapaId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeRadioCapaId() *UeRadioCapaId {
	this := UeRadioCapaId{}
	return &this
}

// NewUeRadioCapaIdWithDefaults instantiates a new UeRadioCapaId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeRadioCapaIdWithDefaults() *UeRadioCapaId {
	this := UeRadioCapaId{}
	return &this
}

// GetPlmnAssiUeRadioCapId returns the PlmnAssiUeRadioCapId field value if set, zero value otherwise.
func (o *UeRadioCapaId) GetPlmnAssiUeRadioCapId() string {
	if o == nil || isNil(o.PlmnAssiUeRadioCapId) {
		var ret string
		return ret
	}
	return *o.PlmnAssiUeRadioCapId
}

// GetPlmnAssiUeRadioCapIdOk returns a tuple with the PlmnAssiUeRadioCapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeRadioCapaId) GetPlmnAssiUeRadioCapIdOk() (*string, bool) {
	if o == nil || isNil(o.PlmnAssiUeRadioCapId) {
    return nil, false
	}
	return o.PlmnAssiUeRadioCapId, true
}

// HasPlmnAssiUeRadioCapId returns a boolean if a field has been set.
func (o *UeRadioCapaId) HasPlmnAssiUeRadioCapId() bool {
	if o != nil && !isNil(o.PlmnAssiUeRadioCapId) {
		return true
	}

	return false
}

// SetPlmnAssiUeRadioCapId gets a reference to the given string and assigns it to the PlmnAssiUeRadioCapId field.
func (o *UeRadioCapaId) SetPlmnAssiUeRadioCapId(v string) {
	o.PlmnAssiUeRadioCapId = &v
}

// GetManAssiUeRadioCapId returns the ManAssiUeRadioCapId field value if set, zero value otherwise.
func (o *UeRadioCapaId) GetManAssiUeRadioCapId() string {
	if o == nil || isNil(o.ManAssiUeRadioCapId) {
		var ret string
		return ret
	}
	return *o.ManAssiUeRadioCapId
}

// GetManAssiUeRadioCapIdOk returns a tuple with the ManAssiUeRadioCapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeRadioCapaId) GetManAssiUeRadioCapIdOk() (*string, bool) {
	if o == nil || isNil(o.ManAssiUeRadioCapId) {
    return nil, false
	}
	return o.ManAssiUeRadioCapId, true
}

// HasManAssiUeRadioCapId returns a boolean if a field has been set.
func (o *UeRadioCapaId) HasManAssiUeRadioCapId() bool {
	if o != nil && !isNil(o.ManAssiUeRadioCapId) {
		return true
	}

	return false
}

// SetManAssiUeRadioCapId gets a reference to the given string and assigns it to the ManAssiUeRadioCapId field.
func (o *UeRadioCapaId) SetManAssiUeRadioCapId(v string) {
	o.ManAssiUeRadioCapId = &v
}

func (o UeRadioCapaId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PlmnAssiUeRadioCapId) {
		toSerialize["plmnAssiUeRadioCapId"] = o.PlmnAssiUeRadioCapId
	}
	if !isNil(o.ManAssiUeRadioCapId) {
		toSerialize["manAssiUeRadioCapId"] = o.ManAssiUeRadioCapId
	}
	return json.Marshal(toSerialize)
}

type NullableUeRadioCapaId struct {
	value *UeRadioCapaId
	isSet bool
}

func (v NullableUeRadioCapaId) Get() *UeRadioCapaId {
	return v.value
}

func (v *NullableUeRadioCapaId) Set(val *UeRadioCapaId) {
	v.value = val
	v.isSet = true
}

func (v NullableUeRadioCapaId) IsSet() bool {
	return v.isSet
}

func (v *NullableUeRadioCapaId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeRadioCapaId(val *UeRadioCapaId) *NullableUeRadioCapaId {
	return &NullableUeRadioCapaId{value: val, isSet: true}
}

func (v NullableUeRadioCapaId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeRadioCapaId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


