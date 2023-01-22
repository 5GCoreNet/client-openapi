# UpfFunctionSingle

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
**EPN3** | Pointer to [**[]EPN3Single**](EPN3Single.md) |  | [optional] 
**EPN4** | Pointer to [**[]EPN4Single**](EPN4Single.md) |  | [optional] 
**EPN6** | Pointer to [**[]EPN6Single**](EPN6Single.md) |  | [optional] 
**EPN9** | Pointer to [**[]EPN9Single**](EPN9Single.md) |  | [optional] 
**EPS5U** | Pointer to [**[]EPS5USingle**](EPS5USingle.md) |  | [optional] 

## Methods

### NewUpfFunctionSingle

`func NewUpfFunctionSingle(id NullableString, ) *UpfFunctionSingle`

NewUpfFunctionSingle instantiates a new UpfFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpfFunctionSingleWithDefaults

`func NewUpfFunctionSingleWithDefaults() *UpfFunctionSingle`

NewUpfFunctionSingleWithDefaults instantiates a new UpfFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpfFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpfFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpfFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *UpfFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpfFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *UpfFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *UpfFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *UpfFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *UpfFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *UpfFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *UpfFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *UpfFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *UpfFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *UpfFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *UpfFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *UpfFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *UpfFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *UpfFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpfFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpfFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpfFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *UpfFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *UpfFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *UpfFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *UpfFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *UpfFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *UpfFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *UpfFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *UpfFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *UpfFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *UpfFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *UpfFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *UpfFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *UpfFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *UpfFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *UpfFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *UpfFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetEPN3

`func (o *UpfFunctionSingle) GetEPN3() []EPN3Single`

GetEPN3 returns the EPN3 field if non-nil, zero value otherwise.

### GetEPN3Ok

`func (o *UpfFunctionSingle) GetEPN3Ok() (*[]EPN3Single, bool)`

GetEPN3Ok returns a tuple with the EPN3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN3

`func (o *UpfFunctionSingle) SetEPN3(v []EPN3Single)`

SetEPN3 sets EPN3 field to given value.

### HasEPN3

`func (o *UpfFunctionSingle) HasEPN3() bool`

HasEPN3 returns a boolean if a field has been set.

### GetEPN4

`func (o *UpfFunctionSingle) GetEPN4() []EPN4Single`

GetEPN4 returns the EPN4 field if non-nil, zero value otherwise.

### GetEPN4Ok

`func (o *UpfFunctionSingle) GetEPN4Ok() (*[]EPN4Single, bool)`

GetEPN4Ok returns a tuple with the EPN4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN4

`func (o *UpfFunctionSingle) SetEPN4(v []EPN4Single)`

SetEPN4 sets EPN4 field to given value.

### HasEPN4

`func (o *UpfFunctionSingle) HasEPN4() bool`

HasEPN4 returns a boolean if a field has been set.

### GetEPN6

`func (o *UpfFunctionSingle) GetEPN6() []EPN6Single`

GetEPN6 returns the EPN6 field if non-nil, zero value otherwise.

### GetEPN6Ok

`func (o *UpfFunctionSingle) GetEPN6Ok() (*[]EPN6Single, bool)`

GetEPN6Ok returns a tuple with the EPN6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN6

`func (o *UpfFunctionSingle) SetEPN6(v []EPN6Single)`

SetEPN6 sets EPN6 field to given value.

### HasEPN6

`func (o *UpfFunctionSingle) HasEPN6() bool`

HasEPN6 returns a boolean if a field has been set.

### GetEPN9

`func (o *UpfFunctionSingle) GetEPN9() []EPN9Single`

GetEPN9 returns the EPN9 field if non-nil, zero value otherwise.

### GetEPN9Ok

`func (o *UpfFunctionSingle) GetEPN9Ok() (*[]EPN9Single, bool)`

GetEPN9Ok returns a tuple with the EPN9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN9

`func (o *UpfFunctionSingle) SetEPN9(v []EPN9Single)`

SetEPN9 sets EPN9 field to given value.

### HasEPN9

`func (o *UpfFunctionSingle) HasEPN9() bool`

HasEPN9 returns a boolean if a field has been set.

### GetEPS5U

`func (o *UpfFunctionSingle) GetEPS5U() []EPS5USingle`

GetEPS5U returns the EPS5U field if non-nil, zero value otherwise.

### GetEPS5UOk

`func (o *UpfFunctionSingle) GetEPS5UOk() (*[]EPS5USingle, bool)`

GetEPS5UOk returns a tuple with the EPS5U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPS5U

`func (o *UpfFunctionSingle) SetEPS5U(v []EPS5USingle)`

SetEPS5U sets EPS5U field to given value.

### HasEPS5U

`func (o *UpfFunctionSingle) HasEPS5U() bool`

HasEPS5U returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


