/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"time"
)

// EventReportingRequirement Represents the type of reporting that the subscription requires.
type EventReportingRequirement struct {
	Accuracy *Accuracy `json:"accuracy,omitempty"`
	// Each element indicates the preferred accuracy level per analytics subset. It may be present if the \"listOfAnaSubsets\" attribute is present in the subscription request when the subscription event is NF_LOAD, UE_COMM, DISPERSION, NETWORK_PERFORMANCE, WLAN_PERFORMANCE, DN_PERFORMANCE or SERVICE_EXPERIENCE. 
	AccPerSubset []Accuracy `json:"accPerSubset,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	StartTs *time.Time `json:"startTs,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	EndTs *time.Time `json:"endTs,omitempty"`
	// Offset period in units of seconds to the reporting time, if the value is negative means  statistics in the past offset period, otherwise a positive value means prediction in the  future offset period. May be present if the \"repPeriod\" attribute is included within the  \"evtReq\" attribute. 
	OffsetPeriod *int32 `json:"offsetPeriod,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	SampRatio *int32 `json:"sampRatio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxObjectNbr *int32 `json:"maxObjectNbr,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxSupiNbr *int32 `json:"maxSupiNbr,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeAnaNeeded *time.Time `json:"timeAnaNeeded,omitempty"`
	AnaMeta []AnalyticsMetadata `json:"anaMeta,omitempty"`
	AnaMetaInd *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
}

// NewEventReportingRequirement instantiates a new EventReportingRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventReportingRequirement() *EventReportingRequirement {
	this := EventReportingRequirement{}
	return &this
}

// NewEventReportingRequirementWithDefaults instantiates a new EventReportingRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventReportingRequirementWithDefaults() *EventReportingRequirement {
	this := EventReportingRequirement{}
	return &this
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *EventReportingRequirement) GetAccuracy() Accuracy {
	if o == nil || isNil(o.Accuracy) {
		var ret Accuracy
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingRequirement) GetAccuracyOk() (*Accuracy, bool) {
	if o == nil || isNil(o.Accuracy) {
    return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *EventReportingRequirement) HasAccuracy() bool {
	if o != nil && !isNil(o.Accuracy) {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given Accuracy and assigns it to the Accuracy field.
func (o *EventReportingRequirement) SetAccuracy(v Accuracy) {
	o.Accuracy = &v
}

// GetAccPerSubset returns the AccPerSubset field value if set, zero value otherwise.
func (o *EventReportingRequirement) GetAccPerSubset() []Accuracy {
	if o == nil || isNil(o.AccPerSubset) {
		var ret []Accuracy
		return ret
	}
	return o.AccPerSubset
}

// GetAccPerSubsetOk returns a tuple with the AccPerSubset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingRequirement) GetAccPerSubsetOk() ([]Accuracy, bool) {
	if o == nil || isNil(o.AccPerSubset) {
    return nil, false
	}
	return o.AccPerSubset, true
}

// HasAccPerSubset returns a boolean if a field has been set.
func (o *EventReportingRequirement) HasAccPerSubset() bool {
	if o != nil && !isNil(o.AccPerSubset) {
		return true
	}

	return false
}

// SetAccPerSubset gets a reference to the given []Accuracy and assigns it to the AccPerSubset field.
func (o *EventReportingRequirement) SetAccPerSubset(v []Accuracy) {
	o.AccPerSubset = v
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *EventReportingRequirement) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingRequirement) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *EventReportingRequirement) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *EventReportingRequirement) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *EventReportingRequirement) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingRequirement) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *EventReportingRequirement) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *EventReportingRequirement) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetOffsetPeriod returns the OffsetPeriod field value if set, zero value otherwise.
func (o *EventReportingRequirement) GetOffsetPeriod() int32 {
	if o == nil || isNil(o.OffsetPeriod) {
		var ret int32
		return ret
	}
	return *o.OffsetPeriod
}

// GetOffsetPeriodOk returns a tuple with the OffsetPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingRequirement) GetOffsetPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.OffsetPeriod) {
    return nil, false
	}
	return o.OffsetPeriod, true
}

// HasOffsetPeriod returns a boolean if a field has been set.
func (o *EventReportingRequirement) HasOffsetPeriod() bool {
	if o != nil && !isNil(o.OffsetPeriod) {
		return true
	}

	return false
}

// SetOffsetPeriod gets a reference to the given int32 and assigns it to the OffsetPeriod field.
func (o *EventReportingRequirement) SetOffsetPeriod(v int32) {
	o.OffsetPeriod = &v
}

// GetSampRatio returns the SampRatio field value if set, zero value otherwise.
func (o *EventReportingRequirement) GetSampRatio() int32 {
	if o == nil || isNil(o.SampRatio) {
		var ret int32
		return ret
	}
	return *o.SampRatio
}

// GetSampRatioOk returns a tuple with the SampRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingRequirement) GetSampRatioOk() (*int32, bool) {
	if o == nil || isNil(o.SampRatio) {
    return nil, false
	}
	return o.SampRatio, true
}

// HasSampRatio returns a boolean if a field has been set.
func (o *EventReportingRequirement) HasSampRatio() bool {
	if o != nil && !isNil(o.SampRatio) {
		return true
	}

	return false
}

// SetSampRatio gets a reference to the given int32 and assigns it to the SampRatio field.
func (o *EventReportingRequirement) SetSampRatio(v int32) {
	o.SampRatio = &v
}

// GetMaxObjectNbr returns the MaxObjectNbr field value if set, zero value otherwise.
func (o *EventReportingRequirement) GetMaxObjectNbr() int32 {
	if o == nil || isNil(o.MaxObjectNbr) {
		var ret int32
		return ret
	}
	return *o.MaxObjectNbr
}

// GetMaxObjectNbrOk returns a tuple with the MaxObjectNbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingRequirement) GetMaxObjectNbrOk() (*int32, bool) {
	if o == nil || isNil(o.MaxObjectNbr) {
    return nil, false
	}
	return o.MaxObjectNbr, true
}

// HasMaxObjectNbr returns a boolean if a field has been set.
func (o *EventReportingRequirement) HasMaxObjectNbr() bool {
	if o != nil && !isNil(o.MaxObjectNbr) {
		return true
	}

	return false
}

// SetMaxObjectNbr gets a reference to the given int32 and assigns it to the MaxObjectNbr field.
func (o *EventReportingRequirement) SetMaxObjectNbr(v int32) {
	o.MaxObjectNbr = &v
}

// GetMaxSupiNbr returns the MaxSupiNbr field value if set, zero value otherwise.
func (o *EventReportingRequirement) GetMaxSupiNbr() int32 {
	if o == nil || isNil(o.MaxSupiNbr) {
		var ret int32
		return ret
	}
	return *o.MaxSupiNbr
}

// GetMaxSupiNbrOk returns a tuple with the MaxSupiNbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingRequirement) GetMaxSupiNbrOk() (*int32, bool) {
	if o == nil || isNil(o.MaxSupiNbr) {
    return nil, false
	}
	return o.MaxSupiNbr, true
}

// HasMaxSupiNbr returns a boolean if a field has been set.
func (o *EventReportingRequirement) HasMaxSupiNbr() bool {
	if o != nil && !isNil(o.MaxSupiNbr) {
		return true
	}

	return false
}

// SetMaxSupiNbr gets a reference to the given int32 and assigns it to the MaxSupiNbr field.
func (o *EventReportingRequirement) SetMaxSupiNbr(v int32) {
	o.MaxSupiNbr = &v
}

// GetTimeAnaNeeded returns the TimeAnaNeeded field value if set, zero value otherwise.
func (o *EventReportingRequirement) GetTimeAnaNeeded() time.Time {
	if o == nil || isNil(o.TimeAnaNeeded) {
		var ret time.Time
		return ret
	}
	return *o.TimeAnaNeeded
}

// GetTimeAnaNeededOk returns a tuple with the TimeAnaNeeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingRequirement) GetTimeAnaNeededOk() (*time.Time, bool) {
	if o == nil || isNil(o.TimeAnaNeeded) {
    return nil, false
	}
	return o.TimeAnaNeeded, true
}

// HasTimeAnaNeeded returns a boolean if a field has been set.
func (o *EventReportingRequirement) HasTimeAnaNeeded() bool {
	if o != nil && !isNil(o.TimeAnaNeeded) {
		return true
	}

	return false
}

// SetTimeAnaNeeded gets a reference to the given time.Time and assigns it to the TimeAnaNeeded field.
func (o *EventReportingRequirement) SetTimeAnaNeeded(v time.Time) {
	o.TimeAnaNeeded = &v
}

// GetAnaMeta returns the AnaMeta field value if set, zero value otherwise.
func (o *EventReportingRequirement) GetAnaMeta() []AnalyticsMetadata {
	if o == nil || isNil(o.AnaMeta) {
		var ret []AnalyticsMetadata
		return ret
	}
	return o.AnaMeta
}

// GetAnaMetaOk returns a tuple with the AnaMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingRequirement) GetAnaMetaOk() ([]AnalyticsMetadata, bool) {
	if o == nil || isNil(o.AnaMeta) {
    return nil, false
	}
	return o.AnaMeta, true
}

// HasAnaMeta returns a boolean if a field has been set.
func (o *EventReportingRequirement) HasAnaMeta() bool {
	if o != nil && !isNil(o.AnaMeta) {
		return true
	}

	return false
}

// SetAnaMeta gets a reference to the given []AnalyticsMetadata and assigns it to the AnaMeta field.
func (o *EventReportingRequirement) SetAnaMeta(v []AnalyticsMetadata) {
	o.AnaMeta = v
}

// GetAnaMetaInd returns the AnaMetaInd field value if set, zero value otherwise.
func (o *EventReportingRequirement) GetAnaMetaInd() AnalyticsMetadataIndication {
	if o == nil || isNil(o.AnaMetaInd) {
		var ret AnalyticsMetadataIndication
		return ret
	}
	return *o.AnaMetaInd
}

// GetAnaMetaIndOk returns a tuple with the AnaMetaInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingRequirement) GetAnaMetaIndOk() (*AnalyticsMetadataIndication, bool) {
	if o == nil || isNil(o.AnaMetaInd) {
    return nil, false
	}
	return o.AnaMetaInd, true
}

// HasAnaMetaInd returns a boolean if a field has been set.
func (o *EventReportingRequirement) HasAnaMetaInd() bool {
	if o != nil && !isNil(o.AnaMetaInd) {
		return true
	}

	return false
}

// SetAnaMetaInd gets a reference to the given AnalyticsMetadataIndication and assigns it to the AnaMetaInd field.
func (o *EventReportingRequirement) SetAnaMetaInd(v AnalyticsMetadataIndication) {
	o.AnaMetaInd = &v
}

func (o EventReportingRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Accuracy) {
		toSerialize["accuracy"] = o.Accuracy
	}
	if !isNil(o.AccPerSubset) {
		toSerialize["accPerSubset"] = o.AccPerSubset
	}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.OffsetPeriod) {
		toSerialize["offsetPeriod"] = o.OffsetPeriod
	}
	if !isNil(o.SampRatio) {
		toSerialize["sampRatio"] = o.SampRatio
	}
	if !isNil(o.MaxObjectNbr) {
		toSerialize["maxObjectNbr"] = o.MaxObjectNbr
	}
	if !isNil(o.MaxSupiNbr) {
		toSerialize["maxSupiNbr"] = o.MaxSupiNbr
	}
	if !isNil(o.TimeAnaNeeded) {
		toSerialize["timeAnaNeeded"] = o.TimeAnaNeeded
	}
	if !isNil(o.AnaMeta) {
		toSerialize["anaMeta"] = o.AnaMeta
	}
	if !isNil(o.AnaMetaInd) {
		toSerialize["anaMetaInd"] = o.AnaMetaInd
	}
	return json.Marshal(toSerialize)
}

type NullableEventReportingRequirement struct {
	value *EventReportingRequirement
	isSet bool
}

func (v NullableEventReportingRequirement) Get() *EventReportingRequirement {
	return v.value
}

func (v *NullableEventReportingRequirement) Set(val *EventReportingRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableEventReportingRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableEventReportingRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventReportingRequirement(val *EventReportingRequirement) *NullableEventReportingRequirement {
	return &NullableEventReportingRequirement{value: val, isSet: true}
}

func (v NullableEventReportingRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventReportingRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


