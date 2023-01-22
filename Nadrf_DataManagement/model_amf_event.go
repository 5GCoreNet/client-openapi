/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nadrf_DataManagement

import (
	"encoding/json"
	"time"
)

// AmfEvent Describes an event to be subscribed
type AmfEvent struct {
	Type AmfEventType `json:"type"`
	ImmediateFlag *bool `json:"immediateFlag,omitempty"`
	AreaList []AmfEventArea `json:"areaList,omitempty"`
	LocationFilterList []LocationFilter `json:"locationFilterList,omitempty"`
	RefId *int32 `json:"refId,omitempty"`
	TrafficDescriptorList []TrafficDescriptor `json:"trafficDescriptorList,omitempty"`
	ReportUeReachable *bool `json:"reportUeReachable,omitempty"`
	ReachabilityFilter *ReachabilityFilter `json:"reachabilityFilter,omitempty"`
	UdmDetectInd *bool `json:"udmDetectInd,omitempty"`
	MaxReports *int32 `json:"maxReports,omitempty"`
	// A map(list of key-value pairs) where praId serves as key.
	PresenceInfoList *map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	// indicating a time in seconds.
	MaxResponseTime *int32 `json:"maxResponseTime,omitempty"`
	TargetArea *TargetArea `json:"targetArea,omitempty"`
	SnssaiFilter []ExtSnssai `json:"snssaiFilter,omitempty"`
	UeInAreaFilter *UeInAreaFilter `json:"ueInAreaFilter,omitempty"`
	// indicating a time in seconds.
	MinInterval *int32 `json:"minInterval,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	NextReport *time.Time `json:"nextReport,omitempty"`
	IdleStatusInd *bool `json:"idleStatusInd,omitempty"`
	DispersionArea *DispersionArea `json:"dispersionArea,omitempty"`
}

// NewAmfEvent instantiates a new AmfEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfEvent(type_ AmfEventType) *AmfEvent {
	this := AmfEvent{}
	this.Type = type_
	var immediateFlag bool = false
	this.ImmediateFlag = &immediateFlag
	var reportUeReachable bool = false
	this.ReportUeReachable = &reportUeReachable
	var udmDetectInd bool = false
	this.UdmDetectInd = &udmDetectInd
	var idleStatusInd bool = false
	this.IdleStatusInd = &idleStatusInd
	return &this
}

// NewAmfEventWithDefaults instantiates a new AmfEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfEventWithDefaults() *AmfEvent {
	this := AmfEvent{}
	var immediateFlag bool = false
	this.ImmediateFlag = &immediateFlag
	var reportUeReachable bool = false
	this.ReportUeReachable = &reportUeReachable
	var udmDetectInd bool = false
	this.UdmDetectInd = &udmDetectInd
	var idleStatusInd bool = false
	this.IdleStatusInd = &idleStatusInd
	return &this
}

// GetType returns the Type field value
func (o *AmfEvent) GetType() AmfEventType {
	if o == nil {
		var ret AmfEventType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetTypeOk() (*AmfEventType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AmfEvent) SetType(v AmfEventType) {
	o.Type = v
}

// GetImmediateFlag returns the ImmediateFlag field value if set, zero value otherwise.
func (o *AmfEvent) GetImmediateFlag() bool {
	if o == nil || isNil(o.ImmediateFlag) {
		var ret bool
		return ret
	}
	return *o.ImmediateFlag
}

// GetImmediateFlagOk returns a tuple with the ImmediateFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetImmediateFlagOk() (*bool, bool) {
	if o == nil || isNil(o.ImmediateFlag) {
    return nil, false
	}
	return o.ImmediateFlag, true
}

// HasImmediateFlag returns a boolean if a field has been set.
func (o *AmfEvent) HasImmediateFlag() bool {
	if o != nil && !isNil(o.ImmediateFlag) {
		return true
	}

	return false
}

// SetImmediateFlag gets a reference to the given bool and assigns it to the ImmediateFlag field.
func (o *AmfEvent) SetImmediateFlag(v bool) {
	o.ImmediateFlag = &v
}

// GetAreaList returns the AreaList field value if set, zero value otherwise.
func (o *AmfEvent) GetAreaList() []AmfEventArea {
	if o == nil || isNil(o.AreaList) {
		var ret []AmfEventArea
		return ret
	}
	return o.AreaList
}

// GetAreaListOk returns a tuple with the AreaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetAreaListOk() ([]AmfEventArea, bool) {
	if o == nil || isNil(o.AreaList) {
    return nil, false
	}
	return o.AreaList, true
}

// HasAreaList returns a boolean if a field has been set.
func (o *AmfEvent) HasAreaList() bool {
	if o != nil && !isNil(o.AreaList) {
		return true
	}

	return false
}

// SetAreaList gets a reference to the given []AmfEventArea and assigns it to the AreaList field.
func (o *AmfEvent) SetAreaList(v []AmfEventArea) {
	o.AreaList = v
}

// GetLocationFilterList returns the LocationFilterList field value if set, zero value otherwise.
func (o *AmfEvent) GetLocationFilterList() []LocationFilter {
	if o == nil || isNil(o.LocationFilterList) {
		var ret []LocationFilter
		return ret
	}
	return o.LocationFilterList
}

// GetLocationFilterListOk returns a tuple with the LocationFilterList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetLocationFilterListOk() ([]LocationFilter, bool) {
	if o == nil || isNil(o.LocationFilterList) {
    return nil, false
	}
	return o.LocationFilterList, true
}

// HasLocationFilterList returns a boolean if a field has been set.
func (o *AmfEvent) HasLocationFilterList() bool {
	if o != nil && !isNil(o.LocationFilterList) {
		return true
	}

	return false
}

// SetLocationFilterList gets a reference to the given []LocationFilter and assigns it to the LocationFilterList field.
func (o *AmfEvent) SetLocationFilterList(v []LocationFilter) {
	o.LocationFilterList = v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *AmfEvent) GetRefId() int32 {
	if o == nil || isNil(o.RefId) {
		var ret int32
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetRefIdOk() (*int32, bool) {
	if o == nil || isNil(o.RefId) {
    return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *AmfEvent) HasRefId() bool {
	if o != nil && !isNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given int32 and assigns it to the RefId field.
func (o *AmfEvent) SetRefId(v int32) {
	o.RefId = &v
}

// GetTrafficDescriptorList returns the TrafficDescriptorList field value if set, zero value otherwise.
func (o *AmfEvent) GetTrafficDescriptorList() []TrafficDescriptor {
	if o == nil || isNil(o.TrafficDescriptorList) {
		var ret []TrafficDescriptor
		return ret
	}
	return o.TrafficDescriptorList
}

// GetTrafficDescriptorListOk returns a tuple with the TrafficDescriptorList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetTrafficDescriptorListOk() ([]TrafficDescriptor, bool) {
	if o == nil || isNil(o.TrafficDescriptorList) {
    return nil, false
	}
	return o.TrafficDescriptorList, true
}

// HasTrafficDescriptorList returns a boolean if a field has been set.
func (o *AmfEvent) HasTrafficDescriptorList() bool {
	if o != nil && !isNil(o.TrafficDescriptorList) {
		return true
	}

	return false
}

// SetTrafficDescriptorList gets a reference to the given []TrafficDescriptor and assigns it to the TrafficDescriptorList field.
func (o *AmfEvent) SetTrafficDescriptorList(v []TrafficDescriptor) {
	o.TrafficDescriptorList = v
}

// GetReportUeReachable returns the ReportUeReachable field value if set, zero value otherwise.
func (o *AmfEvent) GetReportUeReachable() bool {
	if o == nil || isNil(o.ReportUeReachable) {
		var ret bool
		return ret
	}
	return *o.ReportUeReachable
}

// GetReportUeReachableOk returns a tuple with the ReportUeReachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetReportUeReachableOk() (*bool, bool) {
	if o == nil || isNil(o.ReportUeReachable) {
    return nil, false
	}
	return o.ReportUeReachable, true
}

// HasReportUeReachable returns a boolean if a field has been set.
func (o *AmfEvent) HasReportUeReachable() bool {
	if o != nil && !isNil(o.ReportUeReachable) {
		return true
	}

	return false
}

// SetReportUeReachable gets a reference to the given bool and assigns it to the ReportUeReachable field.
func (o *AmfEvent) SetReportUeReachable(v bool) {
	o.ReportUeReachable = &v
}

// GetReachabilityFilter returns the ReachabilityFilter field value if set, zero value otherwise.
func (o *AmfEvent) GetReachabilityFilter() ReachabilityFilter {
	if o == nil || isNil(o.ReachabilityFilter) {
		var ret ReachabilityFilter
		return ret
	}
	return *o.ReachabilityFilter
}

// GetReachabilityFilterOk returns a tuple with the ReachabilityFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetReachabilityFilterOk() (*ReachabilityFilter, bool) {
	if o == nil || isNil(o.ReachabilityFilter) {
    return nil, false
	}
	return o.ReachabilityFilter, true
}

// HasReachabilityFilter returns a boolean if a field has been set.
func (o *AmfEvent) HasReachabilityFilter() bool {
	if o != nil && !isNil(o.ReachabilityFilter) {
		return true
	}

	return false
}

// SetReachabilityFilter gets a reference to the given ReachabilityFilter and assigns it to the ReachabilityFilter field.
func (o *AmfEvent) SetReachabilityFilter(v ReachabilityFilter) {
	o.ReachabilityFilter = &v
}

// GetUdmDetectInd returns the UdmDetectInd field value if set, zero value otherwise.
func (o *AmfEvent) GetUdmDetectInd() bool {
	if o == nil || isNil(o.UdmDetectInd) {
		var ret bool
		return ret
	}
	return *o.UdmDetectInd
}

// GetUdmDetectIndOk returns a tuple with the UdmDetectInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetUdmDetectIndOk() (*bool, bool) {
	if o == nil || isNil(o.UdmDetectInd) {
    return nil, false
	}
	return o.UdmDetectInd, true
}

// HasUdmDetectInd returns a boolean if a field has been set.
func (o *AmfEvent) HasUdmDetectInd() bool {
	if o != nil && !isNil(o.UdmDetectInd) {
		return true
	}

	return false
}

// SetUdmDetectInd gets a reference to the given bool and assigns it to the UdmDetectInd field.
func (o *AmfEvent) SetUdmDetectInd(v bool) {
	o.UdmDetectInd = &v
}

// GetMaxReports returns the MaxReports field value if set, zero value otherwise.
func (o *AmfEvent) GetMaxReports() int32 {
	if o == nil || isNil(o.MaxReports) {
		var ret int32
		return ret
	}
	return *o.MaxReports
}

// GetMaxReportsOk returns a tuple with the MaxReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetMaxReportsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxReports) {
    return nil, false
	}
	return o.MaxReports, true
}

// HasMaxReports returns a boolean if a field has been set.
func (o *AmfEvent) HasMaxReports() bool {
	if o != nil && !isNil(o.MaxReports) {
		return true
	}

	return false
}

// SetMaxReports gets a reference to the given int32 and assigns it to the MaxReports field.
func (o *AmfEvent) SetMaxReports(v int32) {
	o.MaxReports = &v
}

// GetPresenceInfoList returns the PresenceInfoList field value if set, zero value otherwise.
func (o *AmfEvent) GetPresenceInfoList() map[string]PresenceInfo {
	if o == nil || isNil(o.PresenceInfoList) {
		var ret map[string]PresenceInfo
		return ret
	}
	return *o.PresenceInfoList
}

// GetPresenceInfoListOk returns a tuple with the PresenceInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetPresenceInfoListOk() (*map[string]PresenceInfo, bool) {
	if o == nil || isNil(o.PresenceInfoList) {
    return nil, false
	}
	return o.PresenceInfoList, true
}

// HasPresenceInfoList returns a boolean if a field has been set.
func (o *AmfEvent) HasPresenceInfoList() bool {
	if o != nil && !isNil(o.PresenceInfoList) {
		return true
	}

	return false
}

// SetPresenceInfoList gets a reference to the given map[string]PresenceInfo and assigns it to the PresenceInfoList field.
func (o *AmfEvent) SetPresenceInfoList(v map[string]PresenceInfo) {
	o.PresenceInfoList = &v
}

// GetMaxResponseTime returns the MaxResponseTime field value if set, zero value otherwise.
func (o *AmfEvent) GetMaxResponseTime() int32 {
	if o == nil || isNil(o.MaxResponseTime) {
		var ret int32
		return ret
	}
	return *o.MaxResponseTime
}

// GetMaxResponseTimeOk returns a tuple with the MaxResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetMaxResponseTimeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxResponseTime) {
    return nil, false
	}
	return o.MaxResponseTime, true
}

// HasMaxResponseTime returns a boolean if a field has been set.
func (o *AmfEvent) HasMaxResponseTime() bool {
	if o != nil && !isNil(o.MaxResponseTime) {
		return true
	}

	return false
}

// SetMaxResponseTime gets a reference to the given int32 and assigns it to the MaxResponseTime field.
func (o *AmfEvent) SetMaxResponseTime(v int32) {
	o.MaxResponseTime = &v
}

// GetTargetArea returns the TargetArea field value if set, zero value otherwise.
func (o *AmfEvent) GetTargetArea() TargetArea {
	if o == nil || isNil(o.TargetArea) {
		var ret TargetArea
		return ret
	}
	return *o.TargetArea
}

// GetTargetAreaOk returns a tuple with the TargetArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetTargetAreaOk() (*TargetArea, bool) {
	if o == nil || isNil(o.TargetArea) {
    return nil, false
	}
	return o.TargetArea, true
}

// HasTargetArea returns a boolean if a field has been set.
func (o *AmfEvent) HasTargetArea() bool {
	if o != nil && !isNil(o.TargetArea) {
		return true
	}

	return false
}

// SetTargetArea gets a reference to the given TargetArea and assigns it to the TargetArea field.
func (o *AmfEvent) SetTargetArea(v TargetArea) {
	o.TargetArea = &v
}

// GetSnssaiFilter returns the SnssaiFilter field value if set, zero value otherwise.
func (o *AmfEvent) GetSnssaiFilter() []ExtSnssai {
	if o == nil || isNil(o.SnssaiFilter) {
		var ret []ExtSnssai
		return ret
	}
	return o.SnssaiFilter
}

// GetSnssaiFilterOk returns a tuple with the SnssaiFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetSnssaiFilterOk() ([]ExtSnssai, bool) {
	if o == nil || isNil(o.SnssaiFilter) {
    return nil, false
	}
	return o.SnssaiFilter, true
}

// HasSnssaiFilter returns a boolean if a field has been set.
func (o *AmfEvent) HasSnssaiFilter() bool {
	if o != nil && !isNil(o.SnssaiFilter) {
		return true
	}

	return false
}

// SetSnssaiFilter gets a reference to the given []ExtSnssai and assigns it to the SnssaiFilter field.
func (o *AmfEvent) SetSnssaiFilter(v []ExtSnssai) {
	o.SnssaiFilter = v
}

// GetUeInAreaFilter returns the UeInAreaFilter field value if set, zero value otherwise.
func (o *AmfEvent) GetUeInAreaFilter() UeInAreaFilter {
	if o == nil || isNil(o.UeInAreaFilter) {
		var ret UeInAreaFilter
		return ret
	}
	return *o.UeInAreaFilter
}

// GetUeInAreaFilterOk returns a tuple with the UeInAreaFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetUeInAreaFilterOk() (*UeInAreaFilter, bool) {
	if o == nil || isNil(o.UeInAreaFilter) {
    return nil, false
	}
	return o.UeInAreaFilter, true
}

// HasUeInAreaFilter returns a boolean if a field has been set.
func (o *AmfEvent) HasUeInAreaFilter() bool {
	if o != nil && !isNil(o.UeInAreaFilter) {
		return true
	}

	return false
}

// SetUeInAreaFilter gets a reference to the given UeInAreaFilter and assigns it to the UeInAreaFilter field.
func (o *AmfEvent) SetUeInAreaFilter(v UeInAreaFilter) {
	o.UeInAreaFilter = &v
}

// GetMinInterval returns the MinInterval field value if set, zero value otherwise.
func (o *AmfEvent) GetMinInterval() int32 {
	if o == nil || isNil(o.MinInterval) {
		var ret int32
		return ret
	}
	return *o.MinInterval
}

// GetMinIntervalOk returns a tuple with the MinInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetMinIntervalOk() (*int32, bool) {
	if o == nil || isNil(o.MinInterval) {
    return nil, false
	}
	return o.MinInterval, true
}

// HasMinInterval returns a boolean if a field has been set.
func (o *AmfEvent) HasMinInterval() bool {
	if o != nil && !isNil(o.MinInterval) {
		return true
	}

	return false
}

// SetMinInterval gets a reference to the given int32 and assigns it to the MinInterval field.
func (o *AmfEvent) SetMinInterval(v int32) {
	o.MinInterval = &v
}

// GetNextReport returns the NextReport field value if set, zero value otherwise.
func (o *AmfEvent) GetNextReport() time.Time {
	if o == nil || isNil(o.NextReport) {
		var ret time.Time
		return ret
	}
	return *o.NextReport
}

// GetNextReportOk returns a tuple with the NextReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetNextReportOk() (*time.Time, bool) {
	if o == nil || isNil(o.NextReport) {
    return nil, false
	}
	return o.NextReport, true
}

// HasNextReport returns a boolean if a field has been set.
func (o *AmfEvent) HasNextReport() bool {
	if o != nil && !isNil(o.NextReport) {
		return true
	}

	return false
}

// SetNextReport gets a reference to the given time.Time and assigns it to the NextReport field.
func (o *AmfEvent) SetNextReport(v time.Time) {
	o.NextReport = &v
}

// GetIdleStatusInd returns the IdleStatusInd field value if set, zero value otherwise.
func (o *AmfEvent) GetIdleStatusInd() bool {
	if o == nil || isNil(o.IdleStatusInd) {
		var ret bool
		return ret
	}
	return *o.IdleStatusInd
}

// GetIdleStatusIndOk returns a tuple with the IdleStatusInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetIdleStatusIndOk() (*bool, bool) {
	if o == nil || isNil(o.IdleStatusInd) {
    return nil, false
	}
	return o.IdleStatusInd, true
}

// HasIdleStatusInd returns a boolean if a field has been set.
func (o *AmfEvent) HasIdleStatusInd() bool {
	if o != nil && !isNil(o.IdleStatusInd) {
		return true
	}

	return false
}

// SetIdleStatusInd gets a reference to the given bool and assigns it to the IdleStatusInd field.
func (o *AmfEvent) SetIdleStatusInd(v bool) {
	o.IdleStatusInd = &v
}

// GetDispersionArea returns the DispersionArea field value if set, zero value otherwise.
func (o *AmfEvent) GetDispersionArea() DispersionArea {
	if o == nil || isNil(o.DispersionArea) {
		var ret DispersionArea
		return ret
	}
	return *o.DispersionArea
}

// GetDispersionAreaOk returns a tuple with the DispersionArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEvent) GetDispersionAreaOk() (*DispersionArea, bool) {
	if o == nil || isNil(o.DispersionArea) {
    return nil, false
	}
	return o.DispersionArea, true
}

// HasDispersionArea returns a boolean if a field has been set.
func (o *AmfEvent) HasDispersionArea() bool {
	if o != nil && !isNil(o.DispersionArea) {
		return true
	}

	return false
}

// SetDispersionArea gets a reference to the given DispersionArea and assigns it to the DispersionArea field.
func (o *AmfEvent) SetDispersionArea(v DispersionArea) {
	o.DispersionArea = &v
}

func (o AmfEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.ImmediateFlag) {
		toSerialize["immediateFlag"] = o.ImmediateFlag
	}
	if !isNil(o.AreaList) {
		toSerialize["areaList"] = o.AreaList
	}
	if !isNil(o.LocationFilterList) {
		toSerialize["locationFilterList"] = o.LocationFilterList
	}
	if !isNil(o.RefId) {
		toSerialize["refId"] = o.RefId
	}
	if !isNil(o.TrafficDescriptorList) {
		toSerialize["trafficDescriptorList"] = o.TrafficDescriptorList
	}
	if !isNil(o.ReportUeReachable) {
		toSerialize["reportUeReachable"] = o.ReportUeReachable
	}
	if !isNil(o.ReachabilityFilter) {
		toSerialize["reachabilityFilter"] = o.ReachabilityFilter
	}
	if !isNil(o.UdmDetectInd) {
		toSerialize["udmDetectInd"] = o.UdmDetectInd
	}
	if !isNil(o.MaxReports) {
		toSerialize["maxReports"] = o.MaxReports
	}
	if !isNil(o.PresenceInfoList) {
		toSerialize["presenceInfoList"] = o.PresenceInfoList
	}
	if !isNil(o.MaxResponseTime) {
		toSerialize["maxResponseTime"] = o.MaxResponseTime
	}
	if !isNil(o.TargetArea) {
		toSerialize["targetArea"] = o.TargetArea
	}
	if !isNil(o.SnssaiFilter) {
		toSerialize["snssaiFilter"] = o.SnssaiFilter
	}
	if !isNil(o.UeInAreaFilter) {
		toSerialize["ueInAreaFilter"] = o.UeInAreaFilter
	}
	if !isNil(o.MinInterval) {
		toSerialize["minInterval"] = o.MinInterval
	}
	if !isNil(o.NextReport) {
		toSerialize["nextReport"] = o.NextReport
	}
	if !isNil(o.IdleStatusInd) {
		toSerialize["idleStatusInd"] = o.IdleStatusInd
	}
	if !isNil(o.DispersionArea) {
		toSerialize["dispersionArea"] = o.DispersionArea
	}
	return json.Marshal(toSerialize)
}

type NullableAmfEvent struct {
	value *AmfEvent
	isSet bool
}

func (v NullableAmfEvent) Get() *AmfEvent {
	return v.value
}

func (v *NullableAmfEvent) Set(val *AmfEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEvent(val *AmfEvent) *NullableAmfEvent {
	return &NullableAmfEvent{value: val, isSet: true}
}

func (v NullableAmfEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


