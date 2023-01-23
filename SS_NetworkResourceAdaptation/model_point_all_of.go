/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceAdaptation

import (
	"encoding/json"
)

// PointAllOf struct for PointAllOf
type PointAllOf struct {
	Point GeographicalCoordinates `json:"point"`
}

// NewPointAllOf instantiates a new PointAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointAllOf(point GeographicalCoordinates) *PointAllOf {
	this := PointAllOf{}
	this.Point = point
	return &this
}

// NewPointAllOfWithDefaults instantiates a new PointAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointAllOfWithDefaults() *PointAllOf {
	this := PointAllOf{}
	return &this
}

// GetPoint returns the Point field value
func (o *PointAllOf) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *PointAllOf) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *PointAllOf) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

func (o PointAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["point"] = o.Point
	}
	return json.Marshal(toSerialize)
}

type NullablePointAllOf struct {
	value *PointAllOf
	isSet bool
}

func (v NullablePointAllOf) Get() *PointAllOf {
	return v.value
}

func (v *NullablePointAllOf) Set(val *PointAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePointAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePointAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointAllOf(val *PointAllOf) *NullablePointAllOf {
	return &NullablePointAllOf{value: val, isSet: true}
}

func (v NullablePointAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


