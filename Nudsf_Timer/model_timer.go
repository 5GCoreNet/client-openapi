/*
Nudsf_Timer

Nudsf Timer Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudsf_Timer

import (
	"encoding/json"
	"time"
)

// Timer Represents a timer.
type Timer struct {
	// Represents the identifier of a timer.
	TimerId *string `json:"timerId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expires time.Time `json:"expires"`
	// A map (list of key-value pairs where a tagName of type string serves as key) of tagValue lists
	MetaTags *map[string][]string `json:"metaTags,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	CallbackReference *string `json:"callbackReference,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	DeleteAfter *int32 `json:"deleteAfter,omitempty"`
}

// NewTimer instantiates a new Timer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimer(expires time.Time) *Timer {
	this := Timer{}
	this.Expires = expires
	return &this
}

// NewTimerWithDefaults instantiates a new Timer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimerWithDefaults() *Timer {
	this := Timer{}
	return &this
}

// GetTimerId returns the TimerId field value if set, zero value otherwise.
func (o *Timer) GetTimerId() string {
	if o == nil || isNil(o.TimerId) {
		var ret string
		return ret
	}
	return *o.TimerId
}

// GetTimerIdOk returns a tuple with the TimerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timer) GetTimerIdOk() (*string, bool) {
	if o == nil || isNil(o.TimerId) {
    return nil, false
	}
	return o.TimerId, true
}

// HasTimerId returns a boolean if a field has been set.
func (o *Timer) HasTimerId() bool {
	if o != nil && !isNil(o.TimerId) {
		return true
	}

	return false
}

// SetTimerId gets a reference to the given string and assigns it to the TimerId field.
func (o *Timer) SetTimerId(v string) {
	o.TimerId = &v
}

// GetExpires returns the Expires field value
func (o *Timer) GetExpires() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value
// and a boolean to check if the value has been set.
func (o *Timer) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Expires, true
}

// SetExpires sets field value
func (o *Timer) SetExpires(v time.Time) {
	o.Expires = v
}

// GetMetaTags returns the MetaTags field value if set, zero value otherwise.
func (o *Timer) GetMetaTags() map[string][]string {
	if o == nil || isNil(o.MetaTags) {
		var ret map[string][]string
		return ret
	}
	return *o.MetaTags
}

// GetMetaTagsOk returns a tuple with the MetaTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timer) GetMetaTagsOk() (*map[string][]string, bool) {
	if o == nil || isNil(o.MetaTags) {
    return nil, false
	}
	return o.MetaTags, true
}

// HasMetaTags returns a boolean if a field has been set.
func (o *Timer) HasMetaTags() bool {
	if o != nil && !isNil(o.MetaTags) {
		return true
	}

	return false
}

// SetMetaTags gets a reference to the given map[string][]string and assigns it to the MetaTags field.
func (o *Timer) SetMetaTags(v map[string][]string) {
	o.MetaTags = &v
}

// GetCallbackReference returns the CallbackReference field value if set, zero value otherwise.
func (o *Timer) GetCallbackReference() string {
	if o == nil || isNil(o.CallbackReference) {
		var ret string
		return ret
	}
	return *o.CallbackReference
}

// GetCallbackReferenceOk returns a tuple with the CallbackReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timer) GetCallbackReferenceOk() (*string, bool) {
	if o == nil || isNil(o.CallbackReference) {
    return nil, false
	}
	return o.CallbackReference, true
}

// HasCallbackReference returns a boolean if a field has been set.
func (o *Timer) HasCallbackReference() bool {
	if o != nil && !isNil(o.CallbackReference) {
		return true
	}

	return false
}

// SetCallbackReference gets a reference to the given string and assigns it to the CallbackReference field.
func (o *Timer) SetCallbackReference(v string) {
	o.CallbackReference = &v
}

// GetDeleteAfter returns the DeleteAfter field value if set, zero value otherwise.
func (o *Timer) GetDeleteAfter() int32 {
	if o == nil || isNil(o.DeleteAfter) {
		var ret int32
		return ret
	}
	return *o.DeleteAfter
}

// GetDeleteAfterOk returns a tuple with the DeleteAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timer) GetDeleteAfterOk() (*int32, bool) {
	if o == nil || isNil(o.DeleteAfter) {
    return nil, false
	}
	return o.DeleteAfter, true
}

// HasDeleteAfter returns a boolean if a field has been set.
func (o *Timer) HasDeleteAfter() bool {
	if o != nil && !isNil(o.DeleteAfter) {
		return true
	}

	return false
}

// SetDeleteAfter gets a reference to the given int32 and assigns it to the DeleteAfter field.
func (o *Timer) SetDeleteAfter(v int32) {
	o.DeleteAfter = &v
}

func (o Timer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TimerId) {
		toSerialize["timerId"] = o.TimerId
	}
	if true {
		toSerialize["expires"] = o.Expires
	}
	if !isNil(o.MetaTags) {
		toSerialize["metaTags"] = o.MetaTags
	}
	if !isNil(o.CallbackReference) {
		toSerialize["callbackReference"] = o.CallbackReference
	}
	if !isNil(o.DeleteAfter) {
		toSerialize["deleteAfter"] = o.DeleteAfter
	}
	return json.Marshal(toSerialize)
}

type NullableTimer struct {
	value *Timer
	isSet bool
}

func (v NullableTimer) Get() *Timer {
	return v.value
}

func (v *NullableTimer) Set(val *Timer) {
	v.value = val
	v.isSet = true
}

func (v NullableTimer) IsSet() bool {
	return v.isSet
}

func (v *NullableTimer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimer(val *Timer) *NullableTimer {
	return &NullableTimer{value: val, isSet: true}
}

func (v NullableTimer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


