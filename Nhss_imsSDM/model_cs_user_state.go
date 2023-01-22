/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsSDM

import (
	"encoding/json"
)

// CsUserState User state in CS domain
type CsUserState struct {
	MscVlrUserState UserStateCs `json:"mscVlrUserState"`
}

// NewCsUserState instantiates a new CsUserState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsUserState(mscVlrUserState UserStateCs) *CsUserState {
	this := CsUserState{}
	this.MscVlrUserState = mscVlrUserState
	return &this
}

// NewCsUserStateWithDefaults instantiates a new CsUserState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsUserStateWithDefaults() *CsUserState {
	this := CsUserState{}
	return &this
}

// GetMscVlrUserState returns the MscVlrUserState field value
func (o *CsUserState) GetMscVlrUserState() UserStateCs {
	if o == nil {
		var ret UserStateCs
		return ret
	}

	return o.MscVlrUserState
}

// GetMscVlrUserStateOk returns a tuple with the MscVlrUserState field value
// and a boolean to check if the value has been set.
func (o *CsUserState) GetMscVlrUserStateOk() (*UserStateCs, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MscVlrUserState, true
}

// SetMscVlrUserState sets field value
func (o *CsUserState) SetMscVlrUserState(v UserStateCs) {
	o.MscVlrUserState = v
}

func (o CsUserState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mscVlrUserState"] = o.MscVlrUserState
	}
	return json.Marshal(toSerialize)
}

type NullableCsUserState struct {
	value *CsUserState
	isSet bool
}

func (v NullableCsUserState) Get() *CsUserState {
	return v.value
}

func (v *NullableCsUserState) Set(val *CsUserState) {
	v.value = val
	v.isSet = true
}

func (v NullableCsUserState) IsSet() bool {
	return v.isSet
}

func (v *NullableCsUserState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsUserState(val *CsUserState) *NullableCsUserState {
	return &NullableCsUserState{value: val, isSet: true}
}

func (v NullableCsUserState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsUserState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

