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

// SendMtSMS200Response struct for SendMtSMS200Response
type SendMtSMS200Response struct {
	JsonData *SmsDeliveryData `json:"jsonData,omitempty"`
	BinaryPayload **os.File `json:"binaryPayload,omitempty"`
}

// NewSendMtSMS200Response instantiates a new SendMtSMS200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendMtSMS200Response() *SendMtSMS200Response {
	this := SendMtSMS200Response{}
	return &this
}

// NewSendMtSMS200ResponseWithDefaults instantiates a new SendMtSMS200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendMtSMS200ResponseWithDefaults() *SendMtSMS200Response {
	this := SendMtSMS200Response{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *SendMtSMS200Response) GetJsonData() SmsDeliveryData {
	if o == nil || isNil(o.JsonData) {
		var ret SmsDeliveryData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMtSMS200Response) GetJsonDataOk() (*SmsDeliveryData, bool) {
	if o == nil || isNil(o.JsonData) {
    return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *SendMtSMS200Response) HasJsonData() bool {
	if o != nil && !isNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given SmsDeliveryData and assigns it to the JsonData field.
func (o *SendMtSMS200Response) SetJsonData(v SmsDeliveryData) {
	o.JsonData = &v
}

// GetBinaryPayload returns the BinaryPayload field value if set, zero value otherwise.
func (o *SendMtSMS200Response) GetBinaryPayload() *os.File {
	if o == nil || isNil(o.BinaryPayload) {
		var ret *os.File
		return ret
	}
	return *o.BinaryPayload
}

// GetBinaryPayloadOk returns a tuple with the BinaryPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendMtSMS200Response) GetBinaryPayloadOk() (**os.File, bool) {
	if o == nil || isNil(o.BinaryPayload) {
    return nil, false
	}
	return o.BinaryPayload, true
}

// HasBinaryPayload returns a boolean if a field has been set.
func (o *SendMtSMS200Response) HasBinaryPayload() bool {
	if o != nil && !isNil(o.BinaryPayload) {
		return true
	}

	return false
}

// SetBinaryPayload gets a reference to the given *os.File and assigns it to the BinaryPayload field.
func (o *SendMtSMS200Response) SetBinaryPayload(v *os.File) {
	o.BinaryPayload = &v
}

func (o SendMtSMS200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !isNil(o.BinaryPayload) {
		toSerialize["binaryPayload"] = o.BinaryPayload
	}
	return json.Marshal(toSerialize)
}

type NullableSendMtSMS200Response struct {
	value *SendMtSMS200Response
	isSet bool
}

func (v NullableSendMtSMS200Response) Get() *SendMtSMS200Response {
	return v.value
}

func (v *NullableSendMtSMS200Response) Set(val *SendMtSMS200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSendMtSMS200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSendMtSMS200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendMtSMS200Response(val *SendMtSMS200Response) *NullableSendMtSMS200Response {
	return &NullableSendMtSMS200Response{value: val, isSet: true}
}

func (v NullableSendMtSMS200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendMtSMS200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


