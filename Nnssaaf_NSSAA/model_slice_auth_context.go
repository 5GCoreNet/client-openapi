/*
Nnssaaf_NSSAA

Network Slice-Specific Authentication and Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnssaaf_NSSAA

import (
	"encoding/json"
)

// SliceAuthContext struct for SliceAuthContext
type SliceAuthContext struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi"`
	Snssai Snssai `json:"snssai"`
	// contains the resource ID of slice authentication context
	AuthCtxId string `json:"authCtxId"`
	// contains an EAP packet
	EapMessage NullableString `json:"eapMessage"`
}

// NewSliceAuthContext instantiates a new SliceAuthContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSliceAuthContext(gpsi string, snssai Snssai, authCtxId string, eapMessage NullableString) *SliceAuthContext {
	this := SliceAuthContext{}
	this.Gpsi = gpsi
	this.Snssai = snssai
	this.AuthCtxId = authCtxId
	this.EapMessage = eapMessage
	return &this
}

// NewSliceAuthContextWithDefaults instantiates a new SliceAuthContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliceAuthContextWithDefaults() *SliceAuthContext {
	this := SliceAuthContext{}
	return &this
}

// GetGpsi returns the Gpsi field value
func (o *SliceAuthContext) GetGpsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value
// and a boolean to check if the value has been set.
func (o *SliceAuthContext) GetGpsiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Gpsi, true
}

// SetGpsi sets field value
func (o *SliceAuthContext) SetGpsi(v string) {
	o.Gpsi = v
}

// GetSnssai returns the Snssai field value
func (o *SliceAuthContext) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *SliceAuthContext) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *SliceAuthContext) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetAuthCtxId returns the AuthCtxId field value
func (o *SliceAuthContext) GetAuthCtxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthCtxId
}

// GetAuthCtxIdOk returns a tuple with the AuthCtxId field value
// and a boolean to check if the value has been set.
func (o *SliceAuthContext) GetAuthCtxIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuthCtxId, true
}

// SetAuthCtxId sets field value
func (o *SliceAuthContext) SetAuthCtxId(v string) {
	o.AuthCtxId = v
}

// GetEapMessage returns the EapMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SliceAuthContext) GetEapMessage() string {
	if o == nil || o.EapMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.EapMessage.Get()
}

// GetEapMessageOk returns a tuple with the EapMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SliceAuthContext) GetEapMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EapMessage.Get(), o.EapMessage.IsSet()
}

// SetEapMessage sets field value
func (o *SliceAuthContext) SetEapMessage(v string) {
	o.EapMessage.Set(&v)
}

func (o SliceAuthContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["gpsi"] = o.Gpsi
	}
	if true {
		toSerialize["snssai"] = o.Snssai
	}
	if true {
		toSerialize["authCtxId"] = o.AuthCtxId
	}
	if true {
		toSerialize["eapMessage"] = o.EapMessage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSliceAuthContext struct {
	value *SliceAuthContext
	isSet bool
}

func (v NullableSliceAuthContext) Get() *SliceAuthContext {
	return v.value
}

func (v *NullableSliceAuthContext) Set(val *SliceAuthContext) {
	v.value = val
	v.isSet = true
}

func (v NullableSliceAuthContext) IsSet() bool {
	return v.isSet
}

func (v *NullableSliceAuthContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliceAuthContext(val *SliceAuthContext) *NullableSliceAuthContext {
	return &NullableSliceAuthContext{value: val, isSet: true}
}

func (v NullableSliceAuthContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliceAuthContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


