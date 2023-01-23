/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_EE

import (
	"encoding/json"
	"time"
)

// IdleStatusIndication Represents the idle status indication.
type IdleStatusIndication struct {
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp *time.Time `json:"timeStamp,omitempty"`
	// indicating a time in seconds.
	ActiveTime *int32 `json:"activeTime,omitempty"`
	// indicating a time in seconds.
	SubsRegTimer *int32 `json:"subsRegTimer,omitempty"`
	EdrxCycleLength *int32 `json:"edrxCycleLength,omitempty"`
	SuggestedNumOfDlPackets *int32 `json:"suggestedNumOfDlPackets,omitempty"`
}

// NewIdleStatusIndication instantiates a new IdleStatusIndication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdleStatusIndication() *IdleStatusIndication {
	this := IdleStatusIndication{}
	return &this
}

// NewIdleStatusIndicationWithDefaults instantiates a new IdleStatusIndication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdleStatusIndicationWithDefaults() *IdleStatusIndication {
	this := IdleStatusIndication{}
	return &this
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *IdleStatusIndication) GetTimeStamp() time.Time {
	if o == nil || isNil(o.TimeStamp) {
		var ret time.Time
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdleStatusIndication) GetTimeStampOk() (*time.Time, bool) {
	if o == nil || isNil(o.TimeStamp) {
    return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *IdleStatusIndication) HasTimeStamp() bool {
	if o != nil && !isNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given time.Time and assigns it to the TimeStamp field.
func (o *IdleStatusIndication) SetTimeStamp(v time.Time) {
	o.TimeStamp = &v
}

// GetActiveTime returns the ActiveTime field value if set, zero value otherwise.
func (o *IdleStatusIndication) GetActiveTime() int32 {
	if o == nil || isNil(o.ActiveTime) {
		var ret int32
		return ret
	}
	return *o.ActiveTime
}

// GetActiveTimeOk returns a tuple with the ActiveTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdleStatusIndication) GetActiveTimeOk() (*int32, bool) {
	if o == nil || isNil(o.ActiveTime) {
    return nil, false
	}
	return o.ActiveTime, true
}

// HasActiveTime returns a boolean if a field has been set.
func (o *IdleStatusIndication) HasActiveTime() bool {
	if o != nil && !isNil(o.ActiveTime) {
		return true
	}

	return false
}

// SetActiveTime gets a reference to the given int32 and assigns it to the ActiveTime field.
func (o *IdleStatusIndication) SetActiveTime(v int32) {
	o.ActiveTime = &v
}

// GetSubsRegTimer returns the SubsRegTimer field value if set, zero value otherwise.
func (o *IdleStatusIndication) GetSubsRegTimer() int32 {
	if o == nil || isNil(o.SubsRegTimer) {
		var ret int32
		return ret
	}
	return *o.SubsRegTimer
}

// GetSubsRegTimerOk returns a tuple with the SubsRegTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdleStatusIndication) GetSubsRegTimerOk() (*int32, bool) {
	if o == nil || isNil(o.SubsRegTimer) {
    return nil, false
	}
	return o.SubsRegTimer, true
}

// HasSubsRegTimer returns a boolean if a field has been set.
func (o *IdleStatusIndication) HasSubsRegTimer() bool {
	if o != nil && !isNil(o.SubsRegTimer) {
		return true
	}

	return false
}

// SetSubsRegTimer gets a reference to the given int32 and assigns it to the SubsRegTimer field.
func (o *IdleStatusIndication) SetSubsRegTimer(v int32) {
	o.SubsRegTimer = &v
}

// GetEdrxCycleLength returns the EdrxCycleLength field value if set, zero value otherwise.
func (o *IdleStatusIndication) GetEdrxCycleLength() int32 {
	if o == nil || isNil(o.EdrxCycleLength) {
		var ret int32
		return ret
	}
	return *o.EdrxCycleLength
}

// GetEdrxCycleLengthOk returns a tuple with the EdrxCycleLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdleStatusIndication) GetEdrxCycleLengthOk() (*int32, bool) {
	if o == nil || isNil(o.EdrxCycleLength) {
    return nil, false
	}
	return o.EdrxCycleLength, true
}

// HasEdrxCycleLength returns a boolean if a field has been set.
func (o *IdleStatusIndication) HasEdrxCycleLength() bool {
	if o != nil && !isNil(o.EdrxCycleLength) {
		return true
	}

	return false
}

// SetEdrxCycleLength gets a reference to the given int32 and assigns it to the EdrxCycleLength field.
func (o *IdleStatusIndication) SetEdrxCycleLength(v int32) {
	o.EdrxCycleLength = &v
}

// GetSuggestedNumOfDlPackets returns the SuggestedNumOfDlPackets field value if set, zero value otherwise.
func (o *IdleStatusIndication) GetSuggestedNumOfDlPackets() int32 {
	if o == nil || isNil(o.SuggestedNumOfDlPackets) {
		var ret int32
		return ret
	}
	return *o.SuggestedNumOfDlPackets
}

// GetSuggestedNumOfDlPacketsOk returns a tuple with the SuggestedNumOfDlPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdleStatusIndication) GetSuggestedNumOfDlPacketsOk() (*int32, bool) {
	if o == nil || isNil(o.SuggestedNumOfDlPackets) {
    return nil, false
	}
	return o.SuggestedNumOfDlPackets, true
}

// HasSuggestedNumOfDlPackets returns a boolean if a field has been set.
func (o *IdleStatusIndication) HasSuggestedNumOfDlPackets() bool {
	if o != nil && !isNil(o.SuggestedNumOfDlPackets) {
		return true
	}

	return false
}

// SetSuggestedNumOfDlPackets gets a reference to the given int32 and assigns it to the SuggestedNumOfDlPackets field.
func (o *IdleStatusIndication) SetSuggestedNumOfDlPackets(v int32) {
	o.SuggestedNumOfDlPackets = &v
}

func (o IdleStatusIndication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TimeStamp) {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	if !isNil(o.ActiveTime) {
		toSerialize["activeTime"] = o.ActiveTime
	}
	if !isNil(o.SubsRegTimer) {
		toSerialize["subsRegTimer"] = o.SubsRegTimer
	}
	if !isNil(o.EdrxCycleLength) {
		toSerialize["edrxCycleLength"] = o.EdrxCycleLength
	}
	if !isNil(o.SuggestedNumOfDlPackets) {
		toSerialize["suggestedNumOfDlPackets"] = o.SuggestedNumOfDlPackets
	}
	return json.Marshal(toSerialize)
}

type NullableIdleStatusIndication struct {
	value *IdleStatusIndication
	isSet bool
}

func (v NullableIdleStatusIndication) Get() *IdleStatusIndication {
	return v.value
}

func (v *NullableIdleStatusIndication) Set(val *IdleStatusIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableIdleStatusIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableIdleStatusIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdleStatusIndication(val *IdleStatusIndication) *NullableIdleStatusIndication {
	return &NullableIdleStatusIndication{value: val, isSet: true}
}

func (v NullableIdleStatusIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdleStatusIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


