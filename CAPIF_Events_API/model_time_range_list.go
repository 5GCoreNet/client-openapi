/*
CAPIF_Events_API

API for event subscription management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Events_API

import (
	"encoding/json"
	"time"
)

// TimeRangeList Represents the time range during which the invocation of a service API is allowed by the API invoker. 
type TimeRangeList struct {
	// string with format \"date-time\" as defined in OpenAPI.
	StartTime *time.Time `json:"startTime,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	StopTime *time.Time `json:"stopTime,omitempty"`
}

// NewTimeRangeList instantiates a new TimeRangeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeRangeList() *TimeRangeList {
	this := TimeRangeList{}
	return &this
}

// NewTimeRangeListWithDefaults instantiates a new TimeRangeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeRangeListWithDefaults() *TimeRangeList {
	this := TimeRangeList{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TimeRangeList) GetStartTime() time.Time {
	if o == nil || isNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeRangeList) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTime) {
    return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TimeRangeList) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *TimeRangeList) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStopTime returns the StopTime field value if set, zero value otherwise.
func (o *TimeRangeList) GetStopTime() time.Time {
	if o == nil || isNil(o.StopTime) {
		var ret time.Time
		return ret
	}
	return *o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeRangeList) GetStopTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.StopTime) {
    return nil, false
	}
	return o.StopTime, true
}

// HasStopTime returns a boolean if a field has been set.
func (o *TimeRangeList) HasStopTime() bool {
	if o != nil && !isNil(o.StopTime) {
		return true
	}

	return false
}

// SetStopTime gets a reference to the given time.Time and assigns it to the StopTime field.
func (o *TimeRangeList) SetStopTime(v time.Time) {
	o.StopTime = &v
}

func (o TimeRangeList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !isNil(o.StopTime) {
		toSerialize["stopTime"] = o.StopTime
	}
	return json.Marshal(toSerialize)
}

type NullableTimeRangeList struct {
	value *TimeRangeList
	isSet bool
}

func (v NullableTimeRangeList) Get() *TimeRangeList {
	return v.value
}

func (v *NullableTimeRangeList) Set(val *TimeRangeList) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeRangeList) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeRangeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeRangeList(val *TimeRangeList) *NullableTimeRangeList {
	return &NullableTimeRangeList{value: val, isSet: true}
}

func (v NullableTimeRangeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeRangeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


