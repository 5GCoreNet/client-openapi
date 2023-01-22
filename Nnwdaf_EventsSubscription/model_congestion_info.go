/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_EventsSubscription

import (
	"encoding/json"
)

// CongestionInfo Represents the congestion information.
type CongestionInfo struct {
	CongType CongestionType `json:"congType"`
	TimeIntev TimeWindow `json:"timeIntev"`
	Nsi ThresholdLevel `json:"nsi"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
	TopAppListUl []TopApplication `json:"topAppListUl,omitempty"`
	TopAppListDl []TopApplication `json:"topAppListDl,omitempty"`
}

// NewCongestionInfo instantiates a new CongestionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCongestionInfo(congType CongestionType, timeIntev TimeWindow, nsi ThresholdLevel) *CongestionInfo {
	this := CongestionInfo{}
	this.CongType = congType
	this.TimeIntev = timeIntev
	this.Nsi = nsi
	return &this
}

// NewCongestionInfoWithDefaults instantiates a new CongestionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCongestionInfoWithDefaults() *CongestionInfo {
	this := CongestionInfo{}
	return &this
}

// GetCongType returns the CongType field value
func (o *CongestionInfo) GetCongType() CongestionType {
	if o == nil {
		var ret CongestionType
		return ret
	}

	return o.CongType
}

// GetCongTypeOk returns a tuple with the CongType field value
// and a boolean to check if the value has been set.
func (o *CongestionInfo) GetCongTypeOk() (*CongestionType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CongType, true
}

// SetCongType sets field value
func (o *CongestionInfo) SetCongType(v CongestionType) {
	o.CongType = v
}

// GetTimeIntev returns the TimeIntev field value
func (o *CongestionInfo) GetTimeIntev() TimeWindow {
	if o == nil {
		var ret TimeWindow
		return ret
	}

	return o.TimeIntev
}

// GetTimeIntevOk returns a tuple with the TimeIntev field value
// and a boolean to check if the value has been set.
func (o *CongestionInfo) GetTimeIntevOk() (*TimeWindow, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeIntev, true
}

// SetTimeIntev sets field value
func (o *CongestionInfo) SetTimeIntev(v TimeWindow) {
	o.TimeIntev = v
}

// GetNsi returns the Nsi field value
func (o *CongestionInfo) GetNsi() ThresholdLevel {
	if o == nil {
		var ret ThresholdLevel
		return ret
	}

	return o.Nsi
}

// GetNsiOk returns a tuple with the Nsi field value
// and a boolean to check if the value has been set.
func (o *CongestionInfo) GetNsiOk() (*ThresholdLevel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Nsi, true
}

// SetNsi sets field value
func (o *CongestionInfo) SetNsi(v ThresholdLevel) {
	o.Nsi = v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *CongestionInfo) GetConfidence() int32 {
	if o == nil || isNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CongestionInfo) GetConfidenceOk() (*int32, bool) {
	if o == nil || isNil(o.Confidence) {
    return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *CongestionInfo) HasConfidence() bool {
	if o != nil && !isNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *CongestionInfo) SetConfidence(v int32) {
	o.Confidence = &v
}

// GetTopAppListUl returns the TopAppListUl field value if set, zero value otherwise.
func (o *CongestionInfo) GetTopAppListUl() []TopApplication {
	if o == nil || isNil(o.TopAppListUl) {
		var ret []TopApplication
		return ret
	}
	return o.TopAppListUl
}

// GetTopAppListUlOk returns a tuple with the TopAppListUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CongestionInfo) GetTopAppListUlOk() ([]TopApplication, bool) {
	if o == nil || isNil(o.TopAppListUl) {
    return nil, false
	}
	return o.TopAppListUl, true
}

// HasTopAppListUl returns a boolean if a field has been set.
func (o *CongestionInfo) HasTopAppListUl() bool {
	if o != nil && !isNil(o.TopAppListUl) {
		return true
	}

	return false
}

// SetTopAppListUl gets a reference to the given []TopApplication and assigns it to the TopAppListUl field.
func (o *CongestionInfo) SetTopAppListUl(v []TopApplication) {
	o.TopAppListUl = v
}

// GetTopAppListDl returns the TopAppListDl field value if set, zero value otherwise.
func (o *CongestionInfo) GetTopAppListDl() []TopApplication {
	if o == nil || isNil(o.TopAppListDl) {
		var ret []TopApplication
		return ret
	}
	return o.TopAppListDl
}

// GetTopAppListDlOk returns a tuple with the TopAppListDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CongestionInfo) GetTopAppListDlOk() ([]TopApplication, bool) {
	if o == nil || isNil(o.TopAppListDl) {
    return nil, false
	}
	return o.TopAppListDl, true
}

// HasTopAppListDl returns a boolean if a field has been set.
func (o *CongestionInfo) HasTopAppListDl() bool {
	if o != nil && !isNil(o.TopAppListDl) {
		return true
	}

	return false
}

// SetTopAppListDl gets a reference to the given []TopApplication and assigns it to the TopAppListDl field.
func (o *CongestionInfo) SetTopAppListDl(v []TopApplication) {
	o.TopAppListDl = v
}

func (o CongestionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["congType"] = o.CongType
	}
	if true {
		toSerialize["timeIntev"] = o.TimeIntev
	}
	if true {
		toSerialize["nsi"] = o.Nsi
	}
	if !isNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	if !isNil(o.TopAppListUl) {
		toSerialize["topAppListUl"] = o.TopAppListUl
	}
	if !isNil(o.TopAppListDl) {
		toSerialize["topAppListDl"] = o.TopAppListDl
	}
	return json.Marshal(toSerialize)
}

type NullableCongestionInfo struct {
	value *CongestionInfo
	isSet bool
}

func (v NullableCongestionInfo) Get() *CongestionInfo {
	return v.value
}

func (v *NullableCongestionInfo) Set(val *CongestionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCongestionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCongestionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCongestionInfo(val *CongestionInfo) *NullableCongestionInfo {
	return &NullableCongestionInfo{value: val, isSet: true}
}

func (v NullableCongestionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCongestionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


