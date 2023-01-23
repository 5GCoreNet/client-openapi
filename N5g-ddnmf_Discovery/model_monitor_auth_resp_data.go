/*
N5g-ddnmf_Discovery API

N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N5g-ddnmf_Discovery

import (
	"encoding/json"
)

// MonitorAuthRespData Represents the obtained Monitor Authorize Data for a UE
type MonitorAuthRespData struct {
	AuthDataOpen *MonitorAuthDataForOpen `json:"authDataOpen,omitempty"`
	AuthDataRestricted *MonitorAuthDataForRestricted `json:"authDataRestricted,omitempty"`
}

// NewMonitorAuthRespData instantiates a new MonitorAuthRespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorAuthRespData() *MonitorAuthRespData {
	this := MonitorAuthRespData{}
	return &this
}

// NewMonitorAuthRespDataWithDefaults instantiates a new MonitorAuthRespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorAuthRespDataWithDefaults() *MonitorAuthRespData {
	this := MonitorAuthRespData{}
	return &this
}

// GetAuthDataOpen returns the AuthDataOpen field value if set, zero value otherwise.
func (o *MonitorAuthRespData) GetAuthDataOpen() MonitorAuthDataForOpen {
	if o == nil || isNil(o.AuthDataOpen) {
		var ret MonitorAuthDataForOpen
		return ret
	}
	return *o.AuthDataOpen
}

// GetAuthDataOpenOk returns a tuple with the AuthDataOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorAuthRespData) GetAuthDataOpenOk() (*MonitorAuthDataForOpen, bool) {
	if o == nil || isNil(o.AuthDataOpen) {
    return nil, false
	}
	return o.AuthDataOpen, true
}

// HasAuthDataOpen returns a boolean if a field has been set.
func (o *MonitorAuthRespData) HasAuthDataOpen() bool {
	if o != nil && !isNil(o.AuthDataOpen) {
		return true
	}

	return false
}

// SetAuthDataOpen gets a reference to the given MonitorAuthDataForOpen and assigns it to the AuthDataOpen field.
func (o *MonitorAuthRespData) SetAuthDataOpen(v MonitorAuthDataForOpen) {
	o.AuthDataOpen = &v
}

// GetAuthDataRestricted returns the AuthDataRestricted field value if set, zero value otherwise.
func (o *MonitorAuthRespData) GetAuthDataRestricted() MonitorAuthDataForRestricted {
	if o == nil || isNil(o.AuthDataRestricted) {
		var ret MonitorAuthDataForRestricted
		return ret
	}
	return *o.AuthDataRestricted
}

// GetAuthDataRestrictedOk returns a tuple with the AuthDataRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorAuthRespData) GetAuthDataRestrictedOk() (*MonitorAuthDataForRestricted, bool) {
	if o == nil || isNil(o.AuthDataRestricted) {
    return nil, false
	}
	return o.AuthDataRestricted, true
}

// HasAuthDataRestricted returns a boolean if a field has been set.
func (o *MonitorAuthRespData) HasAuthDataRestricted() bool {
	if o != nil && !isNil(o.AuthDataRestricted) {
		return true
	}

	return false
}

// SetAuthDataRestricted gets a reference to the given MonitorAuthDataForRestricted and assigns it to the AuthDataRestricted field.
func (o *MonitorAuthRespData) SetAuthDataRestricted(v MonitorAuthDataForRestricted) {
	o.AuthDataRestricted = &v
}

func (o MonitorAuthRespData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AuthDataOpen) {
		toSerialize["authDataOpen"] = o.AuthDataOpen
	}
	if !isNil(o.AuthDataRestricted) {
		toSerialize["authDataRestricted"] = o.AuthDataRestricted
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorAuthRespData struct {
	value *MonitorAuthRespData
	isSet bool
}

func (v NullableMonitorAuthRespData) Get() *MonitorAuthRespData {
	return v.value
}

func (v *NullableMonitorAuthRespData) Set(val *MonitorAuthRespData) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorAuthRespData) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorAuthRespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorAuthRespData(val *MonitorAuthRespData) *NullableMonitorAuthRespData {
	return &NullableMonitorAuthRespData{value: val, isSet: true}
}

func (v NullableMonitorAuthRespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorAuthRespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


