/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
)

// CongestInfo Represents a UE's user data congestion information.
type CongestInfo struct {
	LocArea LocationArea5G `json:"locArea"`
	CngAnas []CongestionAnalytics `json:"cngAnas"`
}

// NewCongestInfo instantiates a new CongestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCongestInfo(locArea LocationArea5G, cngAnas []CongestionAnalytics) *CongestInfo {
	this := CongestInfo{}
	this.LocArea = locArea
	this.CngAnas = cngAnas
	return &this
}

// NewCongestInfoWithDefaults instantiates a new CongestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCongestInfoWithDefaults() *CongestInfo {
	this := CongestInfo{}
	return &this
}

// GetLocArea returns the LocArea field value
func (o *CongestInfo) GetLocArea() LocationArea5G {
	if o == nil {
		var ret LocationArea5G
		return ret
	}

	return o.LocArea
}

// GetLocAreaOk returns a tuple with the LocArea field value
// and a boolean to check if the value has been set.
func (o *CongestInfo) GetLocAreaOk() (*LocationArea5G, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocArea, true
}

// SetLocArea sets field value
func (o *CongestInfo) SetLocArea(v LocationArea5G) {
	o.LocArea = v
}

// GetCngAnas returns the CngAnas field value
func (o *CongestInfo) GetCngAnas() []CongestionAnalytics {
	if o == nil {
		var ret []CongestionAnalytics
		return ret
	}

	return o.CngAnas
}

// GetCngAnasOk returns a tuple with the CngAnas field value
// and a boolean to check if the value has been set.
func (o *CongestInfo) GetCngAnasOk() ([]CongestionAnalytics, bool) {
	if o == nil {
    return nil, false
	}
	return o.CngAnas, true
}

// SetCngAnas sets field value
func (o *CongestInfo) SetCngAnas(v []CongestionAnalytics) {
	o.CngAnas = v
}

func (o CongestInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["locArea"] = o.LocArea
	}
	if true {
		toSerialize["cngAnas"] = o.CngAnas
	}
	return json.Marshal(toSerialize)
}

type NullableCongestInfo struct {
	value *CongestInfo
	isSet bool
}

func (v NullableCongestInfo) Get() *CongestInfo {
	return v.value
}

func (v *NullableCongestInfo) Set(val *CongestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCongestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCongestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCongestInfo(val *CongestInfo) *NullableCongestInfo {
	return &NullableCongestInfo{value: val, isSet: true}
}

func (v NullableCongestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCongestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


