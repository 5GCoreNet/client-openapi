# ExternalGnbCuCpFunctionSingle

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
**ExternalNrCellCu** | Pointer to [**[]ExternalNrCellCuSingle**](ExternalNrCellCuSingle.md) |  | [optional] 
**EPXnC** | Pointer to [**[]EPXnCSingle**](EPXnCSingle.md) |  | [optional] 
**EPE1** | Pointer to [**[]EPE1Single**](EPE1Single.md) |  | [optional] 
**EPF1C** | Pointer to [**[]EPF1CSingle**](EPF1CSingle.md) |  | [optional] 

## Methods

### NewExternalGnbCuCpFunctionSingle

`func NewExternalGnbCuCpFunctionSingle(id NullableString, ) *ExternalGnbCuCpFunctionSingle`

NewExternalGnbCuCpFunctionSingle instantiates a new ExternalGnbCuCpFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGnbCuCpFunctionSingleWithDefaults

`func NewExternalGnbCuCpFunctionSingleWithDefaults() *ExternalGnbCuCpFunctionSingle`

NewExternalGnbCuCpFunctionSingleWithDefaults instantiates a new ExternalGnbCuCpFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalGnbCuCpFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalGnbCuCpFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalGnbCuCpFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ExternalGnbCuCpFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExternalGnbCuCpFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *ExternalGnbCuCpFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ExternalGnbCuCpFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ExternalGnbCuCpFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ExternalGnbCuCpFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *ExternalGnbCuCpFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *ExternalGnbCuCpFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *ExternalGnbCuCpFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *ExternalGnbCuCpFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *ExternalGnbCuCpFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *ExternalGnbCuCpFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *ExternalGnbCuCpFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *ExternalGnbCuCpFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *ExternalGnbCuCpFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalGnbCuCpFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalGnbCuCpFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ExternalGnbCuCpFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *ExternalGnbCuCpFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *ExternalGnbCuCpFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *ExternalGnbCuCpFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *ExternalGnbCuCpFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *ExternalGnbCuCpFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *ExternalGnbCuCpFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *ExternalGnbCuCpFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *ExternalGnbCuCpFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *ExternalGnbCuCpFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *ExternalGnbCuCpFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *ExternalGnbCuCpFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *ExternalGnbCuCpFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *ExternalGnbCuCpFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *ExternalGnbCuCpFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *ExternalGnbCuCpFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *ExternalGnbCuCpFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetExternalNrCellCu

`func (o *ExternalGnbCuCpFunctionSingle) GetExternalNrCellCu() []ExternalNrCellCuSingle`

GetExternalNrCellCu returns the ExternalNrCellCu field if non-nil, zero value otherwise.

### GetExternalNrCellCuOk

`func (o *ExternalGnbCuCpFunctionSingle) GetExternalNrCellCuOk() (*[]ExternalNrCellCuSingle, bool)`

GetExternalNrCellCuOk returns a tuple with the ExternalNrCellCu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNrCellCu

`func (o *ExternalGnbCuCpFunctionSingle) SetExternalNrCellCu(v []ExternalNrCellCuSingle)`

SetExternalNrCellCu sets ExternalNrCellCu field to given value.

### HasExternalNrCellCu

`func (o *ExternalGnbCuCpFunctionSingle) HasExternalNrCellCu() bool`

HasExternalNrCellCu returns a boolean if a field has been set.

### GetEPXnC

`func (o *ExternalGnbCuCpFunctionSingle) GetEPXnC() []EPXnCSingle`

GetEPXnC returns the EPXnC field if non-nil, zero value otherwise.

### GetEPXnCOk

`func (o *ExternalGnbCuCpFunctionSingle) GetEPXnCOk() (*[]EPXnCSingle, bool)`

GetEPXnCOk returns a tuple with the EPXnC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPXnC

`func (o *ExternalGnbCuCpFunctionSingle) SetEPXnC(v []EPXnCSingle)`

SetEPXnC sets EPXnC field to given value.

### HasEPXnC

`func (o *ExternalGnbCuCpFunctionSingle) HasEPXnC() bool`

HasEPXnC returns a boolean if a field has been set.

### GetEPE1

`func (o *ExternalGnbCuCpFunctionSingle) GetEPE1() []EPE1Single`

GetEPE1 returns the EPE1 field if non-nil, zero value otherwise.

### GetEPE1Ok

`func (o *ExternalGnbCuCpFunctionSingle) GetEPE1Ok() (*[]EPE1Single, bool)`

GetEPE1Ok returns a tuple with the EPE1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPE1

`func (o *ExternalGnbCuCpFunctionSingle) SetEPE1(v []EPE1Single)`

SetEPE1 sets EPE1 field to given value.

### HasEPE1

`func (o *ExternalGnbCuCpFunctionSingle) HasEPE1() bool`

HasEPE1 returns a boolean if a field has been set.

### GetEPF1C

`func (o *ExternalGnbCuCpFunctionSingle) GetEPF1C() []EPF1CSingle`

GetEPF1C returns the EPF1C field if non-nil, zero value otherwise.

### GetEPF1COk

`func (o *ExternalGnbCuCpFunctionSingle) GetEPF1COk() (*[]EPF1CSingle, bool)`

GetEPF1COk returns a tuple with the EPF1C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPF1C

`func (o *ExternalGnbCuCpFunctionSingle) SetEPF1C(v []EPF1CSingle)`

SetEPF1C sets EPF1C field to given value.

### HasEPF1C

`func (o *ExternalGnbCuCpFunctionSingle) HasEPF1C() bool`

HasEPF1C returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


