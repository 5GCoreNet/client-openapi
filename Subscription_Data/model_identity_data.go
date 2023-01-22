/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
)

// IdentityData Identity data corresponds to the provided ueId.
type IdentityData struct {
	SupiList []string `json:"supiList,omitempty"`
	GpsiList []string `json:"gpsiList,omitempty"`
	AllowedAfIds []string `json:"allowedAfIds,omitempty"`
	// A map (list of key-value pairs where AppPortId serves as key) of GPSIs.
	ApplicationPortIds *map[string]string `json:"applicationPortIds,omitempty"`
}

// NewIdentityData instantiates a new IdentityData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityData() *IdentityData {
	this := IdentityData{}
	return &this
}

// NewIdentityDataWithDefaults instantiates a new IdentityData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityDataWithDefaults() *IdentityData {
	this := IdentityData{}
	return &this
}

// GetSupiList returns the SupiList field value if set, zero value otherwise.
func (o *IdentityData) GetSupiList() []string {
	if o == nil || isNil(o.SupiList) {
		var ret []string
		return ret
	}
	return o.SupiList
}

// GetSupiListOk returns a tuple with the SupiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityData) GetSupiListOk() ([]string, bool) {
	if o == nil || isNil(o.SupiList) {
    return nil, false
	}
	return o.SupiList, true
}

// HasSupiList returns a boolean if a field has been set.
func (o *IdentityData) HasSupiList() bool {
	if o != nil && !isNil(o.SupiList) {
		return true
	}

	return false
}

// SetSupiList gets a reference to the given []string and assigns it to the SupiList field.
func (o *IdentityData) SetSupiList(v []string) {
	o.SupiList = v
}

// GetGpsiList returns the GpsiList field value if set, zero value otherwise.
func (o *IdentityData) GetGpsiList() []string {
	if o == nil || isNil(o.GpsiList) {
		var ret []string
		return ret
	}
	return o.GpsiList
}

// GetGpsiListOk returns a tuple with the GpsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityData) GetGpsiListOk() ([]string, bool) {
	if o == nil || isNil(o.GpsiList) {
    return nil, false
	}
	return o.GpsiList, true
}

// HasGpsiList returns a boolean if a field has been set.
func (o *IdentityData) HasGpsiList() bool {
	if o != nil && !isNil(o.GpsiList) {
		return true
	}

	return false
}

// SetGpsiList gets a reference to the given []string and assigns it to the GpsiList field.
func (o *IdentityData) SetGpsiList(v []string) {
	o.GpsiList = v
}

// GetAllowedAfIds returns the AllowedAfIds field value if set, zero value otherwise.
func (o *IdentityData) GetAllowedAfIds() []string {
	if o == nil || isNil(o.AllowedAfIds) {
		var ret []string
		return ret
	}
	return o.AllowedAfIds
}

// GetAllowedAfIdsOk returns a tuple with the AllowedAfIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityData) GetAllowedAfIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedAfIds) {
    return nil, false
	}
	return o.AllowedAfIds, true
}

// HasAllowedAfIds returns a boolean if a field has been set.
func (o *IdentityData) HasAllowedAfIds() bool {
	if o != nil && !isNil(o.AllowedAfIds) {
		return true
	}

	return false
}

// SetAllowedAfIds gets a reference to the given []string and assigns it to the AllowedAfIds field.
func (o *IdentityData) SetAllowedAfIds(v []string) {
	o.AllowedAfIds = v
}

// GetApplicationPortIds returns the ApplicationPortIds field value if set, zero value otherwise.
func (o *IdentityData) GetApplicationPortIds() map[string]string {
	if o == nil || isNil(o.ApplicationPortIds) {
		var ret map[string]string
		return ret
	}
	return *o.ApplicationPortIds
}

// GetApplicationPortIdsOk returns a tuple with the ApplicationPortIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityData) GetApplicationPortIdsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ApplicationPortIds) {
    return nil, false
	}
	return o.ApplicationPortIds, true
}

// HasApplicationPortIds returns a boolean if a field has been set.
func (o *IdentityData) HasApplicationPortIds() bool {
	if o != nil && !isNil(o.ApplicationPortIds) {
		return true
	}

	return false
}

// SetApplicationPortIds gets a reference to the given map[string]string and assigns it to the ApplicationPortIds field.
func (o *IdentityData) SetApplicationPortIds(v map[string]string) {
	o.ApplicationPortIds = &v
}

func (o IdentityData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SupiList) {
		toSerialize["supiList"] = o.SupiList
	}
	if !isNil(o.GpsiList) {
		toSerialize["gpsiList"] = o.GpsiList
	}
	if !isNil(o.AllowedAfIds) {
		toSerialize["allowedAfIds"] = o.AllowedAfIds
	}
	if !isNil(o.ApplicationPortIds) {
		toSerialize["applicationPortIds"] = o.ApplicationPortIds
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityData struct {
	value *IdentityData
	isSet bool
}

func (v NullableIdentityData) Get() *IdentityData {
	return v.value
}

func (v *NullableIdentityData) Set(val *IdentityData) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityData) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityData(val *IdentityData) *NullableIdentityData {
	return &NullableIdentityData{value: val, isSet: true}
}

func (v NullableIdentityData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


