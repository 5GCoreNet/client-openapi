/*
Eees_ACREvents

API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_ACREvents

import (
	"encoding/json"
	"time"
)

// ACREventsSubscriptionPatch An individual ACR events subscription resource to be updated.
type ACREventsSubscriptionPatch struct {
	// string with format \"date-time\" as defined in OpenAPI.
	ExpTime *time.Time `json:"expTime,omitempty"`
	// The list of application identifiers of the EASs.
	EasIds []string `json:"easIds,omitempty"`
	EventIds *ACREventIDs `json:"eventIds,omitempty"`
	// string providing an URI formatted according to IETF RFC 3986.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
}

// NewACREventsSubscriptionPatch instantiates a new ACREventsSubscriptionPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACREventsSubscriptionPatch() *ACREventsSubscriptionPatch {
	this := ACREventsSubscriptionPatch{}
	return &this
}

// NewACREventsSubscriptionPatchWithDefaults instantiates a new ACREventsSubscriptionPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACREventsSubscriptionPatchWithDefaults() *ACREventsSubscriptionPatch {
	this := ACREventsSubscriptionPatch{}
	return &this
}

// GetExpTime returns the ExpTime field value if set, zero value otherwise.
func (o *ACREventsSubscriptionPatch) GetExpTime() time.Time {
	if o == nil || isNil(o.ExpTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpTime
}

// GetExpTimeOk returns a tuple with the ExpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACREventsSubscriptionPatch) GetExpTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpTime) {
    return nil, false
	}
	return o.ExpTime, true
}

// HasExpTime returns a boolean if a field has been set.
func (o *ACREventsSubscriptionPatch) HasExpTime() bool {
	if o != nil && !isNil(o.ExpTime) {
		return true
	}

	return false
}

// SetExpTime gets a reference to the given time.Time and assigns it to the ExpTime field.
func (o *ACREventsSubscriptionPatch) SetExpTime(v time.Time) {
	o.ExpTime = &v
}

// GetEasIds returns the EasIds field value if set, zero value otherwise.
func (o *ACREventsSubscriptionPatch) GetEasIds() []string {
	if o == nil || isNil(o.EasIds) {
		var ret []string
		return ret
	}
	return o.EasIds
}

// GetEasIdsOk returns a tuple with the EasIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACREventsSubscriptionPatch) GetEasIdsOk() ([]string, bool) {
	if o == nil || isNil(o.EasIds) {
    return nil, false
	}
	return o.EasIds, true
}

// HasEasIds returns a boolean if a field has been set.
func (o *ACREventsSubscriptionPatch) HasEasIds() bool {
	if o != nil && !isNil(o.EasIds) {
		return true
	}

	return false
}

// SetEasIds gets a reference to the given []string and assigns it to the EasIds field.
func (o *ACREventsSubscriptionPatch) SetEasIds(v []string) {
	o.EasIds = v
}

// GetEventIds returns the EventIds field value if set, zero value otherwise.
func (o *ACREventsSubscriptionPatch) GetEventIds() ACREventIDs {
	if o == nil || isNil(o.EventIds) {
		var ret ACREventIDs
		return ret
	}
	return *o.EventIds
}

// GetEventIdsOk returns a tuple with the EventIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACREventsSubscriptionPatch) GetEventIdsOk() (*ACREventIDs, bool) {
	if o == nil || isNil(o.EventIds) {
    return nil, false
	}
	return o.EventIds, true
}

// HasEventIds returns a boolean if a field has been set.
func (o *ACREventsSubscriptionPatch) HasEventIds() bool {
	if o != nil && !isNil(o.EventIds) {
		return true
	}

	return false
}

// SetEventIds gets a reference to the given ACREventIDs and assigns it to the EventIds field.
func (o *ACREventsSubscriptionPatch) SetEventIds(v ACREventIDs) {
	o.EventIds = &v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *ACREventsSubscriptionPatch) GetNotificationDestination() string {
	if o == nil || isNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACREventsSubscriptionPatch) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || isNil(o.NotificationDestination) {
    return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *ACREventsSubscriptionPatch) HasNotificationDestination() bool {
	if o != nil && !isNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *ACREventsSubscriptionPatch) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

func (o ACREventsSubscriptionPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExpTime) {
		toSerialize["expTime"] = o.ExpTime
	}
	if !isNil(o.EasIds) {
		toSerialize["easIds"] = o.EasIds
	}
	if !isNil(o.EventIds) {
		toSerialize["eventIds"] = o.EventIds
	}
	if !isNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	return json.Marshal(toSerialize)
}

type NullableACREventsSubscriptionPatch struct {
	value *ACREventsSubscriptionPatch
	isSet bool
}

func (v NullableACREventsSubscriptionPatch) Get() *ACREventsSubscriptionPatch {
	return v.value
}

func (v *NullableACREventsSubscriptionPatch) Set(val *ACREventsSubscriptionPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableACREventsSubscriptionPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableACREventsSubscriptionPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACREventsSubscriptionPatch(val *ACREventsSubscriptionPatch) *NullableACREventsSubscriptionPatch {
	return &NullableACREventsSubscriptionPatch{value: val, isSet: true}
}

func (v NullableACREventsSubscriptionPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACREventsSubscriptionPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


