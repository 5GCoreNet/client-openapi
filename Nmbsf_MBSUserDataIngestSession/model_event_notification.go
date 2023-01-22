/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
	"time"
)

// EventNotification Represents Event Notification.
type EventNotification struct {
	StatusEvent Event `json:"statusEvent"`
	MbsDisSessionId *string `json:"mbsDisSessionId,omitempty"`
	StatusAddInfo *string `json:"statusAddInfo,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
}

// NewEventNotification instantiates a new EventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventNotification(statusEvent Event, timeStamp time.Time) *EventNotification {
	this := EventNotification{}
	this.StatusEvent = statusEvent
	this.TimeStamp = timeStamp
	return &this
}

// NewEventNotificationWithDefaults instantiates a new EventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventNotificationWithDefaults() *EventNotification {
	this := EventNotification{}
	return &this
}

// GetStatusEvent returns the StatusEvent field value
func (o *EventNotification) GetStatusEvent() Event {
	if o == nil {
		var ret Event
		return ret
	}

	return o.StatusEvent
}

// GetStatusEventOk returns a tuple with the StatusEvent field value
// and a boolean to check if the value has been set.
func (o *EventNotification) GetStatusEventOk() (*Event, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StatusEvent, true
}

// SetStatusEvent sets field value
func (o *EventNotification) SetStatusEvent(v Event) {
	o.StatusEvent = v
}

// GetMbsDisSessionId returns the MbsDisSessionId field value if set, zero value otherwise.
func (o *EventNotification) GetMbsDisSessionId() string {
	if o == nil || isNil(o.MbsDisSessionId) {
		var ret string
		return ret
	}
	return *o.MbsDisSessionId
}

// GetMbsDisSessionIdOk returns a tuple with the MbsDisSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetMbsDisSessionIdOk() (*string, bool) {
	if o == nil || isNil(o.MbsDisSessionId) {
    return nil, false
	}
	return o.MbsDisSessionId, true
}

// HasMbsDisSessionId returns a boolean if a field has been set.
func (o *EventNotification) HasMbsDisSessionId() bool {
	if o != nil && !isNil(o.MbsDisSessionId) {
		return true
	}

	return false
}

// SetMbsDisSessionId gets a reference to the given string and assigns it to the MbsDisSessionId field.
func (o *EventNotification) SetMbsDisSessionId(v string) {
	o.MbsDisSessionId = &v
}

// GetStatusAddInfo returns the StatusAddInfo field value if set, zero value otherwise.
func (o *EventNotification) GetStatusAddInfo() string {
	if o == nil || isNil(o.StatusAddInfo) {
		var ret string
		return ret
	}
	return *o.StatusAddInfo
}

// GetStatusAddInfoOk returns a tuple with the StatusAddInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetStatusAddInfoOk() (*string, bool) {
	if o == nil || isNil(o.StatusAddInfo) {
    return nil, false
	}
	return o.StatusAddInfo, true
}

// HasStatusAddInfo returns a boolean if a field has been set.
func (o *EventNotification) HasStatusAddInfo() bool {
	if o != nil && !isNil(o.StatusAddInfo) {
		return true
	}

	return false
}

// SetStatusAddInfo gets a reference to the given string and assigns it to the StatusAddInfo field.
func (o *EventNotification) SetStatusAddInfo(v string) {
	o.StatusAddInfo = &v
}

// GetTimeStamp returns the TimeStamp field value
func (o *EventNotification) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *EventNotification) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *EventNotification) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

func (o EventNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["statusEvent"] = o.StatusEvent
	}
	if !isNil(o.MbsDisSessionId) {
		toSerialize["mbsDisSessionId"] = o.MbsDisSessionId
	}
	if !isNil(o.StatusAddInfo) {
		toSerialize["statusAddInfo"] = o.StatusAddInfo
	}
	if true {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	return json.Marshal(toSerialize)
}

type NullableEventNotification struct {
	value *EventNotification
	isSet bool
}

func (v NullableEventNotification) Get() *EventNotification {
	return v.value
}

func (v *NullableEventNotification) Set(val *EventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventNotification(val *EventNotification) *NullableEventNotification {
	return &NullableEventNotification{value: val, isSet: true}
}

func (v NullableEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


