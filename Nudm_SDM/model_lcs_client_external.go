/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
)

// LcsClientExternal struct for LcsClientExternal
type LcsClientExternal struct {
	AllowedGeographicArea []GeographicArea `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction *PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	ValidTimePeriod *ValidTimePeriod `json:"validTimePeriod,omitempty"`
}

// NewLcsClientExternal instantiates a new LcsClientExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLcsClientExternal() *LcsClientExternal {
	this := LcsClientExternal{}
	return &this
}

// NewLcsClientExternalWithDefaults instantiates a new LcsClientExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLcsClientExternalWithDefaults() *LcsClientExternal {
	this := LcsClientExternal{}
	return &this
}

// GetAllowedGeographicArea returns the AllowedGeographicArea field value if set, zero value otherwise.
func (o *LcsClientExternal) GetAllowedGeographicArea() []GeographicArea {
	if o == nil || isNil(o.AllowedGeographicArea) {
		var ret []GeographicArea
		return ret
	}
	return o.AllowedGeographicArea
}

// GetAllowedGeographicAreaOk returns a tuple with the AllowedGeographicArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsClientExternal) GetAllowedGeographicAreaOk() ([]GeographicArea, bool) {
	if o == nil || isNil(o.AllowedGeographicArea) {
    return nil, false
	}
	return o.AllowedGeographicArea, true
}

// HasAllowedGeographicArea returns a boolean if a field has been set.
func (o *LcsClientExternal) HasAllowedGeographicArea() bool {
	if o != nil && !isNil(o.AllowedGeographicArea) {
		return true
	}

	return false
}

// SetAllowedGeographicArea gets a reference to the given []GeographicArea and assigns it to the AllowedGeographicArea field.
func (o *LcsClientExternal) SetAllowedGeographicArea(v []GeographicArea) {
	o.AllowedGeographicArea = v
}

// GetPrivacyCheckRelatedAction returns the PrivacyCheckRelatedAction field value if set, zero value otherwise.
func (o *LcsClientExternal) GetPrivacyCheckRelatedAction() PrivacyCheckRelatedAction {
	if o == nil || isNil(o.PrivacyCheckRelatedAction) {
		var ret PrivacyCheckRelatedAction
		return ret
	}
	return *o.PrivacyCheckRelatedAction
}

// GetPrivacyCheckRelatedActionOk returns a tuple with the PrivacyCheckRelatedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsClientExternal) GetPrivacyCheckRelatedActionOk() (*PrivacyCheckRelatedAction, bool) {
	if o == nil || isNil(o.PrivacyCheckRelatedAction) {
    return nil, false
	}
	return o.PrivacyCheckRelatedAction, true
}

// HasPrivacyCheckRelatedAction returns a boolean if a field has been set.
func (o *LcsClientExternal) HasPrivacyCheckRelatedAction() bool {
	if o != nil && !isNil(o.PrivacyCheckRelatedAction) {
		return true
	}

	return false
}

// SetPrivacyCheckRelatedAction gets a reference to the given PrivacyCheckRelatedAction and assigns it to the PrivacyCheckRelatedAction field.
func (o *LcsClientExternal) SetPrivacyCheckRelatedAction(v PrivacyCheckRelatedAction) {
	o.PrivacyCheckRelatedAction = &v
}

// GetValidTimePeriod returns the ValidTimePeriod field value if set, zero value otherwise.
func (o *LcsClientExternal) GetValidTimePeriod() ValidTimePeriod {
	if o == nil || isNil(o.ValidTimePeriod) {
		var ret ValidTimePeriod
		return ret
	}
	return *o.ValidTimePeriod
}

// GetValidTimePeriodOk returns a tuple with the ValidTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LcsClientExternal) GetValidTimePeriodOk() (*ValidTimePeriod, bool) {
	if o == nil || isNil(o.ValidTimePeriod) {
    return nil, false
	}
	return o.ValidTimePeriod, true
}

// HasValidTimePeriod returns a boolean if a field has been set.
func (o *LcsClientExternal) HasValidTimePeriod() bool {
	if o != nil && !isNil(o.ValidTimePeriod) {
		return true
	}

	return false
}

// SetValidTimePeriod gets a reference to the given ValidTimePeriod and assigns it to the ValidTimePeriod field.
func (o *LcsClientExternal) SetValidTimePeriod(v ValidTimePeriod) {
	o.ValidTimePeriod = &v
}

func (o LcsClientExternal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableLcsClientExternal struct {
	value *LcsClientExternal
	isSet bool
}

func (v NullableLcsClientExternal) Get() *LcsClientExternal {
	return v.value
}

func (v *NullableLcsClientExternal) Set(val *LcsClientExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsClientExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsClientExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsClientExternal(val *LcsClientExternal) *NullableLcsClientExternal {
	return &NullableLcsClientExternal{value: val, isSet: true}
}

func (v NullableLcsClientExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsClientExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


