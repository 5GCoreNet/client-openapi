/*
ECS EES Registration_API

API for EES Registration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eecs_EESRegistration

import (
	"encoding/json"
)

// ServiceArea Represents a service area information of the EdgeApp entity.
type ServiceArea struct {
	TopServAr *TopologicalServiceArea `json:"topServAr,omitempty"`
	GeoServAr *GeographicalServiceArea `json:"geoServAr,omitempty"`
}

// NewServiceArea instantiates a new ServiceArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceArea() *ServiceArea {
	this := ServiceArea{}
	return &this
}

// NewServiceAreaWithDefaults instantiates a new ServiceArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAreaWithDefaults() *ServiceArea {
	this := ServiceArea{}
	return &this
}

// GetTopServAr returns the TopServAr field value if set, zero value otherwise.
func (o *ServiceArea) GetTopServAr() TopologicalServiceArea {
	if o == nil || isNil(o.TopServAr) {
		var ret TopologicalServiceArea
		return ret
	}
	return *o.TopServAr
}

// GetTopServArOk returns a tuple with the TopServAr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceArea) GetTopServArOk() (*TopologicalServiceArea, bool) {
	if o == nil || isNil(o.TopServAr) {
    return nil, false
	}
	return o.TopServAr, true
}

// HasTopServAr returns a boolean if a field has been set.
func (o *ServiceArea) HasTopServAr() bool {
	if o != nil && !isNil(o.TopServAr) {
		return true
	}

	return false
}

// SetTopServAr gets a reference to the given TopologicalServiceArea and assigns it to the TopServAr field.
func (o *ServiceArea) SetTopServAr(v TopologicalServiceArea) {
	o.TopServAr = &v
}

// GetGeoServAr returns the GeoServAr field value if set, zero value otherwise.
func (o *ServiceArea) GetGeoServAr() GeographicalServiceArea {
	if o == nil || isNil(o.GeoServAr) {
		var ret GeographicalServiceArea
		return ret
	}
	return *o.GeoServAr
}

// GetGeoServArOk returns a tuple with the GeoServAr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceArea) GetGeoServArOk() (*GeographicalServiceArea, bool) {
	if o == nil || isNil(o.GeoServAr) {
    return nil, false
	}
	return o.GeoServAr, true
}

// HasGeoServAr returns a boolean if a field has been set.
func (o *ServiceArea) HasGeoServAr() bool {
	if o != nil && !isNil(o.GeoServAr) {
		return true
	}

	return false
}

// SetGeoServAr gets a reference to the given GeographicalServiceArea and assigns it to the GeoServAr field.
func (o *ServiceArea) SetGeoServAr(v GeographicalServiceArea) {
	o.GeoServAr = &v
}

func (o ServiceArea) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TopServAr) {
		toSerialize["topServAr"] = o.TopServAr
	}
	if !isNil(o.GeoServAr) {
		toSerialize["geoServAr"] = o.GeoServAr
	}
	return json.Marshal(toSerialize)
}

type NullableServiceArea struct {
	value *ServiceArea
	isSet bool
}

func (v NullableServiceArea) Get() *ServiceArea {
	return v.value
}

func (v *NullableServiceArea) Set(val *ServiceArea) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceArea) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceArea(val *ServiceArea) *NullableServiceArea {
	return &NullableServiceArea{value: val, isSet: true}
}

func (v NullableServiceArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


