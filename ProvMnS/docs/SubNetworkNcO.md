# SubNetworkNcO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSubNetworkNcO

`func NewSubNetworkNcO() *SubNetworkNcO`

NewSubNetworkNcO instantiates a new SubNetworkNcO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubNetworkNcOWithDefaults

`func NewSubNetworkNcOWithDefaults() *SubNetworkNcO`

NewSubNetworkNcOWithDefaults instantiates a new SubNetworkNcO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementNode

`func (o *SubNetworkNcO) GetManagementNode() []ManagementNodeSingle`

GetManagementNode returns the ManagementNode field if non-nil, zero value otherwise.

### GetManagementNodeOk

`func (o *SubNetworkNcO) GetManagementNodeOk() (*[]ManagementNodeSingle, bool)`

GetManagementNodeOk returns a tuple with the ManagementNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementNode

`func (o *SubNetworkNcO) SetManagementNode(v []ManagementNodeSingle)`

SetManagementNode sets ManagementNode field to given value.

### HasManagementNode

`func (o *SubNetworkNcO) HasManagementNode() bool`

HasManagementNode returns a boolean if a field has been set.

### GetMnsAgent

`func (o *SubNetworkNcO) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *SubNetworkNcO) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *SubNetworkNcO) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *SubNetworkNcO) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetMeContext

`func (o *SubNetworkNcO) GetMeContext() []MeContextSingle`

GetMeContext returns the MeContext field if non-nil, zero value otherwise.

### GetMeContextOk

`func (o *SubNetworkNcO) GetMeContextOk() (*[]MeContextSingle, bool)`

GetMeContextOk returns a tuple with the MeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeContext

`func (o *SubNetworkNcO) SetMeContext(v []MeContextSingle)`

SetMeContext sets MeContext field to given value.

### HasMeContext

`func (o *SubNetworkNcO) HasMeContext() bool`

HasMeContext returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *SubNetworkNcO) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *SubNetworkNcO) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *SubNetworkNcO) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *SubNetworkNcO) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *SubNetworkNcO) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *SubNetworkNcO) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *SubNetworkNcO) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *SubNetworkNcO) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetTraceJob

`func (o *SubNetworkNcO) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *SubNetworkNcO) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *SubNetworkNcO) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *SubNetworkNcO) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetManagementDataCollection

`func (o *SubNetworkNcO) GetManagementDataCollection() []ManagementDataCollectionSingle`

GetManagementDataCollection returns the ManagementDataCollection field if non-nil, zero value otherwise.

### GetManagementDataCollectionOk

`func (o *SubNetworkNcO) GetManagementDataCollectionOk() (*[]ManagementDataCollectionSingle, bool)`

GetManagementDataCollectionOk returns a tuple with the ManagementDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementDataCollection

`func (o *SubNetworkNcO) SetManagementDataCollection(v []ManagementDataCollectionSingle)`

SetManagementDataCollection sets ManagementDataCollection field to given value.

### HasManagementDataCollection

`func (o *SubNetworkNcO) HasManagementDataCollection() bool`

HasManagementDataCollection returns a boolean if a field has been set.

### GetNtfSubscriptionControl

`func (o *SubNetworkNcO) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle`

GetNtfSubscriptionControl returns the NtfSubscriptionControl field if non-nil, zero value otherwise.

### GetNtfSubscriptionControlOk

`func (o *SubNetworkNcO) GetNtfSubscriptionControlOk() (*[]NtfSubscriptionControlSingle, bool)`

GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfSubscriptionControl

`func (o *SubNetworkNcO) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle)`

SetNtfSubscriptionControl sets NtfSubscriptionControl field to given value.

### HasNtfSubscriptionControl

`func (o *SubNetworkNcO) HasNtfSubscriptionControl() bool`

HasNtfSubscriptionControl returns a boolean if a field has been set.

### GetAlarmList

`func (o *SubNetworkNcO) GetAlarmList() AlarmListSingle`

GetAlarmList returns the AlarmList field if non-nil, zero value otherwise.

### GetAlarmListOk

`func (o *SubNetworkNcO) GetAlarmListOk() (*AlarmListSingle, bool)`

GetAlarmListOk returns a tuple with the AlarmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmList

`func (o *SubNetworkNcO) SetAlarmList(v AlarmListSingle)`

SetAlarmList sets AlarmList field to given value.

### HasAlarmList

`func (o *SubNetworkNcO) HasAlarmList() bool`

HasAlarmList returns a boolean if a field has been set.

### GetFiles

`func (o *SubNetworkNcO) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SubNetworkNcO) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SubNetworkNcO) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SubNetworkNcO) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFileDownloadJob

`func (o *SubNetworkNcO) GetFileDownloadJob() []FileDownloadJobSingle`

GetFileDownloadJob returns the FileDownloadJob field if non-nil, zero value otherwise.

### GetFileDownloadJobOk

`func (o *SubNetworkNcO) GetFileDownloadJobOk() (*[]FileDownloadJobSingle, bool)`

GetFileDownloadJobOk returns a tuple with the FileDownloadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadJob

`func (o *SubNetworkNcO) SetFileDownloadJob(v []FileDownloadJobSingle)`

SetFileDownloadJob sets FileDownloadJob field to given value.

### HasFileDownloadJob

`func (o *SubNetworkNcO) HasFileDownloadJob() bool`

HasFileDownloadJob returns a boolean if a field has been set.

### GetMnsRegistry

`func (o *SubNetworkNcO) GetMnsRegistry() MnsRegistrySingle`

GetMnsRegistry returns the MnsRegistry field if non-nil, zero value otherwise.

### GetMnsRegistryOk

`func (o *SubNetworkNcO) GetMnsRegistryOk() (*MnsRegistrySingle, bool)`

GetMnsRegistryOk returns a tuple with the MnsRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsRegistry

`func (o *SubNetworkNcO) SetMnsRegistry(v MnsRegistrySingle)`

SetMnsRegistry sets MnsRegistry field to given value.

### HasMnsRegistry

`func (o *SubNetworkNcO) HasMnsRegistry() bool`

HasMnsRegistry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


