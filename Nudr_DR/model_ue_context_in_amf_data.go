/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// UeContextInAmfData struct for UeContextInAmfData
type UeContextInAmfData struct {
	EpsInterworkingInfo *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	// AMF information
	AmfInfo []AmfInfo `json:"amfInfo,omitempty"`
}

// NewUeContextInAmfData instantiates a new UeContextInAmfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeContextInAmfData() *UeContextInAmfData {
	this := UeContextInAmfData{}
	return &this
}

// NewUeContextInAmfDataWithDefaults instantiates a new UeContextInAmfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeContextInAmfDataWithDefaults() *UeContextInAmfData {
	this := UeContextInAmfData{}
	return &this
}

// GetEpsInterworkingInfo returns the EpsInterworkingInfo field value if set, zero value otherwise.
func (o *UeContextInAmfData) GetEpsInterworkingInfo() EpsInterworkingInfo {
	if o == nil || isNil(o.EpsInterworkingInfo) {
		var ret EpsInterworkingInfo
		return ret
	}
	return *o.EpsInterworkingInfo
}

// GetEpsInterworkingInfoOk returns a tuple with the EpsInterworkingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInAmfData) GetEpsInterworkingInfoOk() (*EpsInterworkingInfo, bool) {
	if o == nil || isNil(o.EpsInterworkingInfo) {
    return nil, false
	}
	return o.EpsInterworkingInfo, true
}

// HasEpsInterworkingInfo returns a boolean if a field has been set.
func (o *UeContextInAmfData) HasEpsInterworkingInfo() bool {
	if o != nil && !isNil(o.EpsInterworkingInfo) {
		return true
	}

	return false
}

// SetEpsInterworkingInfo gets a reference to the given EpsInterworkingInfo and assigns it to the EpsInterworkingInfo field.
func (o *UeContextInAmfData) SetEpsInterworkingInfo(v EpsInterworkingInfo) {
	o.EpsInterworkingInfo = &v
}

// GetAmfInfo returns the AmfInfo field value if set, zero value otherwise.
func (o *UeContextInAmfData) GetAmfInfo() []AmfInfo {
	if o == nil || isNil(o.AmfInfo) {
		var ret []AmfInfo
		return ret
	}
	return o.AmfInfo
}

// GetAmfInfoOk returns a tuple with the AmfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInAmfData) GetAmfInfoOk() ([]AmfInfo, bool) {
	if o == nil || isNil(o.AmfInfo) {
    return nil, false
	}
	return o.AmfInfo, true
}

// HasAmfInfo returns a boolean if a field has been set.
func (o *UeContextInAmfData) HasAmfInfo() bool {
	if o != nil && !isNil(o.AmfInfo) {
		return true
	}

	return false
}

// SetAmfInfo gets a reference to the given []AmfInfo and assigns it to the AmfInfo field.
func (o *UeContextInAmfData) SetAmfInfo(v []AmfInfo) {
	o.AmfInfo = v
}

func (o UeContextInAmfData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EpsInterworkingInfo) {
		toSerialize["epsInterworkingInfo"] = o.EpsInterworkingInfo
	}
	if !isNil(o.AmfInfo) {
		toSerialize["amfInfo"] = o.AmfInfo
	}
	return json.Marshal(toSerialize)
}

type NullableUeContextInAmfData struct {
	value *UeContextInAmfData
	isSet bool
}

func (v NullableUeContextInAmfData) Get() *UeContextInAmfData {
	return v.value
}

func (v *NullableUeContextInAmfData) Set(val *UeContextInAmfData) {
	v.value = val
	v.isSet = true
}

func (v NullableUeContextInAmfData) IsSet() bool {
	return v.isSet
}

func (v *NullableUeContextInAmfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeContextInAmfData(val *UeContextInAmfData) *NullableUeContextInAmfData {
	return &NullableUeContextInAmfData{value: val, isSet: true}
}

func (v NullableUeContextInAmfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeContextInAmfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


