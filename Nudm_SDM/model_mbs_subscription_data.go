/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_SDM

import (
	"encoding/json"
)

// MbsSubscriptionData Contains the 5MBS Subscription Data.
type MbsSubscriptionData struct {
	MbsAllowed *bool `json:"mbsAllowed,omitempty"`
	MbsSessionIdList []MbsSessionId `json:"mbsSessionIdList,omitempty"`
}

// NewMbsSubscriptionData instantiates a new MbsSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSubscriptionData() *MbsSubscriptionData {
	this := MbsSubscriptionData{}
	var mbsAllowed bool = false
	this.MbsAllowed = &mbsAllowed
	return &this
}

// NewMbsSubscriptionDataWithDefaults instantiates a new MbsSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSubscriptionDataWithDefaults() *MbsSubscriptionData {
	this := MbsSubscriptionData{}
	var mbsAllowed bool = false
	this.MbsAllowed = &mbsAllowed
	return &this
}

// GetMbsAllowed returns the MbsAllowed field value if set, zero value otherwise.
func (o *MbsSubscriptionData) GetMbsAllowed() bool {
	if o == nil || isNil(o.MbsAllowed) {
		var ret bool
		return ret
	}
	return *o.MbsAllowed
}

// GetMbsAllowedOk returns a tuple with the MbsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSubscriptionData) GetMbsAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.MbsAllowed) {
    return nil, false
	}
	return o.MbsAllowed, true
}

// HasMbsAllowed returns a boolean if a field has been set.
func (o *MbsSubscriptionData) HasMbsAllowed() bool {
	if o != nil && !isNil(o.MbsAllowed) {
		return true
	}

	return false
}

// SetMbsAllowed gets a reference to the given bool and assigns it to the MbsAllowed field.
func (o *MbsSubscriptionData) SetMbsAllowed(v bool) {
	o.MbsAllowed = &v
}

// GetMbsSessionIdList returns the MbsSessionIdList field value if set, zero value otherwise.
func (o *MbsSubscriptionData) GetMbsSessionIdList() []MbsSessionId {
	if o == nil || isNil(o.MbsSessionIdList) {
		var ret []MbsSessionId
		return ret
	}
	return o.MbsSessionIdList
}

// GetMbsSessionIdListOk returns a tuple with the MbsSessionIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSubscriptionData) GetMbsSessionIdListOk() ([]MbsSessionId, bool) {
	if o == nil || isNil(o.MbsSessionIdList) {
    return nil, false
	}
	return o.MbsSessionIdList, true
}

// HasMbsSessionIdList returns a boolean if a field has been set.
func (o *MbsSubscriptionData) HasMbsSessionIdList() bool {
	if o != nil && !isNil(o.MbsSessionIdList) {
		return true
	}

	return false
}

// SetMbsSessionIdList gets a reference to the given []MbsSessionId and assigns it to the MbsSessionIdList field.
func (o *MbsSubscriptionData) SetMbsSessionIdList(v []MbsSessionId) {
	o.MbsSessionIdList = v
}

func (o MbsSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MbsAllowed) {
		toSerialize["mbsAllowed"] = o.MbsAllowed
	}
	if !isNil(o.MbsSessionIdList) {
		toSerialize["mbsSessionIdList"] = o.MbsSessionIdList
	}
	return json.Marshal(toSerialize)
}

type NullableMbsSubscriptionData struct {
	value *MbsSubscriptionData
	isSet bool
}

func (v NullableMbsSubscriptionData) Get() *MbsSubscriptionData {
	return v.value
}

func (v *NullableMbsSubscriptionData) Set(val *MbsSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSubscriptionData(val *MbsSubscriptionData) *NullableMbsSubscriptionData {
	return &NullableMbsSubscriptionData{value: val, isSet: true}
}

func (v NullableMbsSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


