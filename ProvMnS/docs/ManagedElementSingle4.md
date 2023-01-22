# ManagedElementSingle4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**ManagedElementAttr**](ManagedElementAttr.md) |  | [optional] 
**MnsAgent** | Pointer to [**[]MnsAgentSingle**](MnsAgentSingle.md) |  | [optional] 
**PerfMetricJob** | Pointer to [**[]PerfMetricJobSingle**](PerfMetricJobSingle.md) |  | [optional] 
**ThresholdMonitor** | Pointer to [**[]ThresholdMonitorSingle**](ThresholdMonitorSingle.md) |  | [optional] 
**TraceJob** | Pointer to [**[]TraceJobSingle**](TraceJobSingle.md) |  | [optional] 
**NtfSubscriptionControl** | Pointer to [**[]NtfSubscriptionControlSingle**](NtfSubscriptionControlSingle.md) |  | [optional] 
**AlarmList** | Pointer to [**AlarmListSingle**](AlarmListSingle.md) |  | [optional] 
**FileDownloadJob** | Pointer to [**[]FileDownloadJobSingle**](FileDownloadJobSingle.md) |  | [optional] 
**Files** | Pointer to [**[]FilesSingle**](FilesSingle.md) |  | [optional] 
**MLTrainingFunction** | Pointer to [**[]MLTrainingFunctionSingle**](MLTrainingFunctionSingle.md) |  | [optional] 

## Methods

### NewManagedElementSingle4

`func NewManagedElementSingle4(id NullableString, ) *ManagedElementSingle4`

NewManagedElementSingle4 instantiates a new ManagedElementSingle4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedElementSingle4WithDefaults

`func NewManagedElementSingle4WithDefaults() *ManagedElementSingle4`

NewManagedElementSingle4WithDefaults instantiates a new ManagedElementSingle4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedElementSingle4) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedElementSingle4) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedElementSingle4) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ManagedElementSingle4) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ManagedElementSingle4) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *ManagedElementSingle4) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ManagedElementSingle4) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ManagedElementSingle4) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ManagedElementSingle4) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *ManagedElementSingle4) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *ManagedElementSingle4) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *ManagedElementSingle4) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *ManagedElementSingle4) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *ManagedElementSingle4) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *ManagedElementSingle4) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *ManagedElementSingle4) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *ManagedElementSingle4) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *ManagedElementSingle4) GetAttributes() ManagedElementAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ManagedElementSingle4) GetAttributesOk() (*ManagedElementAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ManagedElementSingle4) SetAttributes(v ManagedElementAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ManagedElementSingle4) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMnsAgent

`func (o *ManagedElementSingle4) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *ManagedElementSingle4) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *ManagedElementSingle4) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *ManagedElementSingle4) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *ManagedElementSingle4) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *ManagedElementSingle4) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *ManagedElementSingle4) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *ManagedElementSingle4) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *ManagedElementSingle4) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *ManagedElementSingle4) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *ManagedElementSingle4) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *ManagedElementSingle4) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetTraceJob

`func (o *ManagedElementSingle4) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *ManagedElementSingle4) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *ManagedElementSingle4) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *ManagedElementSingle4) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetNtfSubscriptionControl

`func (o *ManagedElementSingle4) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle`

GetNtfSubscriptionControl returns the NtfSubscriptionControl field if non-nil, zero value otherwise.

### GetNtfSubscriptionControlOk

`func (o *ManagedElementSingle4) GetNtfSubscriptionControlOk() (*[]NtfSubscriptionControlSingle, bool)`

GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfSubscriptionControl

`func (o *ManagedElementSingle4) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle)`

SetNtfSubscriptionControl sets NtfSubscriptionControl field to given value.

### HasNtfSubscriptionControl

`func (o *ManagedElementSingle4) HasNtfSubscriptionControl() bool`

HasNtfSubscriptionControl returns a boolean if a field has been set.

### GetAlarmList

`func (o *ManagedElementSingle4) GetAlarmList() AlarmListSingle`

GetAlarmList returns the AlarmList field if non-nil, zero value otherwise.

### GetAlarmListOk

`func (o *ManagedElementSingle4) GetAlarmListOk() (*AlarmListSingle, bool)`

GetAlarmListOk returns a tuple with the AlarmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmList

`func (o *ManagedElementSingle4) SetAlarmList(v AlarmListSingle)`

SetAlarmList sets AlarmList field to given value.

### HasAlarmList

`func (o *ManagedElementSingle4) HasAlarmList() bool`

HasAlarmList returns a boolean if a field has been set.

### GetFileDownloadJob

`func (o *ManagedElementSingle4) GetFileDownloadJob() []FileDownloadJobSingle`

GetFileDownloadJob returns the FileDownloadJob field if non-nil, zero value otherwise.

### GetFileDownloadJobOk

`func (o *ManagedElementSingle4) GetFileDownloadJobOk() (*[]FileDownloadJobSingle, bool)`

GetFileDownloadJobOk returns a tuple with the FileDownloadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadJob

`func (o *ManagedElementSingle4) SetFileDownloadJob(v []FileDownloadJobSingle)`

SetFileDownloadJob sets FileDownloadJob field to given value.

### HasFileDownloadJob

`func (o *ManagedElementSingle4) HasFileDownloadJob() bool`

HasFileDownloadJob returns a boolean if a field has been set.

### GetFiles

`func (o *ManagedElementSingle4) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ManagedElementSingle4) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ManagedElementSingle4) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ManagedElementSingle4) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetMLTrainingFunction

`func (o *ManagedElementSingle4) GetMLTrainingFunction() []MLTrainingFunctionSingle`

GetMLTrainingFunction returns the MLTrainingFunction field if non-nil, zero value otherwise.

### GetMLTrainingFunctionOk

`func (o *ManagedElementSingle4) GetMLTrainingFunctionOk() (*[]MLTrainingFunctionSingle, bool)`

GetMLTrainingFunctionOk returns a tuple with the MLTrainingFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLTrainingFunction

`func (o *ManagedElementSingle4) SetMLTrainingFunction(v []MLTrainingFunctionSingle)`

SetMLTrainingFunction sets MLTrainingFunction field to given value.

### HasMLTrainingFunction

`func (o *ManagedElementSingle4) HasMLTrainingFunction() bool`

HasMLTrainingFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


