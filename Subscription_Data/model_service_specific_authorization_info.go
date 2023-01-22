/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
)

// ServiceSpecificAuthorizationInfo Information related to active Service Specific Authorizations
type ServiceSpecificAuthorizationInfo struct {
	ServiceSpecificAuthorizationList []AuthorizationInfo `json:"serviceSpecificAuthorizationList"`
}

// NewServiceSpecificAuthorizationInfo instantiates a new ServiceSpecificAuthorizationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSpecificAuthorizationInfo(serviceSpecificAuthorizationList []AuthorizationInfo) *ServiceSpecificAuthorizationInfo {
	this := ServiceSpecificAuthorizationInfo{}
	this.ServiceSpecificAuthorizationList = serviceSpecificAuthorizationList
	return &this
}

// NewServiceSpecificAuthorizationInfoWithDefaults instantiates a new ServiceSpecificAuthorizationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSpecificAuthorizationInfoWithDefaults() *ServiceSpecificAuthorizationInfo {
	this := ServiceSpecificAuthorizationInfo{}
	return &this
}

// GetServiceSpecificAuthorizationList returns the ServiceSpecificAuthorizationList field value
func (o *ServiceSpecificAuthorizationInfo) GetServiceSpecificAuthorizationList() []AuthorizationInfo {
	if o == nil {
		var ret []AuthorizationInfo
		return ret
	}

	return o.ServiceSpecificAuthorizationList
}

// GetServiceSpecificAuthorizationListOk returns a tuple with the ServiceSpecificAuthorizationList field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificAuthorizationInfo) GetServiceSpecificAuthorizationListOk() ([]AuthorizationInfo, bool) {
	if o == nil {
    return nil, false
	}
	return o.ServiceSpecificAuthorizationList, true
}

// SetServiceSpecificAuthorizationList sets field value
func (o *ServiceSpecificAuthorizationInfo) SetServiceSpecificAuthorizationList(v []AuthorizationInfo) {
	o.ServiceSpecificAuthorizationList = v
}

func (o ServiceSpecificAuthorizationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serviceSpecificAuthorizationList"] = o.ServiceSpecificAuthorizationList
	}
	return json.Marshal(toSerialize)
}

type NullableServiceSpecificAuthorizationInfo struct {
	value *ServiceSpecificAuthorizationInfo
	isSet bool
}

func (v NullableServiceSpecificAuthorizationInfo) Get() *ServiceSpecificAuthorizationInfo {
	return v.value
}

func (v *NullableServiceSpecificAuthorizationInfo) Set(val *ServiceSpecificAuthorizationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSpecificAuthorizationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSpecificAuthorizationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSpecificAuthorizationInfo(val *ServiceSpecificAuthorizationInfo) *NullableServiceSpecificAuthorizationInfo {
	return &NullableServiceSpecificAuthorizationInfo{value: val, isSet: true}
}

func (v NullableServiceSpecificAuthorizationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSpecificAuthorizationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

