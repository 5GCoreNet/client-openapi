/*
Namf_MBSCommunication

AMF Communication Service for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSCommunication

import (
	"encoding/json"
)

// N2MbsSmInfo N2 MBS Session Management information
type N2MbsSmInfo struct {
	NgapIeType MbsNgapIeType `json:"ngapIeType"`
	NgapData RefToBinaryData `json:"ngapData"`
}

// NewN2MbsSmInfo instantiates a new N2MbsSmInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN2MbsSmInfo(ngapIeType MbsNgapIeType, ngapData RefToBinaryData) *N2MbsSmInfo {
	this := N2MbsSmInfo{}
	this.NgapIeType = ngapIeType
	this.NgapData = ngapData
	return &this
}

// NewN2MbsSmInfoWithDefaults instantiates a new N2MbsSmInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN2MbsSmInfoWithDefaults() *N2MbsSmInfo {
	this := N2MbsSmInfo{}
	return &this
}

// GetNgapIeType returns the NgapIeType field value
func (o *N2MbsSmInfo) GetNgapIeType() MbsNgapIeType {
	if o == nil {
		var ret MbsNgapIeType
		return ret
	}

	return o.NgapIeType
}

// GetNgapIeTypeOk returns a tuple with the NgapIeType field value
// and a boolean to check if the value has been set.
func (o *N2MbsSmInfo) GetNgapIeTypeOk() (*MbsNgapIeType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NgapIeType, true
}

// SetNgapIeType sets field value
func (o *N2MbsSmInfo) SetNgapIeType(v MbsNgapIeType) {
	o.NgapIeType = v
}

// GetNgapData returns the NgapData field value
func (o *N2MbsSmInfo) GetNgapData() RefToBinaryData {
	if o == nil {
		var ret RefToBinaryData
		return ret
	}

	return o.NgapData
}

// GetNgapDataOk returns a tuple with the NgapData field value
// and a boolean to check if the value has been set.
func (o *N2MbsSmInfo) GetNgapDataOk() (*RefToBinaryData, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NgapData, true
}

// SetNgapData sets field value
func (o *N2MbsSmInfo) SetNgapData(v RefToBinaryData) {
	o.NgapData = v
}

func (o N2MbsSmInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ngapIeType"] = o.NgapIeType
	}
	if true {
		toSerialize["ngapData"] = o.NgapData
	}
	return json.Marshal(toSerialize)
}

type NullableN2MbsSmInfo struct {
	value *N2MbsSmInfo
	isSet bool
}

func (v NullableN2MbsSmInfo) Get() *N2MbsSmInfo {
	return v.value
}

func (v *NullableN2MbsSmInfo) Set(val *N2MbsSmInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableN2MbsSmInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableN2MbsSmInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2MbsSmInfo(val *N2MbsSmInfo) *NullableN2MbsSmInfo {
	return &NullableN2MbsSmInfo{value: val, isSet: true}
}

func (v NullableN2MbsSmInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2MbsSmInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


