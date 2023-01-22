/*
3gpp-traffic-influence

API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package TrafficInfluence

import (
	"encoding/json"
)

// PointUncertaintyCircle Ellipsoid point with uncertainty circle.
type PointUncertaintyCircle struct {
	GADShape
	Point GeographicalCoordinates `json:"point"`
	// Indicates value of uncertainty.
	Uncertainty float32 `json:"uncertainty"`
}

// NewPointUncertaintyCircle instantiates a new PointUncertaintyCircle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointUncertaintyCircle(point GeographicalCoordinates, uncertainty float32, shape SupportedGADShapes) *PointUncertaintyCircle {
	this := PointUncertaintyCircle{}
	this.Shape = shape
	this.Point = point
	this.Uncertainty = uncertainty
	return &this
}

// NewPointUncertaintyCircleWithDefaults instantiates a new PointUncertaintyCircle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointUncertaintyCircleWithDefaults() *PointUncertaintyCircle {
	this := PointUncertaintyCircle{}
	return &this
}

// GetPoint returns the Point field value
func (o *PointUncertaintyCircle) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyCircle) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *PointUncertaintyCircle) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetUncertainty returns the Uncertainty field value
func (o *PointUncertaintyCircle) GetUncertainty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Uncertainty
}

// GetUncertaintyOk returns a tuple with the Uncertainty field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyCircle) GetUncertaintyOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Uncertainty, true
}

// SetUncertainty sets field value
func (o *PointUncertaintyCircle) SetUncertainty(v float32) {
	o.Uncertainty = v
}

func (o PointUncertaintyCircle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedGADShape, errGADShape := json.Marshal(o.GADShape)
	if errGADShape != nil {
		return []byte{}, errGADShape
	}
	errGADShape = json.Unmarshal([]byte(serializedGADShape), &toSerialize)
	if errGADShape != nil {
		return []byte{}, errGADShape
	}
	if true {
		toSerialize["point"] = o.Point
	}
	if true {
		toSerialize["uncertainty"] = o.Uncertainty
	}
	return json.Marshal(toSerialize)
}

type NullablePointUncertaintyCircle struct {
	value *PointUncertaintyCircle
	isSet bool
}

func (v NullablePointUncertaintyCircle) Get() *PointUncertaintyCircle {
	return v.value
}

func (v *NullablePointUncertaintyCircle) Set(val *PointUncertaintyCircle) {
	v.value = val
	v.isSet = true
}

func (v NullablePointUncertaintyCircle) IsSet() bool {
	return v.isSet
}

func (v *NullablePointUncertaintyCircle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointUncertaintyCircle(val *PointUncertaintyCircle) *NullablePointUncertaintyCircle {
	return &NullablePointUncertaintyCircle{value: val, isSet: true}
}

func (v NullablePointUncertaintyCircle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointUncertaintyCircle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


