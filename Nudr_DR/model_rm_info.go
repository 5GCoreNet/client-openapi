/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// RmInfo Represents the registration state of a UE for an access type
type RmInfo struct {
	RmState RmState `json:"rmState"`
	AccessType AccessType `json:"accessType"`
}

// NewRmInfo instantiates a new RmInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRmInfo(rmState RmState, accessType AccessType) *RmInfo {
	this := RmInfo{}
	this.RmState = rmState
	this.AccessType = accessType
	return &this
}

// NewRmInfoWithDefaults instantiates a new RmInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRmInfoWithDefaults() *RmInfo {
	this := RmInfo{}
	return &this
}

// GetRmState returns the RmState field value
func (o *RmInfo) GetRmState() RmState {
	if o == nil {
		var ret RmState
		return ret
	}

	return o.RmState
}

// GetRmStateOk returns a tuple with the RmState field value
// and a boolean to check if the value has been set.
func (o *RmInfo) GetRmStateOk() (*RmState, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RmState, true
}

// SetRmState sets field value
func (o *RmInfo) SetRmState(v RmState) {
	o.RmState = v
}

// GetAccessType returns the AccessType field value
func (o *RmInfo) GetAccessType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *RmInfo) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *RmInfo) SetAccessType(v AccessType) {
	o.AccessType = v
}

func (o RmInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rmState"] = o.RmState
	}
	if true {
		toSerialize["accessType"] = o.AccessType
	}
	return json.Marshal(toSerialize)
}

type NullableRmInfo struct {
	value *RmInfo
	isSet bool
}

func (v NullableRmInfo) Get() *RmInfo {
	return v.value
}

func (v *NullableRmInfo) Set(val *RmInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRmInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRmInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmInfo(val *RmInfo) *NullableRmInfo {
	return &NullableRmInfo{value: val, isSet: true}
}

func (v NullableRmInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


