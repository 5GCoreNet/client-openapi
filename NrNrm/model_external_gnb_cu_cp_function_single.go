/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// ExternalGnbCuCpFunctionSingle struct for ExternalGnbCuCpFunctionSingle
type ExternalGnbCuCpFunctionSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`
	ManagedNFService []ManagedNFServiceSingle `json:"ManagedNFService,omitempty"`
	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`
	ExternalNrCellCu []ExternalNrCellCuSingle `json:"ExternalNrCellCu,omitempty"`
	EPXnC []EPXnCSingle `json:"EP_XnC,omitempty"`
	EPE1 []EPE1Single `json:"EP_E1,omitempty"`
	EPF1C []EPF1CSingle `json:"EP_F1C,omitempty"`
}

// NewExternalGnbCuCpFunctionSingle instantiates a new ExternalGnbCuCpFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGnbCuCpFunctionSingle(id NullableString) *ExternalGnbCuCpFunctionSingle {
	this := ExternalGnbCuCpFunctionSingle{}
	this.Id = id
	return &this
}

// NewExternalGnbCuCpFunctionSingleWithDefaults instantiates a new ExternalGnbCuCpFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGnbCuCpFunctionSingleWithDefaults() *ExternalGnbCuCpFunctionSingle {
	this := ExternalGnbCuCpFunctionSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExternalGnbCuCpFunctionSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalGnbCuCpFunctionSingle) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *ExternalGnbCuCpFunctionSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
    return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *ExternalGnbCuCpFunctionSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
    return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *ExternalGnbCuCpFunctionSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
    return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *ExternalGnbCuCpFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingle) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *ExternalGnbCuCpFunctionSingle) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || isNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || isNil(o.PerfMetricJob) {
    return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingle) HasPerfMetricJob() bool {
	if o != nil && !isNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *ExternalGnbCuCpFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || isNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || isNil(o.ThresholdMonitor) {
    return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingle) HasThresholdMonitor() bool {
	if o != nil && !isNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *ExternalGnbCuCpFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetManagedNFService returns the ManagedNFService field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle {
	if o == nil || isNil(o.ManagedNFService) {
		var ret []ManagedNFServiceSingle
		return ret
	}
	return o.ManagedNFService
}

// GetManagedNFServiceOk returns a tuple with the ManagedNFService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingle) GetManagedNFServiceOk() ([]ManagedNFServiceSingle, bool) {
	if o == nil || isNil(o.ManagedNFService) {
    return nil, false
	}
	return o.ManagedNFService, true
}

// HasManagedNFService returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingle) HasManagedNFService() bool {
	if o != nil && !isNil(o.ManagedNFService) {
		return true
	}

	return false
}

// SetManagedNFService gets a reference to the given []ManagedNFServiceSingle and assigns it to the ManagedNFService field.
func (o *ExternalGnbCuCpFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle) {
	o.ManagedNFService = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || isNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || isNil(o.TraceJob) {
    return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingle) HasTraceJob() bool {
	if o != nil && !isNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *ExternalGnbCuCpFunctionSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetExternalNrCellCu returns the ExternalNrCellCu field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingle) GetExternalNrCellCu() []ExternalNrCellCuSingle {
	if o == nil || isNil(o.ExternalNrCellCu) {
		var ret []ExternalNrCellCuSingle
		return ret
	}
	return o.ExternalNrCellCu
}

// GetExternalNrCellCuOk returns a tuple with the ExternalNrCellCu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingle) GetExternalNrCellCuOk() ([]ExternalNrCellCuSingle, bool) {
	if o == nil || isNil(o.ExternalNrCellCu) {
    return nil, false
	}
	return o.ExternalNrCellCu, true
}

// HasExternalNrCellCu returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingle) HasExternalNrCellCu() bool {
	if o != nil && !isNil(o.ExternalNrCellCu) {
		return true
	}

	return false
}

// SetExternalNrCellCu gets a reference to the given []ExternalNrCellCuSingle and assigns it to the ExternalNrCellCu field.
func (o *ExternalGnbCuCpFunctionSingle) SetExternalNrCellCu(v []ExternalNrCellCuSingle) {
	o.ExternalNrCellCu = v
}

// GetEPXnC returns the EPXnC field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingle) GetEPXnC() []EPXnCSingle {
	if o == nil || isNil(o.EPXnC) {
		var ret []EPXnCSingle
		return ret
	}
	return o.EPXnC
}

// GetEPXnCOk returns a tuple with the EPXnC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingle) GetEPXnCOk() ([]EPXnCSingle, bool) {
	if o == nil || isNil(o.EPXnC) {
    return nil, false
	}
	return o.EPXnC, true
}

// HasEPXnC returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingle) HasEPXnC() bool {
	if o != nil && !isNil(o.EPXnC) {
		return true
	}

	return false
}

// SetEPXnC gets a reference to the given []EPXnCSingle and assigns it to the EPXnC field.
func (o *ExternalGnbCuCpFunctionSingle) SetEPXnC(v []EPXnCSingle) {
	o.EPXnC = v
}

// GetEPE1 returns the EPE1 field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingle) GetEPE1() []EPE1Single {
	if o == nil || isNil(o.EPE1) {
		var ret []EPE1Single
		return ret
	}
	return o.EPE1
}

// GetEPE1Ok returns a tuple with the EPE1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingle) GetEPE1Ok() ([]EPE1Single, bool) {
	if o == nil || isNil(o.EPE1) {
    return nil, false
	}
	return o.EPE1, true
}

// HasEPE1 returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingle) HasEPE1() bool {
	if o != nil && !isNil(o.EPE1) {
		return true
	}

	return false
}

// SetEPE1 gets a reference to the given []EPE1Single and assigns it to the EPE1 field.
func (o *ExternalGnbCuCpFunctionSingle) SetEPE1(v []EPE1Single) {
	o.EPE1 = v
}

// GetEPF1C returns the EPF1C field value if set, zero value otherwise.
func (o *ExternalGnbCuCpFunctionSingle) GetEPF1C() []EPF1CSingle {
	if o == nil || isNil(o.EPF1C) {
		var ret []EPF1CSingle
		return ret
	}
	return o.EPF1C
}

// GetEPF1COk returns a tuple with the EPF1C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGnbCuCpFunctionSingle) GetEPF1COk() ([]EPF1CSingle, bool) {
	if o == nil || isNil(o.EPF1C) {
    return nil, false
	}
	return o.EPF1C, true
}

// HasEPF1C returns a boolean if a field has been set.
func (o *ExternalGnbCuCpFunctionSingle) HasEPF1C() bool {
	if o != nil && !isNil(o.EPF1C) {
		return true
	}

	return false
}

// SetEPF1C gets a reference to the given []EPF1CSingle and assigns it to the EPF1C field.
func (o *ExternalGnbCuCpFunctionSingle) SetEPF1C(v []EPF1CSingle) {
	o.EPF1C = v
}

func (o ExternalGnbCuCpFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id.Get()
	}
	if !isNil(o.ObjectClass) {
		toSerialize["objectClass"] = o.ObjectClass
	}
	if !isNil(o.ObjectInstance) {
		toSerialize["objectInstance"] = o.ObjectInstance
	}
	if !isNil(o.VsDataContainer) {
		toSerialize["VsDataContainer"] = o.VsDataContainer
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.PerfMetricJob) {
		toSerialize["PerfMetricJob"] = o.PerfMetricJob
	}
	if !isNil(o.ThresholdMonitor) {
		toSerialize["ThresholdMonitor"] = o.ThresholdMonitor
	}
	if !isNil(o.ManagedNFService) {
		toSerialize["ManagedNFService"] = o.ManagedNFService
	}
	if !isNil(o.TraceJob) {
		toSerialize["TraceJob"] = o.TraceJob
	}
	if !isNil(o.ExternalNrCellCu) {
		toSerialize["ExternalNrCellCu"] = o.ExternalNrCellCu
	}
	if !isNil(o.EPXnC) {
		toSerialize["EP_XnC"] = o.EPXnC
	}
	if !isNil(o.EPE1) {
		toSerialize["EP_E1"] = o.EPE1
	}
	if !isNil(o.EPF1C) {
		toSerialize["EP_F1C"] = o.EPF1C
	}
	return json.Marshal(toSerialize)
}

type NullableExternalGnbCuCpFunctionSingle struct {
	value *ExternalGnbCuCpFunctionSingle
	isSet bool
}

func (v NullableExternalGnbCuCpFunctionSingle) Get() *ExternalGnbCuCpFunctionSingle {
	return v.value
}

func (v *NullableExternalGnbCuCpFunctionSingle) Set(val *ExternalGnbCuCpFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGnbCuCpFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGnbCuCpFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGnbCuCpFunctionSingle(val *ExternalGnbCuCpFunctionSingle) *NullableExternalGnbCuCpFunctionSingle {
	return &NullableExternalGnbCuCpFunctionSingle{value: val, isSet: true}
}

func (v NullableExternalGnbCuCpFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGnbCuCpFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


