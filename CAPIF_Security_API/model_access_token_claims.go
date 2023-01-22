/*
CAPIF_Security_API

API for CAPIF security management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CAPIF_Security_API

import (
	"encoding/json"
)

// AccessTokenClaims Represents the claims data structure for the access token.
type AccessTokenClaims struct {
	Iss string `json:"iss"`
	Scope string `json:"scope"`
	// Unsigned integer identifying a period of time in units of seconds.
	Exp int32 `json:"exp"`
}

// NewAccessTokenClaims instantiates a new AccessTokenClaims object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTokenClaims(iss string, scope string, exp int32) *AccessTokenClaims {
	this := AccessTokenClaims{}
	this.Iss = iss
	this.Scope = scope
	this.Exp = exp
	return &this
}

// NewAccessTokenClaimsWithDefaults instantiates a new AccessTokenClaims object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenClaimsWithDefaults() *AccessTokenClaims {
	this := AccessTokenClaims{}
	return &this
}

// GetIss returns the Iss field value
func (o *AccessTokenClaims) GetIss() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iss
}

// GetIssOk returns a tuple with the Iss field value
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetIssOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Iss, true
}

// SetIss sets field value
func (o *AccessTokenClaims) SetIss(v string) {
	o.Iss = v
}

// GetScope returns the Scope field value
func (o *AccessTokenClaims) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetScopeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *AccessTokenClaims) SetScope(v string) {
	o.Scope = v
}

// GetExp returns the Exp field value
func (o *AccessTokenClaims) GetExp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Exp
}

// GetExpOk returns a tuple with the Exp field value
// and a boolean to check if the value has been set.
func (o *AccessTokenClaims) GetExpOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Exp, true
}

// SetExp sets field value
func (o *AccessTokenClaims) SetExp(v int32) {
	o.Exp = v
}

func (o AccessTokenClaims) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["iss"] = o.Iss
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["exp"] = o.Exp
	}
	return json.Marshal(toSerialize)
}

type NullableAccessTokenClaims struct {
	value *AccessTokenClaims
	isSet bool
}

func (v NullableAccessTokenClaims) Get() *AccessTokenClaims {
	return v.value
}

func (v *NullableAccessTokenClaims) Set(val *AccessTokenClaims) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenClaims) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenClaims) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenClaims(val *AccessTokenClaims) *NullableAccessTokenClaims {
	return &NullableAccessTokenClaims{value: val, isSet: true}
}

func (v NullableAccessTokenClaims) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenClaims) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


