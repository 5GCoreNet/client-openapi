/*
EES Application Client Information_API

API for EES Application Client Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_AppClientInformation

import (
	"encoding/json"
)

// ACInfoNotification AC Information notification.
type ACInfoNotification struct {
	// Identifier of the AC information subscription for which this notification is related to. 
	SubId string `json:"subId"`
	// Notifications that include the ACs information matching filter criteria.
	AcInfs []ACInformation `json:"acInfs"`
}

// NewACInfoNotification instantiates a new ACInfoNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACInfoNotification(subId string, acInfs []ACInformation) *ACInfoNotification {
	this := ACInfoNotification{}
	this.SubId = subId
	this.AcInfs = acInfs
	return &this
}

// NewACInfoNotificationWithDefaults instantiates a new ACInfoNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACInfoNotificationWithDefaults() *ACInfoNotification {
	this := ACInfoNotification{}
	return &this
}

// GetSubId returns the SubId field value
func (o *ACInfoNotification) GetSubId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubId
}

// GetSubIdOk returns a tuple with the SubId field value
// and a boolean to check if the value has been set.
func (o *ACInfoNotification) GetSubIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubId, true
}

// SetSubId sets field value
func (o *ACInfoNotification) SetSubId(v string) {
	o.SubId = v
}

// GetAcInfs returns the AcInfs field value
func (o *ACInfoNotification) GetAcInfs() []ACInformation {
	if o == nil {
		var ret []ACInformation
		return ret
	}

	return o.AcInfs
}

// GetAcInfsOk returns a tuple with the AcInfs field value
// and a boolean to check if the value has been set.
func (o *ACInfoNotification) GetAcInfsOk() ([]ACInformation, bool) {
	if o == nil {
    return nil, false
	}
	return o.AcInfs, true
}

// SetAcInfs sets field value
func (o *ACInfoNotification) SetAcInfs(v []ACInformation) {
	o.AcInfs = v
}

func (o ACInfoNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subId"] = o.SubId
	}
	if true {
		toSerialize["acInfs"] = o.AcInfs
	}
	return json.Marshal(toSerialize)
}

type NullableACInfoNotification struct {
	value *ACInfoNotification
	isSet bool
}

func (v NullableACInfoNotification) Get() *ACInfoNotification {
	return v.value
}

func (v *NullableACInfoNotification) Set(val *ACInfoNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableACInfoNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableACInfoNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACInfoNotification(val *ACInfoNotification) *NullableACInfoNotification {
	return &NullableACInfoNotification{value: val, isSet: true}
}

func (v NullableACInfoNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACInfoNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


