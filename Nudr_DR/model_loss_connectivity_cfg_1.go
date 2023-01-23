/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// LossConnectivityCfg1 struct for LossConnectivityCfg1
type LossConnectivityCfg1 struct {
	// indicating a time in seconds.
	MaxDetectionTime *int32 `json:"maxDetectionTime,omitempty"`
}

// NewLossConnectivityCfg1 instantiates a new LossConnectivityCfg1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLossConnectivityCfg1() *LossConnectivityCfg1 {
	this := LossConnectivityCfg1{}
	return &this
}

// NewLossConnectivityCfg1WithDefaults instantiates a new LossConnectivityCfg1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLossConnectivityCfg1WithDefaults() *LossConnectivityCfg1 {
	this := LossConnectivityCfg1{}
	return &this
}

// GetMaxDetectionTime returns the MaxDetectionTime field value if set, zero value otherwise.
func (o *LossConnectivityCfg1) GetMaxDetectionTime() int32 {
	if o == nil || isNil(o.MaxDetectionTime) {
		var ret int32
		return ret
	}
	return *o.MaxDetectionTime
}

// GetMaxDetectionTimeOk returns a tuple with the MaxDetectionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LossConnectivityCfg1) GetMaxDetectionTimeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxDetectionTime) {
    return nil, false
	}
	return o.MaxDetectionTime, true
}

// HasMaxDetectionTime returns a boolean if a field has been set.
func (o *LossConnectivityCfg1) HasMaxDetectionTime() bool {
	if o != nil && !isNil(o.MaxDetectionTime) {
		return true
	}

	return false
}

// SetMaxDetectionTime gets a reference to the given int32 and assigns it to the MaxDetectionTime field.
func (o *LossConnectivityCfg1) SetMaxDetectionTime(v int32) {
	o.MaxDetectionTime = &v
}

func (o LossConnectivityCfg1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MaxDetectionTime) {
		toSerialize["maxDetectionTime"] = o.MaxDetectionTime
	}
	return json.Marshal(toSerialize)
}

type NullableLossConnectivityCfg1 struct {
	value *LossConnectivityCfg1
	isSet bool
}

func (v NullableLossConnectivityCfg1) Get() *LossConnectivityCfg1 {
	return v.value
}

func (v *NullableLossConnectivityCfg1) Set(val *LossConnectivityCfg1) {
	v.value = val
	v.isSet = true
}

func (v NullableLossConnectivityCfg1) IsSet() bool {
	return v.isSet
}

func (v *NullableLossConnectivityCfg1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLossConnectivityCfg1(val *LossConnectivityCfg1) *NullableLossConnectivityCfg1 {
	return &NullableLossConnectivityCfg1{value: val, isSet: true}
}

func (v NullableLossConnectivityCfg1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLossConnectivityCfg1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


