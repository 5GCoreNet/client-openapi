/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
	"time"
)

// EASDeploymentChargingInformation struct for EASDeploymentChargingInformation
type EASDeploymentChargingInformation struct {
	// string with format 'date-time' as defined in OpenAPI.
	LCMStartTime *time.Time `json:"lCMStartTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	LCMEndTime *time.Time `json:"lCMEndTime,omitempty"`
}

// NewEASDeploymentChargingInformation instantiates a new EASDeploymentChargingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEASDeploymentChargingInformation() *EASDeploymentChargingInformation {
	this := EASDeploymentChargingInformation{}
	return &this
}

// NewEASDeploymentChargingInformationWithDefaults instantiates a new EASDeploymentChargingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEASDeploymentChargingInformationWithDefaults() *EASDeploymentChargingInformation {
	this := EASDeploymentChargingInformation{}
	return &this
}

// GetLCMStartTime returns the LCMStartTime field value if set, zero value otherwise.
func (o *EASDeploymentChargingInformation) GetLCMStartTime() time.Time {
	if o == nil || isNil(o.LCMStartTime) {
		var ret time.Time
		return ret
	}
	return *o.LCMStartTime
}

// GetLCMStartTimeOk returns a tuple with the LCMStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASDeploymentChargingInformation) GetLCMStartTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LCMStartTime) {
    return nil, false
	}
	return o.LCMStartTime, true
}

// HasLCMStartTime returns a boolean if a field has been set.
func (o *EASDeploymentChargingInformation) HasLCMStartTime() bool {
	if o != nil && !isNil(o.LCMStartTime) {
		return true
	}

	return false
}

// SetLCMStartTime gets a reference to the given time.Time and assigns it to the LCMStartTime field.
func (o *EASDeploymentChargingInformation) SetLCMStartTime(v time.Time) {
	o.LCMStartTime = &v
}

// GetLCMEndTime returns the LCMEndTime field value if set, zero value otherwise.
func (o *EASDeploymentChargingInformation) GetLCMEndTime() time.Time {
	if o == nil || isNil(o.LCMEndTime) {
		var ret time.Time
		return ret
	}
	return *o.LCMEndTime
}

// GetLCMEndTimeOk returns a tuple with the LCMEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASDeploymentChargingInformation) GetLCMEndTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LCMEndTime) {
    return nil, false
	}
	return o.LCMEndTime, true
}

// HasLCMEndTime returns a boolean if a field has been set.
func (o *EASDeploymentChargingInformation) HasLCMEndTime() bool {
	if o != nil && !isNil(o.LCMEndTime) {
		return true
	}

	return false
}

// SetLCMEndTime gets a reference to the given time.Time and assigns it to the LCMEndTime field.
func (o *EASDeploymentChargingInformation) SetLCMEndTime(v time.Time) {
	o.LCMEndTime = &v
}

func (o EASDeploymentChargingInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LCMStartTime) {
		toSerialize["lCMStartTime"] = o.LCMStartTime
	}
	if !isNil(o.LCMEndTime) {
		toSerialize["lCMEndTime"] = o.LCMEndTime
	}
	return json.Marshal(toSerialize)
}

type NullableEASDeploymentChargingInformation struct {
	value *EASDeploymentChargingInformation
	isSet bool
}

func (v NullableEASDeploymentChargingInformation) Get() *EASDeploymentChargingInformation {
	return v.value
}

func (v *NullableEASDeploymentChargingInformation) Set(val *EASDeploymentChargingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableEASDeploymentChargingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableEASDeploymentChargingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEASDeploymentChargingInformation(val *EASDeploymentChargingInformation) *NullableEASDeploymentChargingInformation {
	return &NullableEASDeploymentChargingInformation{value: val, isSet: true}
}

func (v NullableEASDeploymentChargingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEASDeploymentChargingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

