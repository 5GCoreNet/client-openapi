/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AnalyticsExposure

import (
	"encoding/json"
	"time"
)

// AnalyticsEventNotif Represents an analytics event to be reported.
type AnalyticsEventNotif struct {
	AnalyEvent AnalyticsEvent `json:"analyEvent"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
	FailNotifyCode *NwdafFailureCode `json:"failNotifyCode,omitempty"`
	// indicating a time in seconds.
	RvWaitTime *int32 `json:"rvWaitTime,omitempty"`
	UeMobilityInfos []UeMobilityExposure `json:"ueMobilityInfos,omitempty"`
	UeCommInfos []UeCommunication `json:"ueCommInfos,omitempty"`
	AbnormalInfos []AbnormalExposure `json:"abnormalInfos,omitempty"`
	CongestInfos []CongestInfo `json:"congestInfos,omitempty"`
	NwPerfInfos []NetworkPerfExposure `json:"nwPerfInfos,omitempty"`
	QosSustainInfos []QosSustainabilityExposure `json:"qosSustainInfos,omitempty"`
	DisperInfos []DispersionInfo `json:"disperInfos,omitempty"`
	DnPerfInfos []DnPerfInfo `json:"dnPerfInfos,omitempty"`
	SvcExps []ServiceExperienceInfo `json:"svcExps,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Start *time.Time `json:"start,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStampGen *time.Time `json:"timeStampGen,omitempty"`
}

// NewAnalyticsEventNotif instantiates a new AnalyticsEventNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsEventNotif(analyEvent AnalyticsEvent, timeStamp time.Time) *AnalyticsEventNotif {
	this := AnalyticsEventNotif{}
	this.AnalyEvent = analyEvent
	this.TimeStamp = timeStamp
	return &this
}

// NewAnalyticsEventNotifWithDefaults instantiates a new AnalyticsEventNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsEventNotifWithDefaults() *AnalyticsEventNotif {
	this := AnalyticsEventNotif{}
	return &this
}

// GetAnalyEvent returns the AnalyEvent field value
func (o *AnalyticsEventNotif) GetAnalyEvent() AnalyticsEvent {
	if o == nil {
		var ret AnalyticsEvent
		return ret
	}

	return o.AnalyEvent
}

// GetAnalyEventOk returns a tuple with the AnalyEvent field value
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetAnalyEventOk() (*AnalyticsEvent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AnalyEvent, true
}

// SetAnalyEvent sets field value
func (o *AnalyticsEventNotif) SetAnalyEvent(v AnalyticsEvent) {
	o.AnalyEvent = v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetExpiry() time.Time {
	if o == nil || isNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetExpiryOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expiry) {
    return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasExpiry() bool {
	if o != nil && !isNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *AnalyticsEventNotif) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetTimeStamp returns the TimeStamp field value
func (o *AnalyticsEventNotif) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *AnalyticsEventNotif) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

// GetFailNotifyCode returns the FailNotifyCode field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetFailNotifyCode() NwdafFailureCode {
	if o == nil || isNil(o.FailNotifyCode) {
		var ret NwdafFailureCode
		return ret
	}
	return *o.FailNotifyCode
}

// GetFailNotifyCodeOk returns a tuple with the FailNotifyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetFailNotifyCodeOk() (*NwdafFailureCode, bool) {
	if o == nil || isNil(o.FailNotifyCode) {
    return nil, false
	}
	return o.FailNotifyCode, true
}

// HasFailNotifyCode returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasFailNotifyCode() bool {
	if o != nil && !isNil(o.FailNotifyCode) {
		return true
	}

	return false
}

// SetFailNotifyCode gets a reference to the given NwdafFailureCode and assigns it to the FailNotifyCode field.
func (o *AnalyticsEventNotif) SetFailNotifyCode(v NwdafFailureCode) {
	o.FailNotifyCode = &v
}

// GetRvWaitTime returns the RvWaitTime field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetRvWaitTime() int32 {
	if o == nil || isNil(o.RvWaitTime) {
		var ret int32
		return ret
	}
	return *o.RvWaitTime
}

// GetRvWaitTimeOk returns a tuple with the RvWaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetRvWaitTimeOk() (*int32, bool) {
	if o == nil || isNil(o.RvWaitTime) {
    return nil, false
	}
	return o.RvWaitTime, true
}

// HasRvWaitTime returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasRvWaitTime() bool {
	if o != nil && !isNil(o.RvWaitTime) {
		return true
	}

	return false
}

// SetRvWaitTime gets a reference to the given int32 and assigns it to the RvWaitTime field.
func (o *AnalyticsEventNotif) SetRvWaitTime(v int32) {
	o.RvWaitTime = &v
}

// GetUeMobilityInfos returns the UeMobilityInfos field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetUeMobilityInfos() []UeMobilityExposure {
	if o == nil || isNil(o.UeMobilityInfos) {
		var ret []UeMobilityExposure
		return ret
	}
	return o.UeMobilityInfos
}

// GetUeMobilityInfosOk returns a tuple with the UeMobilityInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetUeMobilityInfosOk() ([]UeMobilityExposure, bool) {
	if o == nil || isNil(o.UeMobilityInfos) {
    return nil, false
	}
	return o.UeMobilityInfos, true
}

// HasUeMobilityInfos returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasUeMobilityInfos() bool {
	if o != nil && !isNil(o.UeMobilityInfos) {
		return true
	}

	return false
}

// SetUeMobilityInfos gets a reference to the given []UeMobilityExposure and assigns it to the UeMobilityInfos field.
func (o *AnalyticsEventNotif) SetUeMobilityInfos(v []UeMobilityExposure) {
	o.UeMobilityInfos = v
}

// GetUeCommInfos returns the UeCommInfos field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetUeCommInfos() []UeCommunication {
	if o == nil || isNil(o.UeCommInfos) {
		var ret []UeCommunication
		return ret
	}
	return o.UeCommInfos
}

// GetUeCommInfosOk returns a tuple with the UeCommInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetUeCommInfosOk() ([]UeCommunication, bool) {
	if o == nil || isNil(o.UeCommInfos) {
    return nil, false
	}
	return o.UeCommInfos, true
}

// HasUeCommInfos returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasUeCommInfos() bool {
	if o != nil && !isNil(o.UeCommInfos) {
		return true
	}

	return false
}

// SetUeCommInfos gets a reference to the given []UeCommunication and assigns it to the UeCommInfos field.
func (o *AnalyticsEventNotif) SetUeCommInfos(v []UeCommunication) {
	o.UeCommInfos = v
}

// GetAbnormalInfos returns the AbnormalInfos field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetAbnormalInfos() []AbnormalExposure {
	if o == nil || isNil(o.AbnormalInfos) {
		var ret []AbnormalExposure
		return ret
	}
	return o.AbnormalInfos
}

// GetAbnormalInfosOk returns a tuple with the AbnormalInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetAbnormalInfosOk() ([]AbnormalExposure, bool) {
	if o == nil || isNil(o.AbnormalInfos) {
    return nil, false
	}
	return o.AbnormalInfos, true
}

// HasAbnormalInfos returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasAbnormalInfos() bool {
	if o != nil && !isNil(o.AbnormalInfos) {
		return true
	}

	return false
}

// SetAbnormalInfos gets a reference to the given []AbnormalExposure and assigns it to the AbnormalInfos field.
func (o *AnalyticsEventNotif) SetAbnormalInfos(v []AbnormalExposure) {
	o.AbnormalInfos = v
}

// GetCongestInfos returns the CongestInfos field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetCongestInfos() []CongestInfo {
	if o == nil || isNil(o.CongestInfos) {
		var ret []CongestInfo
		return ret
	}
	return o.CongestInfos
}

// GetCongestInfosOk returns a tuple with the CongestInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetCongestInfosOk() ([]CongestInfo, bool) {
	if o == nil || isNil(o.CongestInfos) {
    return nil, false
	}
	return o.CongestInfos, true
}

// HasCongestInfos returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasCongestInfos() bool {
	if o != nil && !isNil(o.CongestInfos) {
		return true
	}

	return false
}

// SetCongestInfos gets a reference to the given []CongestInfo and assigns it to the CongestInfos field.
func (o *AnalyticsEventNotif) SetCongestInfos(v []CongestInfo) {
	o.CongestInfos = v
}

// GetNwPerfInfos returns the NwPerfInfos field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetNwPerfInfos() []NetworkPerfExposure {
	if o == nil || isNil(o.NwPerfInfos) {
		var ret []NetworkPerfExposure
		return ret
	}
	return o.NwPerfInfos
}

// GetNwPerfInfosOk returns a tuple with the NwPerfInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetNwPerfInfosOk() ([]NetworkPerfExposure, bool) {
	if o == nil || isNil(o.NwPerfInfos) {
    return nil, false
	}
	return o.NwPerfInfos, true
}

// HasNwPerfInfos returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasNwPerfInfos() bool {
	if o != nil && !isNil(o.NwPerfInfos) {
		return true
	}

	return false
}

// SetNwPerfInfos gets a reference to the given []NetworkPerfExposure and assigns it to the NwPerfInfos field.
func (o *AnalyticsEventNotif) SetNwPerfInfos(v []NetworkPerfExposure) {
	o.NwPerfInfos = v
}

// GetQosSustainInfos returns the QosSustainInfos field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetQosSustainInfos() []QosSustainabilityExposure {
	if o == nil || isNil(o.QosSustainInfos) {
		var ret []QosSustainabilityExposure
		return ret
	}
	return o.QosSustainInfos
}

// GetQosSustainInfosOk returns a tuple with the QosSustainInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetQosSustainInfosOk() ([]QosSustainabilityExposure, bool) {
	if o == nil || isNil(o.QosSustainInfos) {
    return nil, false
	}
	return o.QosSustainInfos, true
}

// HasQosSustainInfos returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasQosSustainInfos() bool {
	if o != nil && !isNil(o.QosSustainInfos) {
		return true
	}

	return false
}

// SetQosSustainInfos gets a reference to the given []QosSustainabilityExposure and assigns it to the QosSustainInfos field.
func (o *AnalyticsEventNotif) SetQosSustainInfos(v []QosSustainabilityExposure) {
	o.QosSustainInfos = v
}

// GetDisperInfos returns the DisperInfos field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetDisperInfos() []DispersionInfo {
	if o == nil || isNil(o.DisperInfos) {
		var ret []DispersionInfo
		return ret
	}
	return o.DisperInfos
}

// GetDisperInfosOk returns a tuple with the DisperInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetDisperInfosOk() ([]DispersionInfo, bool) {
	if o == nil || isNil(o.DisperInfos) {
    return nil, false
	}
	return o.DisperInfos, true
}

// HasDisperInfos returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasDisperInfos() bool {
	if o != nil && !isNil(o.DisperInfos) {
		return true
	}

	return false
}

// SetDisperInfos gets a reference to the given []DispersionInfo and assigns it to the DisperInfos field.
func (o *AnalyticsEventNotif) SetDisperInfos(v []DispersionInfo) {
	o.DisperInfos = v
}

// GetDnPerfInfos returns the DnPerfInfos field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetDnPerfInfos() []DnPerfInfo {
	if o == nil || isNil(o.DnPerfInfos) {
		var ret []DnPerfInfo
		return ret
	}
	return o.DnPerfInfos
}

// GetDnPerfInfosOk returns a tuple with the DnPerfInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetDnPerfInfosOk() ([]DnPerfInfo, bool) {
	if o == nil || isNil(o.DnPerfInfos) {
    return nil, false
	}
	return o.DnPerfInfos, true
}

// HasDnPerfInfos returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasDnPerfInfos() bool {
	if o != nil && !isNil(o.DnPerfInfos) {
		return true
	}

	return false
}

// SetDnPerfInfos gets a reference to the given []DnPerfInfo and assigns it to the DnPerfInfos field.
func (o *AnalyticsEventNotif) SetDnPerfInfos(v []DnPerfInfo) {
	o.DnPerfInfos = v
}

// GetSvcExps returns the SvcExps field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetSvcExps() []ServiceExperienceInfo {
	if o == nil || isNil(o.SvcExps) {
		var ret []ServiceExperienceInfo
		return ret
	}
	return o.SvcExps
}

// GetSvcExpsOk returns a tuple with the SvcExps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetSvcExpsOk() ([]ServiceExperienceInfo, bool) {
	if o == nil || isNil(o.SvcExps) {
    return nil, false
	}
	return o.SvcExps, true
}

// HasSvcExps returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasSvcExps() bool {
	if o != nil && !isNil(o.SvcExps) {
		return true
	}

	return false
}

// SetSvcExps gets a reference to the given []ServiceExperienceInfo and assigns it to the SvcExps field.
func (o *AnalyticsEventNotif) SetSvcExps(v []ServiceExperienceInfo) {
	o.SvcExps = v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetStart() time.Time {
	if o == nil || isNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetStartOk() (*time.Time, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *AnalyticsEventNotif) SetStart(v time.Time) {
	o.Start = &v
}

// GetTimeStampGen returns the TimeStampGen field value if set, zero value otherwise.
func (o *AnalyticsEventNotif) GetTimeStampGen() time.Time {
	if o == nil || isNil(o.TimeStampGen) {
		var ret time.Time
		return ret
	}
	return *o.TimeStampGen
}

// GetTimeStampGenOk returns a tuple with the TimeStampGen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsEventNotif) GetTimeStampGenOk() (*time.Time, bool) {
	if o == nil || isNil(o.TimeStampGen) {
    return nil, false
	}
	return o.TimeStampGen, true
}

// HasTimeStampGen returns a boolean if a field has been set.
func (o *AnalyticsEventNotif) HasTimeStampGen() bool {
	if o != nil && !isNil(o.TimeStampGen) {
		return true
	}

	return false
}

// SetTimeStampGen gets a reference to the given time.Time and assigns it to the TimeStampGen field.
func (o *AnalyticsEventNotif) SetTimeStampGen(v time.Time) {
	o.TimeStampGen = &v
}

func (o AnalyticsEventNotif) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["analyEvent"] = o.AnalyEvent
	}
	if !isNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if true {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	if !isNil(o.FailNotifyCode) {
		toSerialize["failNotifyCode"] = o.FailNotifyCode
	}
	if !isNil(o.RvWaitTime) {
		toSerialize["rvWaitTime"] = o.RvWaitTime
	}
	if !isNil(o.UeMobilityInfos) {
		toSerialize["ueMobilityInfos"] = o.UeMobilityInfos
	}
	if !isNil(o.UeCommInfos) {
		toSerialize["ueCommInfos"] = o.UeCommInfos
	}
	if !isNil(o.AbnormalInfos) {
		toSerialize["abnormalInfos"] = o.AbnormalInfos
	}
	if !isNil(o.CongestInfos) {
		toSerialize["congestInfos"] = o.CongestInfos
	}
	if !isNil(o.NwPerfInfos) {
		toSerialize["nwPerfInfos"] = o.NwPerfInfos
	}
	if !isNil(o.QosSustainInfos) {
		toSerialize["qosSustainInfos"] = o.QosSustainInfos
	}
	if !isNil(o.DisperInfos) {
		toSerialize["disperInfos"] = o.DisperInfos
	}
	if !isNil(o.DnPerfInfos) {
		toSerialize["dnPerfInfos"] = o.DnPerfInfos
	}
	if !isNil(o.SvcExps) {
		toSerialize["svcExps"] = o.SvcExps
	}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.TimeStampGen) {
		toSerialize["timeStampGen"] = o.TimeStampGen
	}
	return json.Marshal(toSerialize)
}

type NullableAnalyticsEventNotif struct {
	value *AnalyticsEventNotif
	isSet bool
}

func (v NullableAnalyticsEventNotif) Get() *AnalyticsEventNotif {
	return v.value
}

func (v *NullableAnalyticsEventNotif) Set(val *AnalyticsEventNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsEventNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsEventNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsEventNotif(val *AnalyticsEventNotif) *NullableAnalyticsEventNotif {
	return &NullableAnalyticsEventNotif{value: val, isSet: true}
}

func (v NullableAnalyticsEventNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsEventNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


