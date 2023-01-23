/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// SnssaiMbSmfInfoItem Parameters supported by an MB-SMF for a given S-NSSAI
type SnssaiMbSmfInfoItem struct {
	SNssai ExtSnssai `json:"sNssai"`
	DnnInfoList []DnnMbSmfInfoItem `json:"dnnInfoList"`
}

// NewSnssaiMbSmfInfoItem instantiates a new SnssaiMbSmfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiMbSmfInfoItem(sNssai ExtSnssai, dnnInfoList []DnnMbSmfInfoItem) *SnssaiMbSmfInfoItem {
	this := SnssaiMbSmfInfoItem{}
	this.SNssai = sNssai
	this.DnnInfoList = dnnInfoList
	return &this
}

// NewSnssaiMbSmfInfoItemWithDefaults instantiates a new SnssaiMbSmfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiMbSmfInfoItemWithDefaults() *SnssaiMbSmfInfoItem {
	this := SnssaiMbSmfInfoItem{}
	return &this
}

// GetSNssai returns the SNssai field value
func (o *SnssaiMbSmfInfoItem) GetSNssai() ExtSnssai {
	if o == nil {
		var ret ExtSnssai
		return ret
	}

	return o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value
// and a boolean to check if the value has been set.
func (o *SnssaiMbSmfInfoItem) GetSNssaiOk() (*ExtSnssai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SNssai, true
}

// SetSNssai sets field value
func (o *SnssaiMbSmfInfoItem) SetSNssai(v ExtSnssai) {
	o.SNssai = v
}

// GetDnnInfoList returns the DnnInfoList field value
func (o *SnssaiMbSmfInfoItem) GetDnnInfoList() []DnnMbSmfInfoItem {
	if o == nil {
		var ret []DnnMbSmfInfoItem
		return ret
	}

	return o.DnnInfoList
}

// GetDnnInfoListOk returns a tuple with the DnnInfoList field value
// and a boolean to check if the value has been set.
func (o *SnssaiMbSmfInfoItem) GetDnnInfoListOk() ([]DnnMbSmfInfoItem, bool) {
	if o == nil {
    return nil, false
	}
	return o.DnnInfoList, true
}

// SetDnnInfoList sets field value
func (o *SnssaiMbSmfInfoItem) SetDnnInfoList(v []DnnMbSmfInfoItem) {
	o.DnnInfoList = v
}

func (o SnssaiMbSmfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sNssai"] = o.SNssai
	}
	if true {
		toSerialize["dnnInfoList"] = o.DnnInfoList
	}
	return json.Marshal(toSerialize)
}

type NullableSnssaiMbSmfInfoItem struct {
	value *SnssaiMbSmfInfoItem
	isSet bool
}

func (v NullableSnssaiMbSmfInfoItem) Get() *SnssaiMbSmfInfoItem {
	return v.value
}

func (v *NullableSnssaiMbSmfInfoItem) Set(val *SnssaiMbSmfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiMbSmfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiMbSmfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiMbSmfInfoItem(val *SnssaiMbSmfInfoItem) *NullableSnssaiMbSmfInfoItem {
	return &NullableSnssaiMbSmfInfoItem{value: val, isSet: true}
}

func (v NullableSnssaiMbSmfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiMbSmfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


