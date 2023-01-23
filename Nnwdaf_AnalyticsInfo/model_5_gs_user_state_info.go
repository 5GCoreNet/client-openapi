/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// Model5GsUserStateInfo Represents the 5GS User state of the UE for an access type
type Model5GsUserStateInfo struct {
	Var5gsUserState Model5GsUserState `json:"5gsUserState"`
	AccessType AccessType `json:"accessType"`
}

// NewModel5GsUserStateInfo instantiates a new Model5GsUserStateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel5GsUserStateInfo(var5gsUserState Model5GsUserState, accessType AccessType) *Model5GsUserStateInfo {
	this := Model5GsUserStateInfo{}
	this.Var5gsUserState = var5gsUserState
	this.AccessType = accessType
	return &this
}

// NewModel5GsUserStateInfoWithDefaults instantiates a new Model5GsUserStateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel5GsUserStateInfoWithDefaults() *Model5GsUserStateInfo {
	this := Model5GsUserStateInfo{}
	return &this
}

// GetVar5gsUserState returns the Var5gsUserState field value
func (o *Model5GsUserStateInfo) GetVar5gsUserState() Model5GsUserState {
	if o == nil {
		var ret Model5GsUserState
		return ret
	}

	return o.Var5gsUserState
}

// GetVar5gsUserStateOk returns a tuple with the Var5gsUserState field value
// and a boolean to check if the value has been set.
func (o *Model5GsUserStateInfo) GetVar5gsUserStateOk() (*Model5GsUserState, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Var5gsUserState, true
}

// SetVar5gsUserState sets field value
func (o *Model5GsUserStateInfo) SetVar5gsUserState(v Model5GsUserState) {
	o.Var5gsUserState = v
}

// GetAccessType returns the AccessType field value
func (o *Model5GsUserStateInfo) GetAccessType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *Model5GsUserStateInfo) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *Model5GsUserStateInfo) SetAccessType(v AccessType) {
	o.AccessType = v
}

func (o Model5GsUserStateInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["5gsUserState"] = o.Var5gsUserState
	}
	if true {
		toSerialize["accessType"] = o.AccessType
	}
	return json.Marshal(toSerialize)
}

type NullableModel5GsUserStateInfo struct {
	value *Model5GsUserStateInfo
	isSet bool
}

func (v NullableModel5GsUserStateInfo) Get() *Model5GsUserStateInfo {
	return v.value
}

func (v *NullableModel5GsUserStateInfo) Set(val *Model5GsUserStateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModel5GsUserStateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModel5GsUserStateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel5GsUserStateInfo(val *Model5GsUserStateInfo) *NullableModel5GsUserStateInfo {
	return &NullableModel5GsUserStateInfo{value: val, isSet: true}
}

func (v NullableModel5GsUserStateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel5GsUserStateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


