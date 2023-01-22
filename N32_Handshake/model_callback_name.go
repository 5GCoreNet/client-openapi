/*
N32 Handshake API

N32-c Handshake Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package N32_Handshake

import (
	"encoding/json"
)

// CallbackName Callback Name
type CallbackName struct {
	CallbackType string `json:"callbackType"`
}

// NewCallbackName instantiates a new CallbackName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallbackName(callbackType string) *CallbackName {
	this := CallbackName{}
	this.CallbackType = callbackType
	return &this
}

// NewCallbackNameWithDefaults instantiates a new CallbackName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallbackNameWithDefaults() *CallbackName {
	this := CallbackName{}
	return &this
}

// GetCallbackType returns the CallbackType field value
func (o *CallbackName) GetCallbackType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackType
}

// GetCallbackTypeOk returns a tuple with the CallbackType field value
// and a boolean to check if the value has been set.
func (o *CallbackName) GetCallbackTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CallbackType, true
}

// SetCallbackType sets field value
func (o *CallbackName) SetCallbackType(v string) {
	o.CallbackType = v
}

func (o CallbackName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["callbackType"] = o.CallbackType
	}
	return json.Marshal(toSerialize)
}

type NullableCallbackName struct {
	value *CallbackName
	isSet bool
}

func (v NullableCallbackName) Get() *CallbackName {
	return v.value
}

func (v *NullableCallbackName) Set(val *CallbackName) {
	v.value = val
	v.isSet = true
}

func (v NullableCallbackName) IsSet() bool {
	return v.isSet
}

func (v *NullableCallbackName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallbackName(val *CallbackName) *NullableCallbackName {
	return &NullableCallbackName{value: val, isSet: true}
}

func (v NullableCallbackName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallbackName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


