/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// SubNetworkSingle1 struct for SubNetworkSingle1
type SubNetworkSingle1 struct {
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
	ExternalAmfFunction []ExternalAmfFunctionSingle `json:"ExternalAmfFunction,omitempty"`
	ExternalNrfFunction []ExternalNrfFunctionSingle `json:"ExternalNrfFunction,omitempty"`
	ExternalNssfFunction []ExternalNssfFunctionSingle `json:"ExternalNssfFunction,omitempty"`
	AmfSet []AmfSetSingle `json:"AmfSet,omitempty"`
	AmfRegion []AmfRegionSingle `json:"AmfRegion,omitempty"`
	Configurable5QISet []Configurable5QISetSingle `json:"Configurable5QISet,omitempty"`
	Dynamic5QISet []Dynamic5QISetSingle `json:"Dynamic5QISet,omitempty"`
	EcmConnectionInfo []EcmConnectionInfoSingle `json:"EcmConnectionInfo,omitempty"`
}

// NewSubNetworkSingle1 instantiates a new SubNetworkSingle1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingle1(id NullableString) *SubNetworkSingle1 {
	this := SubNetworkSingle1{}
	this.Id = id
	return &this
}

// NewSubNetworkSingle1WithDefaults instantiates a new SubNetworkSingle1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingle1WithDefaults() *SubNetworkSingle1 {
	this := SubNetworkSingle1{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubNetworkSingle1) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubNetworkSingle1) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *SubNetworkSingle1) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
    return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *SubNetworkSingle1) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
    return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *SubNetworkSingle1) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
    return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *SubNetworkSingle1) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetAttributes() SubNetworkAttr {
	if o == nil || isNil(o.Attributes) {
		var ret SubNetworkAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetAttributesOk() (*SubNetworkAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubNetworkAttr and assigns it to the Attributes field.
func (o *SubNetworkSingle1) SetAttributes(v SubNetworkAttr) {
	o.Attributes = &v
}

// GetManagementNode returns the ManagementNode field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetManagementNode() []ManagementNodeSingle {
	if o == nil || isNil(o.ManagementNode) {
		var ret []ManagementNodeSingle
		return ret
	}
	return o.ManagementNode
}

// GetManagementNodeOk returns a tuple with the ManagementNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetManagementNodeOk() ([]ManagementNodeSingle, bool) {
	if o == nil || isNil(o.ManagementNode) {
    return nil, false
	}
	return o.ManagementNode, true
}

// HasManagementNode returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasManagementNode() bool {
	if o != nil && !isNil(o.ManagementNode) {
		return true
	}

	return false
}

// SetManagementNode gets a reference to the given []ManagementNodeSingle and assigns it to the ManagementNode field.
func (o *SubNetworkSingle1) SetManagementNode(v []ManagementNodeSingle) {
	o.ManagementNode = v
}

// GetMnsAgent returns the MnsAgent field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetMnsAgent() []MnsAgentSingle {
	if o == nil || isNil(o.MnsAgent) {
		var ret []MnsAgentSingle
		return ret
	}
	return o.MnsAgent
}

// GetMnsAgentOk returns a tuple with the MnsAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetMnsAgentOk() ([]MnsAgentSingle, bool) {
	if o == nil || isNil(o.MnsAgent) {
    return nil, false
	}
	return o.MnsAgent, true
}

// HasMnsAgent returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasMnsAgent() bool {
	if o != nil && !isNil(o.MnsAgent) {
		return true
	}

	return false
}

// SetMnsAgent gets a reference to the given []MnsAgentSingle and assigns it to the MnsAgent field.
func (o *SubNetworkSingle1) SetMnsAgent(v []MnsAgentSingle) {
	o.MnsAgent = v
}

// GetMeContext returns the MeContext field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetMeContext() []MeContextSingle {
	if o == nil || isNil(o.MeContext) {
		var ret []MeContextSingle
		return ret
	}
	return o.MeContext
}

// GetMeContextOk returns a tuple with the MeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetMeContextOk() ([]MeContextSingle, bool) {
	if o == nil || isNil(o.MeContext) {
    return nil, false
	}
	return o.MeContext, true
}

// HasMeContext returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasMeContext() bool {
	if o != nil && !isNil(o.MeContext) {
		return true
	}

	return false
}

// SetMeContext gets a reference to the given []MeContextSingle and assigns it to the MeContext field.
func (o *SubNetworkSingle1) SetMeContext(v []MeContextSingle) {
	o.MeContext = v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || isNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || isNil(o.PerfMetricJob) {
    return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasPerfMetricJob() bool {
	if o != nil && !isNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *SubNetworkSingle1) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || isNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || isNil(o.ThresholdMonitor) {
    return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasThresholdMonitor() bool {
	if o != nil && !isNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *SubNetworkSingle1) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetTraceJob() []TraceJobSingle {
	if o == nil || isNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || isNil(o.TraceJob) {
    return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasTraceJob() bool {
	if o != nil && !isNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *SubNetworkSingle1) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetManagementDataCollection returns the ManagementDataCollection field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetManagementDataCollection() []ManagementDataCollectionSingle {
	if o == nil || isNil(o.ManagementDataCollection) {
		var ret []ManagementDataCollectionSingle
		return ret
	}
	return o.ManagementDataCollection
}

// GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetManagementDataCollectionOk() ([]ManagementDataCollectionSingle, bool) {
	if o == nil || isNil(o.ManagementDataCollection) {
    return nil, false
	}
	return o.ManagementDataCollection, true
}

// HasManagementDataCollection returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasManagementDataCollection() bool {
	if o != nil && !isNil(o.ManagementDataCollection) {
		return true
	}

	return false
}

// SetManagementDataCollection gets a reference to the given []ManagementDataCollectionSingle and assigns it to the ManagementDataCollection field.
func (o *SubNetworkSingle1) SetManagementDataCollection(v []ManagementDataCollectionSingle) {
	o.ManagementDataCollection = v
}

// GetNtfSubscriptionControl returns the NtfSubscriptionControl field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle {
	if o == nil || isNil(o.NtfSubscriptionControl) {
		var ret []NtfSubscriptionControlSingle
		return ret
	}
	return o.NtfSubscriptionControl
}

// GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetNtfSubscriptionControlOk() ([]NtfSubscriptionControlSingle, bool) {
	if o == nil || isNil(o.NtfSubscriptionControl) {
    return nil, false
	}
	return o.NtfSubscriptionControl, true
}

// HasNtfSubscriptionControl returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasNtfSubscriptionControl() bool {
	if o != nil && !isNil(o.NtfSubscriptionControl) {
		return true
	}

	return false
}

// SetNtfSubscriptionControl gets a reference to the given []NtfSubscriptionControlSingle and assigns it to the NtfSubscriptionControl field.
func (o *SubNetworkSingle1) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle) {
	o.NtfSubscriptionControl = v
}

// GetAlarmList returns the AlarmList field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetAlarmList() AlarmListSingle {
	if o == nil || isNil(o.AlarmList) {
		var ret AlarmListSingle
		return ret
	}
	return *o.AlarmList
}

// GetAlarmListOk returns a tuple with the AlarmList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetAlarmListOk() (*AlarmListSingle, bool) {
	if o == nil || isNil(o.AlarmList) {
    return nil, false
	}
	return o.AlarmList, true
}

// HasAlarmList returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasAlarmList() bool {
	if o != nil && !isNil(o.AlarmList) {
		return true
	}

	return false
}

// SetAlarmList gets a reference to the given AlarmListSingle and assigns it to the AlarmList field.
func (o *SubNetworkSingle1) SetAlarmList(v AlarmListSingle) {
	o.AlarmList = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetFiles() []FilesSingle {
	if o == nil || isNil(o.Files) {
		var ret []FilesSingle
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetFilesOk() ([]FilesSingle, bool) {
	if o == nil || isNil(o.Files) {
    return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasFiles() bool {
	if o != nil && !isNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FilesSingle and assigns it to the Files field.
func (o *SubNetworkSingle1) SetFiles(v []FilesSingle) {
	o.Files = v
}

// GetFileDownloadJob returns the FileDownloadJob field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetFileDownloadJob() []FileDownloadJobSingle {
	if o == nil || isNil(o.FileDownloadJob) {
		var ret []FileDownloadJobSingle
		return ret
	}
	return o.FileDownloadJob
}

// GetFileDownloadJobOk returns a tuple with the FileDownloadJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetFileDownloadJobOk() ([]FileDownloadJobSingle, bool) {
	if o == nil || isNil(o.FileDownloadJob) {
    return nil, false
	}
	return o.FileDownloadJob, true
}

// HasFileDownloadJob returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasFileDownloadJob() bool {
	if o != nil && !isNil(o.FileDownloadJob) {
		return true
	}

	return false
}

// SetFileDownloadJob gets a reference to the given []FileDownloadJobSingle and assigns it to the FileDownloadJob field.
func (o *SubNetworkSingle1) SetFileDownloadJob(v []FileDownloadJobSingle) {
	o.FileDownloadJob = v
}

// GetMnsRegistry returns the MnsRegistry field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetMnsRegistry() MnsRegistrySingle {
	if o == nil || isNil(o.MnsRegistry) {
		var ret MnsRegistrySingle
		return ret
	}
	return *o.MnsRegistry
}

// GetMnsRegistryOk returns a tuple with the MnsRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetMnsRegistryOk() (*MnsRegistrySingle, bool) {
	if o == nil || isNil(o.MnsRegistry) {
    return nil, false
	}
	return o.MnsRegistry, true
}

// HasMnsRegistry returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasMnsRegistry() bool {
	if o != nil && !isNil(o.MnsRegistry) {
		return true
	}

	return false
}

// SetMnsRegistry gets a reference to the given MnsRegistrySingle and assigns it to the MnsRegistry field.
func (o *SubNetworkSingle1) SetMnsRegistry(v MnsRegistrySingle) {
	o.MnsRegistry = &v
}

// GetSubNetwork returns the SubNetwork field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetSubNetwork() []SubNetworkSingle {
	if o == nil || isNil(o.SubNetwork) {
		var ret []SubNetworkSingle
		return ret
	}
	return o.SubNetwork
}

// GetSubNetworkOk returns a tuple with the SubNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetSubNetworkOk() ([]SubNetworkSingle, bool) {
	if o == nil || isNil(o.SubNetwork) {
    return nil, false
	}
	return o.SubNetwork, true
}

// HasSubNetwork returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasSubNetwork() bool {
	if o != nil && !isNil(o.SubNetwork) {
		return true
	}

	return false
}

// SetSubNetwork gets a reference to the given []SubNetworkSingle and assigns it to the SubNetwork field.
func (o *SubNetworkSingle1) SetSubNetwork(v []SubNetworkSingle) {
	o.SubNetwork = v
}

// GetManagedElement returns the ManagedElement field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetManagedElement() []ManagedElementSingle {
	if o == nil || isNil(o.ManagedElement) {
		var ret []ManagedElementSingle
		return ret
	}
	return o.ManagedElement
}

// GetManagedElementOk returns a tuple with the ManagedElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetManagedElementOk() ([]ManagedElementSingle, bool) {
	if o == nil || isNil(o.ManagedElement) {
    return nil, false
	}
	return o.ManagedElement, true
}

// HasManagedElement returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasManagedElement() bool {
	if o != nil && !isNil(o.ManagedElement) {
		return true
	}

	return false
}

// SetManagedElement gets a reference to the given []ManagedElementSingle and assigns it to the ManagedElement field.
func (o *SubNetworkSingle1) SetManagedElement(v []ManagedElementSingle) {
	o.ManagedElement = v
}

// GetExternalAmfFunction returns the ExternalAmfFunction field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetExternalAmfFunction() []ExternalAmfFunctionSingle {
	if o == nil || isNil(o.ExternalAmfFunction) {
		var ret []ExternalAmfFunctionSingle
		return ret
	}
	return o.ExternalAmfFunction
}

// GetExternalAmfFunctionOk returns a tuple with the ExternalAmfFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetExternalAmfFunctionOk() ([]ExternalAmfFunctionSingle, bool) {
	if o == nil || isNil(o.ExternalAmfFunction) {
    return nil, false
	}
	return o.ExternalAmfFunction, true
}

// HasExternalAmfFunction returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasExternalAmfFunction() bool {
	if o != nil && !isNil(o.ExternalAmfFunction) {
		return true
	}

	return false
}

// SetExternalAmfFunction gets a reference to the given []ExternalAmfFunctionSingle and assigns it to the ExternalAmfFunction field.
func (o *SubNetworkSingle1) SetExternalAmfFunction(v []ExternalAmfFunctionSingle) {
	o.ExternalAmfFunction = v
}

// GetExternalNrfFunction returns the ExternalNrfFunction field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetExternalNrfFunction() []ExternalNrfFunctionSingle {
	if o == nil || isNil(o.ExternalNrfFunction) {
		var ret []ExternalNrfFunctionSingle
		return ret
	}
	return o.ExternalNrfFunction
}

// GetExternalNrfFunctionOk returns a tuple with the ExternalNrfFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetExternalNrfFunctionOk() ([]ExternalNrfFunctionSingle, bool) {
	if o == nil || isNil(o.ExternalNrfFunction) {
    return nil, false
	}
	return o.ExternalNrfFunction, true
}

// HasExternalNrfFunction returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasExternalNrfFunction() bool {
	if o != nil && !isNil(o.ExternalNrfFunction) {
		return true
	}

	return false
}

// SetExternalNrfFunction gets a reference to the given []ExternalNrfFunctionSingle and assigns it to the ExternalNrfFunction field.
func (o *SubNetworkSingle1) SetExternalNrfFunction(v []ExternalNrfFunctionSingle) {
	o.ExternalNrfFunction = v
}

// GetExternalNssfFunction returns the ExternalNssfFunction field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetExternalNssfFunction() []ExternalNssfFunctionSingle {
	if o == nil || isNil(o.ExternalNssfFunction) {
		var ret []ExternalNssfFunctionSingle
		return ret
	}
	return o.ExternalNssfFunction
}

// GetExternalNssfFunctionOk returns a tuple with the ExternalNssfFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetExternalNssfFunctionOk() ([]ExternalNssfFunctionSingle, bool) {
	if o == nil || isNil(o.ExternalNssfFunction) {
    return nil, false
	}
	return o.ExternalNssfFunction, true
}

// HasExternalNssfFunction returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasExternalNssfFunction() bool {
	if o != nil && !isNil(o.ExternalNssfFunction) {
		return true
	}

	return false
}

// SetExternalNssfFunction gets a reference to the given []ExternalNssfFunctionSingle and assigns it to the ExternalNssfFunction field.
func (o *SubNetworkSingle1) SetExternalNssfFunction(v []ExternalNssfFunctionSingle) {
	o.ExternalNssfFunction = v
}

// GetAmfSet returns the AmfSet field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetAmfSet() []AmfSetSingle {
	if o == nil || isNil(o.AmfSet) {
		var ret []AmfSetSingle
		return ret
	}
	return o.AmfSet
}

// GetAmfSetOk returns a tuple with the AmfSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetAmfSetOk() ([]AmfSetSingle, bool) {
	if o == nil || isNil(o.AmfSet) {
    return nil, false
	}
	return o.AmfSet, true
}

// HasAmfSet returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasAmfSet() bool {
	if o != nil && !isNil(o.AmfSet) {
		return true
	}

	return false
}

// SetAmfSet gets a reference to the given []AmfSetSingle and assigns it to the AmfSet field.
func (o *SubNetworkSingle1) SetAmfSet(v []AmfSetSingle) {
	o.AmfSet = v
}

// GetAmfRegion returns the AmfRegion field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetAmfRegion() []AmfRegionSingle {
	if o == nil || isNil(o.AmfRegion) {
		var ret []AmfRegionSingle
		return ret
	}
	return o.AmfRegion
}

// GetAmfRegionOk returns a tuple with the AmfRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetAmfRegionOk() ([]AmfRegionSingle, bool) {
	if o == nil || isNil(o.AmfRegion) {
    return nil, false
	}
	return o.AmfRegion, true
}

// HasAmfRegion returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasAmfRegion() bool {
	if o != nil && !isNil(o.AmfRegion) {
		return true
	}

	return false
}

// SetAmfRegion gets a reference to the given []AmfRegionSingle and assigns it to the AmfRegion field.
func (o *SubNetworkSingle1) SetAmfRegion(v []AmfRegionSingle) {
	o.AmfRegion = v
}

// GetConfigurable5QISet returns the Configurable5QISet field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetConfigurable5QISet() []Configurable5QISetSingle {
	if o == nil || isNil(o.Configurable5QISet) {
		var ret []Configurable5QISetSingle
		return ret
	}
	return o.Configurable5QISet
}

// GetConfigurable5QISetOk returns a tuple with the Configurable5QISet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetConfigurable5QISetOk() ([]Configurable5QISetSingle, bool) {
	if o == nil || isNil(o.Configurable5QISet) {
    return nil, false
	}
	return o.Configurable5QISet, true
}

// HasConfigurable5QISet returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasConfigurable5QISet() bool {
	if o != nil && !isNil(o.Configurable5QISet) {
		return true
	}

	return false
}

// SetConfigurable5QISet gets a reference to the given []Configurable5QISetSingle and assigns it to the Configurable5QISet field.
func (o *SubNetworkSingle1) SetConfigurable5QISet(v []Configurable5QISetSingle) {
	o.Configurable5QISet = v
}

// GetDynamic5QISet returns the Dynamic5QISet field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetDynamic5QISet() []Dynamic5QISetSingle {
	if o == nil || isNil(o.Dynamic5QISet) {
		var ret []Dynamic5QISetSingle
		return ret
	}
	return o.Dynamic5QISet
}

// GetDynamic5QISetOk returns a tuple with the Dynamic5QISet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetDynamic5QISetOk() ([]Dynamic5QISetSingle, bool) {
	if o == nil || isNil(o.Dynamic5QISet) {
    return nil, false
	}
	return o.Dynamic5QISet, true
}

// HasDynamic5QISet returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasDynamic5QISet() bool {
	if o != nil && !isNil(o.Dynamic5QISet) {
		return true
	}

	return false
}

// SetDynamic5QISet gets a reference to the given []Dynamic5QISetSingle and assigns it to the Dynamic5QISet field.
func (o *SubNetworkSingle1) SetDynamic5QISet(v []Dynamic5QISetSingle) {
	o.Dynamic5QISet = v
}

// GetEcmConnectionInfo returns the EcmConnectionInfo field value if set, zero value otherwise.
func (o *SubNetworkSingle1) GetEcmConnectionInfo() []EcmConnectionInfoSingle {
	if o == nil || isNil(o.EcmConnectionInfo) {
		var ret []EcmConnectionInfoSingle
		return ret
	}
	return o.EcmConnectionInfo
}

// GetEcmConnectionInfoOk returns a tuple with the EcmConnectionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle1) GetEcmConnectionInfoOk() ([]EcmConnectionInfoSingle, bool) {
	if o == nil || isNil(o.EcmConnectionInfo) {
    return nil, false
	}
	return o.EcmConnectionInfo, true
}

// HasEcmConnectionInfo returns a boolean if a field has been set.
func (o *SubNetworkSingle1) HasEcmConnectionInfo() bool {
	if o != nil && !isNil(o.EcmConnectionInfo) {
		return true
	}

	return false
}

// SetEcmConnectionInfo gets a reference to the given []EcmConnectionInfoSingle and assigns it to the EcmConnectionInfo field.
func (o *SubNetworkSingle1) SetEcmConnectionInfo(v []EcmConnectionInfoSingle) {
	o.EcmConnectionInfo = v
}

func (o SubNetworkSingle1) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.ExternalAmfFunction) {
		toSerialize["ExternalAmfFunction"] = o.ExternalAmfFunction
	}
	if !isNil(o.ExternalNrfFunction) {
		toSerialize["ExternalNrfFunction"] = o.ExternalNrfFunction
	}
	if !isNil(o.ExternalNssfFunction) {
		toSerialize["ExternalNssfFunction"] = o.ExternalNssfFunction
	}
	if !isNil(o.AmfSet) {
		toSerialize["AmfSet"] = o.AmfSet
	}
	if !isNil(o.AmfRegion) {
		toSerialize["AmfRegion"] = o.AmfRegion
	}
	if !isNil(o.Configurable5QISet) {
		toSerialize["Configurable5QISet"] = o.Configurable5QISet
	}
	if !isNil(o.Dynamic5QISet) {
		toSerialize["Dynamic5QISet"] = o.Dynamic5QISet
	}
	if !isNil(o.EcmConnectionInfo) {
		toSerialize["EcmConnectionInfo"] = o.EcmConnectionInfo
	}
	return json.Marshal(toSerialize)
}

type NullableSubNetworkSingle1 struct {
	value *SubNetworkSingle1
	isSet bool
}

func (v NullableSubNetworkSingle1) Get() *SubNetworkSingle1 {
	return v.value
}

func (v *NullableSubNetworkSingle1) Set(val *SubNetworkSingle1) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingle1) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingle1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingle1(val *SubNetworkSingle1) *NullableSubNetworkSingle1 {
	return &NullableSubNetworkSingle1{value: val, isSet: true}
}

func (v NullableSubNetworkSingle1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingle1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


