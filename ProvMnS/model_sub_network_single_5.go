/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// SubNetworkSingle5 struct for SubNetworkSingle5
type SubNetworkSingle5 struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *SubNetworkAttr `json:"attributes,omitempty"`
	ManagementNode []ManagementNodeSingle `json:"ManagementNode,omitempty"`
	MnsAgent []MnsAgentSingle `json:"MnsAgent,omitempty"`
	MeContext []MeContextSingle `json:"MeContext,omitempty"`
	PerfMetricJob []PerfMetricJobSingle `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle `json:"ThresholdMonitor,omitempty"`
	TraceJob []TraceJobSingle `json:"TraceJob,omitempty"`
	ManagementDataCollection []ManagementDataCollectionSingle `json:"ManagementDataCollection,omitempty"`
	NtfSubscriptionControl []NtfSubscriptionControlSingle `json:"NtfSubscriptionControl,omitempty"`
	AlarmList *AlarmListSingle `json:"AlarmList,omitempty"`
	Files []FilesSingle `json:"Files,omitempty"`
	FileDownloadJob []FileDownloadJobSingle `json:"FileDownloadJob,omitempty"`
	MnsRegistry *MnsRegistrySingle `json:"MnsRegistry,omitempty"`
	SubNetwork []SubNetworkSingle `json:"SubNetwork,omitempty"`
	ManagedElement []ManagedElementSingle `json:"ManagedElement,omitempty"`
	MLTrainingFunction []MLTrainingFunctionSingle `json:"MLTrainingFunction,omitempty"`
}

// NewSubNetworkSingle5 instantiates a new SubNetworkSingle5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingle5(id NullableString) *SubNetworkSingle5 {
	this := SubNetworkSingle5{}
	this.Id = id
	return &this
}

// NewSubNetworkSingle5WithDefaults instantiates a new SubNetworkSingle5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingle5WithDefaults() *SubNetworkSingle5 {
	this := SubNetworkSingle5{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubNetworkSingle5) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubNetworkSingle5) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *SubNetworkSingle5) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
    return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *SubNetworkSingle5) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
    return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *SubNetworkSingle5) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
    return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *SubNetworkSingle5) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetAttributes() SubNetworkAttr {
	if o == nil || isNil(o.Attributes) {
		var ret SubNetworkAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetAttributesOk() (*SubNetworkAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubNetworkAttr and assigns it to the Attributes field.
func (o *SubNetworkSingle5) SetAttributes(v SubNetworkAttr) {
	o.Attributes = &v
}

// GetManagementNode returns the ManagementNode field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetManagementNode() []ManagementNodeSingle {
	if o == nil || isNil(o.ManagementNode) {
		var ret []ManagementNodeSingle
		return ret
	}
	return o.ManagementNode
}

// GetManagementNodeOk returns a tuple with the ManagementNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetManagementNodeOk() ([]ManagementNodeSingle, bool) {
	if o == nil || isNil(o.ManagementNode) {
    return nil, false
	}
	return o.ManagementNode, true
}

// HasManagementNode returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasManagementNode() bool {
	if o != nil && !isNil(o.ManagementNode) {
		return true
	}

	return false
}

// SetManagementNode gets a reference to the given []ManagementNodeSingle and assigns it to the ManagementNode field.
func (o *SubNetworkSingle5) SetManagementNode(v []ManagementNodeSingle) {
	o.ManagementNode = v
}

// GetMnsAgent returns the MnsAgent field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetMnsAgent() []MnsAgentSingle {
	if o == nil || isNil(o.MnsAgent) {
		var ret []MnsAgentSingle
		return ret
	}
	return o.MnsAgent
}

// GetMnsAgentOk returns a tuple with the MnsAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetMnsAgentOk() ([]MnsAgentSingle, bool) {
	if o == nil || isNil(o.MnsAgent) {
    return nil, false
	}
	return o.MnsAgent, true
}

// HasMnsAgent returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasMnsAgent() bool {
	if o != nil && !isNil(o.MnsAgent) {
		return true
	}

	return false
}

// SetMnsAgent gets a reference to the given []MnsAgentSingle and assigns it to the MnsAgent field.
func (o *SubNetworkSingle5) SetMnsAgent(v []MnsAgentSingle) {
	o.MnsAgent = v
}

// GetMeContext returns the MeContext field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetMeContext() []MeContextSingle {
	if o == nil || isNil(o.MeContext) {
		var ret []MeContextSingle
		return ret
	}
	return o.MeContext
}

// GetMeContextOk returns a tuple with the MeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetMeContextOk() ([]MeContextSingle, bool) {
	if o == nil || isNil(o.MeContext) {
    return nil, false
	}
	return o.MeContext, true
}

// HasMeContext returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasMeContext() bool {
	if o != nil && !isNil(o.MeContext) {
		return true
	}

	return false
}

// SetMeContext gets a reference to the given []MeContextSingle and assigns it to the MeContext field.
func (o *SubNetworkSingle5) SetMeContext(v []MeContextSingle) {
	o.MeContext = v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || isNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || isNil(o.PerfMetricJob) {
    return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasPerfMetricJob() bool {
	if o != nil && !isNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *SubNetworkSingle5) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || isNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || isNil(o.ThresholdMonitor) {
    return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasThresholdMonitor() bool {
	if o != nil && !isNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *SubNetworkSingle5) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetTraceJob() []TraceJobSingle {
	if o == nil || isNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || isNil(o.TraceJob) {
    return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasTraceJob() bool {
	if o != nil && !isNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *SubNetworkSingle5) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetManagementDataCollection returns the ManagementDataCollection field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetManagementDataCollection() []ManagementDataCollectionSingle {
	if o == nil || isNil(o.ManagementDataCollection) {
		var ret []ManagementDataCollectionSingle
		return ret
	}
	return o.ManagementDataCollection
}

// GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetManagementDataCollectionOk() ([]ManagementDataCollectionSingle, bool) {
	if o == nil || isNil(o.ManagementDataCollection) {
    return nil, false
	}
	return o.ManagementDataCollection, true
}

// HasManagementDataCollection returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasManagementDataCollection() bool {
	if o != nil && !isNil(o.ManagementDataCollection) {
		return true
	}

	return false
}

// SetManagementDataCollection gets a reference to the given []ManagementDataCollectionSingle and assigns it to the ManagementDataCollection field.
func (o *SubNetworkSingle5) SetManagementDataCollection(v []ManagementDataCollectionSingle) {
	o.ManagementDataCollection = v
}

// GetNtfSubscriptionControl returns the NtfSubscriptionControl field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle {
	if o == nil || isNil(o.NtfSubscriptionControl) {
		var ret []NtfSubscriptionControlSingle
		return ret
	}
	return o.NtfSubscriptionControl
}

// GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetNtfSubscriptionControlOk() ([]NtfSubscriptionControlSingle, bool) {
	if o == nil || isNil(o.NtfSubscriptionControl) {
    return nil, false
	}
	return o.NtfSubscriptionControl, true
}

// HasNtfSubscriptionControl returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasNtfSubscriptionControl() bool {
	if o != nil && !isNil(o.NtfSubscriptionControl) {
		return true
	}

	return false
}

// SetNtfSubscriptionControl gets a reference to the given []NtfSubscriptionControlSingle and assigns it to the NtfSubscriptionControl field.
func (o *SubNetworkSingle5) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle) {
	o.NtfSubscriptionControl = v
}

// GetAlarmList returns the AlarmList field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetAlarmList() AlarmListSingle {
	if o == nil || isNil(o.AlarmList) {
		var ret AlarmListSingle
		return ret
	}
	return *o.AlarmList
}

// GetAlarmListOk returns a tuple with the AlarmList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetAlarmListOk() (*AlarmListSingle, bool) {
	if o == nil || isNil(o.AlarmList) {
    return nil, false
	}
	return o.AlarmList, true
}

// HasAlarmList returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasAlarmList() bool {
	if o != nil && !isNil(o.AlarmList) {
		return true
	}

	return false
}

// SetAlarmList gets a reference to the given AlarmListSingle and assigns it to the AlarmList field.
func (o *SubNetworkSingle5) SetAlarmList(v AlarmListSingle) {
	o.AlarmList = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetFiles() []FilesSingle {
	if o == nil || isNil(o.Files) {
		var ret []FilesSingle
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetFilesOk() ([]FilesSingle, bool) {
	if o == nil || isNil(o.Files) {
    return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasFiles() bool {
	if o != nil && !isNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FilesSingle and assigns it to the Files field.
func (o *SubNetworkSingle5) SetFiles(v []FilesSingle) {
	o.Files = v
}

// GetFileDownloadJob returns the FileDownloadJob field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetFileDownloadJob() []FileDownloadJobSingle {
	if o == nil || isNil(o.FileDownloadJob) {
		var ret []FileDownloadJobSingle
		return ret
	}
	return o.FileDownloadJob
}

// GetFileDownloadJobOk returns a tuple with the FileDownloadJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetFileDownloadJobOk() ([]FileDownloadJobSingle, bool) {
	if o == nil || isNil(o.FileDownloadJob) {
    return nil, false
	}
	return o.FileDownloadJob, true
}

// HasFileDownloadJob returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasFileDownloadJob() bool {
	if o != nil && !isNil(o.FileDownloadJob) {
		return true
	}

	return false
}

// SetFileDownloadJob gets a reference to the given []FileDownloadJobSingle and assigns it to the FileDownloadJob field.
func (o *SubNetworkSingle5) SetFileDownloadJob(v []FileDownloadJobSingle) {
	o.FileDownloadJob = v
}

// GetMnsRegistry returns the MnsRegistry field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetMnsRegistry() MnsRegistrySingle {
	if o == nil || isNil(o.MnsRegistry) {
		var ret MnsRegistrySingle
		return ret
	}
	return *o.MnsRegistry
}

// GetMnsRegistryOk returns a tuple with the MnsRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetMnsRegistryOk() (*MnsRegistrySingle, bool) {
	if o == nil || isNil(o.MnsRegistry) {
    return nil, false
	}
	return o.MnsRegistry, true
}

// HasMnsRegistry returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasMnsRegistry() bool {
	if o != nil && !isNil(o.MnsRegistry) {
		return true
	}

	return false
}

// SetMnsRegistry gets a reference to the given MnsRegistrySingle and assigns it to the MnsRegistry field.
func (o *SubNetworkSingle5) SetMnsRegistry(v MnsRegistrySingle) {
	o.MnsRegistry = &v
}

// GetSubNetwork returns the SubNetwork field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetSubNetwork() []SubNetworkSingle {
	if o == nil || isNil(o.SubNetwork) {
		var ret []SubNetworkSingle
		return ret
	}
	return o.SubNetwork
}

// GetSubNetworkOk returns a tuple with the SubNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetSubNetworkOk() ([]SubNetworkSingle, bool) {
	if o == nil || isNil(o.SubNetwork) {
    return nil, false
	}
	return o.SubNetwork, true
}

// HasSubNetwork returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasSubNetwork() bool {
	if o != nil && !isNil(o.SubNetwork) {
		return true
	}

	return false
}

// SetSubNetwork gets a reference to the given []SubNetworkSingle and assigns it to the SubNetwork field.
func (o *SubNetworkSingle5) SetSubNetwork(v []SubNetworkSingle) {
	o.SubNetwork = v
}

// GetManagedElement returns the ManagedElement field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetManagedElement() []ManagedElementSingle {
	if o == nil || isNil(o.ManagedElement) {
		var ret []ManagedElementSingle
		return ret
	}
	return o.ManagedElement
}

// GetManagedElementOk returns a tuple with the ManagedElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetManagedElementOk() ([]ManagedElementSingle, bool) {
	if o == nil || isNil(o.ManagedElement) {
    return nil, false
	}
	return o.ManagedElement, true
}

// HasManagedElement returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasManagedElement() bool {
	if o != nil && !isNil(o.ManagedElement) {
		return true
	}

	return false
}

// SetManagedElement gets a reference to the given []ManagedElementSingle and assigns it to the ManagedElement field.
func (o *SubNetworkSingle5) SetManagedElement(v []ManagedElementSingle) {
	o.ManagedElement = v
}

// GetMLTrainingFunction returns the MLTrainingFunction field value if set, zero value otherwise.
func (o *SubNetworkSingle5) GetMLTrainingFunction() []MLTrainingFunctionSingle {
	if o == nil || isNil(o.MLTrainingFunction) {
		var ret []MLTrainingFunctionSingle
		return ret
	}
	return o.MLTrainingFunction
}

// GetMLTrainingFunctionOk returns a tuple with the MLTrainingFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle5) GetMLTrainingFunctionOk() ([]MLTrainingFunctionSingle, bool) {
	if o == nil || isNil(o.MLTrainingFunction) {
    return nil, false
	}
	return o.MLTrainingFunction, true
}

// HasMLTrainingFunction returns a boolean if a field has been set.
func (o *SubNetworkSingle5) HasMLTrainingFunction() bool {
	if o != nil && !isNil(o.MLTrainingFunction) {
		return true
	}

	return false
}

// SetMLTrainingFunction gets a reference to the given []MLTrainingFunctionSingle and assigns it to the MLTrainingFunction field.
func (o *SubNetworkSingle5) SetMLTrainingFunction(v []MLTrainingFunctionSingle) {
	o.MLTrainingFunction = v
}

func (o SubNetworkSingle5) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.ManagementNode) {
		toSerialize["ManagementNode"] = o.ManagementNode
	}
	if !isNil(o.MnsAgent) {
		toSerialize["MnsAgent"] = o.MnsAgent
	}
	if !isNil(o.MeContext) {
		toSerialize["MeContext"] = o.MeContext
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
	if !isNil(o.ManagementDataCollection) {
		toSerialize["ManagementDataCollection"] = o.ManagementDataCollection
	}
	if !isNil(o.NtfSubscriptionControl) {
		toSerialize["NtfSubscriptionControl"] = o.NtfSubscriptionControl
	}
	if !isNil(o.AlarmList) {
		toSerialize["AlarmList"] = o.AlarmList
	}
	if !isNil(o.Files) {
		toSerialize["Files"] = o.Files
	}
	if !isNil(o.FileDownloadJob) {
		toSerialize["FileDownloadJob"] = o.FileDownloadJob
	}
	if !isNil(o.MnsRegistry) {
		toSerialize["MnsRegistry"] = o.MnsRegistry
	}
	if !isNil(o.SubNetwork) {
		toSerialize["SubNetwork"] = o.SubNetwork
	}
	if !isNil(o.ManagedElement) {
		toSerialize["ManagedElement"] = o.ManagedElement
	}
	if !isNil(o.MLTrainingFunction) {
		toSerialize["MLTrainingFunction"] = o.MLTrainingFunction
	}
	return json.Marshal(toSerialize)
}

type NullableSubNetworkSingle5 struct {
	value *SubNetworkSingle5
	isSet bool
}

func (v NullableSubNetworkSingle5) Get() *SubNetworkSingle5 {
	return v.value
}

func (v *NullableSubNetworkSingle5) Set(val *SubNetworkSingle5) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingle5) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingle5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingle5(val *SubNetworkSingle5) *NullableSubNetworkSingle5 {
	return &NullableSubNetworkSingle5{value: val, isSet: true}
}

func (v NullableSubNetworkSingle5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingle5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


