/*
Nnef_Authentication

NEF Auth Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnef_Authentication

import (
	"encoding/json"
)

// AuthNotification UAV related notification
type AuthNotification struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi"`
	ServiceLevelId string `json:"serviceLevelId"`
	NotifyCorrId string `json:"notifyCorrId"`
	AuthMsg *RefToBinaryData `json:"authMsg,omitempty"`
	NotifType NotifType `json:"notifType"`
}

// NewAuthNotification instantiates a new AuthNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthNotification(gpsi string, serviceLevelId string, notifyCorrId string, notifType NotifType) *AuthNotification {
	this := AuthNotification{}
	this.Gpsi = gpsi
	this.ServiceLevelId = serviceLevelId
	this.NotifyCorrId = notifyCorrId
	this.NotifType = notifType
	return &this
}

// NewAuthNotificationWithDefaults instantiates a new AuthNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthNotificationWithDefaults() *AuthNotification {
	this := AuthNotification{}
	return &this
}

// GetGpsi returns the Gpsi field value
func (o *AuthNotification) GetGpsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value
// and a boolean to check if the value has been set.
func (o *AuthNotification) GetGpsiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Gpsi, true
}

// SetGpsi sets field value
func (o *AuthNotification) SetGpsi(v string) {
	o.Gpsi = v
}

// GetServiceLevelId returns the ServiceLevelId field value
func (o *AuthNotification) GetServiceLevelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceLevelId
}

// GetServiceLevelIdOk returns a tuple with the ServiceLevelId field value
// and a boolean to check if the value has been set.
func (o *AuthNotification) GetServiceLevelIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServiceLevelId, true
}

// SetServiceLevelId sets field value
func (o *AuthNotification) SetServiceLevelId(v string) {
	o.ServiceLevelId = v
}

// GetNotifyCorrId returns the NotifyCorrId field value
func (o *AuthNotification) GetNotifyCorrId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifyCorrId
}

// GetNotifyCorrIdOk returns a tuple with the NotifyCorrId field value
// and a boolean to check if the value has been set.
func (o *AuthNotification) GetNotifyCorrIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifyCorrId, true
}

// SetNotifyCorrId sets field value
func (o *AuthNotification) SetNotifyCorrId(v string) {
	o.NotifyCorrId = v
}

// GetAuthMsg returns the AuthMsg field value if set, zero value otherwise.
func (o *AuthNotification) GetAuthMsg() RefToBinaryData {
	if o == nil || isNil(o.AuthMsg) {
		var ret RefToBinaryData
		return ret
	}
	return *o.AuthMsg
}

// GetAuthMsgOk returns a tuple with the AuthMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthNotification) GetAuthMsgOk() (*RefToBinaryData, bool) {
	if o == nil || isNil(o.AuthMsg) {
    return nil, false
	}
	return o.AuthMsg, true
}

// HasAuthMsg returns a boolean if a field has been set.
func (o *AuthNotification) HasAuthMsg() bool {
	if o != nil && !isNil(o.AuthMsg) {
		return true
	}

	return false
}

// SetAuthMsg gets a reference to the given RefToBinaryData and assigns it to the AuthMsg field.
func (o *AuthNotification) SetAuthMsg(v RefToBinaryData) {
	o.AuthMsg = &v
}

// GetNotifType returns the NotifType field value
func (o *AuthNotification) GetNotifType() NotifType {
	if o == nil {
		var ret NotifType
		return ret
	}

	return o.NotifType
}

// GetNotifTypeOk returns a tuple with the NotifType field value
// and a boolean to check if the value has been set.
func (o *AuthNotification) GetNotifTypeOk() (*NotifType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifType, true
}

// SetNotifType sets field value
func (o *AuthNotification) SetNotifType(v NotifType) {
	o.NotifType = v
}

func (o AuthNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["gpsi"] = o.Gpsi
	}
	if true {
		toSerialize["serviceLevelId"] = o.ServiceLevelId
	}
	if true {
		toSerialize["notifyCorrId"] = o.NotifyCorrId
	}
	if !isNil(o.AuthMsg) {
		toSerialize["authMsg"] = o.AuthMsg
	}
	if true {
		toSerialize["notifType"] = o.NotifType
	}
	return json.Marshal(toSerialize)
}

type NullableAuthNotification struct {
	value *AuthNotification
	isSet bool
}

func (v NullableAuthNotification) Get() *AuthNotification {
	return v.value
}

func (v *NullableAuthNotification) Set(val *AuthNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthNotification(val *AuthNotification) *NullableAuthNotification {
	return &NullableAuthNotification{value: val, isSet: true}
}

func (v NullableAuthNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


