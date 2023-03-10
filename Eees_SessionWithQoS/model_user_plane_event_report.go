/*
EES Session with QoS API

API for EES Session with Qos service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_SessionWithQoS

import (
	"encoding/json"
)

// UserPlaneEventReport Represents an event report for user plane.
type UserPlaneEventReport struct {
	Event UserPlaneEvent `json:"event"`
	AccumulatedUsage *AccumulatedUsage `json:"accumulatedUsage,omitempty"`
	// Identifies the affected flows that were sent during event subscription. It might be omitted when the reported event applies to all the flows sent during the subscription. 
	FlowIds []int32 `json:"flowIds,omitempty"`
	// The currently applied QoS reference. Applicable for event QOS_NOT_GUARANTEED or SUCCESSFUL_RESOURCES_ALLOCATION.
	AppliedQosRef *string `json:"appliedQosRef,omitempty"`
	// Contains the QoS Monitoring Reporting information
	QosMonReports []QosMonitoringReport `json:"qosMonReports,omitempty"`
}

// NewUserPlaneEventReport instantiates a new UserPlaneEventReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPlaneEventReport(event UserPlaneEvent) *UserPlaneEventReport {
	this := UserPlaneEventReport{}
	this.Event = event
	return &this
}

// NewUserPlaneEventReportWithDefaults instantiates a new UserPlaneEventReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPlaneEventReportWithDefaults() *UserPlaneEventReport {
	this := UserPlaneEventReport{}
	return &this
}

// GetEvent returns the Event field value
func (o *UserPlaneEventReport) GetEvent() UserPlaneEvent {
	if o == nil {
		var ret UserPlaneEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *UserPlaneEventReport) GetEventOk() (*UserPlaneEvent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *UserPlaneEventReport) SetEvent(v UserPlaneEvent) {
	o.Event = v
}

// GetAccumulatedUsage returns the AccumulatedUsage field value if set, zero value otherwise.
func (o *UserPlaneEventReport) GetAccumulatedUsage() AccumulatedUsage {
	if o == nil || isNil(o.AccumulatedUsage) {
		var ret AccumulatedUsage
		return ret
	}
	return *o.AccumulatedUsage
}

// GetAccumulatedUsageOk returns a tuple with the AccumulatedUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPlaneEventReport) GetAccumulatedUsageOk() (*AccumulatedUsage, bool) {
	if o == nil || isNil(o.AccumulatedUsage) {
    return nil, false
	}
	return o.AccumulatedUsage, true
}

// HasAccumulatedUsage returns a boolean if a field has been set.
func (o *UserPlaneEventReport) HasAccumulatedUsage() bool {
	if o != nil && !isNil(o.AccumulatedUsage) {
		return true
	}

	return false
}

// SetAccumulatedUsage gets a reference to the given AccumulatedUsage and assigns it to the AccumulatedUsage field.
func (o *UserPlaneEventReport) SetAccumulatedUsage(v AccumulatedUsage) {
	o.AccumulatedUsage = &v
}

// GetFlowIds returns the FlowIds field value if set, zero value otherwise.
func (o *UserPlaneEventReport) GetFlowIds() []int32 {
	if o == nil || isNil(o.FlowIds) {
		var ret []int32
		return ret
	}
	return o.FlowIds
}

// GetFlowIdsOk returns a tuple with the FlowIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPlaneEventReport) GetFlowIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.FlowIds) {
    return nil, false
	}
	return o.FlowIds, true
}

// HasFlowIds returns a boolean if a field has been set.
func (o *UserPlaneEventReport) HasFlowIds() bool {
	if o != nil && !isNil(o.FlowIds) {
		return true
	}

	return false
}

// SetFlowIds gets a reference to the given []int32 and assigns it to the FlowIds field.
func (o *UserPlaneEventReport) SetFlowIds(v []int32) {
	o.FlowIds = v
}

// GetAppliedQosRef returns the AppliedQosRef field value if set, zero value otherwise.
func (o *UserPlaneEventReport) GetAppliedQosRef() string {
	if o == nil || isNil(o.AppliedQosRef) {
		var ret string
		return ret
	}
	return *o.AppliedQosRef
}

// GetAppliedQosRefOk returns a tuple with the AppliedQosRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPlaneEventReport) GetAppliedQosRefOk() (*string, bool) {
	if o == nil || isNil(o.AppliedQosRef) {
    return nil, false
	}
	return o.AppliedQosRef, true
}

// HasAppliedQosRef returns a boolean if a field has been set.
func (o *UserPlaneEventReport) HasAppliedQosRef() bool {
	if o != nil && !isNil(o.AppliedQosRef) {
		return true
	}

	return false
}

// SetAppliedQosRef gets a reference to the given string and assigns it to the AppliedQosRef field.
func (o *UserPlaneEventReport) SetAppliedQosRef(v string) {
	o.AppliedQosRef = &v
}

// GetQosMonReports returns the QosMonReports field value if set, zero value otherwise.
func (o *UserPlaneEventReport) GetQosMonReports() []QosMonitoringReport {
	if o == nil || isNil(o.QosMonReports) {
		var ret []QosMonitoringReport
		return ret
	}
	return o.QosMonReports
}

// GetQosMonReportsOk returns a tuple with the QosMonReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPlaneEventReport) GetQosMonReportsOk() ([]QosMonitoringReport, bool) {
	if o == nil || isNil(o.QosMonReports) {
    return nil, false
	}
	return o.QosMonReports, true
}

// HasQosMonReports returns a boolean if a field has been set.
func (o *UserPlaneEventReport) HasQosMonReports() bool {
	if o != nil && !isNil(o.QosMonReports) {
		return true
	}

	return false
}

// SetQosMonReports gets a reference to the given []QosMonitoringReport and assigns it to the QosMonReports field.
func (o *UserPlaneEventReport) SetQosMonReports(v []QosMonitoringReport) {
	o.QosMonReports = v
}

func (o UserPlaneEventReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event"] = o.Event
	}
	if !isNil(o.AccumulatedUsage) {
		toSerialize["accumulatedUsage"] = o.AccumulatedUsage
	}
	if !isNil(o.FlowIds) {
		toSerialize["flowIds"] = o.FlowIds
	}
	if !isNil(o.AppliedQosRef) {
		toSerialize["appliedQosRef"] = o.AppliedQosRef
	}
	if !isNil(o.QosMonReports) {
		toSerialize["qosMonReports"] = o.QosMonReports
	}
	return json.Marshal(toSerialize)
}

type NullableUserPlaneEventReport struct {
	value *UserPlaneEventReport
	isSet bool
}

func (v NullableUserPlaneEventReport) Get() *UserPlaneEventReport {
	return v.value
}

func (v *NullableUserPlaneEventReport) Set(val *UserPlaneEventReport) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPlaneEventReport) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPlaneEventReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPlaneEventReport(val *UserPlaneEventReport) *NullableUserPlaneEventReport {
	return &NullableUserPlaneEventReport{value: val, isSet: true}
}

func (v NullableUserPlaneEventReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPlaneEventReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


