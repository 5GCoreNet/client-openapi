/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// TimePeriod Contains the periodicity for the defined usage monitoring data limits.
type TimePeriod struct {
	Period Periodicity `json:"period"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxNumPeriod *int32 `json:"maxNumPeriod,omitempty"`
}

// NewTimePeriod instantiates a new TimePeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimePeriod(period Periodicity) *TimePeriod {
	this := TimePeriod{}
	this.Period = period
	return &this
}

// NewTimePeriodWithDefaults instantiates a new TimePeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimePeriodWithDefaults() *TimePeriod {
	this := TimePeriod{}
	return &this
}

// GetPeriod returns the Period field value
func (o *TimePeriod) GetPeriod() Periodicity {
	if o == nil {
		var ret Periodicity
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *TimePeriod) GetPeriodOk() (*Periodicity, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *TimePeriod) SetPeriod(v Periodicity) {
	o.Period = v
}

// GetMaxNumPeriod returns the MaxNumPeriod field value if set, zero value otherwise.
func (o *TimePeriod) GetMaxNumPeriod() int32 {
	if o == nil || isNil(o.MaxNumPeriod) {
		var ret int32
		return ret
	}
	return *o.MaxNumPeriod
}

// GetMaxNumPeriodOk returns a tuple with the MaxNumPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimePeriod) GetMaxNumPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.MaxNumPeriod) {
    return nil, false
	}
	return o.MaxNumPeriod, true
}

// HasMaxNumPeriod returns a boolean if a field has been set.
func (o *TimePeriod) HasMaxNumPeriod() bool {
	if o != nil && !isNil(o.MaxNumPeriod) {
		return true
	}

	return false
}

// SetMaxNumPeriod gets a reference to the given int32 and assigns it to the MaxNumPeriod field.
func (o *TimePeriod) SetMaxNumPeriod(v int32) {
	o.MaxNumPeriod = &v
}

func (o TimePeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["period"] = o.Period
	}
	if !isNil(o.MaxNumPeriod) {
		toSerialize["maxNumPeriod"] = o.MaxNumPeriod
	}
	return json.Marshal(toSerialize)
}

type NullableTimePeriod struct {
	value *TimePeriod
	isSet bool
}

func (v NullableTimePeriod) Get() *TimePeriod {
	return v.value
}

func (v *NullableTimePeriod) Set(val *TimePeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableTimePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableTimePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimePeriod(val *TimePeriod) *NullableTimePeriod {
	return &NullableTimePeriod{value: val, isSet: true}
}

func (v NullableTimePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


