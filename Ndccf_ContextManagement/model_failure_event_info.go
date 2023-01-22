/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_ContextManagement

import (
	"encoding/json"
)

// FailureEventInfo Contains information on the event for which the subscription is not successful.
type FailureEventInfo struct {
	Event NwdafEvent `json:"event"`
	FailureCode NwdafFailureCode `json:"failureCode"`
}

// NewFailureEventInfo instantiates a new FailureEventInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailureEventInfo(event NwdafEvent, failureCode NwdafFailureCode) *FailureEventInfo {
	this := FailureEventInfo{}
	this.Event = event
	this.FailureCode = failureCode
	return &this
}

// NewFailureEventInfoWithDefaults instantiates a new FailureEventInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailureEventInfoWithDefaults() *FailureEventInfo {
	this := FailureEventInfo{}
	return &this
}

// GetEvent returns the Event field value
func (o *FailureEventInfo) GetEvent() NwdafEvent {
	if o == nil {
		var ret NwdafEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *FailureEventInfo) GetEventOk() (*NwdafEvent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *FailureEventInfo) SetEvent(v NwdafEvent) {
	o.Event = v
}

// GetFailureCode returns the FailureCode field value
func (o *FailureEventInfo) GetFailureCode() NwdafFailureCode {
	if o == nil {
		var ret NwdafFailureCode
		return ret
	}

	return o.FailureCode
}

// GetFailureCodeOk returns a tuple with the FailureCode field value
// and a boolean to check if the value has been set.
func (o *FailureEventInfo) GetFailureCodeOk() (*NwdafFailureCode, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FailureCode, true
}

// SetFailureCode sets field value
func (o *FailureEventInfo) SetFailureCode(v NwdafFailureCode) {
	o.FailureCode = v
}

func (o FailureEventInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["failureCode"] = o.FailureCode
	}
	return json.Marshal(toSerialize)
}

type NullableFailureEventInfo struct {
	value *FailureEventInfo
	isSet bool
}

func (v NullableFailureEventInfo) Get() *FailureEventInfo {
	return v.value
}

func (v *NullableFailureEventInfo) Set(val *FailureEventInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureEventInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureEventInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureEventInfo(val *FailureEventInfo) *NullableFailureEventInfo {
	return &NullableFailureEventInfo{value: val, isSet: true}
}

func (v NullableFailureEventInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureEventInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


