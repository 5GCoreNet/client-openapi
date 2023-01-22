/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package StreamingDataMnS

import (
	"encoding/json"
)

// PerformanceInfoType Information specific to performance data reporting
type PerformanceInfoType struct {
	MeasObjDn MeasObjDnType `json:"measObjDn"`
	// an ordered list of performance metric names (see clause 4.4.1 of 3GPP TS 28.622[11]) whose values are to be reported by the Performance Data Stream Units (see Annex C of TS 28.550 [42]) via this stream. Performance metrics include measurement and KPI
	PerformanceMetrics []string `json:"performanceMetrics"`
	JobId *string `json:"jobId,omitempty"`
}

// NewPerformanceInfoType instantiates a new PerformanceInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceInfoType(measObjDn MeasObjDnType, performanceMetrics []string) *PerformanceInfoType {
	this := PerformanceInfoType{}
	this.MeasObjDn = measObjDn
	this.PerformanceMetrics = performanceMetrics
	return &this
}

// NewPerformanceInfoTypeWithDefaults instantiates a new PerformanceInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceInfoTypeWithDefaults() *PerformanceInfoType {
	this := PerformanceInfoType{}
	return &this
}

// GetMeasObjDn returns the MeasObjDn field value
func (o *PerformanceInfoType) GetMeasObjDn() MeasObjDnType {
	if o == nil {
		var ret MeasObjDnType
		return ret
	}

	return o.MeasObjDn
}

// GetMeasObjDnOk returns a tuple with the MeasObjDn field value
// and a boolean to check if the value has been set.
func (o *PerformanceInfoType) GetMeasObjDnOk() (*MeasObjDnType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MeasObjDn, true
}

// SetMeasObjDn sets field value
func (o *PerformanceInfoType) SetMeasObjDn(v MeasObjDnType) {
	o.MeasObjDn = v
}

// GetPerformanceMetrics returns the PerformanceMetrics field value
func (o *PerformanceInfoType) GetPerformanceMetrics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PerformanceMetrics
}

// GetPerformanceMetricsOk returns a tuple with the PerformanceMetrics field value
// and a boolean to check if the value has been set.
func (o *PerformanceInfoType) GetPerformanceMetricsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.PerformanceMetrics, true
}

// SetPerformanceMetrics sets field value
func (o *PerformanceInfoType) SetPerformanceMetrics(v []string) {
	o.PerformanceMetrics = v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *PerformanceInfoType) GetJobId() string {
	if o == nil || isNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceInfoType) GetJobIdOk() (*string, bool) {
	if o == nil || isNil(o.JobId) {
    return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *PerformanceInfoType) HasJobId() bool {
	if o != nil && !isNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *PerformanceInfoType) SetJobId(v string) {
	o.JobId = &v
}

func (o PerformanceInfoType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["measObjDn"] = o.MeasObjDn
	}
	if true {
		toSerialize["performanceMetrics"] = o.PerformanceMetrics
	}
	if !isNil(o.JobId) {
		toSerialize["jobId"] = o.JobId
	}
	return json.Marshal(toSerialize)
}

type NullablePerformanceInfoType struct {
	value *PerformanceInfoType
	isSet bool
}

func (v NullablePerformanceInfoType) Get() *PerformanceInfoType {
	return v.value
}

func (v *NullablePerformanceInfoType) Set(val *PerformanceInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceInfoType(val *PerformanceInfoType) *NullablePerformanceInfoType {
	return &NullablePerformanceInfoType{value: val, isSet: true}
}

func (v NullablePerformanceInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


