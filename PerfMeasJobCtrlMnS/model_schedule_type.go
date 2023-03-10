/*
TS 28.550 Performance Measurement Job Control Service

OAS 3.0.1 specification of the Performance Measurement Job Control Service @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 16.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PerfMeasJobCtrlMnS

import (
	"encoding/json"
)

// ScheduleType struct for ScheduleType
type ScheduleType struct {
	ScheduleOption *ScheduleOptionType `json:"scheduleOption,omitempty"`
	DailySchedule []TimeIntervalType `json:"dailySchedule,omitempty"`
	WeeklySchedule []ScheduleOfDayType `json:"weeklySchedule,omitempty"`
}

// NewScheduleType instantiates a new ScheduleType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleType() *ScheduleType {
	this := ScheduleType{}
	return &this
}

// NewScheduleTypeWithDefaults instantiates a new ScheduleType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleTypeWithDefaults() *ScheduleType {
	this := ScheduleType{}
	return &this
}

// GetScheduleOption returns the ScheduleOption field value if set, zero value otherwise.
func (o *ScheduleType) GetScheduleOption() ScheduleOptionType {
	if o == nil || isNil(o.ScheduleOption) {
		var ret ScheduleOptionType
		return ret
	}
	return *o.ScheduleOption
}

// GetScheduleOptionOk returns a tuple with the ScheduleOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleType) GetScheduleOptionOk() (*ScheduleOptionType, bool) {
	if o == nil || isNil(o.ScheduleOption) {
    return nil, false
	}
	return o.ScheduleOption, true
}

// HasScheduleOption returns a boolean if a field has been set.
func (o *ScheduleType) HasScheduleOption() bool {
	if o != nil && !isNil(o.ScheduleOption) {
		return true
	}

	return false
}

// SetScheduleOption gets a reference to the given ScheduleOptionType and assigns it to the ScheduleOption field.
func (o *ScheduleType) SetScheduleOption(v ScheduleOptionType) {
	o.ScheduleOption = &v
}

// GetDailySchedule returns the DailySchedule field value if set, zero value otherwise.
func (o *ScheduleType) GetDailySchedule() []TimeIntervalType {
	if o == nil || isNil(o.DailySchedule) {
		var ret []TimeIntervalType
		return ret
	}
	return o.DailySchedule
}

// GetDailyScheduleOk returns a tuple with the DailySchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleType) GetDailyScheduleOk() ([]TimeIntervalType, bool) {
	if o == nil || isNil(o.DailySchedule) {
    return nil, false
	}
	return o.DailySchedule, true
}

// HasDailySchedule returns a boolean if a field has been set.
func (o *ScheduleType) HasDailySchedule() bool {
	if o != nil && !isNil(o.DailySchedule) {
		return true
	}

	return false
}

// SetDailySchedule gets a reference to the given []TimeIntervalType and assigns it to the DailySchedule field.
func (o *ScheduleType) SetDailySchedule(v []TimeIntervalType) {
	o.DailySchedule = v
}

// GetWeeklySchedule returns the WeeklySchedule field value if set, zero value otherwise.
func (o *ScheduleType) GetWeeklySchedule() []ScheduleOfDayType {
	if o == nil || isNil(o.WeeklySchedule) {
		var ret []ScheduleOfDayType
		return ret
	}
	return o.WeeklySchedule
}

// GetWeeklyScheduleOk returns a tuple with the WeeklySchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleType) GetWeeklyScheduleOk() ([]ScheduleOfDayType, bool) {
	if o == nil || isNil(o.WeeklySchedule) {
    return nil, false
	}
	return o.WeeklySchedule, true
}

// HasWeeklySchedule returns a boolean if a field has been set.
func (o *ScheduleType) HasWeeklySchedule() bool {
	if o != nil && !isNil(o.WeeklySchedule) {
		return true
	}

	return false
}

// SetWeeklySchedule gets a reference to the given []ScheduleOfDayType and assigns it to the WeeklySchedule field.
func (o *ScheduleType) SetWeeklySchedule(v []ScheduleOfDayType) {
	o.WeeklySchedule = v
}

func (o ScheduleType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ScheduleOption) {
		toSerialize["scheduleOption"] = o.ScheduleOption
	}
	if !isNil(o.DailySchedule) {
		toSerialize["dailySchedule"] = o.DailySchedule
	}
	if !isNil(o.WeeklySchedule) {
		toSerialize["weeklySchedule"] = o.WeeklySchedule
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleType struct {
	value *ScheduleType
	isSet bool
}

func (v NullableScheduleType) Get() *ScheduleType {
	return v.value
}

func (v *NullableScheduleType) Set(val *ScheduleType) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleType) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleType(val *ScheduleType) *NullableScheduleType {
	return &NullableScheduleType{value: val, isSet: true}
}

func (v NullableScheduleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


