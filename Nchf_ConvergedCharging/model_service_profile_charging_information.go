/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
)

// ServiceProfileChargingInformation struct for ServiceProfileChargingInformation
type ServiceProfileChargingInformation struct {
	ServiceProfileIdentifier *string `json:"serviceProfileIdentifier,omitempty"`
	SNSSAIList []Snssai `json:"sNSSAIList,omitempty"`
	Latency *int32 `json:"latency,omitempty"`
	Availability *float32 `json:"availability,omitempty"`
	Jitter *int32 `json:"jitter,omitempty"`
	Reliability *string `json:"reliability,omitempty"`
	MaxNumberofUEs *int32 `json:"maxNumberofUEs,omitempty"`
	CoverageArea *string `json:"coverageArea,omitempty"`
	DLThptPerSlice *Throughput `json:"dLThptPerSlice,omitempty"`
	DLThptPerUE *Throughput `json:"dLThptPerUE,omitempty"`
	ULThptPerSlice *Throughput `json:"uLThptPerSlice,omitempty"`
	ULThptPerUE *Throughput `json:"uLThptPerUE,omitempty"`
	MaxNumberofPDUsessions *int32 `json:"maxNumberofPDUsessions,omitempty"`
	KPIMonitoringList *string `json:"kPIMonitoringList,omitempty"`
	SupportedAccessTechnology *int32 `json:"supportedAccessTechnology,omitempty"`
	AddServiceProfileInfo *string `json:"addServiceProfileInfo,omitempty"`
}

// NewServiceProfileChargingInformation instantiates a new ServiceProfileChargingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProfileChargingInformation() *ServiceProfileChargingInformation {
	this := ServiceProfileChargingInformation{}
	return &this
}

// NewServiceProfileChargingInformationWithDefaults instantiates a new ServiceProfileChargingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProfileChargingInformationWithDefaults() *ServiceProfileChargingInformation {
	this := ServiceProfileChargingInformation{}
	return &this
}

// GetServiceProfileIdentifier returns the ServiceProfileIdentifier field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetServiceProfileIdentifier() string {
	if o == nil || isNil(o.ServiceProfileIdentifier) {
		var ret string
		return ret
	}
	return *o.ServiceProfileIdentifier
}

// GetServiceProfileIdentifierOk returns a tuple with the ServiceProfileIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetServiceProfileIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.ServiceProfileIdentifier) {
    return nil, false
	}
	return o.ServiceProfileIdentifier, true
}

// HasServiceProfileIdentifier returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasServiceProfileIdentifier() bool {
	if o != nil && !isNil(o.ServiceProfileIdentifier) {
		return true
	}

	return false
}

// SetServiceProfileIdentifier gets a reference to the given string and assigns it to the ServiceProfileIdentifier field.
func (o *ServiceProfileChargingInformation) SetServiceProfileIdentifier(v string) {
	o.ServiceProfileIdentifier = &v
}

// GetSNSSAIList returns the SNSSAIList field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetSNSSAIList() []Snssai {
	if o == nil || isNil(o.SNSSAIList) {
		var ret []Snssai
		return ret
	}
	return o.SNSSAIList
}

// GetSNSSAIListOk returns a tuple with the SNSSAIList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetSNSSAIListOk() ([]Snssai, bool) {
	if o == nil || isNil(o.SNSSAIList) {
    return nil, false
	}
	return o.SNSSAIList, true
}

// HasSNSSAIList returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasSNSSAIList() bool {
	if o != nil && !isNil(o.SNSSAIList) {
		return true
	}

	return false
}

// SetSNSSAIList gets a reference to the given []Snssai and assigns it to the SNSSAIList field.
func (o *ServiceProfileChargingInformation) SetSNSSAIList(v []Snssai) {
	o.SNSSAIList = v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetLatency() int32 {
	if o == nil || isNil(o.Latency) {
		var ret int32
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetLatencyOk() (*int32, bool) {
	if o == nil || isNil(o.Latency) {
    return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasLatency() bool {
	if o != nil && !isNil(o.Latency) {
		return true
	}

	return false
}

// SetLatency gets a reference to the given int32 and assigns it to the Latency field.
func (o *ServiceProfileChargingInformation) SetLatency(v int32) {
	o.Latency = &v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetAvailability() float32 {
	if o == nil || isNil(o.Availability) {
		var ret float32
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetAvailabilityOk() (*float32, bool) {
	if o == nil || isNil(o.Availability) {
    return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasAvailability() bool {
	if o != nil && !isNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given float32 and assigns it to the Availability field.
func (o *ServiceProfileChargingInformation) SetAvailability(v float32) {
	o.Availability = &v
}

// GetJitter returns the Jitter field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetJitter() int32 {
	if o == nil || isNil(o.Jitter) {
		var ret int32
		return ret
	}
	return *o.Jitter
}

// GetJitterOk returns a tuple with the Jitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetJitterOk() (*int32, bool) {
	if o == nil || isNil(o.Jitter) {
    return nil, false
	}
	return o.Jitter, true
}

// HasJitter returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasJitter() bool {
	if o != nil && !isNil(o.Jitter) {
		return true
	}

	return false
}

// SetJitter gets a reference to the given int32 and assigns it to the Jitter field.
func (o *ServiceProfileChargingInformation) SetJitter(v int32) {
	o.Jitter = &v
}

// GetReliability returns the Reliability field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetReliability() string {
	if o == nil || isNil(o.Reliability) {
		var ret string
		return ret
	}
	return *o.Reliability
}

// GetReliabilityOk returns a tuple with the Reliability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetReliabilityOk() (*string, bool) {
	if o == nil || isNil(o.Reliability) {
    return nil, false
	}
	return o.Reliability, true
}

// HasReliability returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasReliability() bool {
	if o != nil && !isNil(o.Reliability) {
		return true
	}

	return false
}

// SetReliability gets a reference to the given string and assigns it to the Reliability field.
func (o *ServiceProfileChargingInformation) SetReliability(v string) {
	o.Reliability = &v
}

// GetMaxNumberofUEs returns the MaxNumberofUEs field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetMaxNumberofUEs() int32 {
	if o == nil || isNil(o.MaxNumberofUEs) {
		var ret int32
		return ret
	}
	return *o.MaxNumberofUEs
}

// GetMaxNumberofUEsOk returns a tuple with the MaxNumberofUEs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetMaxNumberofUEsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxNumberofUEs) {
    return nil, false
	}
	return o.MaxNumberofUEs, true
}

// HasMaxNumberofUEs returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasMaxNumberofUEs() bool {
	if o != nil && !isNil(o.MaxNumberofUEs) {
		return true
	}

	return false
}

// SetMaxNumberofUEs gets a reference to the given int32 and assigns it to the MaxNumberofUEs field.
func (o *ServiceProfileChargingInformation) SetMaxNumberofUEs(v int32) {
	o.MaxNumberofUEs = &v
}

// GetCoverageArea returns the CoverageArea field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetCoverageArea() string {
	if o == nil || isNil(o.CoverageArea) {
		var ret string
		return ret
	}
	return *o.CoverageArea
}

// GetCoverageAreaOk returns a tuple with the CoverageArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetCoverageAreaOk() (*string, bool) {
	if o == nil || isNil(o.CoverageArea) {
    return nil, false
	}
	return o.CoverageArea, true
}

// HasCoverageArea returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasCoverageArea() bool {
	if o != nil && !isNil(o.CoverageArea) {
		return true
	}

	return false
}

// SetCoverageArea gets a reference to the given string and assigns it to the CoverageArea field.
func (o *ServiceProfileChargingInformation) SetCoverageArea(v string) {
	o.CoverageArea = &v
}

// GetDLThptPerSlice returns the DLThptPerSlice field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetDLThptPerSlice() Throughput {
	if o == nil || isNil(o.DLThptPerSlice) {
		var ret Throughput
		return ret
	}
	return *o.DLThptPerSlice
}

// GetDLThptPerSliceOk returns a tuple with the DLThptPerSlice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetDLThptPerSliceOk() (*Throughput, bool) {
	if o == nil || isNil(o.DLThptPerSlice) {
    return nil, false
	}
	return o.DLThptPerSlice, true
}

// HasDLThptPerSlice returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasDLThptPerSlice() bool {
	if o != nil && !isNil(o.DLThptPerSlice) {
		return true
	}

	return false
}

// SetDLThptPerSlice gets a reference to the given Throughput and assigns it to the DLThptPerSlice field.
func (o *ServiceProfileChargingInformation) SetDLThptPerSlice(v Throughput) {
	o.DLThptPerSlice = &v
}

// GetDLThptPerUE returns the DLThptPerUE field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetDLThptPerUE() Throughput {
	if o == nil || isNil(o.DLThptPerUE) {
		var ret Throughput
		return ret
	}
	return *o.DLThptPerUE
}

// GetDLThptPerUEOk returns a tuple with the DLThptPerUE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetDLThptPerUEOk() (*Throughput, bool) {
	if o == nil || isNil(o.DLThptPerUE) {
    return nil, false
	}
	return o.DLThptPerUE, true
}

// HasDLThptPerUE returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasDLThptPerUE() bool {
	if o != nil && !isNil(o.DLThptPerUE) {
		return true
	}

	return false
}

// SetDLThptPerUE gets a reference to the given Throughput and assigns it to the DLThptPerUE field.
func (o *ServiceProfileChargingInformation) SetDLThptPerUE(v Throughput) {
	o.DLThptPerUE = &v
}

// GetULThptPerSlice returns the ULThptPerSlice field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetULThptPerSlice() Throughput {
	if o == nil || isNil(o.ULThptPerSlice) {
		var ret Throughput
		return ret
	}
	return *o.ULThptPerSlice
}

// GetULThptPerSliceOk returns a tuple with the ULThptPerSlice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetULThptPerSliceOk() (*Throughput, bool) {
	if o == nil || isNil(o.ULThptPerSlice) {
    return nil, false
	}
	return o.ULThptPerSlice, true
}

// HasULThptPerSlice returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasULThptPerSlice() bool {
	if o != nil && !isNil(o.ULThptPerSlice) {
		return true
	}

	return false
}

// SetULThptPerSlice gets a reference to the given Throughput and assigns it to the ULThptPerSlice field.
func (o *ServiceProfileChargingInformation) SetULThptPerSlice(v Throughput) {
	o.ULThptPerSlice = &v
}

// GetULThptPerUE returns the ULThptPerUE field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetULThptPerUE() Throughput {
	if o == nil || isNil(o.ULThptPerUE) {
		var ret Throughput
		return ret
	}
	return *o.ULThptPerUE
}

// GetULThptPerUEOk returns a tuple with the ULThptPerUE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetULThptPerUEOk() (*Throughput, bool) {
	if o == nil || isNil(o.ULThptPerUE) {
    return nil, false
	}
	return o.ULThptPerUE, true
}

// HasULThptPerUE returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasULThptPerUE() bool {
	if o != nil && !isNil(o.ULThptPerUE) {
		return true
	}

	return false
}

// SetULThptPerUE gets a reference to the given Throughput and assigns it to the ULThptPerUE field.
func (o *ServiceProfileChargingInformation) SetULThptPerUE(v Throughput) {
	o.ULThptPerUE = &v
}

// GetMaxNumberofPDUsessions returns the MaxNumberofPDUsessions field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetMaxNumberofPDUsessions() int32 {
	if o == nil || isNil(o.MaxNumberofPDUsessions) {
		var ret int32
		return ret
	}
	return *o.MaxNumberofPDUsessions
}

// GetMaxNumberofPDUsessionsOk returns a tuple with the MaxNumberofPDUsessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetMaxNumberofPDUsessionsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxNumberofPDUsessions) {
    return nil, false
	}
	return o.MaxNumberofPDUsessions, true
}

// HasMaxNumberofPDUsessions returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasMaxNumberofPDUsessions() bool {
	if o != nil && !isNil(o.MaxNumberofPDUsessions) {
		return true
	}

	return false
}

// SetMaxNumberofPDUsessions gets a reference to the given int32 and assigns it to the MaxNumberofPDUsessions field.
func (o *ServiceProfileChargingInformation) SetMaxNumberofPDUsessions(v int32) {
	o.MaxNumberofPDUsessions = &v
}

// GetKPIMonitoringList returns the KPIMonitoringList field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetKPIMonitoringList() string {
	if o == nil || isNil(o.KPIMonitoringList) {
		var ret string
		return ret
	}
	return *o.KPIMonitoringList
}

// GetKPIMonitoringListOk returns a tuple with the KPIMonitoringList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetKPIMonitoringListOk() (*string, bool) {
	if o == nil || isNil(o.KPIMonitoringList) {
    return nil, false
	}
	return o.KPIMonitoringList, true
}

// HasKPIMonitoringList returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasKPIMonitoringList() bool {
	if o != nil && !isNil(o.KPIMonitoringList) {
		return true
	}

	return false
}

// SetKPIMonitoringList gets a reference to the given string and assigns it to the KPIMonitoringList field.
func (o *ServiceProfileChargingInformation) SetKPIMonitoringList(v string) {
	o.KPIMonitoringList = &v
}

// GetSupportedAccessTechnology returns the SupportedAccessTechnology field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetSupportedAccessTechnology() int32 {
	if o == nil || isNil(o.SupportedAccessTechnology) {
		var ret int32
		return ret
	}
	return *o.SupportedAccessTechnology
}

// GetSupportedAccessTechnologyOk returns a tuple with the SupportedAccessTechnology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetSupportedAccessTechnologyOk() (*int32, bool) {
	if o == nil || isNil(o.SupportedAccessTechnology) {
    return nil, false
	}
	return o.SupportedAccessTechnology, true
}

// HasSupportedAccessTechnology returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasSupportedAccessTechnology() bool {
	if o != nil && !isNil(o.SupportedAccessTechnology) {
		return true
	}

	return false
}

// SetSupportedAccessTechnology gets a reference to the given int32 and assigns it to the SupportedAccessTechnology field.
func (o *ServiceProfileChargingInformation) SetSupportedAccessTechnology(v int32) {
	o.SupportedAccessTechnology = &v
}

// GetAddServiceProfileInfo returns the AddServiceProfileInfo field value if set, zero value otherwise.
func (o *ServiceProfileChargingInformation) GetAddServiceProfileInfo() string {
	if o == nil || isNil(o.AddServiceProfileInfo) {
		var ret string
		return ret
	}
	return *o.AddServiceProfileInfo
}

// GetAddServiceProfileInfoOk returns a tuple with the AddServiceProfileInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileChargingInformation) GetAddServiceProfileInfoOk() (*string, bool) {
	if o == nil || isNil(o.AddServiceProfileInfo) {
    return nil, false
	}
	return o.AddServiceProfileInfo, true
}

// HasAddServiceProfileInfo returns a boolean if a field has been set.
func (o *ServiceProfileChargingInformation) HasAddServiceProfileInfo() bool {
	if o != nil && !isNil(o.AddServiceProfileInfo) {
		return true
	}

	return false
}

// SetAddServiceProfileInfo gets a reference to the given string and assigns it to the AddServiceProfileInfo field.
func (o *ServiceProfileChargingInformation) SetAddServiceProfileInfo(v string) {
	o.AddServiceProfileInfo = &v
}

func (o ServiceProfileChargingInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServiceProfileIdentifier) {
		toSerialize["serviceProfileIdentifier"] = o.ServiceProfileIdentifier
	}
	if !isNil(o.SNSSAIList) {
		toSerialize["sNSSAIList"] = o.SNSSAIList
	}
	if !isNil(o.Latency) {
		toSerialize["latency"] = o.Latency
	}
	if !isNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !isNil(o.Jitter) {
		toSerialize["jitter"] = o.Jitter
	}
	if !isNil(o.Reliability) {
		toSerialize["reliability"] = o.Reliability
	}
	if !isNil(o.MaxNumberofUEs) {
		toSerialize["maxNumberofUEs"] = o.MaxNumberofUEs
	}
	if !isNil(o.CoverageArea) {
		toSerialize["coverageArea"] = o.CoverageArea
	}
	if !isNil(o.DLThptPerSlice) {
		toSerialize["dLThptPerSlice"] = o.DLThptPerSlice
	}
	if !isNil(o.DLThptPerUE) {
		toSerialize["dLThptPerUE"] = o.DLThptPerUE
	}
	if !isNil(o.ULThptPerSlice) {
		toSerialize["uLThptPerSlice"] = o.ULThptPerSlice
	}
	if !isNil(o.ULThptPerUE) {
		toSerialize["uLThptPerUE"] = o.ULThptPerUE
	}
	if !isNil(o.MaxNumberofPDUsessions) {
		toSerialize["maxNumberofPDUsessions"] = o.MaxNumberofPDUsessions
	}
	if !isNil(o.KPIMonitoringList) {
		toSerialize["kPIMonitoringList"] = o.KPIMonitoringList
	}
	if !isNil(o.SupportedAccessTechnology) {
		toSerialize["supportedAccessTechnology"] = o.SupportedAccessTechnology
	}
	if !isNil(o.AddServiceProfileInfo) {
		toSerialize["addServiceProfileInfo"] = o.AddServiceProfileInfo
	}
	return json.Marshal(toSerialize)
}

type NullableServiceProfileChargingInformation struct {
	value *ServiceProfileChargingInformation
	isSet bool
}

func (v NullableServiceProfileChargingInformation) Get() *ServiceProfileChargingInformation {
	return v.value
}

func (v *NullableServiceProfileChargingInformation) Set(val *ServiceProfileChargingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileChargingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileChargingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileChargingInformation(val *ServiceProfileChargingInformation) *NullableServiceProfileChargingInformation {
	return &NullableServiceProfileChargingInformation{value: val, isSet: true}
}

func (v NullableServiceProfileChargingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileChargingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


