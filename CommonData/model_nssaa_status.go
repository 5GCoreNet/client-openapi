/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
)

// NssaaStatus contains the Subscribed S-NSSAI subject to NSSAA procedure and the status.
type NssaaStatus struct {
	Snssai Snssai `json:"snssai"`
	Status AuthStatus `json:"status"`
}

// NewNssaaStatus instantiates a new NssaaStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNssaaStatus(snssai Snssai, status AuthStatus) *NssaaStatus {
	this := NssaaStatus{}
	this.Snssai = snssai
	this.Status = status
	return &this
}

// NewNssaaStatusWithDefaults instantiates a new NssaaStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNssaaStatusWithDefaults() *NssaaStatus {
	this := NssaaStatus{}
	return &this
}

// GetSnssai returns the Snssai field value
func (o *NssaaStatus) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *NssaaStatus) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *NssaaStatus) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetStatus returns the Status field value
func (o *NssaaStatus) GetStatus() AuthStatus {
	if o == nil {
		var ret AuthStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NssaaStatus) GetStatusOk() (*AuthStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NssaaStatus) SetStatus(v AuthStatus) {
	o.Status = v
}

func (o NssaaStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["snssai"] = o.Snssai
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableNssaaStatus struct {
	value *NssaaStatus
	isSet bool
}

func (v NullableNssaaStatus) Get() *NssaaStatus {
	return v.value
}

func (v *NullableNssaaStatus) Set(val *NssaaStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNssaaStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNssaaStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNssaaStatus(val *NssaaStatus) *NullableNssaaStatus {
	return &NullableNssaaStatus{value: val, isSet: true}
}

func (v NullableNssaaStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNssaaStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


