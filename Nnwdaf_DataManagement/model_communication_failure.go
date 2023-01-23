/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// CommunicationFailure Describes a communication failure detected by AMF
type CommunicationFailure struct {
	NasReleaseCode *string `json:"nasReleaseCode,omitempty"`
	RanReleaseCode *NgApCause `json:"ranReleaseCode,omitempty"`
}

// NewCommunicationFailure instantiates a new CommunicationFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationFailure() *CommunicationFailure {
	this := CommunicationFailure{}
	return &this
}

// NewCommunicationFailureWithDefaults instantiates a new CommunicationFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationFailureWithDefaults() *CommunicationFailure {
	this := CommunicationFailure{}
	return &this
}

// GetNasReleaseCode returns the NasReleaseCode field value if set, zero value otherwise.
func (o *CommunicationFailure) GetNasReleaseCode() string {
	if o == nil || isNil(o.NasReleaseCode) {
		var ret string
		return ret
	}
	return *o.NasReleaseCode
}

// GetNasReleaseCodeOk returns a tuple with the NasReleaseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationFailure) GetNasReleaseCodeOk() (*string, bool) {
	if o == nil || isNil(o.NasReleaseCode) {
    return nil, false
	}
	return o.NasReleaseCode, true
}

// HasNasReleaseCode returns a boolean if a field has been set.
func (o *CommunicationFailure) HasNasReleaseCode() bool {
	if o != nil && !isNil(o.NasReleaseCode) {
		return true
	}

	return false
}

// SetNasReleaseCode gets a reference to the given string and assigns it to the NasReleaseCode field.
func (o *CommunicationFailure) SetNasReleaseCode(v string) {
	o.NasReleaseCode = &v
}

// GetRanReleaseCode returns the RanReleaseCode field value if set, zero value otherwise.
func (o *CommunicationFailure) GetRanReleaseCode() NgApCause {
	if o == nil || isNil(o.RanReleaseCode) {
		var ret NgApCause
		return ret
	}
	return *o.RanReleaseCode
}

// GetRanReleaseCodeOk returns a tuple with the RanReleaseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationFailure) GetRanReleaseCodeOk() (*NgApCause, bool) {
	if o == nil || isNil(o.RanReleaseCode) {
    return nil, false
	}
	return o.RanReleaseCode, true
}

// HasRanReleaseCode returns a boolean if a field has been set.
func (o *CommunicationFailure) HasRanReleaseCode() bool {
	if o != nil && !isNil(o.RanReleaseCode) {
		return true
	}

	return false
}

// SetRanReleaseCode gets a reference to the given NgApCause and assigns it to the RanReleaseCode field.
func (o *CommunicationFailure) SetRanReleaseCode(v NgApCause) {
	o.RanReleaseCode = &v
}

func (o CommunicationFailure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NasReleaseCode) {
		toSerialize["nasReleaseCode"] = o.NasReleaseCode
	}
	if !isNil(o.RanReleaseCode) {
		toSerialize["ranReleaseCode"] = o.RanReleaseCode
	}
	return json.Marshal(toSerialize)
}

type NullableCommunicationFailure struct {
	value *CommunicationFailure
	isSet bool
}

func (v NullableCommunicationFailure) Get() *CommunicationFailure {
	return v.value
}

func (v *NullableCommunicationFailure) Set(val *CommunicationFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationFailure(val *CommunicationFailure) *NullableCommunicationFailure {
	return &NullableCommunicationFailure{value: val, isSet: true}
}

func (v NullableCommunicationFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


