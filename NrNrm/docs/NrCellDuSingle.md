# NrCellDuSingle

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
**RRMPolicyRatio** | Pointer to [**[]RRMPolicyRatioSingle**](RRMPolicyRatioSingle.md) |  | [optional] 
**CPCIConfigurationFunction** | Pointer to [**CPCIConfigurationFunctionSingle**](CPCIConfigurationFunctionSingle.md) |  | [optional] 
**DRACHOptimizationFunction** | Pointer to [**DRACHOptimizationFunctionSingle**](DRACHOptimizationFunctionSingle.md) |  | [optional] 
**NrOperatorCellDu** | Pointer to [**[]NrOperatorCellDuSingle**](NrOperatorCellDuSingle.md) |  | [optional] 

## Methods

### NewNrCellDuSingle

`func NewNrCellDuSingle(id NullableString, ) *NrCellDuSingle`

NewNrCellDuSingle instantiates a new NrCellDuSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrCellDuSingleWithDefaults

`func NewNrCellDuSingleWithDefaults() *NrCellDuSingle`

NewNrCellDuSingleWithDefaults instantiates a new NrCellDuSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NrCellDuSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NrCellDuSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NrCellDuSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *NrCellDuSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NrCellDuSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *NrCellDuSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *NrCellDuSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *NrCellDuSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *NrCellDuSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *NrCellDuSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *NrCellDuSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *NrCellDuSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *NrCellDuSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *NrCellDuSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *NrCellDuSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *NrCellDuSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *NrCellDuSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *NrCellDuSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NrCellDuSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NrCellDuSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NrCellDuSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *NrCellDuSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *NrCellDuSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *NrCellDuSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *NrCellDuSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *NrCellDuSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *NrCellDuSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *NrCellDuSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *NrCellDuSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *NrCellDuSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *NrCellDuSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *NrCellDuSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *NrCellDuSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *NrCellDuSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *NrCellDuSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *NrCellDuSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *NrCellDuSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetRRMPolicyRatio

`func (o *NrCellDuSingle) GetRRMPolicyRatio() []RRMPolicyRatioSingle`

GetRRMPolicyRatio returns the RRMPolicyRatio field if non-nil, zero value otherwise.

### GetRRMPolicyRatioOk

`func (o *NrCellDuSingle) GetRRMPolicyRatioOk() (*[]RRMPolicyRatioSingle, bool)`

GetRRMPolicyRatioOk returns a tuple with the RRMPolicyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRRMPolicyRatio

`func (o *NrCellDuSingle) SetRRMPolicyRatio(v []RRMPolicyRatioSingle)`

SetRRMPolicyRatio sets RRMPolicyRatio field to given value.

### HasRRMPolicyRatio

`func (o *NrCellDuSingle) HasRRMPolicyRatio() bool`

HasRRMPolicyRatio returns a boolean if a field has been set.

### GetCPCIConfigurationFunction

`func (o *NrCellDuSingle) GetCPCIConfigurationFunction() CPCIConfigurationFunctionSingle`

GetCPCIConfigurationFunction returns the CPCIConfigurationFunction field if non-nil, zero value otherwise.

### GetCPCIConfigurationFunctionOk

`func (o *NrCellDuSingle) GetCPCIConfigurationFunctionOk() (*CPCIConfigurationFunctionSingle, bool)`

GetCPCIConfigurationFunctionOk returns a tuple with the CPCIConfigurationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPCIConfigurationFunction

`func (o *NrCellDuSingle) SetCPCIConfigurationFunction(v CPCIConfigurationFunctionSingle)`

SetCPCIConfigurationFunction sets CPCIConfigurationFunction field to given value.

### HasCPCIConfigurationFunction

`func (o *NrCellDuSingle) HasCPCIConfigurationFunction() bool`

HasCPCIConfigurationFunction returns a boolean if a field has been set.

### GetDRACHOptimizationFunction

`func (o *NrCellDuSingle) GetDRACHOptimizationFunction() DRACHOptimizationFunctionSingle`

GetDRACHOptimizationFunction returns the DRACHOptimizationFunction field if non-nil, zero value otherwise.

### GetDRACHOptimizationFunctionOk

`func (o *NrCellDuSingle) GetDRACHOptimizationFunctionOk() (*DRACHOptimizationFunctionSingle, bool)`

GetDRACHOptimizationFunctionOk returns a tuple with the DRACHOptimizationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDRACHOptimizationFunction

`func (o *NrCellDuSingle) SetDRACHOptimizationFunction(v DRACHOptimizationFunctionSingle)`

SetDRACHOptimizationFunction sets DRACHOptimizationFunction field to given value.

### HasDRACHOptimizationFunction

`func (o *NrCellDuSingle) HasDRACHOptimizationFunction() bool`

HasDRACHOptimizationFunction returns a boolean if a field has been set.

### GetNrOperatorCellDu

`func (o *NrCellDuSingle) GetNrOperatorCellDu() []NrOperatorCellDuSingle`

GetNrOperatorCellDu returns the NrOperatorCellDu field if non-nil, zero value otherwise.

### GetNrOperatorCellDuOk

`func (o *NrCellDuSingle) GetNrOperatorCellDuOk() (*[]NrOperatorCellDuSingle, bool)`

GetNrOperatorCellDuOk returns a tuple with the NrOperatorCellDu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrOperatorCellDu

`func (o *NrCellDuSingle) SetNrOperatorCellDu(v []NrOperatorCellDuSingle)`

SetNrOperatorCellDu sets NrOperatorCellDu field to given value.

### HasNrOperatorCellDu

`func (o *NrCellDuSingle) HasNrOperatorCellDu() bool`

HasNrOperatorCellDu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


