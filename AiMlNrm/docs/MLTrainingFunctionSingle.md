# MLTrainingFunctionSingle

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
**MLTrainingRequest** | Pointer to [**[]MLTrainingRequestSingle**](MLTrainingRequestSingle.md) |  | [optional] 
**MLTrainingProcess** | Pointer to [**[]MLTrainingProcessSingle**](MLTrainingProcessSingle.md) |  | [optional] 
**MLTrainingReport** | Pointer to [**[]MLTrainingReportSingle**](MLTrainingReportSingle.md) |  | [optional] 

## Methods

### NewMLTrainingFunctionSingle

`func NewMLTrainingFunctionSingle(id NullableString, ) *MLTrainingFunctionSingle`

NewMLTrainingFunctionSingle instantiates a new MLTrainingFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMLTrainingFunctionSingleWithDefaults

`func NewMLTrainingFunctionSingleWithDefaults() *MLTrainingFunctionSingle`

NewMLTrainingFunctionSingleWithDefaults instantiates a new MLTrainingFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MLTrainingFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MLTrainingFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MLTrainingFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *MLTrainingFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MLTrainingFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *MLTrainingFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *MLTrainingFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *MLTrainingFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *MLTrainingFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *MLTrainingFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *MLTrainingFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *MLTrainingFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *MLTrainingFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *MLTrainingFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *MLTrainingFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *MLTrainingFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *MLTrainingFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *MLTrainingFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MLTrainingFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MLTrainingFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MLTrainingFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *MLTrainingFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *MLTrainingFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *MLTrainingFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *MLTrainingFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *MLTrainingFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *MLTrainingFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *MLTrainingFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *MLTrainingFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *MLTrainingFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *MLTrainingFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *MLTrainingFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *MLTrainingFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *MLTrainingFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *MLTrainingFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *MLTrainingFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *MLTrainingFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetMLTrainingRequest

`func (o *MLTrainingFunctionSingle) GetMLTrainingRequest() []MLTrainingRequestSingle`

GetMLTrainingRequest returns the MLTrainingRequest field if non-nil, zero value otherwise.

### GetMLTrainingRequestOk

`func (o *MLTrainingFunctionSingle) GetMLTrainingRequestOk() (*[]MLTrainingRequestSingle, bool)`

GetMLTrainingRequestOk returns a tuple with the MLTrainingRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLTrainingRequest

`func (o *MLTrainingFunctionSingle) SetMLTrainingRequest(v []MLTrainingRequestSingle)`

SetMLTrainingRequest sets MLTrainingRequest field to given value.

### HasMLTrainingRequest

`func (o *MLTrainingFunctionSingle) HasMLTrainingRequest() bool`

HasMLTrainingRequest returns a boolean if a field has been set.

### GetMLTrainingProcess

`func (o *MLTrainingFunctionSingle) GetMLTrainingProcess() []MLTrainingProcessSingle`

GetMLTrainingProcess returns the MLTrainingProcess field if non-nil, zero value otherwise.

### GetMLTrainingProcessOk

`func (o *MLTrainingFunctionSingle) GetMLTrainingProcessOk() (*[]MLTrainingProcessSingle, bool)`

GetMLTrainingProcessOk returns a tuple with the MLTrainingProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLTrainingProcess

`func (o *MLTrainingFunctionSingle) SetMLTrainingProcess(v []MLTrainingProcessSingle)`

SetMLTrainingProcess sets MLTrainingProcess field to given value.

### HasMLTrainingProcess

`func (o *MLTrainingFunctionSingle) HasMLTrainingProcess() bool`

HasMLTrainingProcess returns a boolean if a field has been set.

### GetMLTrainingReport

`func (o *MLTrainingFunctionSingle) GetMLTrainingReport() []MLTrainingReportSingle`

GetMLTrainingReport returns the MLTrainingReport field if non-nil, zero value otherwise.

### GetMLTrainingReportOk

`func (o *MLTrainingFunctionSingle) GetMLTrainingReportOk() (*[]MLTrainingReportSingle, bool)`

GetMLTrainingReportOk returns a tuple with the MLTrainingReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLTrainingReport

`func (o *MLTrainingFunctionSingle) SetMLTrainingReport(v []MLTrainingReportSingle)`

SetMLTrainingReport sets MLTrainingReport field to given value.

### HasMLTrainingReport

`func (o *MLTrainingFunctionSingle) HasMLTrainingReport() bool`

HasMLTrainingReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


