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

// UeContextInSmfDataSubFilter UE Context In Smf Data Subscription Filter.
type UeContextInSmfDataSubFilter struct {
	DnnList []string `json:"dnnList,omitempty"`
	SnssaiList []Snssai `json:"snssaiList,omitempty"`
	EmergencyInd *bool `json:"emergencyInd,omitempty"`
}

// NewUeContextInSmfDataSubFilter instantiates a new UeContextInSmfDataSubFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeContextInSmfDataSubFilter() *UeContextInSmfDataSubFilter {
	this := UeContextInSmfDataSubFilter{}
	var emergencyInd bool = false
	this.EmergencyInd = &emergencyInd
	return &this
}

// NewUeContextInSmfDataSubFilterWithDefaults instantiates a new UeContextInSmfDataSubFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeContextInSmfDataSubFilterWithDefaults() *UeContextInSmfDataSubFilter {
	this := UeContextInSmfDataSubFilter{}
	var emergencyInd bool = false
	this.EmergencyInd = &emergencyInd
	return &this
}

// GetDnnList returns the DnnList field value if set, zero value otherwise.
func (o *UeContextInSmfDataSubFilter) GetDnnList() []string {
	if o == nil || isNil(o.DnnList) {
		var ret []string
		return ret
	}
	return o.DnnList
}

// GetDnnListOk returns a tuple with the DnnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInSmfDataSubFilter) GetDnnListOk() ([]string, bool) {
	if o == nil || isNil(o.DnnList) {
    return nil, false
	}
	return o.DnnList, true
}

// HasDnnList returns a boolean if a field has been set.
func (o *UeContextInSmfDataSubFilter) HasDnnList() bool {
	if o != nil && !isNil(o.DnnList) {
		return true
	}

	return false
}

// SetDnnList gets a reference to the given []string and assigns it to the DnnList field.
func (o *UeContextInSmfDataSubFilter) SetDnnList(v []string) {
	o.DnnList = v
}

// GetSnssaiList returns the SnssaiList field value if set, zero value otherwise.
func (o *UeContextInSmfDataSubFilter) GetSnssaiList() []Snssai {
	if o == nil || isNil(o.SnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInSmfDataSubFilter) GetSnssaiListOk() ([]Snssai, bool) {
	if o == nil || isNil(o.SnssaiList) {
    return nil, false
	}
	return o.SnssaiList, true
}

// HasSnssaiList returns a boolean if a field has been set.
func (o *UeContextInSmfDataSubFilter) HasSnssaiList() bool {
	if o != nil && !isNil(o.SnssaiList) {
		return true
	}

	return false
}

// SetSnssaiList gets a reference to the given []Snssai and assigns it to the SnssaiList field.
func (o *UeContextInSmfDataSubFilter) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

// GetEmergencyInd returns the EmergencyInd field value if set, zero value otherwise.
func (o *UeContextInSmfDataSubFilter) GetEmergencyInd() bool {
	if o == nil || isNil(o.EmergencyInd) {
		var ret bool
		return ret
	}
	return *o.EmergencyInd
}

// GetEmergencyIndOk returns a tuple with the EmergencyInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInSmfDataSubFilter) GetEmergencyIndOk() (*bool, bool) {
	if o == nil || isNil(o.EmergencyInd) {
    return nil, false
	}
	return o.EmergencyInd, true
}

// HasEmergencyInd returns a boolean if a field has been set.
func (o *UeContextInSmfDataSubFilter) HasEmergencyInd() bool {
	if o != nil && !isNil(o.EmergencyInd) {
		return true
	}

	return false
}

// SetEmergencyInd gets a reference to the given bool and assigns it to the EmergencyInd field.
func (o *UeContextInSmfDataSubFilter) SetEmergencyInd(v bool) {
	o.EmergencyInd = &v
}

func (o UeContextInSmfDataSubFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DnnList) {
		toSerialize["dnnList"] = o.DnnList
	}
	if !isNil(o.SnssaiList) {
		toSerialize["snssaiList"] = o.SnssaiList
	}
	if !isNil(o.EmergencyInd) {
		toSerialize["emergencyInd"] = o.EmergencyInd
	}
	return json.Marshal(toSerialize)
}

type NullableUeContextInSmfDataSubFilter struct {
	value *UeContextInSmfDataSubFilter
	isSet bool
}

func (v NullableUeContextInSmfDataSubFilter) Get() *UeContextInSmfDataSubFilter {
	return v.value
}

func (v *NullableUeContextInSmfDataSubFilter) Set(val *UeContextInSmfDataSubFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableUeContextInSmfDataSubFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableUeContextInSmfDataSubFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeContextInSmfDataSubFilter(val *UeContextInSmfDataSubFilter) *NullableUeContextInSmfDataSubFilter {
	return &NullableUeContextInSmfDataSubFilter{value: val, isSet: true}
}

func (v NullableUeContextInSmfDataSubFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeContextInSmfDataSubFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


