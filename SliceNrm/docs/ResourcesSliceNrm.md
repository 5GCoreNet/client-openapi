# ResourcesSliceNrm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubNetwork** | Pointer to [**[]SubNetworkSingle**](SubNetworkSingle.md) |  | [optional] 
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**FeasibilityCheckAndReservationJobSingleAllOfAttributes**](FeasibilityCheckAndReservationJobSingleAllOfAttributes.md) |  | [optional] 
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
**NetworkSlice** | Pointer to [**[]NetworkSliceSingle**](NetworkSliceSingle.md) |  | [optional] 
**NetworkSliceSubnet** | Pointer to [**[]NetworkSliceSubnetSingle**](NetworkSliceSubnetSingle.md) |  | [optional] 
**EPTransport** | Pointer to [**[]EPTransportSingle**](EPTransportSingle.md) |  | [optional] 
**NetworkSliceSubnetProviderCapabilities** | Pointer to [**[]NetworkSliceSubnetProviderCapabilitiesSingle**](NetworkSliceSubnetProviderCapabilitiesSingle.md) |  | [optional] 
**FeasibilityCheckAndReservationJob** | Pointer to [**[]FeasibilityCheckAndReservationJobSingle**](FeasibilityCheckAndReservationJobSingle.md) |  | [optional] 

## Methods

### NewResourcesSliceNrm

`func NewResourcesSliceNrm(id NullableString, ) *ResourcesSliceNrm`

NewResourcesSliceNrm instantiates a new ResourcesSliceNrm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesSliceNrmWithDefaults

`func NewResourcesSliceNrmWithDefaults() *ResourcesSliceNrm`

NewResourcesSliceNrmWithDefaults instantiates a new ResourcesSliceNrm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubNetwork

`func (o *ResourcesSliceNrm) GetSubNetwork() []SubNetworkSingle`

GetSubNetwork returns the SubNetwork field if non-nil, zero value otherwise.

### GetSubNetworkOk

`func (o *ResourcesSliceNrm) GetSubNetworkOk() (*[]SubNetworkSingle, bool)`

GetSubNetworkOk returns a tuple with the SubNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetwork

`func (o *ResourcesSliceNrm) SetSubNetwork(v []SubNetworkSingle)`

SetSubNetwork sets SubNetwork field to given value.

### HasSubNetwork

`func (o *ResourcesSliceNrm) HasSubNetwork() bool`

HasSubNetwork returns a boolean if a field has been set.

### GetId

`func (o *ResourcesSliceNrm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcesSliceNrm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcesSliceNrm) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ResourcesSliceNrm) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ResourcesSliceNrm) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *ResourcesSliceNrm) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ResourcesSliceNrm) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ResourcesSliceNrm) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ResourcesSliceNrm) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *ResourcesSliceNrm) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *ResourcesSliceNrm) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *ResourcesSliceNrm) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *ResourcesSliceNrm) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *ResourcesSliceNrm) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *ResourcesSliceNrm) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *ResourcesSliceNrm) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *ResourcesSliceNrm) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *ResourcesSliceNrm) GetAttributes() FeasibilityCheckAndReservationJobSingleAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ResourcesSliceNrm) GetAttributesOk() (*FeasibilityCheckAndReservationJobSingleAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ResourcesSliceNrm) SetAttributes(v FeasibilityCheckAndReservationJobSingleAllOfAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ResourcesSliceNrm) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetManagementNode

`func (o *ResourcesSliceNrm) GetManagementNode() []ManagementNodeSingle`

GetManagementNode returns the ManagementNode field if non-nil, zero value otherwise.

### GetManagementNodeOk

`func (o *ResourcesSliceNrm) GetManagementNodeOk() (*[]ManagementNodeSingle, bool)`

GetManagementNodeOk returns a tuple with the ManagementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNode

`func (o *ResourcesSliceNrm) SetManagementNode(v []ManagementNodeSingle)`

SetManagementNode sets ManagementNode field to given value.

### HasManagementNode

`func (o *ResourcesSliceNrm) HasManagementNode() bool`

HasManagementNode returns a boolean if a field has been set.

### GetMnsAgent

`func (o *ResourcesSliceNrm) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *ResourcesSliceNrm) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *ResourcesSliceNrm) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *ResourcesSliceNrm) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetMeContext

`func (o *ResourcesSliceNrm) GetMeContext() []MeContextSingle`

GetMeContext returns the MeContext field if non-nil, zero value otherwise.

### GetMeContextOk

`func (o *ResourcesSliceNrm) GetMeContextOk() (*[]MeContextSingle, bool)`

GetMeContextOk returns a tuple with the MeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeContext

`func (o *ResourcesSliceNrm) SetMeContext(v []MeContextSingle)`

SetMeContext sets MeContext field to given value.

### HasMeContext

`func (o *ResourcesSliceNrm) HasMeContext() bool`

HasMeContext returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *ResourcesSliceNrm) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *ResourcesSliceNrm) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *ResourcesSliceNrm) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *ResourcesSliceNrm) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *ResourcesSliceNrm) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *ResourcesSliceNrm) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *ResourcesSliceNrm) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *ResourcesSliceNrm) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetTraceJob

`func (o *ResourcesSliceNrm) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *ResourcesSliceNrm) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *ResourcesSliceNrm) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *ResourcesSliceNrm) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetManagementDataCollection

`func (o *ResourcesSliceNrm) GetManagementDataCollection() []ManagementDataCollectionSingle`

GetManagementDataCollection returns the ManagementDataCollection field if non-nil, zero value otherwise.

### GetManagementDataCollectionOk

`func (o *ResourcesSliceNrm) GetManagementDataCollectionOk() (*[]ManagementDataCollectionSingle, bool)`

GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementDataCollection

`func (o *ResourcesSliceNrm) SetManagementDataCollection(v []ManagementDataCollectionSingle)`

SetManagementDataCollection sets ManagementDataCollection field to given value.

### HasManagementDataCollection

`func (o *ResourcesSliceNrm) HasManagementDataCollection() bool`

HasManagementDataCollection returns a boolean if a field has been set.

### GetNtfSubscriptionControl

`func (o *ResourcesSliceNrm) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle`

GetNtfSubscriptionControl returns the NtfSubscriptionControl field if non-nil, zero value otherwise.

### GetNtfSubscriptionControlOk

`func (o *ResourcesSliceNrm) GetNtfSubscriptionControlOk() (*[]NtfSubscriptionControlSingle, bool)`

GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfSubscriptionControl

`func (o *ResourcesSliceNrm) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle)`

SetNtfSubscriptionControl sets NtfSubscriptionControl field to given value.

### HasNtfSubscriptionControl

`func (o *ResourcesSliceNrm) HasNtfSubscriptionControl() bool`

HasNtfSubscriptionControl returns a boolean if a field has been set.

### GetAlarmList

`func (o *ResourcesSliceNrm) GetAlarmList() AlarmListSingle`

GetAlarmList returns the AlarmList field if non-nil, zero value otherwise.

### GetAlarmListOk

`func (o *ResourcesSliceNrm) GetAlarmListOk() (*AlarmListSingle, bool)`

GetAlarmListOk returns a tuple with the AlarmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmList

`func (o *ResourcesSliceNrm) SetAlarmList(v AlarmListSingle)`

SetAlarmList sets AlarmList field to given value.

### HasAlarmList

`func (o *ResourcesSliceNrm) HasAlarmList() bool`

HasAlarmList returns a boolean if a field has been set.

### GetFiles

`func (o *ResourcesSliceNrm) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ResourcesSliceNrm) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ResourcesSliceNrm) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ResourcesSliceNrm) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFileDownloadJob

`func (o *ResourcesSliceNrm) GetFileDownloadJob() []FileDownloadJobSingle`

GetFileDownloadJob returns the FileDownloadJob field if non-nil, zero value otherwise.

### GetFileDownloadJobOk

`func (o *ResourcesSliceNrm) GetFileDownloadJobOk() (*[]FileDownloadJobSingle, bool)`

GetFileDownloadJobOk returns a tuple with the FileDownloadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadJob

`func (o *ResourcesSliceNrm) SetFileDownloadJob(v []FileDownloadJobSingle)`

SetFileDownloadJob sets FileDownloadJob field to given value.

### HasFileDownloadJob

`func (o *ResourcesSliceNrm) HasFileDownloadJob() bool`

HasFileDownloadJob returns a boolean if a field has been set.

### GetMnsRegistry

`func (o *ResourcesSliceNrm) GetMnsRegistry() MnsRegistrySingle`

GetMnsRegistry returns the MnsRegistry field if non-nil, zero value otherwise.

### GetMnsRegistryOk

`func (o *ResourcesSliceNrm) GetMnsRegistryOk() (*MnsRegistrySingle, bool)`

GetMnsRegistryOk returns a tuple with the MnsRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsRegistry

`func (o *ResourcesSliceNrm) SetMnsRegistry(v MnsRegistrySingle)`

SetMnsRegistry sets MnsRegistry field to given value.

### HasMnsRegistry

`func (o *ResourcesSliceNrm) HasMnsRegistry() bool`

HasMnsRegistry returns a boolean if a field has been set.

### GetNetworkSlice

`func (o *ResourcesSliceNrm) GetNetworkSlice() []NetworkSliceSingle`

GetNetworkSlice returns the NetworkSlice field if non-nil, zero value otherwise.

### GetNetworkSliceOk

`func (o *ResourcesSliceNrm) GetNetworkSliceOk() (*[]NetworkSliceSingle, bool)`

GetNetworkSliceOk returns a tuple with the NetworkSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSlice

`func (o *ResourcesSliceNrm) SetNetworkSlice(v []NetworkSliceSingle)`

SetNetworkSlice sets NetworkSlice field to given value.

### HasNetworkSlice

`func (o *ResourcesSliceNrm) HasNetworkSlice() bool`

HasNetworkSlice returns a boolean if a field has been set.

### GetNetworkSliceSubnet

`func (o *ResourcesSliceNrm) GetNetworkSliceSubnet() []NetworkSliceSubnetSingle`

GetNetworkSliceSubnet returns the NetworkSliceSubnet field if non-nil, zero value otherwise.

### GetNetworkSliceSubnetOk

`func (o *ResourcesSliceNrm) GetNetworkSliceSubnetOk() (*[]NetworkSliceSubnetSingle, bool)`

GetNetworkSliceSubnetOk returns a tuple with the NetworkSliceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSliceSubnet

`func (o *ResourcesSliceNrm) SetNetworkSliceSubnet(v []NetworkSliceSubnetSingle)`

SetNetworkSliceSubnet sets NetworkSliceSubnet field to given value.

### HasNetworkSliceSubnet

`func (o *ResourcesSliceNrm) HasNetworkSliceSubnet() bool`

HasNetworkSliceSubnet returns a boolean if a field has been set.

### GetEPTransport

`func (o *ResourcesSliceNrm) GetEPTransport() []EPTransportSingle`

GetEPTransport returns the EPTransport field if non-nil, zero value otherwise.

### GetEPTransportOk

`func (o *ResourcesSliceNrm) GetEPTransportOk() (*[]EPTransportSingle, bool)`

GetEPTransportOk returns a tuple with the EPTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPTransport

`func (o *ResourcesSliceNrm) SetEPTransport(v []EPTransportSingle)`

SetEPTransport sets EPTransport field to given value.

### HasEPTransport

`func (o *ResourcesSliceNrm) HasEPTransport() bool`

HasEPTransport returns a boolean if a field has been set.

### GetNetworkSliceSubnetProviderCapabilities

`func (o *ResourcesSliceNrm) GetNetworkSliceSubnetProviderCapabilities() []NetworkSliceSubnetProviderCapabilitiesSingle`

GetNetworkSliceSubnetProviderCapabilities returns the NetworkSliceSubnetProviderCapabilities field if non-nil, zero value otherwise.

### GetNetworkSliceSubnetProviderCapabilitiesOk

`func (o *ResourcesSliceNrm) GetNetworkSliceSubnetProviderCapabilitiesOk() (*[]NetworkSliceSubnetProviderCapabilitiesSingle, bool)`

GetNetworkSliceSubnetProviderCapabilitiesOk returns a tuple with the NetworkSliceSubnetProviderCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSliceSubnetProviderCapabilities

`func (o *ResourcesSliceNrm) SetNetworkSliceSubnetProviderCapabilities(v []NetworkSliceSubnetProviderCapabilitiesSingle)`

SetNetworkSliceSubnetProviderCapabilities sets NetworkSliceSubnetProviderCapabilities field to given value.

### HasNetworkSliceSubnetProviderCapabilities

`func (o *ResourcesSliceNrm) HasNetworkSliceSubnetProviderCapabilities() bool`

HasNetworkSliceSubnetProviderCapabilities returns a boolean if a field has been set.

### GetFeasibilityCheckAndReservationJob

`func (o *ResourcesSliceNrm) GetFeasibilityCheckAndReservationJob() []FeasibilityCheckAndReservationJobSingle`

GetFeasibilityCheckAndReservationJob returns the FeasibilityCheckAndReservationJob field if non-nil, zero value otherwise.

### GetFeasibilityCheckAndReservationJobOk

`func (o *ResourcesSliceNrm) GetFeasibilityCheckAndReservationJobOk() (*[]FeasibilityCheckAndReservationJobSingle, bool)`

GetFeasibilityCheckAndReservationJobOk returns a tuple with the FeasibilityCheckAndReservationJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeasibilityCheckAndReservationJob

`func (o *ResourcesSliceNrm) SetFeasibilityCheckAndReservationJob(v []FeasibilityCheckAndReservationJobSingle)`

SetFeasibilityCheckAndReservationJob sets FeasibilityCheckAndReservationJob field to given value.

### HasFeasibilityCheckAndReservationJob

`func (o *ResourcesSliceNrm) HasFeasibilityCheckAndReservationJob() bool`

HasFeasibilityCheckAndReservationJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


