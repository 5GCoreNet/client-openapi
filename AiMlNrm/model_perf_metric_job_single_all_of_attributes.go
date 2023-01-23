/*
AI/ML NRM

OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AiMlNrm

import (
	"encoding/json"
)

// PerfMetricJobSingleAllOfAttributes struct for PerfMetricJobSingleAllOfAttributes
type PerfMetricJobSingleAllOfAttributes struct {
	AdministrativeState *AdministrativeState `json:"administrativeState,omitempty"`
	OperationalState *OperationalState `json:"operationalState,omitempty"`
	JobId *string `json:"jobId,omitempty"`
	PerformanceMetrics []string `json:"performanceMetrics,omitempty"`
	GranularityPeriod *int32 `json:"granularityPeriod,omitempty"`
	ObjectInstances []string `json:"objectInstances,omitempty"`
	RootObjectInstances []string `json:"rootObjectInstances,omitempty"`
	ReportingCtrl *ReportingCtrl `json:"reportingCtrl,omitempty"`
}

// NewPerfMetricJobSingleAllOfAttributes instantiates a new PerfMetricJobSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerfMetricJobSingleAllOfAttributes() *PerfMetricJobSingleAllOfAttributes {
	this := PerfMetricJobSingleAllOfAttributes{}
	return &this
}

// NewPerfMetricJobSingleAllOfAttributesWithDefaults instantiates a new PerfMetricJobSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerfMetricJobSingleAllOfAttributesWithDefaults() *PerfMetricJobSingleAllOfAttributes {
	this := PerfMetricJobSingleAllOfAttributes{}
	return &this
}

// GetAdministrativeState returns the AdministrativeState field value if set, zero value otherwise.
func (o *PerfMetricJobSingleAllOfAttributes) GetAdministrativeState() AdministrativeState {
	if o == nil || isNil(o.AdministrativeState) {
		var ret AdministrativeState
		return ret
	}
	return *o.AdministrativeState
}

// GetAdministrativeStateOk returns a tuple with the AdministrativeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfMetricJobSingleAllOfAttributes) GetAdministrativeStateOk() (*AdministrativeState, bool) {
	if o == nil || isNil(o.AdministrativeState) {
    return nil, false
	}
	return o.AdministrativeState, true
}

// HasAdministrativeState returns a boolean if a field has been set.
func (o *PerfMetricJobSingleAllOfAttributes) HasAdministrativeState() bool {
	if o != nil && !isNil(o.AdministrativeState) {
		return true
	}

	return false
}

// SetAdministrativeState gets a reference to the given AdministrativeState and assigns it to the AdministrativeState field.
func (o *PerfMetricJobSingleAllOfAttributes) SetAdministrativeState(v AdministrativeState) {
	o.AdministrativeState = &v
}

// GetOperationalState returns the OperationalState field value if set, zero value otherwise.
func (o *PerfMetricJobSingleAllOfAttributes) GetOperationalState() OperationalState {
	if o == nil || isNil(o.OperationalState) {
		var ret OperationalState
		return ret
	}
	return *o.OperationalState
}

// GetOperationalStateOk returns a tuple with the OperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfMetricJobSingleAllOfAttributes) GetOperationalStateOk() (*OperationalState, bool) {
	if o == nil || isNil(o.OperationalState) {
    return nil, false
	}
	return o.OperationalState, true
}

// HasOperationalState returns a boolean if a field has been set.
func (o *PerfMetricJobSingleAllOfAttributes) HasOperationalState() bool {
	if o != nil && !isNil(o.OperationalState) {
		return true
	}

	return false
}

// SetOperationalState gets a reference to the given OperationalState and assigns it to the OperationalState field.
func (o *PerfMetricJobSingleAllOfAttributes) SetOperationalState(v OperationalState) {
	o.OperationalState = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *PerfMetricJobSingleAllOfAttributes) GetJobId() string {
	if o == nil || isNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfMetricJobSingleAllOfAttributes) GetJobIdOk() (*string, bool) {
	if o == nil || isNil(o.JobId) {
    return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *PerfMetricJobSingleAllOfAttributes) HasJobId() bool {
	if o != nil && !isNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *PerfMetricJobSingleAllOfAttributes) SetJobId(v string) {
	o.JobId = &v
}

// GetPerformanceMetrics returns the PerformanceMetrics field value if set, zero value otherwise.
func (o *PerfMetricJobSingleAllOfAttributes) GetPerformanceMetrics() []string {
	if o == nil || isNil(o.PerformanceMetrics) {
		var ret []string
		return ret
	}
	return o.PerformanceMetrics
}

// GetPerformanceMetricsOk returns a tuple with the PerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfMetricJobSingleAllOfAttributes) GetPerformanceMetricsOk() ([]string, bool) {
	if o == nil || isNil(o.PerformanceMetrics) {
    return nil, false
	}
	return o.PerformanceMetrics, true
}

// HasPerformanceMetrics returns a boolean if a field has been set.
func (o *PerfMetricJobSingleAllOfAttributes) HasPerformanceMetrics() bool {
	if o != nil && !isNil(o.PerformanceMetrics) {
		return true
	}

	return false
}

// SetPerformanceMetrics gets a reference to the given []string and assigns it to the PerformanceMetrics field.
func (o *PerfMetricJobSingleAllOfAttributes) SetPerformanceMetrics(v []string) {
	o.PerformanceMetrics = v
}

// GetGranularityPeriod returns the GranularityPeriod field value if set, zero value otherwise.
func (o *PerfMetricJobSingleAllOfAttributes) GetGranularityPeriod() int32 {
	if o == nil || isNil(o.GranularityPeriod) {
		var ret int32
		return ret
	}
	return *o.GranularityPeriod
}

// GetGranularityPeriodOk returns a tuple with the GranularityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfMetricJobSingleAllOfAttributes) GetGranularityPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.GranularityPeriod) {
    return nil, false
	}
	return o.GranularityPeriod, true
}

// HasGranularityPeriod returns a boolean if a field has been set.
func (o *PerfMetricJobSingleAllOfAttributes) HasGranularityPeriod() bool {
	if o != nil && !isNil(o.GranularityPeriod) {
		return true
	}

	return false
}

// SetGranularityPeriod gets a reference to the given int32 and assigns it to the GranularityPeriod field.
func (o *PerfMetricJobSingleAllOfAttributes) SetGranularityPeriod(v int32) {
	o.GranularityPeriod = &v
}

// GetObjectInstances returns the ObjectInstances field value if set, zero value otherwise.
func (o *PerfMetricJobSingleAllOfAttributes) GetObjectInstances() []string {
	if o == nil || isNil(o.ObjectInstances) {
		var ret []string
		return ret
	}
	return o.ObjectInstances
}

// GetObjectInstancesOk returns a tuple with the ObjectInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfMetricJobSingleAllOfAttributes) GetObjectInstancesOk() ([]string, bool) {
	if o == nil || isNil(o.ObjectInstances) {
    return nil, false
	}
	return o.ObjectInstances, true
}

// HasObjectInstances returns a boolean if a field has been set.
func (o *PerfMetricJobSingleAllOfAttributes) HasObjectInstances() bool {
	if o != nil && !isNil(o.ObjectInstances) {
		return true
	}

	return false
}

// SetObjectInstances gets a reference to the given []string and assigns it to the ObjectInstances field.
func (o *PerfMetricJobSingleAllOfAttributes) SetObjectInstances(v []string) {
	o.ObjectInstances = v
}

// GetRootObjectInstances returns the RootObjectInstances field value if set, zero value otherwise.
func (o *PerfMetricJobSingleAllOfAttributes) GetRootObjectInstances() []string {
	if o == nil || isNil(o.RootObjectInstances) {
		var ret []string
		return ret
	}
	return o.RootObjectInstances
}

// GetRootObjectInstancesOk returns a tuple with the RootObjectInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfMetricJobSingleAllOfAttributes) GetRootObjectInstancesOk() ([]string, bool) {
	if o == nil || isNil(o.RootObjectInstances) {
    return nil, false
	}
	return o.RootObjectInstances, true
}

// HasRootObjectInstances returns a boolean if a field has been set.
func (o *PerfMetricJobSingleAllOfAttributes) HasRootObjectInstances() bool {
	if o != nil && !isNil(o.RootObjectInstances) {
		return true
	}

	return false
}

// SetRootObjectInstances gets a reference to the given []string and assigns it to the RootObjectInstances field.
func (o *PerfMetricJobSingleAllOfAttributes) SetRootObjectInstances(v []string) {
	o.RootObjectInstances = v
}

// GetReportingCtrl returns the ReportingCtrl field value if set, zero value otherwise.
func (o *PerfMetricJobSingleAllOfAttributes) GetReportingCtrl() ReportingCtrl {
	if o == nil || isNil(o.ReportingCtrl) {
		var ret ReportingCtrl
		return ret
	}
	return *o.ReportingCtrl
}

// GetReportingCtrlOk returns a tuple with the ReportingCtrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfMetricJobSingleAllOfAttributes) GetReportingCtrlOk() (*ReportingCtrl, bool) {
	if o == nil || isNil(o.ReportingCtrl) {
    return nil, false
	}
	return o.ReportingCtrl, true
}

// HasReportingCtrl returns a boolean if a field has been set.
func (o *PerfMetricJobSingleAllOfAttributes) HasReportingCtrl() bool {
	if o != nil && !isNil(o.ReportingCtrl) {
		return true
	}

	return false
}

// SetReportingCtrl gets a reference to the given ReportingCtrl and assigns it to the ReportingCtrl field.
func (o *PerfMetricJobSingleAllOfAttributes) SetReportingCtrl(v ReportingCtrl) {
	o.ReportingCtrl = &v
}

func (o PerfMetricJobSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AdministrativeState) {
		toSerialize["administrativeState"] = o.AdministrativeState
	}
	if !isNil(o.OperationalState) {
		toSerialize["operationalState"] = o.OperationalState
	}
	if !isNil(o.JobId) {
		toSerialize["jobId"] = o.JobId
	}
	if !isNil(o.PerformanceMetrics) {
		toSerialize["performanceMetrics"] = o.PerformanceMetrics
	}
	if !isNil(o.GranularityPeriod) {
		toSerialize["granularityPeriod"] = o.GranularityPeriod
	}
	if !isNil(o.ObjectInstances) {
		toSerialize["objectInstances"] = o.ObjectInstances
	}
	if !isNil(o.RootObjectInstances) {
		toSerialize["rootObjectInstances"] = o.RootObjectInstances
	}
	if !isNil(o.ReportingCtrl) {
		toSerialize["reportingCtrl"] = o.ReportingCtrl
	}
	return json.Marshal(toSerialize)
}

type NullablePerfMetricJobSingleAllOfAttributes struct {
	value *PerfMetricJobSingleAllOfAttributes
	isSet bool
}

func (v NullablePerfMetricJobSingleAllOfAttributes) Get() *PerfMetricJobSingleAllOfAttributes {
	return v.value
}

func (v *NullablePerfMetricJobSingleAllOfAttributes) Set(val *PerfMetricJobSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePerfMetricJobSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePerfMetricJobSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerfMetricJobSingleAllOfAttributes(val *PerfMetricJobSingleAllOfAttributes) *NullablePerfMetricJobSingleAllOfAttributes {
	return &NullablePerfMetricJobSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullablePerfMetricJobSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerfMetricJobSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


