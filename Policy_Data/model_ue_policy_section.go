/*
Unified Data Repository Service API file for policy data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Policy_Data

import (
	"encoding/json"
)

// UePolicySection Contains the UE policy section.
type UePolicySection struct {
	// string with format 'bytes' as defined in OpenAPI
	UePolicySectionInfo string `json:"uePolicySectionInfo"`
	Upsi string `json:"upsi"`
}

// NewUePolicySection instantiates a new UePolicySection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUePolicySection(uePolicySectionInfo string, upsi string) *UePolicySection {
	this := UePolicySection{}
	this.UePolicySectionInfo = uePolicySectionInfo
	this.Upsi = upsi
	return &this
}

// NewUePolicySectionWithDefaults instantiates a new UePolicySection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUePolicySectionWithDefaults() *UePolicySection {
	this := UePolicySection{}
	return &this
}

// GetUePolicySectionInfo returns the UePolicySectionInfo field value
func (o *UePolicySection) GetUePolicySectionInfo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UePolicySectionInfo
}

// GetUePolicySectionInfoOk returns a tuple with the UePolicySectionInfo field value
// and a boolean to check if the value has been set.
func (o *UePolicySection) GetUePolicySectionInfoOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UePolicySectionInfo, true
}

// SetUePolicySectionInfo sets field value
func (o *UePolicySection) SetUePolicySectionInfo(v string) {
	o.UePolicySectionInfo = v
}

// GetUpsi returns the Upsi field value
func (o *UePolicySection) GetUpsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Upsi
}

// GetUpsiOk returns a tuple with the Upsi field value
// and a boolean to check if the value has been set.
func (o *UePolicySection) GetUpsiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Upsi, true
}

// SetUpsi sets field value
func (o *UePolicySection) SetUpsi(v string) {
	o.Upsi = v
}

func (o UePolicySection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uePolicySectionInfo"] = o.UePolicySectionInfo
	}
	if true {
		toSerialize["upsi"] = o.Upsi
	}
	return json.Marshal(toSerialize)
}

type NullableUePolicySection struct {
	value *UePolicySection
	isSet bool
}

func (v NullableUePolicySection) Get() *UePolicySection {
	return v.value
}

func (v *NullableUePolicySection) Set(val *UePolicySection) {
	v.value = val
	v.isSet = true
}

func (v NullableUePolicySection) IsSet() bool {
	return v.isSet
}

func (v *NullableUePolicySection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUePolicySection(val *UePolicySection) *NullableUePolicySection {
	return &NullableUePolicySection{value: val, isSet: true}
}

func (v NullableUePolicySection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUePolicySection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


