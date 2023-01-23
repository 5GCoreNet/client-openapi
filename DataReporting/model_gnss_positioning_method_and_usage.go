/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_DataReporting

import (
	"encoding/json"
)

// GnssPositioningMethodAndUsage Indicates the usage of a Global Navigation Satellite System (GNSS) positioning method.
type GnssPositioningMethodAndUsage struct {
	Mode PositioningMode `json:"mode"`
	Gnss GnssId `json:"gnss"`
	Usage Usage `json:"usage"`
}

// NewGnssPositioningMethodAndUsage instantiates a new GnssPositioningMethodAndUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGnssPositioningMethodAndUsage(mode PositioningMode, gnss GnssId, usage Usage) *GnssPositioningMethodAndUsage {
	this := GnssPositioningMethodAndUsage{}
	this.Mode = mode
	this.Gnss = gnss
	this.Usage = usage
	return &this
}

// NewGnssPositioningMethodAndUsageWithDefaults instantiates a new GnssPositioningMethodAndUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGnssPositioningMethodAndUsageWithDefaults() *GnssPositioningMethodAndUsage {
	this := GnssPositioningMethodAndUsage{}
	return &this
}

// GetMode returns the Mode field value
func (o *GnssPositioningMethodAndUsage) GetMode() PositioningMode {
	if o == nil {
		var ret PositioningMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *GnssPositioningMethodAndUsage) GetModeOk() (*PositioningMode, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *GnssPositioningMethodAndUsage) SetMode(v PositioningMode) {
	o.Mode = v
}

// GetGnss returns the Gnss field value
func (o *GnssPositioningMethodAndUsage) GetGnss() GnssId {
	if o == nil {
		var ret GnssId
		return ret
	}

	return o.Gnss
}

// GetGnssOk returns a tuple with the Gnss field value
// and a boolean to check if the value has been set.
func (o *GnssPositioningMethodAndUsage) GetGnssOk() (*GnssId, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Gnss, true
}

// SetGnss sets field value
func (o *GnssPositioningMethodAndUsage) SetGnss(v GnssId) {
	o.Gnss = v
}

// GetUsage returns the Usage field value
func (o *GnssPositioningMethodAndUsage) GetUsage() Usage {
	if o == nil {
		var ret Usage
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *GnssPositioningMethodAndUsage) GetUsageOk() (*Usage, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *GnssPositioningMethodAndUsage) SetUsage(v Usage) {
	o.Usage = v
}

func (o GnssPositioningMethodAndUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["gnss"] = o.Gnss
	}
	if true {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableGnssPositioningMethodAndUsage struct {
	value *GnssPositioningMethodAndUsage
	isSet bool
}

func (v NullableGnssPositioningMethodAndUsage) Get() *GnssPositioningMethodAndUsage {
	return v.value
}

func (v *NullableGnssPositioningMethodAndUsage) Set(val *GnssPositioningMethodAndUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableGnssPositioningMethodAndUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableGnssPositioningMethodAndUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGnssPositioningMethodAndUsage(val *GnssPositioningMethodAndUsage) *NullableGnssPositioningMethodAndUsage {
	return &NullableGnssPositioningMethodAndUsage{value: val, isSet: true}
}

func (v NullableGnssPositioningMethodAndUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGnssPositioningMethodAndUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


