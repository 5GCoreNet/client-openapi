/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
	"time"
)

// NssaiAckData Used to store the status of the latest NSSAI data update.
type NssaiAckData struct {
	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`
	UeUpdateStatus UeUpdateStatus `json:"ueUpdateStatus"`
}

// NewNssaiAckData instantiates a new NssaiAckData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNssaiAckData(provisioningTime time.Time, ueUpdateStatus UeUpdateStatus) *NssaiAckData {
	this := NssaiAckData{}
	this.ProvisioningTime = provisioningTime
	this.UeUpdateStatus = ueUpdateStatus
	return &this
}

// NewNssaiAckDataWithDefaults instantiates a new NssaiAckData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNssaiAckDataWithDefaults() *NssaiAckData {
	this := NssaiAckData{}
	return &this
}

// GetProvisioningTime returns the ProvisioningTime field value
func (o *NssaiAckData) GetProvisioningTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ProvisioningTime
}

// GetProvisioningTimeOk returns a tuple with the ProvisioningTime field value
// and a boolean to check if the value has been set.
func (o *NssaiAckData) GetProvisioningTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProvisioningTime, true
}

// SetProvisioningTime sets field value
func (o *NssaiAckData) SetProvisioningTime(v time.Time) {
	o.ProvisioningTime = v
}

// GetUeUpdateStatus returns the UeUpdateStatus field value
func (o *NssaiAckData) GetUeUpdateStatus() UeUpdateStatus {
	if o == nil {
		var ret UeUpdateStatus
		return ret
	}

	return o.UeUpdateStatus
}

// GetUeUpdateStatusOk returns a tuple with the UeUpdateStatus field value
// and a boolean to check if the value has been set.
func (o *NssaiAckData) GetUeUpdateStatusOk() (*UeUpdateStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UeUpdateStatus, true
}

// SetUeUpdateStatus sets field value
func (o *NssaiAckData) SetUeUpdateStatus(v UeUpdateStatus) {
	o.UeUpdateStatus = v
}

func (o NssaiAckData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["provisioningTime"] = o.ProvisioningTime
	}
	if true {
		toSerialize["ueUpdateStatus"] = o.UeUpdateStatus
	}
	return json.Marshal(toSerialize)
}

type NullableNssaiAckData struct {
	value *NssaiAckData
	isSet bool
}

func (v NullableNssaiAckData) Get() *NssaiAckData {
	return v.value
}

func (v *NullableNssaiAckData) Set(val *NssaiAckData) {
	v.value = val
	v.isSet = true
}

func (v NullableNssaiAckData) IsSet() bool {
	return v.isSet
}

func (v *NullableNssaiAckData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNssaiAckData(val *NssaiAckData) *NullableNssaiAckData {
	return &NullableNssaiAckData{value: val, isSet: true}
}

func (v NullableNssaiAckData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNssaiAckData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

