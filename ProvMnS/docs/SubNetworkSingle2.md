# SubNetworkSingle2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**SubNetworkAttr**](SubNetwork-Attr.md) |  | [optional] 
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
**SubNetwork** | Pointer to [**[]SubNetworkSingle**](SubNetworkSingle.md) |  | [optional] 
**NetworkSlice** | Pointer to [**[]NetworkSliceSingle**](NetworkSliceSingle.md) |  | [optional] 
**NetworkSliceSubnet** | Pointer to [**[]NetworkSliceSubnetSingle**](NetworkSliceSubnetSingle.md) |  | [optional] 
**EPTransport** | Pointer to [**[]EPTransportSingle**](EPTransportSingle.md) |  | [optional] 
**NetworkSliceSubnetProviderCapabilities** | Pointer to [**[]NetworkSliceSubnetProviderCapabilitiesSingle**](NetworkSliceSubnetProviderCapabilitiesSingle.md) |  | [optional] 
**FeasibilityCheckAndReservationJob** | Pointer to [**[]FeasibilityCheckAndReservationJobSingle**](FeasibilityCheckAndReservationJobSingle.md) |  | [optional] 

## Methods

### NewSubNetworkSingle2

`func NewSubNetworkSingle2(id NullableString, ) *SubNetworkSingle2`

NewSubNetworkSingle2 instantiates a new SubNetworkSingle2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubNetworkSingle2WithDefaults

`func NewSubNetworkSingle2WithDefaults() *SubNetworkSingle2`

NewSubNetworkSingle2WithDefaults instantiates a new SubNetworkSingle2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubNetworkSingle2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubNetworkSingle2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubNetworkSingle2) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *SubNetworkSingle2) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SubNetworkSingle2) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *SubNetworkSingle2) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *SubNetworkSingle2) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *SubNetworkSingle2) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *SubNetworkSingle2) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *SubNetworkSingle2) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *SubNetworkSingle2) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *SubNetworkSingle2) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *SubNetworkSingle2) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *SubNetworkSingle2) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *SubNetworkSingle2) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *SubNetworkSingle2) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *SubNetworkSingle2) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *SubNetworkSingle2) GetAttributes() SubNetworkAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubNetworkSingle2) GetAttributesOk() (*SubNetworkAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubNetworkSingle2) SetAttributes(v SubNetworkAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SubNetworkSingle2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetManagementNode

`func (o *SubNetworkSingle2) GetManagementNode() []ManagementNodeSingle`

GetManagementNode returns the ManagementNode field if non-nil, zero value otherwise.

### GetManagementNodeOk

`func (o *SubNetworkSingle2) GetManagementNodeOk() (*[]ManagementNodeSingle, bool)`

GetManagementNodeOk returns a tuple with the ManagementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNode

`func (o *SubNetworkSingle2) SetManagementNode(v []ManagementNodeSingle)`

SetManagementNode sets ManagementNode field to given value.

### HasManagementNode

`func (o *SubNetworkSingle2) HasManagementNode() bool`

HasManagementNode returns a boolean if a field has been set.

### GetMnsAgent

`func (o *SubNetworkSingle2) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *SubNetworkSingle2) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *SubNetworkSingle2) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *SubNetworkSingle2) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetMeContext

`func (o *SubNetworkSingle2) GetMeContext() []MeContextSingle`

GetMeContext returns the MeContext field if non-nil, zero value otherwise.

### GetMeContextOk

`func (o *SubNetworkSingle2) GetMeContextOk() (*[]MeContextSingle, bool)`

GetMeContextOk returns a tuple with the MeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeContext

`func (o *SubNetworkSingle2) SetMeContext(v []MeContextSingle)`

SetMeContext sets MeContext field to given value.

### HasMeContext

`func (o *SubNetworkSingle2) HasMeContext() bool`

HasMeContext returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *SubNetworkSingle2) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *SubNetworkSingle2) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *SubNetworkSingle2) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *SubNetworkSingle2) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *SubNetworkSingle2) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *SubNetworkSingle2) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *SubNetworkSingle2) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *SubNetworkSingle2) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetTraceJob

`func (o *SubNetworkSingle2) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *SubNetworkSingle2) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *SubNetworkSingle2) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *SubNetworkSingle2) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetManagementDataCollection

`func (o *SubNetworkSingle2) GetManagementDataCollection() []ManagementDataCollectionSingle`

GetManagementDataCollection returns the ManagementDataCollection field if non-nil, zero value otherwise.

### GetManagementDataCollectionOk

`func (o *SubNetworkSingle2) GetManagementDataCollectionOk() (*[]ManagementDataCollectionSingle, bool)`

GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementDataCollection

`func (o *SubNetworkSingle2) SetManagementDataCollection(v []ManagementDataCollectionSingle)`

SetManagementDataCollection sets ManagementDataCollection field to given value.

### HasManagementDataCollection

`func (o *SubNetworkSingle2) HasManagementDataCollection() bool`

HasManagementDataCollection returns a boolean if a field has been set.

### GetNtfSubscriptionControl

`func (o *SubNetworkSingle2) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle`

GetNtfSubscriptionControl returns the NtfSubscriptionControl field if non-nil, zero value otherwise.

### GetNtfSubscriptionControlOk

`func (o *SubNetworkSingle2) GetNtfSubscriptionControlOk() (*[]NtfSubscriptionControlSingle, bool)`

GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfSubscriptionControl

`func (o *SubNetworkSingle2) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle)`

SetNtfSubscriptionControl sets NtfSubscriptionControl field to given value.

### HasNtfSubscriptionControl

`func (o *SubNetworkSingle2) HasNtfSubscriptionControl() bool`

HasNtfSubscriptionControl returns a boolean if a field has been set.

### GetAlarmList

`func (o *SubNetworkSingle2) GetAlarmList() AlarmListSingle`

GetAlarmList returns the AlarmList field if non-nil, zero value otherwise.

### GetAlarmListOk

`func (o *SubNetworkSingle2) GetAlarmListOk() (*AlarmListSingle, bool)`

GetAlarmListOk returns a tuple with the AlarmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmList

`func (o *SubNetworkSingle2) SetAlarmList(v AlarmListSingle)`

SetAlarmList sets AlarmList field to given value.

### HasAlarmList

`func (o *SubNetworkSingle2) HasAlarmList() bool`

HasAlarmList returns a boolean if a field has been set.

### GetFiles

`func (o *SubNetworkSingle2) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SubNetworkSingle2) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SubNetworkSingle2) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SubNetworkSingle2) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFileDownloadJob

`func (o *SubNetworkSingle2) GetFileDownloadJob() []FileDownloadJobSingle`

GetFileDownloadJob returns the FileDownloadJob field if non-nil, zero value otherwise.

### GetFileDownloadJobOk

`func (o *SubNetworkSingle2) GetFileDownloadJobOk() (*[]FileDownloadJobSingle, bool)`

GetFileDownloadJobOk returns a tuple with the FileDownloadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadJob

`func (o *SubNetworkSingle2) SetFileDownloadJob(v []FileDownloadJobSingle)`

SetFileDownloadJob sets FileDownloadJob field to given value.

### HasFileDownloadJob

`func (o *SubNetworkSingle2) HasFileDownloadJob() bool`

HasFileDownloadJob returns a boolean if a field has been set.

### GetMnsRegistry

`func (o *SubNetworkSingle2) GetMnsRegistry() MnsRegistrySingle`

GetMnsRegistry returns the MnsRegistry field if non-nil, zero value otherwise.

### GetMnsRegistryOk

`func (o *SubNetworkSingle2) GetMnsRegistryOk() (*MnsRegistrySingle, bool)`

GetMnsRegistryOk returns a tuple with the MnsRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsRegistry

`func (o *SubNetworkSingle2) SetMnsRegistry(v MnsRegistrySingle)`

SetMnsRegistry sets MnsRegistry field to given value.

### HasMnsRegistry

`func (o *SubNetworkSingle2) HasMnsRegistry() bool`

HasMnsRegistry returns a boolean if a field has been set.

### GetSubNetwork

`func (o *SubNetworkSingle2) GetSubNetwork() []SubNetworkSingle`

GetSubNetwork returns the SubNetwork field if non-nil, zero value otherwise.

### GetSubNetworkOk

`func (o *SubNetworkSingle2) GetSubNetworkOk() (*[]SubNetworkSingle, bool)`

GetSubNetworkOk returns a tuple with the SubNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetwork

`func (o *SubNetworkSingle2) SetSubNetwork(v []SubNetworkSingle)`

SetSubNetwork sets SubNetwork field to given value.

### HasSubNetwork

`func (o *SubNetworkSingle2) HasSubNetwork() bool`

HasSubNetwork returns a boolean if a field has been set.

### GetNetworkSlice

`func (o *SubNetworkSingle2) GetNetworkSlice() []NetworkSliceSingle`

GetNetworkSlice returns the NetworkSlice field if non-nil, zero value otherwise.

### GetNetworkSliceOk

`func (o *SubNetworkSingle2) GetNetworkSliceOk() (*[]NetworkSliceSingle, bool)`

GetNetworkSliceOk returns a tuple with the NetworkSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSlice

`func (o *SubNetworkSingle2) SetNetworkSlice(v []NetworkSliceSingle)`

SetNetworkSlice sets NetworkSlice field to given value.

### HasNetworkSlice

`func (o *SubNetworkSingle2) HasNetworkSlice() bool`

HasNetworkSlice returns a boolean if a field has been set.

### GetNetworkSliceSubnet

`func (o *SubNetworkSingle2) GetNetworkSliceSubnet() []NetworkSliceSubnetSingle`

GetNetworkSliceSubnet returns the NetworkSliceSubnet field if non-nil, zero value otherwise.

### GetNetworkSliceSubnetOk

`func (o *SubNetworkSingle2) GetNetworkSliceSubnetOk() (*[]NetworkSliceSubnetSingle, bool)`

GetNetworkSliceSubnetOk returns a tuple with the NetworkSliceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSliceSubnet

`func (o *SubNetworkSingle2) SetNetworkSliceSubnet(v []NetworkSliceSubnetSingle)`

SetNetworkSliceSubnet sets NetworkSliceSubnet field to given value.

### HasNetworkSliceSubnet

`func (o *SubNetworkSingle2) HasNetworkSliceSubnet() bool`

HasNetworkSliceSubnet returns a boolean if a field has been set.

### GetEPTransport

`func (o *SubNetworkSingle2) GetEPTransport() []EPTransportSingle`

GetEPTransport returns the EPTransport field if non-nil, zero value otherwise.

### GetEPTransportOk

`func (o *SubNetworkSingle2) GetEPTransportOk() (*[]EPTransportSingle, bool)`

GetEPTransportOk returns a tuple with the EPTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPTransport

`func (o *SubNetworkSingle2) SetEPTransport(v []EPTransportSingle)`

SetEPTransport sets EPTransport field to given value.

### HasEPTransport

`func (o *SubNetworkSingle2) HasEPTransport() bool`

HasEPTransport returns a boolean if a field has been set.

### GetNetworkSliceSubnetProviderCapabilities

`func (o *SubNetworkSingle2) GetNetworkSliceSubnetProviderCapabilities() []NetworkSliceSubnetProviderCapabilitiesSingle`

GetNetworkSliceSubnetProviderCapabilities returns the NetworkSliceSubnetProviderCapabilities field if non-nil, zero value otherwise.

### GetNetworkSliceSubnetProviderCapabilitiesOk

`func (o *SubNetworkSingle2) GetNetworkSliceSubnetProviderCapabilitiesOk() (*[]NetworkSliceSubnetProviderCapabilitiesSingle, bool)`

GetNetworkSliceSubnetProviderCapabilitiesOk returns a tuple with the NetworkSliceSubnetProviderCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSliceSubnetProviderCapabilities

`func (o *SubNetworkSingle2) SetNetworkSliceSubnetProviderCapabilities(v []NetworkSliceSubnetProviderCapabilitiesSingle)`

SetNetworkSliceSubnetProviderCapabilities sets NetworkSliceSubnetProviderCapabilities field to given value.

### HasNetworkSliceSubnetProviderCapabilities

`func (o *SubNetworkSingle2) HasNetworkSliceSubnetProviderCapabilities() bool`

HasNetworkSliceSubnetProviderCapabilities returns a boolean if a field has been set.

### GetFeasibilityCheckAndReservationJob

`func (o *SubNetworkSingle2) GetFeasibilityCheckAndReservationJob() []FeasibilityCheckAndReservationJobSingle`

GetFeasibilityCheckAndReservationJob returns the FeasibilityCheckAndReservationJob field if non-nil, zero value otherwise.

### GetFeasibilityCheckAndReservationJobOk

`func (o *SubNetworkSingle2) GetFeasibilityCheckAndReservationJobOk() (*[]FeasibilityCheckAndReservationJobSingle, bool)`

GetFeasibilityCheckAndReservationJobOk returns a tuple with the FeasibilityCheckAndReservationJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeasibilityCheckAndReservationJob

`func (o *SubNetworkSingle2) SetFeasibilityCheckAndReservationJob(v []FeasibilityCheckAndReservationJobSingle)`

SetFeasibilityCheckAndReservationJob sets FeasibilityCheckAndReservationJob field to given value.

### HasFeasibilityCheckAndReservationJob

`func (o *SubNetworkSingle2) HasFeasibilityCheckAndReservationJob() bool`

HasFeasibilityCheckAndReservationJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


