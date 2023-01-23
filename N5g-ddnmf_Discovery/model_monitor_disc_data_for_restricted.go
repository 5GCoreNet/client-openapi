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

// MonitorDiscDataForRestricted Represents Data for restricted discovery used to request the authorization to monitor for a UE
type MonitorDiscDataForRestricted struct {
	// Contains the RPAUID.
	Rpauid string `json:"rpauid"`
	// Contains the PDUID.
	TargetPduid string `json:"targetPduid"`
	// Contains the Application ID.
	AppId string `json:"appId"`
	// Contains the RPAUID.
	TargetRpauid string `json:"targetRpauid"`
}

// NewMonitorDiscDataForRestricted instantiates a new MonitorDiscDataForRestricted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorDiscDataForRestricted(rpauid string, targetPduid string, appId string, targetRpauid string) *MonitorDiscDataForRestricted {
	this := MonitorDiscDataForRestricted{}
	this.Rpauid = rpauid
	this.TargetPduid = targetPduid
	this.AppId = appId
	this.TargetRpauid = targetRpauid
	return &this
}

// NewMonitorDiscDataForRestrictedWithDefaults instantiates a new MonitorDiscDataForRestricted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorDiscDataForRestrictedWithDefaults() *MonitorDiscDataForRestricted {
	this := MonitorDiscDataForRestricted{}
	return &this
}

// GetRpauid returns the Rpauid field value
func (o *MonitorDiscDataForRestricted) GetRpauid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rpauid
}

// GetRpauidOk returns a tuple with the Rpauid field value
// and a boolean to check if the value has been set.
func (o *MonitorDiscDataForRestricted) GetRpauidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Rpauid, true
}

// SetRpauid sets field value
func (o *MonitorDiscDataForRestricted) SetRpauid(v string) {
	o.Rpauid = v
}

// GetTargetPduid returns the TargetPduid field value
func (o *MonitorDiscDataForRestricted) GetTargetPduid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetPduid
}

// GetTargetPduidOk returns a tuple with the TargetPduid field value
// and a boolean to check if the value has been set.
func (o *MonitorDiscDataForRestricted) GetTargetPduidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TargetPduid, true
}

// SetTargetPduid sets field value
func (o *MonitorDiscDataForRestricted) SetTargetPduid(v string) {
	o.TargetPduid = v
}

// GetAppId returns the AppId field value
func (o *MonitorDiscDataForRestricted) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *MonitorDiscDataForRestricted) GetAppIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *MonitorDiscDataForRestricted) SetAppId(v string) {
	o.AppId = v
}

// GetTargetRpauid returns the TargetRpauid field value
func (o *MonitorDiscDataForRestricted) GetTargetRpauid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetRpauid
}

// GetTargetRpauidOk returns a tuple with the TargetRpauid field value
// and a boolean to check if the value has been set.
func (o *MonitorDiscDataForRestricted) GetTargetRpauidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TargetRpauid, true
}

// SetTargetRpauid sets field value
func (o *MonitorDiscDataForRestricted) SetTargetRpauid(v string) {
	o.TargetRpauid = v
}

func (o MonitorDiscDataForRestricted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rpauid"] = o.Rpauid
	}
	if true {
		toSerialize["targetPduid"] = o.TargetPduid
	}
	if true {
		toSerialize["appId"] = o.AppId
	}
	if true {
		toSerialize["targetRpauid"] = o.TargetRpauid
	}
	return json.Marshal(toSerialize)
}

type NullableMonitorDiscDataForRestricted struct {
	value *MonitorDiscDataForRestricted
	isSet bool
}

func (v NullableMonitorDiscDataForRestricted) Get() *MonitorDiscDataForRestricted {
	return v.value
}

func (v *NullableMonitorDiscDataForRestricted) Set(val *MonitorDiscDataForRestricted) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorDiscDataForRestricted) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorDiscDataForRestricted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorDiscDataForRestricted(val *MonitorDiscDataForRestricted) *NullableMonitorDiscDataForRestricted {
	return &NullableMonitorDiscDataForRestricted{value: val, isSet: true}
}

func (v NullableMonitorDiscDataForRestricted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorDiscDataForRestricted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


