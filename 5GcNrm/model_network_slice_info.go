/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// NetworkSliceInfo struct for NetworkSliceInfo
type NetworkSliceInfo struct {
	SNSSAI *Snssai `json:"sNSSAI,omitempty"`
	// CNSI Id is defined in TS 29.531, only for Core Network
	CNSIId *string `json:"cNSIId,omitempty"`
	NetworkSliceRef []string `json:"networkSliceRef,omitempty"`
}

// NewNetworkSliceInfo instantiates a new NetworkSliceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSliceInfo() *NetworkSliceInfo {
	this := NetworkSliceInfo{}
	return &this
}

// NewNetworkSliceInfoWithDefaults instantiates a new NetworkSliceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSliceInfoWithDefaults() *NetworkSliceInfo {
	this := NetworkSliceInfo{}
	return &this
}

// GetSNSSAI returns the SNSSAI field value if set, zero value otherwise.
func (o *NetworkSliceInfo) GetSNSSAI() Snssai {
	if o == nil || isNil(o.SNSSAI) {
		var ret Snssai
		return ret
	}
	return *o.SNSSAI
}

// GetSNSSAIOk returns a tuple with the SNSSAI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceInfo) GetSNSSAIOk() (*Snssai, bool) {
	if o == nil || isNil(o.SNSSAI) {
    return nil, false
	}
	return o.SNSSAI, true
}

// HasSNSSAI returns a boolean if a field has been set.
func (o *NetworkSliceInfo) HasSNSSAI() bool {
	if o != nil && !isNil(o.SNSSAI) {
		return true
	}

	return false
}

// SetSNSSAI gets a reference to the given Snssai and assigns it to the SNSSAI field.
func (o *NetworkSliceInfo) SetSNSSAI(v Snssai) {
	o.SNSSAI = &v
}

// GetCNSIId returns the CNSIId field value if set, zero value otherwise.
func (o *NetworkSliceInfo) GetCNSIId() string {
	if o == nil || isNil(o.CNSIId) {
		var ret string
		return ret
	}
	return *o.CNSIId
}

// GetCNSIIdOk returns a tuple with the CNSIId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceInfo) GetCNSIIdOk() (*string, bool) {
	if o == nil || isNil(o.CNSIId) {
    return nil, false
	}
	return o.CNSIId, true
}

// HasCNSIId returns a boolean if a field has been set.
func (o *NetworkSliceInfo) HasCNSIId() bool {
	if o != nil && !isNil(o.CNSIId) {
		return true
	}

	return false
}

// SetCNSIId gets a reference to the given string and assigns it to the CNSIId field.
func (o *NetworkSliceInfo) SetCNSIId(v string) {
	o.CNSIId = &v
}

// GetNetworkSliceRef returns the NetworkSliceRef field value if set, zero value otherwise.
func (o *NetworkSliceInfo) GetNetworkSliceRef() []string {
	if o == nil || isNil(o.NetworkSliceRef) {
		var ret []string
		return ret
	}
	return o.NetworkSliceRef
}

// GetNetworkSliceRefOk returns a tuple with the NetworkSliceRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceInfo) GetNetworkSliceRefOk() ([]string, bool) {
	if o == nil || isNil(o.NetworkSliceRef) {
    return nil, false
	}
	return o.NetworkSliceRef, true
}

// HasNetworkSliceRef returns a boolean if a field has been set.
func (o *NetworkSliceInfo) HasNetworkSliceRef() bool {
	if o != nil && !isNil(o.NetworkSliceRef) {
		return true
	}

	return false
}

// SetNetworkSliceRef gets a reference to the given []string and assigns it to the NetworkSliceRef field.
func (o *NetworkSliceInfo) SetNetworkSliceRef(v []string) {
	o.NetworkSliceRef = v
}

func (o NetworkSliceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SNSSAI) {
		toSerialize["sNSSAI"] = o.SNSSAI
	}
	if !isNil(o.CNSIId) {
		toSerialize["cNSIId"] = o.CNSIId
	}
	if !isNil(o.NetworkSliceRef) {
		toSerialize["networkSliceRef"] = o.NetworkSliceRef
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkSliceInfo struct {
	value *NetworkSliceInfo
	isSet bool
}

func (v NullableNetworkSliceInfo) Get() *NetworkSliceInfo {
	return v.value
}

func (v *NullableNetworkSliceInfo) Set(val *NetworkSliceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSliceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSliceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSliceInfo(val *NetworkSliceInfo) *NullableNetworkSliceInfo {
	return &NullableNetworkSliceInfo{value: val, isSet: true}
}

func (v NullableNetworkSliceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSliceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


