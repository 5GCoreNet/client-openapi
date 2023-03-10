/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Location

import (
	"encoding/json"
	"os"
)

// DetermineLocationRequest struct for DetermineLocationRequest
type DetermineLocationRequest struct {
	JsonData *InputData `json:"jsonData,omitempty"`
	BinaryDataLppMessage **os.File `json:"binaryDataLppMessage,omitempty"`
}

// NewDetermineLocationRequest instantiates a new DetermineLocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetermineLocationRequest() *DetermineLocationRequest {
	this := DetermineLocationRequest{}
	return &this
}

// NewDetermineLocationRequestWithDefaults instantiates a new DetermineLocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetermineLocationRequestWithDefaults() *DetermineLocationRequest {
	this := DetermineLocationRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *DetermineLocationRequest) GetJsonData() InputData {
	if o == nil || isNil(o.JsonData) {
		var ret InputData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetermineLocationRequest) GetJsonDataOk() (*InputData, bool) {
	if o == nil || isNil(o.JsonData) {
    return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *DetermineLocationRequest) HasJsonData() bool {
	if o != nil && !isNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given InputData and assigns it to the JsonData field.
func (o *DetermineLocationRequest) SetJsonData(v InputData) {
	o.JsonData = &v
}

// GetBinaryDataLppMessage returns the BinaryDataLppMessage field value if set, zero value otherwise.
func (o *DetermineLocationRequest) GetBinaryDataLppMessage() *os.File {
	if o == nil || isNil(o.BinaryDataLppMessage) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataLppMessage
}

// GetBinaryDataLppMessageOk returns a tuple with the BinaryDataLppMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetermineLocationRequest) GetBinaryDataLppMessageOk() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataLppMessage) {
    return nil, false
	}
	return o.BinaryDataLppMessage, true
}

// HasBinaryDataLppMessage returns a boolean if a field has been set.
func (o *DetermineLocationRequest) HasBinaryDataLppMessage() bool {
	if o != nil && !isNil(o.BinaryDataLppMessage) {
		return true
	}

	return false
}

// SetBinaryDataLppMessage gets a reference to the given *os.File and assigns it to the BinaryDataLppMessage field.
func (o *DetermineLocationRequest) SetBinaryDataLppMessage(v *os.File) {
	o.BinaryDataLppMessage = &v
}

func (o DetermineLocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !isNil(o.BinaryDataLppMessage) {
		toSerialize["binaryDataLppMessage"] = o.BinaryDataLppMessage
	}
	return json.Marshal(toSerialize)
}

type NullableDetermineLocationRequest struct {
	value *DetermineLocationRequest
	isSet bool
}

func (v NullableDetermineLocationRequest) Get() *DetermineLocationRequest {
	return v.value
}

func (v *NullableDetermineLocationRequest) Set(val *DetermineLocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDetermineLocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDetermineLocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetermineLocationRequest(val *DetermineLocationRequest) *NullableDetermineLocationRequest {
	return &NullableDetermineLocationRequest{value: val, isSet: true}
}

func (v NullableDetermineLocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetermineLocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


