/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// AppSessionContext Represents an Individual Application Session Context resource.
type AppSessionContext struct {
	AscReqData *AppSessionContextReqData `json:"ascReqData,omitempty"`
	AscRespData *AppSessionContextRespData `json:"ascRespData,omitempty"`
	EvsNotif *EventsNotification `json:"evsNotif,omitempty"`
}

// NewAppSessionContext instantiates a new AppSessionContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppSessionContext() *AppSessionContext {
	this := AppSessionContext{}
	return &this
}

// NewAppSessionContextWithDefaults instantiates a new AppSessionContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppSessionContextWithDefaults() *AppSessionContext {
	this := AppSessionContext{}
	return &this
}

// GetAscReqData returns the AscReqData field value if set, zero value otherwise.
func (o *AppSessionContext) GetAscReqData() AppSessionContextReqData {
	if o == nil || isNil(o.AscReqData) {
		var ret AppSessionContextReqData
		return ret
	}
	return *o.AscReqData
}

// GetAscReqDataOk returns a tuple with the AscReqData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContext) GetAscReqDataOk() (*AppSessionContextReqData, bool) {
	if o == nil || isNil(o.AscReqData) {
    return nil, false
	}
	return o.AscReqData, true
}

// HasAscReqData returns a boolean if a field has been set.
func (o *AppSessionContext) HasAscReqData() bool {
	if o != nil && !isNil(o.AscReqData) {
		return true
	}

	return false
}

// SetAscReqData gets a reference to the given AppSessionContextReqData and assigns it to the AscReqData field.
func (o *AppSessionContext) SetAscReqData(v AppSessionContextReqData) {
	o.AscReqData = &v
}

// GetAscRespData returns the AscRespData field value if set, zero value otherwise.
func (o *AppSessionContext) GetAscRespData() AppSessionContextRespData {
	if o == nil || isNil(o.AscRespData) {
		var ret AppSessionContextRespData
		return ret
	}
	return *o.AscRespData
}

// GetAscRespDataOk returns a tuple with the AscRespData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContext) GetAscRespDataOk() (*AppSessionContextRespData, bool) {
	if o == nil || isNil(o.AscRespData) {
    return nil, false
	}
	return o.AscRespData, true
}

// HasAscRespData returns a boolean if a field has been set.
func (o *AppSessionContext) HasAscRespData() bool {
	if o != nil && !isNil(o.AscRespData) {
		return true
	}

	return false
}

// SetAscRespData gets a reference to the given AppSessionContextRespData and assigns it to the AscRespData field.
func (o *AppSessionContext) SetAscRespData(v AppSessionContextRespData) {
	o.AscRespData = &v
}

// GetEvsNotif returns the EvsNotif field value if set, zero value otherwise.
func (o *AppSessionContext) GetEvsNotif() EventsNotification {
	if o == nil || isNil(o.EvsNotif) {
		var ret EventsNotification
		return ret
	}
	return *o.EvsNotif
}

// GetEvsNotifOk returns a tuple with the EvsNotif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSessionContext) GetEvsNotifOk() (*EventsNotification, bool) {
	if o == nil || isNil(o.EvsNotif) {
    return nil, false
	}
	return o.EvsNotif, true
}

// HasEvsNotif returns a boolean if a field has been set.
func (o *AppSessionContext) HasEvsNotif() bool {
	if o != nil && !isNil(o.EvsNotif) {
		return true
	}

	return false
}

// SetEvsNotif gets a reference to the given EventsNotification and assigns it to the EvsNotif field.
func (o *AppSessionContext) SetEvsNotif(v EventsNotification) {
	o.EvsNotif = &v
}

func (o AppSessionContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AscReqData) {
		toSerialize["ascReqData"] = o.AscReqData
	}
	if !isNil(o.AscRespData) {
		toSerialize["ascRespData"] = o.AscRespData
	}
	if !isNil(o.EvsNotif) {
		toSerialize["evsNotif"] = o.EvsNotif
	}
	return json.Marshal(toSerialize)
}

type NullableAppSessionContext struct {
	value *AppSessionContext
	isSet bool
}

func (v NullableAppSessionContext) Get() *AppSessionContext {
	return v.value
}

func (v *NullableAppSessionContext) Set(val *AppSessionContext) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSessionContext) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSessionContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSessionContext(val *AppSessionContext) *NullableAppSessionContext {
	return &NullableAppSessionContext{value: val, isSet: true}
}

func (v NullableAppSessionContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSessionContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


