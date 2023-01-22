/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
	"time"
)

// TimeWindow1 struct for TimeWindow1
type TimeWindow1 struct {
	StartTime *time.Time `json:"startTime,omitempty"`
	EndTime *time.Time `json:"endTime,omitempty"`
}

// NewTimeWindow1 instantiates a new TimeWindow1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeWindow1() *TimeWindow1 {
	this := TimeWindow1{}
	return &this
}

// NewTimeWindow1WithDefaults instantiates a new TimeWindow1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeWindow1WithDefaults() *TimeWindow1 {
	this := TimeWindow1{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TimeWindow1) GetStartTime() time.Time {
	if o == nil || isNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWindow1) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTime) {
    return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TimeWindow1) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *TimeWindow1) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TimeWindow1) GetEndTime() time.Time {
	if o == nil || isNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWindow1) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTime) {
    return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TimeWindow1) HasEndTime() bool {
	if o != nil && !isNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *TimeWindow1) SetEndTime(v time.Time) {
	o.EndTime = &v
}

func (o TimeWindow1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !isNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	return json.Marshal(toSerialize)
}

type NullableTimeWindow1 struct {
	value *TimeWindow1
	isSet bool
}

func (v NullableTimeWindow1) Get() *TimeWindow1 {
	return v.value
}

func (v *NullableTimeWindow1) Set(val *TimeWindow1) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeWindow1) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeWindow1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeWindow1(val *TimeWindow1) *NullableTimeWindow1 {
	return &NullableTimeWindow1{value: val, isSet: true}
}

func (v NullableTimeWindow1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeWindow1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

