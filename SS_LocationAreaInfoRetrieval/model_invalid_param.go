/*
SS_LocationAreaInfoRetrieval

API for SEAL Location Area Info Retrieval.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_LocationAreaInfoRetrieval

import (
	"encoding/json"
)

// InvalidParam Represents the description of invalid parameters, for a request rejected due to invalid parameters.
type InvalidParam struct {
	// Attribute's name encoded as a JSON Pointer, or header's name.
	Param string `json:"param"`
	// A human-readable reason, e.g. \"must be a positive integer\".
	Reason *string `json:"reason,omitempty"`
}

// NewInvalidParam instantiates a new InvalidParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidParam(param string) *InvalidParam {
	this := InvalidParam{}
	this.Param = param
	return &this
}

// NewInvalidParamWithDefaults instantiates a new InvalidParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidParamWithDefaults() *InvalidParam {
	this := InvalidParam{}
	return &this
}

// GetParam returns the Param field value
func (o *InvalidParam) GetParam() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Param
}

// GetParamOk returns a tuple with the Param field value
// and a boolean to check if the value has been set.
func (o *InvalidParam) GetParamOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Param, true
}

// SetParam sets field value
func (o *InvalidParam) SetParam(v string) {
	o.Param = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InvalidParam) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidParam) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InvalidParam) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InvalidParam) SetReason(v string) {
	o.Reason = &v
}

func (o InvalidParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["param"] = o.Param
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableInvalidParam struct {
	value *InvalidParam
	isSet bool
}

func (v NullableInvalidParam) Get() *InvalidParam {
	return v.value
}

func (v *NullableInvalidParam) Set(val *InvalidParam) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidParam) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidParam(val *InvalidParam) *NullableInvalidParam {
	return &NullableInvalidParam{value: val, isSet: true}
}

func (v NullableInvalidParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


