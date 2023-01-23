/*
Npcf_AMPolicyAuthorization Service API

PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_AMPolicyAuthorization

import (
	"encoding/json"
)

// AmEventsNotification Describes the notification about the events occurred within an Individual Application AM Context resource.
type AmEventsNotification struct {
	// Contains the AM Policy Events Subscription resource identifier related to the event notification.
	AppAmContextId *string `json:"appAmContextId,omitempty"`
	RepEvents []AmEventNotification `json:"repEvents"`
}

// NewAmEventsNotification instantiates a new AmEventsNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmEventsNotification(repEvents []AmEventNotification) *AmEventsNotification {
	this := AmEventsNotification{}
	this.RepEvents = repEvents
	return &this
}

// NewAmEventsNotificationWithDefaults instantiates a new AmEventsNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmEventsNotificationWithDefaults() *AmEventsNotification {
	this := AmEventsNotification{}
	return &this
}

// GetAppAmContextId returns the AppAmContextId field value if set, zero value otherwise.
func (o *AmEventsNotification) GetAppAmContextId() string {
	if o == nil || isNil(o.AppAmContextId) {
		var ret string
		return ret
	}
	return *o.AppAmContextId
}

// GetAppAmContextIdOk returns a tuple with the AppAmContextId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmEventsNotification) GetAppAmContextIdOk() (*string, bool) {
	if o == nil || isNil(o.AppAmContextId) {
    return nil, false
	}
	return o.AppAmContextId, true
}

// HasAppAmContextId returns a boolean if a field has been set.
func (o *AmEventsNotification) HasAppAmContextId() bool {
	if o != nil && !isNil(o.AppAmContextId) {
		return true
	}

	return false
}

// SetAppAmContextId gets a reference to the given string and assigns it to the AppAmContextId field.
func (o *AmEventsNotification) SetAppAmContextId(v string) {
	o.AppAmContextId = &v
}

// GetRepEvents returns the RepEvents field value
func (o *AmEventsNotification) GetRepEvents() []AmEventNotification {
	if o == nil {
		var ret []AmEventNotification
		return ret
	}

	return o.RepEvents
}

// GetRepEventsOk returns a tuple with the RepEvents field value
// and a boolean to check if the value has been set.
func (o *AmEventsNotification) GetRepEventsOk() ([]AmEventNotification, bool) {
	if o == nil {
    return nil, false
	}
	return o.RepEvents, true
}

// SetRepEvents sets field value
func (o *AmEventsNotification) SetRepEvents(v []AmEventNotification) {
	o.RepEvents = v
}

func (o AmEventsNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppAmContextId) {
		toSerialize["appAmContextId"] = o.AppAmContextId
	}
	if true {
		toSerialize["repEvents"] = o.RepEvents
	}
	return json.Marshal(toSerialize)
}

type NullableAmEventsNotification struct {
	value *AmEventsNotification
	isSet bool
}

func (v NullableAmEventsNotification) Get() *AmEventsNotification {
	return v.value
}

func (v *NullableAmEventsNotification) Set(val *AmEventsNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAmEventsNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAmEventsNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmEventsNotification(val *AmEventsNotification) *NullableAmEventsNotification {
	return &NullableAmEventsNotification{value: val, isSet: true}
}

func (v NullableAmEventsNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmEventsNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


