/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_EE

import (
	"encoding/json"
)

// FailedMonitoringConfiguration Contains the event type and failed cause of the failed Monitoring Configuration in the EE subscription
type FailedMonitoringConfiguration struct {
	EventType EventType `json:"eventType"`
	FailedCause FailedCause `json:"failedCause"`
}

// NewFailedMonitoringConfiguration instantiates a new FailedMonitoringConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailedMonitoringConfiguration(eventType EventType, failedCause FailedCause) *FailedMonitoringConfiguration {
	this := FailedMonitoringConfiguration{}
	this.EventType = eventType
	this.FailedCause = failedCause
	return &this
}

// NewFailedMonitoringConfigurationWithDefaults instantiates a new FailedMonitoringConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailedMonitoringConfigurationWithDefaults() *FailedMonitoringConfiguration {
	this := FailedMonitoringConfiguration{}
	return &this
}

// GetEventType returns the EventType field value
func (o *FailedMonitoringConfiguration) GetEventType() EventType {
	if o == nil {
		var ret EventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *FailedMonitoringConfiguration) GetEventTypeOk() (*EventType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *FailedMonitoringConfiguration) SetEventType(v EventType) {
	o.EventType = v
}

// GetFailedCause returns the FailedCause field value
func (o *FailedMonitoringConfiguration) GetFailedCause() FailedCause {
	if o == nil {
		var ret FailedCause
		return ret
	}

	return o.FailedCause
}

// GetFailedCauseOk returns a tuple with the FailedCause field value
// and a boolean to check if the value has been set.
func (o *FailedMonitoringConfiguration) GetFailedCauseOk() (*FailedCause, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FailedCause, true
}

// SetFailedCause sets field value
func (o *FailedMonitoringConfiguration) SetFailedCause(v FailedCause) {
	o.FailedCause = v
}

func (o FailedMonitoringConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["failedCause"] = o.FailedCause
	}
	return json.Marshal(toSerialize)
}

type NullableFailedMonitoringConfiguration struct {
	value *FailedMonitoringConfiguration
	isSet bool
}

func (v NullableFailedMonitoringConfiguration) Get() *FailedMonitoringConfiguration {
	return v.value
}

func (v *NullableFailedMonitoringConfiguration) Set(val *FailedMonitoringConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableFailedMonitoringConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableFailedMonitoringConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailedMonitoringConfiguration(val *FailedMonitoringConfiguration) *NullableFailedMonitoringConfiguration {
	return &NullableFailedMonitoringConfiguration{value: val, isSet: true}
}

func (v NullableFailedMonitoringConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailedMonitoringConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


