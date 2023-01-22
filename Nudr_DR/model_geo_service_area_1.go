/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// GeoServiceArea1 List of geographic area or list of civic address info
type GeoServiceArea1 struct {
	GeographicAreaList []GeographicArea `json:"geographicAreaList,omitempty"`
	CivicAddressList []CivicAddress `json:"civicAddressList,omitempty"`
}

// NewGeoServiceArea1 instantiates a new GeoServiceArea1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoServiceArea1() *GeoServiceArea1 {
	this := GeoServiceArea1{}
	return &this
}

// NewGeoServiceArea1WithDefaults instantiates a new GeoServiceArea1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoServiceArea1WithDefaults() *GeoServiceArea1 {
	this := GeoServiceArea1{}
	return &this
}

// GetGeographicAreaList returns the GeographicAreaList field value if set, zero value otherwise.
func (o *GeoServiceArea1) GetGeographicAreaList() []GeographicArea {
	if o == nil || isNil(o.GeographicAreaList) {
		var ret []GeographicArea
		return ret
	}
	return o.GeographicAreaList
}

// GetGeographicAreaListOk returns a tuple with the GeographicAreaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoServiceArea1) GetGeographicAreaListOk() ([]GeographicArea, bool) {
	if o == nil || isNil(o.GeographicAreaList) {
    return nil, false
	}
	return o.GeographicAreaList, true
}

// HasGeographicAreaList returns a boolean if a field has been set.
func (o *GeoServiceArea1) HasGeographicAreaList() bool {
	if o != nil && !isNil(o.GeographicAreaList) {
		return true
	}

	return false
}

// SetGeographicAreaList gets a reference to the given []GeographicArea and assigns it to the GeographicAreaList field.
func (o *GeoServiceArea1) SetGeographicAreaList(v []GeographicArea) {
	o.GeographicAreaList = v
}

// GetCivicAddressList returns the CivicAddressList field value if set, zero value otherwise.
func (o *GeoServiceArea1) GetCivicAddressList() []CivicAddress {
	if o == nil || isNil(o.CivicAddressList) {
		var ret []CivicAddress
		return ret
	}
	return o.CivicAddressList
}

// GetCivicAddressListOk returns a tuple with the CivicAddressList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoServiceArea1) GetCivicAddressListOk() ([]CivicAddress, bool) {
	if o == nil || isNil(o.CivicAddressList) {
    return nil, false
	}
	return o.CivicAddressList, true
}

// HasCivicAddressList returns a boolean if a field has been set.
func (o *GeoServiceArea1) HasCivicAddressList() bool {
	if o != nil && !isNil(o.CivicAddressList) {
		return true
	}

	return false
}

// SetCivicAddressList gets a reference to the given []CivicAddress and assigns it to the CivicAddressList field.
func (o *GeoServiceArea1) SetCivicAddressList(v []CivicAddress) {
	o.CivicAddressList = v
}

func (o GeoServiceArea1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GeographicAreaList) {
		toSerialize["geographicAreaList"] = o.GeographicAreaList
	}
	if !isNil(o.CivicAddressList) {
		toSerialize["civicAddressList"] = o.CivicAddressList
	}
	return json.Marshal(toSerialize)
}

type NullableGeoServiceArea1 struct {
	value *GeoServiceArea1
	isSet bool
}

func (v NullableGeoServiceArea1) Get() *GeoServiceArea1 {
	return v.value
}

func (v *NullableGeoServiceArea1) Set(val *GeoServiceArea1) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoServiceArea1) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoServiceArea1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoServiceArea1(val *GeoServiceArea1) *NullableGeoServiceArea1 {
	return &NullableGeoServiceArea1{value: val, isSet: true}
}

func (v NullableGeoServiceArea1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoServiceArea1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


