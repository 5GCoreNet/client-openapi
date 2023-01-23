/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
	"time"
)

// TimeWindow struct for TimeWindow
type TimeWindow struct {
	StartTime *time.Time `json:"startTime,omitempty"`
	EndTime *time.Time `json:"endTime,omitempty"`
}

// NewTimeWindow instantiates a new TimeWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeWindow() *TimeWindow {
	this := TimeWindow{}
	return &this
}

// NewTimeWindowWithDefaults instantiates a new TimeWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeWindowWithDefaults() *TimeWindow {
	this := TimeWindow{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TimeWindow) GetStartTime() time.Time {
	if o == nil || isNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWindow) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTime) {
    return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TimeWindow) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *TimeWindow) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TimeWindow) GetEndTime() time.Time {
	if o == nil || isNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWindow) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTime) {
    return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TimeWindow) HasEndTime() bool {
	if o != nil && !isNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *TimeWindow) SetEndTime(v time.Time) {
	o.EndTime = &v
}

func (o TimeWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !isNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	return json.Marshal(toSerialize)
}

type NullableTimeWindow struct {
	value *TimeWindow
	isSet bool
}

func (v NullableTimeWindow) Get() *TimeWindow {
	return v.value
}

func (v *NullableTimeWindow) Set(val *TimeWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeWindow(val *TimeWindow) *NullableTimeWindow {
	return &NullableTimeWindow{value: val, isSet: true}
}

func (v NullableTimeWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


