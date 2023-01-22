/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nlmf_Location

import (
	"encoding/json"
)

// PointAltitudeUncertainty Ellipsoid point with altitude and uncertainty ellipsoid.
type PointAltitudeUncertainty struct {
	GADShape
	Point GeographicalCoordinates `json:"point"`
	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`
	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`
	// Indicates value of uncertainty.
	UncertaintyAltitude float32 `json:"uncertaintyAltitude"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// NewPointAltitudeUncertainty instantiates a new PointAltitudeUncertainty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointAltitudeUncertainty(point GeographicalCoordinates, altitude float64, uncertaintyEllipse UncertaintyEllipse, uncertaintyAltitude float32, confidence int32, shape SupportedGADShapes) *PointAltitudeUncertainty {
	this := PointAltitudeUncertainty{}
	this.Shape = shape
	this.Point = point
	this.Altitude = altitude
	this.UncertaintyEllipse = uncertaintyEllipse
	this.UncertaintyAltitude = uncertaintyAltitude
	this.Confidence = confidence
	return &this
}

// NewPointAltitudeUncertaintyWithDefaults instantiates a new PointAltitudeUncertainty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointAltitudeUncertaintyWithDefaults() *PointAltitudeUncertainty {
	this := PointAltitudeUncertainty{}
	return &this
}

// GetPoint returns the Point field value
func (o *PointAltitudeUncertainty) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *PointAltitudeUncertainty) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *PointAltitudeUncertainty) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetAltitude returns the Altitude field value
func (o *PointAltitudeUncertainty) GetAltitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field value
// and a boolean to check if the value has been set.
func (o *PointAltitudeUncertainty) GetAltitudeOk() (*float64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Altitude, true
}

// SetAltitude sets field value
func (o *PointAltitudeUncertainty) SetAltitude(v float64) {
	o.Altitude = v
}

// GetUncertaintyEllipse returns the UncertaintyEllipse field value
func (o *PointAltitudeUncertainty) GetUncertaintyEllipse() UncertaintyEllipse {
	if o == nil {
		var ret UncertaintyEllipse
		return ret
	}

	return o.UncertaintyEllipse
}

// GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field value
// and a boolean to check if the value has been set.
func (o *PointAltitudeUncertainty) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UncertaintyEllipse, true
}

// SetUncertaintyEllipse sets field value
func (o *PointAltitudeUncertainty) SetUncertaintyEllipse(v UncertaintyEllipse) {
	o.UncertaintyEllipse = v
}

// GetUncertaintyAltitude returns the UncertaintyAltitude field value
func (o *PointAltitudeUncertainty) GetUncertaintyAltitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UncertaintyAltitude
}

// GetUncertaintyAltitudeOk returns a tuple with the UncertaintyAltitude field value
// and a boolean to check if the value has been set.
func (o *PointAltitudeUncertainty) GetUncertaintyAltitudeOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UncertaintyAltitude, true
}

// SetUncertaintyAltitude sets field value
func (o *PointAltitudeUncertainty) SetUncertaintyAltitude(v float32) {
	o.UncertaintyAltitude = v
}

// GetConfidence returns the Confidence field value
func (o *PointAltitudeUncertainty) GetConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *PointAltitudeUncertainty) GetConfidenceOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *PointAltitudeUncertainty) SetConfidence(v int32) {
	o.Confidence = v
}

func (o PointAltitudeUncertainty) MarshalJSON() ([]byte, error) {
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
		toSerialize["altitude"] = o.Altitude
	}
	if true {
		toSerialize["uncertaintyEllipse"] = o.UncertaintyEllipse
	}
	if true {
		toSerialize["uncertaintyAltitude"] = o.UncertaintyAltitude
	}
	if true {
		toSerialize["confidence"] = o.Confidence
	}
	return json.Marshal(toSerialize)
}

type NullablePointAltitudeUncertainty struct {
	value *PointAltitudeUncertainty
	isSet bool
}

func (v NullablePointAltitudeUncertainty) Get() *PointAltitudeUncertainty {
	return v.value
}

func (v *NullablePointAltitudeUncertainty) Set(val *PointAltitudeUncertainty) {
	v.value = val
	v.isSet = true
}

func (v NullablePointAltitudeUncertainty) IsSet() bool {
	return v.isSet
}

func (v *NullablePointAltitudeUncertainty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointAltitudeUncertainty(val *PointAltitudeUncertainty) *NullablePointAltitudeUncertainty {
	return &NullablePointAltitudeUncertainty{value: val, isSet: true}
}

func (v NullablePointAltitudeUncertainty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointAltitudeUncertainty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


