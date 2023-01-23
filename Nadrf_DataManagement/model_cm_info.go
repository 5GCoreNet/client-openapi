/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// CmInfo Represents the connection management state of a UE for an access type
type CmInfo struct {
	CmState CmState `json:"cmState"`
	AccessType AccessType `json:"accessType"`
}

// NewCmInfo instantiates a new CmInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmInfo(cmState CmState, accessType AccessType) *CmInfo {
	this := CmInfo{}
	this.CmState = cmState
	this.AccessType = accessType
	return &this
}

// NewCmInfoWithDefaults instantiates a new CmInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmInfoWithDefaults() *CmInfo {
	this := CmInfo{}
	return &this
}

// GetCmState returns the CmState field value
func (o *CmInfo) GetCmState() CmState {
	if o == nil {
		var ret CmState
		return ret
	}

	return o.CmState
}

// GetCmStateOk returns a tuple with the CmState field value
// and a boolean to check if the value has been set.
func (o *CmInfo) GetCmStateOk() (*CmState, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CmState, true
}

// SetCmState sets field value
func (o *CmInfo) SetCmState(v CmState) {
	o.CmState = v
}

// GetAccessType returns the AccessType field value
func (o *CmInfo) GetAccessType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *CmInfo) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *CmInfo) SetAccessType(v AccessType) {
	o.AccessType = v
}

func (o CmInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cmState"] = o.CmState
	}
	if true {
		toSerialize["accessType"] = o.AccessType
	}
	return json.Marshal(toSerialize)
}

type NullableCmInfo struct {
	value *CmInfo
	isSet bool
}

func (v NullableCmInfo) Get() *CmInfo {
	return v.value
}

func (v *NullableCmInfo) Set(val *CmInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCmInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCmInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmInfo(val *CmInfo) *NullableCmInfo {
	return &NullableCmInfo{value: val, isSet: true}
}

func (v NullableCmInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


