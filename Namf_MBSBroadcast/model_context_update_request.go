/*
Namf_MBSBroadcast

AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSBroadcast

import (
	"encoding/json"
	"os"
)

// ContextUpdateRequest struct for ContextUpdateRequest
type ContextUpdateRequest struct {
	JsonData *ContextUpdateReqData `json:"jsonData,omitempty"`
	BinaryDataN2Information **os.File `json:"binaryDataN2Information,omitempty"`
}

// NewContextUpdateRequest instantiates a new ContextUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextUpdateRequest() *ContextUpdateRequest {
	this := ContextUpdateRequest{}
	return &this
}

// NewContextUpdateRequestWithDefaults instantiates a new ContextUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextUpdateRequestWithDefaults() *ContextUpdateRequest {
	this := ContextUpdateRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *ContextUpdateRequest) GetJsonData() ContextUpdateReqData {
	if o == nil || isNil(o.JsonData) {
		var ret ContextUpdateReqData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateRequest) GetJsonDataOk() (*ContextUpdateReqData, bool) {
	if o == nil || isNil(o.JsonData) {
    return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *ContextUpdateRequest) HasJsonData() bool {
	if o != nil && !isNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given ContextUpdateReqData and assigns it to the JsonData field.
func (o *ContextUpdateRequest) SetJsonData(v ContextUpdateReqData) {
	o.JsonData = &v
}

// GetBinaryDataN2Information returns the BinaryDataN2Information field value if set, zero value otherwise.
func (o *ContextUpdateRequest) GetBinaryDataN2Information() *os.File {
	if o == nil || isNil(o.BinaryDataN2Information) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information
}

// GetBinaryDataN2InformationOk returns a tuple with the BinaryDataN2Information field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdateRequest) GetBinaryDataN2InformationOk() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2Information) {
    return nil, false
	}
	return o.BinaryDataN2Information, true
}

// HasBinaryDataN2Information returns a boolean if a field has been set.
func (o *ContextUpdateRequest) HasBinaryDataN2Information() bool {
	if o != nil && !isNil(o.BinaryDataN2Information) {
		return true
	}

	return false
}

// SetBinaryDataN2Information gets a reference to the given *os.File and assigns it to the BinaryDataN2Information field.
func (o *ContextUpdateRequest) SetBinaryDataN2Information(v *os.File) {
	o.BinaryDataN2Information = &v
}

func (o ContextUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !isNil(o.BinaryDataN2Information) {
		toSerialize["binaryDataN2Information"] = o.BinaryDataN2Information
	}
	return json.Marshal(toSerialize)
}

type NullableContextUpdateRequest struct {
	value *ContextUpdateRequest
	isSet bool
}

func (v NullableContextUpdateRequest) Get() *ContextUpdateRequest {
	return v.value
}

func (v *NullableContextUpdateRequest) Set(val *ContextUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContextUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContextUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextUpdateRequest(val *ContextUpdateRequest) *NullableContextUpdateRequest {
	return &NullableContextUpdateRequest{value: val, isSet: true}
}

func (v NullableContextUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


