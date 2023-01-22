/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
)

// UpSecurity Contains Userplain security information.
type UpSecurity struct {
	UpIntegr UpIntegrity `json:"upIntegr"`
	UpConfid UpConfidentiality `json:"upConfid"`
}

// NewUpSecurity instantiates a new UpSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpSecurity(upIntegr UpIntegrity, upConfid UpConfidentiality) *UpSecurity {
	this := UpSecurity{}
	this.UpIntegr = upIntegr
	this.UpConfid = upConfid
	return &this
}

// NewUpSecurityWithDefaults instantiates a new UpSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpSecurityWithDefaults() *UpSecurity {
	this := UpSecurity{}
	return &this
}

// GetUpIntegr returns the UpIntegr field value
func (o *UpSecurity) GetUpIntegr() UpIntegrity {
	if o == nil {
		var ret UpIntegrity
		return ret
	}

	return o.UpIntegr
}

// GetUpIntegrOk returns a tuple with the UpIntegr field value
// and a boolean to check if the value has been set.
func (o *UpSecurity) GetUpIntegrOk() (*UpIntegrity, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpIntegr, true
}

// SetUpIntegr sets field value
func (o *UpSecurity) SetUpIntegr(v UpIntegrity) {
	o.UpIntegr = v
}

// GetUpConfid returns the UpConfid field value
func (o *UpSecurity) GetUpConfid() UpConfidentiality {
	if o == nil {
		var ret UpConfidentiality
		return ret
	}

	return o.UpConfid
}

// GetUpConfidOk returns a tuple with the UpConfid field value
// and a boolean to check if the value has been set.
func (o *UpSecurity) GetUpConfidOk() (*UpConfidentiality, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpConfid, true
}

// SetUpConfid sets field value
func (o *UpSecurity) SetUpConfid(v UpConfidentiality) {
	o.UpConfid = v
}

func (o UpSecurity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["upIntegr"] = o.UpIntegr
	}
	if true {
		toSerialize["upConfid"] = o.UpConfid
	}
	return json.Marshal(toSerialize)
}

type NullableUpSecurity struct {
	value *UpSecurity
	isSet bool
}

func (v NullableUpSecurity) Get() *UpSecurity {
	return v.value
}

func (v *NullableUpSecurity) Set(val *UpSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableUpSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableUpSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpSecurity(val *UpSecurity) *NullableUpSecurity {
	return &NullableUpSecurity{value: val, isSet: true}
}

func (v NullableUpSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


