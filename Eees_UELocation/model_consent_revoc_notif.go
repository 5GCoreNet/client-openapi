/*
EES UE Location Information_API

API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_UELocation

import (
	"encoding/json"
)

// ConsentRevocNotif Represents the user consent revocation information conveyed in a user consent revocation notification. 
type ConsentRevocNotif struct {
	SubscriptionId string `json:"subscriptionId"`
	ConsentsRevoked []ConsentRevoked `json:"consentsRevoked"`
}

// NewConsentRevocNotif instantiates a new ConsentRevocNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentRevocNotif(subscriptionId string, consentsRevoked []ConsentRevoked) *ConsentRevocNotif {
	this := ConsentRevocNotif{}
	this.SubscriptionId = subscriptionId
	this.ConsentsRevoked = consentsRevoked
	return &this
}

// NewConsentRevocNotifWithDefaults instantiates a new ConsentRevocNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentRevocNotifWithDefaults() *ConsentRevocNotif {
	this := ConsentRevocNotif{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *ConsentRevocNotif) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *ConsentRevocNotif) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *ConsentRevocNotif) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetConsentsRevoked returns the ConsentsRevoked field value
func (o *ConsentRevocNotif) GetConsentsRevoked() []ConsentRevoked {
	if o == nil {
		var ret []ConsentRevoked
		return ret
	}

	return o.ConsentsRevoked
}

// GetConsentsRevokedOk returns a tuple with the ConsentsRevoked field value
// and a boolean to check if the value has been set.
func (o *ConsentRevocNotif) GetConsentsRevokedOk() ([]ConsentRevoked, bool) {
	if o == nil {
    return nil, false
	}
	return o.ConsentsRevoked, true
}

// SetConsentsRevoked sets field value
func (o *ConsentRevocNotif) SetConsentsRevoked(v []ConsentRevoked) {
	o.ConsentsRevoked = v
}

func (o ConsentRevocNotif) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if true {
		toSerialize["consentsRevoked"] = o.ConsentsRevoked
	}
	return json.Marshal(toSerialize)
}

type NullableConsentRevocNotif struct {
	value *ConsentRevocNotif
	isSet bool
}

func (v NullableConsentRevocNotif) Get() *ConsentRevocNotif {
	return v.value
}

func (v *NullableConsentRevocNotif) Set(val *ConsentRevocNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentRevocNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentRevocNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentRevocNotif(val *ConsentRevocNotif) *NullableConsentRevocNotif {
	return &NullableConsentRevocNotif{value: val, isSet: true}
}

func (v NullableConsentRevocNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentRevocNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


