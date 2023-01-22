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
**GnbDuFunction** | Pointer to [**[]GnbDuFunctionSingle**](GnbDuFunctionSingle.md) |  | [optional] 
**GnbCuUpFunction** | Pointer to [**[]GnbCuUpFunctionSingle**](GnbCuUpFunctionSingle.md) |  | [optional] 
**GnbCuCpFunction** | Pointer to [**[]GnbCuCpFunctionSingle**](GnbCuCpFunctionSingle.md) |  | [optional] 
**DESManagementFunction** | Pointer to [**DESManagementFunctionSingle**](DESManagementFunctionSingle.md) |  | [optional] 
**DRACHOptimizationFunction** | Pointer to [**DRACHOptimizationFunctionSingle**](DRACHOptimizationFunctionSingle.md) |  | [optional] 
**DMROFunction** | Pointer to [**DMROFunctionSingle**](DMROFunctionSingle.md) |  | [optional] 
**DLBOFunction** | Pointer to [**DLBOFunctionSingle**](DLBOFunctionSingle.md) |  | [optional] 
**DPCIConfigurationFunction** | Pointer to [**DPCIConfigurationFunctionSingle**](DPCIConfigurationFunctionSingle.md) |  | [optional] 
**CPCIConfigurationFunction** | Pointer to [**CPCIConfigurationFunctionSingle**](CPCIConfigurationFunctionSingle.md) |  | [optional] 
**CESManagementFunction** | Pointer to [**CESManagementFunctionSingle**](CESManagementFunctionSingle.md) |  | [optional] 
**Configurable5QISet** | Pointer to [**[]Configurable5QISetSingle**](Configurable5QISetSingle.md) |  | [optional] 
**Dynamic5QISet** | Pointer to [**[]Dynamic5QISetSingle**](Dynamic5QISetSingle.md) |  | [optional] 

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

### GetGnbDuFunction

`func (o *ManagedElementSingle) GetGnbDuFunction() []GnbDuFunctionSingle`

GetGnbDuFunction returns the GnbDuFunction field if non-nil, zero value otherwise.

### GetGnbDuFunctionOk

`func (o *ManagedElementSingle) GetGnbDuFunctionOk() (*[]GnbDuFunctionSingle, bool)`

GetGnbDuFunctionOk returns a tuple with the GnbDuFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbDuFunction

`func (o *ManagedElementSingle) SetGnbDuFunction(v []GnbDuFunctionSingle)`

SetGnbDuFunction sets GnbDuFunction field to given value.

### HasGnbDuFunction

`func (o *ManagedElementSingle) HasGnbDuFunction() bool`

HasGnbDuFunction returns a boolean if a field has been set.

### GetGnbCuUpFunction

`func (o *ManagedElementSingle) GetGnbCuUpFunction() []GnbCuUpFunctionSingle`

GetGnbCuUpFunction returns the GnbCuUpFunction field if non-nil, zero value otherwise.

### GetGnbCuUpFunctionOk

`func (o *ManagedElementSingle) GetGnbCuUpFunctionOk() (*[]GnbCuUpFunctionSingle, bool)`

GetGnbCuUpFunctionOk returns a tuple with the GnbCuUpFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbCuUpFunction

`func (o *ManagedElementSingle) SetGnbCuUpFunction(v []GnbCuUpFunctionSingle)`

SetGnbCuUpFunction sets GnbCuUpFunction field to given value.

### HasGnbCuUpFunction

`func (o *ManagedElementSingle) HasGnbCuUpFunction() bool`

HasGnbCuUpFunction returns a boolean if a field has been set.

### GetGnbCuCpFunction

`func (o *ManagedElementSingle) GetGnbCuCpFunction() []GnbCuCpFunctionSingle`

GetGnbCuCpFunction returns the GnbCuCpFunction field if non-nil, zero value otherwise.

### GetGnbCuCpFunctionOk

`func (o *ManagedElementSingle) GetGnbCuCpFunctionOk() (*[]GnbCuCpFunctionSingle, bool)`

GetGnbCuCpFunctionOk returns a tuple with the GnbCuCpFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbCuCpFunction

`func (o *ManagedElementSingle) SetGnbCuCpFunction(v []GnbCuCpFunctionSingle)`

SetGnbCuCpFunction sets GnbCuCpFunction field to given value.

### HasGnbCuCpFunction

`func (o *ManagedElementSingle) HasGnbCuCpFunction() bool`

HasGnbCuCpFunction returns a boolean if a field has been set.

### GetDESManagementFunction

`func (o *ManagedElementSingle) GetDESManagementFunction() DESManagementFunctionSingle`

GetDESManagementFunction returns the DESManagementFunction field if non-nil, zero value otherwise.

### GetDESManagementFunctionOk

`func (o *ManagedElementSingle) GetDESManagementFunctionOk() (*DESManagementFunctionSingle, bool)`

GetDESManagementFunctionOk returns a tuple with the DESManagementFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDESManagementFunction

`func (o *ManagedElementSingle) SetDESManagementFunction(v DESManagementFunctionSingle)`

SetDESManagementFunction sets DESManagementFunction field to given value.

### HasDESManagementFunction

`func (o *ManagedElementSingle) HasDESManagementFunction() bool`

HasDESManagementFunction returns a boolean if a field has been set.

### GetDRACHOptimizationFunction

`func (o *ManagedElementSingle) GetDRACHOptimizationFunction() DRACHOptimizationFunctionSingle`

GetDRACHOptimizationFunction returns the DRACHOptimizationFunction field if non-nil, zero value otherwise.

### GetDRACHOptimizationFunctionOk

`func (o *ManagedElementSingle) GetDRACHOptimizationFunctionOk() (*DRACHOptimizationFunctionSingle, bool)`

GetDRACHOptimizationFunctionOk returns a tuple with the DRACHOptimizationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDRACHOptimizationFunction

`func (o *ManagedElementSingle) SetDRACHOptimizationFunction(v DRACHOptimizationFunctionSingle)`

SetDRACHOptimizationFunction sets DRACHOptimizationFunction field to given value.

### HasDRACHOptimizationFunction

`func (o *ManagedElementSingle) HasDRACHOptimizationFunction() bool`

HasDRACHOptimizationFunction returns a boolean if a field has been set.

### GetDMROFunction

`func (o *ManagedElementSingle) GetDMROFunction() DMROFunctionSingle`

GetDMROFunction returns the DMROFunction field if non-nil, zero value otherwise.

### GetDMROFunctionOk

`func (o *ManagedElementSingle) GetDMROFunctionOk() (*DMROFunctionSingle, bool)`

GetDMROFunctionOk returns a tuple with the DMROFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDMROFunction

`func (o *ManagedElementSingle) SetDMROFunction(v DMROFunctionSingle)`

SetDMROFunction sets DMROFunction field to given value.

### HasDMROFunction

`func (o *ManagedElementSingle) HasDMROFunction() bool`

HasDMROFunction returns a boolean if a field has been set.

### GetDLBOFunction

`func (o *ManagedElementSingle) GetDLBOFunction() DLBOFunctionSingle`

GetDLBOFunction returns the DLBOFunction field if non-nil, zero value otherwise.

### GetDLBOFunctionOk

`func (o *ManagedElementSingle) GetDLBOFunctionOk() (*DLBOFunctionSingle, bool)`

GetDLBOFunctionOk returns a tuple with the DLBOFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLBOFunction

`func (o *ManagedElementSingle) SetDLBOFunction(v DLBOFunctionSingle)`

SetDLBOFunction sets DLBOFunction field to given value.

### HasDLBOFunction

`func (o *ManagedElementSingle) HasDLBOFunction() bool`

HasDLBOFunction returns a boolean if a field has been set.

### GetDPCIConfigurationFunction

`func (o *ManagedElementSingle) GetDPCIConfigurationFunction() DPCIConfigurationFunctionSingle`

GetDPCIConfigurationFunction returns the DPCIConfigurationFunction field if non-nil, zero value otherwise.

### GetDPCIConfigurationFunctionOk

`func (o *ManagedElementSingle) GetDPCIConfigurationFunctionOk() (*DPCIConfigurationFunctionSingle, bool)`

GetDPCIConfigurationFunctionOk returns a tuple with the DPCIConfigurationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDPCIConfigurationFunction

`func (o *ManagedElementSingle) SetDPCIConfigurationFunction(v DPCIConfigurationFunctionSingle)`

SetDPCIConfigurationFunction sets DPCIConfigurationFunction field to given value.

### HasDPCIConfigurationFunction

`func (o *ManagedElementSingle) HasDPCIConfigurationFunction() bool`

HasDPCIConfigurationFunction returns a boolean if a field has been set.

### GetCPCIConfigurationFunction

`func (o *ManagedElementSingle) GetCPCIConfigurationFunction() CPCIConfigurationFunctionSingle`

GetCPCIConfigurationFunction returns the CPCIConfigurationFunction field if non-nil, zero value otherwise.

### GetCPCIConfigurationFunctionOk

`func (o *ManagedElementSingle) GetCPCIConfigurationFunctionOk() (*CPCIConfigurationFunctionSingle, bool)`

GetCPCIConfigurationFunctionOk returns a tuple with the CPCIConfigurationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPCIConfigurationFunction

`func (o *ManagedElementSingle) SetCPCIConfigurationFunction(v CPCIConfigurationFunctionSingle)`

SetCPCIConfigurationFunction sets CPCIConfigurationFunction field to given value.

### HasCPCIConfigurationFunction

`func (o *ManagedElementSingle) HasCPCIConfigurationFunction() bool`

HasCPCIConfigurationFunction returns a boolean if a field has been set.

### GetCESManagementFunction

`func (o *ManagedElementSingle) GetCESManagementFunction() CESManagementFunctionSingle`

GetCESManagementFunction returns the CESManagementFunction field if non-nil, zero value otherwise.

### GetCESManagementFunctionOk

`func (o *ManagedElementSingle) GetCESManagementFunctionOk() (*CESManagementFunctionSingle, bool)`

GetCESManagementFunctionOk returns a tuple with the CESManagementFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCESManagementFunction

`func (o *ManagedElementSingle) SetCESManagementFunction(v CESManagementFunctionSingle)`

SetCESManagementFunction sets CESManagementFunction field to given value.

### HasCESManagementFunction

`func (o *ManagedElementSingle) HasCESManagementFunction() bool`

HasCESManagementFunction returns a boolean if a field has been set.

### GetConfigurable5QISet

`func (o *ManagedElementSingle) GetConfigurable5QISet() []Configurable5QISetSingle`

GetConfigurable5QISet returns the Configurable5QISet field if non-nil, zero value otherwise.

### GetConfigurable5QISetOk

`func (o *ManagedElementSingle) GetConfigurable5QISetOk() (*[]Configurable5QISetSingle, bool)`

GetConfigurable5QISetOk returns a tuple with the Configurable5QISet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable5QISet

`func (o *ManagedElementSingle) SetConfigurable5QISet(v []Configurable5QISetSingle)`

SetConfigurable5QISet sets Configurable5QISet field to given value.

### HasConfigurable5QISet

`func (o *ManagedElementSingle) HasConfigurable5QISet() bool`

HasConfigurable5QISet returns a boolean if a field has been set.

### GetDynamic5QISet

`func (o *ManagedElementSingle) GetDynamic5QISet() []Dynamic5QISetSingle`

GetDynamic5QISet returns the Dynamic5QISet field if non-nil, zero value otherwise.

### GetDynamic5QISetOk

`func (o *ManagedElementSingle) GetDynamic5QISetOk() (*[]Dynamic5QISetSingle, bool)`

GetDynamic5QISetOk returns a tuple with the Dynamic5QISet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic5QISet

`func (o *ManagedElementSingle) SetDynamic5QISet(v []Dynamic5QISetSingle)`

SetDynamic5QISet sets Dynamic5QISet field to given value.

### HasDynamic5QISet

`func (o *ManagedElementSingle) HasDynamic5QISet() bool`

HasDynamic5QISet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


