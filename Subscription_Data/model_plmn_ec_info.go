/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
)

// PlmnEcInfo struct for PlmnEcInfo
type PlmnEcInfo struct {
	PlmnId PlmnId `json:"plmnId"`
	EcRestrictionDataWb *EcRestrictionDataWb `json:"ecRestrictionDataWb,omitempty"`
	EcRestrictionDataNb *bool `json:"ecRestrictionDataNb,omitempty"`
}

// NewPlmnEcInfo instantiates a new PlmnEcInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnEcInfo(plmnId PlmnId) *PlmnEcInfo {
	this := PlmnEcInfo{}
	this.PlmnId = plmnId
	var ecRestrictionDataNb bool = false
	this.EcRestrictionDataNb = &ecRestrictionDataNb
	return &this
}

// NewPlmnEcInfoWithDefaults instantiates a new PlmnEcInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnEcInfoWithDefaults() *PlmnEcInfo {
	this := PlmnEcInfo{}
	var ecRestrictionDataNb bool = false
	this.EcRestrictionDataNb = &ecRestrictionDataNb
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *PlmnEcInfo) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *PlmnEcInfo) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *PlmnEcInfo) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetEcRestrictionDataWb returns the EcRestrictionDataWb field value if set, zero value otherwise.
func (o *PlmnEcInfo) GetEcRestrictionDataWb() EcRestrictionDataWb {
	if o == nil || isNil(o.EcRestrictionDataWb) {
		var ret EcRestrictionDataWb
		return ret
	}
	return *o.EcRestrictionDataWb
}

// GetEcRestrictionDataWbOk returns a tuple with the EcRestrictionDataWb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnEcInfo) GetEcRestrictionDataWbOk() (*EcRestrictionDataWb, bool) {
	if o == nil || isNil(o.EcRestrictionDataWb) {
    return nil, false
	}
	return o.EcRestrictionDataWb, true
}

// HasEcRestrictionDataWb returns a boolean if a field has been set.
func (o *PlmnEcInfo) HasEcRestrictionDataWb() bool {
	if o != nil && !isNil(o.EcRestrictionDataWb) {
		return true
	}

	return false
}

// SetEcRestrictionDataWb gets a reference to the given EcRestrictionDataWb and assigns it to the EcRestrictionDataWb field.
func (o *PlmnEcInfo) SetEcRestrictionDataWb(v EcRestrictionDataWb) {
	o.EcRestrictionDataWb = &v
}

// GetEcRestrictionDataNb returns the EcRestrictionDataNb field value if set, zero value otherwise.
func (o *PlmnEcInfo) GetEcRestrictionDataNb() bool {
	if o == nil || isNil(o.EcRestrictionDataNb) {
		var ret bool
		return ret
	}
	return *o.EcRestrictionDataNb
}

// GetEcRestrictionDataNbOk returns a tuple with the EcRestrictionDataNb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnEcInfo) GetEcRestrictionDataNbOk() (*bool, bool) {
	if o == nil || isNil(o.EcRestrictionDataNb) {
    return nil, false
	}
	return o.EcRestrictionDataNb, true
}

// HasEcRestrictionDataNb returns a boolean if a field has been set.
func (o *PlmnEcInfo) HasEcRestrictionDataNb() bool {
	if o != nil && !isNil(o.EcRestrictionDataNb) {
		return true
	}

	return false
}

// SetEcRestrictionDataNb gets a reference to the given bool and assigns it to the EcRestrictionDataNb field.
func (o *PlmnEcInfo) SetEcRestrictionDataNb(v bool) {
	o.EcRestrictionDataNb = &v
}

func (o PlmnEcInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !isNil(o.EcRestrictionDataWb) {
		toSerialize["ecRestrictionDataWb"] = o.EcRestrictionDataWb
	}
	if !isNil(o.EcRestrictionDataNb) {
		toSerialize["ecRestrictionDataNb"] = o.EcRestrictionDataNb
	}
	return json.Marshal(toSerialize)
}

type NullablePlmnEcInfo struct {
	value *PlmnEcInfo
	isSet bool
}

func (v NullablePlmnEcInfo) Get() *PlmnEcInfo {
	return v.value
}

func (v *NullablePlmnEcInfo) Set(val *PlmnEcInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnEcInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnEcInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnEcInfo(val *PlmnEcInfo) *NullablePlmnEcInfo {
	return &NullablePlmnEcInfo{value: val, isSet: true}
}

func (v NullablePlmnEcInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnEcInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


