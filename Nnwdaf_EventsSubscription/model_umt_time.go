/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_EventsSubscription

import (
	"encoding/json"
)

// UmtTime struct for UmtTime
type UmtTime struct {
	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).  
	TimeOfDay string `json:"timeOfDay"`
	// integer between and including 1 and 7 denoting a weekday. 1 shall indicate Monday, and the subsequent weekdays  shall be indicated with the next higher numbers. 7 shall indicate Sunday. 
	DayOfWeek int32 `json:"dayOfWeek"`
}

// NewUmtTime instantiates a new UmtTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmtTime(timeOfDay string, dayOfWeek int32) *UmtTime {
	this := UmtTime{}
	this.TimeOfDay = timeOfDay
	this.DayOfWeek = dayOfWeek
	return &this
}

// NewUmtTimeWithDefaults instantiates a new UmtTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmtTimeWithDefaults() *UmtTime {
	this := UmtTime{}
	return &this
}

// GetTimeOfDay returns the TimeOfDay field value
func (o *UmtTime) GetTimeOfDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeOfDay
}

// GetTimeOfDayOk returns a tuple with the TimeOfDay field value
// and a boolean to check if the value has been set.
func (o *UmtTime) GetTimeOfDayOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeOfDay, true
}

// SetTimeOfDay sets field value
func (o *UmtTime) SetTimeOfDay(v string) {
	o.TimeOfDay = v
}

// GetDayOfWeek returns the DayOfWeek field value
func (o *UmtTime) GetDayOfWeek() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value
// and a boolean to check if the value has been set.
func (o *UmtTime) GetDayOfWeekOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DayOfWeek, true
}

// SetDayOfWeek sets field value
func (o *UmtTime) SetDayOfWeek(v int32) {
	o.DayOfWeek = v
}

func (o UmtTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timeOfDay"] = o.TimeOfDay
	}
	if true {
		toSerialize["dayOfWeek"] = o.DayOfWeek
	}
	return json.Marshal(toSerialize)
}

type NullableUmtTime struct {
	value *UmtTime
	isSet bool
}

func (v NullableUmtTime) Get() *UmtTime {
	return v.value
}

func (v *NullableUmtTime) Set(val *UmtTime) {
	v.value = val
	v.isSet = true
}

func (v NullableUmtTime) IsSet() bool {
	return v.isSet
}

func (v *NullableUmtTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmtTime(val *UmtTime) *NullableUmtTime {
	return &NullableUmtTime{value: val, isSet: true}
}

func (v NullableUmtTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmtTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


