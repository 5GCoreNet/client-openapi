/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// EventsSubs Represents an event to be subscribed and the related event filter information.
type EventsSubs struct {
	Event AfEvent `json:"event"`
	EventFilter EventFilter `json:"eventFilter"`
}

// NewEventsSubs instantiates a new EventsSubs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsSubs(event AfEvent, eventFilter EventFilter) *EventsSubs {
	this := EventsSubs{}
	this.Event = event
	this.EventFilter = eventFilter
	return &this
}

// NewEventsSubsWithDefaults instantiates a new EventsSubs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsSubsWithDefaults() *EventsSubs {
	this := EventsSubs{}
	return &this
}

// GetEvent returns the Event field value
func (o *EventsSubs) GetEvent() AfEvent {
	if o == nil {
		var ret AfEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *EventsSubs) GetEventOk() (*AfEvent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *EventsSubs) SetEvent(v AfEvent) {
	o.Event = v
}

// GetEventFilter returns the EventFilter field value
func (o *EventsSubs) GetEventFilter() EventFilter {
	if o == nil {
		var ret EventFilter
		return ret
	}

	return o.EventFilter
}

// GetEventFilterOk returns a tuple with the EventFilter field value
// and a boolean to check if the value has been set.
func (o *EventsSubs) GetEventFilterOk() (*EventFilter, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventFilter, true
}

// SetEventFilter sets field value
func (o *EventsSubs) SetEventFilter(v EventFilter) {
	o.EventFilter = v
}

func (o EventsSubs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["eventFilter"] = o.EventFilter
	}
	return json.Marshal(toSerialize)
}

type NullableEventsSubs struct {
	value *EventsSubs
	isSet bool
}

func (v NullableEventsSubs) Get() *EventsSubs {
	return v.value
}

func (v *NullableEventsSubs) Set(val *EventsSubs) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsSubs) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsSubs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsSubs(val *EventsSubs) *NullableEventsSubs {
	return &NullableEventsSubs{value: val, isSet: true}
}

func (v NullableEventsSubs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsSubs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


