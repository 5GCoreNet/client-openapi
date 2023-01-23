/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
)

// SessionRuleReport Represents reporting of the status of a session rule.
type SessionRuleReport struct {
	// Contains the identifier of the affected session rule(s).
	RuleIds []string `json:"ruleIds"`
	RuleStatus RuleStatus `json:"ruleStatus"`
	SessRuleFailureCode *SessionRuleFailureCode `json:"sessRuleFailureCode,omitempty"`
	// Contains the type(s) of failed policy decision and/or condition data.
	PolicyDecFailureReports []PolicyDecisionFailureCode `json:"policyDecFailureReports,omitempty"`
}

// NewSessionRuleReport instantiates a new SessionRuleReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionRuleReport(ruleIds []string, ruleStatus RuleStatus) *SessionRuleReport {
	this := SessionRuleReport{}
	this.RuleIds = ruleIds
	this.RuleStatus = ruleStatus
	return &this
}

// NewSessionRuleReportWithDefaults instantiates a new SessionRuleReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionRuleReportWithDefaults() *SessionRuleReport {
	this := SessionRuleReport{}
	return &this
}

// GetRuleIds returns the RuleIds field value
func (o *SessionRuleReport) GetRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value
// and a boolean to check if the value has been set.
func (o *SessionRuleReport) GetRuleIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RuleIds, true
}

// SetRuleIds sets field value
func (o *SessionRuleReport) SetRuleIds(v []string) {
	o.RuleIds = v
}

// GetRuleStatus returns the RuleStatus field value
func (o *SessionRuleReport) GetRuleStatus() RuleStatus {
	if o == nil {
		var ret RuleStatus
		return ret
	}

	return o.RuleStatus
}

// GetRuleStatusOk returns a tuple with the RuleStatus field value
// and a boolean to check if the value has been set.
func (o *SessionRuleReport) GetRuleStatusOk() (*RuleStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RuleStatus, true
}

// SetRuleStatus sets field value
func (o *SessionRuleReport) SetRuleStatus(v RuleStatus) {
	o.RuleStatus = v
}

// GetSessRuleFailureCode returns the SessRuleFailureCode field value if set, zero value otherwise.
func (o *SessionRuleReport) GetSessRuleFailureCode() SessionRuleFailureCode {
	if o == nil || isNil(o.SessRuleFailureCode) {
		var ret SessionRuleFailureCode
		return ret
	}
	return *o.SessRuleFailureCode
}

// GetSessRuleFailureCodeOk returns a tuple with the SessRuleFailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionRuleReport) GetSessRuleFailureCodeOk() (*SessionRuleFailureCode, bool) {
	if o == nil || isNil(o.SessRuleFailureCode) {
    return nil, false
	}
	return o.SessRuleFailureCode, true
}

// HasSessRuleFailureCode returns a boolean if a field has been set.
func (o *SessionRuleReport) HasSessRuleFailureCode() bool {
	if o != nil && !isNil(o.SessRuleFailureCode) {
		return true
	}

	return false
}

// SetSessRuleFailureCode gets a reference to the given SessionRuleFailureCode and assigns it to the SessRuleFailureCode field.
func (o *SessionRuleReport) SetSessRuleFailureCode(v SessionRuleFailureCode) {
	o.SessRuleFailureCode = &v
}

// GetPolicyDecFailureReports returns the PolicyDecFailureReports field value if set, zero value otherwise.
func (o *SessionRuleReport) GetPolicyDecFailureReports() []PolicyDecisionFailureCode {
	if o == nil || isNil(o.PolicyDecFailureReports) {
		var ret []PolicyDecisionFailureCode
		return ret
	}
	return o.PolicyDecFailureReports
}

// GetPolicyDecFailureReportsOk returns a tuple with the PolicyDecFailureReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionRuleReport) GetPolicyDecFailureReportsOk() ([]PolicyDecisionFailureCode, bool) {
	if o == nil || isNil(o.PolicyDecFailureReports) {
    return nil, false
	}
	return o.PolicyDecFailureReports, true
}

// HasPolicyDecFailureReports returns a boolean if a field has been set.
func (o *SessionRuleReport) HasPolicyDecFailureReports() bool {
	if o != nil && !isNil(o.PolicyDecFailureReports) {
		return true
	}

	return false
}

// SetPolicyDecFailureReports gets a reference to the given []PolicyDecisionFailureCode and assigns it to the PolicyDecFailureReports field.
func (o *SessionRuleReport) SetPolicyDecFailureReports(v []PolicyDecisionFailureCode) {
	o.PolicyDecFailureReports = v
}

func (o SessionRuleReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ruleIds"] = o.RuleIds
	}
	if true {
		toSerialize["ruleStatus"] = o.RuleStatus
	}
	if !isNil(o.SessRuleFailureCode) {
		toSerialize["sessRuleFailureCode"] = o.SessRuleFailureCode
	}
	if !isNil(o.PolicyDecFailureReports) {
		toSerialize["policyDecFailureReports"] = o.PolicyDecFailureReports
	}
	return json.Marshal(toSerialize)
}

type NullableSessionRuleReport struct {
	value *SessionRuleReport
	isSet bool
}

func (v NullableSessionRuleReport) Get() *SessionRuleReport {
	return v.value
}

func (v *NullableSessionRuleReport) Set(val *SessionRuleReport) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionRuleReport) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionRuleReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionRuleReport(val *SessionRuleReport) *NullableSessionRuleReport {
	return &NullableSessionRuleReport{value: val, isSet: true}
}

func (v NullableSessionRuleReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionRuleReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


