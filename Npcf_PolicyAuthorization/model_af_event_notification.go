/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// AfEventNotification Describes the event information delivered in the notification.
type AfEventNotification struct {
	Event AfEvent `json:"event"`
	Flows []Flows `json:"flows,omitempty"`
}

// NewAfEventNotification instantiates a new AfEventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfEventNotification(event AfEvent) *AfEventNotification {
	this := AfEventNotification{}
	this.Event = event
	return &this
}

// NewAfEventNotificationWithDefaults instantiates a new AfEventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfEventNotificationWithDefaults() *AfEventNotification {
	this := AfEventNotification{}
	return &this
}

// GetEvent returns the Event field value
func (o *AfEventNotification) GetEvent() AfEvent {
	if o == nil {
		var ret AfEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetEventOk() (*AfEvent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AfEventNotification) SetEvent(v AfEvent) {
	o.Event = v
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *AfEventNotification) GetFlows() []Flows {
	if o == nil || isNil(o.Flows) {
		var ret []Flows
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetFlowsOk() ([]Flows, bool) {
	if o == nil || isNil(o.Flows) {
    return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *AfEventNotification) HasFlows() bool {
	if o != nil && !isNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []Flows and assigns it to the Flows field.
func (o *AfEventNotification) SetFlows(v []Flows) {
	o.Flows = v
}

func (o AfEventNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event"] = o.Event
	}
	if !isNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	return json.Marshal(toSerialize)
}

type NullableAfEventNotification struct {
	value *AfEventNotification
	isSet bool
}

func (v NullableAfEventNotification) Get() *AfEventNotification {
	return v.value
}

func (v *NullableAfEventNotification) Set(val *AfEventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAfEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAfEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfEventNotification(val *AfEventNotification) *NullableAfEventNotification {
	return &NullableAfEventNotification{value: val, isSet: true}
}

func (v NullableAfEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


