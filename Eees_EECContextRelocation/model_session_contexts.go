/*
EES EEC Context Relocation API

API for EEC Context Relocation.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_EECContextRelocation

import (
	"encoding/json"
)

// SessionContexts Represents the list of service session contexts information.
type SessionContexts struct {
	// List of service session contexts information.
	SessCntxs []IndividualSessionContext `json:"sessCntxs"`
}

// NewSessionContexts instantiates a new SessionContexts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionContexts(sessCntxs []IndividualSessionContext) *SessionContexts {
	this := SessionContexts{}
	this.SessCntxs = sessCntxs
	return &this
}

// NewSessionContextsWithDefaults instantiates a new SessionContexts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionContextsWithDefaults() *SessionContexts {
	this := SessionContexts{}
	return &this
}

// GetSessCntxs returns the SessCntxs field value
func (o *SessionContexts) GetSessCntxs() []IndividualSessionContext {
	if o == nil {
		var ret []IndividualSessionContext
		return ret
	}

	return o.SessCntxs
}

// GetSessCntxsOk returns a tuple with the SessCntxs field value
// and a boolean to check if the value has been set.
func (o *SessionContexts) GetSessCntxsOk() ([]IndividualSessionContext, bool) {
	if o == nil {
    return nil, false
	}
	return o.SessCntxs, true
}

// SetSessCntxs sets field value
func (o *SessionContexts) SetSessCntxs(v []IndividualSessionContext) {
	o.SessCntxs = v
}

func (o SessionContexts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sessCntxs"] = o.SessCntxs
	}
	return json.Marshal(toSerialize)
}

type NullableSessionContexts struct {
	value *SessionContexts
	isSet bool
}

func (v NullableSessionContexts) Get() *SessionContexts {
	return v.value
}

func (v *NullableSessionContexts) Set(val *SessionContexts) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionContexts) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionContexts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionContexts(val *SessionContexts) *NullableSessionContexts {
	return &NullableSessionContexts{value: val, isSet: true}
}

func (v NullableSessionContexts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionContexts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


