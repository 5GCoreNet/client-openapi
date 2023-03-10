/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"os"
)

// PostSmContextsRequest struct for PostSmContextsRequest
type PostSmContextsRequest struct {
	JsonData *SmContextCreateData `json:"jsonData,omitempty"`
	BinaryDataN1SmMessage **os.File `json:"binaryDataN1SmMessage,omitempty"`
	BinaryDataN2SmInformation **os.File `json:"binaryDataN2SmInformation,omitempty"`
	BinaryDataN2SmInformationExt1 **os.File `json:"binaryDataN2SmInformationExt1,omitempty"`
}

// NewPostSmContextsRequest instantiates a new PostSmContextsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSmContextsRequest() *PostSmContextsRequest {
	this := PostSmContextsRequest{}
	return &this
}

// NewPostSmContextsRequestWithDefaults instantiates a new PostSmContextsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSmContextsRequestWithDefaults() *PostSmContextsRequest {
	this := PostSmContextsRequest{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *PostSmContextsRequest) GetJsonData() SmContextCreateData {
	if o == nil || isNil(o.JsonData) {
		var ret SmContextCreateData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSmContextsRequest) GetJsonDataOk() (*SmContextCreateData, bool) {
	if o == nil || isNil(o.JsonData) {
    return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *PostSmContextsRequest) HasJsonData() bool {
	if o != nil && !isNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given SmContextCreateData and assigns it to the JsonData field.
func (o *PostSmContextsRequest) SetJsonData(v SmContextCreateData) {
	o.JsonData = &v
}

// GetBinaryDataN1SmMessage returns the BinaryDataN1SmMessage field value if set, zero value otherwise.
func (o *PostSmContextsRequest) GetBinaryDataN1SmMessage() *os.File {
	if o == nil || isNil(o.BinaryDataN1SmMessage) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN1SmMessage
}

// GetBinaryDataN1SmMessageOk returns a tuple with the BinaryDataN1SmMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSmContextsRequest) GetBinaryDataN1SmMessageOk() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN1SmMessage) {
    return nil, false
	}
	return o.BinaryDataN1SmMessage, true
}

// HasBinaryDataN1SmMessage returns a boolean if a field has been set.
func (o *PostSmContextsRequest) HasBinaryDataN1SmMessage() bool {
	if o != nil && !isNil(o.BinaryDataN1SmMessage) {
		return true
	}

	return false
}

// SetBinaryDataN1SmMessage gets a reference to the given *os.File and assigns it to the BinaryDataN1SmMessage field.
func (o *PostSmContextsRequest) SetBinaryDataN1SmMessage(v *os.File) {
	o.BinaryDataN1SmMessage = &v
}

// GetBinaryDataN2SmInformation returns the BinaryDataN2SmInformation field value if set, zero value otherwise.
func (o *PostSmContextsRequest) GetBinaryDataN2SmInformation() *os.File {
	if o == nil || isNil(o.BinaryDataN2SmInformation) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2SmInformation
}

// GetBinaryDataN2SmInformationOk returns a tuple with the BinaryDataN2SmInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSmContextsRequest) GetBinaryDataN2SmInformationOk() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2SmInformation) {
    return nil, false
	}
	return o.BinaryDataN2SmInformation, true
}

// HasBinaryDataN2SmInformation returns a boolean if a field has been set.
func (o *PostSmContextsRequest) HasBinaryDataN2SmInformation() bool {
	if o != nil && !isNil(o.BinaryDataN2SmInformation) {
		return true
	}

	return false
}

// SetBinaryDataN2SmInformation gets a reference to the given *os.File and assigns it to the BinaryDataN2SmInformation field.
func (o *PostSmContextsRequest) SetBinaryDataN2SmInformation(v *os.File) {
	o.BinaryDataN2SmInformation = &v
}

// GetBinaryDataN2SmInformationExt1 returns the BinaryDataN2SmInformationExt1 field value if set, zero value otherwise.
func (o *PostSmContextsRequest) GetBinaryDataN2SmInformationExt1() *os.File {
	if o == nil || isNil(o.BinaryDataN2SmInformationExt1) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2SmInformationExt1
}

// GetBinaryDataN2SmInformationExt1Ok returns a tuple with the BinaryDataN2SmInformationExt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSmContextsRequest) GetBinaryDataN2SmInformationExt1Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2SmInformationExt1) {
    return nil, false
	}
	return o.BinaryDataN2SmInformationExt1, true
}

// HasBinaryDataN2SmInformationExt1 returns a boolean if a field has been set.
func (o *PostSmContextsRequest) HasBinaryDataN2SmInformationExt1() bool {
	if o != nil && !isNil(o.BinaryDataN2SmInformationExt1) {
		return true
	}

	return false
}

// SetBinaryDataN2SmInformationExt1 gets a reference to the given *os.File and assigns it to the BinaryDataN2SmInformationExt1 field.
func (o *PostSmContextsRequest) SetBinaryDataN2SmInformationExt1(v *os.File) {
	o.BinaryDataN2SmInformationExt1 = &v
}

func (o PostSmContextsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !isNil(o.BinaryDataN1SmMessage) {
		toSerialize["binaryDataN1SmMessage"] = o.BinaryDataN1SmMessage
	}
	if !isNil(o.BinaryDataN2SmInformation) {
		toSerialize["binaryDataN2SmInformation"] = o.BinaryDataN2SmInformation
	}
	if !isNil(o.BinaryDataN2SmInformationExt1) {
		toSerialize["binaryDataN2SmInformationExt1"] = o.BinaryDataN2SmInformationExt1
	}
	return json.Marshal(toSerialize)
}

type NullablePostSmContextsRequest struct {
	value *PostSmContextsRequest
	isSet bool
}

func (v NullablePostSmContextsRequest) Get() *PostSmContextsRequest {
	return v.value
}

func (v *NullablePostSmContextsRequest) Set(val *PostSmContextsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSmContextsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSmContextsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSmContextsRequest(val *PostSmContextsRequest) *NullablePostSmContextsRequest {
	return &NullablePostSmContextsRequest{value: val, isSet: true}
}

func (v NullablePostSmContextsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSmContextsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


