/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MonitoringEvent

import (
	"encoding/json"
)

// MonitoringEventReports Represents a set of event monitoring reports.
type MonitoringEventReports struct {
	MonitoringEventReports []MonitoringEventReport `json:"monitoringEventReports"`
}

// NewMonitoringEventReports instantiates a new MonitoringEventReports object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringEventReports(monitoringEventReports []MonitoringEventReport) *MonitoringEventReports {
	this := MonitoringEventReports{}
	this.MonitoringEventReports = monitoringEventReports
	return &this
}

// NewMonitoringEventReportsWithDefaults instantiates a new MonitoringEventReports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringEventReportsWithDefaults() *MonitoringEventReports {
	this := MonitoringEventReports{}
	return &this
}

// GetMonitoringEventReports returns the MonitoringEventReports field value
func (o *MonitoringEventReports) GetMonitoringEventReports() []MonitoringEventReport {
	if o == nil {
		var ret []MonitoringEventReport
		return ret
	}

	return o.MonitoringEventReports
}

// GetMonitoringEventReportsOk returns a tuple with the MonitoringEventReports field value
// and a boolean to check if the value has been set.
func (o *MonitoringEventReports) GetMonitoringEventReportsOk() ([]MonitoringEventReport, bool) {
	if o == nil {
    return nil, false
	}
	return o.MonitoringEventReports, true
}

// SetMonitoringEventReports sets field value
func (o *MonitoringEventReports) SetMonitoringEventReports(v []MonitoringEventReport) {
	o.MonitoringEventReports = v
}

func (o MonitoringEventReports) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["monitoringEventReports"] = o.MonitoringEventReports
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringEventReports struct {
	value *MonitoringEventReports
	isSet bool
}

func (v NullableMonitoringEventReports) Get() *MonitoringEventReports {
	return v.value
}

func (v *NullableMonitoringEventReports) Set(val *MonitoringEventReports) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringEventReports) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringEventReports) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringEventReports(val *MonitoringEventReports) *NullableMonitoringEventReports {
	return &NullableMonitoringEventReports{value: val, isSet: true}
}

func (v NullableMonitoringEventReports) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringEventReports) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

