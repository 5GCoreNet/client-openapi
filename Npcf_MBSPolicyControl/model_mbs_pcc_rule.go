/*
Npcf_MBSPolicyControl API

MBS Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_MBSPolicyControl

import (
	"encoding/json"
)

// MbsPccRule Represents the parameters constituting an MBS PCC rule.
type MbsPccRule struct {
	MbsPccRuleId string `json:"mbsPccRuleId"`
	MbsDlIpFlowInfo []string `json:"mbsDlIpFlowInfo,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Precedence *int32 `json:"precedence,omitempty"`
	RefMbsQosDec []string `json:"refMbsQosDec,omitempty"`
}

// NewMbsPccRule instantiates a new MbsPccRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsPccRule(mbsPccRuleId string) *MbsPccRule {
	this := MbsPccRule{}
	this.MbsPccRuleId = mbsPccRuleId
	return &this
}

// NewMbsPccRuleWithDefaults instantiates a new MbsPccRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsPccRuleWithDefaults() *MbsPccRule {
	this := MbsPccRule{}
	return &this
}

// GetMbsPccRuleId returns the MbsPccRuleId field value
func (o *MbsPccRule) GetMbsPccRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MbsPccRuleId
}

// GetMbsPccRuleIdOk returns a tuple with the MbsPccRuleId field value
// and a boolean to check if the value has been set.
func (o *MbsPccRule) GetMbsPccRuleIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MbsPccRuleId, true
}

// SetMbsPccRuleId sets field value
func (o *MbsPccRule) SetMbsPccRuleId(v string) {
	o.MbsPccRuleId = v
}

// GetMbsDlIpFlowInfo returns the MbsDlIpFlowInfo field value if set, zero value otherwise.
func (o *MbsPccRule) GetMbsDlIpFlowInfo() []string {
	if o == nil || isNil(o.MbsDlIpFlowInfo) {
		var ret []string
		return ret
	}
	return o.MbsDlIpFlowInfo
}

// GetMbsDlIpFlowInfoOk returns a tuple with the MbsDlIpFlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsPccRule) GetMbsDlIpFlowInfoOk() ([]string, bool) {
	if o == nil || isNil(o.MbsDlIpFlowInfo) {
    return nil, false
	}
	return o.MbsDlIpFlowInfo, true
}

// HasMbsDlIpFlowInfo returns a boolean if a field has been set.
func (o *MbsPccRule) HasMbsDlIpFlowInfo() bool {
	if o != nil && !isNil(o.MbsDlIpFlowInfo) {
		return true
	}

	return false
}

// SetMbsDlIpFlowInfo gets a reference to the given []string and assigns it to the MbsDlIpFlowInfo field.
func (o *MbsPccRule) SetMbsDlIpFlowInfo(v []string) {
	o.MbsDlIpFlowInfo = v
}

// GetPrecedence returns the Precedence field value if set, zero value otherwise.
func (o *MbsPccRule) GetPrecedence() int32 {
	if o == nil || isNil(o.Precedence) {
		var ret int32
		return ret
	}
	return *o.Precedence
}

// GetPrecedenceOk returns a tuple with the Precedence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsPccRule) GetPrecedenceOk() (*int32, bool) {
	if o == nil || isNil(o.Precedence) {
    return nil, false
	}
	return o.Precedence, true
}

// HasPrecedence returns a boolean if a field has been set.
func (o *MbsPccRule) HasPrecedence() bool {
	if o != nil && !isNil(o.Precedence) {
		return true
	}

	return false
}

// SetPrecedence gets a reference to the given int32 and assigns it to the Precedence field.
func (o *MbsPccRule) SetPrecedence(v int32) {
	o.Precedence = &v
}

// GetRefMbsQosDec returns the RefMbsQosDec field value if set, zero value otherwise.
func (o *MbsPccRule) GetRefMbsQosDec() []string {
	if o == nil || isNil(o.RefMbsQosDec) {
		var ret []string
		return ret
	}
	return o.RefMbsQosDec
}

// GetRefMbsQosDecOk returns a tuple with the RefMbsQosDec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsPccRule) GetRefMbsQosDecOk() ([]string, bool) {
	if o == nil || isNil(o.RefMbsQosDec) {
    return nil, false
	}
	return o.RefMbsQosDec, true
}

// HasRefMbsQosDec returns a boolean if a field has been set.
func (o *MbsPccRule) HasRefMbsQosDec() bool {
	if o != nil && !isNil(o.RefMbsQosDec) {
		return true
	}

	return false
}

// SetRefMbsQosDec gets a reference to the given []string and assigns it to the RefMbsQosDec field.
func (o *MbsPccRule) SetRefMbsQosDec(v []string) {
	o.RefMbsQosDec = v
}

func (o MbsPccRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mbsPccRuleId"] = o.MbsPccRuleId
	}
	if !isNil(o.MbsDlIpFlowInfo) {
		toSerialize["mbsDlIpFlowInfo"] = o.MbsDlIpFlowInfo
	}
	if !isNil(o.Precedence) {
		toSerialize["precedence"] = o.Precedence
	}
	if !isNil(o.RefMbsQosDec) {
		toSerialize["refMbsQosDec"] = o.RefMbsQosDec
	}
	return json.Marshal(toSerialize)
}

type NullableMbsPccRule struct {
	value *MbsPccRule
	isSet bool
}

func (v NullableMbsPccRule) Get() *MbsPccRule {
	return v.value
}

func (v *NullableMbsPccRule) Set(val *MbsPccRule) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsPccRule) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsPccRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsPccRule(val *MbsPccRule) *NullableMbsPccRule {
	return &NullableMbsPccRule{value: val, isSet: true}
}

func (v NullableMbsPccRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsPccRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


