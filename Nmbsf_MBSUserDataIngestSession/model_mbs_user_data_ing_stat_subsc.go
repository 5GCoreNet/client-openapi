/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
)

// MBSUserDataIngStatSubsc Represents an MBS User Data Ingest Session Status Subscription. 
type MBSUserDataIngStatSubsc struct {
	MbsIngSessionId string `json:"mbsIngSessionId"`
	EventSubscs []SubscribedEvent `json:"eventSubscs"`
	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`
}

// NewMBSUserDataIngStatSubsc instantiates a new MBSUserDataIngStatSubsc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMBSUserDataIngStatSubsc(mbsIngSessionId string, eventSubscs []SubscribedEvent, notifUri string) *MBSUserDataIngStatSubsc {
	this := MBSUserDataIngStatSubsc{}
	this.MbsIngSessionId = mbsIngSessionId
	this.EventSubscs = eventSubscs
	this.NotifUri = notifUri
	return &this
}

// NewMBSUserDataIngStatSubscWithDefaults instantiates a new MBSUserDataIngStatSubsc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMBSUserDataIngStatSubscWithDefaults() *MBSUserDataIngStatSubsc {
	this := MBSUserDataIngStatSubsc{}
	return &this
}

// GetMbsIngSessionId returns the MbsIngSessionId field value
func (o *MBSUserDataIngStatSubsc) GetMbsIngSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MbsIngSessionId
}

// GetMbsIngSessionIdOk returns a tuple with the MbsIngSessionId field value
// and a boolean to check if the value has been set.
func (o *MBSUserDataIngStatSubsc) GetMbsIngSessionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MbsIngSessionId, true
}

// SetMbsIngSessionId sets field value
func (o *MBSUserDataIngStatSubsc) SetMbsIngSessionId(v string) {
	o.MbsIngSessionId = v
}

// GetEventSubscs returns the EventSubscs field value
func (o *MBSUserDataIngStatSubsc) GetEventSubscs() []SubscribedEvent {
	if o == nil {
		var ret []SubscribedEvent
		return ret
	}

	return o.EventSubscs
}

// GetEventSubscsOk returns a tuple with the EventSubscs field value
// and a boolean to check if the value has been set.
func (o *MBSUserDataIngStatSubsc) GetEventSubscsOk() ([]SubscribedEvent, bool) {
	if o == nil {
    return nil, false
	}
	return o.EventSubscs, true
}

// SetEventSubscs sets field value
func (o *MBSUserDataIngStatSubsc) SetEventSubscs(v []SubscribedEvent) {
	o.EventSubscs = v
}

// GetNotifUri returns the NotifUri field value
func (o *MBSUserDataIngStatSubsc) GetNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value
// and a boolean to check if the value has been set.
func (o *MBSUserDataIngStatSubsc) GetNotifUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifUri, true
}

// SetNotifUri sets field value
func (o *MBSUserDataIngStatSubsc) SetNotifUri(v string) {
	o.NotifUri = v
}

func (o MBSUserDataIngStatSubsc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mbsIngSessionId"] = o.MbsIngSessionId
	}
	if true {
		toSerialize["eventSubscs"] = o.EventSubscs
	}
	if true {
		toSerialize["notifUri"] = o.NotifUri
	}
	return json.Marshal(toSerialize)
}

type NullableMBSUserDataIngStatSubsc struct {
	value *MBSUserDataIngStatSubsc
	isSet bool
}

func (v NullableMBSUserDataIngStatSubsc) Get() *MBSUserDataIngStatSubsc {
	return v.value
}

func (v *NullableMBSUserDataIngStatSubsc) Set(val *MBSUserDataIngStatSubsc) {
	v.value = val
	v.isSet = true
}

func (v NullableMBSUserDataIngStatSubsc) IsSet() bool {
	return v.isSet
}

func (v *NullableMBSUserDataIngStatSubsc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMBSUserDataIngStatSubsc(val *MBSUserDataIngStatSubsc) *NullableMBSUserDataIngStatSubsc {
	return &NullableMBSUserDataIngStatSubsc{value: val, isSet: true}
}

func (v NullableMBSUserDataIngStatSubsc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMBSUserDataIngStatSubsc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


