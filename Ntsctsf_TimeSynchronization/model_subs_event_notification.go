/*
Ntsctsf_TimeSynchronization Service API

TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ntsctsf_TimeSynchronization

import (
	"encoding/json"
)

// SubsEventNotification Contains the notification of capability of time synchronization for a list of UEs. 
type SubsEventNotification struct {
	Event *SubscribedEvent `json:"event,omitempty"`
	TimeSyncCapas []TimeSyncCapability `json:"timeSyncCapas,omitempty"`
}

// NewSubsEventNotification instantiates a new SubsEventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubsEventNotification() *SubsEventNotification {
	this := SubsEventNotification{}
	return &this
}

// NewSubsEventNotificationWithDefaults instantiates a new SubsEventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubsEventNotificationWithDefaults() *SubsEventNotification {
	this := SubsEventNotification{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *SubsEventNotification) GetEvent() SubscribedEvent {
	if o == nil || isNil(o.Event) {
		var ret SubscribedEvent
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubsEventNotification) GetEventOk() (*SubscribedEvent, bool) {
	if o == nil || isNil(o.Event) {
    return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *SubsEventNotification) HasEvent() bool {
	if o != nil && !isNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given SubscribedEvent and assigns it to the Event field.
func (o *SubsEventNotification) SetEvent(v SubscribedEvent) {
	o.Event = &v
}

// GetTimeSyncCapas returns the TimeSyncCapas field value if set, zero value otherwise.
func (o *SubsEventNotification) GetTimeSyncCapas() []TimeSyncCapability {
	if o == nil || isNil(o.TimeSyncCapas) {
		var ret []TimeSyncCapability
		return ret
	}
	return o.TimeSyncCapas
}

// GetTimeSyncCapasOk returns a tuple with the TimeSyncCapas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubsEventNotification) GetTimeSyncCapasOk() ([]TimeSyncCapability, bool) {
	if o == nil || isNil(o.TimeSyncCapas) {
    return nil, false
	}
	return o.TimeSyncCapas, true
}

// HasTimeSyncCapas returns a boolean if a field has been set.
func (o *SubsEventNotification) HasTimeSyncCapas() bool {
	if o != nil && !isNil(o.TimeSyncCapas) {
		return true
	}

	return false
}

// SetTimeSyncCapas gets a reference to the given []TimeSyncCapability and assigns it to the TimeSyncCapas field.
func (o *SubsEventNotification) SetTimeSyncCapas(v []TimeSyncCapability) {
	o.TimeSyncCapas = v
}

func (o SubsEventNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !isNil(o.TimeSyncCapas) {
		toSerialize["timeSyncCapas"] = o.TimeSyncCapas
	}
	return json.Marshal(toSerialize)
}

type NullableSubsEventNotification struct {
	value *SubsEventNotification
	isSet bool
}

func (v NullableSubsEventNotification) Get() *SubsEventNotification {
	return v.value
}

func (v *NullableSubsEventNotification) Set(val *SubsEventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableSubsEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableSubsEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubsEventNotification(val *SubsEventNotification) *NullableSubsEventNotification {
	return &NullableSubsEventNotification{value: val, isSet: true}
}

func (v NullableSubsEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubsEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


