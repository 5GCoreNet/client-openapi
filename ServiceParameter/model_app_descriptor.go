/*
3gpp-service-parameter

API for AF service paramter   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ServiceParameter

import (
	"encoding/json"
)

// AppDescriptor Represents an operation system and the corresponding applications.
type AppDescriptor struct {
	// Represents the Operating System of the served UE.
	OsId string `json:"osId"`
	// Identifies applications that are running on the UE's operating system. Any string value can be used as a key of the map. 
	AppIds map[string]string `json:"appIds"`
}

// NewAppDescriptor instantiates a new AppDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDescriptor(osId string, appIds map[string]string) *AppDescriptor {
	this := AppDescriptor{}
	this.OsId = osId
	this.AppIds = appIds
	return &this
}

// NewAppDescriptorWithDefaults instantiates a new AppDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDescriptorWithDefaults() *AppDescriptor {
	this := AppDescriptor{}
	return &this
}

// GetOsId returns the OsId field value
func (o *AppDescriptor) GetOsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsId
}

// GetOsIdOk returns a tuple with the OsId field value
// and a boolean to check if the value has been set.
func (o *AppDescriptor) GetOsIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OsId, true
}

// SetOsId sets field value
func (o *AppDescriptor) SetOsId(v string) {
	o.OsId = v
}

// GetAppIds returns the AppIds field value
func (o *AppDescriptor) GetAppIds() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value
// and a boolean to check if the value has been set.
func (o *AppDescriptor) GetAppIdsOk() (*map[string]string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppIds, true
}

// SetAppIds sets field value
func (o *AppDescriptor) SetAppIds(v map[string]string) {
	o.AppIds = v
}

func (o AppDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["osId"] = o.OsId
	}
	if true {
		toSerialize["appIds"] = o.AppIds
	}
	return json.Marshal(toSerialize)
}

type NullableAppDescriptor struct {
	value *AppDescriptor
	isSet bool
}

func (v NullableAppDescriptor) Get() *AppDescriptor {
	return v.value
}

func (v *NullableAppDescriptor) Set(val *AppDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDescriptor(val *AppDescriptor) *NullableAppDescriptor {
	return &NullableAppDescriptor{value: val, isSet: true}
}

func (v NullableAppDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


