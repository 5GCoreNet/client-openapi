/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Location

import (
	"encoding/json"
)

// UeConnectivityState Indicates the connectivity state of a UE.
type UeConnectivityState struct {
	AccessType AccessType `json:"accessType"`
	Connectivitystate *CmState `json:"connectivitystate,omitempty"`
}

// NewUeConnectivityState instantiates a new UeConnectivityState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeConnectivityState(accessType AccessType) *UeConnectivityState {
	this := UeConnectivityState{}
	this.AccessType = accessType
	return &this
}

// NewUeConnectivityStateWithDefaults instantiates a new UeConnectivityState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeConnectivityStateWithDefaults() *UeConnectivityState {
	this := UeConnectivityState{}
	return &this
}

// GetAccessType returns the AccessType field value
func (o *UeConnectivityState) GetAccessType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *UeConnectivityState) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *UeConnectivityState) SetAccessType(v AccessType) {
	o.AccessType = v
}

// GetConnectivitystate returns the Connectivitystate field value if set, zero value otherwise.
func (o *UeConnectivityState) GetConnectivitystate() CmState {
	if o == nil || isNil(o.Connectivitystate) {
		var ret CmState
		return ret
	}
	return *o.Connectivitystate
}

// GetConnectivitystateOk returns a tuple with the Connectivitystate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeConnectivityState) GetConnectivitystateOk() (*CmState, bool) {
	if o == nil || isNil(o.Connectivitystate) {
    return nil, false
	}
	return o.Connectivitystate, true
}

// HasConnectivitystate returns a boolean if a field has been set.
func (o *UeConnectivityState) HasConnectivitystate() bool {
	if o != nil && !isNil(o.Connectivitystate) {
		return true
	}

	return false
}

// SetConnectivitystate gets a reference to the given CmState and assigns it to the Connectivitystate field.
func (o *UeConnectivityState) SetConnectivitystate(v CmState) {
	o.Connectivitystate = &v
}

func (o UeConnectivityState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accessType"] = o.AccessType
	}
	if !isNil(o.Connectivitystate) {
		toSerialize["connectivitystate"] = o.Connectivitystate
	}
	return json.Marshal(toSerialize)
}

type NullableUeConnectivityState struct {
	value *UeConnectivityState
	isSet bool
}

func (v NullableUeConnectivityState) Get() *UeConnectivityState {
	return v.value
}

func (v *NullableUeConnectivityState) Set(val *UeConnectivityState) {
	v.value = val
	v.isSet = true
}

func (v NullableUeConnectivityState) IsSet() bool {
	return v.isSet
}

func (v *NullableUeConnectivityState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeConnectivityState(val *UeConnectivityState) *NullableUeConnectivityState {
	return &NullableUeConnectivityState{value: val, isSet: true}
}

func (v NullableUeConnectivityState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeConnectivityState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


