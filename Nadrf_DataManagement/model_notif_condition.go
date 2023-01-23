/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// NotifCondition Condition (list of attributes in the NF Profile) to determine whether a notification must be sent by NRF 
type NotifCondition struct {
	MonitoredAttributes []string `json:"monitoredAttributes,omitempty"`
	UnmonitoredAttributes []string `json:"unmonitoredAttributes,omitempty"`
}

// NewNotifCondition instantiates a new NotifCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifCondition() *NotifCondition {
	this := NotifCondition{}
	return &this
}

// NewNotifConditionWithDefaults instantiates a new NotifCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifConditionWithDefaults() *NotifCondition {
	this := NotifCondition{}
	return &this
}

// GetMonitoredAttributes returns the MonitoredAttributes field value if set, zero value otherwise.
func (o *NotifCondition) GetMonitoredAttributes() []string {
	if o == nil || isNil(o.MonitoredAttributes) {
		var ret []string
		return ret
	}
	return o.MonitoredAttributes
}

// GetMonitoredAttributesOk returns a tuple with the MonitoredAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifCondition) GetMonitoredAttributesOk() ([]string, bool) {
	if o == nil || isNil(o.MonitoredAttributes) {
    return nil, false
	}
	return o.MonitoredAttributes, true
}

// HasMonitoredAttributes returns a boolean if a field has been set.
func (o *NotifCondition) HasMonitoredAttributes() bool {
	if o != nil && !isNil(o.MonitoredAttributes) {
		return true
	}

	return false
}

// SetMonitoredAttributes gets a reference to the given []string and assigns it to the MonitoredAttributes field.
func (o *NotifCondition) SetMonitoredAttributes(v []string) {
	o.MonitoredAttributes = v
}

// GetUnmonitoredAttributes returns the UnmonitoredAttributes field value if set, zero value otherwise.
func (o *NotifCondition) GetUnmonitoredAttributes() []string {
	if o == nil || isNil(o.UnmonitoredAttributes) {
		var ret []string
		return ret
	}
	return o.UnmonitoredAttributes
}

// GetUnmonitoredAttributesOk returns a tuple with the UnmonitoredAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifCondition) GetUnmonitoredAttributesOk() ([]string, bool) {
	if o == nil || isNil(o.UnmonitoredAttributes) {
    return nil, false
	}
	return o.UnmonitoredAttributes, true
}

// HasUnmonitoredAttributes returns a boolean if a field has been set.
func (o *NotifCondition) HasUnmonitoredAttributes() bool {
	if o != nil && !isNil(o.UnmonitoredAttributes) {
		return true
	}

	return false
}

// SetUnmonitoredAttributes gets a reference to the given []string and assigns it to the UnmonitoredAttributes field.
func (o *NotifCondition) SetUnmonitoredAttributes(v []string) {
	o.UnmonitoredAttributes = v
}

func (o NotifCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MonitoredAttributes) {
		toSerialize["monitoredAttributes"] = o.MonitoredAttributes
	}
	if !isNil(o.UnmonitoredAttributes) {
		toSerialize["unmonitoredAttributes"] = o.UnmonitoredAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableNotifCondition struct {
	value *NotifCondition
	isSet bool
}

func (v NullableNotifCondition) Get() *NotifCondition {
	return v.value
}

func (v *NullableNotifCondition) Set(val *NotifCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifCondition(val *NotifCondition) *NullableNotifCondition {
	return &NullableNotifCondition{value: val, isSet: true}
}

func (v NullableNotifCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


