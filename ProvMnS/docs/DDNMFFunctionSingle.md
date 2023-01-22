# DDNMFFunctionSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**ManagedFunctionAttr**](ManagedFunction-Attr.md) |  | [optional] 
**PerfMetricJob** | Pointer to [**[]PerfMetricJobSingle**](PerfMetricJobSingle.md) |  | [optional] 
**ThresholdMonitor** | Pointer to [**[]ThresholdMonitorSingle**](ThresholdMonitorSingle.md) |  | [optional] 
**ManagedNFService** | Pointer to [**[]ManagedNFServiceSingle**](ManagedNFServiceSingle.md) |  | [optional] 
**TraceJob** | Pointer to [**[]TraceJobSingle**](TraceJobSingle.md) |  | [optional] 
**EPNpc4** | Pointer to [**[]EPNpc4Single**](EPNpc4Single.md) |  | [optional] 
**EPNpc6** | Pointer to [**[]EPNpc6Single**](EPNpc6Single.md) |  | [optional] 
**EPNpc7** | Pointer to [**[]EPNpc7Single**](EPNpc7Single.md) |  | [optional] 
**EPNpc8** | Pointer to [**[]EPNpc8Single**](EPNpc8Single.md) |  | [optional] 

## Methods

### NewDDNMFFunctionSingle

`func NewDDNMFFunctionSingle(id NullableString, ) *DDNMFFunctionSingle`

NewDDNMFFunctionSingle instantiates a new DDNMFFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDDNMFFunctionSingleWithDefaults

`func NewDDNMFFunctionSingleWithDefaults() *DDNMFFunctionSingle`

NewDDNMFFunctionSingleWithDefaults instantiates a new DDNMFFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DDNMFFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DDNMFFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DDNMFFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *DDNMFFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DDNMFFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *DDNMFFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *DDNMFFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *DDNMFFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *DDNMFFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *DDNMFFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *DDNMFFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *DDNMFFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *DDNMFFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *DDNMFFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *DDNMFFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *DDNMFFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *DDNMFFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *DDNMFFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DDNMFFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DDNMFFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DDNMFFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *DDNMFFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *DDNMFFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *DDNMFFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *DDNMFFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *DDNMFFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *DDNMFFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *DDNMFFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *DDNMFFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *DDNMFFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *DDNMFFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *DDNMFFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *DDNMFFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *DDNMFFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *DDNMFFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *DDNMFFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *DDNMFFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetEPNpc4

`func (o *DDNMFFunctionSingle) GetEPNpc4() []EPNpc4Single`

GetEPNpc4 returns the EPNpc4 field if non-nil, zero value otherwise.

### GetEPNpc4Ok

`func (o *DDNMFFunctionSingle) GetEPNpc4Ok() (*[]EPNpc4Single, bool)`

GetEPNpc4Ok returns a tuple with the EPNpc4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPNpc4

`func (o *DDNMFFunctionSingle) SetEPNpc4(v []EPNpc4Single)`

SetEPNpc4 sets EPNpc4 field to given value.

### HasEPNpc4

`func (o *DDNMFFunctionSingle) HasEPNpc4() bool`

HasEPNpc4 returns a boolean if a field has been set.

### GetEPNpc6

`func (o *DDNMFFunctionSingle) GetEPNpc6() []EPNpc6Single`

GetEPNpc6 returns the EPNpc6 field if non-nil, zero value otherwise.

### GetEPNpc6Ok

`func (o *DDNMFFunctionSingle) GetEPNpc6Ok() (*[]EPNpc6Single, bool)`

GetEPNpc6Ok returns a tuple with the EPNpc6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPNpc6

`func (o *DDNMFFunctionSingle) SetEPNpc6(v []EPNpc6Single)`

SetEPNpc6 sets EPNpc6 field to given value.

### HasEPNpc6

`func (o *DDNMFFunctionSingle) HasEPNpc6() bool`

HasEPNpc6 returns a boolean if a field has been set.

### GetEPNpc7

`func (o *DDNMFFunctionSingle) GetEPNpc7() []EPNpc7Single`

GetEPNpc7 returns the EPNpc7 field if non-nil, zero value otherwise.

### GetEPNpc7Ok

`func (o *DDNMFFunctionSingle) GetEPNpc7Ok() (*[]EPNpc7Single, bool)`

GetEPNpc7Ok returns a tuple with the EPNpc7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPNpc7

`func (o *DDNMFFunctionSingle) SetEPNpc7(v []EPNpc7Single)`

SetEPNpc7 sets EPNpc7 field to given value.

### HasEPNpc7

`func (o *DDNMFFunctionSingle) HasEPNpc7() bool`

HasEPNpc7 returns a boolean if a field has been set.

### GetEPNpc8

`func (o *DDNMFFunctionSingle) GetEPNpc8() []EPNpc8Single`

GetEPNpc8 returns the EPNpc8 field if non-nil, zero value otherwise.

### GetEPNpc8Ok

`func (o *DDNMFFunctionSingle) GetEPNpc8Ok() (*[]EPNpc8Single, bool)`

GetEPNpc8Ok returns a tuple with the EPNpc8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPNpc8

`func (o *DDNMFFunctionSingle) SetEPNpc8(v []EPNpc8Single)`

SetEPNpc8 sets EPNpc8 field to given value.

### HasEPNpc8

`func (o *DDNMFFunctionSingle) HasEPNpc8() bool`

HasEPNpc8 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


