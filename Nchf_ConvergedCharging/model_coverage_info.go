/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"time"
)

// CoverageInfo struct for CoverageInfo
type CoverageInfo struct {
	CoverageStatus *bool `json:"coverageStatus,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ChangeTime *time.Time `json:"changeTime,omitempty"`
	LocationInfo []UserLocation `json:"locationInfo,omitempty"`
}

// NewCoverageInfo instantiates a new CoverageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoverageInfo() *CoverageInfo {
	this := CoverageInfo{}
	return &this
}

// NewCoverageInfoWithDefaults instantiates a new CoverageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoverageInfoWithDefaults() *CoverageInfo {
	this := CoverageInfo{}
	return &this
}

// GetCoverageStatus returns the CoverageStatus field value if set, zero value otherwise.
func (o *CoverageInfo) GetCoverageStatus() bool {
	if o == nil || isNil(o.CoverageStatus) {
		var ret bool
		return ret
	}
	return *o.CoverageStatus
}

// GetCoverageStatusOk returns a tuple with the CoverageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageInfo) GetCoverageStatusOk() (*bool, bool) {
	if o == nil || isNil(o.CoverageStatus) {
    return nil, false
	}
	return o.CoverageStatus, true
}

// HasCoverageStatus returns a boolean if a field has been set.
func (o *CoverageInfo) HasCoverageStatus() bool {
	if o != nil && !isNil(o.CoverageStatus) {
		return true
	}

	return false
}

// SetCoverageStatus gets a reference to the given bool and assigns it to the CoverageStatus field.
func (o *CoverageInfo) SetCoverageStatus(v bool) {
	o.CoverageStatus = &v
}

// GetChangeTime returns the ChangeTime field value if set, zero value otherwise.
func (o *CoverageInfo) GetChangeTime() time.Time {
	if o == nil || isNil(o.ChangeTime) {
		var ret time.Time
		return ret
	}
	return *o.ChangeTime
}

// GetChangeTimeOk returns a tuple with the ChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageInfo) GetChangeTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ChangeTime) {
    return nil, false
	}
	return o.ChangeTime, true
}

// HasChangeTime returns a boolean if a field has been set.
func (o *CoverageInfo) HasChangeTime() bool {
	if o != nil && !isNil(o.ChangeTime) {
		return true
	}

	return false
}

// SetChangeTime gets a reference to the given time.Time and assigns it to the ChangeTime field.
func (o *CoverageInfo) SetChangeTime(v time.Time) {
	o.ChangeTime = &v
}

// GetLocationInfo returns the LocationInfo field value if set, zero value otherwise.
func (o *CoverageInfo) GetLocationInfo() []UserLocation {
	if o == nil || isNil(o.LocationInfo) {
		var ret []UserLocation
		return ret
	}
	return o.LocationInfo
}

// GetLocationInfoOk returns a tuple with the LocationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageInfo) GetLocationInfoOk() ([]UserLocation, bool) {
	if o == nil || isNil(o.LocationInfo) {
    return nil, false
	}
	return o.LocationInfo, true
}

// HasLocationInfo returns a boolean if a field has been set.
func (o *CoverageInfo) HasLocationInfo() bool {
	if o != nil && !isNil(o.LocationInfo) {
		return true
	}

	return false
}

// SetLocationInfo gets a reference to the given []UserLocation and assigns it to the LocationInfo field.
func (o *CoverageInfo) SetLocationInfo(v []UserLocation) {
	o.LocationInfo = v
}

func (o CoverageInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CoverageStatus) {
		toSerialize["coverageStatus"] = o.CoverageStatus
	}
	if !isNil(o.ChangeTime) {
		toSerialize["changeTime"] = o.ChangeTime
	}
	if !isNil(o.LocationInfo) {
		toSerialize["locationInfo"] = o.LocationInfo
	}
	return json.Marshal(toSerialize)
}

type NullableCoverageInfo struct {
	value *CoverageInfo
	isSet bool
}

func (v NullableCoverageInfo) Get() *CoverageInfo {
	return v.value
}

func (v *NullableCoverageInfo) Set(val *CoverageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCoverageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCoverageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoverageInfo(val *CoverageInfo) *NullableCoverageInfo {
	return &NullableCoverageInfo{value: val, isSet: true}
}

func (v NullableCoverageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoverageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


