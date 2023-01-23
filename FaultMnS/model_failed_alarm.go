/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
)

// FailedAlarm struct for FailedAlarm
type FailedAlarm struct {
	AlarmId string `json:"alarmId"`
	FailureReason string `json:"failureReason"`
}

// NewFailedAlarm instantiates a new FailedAlarm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailedAlarm(alarmId string, failureReason string) *FailedAlarm {
	this := FailedAlarm{}
	this.AlarmId = alarmId
	this.FailureReason = failureReason
	return &this
}

// NewFailedAlarmWithDefaults instantiates a new FailedAlarm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailedAlarmWithDefaults() *FailedAlarm {
	this := FailedAlarm{}
	return &this
}

// GetAlarmId returns the AlarmId field value
func (o *FailedAlarm) GetAlarmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlarmId
}

// GetAlarmIdOk returns a tuple with the AlarmId field value
// and a boolean to check if the value has been set.
func (o *FailedAlarm) GetAlarmIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AlarmId, true
}

// SetAlarmId sets field value
func (o *FailedAlarm) SetAlarmId(v string) {
	o.AlarmId = v
}

// GetFailureReason returns the FailureReason field value
func (o *FailedAlarm) GetFailureReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value
// and a boolean to check if the value has been set.
func (o *FailedAlarm) GetFailureReasonOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FailureReason, true
}

// SetFailureReason sets field value
func (o *FailedAlarm) SetFailureReason(v string) {
	o.FailureReason = v
}

func (o FailedAlarm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alarmId"] = o.AlarmId
	}
	if true {
		toSerialize["failureReason"] = o.FailureReason
	}
	return json.Marshal(toSerialize)
}

type NullableFailedAlarm struct {
	value *FailedAlarm
	isSet bool
}

func (v NullableFailedAlarm) Get() *FailedAlarm {
	return v.value
}

func (v *NullableFailedAlarm) Set(val *FailedAlarm) {
	v.value = val
	v.isSet = true
}

func (v NullableFailedAlarm) IsSet() bool {
	return v.isSet
}

func (v *NullableFailedAlarm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailedAlarm(val *FailedAlarm) *NullableFailedAlarm {
	return &NullableFailedAlarm{value: val, isSet: true}
}

func (v NullableFailedAlarm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailedAlarm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


