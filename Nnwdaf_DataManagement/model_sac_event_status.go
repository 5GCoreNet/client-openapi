/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_DataManagement

import (
	"encoding/json"
)

// SACEventStatus Contains the network slice status information in terms of the current number of UEs registered  with a network slice, the current number of PDU Sessions established on a network slice or both. 
type SACEventStatus struct {
	ReachedNumUes *SACInfo `json:"reachedNumUes,omitempty"`
	ReachedNumPduSess *SACInfo `json:"reachedNumPduSess,omitempty"`
}

// NewSACEventStatus instantiates a new SACEventStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSACEventStatus() *SACEventStatus {
	this := SACEventStatus{}
	return &this
}

// NewSACEventStatusWithDefaults instantiates a new SACEventStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSACEventStatusWithDefaults() *SACEventStatus {
	this := SACEventStatus{}
	return &this
}

// GetReachedNumUes returns the ReachedNumUes field value if set, zero value otherwise.
func (o *SACEventStatus) GetReachedNumUes() SACInfo {
	if o == nil || isNil(o.ReachedNumUes) {
		var ret SACInfo
		return ret
	}
	return *o.ReachedNumUes
}

// GetReachedNumUesOk returns a tuple with the ReachedNumUes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACEventStatus) GetReachedNumUesOk() (*SACInfo, bool) {
	if o == nil || isNil(o.ReachedNumUes) {
    return nil, false
	}
	return o.ReachedNumUes, true
}

// HasReachedNumUes returns a boolean if a field has been set.
func (o *SACEventStatus) HasReachedNumUes() bool {
	if o != nil && !isNil(o.ReachedNumUes) {
		return true
	}

	return false
}

// SetReachedNumUes gets a reference to the given SACInfo and assigns it to the ReachedNumUes field.
func (o *SACEventStatus) SetReachedNumUes(v SACInfo) {
	o.ReachedNumUes = &v
}

// GetReachedNumPduSess returns the ReachedNumPduSess field value if set, zero value otherwise.
func (o *SACEventStatus) GetReachedNumPduSess() SACInfo {
	if o == nil || isNil(o.ReachedNumPduSess) {
		var ret SACInfo
		return ret
	}
	return *o.ReachedNumPduSess
}

// GetReachedNumPduSessOk returns a tuple with the ReachedNumPduSess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACEventStatus) GetReachedNumPduSessOk() (*SACInfo, bool) {
	if o == nil || isNil(o.ReachedNumPduSess) {
    return nil, false
	}
	return o.ReachedNumPduSess, true
}

// HasReachedNumPduSess returns a boolean if a field has been set.
func (o *SACEventStatus) HasReachedNumPduSess() bool {
	if o != nil && !isNil(o.ReachedNumPduSess) {
		return true
	}

	return false
}

// SetReachedNumPduSess gets a reference to the given SACInfo and assigns it to the ReachedNumPduSess field.
func (o *SACEventStatus) SetReachedNumPduSess(v SACInfo) {
	o.ReachedNumPduSess = &v
}

func (o SACEventStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReachedNumUes) {
		toSerialize["reachedNumUes"] = o.ReachedNumUes
	}
	if !isNil(o.ReachedNumPduSess) {
		toSerialize["reachedNumPduSess"] = o.ReachedNumPduSess
	}
	return json.Marshal(toSerialize)
}

type NullableSACEventStatus struct {
	value *SACEventStatus
	isSet bool
}

func (v NullableSACEventStatus) Get() *SACEventStatus {
	return v.value
}

func (v *NullableSACEventStatus) Set(val *SACEventStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSACEventStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSACEventStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSACEventStatus(val *SACEventStatus) *NullableSACEventStatus {
	return &NullableSACEventStatus{value: val, isSet: true}
}

func (v NullableSACEventStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSACEventStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


