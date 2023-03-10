/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudsf_DataRepository

import (
	"encoding/json"
)

// NotificationInfo struct for NotificationInfo
type NotificationInfo struct {
	ExpiredSubscriptions []NotificationSubscription `json:"expiredSubscriptions"`
}

// NewNotificationInfo instantiates a new NotificationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationInfo(expiredSubscriptions []NotificationSubscription) *NotificationInfo {
	this := NotificationInfo{}
	this.ExpiredSubscriptions = expiredSubscriptions
	return &this
}

// NewNotificationInfoWithDefaults instantiates a new NotificationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationInfoWithDefaults() *NotificationInfo {
	this := NotificationInfo{}
	return &this
}

// GetExpiredSubscriptions returns the ExpiredSubscriptions field value
func (o *NotificationInfo) GetExpiredSubscriptions() []NotificationSubscription {
	if o == nil {
		var ret []NotificationSubscription
		return ret
	}

	return o.ExpiredSubscriptions
}

// GetExpiredSubscriptionsOk returns a tuple with the ExpiredSubscriptions field value
// and a boolean to check if the value has been set.
func (o *NotificationInfo) GetExpiredSubscriptionsOk() ([]NotificationSubscription, bool) {
	if o == nil {
    return nil, false
	}
	return o.ExpiredSubscriptions, true
}

// SetExpiredSubscriptions sets field value
func (o *NotificationInfo) SetExpiredSubscriptions(v []NotificationSubscription) {
	o.ExpiredSubscriptions = v
}

func (o NotificationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["expiredSubscriptions"] = o.ExpiredSubscriptions
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationInfo struct {
	value *NotificationInfo
	isSet bool
}

func (v NullableNotificationInfo) Get() *NotificationInfo {
	return v.value
}

func (v *NullableNotificationInfo) Set(val *NotificationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationInfo(val *NotificationInfo) *NullableNotificationInfo {
	return &NullableNotificationInfo{value: val, isSet: true}
}

func (v NullableNotificationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


