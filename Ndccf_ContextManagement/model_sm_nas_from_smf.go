/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_ContextManagement

import (
	"encoding/json"
	"time"
)

// SmNasFromSmf Represents information on the SM congestion control applied SM NAS messages that SMF sends to UE for PDU Session. 
type SmNasFromSmf struct {
	SmNasType string `json:"smNasType"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
	// indicating a time in seconds.
	BackoffTimer int32 `json:"backoffTimer"`
	AppliedSmccType AppliedSmccType `json:"appliedSmccType"`
}

// NewSmNasFromSmf instantiates a new SmNasFromSmf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmNasFromSmf(smNasType string, timeStamp time.Time, backoffTimer int32, appliedSmccType AppliedSmccType) *SmNasFromSmf {
	this := SmNasFromSmf{}
	this.SmNasType = smNasType
	this.TimeStamp = timeStamp
	this.BackoffTimer = backoffTimer
	this.AppliedSmccType = appliedSmccType
	return &this
}

// NewSmNasFromSmfWithDefaults instantiates a new SmNasFromSmf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmNasFromSmfWithDefaults() *SmNasFromSmf {
	this := SmNasFromSmf{}
	return &this
}

// GetSmNasType returns the SmNasType field value
func (o *SmNasFromSmf) GetSmNasType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SmNasType
}

// GetSmNasTypeOk returns a tuple with the SmNasType field value
// and a boolean to check if the value has been set.
func (o *SmNasFromSmf) GetSmNasTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SmNasType, true
}

// SetSmNasType sets field value
func (o *SmNasFromSmf) SetSmNasType(v string) {
	o.SmNasType = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *SmNasFromSmf) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *SmNasFromSmf) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *SmNasFromSmf) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

// GetBackoffTimer returns the BackoffTimer field value
func (o *SmNasFromSmf) GetBackoffTimer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BackoffTimer
}

// GetBackoffTimerOk returns a tuple with the BackoffTimer field value
// and a boolean to check if the value has been set.
func (o *SmNasFromSmf) GetBackoffTimerOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BackoffTimer, true
}

// SetBackoffTimer sets field value
func (o *SmNasFromSmf) SetBackoffTimer(v int32) {
	o.BackoffTimer = v
}

// GetAppliedSmccType returns the AppliedSmccType field value
func (o *SmNasFromSmf) GetAppliedSmccType() AppliedSmccType {
	if o == nil {
		var ret AppliedSmccType
		return ret
	}

	return o.AppliedSmccType
}

// GetAppliedSmccTypeOk returns a tuple with the AppliedSmccType field value
// and a boolean to check if the value has been set.
func (o *SmNasFromSmf) GetAppliedSmccTypeOk() (*AppliedSmccType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppliedSmccType, true
}

// SetAppliedSmccType sets field value
func (o *SmNasFromSmf) SetAppliedSmccType(v AppliedSmccType) {
	o.AppliedSmccType = v
}

func (o SmNasFromSmf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["smNasType"] = o.SmNasType
	}
	if true {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	if true {
		toSerialize["backoffTimer"] = o.BackoffTimer
	}
	if true {
		toSerialize["appliedSmccType"] = o.AppliedSmccType
	}
	return json.Marshal(toSerialize)
}

type NullableSmNasFromSmf struct {
	value *SmNasFromSmf
	isSet bool
}

func (v NullableSmNasFromSmf) Get() *SmNasFromSmf {
	return v.value
}

func (v *NullableSmNasFromSmf) Set(val *SmNasFromSmf) {
	v.value = val
	v.isSet = true
}

func (v NullableSmNasFromSmf) IsSet() bool {
	return v.isSet
}

func (v *NullableSmNasFromSmf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmNasFromSmf(val *SmNasFromSmf) *NullableSmNasFromSmf {
	return &NullableSmNasFromSmf{value: val, isSet: true}
}

func (v NullableSmNasFromSmf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmNasFromSmf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


