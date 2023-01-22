/*
Nnef_Authentication

NEF Auth Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnef_Authentication

import (
	"encoding/json"
)

// UAVAuthFailure UAV auth failure
type UAVAuthFailure struct {
	Error ProblemDetails `json:"error"`
	UasResourceRelease *bool `json:"uasResourceRelease,omitempty"`
}

// NewUAVAuthFailure instantiates a new UAVAuthFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUAVAuthFailure(error_ ProblemDetails) *UAVAuthFailure {
	this := UAVAuthFailure{}
	this.Error = error_
	var uasResourceRelease bool = false
	this.UasResourceRelease = &uasResourceRelease
	return &this
}

// NewUAVAuthFailureWithDefaults instantiates a new UAVAuthFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUAVAuthFailureWithDefaults() *UAVAuthFailure {
	this := UAVAuthFailure{}
	var uasResourceRelease bool = false
	this.UasResourceRelease = &uasResourceRelease
	return &this
}

// GetError returns the Error field value
func (o *UAVAuthFailure) GetError() ProblemDetails {
	if o == nil {
		var ret ProblemDetails
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *UAVAuthFailure) GetErrorOk() (*ProblemDetails, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *UAVAuthFailure) SetError(v ProblemDetails) {
	o.Error = v
}

// GetUasResourceRelease returns the UasResourceRelease field value if set, zero value otherwise.
func (o *UAVAuthFailure) GetUasResourceRelease() bool {
	if o == nil || isNil(o.UasResourceRelease) {
		var ret bool
		return ret
	}
	return *o.UasResourceRelease
}

// GetUasResourceReleaseOk returns a tuple with the UasResourceRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthFailure) GetUasResourceReleaseOk() (*bool, bool) {
	if o == nil || isNil(o.UasResourceRelease) {
    return nil, false
	}
	return o.UasResourceRelease, true
}

// HasUasResourceRelease returns a boolean if a field has been set.
func (o *UAVAuthFailure) HasUasResourceRelease() bool {
	if o != nil && !isNil(o.UasResourceRelease) {
		return true
	}

	return false
}

// SetUasResourceRelease gets a reference to the given bool and assigns it to the UasResourceRelease field.
func (o *UAVAuthFailure) SetUasResourceRelease(v bool) {
	o.UasResourceRelease = &v
}

func (o UAVAuthFailure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.UasResourceRelease) {
		toSerialize["uasResourceRelease"] = o.UasResourceRelease
	}
	return json.Marshal(toSerialize)
}

type NullableUAVAuthFailure struct {
	value *UAVAuthFailure
	isSet bool
}

func (v NullableUAVAuthFailure) Get() *UAVAuthFailure {
	return v.value
}

func (v *NullableUAVAuthFailure) Set(val *UAVAuthFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableUAVAuthFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableUAVAuthFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUAVAuthFailure(val *UAVAuthFailure) *NullableUAVAuthFailure {
	return &NullableUAVAuthFailure{value: val, isSet: true}
}

func (v NullableUAVAuthFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUAVAuthFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


