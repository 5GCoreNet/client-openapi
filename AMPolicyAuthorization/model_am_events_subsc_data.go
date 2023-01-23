/*
3gpp-am-policyauthorization

API for AM policy authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AMPolicyAuthorization

import (
	"encoding/json"
)

// AmEventsSubscData It represents the AM Policy Events Subscription subresource and identifies the events the application subscribes to.
type AmEventsSubscData struct {
	// String providing an URI formatted according to RFC 3986.
	EventNotifUri string `json:"eventNotifUri"`
	Events []AmEventData `json:"events,omitempty"`
}

// NewAmEventsSubscData instantiates a new AmEventsSubscData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmEventsSubscData(eventNotifUri string) *AmEventsSubscData {
	this := AmEventsSubscData{}
	this.EventNotifUri = eventNotifUri
	return &this
}

// NewAmEventsSubscDataWithDefaults instantiates a new AmEventsSubscData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmEventsSubscDataWithDefaults() *AmEventsSubscData {
	this := AmEventsSubscData{}
	return &this
}

// GetEventNotifUri returns the EventNotifUri field value
func (o *AmEventsSubscData) GetEventNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventNotifUri
}

// GetEventNotifUriOk returns a tuple with the EventNotifUri field value
// and a boolean to check if the value has been set.
func (o *AmEventsSubscData) GetEventNotifUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventNotifUri, true
}

// SetEventNotifUri sets field value
func (o *AmEventsSubscData) SetEventNotifUri(v string) {
	o.EventNotifUri = v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *AmEventsSubscData) GetEvents() []AmEventData {
	if o == nil || isNil(o.Events) {
		var ret []AmEventData
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmEventsSubscData) GetEventsOk() ([]AmEventData, bool) {
	if o == nil || isNil(o.Events) {
    return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *AmEventsSubscData) HasEvents() bool {
	if o != nil && !isNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []AmEventData and assigns it to the Events field.
func (o *AmEventsSubscData) SetEvents(v []AmEventData) {
	o.Events = v
}

func (o AmEventsSubscData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventNotifUri"] = o.EventNotifUri
	}
	if !isNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableAmEventsSubscData struct {
	value *AmEventsSubscData
	isSet bool
}

func (v NullableAmEventsSubscData) Get() *AmEventsSubscData {
	return v.value
}

func (v *NullableAmEventsSubscData) Set(val *AmEventsSubscData) {
	v.value = val
	v.isSet = true
}

func (v NullableAmEventsSubscData) IsSet() bool {
	return v.isSet
}

func (v *NullableAmEventsSubscData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmEventsSubscData(val *AmEventsSubscData) *NullableAmEventsSubscData {
	return &NullableAmEventsSubscData{value: val, isSet: true}
}

func (v NullableAmEventsSubscData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmEventsSubscData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


