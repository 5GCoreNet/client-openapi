/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MBSUserDataIngestSession

import (
	"encoding/json"
)

// MBSUserDataIngStatSubscPatch Represents the requested modifications to an MBS User Data Ingest Session Status  Subscription. 
type MBSUserDataIngStatSubscPatch struct {
	EventSubscs []SubscribedEvent `json:"eventSubscs,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NotifUri *string `json:"notifUri,omitempty"`
}

// NewMBSUserDataIngStatSubscPatch instantiates a new MBSUserDataIngStatSubscPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMBSUserDataIngStatSubscPatch() *MBSUserDataIngStatSubscPatch {
	this := MBSUserDataIngStatSubscPatch{}
	return &this
}

// NewMBSUserDataIngStatSubscPatchWithDefaults instantiates a new MBSUserDataIngStatSubscPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMBSUserDataIngStatSubscPatchWithDefaults() *MBSUserDataIngStatSubscPatch {
	this := MBSUserDataIngStatSubscPatch{}
	return &this
}

// GetEventSubscs returns the EventSubscs field value if set, zero value otherwise.
func (o *MBSUserDataIngStatSubscPatch) GetEventSubscs() []SubscribedEvent {
	if o == nil || isNil(o.EventSubscs) {
		var ret []SubscribedEvent
		return ret
	}
	return o.EventSubscs
}

// GetEventSubscsOk returns a tuple with the EventSubscs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSUserDataIngStatSubscPatch) GetEventSubscsOk() ([]SubscribedEvent, bool) {
	if o == nil || isNil(o.EventSubscs) {
    return nil, false
	}
	return o.EventSubscs, true
}

// HasEventSubscs returns a boolean if a field has been set.
func (o *MBSUserDataIngStatSubscPatch) HasEventSubscs() bool {
	if o != nil && !isNil(o.EventSubscs) {
		return true
	}

	return false
}

// SetEventSubscs gets a reference to the given []SubscribedEvent and assigns it to the EventSubscs field.
func (o *MBSUserDataIngStatSubscPatch) SetEventSubscs(v []SubscribedEvent) {
	o.EventSubscs = v
}

// GetNotifUri returns the NotifUri field value if set, zero value otherwise.
func (o *MBSUserDataIngStatSubscPatch) GetNotifUri() string {
	if o == nil || isNil(o.NotifUri) {
		var ret string
		return ret
	}
	return *o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSUserDataIngStatSubscPatch) GetNotifUriOk() (*string, bool) {
	if o == nil || isNil(o.NotifUri) {
    return nil, false
	}
	return o.NotifUri, true
}

// HasNotifUri returns a boolean if a field has been set.
func (o *MBSUserDataIngStatSubscPatch) HasNotifUri() bool {
	if o != nil && !isNil(o.NotifUri) {
		return true
	}

	return false
}

// SetNotifUri gets a reference to the given string and assigns it to the NotifUri field.
func (o *MBSUserDataIngStatSubscPatch) SetNotifUri(v string) {
	o.NotifUri = &v
}

func (o MBSUserDataIngStatSubscPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EventSubscs) {
		toSerialize["eventSubscs"] = o.EventSubscs
	}
	if !isNil(o.NotifUri) {
		toSerialize["notifUri"] = o.NotifUri
	}
	return json.Marshal(toSerialize)
}

type NullableMBSUserDataIngStatSubscPatch struct {
	value *MBSUserDataIngStatSubscPatch
	isSet bool
}

func (v NullableMBSUserDataIngStatSubscPatch) Get() *MBSUserDataIngStatSubscPatch {
	return v.value
}

func (v *NullableMBSUserDataIngStatSubscPatch) Set(val *MBSUserDataIngStatSubscPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableMBSUserDataIngStatSubscPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableMBSUserDataIngStatSubscPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMBSUserDataIngStatSubscPatch(val *MBSUserDataIngStatSubscPatch) *NullableMBSUserDataIngStatSubscPatch {
	return &NullableMBSUserDataIngStatSubscPatch{value: val, isSet: true}
}

func (v NullableMBSUserDataIngStatSubscPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMBSUserDataIngStatSubscPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


