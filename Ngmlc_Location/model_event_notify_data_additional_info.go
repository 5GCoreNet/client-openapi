/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
)

// EventNotifyDataAdditionalInfo Additional information to Event Notify Data
type EventNotifyDataAdditionalInfo struct {
	AddEventDataList []EventNotifyData `json:"addEventDataList,omitempty"`
}

// NewEventNotifyDataAdditionalInfo instantiates a new EventNotifyDataAdditionalInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventNotifyDataAdditionalInfo() *EventNotifyDataAdditionalInfo {
	this := EventNotifyDataAdditionalInfo{}
	return &this
}

// NewEventNotifyDataAdditionalInfoWithDefaults instantiates a new EventNotifyDataAdditionalInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventNotifyDataAdditionalInfoWithDefaults() *EventNotifyDataAdditionalInfo {
	this := EventNotifyDataAdditionalInfo{}
	return &this
}

// GetAddEventDataList returns the AddEventDataList field value if set, zero value otherwise.
func (o *EventNotifyDataAdditionalInfo) GetAddEventDataList() []EventNotifyData {
	if o == nil || isNil(o.AddEventDataList) {
		var ret []EventNotifyData
		return ret
	}
	return o.AddEventDataList
}

// GetAddEventDataListOk returns a tuple with the AddEventDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotifyDataAdditionalInfo) GetAddEventDataListOk() ([]EventNotifyData, bool) {
	if o == nil || isNil(o.AddEventDataList) {
    return nil, false
	}
	return o.AddEventDataList, true
}

// HasAddEventDataList returns a boolean if a field has been set.
func (o *EventNotifyDataAdditionalInfo) HasAddEventDataList() bool {
	if o != nil && !isNil(o.AddEventDataList) {
		return true
	}

	return false
}

// SetAddEventDataList gets a reference to the given []EventNotifyData and assigns it to the AddEventDataList field.
func (o *EventNotifyDataAdditionalInfo) SetAddEventDataList(v []EventNotifyData) {
	o.AddEventDataList = v
}

func (o EventNotifyDataAdditionalInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AddEventDataList) {
		toSerialize["addEventDataList"] = o.AddEventDataList
	}
	return json.Marshal(toSerialize)
}

type NullableEventNotifyDataAdditionalInfo struct {
	value *EventNotifyDataAdditionalInfo
	isSet bool
}

func (v NullableEventNotifyDataAdditionalInfo) Get() *EventNotifyDataAdditionalInfo {
	return v.value
}

func (v *NullableEventNotifyDataAdditionalInfo) Set(val *EventNotifyDataAdditionalInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEventNotifyDataAdditionalInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEventNotifyDataAdditionalInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventNotifyDataAdditionalInfo(val *EventNotifyDataAdditionalInfo) *NullableEventNotifyDataAdditionalInfo {
	return &NullableEventNotifyDataAdditionalInfo{value: val, isSet: true}
}

func (v NullableEventNotifyDataAdditionalInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventNotifyDataAdditionalInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


