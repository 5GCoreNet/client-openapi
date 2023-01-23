/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// LocationArea1 struct for LocationArea1
type LocationArea1 struct {
	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`
	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []CivicAddress `json:"civicAddresses,omitempty"`
	NwAreaInfo *NetworkAreaInfo1 `json:"nwAreaInfo,omitempty"`
	UmtTime *UmtTime1 `json:"umtTime,omitempty"`
}

// NewLocationArea1 instantiates a new LocationArea1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationArea1() *LocationArea1 {
	this := LocationArea1{}
	return &this
}

// NewLocationArea1WithDefaults instantiates a new LocationArea1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationArea1WithDefaults() *LocationArea1 {
	this := LocationArea1{}
	return &this
}

// GetGeographicAreas returns the GeographicAreas field value if set, zero value otherwise.
func (o *LocationArea1) GetGeographicAreas() []GeographicArea {
	if o == nil || isNil(o.GeographicAreas) {
		var ret []GeographicArea
		return ret
	}
	return o.GeographicAreas
}

// GetGeographicAreasOk returns a tuple with the GeographicAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea1) GetGeographicAreasOk() ([]GeographicArea, bool) {
	if o == nil || isNil(o.GeographicAreas) {
    return nil, false
	}
	return o.GeographicAreas, true
}

// HasGeographicAreas returns a boolean if a field has been set.
func (o *LocationArea1) HasGeographicAreas() bool {
	if o != nil && !isNil(o.GeographicAreas) {
		return true
	}

	return false
}

// SetGeographicAreas gets a reference to the given []GeographicArea and assigns it to the GeographicAreas field.
func (o *LocationArea1) SetGeographicAreas(v []GeographicArea) {
	o.GeographicAreas = v
}

// GetCivicAddresses returns the CivicAddresses field value if set, zero value otherwise.
func (o *LocationArea1) GetCivicAddresses() []CivicAddress {
	if o == nil || isNil(o.CivicAddresses) {
		var ret []CivicAddress
		return ret
	}
	return o.CivicAddresses
}

// GetCivicAddressesOk returns a tuple with the CivicAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea1) GetCivicAddressesOk() ([]CivicAddress, bool) {
	if o == nil || isNil(o.CivicAddresses) {
    return nil, false
	}
	return o.CivicAddresses, true
}

// HasCivicAddresses returns a boolean if a field has been set.
func (o *LocationArea1) HasCivicAddresses() bool {
	if o != nil && !isNil(o.CivicAddresses) {
		return true
	}

	return false
}

// SetCivicAddresses gets a reference to the given []CivicAddress and assigns it to the CivicAddresses field.
func (o *LocationArea1) SetCivicAddresses(v []CivicAddress) {
	o.CivicAddresses = v
}

// GetNwAreaInfo returns the NwAreaInfo field value if set, zero value otherwise.
func (o *LocationArea1) GetNwAreaInfo() NetworkAreaInfo1 {
	if o == nil || isNil(o.NwAreaInfo) {
		var ret NetworkAreaInfo1
		return ret
	}
	return *o.NwAreaInfo
}

// GetNwAreaInfoOk returns a tuple with the NwAreaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea1) GetNwAreaInfoOk() (*NetworkAreaInfo1, bool) {
	if o == nil || isNil(o.NwAreaInfo) {
    return nil, false
	}
	return o.NwAreaInfo, true
}

// HasNwAreaInfo returns a boolean if a field has been set.
func (o *LocationArea1) HasNwAreaInfo() bool {
	if o != nil && !isNil(o.NwAreaInfo) {
		return true
	}

	return false
}

// SetNwAreaInfo gets a reference to the given NetworkAreaInfo1 and assigns it to the NwAreaInfo field.
func (o *LocationArea1) SetNwAreaInfo(v NetworkAreaInfo1) {
	o.NwAreaInfo = &v
}

// GetUmtTime returns the UmtTime field value if set, zero value otherwise.
func (o *LocationArea1) GetUmtTime() UmtTime1 {
	if o == nil || isNil(o.UmtTime) {
		var ret UmtTime1
		return ret
	}
	return *o.UmtTime
}

// GetUmtTimeOk returns a tuple with the UmtTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea1) GetUmtTimeOk() (*UmtTime1, bool) {
	if o == nil || isNil(o.UmtTime) {
    return nil, false
	}
	return o.UmtTime, true
}

// HasUmtTime returns a boolean if a field has been set.
func (o *LocationArea1) HasUmtTime() bool {
	if o != nil && !isNil(o.UmtTime) {
		return true
	}

	return false
}

// SetUmtTime gets a reference to the given UmtTime1 and assigns it to the UmtTime field.
func (o *LocationArea1) SetUmtTime(v UmtTime1) {
	o.UmtTime = &v
}

func (o LocationArea1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GeographicAreas) {
		toSerialize["geographicAreas"] = o.GeographicAreas
	}
	if !isNil(o.CivicAddresses) {
		toSerialize["civicAddresses"] = o.CivicAddresses
	}
	if !isNil(o.NwAreaInfo) {
		toSerialize["nwAreaInfo"] = o.NwAreaInfo
	}
	if !isNil(o.UmtTime) {
		toSerialize["umtTime"] = o.UmtTime
	}
	return json.Marshal(toSerialize)
}

type NullableLocationArea1 struct {
	value *LocationArea1
	isSet bool
}

func (v NullableLocationArea1) Get() *LocationArea1 {
	return v.value
}

func (v *NullableLocationArea1) Set(val *LocationArea1) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationArea1) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationArea1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationArea1(val *LocationArea1) *NullableLocationArea1 {
	return &NullableLocationArea1{value: val, isSet: true}
}

func (v NullableLocationArea1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationArea1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


