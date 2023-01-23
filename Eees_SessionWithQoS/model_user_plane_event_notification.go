/*
EES Session with QoS API

API for EES Session with Qos service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_SessionWithQoS

import (
	"encoding/json"
)

// UserPlaneEventNotification Represents the user plane event notification.
type UserPlaneEventNotification struct {
	// String identifying the individual data session information for which the QoS event notification is delivered. 
	SessionId string `json:"sessionId"`
	// Contains the flow description for the Uplink and/or Downlink IP flows. 
	EventReports []UserPlaneEventReport `json:"eventReports"`
}

// NewUserPlaneEventNotification instantiates a new UserPlaneEventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPlaneEventNotification(sessionId string, eventReports []UserPlaneEventReport) *UserPlaneEventNotification {
	this := UserPlaneEventNotification{}
	this.SessionId = sessionId
	this.EventReports = eventReports
	return &this
}

// NewUserPlaneEventNotificationWithDefaults instantiates a new UserPlaneEventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPlaneEventNotificationWithDefaults() *UserPlaneEventNotification {
	this := UserPlaneEventNotification{}
	return &this
}

// GetSessionId returns the SessionId field value
func (o *UserPlaneEventNotification) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *UserPlaneEventNotification) GetSessionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *UserPlaneEventNotification) SetSessionId(v string) {
	o.SessionId = v
}

// GetEventReports returns the EventReports field value
func (o *UserPlaneEventNotification) GetEventReports() []UserPlaneEventReport {
	if o == nil {
		var ret []UserPlaneEventReport
		return ret
	}

	return o.EventReports
}

// GetEventReportsOk returns a tuple with the EventReports field value
// and a boolean to check if the value has been set.
func (o *UserPlaneEventNotification) GetEventReportsOk() ([]UserPlaneEventReport, bool) {
	if o == nil {
    return nil, false
	}
	return o.EventReports, true
}

// SetEventReports sets field value
func (o *UserPlaneEventNotification) SetEventReports(v []UserPlaneEventReport) {
	o.EventReports = v
}

func (o UserPlaneEventNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sessionId"] = o.SessionId
	}
	if true {
		toSerialize["eventReports"] = o.EventReports
	}
	return json.Marshal(toSerialize)
}

type NullableUserPlaneEventNotification struct {
	value *UserPlaneEventNotification
	isSet bool
}

func (v NullableUserPlaneEventNotification) Get() *UserPlaneEventNotification {
	return v.value
}

func (v *NullableUserPlaneEventNotification) Set(val *UserPlaneEventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPlaneEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPlaneEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPlaneEventNotification(val *UserPlaneEventNotification) *NullableUserPlaneEventNotification {
	return &NullableUserPlaneEventNotification{value: val, isSet: true}
}

func (v NullableUserPlaneEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPlaneEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


