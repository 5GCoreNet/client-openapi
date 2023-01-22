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

// SessionRule Contains session level policy information.
type SessionRule struct {
	AuthSessAmbr *Ambr `json:"authSessAmbr,omitempty"`
	AuthDefQos *AuthorizedDefaultQos `json:"authDefQos,omitempty"`
	// Univocally identifies the session rule within a PDU session.
	SessRuleId string `json:"sessRuleId"`
	// A reference to UsageMonitoringData policy decision type. It is the umId described in clause 5.6.2.12.
	RefUmData NullableString `json:"refUmData,omitempty"`
	// A reference to UsageMonitoringData policy decision type to apply for Non-3GPP access. It is the umId described in clause 5.6.2.12.
	RefUmN3gData NullableString `json:"refUmN3gData,omitempty"`
	// A reference to the condition data. It is the condId described in clause 5.6.2.9.
	RefCondData NullableString `json:"refCondData,omitempty"`
}

// NewSessionRule instantiates a new SessionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionRule(sessRuleId string) *SessionRule {
	this := SessionRule{}
	this.SessRuleId = sessRuleId
	return &this
}

// NewSessionRuleWithDefaults instantiates a new SessionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionRuleWithDefaults() *SessionRule {
	this := SessionRule{}
	return &this
}

// GetAuthSessAmbr returns the AuthSessAmbr field value if set, zero value otherwise.
func (o *SessionRule) GetAuthSessAmbr() Ambr {
	if o == nil || isNil(o.AuthSessAmbr) {
		var ret Ambr
		return ret
	}
	return *o.AuthSessAmbr
}

// GetAuthSessAmbrOk returns a tuple with the AuthSessAmbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionRule) GetAuthSessAmbrOk() (*Ambr, bool) {
	if o == nil || isNil(o.AuthSessAmbr) {
    return nil, false
	}
	return o.AuthSessAmbr, true
}

// HasAuthSessAmbr returns a boolean if a field has been set.
func (o *SessionRule) HasAuthSessAmbr() bool {
	if o != nil && !isNil(o.AuthSessAmbr) {
		return true
	}

	return false
}

// SetAuthSessAmbr gets a reference to the given Ambr and assigns it to the AuthSessAmbr field.
func (o *SessionRule) SetAuthSessAmbr(v Ambr) {
	o.AuthSessAmbr = &v
}

// GetAuthDefQos returns the AuthDefQos field value if set, zero value otherwise.
func (o *SessionRule) GetAuthDefQos() AuthorizedDefaultQos {
	if o == nil || isNil(o.AuthDefQos) {
		var ret AuthorizedDefaultQos
		return ret
	}
	return *o.AuthDefQos
}

// GetAuthDefQosOk returns a tuple with the AuthDefQos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionRule) GetAuthDefQosOk() (*AuthorizedDefaultQos, bool) {
	if o == nil || isNil(o.AuthDefQos) {
    return nil, false
	}
	return o.AuthDefQos, true
}

// HasAuthDefQos returns a boolean if a field has been set.
func (o *SessionRule) HasAuthDefQos() bool {
	if o != nil && !isNil(o.AuthDefQos) {
		return true
	}

	return false
}

// SetAuthDefQos gets a reference to the given AuthorizedDefaultQos and assigns it to the AuthDefQos field.
func (o *SessionRule) SetAuthDefQos(v AuthorizedDefaultQos) {
	o.AuthDefQos = &v
}

// GetSessRuleId returns the SessRuleId field value
func (o *SessionRule) GetSessRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessRuleId
}

// GetSessRuleIdOk returns a tuple with the SessRuleId field value
// and a boolean to check if the value has been set.
func (o *SessionRule) GetSessRuleIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SessRuleId, true
}

// SetSessRuleId sets field value
func (o *SessionRule) SetSessRuleId(v string) {
	o.SessRuleId = v
}

// GetRefUmData returns the RefUmData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionRule) GetRefUmData() string {
	if o == nil || isNil(o.RefUmData.Get()) {
		var ret string
		return ret
	}
	return *o.RefUmData.Get()
}

// GetRefUmDataOk returns a tuple with the RefUmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionRule) GetRefUmDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RefUmData.Get(), o.RefUmData.IsSet()
}

// HasRefUmData returns a boolean if a field has been set.
func (o *SessionRule) HasRefUmData() bool {
	if o != nil && o.RefUmData.IsSet() {
		return true
	}

	return false
}

// SetRefUmData gets a reference to the given NullableString and assigns it to the RefUmData field.
func (o *SessionRule) SetRefUmData(v string) {
	o.RefUmData.Set(&v)
}
// SetRefUmDataNil sets the value for RefUmData to be an explicit nil
func (o *SessionRule) SetRefUmDataNil() {
	o.RefUmData.Set(nil)
}

// UnsetRefUmData ensures that no value is present for RefUmData, not even an explicit nil
func (o *SessionRule) UnsetRefUmData() {
	o.RefUmData.Unset()
}

// GetRefUmN3gData returns the RefUmN3gData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionRule) GetRefUmN3gData() string {
	if o == nil || isNil(o.RefUmN3gData.Get()) {
		var ret string
		return ret
	}
	return *o.RefUmN3gData.Get()
}

// GetRefUmN3gDataOk returns a tuple with the RefUmN3gData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionRule) GetRefUmN3gDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RefUmN3gData.Get(), o.RefUmN3gData.IsSet()
}

// HasRefUmN3gData returns a boolean if a field has been set.
func (o *SessionRule) HasRefUmN3gData() bool {
	if o != nil && o.RefUmN3gData.IsSet() {
		return true
	}

	return false
}

// SetRefUmN3gData gets a reference to the given NullableString and assigns it to the RefUmN3gData field.
func (o *SessionRule) SetRefUmN3gData(v string) {
	o.RefUmN3gData.Set(&v)
}
// SetRefUmN3gDataNil sets the value for RefUmN3gData to be an explicit nil
func (o *SessionRule) SetRefUmN3gDataNil() {
	o.RefUmN3gData.Set(nil)
}

// UnsetRefUmN3gData ensures that no value is present for RefUmN3gData, not even an explicit nil
func (o *SessionRule) UnsetRefUmN3gData() {
	o.RefUmN3gData.Unset()
}

// GetRefCondData returns the RefCondData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionRule) GetRefCondData() string {
	if o == nil || isNil(o.RefCondData.Get()) {
		var ret string
		return ret
	}
	return *o.RefCondData.Get()
}

// GetRefCondDataOk returns a tuple with the RefCondData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionRule) GetRefCondDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RefCondData.Get(), o.RefCondData.IsSet()
}

// HasRefCondData returns a boolean if a field has been set.
func (o *SessionRule) HasRefCondData() bool {
	if o != nil && o.RefCondData.IsSet() {
		return true
	}

	return false
}

// SetRefCondData gets a reference to the given NullableString and assigns it to the RefCondData field.
func (o *SessionRule) SetRefCondData(v string) {
	o.RefCondData.Set(&v)
}
// SetRefCondDataNil sets the value for RefCondData to be an explicit nil
func (o *SessionRule) SetRefCondDataNil() {
	o.RefCondData.Set(nil)
}

// UnsetRefCondData ensures that no value is present for RefCondData, not even an explicit nil
func (o *SessionRule) UnsetRefCondData() {
	o.RefCondData.Unset()
}

func (o SessionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AuthSessAmbr) {
		toSerialize["authSessAmbr"] = o.AuthSessAmbr
	}
	if !isNil(o.AuthDefQos) {
		toSerialize["authDefQos"] = o.AuthDefQos
	}
	if true {
		toSerialize["sessRuleId"] = o.SessRuleId
	}
	if o.RefUmData.IsSet() {
		toSerialize["refUmData"] = o.RefUmData.Get()
	}
	if o.RefUmN3gData.IsSet() {
		toSerialize["refUmN3gData"] = o.RefUmN3gData.Get()
	}
	if o.RefCondData.IsSet() {
		toSerialize["refCondData"] = o.RefCondData.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSessionRule struct {
	value *SessionRule
	isSet bool
}

func (v NullableSessionRule) Get() *SessionRule {
	return v.value
}

func (v *NullableSessionRule) Set(val *SessionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionRule(val *SessionRule) *NullableSessionRule {
	return &NullableSessionRule{value: val, isSet: true}
}

func (v NullableSessionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


