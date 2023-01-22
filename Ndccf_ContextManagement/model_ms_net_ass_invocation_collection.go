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

// MsNetAssInvocationCollection Contains the Media Streaming Network Assistance invocation collected for an UE Application  via AF. 
type MsNetAssInvocationCollection struct {
	MsNetAssInvocs []NetworkAssistanceSession `json:"msNetAssInvocs"`
}

// NewMsNetAssInvocationCollection instantiates a new MsNetAssInvocationCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsNetAssInvocationCollection(msNetAssInvocs []NetworkAssistanceSession) *MsNetAssInvocationCollection {
	this := MsNetAssInvocationCollection{}
	this.MsNetAssInvocs = msNetAssInvocs
	return &this
}

// NewMsNetAssInvocationCollectionWithDefaults instantiates a new MsNetAssInvocationCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsNetAssInvocationCollectionWithDefaults() *MsNetAssInvocationCollection {
	this := MsNetAssInvocationCollection{}
	return &this
}

// GetMsNetAssInvocs returns the MsNetAssInvocs field value
func (o *MsNetAssInvocationCollection) GetMsNetAssInvocs() []NetworkAssistanceSession {
	if o == nil {
		var ret []NetworkAssistanceSession
		return ret
	}

	return o.MsNetAssInvocs
}

// GetMsNetAssInvocsOk returns a tuple with the MsNetAssInvocs field value
// and a boolean to check if the value has been set.
func (o *MsNetAssInvocationCollection) GetMsNetAssInvocsOk() ([]NetworkAssistanceSession, bool) {
	if o == nil {
    return nil, false
	}
	return o.MsNetAssInvocs, true
}

// SetMsNetAssInvocs sets field value
func (o *MsNetAssInvocationCollection) SetMsNetAssInvocs(v []NetworkAssistanceSession) {
	o.MsNetAssInvocs = v
}

func (o MsNetAssInvocationCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["msNetAssInvocs"] = o.MsNetAssInvocs
	}
	return json.Marshal(toSerialize)
}

type NullableMsNetAssInvocationCollection struct {
	value *MsNetAssInvocationCollection
	isSet bool
}

func (v NullableMsNetAssInvocationCollection) Get() *MsNetAssInvocationCollection {
	return v.value
}

func (v *NullableMsNetAssInvocationCollection) Set(val *MsNetAssInvocationCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableMsNetAssInvocationCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableMsNetAssInvocationCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsNetAssInvocationCollection(val *MsNetAssInvocationCollection) *NullableMsNetAssInvocationCollection {
	return &NullableMsNetAssInvocationCollection{value: val, isSet: true}
}

func (v NullableMsNetAssInvocationCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsNetAssInvocationCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


