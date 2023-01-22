/*
Nnef_EventExposure

NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnef_EventExposure

import (
	"encoding/json"
)

// NefEventExposureNotif Represents notifications on network exposure event(s) that occurred for an Individual Network Exposure Event Subscription resource. 
type NefEventExposureNotif struct {
	NotifId string `json:"notifId"`
	EventNotifs []NefEventNotification `json:"eventNotifs"`
}

// NewNefEventExposureNotif instantiates a new NefEventExposureNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNefEventExposureNotif(notifId string, eventNotifs []NefEventNotification) *NefEventExposureNotif {
	this := NefEventExposureNotif{}
	this.NotifId = notifId
	this.EventNotifs = eventNotifs
	return &this
}

// NewNefEventExposureNotifWithDefaults instantiates a new NefEventExposureNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNefEventExposureNotifWithDefaults() *NefEventExposureNotif {
	this := NefEventExposureNotif{}
	return &this
}

// GetNotifId returns the NotifId field value
func (o *NefEventExposureNotif) GetNotifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifId
}

// GetNotifIdOk returns a tuple with the NotifId field value
// and a boolean to check if the value has been set.
func (o *NefEventExposureNotif) GetNotifIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifId, true
}

// SetNotifId sets field value
func (o *NefEventExposureNotif) SetNotifId(v string) {
	o.NotifId = v
}

// GetEventNotifs returns the EventNotifs field value
func (o *NefEventExposureNotif) GetEventNotifs() []NefEventNotification {
	if o == nil {
		var ret []NefEventNotification
		return ret
	}

	return o.EventNotifs
}

// GetEventNotifsOk returns a tuple with the EventNotifs field value
// and a boolean to check if the value has been set.
func (o *NefEventExposureNotif) GetEventNotifsOk() ([]NefEventNotification, bool) {
	if o == nil {
    return nil, false
	}
	return o.EventNotifs, true
}

// SetEventNotifs sets field value
func (o *NefEventExposureNotif) SetEventNotifs(v []NefEventNotification) {
	o.EventNotifs = v
}

func (o NefEventExposureNotif) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["notifId"] = o.NotifId
	}
	if true {
		toSerialize["eventNotifs"] = o.EventNotifs
	}
	return json.Marshal(toSerialize)
}

type NullableNefEventExposureNotif struct {
	value *NefEventExposureNotif
	isSet bool
}

func (v NullableNefEventExposureNotif) Get() *NefEventExposureNotif {
	return v.value
}

func (v *NullableNefEventExposureNotif) Set(val *NefEventExposureNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableNefEventExposureNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableNefEventExposureNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNefEventExposureNotif(val *NefEventExposureNotif) *NullableNefEventExposureNotif {
	return &NullableNefEventExposureNotif{value: val, isSet: true}
}

func (v NullableNefEventExposureNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNefEventExposureNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


