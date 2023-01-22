/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// EsNotAllowedTimePeriod struct for EsNotAllowedTimePeriod
type EsNotAllowedTimePeriod struct {
	StartTimeandendTime *string `json:"startTimeandendTime,omitempty"`
	PeriodOfDay *string `json:"periodOfDay,omitempty"`
	DaysOfWeekList *string `json:"daysOfWeekList,omitempty"`
	Listoftimeperiods *string `json:"listoftimeperiods,omitempty"`
}

// NewEsNotAllowedTimePeriod instantiates a new EsNotAllowedTimePeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsNotAllowedTimePeriod() *EsNotAllowedTimePeriod {
	this := EsNotAllowedTimePeriod{}
	return &this
}

// NewEsNotAllowedTimePeriodWithDefaults instantiates a new EsNotAllowedTimePeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsNotAllowedTimePeriodWithDefaults() *EsNotAllowedTimePeriod {
	this := EsNotAllowedTimePeriod{}
	return &this
}

// GetStartTimeandendTime returns the StartTimeandendTime field value if set, zero value otherwise.
func (o *EsNotAllowedTimePeriod) GetStartTimeandendTime() string {
	if o == nil || isNil(o.StartTimeandendTime) {
		var ret string
		return ret
	}
	return *o.StartTimeandendTime
}

// GetStartTimeandendTimeOk returns a tuple with the StartTimeandendTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsNotAllowedTimePeriod) GetStartTimeandendTimeOk() (*string, bool) {
	if o == nil || isNil(o.StartTimeandendTime) {
    return nil, false
	}
	return o.StartTimeandendTime, true
}

// HasStartTimeandendTime returns a boolean if a field has been set.
func (o *EsNotAllowedTimePeriod) HasStartTimeandendTime() bool {
	if o != nil && !isNil(o.StartTimeandendTime) {
		return true
	}

	return false
}

// SetStartTimeandendTime gets a reference to the given string and assigns it to the StartTimeandendTime field.
func (o *EsNotAllowedTimePeriod) SetStartTimeandendTime(v string) {
	o.StartTimeandendTime = &v
}

// GetPeriodOfDay returns the PeriodOfDay field value if set, zero value otherwise.
func (o *EsNotAllowedTimePeriod) GetPeriodOfDay() string {
	if o == nil || isNil(o.PeriodOfDay) {
		var ret string
		return ret
	}
	return *o.PeriodOfDay
}

// GetPeriodOfDayOk returns a tuple with the PeriodOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsNotAllowedTimePeriod) GetPeriodOfDayOk() (*string, bool) {
	if o == nil || isNil(o.PeriodOfDay) {
    return nil, false
	}
	return o.PeriodOfDay, true
}

// HasPeriodOfDay returns a boolean if a field has been set.
func (o *EsNotAllowedTimePeriod) HasPeriodOfDay() bool {
	if o != nil && !isNil(o.PeriodOfDay) {
		return true
	}

	return false
}

// SetPeriodOfDay gets a reference to the given string and assigns it to the PeriodOfDay field.
func (o *EsNotAllowedTimePeriod) SetPeriodOfDay(v string) {
	o.PeriodOfDay = &v
}

// GetDaysOfWeekList returns the DaysOfWeekList field value if set, zero value otherwise.
func (o *EsNotAllowedTimePeriod) GetDaysOfWeekList() string {
	if o == nil || isNil(o.DaysOfWeekList) {
		var ret string
		return ret
	}
	return *o.DaysOfWeekList
}

// GetDaysOfWeekListOk returns a tuple with the DaysOfWeekList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsNotAllowedTimePeriod) GetDaysOfWeekListOk() (*string, bool) {
	if o == nil || isNil(o.DaysOfWeekList) {
    return nil, false
	}
	return o.DaysOfWeekList, true
}

// HasDaysOfWeekList returns a boolean if a field has been set.
func (o *EsNotAllowedTimePeriod) HasDaysOfWeekList() bool {
	if o != nil && !isNil(o.DaysOfWeekList) {
		return true
	}

	return false
}

// SetDaysOfWeekList gets a reference to the given string and assigns it to the DaysOfWeekList field.
func (o *EsNotAllowedTimePeriod) SetDaysOfWeekList(v string) {
	o.DaysOfWeekList = &v
}

// GetListoftimeperiods returns the Listoftimeperiods field value if set, zero value otherwise.
func (o *EsNotAllowedTimePeriod) GetListoftimeperiods() string {
	if o == nil || isNil(o.Listoftimeperiods) {
		var ret string
		return ret
	}
	return *o.Listoftimeperiods
}

// GetListoftimeperiodsOk returns a tuple with the Listoftimeperiods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsNotAllowedTimePeriod) GetListoftimeperiodsOk() (*string, bool) {
	if o == nil || isNil(o.Listoftimeperiods) {
    return nil, false
	}
	return o.Listoftimeperiods, true
}

// HasListoftimeperiods returns a boolean if a field has been set.
func (o *EsNotAllowedTimePeriod) HasListoftimeperiods() bool {
	if o != nil && !isNil(o.Listoftimeperiods) {
		return true
	}

	return false
}

// SetListoftimeperiods gets a reference to the given string and assigns it to the Listoftimeperiods field.
func (o *EsNotAllowedTimePeriod) SetListoftimeperiods(v string) {
	o.Listoftimeperiods = &v
}

func (o EsNotAllowedTimePeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTimeandendTime) {
		toSerialize["startTimeandendTime"] = o.StartTimeandendTime
	}
	if !isNil(o.PeriodOfDay) {
		toSerialize["periodOfDay"] = o.PeriodOfDay
	}
	if !isNil(o.DaysOfWeekList) {
		toSerialize["daysOfWeekList"] = o.DaysOfWeekList
	}
	if !isNil(o.Listoftimeperiods) {
		toSerialize["listoftimeperiods"] = o.Listoftimeperiods
	}
	return json.Marshal(toSerialize)
}

type NullableEsNotAllowedTimePeriod struct {
	value *EsNotAllowedTimePeriod
	isSet bool
}

func (v NullableEsNotAllowedTimePeriod) Get() *EsNotAllowedTimePeriod {
	return v.value
}

func (v *NullableEsNotAllowedTimePeriod) Set(val *EsNotAllowedTimePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableEsNotAllowedTimePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableEsNotAllowedTimePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsNotAllowedTimePeriod(val *EsNotAllowedTimePeriod) *NullableEsNotAllowedTimePeriod {
	return &NullableEsNotAllowedTimePeriod{value: val, isSet: true}
}

func (v NullableEsNotAllowedTimePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsNotAllowedTimePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

