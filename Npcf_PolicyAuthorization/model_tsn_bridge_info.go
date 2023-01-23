/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// TsnBridgeInfo Contains parameters that describe and identify the TSC user plane node.
type TsnBridgeInfo struct {
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	BridgeId *int32 `json:"bridgeId,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	DsttAddr *string `json:"dsttAddr,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	DsttPortNum *int32 `json:"dsttPortNum,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	DsttResidTime *int32 `json:"dsttResidTime,omitempty"`
}

// NewTsnBridgeInfo instantiates a new TsnBridgeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTsnBridgeInfo() *TsnBridgeInfo {
	this := TsnBridgeInfo{}
	return &this
}

// NewTsnBridgeInfoWithDefaults instantiates a new TsnBridgeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTsnBridgeInfoWithDefaults() *TsnBridgeInfo {
	this := TsnBridgeInfo{}
	return &this
}

// GetBridgeId returns the BridgeId field value if set, zero value otherwise.
func (o *TsnBridgeInfo) GetBridgeId() int32 {
	if o == nil || isNil(o.BridgeId) {
		var ret int32
		return ret
	}
	return *o.BridgeId
}

// GetBridgeIdOk returns a tuple with the BridgeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnBridgeInfo) GetBridgeIdOk() (*int32, bool) {
	if o == nil || isNil(o.BridgeId) {
    return nil, false
	}
	return o.BridgeId, true
}

// HasBridgeId returns a boolean if a field has been set.
func (o *TsnBridgeInfo) HasBridgeId() bool {
	if o != nil && !isNil(o.BridgeId) {
		return true
	}

	return false
}

// SetBridgeId gets a reference to the given int32 and assigns it to the BridgeId field.
func (o *TsnBridgeInfo) SetBridgeId(v int32) {
	o.BridgeId = &v
}

// GetDsttAddr returns the DsttAddr field value if set, zero value otherwise.
func (o *TsnBridgeInfo) GetDsttAddr() string {
	if o == nil || isNil(o.DsttAddr) {
		var ret string
		return ret
	}
	return *o.DsttAddr
}

// GetDsttAddrOk returns a tuple with the DsttAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnBridgeInfo) GetDsttAddrOk() (*string, bool) {
	if o == nil || isNil(o.DsttAddr) {
    return nil, false
	}
	return o.DsttAddr, true
}

// HasDsttAddr returns a boolean if a field has been set.
func (o *TsnBridgeInfo) HasDsttAddr() bool {
	if o != nil && !isNil(o.DsttAddr) {
		return true
	}

	return false
}

// SetDsttAddr gets a reference to the given string and assigns it to the DsttAddr field.
func (o *TsnBridgeInfo) SetDsttAddr(v string) {
	o.DsttAddr = &v
}

// GetDsttPortNum returns the DsttPortNum field value if set, zero value otherwise.
func (o *TsnBridgeInfo) GetDsttPortNum() int32 {
	if o == nil || isNil(o.DsttPortNum) {
		var ret int32
		return ret
	}
	return *o.DsttPortNum
}

// GetDsttPortNumOk returns a tuple with the DsttPortNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnBridgeInfo) GetDsttPortNumOk() (*int32, bool) {
	if o == nil || isNil(o.DsttPortNum) {
    return nil, false
	}
	return o.DsttPortNum, true
}

// HasDsttPortNum returns a boolean if a field has been set.
func (o *TsnBridgeInfo) HasDsttPortNum() bool {
	if o != nil && !isNil(o.DsttPortNum) {
		return true
	}

	return false
}

// SetDsttPortNum gets a reference to the given int32 and assigns it to the DsttPortNum field.
func (o *TsnBridgeInfo) SetDsttPortNum(v int32) {
	o.DsttPortNum = &v
}

// GetDsttResidTime returns the DsttResidTime field value if set, zero value otherwise.
func (o *TsnBridgeInfo) GetDsttResidTime() int32 {
	if o == nil || isNil(o.DsttResidTime) {
		var ret int32
		return ret
	}
	return *o.DsttResidTime
}

// GetDsttResidTimeOk returns a tuple with the DsttResidTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnBridgeInfo) GetDsttResidTimeOk() (*int32, bool) {
	if o == nil || isNil(o.DsttResidTime) {
    return nil, false
	}
	return o.DsttResidTime, true
}

// HasDsttResidTime returns a boolean if a field has been set.
func (o *TsnBridgeInfo) HasDsttResidTime() bool {
	if o != nil && !isNil(o.DsttResidTime) {
		return true
	}

	return false
}

// SetDsttResidTime gets a reference to the given int32 and assigns it to the DsttResidTime field.
func (o *TsnBridgeInfo) SetDsttResidTime(v int32) {
	o.DsttResidTime = &v
}

func (o TsnBridgeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BridgeId) {
		toSerialize["bridgeId"] = o.BridgeId
	}
	if !isNil(o.DsttAddr) {
		toSerialize["dsttAddr"] = o.DsttAddr
	}
	if !isNil(o.DsttPortNum) {
		toSerialize["dsttPortNum"] = o.DsttPortNum
	}
	if !isNil(o.DsttResidTime) {
		toSerialize["dsttResidTime"] = o.DsttResidTime
	}
	return json.Marshal(toSerialize)
}

type NullableTsnBridgeInfo struct {
	value *TsnBridgeInfo
	isSet bool
}

func (v NullableTsnBridgeInfo) Get() *TsnBridgeInfo {
	return v.value
}

func (v *NullableTsnBridgeInfo) Set(val *TsnBridgeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTsnBridgeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTsnBridgeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTsnBridgeInfo(val *TsnBridgeInfo) *NullableTsnBridgeInfo {
	return &NullableTsnBridgeInfo{value: val, isSet: true}
}

func (v NullableTsnBridgeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTsnBridgeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


