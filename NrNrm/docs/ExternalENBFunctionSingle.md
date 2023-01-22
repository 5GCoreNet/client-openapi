# ExternalENBFunctionSingle

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
**ExternalEUTranCell** | Pointer to [**[]ExternalEUTranCellSingle**](ExternalEUTranCellSingle.md) |  | [optional] 

## Methods

### NewExternalENBFunctionSingle

`func NewExternalENBFunctionSingle(id NullableString, ) *ExternalENBFunctionSingle`

NewExternalENBFunctionSingle instantiates a new ExternalENBFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalENBFunctionSingleWithDefaults

`func NewExternalENBFunctionSingleWithDefaults() *ExternalENBFunctionSingle`

NewExternalENBFunctionSingleWithDefaults instantiates a new ExternalENBFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalENBFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalENBFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalENBFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ExternalENBFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExternalENBFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *ExternalENBFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ExternalENBFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ExternalENBFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ExternalENBFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *ExternalENBFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *ExternalENBFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *ExternalENBFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *ExternalENBFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *ExternalENBFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *ExternalENBFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *ExternalENBFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *ExternalENBFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *ExternalENBFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalENBFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalENBFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ExternalENBFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *ExternalENBFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *ExternalENBFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *ExternalENBFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *ExternalENBFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *ExternalENBFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *ExternalENBFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *ExternalENBFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *ExternalENBFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *ExternalENBFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *ExternalENBFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *ExternalENBFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *ExternalENBFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *ExternalENBFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *ExternalENBFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *ExternalENBFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *ExternalENBFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetExternalEUTranCell

`func (o *ExternalENBFunctionSingle) GetExternalEUTranCell() []ExternalEUTranCellSingle`

GetExternalEUTranCell returns the ExternalEUTranCell field if non-nil, zero value otherwise.

### GetExternalEUTranCellOk

`func (o *ExternalENBFunctionSingle) GetExternalEUTranCellOk() (*[]ExternalEUTranCellSingle, bool)`

GetExternalEUTranCellOk returns a tuple with the ExternalEUTranCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEUTranCell

`func (o *ExternalENBFunctionSingle) SetExternalEUTranCell(v []ExternalEUTranCellSingle)`

SetExternalEUTranCell sets ExternalEUTranCell field to given value.

### HasExternalEUTranCell

`func (o *ExternalENBFunctionSingle) HasExternalEUTranCell() bool`

HasExternalEUTranCell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


