/*
EES EAS Registration_API

API for EAS Registration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_EASRegistration

import (
	"encoding/json"
)

// GeographicalServiceArea Represents geographical service area information.
type GeographicalServiceArea struct {
	// A list of geographic area information.
	GeoArs []GeographicArea `json:"geoArs,omitempty"`
	// A list of civic address information.
	CivicAddrs []CivicAddress `json:"civicAddrs,omitempty"`
}

// NewGeographicalServiceArea instantiates a new GeographicalServiceArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeographicalServiceArea() *GeographicalServiceArea {
	this := GeographicalServiceArea{}
	return &this
}

// NewGeographicalServiceAreaWithDefaults instantiates a new GeographicalServiceArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeographicalServiceAreaWithDefaults() *GeographicalServiceArea {
	this := GeographicalServiceArea{}
	return &this
}

// GetGeoArs returns the GeoArs field value if set, zero value otherwise.
func (o *GeographicalServiceArea) GetGeoArs() []GeographicArea {
	if o == nil || isNil(o.GeoArs) {
		var ret []GeographicArea
		return ret
	}
	return o.GeoArs
}

// GetGeoArsOk returns a tuple with the GeoArs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicalServiceArea) GetGeoArsOk() ([]GeographicArea, bool) {
	if o == nil || isNil(o.GeoArs) {
    return nil, false
	}
	return o.GeoArs, true
}

// HasGeoArs returns a boolean if a field has been set.
func (o *GeographicalServiceArea) HasGeoArs() bool {
	if o != nil && !isNil(o.GeoArs) {
		return true
	}

	return false
}

// SetGeoArs gets a reference to the given []GeographicArea and assigns it to the GeoArs field.
func (o *GeographicalServiceArea) SetGeoArs(v []GeographicArea) {
	o.GeoArs = v
}

// GetCivicAddrs returns the CivicAddrs field value if set, zero value otherwise.
func (o *GeographicalServiceArea) GetCivicAddrs() []CivicAddress {
	if o == nil || isNil(o.CivicAddrs) {
		var ret []CivicAddress
		return ret
	}
	return o.CivicAddrs
}

// GetCivicAddrsOk returns a tuple with the CivicAddrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeographicalServiceArea) GetCivicAddrsOk() ([]CivicAddress, bool) {
	if o == nil || isNil(o.CivicAddrs) {
    return nil, false
	}
	return o.CivicAddrs, true
}

// HasCivicAddrs returns a boolean if a field has been set.
func (o *GeographicalServiceArea) HasCivicAddrs() bool {
	if o != nil && !isNil(o.CivicAddrs) {
		return true
	}

	return false
}

// SetCivicAddrs gets a reference to the given []CivicAddress and assigns it to the CivicAddrs field.
func (o *GeographicalServiceArea) SetCivicAddrs(v []CivicAddress) {
	o.CivicAddrs = v
}

func (o GeographicalServiceArea) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GeoArs) {
		toSerialize["geoArs"] = o.GeoArs
	}
	if !isNil(o.CivicAddrs) {
		toSerialize["civicAddrs"] = o.CivicAddrs
	}
	return json.Marshal(toSerialize)
}

type NullableGeographicalServiceArea struct {
	value *GeographicalServiceArea
	isSet bool
}

func (v NullableGeographicalServiceArea) Get() *GeographicalServiceArea {
	return v.value
}

func (v *NullableGeographicalServiceArea) Set(val *GeographicalServiceArea) {
	v.value = val
	v.isSet = true
}

func (v NullableGeographicalServiceArea) IsSet() bool {
	return v.isSet
}

func (v *NullableGeographicalServiceArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeographicalServiceArea(val *GeographicalServiceArea) *NullableGeographicalServiceArea {
	return &NullableGeographicalServiceArea{value: val, isSet: true}
}

func (v NullableGeographicalServiceArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeographicalServiceArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


