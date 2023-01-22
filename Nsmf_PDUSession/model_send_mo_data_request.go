/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
	"os"
)

// SendMoDataRequest struct for SendMoDataRequest
type SendMoDataRequest struct {
	JsonData *SendMoDataReqData `json:"jsonData,omitempty"`
	BinaryMoData **os.File `json:"binaryMoData,omitempty"`
}

// NewSendMoDataRequest instantiates a new SendMoDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendMoDataRequest() *SendMoDataRequest {
	this := SendMoDataRequest{}
	return &this
}

// NewSendMoDataRequestWithDefaults instantiates a new SendMoDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendMoDataRequestWithDefaults() *SendMoDataRequest {
	this := SendMoDataRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *SendMoDataRequest) GetJsonData() SendMoDataReqData {
	if o == nil || isNil(o.JsonData) {
		var ret SendMoDataReqData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMoDataRequest) GetJsonDataOk() (*SendMoDataReqData, bool) {
	if o == nil || isNil(o.JsonData) {
    return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *SendMoDataRequest) HasJsonData() bool {
	if o != nil && !isNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given SendMoDataReqData and assigns it to the JsonData field.
func (o *SendMoDataRequest) SetJsonData(v SendMoDataReqData) {
	o.JsonData = &v
}

// GetBinaryMoData returns the BinaryMoData field value if set, zero value otherwise.
func (o *SendMoDataRequest) GetBinaryMoData() *os.File {
	if o == nil || isNil(o.BinaryMoData) {
		var ret *os.File
		return ret
	}
	return *o.BinaryMoData
}

// GetBinaryMoDataOk returns a tuple with the BinaryMoData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMoDataRequest) GetBinaryMoDataOk() (**os.File, bool) {
	if o == nil || isNil(o.BinaryMoData) {
    return nil, false
	}
	return o.BinaryMoData, true
}

// HasBinaryMoData returns a boolean if a field has been set.
func (o *SendMoDataRequest) HasBinaryMoData() bool {
	if o != nil && !isNil(o.BinaryMoData) {
		return true
	}

	return false
}

// SetBinaryMoData gets a reference to the given *os.File and assigns it to the BinaryMoData field.
func (o *SendMoDataRequest) SetBinaryMoData(v *os.File) {
	o.BinaryMoData = &v
}

func (o SendMoDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !isNil(o.BinaryMoData) {
		toSerialize["binaryMoData"] = o.BinaryMoData
	}
	return json.Marshal(toSerialize)
}

type NullableSendMoDataRequest struct {
	value *SendMoDataRequest
	isSet bool
}

func (v NullableSendMoDataRequest) Get() *SendMoDataRequest {
	return v.value
}

func (v *NullableSendMoDataRequest) Set(val *SendMoDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendMoDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendMoDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendMoDataRequest(val *SendMoDataRequest) *NullableSendMoDataRequest {
	return &NullableSendMoDataRequest{value: val, isSet: true}
}

func (v NullableSendMoDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendMoDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


