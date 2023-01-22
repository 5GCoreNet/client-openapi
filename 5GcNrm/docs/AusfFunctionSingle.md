# AusfFunctionSingle

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
**EPN12** | Pointer to [**[]EPN12Single**](EPN12Single.md) |  | [optional] 
**EPN13** | Pointer to [**[]EPN13Single**](EPN13Single.md) |  | [optional] 

## Methods

### NewAusfFunctionSingle

`func NewAusfFunctionSingle(id NullableString, ) *AusfFunctionSingle`

NewAusfFunctionSingle instantiates a new AusfFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAusfFunctionSingleWithDefaults

`func NewAusfFunctionSingleWithDefaults() *AusfFunctionSingle`

NewAusfFunctionSingleWithDefaults instantiates a new AusfFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AusfFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AusfFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AusfFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AusfFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AusfFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *AusfFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *AusfFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *AusfFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *AusfFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *AusfFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *AusfFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *AusfFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *AusfFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *AusfFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *AusfFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *AusfFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *AusfFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *AusfFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AusfFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AusfFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AusfFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *AusfFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *AusfFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *AusfFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *AusfFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *AusfFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *AusfFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *AusfFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *AusfFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *AusfFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *AusfFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *AusfFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *AusfFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *AusfFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *AusfFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *AusfFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *AusfFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetEPN12

`func (o *AusfFunctionSingle) GetEPN12() []EPN12Single`

GetEPN12 returns the EPN12 field if non-nil, zero value otherwise.

### GetEPN12Ok

`func (o *AusfFunctionSingle) GetEPN12Ok() (*[]EPN12Single, bool)`

GetEPN12Ok returns a tuple with the EPN12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN12

`func (o *AusfFunctionSingle) SetEPN12(v []EPN12Single)`

SetEPN12 sets EPN12 field to given value.

### HasEPN12

`func (o *AusfFunctionSingle) HasEPN12() bool`

HasEPN12 returns a boolean if a field has been set.

### GetEPN13

`func (o *AusfFunctionSingle) GetEPN13() []EPN13Single`

GetEPN13 returns the EPN13 field if non-nil, zero value otherwise.

### GetEPN13Ok

`func (o *AusfFunctionSingle) GetEPN13Ok() (*[]EPN13Single, bool)`

GetEPN13Ok returns a tuple with the EPN13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN13

`func (o *AusfFunctionSingle) SetEPN13(v []EPN13Single)`

SetEPN13 sets EPN13 field to given value.

### HasEPN13

`func (o *AusfFunctionSingle) HasEPN13() bool`

HasEPN13 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


