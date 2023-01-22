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

// PsiActivationState Public Service Identity activation state
type PsiActivationState struct {
	ActivationState ActivationState `json:"activationState"`
}

// NewPsiActivationState instantiates a new PsiActivationState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPsiActivationState(activationState ActivationState) *PsiActivationState {
	this := PsiActivationState{}
	this.ActivationState = activationState
	return &this
}

// NewPsiActivationStateWithDefaults instantiates a new PsiActivationState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPsiActivationStateWithDefaults() *PsiActivationState {
	this := PsiActivationState{}
	return &this
}

// GetActivationState returns the ActivationState field value
func (o *PsiActivationState) GetActivationState() ActivationState {
	if o == nil {
		var ret ActivationState
		return ret
	}

	return o.ActivationState
}

// GetActivationStateOk returns a tuple with the ActivationState field value
// and a boolean to check if the value has been set.
func (o *PsiActivationState) GetActivationStateOk() (*ActivationState, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ActivationState, true
}

// SetActivationState sets field value
func (o *PsiActivationState) SetActivationState(v ActivationState) {
	o.ActivationState = v
}

func (o PsiActivationState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["activationState"] = o.ActivationState
	}
	return json.Marshal(toSerialize)
}

type NullablePsiActivationState struct {
	value *PsiActivationState
	isSet bool
}

func (v NullablePsiActivationState) Get() *PsiActivationState {
	return v.value
}

func (v *NullablePsiActivationState) Set(val *PsiActivationState) {
	v.value = val
	v.isSet = true
}

func (v NullablePsiActivationState) IsSet() bool {
	return v.isSet
}

func (v *NullablePsiActivationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePsiActivationState(val *PsiActivationState) *NullablePsiActivationState {
	return &NullablePsiActivationState{value: val, isSet: true}
}

func (v NullablePsiActivationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePsiActivationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


