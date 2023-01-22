/*
Nnwdaf_MLModelProvision

Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_MLModelProvision

import (
	"encoding/json"
)

// UncertaintyEllipse Ellipse with uncertainty.
type UncertaintyEllipse struct {
	// Indicates value of uncertainty.
	SemiMajor float32 `json:"semiMajor"`
	// Indicates value of uncertainty.
	SemiMinor float32 `json:"semiMinor"`
	// Indicates value of orientation angle.
	OrientationMajor int32 `json:"orientationMajor"`
}

// NewUncertaintyEllipse instantiates a new UncertaintyEllipse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUncertaintyEllipse(semiMajor float32, semiMinor float32, orientationMajor int32) *UncertaintyEllipse {
	this := UncertaintyEllipse{}
	this.SemiMajor = semiMajor
	this.SemiMinor = semiMinor
	this.OrientationMajor = orientationMajor
	return &this
}

// NewUncertaintyEllipseWithDefaults instantiates a new UncertaintyEllipse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUncertaintyEllipseWithDefaults() *UncertaintyEllipse {
	this := UncertaintyEllipse{}
	return &this
}

// GetSemiMajor returns the SemiMajor field value
func (o *UncertaintyEllipse) GetSemiMajor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SemiMajor
}

// GetSemiMajorOk returns a tuple with the SemiMajor field value
// and a boolean to check if the value has been set.
func (o *UncertaintyEllipse) GetSemiMajorOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SemiMajor, true
}

// SetSemiMajor sets field value
func (o *UncertaintyEllipse) SetSemiMajor(v float32) {
	o.SemiMajor = v
}

// GetSemiMinor returns the SemiMinor field value
func (o *UncertaintyEllipse) GetSemiMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SemiMinor
}

// GetSemiMinorOk returns a tuple with the SemiMinor field value
// and a boolean to check if the value has been set.
func (o *UncertaintyEllipse) GetSemiMinorOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SemiMinor, true
}

// SetSemiMinor sets field value
func (o *UncertaintyEllipse) SetSemiMinor(v float32) {
	o.SemiMinor = v
}

// GetOrientationMajor returns the OrientationMajor field value
func (o *UncertaintyEllipse) GetOrientationMajor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrientationMajor
}

// GetOrientationMajorOk returns a tuple with the OrientationMajor field value
// and a boolean to check if the value has been set.
func (o *UncertaintyEllipse) GetOrientationMajorOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrientationMajor, true
}

// SetOrientationMajor sets field value
func (o *UncertaintyEllipse) SetOrientationMajor(v int32) {
	o.OrientationMajor = v
}

func (o UncertaintyEllipse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["semiMajor"] = o.SemiMajor
	}
	if true {
		toSerialize["semiMinor"] = o.SemiMinor
	}
	if true {
		toSerialize["orientationMajor"] = o.OrientationMajor
	}
	return json.Marshal(toSerialize)
}

type NullableUncertaintyEllipse struct {
	value *UncertaintyEllipse
	isSet bool
}

func (v NullableUncertaintyEllipse) Get() *UncertaintyEllipse {
	return v.value
}

func (v *NullableUncertaintyEllipse) Set(val *UncertaintyEllipse) {
	v.value = val
	v.isSet = true
}

func (v NullableUncertaintyEllipse) IsSet() bool {
	return v.isSet
}

func (v *NullableUncertaintyEllipse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUncertaintyEllipse(val *UncertaintyEllipse) *NullableUncertaintyEllipse {
	return &NullableUncertaintyEllipse{value: val, isSet: true}
}

func (v NullableUncertaintyEllipse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUncertaintyEllipse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


