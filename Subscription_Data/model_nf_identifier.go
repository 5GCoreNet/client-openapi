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

// NfIdentifier struct for NfIdentifier
type NfIdentifier struct {
	NfType NFType `json:"nfType"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId *string `json:"nfInstanceId,omitempty"`
}

// NewNfIdentifier instantiates a new NfIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfIdentifier(nfType NFType) *NfIdentifier {
	this := NfIdentifier{}
	this.NfType = nfType
	return &this
}

// NewNfIdentifierWithDefaults instantiates a new NfIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfIdentifierWithDefaults() *NfIdentifier {
	this := NfIdentifier{}
	return &this
}

// GetNfType returns the NfType field value
func (o *NfIdentifier) GetNfType() NFType {
	if o == nil {
		var ret NFType
		return ret
	}

	return o.NfType
}

// GetNfTypeOk returns a tuple with the NfType field value
// and a boolean to check if the value has been set.
func (o *NfIdentifier) GetNfTypeOk() (*NFType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NfType, true
}

// SetNfType sets field value
func (o *NfIdentifier) SetNfType(v NFType) {
	o.NfType = v
}

// GetNfInstanceId returns the NfInstanceId field value if set, zero value otherwise.
func (o *NfIdentifier) GetNfInstanceId() string {
	if o == nil || isNil(o.NfInstanceId) {
		var ret string
		return ret
	}
	return *o.NfInstanceId
}

// GetNfInstanceIdOk returns a tuple with the NfInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfIdentifier) GetNfInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.NfInstanceId) {
    return nil, false
	}
	return o.NfInstanceId, true
}

// HasNfInstanceId returns a boolean if a field has been set.
func (o *NfIdentifier) HasNfInstanceId() bool {
	if o != nil && !isNil(o.NfInstanceId) {
		return true
	}

	return false
}

// SetNfInstanceId gets a reference to the given string and assigns it to the NfInstanceId field.
func (o *NfIdentifier) SetNfInstanceId(v string) {
	o.NfInstanceId = &v
}

func (o NfIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nfType"] = o.NfType
	}
	if !isNil(o.NfInstanceId) {
		toSerialize["nfInstanceId"] = o.NfInstanceId
	}
	return json.Marshal(toSerialize)
}

type NullableNfIdentifier struct {
	value *NfIdentifier
	isSet bool
}

func (v NullableNfIdentifier) Get() *NfIdentifier {
	return v.value
}

func (v *NullableNfIdentifier) Set(val *NfIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableNfIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableNfIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfIdentifier(val *NfIdentifier) *NullableNfIdentifier {
	return &NullableNfIdentifier{value: val, isSet: true}
}

func (v NullableNfIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


