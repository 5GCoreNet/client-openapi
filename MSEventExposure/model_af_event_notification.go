/*
3gpp-ms-event-exposure

API for Media Streaming Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSEventExposure

import (
	"encoding/json"
	"time"
)

// AfEventNotification Represents information related to an event to be reported.
type AfEventNotification struct {
	Event AfEvent `json:"event"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
	SvcExprcInfos []ServiceExperienceInfoPerApp `json:"svcExprcInfos,omitempty"`
	UeMobilityInfos []UeMobilityCollection `json:"ueMobilityInfos,omitempty"`
	UeCommInfos []UeCommunicationCollection `json:"ueCommInfos,omitempty"`
	ExcepInfos []ExceptionInfo `json:"excepInfos,omitempty"`
	CongestionInfos []UserDataCongestionCollection `json:"congestionInfos,omitempty"`
	PerfDataInfos []PerformanceDataCollection `json:"perfDataInfos,omitempty"`
	DispersionInfos []DispersionCollection `json:"dispersionInfos,omitempty"`
	CollBhvrInfs []CollectiveBehaviourInfo `json:"collBhvrInfs,omitempty"`
	MsQoeMetrInfos []MsQoeMetricsCollection `json:"msQoeMetrInfos,omitempty"`
	MsConsumpInfos []MsConsumptionCollection `json:"msConsumpInfos,omitempty"`
	MsNetAssInvInfos []MsNetAssInvocationCollection `json:"msNetAssInvInfos,omitempty"`
	MsDynPlyInvInfos []MsDynPolicyInvocationCollection `json:"msDynPlyInvInfos,omitempty"`
	MsAccActInfos []MSAccessActivityCollection `json:"msAccActInfos,omitempty"`
}

// NewAfEventNotification instantiates a new AfEventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfEventNotification(event AfEvent, timeStamp time.Time) *AfEventNotification {
	this := AfEventNotification{}
	this.Event = event
	this.TimeStamp = timeStamp
	return &this
}

// NewAfEventNotificationWithDefaults instantiates a new AfEventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfEventNotificationWithDefaults() *AfEventNotification {
	this := AfEventNotification{}
	return &this
}

// GetEvent returns the Event field value
func (o *AfEventNotification) GetEvent() AfEvent {
	if o == nil {
		var ret AfEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetEventOk() (*AfEvent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *AfEventNotification) SetEvent(v AfEvent) {
	o.Event = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *AfEventNotification) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *AfEventNotification) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

// GetSvcExprcInfos returns the SvcExprcInfos field value if set, zero value otherwise.
func (o *AfEventNotification) GetSvcExprcInfos() []ServiceExperienceInfoPerApp {
	if o == nil || isNil(o.SvcExprcInfos) {
		var ret []ServiceExperienceInfoPerApp
		return ret
	}
	return o.SvcExprcInfos
}

// GetSvcExprcInfosOk returns a tuple with the SvcExprcInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetSvcExprcInfosOk() ([]ServiceExperienceInfoPerApp, bool) {
	if o == nil || isNil(o.SvcExprcInfos) {
    return nil, false
	}
	return o.SvcExprcInfos, true
}

// HasSvcExprcInfos returns a boolean if a field has been set.
func (o *AfEventNotification) HasSvcExprcInfos() bool {
	if o != nil && !isNil(o.SvcExprcInfos) {
		return true
	}

	return false
}

// SetSvcExprcInfos gets a reference to the given []ServiceExperienceInfoPerApp and assigns it to the SvcExprcInfos field.
func (o *AfEventNotification) SetSvcExprcInfos(v []ServiceExperienceInfoPerApp) {
	o.SvcExprcInfos = v
}

// GetUeMobilityInfos returns the UeMobilityInfos field value if set, zero value otherwise.
func (o *AfEventNotification) GetUeMobilityInfos() []UeMobilityCollection {
	if o == nil || isNil(o.UeMobilityInfos) {
		var ret []UeMobilityCollection
		return ret
	}
	return o.UeMobilityInfos
}

// GetUeMobilityInfosOk returns a tuple with the UeMobilityInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetUeMobilityInfosOk() ([]UeMobilityCollection, bool) {
	if o == nil || isNil(o.UeMobilityInfos) {
    return nil, false
	}
	return o.UeMobilityInfos, true
}

// HasUeMobilityInfos returns a boolean if a field has been set.
func (o *AfEventNotification) HasUeMobilityInfos() bool {
	if o != nil && !isNil(o.UeMobilityInfos) {
		return true
	}

	return false
}

// SetUeMobilityInfos gets a reference to the given []UeMobilityCollection and assigns it to the UeMobilityInfos field.
func (o *AfEventNotification) SetUeMobilityInfos(v []UeMobilityCollection) {
	o.UeMobilityInfos = v
}

// GetUeCommInfos returns the UeCommInfos field value if set, zero value otherwise.
func (o *AfEventNotification) GetUeCommInfos() []UeCommunicationCollection {
	if o == nil || isNil(o.UeCommInfos) {
		var ret []UeCommunicationCollection
		return ret
	}
	return o.UeCommInfos
}

// GetUeCommInfosOk returns a tuple with the UeCommInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetUeCommInfosOk() ([]UeCommunicationCollection, bool) {
	if o == nil || isNil(o.UeCommInfos) {
    return nil, false
	}
	return o.UeCommInfos, true
}

// HasUeCommInfos returns a boolean if a field has been set.
func (o *AfEventNotification) HasUeCommInfos() bool {
	if o != nil && !isNil(o.UeCommInfos) {
		return true
	}

	return false
}

// SetUeCommInfos gets a reference to the given []UeCommunicationCollection and assigns it to the UeCommInfos field.
func (o *AfEventNotification) SetUeCommInfos(v []UeCommunicationCollection) {
	o.UeCommInfos = v
}

// GetExcepInfos returns the ExcepInfos field value if set, zero value otherwise.
func (o *AfEventNotification) GetExcepInfos() []ExceptionInfo {
	if o == nil || isNil(o.ExcepInfos) {
		var ret []ExceptionInfo
		return ret
	}
	return o.ExcepInfos
}

// GetExcepInfosOk returns a tuple with the ExcepInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetExcepInfosOk() ([]ExceptionInfo, bool) {
	if o == nil || isNil(o.ExcepInfos) {
    return nil, false
	}
	return o.ExcepInfos, true
}

// HasExcepInfos returns a boolean if a field has been set.
func (o *AfEventNotification) HasExcepInfos() bool {
	if o != nil && !isNil(o.ExcepInfos) {
		return true
	}

	return false
}

// SetExcepInfos gets a reference to the given []ExceptionInfo and assigns it to the ExcepInfos field.
func (o *AfEventNotification) SetExcepInfos(v []ExceptionInfo) {
	o.ExcepInfos = v
}

// GetCongestionInfos returns the CongestionInfos field value if set, zero value otherwise.
func (o *AfEventNotification) GetCongestionInfos() []UserDataCongestionCollection {
	if o == nil || isNil(o.CongestionInfos) {
		var ret []UserDataCongestionCollection
		return ret
	}
	return o.CongestionInfos
}

// GetCongestionInfosOk returns a tuple with the CongestionInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetCongestionInfosOk() ([]UserDataCongestionCollection, bool) {
	if o == nil || isNil(o.CongestionInfos) {
    return nil, false
	}
	return o.CongestionInfos, true
}

// HasCongestionInfos returns a boolean if a field has been set.
func (o *AfEventNotification) HasCongestionInfos() bool {
	if o != nil && !isNil(o.CongestionInfos) {
		return true
	}

	return false
}

// SetCongestionInfos gets a reference to the given []UserDataCongestionCollection and assigns it to the CongestionInfos field.
func (o *AfEventNotification) SetCongestionInfos(v []UserDataCongestionCollection) {
	o.CongestionInfos = v
}

// GetPerfDataInfos returns the PerfDataInfos field value if set, zero value otherwise.
func (o *AfEventNotification) GetPerfDataInfos() []PerformanceDataCollection {
	if o == nil || isNil(o.PerfDataInfos) {
		var ret []PerformanceDataCollection
		return ret
	}
	return o.PerfDataInfos
}

// GetPerfDataInfosOk returns a tuple with the PerfDataInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetPerfDataInfosOk() ([]PerformanceDataCollection, bool) {
	if o == nil || isNil(o.PerfDataInfos) {
    return nil, false
	}
	return o.PerfDataInfos, true
}

// HasPerfDataInfos returns a boolean if a field has been set.
func (o *AfEventNotification) HasPerfDataInfos() bool {
	if o != nil && !isNil(o.PerfDataInfos) {
		return true
	}

	return false
}

// SetPerfDataInfos gets a reference to the given []PerformanceDataCollection and assigns it to the PerfDataInfos field.
func (o *AfEventNotification) SetPerfDataInfos(v []PerformanceDataCollection) {
	o.PerfDataInfos = v
}

// GetDispersionInfos returns the DispersionInfos field value if set, zero value otherwise.
func (o *AfEventNotification) GetDispersionInfos() []DispersionCollection {
	if o == nil || isNil(o.DispersionInfos) {
		var ret []DispersionCollection
		return ret
	}
	return o.DispersionInfos
}

// GetDispersionInfosOk returns a tuple with the DispersionInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetDispersionInfosOk() ([]DispersionCollection, bool) {
	if o == nil || isNil(o.DispersionInfos) {
    return nil, false
	}
	return o.DispersionInfos, true
}

// HasDispersionInfos returns a boolean if a field has been set.
func (o *AfEventNotification) HasDispersionInfos() bool {
	if o != nil && !isNil(o.DispersionInfos) {
		return true
	}

	return false
}

// SetDispersionInfos gets a reference to the given []DispersionCollection and assigns it to the DispersionInfos field.
func (o *AfEventNotification) SetDispersionInfos(v []DispersionCollection) {
	o.DispersionInfos = v
}

// GetCollBhvrInfs returns the CollBhvrInfs field value if set, zero value otherwise.
func (o *AfEventNotification) GetCollBhvrInfs() []CollectiveBehaviourInfo {
	if o == nil || isNil(o.CollBhvrInfs) {
		var ret []CollectiveBehaviourInfo
		return ret
	}
	return o.CollBhvrInfs
}

// GetCollBhvrInfsOk returns a tuple with the CollBhvrInfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetCollBhvrInfsOk() ([]CollectiveBehaviourInfo, bool) {
	if o == nil || isNil(o.CollBhvrInfs) {
    return nil, false
	}
	return o.CollBhvrInfs, true
}

// HasCollBhvrInfs returns a boolean if a field has been set.
func (o *AfEventNotification) HasCollBhvrInfs() bool {
	if o != nil && !isNil(o.CollBhvrInfs) {
		return true
	}

	return false
}

// SetCollBhvrInfs gets a reference to the given []CollectiveBehaviourInfo and assigns it to the CollBhvrInfs field.
func (o *AfEventNotification) SetCollBhvrInfs(v []CollectiveBehaviourInfo) {
	o.CollBhvrInfs = v
}

// GetMsQoeMetrInfos returns the MsQoeMetrInfos field value if set, zero value otherwise.
func (o *AfEventNotification) GetMsQoeMetrInfos() []MsQoeMetricsCollection {
	if o == nil || isNil(o.MsQoeMetrInfos) {
		var ret []MsQoeMetricsCollection
		return ret
	}
	return o.MsQoeMetrInfos
}

// GetMsQoeMetrInfosOk returns a tuple with the MsQoeMetrInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetMsQoeMetrInfosOk() ([]MsQoeMetricsCollection, bool) {
	if o == nil || isNil(o.MsQoeMetrInfos) {
    return nil, false
	}
	return o.MsQoeMetrInfos, true
}

// HasMsQoeMetrInfos returns a boolean if a field has been set.
func (o *AfEventNotification) HasMsQoeMetrInfos() bool {
	if o != nil && !isNil(o.MsQoeMetrInfos) {
		return true
	}

	return false
}

// SetMsQoeMetrInfos gets a reference to the given []MsQoeMetricsCollection and assigns it to the MsQoeMetrInfos field.
func (o *AfEventNotification) SetMsQoeMetrInfos(v []MsQoeMetricsCollection) {
	o.MsQoeMetrInfos = v
}

// GetMsConsumpInfos returns the MsConsumpInfos field value if set, zero value otherwise.
func (o *AfEventNotification) GetMsConsumpInfos() []MsConsumptionCollection {
	if o == nil || isNil(o.MsConsumpInfos) {
		var ret []MsConsumptionCollection
		return ret
	}
	return o.MsConsumpInfos
}

// GetMsConsumpInfosOk returns a tuple with the MsConsumpInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetMsConsumpInfosOk() ([]MsConsumptionCollection, bool) {
	if o == nil || isNil(o.MsConsumpInfos) {
    return nil, false
	}
	return o.MsConsumpInfos, true
}

// HasMsConsumpInfos returns a boolean if a field has been set.
func (o *AfEventNotification) HasMsConsumpInfos() bool {
	if o != nil && !isNil(o.MsConsumpInfos) {
		return true
	}

	return false
}

// SetMsConsumpInfos gets a reference to the given []MsConsumptionCollection and assigns it to the MsConsumpInfos field.
func (o *AfEventNotification) SetMsConsumpInfos(v []MsConsumptionCollection) {
	o.MsConsumpInfos = v
}

// GetMsNetAssInvInfos returns the MsNetAssInvInfos field value if set, zero value otherwise.
func (o *AfEventNotification) GetMsNetAssInvInfos() []MsNetAssInvocationCollection {
	if o == nil || isNil(o.MsNetAssInvInfos) {
		var ret []MsNetAssInvocationCollection
		return ret
	}
	return o.MsNetAssInvInfos
}

// GetMsNetAssInvInfosOk returns a tuple with the MsNetAssInvInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetMsNetAssInvInfosOk() ([]MsNetAssInvocationCollection, bool) {
	if o == nil || isNil(o.MsNetAssInvInfos) {
    return nil, false
	}
	return o.MsNetAssInvInfos, true
}

// HasMsNetAssInvInfos returns a boolean if a field has been set.
func (o *AfEventNotification) HasMsNetAssInvInfos() bool {
	if o != nil && !isNil(o.MsNetAssInvInfos) {
		return true
	}

	return false
}

// SetMsNetAssInvInfos gets a reference to the given []MsNetAssInvocationCollection and assigns it to the MsNetAssInvInfos field.
func (o *AfEventNotification) SetMsNetAssInvInfos(v []MsNetAssInvocationCollection) {
	o.MsNetAssInvInfos = v
}

// GetMsDynPlyInvInfos returns the MsDynPlyInvInfos field value if set, zero value otherwise.
func (o *AfEventNotification) GetMsDynPlyInvInfos() []MsDynPolicyInvocationCollection {
	if o == nil || isNil(o.MsDynPlyInvInfos) {
		var ret []MsDynPolicyInvocationCollection
		return ret
	}
	return o.MsDynPlyInvInfos
}

// GetMsDynPlyInvInfosOk returns a tuple with the MsDynPlyInvInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetMsDynPlyInvInfosOk() ([]MsDynPolicyInvocationCollection, bool) {
	if o == nil || isNil(o.MsDynPlyInvInfos) {
    return nil, false
	}
	return o.MsDynPlyInvInfos, true
}

// HasMsDynPlyInvInfos returns a boolean if a field has been set.
func (o *AfEventNotification) HasMsDynPlyInvInfos() bool {
	if o != nil && !isNil(o.MsDynPlyInvInfos) {
		return true
	}

	return false
}

// SetMsDynPlyInvInfos gets a reference to the given []MsDynPolicyInvocationCollection and assigns it to the MsDynPlyInvInfos field.
func (o *AfEventNotification) SetMsDynPlyInvInfos(v []MsDynPolicyInvocationCollection) {
	o.MsDynPlyInvInfos = v
}

// GetMsAccActInfos returns the MsAccActInfos field value if set, zero value otherwise.
func (o *AfEventNotification) GetMsAccActInfos() []MSAccessActivityCollection {
	if o == nil || isNil(o.MsAccActInfos) {
		var ret []MSAccessActivityCollection
		return ret
	}
	return o.MsAccActInfos
}

// GetMsAccActInfosOk returns a tuple with the MsAccActInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfEventNotification) GetMsAccActInfosOk() ([]MSAccessActivityCollection, bool) {
	if o == nil || isNil(o.MsAccActInfos) {
    return nil, false
	}
	return o.MsAccActInfos, true
}

// HasMsAccActInfos returns a boolean if a field has been set.
func (o *AfEventNotification) HasMsAccActInfos() bool {
	if o != nil && !isNil(o.MsAccActInfos) {
		return true
	}

	return false
}

// SetMsAccActInfos gets a reference to the given []MSAccessActivityCollection and assigns it to the MsAccActInfos field.
func (o *AfEventNotification) SetMsAccActInfos(v []MSAccessActivityCollection) {
	o.MsAccActInfos = v
}

func (o AfEventNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	if !isNil(o.SvcExprcInfos) {
		toSerialize["svcExprcInfos"] = o.SvcExprcInfos
	}
	if !isNil(o.UeMobilityInfos) {
		toSerialize["ueMobilityInfos"] = o.UeMobilityInfos
	}
	if !isNil(o.UeCommInfos) {
		toSerialize["ueCommInfos"] = o.UeCommInfos
	}
	if !isNil(o.ExcepInfos) {
		toSerialize["excepInfos"] = o.ExcepInfos
	}
	if !isNil(o.CongestionInfos) {
		toSerialize["congestionInfos"] = o.CongestionInfos
	}
	if !isNil(o.PerfDataInfos) {
		toSerialize["perfDataInfos"] = o.PerfDataInfos
	}
	if !isNil(o.DispersionInfos) {
		toSerialize["dispersionInfos"] = o.DispersionInfos
	}
	if !isNil(o.CollBhvrInfs) {
		toSerialize["collBhvrInfs"] = o.CollBhvrInfs
	}
	if !isNil(o.MsQoeMetrInfos) {
		toSerialize["msQoeMetrInfos"] = o.MsQoeMetrInfos
	}
	if !isNil(o.MsConsumpInfos) {
		toSerialize["msConsumpInfos"] = o.MsConsumpInfos
	}
	if !isNil(o.MsNetAssInvInfos) {
		toSerialize["msNetAssInvInfos"] = o.MsNetAssInvInfos
	}
	if !isNil(o.MsDynPlyInvInfos) {
		toSerialize["msDynPlyInvInfos"] = o.MsDynPlyInvInfos
	}
	if !isNil(o.MsAccActInfos) {
		toSerialize["msAccActInfos"] = o.MsAccActInfos
	}
	return json.Marshal(toSerialize)
}

type NullableAfEventNotification struct {
	value *AfEventNotification
	isSet bool
}

func (v NullableAfEventNotification) Get() *AfEventNotification {
	return v.value
}

func (v *NullableAfEventNotification) Set(val *AfEventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAfEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAfEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfEventNotification(val *AfEventNotification) *NullableAfEventNotification {
	return &NullableAfEventNotification{value: val, isSet: true}
}

func (v NullableAfEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


