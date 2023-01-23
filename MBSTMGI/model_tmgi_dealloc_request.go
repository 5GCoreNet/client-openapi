/*
3gpp-mbs-tmgi

API for the allocation, deallocation and management of TMGI(s) for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSTMGI

import (
	"encoding/json"
)

// TmgiDeallocRequest Represents information to request the deallocation of MBS TMGI(s).
type TmgiDeallocRequest struct {
	AfId string `json:"afId"`
	Tmgis []Tmgi `json:"tmgis"`
}

// NewTmgiDeallocRequest instantiates a new TmgiDeallocRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTmgiDeallocRequest(afId string, tmgis []Tmgi) *TmgiDeallocRequest {
	this := TmgiDeallocRequest{}
	this.AfId = afId
	this.Tmgis = tmgis
	return &this
}

// NewTmgiDeallocRequestWithDefaults instantiates a new TmgiDeallocRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTmgiDeallocRequestWithDefaults() *TmgiDeallocRequest {
	this := TmgiDeallocRequest{}
	return &this
}

// GetAfId returns the AfId field value
func (o *TmgiDeallocRequest) GetAfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value
// and a boolean to check if the value has been set.
func (o *TmgiDeallocRequest) GetAfIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AfId, true
}

// SetAfId sets field value
func (o *TmgiDeallocRequest) SetAfId(v string) {
	o.AfId = v
}

// GetTmgis returns the Tmgis field value
func (o *TmgiDeallocRequest) GetTmgis() []Tmgi {
	if o == nil {
		var ret []Tmgi
		return ret
	}

	return o.Tmgis
}

// GetTmgisOk returns a tuple with the Tmgis field value
// and a boolean to check if the value has been set.
func (o *TmgiDeallocRequest) GetTmgisOk() ([]Tmgi, bool) {
	if o == nil {
    return nil, false
	}
	return o.Tmgis, true
}

// SetTmgis sets field value
func (o *TmgiDeallocRequest) SetTmgis(v []Tmgi) {
	o.Tmgis = v
}

func (o TmgiDeallocRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["afId"] = o.AfId
	}
	if true {
		toSerialize["tmgis"] = o.Tmgis
	}
	return json.Marshal(toSerialize)
}

type NullableTmgiDeallocRequest struct {
	value *TmgiDeallocRequest
	isSet bool
}

func (v NullableTmgiDeallocRequest) Get() *TmgiDeallocRequest {
	return v.value
}

func (v *NullableTmgiDeallocRequest) Set(val *TmgiDeallocRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTmgiDeallocRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTmgiDeallocRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTmgiDeallocRequest(val *TmgiDeallocRequest) *NullableTmgiDeallocRequest {
	return &NullableTmgiDeallocRequest{value: val, isSet: true}
}

func (v NullableTmgiDeallocRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTmgiDeallocRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


