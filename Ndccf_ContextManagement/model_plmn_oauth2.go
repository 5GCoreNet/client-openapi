/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_ContextManagement

import (
	"encoding/json"
)

// PlmnOauth2 Oauth2.0 required indication for a given PLMN ID
type PlmnOauth2 struct {
	Oauth2RequiredPlmnIdList []PlmnId `json:"oauth2RequiredPlmnIdList,omitempty"`
	Oauth2NotRequiredPlmnIdList []PlmnId `json:"oauth2NotRequiredPlmnIdList,omitempty"`
}

// NewPlmnOauth2 instantiates a new PlmnOauth2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnOauth2() *PlmnOauth2 {
	this := PlmnOauth2{}
	return &this
}

// NewPlmnOauth2WithDefaults instantiates a new PlmnOauth2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnOauth2WithDefaults() *PlmnOauth2 {
	this := PlmnOauth2{}
	return &this
}

// GetOauth2RequiredPlmnIdList returns the Oauth2RequiredPlmnIdList field value if set, zero value otherwise.
func (o *PlmnOauth2) GetOauth2RequiredPlmnIdList() []PlmnId {
	if o == nil || isNil(o.Oauth2RequiredPlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.Oauth2RequiredPlmnIdList
}

// GetOauth2RequiredPlmnIdListOk returns a tuple with the Oauth2RequiredPlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnOauth2) GetOauth2RequiredPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || isNil(o.Oauth2RequiredPlmnIdList) {
    return nil, false
	}
	return o.Oauth2RequiredPlmnIdList, true
}

// HasOauth2RequiredPlmnIdList returns a boolean if a field has been set.
func (o *PlmnOauth2) HasOauth2RequiredPlmnIdList() bool {
	if o != nil && !isNil(o.Oauth2RequiredPlmnIdList) {
		return true
	}

	return false
}

// SetOauth2RequiredPlmnIdList gets a reference to the given []PlmnId and assigns it to the Oauth2RequiredPlmnIdList field.
func (o *PlmnOauth2) SetOauth2RequiredPlmnIdList(v []PlmnId) {
	o.Oauth2RequiredPlmnIdList = v
}

// GetOauth2NotRequiredPlmnIdList returns the Oauth2NotRequiredPlmnIdList field value if set, zero value otherwise.
func (o *PlmnOauth2) GetOauth2NotRequiredPlmnIdList() []PlmnId {
	if o == nil || isNil(o.Oauth2NotRequiredPlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.Oauth2NotRequiredPlmnIdList
}

// GetOauth2NotRequiredPlmnIdListOk returns a tuple with the Oauth2NotRequiredPlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnOauth2) GetOauth2NotRequiredPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || isNil(o.Oauth2NotRequiredPlmnIdList) {
    return nil, false
	}
	return o.Oauth2NotRequiredPlmnIdList, true
}

// HasOauth2NotRequiredPlmnIdList returns a boolean if a field has been set.
func (o *PlmnOauth2) HasOauth2NotRequiredPlmnIdList() bool {
	if o != nil && !isNil(o.Oauth2NotRequiredPlmnIdList) {
		return true
	}

	return false
}

// SetOauth2NotRequiredPlmnIdList gets a reference to the given []PlmnId and assigns it to the Oauth2NotRequiredPlmnIdList field.
func (o *PlmnOauth2) SetOauth2NotRequiredPlmnIdList(v []PlmnId) {
	o.Oauth2NotRequiredPlmnIdList = v
}

func (o PlmnOauth2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Oauth2RequiredPlmnIdList) {
		toSerialize["oauth2RequiredPlmnIdList"] = o.Oauth2RequiredPlmnIdList
	}
	if !isNil(o.Oauth2NotRequiredPlmnIdList) {
		toSerialize["oauth2NotRequiredPlmnIdList"] = o.Oauth2NotRequiredPlmnIdList
	}
	return json.Marshal(toSerialize)
}

type NullablePlmnOauth2 struct {
	value *PlmnOauth2
	isSet bool
}

func (v NullablePlmnOauth2) Get() *PlmnOauth2 {
	return v.value
}

func (v *NullablePlmnOauth2) Set(val *PlmnOauth2) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnOauth2) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnOauth2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnOauth2(val *PlmnOauth2) *NullablePlmnOauth2 {
	return &NullablePlmnOauth2{value: val, isSet: true}
}

func (v NullablePlmnOauth2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnOauth2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


