/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nmbsmf_MBSSession

import (
	"encoding/json"
	"time"
)

// ContextStatusEventReport Context Status Event Report
type ContextStatusEventReport struct {
	EventType ContextStatusEventType `json:"eventType"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
	QosInfo *QosInfo `json:"qosInfo,omitempty"`
	StatusInfo *MbsSessionActivityStatus `json:"statusInfo,omitempty"`
	MbsServiceArea *MbsServiceArea `json:"mbsServiceArea,omitempty"`
	// A map (list of key-value pairs) where the key identifies an areaSessionId 
	MbsServiceAreaInfoList *map[string]MbsServiceAreaInfo `json:"mbsServiceAreaInfoList,omitempty"`
	MulticastTransAddInfo *MulticastTransportAddressChangeInfo `json:"multicastTransAddInfo,omitempty"`
	MbsSecurityContext *MbsSecurityContext `json:"mbsSecurityContext,omitempty"`
}

// NewContextStatusEventReport instantiates a new ContextStatusEventReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextStatusEventReport(eventType ContextStatusEventType, timeStamp time.Time) *ContextStatusEventReport {
	this := ContextStatusEventReport{}
	this.EventType = eventType
	this.TimeStamp = timeStamp
	return &this
}

// NewContextStatusEventReportWithDefaults instantiates a new ContextStatusEventReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextStatusEventReportWithDefaults() *ContextStatusEventReport {
	this := ContextStatusEventReport{}
	return &this
}

// GetEventType returns the EventType field value
func (o *ContextStatusEventReport) GetEventType() ContextStatusEventType {
	if o == nil {
		var ret ContextStatusEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ContextStatusEventReport) GetEventTypeOk() (*ContextStatusEventType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *ContextStatusEventReport) SetEventType(v ContextStatusEventType) {
	o.EventType = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *ContextStatusEventReport) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *ContextStatusEventReport) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *ContextStatusEventReport) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

// GetQosInfo returns the QosInfo field value if set, zero value otherwise.
func (o *ContextStatusEventReport) GetQosInfo() QosInfo {
	if o == nil || isNil(o.QosInfo) {
		var ret QosInfo
		return ret
	}
	return *o.QosInfo
}

// GetQosInfoOk returns a tuple with the QosInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextStatusEventReport) GetQosInfoOk() (*QosInfo, bool) {
	if o == nil || isNil(o.QosInfo) {
    return nil, false
	}
	return o.QosInfo, true
}

// HasQosInfo returns a boolean if a field has been set.
func (o *ContextStatusEventReport) HasQosInfo() bool {
	if o != nil && !isNil(o.QosInfo) {
		return true
	}

	return false
}

// SetQosInfo gets a reference to the given QosInfo and assigns it to the QosInfo field.
func (o *ContextStatusEventReport) SetQosInfo(v QosInfo) {
	o.QosInfo = &v
}

// GetStatusInfo returns the StatusInfo field value if set, zero value otherwise.
func (o *ContextStatusEventReport) GetStatusInfo() MbsSessionActivityStatus {
	if o == nil || isNil(o.StatusInfo) {
		var ret MbsSessionActivityStatus
		return ret
	}
	return *o.StatusInfo
}

// GetStatusInfoOk returns a tuple with the StatusInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextStatusEventReport) GetStatusInfoOk() (*MbsSessionActivityStatus, bool) {
	if o == nil || isNil(o.StatusInfo) {
    return nil, false
	}
	return o.StatusInfo, true
}

// HasStatusInfo returns a boolean if a field has been set.
func (o *ContextStatusEventReport) HasStatusInfo() bool {
	if o != nil && !isNil(o.StatusInfo) {
		return true
	}

	return false
}

// SetStatusInfo gets a reference to the given MbsSessionActivityStatus and assigns it to the StatusInfo field.
func (o *ContextStatusEventReport) SetStatusInfo(v MbsSessionActivityStatus) {
	o.StatusInfo = &v
}

// GetMbsServiceArea returns the MbsServiceArea field value if set, zero value otherwise.
func (o *ContextStatusEventReport) GetMbsServiceArea() MbsServiceArea {
	if o == nil || isNil(o.MbsServiceArea) {
		var ret MbsServiceArea
		return ret
	}
	return *o.MbsServiceArea
}

// GetMbsServiceAreaOk returns a tuple with the MbsServiceArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextStatusEventReport) GetMbsServiceAreaOk() (*MbsServiceArea, bool) {
	if o == nil || isNil(o.MbsServiceArea) {
    return nil, false
	}
	return o.MbsServiceArea, true
}

// HasMbsServiceArea returns a boolean if a field has been set.
func (o *ContextStatusEventReport) HasMbsServiceArea() bool {
	if o != nil && !isNil(o.MbsServiceArea) {
		return true
	}

	return false
}

// SetMbsServiceArea gets a reference to the given MbsServiceArea and assigns it to the MbsServiceArea field.
func (o *ContextStatusEventReport) SetMbsServiceArea(v MbsServiceArea) {
	o.MbsServiceArea = &v
}

// GetMbsServiceAreaInfoList returns the MbsServiceAreaInfoList field value if set, zero value otherwise.
func (o *ContextStatusEventReport) GetMbsServiceAreaInfoList() map[string]MbsServiceAreaInfo {
	if o == nil || isNil(o.MbsServiceAreaInfoList) {
		var ret map[string]MbsServiceAreaInfo
		return ret
	}
	return *o.MbsServiceAreaInfoList
}

// GetMbsServiceAreaInfoListOk returns a tuple with the MbsServiceAreaInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextStatusEventReport) GetMbsServiceAreaInfoListOk() (*map[string]MbsServiceAreaInfo, bool) {
	if o == nil || isNil(o.MbsServiceAreaInfoList) {
    return nil, false
	}
	return o.MbsServiceAreaInfoList, true
}

// HasMbsServiceAreaInfoList returns a boolean if a field has been set.
func (o *ContextStatusEventReport) HasMbsServiceAreaInfoList() bool {
	if o != nil && !isNil(o.MbsServiceAreaInfoList) {
		return true
	}

	return false
}

// SetMbsServiceAreaInfoList gets a reference to the given map[string]MbsServiceAreaInfo and assigns it to the MbsServiceAreaInfoList field.
func (o *ContextStatusEventReport) SetMbsServiceAreaInfoList(v map[string]MbsServiceAreaInfo) {
	o.MbsServiceAreaInfoList = &v
}

// GetMulticastTransAddInfo returns the MulticastTransAddInfo field value if set, zero value otherwise.
func (o *ContextStatusEventReport) GetMulticastTransAddInfo() MulticastTransportAddressChangeInfo {
	if o == nil || isNil(o.MulticastTransAddInfo) {
		var ret MulticastTransportAddressChangeInfo
		return ret
	}
	return *o.MulticastTransAddInfo
}

// GetMulticastTransAddInfoOk returns a tuple with the MulticastTransAddInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextStatusEventReport) GetMulticastTransAddInfoOk() (*MulticastTransportAddressChangeInfo, bool) {
	if o == nil || isNil(o.MulticastTransAddInfo) {
    return nil, false
	}
	return o.MulticastTransAddInfo, true
}

// HasMulticastTransAddInfo returns a boolean if a field has been set.
func (o *ContextStatusEventReport) HasMulticastTransAddInfo() bool {
	if o != nil && !isNil(o.MulticastTransAddInfo) {
		return true
	}

	return false
}

// SetMulticastTransAddInfo gets a reference to the given MulticastTransportAddressChangeInfo and assigns it to the MulticastTransAddInfo field.
func (o *ContextStatusEventReport) SetMulticastTransAddInfo(v MulticastTransportAddressChangeInfo) {
	o.MulticastTransAddInfo = &v
}

// GetMbsSecurityContext returns the MbsSecurityContext field value if set, zero value otherwise.
func (o *ContextStatusEventReport) GetMbsSecurityContext() MbsSecurityContext {
	if o == nil || isNil(o.MbsSecurityContext) {
		var ret MbsSecurityContext
		return ret
	}
	return *o.MbsSecurityContext
}

// GetMbsSecurityContextOk returns a tuple with the MbsSecurityContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextStatusEventReport) GetMbsSecurityContextOk() (*MbsSecurityContext, bool) {
	if o == nil || isNil(o.MbsSecurityContext) {
    return nil, false
	}
	return o.MbsSecurityContext, true
}

// HasMbsSecurityContext returns a boolean if a field has been set.
func (o *ContextStatusEventReport) HasMbsSecurityContext() bool {
	if o != nil && !isNil(o.MbsSecurityContext) {
		return true
	}

	return false
}

// SetMbsSecurityContext gets a reference to the given MbsSecurityContext and assigns it to the MbsSecurityContext field.
func (o *ContextStatusEventReport) SetMbsSecurityContext(v MbsSecurityContext) {
	o.MbsSecurityContext = &v
}

func (o ContextStatusEventReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	if !isNil(o.QosInfo) {
		toSerialize["qosInfo"] = o.QosInfo
	}
	if !isNil(o.StatusInfo) {
		toSerialize["statusInfo"] = o.StatusInfo
	}
	if !isNil(o.MbsServiceArea) {
		toSerialize["mbsServiceArea"] = o.MbsServiceArea
	}
	if !isNil(o.MbsServiceAreaInfoList) {
		toSerialize["mbsServiceAreaInfoList"] = o.MbsServiceAreaInfoList
	}
	if !isNil(o.MulticastTransAddInfo) {
		toSerialize["multicastTransAddInfo"] = o.MulticastTransAddInfo
	}
	if !isNil(o.MbsSecurityContext) {
		toSerialize["mbsSecurityContext"] = o.MbsSecurityContext
	}
	return json.Marshal(toSerialize)
}

type NullableContextStatusEventReport struct {
	value *ContextStatusEventReport
	isSet bool
}

func (v NullableContextStatusEventReport) Get() *ContextStatusEventReport {
	return v.value
}

func (v *NullableContextStatusEventReport) Set(val *ContextStatusEventReport) {
	v.value = val
	v.isSet = true
}

func (v NullableContextStatusEventReport) IsSet() bool {
	return v.isSet
}

func (v *NullableContextStatusEventReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextStatusEventReport(val *ContextStatusEventReport) *NullableContextStatusEventReport {
	return &NullableContextStatusEventReport{value: val, isSet: true}
}

func (v NullableContextStatusEventReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextStatusEventReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


