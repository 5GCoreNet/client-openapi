/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_DataManagement

import (
	"encoding/json"
	"time"
)

// ReportingOptions1 struct for ReportingOptions1
type ReportingOptions1 struct {
	ReportMode *EventReportMode `json:"reportMode,omitempty"`
	MaxNumOfReports *int32 `json:"maxNumOfReports,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	SamplingRatio *int32 `json:"samplingRatio,omitempty"`
	// indicating a time in seconds.
	GuardTime *int32 `json:"guardTime,omitempty"`
	// indicating a time in seconds.
	ReportPeriod *int32 `json:"reportPeriod,omitempty"`
	NotifFlag *NotificationFlag `json:"notifFlag,omitempty"`
}

// NewReportingOptions1 instantiates a new ReportingOptions1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingOptions1() *ReportingOptions1 {
	this := ReportingOptions1{}
	return &this
}

// NewReportingOptions1WithDefaults instantiates a new ReportingOptions1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingOptions1WithDefaults() *ReportingOptions1 {
	this := ReportingOptions1{}
	return &this
}

// GetReportMode returns the ReportMode field value if set, zero value otherwise.
func (o *ReportingOptions1) GetReportMode() EventReportMode {
	if o == nil || isNil(o.ReportMode) {
		var ret EventReportMode
		return ret
	}
	return *o.ReportMode
}

// GetReportModeOk returns a tuple with the ReportMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions1) GetReportModeOk() (*EventReportMode, bool) {
	if o == nil || isNil(o.ReportMode) {
    return nil, false
	}
	return o.ReportMode, true
}

// HasReportMode returns a boolean if a field has been set.
func (o *ReportingOptions1) HasReportMode() bool {
	if o != nil && !isNil(o.ReportMode) {
		return true
	}

	return false
}

// SetReportMode gets a reference to the given EventReportMode and assigns it to the ReportMode field.
func (o *ReportingOptions1) SetReportMode(v EventReportMode) {
	o.ReportMode = &v
}

// GetMaxNumOfReports returns the MaxNumOfReports field value if set, zero value otherwise.
func (o *ReportingOptions1) GetMaxNumOfReports() int32 {
	if o == nil || isNil(o.MaxNumOfReports) {
		var ret int32
		return ret
	}
	return *o.MaxNumOfReports
}

// GetMaxNumOfReportsOk returns a tuple with the MaxNumOfReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions1) GetMaxNumOfReportsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxNumOfReports) {
    return nil, false
	}
	return o.MaxNumOfReports, true
}

// HasMaxNumOfReports returns a boolean if a field has been set.
func (o *ReportingOptions1) HasMaxNumOfReports() bool {
	if o != nil && !isNil(o.MaxNumOfReports) {
		return true
	}

	return false
}

// SetMaxNumOfReports gets a reference to the given int32 and assigns it to the MaxNumOfReports field.
func (o *ReportingOptions1) SetMaxNumOfReports(v int32) {
	o.MaxNumOfReports = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *ReportingOptions1) GetExpiry() time.Time {
	if o == nil || isNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions1) GetExpiryOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expiry) {
    return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *ReportingOptions1) HasExpiry() bool {
	if o != nil && !isNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *ReportingOptions1) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetSamplingRatio returns the SamplingRatio field value if set, zero value otherwise.
func (o *ReportingOptions1) GetSamplingRatio() int32 {
	if o == nil || isNil(o.SamplingRatio) {
		var ret int32
		return ret
	}
	return *o.SamplingRatio
}

// GetSamplingRatioOk returns a tuple with the SamplingRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions1) GetSamplingRatioOk() (*int32, bool) {
	if o == nil || isNil(o.SamplingRatio) {
    return nil, false
	}
	return o.SamplingRatio, true
}

// HasSamplingRatio returns a boolean if a field has been set.
func (o *ReportingOptions1) HasSamplingRatio() bool {
	if o != nil && !isNil(o.SamplingRatio) {
		return true
	}

	return false
}

// SetSamplingRatio gets a reference to the given int32 and assigns it to the SamplingRatio field.
func (o *ReportingOptions1) SetSamplingRatio(v int32) {
	o.SamplingRatio = &v
}

// GetGuardTime returns the GuardTime field value if set, zero value otherwise.
func (o *ReportingOptions1) GetGuardTime() int32 {
	if o == nil || isNil(o.GuardTime) {
		var ret int32
		return ret
	}
	return *o.GuardTime
}

// GetGuardTimeOk returns a tuple with the GuardTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions1) GetGuardTimeOk() (*int32, bool) {
	if o == nil || isNil(o.GuardTime) {
    return nil, false
	}
	return o.GuardTime, true
}

// HasGuardTime returns a boolean if a field has been set.
func (o *ReportingOptions1) HasGuardTime() bool {
	if o != nil && !isNil(o.GuardTime) {
		return true
	}

	return false
}

// SetGuardTime gets a reference to the given int32 and assigns it to the GuardTime field.
func (o *ReportingOptions1) SetGuardTime(v int32) {
	o.GuardTime = &v
}

// GetReportPeriod returns the ReportPeriod field value if set, zero value otherwise.
func (o *ReportingOptions1) GetReportPeriod() int32 {
	if o == nil || isNil(o.ReportPeriod) {
		var ret int32
		return ret
	}
	return *o.ReportPeriod
}

// GetReportPeriodOk returns a tuple with the ReportPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions1) GetReportPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.ReportPeriod) {
    return nil, false
	}
	return o.ReportPeriod, true
}

// HasReportPeriod returns a boolean if a field has been set.
func (o *ReportingOptions1) HasReportPeriod() bool {
	if o != nil && !isNil(o.ReportPeriod) {
		return true
	}

	return false
}

// SetReportPeriod gets a reference to the given int32 and assigns it to the ReportPeriod field.
func (o *ReportingOptions1) SetReportPeriod(v int32) {
	o.ReportPeriod = &v
}

// GetNotifFlag returns the NotifFlag field value if set, zero value otherwise.
func (o *ReportingOptions1) GetNotifFlag() NotificationFlag {
	if o == nil || isNil(o.NotifFlag) {
		var ret NotificationFlag
		return ret
	}
	return *o.NotifFlag
}

// GetNotifFlagOk returns a tuple with the NotifFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions1) GetNotifFlagOk() (*NotificationFlag, bool) {
	if o == nil || isNil(o.NotifFlag) {
    return nil, false
	}
	return o.NotifFlag, true
}

// HasNotifFlag returns a boolean if a field has been set.
func (o *ReportingOptions1) HasNotifFlag() bool {
	if o != nil && !isNil(o.NotifFlag) {
		return true
	}

	return false
}

// SetNotifFlag gets a reference to the given NotificationFlag and assigns it to the NotifFlag field.
func (o *ReportingOptions1) SetNotifFlag(v NotificationFlag) {
	o.NotifFlag = &v
}

func (o ReportingOptions1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReportMode) {
		toSerialize["reportMode"] = o.ReportMode
	}
	if !isNil(o.MaxNumOfReports) {
		toSerialize["maxNumOfReports"] = o.MaxNumOfReports
	}
	if !isNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !isNil(o.SamplingRatio) {
		toSerialize["samplingRatio"] = o.SamplingRatio
	}
	if !isNil(o.GuardTime) {
		toSerialize["guardTime"] = o.GuardTime
	}
	if !isNil(o.ReportPeriod) {
		toSerialize["reportPeriod"] = o.ReportPeriod
	}
	if !isNil(o.NotifFlag) {
		toSerialize["notifFlag"] = o.NotifFlag
	}
	return json.Marshal(toSerialize)
}

type NullableReportingOptions1 struct {
	value *ReportingOptions1
	isSet bool
}

func (v NullableReportingOptions1) Get() *ReportingOptions1 {
	return v.value
}

func (v *NullableReportingOptions1) Set(val *ReportingOptions1) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingOptions1) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingOptions1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingOptions1(val *ReportingOptions1) *NullableReportingOptions1 {
	return &NullableReportingOptions1{value: val, isSet: true}
}

func (v NullableReportingOptions1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingOptions1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


