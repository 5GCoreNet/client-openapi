/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
)

// ManagedElementSingle struct for ManagedElementSingle
type ManagedElementSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *ManagedElementAttr `json:"attributes,omitempty"`
	MnsAgent []MnsAgentSingle `json:"MnsAgent,omitempty"`
	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`
	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`
	NtfSubscriptionControl []NtfSubscriptionControlSingle `json:"NtfSubscriptionControl,omitempty"`
	AlarmList *AlarmListSingle `json:"AlarmList,omitempty"`
	FileDownloadJob []FileDownloadJobSingle `json:"FileDownloadJob,omitempty"`
	Files []FilesSingle `json:"Files,omitempty"`
	MDAFunction []MDAFunctionSingle `json:"MDAFunction,omitempty"`
}

// NewManagedElementSingle instantiates a new ManagedElementSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedElementSingle(id NullableString) *ManagedElementSingle {
	this := ManagedElementSingle{}
	this.Id = id
	return &this
}

// NewManagedElementSingleWithDefaults instantiates a new ManagedElementSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedElementSingleWithDefaults() *ManagedElementSingle {
	this := ManagedElementSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ManagedElementSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedElementSingle) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *ManagedElementSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
    return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *ManagedElementSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
    return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *ManagedElementSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
    return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *ManagedElementSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetAttributes() ManagedElementAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedElementAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetAttributesOk() (*ManagedElementAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedElementAttr and assigns it to the Attributes field.
func (o *ManagedElementSingle) SetAttributes(v ManagedElementAttr) {
	o.Attributes = &v
}

// GetMnsAgent returns the MnsAgent field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetMnsAgent() []MnsAgentSingle {
	if o == nil || isNil(o.MnsAgent) {
		var ret []MnsAgentSingle
		return ret
	}
	return o.MnsAgent
}

// GetMnsAgentOk returns a tuple with the MnsAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetMnsAgentOk() ([]MnsAgentSingle, bool) {
	if o == nil || isNil(o.MnsAgent) {
    return nil, false
	}
	return o.MnsAgent, true
}

// HasMnsAgent returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasMnsAgent() bool {
	if o != nil && !isNil(o.MnsAgent) {
		return true
	}

	return false
}

// SetMnsAgent gets a reference to the given []MnsAgentSingle and assigns it to the MnsAgent field.
func (o *ManagedElementSingle) SetMnsAgent(v []MnsAgentSingle) {
	o.MnsAgent = v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || isNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || isNil(o.PerfMetricJob) {
    return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasPerfMetricJob() bool {
	if o != nil && !isNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *ManagedElementSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || isNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || isNil(o.ThresholdMonitor) {
    return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasThresholdMonitor() bool {
	if o != nil && !isNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *ManagedElementSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || isNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || isNil(o.TraceJob) {
    return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasTraceJob() bool {
	if o != nil && !isNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *ManagedElementSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetNtfSubscriptionControl returns the NtfSubscriptionControl field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle {
	if o == nil || isNil(o.NtfSubscriptionControl) {
		var ret []NtfSubscriptionControlSingle
		return ret
	}
	return o.NtfSubscriptionControl
}

// GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetNtfSubscriptionControlOk() ([]NtfSubscriptionControlSingle, bool) {
	if o == nil || isNil(o.NtfSubscriptionControl) {
    return nil, false
	}
	return o.NtfSubscriptionControl, true
}

// HasNtfSubscriptionControl returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasNtfSubscriptionControl() bool {
	if o != nil && !isNil(o.NtfSubscriptionControl) {
		return true
	}

	return false
}

// SetNtfSubscriptionControl gets a reference to the given []NtfSubscriptionControlSingle and assigns it to the NtfSubscriptionControl field.
func (o *ManagedElementSingle) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle) {
	o.NtfSubscriptionControl = v
}

// GetAlarmList returns the AlarmList field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetAlarmList() AlarmListSingle {
	if o == nil || isNil(o.AlarmList) {
		var ret AlarmListSingle
		return ret
	}
	return *o.AlarmList
}

// GetAlarmListOk returns a tuple with the AlarmList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetAlarmListOk() (*AlarmListSingle, bool) {
	if o == nil || isNil(o.AlarmList) {
    return nil, false
	}
	return o.AlarmList, true
}

// HasAlarmList returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasAlarmList() bool {
	if o != nil && !isNil(o.AlarmList) {
		return true
	}

	return false
}

// SetAlarmList gets a reference to the given AlarmListSingle and assigns it to the AlarmList field.
func (o *ManagedElementSingle) SetAlarmList(v AlarmListSingle) {
	o.AlarmList = &v
}

// GetFileDownloadJob returns the FileDownloadJob field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetFileDownloadJob() []FileDownloadJobSingle {
	if o == nil || isNil(o.FileDownloadJob) {
		var ret []FileDownloadJobSingle
		return ret
	}
	return o.FileDownloadJob
}

// GetFileDownloadJobOk returns a tuple with the FileDownloadJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetFileDownloadJobOk() ([]FileDownloadJobSingle, bool) {
	if o == nil || isNil(o.FileDownloadJob) {
    return nil, false
	}
	return o.FileDownloadJob, true
}

// HasFileDownloadJob returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasFileDownloadJob() bool {
	if o != nil && !isNil(o.FileDownloadJob) {
		return true
	}

	return false
}

// SetFileDownloadJob gets a reference to the given []FileDownloadJobSingle and assigns it to the FileDownloadJob field.
func (o *ManagedElementSingle) SetFileDownloadJob(v []FileDownloadJobSingle) {
	o.FileDownloadJob = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetFiles() []FilesSingle {
	if o == nil || isNil(o.Files) {
		var ret []FilesSingle
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetFilesOk() ([]FilesSingle, bool) {
	if o == nil || isNil(o.Files) {
    return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasFiles() bool {
	if o != nil && !isNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FilesSingle and assigns it to the Files field.
func (o *ManagedElementSingle) SetFiles(v []FilesSingle) {
	o.Files = v
}

// GetMDAFunction returns the MDAFunction field value if set, zero value otherwise.
func (o *ManagedElementSingle) GetMDAFunction() []MDAFunctionSingle {
	if o == nil || isNil(o.MDAFunction) {
		var ret []MDAFunctionSingle
		return ret
	}
	return o.MDAFunction
}

// GetMDAFunctionOk returns a tuple with the MDAFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingle) GetMDAFunctionOk() ([]MDAFunctionSingle, bool) {
	if o == nil || isNil(o.MDAFunction) {
    return nil, false
	}
	return o.MDAFunction, true
}

// HasMDAFunction returns a boolean if a field has been set.
func (o *ManagedElementSingle) HasMDAFunction() bool {
	if o != nil && !isNil(o.MDAFunction) {
		return true
	}

	return false
}

// SetMDAFunction gets a reference to the given []MDAFunctionSingle and assigns it to the MDAFunction field.
func (o *ManagedElementSingle) SetMDAFunction(v []MDAFunctionSingle) {
	o.MDAFunction = v
}

func (o ManagedElementSingle) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.MnsAgent) {
		toSerialize["MnsAgent"] = o.MnsAgent
	}
	if !isNil(o.PerfMetricJob) {
		toSerialize["PerfMetricJob"] = o.PerfMetricJob
	}
	if !isNil(o.ThresholdMonitor) {
		toSerialize["ThresholdMonitor"] = o.ThresholdMonitor
	}
	if !isNil(o.TraceJob) {
		toSerialize["TraceJob"] = o.TraceJob
	}
	if !isNil(o.NtfSubscriptionControl) {
		toSerialize["NtfSubscriptionControl"] = o.NtfSubscriptionControl
	}
	if !isNil(o.AlarmList) {
		toSerialize["AlarmList"] = o.AlarmList
	}
	if !isNil(o.FileDownloadJob) {
		toSerialize["FileDownloadJob"] = o.FileDownloadJob
	}
	if !isNil(o.Files) {
		toSerialize["Files"] = o.Files
	}
	if !isNil(o.MDAFunction) {
		toSerialize["MDAFunction"] = o.MDAFunction
	}
	return json.Marshal(toSerialize)
}

type NullableManagedElementSingle struct {
	value *ManagedElementSingle
	isSet bool
}

func (v NullableManagedElementSingle) Get() *ManagedElementSingle {
	return v.value
}

func (v *NullableManagedElementSingle) Set(val *ManagedElementSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedElementSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedElementSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedElementSingle(val *ManagedElementSingle) *NullableManagedElementSingle {
	return &NullableManagedElementSingle{value: val, isSet: true}
}

func (v NullableManagedElementSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedElementSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


