/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_EE

import (
	"encoding/json"
)

// MonitoringEvent struct for MonitoringEvent
type MonitoringEvent struct {
	EventType EventType `json:"eventType"`
	RevokedCause *RevokedCause `json:"revokedCause,omitempty"`
}

// NewMonitoringEvent instantiates a new MonitoringEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringEvent(eventType EventType) *MonitoringEvent {
	this := MonitoringEvent{}
	this.EventType = eventType
	return &this
}

// NewMonitoringEventWithDefaults instantiates a new MonitoringEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringEventWithDefaults() *MonitoringEvent {
	this := MonitoringEvent{}
	return &this
}

// GetEventType returns the EventType field value
func (o *MonitoringEvent) GetEventType() EventType {
	if o == nil {
		var ret EventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *MonitoringEvent) GetEventTypeOk() (*EventType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *MonitoringEvent) SetEventType(v EventType) {
	o.EventType = v
}

// GetRevokedCause returns the RevokedCause field value if set, zero value otherwise.
func (o *MonitoringEvent) GetRevokedCause() RevokedCause {
	if o == nil || isNil(o.RevokedCause) {
		var ret RevokedCause
		return ret
	}
	return *o.RevokedCause
}

// GetRevokedCauseOk returns a tuple with the RevokedCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringEvent) GetRevokedCauseOk() (*RevokedCause, bool) {
	if o == nil || isNil(o.RevokedCause) {
    return nil, false
	}
	return o.RevokedCause, true
}

// HasRevokedCause returns a boolean if a field has been set.
func (o *MonitoringEvent) HasRevokedCause() bool {
	if o != nil && !isNil(o.RevokedCause) {
		return true
	}

	return false
}

// SetRevokedCause gets a reference to the given RevokedCause and assigns it to the RevokedCause field.
func (o *MonitoringEvent) SetRevokedCause(v RevokedCause) {
	o.RevokedCause = &v
}

func (o MonitoringEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if !isNil(o.RevokedCause) {
		toSerialize["revokedCause"] = o.RevokedCause
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringEvent struct {
	value *MonitoringEvent
	isSet bool
}

func (v NullableMonitoringEvent) Get() *MonitoringEvent {
	return v.value
}

func (v *NullableMonitoringEvent) Set(val *MonitoringEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringEvent(val *MonitoringEvent) *NullableMonitoringEvent {
	return &NullableMonitoringEvent{value: val, isSet: true}
}

func (v NullableMonitoringEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


