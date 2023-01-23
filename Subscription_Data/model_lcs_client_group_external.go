/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// LcsClientGroupExternal struct for LcsClientGroupExternal
type LcsClientGroupExternal struct {
	LcsClientGroupId *string `json:"lcsClientGroupId,omitempty"`
	AllowedGeographicArea []GeographicArea `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction *PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	ValidTimePeriod *ValidTimePeriod `json:"validTimePeriod,omitempty"`
}

// NewLcsClientGroupExternal instantiates a new LcsClientGroupExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLcsClientGroupExternal() *LcsClientGroupExternal {
	this := LcsClientGroupExternal{}
	return &this
}

// NewLcsClientGroupExternalWithDefaults instantiates a new LcsClientGroupExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLcsClientGroupExternalWithDefaults() *LcsClientGroupExternal {
	this := LcsClientGroupExternal{}
	return &this
}

// GetLcsClientGroupId returns the LcsClientGroupId field value if set, zero value otherwise.
func (o *LcsClientGroupExternal) GetLcsClientGroupId() string {
	if o == nil || isNil(o.LcsClientGroupId) {
		var ret string
		return ret
	}
	return *o.LcsClientGroupId
}

// GetLcsClientGroupIdOk returns a tuple with the LcsClientGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsClientGroupExternal) GetLcsClientGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.LcsClientGroupId) {
    return nil, false
	}
	return o.LcsClientGroupId, true
}

// HasLcsClientGroupId returns a boolean if a field has been set.
func (o *LcsClientGroupExternal) HasLcsClientGroupId() bool {
	if o != nil && !isNil(o.LcsClientGroupId) {
		return true
	}

	return false
}

// SetLcsClientGroupId gets a reference to the given string and assigns it to the LcsClientGroupId field.
func (o *LcsClientGroupExternal) SetLcsClientGroupId(v string) {
	o.LcsClientGroupId = &v
}

// GetAllowedGeographicArea returns the AllowedGeographicArea field value if set, zero value otherwise.
func (o *LcsClientGroupExternal) GetAllowedGeographicArea() []GeographicArea {
	if o == nil || isNil(o.AllowedGeographicArea) {
		var ret []GeographicArea
		return ret
	}
	return o.AllowedGeographicArea
}

// GetAllowedGeographicAreaOk returns a tuple with the AllowedGeographicArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsClientGroupExternal) GetAllowedGeographicAreaOk() ([]GeographicArea, bool) {
	if o == nil || isNil(o.AllowedGeographicArea) {
    return nil, false
	}
	return o.AllowedGeographicArea, true
}

// HasAllowedGeographicArea returns a boolean if a field has been set.
func (o *LcsClientGroupExternal) HasAllowedGeographicArea() bool {
	if o != nil && !isNil(o.AllowedGeographicArea) {
		return true
	}

	return false
}

// SetAllowedGeographicArea gets a reference to the given []GeographicArea and assigns it to the AllowedGeographicArea field.
func (o *LcsClientGroupExternal) SetAllowedGeographicArea(v []GeographicArea) {
	o.AllowedGeographicArea = v
}

// GetPrivacyCheckRelatedAction returns the PrivacyCheckRelatedAction field value if set, zero value otherwise.
func (o *LcsClientGroupExternal) GetPrivacyCheckRelatedAction() PrivacyCheckRelatedAction {
	if o == nil || isNil(o.PrivacyCheckRelatedAction) {
		var ret PrivacyCheckRelatedAction
		return ret
	}
	return *o.PrivacyCheckRelatedAction
}

// GetPrivacyCheckRelatedActionOk returns a tuple with the PrivacyCheckRelatedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsClientGroupExternal) GetPrivacyCheckRelatedActionOk() (*PrivacyCheckRelatedAction, bool) {
	if o == nil || isNil(o.PrivacyCheckRelatedAction) {
    return nil, false
	}
	return o.PrivacyCheckRelatedAction, true
}

// HasPrivacyCheckRelatedAction returns a boolean if a field has been set.
func (o *LcsClientGroupExternal) HasPrivacyCheckRelatedAction() bool {
	if o != nil && !isNil(o.PrivacyCheckRelatedAction) {
		return true
	}

	return false
}

// SetPrivacyCheckRelatedAction gets a reference to the given PrivacyCheckRelatedAction and assigns it to the PrivacyCheckRelatedAction field.
func (o *LcsClientGroupExternal) SetPrivacyCheckRelatedAction(v PrivacyCheckRelatedAction) {
	o.PrivacyCheckRelatedAction = &v
}

// GetValidTimePeriod returns the ValidTimePeriod field value if set, zero value otherwise.
func (o *LcsClientGroupExternal) GetValidTimePeriod() ValidTimePeriod {
	if o == nil || isNil(o.ValidTimePeriod) {
		var ret ValidTimePeriod
		return ret
	}
	return *o.ValidTimePeriod
}

// GetValidTimePeriodOk returns a tuple with the ValidTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsClientGroupExternal) GetValidTimePeriodOk() (*ValidTimePeriod, bool) {
	if o == nil || isNil(o.ValidTimePeriod) {
    return nil, false
	}
	return o.ValidTimePeriod, true
}

// HasValidTimePeriod returns a boolean if a field has been set.
func (o *LcsClientGroupExternal) HasValidTimePeriod() bool {
	if o != nil && !isNil(o.ValidTimePeriod) {
		return true
	}

	return false
}

// SetValidTimePeriod gets a reference to the given ValidTimePeriod and assigns it to the ValidTimePeriod field.
func (o *LcsClientGroupExternal) SetValidTimePeriod(v ValidTimePeriod) {
	o.ValidTimePeriod = &v
}

func (o LcsClientGroupExternal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LcsClientGroupId) {
		toSerialize["lcsClientGroupId"] = o.LcsClientGroupId
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

type NullableLcsClientGroupExternal struct {
	value *LcsClientGroupExternal
	isSet bool
}

func (v NullableLcsClientGroupExternal) Get() *LcsClientGroupExternal {
	return v.value
}

func (v *NullableLcsClientGroupExternal) Set(val *LcsClientGroupExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsClientGroupExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsClientGroupExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsClientGroupExternal(val *LcsClientGroupExternal) *NullableLcsClientGroupExternal {
	return &NullableLcsClientGroupExternal{value: val, isSet: true}
}

func (v NullableLcsClientGroupExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsClientGroupExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


