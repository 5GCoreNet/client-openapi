/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_SMPolicyControl

import (
	"encoding/json"
)

// RequestedRuleData Contains rule data requested by the PCF to receive information associated with PCC rule(s).
type RequestedRuleData struct {
	// An array of PCC rule id references to the PCC rules associated with the control data.
	RefPccRuleIds []string `json:"refPccRuleIds"`
	// Array of requested rule data type elements indicating what type of rule data is requested for the corresponding referenced PCC rules.
	ReqData []RequestedRuleDataType `json:"reqData"`
}

// NewRequestedRuleData instantiates a new RequestedRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestedRuleData(refPccRuleIds []string, reqData []RequestedRuleDataType) *RequestedRuleData {
	this := RequestedRuleData{}
	this.RefPccRuleIds = refPccRuleIds
	this.ReqData = reqData
	return &this
}

// NewRequestedRuleDataWithDefaults instantiates a new RequestedRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestedRuleDataWithDefaults() *RequestedRuleData {
	this := RequestedRuleData{}
	return &this
}

// GetRefPccRuleIds returns the RefPccRuleIds field value
func (o *RequestedRuleData) GetRefPccRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RefPccRuleIds
}

// GetRefPccRuleIdsOk returns a tuple with the RefPccRuleIds field value
// and a boolean to check if the value has been set.
func (o *RequestedRuleData) GetRefPccRuleIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RefPccRuleIds, true
}

// SetRefPccRuleIds sets field value
func (o *RequestedRuleData) SetRefPccRuleIds(v []string) {
	o.RefPccRuleIds = v
}

// GetReqData returns the ReqData field value
func (o *RequestedRuleData) GetReqData() []RequestedRuleDataType {
	if o == nil {
		var ret []RequestedRuleDataType
		return ret
	}

	return o.ReqData
}

// GetReqDataOk returns a tuple with the ReqData field value
// and a boolean to check if the value has been set.
func (o *RequestedRuleData) GetReqDataOk() ([]RequestedRuleDataType, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReqData, true
}

// SetReqData sets field value
func (o *RequestedRuleData) SetReqData(v []RequestedRuleDataType) {
	o.ReqData = v
}

func (o RequestedRuleData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["refPccRuleIds"] = o.RefPccRuleIds
	}
	if true {
		toSerialize["reqData"] = o.ReqData
	}
	return json.Marshal(toSerialize)
}

type NullableRequestedRuleData struct {
	value *RequestedRuleData
	isSet bool
}

func (v NullableRequestedRuleData) Get() *RequestedRuleData {
	return v.value
}

func (v *NullableRequestedRuleData) Set(val *RequestedRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedRuleData(val *RequestedRuleData) *NullableRequestedRuleData {
	return &NullableRequestedRuleData{value: val, isSet: true}
}

func (v NullableRequestedRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestedRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


