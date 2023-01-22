/*
Nnsacf_SliceEventExposure

Nnsacf_SliceEventExposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnsacf_SliceEventExposure

import (
	"encoding/json"
)

// SACInfo Represents threshold(s) to control the triggering of network slice reporting notifications or the information contained in the network slice reporting notification. 
type SACInfo struct {
	NumericValNumUes *int32 `json:"numericValNumUes,omitempty"`
	NumericValNumPduSess *int32 `json:"numericValNumPduSess,omitempty"`
	PercValueNumUes *int32 `json:"percValueNumUes,omitempty"`
	PercValueNumPduSess *int32 `json:"percValueNumPduSess,omitempty"`
}

// NewSACInfo instantiates a new SACInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSACInfo() *SACInfo {
	this := SACInfo{}
	return &this
}

// NewSACInfoWithDefaults instantiates a new SACInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSACInfoWithDefaults() *SACInfo {
	this := SACInfo{}
	return &this
}

// GetNumericValNumUes returns the NumericValNumUes field value if set, zero value otherwise.
func (o *SACInfo) GetNumericValNumUes() int32 {
	if o == nil || isNil(o.NumericValNumUes) {
		var ret int32
		return ret
	}
	return *o.NumericValNumUes
}

// GetNumericValNumUesOk returns a tuple with the NumericValNumUes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACInfo) GetNumericValNumUesOk() (*int32, bool) {
	if o == nil || isNil(o.NumericValNumUes) {
    return nil, false
	}
	return o.NumericValNumUes, true
}

// HasNumericValNumUes returns a boolean if a field has been set.
func (o *SACInfo) HasNumericValNumUes() bool {
	if o != nil && !isNil(o.NumericValNumUes) {
		return true
	}

	return false
}

// SetNumericValNumUes gets a reference to the given int32 and assigns it to the NumericValNumUes field.
func (o *SACInfo) SetNumericValNumUes(v int32) {
	o.NumericValNumUes = &v
}

// GetNumericValNumPduSess returns the NumericValNumPduSess field value if set, zero value otherwise.
func (o *SACInfo) GetNumericValNumPduSess() int32 {
	if o == nil || isNil(o.NumericValNumPduSess) {
		var ret int32
		return ret
	}
	return *o.NumericValNumPduSess
}

// GetNumericValNumPduSessOk returns a tuple with the NumericValNumPduSess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACInfo) GetNumericValNumPduSessOk() (*int32, bool) {
	if o == nil || isNil(o.NumericValNumPduSess) {
    return nil, false
	}
	return o.NumericValNumPduSess, true
}

// HasNumericValNumPduSess returns a boolean if a field has been set.
func (o *SACInfo) HasNumericValNumPduSess() bool {
	if o != nil && !isNil(o.NumericValNumPduSess) {
		return true
	}

	return false
}

// SetNumericValNumPduSess gets a reference to the given int32 and assigns it to the NumericValNumPduSess field.
func (o *SACInfo) SetNumericValNumPduSess(v int32) {
	o.NumericValNumPduSess = &v
}

// GetPercValueNumUes returns the PercValueNumUes field value if set, zero value otherwise.
func (o *SACInfo) GetPercValueNumUes() int32 {
	if o == nil || isNil(o.PercValueNumUes) {
		var ret int32
		return ret
	}
	return *o.PercValueNumUes
}

// GetPercValueNumUesOk returns a tuple with the PercValueNumUes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACInfo) GetPercValueNumUesOk() (*int32, bool) {
	if o == nil || isNil(o.PercValueNumUes) {
    return nil, false
	}
	return o.PercValueNumUes, true
}

// HasPercValueNumUes returns a boolean if a field has been set.
func (o *SACInfo) HasPercValueNumUes() bool {
	if o != nil && !isNil(o.PercValueNumUes) {
		return true
	}

	return false
}

// SetPercValueNumUes gets a reference to the given int32 and assigns it to the PercValueNumUes field.
func (o *SACInfo) SetPercValueNumUes(v int32) {
	o.PercValueNumUes = &v
}

// GetPercValueNumPduSess returns the PercValueNumPduSess field value if set, zero value otherwise.
func (o *SACInfo) GetPercValueNumPduSess() int32 {
	if o == nil || isNil(o.PercValueNumPduSess) {
		var ret int32
		return ret
	}
	return *o.PercValueNumPduSess
}

// GetPercValueNumPduSessOk returns a tuple with the PercValueNumPduSess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACInfo) GetPercValueNumPduSessOk() (*int32, bool) {
	if o == nil || isNil(o.PercValueNumPduSess) {
    return nil, false
	}
	return o.PercValueNumPduSess, true
}

// HasPercValueNumPduSess returns a boolean if a field has been set.
func (o *SACInfo) HasPercValueNumPduSess() bool {
	if o != nil && !isNil(o.PercValueNumPduSess) {
		return true
	}

	return false
}

// SetPercValueNumPduSess gets a reference to the given int32 and assigns it to the PercValueNumPduSess field.
func (o *SACInfo) SetPercValueNumPduSess(v int32) {
	o.PercValueNumPduSess = &v
}

func (o SACInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NumericValNumUes) {
		toSerialize["numericValNumUes"] = o.NumericValNumUes
	}
	if !isNil(o.NumericValNumPduSess) {
		toSerialize["numericValNumPduSess"] = o.NumericValNumPduSess
	}
	if !isNil(o.PercValueNumUes) {
		toSerialize["percValueNumUes"] = o.PercValueNumUes
	}
	if !isNil(o.PercValueNumPduSess) {
		toSerialize["percValueNumPduSess"] = o.PercValueNumPduSess
	}
	return json.Marshal(toSerialize)
}

type NullableSACInfo struct {
	value *SACInfo
	isSet bool
}

func (v NullableSACInfo) Get() *SACInfo {
	return v.value
}

func (v *NullableSACInfo) Set(val *SACInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSACInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSACInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSACInfo(val *SACInfo) *NullableSACInfo {
	return &NullableSACInfo{value: val, isSet: true}
}

func (v NullableSACInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSACInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


