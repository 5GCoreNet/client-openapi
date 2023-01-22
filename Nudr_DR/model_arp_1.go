/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// Arp1 Contains Allocation and Retention Priority information.
type Arp1 struct {
	// nullable true shall not be used for this attribute. Unsigned integer indicating the ARP Priority Level (see clause 5.7.2.2 of 3GPP TS 23.501, within the range 1 to 15.Values are ordered in decreasing order of priority, i.e. with 1 as the highest priority and 15 as the lowest priority.  
	PriorityLevel NullableInt32 `json:"priorityLevel"`
	PreemptCap PreemptionCapability `json:"preemptCap"`
	PreemptVuln PreemptionVulnerability `json:"preemptVuln"`
}

// NewArp1 instantiates a new Arp1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArp1(priorityLevel NullableInt32, preemptCap PreemptionCapability, preemptVuln PreemptionVulnerability) *Arp1 {
	this := Arp1{}
	this.PriorityLevel = priorityLevel
	this.PreemptCap = preemptCap
	this.PreemptVuln = preemptVuln
	return &this
}

// NewArp1WithDefaults instantiates a new Arp1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArp1WithDefaults() *Arp1 {
	this := Arp1{}
	return &this
}

// GetPriorityLevel returns the PriorityLevel field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Arp1) GetPriorityLevel() int32 {
	if o == nil || o.PriorityLevel.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PriorityLevel.Get()
}

// GetPriorityLevelOk returns a tuple with the PriorityLevel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Arp1) GetPriorityLevelOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.PriorityLevel.Get(), o.PriorityLevel.IsSet()
}

// SetPriorityLevel sets field value
func (o *Arp1) SetPriorityLevel(v int32) {
	o.PriorityLevel.Set(&v)
}

// GetPreemptCap returns the PreemptCap field value
func (o *Arp1) GetPreemptCap() PreemptionCapability {
	if o == nil {
		var ret PreemptionCapability
		return ret
	}

	return o.PreemptCap
}

// GetPreemptCapOk returns a tuple with the PreemptCap field value
// and a boolean to check if the value has been set.
func (o *Arp1) GetPreemptCapOk() (*PreemptionCapability, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreemptCap, true
}

// SetPreemptCap sets field value
func (o *Arp1) SetPreemptCap(v PreemptionCapability) {
	o.PreemptCap = v
}

// GetPreemptVuln returns the PreemptVuln field value
func (o *Arp1) GetPreemptVuln() PreemptionVulnerability {
	if o == nil {
		var ret PreemptionVulnerability
		return ret
	}

	return o.PreemptVuln
}

// GetPreemptVulnOk returns a tuple with the PreemptVuln field value
// and a boolean to check if the value has been set.
func (o *Arp1) GetPreemptVulnOk() (*PreemptionVulnerability, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreemptVuln, true
}

// SetPreemptVuln sets field value
func (o *Arp1) SetPreemptVuln(v PreemptionVulnerability) {
	o.PreemptVuln = v
}

func (o Arp1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["priorityLevel"] = o.PriorityLevel.Get()
	}
	if true {
		toSerialize["preemptCap"] = o.PreemptCap
	}
	if true {
		toSerialize["preemptVuln"] = o.PreemptVuln
	}
	return json.Marshal(toSerialize)
}

type NullableArp1 struct {
	value *Arp1
	isSet bool
}

func (v NullableArp1) Get() *Arp1 {
	return v.value
}

func (v *NullableArp1) Set(val *Arp1) {
	v.value = val
	v.isSet = true
}

func (v NullableArp1) IsSet() bool {
	return v.isSet
}

func (v *NullableArp1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArp1(val *Arp1) *NullableArp1 {
	return &NullableArp1{value: val, isSet: true}
}

func (v NullableArp1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArp1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


