/*
Npcf_AMPolicyAuthorization Service API

PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_AMPolicyAuthorization

import (
	"encoding/json"
)

// AmTerminationInfo Includes information related to the termination of the Individual Application AM Context resource.
type AmTerminationInfo struct {
	// Contains the Individual application AM context resource identifier related to the termination notification.
	AppAmContextId string `json:"appAmContextId"`
	TermCause AmTerminationCause `json:"termCause"`
}

// NewAmTerminationInfo instantiates a new AmTerminationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmTerminationInfo(appAmContextId string, termCause AmTerminationCause) *AmTerminationInfo {
	this := AmTerminationInfo{}
	this.AppAmContextId = appAmContextId
	this.TermCause = termCause
	return &this
}

// NewAmTerminationInfoWithDefaults instantiates a new AmTerminationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmTerminationInfoWithDefaults() *AmTerminationInfo {
	this := AmTerminationInfo{}
	return &this
}

// GetAppAmContextId returns the AppAmContextId field value
func (o *AmTerminationInfo) GetAppAmContextId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppAmContextId
}

// GetAppAmContextIdOk returns a tuple with the AppAmContextId field value
// and a boolean to check if the value has been set.
func (o *AmTerminationInfo) GetAppAmContextIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppAmContextId, true
}

// SetAppAmContextId sets field value
func (o *AmTerminationInfo) SetAppAmContextId(v string) {
	o.AppAmContextId = v
}

// GetTermCause returns the TermCause field value
func (o *AmTerminationInfo) GetTermCause() AmTerminationCause {
	if o == nil {
		var ret AmTerminationCause
		return ret
	}

	return o.TermCause
}

// GetTermCauseOk returns a tuple with the TermCause field value
// and a boolean to check if the value has been set.
func (o *AmTerminationInfo) GetTermCauseOk() (*AmTerminationCause, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TermCause, true
}

// SetTermCause sets field value
func (o *AmTerminationInfo) SetTermCause(v AmTerminationCause) {
	o.TermCause = v
}

func (o AmTerminationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appAmContextId"] = o.AppAmContextId
	}
	if true {
		toSerialize["termCause"] = o.TermCause
	}
	return json.Marshal(toSerialize)
}

type NullableAmTerminationInfo struct {
	value *AmTerminationInfo
	isSet bool
}

func (v NullableAmTerminationInfo) Get() *AmTerminationInfo {
	return v.value
}

func (v *NullableAmTerminationInfo) Set(val *AmTerminationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAmTerminationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAmTerminationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmTerminationInfo(val *AmTerminationInfo) *NullableAmTerminationInfo {
	return &NullableAmTerminationInfo{value: val, isSet: true}
}

func (v NullableAmTerminationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmTerminationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


