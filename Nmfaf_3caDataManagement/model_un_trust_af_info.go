/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// UnTrustAfInfo Information of a untrusted AF Instance
type UnTrustAfInfo struct {
	AfId string `json:"afId"`
	SNssaiInfoList []SnssaiInfoItem `json:"sNssaiInfoList,omitempty"`
	MappingInd *bool `json:"mappingInd,omitempty"`
}

// NewUnTrustAfInfo instantiates a new UnTrustAfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnTrustAfInfo(afId string) *UnTrustAfInfo {
	this := UnTrustAfInfo{}
	this.AfId = afId
	var mappingInd bool = false
	this.MappingInd = &mappingInd
	return &this
}

// NewUnTrustAfInfoWithDefaults instantiates a new UnTrustAfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnTrustAfInfoWithDefaults() *UnTrustAfInfo {
	this := UnTrustAfInfo{}
	var mappingInd bool = false
	this.MappingInd = &mappingInd
	return &this
}

// GetAfId returns the AfId field value
func (o *UnTrustAfInfo) GetAfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value
// and a boolean to check if the value has been set.
func (o *UnTrustAfInfo) GetAfIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AfId, true
}

// SetAfId sets field value
func (o *UnTrustAfInfo) SetAfId(v string) {
	o.AfId = v
}

// GetSNssaiInfoList returns the SNssaiInfoList field value if set, zero value otherwise.
func (o *UnTrustAfInfo) GetSNssaiInfoList() []SnssaiInfoItem {
	if o == nil || isNil(o.SNssaiInfoList) {
		var ret []SnssaiInfoItem
		return ret
	}
	return o.SNssaiInfoList
}

// GetSNssaiInfoListOk returns a tuple with the SNssaiInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnTrustAfInfo) GetSNssaiInfoListOk() ([]SnssaiInfoItem, bool) {
	if o == nil || isNil(o.SNssaiInfoList) {
    return nil, false
	}
	return o.SNssaiInfoList, true
}

// HasSNssaiInfoList returns a boolean if a field has been set.
func (o *UnTrustAfInfo) HasSNssaiInfoList() bool {
	if o != nil && !isNil(o.SNssaiInfoList) {
		return true
	}

	return false
}

// SetSNssaiInfoList gets a reference to the given []SnssaiInfoItem and assigns it to the SNssaiInfoList field.
func (o *UnTrustAfInfo) SetSNssaiInfoList(v []SnssaiInfoItem) {
	o.SNssaiInfoList = v
}

// GetMappingInd returns the MappingInd field value if set, zero value otherwise.
func (o *UnTrustAfInfo) GetMappingInd() bool {
	if o == nil || isNil(o.MappingInd) {
		var ret bool
		return ret
	}
	return *o.MappingInd
}

// GetMappingIndOk returns a tuple with the MappingInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnTrustAfInfo) GetMappingIndOk() (*bool, bool) {
	if o == nil || isNil(o.MappingInd) {
    return nil, false
	}
	return o.MappingInd, true
}

// HasMappingInd returns a boolean if a field has been set.
func (o *UnTrustAfInfo) HasMappingInd() bool {
	if o != nil && !isNil(o.MappingInd) {
		return true
	}

	return false
}

// SetMappingInd gets a reference to the given bool and assigns it to the MappingInd field.
func (o *UnTrustAfInfo) SetMappingInd(v bool) {
	o.MappingInd = &v
}

func (o UnTrustAfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["afId"] = o.AfId
	}
	if !isNil(o.SNssaiInfoList) {
		toSerialize["sNssaiInfoList"] = o.SNssaiInfoList
	}
	if !isNil(o.MappingInd) {
		toSerialize["mappingInd"] = o.MappingInd
	}
	return json.Marshal(toSerialize)
}

type NullableUnTrustAfInfo struct {
	value *UnTrustAfInfo
	isSet bool
}

func (v NullableUnTrustAfInfo) Get() *UnTrustAfInfo {
	return v.value
}

func (v *NullableUnTrustAfInfo) Set(val *UnTrustAfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUnTrustAfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUnTrustAfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnTrustAfInfo(val *UnTrustAfInfo) *NullableUnTrustAfInfo {
	return &NullableUnTrustAfInfo{value: val, isSet: true}
}

func (v NullableUnTrustAfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnTrustAfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


