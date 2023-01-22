# ResourcesEdgeNrm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubNetwork** | Pointer to [**[]SubNetworkSingle**](SubNetworkSingle.md) |  | [optional] 
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**ManagedFunctionAttr**](ManagedFunction-Attr.md) |  | [optional] 
**Subnetwork** | Pointer to [**[]SubNetworkSingle**](SubNetworkSingle.md) |  | [optional] 
**ECSFunction** | Pointer to [**[]ECSFunctionSingle**](ECSFunctionSingle.md) |  | [optional] 
**EdgeDataNetwork** | Pointer to [**[]EdgeDataNetworkSingle**](EdgeDataNetworkSingle.md) |  | [optional] 
**ManagementNode** | Pointer to [**[]ManagementNodeSingle**](ManagementNodeSingle.md) |  | [optional] 
**MnsAgent** | Pointer to [**[]MnsAgentSingle**](MnsAgentSingle.md) |  | [optional] 
**MeContext** | Pointer to [**[]MeContextSingle**](MeContextSingle.md) |  | [optional] 
**PerfMetricJob** | Pointer to [**[]PerfMetricJobSingle**](PerfMetricJobSingle.md) |  | [optional] 
**ThresholdMonitor** | Pointer to [**[]ThresholdMonitorSingle**](ThresholdMonitorSingle.md) |  | [optional] 
**TraceJob** | Pointer to [**[]TraceJobSingle**](TraceJobSingle.md) |  | [optional] 
**ManagementDataCollection** | Pointer to [**[]ManagementDataCollectionSingle**](ManagementDataCollectionSingle.md) |  | [optional] 
**NtfSubscriptionControl** | Pointer to [**[]NtfSubscriptionControlSingle**](NtfSubscriptionControlSingle.md) |  | [optional] 
**AlarmList** | Pointer to [**AlarmListSingle**](AlarmListSingle.md) |  | [optional] 
**Files** | Pointer to [**[]FilesSingle**](FilesSingle.md) |  | [optional] 
**FileDownloadJob** | Pointer to [**[]FileDownloadJobSingle**](FileDownloadJobSingle.md) |  | [optional] 
**MnsRegistry** | Pointer to [**MnsRegistrySingle**](MnsRegistrySingle.md) |  | [optional] 
**ManagedNFService** | Pointer to [**[]ManagedNFServiceSingle**](ManagedNFServiceSingle.md) |  | [optional] 
**EdnIdentifier** | Pointer to **string** |  | [optional] 
**EDNConnectionInfo** | Pointer to [**EDNConnectionInfo**](EDNConnectionInfo.md) |  | [optional] 
**EASFunction** | Pointer to [**[]EASFunctionSingle**](EASFunctionSingle.md) |  | [optional] 
**EESFunction** | Pointer to [**[]EESFunctionSingle**](EESFunctionSingle.md) |  | [optional] 
**RequiredEASservingLocation** | Pointer to [**ServingLocation**](ServingLocation.md) |  | [optional] 
**AffinityAntiAffinity** | Pointer to [**AffinityAntiAffinity**](AffinityAntiAffinity.md) |  | [optional] 
**ServiceContinuity** | Pointer to **bool** |  | [optional] 
**VirtualResource** | Pointer to [**VirtualResource**](VirtualResource.md) |  | [optional] 
**SoftwareImageInfo** | Pointer to [**SoftwareImageInfo**](SoftwareImageInfo.md) |  | [optional] 

## Methods

### NewResourcesEdgeNrm

`func NewResourcesEdgeNrm(id NullableString, ) *ResourcesEdgeNrm`

NewResourcesEdgeNrm instantiates a new ResourcesEdgeNrm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesEdgeNrmWithDefaults

`func NewResourcesEdgeNrmWithDefaults() *ResourcesEdgeNrm`

NewResourcesEdgeNrmWithDefaults instantiates a new ResourcesEdgeNrm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubNetwork

`func (o *ResourcesEdgeNrm) GetSubNetwork() []SubNetworkSingle`

GetSubNetwork returns the SubNetwork field if non-nil, zero value otherwise.

### GetSubNetworkOk

`func (o *ResourcesEdgeNrm) GetSubNetworkOk() (*[]SubNetworkSingle, bool)`

GetSubNetworkOk returns a tuple with the SubNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetwork

`func (o *ResourcesEdgeNrm) SetSubNetwork(v []SubNetworkSingle)`

SetSubNetwork sets SubNetwork field to given value.

### HasSubNetwork

`func (o *ResourcesEdgeNrm) HasSubNetwork() bool`

HasSubNetwork returns a boolean if a field has been set.

### GetId

`func (o *ResourcesEdgeNrm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcesEdgeNrm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcesEdgeNrm) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ResourcesEdgeNrm) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ResourcesEdgeNrm) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *ResourcesEdgeNrm) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ResourcesEdgeNrm) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ResourcesEdgeNrm) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ResourcesEdgeNrm) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *ResourcesEdgeNrm) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *ResourcesEdgeNrm) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *ResourcesEdgeNrm) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *ResourcesEdgeNrm) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *ResourcesEdgeNrm) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *ResourcesEdgeNrm) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *ResourcesEdgeNrm) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *ResourcesEdgeNrm) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *ResourcesEdgeNrm) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ResourcesEdgeNrm) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ResourcesEdgeNrm) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ResourcesEdgeNrm) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetSubnetwork

`func (o *ResourcesEdgeNrm) GetSubnetwork() []SubNetworkSingle`

GetSubnetwork returns the Subnetwork field if non-nil, zero value otherwise.

### GetSubnetworkOk

`func (o *ResourcesEdgeNrm) GetSubnetworkOk() (*[]SubNetworkSingle, bool)`

GetSubnetworkOk returns a tuple with the Subnetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetwork

`func (o *ResourcesEdgeNrm) SetSubnetwork(v []SubNetworkSingle)`

SetSubnetwork sets Subnetwork field to given value.

### HasSubnetwork

`func (o *ResourcesEdgeNrm) HasSubnetwork() bool`

HasSubnetwork returns a boolean if a field has been set.

### GetECSFunction

`func (o *ResourcesEdgeNrm) GetECSFunction() []ECSFunctionSingle`

GetECSFunction returns the ECSFunction field if non-nil, zero value otherwise.

### GetECSFunctionOk

`func (o *ResourcesEdgeNrm) GetECSFunctionOk() (*[]ECSFunctionSingle, bool)`

GetECSFunctionOk returns a tuple with the ECSFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECSFunction

`func (o *ResourcesEdgeNrm) SetECSFunction(v []ECSFunctionSingle)`

SetECSFunction sets ECSFunction field to given value.

### HasECSFunction

`func (o *ResourcesEdgeNrm) HasECSFunction() bool`

HasECSFunction returns a boolean if a field has been set.

### GetEdgeDataNetwork

`func (o *ResourcesEdgeNrm) GetEdgeDataNetwork() []EdgeDataNetworkSingle`

GetEdgeDataNetwork returns the EdgeDataNetwork field if non-nil, zero value otherwise.

### GetEdgeDataNetworkOk

`func (o *ResourcesEdgeNrm) GetEdgeDataNetworkOk() (*[]EdgeDataNetworkSingle, bool)`

GetEdgeDataNetworkOk returns a tuple with the EdgeDataNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeDataNetwork

`func (o *ResourcesEdgeNrm) SetEdgeDataNetwork(v []EdgeDataNetworkSingle)`

SetEdgeDataNetwork sets EdgeDataNetwork field to given value.

### HasEdgeDataNetwork

`func (o *ResourcesEdgeNrm) HasEdgeDataNetwork() bool`

HasEdgeDataNetwork returns a boolean if a field has been set.

### GetManagementNode

`func (o *ResourcesEdgeNrm) GetManagementNode() []ManagementNodeSingle`

GetManagementNode returns the ManagementNode field if non-nil, zero value otherwise.

### GetManagementNodeOk

`func (o *ResourcesEdgeNrm) GetManagementNodeOk() (*[]ManagementNodeSingle, bool)`

GetManagementNodeOk returns a tuple with the ManagementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNode

`func (o *ResourcesEdgeNrm) SetManagementNode(v []ManagementNodeSingle)`

SetManagementNode sets ManagementNode field to given value.

### HasManagementNode

`func (o *ResourcesEdgeNrm) HasManagementNode() bool`

HasManagementNode returns a boolean if a field has been set.

### GetMnsAgent

`func (o *ResourcesEdgeNrm) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *ResourcesEdgeNrm) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *ResourcesEdgeNrm) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *ResourcesEdgeNrm) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetMeContext

`func (o *ResourcesEdgeNrm) GetMeContext() []MeContextSingle`

GetMeContext returns the MeContext field if non-nil, zero value otherwise.

### GetMeContextOk

`func (o *ResourcesEdgeNrm) GetMeContextOk() (*[]MeContextSingle, bool)`

GetMeContextOk returns a tuple with the MeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeContext

`func (o *ResourcesEdgeNrm) SetMeContext(v []MeContextSingle)`

SetMeContext sets MeContext field to given value.

### HasMeContext

`func (o *ResourcesEdgeNrm) HasMeContext() bool`

HasMeContext returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *ResourcesEdgeNrm) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *ResourcesEdgeNrm) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *ResourcesEdgeNrm) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *ResourcesEdgeNrm) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *ResourcesEdgeNrm) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *ResourcesEdgeNrm) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *ResourcesEdgeNrm) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *ResourcesEdgeNrm) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetTraceJob

`func (o *ResourcesEdgeNrm) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *ResourcesEdgeNrm) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *ResourcesEdgeNrm) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *ResourcesEdgeNrm) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetManagementDataCollection

`func (o *ResourcesEdgeNrm) GetManagementDataCollection() []ManagementDataCollectionSingle`

GetManagementDataCollection returns the ManagementDataCollection field if non-nil, zero value otherwise.

### GetManagementDataCollectionOk

`func (o *ResourcesEdgeNrm) GetManagementDataCollectionOk() (*[]ManagementDataCollectionSingle, bool)`

GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementDataCollection

`func (o *ResourcesEdgeNrm) SetManagementDataCollection(v []ManagementDataCollectionSingle)`

SetManagementDataCollection sets ManagementDataCollection field to given value.

### HasManagementDataCollection

`func (o *ResourcesEdgeNrm) HasManagementDataCollection() bool`

HasManagementDataCollection returns a boolean if a field has been set.

### GetNtfSubscriptionControl

`func (o *ResourcesEdgeNrm) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle`

GetNtfSubscriptionControl returns the NtfSubscriptionControl field if non-nil, zero value otherwise.

### GetNtfSubscriptionControlOk

`func (o *ResourcesEdgeNrm) GetNtfSubscriptionControlOk() (*[]NtfSubscriptionControlSingle, bool)`

GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfSubscriptionControl

`func (o *ResourcesEdgeNrm) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle)`

SetNtfSubscriptionControl sets NtfSubscriptionControl field to given value.

### HasNtfSubscriptionControl

`func (o *ResourcesEdgeNrm) HasNtfSubscriptionControl() bool`

HasNtfSubscriptionControl returns a boolean if a field has been set.

### GetAlarmList

`func (o *ResourcesEdgeNrm) GetAlarmList() AlarmListSingle`

GetAlarmList returns the AlarmList field if non-nil, zero value otherwise.

### GetAlarmListOk

`func (o *ResourcesEdgeNrm) GetAlarmListOk() (*AlarmListSingle, bool)`

GetAlarmListOk returns a tuple with the AlarmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmList

`func (o *ResourcesEdgeNrm) SetAlarmList(v AlarmListSingle)`

SetAlarmList sets AlarmList field to given value.

### HasAlarmList

`func (o *ResourcesEdgeNrm) HasAlarmList() bool`

HasAlarmList returns a boolean if a field has been set.

### GetFiles

`func (o *ResourcesEdgeNrm) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ResourcesEdgeNrm) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ResourcesEdgeNrm) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ResourcesEdgeNrm) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFileDownloadJob

`func (o *ResourcesEdgeNrm) GetFileDownloadJob() []FileDownloadJobSingle`

GetFileDownloadJob returns the FileDownloadJob field if non-nil, zero value otherwise.

### GetFileDownloadJobOk

`func (o *ResourcesEdgeNrm) GetFileDownloadJobOk() (*[]FileDownloadJobSingle, bool)`

GetFileDownloadJobOk returns a tuple with the FileDownloadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadJob

`func (o *ResourcesEdgeNrm) SetFileDownloadJob(v []FileDownloadJobSingle)`

SetFileDownloadJob sets FileDownloadJob field to given value.

### HasFileDownloadJob

`func (o *ResourcesEdgeNrm) HasFileDownloadJob() bool`

HasFileDownloadJob returns a boolean if a field has been set.

### GetMnsRegistry

`func (o *ResourcesEdgeNrm) GetMnsRegistry() MnsRegistrySingle`

GetMnsRegistry returns the MnsRegistry field if non-nil, zero value otherwise.

### GetMnsRegistryOk

`func (o *ResourcesEdgeNrm) GetMnsRegistryOk() (*MnsRegistrySingle, bool)`

GetMnsRegistryOk returns a tuple with the MnsRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsRegistry

`func (o *ResourcesEdgeNrm) SetMnsRegistry(v MnsRegistrySingle)`

SetMnsRegistry sets MnsRegistry field to given value.

### HasMnsRegistry

`func (o *ResourcesEdgeNrm) HasMnsRegistry() bool`

HasMnsRegistry returns a boolean if a field has been set.

### GetManagedNFService

`func (o *ResourcesEdgeNrm) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *ResourcesEdgeNrm) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *ResourcesEdgeNrm) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *ResourcesEdgeNrm) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetEdnIdentifier

`func (o *ResourcesEdgeNrm) GetEdnIdentifier() string`

GetEdnIdentifier returns the EdnIdentifier field if non-nil, zero value otherwise.

### GetEdnIdentifierOk

`func (o *ResourcesEdgeNrm) GetEdnIdentifierOk() (*string, bool)`

GetEdnIdentifierOk returns a tuple with the EdnIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnIdentifier

`func (o *ResourcesEdgeNrm) SetEdnIdentifier(v string)`

SetEdnIdentifier sets EdnIdentifier field to given value.

### HasEdnIdentifier

`func (o *ResourcesEdgeNrm) HasEdnIdentifier() bool`

HasEdnIdentifier returns a boolean if a field has been set.

### GetEDNConnectionInfo

`func (o *ResourcesEdgeNrm) GetEDNConnectionInfo() EDNConnectionInfo`

GetEDNConnectionInfo returns the EDNConnectionInfo field if non-nil, zero value otherwise.

### GetEDNConnectionInfoOk

`func (o *ResourcesEdgeNrm) GetEDNConnectionInfoOk() (*EDNConnectionInfo, bool)`

GetEDNConnectionInfoOk returns a tuple with the EDNConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDNConnectionInfo

`func (o *ResourcesEdgeNrm) SetEDNConnectionInfo(v EDNConnectionInfo)`

SetEDNConnectionInfo sets EDNConnectionInfo field to given value.

### HasEDNConnectionInfo

`func (o *ResourcesEdgeNrm) HasEDNConnectionInfo() bool`

HasEDNConnectionInfo returns a boolean if a field has been set.

### GetEASFunction

`func (o *ResourcesEdgeNrm) GetEASFunction() []EASFunctionSingle`

GetEASFunction returns the EASFunction field if non-nil, zero value otherwise.

### GetEASFunctionOk

`func (o *ResourcesEdgeNrm) GetEASFunctionOk() (*[]EASFunctionSingle, bool)`

GetEASFunctionOk returns a tuple with the EASFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEASFunction

`func (o *ResourcesEdgeNrm) SetEASFunction(v []EASFunctionSingle)`

SetEASFunction sets EASFunction field to given value.

### HasEASFunction

`func (o *ResourcesEdgeNrm) HasEASFunction() bool`

HasEASFunction returns a boolean if a field has been set.

### GetEESFunction

`func (o *ResourcesEdgeNrm) GetEESFunction() []EESFunctionSingle`

GetEESFunction returns the EESFunction field if non-nil, zero value otherwise.

### GetEESFunctionOk

`func (o *ResourcesEdgeNrm) GetEESFunctionOk() (*[]EESFunctionSingle, bool)`

GetEESFunctionOk returns a tuple with the EESFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEESFunction

`func (o *ResourcesEdgeNrm) SetEESFunction(v []EESFunctionSingle)`

SetEESFunction sets EESFunction field to given value.

### HasEESFunction

`func (o *ResourcesEdgeNrm) HasEESFunction() bool`

HasEESFunction returns a boolean if a field has been set.

### GetRequiredEASservingLocation

`func (o *ResourcesEdgeNrm) GetRequiredEASservingLocation() ServingLocation`

GetRequiredEASservingLocation returns the RequiredEASservingLocation field if non-nil, zero value otherwise.

### GetRequiredEASservingLocationOk

`func (o *ResourcesEdgeNrm) GetRequiredEASservingLocationOk() (*ServingLocation, bool)`

GetRequiredEASservingLocationOk returns a tuple with the RequiredEASservingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredEASservingLocation

`func (o *ResourcesEdgeNrm) SetRequiredEASservingLocation(v ServingLocation)`

SetRequiredEASservingLocation sets RequiredEASservingLocation field to given value.

### HasRequiredEASservingLocation

`func (o *ResourcesEdgeNrm) HasRequiredEASservingLocation() bool`

HasRequiredEASservingLocation returns a boolean if a field has been set.

### GetAffinityAntiAffinity

`func (o *ResourcesEdgeNrm) GetAffinityAntiAffinity() AffinityAntiAffinity`

GetAffinityAntiAffinity returns the AffinityAntiAffinity field if non-nil, zero value otherwise.

### GetAffinityAntiAffinityOk

`func (o *ResourcesEdgeNrm) GetAffinityAntiAffinityOk() (*AffinityAntiAffinity, bool)`

GetAffinityAntiAffinityOk returns a tuple with the AffinityAntiAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityAntiAffinity

`func (o *ResourcesEdgeNrm) SetAffinityAntiAffinity(v AffinityAntiAffinity)`

SetAffinityAntiAffinity sets AffinityAntiAffinity field to given value.

### HasAffinityAntiAffinity

`func (o *ResourcesEdgeNrm) HasAffinityAntiAffinity() bool`

HasAffinityAntiAffinity returns a boolean if a field has been set.

### GetServiceContinuity

`func (o *ResourcesEdgeNrm) GetServiceContinuity() bool`

GetServiceContinuity returns the ServiceContinuity field if non-nil, zero value otherwise.

### GetServiceContinuityOk

`func (o *ResourcesEdgeNrm) GetServiceContinuityOk() (*bool, bool)`

GetServiceContinuityOk returns a tuple with the ServiceContinuity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceContinuity

`func (o *ResourcesEdgeNrm) SetServiceContinuity(v bool)`

SetServiceContinuity sets ServiceContinuity field to given value.

### HasServiceContinuity

`func (o *ResourcesEdgeNrm) HasServiceContinuity() bool`

HasServiceContinuity returns a boolean if a field has been set.

### GetVirtualResource

`func (o *ResourcesEdgeNrm) GetVirtualResource() VirtualResource`

GetVirtualResource returns the VirtualResource field if non-nil, zero value otherwise.

### GetVirtualResourceOk

`func (o *ResourcesEdgeNrm) GetVirtualResourceOk() (*VirtualResource, bool)`

GetVirtualResourceOk returns a tuple with the VirtualResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualResource

`func (o *ResourcesEdgeNrm) SetVirtualResource(v VirtualResource)`

SetVirtualResource sets VirtualResource field to given value.

### HasVirtualResource

`func (o *ResourcesEdgeNrm) HasVirtualResource() bool`

HasVirtualResource returns a boolean if a field has been set.

### GetSoftwareImageInfo

`func (o *ResourcesEdgeNrm) GetSoftwareImageInfo() SoftwareImageInfo`

GetSoftwareImageInfo returns the SoftwareImageInfo field if non-nil, zero value otherwise.

### GetSoftwareImageInfoOk

`func (o *ResourcesEdgeNrm) GetSoftwareImageInfoOk() (*SoftwareImageInfo, bool)`

GetSoftwareImageInfoOk returns a tuple with the SoftwareImageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareImageInfo

`func (o *ResourcesEdgeNrm) SetSoftwareImageInfo(v SoftwareImageInfo)`

SetSoftwareImageInfo sets SoftwareImageInfo field to given value.

### HasSoftwareImageInfo

`func (o *ResourcesEdgeNrm) HasSoftwareImageInfo() bool`

HasSoftwareImageInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


