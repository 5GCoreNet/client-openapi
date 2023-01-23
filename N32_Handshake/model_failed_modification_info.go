/*
N32 Handshake API

N32-c Handshake Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N32_Handshake

import (
	"encoding/json"
)

// FailedModificationInfo Information on N32-f modifications block that failed to process
type FailedModificationInfo struct {
	// Fully Qualified Domain Name
	IpxId string `json:"ipxId"`
	N32fErrorType N32fErrorType `json:"n32fErrorType"`
}

// NewFailedModificationInfo instantiates a new FailedModificationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailedModificationInfo(ipxId string, n32fErrorType N32fErrorType) *FailedModificationInfo {
	this := FailedModificationInfo{}
	this.IpxId = ipxId
	this.N32fErrorType = n32fErrorType
	return &this
}

// NewFailedModificationInfoWithDefaults instantiates a new FailedModificationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailedModificationInfoWithDefaults() *FailedModificationInfo {
	this := FailedModificationInfo{}
	return &this
}

// GetIpxId returns the IpxId field value
func (o *FailedModificationInfo) GetIpxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpxId
}

// GetIpxIdOk returns a tuple with the IpxId field value
// and a boolean to check if the value has been set.
func (o *FailedModificationInfo) GetIpxIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IpxId, true
}

// SetIpxId sets field value
func (o *FailedModificationInfo) SetIpxId(v string) {
	o.IpxId = v
}

// GetN32fErrorType returns the N32fErrorType field value
func (o *FailedModificationInfo) GetN32fErrorType() N32fErrorType {
	if o == nil {
		var ret N32fErrorType
		return ret
	}

	return o.N32fErrorType
}

// GetN32fErrorTypeOk returns a tuple with the N32fErrorType field value
// and a boolean to check if the value has been set.
func (o *FailedModificationInfo) GetN32fErrorTypeOk() (*N32fErrorType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.N32fErrorType, true
}

// SetN32fErrorType sets field value
func (o *FailedModificationInfo) SetN32fErrorType(v N32fErrorType) {
	o.N32fErrorType = v
}

func (o FailedModificationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ipxId"] = o.IpxId
	}
	if true {
		toSerialize["n32fErrorType"] = o.N32fErrorType
	}
	return json.Marshal(toSerialize)
}

type NullableFailedModificationInfo struct {
	value *FailedModificationInfo
	isSet bool
}

func (v NullableFailedModificationInfo) Get() *FailedModificationInfo {
	return v.value
}

func (v *NullableFailedModificationInfo) Set(val *FailedModificationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFailedModificationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFailedModificationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailedModificationInfo(val *FailedModificationInfo) *NullableFailedModificationInfo {
	return &NullableFailedModificationInfo{value: val, isSet: true}
}

func (v NullableFailedModificationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailedModificationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


