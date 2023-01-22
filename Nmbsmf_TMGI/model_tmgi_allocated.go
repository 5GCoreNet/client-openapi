/*
Nmbsmf_TMGI

MB-SMF TMGI Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsmf_TMGI

import (
	"encoding/json"
	"time"
)

// TmgiAllocated Data within TMGI Allocate Response
type TmgiAllocated struct {
	// One or more TMGI values
	TmgiList []Tmgi `json:"tmgiList"`
	// string with format 'date-time' as defined in OpenAPI.
	ExpirationTime time.Time `json:"expirationTime"`
}

// NewTmgiAllocated instantiates a new TmgiAllocated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTmgiAllocated(tmgiList []Tmgi, expirationTime time.Time) *TmgiAllocated {
	this := TmgiAllocated{}
	this.TmgiList = tmgiList
	this.ExpirationTime = expirationTime
	return &this
}

// NewTmgiAllocatedWithDefaults instantiates a new TmgiAllocated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTmgiAllocatedWithDefaults() *TmgiAllocated {
	this := TmgiAllocated{}
	return &this
}

// GetTmgiList returns the TmgiList field value
func (o *TmgiAllocated) GetTmgiList() []Tmgi {
	if o == nil {
		var ret []Tmgi
		return ret
	}

	return o.TmgiList
}

// GetTmgiListOk returns a tuple with the TmgiList field value
// and a boolean to check if the value has been set.
func (o *TmgiAllocated) GetTmgiListOk() ([]Tmgi, bool) {
	if o == nil {
    return nil, false
	}
	return o.TmgiList, true
}

// SetTmgiList sets field value
func (o *TmgiAllocated) SetTmgiList(v []Tmgi) {
	o.TmgiList = v
}

// GetExpirationTime returns the ExpirationTime field value
func (o *TmgiAllocated) GetExpirationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value
// and a boolean to check if the value has been set.
func (o *TmgiAllocated) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExpirationTime, true
}

// SetExpirationTime sets field value
func (o *TmgiAllocated) SetExpirationTime(v time.Time) {
	o.ExpirationTime = v
}

func (o TmgiAllocated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tmgiList"] = o.TmgiList
	}
	if true {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	return json.Marshal(toSerialize)
}

type NullableTmgiAllocated struct {
	value *TmgiAllocated
	isSet bool
}

func (v NullableTmgiAllocated) Get() *TmgiAllocated {
	return v.value
}

func (v *NullableTmgiAllocated) Set(val *TmgiAllocated) {
	v.value = val
	v.isSet = true
}

func (v NullableTmgiAllocated) IsSet() bool {
	return v.isSet
}

func (v *NullableTmgiAllocated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTmgiAllocated(val *TmgiAllocated) *NullableTmgiAllocated {
	return &NullableTmgiAllocated{value: val, isSet: true}
}

func (v NullableTmgiAllocated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTmgiAllocated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


