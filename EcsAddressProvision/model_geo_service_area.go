/*
3gpp-ecs-address-provision

API for ECS Address Provisioning.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EcsAddressProvision

import (
	"encoding/json"
)

// GeoServiceArea List of geographic area or list of civic address info
type GeoServiceArea struct {
	GeographicAreaList []GeographicArea `json:"geographicAreaList,omitempty"`
	CivicAddressList []CivicAddress `json:"civicAddressList,omitempty"`
}

// NewGeoServiceArea instantiates a new GeoServiceArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoServiceArea() *GeoServiceArea {
	this := GeoServiceArea{}
	return &this
}

// NewGeoServiceAreaWithDefaults instantiates a new GeoServiceArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoServiceAreaWithDefaults() *GeoServiceArea {
	this := GeoServiceArea{}
	return &this
}

// GetGeographicAreaList returns the GeographicAreaList field value if set, zero value otherwise.
func (o *GeoServiceArea) GetGeographicAreaList() []GeographicArea {
	if o == nil || isNil(o.GeographicAreaList) {
		var ret []GeographicArea
		return ret
	}
	return o.GeographicAreaList
}

// GetGeographicAreaListOk returns a tuple with the GeographicAreaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoServiceArea) GetGeographicAreaListOk() ([]GeographicArea, bool) {
	if o == nil || isNil(o.GeographicAreaList) {
    return nil, false
	}
	return o.GeographicAreaList, true
}

// HasGeographicAreaList returns a boolean if a field has been set.
func (o *GeoServiceArea) HasGeographicAreaList() bool {
	if o != nil && !isNil(o.GeographicAreaList) {
		return true
	}

	return false
}

// SetGeographicAreaList gets a reference to the given []GeographicArea and assigns it to the GeographicAreaList field.
func (o *GeoServiceArea) SetGeographicAreaList(v []GeographicArea) {
	o.GeographicAreaList = v
}

// GetCivicAddressList returns the CivicAddressList field value if set, zero value otherwise.
func (o *GeoServiceArea) GetCivicAddressList() []CivicAddress {
	if o == nil || isNil(o.CivicAddressList) {
		var ret []CivicAddress
		return ret
	}
	return o.CivicAddressList
}

// GetCivicAddressListOk returns a tuple with the CivicAddressList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoServiceArea) GetCivicAddressListOk() ([]CivicAddress, bool) {
	if o == nil || isNil(o.CivicAddressList) {
    return nil, false
	}
	return o.CivicAddressList, true
}

// HasCivicAddressList returns a boolean if a field has been set.
func (o *GeoServiceArea) HasCivicAddressList() bool {
	if o != nil && !isNil(o.CivicAddressList) {
		return true
	}

	return false
}

// SetCivicAddressList gets a reference to the given []CivicAddress and assigns it to the CivicAddressList field.
func (o *GeoServiceArea) SetCivicAddressList(v []CivicAddress) {
	o.CivicAddressList = v
}

func (o GeoServiceArea) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GeographicAreaList) {
		toSerialize["geographicAreaList"] = o.GeographicAreaList
	}
	if !isNil(o.CivicAddressList) {
		toSerialize["civicAddressList"] = o.CivicAddressList
	}
	return json.Marshal(toSerialize)
}

type NullableGeoServiceArea struct {
	value *GeoServiceArea
	isSet bool
}

func (v NullableGeoServiceArea) Get() *GeoServiceArea {
	return v.value
}

func (v *NullableGeoServiceArea) Set(val *GeoServiceArea) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoServiceArea) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoServiceArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoServiceArea(val *GeoServiceArea) *NullableGeoServiceArea {
	return &NullableGeoServiceArea{value: val, isSet: true}
}

func (v NullableGeoServiceArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoServiceArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


