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

// SuciInfo SUCI information containing Routing Indicator and Home Network Public Key ID
type SuciInfo struct {
	RoutingInds []string `json:"routingInds,omitempty"`
	HNwPubKeyIds []int32 `json:"hNwPubKeyIds,omitempty"`
}

// NewSuciInfo instantiates a new SuciInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuciInfo() *SuciInfo {
	this := SuciInfo{}
	return &this
}

// NewSuciInfoWithDefaults instantiates a new SuciInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuciInfoWithDefaults() *SuciInfo {
	this := SuciInfo{}
	return &this
}

// GetRoutingInds returns the RoutingInds field value if set, zero value otherwise.
func (o *SuciInfo) GetRoutingInds() []string {
	if o == nil || isNil(o.RoutingInds) {
		var ret []string
		return ret
	}
	return o.RoutingInds
}

// GetRoutingIndsOk returns a tuple with the RoutingInds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuciInfo) GetRoutingIndsOk() ([]string, bool) {
	if o == nil || isNil(o.RoutingInds) {
    return nil, false
	}
	return o.RoutingInds, true
}

// HasRoutingInds returns a boolean if a field has been set.
func (o *SuciInfo) HasRoutingInds() bool {
	if o != nil && !isNil(o.RoutingInds) {
		return true
	}

	return false
}

// SetRoutingInds gets a reference to the given []string and assigns it to the RoutingInds field.
func (o *SuciInfo) SetRoutingInds(v []string) {
	o.RoutingInds = v
}

// GetHNwPubKeyIds returns the HNwPubKeyIds field value if set, zero value otherwise.
func (o *SuciInfo) GetHNwPubKeyIds() []int32 {
	if o == nil || isNil(o.HNwPubKeyIds) {
		var ret []int32
		return ret
	}
	return o.HNwPubKeyIds
}

// GetHNwPubKeyIdsOk returns a tuple with the HNwPubKeyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuciInfo) GetHNwPubKeyIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.HNwPubKeyIds) {
    return nil, false
	}
	return o.HNwPubKeyIds, true
}

// HasHNwPubKeyIds returns a boolean if a field has been set.
func (o *SuciInfo) HasHNwPubKeyIds() bool {
	if o != nil && !isNil(o.HNwPubKeyIds) {
		return true
	}

	return false
}

// SetHNwPubKeyIds gets a reference to the given []int32 and assigns it to the HNwPubKeyIds field.
func (o *SuciInfo) SetHNwPubKeyIds(v []int32) {
	o.HNwPubKeyIds = v
}

func (o SuciInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RoutingInds) {
		toSerialize["routingInds"] = o.RoutingInds
	}
	if !isNil(o.HNwPubKeyIds) {
		toSerialize["hNwPubKeyIds"] = o.HNwPubKeyIds
	}
	return json.Marshal(toSerialize)
}

type NullableSuciInfo struct {
	value *SuciInfo
	isSet bool
}

func (v NullableSuciInfo) Get() *SuciInfo {
	return v.value
}

func (v *NullableSuciInfo) Set(val *SuciInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSuciInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSuciInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuciInfo(val *SuciInfo) *NullableSuciInfo {
	return &NullableSuciInfo{value: val, isSet: true}
}

func (v NullableSuciInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuciInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


