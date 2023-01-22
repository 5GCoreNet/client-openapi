/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_SMPolicyControl

import (
	"encoding/json"
)

// QosNotificationControlInfo Contains the QoS Notification Control Information.
type QosNotificationControlInfo struct {
	// An array of PCC rule id references to the PCC rules associated with the QoS notification control info.
	RefPccRuleIds []string `json:"refPccRuleIds"`
	NotifType QosNotifType `json:"notifType"`
	// Represents the content version of some content.
	ContVer *int32 `json:"contVer,omitempty"`
	AltQosParamId *string `json:"altQosParamId,omitempty"`
}

// NewQosNotificationControlInfo instantiates a new QosNotificationControlInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosNotificationControlInfo(refPccRuleIds []string, notifType QosNotifType) *QosNotificationControlInfo {
	this := QosNotificationControlInfo{}
	this.RefPccRuleIds = refPccRuleIds
	this.NotifType = notifType
	return &this
}

// NewQosNotificationControlInfoWithDefaults instantiates a new QosNotificationControlInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosNotificationControlInfoWithDefaults() *QosNotificationControlInfo {
	this := QosNotificationControlInfo{}
	return &this
}

// GetRefPccRuleIds returns the RefPccRuleIds field value
func (o *QosNotificationControlInfo) GetRefPccRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RefPccRuleIds
}

// GetRefPccRuleIdsOk returns a tuple with the RefPccRuleIds field value
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetRefPccRuleIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RefPccRuleIds, true
}

// SetRefPccRuleIds sets field value
func (o *QosNotificationControlInfo) SetRefPccRuleIds(v []string) {
	o.RefPccRuleIds = v
}

// GetNotifType returns the NotifType field value
func (o *QosNotificationControlInfo) GetNotifType() QosNotifType {
	if o == nil {
		var ret QosNotifType
		return ret
	}

	return o.NotifType
}

// GetNotifTypeOk returns a tuple with the NotifType field value
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetNotifTypeOk() (*QosNotifType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifType, true
}

// SetNotifType sets field value
func (o *QosNotificationControlInfo) SetNotifType(v QosNotifType) {
	o.NotifType = v
}

// GetContVer returns the ContVer field value if set, zero value otherwise.
func (o *QosNotificationControlInfo) GetContVer() int32 {
	if o == nil || isNil(o.ContVer) {
		var ret int32
		return ret
	}
	return *o.ContVer
}

// GetContVerOk returns a tuple with the ContVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetContVerOk() (*int32, bool) {
	if o == nil || isNil(o.ContVer) {
    return nil, false
	}
	return o.ContVer, true
}

// HasContVer returns a boolean if a field has been set.
func (o *QosNotificationControlInfo) HasContVer() bool {
	if o != nil && !isNil(o.ContVer) {
		return true
	}

	return false
}

// SetContVer gets a reference to the given int32 and assigns it to the ContVer field.
func (o *QosNotificationControlInfo) SetContVer(v int32) {
	o.ContVer = &v
}

// GetAltQosParamId returns the AltQosParamId field value if set, zero value otherwise.
func (o *QosNotificationControlInfo) GetAltQosParamId() string {
	if o == nil || isNil(o.AltQosParamId) {
		var ret string
		return ret
	}
	return *o.AltQosParamId
}

// GetAltQosParamIdOk returns a tuple with the AltQosParamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosNotificationControlInfo) GetAltQosParamIdOk() (*string, bool) {
	if o == nil || isNil(o.AltQosParamId) {
    return nil, false
	}
	return o.AltQosParamId, true
}

// HasAltQosParamId returns a boolean if a field has been set.
func (o *QosNotificationControlInfo) HasAltQosParamId() bool {
	if o != nil && !isNil(o.AltQosParamId) {
		return true
	}

	return false
}

// SetAltQosParamId gets a reference to the given string and assigns it to the AltQosParamId field.
func (o *QosNotificationControlInfo) SetAltQosParamId(v string) {
	o.AltQosParamId = &v
}

func (o QosNotificationControlInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["refPccRuleIds"] = o.RefPccRuleIds
	}
	if true {
		toSerialize["notifType"] = o.NotifType
	}
	if !isNil(o.ContVer) {
		toSerialize["contVer"] = o.ContVer
	}
	if !isNil(o.AltQosParamId) {
		toSerialize["altQosParamId"] = o.AltQosParamId
	}
	return json.Marshal(toSerialize)
}

type NullableQosNotificationControlInfo struct {
	value *QosNotificationControlInfo
	isSet bool
}

func (v NullableQosNotificationControlInfo) Get() *QosNotificationControlInfo {
	return v.value
}

func (v *NullableQosNotificationControlInfo) Set(val *QosNotificationControlInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableQosNotificationControlInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableQosNotificationControlInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosNotificationControlInfo(val *QosNotificationControlInfo) *NullableQosNotificationControlInfo {
	return &NullableQosNotificationControlInfo{value: val, isSet: true}
}

func (v NullableQosNotificationControlInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosNotificationControlInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


