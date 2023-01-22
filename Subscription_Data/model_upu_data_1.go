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

// UpuData1 Contains UE parameters update data set (e.g., the updated Routing ID Data or the Default configured NSSAI).
type UpuData1 struct {
	// Contains a secure packet.
	SecPacket *string `json:"secPacket,omitempty"`
	DefaultConfNssai []Snssai `json:"defaultConfNssai,omitempty"`
	// Represents a routing indicator.
	RoutingId *string `json:"routingId,omitempty"`
}

// NewUpuData1 instantiates a new UpuData1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpuData1() *UpuData1 {
	this := UpuData1{}
	return &this
}

// NewUpuData1WithDefaults instantiates a new UpuData1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpuData1WithDefaults() *UpuData1 {
	this := UpuData1{}
	return &this
}

// GetSecPacket returns the SecPacket field value if set, zero value otherwise.
func (o *UpuData1) GetSecPacket() string {
	if o == nil || isNil(o.SecPacket) {
		var ret string
		return ret
	}
	return *o.SecPacket
}

// GetSecPacketOk returns a tuple with the SecPacket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuData1) GetSecPacketOk() (*string, bool) {
	if o == nil || isNil(o.SecPacket) {
    return nil, false
	}
	return o.SecPacket, true
}

// HasSecPacket returns a boolean if a field has been set.
func (o *UpuData1) HasSecPacket() bool {
	if o != nil && !isNil(o.SecPacket) {
		return true
	}

	return false
}

// SetSecPacket gets a reference to the given string and assigns it to the SecPacket field.
func (o *UpuData1) SetSecPacket(v string) {
	o.SecPacket = &v
}

// GetDefaultConfNssai returns the DefaultConfNssai field value if set, zero value otherwise.
func (o *UpuData1) GetDefaultConfNssai() []Snssai {
	if o == nil || isNil(o.DefaultConfNssai) {
		var ret []Snssai
		return ret
	}
	return o.DefaultConfNssai
}

// GetDefaultConfNssaiOk returns a tuple with the DefaultConfNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuData1) GetDefaultConfNssaiOk() ([]Snssai, bool) {
	if o == nil || isNil(o.DefaultConfNssai) {
    return nil, false
	}
	return o.DefaultConfNssai, true
}

// HasDefaultConfNssai returns a boolean if a field has been set.
func (o *UpuData1) HasDefaultConfNssai() bool {
	if o != nil && !isNil(o.DefaultConfNssai) {
		return true
	}

	return false
}

// SetDefaultConfNssai gets a reference to the given []Snssai and assigns it to the DefaultConfNssai field.
func (o *UpuData1) SetDefaultConfNssai(v []Snssai) {
	o.DefaultConfNssai = v
}

// GetRoutingId returns the RoutingId field value if set, zero value otherwise.
func (o *UpuData1) GetRoutingId() string {
	if o == nil || isNil(o.RoutingId) {
		var ret string
		return ret
	}
	return *o.RoutingId
}

// GetRoutingIdOk returns a tuple with the RoutingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpuData1) GetRoutingIdOk() (*string, bool) {
	if o == nil || isNil(o.RoutingId) {
    return nil, false
	}
	return o.RoutingId, true
}

// HasRoutingId returns a boolean if a field has been set.
func (o *UpuData1) HasRoutingId() bool {
	if o != nil && !isNil(o.RoutingId) {
		return true
	}

	return false
}

// SetRoutingId gets a reference to the given string and assigns it to the RoutingId field.
func (o *UpuData1) SetRoutingId(v string) {
	o.RoutingId = &v
}

func (o UpuData1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SecPacket) {
		toSerialize["secPacket"] = o.SecPacket
	}
	if !isNil(o.DefaultConfNssai) {
		toSerialize["defaultConfNssai"] = o.DefaultConfNssai
	}
	if !isNil(o.RoutingId) {
		toSerialize["routingId"] = o.RoutingId
	}
	return json.Marshal(toSerialize)
}

type NullableUpuData1 struct {
	value *UpuData1
	isSet bool
}

func (v NullableUpuData1) Get() *UpuData1 {
	return v.value
}

func (v *NullableUpuData1) Set(val *UpuData1) {
	v.value = val
	v.isSet = true
}

func (v NullableUpuData1) IsSet() bool {
	return v.isSet
}

func (v *NullableUpuData1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpuData1(val *UpuData1) *NullableUpuData1 {
	return &NullableUpuData1{value: val, isSet: true}
}

func (v NullableUpuData1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpuData1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


