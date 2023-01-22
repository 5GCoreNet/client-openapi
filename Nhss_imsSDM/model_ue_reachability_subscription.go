/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nhss_imsSDM

import (
	"encoding/json"
	"time"
)

// UeReachabilitySubscription Contains the request parameters received by the HSS for a subscription to notifications of UE reachability for IP 
type UeReachabilitySubscription struct {
	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry"`
	// String providing an URI formatted according to RFC 3986.
	CallbackReference string `json:"callbackReference"`
}

// NewUeReachabilitySubscription instantiates a new UeReachabilitySubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeReachabilitySubscription(expiry time.Time, callbackReference string) *UeReachabilitySubscription {
	this := UeReachabilitySubscription{}
	this.Expiry = expiry
	this.CallbackReference = callbackReference
	return &this
}

// NewUeReachabilitySubscriptionWithDefaults instantiates a new UeReachabilitySubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeReachabilitySubscriptionWithDefaults() *UeReachabilitySubscription {
	this := UeReachabilitySubscription{}
	return &this
}

// GetExpiry returns the Expiry field value
func (o *UeReachabilitySubscription) GetExpiry() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value
// and a boolean to check if the value has been set.
func (o *UeReachabilitySubscription) GetExpiryOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Expiry, true
}

// SetExpiry sets field value
func (o *UeReachabilitySubscription) SetExpiry(v time.Time) {
	o.Expiry = v
}

// GetCallbackReference returns the CallbackReference field value
func (o *UeReachabilitySubscription) GetCallbackReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackReference
}

// GetCallbackReferenceOk returns a tuple with the CallbackReference field value
// and a boolean to check if the value has been set.
func (o *UeReachabilitySubscription) GetCallbackReferenceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CallbackReference, true
}

// SetCallbackReference sets field value
func (o *UeReachabilitySubscription) SetCallbackReference(v string) {
	o.CallbackReference = v
}

func (o UeReachabilitySubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["expiry"] = o.Expiry
	}
	if true {
		toSerialize["callbackReference"] = o.CallbackReference
	}
	return json.Marshal(toSerialize)
}

type NullableUeReachabilitySubscription struct {
	value *UeReachabilitySubscription
	isSet bool
}

func (v NullableUeReachabilitySubscription) Get() *UeReachabilitySubscription {
	return v.value
}

func (v *NullableUeReachabilitySubscription) Set(val *UeReachabilitySubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableUeReachabilitySubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableUeReachabilitySubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeReachabilitySubscription(val *UeReachabilitySubscription) *NullableUeReachabilitySubscription {
	return &NullableUeReachabilitySubscription{value: val, isSet: true}
}

func (v NullableUeReachabilitySubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeReachabilitySubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

