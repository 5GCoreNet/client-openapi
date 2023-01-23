/*
CAPIF_API_Invoker_Management_API

API for API invoker management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_API_Invoker_Management_API

import (
	"encoding/json"
)

// AefLocation The location information (e.g. civic address, GPS coordinates, data center ID) where the AEF providing the service API is located. 
type AefLocation struct {
	CivicAddr *CivicAddress `json:"civicAddr,omitempty"`
	GeoArea *GeographicArea `json:"geoArea,omitempty"`
	// Identifies the data center where the AEF providing the service API is located. 
	DcId *string `json:"dcId,omitempty"`
}

// NewAefLocation instantiates a new AefLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAefLocation() *AefLocation {
	this := AefLocation{}
	return &this
}

// NewAefLocationWithDefaults instantiates a new AefLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAefLocationWithDefaults() *AefLocation {
	this := AefLocation{}
	return &this
}

// GetCivicAddr returns the CivicAddr field value if set, zero value otherwise.
func (o *AefLocation) GetCivicAddr() CivicAddress {
	if o == nil || isNil(o.CivicAddr) {
		var ret CivicAddress
		return ret
	}
	return *o.CivicAddr
}

// GetCivicAddrOk returns a tuple with the CivicAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AefLocation) GetCivicAddrOk() (*CivicAddress, bool) {
	if o == nil || isNil(o.CivicAddr) {
    return nil, false
	}
	return o.CivicAddr, true
}

// HasCivicAddr returns a boolean if a field has been set.
func (o *AefLocation) HasCivicAddr() bool {
	if o != nil && !isNil(o.CivicAddr) {
		return true
	}

	return false
}

// SetCivicAddr gets a reference to the given CivicAddress and assigns it to the CivicAddr field.
func (o *AefLocation) SetCivicAddr(v CivicAddress) {
	o.CivicAddr = &v
}

// GetGeoArea returns the GeoArea field value if set, zero value otherwise.
func (o *AefLocation) GetGeoArea() GeographicArea {
	if o == nil || isNil(o.GeoArea) {
		var ret GeographicArea
		return ret
	}
	return *o.GeoArea
}

// GetGeoAreaOk returns a tuple with the GeoArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AefLocation) GetGeoAreaOk() (*GeographicArea, bool) {
	if o == nil || isNil(o.GeoArea) {
    return nil, false
	}
	return o.GeoArea, true
}

// HasGeoArea returns a boolean if a field has been set.
func (o *AefLocation) HasGeoArea() bool {
	if o != nil && !isNil(o.GeoArea) {
		return true
	}

	return false
}

// SetGeoArea gets a reference to the given GeographicArea and assigns it to the GeoArea field.
func (o *AefLocation) SetGeoArea(v GeographicArea) {
	o.GeoArea = &v
}

// GetDcId returns the DcId field value if set, zero value otherwise.
func (o *AefLocation) GetDcId() string {
	if o == nil || isNil(o.DcId) {
		var ret string
		return ret
	}
	return *o.DcId
}

// GetDcIdOk returns a tuple with the DcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AefLocation) GetDcIdOk() (*string, bool) {
	if o == nil || isNil(o.DcId) {
    return nil, false
	}
	return o.DcId, true
}

// HasDcId returns a boolean if a field has been set.
func (o *AefLocation) HasDcId() bool {
	if o != nil && !isNil(o.DcId) {
		return true
	}

	return false
}

// SetDcId gets a reference to the given string and assigns it to the DcId field.
func (o *AefLocation) SetDcId(v string) {
	o.DcId = &v
}

func (o AefLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CivicAddr) {
		toSerialize["civicAddr"] = o.CivicAddr
	}
	if !isNil(o.GeoArea) {
		toSerialize["geoArea"] = o.GeoArea
	}
	if !isNil(o.DcId) {
		toSerialize["dcId"] = o.DcId
	}
	return json.Marshal(toSerialize)
}

type NullableAefLocation struct {
	value *AefLocation
	isSet bool
}

func (v NullableAefLocation) Get() *AefLocation {
	return v.value
}

func (v *NullableAefLocation) Set(val *AefLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableAefLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableAefLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAefLocation(val *AefLocation) *NullableAefLocation {
	return &NullableAefLocation{value: val, isSet: true}
}

func (v NullableAefLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAefLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


