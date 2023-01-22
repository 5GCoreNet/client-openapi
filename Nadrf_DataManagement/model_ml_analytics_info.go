/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nadrf_DataManagement

import (
	"encoding/json"
)

// MlAnalyticsInfo ML Analytics Filter information supported by the Nnwdaf_MLModelProvision service
type MlAnalyticsInfo struct {
	MlAnalyticsIds []NwdafEvent `json:"mlAnalyticsIds,omitempty"`
	SnssaiList []Snssai `json:"snssaiList,omitempty"`
	TrackingAreaList []Tai `json:"trackingAreaList,omitempty"`
}

// NewMlAnalyticsInfo instantiates a new MlAnalyticsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMlAnalyticsInfo() *MlAnalyticsInfo {
	this := MlAnalyticsInfo{}
	return &this
}

// NewMlAnalyticsInfoWithDefaults instantiates a new MlAnalyticsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMlAnalyticsInfoWithDefaults() *MlAnalyticsInfo {
	this := MlAnalyticsInfo{}
	return &this
}

// GetMlAnalyticsIds returns the MlAnalyticsIds field value if set, zero value otherwise.
func (o *MlAnalyticsInfo) GetMlAnalyticsIds() []NwdafEvent {
	if o == nil || isNil(o.MlAnalyticsIds) {
		var ret []NwdafEvent
		return ret
	}
	return o.MlAnalyticsIds
}

// GetMlAnalyticsIdsOk returns a tuple with the MlAnalyticsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MlAnalyticsInfo) GetMlAnalyticsIdsOk() ([]NwdafEvent, bool) {
	if o == nil || isNil(o.MlAnalyticsIds) {
    return nil, false
	}
	return o.MlAnalyticsIds, true
}

// HasMlAnalyticsIds returns a boolean if a field has been set.
func (o *MlAnalyticsInfo) HasMlAnalyticsIds() bool {
	if o != nil && !isNil(o.MlAnalyticsIds) {
		return true
	}

	return false
}

// SetMlAnalyticsIds gets a reference to the given []NwdafEvent and assigns it to the MlAnalyticsIds field.
func (o *MlAnalyticsInfo) SetMlAnalyticsIds(v []NwdafEvent) {
	o.MlAnalyticsIds = v
}

// GetSnssaiList returns the SnssaiList field value if set, zero value otherwise.
func (o *MlAnalyticsInfo) GetSnssaiList() []Snssai {
	if o == nil || isNil(o.SnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MlAnalyticsInfo) GetSnssaiListOk() ([]Snssai, bool) {
	if o == nil || isNil(o.SnssaiList) {
    return nil, false
	}
	return o.SnssaiList, true
}

// HasSnssaiList returns a boolean if a field has been set.
func (o *MlAnalyticsInfo) HasSnssaiList() bool {
	if o != nil && !isNil(o.SnssaiList) {
		return true
	}

	return false
}

// SetSnssaiList gets a reference to the given []Snssai and assigns it to the SnssaiList field.
func (o *MlAnalyticsInfo) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

// GetTrackingAreaList returns the TrackingAreaList field value if set, zero value otherwise.
func (o *MlAnalyticsInfo) GetTrackingAreaList() []Tai {
	if o == nil || isNil(o.TrackingAreaList) {
		var ret []Tai
		return ret
	}
	return o.TrackingAreaList
}

// GetTrackingAreaListOk returns a tuple with the TrackingAreaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MlAnalyticsInfo) GetTrackingAreaListOk() ([]Tai, bool) {
	if o == nil || isNil(o.TrackingAreaList) {
    return nil, false
	}
	return o.TrackingAreaList, true
}

// HasTrackingAreaList returns a boolean if a field has been set.
func (o *MlAnalyticsInfo) HasTrackingAreaList() bool {
	if o != nil && !isNil(o.TrackingAreaList) {
		return true
	}

	return false
}

// SetTrackingAreaList gets a reference to the given []Tai and assigns it to the TrackingAreaList field.
func (o *MlAnalyticsInfo) SetTrackingAreaList(v []Tai) {
	o.TrackingAreaList = v
}

func (o MlAnalyticsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MlAnalyticsIds) {
		toSerialize["mlAnalyticsIds"] = o.MlAnalyticsIds
	}
	if !isNil(o.SnssaiList) {
		toSerialize["snssaiList"] = o.SnssaiList
	}
	if !isNil(o.TrackingAreaList) {
		toSerialize["trackingAreaList"] = o.TrackingAreaList
	}
	return json.Marshal(toSerialize)
}

type NullableMlAnalyticsInfo struct {
	value *MlAnalyticsInfo
	isSet bool
}

func (v NullableMlAnalyticsInfo) Get() *MlAnalyticsInfo {
	return v.value
}

func (v *NullableMlAnalyticsInfo) Set(val *MlAnalyticsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMlAnalyticsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMlAnalyticsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMlAnalyticsInfo(val *MlAnalyticsInfo) *NullableMlAnalyticsInfo {
	return &NullableMlAnalyticsInfo{value: val, isSet: true}
}

func (v NullableMlAnalyticsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMlAnalyticsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


