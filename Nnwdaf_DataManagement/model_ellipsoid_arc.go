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

// EllipsoidArc Ellipsoid Arc.
type EllipsoidArc struct {
	GADShape
	Point GeographicalCoordinates `json:"point"`
	// Indicates value of the inner radius.
	InnerRadius int32 `json:"innerRadius"`
	// Indicates value of uncertainty.
	UncertaintyRadius float32 `json:"uncertaintyRadius"`
	// Indicates value of angle.
	OffsetAngle int32 `json:"offsetAngle"`
	// Indicates value of angle.
	IncludedAngle int32 `json:"includedAngle"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// NewEllipsoidArc instantiates a new EllipsoidArc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEllipsoidArc(point GeographicalCoordinates, innerRadius int32, uncertaintyRadius float32, offsetAngle int32, includedAngle int32, confidence int32, shape SupportedGADShapes) *EllipsoidArc {
	this := EllipsoidArc{}
	this.Shape = shape
	this.Point = point
	this.InnerRadius = innerRadius
	this.UncertaintyRadius = uncertaintyRadius
	this.OffsetAngle = offsetAngle
	this.IncludedAngle = includedAngle
	this.Confidence = confidence
	return &this
}

// NewEllipsoidArcWithDefaults instantiates a new EllipsoidArc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEllipsoidArcWithDefaults() *EllipsoidArc {
	this := EllipsoidArc{}
	return &this
}

// GetPoint returns the Point field value
func (o *EllipsoidArc) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *EllipsoidArc) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *EllipsoidArc) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetInnerRadius returns the InnerRadius field value
func (o *EllipsoidArc) GetInnerRadius() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InnerRadius
}

// GetInnerRadiusOk returns a tuple with the InnerRadius field value
// and a boolean to check if the value has been set.
func (o *EllipsoidArc) GetInnerRadiusOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.InnerRadius, true
}

// SetInnerRadius sets field value
func (o *EllipsoidArc) SetInnerRadius(v int32) {
	o.InnerRadius = v
}

// GetUncertaintyRadius returns the UncertaintyRadius field value
func (o *EllipsoidArc) GetUncertaintyRadius() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UncertaintyRadius
}

// GetUncertaintyRadiusOk returns a tuple with the UncertaintyRadius field value
// and a boolean to check if the value has been set.
func (o *EllipsoidArc) GetUncertaintyRadiusOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UncertaintyRadius, true
}

// SetUncertaintyRadius sets field value
func (o *EllipsoidArc) SetUncertaintyRadius(v float32) {
	o.UncertaintyRadius = v
}

// GetOffsetAngle returns the OffsetAngle field value
func (o *EllipsoidArc) GetOffsetAngle() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OffsetAngle
}

// GetOffsetAngleOk returns a tuple with the OffsetAngle field value
// and a boolean to check if the value has been set.
func (o *EllipsoidArc) GetOffsetAngleOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OffsetAngle, true
}

// SetOffsetAngle sets field value
func (o *EllipsoidArc) SetOffsetAngle(v int32) {
	o.OffsetAngle = v
}

// GetIncludedAngle returns the IncludedAngle field value
func (o *EllipsoidArc) GetIncludedAngle() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IncludedAngle
}

// GetIncludedAngleOk returns a tuple with the IncludedAngle field value
// and a boolean to check if the value has been set.
func (o *EllipsoidArc) GetIncludedAngleOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IncludedAngle, true
}

// SetIncludedAngle sets field value
func (o *EllipsoidArc) SetIncludedAngle(v int32) {
	o.IncludedAngle = v
}

// GetConfidence returns the Confidence field value
func (o *EllipsoidArc) GetConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *EllipsoidArc) GetConfidenceOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *EllipsoidArc) SetConfidence(v int32) {
	o.Confidence = v
}

func (o EllipsoidArc) MarshalJSON() ([]byte, error) {
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
		toSerialize["innerRadius"] = o.InnerRadius
	}
	if true {
		toSerialize["uncertaintyRadius"] = o.UncertaintyRadius
	}
	if true {
		toSerialize["offsetAngle"] = o.OffsetAngle
	}
	if true {
		toSerialize["includedAngle"] = o.IncludedAngle
	}
	if true {
		toSerialize["confidence"] = o.Confidence
	}
	return json.Marshal(toSerialize)
}

type NullableEllipsoidArc struct {
	value *EllipsoidArc
	isSet bool
}

func (v NullableEllipsoidArc) Get() *EllipsoidArc {
	return v.value
}

func (v *NullableEllipsoidArc) Set(val *EllipsoidArc) {
	v.value = val
	v.isSet = true
}

func (v NullableEllipsoidArc) IsSet() bool {
	return v.isSet
}

func (v *NullableEllipsoidArc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEllipsoidArc(val *EllipsoidArc) *NullableEllipsoidArc {
	return &NullableEllipsoidArc{value: val, isSet: true}
}

func (v NullableEllipsoidArc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEllipsoidArc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


