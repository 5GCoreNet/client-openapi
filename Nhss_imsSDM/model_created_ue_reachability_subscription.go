/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"time"
)

// CreatedUeReachabilitySubscription Contains the response data returned by HSS after the subscription to  notifications of UE reachability for IP was created 
type CreatedUeReachabilitySubscription struct {
	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry"`
}

// NewCreatedUeReachabilitySubscription instantiates a new CreatedUeReachabilitySubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedUeReachabilitySubscription(expiry time.Time) *CreatedUeReachabilitySubscription {
	this := CreatedUeReachabilitySubscription{}
	this.Expiry = expiry
	return &this
}

// NewCreatedUeReachabilitySubscriptionWithDefaults instantiates a new CreatedUeReachabilitySubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedUeReachabilitySubscriptionWithDefaults() *CreatedUeReachabilitySubscription {
	this := CreatedUeReachabilitySubscription{}
	return &this
}

// GetExpiry returns the Expiry field value
func (o *CreatedUeReachabilitySubscription) GetExpiry() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value
// and a boolean to check if the value has been set.
func (o *CreatedUeReachabilitySubscription) GetExpiryOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Expiry, true
}

// SetExpiry sets field value
func (o *CreatedUeReachabilitySubscription) SetExpiry(v time.Time) {
	o.Expiry = v
}

func (o CreatedUeReachabilitySubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["expiry"] = o.Expiry
	}
	return json.Marshal(toSerialize)
}

type NullableCreatedUeReachabilitySubscription struct {
	value *CreatedUeReachabilitySubscription
	isSet bool
}

func (v NullableCreatedUeReachabilitySubscription) Get() *CreatedUeReachabilitySubscription {
	return v.value
}

func (v *NullableCreatedUeReachabilitySubscription) Set(val *CreatedUeReachabilitySubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedUeReachabilitySubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedUeReachabilitySubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedUeReachabilitySubscription(val *CreatedUeReachabilitySubscription) *NullableCreatedUeReachabilitySubscription {
	return &NullableCreatedUeReachabilitySubscription{value: val, isSet: true}
}

func (v NullableCreatedUeReachabilitySubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedUeReachabilitySubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


