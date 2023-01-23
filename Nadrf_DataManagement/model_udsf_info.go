/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// UdsfInfo Information related to UDSF
type UdsfInfo struct {
	// Identifier of a group of NFs.
	GroupId *string `json:"groupId,omitempty"`
	SupiRanges []SupiRange `json:"supiRanges,omitempty"`
	// A map (list of key-value pairs) where realmId serves as key and each value in the map is an array of IdentityRanges. Each IdentityRange is a range of storageIds. 
	StorageIdRanges *map[string][]IdentityRange `json:"storageIdRanges,omitempty"`
}

// NewUdsfInfo instantiates a new UdsfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUdsfInfo() *UdsfInfo {
	this := UdsfInfo{}
	return &this
}

// NewUdsfInfoWithDefaults instantiates a new UdsfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUdsfInfoWithDefaults() *UdsfInfo {
	this := UdsfInfo{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *UdsfInfo) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdsfInfo) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *UdsfInfo) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *UdsfInfo) SetGroupId(v string) {
	o.GroupId = &v
}

// GetSupiRanges returns the SupiRanges field value if set, zero value otherwise.
func (o *UdsfInfo) GetSupiRanges() []SupiRange {
	if o == nil || isNil(o.SupiRanges) {
		var ret []SupiRange
		return ret
	}
	return o.SupiRanges
}

// GetSupiRangesOk returns a tuple with the SupiRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdsfInfo) GetSupiRangesOk() ([]SupiRange, bool) {
	if o == nil || isNil(o.SupiRanges) {
    return nil, false
	}
	return o.SupiRanges, true
}

// HasSupiRanges returns a boolean if a field has been set.
func (o *UdsfInfo) HasSupiRanges() bool {
	if o != nil && !isNil(o.SupiRanges) {
		return true
	}

	return false
}

// SetSupiRanges gets a reference to the given []SupiRange and assigns it to the SupiRanges field.
func (o *UdsfInfo) SetSupiRanges(v []SupiRange) {
	o.SupiRanges = v
}

// GetStorageIdRanges returns the StorageIdRanges field value if set, zero value otherwise.
func (o *UdsfInfo) GetStorageIdRanges() map[string][]IdentityRange {
	if o == nil || isNil(o.StorageIdRanges) {
		var ret map[string][]IdentityRange
		return ret
	}
	return *o.StorageIdRanges
}

// GetStorageIdRangesOk returns a tuple with the StorageIdRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UdsfInfo) GetStorageIdRangesOk() (*map[string][]IdentityRange, bool) {
	if o == nil || isNil(o.StorageIdRanges) {
    return nil, false
	}
	return o.StorageIdRanges, true
}

// HasStorageIdRanges returns a boolean if a field has been set.
func (o *UdsfInfo) HasStorageIdRanges() bool {
	if o != nil && !isNil(o.StorageIdRanges) {
		return true
	}

	return false
}

// SetStorageIdRanges gets a reference to the given map[string][]IdentityRange and assigns it to the StorageIdRanges field.
func (o *UdsfInfo) SetStorageIdRanges(v map[string][]IdentityRange) {
	o.StorageIdRanges = &v
}

func (o UdsfInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.SupiRanges) {
		toSerialize["supiRanges"] = o.SupiRanges
	}
	if !isNil(o.StorageIdRanges) {
		toSerialize["storageIdRanges"] = o.StorageIdRanges
	}
	return json.Marshal(toSerialize)
}

type NullableUdsfInfo struct {
	value *UdsfInfo
	isSet bool
}

func (v NullableUdsfInfo) Get() *UdsfInfo {
	return v.value
}

func (v *NullableUdsfInfo) Set(val *UdsfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUdsfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUdsfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUdsfInfo(val *UdsfInfo) *NullableUdsfInfo {
	return &NullableUdsfInfo{value: val, isSet: true}
}

func (v NullableUdsfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUdsfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


