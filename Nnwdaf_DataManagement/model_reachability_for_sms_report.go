/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"time"
)

// ReachabilityForSmsReport struct for ReachabilityForSmsReport
type ReachabilityForSmsReport struct {
	SmsfAccessType AccessType `json:"smsfAccessType"`
	// string with format 'date-time' as defined in OpenAPI.
	MaxAvailabilityTime *time.Time `json:"maxAvailabilityTime,omitempty"`
}

// NewReachabilityForSmsReport instantiates a new ReachabilityForSmsReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReachabilityForSmsReport(smsfAccessType AccessType) *ReachabilityForSmsReport {
	this := ReachabilityForSmsReport{}
	this.SmsfAccessType = smsfAccessType
	return &this
}

// NewReachabilityForSmsReportWithDefaults instantiates a new ReachabilityForSmsReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReachabilityForSmsReportWithDefaults() *ReachabilityForSmsReport {
	this := ReachabilityForSmsReport{}
	return &this
}

// GetSmsfAccessType returns the SmsfAccessType field value
func (o *ReachabilityForSmsReport) GetSmsfAccessType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.SmsfAccessType
}

// GetSmsfAccessTypeOk returns a tuple with the SmsfAccessType field value
// and a boolean to check if the value has been set.
func (o *ReachabilityForSmsReport) GetSmsfAccessTypeOk() (*AccessType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SmsfAccessType, true
}

// SetSmsfAccessType sets field value
func (o *ReachabilityForSmsReport) SetSmsfAccessType(v AccessType) {
	o.SmsfAccessType = v
}

// GetMaxAvailabilityTime returns the MaxAvailabilityTime field value if set, zero value otherwise.
func (o *ReachabilityForSmsReport) GetMaxAvailabilityTime() time.Time {
	if o == nil || isNil(o.MaxAvailabilityTime) {
		var ret time.Time
		return ret
	}
	return *o.MaxAvailabilityTime
}

// GetMaxAvailabilityTimeOk returns a tuple with the MaxAvailabilityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReachabilityForSmsReport) GetMaxAvailabilityTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.MaxAvailabilityTime) {
    return nil, false
	}
	return o.MaxAvailabilityTime, true
}

// HasMaxAvailabilityTime returns a boolean if a field has been set.
func (o *ReachabilityForSmsReport) HasMaxAvailabilityTime() bool {
	if o != nil && !isNil(o.MaxAvailabilityTime) {
		return true
	}

	return false
}

// SetMaxAvailabilityTime gets a reference to the given time.Time and assigns it to the MaxAvailabilityTime field.
func (o *ReachabilityForSmsReport) SetMaxAvailabilityTime(v time.Time) {
	o.MaxAvailabilityTime = &v
}

func (o ReachabilityForSmsReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["smsfAccessType"] = o.SmsfAccessType
	}
	if !isNil(o.MaxAvailabilityTime) {
		toSerialize["maxAvailabilityTime"] = o.MaxAvailabilityTime
	}
	return json.Marshal(toSerialize)
}

type NullableReachabilityForSmsReport struct {
	value *ReachabilityForSmsReport
	isSet bool
}

func (v NullableReachabilityForSmsReport) Get() *ReachabilityForSmsReport {
	return v.value
}

func (v *NullableReachabilityForSmsReport) Set(val *ReachabilityForSmsReport) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityForSmsReport) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityForSmsReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityForSmsReport(val *ReachabilityForSmsReport) *NullableReachabilityForSmsReport {
	return &NullableReachabilityForSmsReport{value: val, isSet: true}
}

func (v NullableReachabilityForSmsReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityForSmsReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


