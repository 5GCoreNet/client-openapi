/*
Namf_EventExposure

AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_EventExposure

import (
	"encoding/json"
)

// AmfUpdateEventSubscriptionItem Document describing the modification(s) to an AMF Event Subscription
type AmfUpdateEventSubscriptionItem struct {
	Op string `json:"op"`
	Path string `json:"path"`
	Value *AmfEvent `json:"value,omitempty"`
	PresenceInfo *PresenceInfo `json:"presenceInfo,omitempty"`
	ExcludeSupiList []string `json:"excludeSupiList,omitempty"`
	ExcludeGpsiList []string `json:"excludeGpsiList,omitempty"`
	IncludeSupiList []string `json:"includeSupiList,omitempty"`
	IncludeGpsiList []string `json:"includeGpsiList,omitempty"`
}

// NewAmfUpdateEventSubscriptionItem instantiates a new AmfUpdateEventSubscriptionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfUpdateEventSubscriptionItem(op string, path string) *AmfUpdateEventSubscriptionItem {
	this := AmfUpdateEventSubscriptionItem{}
	this.Op = op
	this.Path = path
	return &this
}

// NewAmfUpdateEventSubscriptionItemWithDefaults instantiates a new AmfUpdateEventSubscriptionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfUpdateEventSubscriptionItemWithDefaults() *AmfUpdateEventSubscriptionItem {
	this := AmfUpdateEventSubscriptionItem{}
	return &this
}

// GetOp returns the Op field value
func (o *AmfUpdateEventSubscriptionItem) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *AmfUpdateEventSubscriptionItem) GetOpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *AmfUpdateEventSubscriptionItem) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *AmfUpdateEventSubscriptionItem) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *AmfUpdateEventSubscriptionItem) GetPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *AmfUpdateEventSubscriptionItem) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AmfUpdateEventSubscriptionItem) GetValue() AmfEvent {
	if o == nil || isNil(o.Value) {
		var ret AmfEvent
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfUpdateEventSubscriptionItem) GetValueOk() (*AmfEvent, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AmfUpdateEventSubscriptionItem) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given AmfEvent and assigns it to the Value field.
func (o *AmfUpdateEventSubscriptionItem) SetValue(v AmfEvent) {
	o.Value = &v
}

// GetPresenceInfo returns the PresenceInfo field value if set, zero value otherwise.
func (o *AmfUpdateEventSubscriptionItem) GetPresenceInfo() PresenceInfo {
	if o == nil || isNil(o.PresenceInfo) {
		var ret PresenceInfo
		return ret
	}
	return *o.PresenceInfo
}

// GetPresenceInfoOk returns a tuple with the PresenceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfUpdateEventSubscriptionItem) GetPresenceInfoOk() (*PresenceInfo, bool) {
	if o == nil || isNil(o.PresenceInfo) {
    return nil, false
	}
	return o.PresenceInfo, true
}

// HasPresenceInfo returns a boolean if a field has been set.
func (o *AmfUpdateEventSubscriptionItem) HasPresenceInfo() bool {
	if o != nil && !isNil(o.PresenceInfo) {
		return true
	}

	return false
}

// SetPresenceInfo gets a reference to the given PresenceInfo and assigns it to the PresenceInfo field.
func (o *AmfUpdateEventSubscriptionItem) SetPresenceInfo(v PresenceInfo) {
	o.PresenceInfo = &v
}

// GetExcludeSupiList returns the ExcludeSupiList field value if set, zero value otherwise.
func (o *AmfUpdateEventSubscriptionItem) GetExcludeSupiList() []string {
	if o == nil || isNil(o.ExcludeSupiList) {
		var ret []string
		return ret
	}
	return o.ExcludeSupiList
}

// GetExcludeSupiListOk returns a tuple with the ExcludeSupiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfUpdateEventSubscriptionItem) GetExcludeSupiListOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludeSupiList) {
    return nil, false
	}
	return o.ExcludeSupiList, true
}

// HasExcludeSupiList returns a boolean if a field has been set.
func (o *AmfUpdateEventSubscriptionItem) HasExcludeSupiList() bool {
	if o != nil && !isNil(o.ExcludeSupiList) {
		return true
	}

	return false
}

// SetExcludeSupiList gets a reference to the given []string and assigns it to the ExcludeSupiList field.
func (o *AmfUpdateEventSubscriptionItem) SetExcludeSupiList(v []string) {
	o.ExcludeSupiList = v
}

// GetExcludeGpsiList returns the ExcludeGpsiList field value if set, zero value otherwise.
func (o *AmfUpdateEventSubscriptionItem) GetExcludeGpsiList() []string {
	if o == nil || isNil(o.ExcludeGpsiList) {
		var ret []string
		return ret
	}
	return o.ExcludeGpsiList
}

// GetExcludeGpsiListOk returns a tuple with the ExcludeGpsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfUpdateEventSubscriptionItem) GetExcludeGpsiListOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludeGpsiList) {
    return nil, false
	}
	return o.ExcludeGpsiList, true
}

// HasExcludeGpsiList returns a boolean if a field has been set.
func (o *AmfUpdateEventSubscriptionItem) HasExcludeGpsiList() bool {
	if o != nil && !isNil(o.ExcludeGpsiList) {
		return true
	}

	return false
}

// SetExcludeGpsiList gets a reference to the given []string and assigns it to the ExcludeGpsiList field.
func (o *AmfUpdateEventSubscriptionItem) SetExcludeGpsiList(v []string) {
	o.ExcludeGpsiList = v
}

// GetIncludeSupiList returns the IncludeSupiList field value if set, zero value otherwise.
func (o *AmfUpdateEventSubscriptionItem) GetIncludeSupiList() []string {
	if o == nil || isNil(o.IncludeSupiList) {
		var ret []string
		return ret
	}
	return o.IncludeSupiList
}

// GetIncludeSupiListOk returns a tuple with the IncludeSupiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfUpdateEventSubscriptionItem) GetIncludeSupiListOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeSupiList) {
    return nil, false
	}
	return o.IncludeSupiList, true
}

// HasIncludeSupiList returns a boolean if a field has been set.
func (o *AmfUpdateEventSubscriptionItem) HasIncludeSupiList() bool {
	if o != nil && !isNil(o.IncludeSupiList) {
		return true
	}

	return false
}

// SetIncludeSupiList gets a reference to the given []string and assigns it to the IncludeSupiList field.
func (o *AmfUpdateEventSubscriptionItem) SetIncludeSupiList(v []string) {
	o.IncludeSupiList = v
}

// GetIncludeGpsiList returns the IncludeGpsiList field value if set, zero value otherwise.
func (o *AmfUpdateEventSubscriptionItem) GetIncludeGpsiList() []string {
	if o == nil || isNil(o.IncludeGpsiList) {
		var ret []string
		return ret
	}
	return o.IncludeGpsiList
}

// GetIncludeGpsiListOk returns a tuple with the IncludeGpsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfUpdateEventSubscriptionItem) GetIncludeGpsiListOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeGpsiList) {
    return nil, false
	}
	return o.IncludeGpsiList, true
}

// HasIncludeGpsiList returns a boolean if a field has been set.
func (o *AmfUpdateEventSubscriptionItem) HasIncludeGpsiList() bool {
	if o != nil && !isNil(o.IncludeGpsiList) {
		return true
	}

	return false
}

// SetIncludeGpsiList gets a reference to the given []string and assigns it to the IncludeGpsiList field.
func (o *AmfUpdateEventSubscriptionItem) SetIncludeGpsiList(v []string) {
	o.IncludeGpsiList = v
}

func (o AmfUpdateEventSubscriptionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["op"] = o.Op
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.PresenceInfo) {
		toSerialize["presenceInfo"] = o.PresenceInfo
	}
	if !isNil(o.ExcludeSupiList) {
		toSerialize["excludeSupiList"] = o.ExcludeSupiList
	}
	if !isNil(o.ExcludeGpsiList) {
		toSerialize["excludeGpsiList"] = o.ExcludeGpsiList
	}
	if !isNil(o.IncludeSupiList) {
		toSerialize["includeSupiList"] = o.IncludeSupiList
	}
	if !isNil(o.IncludeGpsiList) {
		toSerialize["includeGpsiList"] = o.IncludeGpsiList
	}
	return json.Marshal(toSerialize)
}

type NullableAmfUpdateEventSubscriptionItem struct {
	value *AmfUpdateEventSubscriptionItem
	isSet bool
}

func (v NullableAmfUpdateEventSubscriptionItem) Get() *AmfUpdateEventSubscriptionItem {
	return v.value
}

func (v *NullableAmfUpdateEventSubscriptionItem) Set(val *AmfUpdateEventSubscriptionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfUpdateEventSubscriptionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfUpdateEventSubscriptionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfUpdateEventSubscriptionItem(val *AmfUpdateEventSubscriptionItem) *NullableAmfUpdateEventSubscriptionItem {
	return &NullableAmfUpdateEventSubscriptionItem{value: val, isSet: true}
}

func (v NullableAmfUpdateEventSubscriptionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfUpdateEventSubscriptionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


