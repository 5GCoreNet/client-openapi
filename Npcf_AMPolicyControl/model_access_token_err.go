/*
Npcf_AMPolicyControl

Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_AMPolicyControl

import (
	"encoding/json"
)

// AccessTokenErr Error returned in the access token response message
type AccessTokenErr struct {
	Error string `json:"error"`
	ErrorDescription *string `json:"error_description,omitempty"`
	ErrorUri *string `json:"error_uri,omitempty"`
}

// NewAccessTokenErr instantiates a new AccessTokenErr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenErr(error_ string) *AccessTokenErr {
	this := AccessTokenErr{}
	this.Error = error_
	return &this
}

// NewAccessTokenErrWithDefaults instantiates a new AccessTokenErr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenErrWithDefaults() *AccessTokenErr {
	this := AccessTokenErr{}
	return &this
}

// GetError returns the Error field value
func (o *AccessTokenErr) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AccessTokenErr) GetErrorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AccessTokenErr) SetError(v string) {
	o.Error = v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *AccessTokenErr) GetErrorDescription() string {
	if o == nil || isNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenErr) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ErrorDescription) {
    return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *AccessTokenErr) HasErrorDescription() bool {
	if o != nil && !isNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *AccessTokenErr) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorUri returns the ErrorUri field value if set, zero value otherwise.
func (o *AccessTokenErr) GetErrorUri() string {
	if o == nil || isNil(o.ErrorUri) {
		var ret string
		return ret
	}
	return *o.ErrorUri
}

// GetErrorUriOk returns a tuple with the ErrorUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessTokenErr) GetErrorUriOk() (*string, bool) {
	if o == nil || isNil(o.ErrorUri) {
    return nil, false
	}
	return o.ErrorUri, true
}

// HasErrorUri returns a boolean if a field has been set.
func (o *AccessTokenErr) HasErrorUri() bool {
	if o != nil && !isNil(o.ErrorUri) {
		return true
	}

	return false
}

// SetErrorUri gets a reference to the given string and assigns it to the ErrorUri field.
func (o *AccessTokenErr) SetErrorUri(v string) {
	o.ErrorUri = &v
}

func (o AccessTokenErr) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.ErrorDescription) {
		toSerialize["error_description"] = o.ErrorDescription
	}
	if !isNil(o.ErrorUri) {
		toSerialize["error_uri"] = o.ErrorUri
	}
	return json.Marshal(toSerialize)
}

type NullableAccessTokenErr struct {
	value *AccessTokenErr
	isSet bool
}

func (v NullableAccessTokenErr) Get() *AccessTokenErr {
	return v.value
}

func (v *NullableAccessTokenErr) Set(val *AccessTokenErr) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenErr) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenErr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenErr(val *AccessTokenErr) *NullableAccessTokenErr {
	return &NullableAccessTokenErr{value: val, isSet: true}
}

func (v NullableAccessTokenErr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenErr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


