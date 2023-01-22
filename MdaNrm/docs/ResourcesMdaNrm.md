# ResourcesMdaNrm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to **interface{}** |  | [optional] 
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
**ManagedElement** | Pointer to [**[]ManagedElementSingle**](ManagedElementSingle.md) |  | [optional] 
**MDAFunction** | Pointer to [**[]MDAFunctionSingle**](MDAFunctionSingle.md) |  | [optional] 
**MDAReport** | Pointer to [**[]MDAReportSingle**](MDAReportSingle.md) |  | [optional] 
**ManagedNFService** | Pointer to [**[]ManagedNFServiceSingle**](ManagedNFServiceSingle.md) |  | [optional] 
**MDARequest** | Pointer to [**[]MDARequestSingle**](MDARequestSingle.md) |  | [optional] 

## Methods

### NewResourcesMdaNrm

`func NewResourcesMdaNrm(id NullableString, ) *ResourcesMdaNrm`

NewResourcesMdaNrm instantiates a new ResourcesMdaNrm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesMdaNrmWithDefaults

`func NewResourcesMdaNrmWithDefaults() *ResourcesMdaNrm`

NewResourcesMdaNrmWithDefaults instantiates a new ResourcesMdaNrm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourcesMdaNrm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcesMdaNrm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcesMdaNrm) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ResourcesMdaNrm) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ResourcesMdaNrm) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *ResourcesMdaNrm) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ResourcesMdaNrm) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ResourcesMdaNrm) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ResourcesMdaNrm) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *ResourcesMdaNrm) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *ResourcesMdaNrm) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *ResourcesMdaNrm) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *ResourcesMdaNrm) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *ResourcesMdaNrm) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *ResourcesMdaNrm) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *ResourcesMdaNrm) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *ResourcesMdaNrm) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *ResourcesMdaNrm) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ResourcesMdaNrm) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ResourcesMdaNrm) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ResourcesMdaNrm) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetManagementNode

`func (o *ResourcesMdaNrm) GetManagementNode() []ManagementNodeSingle`

GetManagementNode returns the ManagementNode field if non-nil, zero value otherwise.

### GetManagementNodeOk

`func (o *ResourcesMdaNrm) GetManagementNodeOk() (*[]ManagementNodeSingle, bool)`

GetManagementNodeOk returns a tuple with the ManagementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNode

`func (o *ResourcesMdaNrm) SetManagementNode(v []ManagementNodeSingle)`

SetManagementNode sets ManagementNode field to given value.

### HasManagementNode

`func (o *ResourcesMdaNrm) HasManagementNode() bool`

HasManagementNode returns a boolean if a field has been set.

### GetMnsAgent

`func (o *ResourcesMdaNrm) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *ResourcesMdaNrm) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *ResourcesMdaNrm) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *ResourcesMdaNrm) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetMeContext

`func (o *ResourcesMdaNrm) GetMeContext() []MeContextSingle`

GetMeContext returns the MeContext field if non-nil, zero value otherwise.

### GetMeContextOk

`func (o *ResourcesMdaNrm) GetMeContextOk() (*[]MeContextSingle, bool)`

GetMeContextOk returns a tuple with the MeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeContext

`func (o *ResourcesMdaNrm) SetMeContext(v []MeContextSingle)`

SetMeContext sets MeContext field to given value.

### HasMeContext

`func (o *ResourcesMdaNrm) HasMeContext() bool`

HasMeContext returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *ResourcesMdaNrm) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *ResourcesMdaNrm) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *ResourcesMdaNrm) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *ResourcesMdaNrm) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *ResourcesMdaNrm) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *ResourcesMdaNrm) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *ResourcesMdaNrm) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *ResourcesMdaNrm) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetTraceJob

`func (o *ResourcesMdaNrm) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *ResourcesMdaNrm) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *ResourcesMdaNrm) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *ResourcesMdaNrm) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetManagementDataCollection

`func (o *ResourcesMdaNrm) GetManagementDataCollection() []ManagementDataCollectionSingle`

GetManagementDataCollection returns the ManagementDataCollection field if non-nil, zero value otherwise.

### GetManagementDataCollectionOk

`func (o *ResourcesMdaNrm) GetManagementDataCollectionOk() (*[]ManagementDataCollectionSingle, bool)`

GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementDataCollection

`func (o *ResourcesMdaNrm) SetManagementDataCollection(v []ManagementDataCollectionSingle)`

SetManagementDataCollection sets ManagementDataCollection field to given value.

### HasManagementDataCollection

`func (o *ResourcesMdaNrm) HasManagementDataCollection() bool`

HasManagementDataCollection returns a boolean if a field has been set.

### GetNtfSubscriptionControl

`func (o *ResourcesMdaNrm) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle`

GetNtfSubscriptionControl returns the NtfSubscriptionControl field if non-nil, zero value otherwise.

### GetNtfSubscriptionControlOk

`func (o *ResourcesMdaNrm) GetNtfSubscriptionControlOk() (*[]NtfSubscriptionControlSingle, bool)`

GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfSubscriptionControl

`func (o *ResourcesMdaNrm) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle)`

SetNtfSubscriptionControl sets NtfSubscriptionControl field to given value.

### HasNtfSubscriptionControl

`func (o *ResourcesMdaNrm) HasNtfSubscriptionControl() bool`

HasNtfSubscriptionControl returns a boolean if a field has been set.

### GetAlarmList

`func (o *ResourcesMdaNrm) GetAlarmList() AlarmListSingle`

GetAlarmList returns the AlarmList field if non-nil, zero value otherwise.

### GetAlarmListOk

`func (o *ResourcesMdaNrm) GetAlarmListOk() (*AlarmListSingle, bool)`

GetAlarmListOk returns a tuple with the AlarmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmList

`func (o *ResourcesMdaNrm) SetAlarmList(v AlarmListSingle)`

SetAlarmList sets AlarmList field to given value.

### HasAlarmList

`func (o *ResourcesMdaNrm) HasAlarmList() bool`

HasAlarmList returns a boolean if a field has been set.

### GetFiles

`func (o *ResourcesMdaNrm) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ResourcesMdaNrm) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ResourcesMdaNrm) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ResourcesMdaNrm) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFileDownloadJob

`func (o *ResourcesMdaNrm) GetFileDownloadJob() []FileDownloadJobSingle`

GetFileDownloadJob returns the FileDownloadJob field if non-nil, zero value otherwise.

### GetFileDownloadJobOk

`func (o *ResourcesMdaNrm) GetFileDownloadJobOk() (*[]FileDownloadJobSingle, bool)`

GetFileDownloadJobOk returns a tuple with the FileDownloadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadJob

`func (o *ResourcesMdaNrm) SetFileDownloadJob(v []FileDownloadJobSingle)`

SetFileDownloadJob sets FileDownloadJob field to given value.

### HasFileDownloadJob

`func (o *ResourcesMdaNrm) HasFileDownloadJob() bool`

HasFileDownloadJob returns a boolean if a field has been set.

### GetMnsRegistry

`func (o *ResourcesMdaNrm) GetMnsRegistry() MnsRegistrySingle`

GetMnsRegistry returns the MnsRegistry field if non-nil, zero value otherwise.

### GetMnsRegistryOk

`func (o *ResourcesMdaNrm) GetMnsRegistryOk() (*MnsRegistrySingle, bool)`

GetMnsRegistryOk returns a tuple with the MnsRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsRegistry

`func (o *ResourcesMdaNrm) SetMnsRegistry(v MnsRegistrySingle)`

SetMnsRegistry sets MnsRegistry field to given value.

### HasMnsRegistry

`func (o *ResourcesMdaNrm) HasMnsRegistry() bool`

HasMnsRegistry returns a boolean if a field has been set.

### GetSubNetwork

`func (o *ResourcesMdaNrm) GetSubNetwork() []SubNetworkSingle`

GetSubNetwork returns the SubNetwork field if non-nil, zero value otherwise.

### GetSubNetworkOk

`func (o *ResourcesMdaNrm) GetSubNetworkOk() (*[]SubNetworkSingle, bool)`

GetSubNetworkOk returns a tuple with the SubNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetwork

`func (o *ResourcesMdaNrm) SetSubNetwork(v []SubNetworkSingle)`

SetSubNetwork sets SubNetwork field to given value.

### HasSubNetwork

`func (o *ResourcesMdaNrm) HasSubNetwork() bool`

HasSubNetwork returns a boolean if a field has been set.

### GetManagedElement

`func (o *ResourcesMdaNrm) GetManagedElement() []ManagedElementSingle`

GetManagedElement returns the ManagedElement field if non-nil, zero value otherwise.

### GetManagedElementOk

`func (o *ResourcesMdaNrm) GetManagedElementOk() (*[]ManagedElementSingle, bool)`

GetManagedElementOk returns a tuple with the ManagedElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedElement

`func (o *ResourcesMdaNrm) SetManagedElement(v []ManagedElementSingle)`

SetManagedElement sets ManagedElement field to given value.

### HasManagedElement

`func (o *ResourcesMdaNrm) HasManagedElement() bool`

HasManagedElement returns a boolean if a field has been set.

### GetMDAFunction

`func (o *ResourcesMdaNrm) GetMDAFunction() []MDAFunctionSingle`

GetMDAFunction returns the MDAFunction field if non-nil, zero value otherwise.

### GetMDAFunctionOk

`func (o *ResourcesMdaNrm) GetMDAFunctionOk() (*[]MDAFunctionSingle, bool)`

GetMDAFunctionOk returns a tuple with the MDAFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMDAFunction

`func (o *ResourcesMdaNrm) SetMDAFunction(v []MDAFunctionSingle)`

SetMDAFunction sets MDAFunction field to given value.

### HasMDAFunction

`func (o *ResourcesMdaNrm) HasMDAFunction() bool`

HasMDAFunction returns a boolean if a field has been set.

### GetMDAReport

`func (o *ResourcesMdaNrm) GetMDAReport() []MDAReportSingle`

GetMDAReport returns the MDAReport field if non-nil, zero value otherwise.

### GetMDAReportOk

`func (o *ResourcesMdaNrm) GetMDAReportOk() (*[]MDAReportSingle, bool)`

GetMDAReportOk returns a tuple with the MDAReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMDAReport

`func (o *ResourcesMdaNrm) SetMDAReport(v []MDAReportSingle)`

SetMDAReport sets MDAReport field to given value.

### HasMDAReport

`func (o *ResourcesMdaNrm) HasMDAReport() bool`

HasMDAReport returns a boolean if a field has been set.

### GetManagedNFService

`func (o *ResourcesMdaNrm) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *ResourcesMdaNrm) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *ResourcesMdaNrm) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *ResourcesMdaNrm) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetMDARequest

`func (o *ResourcesMdaNrm) GetMDARequest() []MDARequestSingle`

GetMDARequest returns the MDARequest field if non-nil, zero value otherwise.

### GetMDARequestOk

`func (o *ResourcesMdaNrm) GetMDARequestOk() (*[]MDARequestSingle, bool)`

GetMDARequestOk returns a tuple with the MDARequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMDARequest

`func (o *ResourcesMdaNrm) SetMDARequest(v []MDARequestSingle)`

SetMDARequest sets MDARequest field to given value.

### HasMDARequest

`func (o *ResourcesMdaNrm) HasMDARequest() bool`

HasMDARequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


