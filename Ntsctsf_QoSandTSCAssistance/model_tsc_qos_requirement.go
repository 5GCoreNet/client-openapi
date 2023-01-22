/*
Ntsctsf_QoSandTSCAssistance Service API

TSCTSF QoS and TSC Assistance Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ntsctsf_QoSandTSCAssistance

import (
	"encoding/json"
)

// TscQosRequirement Represents QoS requirements for time sensitive communication.
type TscQosRequirement struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	ReqGbrDl *string `json:"reqGbrDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	ReqGbrUl *string `json:"reqGbrUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	ReqMbrDl *string `json:"reqMbrDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	ReqMbrUl *string `json:"reqMbrUl,omitempty"`
	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.  
	MaxTscBurstSize *int32 `json:"maxTscBurstSize,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	Req5Gsdelay *int32 `json:"req5Gsdelay,omitempty"`
	// Represents the priority level of TSC Flows.
	Priority *int32 `json:"priority,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TscaiTimeDom *int32 `json:"tscaiTimeDom,omitempty"`
	TscaiInputDl NullableTscaiInputContainer `json:"tscaiInputDl,omitempty"`
	TscaiInputUl NullableTscaiInputContainer `json:"tscaiInputUl,omitempty"`
}

// NewTscQosRequirement instantiates a new TscQosRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTscQosRequirement() *TscQosRequirement {
	this := TscQosRequirement{}
	return &this
}

// NewTscQosRequirementWithDefaults instantiates a new TscQosRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTscQosRequirementWithDefaults() *TscQosRequirement {
	this := TscQosRequirement{}
	return &this
}

// GetReqGbrDl returns the ReqGbrDl field value if set, zero value otherwise.
func (o *TscQosRequirement) GetReqGbrDl() string {
	if o == nil || isNil(o.ReqGbrDl) {
		var ret string
		return ret
	}
	return *o.ReqGbrDl
}

// GetReqGbrDlOk returns a tuple with the ReqGbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscQosRequirement) GetReqGbrDlOk() (*string, bool) {
	if o == nil || isNil(o.ReqGbrDl) {
    return nil, false
	}
	return o.ReqGbrDl, true
}

// HasReqGbrDl returns a boolean if a field has been set.
func (o *TscQosRequirement) HasReqGbrDl() bool {
	if o != nil && !isNil(o.ReqGbrDl) {
		return true
	}

	return false
}

// SetReqGbrDl gets a reference to the given string and assigns it to the ReqGbrDl field.
func (o *TscQosRequirement) SetReqGbrDl(v string) {
	o.ReqGbrDl = &v
}

// GetReqGbrUl returns the ReqGbrUl field value if set, zero value otherwise.
func (o *TscQosRequirement) GetReqGbrUl() string {
	if o == nil || isNil(o.ReqGbrUl) {
		var ret string
		return ret
	}
	return *o.ReqGbrUl
}

// GetReqGbrUlOk returns a tuple with the ReqGbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscQosRequirement) GetReqGbrUlOk() (*string, bool) {
	if o == nil || isNil(o.ReqGbrUl) {
    return nil, false
	}
	return o.ReqGbrUl, true
}

// HasReqGbrUl returns a boolean if a field has been set.
func (o *TscQosRequirement) HasReqGbrUl() bool {
	if o != nil && !isNil(o.ReqGbrUl) {
		return true
	}

	return false
}

// SetReqGbrUl gets a reference to the given string and assigns it to the ReqGbrUl field.
func (o *TscQosRequirement) SetReqGbrUl(v string) {
	o.ReqGbrUl = &v
}

// GetReqMbrDl returns the ReqMbrDl field value if set, zero value otherwise.
func (o *TscQosRequirement) GetReqMbrDl() string {
	if o == nil || isNil(o.ReqMbrDl) {
		var ret string
		return ret
	}
	return *o.ReqMbrDl
}

// GetReqMbrDlOk returns a tuple with the ReqMbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscQosRequirement) GetReqMbrDlOk() (*string, bool) {
	if o == nil || isNil(o.ReqMbrDl) {
    return nil, false
	}
	return o.ReqMbrDl, true
}

// HasReqMbrDl returns a boolean if a field has been set.
func (o *TscQosRequirement) HasReqMbrDl() bool {
	if o != nil && !isNil(o.ReqMbrDl) {
		return true
	}

	return false
}

// SetReqMbrDl gets a reference to the given string and assigns it to the ReqMbrDl field.
func (o *TscQosRequirement) SetReqMbrDl(v string) {
	o.ReqMbrDl = &v
}

// GetReqMbrUl returns the ReqMbrUl field value if set, zero value otherwise.
func (o *TscQosRequirement) GetReqMbrUl() string {
	if o == nil || isNil(o.ReqMbrUl) {
		var ret string
		return ret
	}
	return *o.ReqMbrUl
}

// GetReqMbrUlOk returns a tuple with the ReqMbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscQosRequirement) GetReqMbrUlOk() (*string, bool) {
	if o == nil || isNil(o.ReqMbrUl) {
    return nil, false
	}
	return o.ReqMbrUl, true
}

// HasReqMbrUl returns a boolean if a field has been set.
func (o *TscQosRequirement) HasReqMbrUl() bool {
	if o != nil && !isNil(o.ReqMbrUl) {
		return true
	}

	return false
}

// SetReqMbrUl gets a reference to the given string and assigns it to the ReqMbrUl field.
func (o *TscQosRequirement) SetReqMbrUl(v string) {
	o.ReqMbrUl = &v
}

// GetMaxTscBurstSize returns the MaxTscBurstSize field value if set, zero value otherwise.
func (o *TscQosRequirement) GetMaxTscBurstSize() int32 {
	if o == nil || isNil(o.MaxTscBurstSize) {
		var ret int32
		return ret
	}
	return *o.MaxTscBurstSize
}

// GetMaxTscBurstSizeOk returns a tuple with the MaxTscBurstSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscQosRequirement) GetMaxTscBurstSizeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTscBurstSize) {
    return nil, false
	}
	return o.MaxTscBurstSize, true
}

// HasMaxTscBurstSize returns a boolean if a field has been set.
func (o *TscQosRequirement) HasMaxTscBurstSize() bool {
	if o != nil && !isNil(o.MaxTscBurstSize) {
		return true
	}

	return false
}

// SetMaxTscBurstSize gets a reference to the given int32 and assigns it to the MaxTscBurstSize field.
func (o *TscQosRequirement) SetMaxTscBurstSize(v int32) {
	o.MaxTscBurstSize = &v
}

// GetReq5Gsdelay returns the Req5Gsdelay field value if set, zero value otherwise.
func (o *TscQosRequirement) GetReq5Gsdelay() int32 {
	if o == nil || isNil(o.Req5Gsdelay) {
		var ret int32
		return ret
	}
	return *o.Req5Gsdelay
}

// GetReq5GsdelayOk returns a tuple with the Req5Gsdelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscQosRequirement) GetReq5GsdelayOk() (*int32, bool) {
	if o == nil || isNil(o.Req5Gsdelay) {
    return nil, false
	}
	return o.Req5Gsdelay, true
}

// HasReq5Gsdelay returns a boolean if a field has been set.
func (o *TscQosRequirement) HasReq5Gsdelay() bool {
	if o != nil && !isNil(o.Req5Gsdelay) {
		return true
	}

	return false
}

// SetReq5Gsdelay gets a reference to the given int32 and assigns it to the Req5Gsdelay field.
func (o *TscQosRequirement) SetReq5Gsdelay(v int32) {
	o.Req5Gsdelay = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *TscQosRequirement) GetPriority() int32 {
	if o == nil || isNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscQosRequirement) GetPriorityOk() (*int32, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *TscQosRequirement) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *TscQosRequirement) SetPriority(v int32) {
	o.Priority = &v
}

// GetTscaiTimeDom returns the TscaiTimeDom field value if set, zero value otherwise.
func (o *TscQosRequirement) GetTscaiTimeDom() int32 {
	if o == nil || isNil(o.TscaiTimeDom) {
		var ret int32
		return ret
	}
	return *o.TscaiTimeDom
}

// GetTscaiTimeDomOk returns a tuple with the TscaiTimeDom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TscQosRequirement) GetTscaiTimeDomOk() (*int32, bool) {
	if o == nil || isNil(o.TscaiTimeDom) {
    return nil, false
	}
	return o.TscaiTimeDom, true
}

// HasTscaiTimeDom returns a boolean if a field has been set.
func (o *TscQosRequirement) HasTscaiTimeDom() bool {
	if o != nil && !isNil(o.TscaiTimeDom) {
		return true
	}

	return false
}

// SetTscaiTimeDom gets a reference to the given int32 and assigns it to the TscaiTimeDom field.
func (o *TscQosRequirement) SetTscaiTimeDom(v int32) {
	o.TscaiTimeDom = &v
}

// GetTscaiInputDl returns the TscaiInputDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TscQosRequirement) GetTscaiInputDl() TscaiInputContainer {
	if o == nil || isNil(o.TscaiInputDl.Get()) {
		var ret TscaiInputContainer
		return ret
	}
	return *o.TscaiInputDl.Get()
}

// GetTscaiInputDlOk returns a tuple with the TscaiInputDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TscQosRequirement) GetTscaiInputDlOk() (*TscaiInputContainer, bool) {
	if o == nil {
    return nil, false
	}
	return o.TscaiInputDl.Get(), o.TscaiInputDl.IsSet()
}

// HasTscaiInputDl returns a boolean if a field has been set.
func (o *TscQosRequirement) HasTscaiInputDl() bool {
	if o != nil && o.TscaiInputDl.IsSet() {
		return true
	}

	return false
}

// SetTscaiInputDl gets a reference to the given NullableTscaiInputContainer and assigns it to the TscaiInputDl field.
func (o *TscQosRequirement) SetTscaiInputDl(v TscaiInputContainer) {
	o.TscaiInputDl.Set(&v)
}
// SetTscaiInputDlNil sets the value for TscaiInputDl to be an explicit nil
func (o *TscQosRequirement) SetTscaiInputDlNil() {
	o.TscaiInputDl.Set(nil)
}

// UnsetTscaiInputDl ensures that no value is present for TscaiInputDl, not even an explicit nil
func (o *TscQosRequirement) UnsetTscaiInputDl() {
	o.TscaiInputDl.Unset()
}

// GetTscaiInputUl returns the TscaiInputUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TscQosRequirement) GetTscaiInputUl() TscaiInputContainer {
	if o == nil || isNil(o.TscaiInputUl.Get()) {
		var ret TscaiInputContainer
		return ret
	}
	return *o.TscaiInputUl.Get()
}

// GetTscaiInputUlOk returns a tuple with the TscaiInputUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TscQosRequirement) GetTscaiInputUlOk() (*TscaiInputContainer, bool) {
	if o == nil {
    return nil, false
	}
	return o.TscaiInputUl.Get(), o.TscaiInputUl.IsSet()
}

// HasTscaiInputUl returns a boolean if a field has been set.
func (o *TscQosRequirement) HasTscaiInputUl() bool {
	if o != nil && o.TscaiInputUl.IsSet() {
		return true
	}

	return false
}

// SetTscaiInputUl gets a reference to the given NullableTscaiInputContainer and assigns it to the TscaiInputUl field.
func (o *TscQosRequirement) SetTscaiInputUl(v TscaiInputContainer) {
	o.TscaiInputUl.Set(&v)
}
// SetTscaiInputUlNil sets the value for TscaiInputUl to be an explicit nil
func (o *TscQosRequirement) SetTscaiInputUlNil() {
	o.TscaiInputUl.Set(nil)
}

// UnsetTscaiInputUl ensures that no value is present for TscaiInputUl, not even an explicit nil
func (o *TscQosRequirement) UnsetTscaiInputUl() {
	o.TscaiInputUl.Unset()
}

func (o TscQosRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReqGbrDl) {
		toSerialize["reqGbrDl"] = o.ReqGbrDl
	}
	if !isNil(o.ReqGbrUl) {
		toSerialize["reqGbrUl"] = o.ReqGbrUl
	}
	if !isNil(o.ReqMbrDl) {
		toSerialize["reqMbrDl"] = o.ReqMbrDl
	}
	if !isNil(o.ReqMbrUl) {
		toSerialize["reqMbrUl"] = o.ReqMbrUl
	}
	if !isNil(o.MaxTscBurstSize) {
		toSerialize["maxTscBurstSize"] = o.MaxTscBurstSize
	}
	if !isNil(o.Req5Gsdelay) {
		toSerialize["req5Gsdelay"] = o.Req5Gsdelay
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !isNil(o.TscaiTimeDom) {
		toSerialize["tscaiTimeDom"] = o.TscaiTimeDom
	}
	if o.TscaiInputDl.IsSet() {
		toSerialize["tscaiInputDl"] = o.TscaiInputDl.Get()
	}
	if o.TscaiInputUl.IsSet() {
		toSerialize["tscaiInputUl"] = o.TscaiInputUl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTscQosRequirement struct {
	value *TscQosRequirement
	isSet bool
}

func (v NullableTscQosRequirement) Get() *TscQosRequirement {
	return v.value
}

func (v *NullableTscQosRequirement) Set(val *TscQosRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableTscQosRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableTscQosRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTscQosRequirement(val *TscQosRequirement) *NullableTscQosRequirement {
	return &NullableTscQosRequirement{value: val, isSet: true}
}

func (v NullableTscQosRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTscQosRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


