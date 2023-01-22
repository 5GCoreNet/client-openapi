/*
Naf_Authentication

AF Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Naf_Authentication

import (
	"encoding/json"
)

// ReauthRevokeNotify UAV related notification
type ReauthRevokeNotify struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi"`
	ServiceLevelId string `json:"serviceLevelId"`
	NotifyCorrId *string `json:"notifyCorrId,omitempty"`
	AuthMsg *string `json:"authMsg,omitempty"`
	NotifyType NotifyType `json:"notifyType"`
	IpAddr *IpAddr `json:"ipAddr,omitempty"`
}

// NewReauthRevokeNotify instantiates a new ReauthRevokeNotify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReauthRevokeNotify(gpsi string, serviceLevelId string, notifyType NotifyType) *ReauthRevokeNotify {
	this := ReauthRevokeNotify{}
	this.Gpsi = gpsi
	this.ServiceLevelId = serviceLevelId
	this.NotifyType = notifyType
	return &this
}

// NewReauthRevokeNotifyWithDefaults instantiates a new ReauthRevokeNotify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReauthRevokeNotifyWithDefaults() *ReauthRevokeNotify {
	this := ReauthRevokeNotify{}
	return &this
}

// GetGpsi returns the Gpsi field value
func (o *ReauthRevokeNotify) GetGpsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value
// and a boolean to check if the value has been set.
func (o *ReauthRevokeNotify) GetGpsiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Gpsi, true
}

// SetGpsi sets field value
func (o *ReauthRevokeNotify) SetGpsi(v string) {
	o.Gpsi = v
}

// GetServiceLevelId returns the ServiceLevelId field value
func (o *ReauthRevokeNotify) GetServiceLevelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceLevelId
}

// GetServiceLevelIdOk returns a tuple with the ServiceLevelId field value
// and a boolean to check if the value has been set.
func (o *ReauthRevokeNotify) GetServiceLevelIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServiceLevelId, true
}

// SetServiceLevelId sets field value
func (o *ReauthRevokeNotify) SetServiceLevelId(v string) {
	o.ServiceLevelId = v
}

// GetNotifyCorrId returns the NotifyCorrId field value if set, zero value otherwise.
func (o *ReauthRevokeNotify) GetNotifyCorrId() string {
	if o == nil || isNil(o.NotifyCorrId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrId
}

// GetNotifyCorrIdOk returns a tuple with the NotifyCorrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReauthRevokeNotify) GetNotifyCorrIdOk() (*string, bool) {
	if o == nil || isNil(o.NotifyCorrId) {
    return nil, false
	}
	return o.NotifyCorrId, true
}

// HasNotifyCorrId returns a boolean if a field has been set.
func (o *ReauthRevokeNotify) HasNotifyCorrId() bool {
	if o != nil && !isNil(o.NotifyCorrId) {
		return true
	}

	return false
}

// SetNotifyCorrId gets a reference to the given string and assigns it to the NotifyCorrId field.
func (o *ReauthRevokeNotify) SetNotifyCorrId(v string) {
	o.NotifyCorrId = &v
}

// GetAuthMsg returns the AuthMsg field value if set, zero value otherwise.
func (o *ReauthRevokeNotify) GetAuthMsg() string {
	if o == nil || isNil(o.AuthMsg) {
		var ret string
		return ret
	}
	return *o.AuthMsg
}

// GetAuthMsgOk returns a tuple with the AuthMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReauthRevokeNotify) GetAuthMsgOk() (*string, bool) {
	if o == nil || isNil(o.AuthMsg) {
    return nil, false
	}
	return o.AuthMsg, true
}

// HasAuthMsg returns a boolean if a field has been set.
func (o *ReauthRevokeNotify) HasAuthMsg() bool {
	if o != nil && !isNil(o.AuthMsg) {
		return true
	}

	return false
}

// SetAuthMsg gets a reference to the given string and assigns it to the AuthMsg field.
func (o *ReauthRevokeNotify) SetAuthMsg(v string) {
	o.AuthMsg = &v
}

// GetNotifyType returns the NotifyType field value
func (o *ReauthRevokeNotify) GetNotifyType() NotifyType {
	if o == nil {
		var ret NotifyType
		return ret
	}

	return o.NotifyType
}

// GetNotifyTypeOk returns a tuple with the NotifyType field value
// and a boolean to check if the value has been set.
func (o *ReauthRevokeNotify) GetNotifyTypeOk() (*NotifyType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifyType, true
}

// SetNotifyType sets field value
func (o *ReauthRevokeNotify) SetNotifyType(v NotifyType) {
	o.NotifyType = v
}

// GetIpAddr returns the IpAddr field value if set, zero value otherwise.
func (o *ReauthRevokeNotify) GetIpAddr() IpAddr {
	if o == nil || isNil(o.IpAddr) {
		var ret IpAddr
		return ret
	}
	return *o.IpAddr
}

// GetIpAddrOk returns a tuple with the IpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReauthRevokeNotify) GetIpAddrOk() (*IpAddr, bool) {
	if o == nil || isNil(o.IpAddr) {
    return nil, false
	}
	return o.IpAddr, true
}

// HasIpAddr returns a boolean if a field has been set.
func (o *ReauthRevokeNotify) HasIpAddr() bool {
	if o != nil && !isNil(o.IpAddr) {
		return true
	}

	return false
}

// SetIpAddr gets a reference to the given IpAddr and assigns it to the IpAddr field.
func (o *ReauthRevokeNotify) SetIpAddr(v IpAddr) {
	o.IpAddr = &v
}

func (o ReauthRevokeNotify) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["gpsi"] = o.Gpsi
	}
	if true {
		toSerialize["serviceLevelId"] = o.ServiceLevelId
	}
	if !isNil(o.NotifyCorrId) {
		toSerialize["notifyCorrId"] = o.NotifyCorrId
	}
	if !isNil(o.AuthMsg) {
		toSerialize["authMsg"] = o.AuthMsg
	}
	if true {
		toSerialize["notifyType"] = o.NotifyType
	}
	if !isNil(o.IpAddr) {
		toSerialize["ipAddr"] = o.IpAddr
	}
	return json.Marshal(toSerialize)
}

type NullableReauthRevokeNotify struct {
	value *ReauthRevokeNotify
	isSet bool
}

func (v NullableReauthRevokeNotify) Get() *ReauthRevokeNotify {
	return v.value
}

func (v *NullableReauthRevokeNotify) Set(val *ReauthRevokeNotify) {
	v.value = val
	v.isSet = true
}

func (v NullableReauthRevokeNotify) IsSet() bool {
	return v.isSet
}

func (v *NullableReauthRevokeNotify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReauthRevokeNotify(val *ReauthRevokeNotify) *NullableReauthRevokeNotify {
	return &NullableReauthRevokeNotify{value: val, isSet: true}
}

func (v NullableReauthRevokeNotify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReauthRevokeNotify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


