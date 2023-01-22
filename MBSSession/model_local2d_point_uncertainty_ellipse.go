/*
3gpp-mbs-session

API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MBSSession

import (
	"encoding/json"
)

// Local2dPointUncertaintyEllipse Local 2D point with uncertainty ellipse
type Local2dPointUncertaintyEllipse struct {
	GADShape
	LocalOrigin LocalOrigin `json:"localOrigin"`
	Point RelativeCartesianLocation `json:"point"`
	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// NewLocal2dPointUncertaintyEllipse instantiates a new Local2dPointUncertaintyEllipse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocal2dPointUncertaintyEllipse(localOrigin LocalOrigin, point RelativeCartesianLocation, uncertaintyEllipse UncertaintyEllipse, confidence int32, shape SupportedGADShapes) *Local2dPointUncertaintyEllipse {
	this := Local2dPointUncertaintyEllipse{}
	this.Shape = shape
	this.LocalOrigin = localOrigin
	this.Point = point
	this.UncertaintyEllipse = uncertaintyEllipse
	this.Confidence = confidence
	return &this
}

// NewLocal2dPointUncertaintyEllipseWithDefaults instantiates a new Local2dPointUncertaintyEllipse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocal2dPointUncertaintyEllipseWithDefaults() *Local2dPointUncertaintyEllipse {
	this := Local2dPointUncertaintyEllipse{}
	return &this
}

// GetLocalOrigin returns the LocalOrigin field value
func (o *Local2dPointUncertaintyEllipse) GetLocalOrigin() LocalOrigin {
	if o == nil {
		var ret LocalOrigin
		return ret
	}

	return o.LocalOrigin
}

// GetLocalOriginOk returns a tuple with the LocalOrigin field value
// and a boolean to check if the value has been set.
func (o *Local2dPointUncertaintyEllipse) GetLocalOriginOk() (*LocalOrigin, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocalOrigin, true
}

// SetLocalOrigin sets field value
func (o *Local2dPointUncertaintyEllipse) SetLocalOrigin(v LocalOrigin) {
	o.LocalOrigin = v
}

// GetPoint returns the Point field value
func (o *Local2dPointUncertaintyEllipse) GetPoint() RelativeCartesianLocation {
	if o == nil {
		var ret RelativeCartesianLocation
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *Local2dPointUncertaintyEllipse) GetPointOk() (*RelativeCartesianLocation, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *Local2dPointUncertaintyEllipse) SetPoint(v RelativeCartesianLocation) {
	o.Point = v
}

// GetUncertaintyEllipse returns the UncertaintyEllipse field value
func (o *Local2dPointUncertaintyEllipse) GetUncertaintyEllipse() UncertaintyEllipse {
	if o == nil {
		var ret UncertaintyEllipse
		return ret
	}

	return o.UncertaintyEllipse
}

// GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field value
// and a boolean to check if the value has been set.
func (o *Local2dPointUncertaintyEllipse) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UncertaintyEllipse, true
}

// SetUncertaintyEllipse sets field value
func (o *Local2dPointUncertaintyEllipse) SetUncertaintyEllipse(v UncertaintyEllipse) {
	o.UncertaintyEllipse = v
}

// GetConfidence returns the Confidence field value
func (o *Local2dPointUncertaintyEllipse) GetConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *Local2dPointUncertaintyEllipse) GetConfidenceOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *Local2dPointUncertaintyEllipse) SetConfidence(v int32) {
	o.Confidence = v
}

func (o Local2dPointUncertaintyEllipse) MarshalJSON() ([]byte, error) {
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
		toSerialize["localOrigin"] = o.LocalOrigin
	}
	if true {
		toSerialize["point"] = o.Point
	}
	if true {
		toSerialize["uncertaintyEllipse"] = o.UncertaintyEllipse
	}
	if true {
		toSerialize["confidence"] = o.Confidence
	}
	return json.Marshal(toSerialize)
}

type NullableLocal2dPointUncertaintyEllipse struct {
	value *Local2dPointUncertaintyEllipse
	isSet bool
}

func (v NullableLocal2dPointUncertaintyEllipse) Get() *Local2dPointUncertaintyEllipse {
	return v.value
}

func (v *NullableLocal2dPointUncertaintyEllipse) Set(val *Local2dPointUncertaintyEllipse) {
	v.value = val
	v.isSet = true
}

func (v NullableLocal2dPointUncertaintyEllipse) IsSet() bool {
	return v.isSet
}

func (v *NullableLocal2dPointUncertaintyEllipse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocal2dPointUncertaintyEllipse(val *Local2dPointUncertaintyEllipse) *NullableLocal2dPointUncertaintyEllipse {
	return &NullableLocal2dPointUncertaintyEllipse{value: val, isSet: true}
}

func (v NullableLocal2dPointUncertaintyEllipse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocal2dPointUncertaintyEllipse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

