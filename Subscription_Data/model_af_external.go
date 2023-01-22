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

// AfExternal struct for AfExternal
type AfExternal struct {
	AfId *string `json:"afId,omitempty"`
	AllowedGeographicArea []GeographicArea `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction *PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	ValidTimePeriod *ValidTimePeriod `json:"validTimePeriod,omitempty"`
}

// NewAfExternal instantiates a new AfExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfExternal() *AfExternal {
	this := AfExternal{}
	return &this
}

// NewAfExternalWithDefaults instantiates a new AfExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfExternalWithDefaults() *AfExternal {
	this := AfExternal{}
	return &this
}

// GetAfId returns the AfId field value if set, zero value otherwise.
func (o *AfExternal) GetAfId() string {
	if o == nil || isNil(o.AfId) {
		var ret string
		return ret
	}
	return *o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfExternal) GetAfIdOk() (*string, bool) {
	if o == nil || isNil(o.AfId) {
    return nil, false
	}
	return o.AfId, true
}

// HasAfId returns a boolean if a field has been set.
func (o *AfExternal) HasAfId() bool {
	if o != nil && !isNil(o.AfId) {
		return true
	}

	return false
}

// SetAfId gets a reference to the given string and assigns it to the AfId field.
func (o *AfExternal) SetAfId(v string) {
	o.AfId = &v
}

// GetAllowedGeographicArea returns the AllowedGeographicArea field value if set, zero value otherwise.
func (o *AfExternal) GetAllowedGeographicArea() []GeographicArea {
	if o == nil || isNil(o.AllowedGeographicArea) {
		var ret []GeographicArea
		return ret
	}
	return o.AllowedGeographicArea
}

// GetAllowedGeographicAreaOk returns a tuple with the AllowedGeographicArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfExternal) GetAllowedGeographicAreaOk() ([]GeographicArea, bool) {
	if o == nil || isNil(o.AllowedGeographicArea) {
    return nil, false
	}
	return o.AllowedGeographicArea, true
}

// HasAllowedGeographicArea returns a boolean if a field has been set.
func (o *AfExternal) HasAllowedGeographicArea() bool {
	if o != nil && !isNil(o.AllowedGeographicArea) {
		return true
	}

	return false
}

// SetAllowedGeographicArea gets a reference to the given []GeographicArea and assigns it to the AllowedGeographicArea field.
func (o *AfExternal) SetAllowedGeographicArea(v []GeographicArea) {
	o.AllowedGeographicArea = v
}

// GetPrivacyCheckRelatedAction returns the PrivacyCheckRelatedAction field value if set, zero value otherwise.
func (o *AfExternal) GetPrivacyCheckRelatedAction() PrivacyCheckRelatedAction {
	if o == nil || isNil(o.PrivacyCheckRelatedAction) {
		var ret PrivacyCheckRelatedAction
		return ret
	}
	return *o.PrivacyCheckRelatedAction
}

// GetPrivacyCheckRelatedActionOk returns a tuple with the PrivacyCheckRelatedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfExternal) GetPrivacyCheckRelatedActionOk() (*PrivacyCheckRelatedAction, bool) {
	if o == nil || isNil(o.PrivacyCheckRelatedAction) {
    return nil, false
	}
	return o.PrivacyCheckRelatedAction, true
}

// HasPrivacyCheckRelatedAction returns a boolean if a field has been set.
func (o *AfExternal) HasPrivacyCheckRelatedAction() bool {
	if o != nil && !isNil(o.PrivacyCheckRelatedAction) {
		return true
	}

	return false
}

// SetPrivacyCheckRelatedAction gets a reference to the given PrivacyCheckRelatedAction and assigns it to the PrivacyCheckRelatedAction field.
func (o *AfExternal) SetPrivacyCheckRelatedAction(v PrivacyCheckRelatedAction) {
	o.PrivacyCheckRelatedAction = &v
}

// GetValidTimePeriod returns the ValidTimePeriod field value if set, zero value otherwise.
func (o *AfExternal) GetValidTimePeriod() ValidTimePeriod {
	if o == nil || isNil(o.ValidTimePeriod) {
		var ret ValidTimePeriod
		return ret
	}
	return *o.ValidTimePeriod
}

// GetValidTimePeriodOk returns a tuple with the ValidTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfExternal) GetValidTimePeriodOk() (*ValidTimePeriod, bool) {
	if o == nil || isNil(o.ValidTimePeriod) {
    return nil, false
	}
	return o.ValidTimePeriod, true
}

// HasValidTimePeriod returns a boolean if a field has been set.
func (o *AfExternal) HasValidTimePeriod() bool {
	if o != nil && !isNil(o.ValidTimePeriod) {
		return true
	}

	return false
}

// SetValidTimePeriod gets a reference to the given ValidTimePeriod and assigns it to the ValidTimePeriod field.
func (o *AfExternal) SetValidTimePeriod(v ValidTimePeriod) {
	o.ValidTimePeriod = &v
}

func (o AfExternal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AfId) {
		toSerialize["afId"] = o.AfId
	}
	if !isNil(o.AllowedGeographicArea) {
		toSerialize["allowedGeographicArea"] = o.AllowedGeographicArea
	}
	if !isNil(o.PrivacyCheckRelatedAction) {
		toSerialize["privacyCheckRelatedAction"] = o.PrivacyCheckRelatedAction
	}
	if !isNil(o.ValidTimePeriod) {
		toSerialize["validTimePeriod"] = o.ValidTimePeriod
	}
	return json.Marshal(toSerialize)
}

type NullableAfExternal struct {
	value *AfExternal
	isSet bool
}

func (v NullableAfExternal) Get() *AfExternal {
	return v.value
}

func (v *NullableAfExternal) Set(val *AfExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableAfExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableAfExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfExternal(val *AfExternal) *NullableAfExternal {
	return &NullableAfExternal{value: val, isSet: true}
}

func (v NullableAfExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


