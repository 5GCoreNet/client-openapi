/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// WlanPerSsIdPerformanceInfo The WLAN performance per SSID.
type WlanPerSsIdPerformanceInfo struct {
	SsId string `json:"ssId"`
	WlanPerTsInfos []WlanPerTsPerformanceInfo `json:"wlanPerTsInfos"`
}

// NewWlanPerSsIdPerformanceInfo instantiates a new WlanPerSsIdPerformanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWlanPerSsIdPerformanceInfo(ssId string, wlanPerTsInfos []WlanPerTsPerformanceInfo) *WlanPerSsIdPerformanceInfo {
	this := WlanPerSsIdPerformanceInfo{}
	this.SsId = ssId
	this.WlanPerTsInfos = wlanPerTsInfos
	return &this
}

// NewWlanPerSsIdPerformanceInfoWithDefaults instantiates a new WlanPerSsIdPerformanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWlanPerSsIdPerformanceInfoWithDefaults() *WlanPerSsIdPerformanceInfo {
	this := WlanPerSsIdPerformanceInfo{}
	return &this
}

// GetSsId returns the SsId field value
func (o *WlanPerSsIdPerformanceInfo) GetSsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SsId
}

// GetSsIdOk returns a tuple with the SsId field value
// and a boolean to check if the value has been set.
func (o *WlanPerSsIdPerformanceInfo) GetSsIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SsId, true
}

// SetSsId sets field value
func (o *WlanPerSsIdPerformanceInfo) SetSsId(v string) {
	o.SsId = v
}

// GetWlanPerTsInfos returns the WlanPerTsInfos field value
func (o *WlanPerSsIdPerformanceInfo) GetWlanPerTsInfos() []WlanPerTsPerformanceInfo {
	if o == nil {
		var ret []WlanPerTsPerformanceInfo
		return ret
	}

	return o.WlanPerTsInfos
}

// GetWlanPerTsInfosOk returns a tuple with the WlanPerTsInfos field value
// and a boolean to check if the value has been set.
func (o *WlanPerSsIdPerformanceInfo) GetWlanPerTsInfosOk() ([]WlanPerTsPerformanceInfo, bool) {
	if o == nil {
    return nil, false
	}
	return o.WlanPerTsInfos, true
}

// SetWlanPerTsInfos sets field value
func (o *WlanPerSsIdPerformanceInfo) SetWlanPerTsInfos(v []WlanPerTsPerformanceInfo) {
	o.WlanPerTsInfos = v
}

func (o WlanPerSsIdPerformanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ssId"] = o.SsId
	}
	if true {
		toSerialize["wlanPerTsInfos"] = o.WlanPerTsInfos
	}
	return json.Marshal(toSerialize)
}

type NullableWlanPerSsIdPerformanceInfo struct {
	value *WlanPerSsIdPerformanceInfo
	isSet bool
}

func (v NullableWlanPerSsIdPerformanceInfo) Get() *WlanPerSsIdPerformanceInfo {
	return v.value
}

func (v *NullableWlanPerSsIdPerformanceInfo) Set(val *WlanPerSsIdPerformanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanPerSsIdPerformanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanPerSsIdPerformanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanPerSsIdPerformanceInfo(val *WlanPerSsIdPerformanceInfo) *NullableWlanPerSsIdPerformanceInfo {
	return &NullableWlanPerSsIdPerformanceInfo{value: val, isSet: true}
}

func (v NullableWlanPerSsIdPerformanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanPerSsIdPerformanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


