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

// MbsReport Contains information about the MBS Policy Decision level failure(s) and/or the MBS PCC rule level failure(s). 
type MbsReport struct {
	MbsPccRuleIds []string `json:"mbsPccRuleIds,omitempty"`
	MbsPccRuleStatus *MbsPccRuleStatus `json:"mbsPccRuleStatus,omitempty"`
	FailureCode *MbsFailureCode `json:"failureCode,omitempty"`
}

// NewMbsReport instantiates a new MbsReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsReport() *MbsReport {
	this := MbsReport{}
	return &this
}

// NewMbsReportWithDefaults instantiates a new MbsReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsReportWithDefaults() *MbsReport {
	this := MbsReport{}
	return &this
}

// GetMbsPccRuleIds returns the MbsPccRuleIds field value if set, zero value otherwise.
func (o *MbsReport) GetMbsPccRuleIds() []string {
	if o == nil || isNil(o.MbsPccRuleIds) {
		var ret []string
		return ret
	}
	return o.MbsPccRuleIds
}

// GetMbsPccRuleIdsOk returns a tuple with the MbsPccRuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsReport) GetMbsPccRuleIdsOk() ([]string, bool) {
	if o == nil || isNil(o.MbsPccRuleIds) {
    return nil, false
	}
	return o.MbsPccRuleIds, true
}

// HasMbsPccRuleIds returns a boolean if a field has been set.
func (o *MbsReport) HasMbsPccRuleIds() bool {
	if o != nil && !isNil(o.MbsPccRuleIds) {
		return true
	}

	return false
}

// SetMbsPccRuleIds gets a reference to the given []string and assigns it to the MbsPccRuleIds field.
func (o *MbsReport) SetMbsPccRuleIds(v []string) {
	o.MbsPccRuleIds = v
}

// GetMbsPccRuleStatus returns the MbsPccRuleStatus field value if set, zero value otherwise.
func (o *MbsReport) GetMbsPccRuleStatus() MbsPccRuleStatus {
	if o == nil || isNil(o.MbsPccRuleStatus) {
		var ret MbsPccRuleStatus
		return ret
	}
	return *o.MbsPccRuleStatus
}

// GetMbsPccRuleStatusOk returns a tuple with the MbsPccRuleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsReport) GetMbsPccRuleStatusOk() (*MbsPccRuleStatus, bool) {
	if o == nil || isNil(o.MbsPccRuleStatus) {
    return nil, false
	}
	return o.MbsPccRuleStatus, true
}

// HasMbsPccRuleStatus returns a boolean if a field has been set.
func (o *MbsReport) HasMbsPccRuleStatus() bool {
	if o != nil && !isNil(o.MbsPccRuleStatus) {
		return true
	}

	return false
}

// SetMbsPccRuleStatus gets a reference to the given MbsPccRuleStatus and assigns it to the MbsPccRuleStatus field.
func (o *MbsReport) SetMbsPccRuleStatus(v MbsPccRuleStatus) {
	o.MbsPccRuleStatus = &v
}

// GetFailureCode returns the FailureCode field value if set, zero value otherwise.
func (o *MbsReport) GetFailureCode() MbsFailureCode {
	if o == nil || isNil(o.FailureCode) {
		var ret MbsFailureCode
		return ret
	}
	return *o.FailureCode
}

// GetFailureCodeOk returns a tuple with the FailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsReport) GetFailureCodeOk() (*MbsFailureCode, bool) {
	if o == nil || isNil(o.FailureCode) {
    return nil, false
	}
	return o.FailureCode, true
}

// HasFailureCode returns a boolean if a field has been set.
func (o *MbsReport) HasFailureCode() bool {
	if o != nil && !isNil(o.FailureCode) {
		return true
	}

	return false
}

// SetFailureCode gets a reference to the given MbsFailureCode and assigns it to the FailureCode field.
func (o *MbsReport) SetFailureCode(v MbsFailureCode) {
	o.FailureCode = &v
}

func (o MbsReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MbsPccRuleIds) {
		toSerialize["mbsPccRuleIds"] = o.MbsPccRuleIds
	}
	if !isNil(o.MbsPccRuleStatus) {
		toSerialize["mbsPccRuleStatus"] = o.MbsPccRuleStatus
	}
	if !isNil(o.FailureCode) {
		toSerialize["failureCode"] = o.FailureCode
	}
	return json.Marshal(toSerialize)
}

type NullableMbsReport struct {
	value *MbsReport
	isSet bool
}

func (v NullableMbsReport) Get() *MbsReport {
	return v.value
}

func (v *NullableMbsReport) Set(val *MbsReport) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsReport) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsReport(val *MbsReport) *NullableMbsReport {
	return &NullableMbsReport{value: val, isSet: true}
}

func (v NullableMbsReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


