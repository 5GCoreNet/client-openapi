/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// CagInfo struct for CagInfo
type CagInfo struct {
	AllowedCagList []string `json:"allowedCagList"`
	CagOnlyIndicator *bool `json:"cagOnlyIndicator,omitempty"`
}

// NewCagInfo instantiates a new CagInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCagInfo(allowedCagList []string) *CagInfo {
	this := CagInfo{}
	this.AllowedCagList = allowedCagList
	return &this
}

// NewCagInfoWithDefaults instantiates a new CagInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCagInfoWithDefaults() *CagInfo {
	this := CagInfo{}
	return &this
}

// GetAllowedCagList returns the AllowedCagList field value
func (o *CagInfo) GetAllowedCagList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedCagList
}

// GetAllowedCagListOk returns a tuple with the AllowedCagList field value
// and a boolean to check if the value has been set.
func (o *CagInfo) GetAllowedCagListOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowedCagList, true
}

// SetAllowedCagList sets field value
func (o *CagInfo) SetAllowedCagList(v []string) {
	o.AllowedCagList = v
}

// GetCagOnlyIndicator returns the CagOnlyIndicator field value if set, zero value otherwise.
func (o *CagInfo) GetCagOnlyIndicator() bool {
	if o == nil || isNil(o.CagOnlyIndicator) {
		var ret bool
		return ret
	}
	return *o.CagOnlyIndicator
}

// GetCagOnlyIndicatorOk returns a tuple with the CagOnlyIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CagInfo) GetCagOnlyIndicatorOk() (*bool, bool) {
	if o == nil || isNil(o.CagOnlyIndicator) {
    return nil, false
	}
	return o.CagOnlyIndicator, true
}

// HasCagOnlyIndicator returns a boolean if a field has been set.
func (o *CagInfo) HasCagOnlyIndicator() bool {
	if o != nil && !isNil(o.CagOnlyIndicator) {
		return true
	}

	return false
}

// SetCagOnlyIndicator gets a reference to the given bool and assigns it to the CagOnlyIndicator field.
func (o *CagInfo) SetCagOnlyIndicator(v bool) {
	o.CagOnlyIndicator = &v
}

func (o CagInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allowedCagList"] = o.AllowedCagList
	}
	if !isNil(o.CagOnlyIndicator) {
		toSerialize["cagOnlyIndicator"] = o.CagOnlyIndicator
	}
	return json.Marshal(toSerialize)
}

type NullableCagInfo struct {
	value *CagInfo
	isSet bool
}

func (v NullableCagInfo) Get() *CagInfo {
	return v.value
}

func (v *NullableCagInfo) Set(val *CagInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCagInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCagInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCagInfo(val *CagInfo) *NullableCagInfo {
	return &NullableCagInfo{value: val, isSet: true}
}

func (v NullableCagInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCagInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


