/*
Nhss_gbaSDM

Nhss Subscriber Data Management Service for GBA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_gbaSDM

import (
	"encoding/json"
)

// UeIdsItem Data item in a UE ID array list
type UeIdsItem struct {
	// Public Identity of the UE
	UeId string `json:"ueId"`
}

// NewUeIdsItem instantiates a new UeIdsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeIdsItem(ueId string) *UeIdsItem {
	this := UeIdsItem{}
	this.UeId = ueId
	return &this
}

// NewUeIdsItemWithDefaults instantiates a new UeIdsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeIdsItemWithDefaults() *UeIdsItem {
	this := UeIdsItem{}
	return &this
}

// GetUeId returns the UeId field value
func (o *UeIdsItem) GetUeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value
// and a boolean to check if the value has been set.
func (o *UeIdsItem) GetUeIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UeId, true
}

// SetUeId sets field value
func (o *UeIdsItem) SetUeId(v string) {
	o.UeId = v
}

func (o UeIdsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ueId"] = o.UeId
	}
	return json.Marshal(toSerialize)
}

type NullableUeIdsItem struct {
	value *UeIdsItem
	isSet bool
}

func (v NullableUeIdsItem) Get() *UeIdsItem {
	return v.value
}

func (v *NullableUeIdsItem) Set(val *UeIdsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUeIdsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUeIdsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeIdsItem(val *UeIdsItem) *NullableUeIdsItem {
	return &NullableUeIdsItem{value: val, isSet: true}
}

func (v NullableUeIdsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeIdsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


