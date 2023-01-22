/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_SDM

import (
	"encoding/json"
)

// SubscribedDefaultQos Provides the subsribed 5QI and the ARP, it may contain the priority level.
type SubscribedDefaultQos struct {
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255. 
	Var5qi int32 `json:"5qi"`
	Arp Arp `json:"arp"`
	// Unsigned integer indicating the 5QI Priority Level (see clauses 5.7.3.3 and 5.7.4 of 3GPP TS 23.501, within the range 1 to 127.Values are ordered in decreasing order of priority,  i.e. with 1 as the highest priority and 127 as the lowest priority.  
	PriorityLevel *int32 `json:"priorityLevel,omitempty"`
}

// NewSubscribedDefaultQos instantiates a new SubscribedDefaultQos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribedDefaultQos(var5qi int32, arp Arp) *SubscribedDefaultQos {
	this := SubscribedDefaultQos{}
	this.Var5qi = var5qi
	this.Arp = arp
	return &this
}

// NewSubscribedDefaultQosWithDefaults instantiates a new SubscribedDefaultQos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribedDefaultQosWithDefaults() *SubscribedDefaultQos {
	this := SubscribedDefaultQos{}
	return &this
}

// GetVar5qi returns the Var5qi field value
func (o *SubscribedDefaultQos) GetVar5qi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var5qi
}

// GetVar5qiOk returns a tuple with the Var5qi field value
// and a boolean to check if the value has been set.
func (o *SubscribedDefaultQos) GetVar5qiOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Var5qi, true
}

// SetVar5qi sets field value
func (o *SubscribedDefaultQos) SetVar5qi(v int32) {
	o.Var5qi = v
}

// GetArp returns the Arp field value
func (o *SubscribedDefaultQos) GetArp() Arp {
	if o == nil {
		var ret Arp
		return ret
	}

	return o.Arp
}

// GetArpOk returns a tuple with the Arp field value
// and a boolean to check if the value has been set.
func (o *SubscribedDefaultQos) GetArpOk() (*Arp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Arp, true
}

// SetArp sets field value
func (o *SubscribedDefaultQos) SetArp(v Arp) {
	o.Arp = v
}

// GetPriorityLevel returns the PriorityLevel field value if set, zero value otherwise.
func (o *SubscribedDefaultQos) GetPriorityLevel() int32 {
	if o == nil || isNil(o.PriorityLevel) {
		var ret int32
		return ret
	}
	return *o.PriorityLevel
}

// GetPriorityLevelOk returns a tuple with the PriorityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribedDefaultQos) GetPriorityLevelOk() (*int32, bool) {
	if o == nil || isNil(o.PriorityLevel) {
    return nil, false
	}
	return o.PriorityLevel, true
}

// HasPriorityLevel returns a boolean if a field has been set.
func (o *SubscribedDefaultQos) HasPriorityLevel() bool {
	if o != nil && !isNil(o.PriorityLevel) {
		return true
	}

	return false
}

// SetPriorityLevel gets a reference to the given int32 and assigns it to the PriorityLevel field.
func (o *SubscribedDefaultQos) SetPriorityLevel(v int32) {
	o.PriorityLevel = &v
}

func (o SubscribedDefaultQos) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["5qi"] = o.Var5qi
	}
	if true {
		toSerialize["arp"] = o.Arp
	}
	if !isNil(o.PriorityLevel) {
		toSerialize["priorityLevel"] = o.PriorityLevel
	}
	return json.Marshal(toSerialize)
}

type NullableSubscribedDefaultQos struct {
	value *SubscribedDefaultQos
	isSet bool
}

func (v NullableSubscribedDefaultQos) Get() *SubscribedDefaultQos {
	return v.value
}

func (v *NullableSubscribedDefaultQos) Set(val *SubscribedDefaultQos) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribedDefaultQos) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribedDefaultQos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribedDefaultQos(val *SubscribedDefaultQos) *NullableSubscribedDefaultQos {
	return &NullableSubscribedDefaultQos{value: val, isSet: true}
}

func (v NullableSubscribedDefaultQos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribedDefaultQos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


