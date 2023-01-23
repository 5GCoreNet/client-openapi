/*
NSSF NSSAI Availability

NSSF NSSAI Availability Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssf_NSSAIAvailability

import (
	"encoding/json"
)

// NssfEventNotification This contains the notification for created event subscription
type NssfEventNotification struct {
	SubscriptionId string `json:"subscriptionId"`
	AuthorizedNssaiAvailabilityData []AuthorizedNssaiAvailabilityData `json:"authorizedNssaiAvailabilityData"`
}

// NewNssfEventNotification instantiates a new NssfEventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNssfEventNotification(subscriptionId string, authorizedNssaiAvailabilityData []AuthorizedNssaiAvailabilityData) *NssfEventNotification {
	this := NssfEventNotification{}
	this.SubscriptionId = subscriptionId
	this.AuthorizedNssaiAvailabilityData = authorizedNssaiAvailabilityData
	return &this
}

// NewNssfEventNotificationWithDefaults instantiates a new NssfEventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNssfEventNotificationWithDefaults() *NssfEventNotification {
	this := NssfEventNotification{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *NssfEventNotification) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *NssfEventNotification) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *NssfEventNotification) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetAuthorizedNssaiAvailabilityData returns the AuthorizedNssaiAvailabilityData field value
func (o *NssfEventNotification) GetAuthorizedNssaiAvailabilityData() []AuthorizedNssaiAvailabilityData {
	if o == nil {
		var ret []AuthorizedNssaiAvailabilityData
		return ret
	}

	return o.AuthorizedNssaiAvailabilityData
}

// GetAuthorizedNssaiAvailabilityDataOk returns a tuple with the AuthorizedNssaiAvailabilityData field value
// and a boolean to check if the value has been set.
func (o *NssfEventNotification) GetAuthorizedNssaiAvailabilityDataOk() ([]AuthorizedNssaiAvailabilityData, bool) {
	if o == nil {
    return nil, false
	}
	return o.AuthorizedNssaiAvailabilityData, true
}

// SetAuthorizedNssaiAvailabilityData sets field value
func (o *NssfEventNotification) SetAuthorizedNssaiAvailabilityData(v []AuthorizedNssaiAvailabilityData) {
	o.AuthorizedNssaiAvailabilityData = v
}

func (o NssfEventNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if true {
		toSerialize["authorizedNssaiAvailabilityData"] = o.AuthorizedNssaiAvailabilityData
	}
	return json.Marshal(toSerialize)
}

type NullableNssfEventNotification struct {
	value *NssfEventNotification
	isSet bool
}

func (v NullableNssfEventNotification) Get() *NssfEventNotification {
	return v.value
}

func (v *NullableNssfEventNotification) Set(val *NssfEventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableNssfEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNssfEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNssfEventNotification(val *NssfEventNotification) *NullableNssfEventNotification {
	return &NullableNssfEventNotification{value: val, isSet: true}
}

func (v NullableNssfEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNssfEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


