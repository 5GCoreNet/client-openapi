/*
3gpp-chargeable-party

API for Chargeable Party management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ChargeableParty

import (
	"encoding/json"
)

// EventReport Represents an event report.
type EventReport struct {
	Event Event `json:"event"`
	AccumulatedUsage *AccumulatedUsage `json:"accumulatedUsage,omitempty"`
	// Identifies the IP flows that were sent during event subscription
	FlowIds []int32 `json:"flowIds,omitempty"`
}

// NewEventReport instantiates a new EventReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventReport(event Event) *EventReport {
	this := EventReport{}
	this.Event = event
	return &this
}

// NewEventReportWithDefaults instantiates a new EventReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventReportWithDefaults() *EventReport {
	this := EventReport{}
	return &this
}

// GetEvent returns the Event field value
func (o *EventReport) GetEvent() Event {
	if o == nil {
		var ret Event
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *EventReport) GetEventOk() (*Event, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *EventReport) SetEvent(v Event) {
	o.Event = v
}

// GetAccumulatedUsage returns the AccumulatedUsage field value if set, zero value otherwise.
func (o *EventReport) GetAccumulatedUsage() AccumulatedUsage {
	if o == nil || isNil(o.AccumulatedUsage) {
		var ret AccumulatedUsage
		return ret
	}
	return *o.AccumulatedUsage
}

// GetAccumulatedUsageOk returns a tuple with the AccumulatedUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReport) GetAccumulatedUsageOk() (*AccumulatedUsage, bool) {
	if o == nil || isNil(o.AccumulatedUsage) {
    return nil, false
	}
	return o.AccumulatedUsage, true
}

// HasAccumulatedUsage returns a boolean if a field has been set.
func (o *EventReport) HasAccumulatedUsage() bool {
	if o != nil && !isNil(o.AccumulatedUsage) {
		return true
	}

	return false
}

// SetAccumulatedUsage gets a reference to the given AccumulatedUsage and assigns it to the AccumulatedUsage field.
func (o *EventReport) SetAccumulatedUsage(v AccumulatedUsage) {
	o.AccumulatedUsage = &v
}

// GetFlowIds returns the FlowIds field value if set, zero value otherwise.
func (o *EventReport) GetFlowIds() []int32 {
	if o == nil || isNil(o.FlowIds) {
		var ret []int32
		return ret
	}
	return o.FlowIds
}

// GetFlowIdsOk returns a tuple with the FlowIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReport) GetFlowIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.FlowIds) {
    return nil, false
	}
	return o.FlowIds, true
}

// HasFlowIds returns a boolean if a field has been set.
func (o *EventReport) HasFlowIds() bool {
	if o != nil && !isNil(o.FlowIds) {
		return true
	}

	return false
}

// SetFlowIds gets a reference to the given []int32 and assigns it to the FlowIds field.
func (o *EventReport) SetFlowIds(v []int32) {
	o.FlowIds = v
}

func (o EventReport) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableEventReport struct {
	value *EventReport
	isSet bool
}

func (v NullableEventReport) Get() *EventReport {
	return v.value
}

func (v *NullableEventReport) Set(val *EventReport) {
	v.value = val
	v.isSet = true
}

func (v NullableEventReport) IsSet() bool {
	return v.isSet
}

func (v *NullableEventReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventReport(val *EventReport) *NullableEventReport {
	return &NullableEventReport{value: val, isSet: true}
}

func (v NullableEventReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


