/*
Nmbstf-distsession

MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbstf_DistSession

import (
	"encoding/json"
)

// DistSessionEventReportList List of Event Reports
type DistSessionEventReportList struct {
	EventReportList []DistSessionEventReport `json:"eventReportList"`
	NotifyCorrelationId *string `json:"notifyCorrelationId,omitempty"`
}

// NewDistSessionEventReportList instantiates a new DistSessionEventReportList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistSessionEventReportList(eventReportList []DistSessionEventReport) *DistSessionEventReportList {
	this := DistSessionEventReportList{}
	this.EventReportList = eventReportList
	return &this
}

// NewDistSessionEventReportListWithDefaults instantiates a new DistSessionEventReportList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistSessionEventReportListWithDefaults() *DistSessionEventReportList {
	this := DistSessionEventReportList{}
	return &this
}

// GetEventReportList returns the EventReportList field value
func (o *DistSessionEventReportList) GetEventReportList() []DistSessionEventReport {
	if o == nil {
		var ret []DistSessionEventReport
		return ret
	}

	return o.EventReportList
}

// GetEventReportListOk returns a tuple with the EventReportList field value
// and a boolean to check if the value has been set.
func (o *DistSessionEventReportList) GetEventReportListOk() ([]DistSessionEventReport, bool) {
	if o == nil {
    return nil, false
	}
	return o.EventReportList, true
}

// SetEventReportList sets field value
func (o *DistSessionEventReportList) SetEventReportList(v []DistSessionEventReport) {
	o.EventReportList = v
}

// GetNotifyCorrelationId returns the NotifyCorrelationId field value if set, zero value otherwise.
func (o *DistSessionEventReportList) GetNotifyCorrelationId() string {
	if o == nil || isNil(o.NotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrelationId
}

// GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistSessionEventReportList) GetNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || isNil(o.NotifyCorrelationId) {
    return nil, false
	}
	return o.NotifyCorrelationId, true
}

// HasNotifyCorrelationId returns a boolean if a field has been set.
func (o *DistSessionEventReportList) HasNotifyCorrelationId() bool {
	if o != nil && !isNil(o.NotifyCorrelationId) {
		return true
	}

	return false
}

// SetNotifyCorrelationId gets a reference to the given string and assigns it to the NotifyCorrelationId field.
func (o *DistSessionEventReportList) SetNotifyCorrelationId(v string) {
	o.NotifyCorrelationId = &v
}

func (o DistSessionEventReportList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventReportList"] = o.EventReportList
	}
	if !isNil(o.NotifyCorrelationId) {
		toSerialize["notifyCorrelationId"] = o.NotifyCorrelationId
	}
	return json.Marshal(toSerialize)
}

type NullableDistSessionEventReportList struct {
	value *DistSessionEventReportList
	isSet bool
}

func (v NullableDistSessionEventReportList) Get() *DistSessionEventReportList {
	return v.value
}

func (v *NullableDistSessionEventReportList) Set(val *DistSessionEventReportList) {
	v.value = val
	v.isSet = true
}

func (v NullableDistSessionEventReportList) IsSet() bool {
	return v.isSet
}

func (v *NullableDistSessionEventReportList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistSessionEventReportList(val *DistSessionEventReportList) *NullableDistSessionEventReportList {
	return &NullableDistSessionEventReportList{value: val, isSet: true}
}

func (v NullableDistSessionEventReportList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistSessionEventReportList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


