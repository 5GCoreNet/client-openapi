/*
GMDviaMBMSbyxMB

API for Group Message Delivery via MBMS by xMB   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GMDviaMBMSbyxMB

import (
	"encoding/json"
)

// GeographicalCoordinates Geographical coordinates.
type GeographicalCoordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// NewGeographicalCoordinates instantiates a new GeographicalCoordinates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeographicalCoordinates(lon float64, lat float64) *GeographicalCoordinates {
	this := GeographicalCoordinates{}
	this.Lon = lon
	this.Lat = lat
	return &this
}

// NewGeographicalCoordinatesWithDefaults instantiates a new GeographicalCoordinates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeographicalCoordinatesWithDefaults() *GeographicalCoordinates {
	this := GeographicalCoordinates{}
	return &this
}

// GetLon returns the Lon field value
func (o *GeographicalCoordinates) GetLon() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Lon
}

// GetLonOk returns a tuple with the Lon field value
// and a boolean to check if the value has been set.
func (o *GeographicalCoordinates) GetLonOk() (*float64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Lon, true
}

// SetLon sets field value
func (o *GeographicalCoordinates) SetLon(v float64) {
	o.Lon = v
}

// GetLat returns the Lat field value
func (o *GeographicalCoordinates) GetLat() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Lat
}

// GetLatOk returns a tuple with the Lat field value
// and a boolean to check if the value has been set.
func (o *GeographicalCoordinates) GetLatOk() (*float64, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Lat, true
}

// SetLat sets field value
func (o *GeographicalCoordinates) SetLat(v float64) {
	o.Lat = v
}

func (o GeographicalCoordinates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["lon"] = o.Lon
	}
	if true {
		toSerialize["lat"] = o.Lat
	}
	return json.Marshal(toSerialize)
}

type NullableGeographicalCoordinates struct {
	value *GeographicalCoordinates
	isSet bool
}

func (v NullableGeographicalCoordinates) Get() *GeographicalCoordinates {
	return v.value
}

func (v *NullableGeographicalCoordinates) Set(val *GeographicalCoordinates) {
	v.value = val
	v.isSet = true
}

func (v NullableGeographicalCoordinates) IsSet() bool {
	return v.isSet
}

func (v *NullableGeographicalCoordinates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeographicalCoordinates(val *GeographicalCoordinates) *NullableGeographicalCoordinates {
	return &NullableGeographicalCoordinates{value: val, isSet: true}
}

func (v NullableGeographicalCoordinates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeographicalCoordinates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


