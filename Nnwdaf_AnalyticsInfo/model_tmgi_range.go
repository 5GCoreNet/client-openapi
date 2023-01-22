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

// TmgiRange Range of TMGIs
type TmgiRange struct {
	MbsServiceIdStart string `json:"mbsServiceIdStart"`
	MbsServiceIdEnd string `json:"mbsServiceIdEnd"`
	PlmnId PlmnId `json:"plmnId"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid *string `json:"nid,omitempty"`
}

// NewTmgiRange instantiates a new TmgiRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTmgiRange(mbsServiceIdStart string, mbsServiceIdEnd string, plmnId PlmnId) *TmgiRange {
	this := TmgiRange{}
	this.MbsServiceIdStart = mbsServiceIdStart
	this.MbsServiceIdEnd = mbsServiceIdEnd
	this.PlmnId = plmnId
	return &this
}

// NewTmgiRangeWithDefaults instantiates a new TmgiRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTmgiRangeWithDefaults() *TmgiRange {
	this := TmgiRange{}
	return &this
}

// GetMbsServiceIdStart returns the MbsServiceIdStart field value
func (o *TmgiRange) GetMbsServiceIdStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MbsServiceIdStart
}

// GetMbsServiceIdStartOk returns a tuple with the MbsServiceIdStart field value
// and a boolean to check if the value has been set.
func (o *TmgiRange) GetMbsServiceIdStartOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MbsServiceIdStart, true
}

// SetMbsServiceIdStart sets field value
func (o *TmgiRange) SetMbsServiceIdStart(v string) {
	o.MbsServiceIdStart = v
}

// GetMbsServiceIdEnd returns the MbsServiceIdEnd field value
func (o *TmgiRange) GetMbsServiceIdEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MbsServiceIdEnd
}

// GetMbsServiceIdEndOk returns a tuple with the MbsServiceIdEnd field value
// and a boolean to check if the value has been set.
func (o *TmgiRange) GetMbsServiceIdEndOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MbsServiceIdEnd, true
}

// SetMbsServiceIdEnd sets field value
func (o *TmgiRange) SetMbsServiceIdEnd(v string) {
	o.MbsServiceIdEnd = v
}

// GetPlmnId returns the PlmnId field value
func (o *TmgiRange) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *TmgiRange) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *TmgiRange) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *TmgiRange) GetNid() string {
	if o == nil || isNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmgiRange) GetNidOk() (*string, bool) {
	if o == nil || isNil(o.Nid) {
    return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *TmgiRange) HasNid() bool {
	if o != nil && !isNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *TmgiRange) SetNid(v string) {
	o.Nid = &v
}

func (o TmgiRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mbsServiceIdStart"] = o.MbsServiceIdStart
	}
	if true {
		toSerialize["mbsServiceIdEnd"] = o.MbsServiceIdEnd
	}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !isNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	return json.Marshal(toSerialize)
}

type NullableTmgiRange struct {
	value *TmgiRange
	isSet bool
}

func (v NullableTmgiRange) Get() *TmgiRange {
	return v.value
}

func (v *NullableTmgiRange) Set(val *TmgiRange) {
	v.value = val
	v.isSet = true
}

func (v NullableTmgiRange) IsSet() bool {
	return v.isSet
}

func (v *NullableTmgiRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTmgiRange(val *TmgiRange) *NullableTmgiRange {
	return &NullableTmgiRange{value: val, isSet: true}
}

func (v NullableTmgiRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTmgiRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


