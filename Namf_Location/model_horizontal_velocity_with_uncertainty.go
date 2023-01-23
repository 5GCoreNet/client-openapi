/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Location

import (
	"encoding/json"
)

// HorizontalVelocityWithUncertainty Horizontal velocity with speed uncertainty.
type HorizontalVelocityWithUncertainty struct {
	// Indicates value of horizontal speed.
	HSpeed float32 `json:"hSpeed"`
	// Indicates value of angle.
	Bearing int32 `json:"bearing"`
	// Indicates value of speed uncertainty.
	HUncertainty float32 `json:"hUncertainty"`
}

// NewHorizontalVelocityWithUncertainty instantiates a new HorizontalVelocityWithUncertainty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHorizontalVelocityWithUncertainty(hSpeed float32, bearing int32, hUncertainty float32) *HorizontalVelocityWithUncertainty {
	this := HorizontalVelocityWithUncertainty{}
	this.HSpeed = hSpeed
	this.Bearing = bearing
	this.HUncertainty = hUncertainty
	return &this
}

// NewHorizontalVelocityWithUncertaintyWithDefaults instantiates a new HorizontalVelocityWithUncertainty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHorizontalVelocityWithUncertaintyWithDefaults() *HorizontalVelocityWithUncertainty {
	this := HorizontalVelocityWithUncertainty{}
	return &this
}

// GetHSpeed returns the HSpeed field value
func (o *HorizontalVelocityWithUncertainty) GetHSpeed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HSpeed
}

// GetHSpeedOk returns a tuple with the HSpeed field value
// and a boolean to check if the value has been set.
func (o *HorizontalVelocityWithUncertainty) GetHSpeedOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HSpeed, true
}

// SetHSpeed sets field value
func (o *HorizontalVelocityWithUncertainty) SetHSpeed(v float32) {
	o.HSpeed = v
}

// GetBearing returns the Bearing field value
func (o *HorizontalVelocityWithUncertainty) GetBearing() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bearing
}

// GetBearingOk returns a tuple with the Bearing field value
// and a boolean to check if the value has been set.
func (o *HorizontalVelocityWithUncertainty) GetBearingOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Bearing, true
}

// SetBearing sets field value
func (o *HorizontalVelocityWithUncertainty) SetBearing(v int32) {
	o.Bearing = v
}

// GetHUncertainty returns the HUncertainty field value
func (o *HorizontalVelocityWithUncertainty) GetHUncertainty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HUncertainty
}

// GetHUncertaintyOk returns a tuple with the HUncertainty field value
// and a boolean to check if the value has been set.
func (o *HorizontalVelocityWithUncertainty) GetHUncertaintyOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HUncertainty, true
}

// SetHUncertainty sets field value
func (o *HorizontalVelocityWithUncertainty) SetHUncertainty(v float32) {
	o.HUncertainty = v
}

func (o HorizontalVelocityWithUncertainty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hSpeed"] = o.HSpeed
	}
	if true {
		toSerialize["bearing"] = o.Bearing
	}
	if true {
		toSerialize["hUncertainty"] = o.HUncertainty
	}
	return json.Marshal(toSerialize)
}

type NullableHorizontalVelocityWithUncertainty struct {
	value *HorizontalVelocityWithUncertainty
	isSet bool
}

func (v NullableHorizontalVelocityWithUncertainty) Get() *HorizontalVelocityWithUncertainty {
	return v.value
}

func (v *NullableHorizontalVelocityWithUncertainty) Set(val *HorizontalVelocityWithUncertainty) {
	v.value = val
	v.isSet = true
}

func (v NullableHorizontalVelocityWithUncertainty) IsSet() bool {
	return v.isSet
}

func (v *NullableHorizontalVelocityWithUncertainty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHorizontalVelocityWithUncertainty(val *HorizontalVelocityWithUncertainty) *NullableHorizontalVelocityWithUncertainty {
	return &NullableHorizontalVelocityWithUncertainty{value: val, isSet: true}
}

func (v NullableHorizontalVelocityWithUncertainty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHorizontalVelocityWithUncertainty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


