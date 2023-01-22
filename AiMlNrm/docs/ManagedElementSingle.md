# ManagedElementSingle

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

### NewManagedElementSingle

`func NewManagedElementSingle(id NullableString, ) *ManagedElementSingle`

NewManagedElementSingle instantiates a new ManagedElementSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedElementSingleWithDefaults

`func NewManagedElementSingleWithDefaults() *ManagedElementSingle`

NewManagedElementSingleWithDefaults instantiates a new ManagedElementSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedElementSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedElementSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedElementSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ManagedElementSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ManagedElementSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *ManagedElementSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ManagedElementSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ManagedElementSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ManagedElementSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *ManagedElementSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *ManagedElementSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *ManagedElementSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *ManagedElementSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *ManagedElementSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *ManagedElementSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *ManagedElementSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *ManagedElementSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *ManagedElementSingle) GetAttributes() ManagedElementAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ManagedElementSingle) GetAttributesOk() (*ManagedElementAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ManagedElementSingle) SetAttributes(v ManagedElementAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ManagedElementSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMnsAgent

`func (o *ManagedElementSingle) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *ManagedElementSingle) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *ManagedElementSingle) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *ManagedElementSingle) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *ManagedElementSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *ManagedElementSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *ManagedElementSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *ManagedElementSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *ManagedElementSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *ManagedElementSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *ManagedElementSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *ManagedElementSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetTraceJob

`func (o *ManagedElementSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *ManagedElementSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *ManagedElementSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *ManagedElementSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetNtfSubscriptionControl

`func (o *ManagedElementSingle) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle`

GetNtfSubscriptionControl returns the NtfSubscriptionControl field if non-nil, zero value otherwise.

### GetNtfSubscriptionControlOk

`func (o *ManagedElementSingle) GetNtfSubscriptionControlOk() (*[]NtfSubscriptionControlSingle, bool)`

GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfSubscriptionControl

`func (o *ManagedElementSingle) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle)`

SetNtfSubscriptionControl sets NtfSubscriptionControl field to given value.

### HasNtfSubscriptionControl

`func (o *ManagedElementSingle) HasNtfSubscriptionControl() bool`

HasNtfSubscriptionControl returns a boolean if a field has been set.

### GetAlarmList

`func (o *ManagedElementSingle) GetAlarmList() AlarmListSingle`

GetAlarmList returns the AlarmList field if non-nil, zero value otherwise.

### GetAlarmListOk

`func (o *ManagedElementSingle) GetAlarmListOk() (*AlarmListSingle, bool)`

GetAlarmListOk returns a tuple with the AlarmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmList

`func (o *ManagedElementSingle) SetAlarmList(v AlarmListSingle)`

SetAlarmList sets AlarmList field to given value.

### HasAlarmList

`func (o *ManagedElementSingle) HasAlarmList() bool`

HasAlarmList returns a boolean if a field has been set.

### GetFileDownloadJob

`func (o *ManagedElementSingle) GetFileDownloadJob() []FileDownloadJobSingle`

GetFileDownloadJob returns the FileDownloadJob field if non-nil, zero value otherwise.

### GetFileDownloadJobOk

`func (o *ManagedElementSingle) GetFileDownloadJobOk() (*[]FileDownloadJobSingle, bool)`

GetFileDownloadJobOk returns a tuple with the FileDownloadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadJob

`func (o *ManagedElementSingle) SetFileDownloadJob(v []FileDownloadJobSingle)`

SetFileDownloadJob sets FileDownloadJob field to given value.

### HasFileDownloadJob

`func (o *ManagedElementSingle) HasFileDownloadJob() bool`

HasFileDownloadJob returns a boolean if a field has been set.

### GetFiles

`func (o *ManagedElementSingle) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ManagedElementSingle) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ManagedElementSingle) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ManagedElementSingle) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetMLTrainingFunction

`func (o *ManagedElementSingle) GetMLTrainingFunction() []MLTrainingFunctionSingle`

GetMLTrainingFunction returns the MLTrainingFunction field if non-nil, zero value otherwise.

### GetMLTrainingFunctionOk

`func (o *ManagedElementSingle) GetMLTrainingFunctionOk() (*[]MLTrainingFunctionSingle, bool)`

GetMLTrainingFunctionOk returns a tuple with the MLTrainingFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLTrainingFunction

`func (o *ManagedElementSingle) SetMLTrainingFunction(v []MLTrainingFunctionSingle)`

SetMLTrainingFunction sets MLTrainingFunction field to given value.

### HasMLTrainingFunction

`func (o *ManagedElementSingle) HasMLTrainingFunction() bool`

HasMLTrainingFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


