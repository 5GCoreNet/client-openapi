/*
Nmbstf-distsession

MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbstf_DistSession

import (
	"encoding/json"
)

// StatusSubscribeRspData Data within StatusSubscribe Response
type StatusSubscribeRspData struct {
	Subscription DistSessionSubscription `json:"subscription"`
	ReportList *DistSessionEventReportList `json:"reportList,omitempty"`
}

// NewStatusSubscribeRspData instantiates a new StatusSubscribeRspData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusSubscribeRspData(subscription DistSessionSubscription) *StatusSubscribeRspData {
	this := StatusSubscribeRspData{}
	this.Subscription = subscription
	return &this
}

// NewStatusSubscribeRspDataWithDefaults instantiates a new StatusSubscribeRspData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusSubscribeRspDataWithDefaults() *StatusSubscribeRspData {
	this := StatusSubscribeRspData{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *StatusSubscribeRspData) GetSubscription() DistSessionSubscription {
	if o == nil {
		var ret DistSessionSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *StatusSubscribeRspData) GetSubscriptionOk() (*DistSessionSubscription, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *StatusSubscribeRspData) SetSubscription(v DistSessionSubscription) {
	o.Subscription = v
}

// GetReportList returns the ReportList field value if set, zero value otherwise.
func (o *StatusSubscribeRspData) GetReportList() DistSessionEventReportList {
	if o == nil || isNil(o.ReportList) {
		var ret DistSessionEventReportList
		return ret
	}
	return *o.ReportList
}

// GetReportListOk returns a tuple with the ReportList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusSubscribeRspData) GetReportListOk() (*DistSessionEventReportList, bool) {
	if o == nil || isNil(o.ReportList) {
    return nil, false
	}
	return o.ReportList, true
}

// HasReportList returns a boolean if a field has been set.
func (o *StatusSubscribeRspData) HasReportList() bool {
	if o != nil && !isNil(o.ReportList) {
		return true
	}

	return false
}

// SetReportList gets a reference to the given DistSessionEventReportList and assigns it to the ReportList field.
func (o *StatusSubscribeRspData) SetReportList(v DistSessionEventReportList) {
	o.ReportList = &v
}

func (o StatusSubscribeRspData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscription"] = o.Subscription
	}
	if !isNil(o.ReportList) {
		toSerialize["reportList"] = o.ReportList
	}
	return json.Marshal(toSerialize)
}

type NullableStatusSubscribeRspData struct {
	value *StatusSubscribeRspData
	isSet bool
}

func (v NullableStatusSubscribeRspData) Get() *StatusSubscribeRspData {
	return v.value
}

func (v *NullableStatusSubscribeRspData) Set(val *StatusSubscribeRspData) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusSubscribeRspData) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusSubscribeRspData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusSubscribeRspData(val *StatusSubscribeRspData) *NullableStatusSubscribeRspData {
	return &NullableStatusSubscribeRspData{value: val, isSet: true}
}

func (v NullableStatusSubscribeRspData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusSubscribeRspData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


