/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
)

// SubNetworkSingle struct for SubNetworkSingle
type SubNetworkSingle struct {
	Id NullableString `json:"id"`
	ObjectClass *string `json:"objectClass,omitempty"`
	ObjectInstance *string `json:"objectInstance,omitempty"`
	VsDataContainer []VsDataContainerSingle `json:"VsDataContainer,omitempty"`
	Attributes *SubNetworkAttr `json:"attributes,omitempty"`
	Subnetwork []SubNetworkSingle `json:"Subnetwork,omitempty"`
	ECSFunction []ECSFunctionSingle `json:"ECSFunction,omitempty"`
	EdgeDataNetwork []EdgeDataNetworkSingle `json:"EdgeDataNetwork,omitempty"`
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
}

// NewSubNetworkSingle instantiates a new SubNetworkSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubNetworkSingle(id NullableString) *SubNetworkSingle {
	this := SubNetworkSingle{}
	this.Id = id
	return &this
}

// NewSubNetworkSingleWithDefaults instantiates a new SubNetworkSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubNetworkSingleWithDefaults() *SubNetworkSingle {
	this := SubNetworkSingle{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubNetworkSingle) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubNetworkSingle) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *SubNetworkSingle) SetId(v string) {
	o.Id.Set(&v)
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetObjectClass() string {
	if o == nil || isNil(o.ObjectClass) {
		var ret string
		return ret
	}
	return *o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetObjectClassOk() (*string, bool) {
	if o == nil || isNil(o.ObjectClass) {
    return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasObjectClass() bool {
	if o != nil && !isNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given string and assigns it to the ObjectClass field.
func (o *SubNetworkSingle) SetObjectClass(v string) {
	o.ObjectClass = &v
}

// GetObjectInstance returns the ObjectInstance field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetObjectInstance() string {
	if o == nil || isNil(o.ObjectInstance) {
		var ret string
		return ret
	}
	return *o.ObjectInstance
}

// GetObjectInstanceOk returns a tuple with the ObjectInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetObjectInstanceOk() (*string, bool) {
	if o == nil || isNil(o.ObjectInstance) {
    return nil, false
	}
	return o.ObjectInstance, true
}

// HasObjectInstance returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasObjectInstance() bool {
	if o != nil && !isNil(o.ObjectInstance) {
		return true
	}

	return false
}

// SetObjectInstance gets a reference to the given string and assigns it to the ObjectInstance field.
func (o *SubNetworkSingle) SetObjectInstance(v string) {
	o.ObjectInstance = &v
}

// GetVsDataContainer returns the VsDataContainer field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetVsDataContainer() []VsDataContainerSingle {
	if o == nil || isNil(o.VsDataContainer) {
		var ret []VsDataContainerSingle
		return ret
	}
	return o.VsDataContainer
}

// GetVsDataContainerOk returns a tuple with the VsDataContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetVsDataContainerOk() ([]VsDataContainerSingle, bool) {
	if o == nil || isNil(o.VsDataContainer) {
    return nil, false
	}
	return o.VsDataContainer, true
}

// HasVsDataContainer returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasVsDataContainer() bool {
	if o != nil && !isNil(o.VsDataContainer) {
		return true
	}

	return false
}

// SetVsDataContainer gets a reference to the given []VsDataContainerSingle and assigns it to the VsDataContainer field.
func (o *SubNetworkSingle) SetVsDataContainer(v []VsDataContainerSingle) {
	o.VsDataContainer = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetAttributes() SubNetworkAttr {
	if o == nil || isNil(o.Attributes) {
		var ret SubNetworkAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetAttributesOk() (*SubNetworkAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubNetworkAttr and assigns it to the Attributes field.
func (o *SubNetworkSingle) SetAttributes(v SubNetworkAttr) {
	o.Attributes = &v
}

// GetSubnetwork returns the Subnetwork field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetSubnetwork() []SubNetworkSingle {
	if o == nil || isNil(o.Subnetwork) {
		var ret []SubNetworkSingle
		return ret
	}
	return o.Subnetwork
}

// GetSubnetworkOk returns a tuple with the Subnetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetSubnetworkOk() ([]SubNetworkSingle, bool) {
	if o == nil || isNil(o.Subnetwork) {
    return nil, false
	}
	return o.Subnetwork, true
}

// HasSubnetwork returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasSubnetwork() bool {
	if o != nil && !isNil(o.Subnetwork) {
		return true
	}

	return false
}

// SetSubnetwork gets a reference to the given []SubNetworkSingle and assigns it to the Subnetwork field.
func (o *SubNetworkSingle) SetSubnetwork(v []SubNetworkSingle) {
	o.Subnetwork = v
}

// GetECSFunction returns the ECSFunction field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetECSFunction() []ECSFunctionSingle {
	if o == nil || isNil(o.ECSFunction) {
		var ret []ECSFunctionSingle
		return ret
	}
	return o.ECSFunction
}

// GetECSFunctionOk returns a tuple with the ECSFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetECSFunctionOk() ([]ECSFunctionSingle, bool) {
	if o == nil || isNil(o.ECSFunction) {
    return nil, false
	}
	return o.ECSFunction, true
}

// HasECSFunction returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasECSFunction() bool {
	if o != nil && !isNil(o.ECSFunction) {
		return true
	}

	return false
}

// SetECSFunction gets a reference to the given []ECSFunctionSingle and assigns it to the ECSFunction field.
func (o *SubNetworkSingle) SetECSFunction(v []ECSFunctionSingle) {
	o.ECSFunction = v
}

// GetEdgeDataNetwork returns the EdgeDataNetwork field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetEdgeDataNetwork() []EdgeDataNetworkSingle {
	if o == nil || isNil(o.EdgeDataNetwork) {
		var ret []EdgeDataNetworkSingle
		return ret
	}
	return o.EdgeDataNetwork
}

// GetEdgeDataNetworkOk returns a tuple with the EdgeDataNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetEdgeDataNetworkOk() ([]EdgeDataNetworkSingle, bool) {
	if o == nil || isNil(o.EdgeDataNetwork) {
    return nil, false
	}
	return o.EdgeDataNetwork, true
}

// HasEdgeDataNetwork returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasEdgeDataNetwork() bool {
	if o != nil && !isNil(o.EdgeDataNetwork) {
		return true
	}

	return false
}

// SetEdgeDataNetwork gets a reference to the given []EdgeDataNetworkSingle and assigns it to the EdgeDataNetwork field.
func (o *SubNetworkSingle) SetEdgeDataNetwork(v []EdgeDataNetworkSingle) {
	o.EdgeDataNetwork = v
}

// GetManagementNode returns the ManagementNode field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetManagementNode() []ManagementNodeSingle {
	if o == nil || isNil(o.ManagementNode) {
		var ret []ManagementNodeSingle
		return ret
	}
	return o.ManagementNode
}

// GetManagementNodeOk returns a tuple with the ManagementNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetManagementNodeOk() ([]ManagementNodeSingle, bool) {
	if o == nil || isNil(o.ManagementNode) {
    return nil, false
	}
	return o.ManagementNode, true
}

// HasManagementNode returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasManagementNode() bool {
	if o != nil && !isNil(o.ManagementNode) {
		return true
	}

	return false
}

// SetManagementNode gets a reference to the given []ManagementNodeSingle and assigns it to the ManagementNode field.
func (o *SubNetworkSingle) SetManagementNode(v []ManagementNodeSingle) {
	o.ManagementNode = v
}

// GetMnsAgent returns the MnsAgent field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetMnsAgent() []MnsAgentSingle {
	if o == nil || isNil(o.MnsAgent) {
		var ret []MnsAgentSingle
		return ret
	}
	return o.MnsAgent
}

// GetMnsAgentOk returns a tuple with the MnsAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetMnsAgentOk() ([]MnsAgentSingle, bool) {
	if o == nil || isNil(o.MnsAgent) {
    return nil, false
	}
	return o.MnsAgent, true
}

// HasMnsAgent returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasMnsAgent() bool {
	if o != nil && !isNil(o.MnsAgent) {
		return true
	}

	return false
}

// SetMnsAgent gets a reference to the given []MnsAgentSingle and assigns it to the MnsAgent field.
func (o *SubNetworkSingle) SetMnsAgent(v []MnsAgentSingle) {
	o.MnsAgent = v
}

// GetMeContext returns the MeContext field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetMeContext() []MeContextSingle {
	if o == nil || isNil(o.MeContext) {
		var ret []MeContextSingle
		return ret
	}
	return o.MeContext
}

// GetMeContextOk returns a tuple with the MeContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetMeContextOk() ([]MeContextSingle, bool) {
	if o == nil || isNil(o.MeContext) {
    return nil, false
	}
	return o.MeContext, true
}

// HasMeContext returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasMeContext() bool {
	if o != nil && !isNil(o.MeContext) {
		return true
	}

	return false
}

// SetMeContext gets a reference to the given []MeContextSingle and assigns it to the MeContext field.
func (o *SubNetworkSingle) SetMeContext(v []MeContextSingle) {
	o.MeContext = v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || isNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || isNil(o.PerfMetricJob) {
    return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasPerfMetricJob() bool {
	if o != nil && !isNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *SubNetworkSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || isNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || isNil(o.ThresholdMonitor) {
    return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasThresholdMonitor() bool {
	if o != nil && !isNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *SubNetworkSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || isNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || isNil(o.TraceJob) {
    return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasTraceJob() bool {
	if o != nil && !isNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *SubNetworkSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetManagementDataCollection returns the ManagementDataCollection field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetManagementDataCollection() []ManagementDataCollectionSingle {
	if o == nil || isNil(o.ManagementDataCollection) {
		var ret []ManagementDataCollectionSingle
		return ret
	}
	return o.ManagementDataCollection
}

// GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetManagementDataCollectionOk() ([]ManagementDataCollectionSingle, bool) {
	if o == nil || isNil(o.ManagementDataCollection) {
    return nil, false
	}
	return o.ManagementDataCollection, true
}

// HasManagementDataCollection returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasManagementDataCollection() bool {
	if o != nil && !isNil(o.ManagementDataCollection) {
		return true
	}

	return false
}

// SetManagementDataCollection gets a reference to the given []ManagementDataCollectionSingle and assigns it to the ManagementDataCollection field.
func (o *SubNetworkSingle) SetManagementDataCollection(v []ManagementDataCollectionSingle) {
	o.ManagementDataCollection = v
}

// GetNtfSubscriptionControl returns the NtfSubscriptionControl field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle {
	if o == nil || isNil(o.NtfSubscriptionControl) {
		var ret []NtfSubscriptionControlSingle
		return ret
	}
	return o.NtfSubscriptionControl
}

// GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetNtfSubscriptionControlOk() ([]NtfSubscriptionControlSingle, bool) {
	if o == nil || isNil(o.NtfSubscriptionControl) {
    return nil, false
	}
	return o.NtfSubscriptionControl, true
}

// HasNtfSubscriptionControl returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasNtfSubscriptionControl() bool {
	if o != nil && !isNil(o.NtfSubscriptionControl) {
		return true
	}

	return false
}

// SetNtfSubscriptionControl gets a reference to the given []NtfSubscriptionControlSingle and assigns it to the NtfSubscriptionControl field.
func (o *SubNetworkSingle) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle) {
	o.NtfSubscriptionControl = v
}

// GetAlarmList returns the AlarmList field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetAlarmList() AlarmListSingle {
	if o == nil || isNil(o.AlarmList) {
		var ret AlarmListSingle
		return ret
	}
	return *o.AlarmList
}

// GetAlarmListOk returns a tuple with the AlarmList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetAlarmListOk() (*AlarmListSingle, bool) {
	if o == nil || isNil(o.AlarmList) {
    return nil, false
	}
	return o.AlarmList, true
}

// HasAlarmList returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasAlarmList() bool {
	if o != nil && !isNil(o.AlarmList) {
		return true
	}

	return false
}

// SetAlarmList gets a reference to the given AlarmListSingle and assigns it to the AlarmList field.
func (o *SubNetworkSingle) SetAlarmList(v AlarmListSingle) {
	o.AlarmList = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetFiles() []FilesSingle {
	if o == nil || isNil(o.Files) {
		var ret []FilesSingle
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetFilesOk() ([]FilesSingle, bool) {
	if o == nil || isNil(o.Files) {
    return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasFiles() bool {
	if o != nil && !isNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FilesSingle and assigns it to the Files field.
func (o *SubNetworkSingle) SetFiles(v []FilesSingle) {
	o.Files = v
}

// GetFileDownloadJob returns the FileDownloadJob field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetFileDownloadJob() []FileDownloadJobSingle {
	if o == nil || isNil(o.FileDownloadJob) {
		var ret []FileDownloadJobSingle
		return ret
	}
	return o.FileDownloadJob
}

// GetFileDownloadJobOk returns a tuple with the FileDownloadJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetFileDownloadJobOk() ([]FileDownloadJobSingle, bool) {
	if o == nil || isNil(o.FileDownloadJob) {
    return nil, false
	}
	return o.FileDownloadJob, true
}

// HasFileDownloadJob returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasFileDownloadJob() bool {
	if o != nil && !isNil(o.FileDownloadJob) {
		return true
	}

	return false
}

// SetFileDownloadJob gets a reference to the given []FileDownloadJobSingle and assigns it to the FileDownloadJob field.
func (o *SubNetworkSingle) SetFileDownloadJob(v []FileDownloadJobSingle) {
	o.FileDownloadJob = v
}

// GetMnsRegistry returns the MnsRegistry field value if set, zero value otherwise.
func (o *SubNetworkSingle) GetMnsRegistry() MnsRegistrySingle {
	if o == nil || isNil(o.MnsRegistry) {
		var ret MnsRegistrySingle
		return ret
	}
	return *o.MnsRegistry
}

// GetMnsRegistryOk returns a tuple with the MnsRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubNetworkSingle) GetMnsRegistryOk() (*MnsRegistrySingle, bool) {
	if o == nil || isNil(o.MnsRegistry) {
    return nil, false
	}
	return o.MnsRegistry, true
}

// HasMnsRegistry returns a boolean if a field has been set.
func (o *SubNetworkSingle) HasMnsRegistry() bool {
	if o != nil && !isNil(o.MnsRegistry) {
		return true
	}

	return false
}

// SetMnsRegistry gets a reference to the given MnsRegistrySingle and assigns it to the MnsRegistry field.
func (o *SubNetworkSingle) SetMnsRegistry(v MnsRegistrySingle) {
	o.MnsRegistry = &v
}

func (o SubNetworkSingle) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Subnetwork) {
		toSerialize["Subnetwork"] = o.Subnetwork
	}
	if !isNil(o.ECSFunction) {
		toSerialize["ECSFunction"] = o.ECSFunction
	}
	if !isNil(o.EdgeDataNetwork) {
		toSerialize["EdgeDataNetwork"] = o.EdgeDataNetwork
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
	return json.Marshal(toSerialize)
}

type NullableSubNetworkSingle struct {
	value *SubNetworkSingle
	isSet bool
}

func (v NullableSubNetworkSingle) Get() *SubNetworkSingle {
	return v.value
}

func (v *NullableSubNetworkSingle) Set(val *SubNetworkSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableSubNetworkSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableSubNetworkSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubNetworkSingle(val *SubNetworkSingle) *NullableSubNetworkSingle {
	return &NullableSubNetworkSingle{value: val, isSet: true}
}

func (v NullableSubNetworkSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubNetworkSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


