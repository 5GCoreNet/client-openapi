/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// EcsAddrConfigInfo1 struct for EcsAddrConfigInfo1
type EcsAddrConfigInfo1 struct {
	EcsServerAddr *EcsServerAddr `json:"ecsServerAddr,omitempty"`
	SpatialValidityCond *SpatialValidityCond1 `json:"spatialValidityCond,omitempty"`
}

// NewEcsAddrConfigInfo1 instantiates a new EcsAddrConfigInfo1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcsAddrConfigInfo1() *EcsAddrConfigInfo1 {
	this := EcsAddrConfigInfo1{}
	return &this
}

// NewEcsAddrConfigInfo1WithDefaults instantiates a new EcsAddrConfigInfo1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcsAddrConfigInfo1WithDefaults() *EcsAddrConfigInfo1 {
	this := EcsAddrConfigInfo1{}
	return &this
}

// GetEcsServerAddr returns the EcsServerAddr field value if set, zero value otherwise.
func (o *EcsAddrConfigInfo1) GetEcsServerAddr() EcsServerAddr {
	if o == nil || isNil(o.EcsServerAddr) {
		var ret EcsServerAddr
		return ret
	}
	return *o.EcsServerAddr
}

// GetEcsServerAddrOk returns a tuple with the EcsServerAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcsAddrConfigInfo1) GetEcsServerAddrOk() (*EcsServerAddr, bool) {
	if o == nil || isNil(o.EcsServerAddr) {
    return nil, false
	}
	return o.EcsServerAddr, true
}

// HasEcsServerAddr returns a boolean if a field has been set.
func (o *EcsAddrConfigInfo1) HasEcsServerAddr() bool {
	if o != nil && !isNil(o.EcsServerAddr) {
		return true
	}

	return false
}

// SetEcsServerAddr gets a reference to the given EcsServerAddr and assigns it to the EcsServerAddr field.
func (o *EcsAddrConfigInfo1) SetEcsServerAddr(v EcsServerAddr) {
	o.EcsServerAddr = &v
}

// GetSpatialValidityCond returns the SpatialValidityCond field value if set, zero value otherwise.
func (o *EcsAddrConfigInfo1) GetSpatialValidityCond() SpatialValidityCond1 {
	if o == nil || isNil(o.SpatialValidityCond) {
		var ret SpatialValidityCond1
		return ret
	}
	return *o.SpatialValidityCond
}

// GetSpatialValidityCondOk returns a tuple with the SpatialValidityCond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcsAddrConfigInfo1) GetSpatialValidityCondOk() (*SpatialValidityCond1, bool) {
	if o == nil || isNil(o.SpatialValidityCond) {
    return nil, false
	}
	return o.SpatialValidityCond, true
}

// HasSpatialValidityCond returns a boolean if a field has been set.
func (o *EcsAddrConfigInfo1) HasSpatialValidityCond() bool {
	if o != nil && !isNil(o.SpatialValidityCond) {
		return true
	}

	return false
}

// SetSpatialValidityCond gets a reference to the given SpatialValidityCond1 and assigns it to the SpatialValidityCond field.
func (o *EcsAddrConfigInfo1) SetSpatialValidityCond(v SpatialValidityCond1) {
	o.SpatialValidityCond = &v
}

func (o EcsAddrConfigInfo1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EcsServerAddr) {
		toSerialize["ecsServerAddr"] = o.EcsServerAddr
	}
	if !isNil(o.SpatialValidityCond) {
		toSerialize["spatialValidityCond"] = o.SpatialValidityCond
	}
	return json.Marshal(toSerialize)
}

type NullableEcsAddrConfigInfo1 struct {
	value *EcsAddrConfigInfo1
	isSet bool
}

func (v NullableEcsAddrConfigInfo1) Get() *EcsAddrConfigInfo1 {
	return v.value
}

func (v *NullableEcsAddrConfigInfo1) Set(val *EcsAddrConfigInfo1) {
	v.value = val
	v.isSet = true
}

func (v NullableEcsAddrConfigInfo1) IsSet() bool {
	return v.isSet
}

func (v *NullableEcsAddrConfigInfo1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcsAddrConfigInfo1(val *EcsAddrConfigInfo1) *NullableEcsAddrConfigInfo1 {
	return &NullableEcsAddrConfigInfo1{value: val, isSet: true}
}

func (v NullableEcsAddrConfigInfo1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcsAddrConfigInfo1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


