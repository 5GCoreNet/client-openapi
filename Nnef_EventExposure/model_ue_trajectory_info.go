/*
Nnef_EventExposure

NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_EventExposure

import (
	"encoding/json"
	"time"
)

// UeTrajectoryInfo Contains UE trajectory information.
type UeTrajectoryInfo struct {
	// string with format 'date-time' as defined in OpenAPI.
	Ts time.Time `json:"ts"`
	Location UserLocation `json:"location"`
}

// NewUeTrajectoryInfo instantiates a new UeTrajectoryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeTrajectoryInfo(ts time.Time, location UserLocation) *UeTrajectoryInfo {
	this := UeTrajectoryInfo{}
	this.Ts = ts
	this.Location = location
	return &this
}

// NewUeTrajectoryInfoWithDefaults instantiates a new UeTrajectoryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeTrajectoryInfoWithDefaults() *UeTrajectoryInfo {
	this := UeTrajectoryInfo{}
	return &this
}

// GetTs returns the Ts field value
func (o *UeTrajectoryInfo) GetTs() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Ts
}

// GetTsOk returns a tuple with the Ts field value
// and a boolean to check if the value has been set.
func (o *UeTrajectoryInfo) GetTsOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Ts, true
}

// SetTs sets field value
func (o *UeTrajectoryInfo) SetTs(v time.Time) {
	o.Ts = v
}

// GetLocation returns the Location field value
func (o *UeTrajectoryInfo) GetLocation() UserLocation {
	if o == nil {
		var ret UserLocation
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *UeTrajectoryInfo) GetLocationOk() (*UserLocation, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *UeTrajectoryInfo) SetLocation(v UserLocation) {
	o.Location = v
}

func (o UeTrajectoryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ts"] = o.Ts
	}
	if true {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableUeTrajectoryInfo struct {
	value *UeTrajectoryInfo
	isSet bool
}

func (v NullableUeTrajectoryInfo) Get() *UeTrajectoryInfo {
	return v.value
}

func (v *NullableUeTrajectoryInfo) Set(val *UeTrajectoryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUeTrajectoryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUeTrajectoryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeTrajectoryInfo(val *UeTrajectoryInfo) *NullableUeTrajectoryInfo {
	return &NullableUeTrajectoryInfo{value: val, isSet: true}
}

func (v NullableUeTrajectoryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeTrajectoryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


