# SubNetworkSingle1

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
**ManagedElement** | Pointer to [**[]ManagedElementSingle**](ManagedElementSingle.md) |  | [optional] 
**ExternalAmfFunction** | Pointer to [**[]ExternalAmfFunctionSingle**](ExternalAmfFunctionSingle.md) |  | [optional] 
**ExternalNrfFunction** | Pointer to [**[]ExternalNrfFunctionSingle**](ExternalNrfFunctionSingle.md) |  | [optional] 
**ExternalNssfFunction** | Pointer to [**[]ExternalNssfFunctionSingle**](ExternalNssfFunctionSingle.md) |  | [optional] 
**AmfSet** | Pointer to [**[]AmfSetSingle**](AmfSetSingle.md) |  | [optional] 
**AmfRegion** | Pointer to [**[]AmfRegionSingle**](AmfRegionSingle.md) |  | [optional] 
**Configurable5QISet** | Pointer to [**[]Configurable5QISetSingle**](Configurable5QISetSingle.md) |  | [optional] 
**Dynamic5QISet** | Pointer to [**[]Dynamic5QISetSingle**](Dynamic5QISetSingle.md) |  | [optional] 
**EcmConnectionInfo** | Pointer to [**[]EcmConnectionInfoSingle**](EcmConnectionInfoSingle.md) |  | [optional] 

## Methods

### NewSubNetworkSingle1

`func NewSubNetworkSingle1(id NullableString, ) *SubNetworkSingle1`

NewSubNetworkSingle1 instantiates a new SubNetworkSingle1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubNetworkSingle1WithDefaults

`func NewSubNetworkSingle1WithDefaults() *SubNetworkSingle1`

NewSubNetworkSingle1WithDefaults instantiates a new SubNetworkSingle1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubNetworkSingle1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubNetworkSingle1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubNetworkSingle1) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *SubNetworkSingle1) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SubNetworkSingle1) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *SubNetworkSingle1) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *SubNetworkSingle1) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *SubNetworkSingle1) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *SubNetworkSingle1) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *SubNetworkSingle1) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *SubNetworkSingle1) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *SubNetworkSingle1) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *SubNetworkSingle1) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *SubNetworkSingle1) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *SubNetworkSingle1) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *SubNetworkSingle1) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *SubNetworkSingle1) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *SubNetworkSingle1) GetAttributes() SubNetworkAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubNetworkSingle1) GetAttributesOk() (*SubNetworkAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubNetworkSingle1) SetAttributes(v SubNetworkAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SubNetworkSingle1) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetManagementNode

`func (o *SubNetworkSingle1) GetManagementNode() []ManagementNodeSingle`

GetManagementNode returns the ManagementNode field if non-nil, zero value otherwise.

### GetManagementNodeOk

`func (o *SubNetworkSingle1) GetManagementNodeOk() (*[]ManagementNodeSingle, bool)`

GetManagementNodeOk returns a tuple with the ManagementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNode

`func (o *SubNetworkSingle1) SetManagementNode(v []ManagementNodeSingle)`

SetManagementNode sets ManagementNode field to given value.

### HasManagementNode

`func (o *SubNetworkSingle1) HasManagementNode() bool`

HasManagementNode returns a boolean if a field has been set.

### GetMnsAgent

`func (o *SubNetworkSingle1) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *SubNetworkSingle1) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *SubNetworkSingle1) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *SubNetworkSingle1) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetMeContext

`func (o *SubNetworkSingle1) GetMeContext() []MeContextSingle`

GetMeContext returns the MeContext field if non-nil, zero value otherwise.

### GetMeContextOk

`func (o *SubNetworkSingle1) GetMeContextOk() (*[]MeContextSingle, bool)`

GetMeContextOk returns a tuple with the MeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeContext

`func (o *SubNetworkSingle1) SetMeContext(v []MeContextSingle)`

SetMeContext sets MeContext field to given value.

### HasMeContext

`func (o *SubNetworkSingle1) HasMeContext() bool`

HasMeContext returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *SubNetworkSingle1) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *SubNetworkSingle1) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *SubNetworkSingle1) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *SubNetworkSingle1) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *SubNetworkSingle1) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *SubNetworkSingle1) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *SubNetworkSingle1) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *SubNetworkSingle1) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetTraceJob

`func (o *SubNetworkSingle1) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *SubNetworkSingle1) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *SubNetworkSingle1) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *SubNetworkSingle1) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetManagementDataCollection

`func (o *SubNetworkSingle1) GetManagementDataCollection() []ManagementDataCollectionSingle`

GetManagementDataCollection returns the ManagementDataCollection field if non-nil, zero value otherwise.

### GetManagementDataCollectionOk

`func (o *SubNetworkSingle1) GetManagementDataCollectionOk() (*[]ManagementDataCollectionSingle, bool)`

GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementDataCollection

`func (o *SubNetworkSingle1) SetManagementDataCollection(v []ManagementDataCollectionSingle)`

SetManagementDataCollection sets ManagementDataCollection field to given value.

### HasManagementDataCollection

`func (o *SubNetworkSingle1) HasManagementDataCollection() bool`

HasManagementDataCollection returns a boolean if a field has been set.

### GetNtfSubscriptionControl

`func (o *SubNetworkSingle1) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle`

GetNtfSubscriptionControl returns the NtfSubscriptionControl field if non-nil, zero value otherwise.

### GetNtfSubscriptionControlOk

`func (o *SubNetworkSingle1) GetNtfSubscriptionControlOk() (*[]NtfSubscriptionControlSingle, bool)`

GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfSubscriptionControl

`func (o *SubNetworkSingle1) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle)`

SetNtfSubscriptionControl sets NtfSubscriptionControl field to given value.

### HasNtfSubscriptionControl

`func (o *SubNetworkSingle1) HasNtfSubscriptionControl() bool`

HasNtfSubscriptionControl returns a boolean if a field has been set.

### GetAlarmList

`func (o *SubNetworkSingle1) GetAlarmList() AlarmListSingle`

GetAlarmList returns the AlarmList field if non-nil, zero value otherwise.

### GetAlarmListOk

`func (o *SubNetworkSingle1) GetAlarmListOk() (*AlarmListSingle, bool)`

GetAlarmListOk returns a tuple with the AlarmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmList

`func (o *SubNetworkSingle1) SetAlarmList(v AlarmListSingle)`

SetAlarmList sets AlarmList field to given value.

### HasAlarmList

`func (o *SubNetworkSingle1) HasAlarmList() bool`

HasAlarmList returns a boolean if a field has been set.

### GetFiles

`func (o *SubNetworkSingle1) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SubNetworkSingle1) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SubNetworkSingle1) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SubNetworkSingle1) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFileDownloadJob

`func (o *SubNetworkSingle1) GetFileDownloadJob() []FileDownloadJobSingle`

GetFileDownloadJob returns the FileDownloadJob field if non-nil, zero value otherwise.

### GetFileDownloadJobOk

`func (o *SubNetworkSingle1) GetFileDownloadJobOk() (*[]FileDownloadJobSingle, bool)`

GetFileDownloadJobOk returns a tuple with the FileDownloadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadJob

`func (o *SubNetworkSingle1) SetFileDownloadJob(v []FileDownloadJobSingle)`

SetFileDownloadJob sets FileDownloadJob field to given value.

### HasFileDownloadJob

`func (o *SubNetworkSingle1) HasFileDownloadJob() bool`

HasFileDownloadJob returns a boolean if a field has been set.

### GetMnsRegistry

`func (o *SubNetworkSingle1) GetMnsRegistry() MnsRegistrySingle`

GetMnsRegistry returns the MnsRegistry field if non-nil, zero value otherwise.

### GetMnsRegistryOk

`func (o *SubNetworkSingle1) GetMnsRegistryOk() (*MnsRegistrySingle, bool)`

GetMnsRegistryOk returns a tuple with the MnsRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsRegistry

`func (o *SubNetworkSingle1) SetMnsRegistry(v MnsRegistrySingle)`

SetMnsRegistry sets MnsRegistry field to given value.

### HasMnsRegistry

`func (o *SubNetworkSingle1) HasMnsRegistry() bool`

HasMnsRegistry returns a boolean if a field has been set.

### GetSubNetwork

`func (o *SubNetworkSingle1) GetSubNetwork() []SubNetworkSingle`

GetSubNetwork returns the SubNetwork field if non-nil, zero value otherwise.

### GetSubNetworkOk

`func (o *SubNetworkSingle1) GetSubNetworkOk() (*[]SubNetworkSingle, bool)`

GetSubNetworkOk returns a tuple with the SubNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetwork

`func (o *SubNetworkSingle1) SetSubNetwork(v []SubNetworkSingle)`

SetSubNetwork sets SubNetwork field to given value.

### HasSubNetwork

`func (o *SubNetworkSingle1) HasSubNetwork() bool`

HasSubNetwork returns a boolean if a field has been set.

### GetManagedElement

`func (o *SubNetworkSingle1) GetManagedElement() []ManagedElementSingle`

GetManagedElement returns the ManagedElement field if non-nil, zero value otherwise.

### GetManagedElementOk

`func (o *SubNetworkSingle1) GetManagedElementOk() (*[]ManagedElementSingle, bool)`

GetManagedElementOk returns a tuple with the ManagedElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedElement

`func (o *SubNetworkSingle1) SetManagedElement(v []ManagedElementSingle)`

SetManagedElement sets ManagedElement field to given value.

### HasManagedElement

`func (o *SubNetworkSingle1) HasManagedElement() bool`

HasManagedElement returns a boolean if a field has been set.

### GetExternalAmfFunction

`func (o *SubNetworkSingle1) GetExternalAmfFunction() []ExternalAmfFunctionSingle`

GetExternalAmfFunction returns the ExternalAmfFunction field if non-nil, zero value otherwise.

### GetExternalAmfFunctionOk

`func (o *SubNetworkSingle1) GetExternalAmfFunctionOk() (*[]ExternalAmfFunctionSingle, bool)`

GetExternalAmfFunctionOk returns a tuple with the ExternalAmfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAmfFunction

`func (o *SubNetworkSingle1) SetExternalAmfFunction(v []ExternalAmfFunctionSingle)`

SetExternalAmfFunction sets ExternalAmfFunction field to given value.

### HasExternalAmfFunction

`func (o *SubNetworkSingle1) HasExternalAmfFunction() bool`

HasExternalAmfFunction returns a boolean if a field has been set.

### GetExternalNrfFunction

`func (o *SubNetworkSingle1) GetExternalNrfFunction() []ExternalNrfFunctionSingle`

GetExternalNrfFunction returns the ExternalNrfFunction field if non-nil, zero value otherwise.

### GetExternalNrfFunctionOk

`func (o *SubNetworkSingle1) GetExternalNrfFunctionOk() (*[]ExternalNrfFunctionSingle, bool)`

GetExternalNrfFunctionOk returns a tuple with the ExternalNrfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNrfFunction

`func (o *SubNetworkSingle1) SetExternalNrfFunction(v []ExternalNrfFunctionSingle)`

SetExternalNrfFunction sets ExternalNrfFunction field to given value.

### HasExternalNrfFunction

`func (o *SubNetworkSingle1) HasExternalNrfFunction() bool`

HasExternalNrfFunction returns a boolean if a field has been set.

### GetExternalNssfFunction

`func (o *SubNetworkSingle1) GetExternalNssfFunction() []ExternalNssfFunctionSingle`

GetExternalNssfFunction returns the ExternalNssfFunction field if non-nil, zero value otherwise.

### GetExternalNssfFunctionOk

`func (o *SubNetworkSingle1) GetExternalNssfFunctionOk() (*[]ExternalNssfFunctionSingle, bool)`

GetExternalNssfFunctionOk returns a tuple with the ExternalNssfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNssfFunction

`func (o *SubNetworkSingle1) SetExternalNssfFunction(v []ExternalNssfFunctionSingle)`

SetExternalNssfFunction sets ExternalNssfFunction field to given value.

### HasExternalNssfFunction

`func (o *SubNetworkSingle1) HasExternalNssfFunction() bool`

HasExternalNssfFunction returns a boolean if a field has been set.

### GetAmfSet

`func (o *SubNetworkSingle1) GetAmfSet() []AmfSetSingle`

GetAmfSet returns the AmfSet field if non-nil, zero value otherwise.

### GetAmfSetOk

`func (o *SubNetworkSingle1) GetAmfSetOk() (*[]AmfSetSingle, bool)`

GetAmfSetOk returns a tuple with the AmfSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfSet

`func (o *SubNetworkSingle1) SetAmfSet(v []AmfSetSingle)`

SetAmfSet sets AmfSet field to given value.

### HasAmfSet

`func (o *SubNetworkSingle1) HasAmfSet() bool`

HasAmfSet returns a boolean if a field has been set.

### GetAmfRegion

`func (o *SubNetworkSingle1) GetAmfRegion() []AmfRegionSingle`

GetAmfRegion returns the AmfRegion field if non-nil, zero value otherwise.

### GetAmfRegionOk

`func (o *SubNetworkSingle1) GetAmfRegionOk() (*[]AmfRegionSingle, bool)`

GetAmfRegionOk returns a tuple with the AmfRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfRegion

`func (o *SubNetworkSingle1) SetAmfRegion(v []AmfRegionSingle)`

SetAmfRegion sets AmfRegion field to given value.

### HasAmfRegion

`func (o *SubNetworkSingle1) HasAmfRegion() bool`

HasAmfRegion returns a boolean if a field has been set.

### GetConfigurable5QISet

`func (o *SubNetworkSingle1) GetConfigurable5QISet() []Configurable5QISetSingle`

GetConfigurable5QISet returns the Configurable5QISet field if non-nil, zero value otherwise.

### GetConfigurable5QISetOk

`func (o *SubNetworkSingle1) GetConfigurable5QISetOk() (*[]Configurable5QISetSingle, bool)`

GetConfigurable5QISetOk returns a tuple with the Configurable5QISet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable5QISet

`func (o *SubNetworkSingle1) SetConfigurable5QISet(v []Configurable5QISetSingle)`

SetConfigurable5QISet sets Configurable5QISet field to given value.

### HasConfigurable5QISet

`func (o *SubNetworkSingle1) HasConfigurable5QISet() bool`

HasConfigurable5QISet returns a boolean if a field has been set.

### GetDynamic5QISet

`func (o *SubNetworkSingle1) GetDynamic5QISet() []Dynamic5QISetSingle`

GetDynamic5QISet returns the Dynamic5QISet field if non-nil, zero value otherwise.

### GetDynamic5QISetOk

`func (o *SubNetworkSingle1) GetDynamic5QISetOk() (*[]Dynamic5QISetSingle, bool)`

GetDynamic5QISetOk returns a tuple with the Dynamic5QISet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic5QISet

`func (o *SubNetworkSingle1) SetDynamic5QISet(v []Dynamic5QISetSingle)`

SetDynamic5QISet sets Dynamic5QISet field to given value.

### HasDynamic5QISet

`func (o *SubNetworkSingle1) HasDynamic5QISet() bool`

HasDynamic5QISet returns a boolean if a field has been set.

### GetEcmConnectionInfo

`func (o *SubNetworkSingle1) GetEcmConnectionInfo() []EcmConnectionInfoSingle`

GetEcmConnectionInfo returns the EcmConnectionInfo field if non-nil, zero value otherwise.

### GetEcmConnectionInfoOk

`func (o *SubNetworkSingle1) GetEcmConnectionInfoOk() (*[]EcmConnectionInfoSingle, bool)`

GetEcmConnectionInfoOk returns a tuple with the EcmConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcmConnectionInfo

`func (o *SubNetworkSingle1) SetEcmConnectionInfo(v []EcmConnectionInfoSingle)`

SetEcmConnectionInfo sets EcmConnectionInfo field to given value.

### HasEcmConnectionInfo

`func (o *SubNetworkSingle1) HasEcmConnectionInfo() bool`

HasEcmConnectionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


