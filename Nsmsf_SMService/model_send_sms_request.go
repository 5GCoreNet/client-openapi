/*
Nsmsf_SMService Service API

SMSF SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmsf_SMService

import (
	"encoding/json"
	"os"
)

// SendSMSRequest struct for SendSMSRequest
type SendSMSRequest struct {
	JsonData *SmsRecordData `json:"jsonData,omitempty"`
	BinaryPayload **os.File `json:"binaryPayload,omitempty"`
}

// NewSendSMSRequest instantiates a new SendSMSRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendSMSRequest() *SendSMSRequest {
	this := SendSMSRequest{}
	return &this
}

// NewSendSMSRequestWithDefaults instantiates a new SendSMSRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendSMSRequestWithDefaults() *SendSMSRequest {
	this := SendSMSRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *SendSMSRequest) GetJsonData() SmsRecordData {
	if o == nil || isNil(o.JsonData) {
		var ret SmsRecordData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendSMSRequest) GetJsonDataOk() (*SmsRecordData, bool) {
	if o == nil || isNil(o.JsonData) {
    return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *SendSMSRequest) HasJsonData() bool {
	if o != nil && !isNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given SmsRecordData and assigns it to the JsonData field.
func (o *SendSMSRequest) SetJsonData(v SmsRecordData) {
	o.JsonData = &v
}

// GetBinaryPayload returns the BinaryPayload field value if set, zero value otherwise.
func (o *SendSMSRequest) GetBinaryPayload() *os.File {
	if o == nil || isNil(o.BinaryPayload) {
		var ret *os.File
		return ret
	}
	return *o.BinaryPayload
}

// GetBinaryPayloadOk returns a tuple with the BinaryPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendSMSRequest) GetBinaryPayloadOk() (**os.File, bool) {
	if o == nil || isNil(o.BinaryPayload) {
    return nil, false
	}
	return o.BinaryPayload, true
}

// HasBinaryPayload returns a boolean if a field has been set.
func (o *SendSMSRequest) HasBinaryPayload() bool {
	if o != nil && !isNil(o.BinaryPayload) {
		return true
	}

	return false
}

// SetBinaryPayload gets a reference to the given *os.File and assigns it to the BinaryPayload field.
func (o *SendSMSRequest) SetBinaryPayload(v *os.File) {
	o.BinaryPayload = &v
}

func (o SendSMSRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !isNil(o.BinaryPayload) {
		toSerialize["binaryPayload"] = o.BinaryPayload
	}
	return json.Marshal(toSerialize)
}

type NullableSendSMSRequest struct {
	value *SendSMSRequest
	isSet bool
}

func (v NullableSendSMSRequest) Get() *SendSMSRequest {
	return v.value
}

func (v *NullableSendSMSRequest) Set(val *SendSMSRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendSMSRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendSMSRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendSMSRequest(val *SendSMSRequest) *NullableSendSMSRequest {
	return &NullableSendSMSRequest{value: val, isSet: true}
}

func (v NullableSendSMSRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendSMSRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


