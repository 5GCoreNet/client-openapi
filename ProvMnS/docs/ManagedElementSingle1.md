# ManagedElementSingle1

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

### NewManagedElementSingle1

`func NewManagedElementSingle1(id NullableString, ) *ManagedElementSingle1`

NewManagedElementSingle1 instantiates a new ManagedElementSingle1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedElementSingle1WithDefaults

`func NewManagedElementSingle1WithDefaults() *ManagedElementSingle1`

NewManagedElementSingle1WithDefaults instantiates a new ManagedElementSingle1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedElementSingle1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedElementSingle1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedElementSingle1) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ManagedElementSingle1) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ManagedElementSingle1) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *ManagedElementSingle1) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ManagedElementSingle1) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ManagedElementSingle1) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ManagedElementSingle1) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *ManagedElementSingle1) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *ManagedElementSingle1) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *ManagedElementSingle1) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *ManagedElementSingle1) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *ManagedElementSingle1) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *ManagedElementSingle1) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *ManagedElementSingle1) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *ManagedElementSingle1) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *ManagedElementSingle1) GetAttributes() ManagedElementAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ManagedElementSingle1) GetAttributesOk() (*ManagedElementAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ManagedElementSingle1) SetAttributes(v ManagedElementAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ManagedElementSingle1) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetMnsAgent

`func (o *ManagedElementSingle1) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *ManagedElementSingle1) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *ManagedElementSingle1) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *ManagedElementSingle1) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *ManagedElementSingle1) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *ManagedElementSingle1) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *ManagedElementSingle1) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *ManagedElementSingle1) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *ManagedElementSingle1) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *ManagedElementSingle1) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *ManagedElementSingle1) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *ManagedElementSingle1) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetTraceJob

`func (o *ManagedElementSingle1) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *ManagedElementSingle1) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *ManagedElementSingle1) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *ManagedElementSingle1) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetNtfSubscriptionControl

`func (o *ManagedElementSingle1) GetNtfSubscriptionControl() []NtfSubscriptionControlSingle`

GetNtfSubscriptionControl returns the NtfSubscriptionControl field if non-nil, zero value otherwise.

### GetNtfSubscriptionControlOk

`func (o *ManagedElementSingle1) GetNtfSubscriptionControlOk() (*[]NtfSubscriptionControlSingle, bool)`

GetNtfSubscriptionControlOk returns a tuple with the NtfSubscriptionControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtfSubscriptionControl

`func (o *ManagedElementSingle1) SetNtfSubscriptionControl(v []NtfSubscriptionControlSingle)`

SetNtfSubscriptionControl sets NtfSubscriptionControl field to given value.

### HasNtfSubscriptionControl

`func (o *ManagedElementSingle1) HasNtfSubscriptionControl() bool`

HasNtfSubscriptionControl returns a boolean if a field has been set.

### GetAlarmList

`func (o *ManagedElementSingle1) GetAlarmList() AlarmListSingle`

GetAlarmList returns the AlarmList field if non-nil, zero value otherwise.

### GetAlarmListOk

`func (o *ManagedElementSingle1) GetAlarmListOk() (*AlarmListSingle, bool)`

GetAlarmListOk returns a tuple with the AlarmList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmList

`func (o *ManagedElementSingle1) SetAlarmList(v AlarmListSingle)`

SetAlarmList sets AlarmList field to given value.

### HasAlarmList

`func (o *ManagedElementSingle1) HasAlarmList() bool`

HasAlarmList returns a boolean if a field has been set.

### GetFileDownloadJob

`func (o *ManagedElementSingle1) GetFileDownloadJob() []FileDownloadJobSingle`

GetFileDownloadJob returns the FileDownloadJob field if non-nil, zero value otherwise.

### GetFileDownloadJobOk

`func (o *ManagedElementSingle1) GetFileDownloadJobOk() (*[]FileDownloadJobSingle, bool)`

GetFileDownloadJobOk returns a tuple with the FileDownloadJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadJob

`func (o *ManagedElementSingle1) SetFileDownloadJob(v []FileDownloadJobSingle)`

SetFileDownloadJob sets FileDownloadJob field to given value.

### HasFileDownloadJob

`func (o *ManagedElementSingle1) HasFileDownloadJob() bool`

HasFileDownloadJob returns a boolean if a field has been set.

### GetFiles

`func (o *ManagedElementSingle1) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ManagedElementSingle1) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ManagedElementSingle1) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ManagedElementSingle1) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetAmfFunction

`func (o *ManagedElementSingle1) GetAmfFunction() []AmfFunctionSingle`

GetAmfFunction returns the AmfFunction field if non-nil, zero value otherwise.

### GetAmfFunctionOk

`func (o *ManagedElementSingle1) GetAmfFunctionOk() (*[]AmfFunctionSingle, bool)`

GetAmfFunctionOk returns a tuple with the AmfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfFunction

`func (o *ManagedElementSingle1) SetAmfFunction(v []AmfFunctionSingle)`

SetAmfFunction sets AmfFunction field to given value.

### HasAmfFunction

`func (o *ManagedElementSingle1) HasAmfFunction() bool`

HasAmfFunction returns a boolean if a field has been set.

### GetSmfFunction

`func (o *ManagedElementSingle1) GetSmfFunction() []SmfFunctionSingle`

GetSmfFunction returns the SmfFunction field if non-nil, zero value otherwise.

### GetSmfFunctionOk

`func (o *ManagedElementSingle1) GetSmfFunctionOk() (*[]SmfFunctionSingle, bool)`

GetSmfFunctionOk returns a tuple with the SmfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfFunction

`func (o *ManagedElementSingle1) SetSmfFunction(v []SmfFunctionSingle)`

SetSmfFunction sets SmfFunction field to given value.

### HasSmfFunction

`func (o *ManagedElementSingle1) HasSmfFunction() bool`

HasSmfFunction returns a boolean if a field has been set.

### GetUpfFunction

`func (o *ManagedElementSingle1) GetUpfFunction() []UpfFunctionSingle`

GetUpfFunction returns the UpfFunction field if non-nil, zero value otherwise.

### GetUpfFunctionOk

`func (o *ManagedElementSingle1) GetUpfFunctionOk() (*[]UpfFunctionSingle, bool)`

GetUpfFunctionOk returns a tuple with the UpfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfFunction

`func (o *ManagedElementSingle1) SetUpfFunction(v []UpfFunctionSingle)`

SetUpfFunction sets UpfFunction field to given value.

### HasUpfFunction

`func (o *ManagedElementSingle1) HasUpfFunction() bool`

HasUpfFunction returns a boolean if a field has been set.

### GetN3iwfFunction

`func (o *ManagedElementSingle1) GetN3iwfFunction() []N3iwfFunctionSingle`

GetN3iwfFunction returns the N3iwfFunction field if non-nil, zero value otherwise.

### GetN3iwfFunctionOk

`func (o *ManagedElementSingle1) GetN3iwfFunctionOk() (*[]N3iwfFunctionSingle, bool)`

GetN3iwfFunctionOk returns a tuple with the N3iwfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3iwfFunction

`func (o *ManagedElementSingle1) SetN3iwfFunction(v []N3iwfFunctionSingle)`

SetN3iwfFunction sets N3iwfFunction field to given value.

### HasN3iwfFunction

`func (o *ManagedElementSingle1) HasN3iwfFunction() bool`

HasN3iwfFunction returns a boolean if a field has been set.

### GetPcfFunction

`func (o *ManagedElementSingle1) GetPcfFunction() []PcfFunctionSingle`

GetPcfFunction returns the PcfFunction field if non-nil, zero value otherwise.

### GetPcfFunctionOk

`func (o *ManagedElementSingle1) GetPcfFunctionOk() (*[]PcfFunctionSingle, bool)`

GetPcfFunctionOk returns a tuple with the PcfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfFunction

`func (o *ManagedElementSingle1) SetPcfFunction(v []PcfFunctionSingle)`

SetPcfFunction sets PcfFunction field to given value.

### HasPcfFunction

`func (o *ManagedElementSingle1) HasPcfFunction() bool`

HasPcfFunction returns a boolean if a field has been set.

### GetAusfFunction

`func (o *ManagedElementSingle1) GetAusfFunction() []AusfFunctionSingle`

GetAusfFunction returns the AusfFunction field if non-nil, zero value otherwise.

### GetAusfFunctionOk

`func (o *ManagedElementSingle1) GetAusfFunctionOk() (*[]AusfFunctionSingle, bool)`

GetAusfFunctionOk returns a tuple with the AusfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAusfFunction

`func (o *ManagedElementSingle1) SetAusfFunction(v []AusfFunctionSingle)`

SetAusfFunction sets AusfFunction field to given value.

### HasAusfFunction

`func (o *ManagedElementSingle1) HasAusfFunction() bool`

HasAusfFunction returns a boolean if a field has been set.

### GetUdmFunction

`func (o *ManagedElementSingle1) GetUdmFunction() []UdmFunctionSingle`

GetUdmFunction returns the UdmFunction field if non-nil, zero value otherwise.

### GetUdmFunctionOk

`func (o *ManagedElementSingle1) GetUdmFunctionOk() (*[]UdmFunctionSingle, bool)`

GetUdmFunctionOk returns a tuple with the UdmFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmFunction

`func (o *ManagedElementSingle1) SetUdmFunction(v []UdmFunctionSingle)`

SetUdmFunction sets UdmFunction field to given value.

### HasUdmFunction

`func (o *ManagedElementSingle1) HasUdmFunction() bool`

HasUdmFunction returns a boolean if a field has been set.

### GetUdrFunction

`func (o *ManagedElementSingle1) GetUdrFunction() []UdrFunctionSingle`

GetUdrFunction returns the UdrFunction field if non-nil, zero value otherwise.

### GetUdrFunctionOk

`func (o *ManagedElementSingle1) GetUdrFunctionOk() (*[]UdrFunctionSingle, bool)`

GetUdrFunctionOk returns a tuple with the UdrFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrFunction

`func (o *ManagedElementSingle1) SetUdrFunction(v []UdrFunctionSingle)`

SetUdrFunction sets UdrFunction field to given value.

### HasUdrFunction

`func (o *ManagedElementSingle1) HasUdrFunction() bool`

HasUdrFunction returns a boolean if a field has been set.

### GetUdsfFunction

`func (o *ManagedElementSingle1) GetUdsfFunction() []UdsfFunctionSingle`

GetUdsfFunction returns the UdsfFunction field if non-nil, zero value otherwise.

### GetUdsfFunctionOk

`func (o *ManagedElementSingle1) GetUdsfFunctionOk() (*[]UdsfFunctionSingle, bool)`

GetUdsfFunctionOk returns a tuple with the UdsfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdsfFunction

`func (o *ManagedElementSingle1) SetUdsfFunction(v []UdsfFunctionSingle)`

SetUdsfFunction sets UdsfFunction field to given value.

### HasUdsfFunction

`func (o *ManagedElementSingle1) HasUdsfFunction() bool`

HasUdsfFunction returns a boolean if a field has been set.

### GetNrfFunction

`func (o *ManagedElementSingle1) GetNrfFunction() []NrfFunctionSingle`

GetNrfFunction returns the NrfFunction field if non-nil, zero value otherwise.

### GetNrfFunctionOk

`func (o *ManagedElementSingle1) GetNrfFunctionOk() (*[]NrfFunctionSingle, bool)`

GetNrfFunctionOk returns a tuple with the NrfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfFunction

`func (o *ManagedElementSingle1) SetNrfFunction(v []NrfFunctionSingle)`

SetNrfFunction sets NrfFunction field to given value.

### HasNrfFunction

`func (o *ManagedElementSingle1) HasNrfFunction() bool`

HasNrfFunction returns a boolean if a field has been set.

### GetNssfFunction

`func (o *ManagedElementSingle1) GetNssfFunction() []NssfFunctionSingle`

GetNssfFunction returns the NssfFunction field if non-nil, zero value otherwise.

### GetNssfFunctionOk

`func (o *ManagedElementSingle1) GetNssfFunctionOk() (*[]NssfFunctionSingle, bool)`

GetNssfFunctionOk returns a tuple with the NssfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssfFunction

`func (o *ManagedElementSingle1) SetNssfFunction(v []NssfFunctionSingle)`

SetNssfFunction sets NssfFunction field to given value.

### HasNssfFunction

`func (o *ManagedElementSingle1) HasNssfFunction() bool`

HasNssfFunction returns a boolean if a field has been set.

### GetSmsfFunction

`func (o *ManagedElementSingle1) GetSmsfFunction() []SmsfFunctionSingle`

GetSmsfFunction returns the SmsfFunction field if non-nil, zero value otherwise.

### GetSmsfFunctionOk

`func (o *ManagedElementSingle1) GetSmsfFunctionOk() (*[]SmsfFunctionSingle, bool)`

GetSmsfFunctionOk returns a tuple with the SmsfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfFunction

`func (o *ManagedElementSingle1) SetSmsfFunction(v []SmsfFunctionSingle)`

SetSmsfFunction sets SmsfFunction field to given value.

### HasSmsfFunction

`func (o *ManagedElementSingle1) HasSmsfFunction() bool`

HasSmsfFunction returns a boolean if a field has been set.

### GetLmfFunction

`func (o *ManagedElementSingle1) GetLmfFunction() []LmfFunctionSingle`

GetLmfFunction returns the LmfFunction field if non-nil, zero value otherwise.

### GetLmfFunctionOk

`func (o *ManagedElementSingle1) GetLmfFunctionOk() (*[]LmfFunctionSingle, bool)`

GetLmfFunctionOk returns a tuple with the LmfFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmfFunction

`func (o *ManagedElementSingle1) SetLmfFunction(v []LmfFunctionSingle)`

SetLmfFunction sets LmfFunction field to given value.

### HasLmfFunction

`func (o *ManagedElementSingle1) HasLmfFunction() bool`

HasLmfFunction returns a boolean if a field has been set.

### GetNgeirFunction

`func (o *ManagedElementSingle1) GetNgeirFunction() []NgeirFunctionSingle`

GetNgeirFunction returns the NgeirFunction field if non-nil, zero value otherwise.

### GetNgeirFunctionOk

`func (o *ManagedElementSingle1) GetNgeirFunctionOk() (*[]NgeirFunctionSingle, bool)`

GetNgeirFunctionOk returns a tuple with the NgeirFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgeirFunction

`func (o *ManagedElementSingle1) SetNgeirFunction(v []NgeirFunctionSingle)`

SetNgeirFunction sets NgeirFunction field to given value.

### HasNgeirFunction

`func (o *ManagedElementSingle1) HasNgeirFunction() bool`

HasNgeirFunction returns a boolean if a field has been set.

### GetSeppFunction

`func (o *ManagedElementSingle1) GetSeppFunction() []SeppFunctionSingle`

GetSeppFunction returns the SeppFunction field if non-nil, zero value otherwise.

### GetSeppFunctionOk

`func (o *ManagedElementSingle1) GetSeppFunctionOk() (*[]SeppFunctionSingle, bool)`

GetSeppFunctionOk returns a tuple with the SeppFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeppFunction

`func (o *ManagedElementSingle1) SetSeppFunction(v []SeppFunctionSingle)`

SetSeppFunction sets SeppFunction field to given value.

### HasSeppFunction

`func (o *ManagedElementSingle1) HasSeppFunction() bool`

HasSeppFunction returns a boolean if a field has been set.

### GetNwdafFunction

`func (o *ManagedElementSingle1) GetNwdafFunction() []NwdafFunctionSingle`

GetNwdafFunction returns the NwdafFunction field if non-nil, zero value otherwise.

### GetNwdafFunctionOk

`func (o *ManagedElementSingle1) GetNwdafFunctionOk() (*[]NwdafFunctionSingle, bool)`

GetNwdafFunctionOk returns a tuple with the NwdafFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafFunction

`func (o *ManagedElementSingle1) SetNwdafFunction(v []NwdafFunctionSingle)`

SetNwdafFunction sets NwdafFunction field to given value.

### HasNwdafFunction

`func (o *ManagedElementSingle1) HasNwdafFunction() bool`

HasNwdafFunction returns a boolean if a field has been set.

### GetScpFunction

`func (o *ManagedElementSingle1) GetScpFunction() []ScpFunctionSingle`

GetScpFunction returns the ScpFunction field if non-nil, zero value otherwise.

### GetScpFunctionOk

`func (o *ManagedElementSingle1) GetScpFunctionOk() (*[]ScpFunctionSingle, bool)`

GetScpFunctionOk returns a tuple with the ScpFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpFunction

`func (o *ManagedElementSingle1) SetScpFunction(v []ScpFunctionSingle)`

SetScpFunction sets ScpFunction field to given value.

### HasScpFunction

`func (o *ManagedElementSingle1) HasScpFunction() bool`

HasScpFunction returns a boolean if a field has been set.

### GetNefFunction

`func (o *ManagedElementSingle1) GetNefFunction() []NefFunctionSingle`

GetNefFunction returns the NefFunction field if non-nil, zero value otherwise.

### GetNefFunctionOk

`func (o *ManagedElementSingle1) GetNefFunctionOk() (*[]NefFunctionSingle, bool)`

GetNefFunctionOk returns a tuple with the NefFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefFunction

`func (o *ManagedElementSingle1) SetNefFunction(v []NefFunctionSingle)`

SetNefFunction sets NefFunction field to given value.

### HasNefFunction

`func (o *ManagedElementSingle1) HasNefFunction() bool`

HasNefFunction returns a boolean if a field has been set.

### GetConfigurable5QISet

`func (o *ManagedElementSingle1) GetConfigurable5QISet() []Configurable5QISetSingle`

GetConfigurable5QISet returns the Configurable5QISet field if non-nil, zero value otherwise.

### GetConfigurable5QISetOk

`func (o *ManagedElementSingle1) GetConfigurable5QISetOk() (*[]Configurable5QISetSingle, bool)`

GetConfigurable5QISetOk returns a tuple with the Configurable5QISet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable5QISet

`func (o *ManagedElementSingle1) SetConfigurable5QISet(v []Configurable5QISetSingle)`

SetConfigurable5QISet sets Configurable5QISet field to given value.

### HasConfigurable5QISet

`func (o *ManagedElementSingle1) HasConfigurable5QISet() bool`

HasConfigurable5QISet returns a boolean if a field has been set.

### GetDynamic5QISet

`func (o *ManagedElementSingle1) GetDynamic5QISet() []Dynamic5QISetSingle`

GetDynamic5QISet returns the Dynamic5QISet field if non-nil, zero value otherwise.

### GetDynamic5QISetOk

`func (o *ManagedElementSingle1) GetDynamic5QISetOk() (*[]Dynamic5QISetSingle, bool)`

GetDynamic5QISetOk returns a tuple with the Dynamic5QISet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic5QISet

`func (o *ManagedElementSingle1) SetDynamic5QISet(v []Dynamic5QISetSingle)`

SetDynamic5QISet sets Dynamic5QISet field to given value.

### HasDynamic5QISet

`func (o *ManagedElementSingle1) HasDynamic5QISet() bool`

HasDynamic5QISet returns a boolean if a field has been set.

### GetEcmConnectionInfo

`func (o *ManagedElementSingle1) GetEcmConnectionInfo() []EcmConnectionInfoSingle`

GetEcmConnectionInfo returns the EcmConnectionInfo field if non-nil, zero value otherwise.

### GetEcmConnectionInfoOk

`func (o *ManagedElementSingle1) GetEcmConnectionInfoOk() (*[]EcmConnectionInfoSingle, bool)`

GetEcmConnectionInfoOk returns a tuple with the EcmConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcmConnectionInfo

`func (o *ManagedElementSingle1) SetEcmConnectionInfo(v []EcmConnectionInfoSingle)`

SetEcmConnectionInfo sets EcmConnectionInfo field to given value.

### HasEcmConnectionInfo

`func (o *ManagedElementSingle1) HasEcmConnectionInfo() bool`

HasEcmConnectionInfo returns a boolean if a field has been set.

### GetEASDFFunction

`func (o *ManagedElementSingle1) GetEASDFFunction() []EASDFFunctionSingle`

GetEASDFFunction returns the EASDFFunction field if non-nil, zero value otherwise.

### GetEASDFFunctionOk

`func (o *ManagedElementSingle1) GetEASDFFunctionOk() (*[]EASDFFunctionSingle, bool)`

GetEASDFFunctionOk returns a tuple with the EASDFFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEASDFFunction

`func (o *ManagedElementSingle1) SetEASDFFunction(v []EASDFFunctionSingle)`

SetEASDFFunction sets EASDFFunction field to given value.

### HasEASDFFunction

`func (o *ManagedElementSingle1) HasEASDFFunction() bool`

HasEASDFFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


