/*
Namf_MBSCommunication

AMF Communication Service for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSCommunication

import (
	"encoding/json"
	"os"
)

// N2MessageTransferRequest struct for N2MessageTransferRequest
type N2MessageTransferRequest struct {
	JsonData *MbsN2MessageTransferReqData `json:"jsonData,omitempty"`
	BinaryDataN2Information **os.File `json:"binaryDataN2Information,omitempty"`
}

// NewN2MessageTransferRequest instantiates a new N2MessageTransferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN2MessageTransferRequest() *N2MessageTransferRequest {
	this := N2MessageTransferRequest{}
	return &this
}

// NewN2MessageTransferRequestWithDefaults instantiates a new N2MessageTransferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN2MessageTransferRequestWithDefaults() *N2MessageTransferRequest {
	this := N2MessageTransferRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *N2MessageTransferRequest) GetJsonData() MbsN2MessageTransferReqData {
	if o == nil || isNil(o.JsonData) {
		var ret MbsN2MessageTransferReqData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2MessageTransferRequest) GetJsonDataOk() (*MbsN2MessageTransferReqData, bool) {
	if o == nil || isNil(o.JsonData) {
    return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *N2MessageTransferRequest) HasJsonData() bool {
	if o != nil && !isNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given MbsN2MessageTransferReqData and assigns it to the JsonData field.
func (o *N2MessageTransferRequest) SetJsonData(v MbsN2MessageTransferReqData) {
	o.JsonData = &v
}

// GetBinaryDataN2Information returns the BinaryDataN2Information field value if set, zero value otherwise.
func (o *N2MessageTransferRequest) GetBinaryDataN2Information() *os.File {
	if o == nil || isNil(o.BinaryDataN2Information) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information
}

// GetBinaryDataN2InformationOk returns a tuple with the BinaryDataN2Information field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2MessageTransferRequest) GetBinaryDataN2InformationOk() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2Information) {
    return nil, false
	}
	return o.BinaryDataN2Information, true
}

// HasBinaryDataN2Information returns a boolean if a field has been set.
func (o *N2MessageTransferRequest) HasBinaryDataN2Information() bool {
	if o != nil && !isNil(o.BinaryDataN2Information) {
		return true
	}

	return false
}

// SetBinaryDataN2Information gets a reference to the given *os.File and assigns it to the BinaryDataN2Information field.
func (o *N2MessageTransferRequest) SetBinaryDataN2Information(v *os.File) {
	o.BinaryDataN2Information = &v
}

func (o N2MessageTransferRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !isNil(o.BinaryDataN2Information) {
		toSerialize["binaryDataN2Information"] = o.BinaryDataN2Information
	}
	return json.Marshal(toSerialize)
}

type NullableN2MessageTransferRequest struct {
	value *N2MessageTransferRequest
	isSet bool
}

func (v NullableN2MessageTransferRequest) Get() *N2MessageTransferRequest {
	return v.value
}

func (v *NullableN2MessageTransferRequest) Set(val *N2MessageTransferRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableN2MessageTransferRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableN2MessageTransferRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2MessageTransferRequest(val *N2MessageTransferRequest) *NullableN2MessageTransferRequest {
	return &NullableN2MessageTransferRequest{value: val, isSet: true}
}

func (v NullableN2MessageTransferRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2MessageTransferRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


