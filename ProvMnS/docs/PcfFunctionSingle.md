# PcfFunctionSingle

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
**EPN5** | Pointer to [**[]EPN5Single**](EPN5Single.md) |  | [optional] 
**EPN7** | Pointer to [**[]EPN7Single**](EPN7Single.md) |  | [optional] 
**EPN15** | Pointer to [**[]EPN15Single**](EPN15Single.md) |  | [optional] 
**EPN16** | Pointer to [**[]EPN16Single**](EPN16Single.md) |  | [optional] 
**EPRx** | Pointer to [**[]EPRxSingle**](EPRxSingle.md) |  | [optional] 
**PredefinedPccRuleSet** | Pointer to [**PredefinedPccRuleSetSingle**](PredefinedPccRuleSetSingle.md) |  | [optional] 

## Methods

### NewPcfFunctionSingle

`func NewPcfFunctionSingle(id NullableString, ) *PcfFunctionSingle`

NewPcfFunctionSingle instantiates a new PcfFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcfFunctionSingleWithDefaults

`func NewPcfFunctionSingleWithDefaults() *PcfFunctionSingle`

NewPcfFunctionSingleWithDefaults instantiates a new PcfFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PcfFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PcfFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PcfFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *PcfFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PcfFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *PcfFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *PcfFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *PcfFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *PcfFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *PcfFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *PcfFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *PcfFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *PcfFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *PcfFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *PcfFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *PcfFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *PcfFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *PcfFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PcfFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PcfFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PcfFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *PcfFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *PcfFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *PcfFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *PcfFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *PcfFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *PcfFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *PcfFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *PcfFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *PcfFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *PcfFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *PcfFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *PcfFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *PcfFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *PcfFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *PcfFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *PcfFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetEPN5

`func (o *PcfFunctionSingle) GetEPN5() []EPN5Single`

GetEPN5 returns the EPN5 field if non-nil, zero value otherwise.

### GetEPN5Ok

`func (o *PcfFunctionSingle) GetEPN5Ok() (*[]EPN5Single, bool)`

GetEPN5Ok returns a tuple with the EPN5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN5

`func (o *PcfFunctionSingle) SetEPN5(v []EPN5Single)`

SetEPN5 sets EPN5 field to given value.

### HasEPN5

`func (o *PcfFunctionSingle) HasEPN5() bool`

HasEPN5 returns a boolean if a field has been set.

### GetEPN7

`func (o *PcfFunctionSingle) GetEPN7() []EPN7Single`

GetEPN7 returns the EPN7 field if non-nil, zero value otherwise.

### GetEPN7Ok

`func (o *PcfFunctionSingle) GetEPN7Ok() (*[]EPN7Single, bool)`

GetEPN7Ok returns a tuple with the EPN7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN7

`func (o *PcfFunctionSingle) SetEPN7(v []EPN7Single)`

SetEPN7 sets EPN7 field to given value.

### HasEPN7

`func (o *PcfFunctionSingle) HasEPN7() bool`

HasEPN7 returns a boolean if a field has been set.

### GetEPN15

`func (o *PcfFunctionSingle) GetEPN15() []EPN15Single`

GetEPN15 returns the EPN15 field if non-nil, zero value otherwise.

### GetEPN15Ok

`func (o *PcfFunctionSingle) GetEPN15Ok() (*[]EPN15Single, bool)`

GetEPN15Ok returns a tuple with the EPN15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN15

`func (o *PcfFunctionSingle) SetEPN15(v []EPN15Single)`

SetEPN15 sets EPN15 field to given value.

### HasEPN15

`func (o *PcfFunctionSingle) HasEPN15() bool`

HasEPN15 returns a boolean if a field has been set.

### GetEPN16

`func (o *PcfFunctionSingle) GetEPN16() []EPN16Single`

GetEPN16 returns the EPN16 field if non-nil, zero value otherwise.

### GetEPN16Ok

`func (o *PcfFunctionSingle) GetEPN16Ok() (*[]EPN16Single, bool)`

GetEPN16Ok returns a tuple with the EPN16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN16

`func (o *PcfFunctionSingle) SetEPN16(v []EPN16Single)`

SetEPN16 sets EPN16 field to given value.

### HasEPN16

`func (o *PcfFunctionSingle) HasEPN16() bool`

HasEPN16 returns a boolean if a field has been set.

### GetEPRx

`func (o *PcfFunctionSingle) GetEPRx() []EPRxSingle`

GetEPRx returns the EPRx field if non-nil, zero value otherwise.

### GetEPRxOk

`func (o *PcfFunctionSingle) GetEPRxOk() (*[]EPRxSingle, bool)`

GetEPRxOk returns a tuple with the EPRx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPRx

`func (o *PcfFunctionSingle) SetEPRx(v []EPRxSingle)`

SetEPRx sets EPRx field to given value.

### HasEPRx

`func (o *PcfFunctionSingle) HasEPRx() bool`

HasEPRx returns a boolean if a field has been set.

### GetPredefinedPccRuleSet

`func (o *PcfFunctionSingle) GetPredefinedPccRuleSet() PredefinedPccRuleSetSingle`

GetPredefinedPccRuleSet returns the PredefinedPccRuleSet field if non-nil, zero value otherwise.

### GetPredefinedPccRuleSetOk

`func (o *PcfFunctionSingle) GetPredefinedPccRuleSetOk() (*PredefinedPccRuleSetSingle, bool)`

GetPredefinedPccRuleSetOk returns a tuple with the PredefinedPccRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedPccRuleSet

`func (o *PcfFunctionSingle) SetPredefinedPccRuleSet(v PredefinedPccRuleSetSingle)`

SetPredefinedPccRuleSet sets PredefinedPccRuleSet field to given value.

### HasPredefinedPccRuleSet

`func (o *PcfFunctionSingle) HasPredefinedPccRuleSet() bool`

HasPredefinedPccRuleSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


