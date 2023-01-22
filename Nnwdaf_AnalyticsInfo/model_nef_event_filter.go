/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// NefEventFilter Represents event filter information for an event.
type NefEventFilter struct {
	TgtUe TargetUeIdentification `json:"tgtUe"`
	AppIds []string `json:"appIds,omitempty"`
	LocArea *NetworkAreaInfo `json:"locArea,omitempty"`
	CollAttrs []CollectiveBehaviourFilter `json:"collAttrs,omitempty"`
}

// NewNefEventFilter instantiates a new NefEventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNefEventFilter(tgtUe TargetUeIdentification) *NefEventFilter {
	this := NefEventFilter{}
	this.TgtUe = tgtUe
	return &this
}

// NewNefEventFilterWithDefaults instantiates a new NefEventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNefEventFilterWithDefaults() *NefEventFilter {
	this := NefEventFilter{}
	return &this
}

// GetTgtUe returns the TgtUe field value
func (o *NefEventFilter) GetTgtUe() TargetUeIdentification {
	if o == nil {
		var ret TargetUeIdentification
		return ret
	}

	return o.TgtUe
}

// GetTgtUeOk returns a tuple with the TgtUe field value
// and a boolean to check if the value has been set.
func (o *NefEventFilter) GetTgtUeOk() (*TargetUeIdentification, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TgtUe, true
}

// SetTgtUe sets field value
func (o *NefEventFilter) SetTgtUe(v TargetUeIdentification) {
	o.TgtUe = v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *NefEventFilter) GetAppIds() []string {
	if o == nil || isNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventFilter) GetAppIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AppIds) {
    return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *NefEventFilter) HasAppIds() bool {
	if o != nil && !isNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *NefEventFilter) SetAppIds(v []string) {
	o.AppIds = v
}

// GetLocArea returns the LocArea field value if set, zero value otherwise.
func (o *NefEventFilter) GetLocArea() NetworkAreaInfo {
	if o == nil || isNil(o.LocArea) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.LocArea
}

// GetLocAreaOk returns a tuple with the LocArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventFilter) GetLocAreaOk() (*NetworkAreaInfo, bool) {
	if o == nil || isNil(o.LocArea) {
    return nil, false
	}
	return o.LocArea, true
}

// HasLocArea returns a boolean if a field has been set.
func (o *NefEventFilter) HasLocArea() bool {
	if o != nil && !isNil(o.LocArea) {
		return true
	}

	return false
}

// SetLocArea gets a reference to the given NetworkAreaInfo and assigns it to the LocArea field.
func (o *NefEventFilter) SetLocArea(v NetworkAreaInfo) {
	o.LocArea = &v
}

// GetCollAttrs returns the CollAttrs field value if set, zero value otherwise.
func (o *NefEventFilter) GetCollAttrs() []CollectiveBehaviourFilter {
	if o == nil || isNil(o.CollAttrs) {
		var ret []CollectiveBehaviourFilter
		return ret
	}
	return o.CollAttrs
}

// GetCollAttrsOk returns a tuple with the CollAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NefEventFilter) GetCollAttrsOk() ([]CollectiveBehaviourFilter, bool) {
	if o == nil || isNil(o.CollAttrs) {
    return nil, false
	}
	return o.CollAttrs, true
}

// HasCollAttrs returns a boolean if a field has been set.
func (o *NefEventFilter) HasCollAttrs() bool {
	if o != nil && !isNil(o.CollAttrs) {
		return true
	}

	return false
}

// SetCollAttrs gets a reference to the given []CollectiveBehaviourFilter and assigns it to the CollAttrs field.
func (o *NefEventFilter) SetCollAttrs(v []CollectiveBehaviourFilter) {
	o.CollAttrs = v
}

func (o NefEventFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tgtUe"] = o.TgtUe
	}
	if !isNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	if !isNil(o.LocArea) {
		toSerialize["locArea"] = o.LocArea
	}
	if !isNil(o.CollAttrs) {
		toSerialize["collAttrs"] = o.CollAttrs
	}
	return json.Marshal(toSerialize)
}

type NullableNefEventFilter struct {
	value *NefEventFilter
	isSet bool
}

func (v NullableNefEventFilter) Get() *NefEventFilter {
	return v.value
}

func (v *NullableNefEventFilter) Set(val *NefEventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableNefEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableNefEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNefEventFilter(val *NefEventFilter) *NullableNefEventFilter {
	return &NullableNefEventFilter{value: val, isSet: true}
}

func (v NullableNefEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNefEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


