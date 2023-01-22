/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// QosMonitoringReport QoS Monitoring reporting information.
type QosMonitoringReport struct {
	Flows []Flows `json:"flows,omitempty"`
	UlDelays []int32 `json:"ulDelays,omitempty"`
	DlDelays []int32 `json:"dlDelays,omitempty"`
	RtDelays []int32 `json:"rtDelays,omitempty"`
}

// NewQosMonitoringReport instantiates a new QosMonitoringReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosMonitoringReport() *QosMonitoringReport {
	this := QosMonitoringReport{}
	return &this
}

// NewQosMonitoringReportWithDefaults instantiates a new QosMonitoringReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosMonitoringReportWithDefaults() *QosMonitoringReport {
	this := QosMonitoringReport{}
	return &this
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetFlows() []Flows {
	if o == nil || isNil(o.Flows) {
		var ret []Flows
		return ret
	}
	return o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetFlowsOk() ([]Flows, bool) {
	if o == nil || isNil(o.Flows) {
    return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasFlows() bool {
	if o != nil && !isNil(o.Flows) {
		return true
	}

	return false
}

// SetFlows gets a reference to the given []Flows and assigns it to the Flows field.
func (o *QosMonitoringReport) SetFlows(v []Flows) {
	o.Flows = v
}

// GetUlDelays returns the UlDelays field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetUlDelays() []int32 {
	if o == nil || isNil(o.UlDelays) {
		var ret []int32
		return ret
	}
	return o.UlDelays
}

// GetUlDelaysOk returns a tuple with the UlDelays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetUlDelaysOk() ([]int32, bool) {
	if o == nil || isNil(o.UlDelays) {
    return nil, false
	}
	return o.UlDelays, true
}

// HasUlDelays returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasUlDelays() bool {
	if o != nil && !isNil(o.UlDelays) {
		return true
	}

	return false
}

// SetUlDelays gets a reference to the given []int32 and assigns it to the UlDelays field.
func (o *QosMonitoringReport) SetUlDelays(v []int32) {
	o.UlDelays = v
}

// GetDlDelays returns the DlDelays field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetDlDelays() []int32 {
	if o == nil || isNil(o.DlDelays) {
		var ret []int32
		return ret
	}
	return o.DlDelays
}

// GetDlDelaysOk returns a tuple with the DlDelays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetDlDelaysOk() ([]int32, bool) {
	if o == nil || isNil(o.DlDelays) {
    return nil, false
	}
	return o.DlDelays, true
}

// HasDlDelays returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasDlDelays() bool {
	if o != nil && !isNil(o.DlDelays) {
		return true
	}

	return false
}

// SetDlDelays gets a reference to the given []int32 and assigns it to the DlDelays field.
func (o *QosMonitoringReport) SetDlDelays(v []int32) {
	o.DlDelays = v
}

// GetRtDelays returns the RtDelays field value if set, zero value otherwise.
func (o *QosMonitoringReport) GetRtDelays() []int32 {
	if o == nil || isNil(o.RtDelays) {
		var ret []int32
		return ret
	}
	return o.RtDelays
}

// GetRtDelaysOk returns a tuple with the RtDelays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringReport) GetRtDelaysOk() ([]int32, bool) {
	if o == nil || isNil(o.RtDelays) {
    return nil, false
	}
	return o.RtDelays, true
}

// HasRtDelays returns a boolean if a field has been set.
func (o *QosMonitoringReport) HasRtDelays() bool {
	if o != nil && !isNil(o.RtDelays) {
		return true
	}

	return false
}

// SetRtDelays gets a reference to the given []int32 and assigns it to the RtDelays field.
func (o *QosMonitoringReport) SetRtDelays(v []int32) {
	o.RtDelays = v
}

func (o QosMonitoringReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Flows) {
		toSerialize["flows"] = o.Flows
	}
	if !isNil(o.UlDelays) {
		toSerialize["ulDelays"] = o.UlDelays
	}
	if !isNil(o.DlDelays) {
		toSerialize["dlDelays"] = o.DlDelays
	}
	if !isNil(o.RtDelays) {
		toSerialize["rtDelays"] = o.RtDelays
	}
	return json.Marshal(toSerialize)
}

type NullableQosMonitoringReport struct {
	value *QosMonitoringReport
	isSet bool
}

func (v NullableQosMonitoringReport) Get() *QosMonitoringReport {
	return v.value
}

func (v *NullableQosMonitoringReport) Set(val *QosMonitoringReport) {
	v.value = val
	v.isSet = true
}

func (v NullableQosMonitoringReport) IsSet() bool {
	return v.isSet
}

func (v *NullableQosMonitoringReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosMonitoringReport(val *QosMonitoringReport) *NullableQosMonitoringReport {
	return &NullableQosMonitoringReport{value: val, isSet: true}
}

func (v NullableQosMonitoringReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosMonitoringReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

