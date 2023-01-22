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

// UeContextInSmfDataSubFilter1 UE Context In Smf Data Subscription Filter.
type UeContextInSmfDataSubFilter1 struct {
	DnnList []string `json:"dnnList,omitempty"`
	SnssaiList []Snssai `json:"snssaiList,omitempty"`
	EmergencyInd *bool `json:"emergencyInd,omitempty"`
}

// NewUeContextInSmfDataSubFilter1 instantiates a new UeContextInSmfDataSubFilter1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeContextInSmfDataSubFilter1() *UeContextInSmfDataSubFilter1 {
	this := UeContextInSmfDataSubFilter1{}
	var emergencyInd bool = false
	this.EmergencyInd = &emergencyInd
	return &this
}

// NewUeContextInSmfDataSubFilter1WithDefaults instantiates a new UeContextInSmfDataSubFilter1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeContextInSmfDataSubFilter1WithDefaults() *UeContextInSmfDataSubFilter1 {
	this := UeContextInSmfDataSubFilter1{}
	var emergencyInd bool = false
	this.EmergencyInd = &emergencyInd
	return &this
}

// GetDnnList returns the DnnList field value if set, zero value otherwise.
func (o *UeContextInSmfDataSubFilter1) GetDnnList() []string {
	if o == nil || isNil(o.DnnList) {
		var ret []string
		return ret
	}
	return o.DnnList
}

// GetDnnListOk returns a tuple with the DnnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInSmfDataSubFilter1) GetDnnListOk() ([]string, bool) {
	if o == nil || isNil(o.DnnList) {
    return nil, false
	}
	return o.DnnList, true
}

// HasDnnList returns a boolean if a field has been set.
func (o *UeContextInSmfDataSubFilter1) HasDnnList() bool {
	if o != nil && !isNil(o.DnnList) {
		return true
	}

	return false
}

// SetDnnList gets a reference to the given []string and assigns it to the DnnList field.
func (o *UeContextInSmfDataSubFilter1) SetDnnList(v []string) {
	o.DnnList = v
}

// GetSnssaiList returns the SnssaiList field value if set, zero value otherwise.
func (o *UeContextInSmfDataSubFilter1) GetSnssaiList() []Snssai {
	if o == nil || isNil(o.SnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInSmfDataSubFilter1) GetSnssaiListOk() ([]Snssai, bool) {
	if o == nil || isNil(o.SnssaiList) {
    return nil, false
	}
	return o.SnssaiList, true
}

// HasSnssaiList returns a boolean if a field has been set.
func (o *UeContextInSmfDataSubFilter1) HasSnssaiList() bool {
	if o != nil && !isNil(o.SnssaiList) {
		return true
	}

	return false
}

// SetSnssaiList gets a reference to the given []Snssai and assigns it to the SnssaiList field.
func (o *UeContextInSmfDataSubFilter1) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

// GetEmergencyInd returns the EmergencyInd field value if set, zero value otherwise.
func (o *UeContextInSmfDataSubFilter1) GetEmergencyInd() bool {
	if o == nil || isNil(o.EmergencyInd) {
		var ret bool
		return ret
	}
	return *o.EmergencyInd
}

// GetEmergencyIndOk returns a tuple with the EmergencyInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInSmfDataSubFilter1) GetEmergencyIndOk() (*bool, bool) {
	if o == nil || isNil(o.EmergencyInd) {
    return nil, false
	}
	return o.EmergencyInd, true
}

// HasEmergencyInd returns a boolean if a field has been set.
func (o *UeContextInSmfDataSubFilter1) HasEmergencyInd() bool {
	if o != nil && !isNil(o.EmergencyInd) {
		return true
	}

	return false
}

// SetEmergencyInd gets a reference to the given bool and assigns it to the EmergencyInd field.
func (o *UeContextInSmfDataSubFilter1) SetEmergencyInd(v bool) {
	o.EmergencyInd = &v
}

func (o UeContextInSmfDataSubFilter1) MarshalJSON() ([]byte, error) {
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

type NullableUeContextInSmfDataSubFilter1 struct {
	value *UeContextInSmfDataSubFilter1
	isSet bool
}

func (v NullableUeContextInSmfDataSubFilter1) Get() *UeContextInSmfDataSubFilter1 {
	return v.value
}

func (v *NullableUeContextInSmfDataSubFilter1) Set(val *UeContextInSmfDataSubFilter1) {
	v.value = val
	v.isSet = true
}

func (v NullableUeContextInSmfDataSubFilter1) IsSet() bool {
	return v.isSet
}

func (v *NullableUeContextInSmfDataSubFilter1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeContextInSmfDataSubFilter1(val *UeContextInSmfDataSubFilter1) *NullableUeContextInSmfDataSubFilter1 {
	return &NullableUeContextInSmfDataSubFilter1{value: val, isSet: true}
}

func (v NullableUeContextInSmfDataSubFilter1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeContextInSmfDataSubFilter1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


