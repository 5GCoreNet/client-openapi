# SmfFunctionSingle

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
**EPN4** | Pointer to [**[]EPN4Single**](EPN4Single.md) |  | [optional] 
**EPN7** | Pointer to [**[]EPN7Single**](EPN7Single.md) |  | [optional] 
**EPN10** | Pointer to [**[]EPN10Single**](EPN10Single.md) |  | [optional] 
**EPN11** | Pointer to [**[]EPN11Single**](EPN11Single.md) |  | [optional] 
**EPN16** | Pointer to [**[]EPN16Single**](EPN16Single.md) |  | [optional] 
**EPS5C** | Pointer to [**[]EPS5CSingle**](EPS5CSingle.md) |  | [optional] 
**FiveQiDscpMappingSet** | Pointer to [**FiveQiDscpMappingSetSingle**](FiveQiDscpMappingSetSingle.md) |  | [optional] 
**GtpUPathQoSMonitoringControl** | Pointer to [**GtpUPathQoSMonitoringControlSingle**](GtpUPathQoSMonitoringControlSingle.md) |  | [optional] 
**QFQoSMonitoringControl** | Pointer to [**QFQoSMonitoringControlSingle**](QFQoSMonitoringControlSingle.md) |  | [optional] 
**PredefinedPccRuleSet** | Pointer to [**PredefinedPccRuleSetSingle**](PredefinedPccRuleSetSingle.md) |  | [optional] 

## Methods

### NewSmfFunctionSingle

`func NewSmfFunctionSingle(id NullableString, ) *SmfFunctionSingle`

NewSmfFunctionSingle instantiates a new SmfFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmfFunctionSingleWithDefaults

`func NewSmfFunctionSingleWithDefaults() *SmfFunctionSingle`

NewSmfFunctionSingleWithDefaults instantiates a new SmfFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmfFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmfFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmfFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *SmfFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SmfFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *SmfFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *SmfFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *SmfFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *SmfFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *SmfFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *SmfFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *SmfFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *SmfFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *SmfFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *SmfFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *SmfFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *SmfFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *SmfFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SmfFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SmfFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SmfFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *SmfFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *SmfFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *SmfFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *SmfFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *SmfFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *SmfFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *SmfFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *SmfFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *SmfFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *SmfFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *SmfFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *SmfFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *SmfFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *SmfFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *SmfFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *SmfFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetEPN4

`func (o *SmfFunctionSingle) GetEPN4() []EPN4Single`

GetEPN4 returns the EPN4 field if non-nil, zero value otherwise.

### GetEPN4Ok

`func (o *SmfFunctionSingle) GetEPN4Ok() (*[]EPN4Single, bool)`

GetEPN4Ok returns a tuple with the EPN4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN4

`func (o *SmfFunctionSingle) SetEPN4(v []EPN4Single)`

SetEPN4 sets EPN4 field to given value.

### HasEPN4

`func (o *SmfFunctionSingle) HasEPN4() bool`

HasEPN4 returns a boolean if a field has been set.

### GetEPN7

`func (o *SmfFunctionSingle) GetEPN7() []EPN7Single`

GetEPN7 returns the EPN7 field if non-nil, zero value otherwise.

### GetEPN7Ok

`func (o *SmfFunctionSingle) GetEPN7Ok() (*[]EPN7Single, bool)`

GetEPN7Ok returns a tuple with the EPN7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN7

`func (o *SmfFunctionSingle) SetEPN7(v []EPN7Single)`

SetEPN7 sets EPN7 field to given value.

### HasEPN7

`func (o *SmfFunctionSingle) HasEPN7() bool`

HasEPN7 returns a boolean if a field has been set.

### GetEPN10

`func (o *SmfFunctionSingle) GetEPN10() []EPN10Single`

GetEPN10 returns the EPN10 field if non-nil, zero value otherwise.

### GetEPN10Ok

`func (o *SmfFunctionSingle) GetEPN10Ok() (*[]EPN10Single, bool)`

GetEPN10Ok returns a tuple with the EPN10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN10

`func (o *SmfFunctionSingle) SetEPN10(v []EPN10Single)`

SetEPN10 sets EPN10 field to given value.

### HasEPN10

`func (o *SmfFunctionSingle) HasEPN10() bool`

HasEPN10 returns a boolean if a field has been set.

### GetEPN11

`func (o *SmfFunctionSingle) GetEPN11() []EPN11Single`

GetEPN11 returns the EPN11 field if non-nil, zero value otherwise.

### GetEPN11Ok

`func (o *SmfFunctionSingle) GetEPN11Ok() (*[]EPN11Single, bool)`

GetEPN11Ok returns a tuple with the EPN11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN11

`func (o *SmfFunctionSingle) SetEPN11(v []EPN11Single)`

SetEPN11 sets EPN11 field to given value.

### HasEPN11

`func (o *SmfFunctionSingle) HasEPN11() bool`

HasEPN11 returns a boolean if a field has been set.

### GetEPN16

`func (o *SmfFunctionSingle) GetEPN16() []EPN16Single`

GetEPN16 returns the EPN16 field if non-nil, zero value otherwise.

### GetEPN16Ok

`func (o *SmfFunctionSingle) GetEPN16Ok() (*[]EPN16Single, bool)`

GetEPN16Ok returns a tuple with the EPN16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN16

`func (o *SmfFunctionSingle) SetEPN16(v []EPN16Single)`

SetEPN16 sets EPN16 field to given value.

### HasEPN16

`func (o *SmfFunctionSingle) HasEPN16() bool`

HasEPN16 returns a boolean if a field has been set.

### GetEPS5C

`func (o *SmfFunctionSingle) GetEPS5C() []EPS5CSingle`

GetEPS5C returns the EPS5C field if non-nil, zero value otherwise.

### GetEPS5COk

`func (o *SmfFunctionSingle) GetEPS5COk() (*[]EPS5CSingle, bool)`

GetEPS5COk returns a tuple with the EPS5C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPS5C

`func (o *SmfFunctionSingle) SetEPS5C(v []EPS5CSingle)`

SetEPS5C sets EPS5C field to given value.

### HasEPS5C

`func (o *SmfFunctionSingle) HasEPS5C() bool`

HasEPS5C returns a boolean if a field has been set.

### GetFiveQiDscpMappingSet

`func (o *SmfFunctionSingle) GetFiveQiDscpMappingSet() FiveQiDscpMappingSetSingle`

GetFiveQiDscpMappingSet returns the FiveQiDscpMappingSet field if non-nil, zero value otherwise.

### GetFiveQiDscpMappingSetOk

`func (o *SmfFunctionSingle) GetFiveQiDscpMappingSetOk() (*FiveQiDscpMappingSetSingle, bool)`

GetFiveQiDscpMappingSetOk returns a tuple with the FiveQiDscpMappingSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveQiDscpMappingSet

`func (o *SmfFunctionSingle) SetFiveQiDscpMappingSet(v FiveQiDscpMappingSetSingle)`

SetFiveQiDscpMappingSet sets FiveQiDscpMappingSet field to given value.

### HasFiveQiDscpMappingSet

`func (o *SmfFunctionSingle) HasFiveQiDscpMappingSet() bool`

HasFiveQiDscpMappingSet returns a boolean if a field has been set.

### GetGtpUPathQoSMonitoringControl

`func (o *SmfFunctionSingle) GetGtpUPathQoSMonitoringControl() GtpUPathQoSMonitoringControlSingle`

GetGtpUPathQoSMonitoringControl returns the GtpUPathQoSMonitoringControl field if non-nil, zero value otherwise.

### GetGtpUPathQoSMonitoringControlOk

`func (o *SmfFunctionSingle) GetGtpUPathQoSMonitoringControlOk() (*GtpUPathQoSMonitoringControlSingle, bool)`

GetGtpUPathQoSMonitoringControlOk returns a tuple with the GtpUPathQoSMonitoringControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtpUPathQoSMonitoringControl

`func (o *SmfFunctionSingle) SetGtpUPathQoSMonitoringControl(v GtpUPathQoSMonitoringControlSingle)`

SetGtpUPathQoSMonitoringControl sets GtpUPathQoSMonitoringControl field to given value.

### HasGtpUPathQoSMonitoringControl

`func (o *SmfFunctionSingle) HasGtpUPathQoSMonitoringControl() bool`

HasGtpUPathQoSMonitoringControl returns a boolean if a field has been set.

### GetQFQoSMonitoringControl

`func (o *SmfFunctionSingle) GetQFQoSMonitoringControl() QFQoSMonitoringControlSingle`

GetQFQoSMonitoringControl returns the QFQoSMonitoringControl field if non-nil, zero value otherwise.

### GetQFQoSMonitoringControlOk

`func (o *SmfFunctionSingle) GetQFQoSMonitoringControlOk() (*QFQoSMonitoringControlSingle, bool)`

GetQFQoSMonitoringControlOk returns a tuple with the QFQoSMonitoringControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQFQoSMonitoringControl

`func (o *SmfFunctionSingle) SetQFQoSMonitoringControl(v QFQoSMonitoringControlSingle)`

SetQFQoSMonitoringControl sets QFQoSMonitoringControl field to given value.

### HasQFQoSMonitoringControl

`func (o *SmfFunctionSingle) HasQFQoSMonitoringControl() bool`

HasQFQoSMonitoringControl returns a boolean if a field has been set.

### GetPredefinedPccRuleSet

`func (o *SmfFunctionSingle) GetPredefinedPccRuleSet() PredefinedPccRuleSetSingle`

GetPredefinedPccRuleSet returns the PredefinedPccRuleSet field if non-nil, zero value otherwise.

### GetPredefinedPccRuleSetOk

`func (o *SmfFunctionSingle) GetPredefinedPccRuleSetOk() (*PredefinedPccRuleSetSingle, bool)`

GetPredefinedPccRuleSetOk returns a tuple with the PredefinedPccRuleSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedPccRuleSet

`func (o *SmfFunctionSingle) SetPredefinedPccRuleSet(v PredefinedPccRuleSetSingle)`

SetPredefinedPccRuleSet sets PredefinedPccRuleSet field to given value.

### HasPredefinedPccRuleSet

`func (o *SmfFunctionSingle) HasPredefinedPccRuleSet() bool`

HasPredefinedPccRuleSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


