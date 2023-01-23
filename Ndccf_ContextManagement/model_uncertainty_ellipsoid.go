/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// UncertaintyEllipsoid Ellipsoid with uncertainty
type UncertaintyEllipsoid struct {
	// Indicates value of uncertainty.
	SemiMajor float32 `json:"semiMajor"`
	// Indicates value of uncertainty.
	SemiMinor float32 `json:"semiMinor"`
	// Indicates value of uncertainty.
	Vertical float32 `json:"vertical"`
	// Indicates value of orientation angle.
	OrientationMajor int32 `json:"orientationMajor"`
}

// NewUncertaintyEllipsoid instantiates a new UncertaintyEllipsoid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUncertaintyEllipsoid(semiMajor float32, semiMinor float32, vertical float32, orientationMajor int32) *UncertaintyEllipsoid {
	this := UncertaintyEllipsoid{}
	this.SemiMajor = semiMajor
	this.SemiMinor = semiMinor
	this.Vertical = vertical
	this.OrientationMajor = orientationMajor
	return &this
}

// NewUncertaintyEllipsoidWithDefaults instantiates a new UncertaintyEllipsoid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUncertaintyEllipsoidWithDefaults() *UncertaintyEllipsoid {
	this := UncertaintyEllipsoid{}
	return &this
}

// GetSemiMajor returns the SemiMajor field value
func (o *UncertaintyEllipsoid) GetSemiMajor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SemiMajor
}

// GetSemiMajorOk returns a tuple with the SemiMajor field value
// and a boolean to check if the value has been set.
func (o *UncertaintyEllipsoid) GetSemiMajorOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SemiMajor, true
}

// SetSemiMajor sets field value
func (o *UncertaintyEllipsoid) SetSemiMajor(v float32) {
	o.SemiMajor = v
}

// GetSemiMinor returns the SemiMinor field value
func (o *UncertaintyEllipsoid) GetSemiMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SemiMinor
}

// GetSemiMinorOk returns a tuple with the SemiMinor field value
// and a boolean to check if the value has been set.
func (o *UncertaintyEllipsoid) GetSemiMinorOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SemiMinor, true
}

// SetSemiMinor sets field value
func (o *UncertaintyEllipsoid) SetSemiMinor(v float32) {
	o.SemiMinor = v
}

// GetVertical returns the Vertical field value
func (o *UncertaintyEllipsoid) GetVertical() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Vertical
}

// GetVerticalOk returns a tuple with the Vertical field value
// and a boolean to check if the value has been set.
func (o *UncertaintyEllipsoid) GetVerticalOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Vertical, true
}

// SetVertical sets field value
func (o *UncertaintyEllipsoid) SetVertical(v float32) {
	o.Vertical = v
}

// GetOrientationMajor returns the OrientationMajor field value
func (o *UncertaintyEllipsoid) GetOrientationMajor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrientationMajor
}

// GetOrientationMajorOk returns a tuple with the OrientationMajor field value
// and a boolean to check if the value has been set.
func (o *UncertaintyEllipsoid) GetOrientationMajorOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrientationMajor, true
}

// SetOrientationMajor sets field value
func (o *UncertaintyEllipsoid) SetOrientationMajor(v int32) {
	o.OrientationMajor = v
}

func (o UncertaintyEllipsoid) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["semiMajor"] = o.SemiMajor
	}
	if true {
		toSerialize["semiMinor"] = o.SemiMinor
	}
	if true {
		toSerialize["vertical"] = o.Vertical
	}
	if true {
		toSerialize["orientationMajor"] = o.OrientationMajor
	}
	return json.Marshal(toSerialize)
}

type NullableUncertaintyEllipsoid struct {
	value *UncertaintyEllipsoid
	isSet bool
}

func (v NullableUncertaintyEllipsoid) Get() *UncertaintyEllipsoid {
	return v.value
}

func (v *NullableUncertaintyEllipsoid) Set(val *UncertaintyEllipsoid) {
	v.value = val
	v.isSet = true
}

func (v NullableUncertaintyEllipsoid) IsSet() bool {
	return v.isSet
}

func (v *NullableUncertaintyEllipsoid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUncertaintyEllipsoid(val *UncertaintyEllipsoid) *NullableUncertaintyEllipsoid {
	return &NullableUncertaintyEllipsoid{value: val, isSet: true}
}

func (v NullableUncertaintyEllipsoid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUncertaintyEllipsoid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


