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

// LossConnectivityCfg struct for LossConnectivityCfg
type LossConnectivityCfg struct {
	// indicating a time in seconds.
	MaxDetectionTime *int32 `json:"maxDetectionTime,omitempty"`
}

// NewLossConnectivityCfg instantiates a new LossConnectivityCfg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLossConnectivityCfg() *LossConnectivityCfg {
	this := LossConnectivityCfg{}
	return &this
}

// NewLossConnectivityCfgWithDefaults instantiates a new LossConnectivityCfg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLossConnectivityCfgWithDefaults() *LossConnectivityCfg {
	this := LossConnectivityCfg{}
	return &this
}

// GetMaxDetectionTime returns the MaxDetectionTime field value if set, zero value otherwise.
func (o *LossConnectivityCfg) GetMaxDetectionTime() int32 {
	if o == nil || isNil(o.MaxDetectionTime) {
		var ret int32
		return ret
	}
	return *o.MaxDetectionTime
}

// GetMaxDetectionTimeOk returns a tuple with the MaxDetectionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LossConnectivityCfg) GetMaxDetectionTimeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxDetectionTime) {
    return nil, false
	}
	return o.MaxDetectionTime, true
}

// HasMaxDetectionTime returns a boolean if a field has been set.
func (o *LossConnectivityCfg) HasMaxDetectionTime() bool {
	if o != nil && !isNil(o.MaxDetectionTime) {
		return true
	}

	return false
}

// SetMaxDetectionTime gets a reference to the given int32 and assigns it to the MaxDetectionTime field.
func (o *LossConnectivityCfg) SetMaxDetectionTime(v int32) {
	o.MaxDetectionTime = &v
}

func (o LossConnectivityCfg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MaxDetectionTime) {
		toSerialize["maxDetectionTime"] = o.MaxDetectionTime
	}
	return json.Marshal(toSerialize)
}

type NullableLossConnectivityCfg struct {
	value *LossConnectivityCfg
	isSet bool
}

func (v NullableLossConnectivityCfg) Get() *LossConnectivityCfg {
	return v.value
}

func (v *NullableLossConnectivityCfg) Set(val *LossConnectivityCfg) {
	v.value = val
	v.isSet = true
}

func (v NullableLossConnectivityCfg) IsSet() bool {
	return v.isSet
}

func (v *NullableLossConnectivityCfg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLossConnectivityCfg(val *LossConnectivityCfg) *NullableLossConnectivityCfg {
	return &NullableLossConnectivityCfg{value: val, isSet: true}
}

func (v NullableLossConnectivityCfg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLossConnectivityCfg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


