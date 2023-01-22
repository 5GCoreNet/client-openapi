# SubNetworkSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**SubNetworkAttr**](SubNetwork-Attr.md) |  | [optional] 
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

## Methods

### NewSubNetworkSingle

`func NewSubNetworkSingle(id NullableString, ) *SubNetworkSingle`

NewSubNetworkSingle instantiates a new SubNetworkSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubNetworkSingleWithDefaults

`func NewSubNetworkSingleWithDefaults() *SubNetworkSingle`

NewSubNetworkSingleWithDefaults instantiates a new SubNetworkSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubNetworkSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubNetworkSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubNetworkSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *SubNetworkSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SubNetworkSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *SubNetworkSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *SubNetworkSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *SubNetworkSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *SubNetworkSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *SubNetworkSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *SubNetworkSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *SubNetworkSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *SubNetworkSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *SubNetworkSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *SubNetworkSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *SubNetworkSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *SubNetworkSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *SubNetworkSingle) GetAttributes() SubNetworkAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubNetworkSingle) GetAttributesOk() (*SubNetworkAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubNetworkSingle) SetAttributes(v SubNetworkAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SubNetworkSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetSubnetwork

`func (o *SubNetworkSingle) GetSubnetwork() []SubNetworkSingle`

GetSubnetwork returns the Subnetwork field if non-nil, zero value otherwise.

### GetSubnetworkOk

`func (o *SubNetworkSingle) GetSubnetworkOk() (*[]SubNetworkSingle, bool)`

GetSubnetworkOk returns a tuple with the Subnetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetwork

`func (o *SubNetworkSingle) SetSubnetwork(v []SubNetworkSingle)`

SetSubnetwork sets Subnetwork field to given value.

### HasSubnetwork

`func (o *SubNetworkSingle) HasSubnetwork() bool`

HasSubnetwork returns a boolean if a field has been set.

### GetECSFunction

`func (o *SubNetworkSingle) GetECSFunction() []ECSFunctionSingle`

GetECSFunction returns the ECSFunction field if non-nil, zero value otherwise.

### GetECSFunctionOk

`func (o *SubNetworkSingle) GetECSFunctionOk() (*[]ECSFunctionSingle, bool)`

GetECSFunctionOk returns a tuple with the ECSFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECSFunction

`func (o *SubNetworkSingle) SetECSFunction(v []ECSFunctionSingle)`

SetECSFunction sets ECSFunction field to given value.

### HasECSFunction

`func (o *SubNetworkSingle) HasECSFunction() bool`

HasECSFunction returns a boolean if a field has been set.

### GetEdgeDataNetwork

`func (o *SubNetworkSingle) GetEdgeDataNetwork() []EdgeDataNetworkSingle`

GetEdgeDataNetwork returns the EdgeDataNetwork field if non-nil, zero value otherwise.

### GetEdgeDataNetworkOk

`func (o *SubNetworkSingle) GetEdgeDataNetworkOk() (*[]EdgeDataNetworkSingle, bool)`

GetEdgeDataNetworkOk returns a tuple with the EdgeDataNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeDataNetwork

`func (o *SubNetworkSingle) SetEdgeDataNetwork(v []EdgeDataNetworkSingle)`

SetEdgeDataNetwork sets EdgeDataNetwork field to given value.

### HasEdgeDataNetwork

`func (o *SubNetworkSingle) HasEdgeDataNetwork() bool`

HasEdgeDataNetwork returns a boolean if a field has been set.

### GetManagementNode

`func (o *SubNetworkSingle) GetManagementNode() []ManagementNodeSingle`

GetManagementNode returns the ManagementNode field if non-nil, zero value otherwise.

### GetManagementNodeOk

`func (o *SubNetworkSingle) GetManagementNodeOk() (*[]ManagementNodeSingle, bool)`

GetManagementNodeOk returns a tuple with the ManagementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNode

`func (o *SubNetworkSingle) SetManagementNode(v []ManagementNodeSingle)`

SetManagementNode sets ManagementNode field to given value.

### HasManagementNode

`func (o *SubNetworkSingle) HasManagementNode() bool`

HasManagementNode returns a boolean if a field has been set.

### GetMnsAgent

`func (o *SubNetworkSingle) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *SubNetworkSingle) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *SubNetworkSingle) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *SubNetworkSingle) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetMeContext

`func (o *SubNetworkSingle) GetMeContext() []MeContextSingle`

GetMeContext returns the MeContext field if non-nil, zero value otherwise.

### GetMeContextOk

`func (o *SubNetworkSingle) GetMeContextOk() (*[]MeContextSingle, bool)`

GetMeContextOk returns a tuple with the MeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeContext

`func (o *SubNetworkSingle) SetMeContext(v []MeContextSingle)`

SetMeContext sets MeContext field to given value.

### HasMeContext

`func (o *SubNetworkSingle) HasMeContext() bool`

HasMeContext returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *SubNetworkSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *SubNetworkSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *SubNetworkSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *SubNetworkSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *SubNetworkSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *SubNetworkSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *SubNetworkSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *SubNetworkSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetTraceJob

`func (o *SubNetworkSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *SubNetworkSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *SubNetworkSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *SubNetworkSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetManagementDataCollection

`func (o *SubNetworkSingle) GetManagementDataCollection() []ManagementDataCollectionSingle`

GetManagementDataCollection returns the ManagementDataCollection field if non-nil, zero value otherwise.

### GetManagementDataCollectionOk

`func (o *SubNetworkSingle) GetManagementDataCollectionOk() (*[]ManagementDataCollectionSingle, bool)`

GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementDataCollection

`func (o *SubNetworkSingle) SetManagementDataCollection(v []ManagementDataCollectionSingle)`

SetManagementDataCollection sets ManagementDataCollection field to given value.

### HasManagementDataCollection

`func (o *SubNetworkSingle) HasManagementDataCollection() bool`

HasManagementDataCollection returns a boolean if a field has been set.

### GetNtfSubscriptionControl

`func (o *SubNetworkSingle) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle`

GetNtfSubscriptionControl returns the NtfSubscriptionControl field if non-nil, zero value otherwise.

### GetNtfSubscriptionControlOk

`func (o *SubNetworkSingle) GetNtfSubscriptionControlOk() (*[]NtfSubscriptionControlSingle, bool)`

GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfSubscriptionControl

`func (o *SubNetworkSingle) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle)`

SetNtfSubscriptionControl sets NtfSubscriptionControl field to given value.

### HasNtfSubscriptionControl

`func (o *SubNetworkSingle) HasNtfSubscriptionControl() bool`

HasNtfSubscriptionControl returns a boolean if a field has been set.

### GetAlarmList

`func (o *SubNetworkSingle) GetAlarmList() AlarmListSingle`

GetAlarmList returns the AlarmList field if non-nil, zero value otherwise.

### GetAlarmListOk

`func (o *SubNetworkSingle) GetAlarmListOk() (*AlarmListSingle, bool)`

GetAlarmListOk returns a tuple with the AlarmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmList

`func (o *SubNetworkSingle) SetAlarmList(v AlarmListSingle)`

SetAlarmList sets AlarmList field to given value.

### HasAlarmList

`func (o *SubNetworkSingle) HasAlarmList() bool`

HasAlarmList returns a boolean if a field has been set.

### GetFiles

`func (o *SubNetworkSingle) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SubNetworkSingle) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SubNetworkSingle) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SubNetworkSingle) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFileDownloadJob

`func (o *SubNetworkSingle) GetFileDownloadJob() []FileDownloadJobSingle`

GetFileDownloadJob returns the FileDownloadJob field if non-nil, zero value otherwise.

### GetFileDownloadJobOk

`func (o *SubNetworkSingle) GetFileDownloadJobOk() (*[]FileDownloadJobSingle, bool)`

GetFileDownloadJobOk returns a tuple with the FileDownloadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadJob

`func (o *SubNetworkSingle) SetFileDownloadJob(v []FileDownloadJobSingle)`

SetFileDownloadJob sets FileDownloadJob field to given value.

### HasFileDownloadJob

`func (o *SubNetworkSingle) HasFileDownloadJob() bool`

HasFileDownloadJob returns a boolean if a field has been set.

### GetMnsRegistry

`func (o *SubNetworkSingle) GetMnsRegistry() MnsRegistrySingle`

GetMnsRegistry returns the MnsRegistry field if non-nil, zero value otherwise.

### GetMnsRegistryOk

`func (o *SubNetworkSingle) GetMnsRegistryOk() (*MnsRegistrySingle, bool)`

GetMnsRegistryOk returns a tuple with the MnsRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsRegistry

`func (o *SubNetworkSingle) SetMnsRegistry(v MnsRegistrySingle)`

SetMnsRegistry sets MnsRegistry field to given value.

### HasMnsRegistry

`func (o *SubNetworkSingle) HasMnsRegistry() bool`

HasMnsRegistry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

