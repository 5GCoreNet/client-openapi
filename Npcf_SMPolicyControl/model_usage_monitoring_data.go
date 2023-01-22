/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_SMPolicyControl

import (
	"encoding/json"
	"time"
)

// UsageMonitoringData Contains usage monitoring related control information.
type UsageMonitoringData struct {
	// Univocally identifies the usage monitoring policy data within a PDU session.
	UmId string `json:"umId"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	VolumeThreshold NullableInt64 `json:"volumeThreshold,omitempty"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	VolumeThresholdUplink NullableInt64 `json:"volumeThresholdUplink,omitempty"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	VolumeThresholdDownlink NullableInt64 `json:"volumeThresholdDownlink,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	TimeThreshold NullableInt32 `json:"timeThreshold,omitempty"`
	// string with format 'date-time' as defined in OpenAPI with 'nullable:true' property.  
	MonitoringTime NullableTime `json:"monitoringTime,omitempty"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	NextVolThreshold NullableInt64 `json:"nextVolThreshold,omitempty"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	NextVolThresholdUplink NullableInt64 `json:"nextVolThresholdUplink,omitempty"`
	// Unsigned integer identifying a volume in units of bytes with \"nullable=true\" property.
	NextVolThresholdDownlink NullableInt64 `json:"nextVolThresholdDownlink,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	NextTimeThreshold NullableInt32 `json:"nextTimeThreshold,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	InactivityTime NullableInt32 `json:"inactivityTime,omitempty"`
	// Contains the PCC rule identifier(s) which corresponding service data flow(s) shall be excluded from PDU Session usage monitoring. It is only included in the UsageMonitoringData instance for session level usage monitoring. 
	ExUsagePccRuleIds []string `json:"exUsagePccRuleIds,omitempty"`
}

// NewUsageMonitoringData instantiates a new UsageMonitoringData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageMonitoringData(umId string) *UsageMonitoringData {
	this := UsageMonitoringData{}
	this.UmId = umId
	return &this
}

// NewUsageMonitoringDataWithDefaults instantiates a new UsageMonitoringData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageMonitoringDataWithDefaults() *UsageMonitoringData {
	this := UsageMonitoringData{}
	return &this
}

// GetUmId returns the UmId field value
func (o *UsageMonitoringData) GetUmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UmId
}

// GetUmIdOk returns a tuple with the UmId field value
// and a boolean to check if the value has been set.
func (o *UsageMonitoringData) GetUmIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UmId, true
}

// SetUmId sets field value
func (o *UsageMonitoringData) SetUmId(v string) {
	o.UmId = v
}

// GetVolumeThreshold returns the VolumeThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageMonitoringData) GetVolumeThreshold() int64 {
	if o == nil || isNil(o.VolumeThreshold.Get()) {
		var ret int64
		return ret
	}
	return *o.VolumeThreshold.Get()
}

// GetVolumeThresholdOk returns a tuple with the VolumeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageMonitoringData) GetVolumeThresholdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return o.VolumeThreshold.Get(), o.VolumeThreshold.IsSet()
}

// HasVolumeThreshold returns a boolean if a field has been set.
func (o *UsageMonitoringData) HasVolumeThreshold() bool {
	if o != nil && o.VolumeThreshold.IsSet() {
		return true
	}

	return false
}

// SetVolumeThreshold gets a reference to the given NullableInt64 and assigns it to the VolumeThreshold field.
func (o *UsageMonitoringData) SetVolumeThreshold(v int64) {
	o.VolumeThreshold.Set(&v)
}
// SetVolumeThresholdNil sets the value for VolumeThreshold to be an explicit nil
func (o *UsageMonitoringData) SetVolumeThresholdNil() {
	o.VolumeThreshold.Set(nil)
}

// UnsetVolumeThreshold ensures that no value is present for VolumeThreshold, not even an explicit nil
func (o *UsageMonitoringData) UnsetVolumeThreshold() {
	o.VolumeThreshold.Unset()
}

// GetVolumeThresholdUplink returns the VolumeThresholdUplink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageMonitoringData) GetVolumeThresholdUplink() int64 {
	if o == nil || isNil(o.VolumeThresholdUplink.Get()) {
		var ret int64
		return ret
	}
	return *o.VolumeThresholdUplink.Get()
}

// GetVolumeThresholdUplinkOk returns a tuple with the VolumeThresholdUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageMonitoringData) GetVolumeThresholdUplinkOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return o.VolumeThresholdUplink.Get(), o.VolumeThresholdUplink.IsSet()
}

// HasVolumeThresholdUplink returns a boolean if a field has been set.
func (o *UsageMonitoringData) HasVolumeThresholdUplink() bool {
	if o != nil && o.VolumeThresholdUplink.IsSet() {
		return true
	}

	return false
}

// SetVolumeThresholdUplink gets a reference to the given NullableInt64 and assigns it to the VolumeThresholdUplink field.
func (o *UsageMonitoringData) SetVolumeThresholdUplink(v int64) {
	o.VolumeThresholdUplink.Set(&v)
}
// SetVolumeThresholdUplinkNil sets the value for VolumeThresholdUplink to be an explicit nil
func (o *UsageMonitoringData) SetVolumeThresholdUplinkNil() {
	o.VolumeThresholdUplink.Set(nil)
}

// UnsetVolumeThresholdUplink ensures that no value is present for VolumeThresholdUplink, not even an explicit nil
func (o *UsageMonitoringData) UnsetVolumeThresholdUplink() {
	o.VolumeThresholdUplink.Unset()
}

// GetVolumeThresholdDownlink returns the VolumeThresholdDownlink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageMonitoringData) GetVolumeThresholdDownlink() int64 {
	if o == nil || isNil(o.VolumeThresholdDownlink.Get()) {
		var ret int64
		return ret
	}
	return *o.VolumeThresholdDownlink.Get()
}

// GetVolumeThresholdDownlinkOk returns a tuple with the VolumeThresholdDownlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageMonitoringData) GetVolumeThresholdDownlinkOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return o.VolumeThresholdDownlink.Get(), o.VolumeThresholdDownlink.IsSet()
}

// HasVolumeThresholdDownlink returns a boolean if a field has been set.
func (o *UsageMonitoringData) HasVolumeThresholdDownlink() bool {
	if o != nil && o.VolumeThresholdDownlink.IsSet() {
		return true
	}

	return false
}

// SetVolumeThresholdDownlink gets a reference to the given NullableInt64 and assigns it to the VolumeThresholdDownlink field.
func (o *UsageMonitoringData) SetVolumeThresholdDownlink(v int64) {
	o.VolumeThresholdDownlink.Set(&v)
}
// SetVolumeThresholdDownlinkNil sets the value for VolumeThresholdDownlink to be an explicit nil
func (o *UsageMonitoringData) SetVolumeThresholdDownlinkNil() {
	o.VolumeThresholdDownlink.Set(nil)
}

// UnsetVolumeThresholdDownlink ensures that no value is present for VolumeThresholdDownlink, not even an explicit nil
func (o *UsageMonitoringData) UnsetVolumeThresholdDownlink() {
	o.VolumeThresholdDownlink.Unset()
}

// GetTimeThreshold returns the TimeThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageMonitoringData) GetTimeThreshold() int32 {
	if o == nil || isNil(o.TimeThreshold.Get()) {
		var ret int32
		return ret
	}
	return *o.TimeThreshold.Get()
}

// GetTimeThresholdOk returns a tuple with the TimeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageMonitoringData) GetTimeThresholdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.TimeThreshold.Get(), o.TimeThreshold.IsSet()
}

// HasTimeThreshold returns a boolean if a field has been set.
func (o *UsageMonitoringData) HasTimeThreshold() bool {
	if o != nil && o.TimeThreshold.IsSet() {
		return true
	}

	return false
}

// SetTimeThreshold gets a reference to the given NullableInt32 and assigns it to the TimeThreshold field.
func (o *UsageMonitoringData) SetTimeThreshold(v int32) {
	o.TimeThreshold.Set(&v)
}
// SetTimeThresholdNil sets the value for TimeThreshold to be an explicit nil
func (o *UsageMonitoringData) SetTimeThresholdNil() {
	o.TimeThreshold.Set(nil)
}

// UnsetTimeThreshold ensures that no value is present for TimeThreshold, not even an explicit nil
func (o *UsageMonitoringData) UnsetTimeThreshold() {
	o.TimeThreshold.Unset()
}

// GetMonitoringTime returns the MonitoringTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageMonitoringData) GetMonitoringTime() time.Time {
	if o == nil || isNil(o.MonitoringTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.MonitoringTime.Get()
}

// GetMonitoringTimeOk returns a tuple with the MonitoringTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageMonitoringData) GetMonitoringTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.MonitoringTime.Get(), o.MonitoringTime.IsSet()
}

// HasMonitoringTime returns a boolean if a field has been set.
func (o *UsageMonitoringData) HasMonitoringTime() bool {
	if o != nil && o.MonitoringTime.IsSet() {
		return true
	}

	return false
}

// SetMonitoringTime gets a reference to the given NullableTime and assigns it to the MonitoringTime field.
func (o *UsageMonitoringData) SetMonitoringTime(v time.Time) {
	o.MonitoringTime.Set(&v)
}
// SetMonitoringTimeNil sets the value for MonitoringTime to be an explicit nil
func (o *UsageMonitoringData) SetMonitoringTimeNil() {
	o.MonitoringTime.Set(nil)
}

// UnsetMonitoringTime ensures that no value is present for MonitoringTime, not even an explicit nil
func (o *UsageMonitoringData) UnsetMonitoringTime() {
	o.MonitoringTime.Unset()
}

// GetNextVolThreshold returns the NextVolThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageMonitoringData) GetNextVolThreshold() int64 {
	if o == nil || isNil(o.NextVolThreshold.Get()) {
		var ret int64
		return ret
	}
	return *o.NextVolThreshold.Get()
}

// GetNextVolThresholdOk returns a tuple with the NextVolThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageMonitoringData) GetNextVolThresholdOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return o.NextVolThreshold.Get(), o.NextVolThreshold.IsSet()
}

// HasNextVolThreshold returns a boolean if a field has been set.
func (o *UsageMonitoringData) HasNextVolThreshold() bool {
	if o != nil && o.NextVolThreshold.IsSet() {
		return true
	}

	return false
}

// SetNextVolThreshold gets a reference to the given NullableInt64 and assigns it to the NextVolThreshold field.
func (o *UsageMonitoringData) SetNextVolThreshold(v int64) {
	o.NextVolThreshold.Set(&v)
}
// SetNextVolThresholdNil sets the value for NextVolThreshold to be an explicit nil
func (o *UsageMonitoringData) SetNextVolThresholdNil() {
	o.NextVolThreshold.Set(nil)
}

// UnsetNextVolThreshold ensures that no value is present for NextVolThreshold, not even an explicit nil
func (o *UsageMonitoringData) UnsetNextVolThreshold() {
	o.NextVolThreshold.Unset()
}

// GetNextVolThresholdUplink returns the NextVolThresholdUplink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageMonitoringData) GetNextVolThresholdUplink() int64 {
	if o == nil || isNil(o.NextVolThresholdUplink.Get()) {
		var ret int64
		return ret
	}
	return *o.NextVolThresholdUplink.Get()
}

// GetNextVolThresholdUplinkOk returns a tuple with the NextVolThresholdUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageMonitoringData) GetNextVolThresholdUplinkOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return o.NextVolThresholdUplink.Get(), o.NextVolThresholdUplink.IsSet()
}

// HasNextVolThresholdUplink returns a boolean if a field has been set.
func (o *UsageMonitoringData) HasNextVolThresholdUplink() bool {
	if o != nil && o.NextVolThresholdUplink.IsSet() {
		return true
	}

	return false
}

// SetNextVolThresholdUplink gets a reference to the given NullableInt64 and assigns it to the NextVolThresholdUplink field.
func (o *UsageMonitoringData) SetNextVolThresholdUplink(v int64) {
	o.NextVolThresholdUplink.Set(&v)
}
// SetNextVolThresholdUplinkNil sets the value for NextVolThresholdUplink to be an explicit nil
func (o *UsageMonitoringData) SetNextVolThresholdUplinkNil() {
	o.NextVolThresholdUplink.Set(nil)
}

// UnsetNextVolThresholdUplink ensures that no value is present for NextVolThresholdUplink, not even an explicit nil
func (o *UsageMonitoringData) UnsetNextVolThresholdUplink() {
	o.NextVolThresholdUplink.Unset()
}

// GetNextVolThresholdDownlink returns the NextVolThresholdDownlink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageMonitoringData) GetNextVolThresholdDownlink() int64 {
	if o == nil || isNil(o.NextVolThresholdDownlink.Get()) {
		var ret int64
		return ret
	}
	return *o.NextVolThresholdDownlink.Get()
}

// GetNextVolThresholdDownlinkOk returns a tuple with the NextVolThresholdDownlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageMonitoringData) GetNextVolThresholdDownlinkOk() (*int64, bool) {
	if o == nil {
    return nil, false
	}
	return o.NextVolThresholdDownlink.Get(), o.NextVolThresholdDownlink.IsSet()
}

// HasNextVolThresholdDownlink returns a boolean if a field has been set.
func (o *UsageMonitoringData) HasNextVolThresholdDownlink() bool {
	if o != nil && o.NextVolThresholdDownlink.IsSet() {
		return true
	}

	return false
}

// SetNextVolThresholdDownlink gets a reference to the given NullableInt64 and assigns it to the NextVolThresholdDownlink field.
func (o *UsageMonitoringData) SetNextVolThresholdDownlink(v int64) {
	o.NextVolThresholdDownlink.Set(&v)
}
// SetNextVolThresholdDownlinkNil sets the value for NextVolThresholdDownlink to be an explicit nil
func (o *UsageMonitoringData) SetNextVolThresholdDownlinkNil() {
	o.NextVolThresholdDownlink.Set(nil)
}

// UnsetNextVolThresholdDownlink ensures that no value is present for NextVolThresholdDownlink, not even an explicit nil
func (o *UsageMonitoringData) UnsetNextVolThresholdDownlink() {
	o.NextVolThresholdDownlink.Unset()
}

// GetNextTimeThreshold returns the NextTimeThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageMonitoringData) GetNextTimeThreshold() int32 {
	if o == nil || isNil(o.NextTimeThreshold.Get()) {
		var ret int32
		return ret
	}
	return *o.NextTimeThreshold.Get()
}

// GetNextTimeThresholdOk returns a tuple with the NextTimeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageMonitoringData) GetNextTimeThresholdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.NextTimeThreshold.Get(), o.NextTimeThreshold.IsSet()
}

// HasNextTimeThreshold returns a boolean if a field has been set.
func (o *UsageMonitoringData) HasNextTimeThreshold() bool {
	if o != nil && o.NextTimeThreshold.IsSet() {
		return true
	}

	return false
}

// SetNextTimeThreshold gets a reference to the given NullableInt32 and assigns it to the NextTimeThreshold field.
func (o *UsageMonitoringData) SetNextTimeThreshold(v int32) {
	o.NextTimeThreshold.Set(&v)
}
// SetNextTimeThresholdNil sets the value for NextTimeThreshold to be an explicit nil
func (o *UsageMonitoringData) SetNextTimeThresholdNil() {
	o.NextTimeThreshold.Set(nil)
}

// UnsetNextTimeThreshold ensures that no value is present for NextTimeThreshold, not even an explicit nil
func (o *UsageMonitoringData) UnsetNextTimeThreshold() {
	o.NextTimeThreshold.Unset()
}

// GetInactivityTime returns the InactivityTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageMonitoringData) GetInactivityTime() int32 {
	if o == nil || isNil(o.InactivityTime.Get()) {
		var ret int32
		return ret
	}
	return *o.InactivityTime.Get()
}

// GetInactivityTimeOk returns a tuple with the InactivityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageMonitoringData) GetInactivityTimeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.InactivityTime.Get(), o.InactivityTime.IsSet()
}

// HasInactivityTime returns a boolean if a field has been set.
func (o *UsageMonitoringData) HasInactivityTime() bool {
	if o != nil && o.InactivityTime.IsSet() {
		return true
	}

	return false
}

// SetInactivityTime gets a reference to the given NullableInt32 and assigns it to the InactivityTime field.
func (o *UsageMonitoringData) SetInactivityTime(v int32) {
	o.InactivityTime.Set(&v)
}
// SetInactivityTimeNil sets the value for InactivityTime to be an explicit nil
func (o *UsageMonitoringData) SetInactivityTimeNil() {
	o.InactivityTime.Set(nil)
}

// UnsetInactivityTime ensures that no value is present for InactivityTime, not even an explicit nil
func (o *UsageMonitoringData) UnsetInactivityTime() {
	o.InactivityTime.Unset()
}

// GetExUsagePccRuleIds returns the ExUsagePccRuleIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageMonitoringData) GetExUsagePccRuleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExUsagePccRuleIds
}

// GetExUsagePccRuleIdsOk returns a tuple with the ExUsagePccRuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageMonitoringData) GetExUsagePccRuleIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ExUsagePccRuleIds) {
    return nil, false
	}
	return o.ExUsagePccRuleIds, true
}

// HasExUsagePccRuleIds returns a boolean if a field has been set.
func (o *UsageMonitoringData) HasExUsagePccRuleIds() bool {
	if o != nil && isNil(o.ExUsagePccRuleIds) {
		return true
	}

	return false
}

// SetExUsagePccRuleIds gets a reference to the given []string and assigns it to the ExUsagePccRuleIds field.
func (o *UsageMonitoringData) SetExUsagePccRuleIds(v []string) {
	o.ExUsagePccRuleIds = v
}

func (o UsageMonitoringData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["umId"] = o.UmId
	}
	if o.VolumeThreshold.IsSet() {
		toSerialize["volumeThreshold"] = o.VolumeThreshold.Get()
	}
	if o.VolumeThresholdUplink.IsSet() {
		toSerialize["volumeThresholdUplink"] = o.VolumeThresholdUplink.Get()
	}
	if o.VolumeThresholdDownlink.IsSet() {
		toSerialize["volumeThresholdDownlink"] = o.VolumeThresholdDownlink.Get()
	}
	if o.TimeThreshold.IsSet() {
		toSerialize["timeThreshold"] = o.TimeThreshold.Get()
	}
	if o.MonitoringTime.IsSet() {
		toSerialize["monitoringTime"] = o.MonitoringTime.Get()
	}
	if o.NextVolThreshold.IsSet() {
		toSerialize["nextVolThreshold"] = o.NextVolThreshold.Get()
	}
	if o.NextVolThresholdUplink.IsSet() {
		toSerialize["nextVolThresholdUplink"] = o.NextVolThresholdUplink.Get()
	}
	if o.NextVolThresholdDownlink.IsSet() {
		toSerialize["nextVolThresholdDownlink"] = o.NextVolThresholdDownlink.Get()
	}
	if o.NextTimeThreshold.IsSet() {
		toSerialize["nextTimeThreshold"] = o.NextTimeThreshold.Get()
	}
	if o.InactivityTime.IsSet() {
		toSerialize["inactivityTime"] = o.InactivityTime.Get()
	}
	if o.ExUsagePccRuleIds != nil {
		toSerialize["exUsagePccRuleIds"] = o.ExUsagePccRuleIds
	}
	return json.Marshal(toSerialize)
}

type NullableUsageMonitoringData struct {
	value *UsageMonitoringData
	isSet bool
}

func (v NullableUsageMonitoringData) Get() *UsageMonitoringData {
	return v.value
}

func (v *NullableUsageMonitoringData) Set(val *UsageMonitoringData) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageMonitoringData) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageMonitoringData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageMonitoringData(val *UsageMonitoringData) *NullableUsageMonitoringData {
	return &NullableUsageMonitoringData{value: val, isSet: true}
}

func (v NullableUsageMonitoringData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageMonitoringData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


