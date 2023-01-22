/*
Nnwdaf_MLModelProvision

Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_MLModelProvision

import (
	"encoding/json"
)

// ScheduledCommunicationTime Identifies time and day of the week when the UE is available for communication.
type ScheduledCommunicationTime struct {
	// Identifies the day(s) of the week. If absent, it indicates every day of the week. 
	DaysOfWeek []int32 `json:"daysOfWeek,omitempty"`
	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).  
	TimeOfDayStart *string `json:"timeOfDayStart,omitempty"`
	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).  
	TimeOfDayEnd *string `json:"timeOfDayEnd,omitempty"`
}

// NewScheduledCommunicationTime instantiates a new ScheduledCommunicationTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledCommunicationTime() *ScheduledCommunicationTime {
	this := ScheduledCommunicationTime{}
	return &this
}

// NewScheduledCommunicationTimeWithDefaults instantiates a new ScheduledCommunicationTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledCommunicationTimeWithDefaults() *ScheduledCommunicationTime {
	this := ScheduledCommunicationTime{}
	return &this
}

// GetDaysOfWeek returns the DaysOfWeek field value if set, zero value otherwise.
func (o *ScheduledCommunicationTime) GetDaysOfWeek() []int32 {
	if o == nil || isNil(o.DaysOfWeek) {
		var ret []int32
		return ret
	}
	return o.DaysOfWeek
}

// GetDaysOfWeekOk returns a tuple with the DaysOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledCommunicationTime) GetDaysOfWeekOk() ([]int32, bool) {
	if o == nil || isNil(o.DaysOfWeek) {
    return nil, false
	}
	return o.DaysOfWeek, true
}

// HasDaysOfWeek returns a boolean if a field has been set.
func (o *ScheduledCommunicationTime) HasDaysOfWeek() bool {
	if o != nil && !isNil(o.DaysOfWeek) {
		return true
	}

	return false
}

// SetDaysOfWeek gets a reference to the given []int32 and assigns it to the DaysOfWeek field.
func (o *ScheduledCommunicationTime) SetDaysOfWeek(v []int32) {
	o.DaysOfWeek = v
}

// GetTimeOfDayStart returns the TimeOfDayStart field value if set, zero value otherwise.
func (o *ScheduledCommunicationTime) GetTimeOfDayStart() string {
	if o == nil || isNil(o.TimeOfDayStart) {
		var ret string
		return ret
	}
	return *o.TimeOfDayStart
}

// GetTimeOfDayStartOk returns a tuple with the TimeOfDayStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledCommunicationTime) GetTimeOfDayStartOk() (*string, bool) {
	if o == nil || isNil(o.TimeOfDayStart) {
    return nil, false
	}
	return o.TimeOfDayStart, true
}

// HasTimeOfDayStart returns a boolean if a field has been set.
func (o *ScheduledCommunicationTime) HasTimeOfDayStart() bool {
	if o != nil && !isNil(o.TimeOfDayStart) {
		return true
	}

	return false
}

// SetTimeOfDayStart gets a reference to the given string and assigns it to the TimeOfDayStart field.
func (o *ScheduledCommunicationTime) SetTimeOfDayStart(v string) {
	o.TimeOfDayStart = &v
}

// GetTimeOfDayEnd returns the TimeOfDayEnd field value if set, zero value otherwise.
func (o *ScheduledCommunicationTime) GetTimeOfDayEnd() string {
	if o == nil || isNil(o.TimeOfDayEnd) {
		var ret string
		return ret
	}
	return *o.TimeOfDayEnd
}

// GetTimeOfDayEndOk returns a tuple with the TimeOfDayEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledCommunicationTime) GetTimeOfDayEndOk() (*string, bool) {
	if o == nil || isNil(o.TimeOfDayEnd) {
    return nil, false
	}
	return o.TimeOfDayEnd, true
}

// HasTimeOfDayEnd returns a boolean if a field has been set.
func (o *ScheduledCommunicationTime) HasTimeOfDayEnd() bool {
	if o != nil && !isNil(o.TimeOfDayEnd) {
		return true
	}

	return false
}

// SetTimeOfDayEnd gets a reference to the given string and assigns it to the TimeOfDayEnd field.
func (o *ScheduledCommunicationTime) SetTimeOfDayEnd(v string) {
	o.TimeOfDayEnd = &v
}

func (o ScheduledCommunicationTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DaysOfWeek) {
		toSerialize["daysOfWeek"] = o.DaysOfWeek
	}
	if !isNil(o.TimeOfDayStart) {
		toSerialize["timeOfDayStart"] = o.TimeOfDayStart
	}
	if !isNil(o.TimeOfDayEnd) {
		toSerialize["timeOfDayEnd"] = o.TimeOfDayEnd
	}
	return json.Marshal(toSerialize)
}

type NullableScheduledCommunicationTime struct {
	value *ScheduledCommunicationTime
	isSet bool
}

func (v NullableScheduledCommunicationTime) Get() *ScheduledCommunicationTime {
	return v.value
}

func (v *NullableScheduledCommunicationTime) Set(val *ScheduledCommunicationTime) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledCommunicationTime) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledCommunicationTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledCommunicationTime(val *ScheduledCommunicationTime) *NullableScheduledCommunicationTime {
	return &NullableScheduledCommunicationTime{value: val, isSet: true}
}

func (v NullableScheduledCommunicationTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledCommunicationTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

