/*
3gpp-ms-event-exposure

API for Media Streaming Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSEventExposure

import (
	"encoding/json"
)

// EventFilter Represents event filter information for an event.
type EventFilter struct {
	Gpsis []string `json:"gpsis,omitempty"`
	Supis []string `json:"supis,omitempty"`
	ExterGroupIds []string `json:"exterGroupIds,omitempty"`
	InterGroupIds []string `json:"interGroupIds,omitempty"`
	AnyUeInd *bool `json:"anyUeInd,omitempty"`
	AppIds []string `json:"appIds,omitempty"`
	LocArea *LocationArea5G `json:"locArea,omitempty"`
	CollAttrs []CollectiveBehaviourFilter `json:"collAttrs,omitempty"`
}

// NewEventFilter instantiates a new EventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventFilter() *EventFilter {
	this := EventFilter{}
	return &this
}

// NewEventFilterWithDefaults instantiates a new EventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventFilterWithDefaults() *EventFilter {
	this := EventFilter{}
	return &this
}

// GetGpsis returns the Gpsis field value if set, zero value otherwise.
func (o *EventFilter) GetGpsis() []string {
	if o == nil || isNil(o.Gpsis) {
		var ret []string
		return ret
	}
	return o.Gpsis
}

// GetGpsisOk returns a tuple with the Gpsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetGpsisOk() ([]string, bool) {
	if o == nil || isNil(o.Gpsis) {
    return nil, false
	}
	return o.Gpsis, true
}

// HasGpsis returns a boolean if a field has been set.
func (o *EventFilter) HasGpsis() bool {
	if o != nil && !isNil(o.Gpsis) {
		return true
	}

	return false
}

// SetGpsis gets a reference to the given []string and assigns it to the Gpsis field.
func (o *EventFilter) SetGpsis(v []string) {
	o.Gpsis = v
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *EventFilter) GetSupis() []string {
	if o == nil || isNil(o.Supis) {
		var ret []string
		return ret
	}
	return o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetSupisOk() ([]string, bool) {
	if o == nil || isNil(o.Supis) {
    return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *EventFilter) HasSupis() bool {
	if o != nil && !isNil(o.Supis) {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []string and assigns it to the Supis field.
func (o *EventFilter) SetSupis(v []string) {
	o.Supis = v
}

// GetExterGroupIds returns the ExterGroupIds field value if set, zero value otherwise.
func (o *EventFilter) GetExterGroupIds() []string {
	if o == nil || isNil(o.ExterGroupIds) {
		var ret []string
		return ret
	}
	return o.ExterGroupIds
}

// GetExterGroupIdsOk returns a tuple with the ExterGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetExterGroupIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ExterGroupIds) {
    return nil, false
	}
	return o.ExterGroupIds, true
}

// HasExterGroupIds returns a boolean if a field has been set.
func (o *EventFilter) HasExterGroupIds() bool {
	if o != nil && !isNil(o.ExterGroupIds) {
		return true
	}

	return false
}

// SetExterGroupIds gets a reference to the given []string and assigns it to the ExterGroupIds field.
func (o *EventFilter) SetExterGroupIds(v []string) {
	o.ExterGroupIds = v
}

// GetInterGroupIds returns the InterGroupIds field value if set, zero value otherwise.
func (o *EventFilter) GetInterGroupIds() []string {
	if o == nil || isNil(o.InterGroupIds) {
		var ret []string
		return ret
	}
	return o.InterGroupIds
}

// GetInterGroupIdsOk returns a tuple with the InterGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetInterGroupIdsOk() ([]string, bool) {
	if o == nil || isNil(o.InterGroupIds) {
    return nil, false
	}
	return o.InterGroupIds, true
}

// HasInterGroupIds returns a boolean if a field has been set.
func (o *EventFilter) HasInterGroupIds() bool {
	if o != nil && !isNil(o.InterGroupIds) {
		return true
	}

	return false
}

// SetInterGroupIds gets a reference to the given []string and assigns it to the InterGroupIds field.
func (o *EventFilter) SetInterGroupIds(v []string) {
	o.InterGroupIds = v
}

// GetAnyUeInd returns the AnyUeInd field value if set, zero value otherwise.
func (o *EventFilter) GetAnyUeInd() bool {
	if o == nil || isNil(o.AnyUeInd) {
		var ret bool
		return ret
	}
	return *o.AnyUeInd
}

// GetAnyUeIndOk returns a tuple with the AnyUeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetAnyUeIndOk() (*bool, bool) {
	if o == nil || isNil(o.AnyUeInd) {
    return nil, false
	}
	return o.AnyUeInd, true
}

// HasAnyUeInd returns a boolean if a field has been set.
func (o *EventFilter) HasAnyUeInd() bool {
	if o != nil && !isNil(o.AnyUeInd) {
		return true
	}

	return false
}

// SetAnyUeInd gets a reference to the given bool and assigns it to the AnyUeInd field.
func (o *EventFilter) SetAnyUeInd(v bool) {
	o.AnyUeInd = &v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *EventFilter) GetAppIds() []string {
	if o == nil || isNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetAppIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AppIds) {
    return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *EventFilter) HasAppIds() bool {
	if o != nil && !isNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *EventFilter) SetAppIds(v []string) {
	o.AppIds = v
}

// GetLocArea returns the LocArea field value if set, zero value otherwise.
func (o *EventFilter) GetLocArea() LocationArea5G {
	if o == nil || isNil(o.LocArea) {
		var ret LocationArea5G
		return ret
	}
	return *o.LocArea
}

// GetLocAreaOk returns a tuple with the LocArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetLocAreaOk() (*LocationArea5G, bool) {
	if o == nil || isNil(o.LocArea) {
    return nil, false
	}
	return o.LocArea, true
}

// HasLocArea returns a boolean if a field has been set.
func (o *EventFilter) HasLocArea() bool {
	if o != nil && !isNil(o.LocArea) {
		return true
	}

	return false
}

// SetLocArea gets a reference to the given LocationArea5G and assigns it to the LocArea field.
func (o *EventFilter) SetLocArea(v LocationArea5G) {
	o.LocArea = &v
}

// GetCollAttrs returns the CollAttrs field value if set, zero value otherwise.
func (o *EventFilter) GetCollAttrs() []CollectiveBehaviourFilter {
	if o == nil || isNil(o.CollAttrs) {
		var ret []CollectiveBehaviourFilter
		return ret
	}
	return o.CollAttrs
}

// GetCollAttrsOk returns a tuple with the CollAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetCollAttrsOk() ([]CollectiveBehaviourFilter, bool) {
	if o == nil || isNil(o.CollAttrs) {
    return nil, false
	}
	return o.CollAttrs, true
}

// HasCollAttrs returns a boolean if a field has been set.
func (o *EventFilter) HasCollAttrs() bool {
	if o != nil && !isNil(o.CollAttrs) {
		return true
	}

	return false
}

// SetCollAttrs gets a reference to the given []CollectiveBehaviourFilter and assigns it to the CollAttrs field.
func (o *EventFilter) SetCollAttrs(v []CollectiveBehaviourFilter) {
	o.CollAttrs = v
}

func (o EventFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Gpsis) {
		toSerialize["gpsis"] = o.Gpsis
	}
	if !isNil(o.Supis) {
		toSerialize["supis"] = o.Supis
	}
	if !isNil(o.ExterGroupIds) {
		toSerialize["exterGroupIds"] = o.ExterGroupIds
	}
	if !isNil(o.InterGroupIds) {
		toSerialize["interGroupIds"] = o.InterGroupIds
	}
	if !isNil(o.AnyUeInd) {
		toSerialize["anyUeInd"] = o.AnyUeInd
	}
	if !isNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	if !isNil(o.LocArea) {
		toSerialize["locArea"] = o.LocArea
	}
	if !isNil(o.CollAttrs) {
		toSerialize["collAttrs"] = o.CollAttrs
	}
	return json.Marshal(toSerialize)
}

type NullableEventFilter struct {
	value *EventFilter
	isSet bool
}

func (v NullableEventFilter) Get() *EventFilter {
	return v.value
}

func (v *NullableEventFilter) Set(val *EventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventFilter(val *EventFilter) *NullableEventFilter {
	return &NullableEventFilter{value: val, isSet: true}
}

func (v NullableEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


