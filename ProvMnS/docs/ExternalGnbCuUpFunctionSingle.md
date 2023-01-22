# ExternalGnbCuUpFunctionSingle

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
**EPE1** | Pointer to [**[]EPE1Single**](EPE1Single.md) |  | [optional] 
**EPF1U** | Pointer to [**[]EPF1USingle**](EPF1USingle.md) |  | [optional] 
**EPXnU** | Pointer to [**[]EPXnUSingle**](EPXnUSingle.md) |  | [optional] 

## Methods

### NewExternalGnbCuUpFunctionSingle

`func NewExternalGnbCuUpFunctionSingle(id NullableString, ) *ExternalGnbCuUpFunctionSingle`

NewExternalGnbCuUpFunctionSingle instantiates a new ExternalGnbCuUpFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGnbCuUpFunctionSingleWithDefaults

`func NewExternalGnbCuUpFunctionSingleWithDefaults() *ExternalGnbCuUpFunctionSingle`

NewExternalGnbCuUpFunctionSingleWithDefaults instantiates a new ExternalGnbCuUpFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalGnbCuUpFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalGnbCuUpFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalGnbCuUpFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ExternalGnbCuUpFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExternalGnbCuUpFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *ExternalGnbCuUpFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ExternalGnbCuUpFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ExternalGnbCuUpFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ExternalGnbCuUpFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *ExternalGnbCuUpFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *ExternalGnbCuUpFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *ExternalGnbCuUpFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *ExternalGnbCuUpFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *ExternalGnbCuUpFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *ExternalGnbCuUpFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *ExternalGnbCuUpFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *ExternalGnbCuUpFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *ExternalGnbCuUpFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalGnbCuUpFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalGnbCuUpFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ExternalGnbCuUpFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *ExternalGnbCuUpFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *ExternalGnbCuUpFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *ExternalGnbCuUpFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *ExternalGnbCuUpFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *ExternalGnbCuUpFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *ExternalGnbCuUpFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *ExternalGnbCuUpFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *ExternalGnbCuUpFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *ExternalGnbCuUpFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *ExternalGnbCuUpFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *ExternalGnbCuUpFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *ExternalGnbCuUpFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *ExternalGnbCuUpFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *ExternalGnbCuUpFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *ExternalGnbCuUpFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *ExternalGnbCuUpFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetEPE1

`func (o *ExternalGnbCuUpFunctionSingle) GetEPE1() []EPE1Single`

GetEPE1 returns the EPE1 field if non-nil, zero value otherwise.

### GetEPE1Ok

`func (o *ExternalGnbCuUpFunctionSingle) GetEPE1Ok() (*[]EPE1Single, bool)`

GetEPE1Ok returns a tuple with the EPE1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPE1

`func (o *ExternalGnbCuUpFunctionSingle) SetEPE1(v []EPE1Single)`

SetEPE1 sets EPE1 field to given value.

### HasEPE1

`func (o *ExternalGnbCuUpFunctionSingle) HasEPE1() bool`

HasEPE1 returns a boolean if a field has been set.

### GetEPF1U

`func (o *ExternalGnbCuUpFunctionSingle) GetEPF1U() []EPF1USingle`

GetEPF1U returns the EPF1U field if non-nil, zero value otherwise.

### GetEPF1UOk

`func (o *ExternalGnbCuUpFunctionSingle) GetEPF1UOk() (*[]EPF1USingle, bool)`

GetEPF1UOk returns a tuple with the EPF1U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPF1U

`func (o *ExternalGnbCuUpFunctionSingle) SetEPF1U(v []EPF1USingle)`

SetEPF1U sets EPF1U field to given value.

### HasEPF1U

`func (o *ExternalGnbCuUpFunctionSingle) HasEPF1U() bool`

HasEPF1U returns a boolean if a field has been set.

### GetEPXnU

`func (o *ExternalGnbCuUpFunctionSingle) GetEPXnU() []EPXnUSingle`

GetEPXnU returns the EPXnU field if non-nil, zero value otherwise.

### GetEPXnUOk

`func (o *ExternalGnbCuUpFunctionSingle) GetEPXnUOk() (*[]EPXnUSingle, bool)`

GetEPXnUOk returns a tuple with the EPXnU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPXnU

`func (o *ExternalGnbCuUpFunctionSingle) SetEPXnU(v []EPXnUSingle)`

SetEPXnU sets EPXnU field to given value.

### HasEPXnU

`func (o *ExternalGnbCuUpFunctionSingle) HasEPXnU() bool`

HasEPXnU returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


