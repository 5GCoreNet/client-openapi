# SubNetworkSingle5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**SubNetworkAttr**](SubNetworkAttr.md) |  | [optional] 
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
**MLTrainingFunction** | Pointer to [**[]MLTrainingFunctionSingle**](MLTrainingFunctionSingle.md) |  | [optional] 

## Methods

### NewSubNetworkSingle5

`func NewSubNetworkSingle5(id NullableString, ) *SubNetworkSingle5`

NewSubNetworkSingle5 instantiates a new SubNetworkSingle5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubNetworkSingle5WithDefaults

`func NewSubNetworkSingle5WithDefaults() *SubNetworkSingle5`

NewSubNetworkSingle5WithDefaults instantiates a new SubNetworkSingle5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubNetworkSingle5) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubNetworkSingle5) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubNetworkSingle5) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *SubNetworkSingle5) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SubNetworkSingle5) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *SubNetworkSingle5) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *SubNetworkSingle5) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *SubNetworkSingle5) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *SubNetworkSingle5) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *SubNetworkSingle5) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *SubNetworkSingle5) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *SubNetworkSingle5) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *SubNetworkSingle5) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *SubNetworkSingle5) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *SubNetworkSingle5) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *SubNetworkSingle5) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *SubNetworkSingle5) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *SubNetworkSingle5) GetAttributes() SubNetworkAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubNetworkSingle5) GetAttributesOk() (*SubNetworkAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubNetworkSingle5) SetAttributes(v SubNetworkAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SubNetworkSingle5) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetManagementNode

`func (o *SubNetworkSingle5) GetManagementNode() []ManagementNodeSingle`

GetManagementNode returns the ManagementNode field if non-nil, zero value otherwise.

### GetManagementNodeOk

`func (o *SubNetworkSingle5) GetManagementNodeOk() (*[]ManagementNodeSingle, bool)`

GetManagementNodeOk returns a tuple with the ManagementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNode

`func (o *SubNetworkSingle5) SetManagementNode(v []ManagementNodeSingle)`

SetManagementNode sets ManagementNode field to given value.

### HasManagementNode

`func (o *SubNetworkSingle5) HasManagementNode() bool`

HasManagementNode returns a boolean if a field has been set.

### GetMnsAgent

`func (o *SubNetworkSingle5) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *SubNetworkSingle5) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *SubNetworkSingle5) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *SubNetworkSingle5) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetMeContext

`func (o *SubNetworkSingle5) GetMeContext() []MeContextSingle`

GetMeContext returns the MeContext field if non-nil, zero value otherwise.

### GetMeContextOk

`func (o *SubNetworkSingle5) GetMeContextOk() (*[]MeContextSingle, bool)`

GetMeContextOk returns a tuple with the MeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeContext

`func (o *SubNetworkSingle5) SetMeContext(v []MeContextSingle)`

SetMeContext sets MeContext field to given value.

### HasMeContext

`func (o *SubNetworkSingle5) HasMeContext() bool`

HasMeContext returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *SubNetworkSingle5) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *SubNetworkSingle5) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *SubNetworkSingle5) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *SubNetworkSingle5) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *SubNetworkSingle5) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *SubNetworkSingle5) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *SubNetworkSingle5) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *SubNetworkSingle5) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetTraceJob

`func (o *SubNetworkSingle5) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *SubNetworkSingle5) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *SubNetworkSingle5) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *SubNetworkSingle5) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetManagementDataCollection

`func (o *SubNetworkSingle5) GetManagementDataCollection() []ManagementDataCollectionSingle`

GetManagementDataCollection returns the ManagementDataCollection field if non-nil, zero value otherwise.

### GetManagementDataCollectionOk

`func (o *SubNetworkSingle5) GetManagementDataCollectionOk() (*[]ManagementDataCollectionSingle, bool)`

GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementDataCollection

`func (o *SubNetworkSingle5) SetManagementDataCollection(v []ManagementDataCollectionSingle)`

SetManagementDataCollection sets ManagementDataCollection field to given value.

### HasManagementDataCollection

`func (o *SubNetworkSingle5) HasManagementDataCollection() bool`

HasManagementDataCollection returns a boolean if a field has been set.

### GetNtfSubscriptionControl

`func (o *SubNetworkSingle5) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle`

GetNtfSubscriptionControl returns the NtfSubscriptionControl field if non-nil, zero value otherwise.

### GetNtfSubscriptionControlOk

`func (o *SubNetworkSingle5) GetNtfSubscriptionControlOk() (*[]NtfSubscriptionControlSingle, bool)`

GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfSubscriptionControl

`func (o *SubNetworkSingle5) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle)`

SetNtfSubscriptionControl sets NtfSubscriptionControl field to given value.

### HasNtfSubscriptionControl

`func (o *SubNetworkSingle5) HasNtfSubscriptionControl() bool`

HasNtfSubscriptionControl returns a boolean if a field has been set.

### GetAlarmList

`func (o *SubNetworkSingle5) GetAlarmList() AlarmListSingle`

GetAlarmList returns the AlarmList field if non-nil, zero value otherwise.

### GetAlarmListOk

`func (o *SubNetworkSingle5) GetAlarmListOk() (*AlarmListSingle, bool)`

GetAlarmListOk returns a tuple with the AlarmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmList

`func (o *SubNetworkSingle5) SetAlarmList(v AlarmListSingle)`

SetAlarmList sets AlarmList field to given value.

### HasAlarmList

`func (o *SubNetworkSingle5) HasAlarmList() bool`

HasAlarmList returns a boolean if a field has been set.

### GetFiles

`func (o *SubNetworkSingle5) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SubNetworkSingle5) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SubNetworkSingle5) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SubNetworkSingle5) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFileDownloadJob

`func (o *SubNetworkSingle5) GetFileDownloadJob() []FileDownloadJobSingle`

GetFileDownloadJob returns the FileDownloadJob field if non-nil, zero value otherwise.

### GetFileDownloadJobOk

`func (o *SubNetworkSingle5) GetFileDownloadJobOk() (*[]FileDownloadJobSingle, bool)`

GetFileDownloadJobOk returns a tuple with the FileDownloadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadJob

`func (o *SubNetworkSingle5) SetFileDownloadJob(v []FileDownloadJobSingle)`

SetFileDownloadJob sets FileDownloadJob field to given value.

### HasFileDownloadJob

`func (o *SubNetworkSingle5) HasFileDownloadJob() bool`

HasFileDownloadJob returns a boolean if a field has been set.

### GetMnsRegistry

`func (o *SubNetworkSingle5) GetMnsRegistry() MnsRegistrySingle`

GetMnsRegistry returns the MnsRegistry field if non-nil, zero value otherwise.

### GetMnsRegistryOk

`func (o *SubNetworkSingle5) GetMnsRegistryOk() (*MnsRegistrySingle, bool)`

GetMnsRegistryOk returns a tuple with the MnsRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsRegistry

`func (o *SubNetworkSingle5) SetMnsRegistry(v MnsRegistrySingle)`

SetMnsRegistry sets MnsRegistry field to given value.

### HasMnsRegistry

`func (o *SubNetworkSingle5) HasMnsRegistry() bool`

HasMnsRegistry returns a boolean if a field has been set.

### GetSubNetwork

`func (o *SubNetworkSingle5) GetSubNetwork() []SubNetworkSingle`

GetSubNetwork returns the SubNetwork field if non-nil, zero value otherwise.

### GetSubNetworkOk

`func (o *SubNetworkSingle5) GetSubNetworkOk() (*[]SubNetworkSingle, bool)`

GetSubNetworkOk returns a tuple with the SubNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetwork

`func (o *SubNetworkSingle5) SetSubNetwork(v []SubNetworkSingle)`

SetSubNetwork sets SubNetwork field to given value.

### HasSubNetwork

`func (o *SubNetworkSingle5) HasSubNetwork() bool`

HasSubNetwork returns a boolean if a field has been set.

### GetManagedElement

`func (o *SubNetworkSingle5) GetManagedElement() []ManagedElementSingle`

GetManagedElement returns the ManagedElement field if non-nil, zero value otherwise.

### GetManagedElementOk

`func (o *SubNetworkSingle5) GetManagedElementOk() (*[]ManagedElementSingle, bool)`

GetManagedElementOk returns a tuple with the ManagedElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedElement

`func (o *SubNetworkSingle5) SetManagedElement(v []ManagedElementSingle)`

SetManagedElement sets ManagedElement field to given value.

### HasManagedElement

`func (o *SubNetworkSingle5) HasManagedElement() bool`

HasManagedElement returns a boolean if a field has been set.

### GetMLTrainingFunction

`func (o *SubNetworkSingle5) GetMLTrainingFunction() []MLTrainingFunctionSingle`

GetMLTrainingFunction returns the MLTrainingFunction field if non-nil, zero value otherwise.

### GetMLTrainingFunctionOk

`func (o *SubNetworkSingle5) GetMLTrainingFunctionOk() (*[]MLTrainingFunctionSingle, bool)`

GetMLTrainingFunctionOk returns a tuple with the MLTrainingFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLTrainingFunction

`func (o *SubNetworkSingle5) SetMLTrainingFunction(v []MLTrainingFunctionSingle)`

SetMLTrainingFunction sets MLTrainingFunction field to given value.

### HasMLTrainingFunction

`func (o *SubNetworkSingle5) HasMLTrainingFunction() bool`

HasMLTrainingFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


