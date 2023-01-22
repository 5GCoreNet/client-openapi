/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_DataManagement

import (
	"encoding/json"
)

// LocationReport struct for LocationReport
type LocationReport struct {
	Location UserLocation `json:"location"`
}

// NewLocationReport instantiates a new LocationReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationReport(location UserLocation) *LocationReport {
	this := LocationReport{}
	this.Location = location
	return &this
}

// NewLocationReportWithDefaults instantiates a new LocationReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationReportWithDefaults() *LocationReport {
	this := LocationReport{}
	return &this
}

// GetLocation returns the Location field value
func (o *LocationReport) GetLocation() UserLocation {
	if o == nil {
		var ret UserLocation
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *LocationReport) GetLocationOk() (*UserLocation, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *LocationReport) SetLocation(v UserLocation) {
	o.Location = v
}

func (o LocationReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableLocationReport struct {
	value *LocationReport
	isSet bool
}

func (v NullableLocationReport) Get() *LocationReport {
	return v.value
}

func (v *NullableLocationReport) Set(val *LocationReport) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationReport) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationReport(val *LocationReport) *NullableLocationReport {
	return &NullableLocationReport{value: val, isSet: true}
}

func (v NullableLocationReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


