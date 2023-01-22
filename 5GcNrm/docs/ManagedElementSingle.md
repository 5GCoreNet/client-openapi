# ManagedElementSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**ManagedElementAttr**](ManagedElement-Attr.md) |  | [optional] 
**MnsAgent** | Pointer to [**[]MnsAgentSingle**](MnsAgentSingle.md) |  | [optional] 
**PerfMetricJob** | Pointer to [**[]PerfMetricJobSingle**](PerfMetricJobSingle.md) |  | [optional] 
**ThresholdMonitor** | Pointer to [**[]ThresholdMonitorSingle**](ThresholdMonitorSingle.md) |  | [optional] 
**TraceJob** | Pointer to [**[]TraceJobSingle**](TraceJobSingle.md) |  | [optional] 
**NtfSubscriptionControl** | Pointer to [**[]NtfSubscriptionControlSingle**](NtfSubscriptionControlSingle.md) |  | [optional] 
**AlarmList** | Pointer to [**AlarmListSingle**](AlarmListSingle.md) |  | [optional] 
**FileDownloadJob** | Pointer to [**[]FileDownloadJobSingle**](FileDownloadJobSingle.md) |  | [optional] 
**Files** | Pointer to [**[]FilesSingle**](FilesSingle.md) |  | [optional] 
**AmfFunction** | Pointer to [**[]AmfFunctionSingle**](AmfFunctionSingle.md) |  | [optional] 
**SmfFunction** | Pointer to [**[]SmfFunctionSingle**](SmfFunctionSingle.md) |  | [optional] 
**UpfFunction** | Pointer to [**[]UpfFunctionSingle**](UpfFunctionSingle.md) |  | [optional] 
**N3iwfFunction** | Pointer to [**[]N3iwfFunctionSingle**](N3iwfFunctionSingle.md) |  | [optional] 
**PcfFunction** | Pointer to [**[]PcfFunctionSingle**](PcfFunctionSingle.md) |  | [optional] 
**AusfFunction** | Pointer to [**[]AusfFunctionSingle**](AusfFunctionSingle.md) |  | [optional] 
**UdmFunction** | Pointer to [**[]UdmFunctionSingle**](UdmFunctionSingle.md) |  | [optional] 
**UdrFunction** | Pointer to [**[]UdrFunctionSingle**](UdrFunctionSingle.md) |  | [optional] 
**UdsfFunction** | Pointer to [**[]UdsfFunctionSingle**](UdsfFunctionSingle.md) |  | [optional] 
**NrfFunction** | Pointer to [**[]NrfFunctionSingle**](NrfFunctionSingle.md) |  | [optional] 
**NssfFunction** | Pointer to [**[]NssfFunctionSingle**](NssfFunctionSingle.md) |  | [optional] 
**SmsfFunction** | Pointer to [**[]SmsfFunctionSingle**](SmsfFunctionSingle.md) |  | [optional] 
**LmfFunction** | Pointer to [**[]LmfFunctionSingle**](LmfFunctionSingle.md) |  | [optional] 
**NgeirFunction** | Pointer to [**[]NgeirFunctionSingle**](NgeirFunctionSingle.md) |  | [optional] 
**SeppFunction** | Pointer to [**[]SeppFunctionSingle**](SeppFunctionSingle.md) |  | [optional] 
**NwdafFunction** | Pointer to [**[]NwdafFunctionSingle**](NwdafFunctionSingle.md) |  | [optional] 
**ScpFunction** | Pointer to [**[]ScpFunctionSingle**](ScpFunctionSingle.md) |  | [optional] 
**NefFunction** | Pointer to [**[]NefFunctionSingle**](NefFunctionSingle.md) |  | [optional] 
**Configurable5QISet** | Pointer to [**[]Configurable5QISetSingle**](Configurable5QISetSingle.md) |  | [optional] 
**Dynamic5QISet** | Pointer to [**[]Dynamic5QISetSingle**](Dynamic5QISetSingle.md) |  | [optional] 
**EcmConnectionInfo** | Pointer to [**[]EcmConnectionInfoSingle**](EcmConnectionInfoSingle.md) |  | [optional] 
**EASDFFunction** | Pointer to [**[]EASDFFunctionSingle**](EASDFFunctionSingle.md) |  | [optional] 

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

### GetAmfFunction

`func (o *ManagedElementSingle) GetAmfFunction() []AmfFunctionSingle`

GetAmfFunction returns the AmfFunction field if non-nil, zero value otherwise.

### GetAmfFunctionOk

`func (o *ManagedElementSingle) GetAmfFunctionOk() (*[]AmfFunctionSingle, bool)`

GetAmfFunctionOk returns a tuple with the AmfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfFunction

`func (o *ManagedElementSingle) SetAmfFunction(v []AmfFunctionSingle)`

SetAmfFunction sets AmfFunction field to given value.

### HasAmfFunction

`func (o *ManagedElementSingle) HasAmfFunction() bool`

HasAmfFunction returns a boolean if a field has been set.

### GetSmfFunction

`func (o *ManagedElementSingle) GetSmfFunction() []SmfFunctionSingle`

GetSmfFunction returns the SmfFunction field if non-nil, zero value otherwise.

### GetSmfFunctionOk

`func (o *ManagedElementSingle) GetSmfFunctionOk() (*[]SmfFunctionSingle, bool)`

GetSmfFunctionOk returns a tuple with the SmfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfFunction

`func (o *ManagedElementSingle) SetSmfFunction(v []SmfFunctionSingle)`

SetSmfFunction sets SmfFunction field to given value.

### HasSmfFunction

`func (o *ManagedElementSingle) HasSmfFunction() bool`

HasSmfFunction returns a boolean if a field has been set.

### GetUpfFunction

`func (o *ManagedElementSingle) GetUpfFunction() []UpfFunctionSingle`

GetUpfFunction returns the UpfFunction field if non-nil, zero value otherwise.

### GetUpfFunctionOk

`func (o *ManagedElementSingle) GetUpfFunctionOk() (*[]UpfFunctionSingle, bool)`

GetUpfFunctionOk returns a tuple with the UpfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfFunction

`func (o *ManagedElementSingle) SetUpfFunction(v []UpfFunctionSingle)`

SetUpfFunction sets UpfFunction field to given value.

### HasUpfFunction

`func (o *ManagedElementSingle) HasUpfFunction() bool`

HasUpfFunction returns a boolean if a field has been set.

### GetN3iwfFunction

`func (o *ManagedElementSingle) GetN3iwfFunction() []N3iwfFunctionSingle`

GetN3iwfFunction returns the N3iwfFunction field if non-nil, zero value otherwise.

### GetN3iwfFunctionOk

`func (o *ManagedElementSingle) GetN3iwfFunctionOk() (*[]N3iwfFunctionSingle, bool)`

GetN3iwfFunctionOk returns a tuple with the N3iwfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3iwfFunction

`func (o *ManagedElementSingle) SetN3iwfFunction(v []N3iwfFunctionSingle)`

SetN3iwfFunction sets N3iwfFunction field to given value.

### HasN3iwfFunction

`func (o *ManagedElementSingle) HasN3iwfFunction() bool`

HasN3iwfFunction returns a boolean if a field has been set.

### GetPcfFunction

`func (o *ManagedElementSingle) GetPcfFunction() []PcfFunctionSingle`

GetPcfFunction returns the PcfFunction field if non-nil, zero value otherwise.

### GetPcfFunctionOk

`func (o *ManagedElementSingle) GetPcfFunctionOk() (*[]PcfFunctionSingle, bool)`

GetPcfFunctionOk returns a tuple with the PcfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfFunction

`func (o *ManagedElementSingle) SetPcfFunction(v []PcfFunctionSingle)`

SetPcfFunction sets PcfFunction field to given value.

### HasPcfFunction

`func (o *ManagedElementSingle) HasPcfFunction() bool`

HasPcfFunction returns a boolean if a field has been set.

### GetAusfFunction

`func (o *ManagedElementSingle) GetAusfFunction() []AusfFunctionSingle`

GetAusfFunction returns the AusfFunction field if non-nil, zero value otherwise.

### GetAusfFunctionOk

`func (o *ManagedElementSingle) GetAusfFunctionOk() (*[]AusfFunctionSingle, bool)`

GetAusfFunctionOk returns a tuple with the AusfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAusfFunction

`func (o *ManagedElementSingle) SetAusfFunction(v []AusfFunctionSingle)`

SetAusfFunction sets AusfFunction field to given value.

### HasAusfFunction

`func (o *ManagedElementSingle) HasAusfFunction() bool`

HasAusfFunction returns a boolean if a field has been set.

### GetUdmFunction

`func (o *ManagedElementSingle) GetUdmFunction() []UdmFunctionSingle`

GetUdmFunction returns the UdmFunction field if non-nil, zero value otherwise.

### GetUdmFunctionOk

`func (o *ManagedElementSingle) GetUdmFunctionOk() (*[]UdmFunctionSingle, bool)`

GetUdmFunctionOk returns a tuple with the UdmFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmFunction

`func (o *ManagedElementSingle) SetUdmFunction(v []UdmFunctionSingle)`

SetUdmFunction sets UdmFunction field to given value.

### HasUdmFunction

`func (o *ManagedElementSingle) HasUdmFunction() bool`

HasUdmFunction returns a boolean if a field has been set.

### GetUdrFunction

`func (o *ManagedElementSingle) GetUdrFunction() []UdrFunctionSingle`

GetUdrFunction returns the UdrFunction field if non-nil, zero value otherwise.

### GetUdrFunctionOk

`func (o *ManagedElementSingle) GetUdrFunctionOk() (*[]UdrFunctionSingle, bool)`

GetUdrFunctionOk returns a tuple with the UdrFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrFunction

`func (o *ManagedElementSingle) SetUdrFunction(v []UdrFunctionSingle)`

SetUdrFunction sets UdrFunction field to given value.

### HasUdrFunction

`func (o *ManagedElementSingle) HasUdrFunction() bool`

HasUdrFunction returns a boolean if a field has been set.

### GetUdsfFunction

`func (o *ManagedElementSingle) GetUdsfFunction() []UdsfFunctionSingle`

GetUdsfFunction returns the UdsfFunction field if non-nil, zero value otherwise.

### GetUdsfFunctionOk

`func (o *ManagedElementSingle) GetUdsfFunctionOk() (*[]UdsfFunctionSingle, bool)`

GetUdsfFunctionOk returns a tuple with the UdsfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdsfFunction

`func (o *ManagedElementSingle) SetUdsfFunction(v []UdsfFunctionSingle)`

SetUdsfFunction sets UdsfFunction field to given value.

### HasUdsfFunction

`func (o *ManagedElementSingle) HasUdsfFunction() bool`

HasUdsfFunction returns a boolean if a field has been set.

### GetNrfFunction

`func (o *ManagedElementSingle) GetNrfFunction() []NrfFunctionSingle`

GetNrfFunction returns the NrfFunction field if non-nil, zero value otherwise.

### GetNrfFunctionOk

`func (o *ManagedElementSingle) GetNrfFunctionOk() (*[]NrfFunctionSingle, bool)`

GetNrfFunctionOk returns a tuple with the NrfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfFunction

`func (o *ManagedElementSingle) SetNrfFunction(v []NrfFunctionSingle)`

SetNrfFunction sets NrfFunction field to given value.

### HasNrfFunction

`func (o *ManagedElementSingle) HasNrfFunction() bool`

HasNrfFunction returns a boolean if a field has been set.

### GetNssfFunction

`func (o *ManagedElementSingle) GetNssfFunction() []NssfFunctionSingle`

GetNssfFunction returns the NssfFunction field if non-nil, zero value otherwise.

### GetNssfFunctionOk

`func (o *ManagedElementSingle) GetNssfFunctionOk() (*[]NssfFunctionSingle, bool)`

GetNssfFunctionOk returns a tuple with the NssfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssfFunction

`func (o *ManagedElementSingle) SetNssfFunction(v []NssfFunctionSingle)`

SetNssfFunction sets NssfFunction field to given value.

### HasNssfFunction

`func (o *ManagedElementSingle) HasNssfFunction() bool`

HasNssfFunction returns a boolean if a field has been set.

### GetSmsfFunction

`func (o *ManagedElementSingle) GetSmsfFunction() []SmsfFunctionSingle`

GetSmsfFunction returns the SmsfFunction field if non-nil, zero value otherwise.

### GetSmsfFunctionOk

`func (o *ManagedElementSingle) GetSmsfFunctionOk() (*[]SmsfFunctionSingle, bool)`

GetSmsfFunctionOk returns a tuple with the SmsfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfFunction

`func (o *ManagedElementSingle) SetSmsfFunction(v []SmsfFunctionSingle)`

SetSmsfFunction sets SmsfFunction field to given value.

### HasSmsfFunction

`func (o *ManagedElementSingle) HasSmsfFunction() bool`

HasSmsfFunction returns a boolean if a field has been set.

### GetLmfFunction

`func (o *ManagedElementSingle) GetLmfFunction() []LmfFunctionSingle`

GetLmfFunction returns the LmfFunction field if non-nil, zero value otherwise.

### GetLmfFunctionOk

`func (o *ManagedElementSingle) GetLmfFunctionOk() (*[]LmfFunctionSingle, bool)`

GetLmfFunctionOk returns a tuple with the LmfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmfFunction

`func (o *ManagedElementSingle) SetLmfFunction(v []LmfFunctionSingle)`

SetLmfFunction sets LmfFunction field to given value.

### HasLmfFunction

`func (o *ManagedElementSingle) HasLmfFunction() bool`

HasLmfFunction returns a boolean if a field has been set.

### GetNgeirFunction

`func (o *ManagedElementSingle) GetNgeirFunction() []NgeirFunctionSingle`

GetNgeirFunction returns the NgeirFunction field if non-nil, zero value otherwise.

### GetNgeirFunctionOk

`func (o *ManagedElementSingle) GetNgeirFunctionOk() (*[]NgeirFunctionSingle, bool)`

GetNgeirFunctionOk returns a tuple with the NgeirFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgeirFunction

`func (o *ManagedElementSingle) SetNgeirFunction(v []NgeirFunctionSingle)`

SetNgeirFunction sets NgeirFunction field to given value.

### HasNgeirFunction

`func (o *ManagedElementSingle) HasNgeirFunction() bool`

HasNgeirFunction returns a boolean if a field has been set.

### GetSeppFunction

`func (o *ManagedElementSingle) GetSeppFunction() []SeppFunctionSingle`

GetSeppFunction returns the SeppFunction field if non-nil, zero value otherwise.

### GetSeppFunctionOk

`func (o *ManagedElementSingle) GetSeppFunctionOk() (*[]SeppFunctionSingle, bool)`

GetSeppFunctionOk returns a tuple with the SeppFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeppFunction

`func (o *ManagedElementSingle) SetSeppFunction(v []SeppFunctionSingle)`

SetSeppFunction sets SeppFunction field to given value.

### HasSeppFunction

`func (o *ManagedElementSingle) HasSeppFunction() bool`

HasSeppFunction returns a boolean if a field has been set.

### GetNwdafFunction

`func (o *ManagedElementSingle) GetNwdafFunction() []NwdafFunctionSingle`

GetNwdafFunction returns the NwdafFunction field if non-nil, zero value otherwise.

### GetNwdafFunctionOk

`func (o *ManagedElementSingle) GetNwdafFunctionOk() (*[]NwdafFunctionSingle, bool)`

GetNwdafFunctionOk returns a tuple with the NwdafFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafFunction

`func (o *ManagedElementSingle) SetNwdafFunction(v []NwdafFunctionSingle)`

SetNwdafFunction sets NwdafFunction field to given value.

### HasNwdafFunction

`func (o *ManagedElementSingle) HasNwdafFunction() bool`

HasNwdafFunction returns a boolean if a field has been set.

### GetScpFunction

`func (o *ManagedElementSingle) GetScpFunction() []ScpFunctionSingle`

GetScpFunction returns the ScpFunction field if non-nil, zero value otherwise.

### GetScpFunctionOk

`func (o *ManagedElementSingle) GetScpFunctionOk() (*[]ScpFunctionSingle, bool)`

GetScpFunctionOk returns a tuple with the ScpFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpFunction

`func (o *ManagedElementSingle) SetScpFunction(v []ScpFunctionSingle)`

SetScpFunction sets ScpFunction field to given value.

### HasScpFunction

`func (o *ManagedElementSingle) HasScpFunction() bool`

HasScpFunction returns a boolean if a field has been set.

### GetNefFunction

`func (o *ManagedElementSingle) GetNefFunction() []NefFunctionSingle`

GetNefFunction returns the NefFunction field if non-nil, zero value otherwise.

### GetNefFunctionOk

`func (o *ManagedElementSingle) GetNefFunctionOk() (*[]NefFunctionSingle, bool)`

GetNefFunctionOk returns a tuple with the NefFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefFunction

`func (o *ManagedElementSingle) SetNefFunction(v []NefFunctionSingle)`

SetNefFunction sets NefFunction field to given value.

### HasNefFunction

`func (o *ManagedElementSingle) HasNefFunction() bool`

HasNefFunction returns a boolean if a field has been set.

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

### GetEcmConnectionInfo

`func (o *ManagedElementSingle) GetEcmConnectionInfo() []EcmConnectionInfoSingle`

GetEcmConnectionInfo returns the EcmConnectionInfo field if non-nil, zero value otherwise.

### GetEcmConnectionInfoOk

`func (o *ManagedElementSingle) GetEcmConnectionInfoOk() (*[]EcmConnectionInfoSingle, bool)`

GetEcmConnectionInfoOk returns a tuple with the EcmConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcmConnectionInfo

`func (o *ManagedElementSingle) SetEcmConnectionInfo(v []EcmConnectionInfoSingle)`

SetEcmConnectionInfo sets EcmConnectionInfo field to given value.

### HasEcmConnectionInfo

`func (o *ManagedElementSingle) HasEcmConnectionInfo() bool`

HasEcmConnectionInfo returns a boolean if a field has been set.

### GetEASDFFunction

`func (o *ManagedElementSingle) GetEASDFFunction() []EASDFFunctionSingle`

GetEASDFFunction returns the EASDFFunction field if non-nil, zero value otherwise.

### GetEASDFFunctionOk

`func (o *ManagedElementSingle) GetEASDFFunctionOk() (*[]EASDFFunctionSingle, bool)`

GetEASDFFunctionOk returns a tuple with the EASDFFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEASDFFunction

`func (o *ManagedElementSingle) SetEASDFFunction(v []EASDFFunctionSingle)`

SetEASDFFunction sets EASDFFunction field to given value.

### HasEASDFFunction

`func (o *ManagedElementSingle) HasEASDFFunction() bool`

HasEASDFFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


