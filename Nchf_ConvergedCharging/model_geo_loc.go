/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// GeoLoc struct for GeoLoc
type GeoLoc struct {
	GeographicalCoordinates *GeographicalCoordinates `json:"geographicalCoordinates,omitempty"`
	CivicLocation *string `json:"civicLocation,omitempty"`
}

// NewGeoLoc instantiates a new GeoLoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoLoc() *GeoLoc {
	this := GeoLoc{}
	return &this
}

// NewGeoLocWithDefaults instantiates a new GeoLoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoLocWithDefaults() *GeoLoc {
	this := GeoLoc{}
	return &this
}

// GetGeographicalCoordinates returns the GeographicalCoordinates field value if set, zero value otherwise.
func (o *GeoLoc) GetGeographicalCoordinates() GeographicalCoordinates {
	if o == nil || isNil(o.GeographicalCoordinates) {
		var ret GeographicalCoordinates
		return ret
	}
	return *o.GeographicalCoordinates
}

// GetGeographicalCoordinatesOk returns a tuple with the GeographicalCoordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLoc) GetGeographicalCoordinatesOk() (*GeographicalCoordinates, bool) {
	if o == nil || isNil(o.GeographicalCoordinates) {
    return nil, false
	}
	return o.GeographicalCoordinates, true
}

// HasGeographicalCoordinates returns a boolean if a field has been set.
func (o *GeoLoc) HasGeographicalCoordinates() bool {
	if o != nil && !isNil(o.GeographicalCoordinates) {
		return true
	}

	return false
}

// SetGeographicalCoordinates gets a reference to the given GeographicalCoordinates and assigns it to the GeographicalCoordinates field.
func (o *GeoLoc) SetGeographicalCoordinates(v GeographicalCoordinates) {
	o.GeographicalCoordinates = &v
}

// GetCivicLocation returns the CivicLocation field value if set, zero value otherwise.
func (o *GeoLoc) GetCivicLocation() string {
	if o == nil || isNil(o.CivicLocation) {
		var ret string
		return ret
	}
	return *o.CivicLocation
}

// GetCivicLocationOk returns a tuple with the CivicLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoLoc) GetCivicLocationOk() (*string, bool) {
	if o == nil || isNil(o.CivicLocation) {
    return nil, false
	}
	return o.CivicLocation, true
}

// HasCivicLocation returns a boolean if a field has been set.
func (o *GeoLoc) HasCivicLocation() bool {
	if o != nil && !isNil(o.CivicLocation) {
		return true
	}

	return false
}

// SetCivicLocation gets a reference to the given string and assigns it to the CivicLocation field.
func (o *GeoLoc) SetCivicLocation(v string) {
	o.CivicLocation = &v
}

func (o GeoLoc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GeographicalCoordinates) {
		toSerialize["geographicalCoordinates"] = o.GeographicalCoordinates
	}
	if !isNil(o.CivicLocation) {
		toSerialize["civicLocation"] = o.CivicLocation
	}
	return json.Marshal(toSerialize)
}

type NullableGeoLoc struct {
	value *GeoLoc
	isSet bool
}

func (v NullableGeoLoc) Get() *GeoLoc {
	return v.value
}

func (v *NullableGeoLoc) Set(val *GeoLoc) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoLoc) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoLoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoLoc(val *GeoLoc) *NullableGeoLoc {
	return &NullableGeoLoc{value: val, isSet: true}
}

func (v NullableGeoLoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoLoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


