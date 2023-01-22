/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
)

// NrCellCuSingle struct for NrCellCuSingle
type NrCellCuSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *ManagedFunctionAttr `json:"attributes,omitempty"`
	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`
	ManagedNFService []ManagedNFServiceSingle `json:"ManagedNFService,omitempty"`
	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`
	RRMPolicyRatio []RRMPolicyRatioSingle `json:"RRMPolicyRatio,omitempty"`
	NRCellRelation []NRCellRelationSingle `json:"NRCellRelation,omitempty"`
	EUtranCellRelation []EUtranCellRelationSingle `json:"EUtranCellRelation,omitempty"`
	NRFreqRelation []NRFreqRelationSingle `json:"NRFreqRelation,omitempty"`
	EUtranFreqRelation []EUtranFreqRelationSingle `json:"EUtranFreqRelation,omitempty"`
	DESManagementFunction *DESManagementFunctionSingle `json:"DESManagementFunction,omitempty"`
	DMROFunction *DMROFunctionSingle `json:"DMROFunction,omitempty"`
	DLBOFunction *DLBOFunctionSingle `json:"DLBOFunction,omitempty"`
	CESManagementFunction *CESManagementFunctionSingle `json:"CESManagementFunction,omitempty"`
	DPCIConfigurationFunction *DPCIConfigurationFunctionSingle `json:"DPCIConfigurationFunction,omitempty"`
}

// NewNrCellCuSingle instantiates a new NrCellCuSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrCellCuSingle(id NullableString) *NrCellCuSingle {
	this := NrCellCuSingle{}
	this.Id = id
	return &this
}

// NewNrCellCuSingleWithDefaults instantiates a new NrCellCuSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrCellCuSingleWithDefaults() *NrCellCuSingle {
	this := NrCellCuSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NrCellCuSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NrCellCuSingle) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *NrCellCuSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
    return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *NrCellCuSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
    return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *NrCellCuSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
    return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *NrCellCuSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetAttributes() ManagedFunctionAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedFunctionAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetAttributesOk() (*ManagedFunctionAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedFunctionAttr and assigns it to the Attributes field.
func (o *NrCellCuSingle) SetAttributes(v ManagedFunctionAttr) {
	o.Attributes = &v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || isNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || isNil(o.PerfMetricJob) {
    return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasPerfMetricJob() bool {
	if o != nil && !isNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *NrCellCuSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || isNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || isNil(o.ThresholdMonitor) {
    return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasThresholdMonitor() bool {
	if o != nil && !isNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *NrCellCuSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetManagedNFService returns the ManagedNFService field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetManagedNFService() []ManagedNFServiceSingle {
	if o == nil || isNil(o.ManagedNFService) {
		var ret []ManagedNFServiceSingle
		return ret
	}
	return o.ManagedNFService
}

// GetManagedNFServiceOk returns a tuple with the ManagedNFService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetManagedNFServiceOk() ([]ManagedNFServiceSingle, bool) {
	if o == nil || isNil(o.ManagedNFService) {
    return nil, false
	}
	return o.ManagedNFService, true
}

// HasManagedNFService returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasManagedNFService() bool {
	if o != nil && !isNil(o.ManagedNFService) {
		return true
	}

	return false
}

// SetManagedNFService gets a reference to the given []ManagedNFServiceSingle and assigns it to the ManagedNFService field.
func (o *NrCellCuSingle) SetManagedNFService(v []ManagedNFServiceSingle) {
	o.ManagedNFService = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || isNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || isNil(o.TraceJob) {
    return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasTraceJob() bool {
	if o != nil && !isNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *NrCellCuSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetRRMPolicyRatio returns the RRMPolicyRatio field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetRRMPolicyRatio() []RRMPolicyRatioSingle {
	if o == nil || isNil(o.RRMPolicyRatio) {
		var ret []RRMPolicyRatioSingle
		return ret
	}
	return o.RRMPolicyRatio
}

// GetRRMPolicyRatioOk returns a tuple with the RRMPolicyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetRRMPolicyRatioOk() ([]RRMPolicyRatioSingle, bool) {
	if o == nil || isNil(o.RRMPolicyRatio) {
    return nil, false
	}
	return o.RRMPolicyRatio, true
}

// HasRRMPolicyRatio returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasRRMPolicyRatio() bool {
	if o != nil && !isNil(o.RRMPolicyRatio) {
		return true
	}

	return false
}

// SetRRMPolicyRatio gets a reference to the given []RRMPolicyRatioSingle and assigns it to the RRMPolicyRatio field.
func (o *NrCellCuSingle) SetRRMPolicyRatio(v []RRMPolicyRatioSingle) {
	o.RRMPolicyRatio = v
}

// GetNRCellRelation returns the NRCellRelation field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetNRCellRelation() []NRCellRelationSingle {
	if o == nil || isNil(o.NRCellRelation) {
		var ret []NRCellRelationSingle
		return ret
	}
	return o.NRCellRelation
}

// GetNRCellRelationOk returns a tuple with the NRCellRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetNRCellRelationOk() ([]NRCellRelationSingle, bool) {
	if o == nil || isNil(o.NRCellRelation) {
    return nil, false
	}
	return o.NRCellRelation, true
}

// HasNRCellRelation returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasNRCellRelation() bool {
	if o != nil && !isNil(o.NRCellRelation) {
		return true
	}

	return false
}

// SetNRCellRelation gets a reference to the given []NRCellRelationSingle and assigns it to the NRCellRelation field.
func (o *NrCellCuSingle) SetNRCellRelation(v []NRCellRelationSingle) {
	o.NRCellRelation = v
}

// GetEUtranCellRelation returns the EUtranCellRelation field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetEUtranCellRelation() []EUtranCellRelationSingle {
	if o == nil || isNil(o.EUtranCellRelation) {
		var ret []EUtranCellRelationSingle
		return ret
	}
	return o.EUtranCellRelation
}

// GetEUtranCellRelationOk returns a tuple with the EUtranCellRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetEUtranCellRelationOk() ([]EUtranCellRelationSingle, bool) {
	if o == nil || isNil(o.EUtranCellRelation) {
    return nil, false
	}
	return o.EUtranCellRelation, true
}

// HasEUtranCellRelation returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasEUtranCellRelation() bool {
	if o != nil && !isNil(o.EUtranCellRelation) {
		return true
	}

	return false
}

// SetEUtranCellRelation gets a reference to the given []EUtranCellRelationSingle and assigns it to the EUtranCellRelation field.
func (o *NrCellCuSingle) SetEUtranCellRelation(v []EUtranCellRelationSingle) {
	o.EUtranCellRelation = v
}

// GetNRFreqRelation returns the NRFreqRelation field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetNRFreqRelation() []NRFreqRelationSingle {
	if o == nil || isNil(o.NRFreqRelation) {
		var ret []NRFreqRelationSingle
		return ret
	}
	return o.NRFreqRelation
}

// GetNRFreqRelationOk returns a tuple with the NRFreqRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetNRFreqRelationOk() ([]NRFreqRelationSingle, bool) {
	if o == nil || isNil(o.NRFreqRelation) {
    return nil, false
	}
	return o.NRFreqRelation, true
}

// HasNRFreqRelation returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasNRFreqRelation() bool {
	if o != nil && !isNil(o.NRFreqRelation) {
		return true
	}

	return false
}

// SetNRFreqRelation gets a reference to the given []NRFreqRelationSingle and assigns it to the NRFreqRelation field.
func (o *NrCellCuSingle) SetNRFreqRelation(v []NRFreqRelationSingle) {
	o.NRFreqRelation = v
}

// GetEUtranFreqRelation returns the EUtranFreqRelation field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetEUtranFreqRelation() []EUtranFreqRelationSingle {
	if o == nil || isNil(o.EUtranFreqRelation) {
		var ret []EUtranFreqRelationSingle
		return ret
	}
	return o.EUtranFreqRelation
}

// GetEUtranFreqRelationOk returns a tuple with the EUtranFreqRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetEUtranFreqRelationOk() ([]EUtranFreqRelationSingle, bool) {
	if o == nil || isNil(o.EUtranFreqRelation) {
    return nil, false
	}
	return o.EUtranFreqRelation, true
}

// HasEUtranFreqRelation returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasEUtranFreqRelation() bool {
	if o != nil && !isNil(o.EUtranFreqRelation) {
		return true
	}

	return false
}

// SetEUtranFreqRelation gets a reference to the given []EUtranFreqRelationSingle and assigns it to the EUtranFreqRelation field.
func (o *NrCellCuSingle) SetEUtranFreqRelation(v []EUtranFreqRelationSingle) {
	o.EUtranFreqRelation = v
}

// GetDESManagementFunction returns the DESManagementFunction field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetDESManagementFunction() DESManagementFunctionSingle {
	if o == nil || isNil(o.DESManagementFunction) {
		var ret DESManagementFunctionSingle
		return ret
	}
	return *o.DESManagementFunction
}

// GetDESManagementFunctionOk returns a tuple with the DESManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetDESManagementFunctionOk() (*DESManagementFunctionSingle, bool) {
	if o == nil || isNil(o.DESManagementFunction) {
    return nil, false
	}
	return o.DESManagementFunction, true
}

// HasDESManagementFunction returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasDESManagementFunction() bool {
	if o != nil && !isNil(o.DESManagementFunction) {
		return true
	}

	return false
}

// SetDESManagementFunction gets a reference to the given DESManagementFunctionSingle and assigns it to the DESManagementFunction field.
func (o *NrCellCuSingle) SetDESManagementFunction(v DESManagementFunctionSingle) {
	o.DESManagementFunction = &v
}

// GetDMROFunction returns the DMROFunction field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetDMROFunction() DMROFunctionSingle {
	if o == nil || isNil(o.DMROFunction) {
		var ret DMROFunctionSingle
		return ret
	}
	return *o.DMROFunction
}

// GetDMROFunctionOk returns a tuple with the DMROFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetDMROFunctionOk() (*DMROFunctionSingle, bool) {
	if o == nil || isNil(o.DMROFunction) {
    return nil, false
	}
	return o.DMROFunction, true
}

// HasDMROFunction returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasDMROFunction() bool {
	if o != nil && !isNil(o.DMROFunction) {
		return true
	}

	return false
}

// SetDMROFunction gets a reference to the given DMROFunctionSingle and assigns it to the DMROFunction field.
func (o *NrCellCuSingle) SetDMROFunction(v DMROFunctionSingle) {
	o.DMROFunction = &v
}

// GetDLBOFunction returns the DLBOFunction field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetDLBOFunction() DLBOFunctionSingle {
	if o == nil || isNil(o.DLBOFunction) {
		var ret DLBOFunctionSingle
		return ret
	}
	return *o.DLBOFunction
}

// GetDLBOFunctionOk returns a tuple with the DLBOFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetDLBOFunctionOk() (*DLBOFunctionSingle, bool) {
	if o == nil || isNil(o.DLBOFunction) {
    return nil, false
	}
	return o.DLBOFunction, true
}

// HasDLBOFunction returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasDLBOFunction() bool {
	if o != nil && !isNil(o.DLBOFunction) {
		return true
	}

	return false
}

// SetDLBOFunction gets a reference to the given DLBOFunctionSingle and assigns it to the DLBOFunction field.
func (o *NrCellCuSingle) SetDLBOFunction(v DLBOFunctionSingle) {
	o.DLBOFunction = &v
}

// GetCESManagementFunction returns the CESManagementFunction field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetCESManagementFunction() CESManagementFunctionSingle {
	if o == nil || isNil(o.CESManagementFunction) {
		var ret CESManagementFunctionSingle
		return ret
	}
	return *o.CESManagementFunction
}

// GetCESManagementFunctionOk returns a tuple with the CESManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetCESManagementFunctionOk() (*CESManagementFunctionSingle, bool) {
	if o == nil || isNil(o.CESManagementFunction) {
    return nil, false
	}
	return o.CESManagementFunction, true
}

// HasCESManagementFunction returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasCESManagementFunction() bool {
	if o != nil && !isNil(o.CESManagementFunction) {
		return true
	}

	return false
}

// SetCESManagementFunction gets a reference to the given CESManagementFunctionSingle and assigns it to the CESManagementFunction field.
func (o *NrCellCuSingle) SetCESManagementFunction(v CESManagementFunctionSingle) {
	o.CESManagementFunction = &v
}

// GetDPCIConfigurationFunction returns the DPCIConfigurationFunction field value if set, zero value otherwise.
func (o *NrCellCuSingle) GetDPCIConfigurationFunction() DPCIConfigurationFunctionSingle {
	if o == nil || isNil(o.DPCIConfigurationFunction) {
		var ret DPCIConfigurationFunctionSingle
		return ret
	}
	return *o.DPCIConfigurationFunction
}

// GetDPCIConfigurationFunctionOk returns a tuple with the DPCIConfigurationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingle) GetDPCIConfigurationFunctionOk() (*DPCIConfigurationFunctionSingle, bool) {
	if o == nil || isNil(o.DPCIConfigurationFunction) {
    return nil, false
	}
	return o.DPCIConfigurationFunction, true
}

// HasDPCIConfigurationFunction returns a boolean if a field has been set.
func (o *NrCellCuSingle) HasDPCIConfigurationFunction() bool {
	if o != nil && !isNil(o.DPCIConfigurationFunction) {
		return true
	}

	return false
}

// SetDPCIConfigurationFunction gets a reference to the given DPCIConfigurationFunctionSingle and assigns it to the DPCIConfigurationFunction field.
func (o *NrCellCuSingle) SetDPCIConfigurationFunction(v DPCIConfigurationFunctionSingle) {
	o.DPCIConfigurationFunction = &v
}

func (o NrCellCuSingle) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.RRMPolicyRatio) {
		toSerialize["RRMPolicyRatio"] = o.RRMPolicyRatio
	}
	if !isNil(o.NRCellRelation) {
		toSerialize["NRCellRelation"] = o.NRCellRelation
	}
	if !isNil(o.EUtranCellRelation) {
		toSerialize["EUtranCellRelation"] = o.EUtranCellRelation
	}
	if !isNil(o.NRFreqRelation) {
		toSerialize["NRFreqRelation"] = o.NRFreqRelation
	}
	if !isNil(o.EUtranFreqRelation) {
		toSerialize["EUtranFreqRelation"] = o.EUtranFreqRelation
	}
	if !isNil(o.DESManagementFunction) {
		toSerialize["DESManagementFunction"] = o.DESManagementFunction
	}
	if !isNil(o.DMROFunction) {
		toSerialize["DMROFunction"] = o.DMROFunction
	}
	if !isNil(o.DLBOFunction) {
		toSerialize["DLBOFunction"] = o.DLBOFunction
	}
	if !isNil(o.CESManagementFunction) {
		toSerialize["CESManagementFunction"] = o.CESManagementFunction
	}
	if !isNil(o.DPCIConfigurationFunction) {
		toSerialize["DPCIConfigurationFunction"] = o.DPCIConfigurationFunction
	}
	return json.Marshal(toSerialize)
}

type NullableNrCellCuSingle struct {
	value *NrCellCuSingle
	isSet bool
}

func (v NullableNrCellCuSingle) Get() *NrCellCuSingle {
	return v.value
}

func (v *NullableNrCellCuSingle) Set(val *NrCellCuSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableNrCellCuSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableNrCellCuSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrCellCuSingle(val *NrCellCuSingle) *NullableNrCellCuSingle {
	return &NullableNrCellCuSingle{value: val, isSet: true}
}

func (v NullableNrCellCuSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrCellCuSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


