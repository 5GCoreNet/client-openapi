/*
Nnwdaf_MLModelProvision

Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_MLModelProvision

import (
	"encoding/json"
)

// Polygon Polygon.
type Polygon struct {
	GADShape
	// List of points.
	PointList []GeographicalCoordinates `json:"pointList"`
}

// NewPolygon instantiates a new Polygon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolygon(pointList []GeographicalCoordinates, shape SupportedGADShapes) *Polygon {
	this := Polygon{}
	this.Shape = shape
	this.PointList = pointList
	return &this
}

// NewPolygonWithDefaults instantiates a new Polygon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolygonWithDefaults() *Polygon {
	this := Polygon{}
	return &this
}

// GetPointList returns the PointList field value
func (o *Polygon) GetPointList() []GeographicalCoordinates {
	if o == nil {
		var ret []GeographicalCoordinates
		return ret
	}

	return o.PointList
}

// GetPointListOk returns a tuple with the PointList field value
// and a boolean to check if the value has been set.
func (o *Polygon) GetPointListOk() ([]GeographicalCoordinates, bool) {
	if o == nil {
    return nil, false
	}
	return o.PointList, true
}

// SetPointList sets field value
func (o *Polygon) SetPointList(v []GeographicalCoordinates) {
	o.PointList = v
}

func (o Polygon) MarshalJSON() ([]byte, error) {
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
		toSerialize["pointList"] = o.PointList
	}
	return json.Marshal(toSerialize)
}

type NullablePolygon struct {
	value *Polygon
	isSet bool
}

func (v NullablePolygon) Get() *Polygon {
	return v.value
}

func (v *NullablePolygon) Set(val *Polygon) {
	v.value = val
	v.isSet = true
}

func (v NullablePolygon) IsSet() bool {
	return v.isSet
}

func (v *NullablePolygon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolygon(val *Polygon) *NullablePolygon {
	return &NullablePolygon{value: val, isSet: true}
}

func (v NullablePolygon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolygon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


