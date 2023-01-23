/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
)

// Arp Contains Allocation and Retention Priority information.
type Arp struct {
	// nullable true shall not be used for this attribute. Unsigned integer indicating the ARP Priority Level (see clause 5.7.2.2 of 3GPP TS 23.501, within the range 1 to 15.Values are ordered in decreasing order of priority, i.e. with 1 as the highest priority and 15 as the lowest priority.  
	PriorityLevel NullableInt32 `json:"priorityLevel"`
	PreemptCap PreemptionCapability `json:"preemptCap"`
	PreemptVuln PreemptionVulnerability `json:"preemptVuln"`
}

// NewArp instantiates a new Arp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArp(priorityLevel NullableInt32, preemptCap PreemptionCapability, preemptVuln PreemptionVulnerability) *Arp {
	this := Arp{}
	this.PriorityLevel = priorityLevel
	this.PreemptCap = preemptCap
	this.PreemptVuln = preemptVuln
	return &this
}

// NewArpWithDefaults instantiates a new Arp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArpWithDefaults() *Arp {
	this := Arp{}
	return &this
}

// GetPriorityLevel returns the PriorityLevel field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Arp) GetPriorityLevel() int32 {
	if o == nil || o.PriorityLevel.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PriorityLevel.Get()
}

// GetPriorityLevelOk returns a tuple with the PriorityLevel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Arp) GetPriorityLevelOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.PriorityLevel.Get(), o.PriorityLevel.IsSet()
}

// SetPriorityLevel sets field value
func (o *Arp) SetPriorityLevel(v int32) {
	o.PriorityLevel.Set(&v)
}

// GetPreemptCap returns the PreemptCap field value
func (o *Arp) GetPreemptCap() PreemptionCapability {
	if o == nil {
		var ret PreemptionCapability
		return ret
	}

	return o.PreemptCap
}

// GetPreemptCapOk returns a tuple with the PreemptCap field value
// and a boolean to check if the value has been set.
func (o *Arp) GetPreemptCapOk() (*PreemptionCapability, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreemptCap, true
}

// SetPreemptCap sets field value
func (o *Arp) SetPreemptCap(v PreemptionCapability) {
	o.PreemptCap = v
}

// GetPreemptVuln returns the PreemptVuln field value
func (o *Arp) GetPreemptVuln() PreemptionVulnerability {
	if o == nil {
		var ret PreemptionVulnerability
		return ret
	}

	return o.PreemptVuln
}

// GetPreemptVulnOk returns a tuple with the PreemptVuln field value
// and a boolean to check if the value has been set.
func (o *Arp) GetPreemptVulnOk() (*PreemptionVulnerability, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PreemptVuln, true
}

// SetPreemptVuln sets field value
func (o *Arp) SetPreemptVuln(v PreemptionVulnerability) {
	o.PreemptVuln = v
}

func (o Arp) MarshalJSON() ([]byte, error) {
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

type NullableArp struct {
	value *Arp
	isSet bool
}

func (v NullableArp) Get() *Arp {
	return v.value
}

func (v *NullableArp) Set(val *Arp) {
	v.value = val
	v.isSet = true
}

func (v NullableArp) IsSet() bool {
	return v.isSet
}

func (v *NullableArp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArp(val *Arp) *NullableArp {
	return &NullableArp{value: val, isSet: true}
}

func (v NullableArp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


