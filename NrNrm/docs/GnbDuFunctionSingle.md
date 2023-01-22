# GnbDuFunctionSingle

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
**NrCellDu** | Pointer to [**[]NrCellDuSingle**](NrCellDuSingle.md) |  | [optional] 
**BwpMultiple** | Pointer to [**[]BwpSingle**](BwpSingle.md) |  | [optional] 
**NrSectorCarrierMultiple** | Pointer to [**[]NrSectorCarrierSingle**](NrSectorCarrierSingle.md) |  | [optional] 
**EPF1C** | Pointer to [**EPF1CSingle**](EPF1CSingle.md) |  | [optional] 
**EPF1U** | Pointer to [**[]EPF1USingle**](EPF1USingle.md) |  | [optional] 
**DRACHOptimizationFunction** | Pointer to [**DRACHOptimizationFunctionSingle**](DRACHOptimizationFunctionSingle.md) |  | [optional] 
**OperatorDU** | Pointer to [**[]OperatorDuSingle**](OperatorDuSingle.md) |  | [optional] 
**BWPSet** | Pointer to [**[]BWPSetSingle**](BWPSetSingle.md) |  | [optional] 

## Methods

### NewGnbDuFunctionSingle

`func NewGnbDuFunctionSingle(id NullableString, ) *GnbDuFunctionSingle`

NewGnbDuFunctionSingle instantiates a new GnbDuFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGnbDuFunctionSingleWithDefaults

`func NewGnbDuFunctionSingleWithDefaults() *GnbDuFunctionSingle`

NewGnbDuFunctionSingleWithDefaults instantiates a new GnbDuFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GnbDuFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GnbDuFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GnbDuFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *GnbDuFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GnbDuFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *GnbDuFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *GnbDuFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *GnbDuFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *GnbDuFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *GnbDuFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *GnbDuFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *GnbDuFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *GnbDuFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *GnbDuFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *GnbDuFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *GnbDuFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *GnbDuFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *GnbDuFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GnbDuFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GnbDuFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GnbDuFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *GnbDuFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *GnbDuFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *GnbDuFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *GnbDuFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *GnbDuFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *GnbDuFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *GnbDuFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *GnbDuFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *GnbDuFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *GnbDuFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *GnbDuFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *GnbDuFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *GnbDuFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *GnbDuFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *GnbDuFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *GnbDuFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetRRMPolicyRatio

`func (o *GnbDuFunctionSingle) GetRRMPolicyRatio() []RRMPolicyRatioSingle`

GetRRMPolicyRatio returns the RRMPolicyRatio field if non-nil, zero value otherwise.

### GetRRMPolicyRatioOk

`func (o *GnbDuFunctionSingle) GetRRMPolicyRatioOk() (*[]RRMPolicyRatioSingle, bool)`

GetRRMPolicyRatioOk returns a tuple with the RRMPolicyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRRMPolicyRatio

`func (o *GnbDuFunctionSingle) SetRRMPolicyRatio(v []RRMPolicyRatioSingle)`

SetRRMPolicyRatio sets RRMPolicyRatio field to given value.

### HasRRMPolicyRatio

`func (o *GnbDuFunctionSingle) HasRRMPolicyRatio() bool`

HasRRMPolicyRatio returns a boolean if a field has been set.

### GetNrCellDu

`func (o *GnbDuFunctionSingle) GetNrCellDu() []NrCellDuSingle`

GetNrCellDu returns the NrCellDu field if non-nil, zero value otherwise.

### GetNrCellDuOk

`func (o *GnbDuFunctionSingle) GetNrCellDuOk() (*[]NrCellDuSingle, bool)`

GetNrCellDuOk returns a tuple with the NrCellDu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrCellDu

`func (o *GnbDuFunctionSingle) SetNrCellDu(v []NrCellDuSingle)`

SetNrCellDu sets NrCellDu field to given value.

### HasNrCellDu

`func (o *GnbDuFunctionSingle) HasNrCellDu() bool`

HasNrCellDu returns a boolean if a field has been set.

### GetBwpMultiple

`func (o *GnbDuFunctionSingle) GetBwpMultiple() []BwpSingle`

GetBwpMultiple returns the BwpMultiple field if non-nil, zero value otherwise.

### GetBwpMultipleOk

`func (o *GnbDuFunctionSingle) GetBwpMultipleOk() (*[]BwpSingle, bool)`

GetBwpMultipleOk returns a tuple with the BwpMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwpMultiple

`func (o *GnbDuFunctionSingle) SetBwpMultiple(v []BwpSingle)`

SetBwpMultiple sets BwpMultiple field to given value.

### HasBwpMultiple

`func (o *GnbDuFunctionSingle) HasBwpMultiple() bool`

HasBwpMultiple returns a boolean if a field has been set.

### GetNrSectorCarrierMultiple

`func (o *GnbDuFunctionSingle) GetNrSectorCarrierMultiple() []NrSectorCarrierSingle`

GetNrSectorCarrierMultiple returns the NrSectorCarrierMultiple field if non-nil, zero value otherwise.

### GetNrSectorCarrierMultipleOk

`func (o *GnbDuFunctionSingle) GetNrSectorCarrierMultipleOk() (*[]NrSectorCarrierSingle, bool)`

GetNrSectorCarrierMultipleOk returns a tuple with the NrSectorCarrierMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrSectorCarrierMultiple

`func (o *GnbDuFunctionSingle) SetNrSectorCarrierMultiple(v []NrSectorCarrierSingle)`

SetNrSectorCarrierMultiple sets NrSectorCarrierMultiple field to given value.

### HasNrSectorCarrierMultiple

`func (o *GnbDuFunctionSingle) HasNrSectorCarrierMultiple() bool`

HasNrSectorCarrierMultiple returns a boolean if a field has been set.

### GetEPF1C

`func (o *GnbDuFunctionSingle) GetEPF1C() EPF1CSingle`

GetEPF1C returns the EPF1C field if non-nil, zero value otherwise.

### GetEPF1COk

`func (o *GnbDuFunctionSingle) GetEPF1COk() (*EPF1CSingle, bool)`

GetEPF1COk returns a tuple with the EPF1C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPF1C

`func (o *GnbDuFunctionSingle) SetEPF1C(v EPF1CSingle)`

SetEPF1C sets EPF1C field to given value.

### HasEPF1C

`func (o *GnbDuFunctionSingle) HasEPF1C() bool`

HasEPF1C returns a boolean if a field has been set.

### GetEPF1U

`func (o *GnbDuFunctionSingle) GetEPF1U() []EPF1USingle`

GetEPF1U returns the EPF1U field if non-nil, zero value otherwise.

### GetEPF1UOk

`func (o *GnbDuFunctionSingle) GetEPF1UOk() (*[]EPF1USingle, bool)`

GetEPF1UOk returns a tuple with the EPF1U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPF1U

`func (o *GnbDuFunctionSingle) SetEPF1U(v []EPF1USingle)`

SetEPF1U sets EPF1U field to given value.

### HasEPF1U

`func (o *GnbDuFunctionSingle) HasEPF1U() bool`

HasEPF1U returns a boolean if a field has been set.

### GetDRACHOptimizationFunction

`func (o *GnbDuFunctionSingle) GetDRACHOptimizationFunction() DRACHOptimizationFunctionSingle`

GetDRACHOptimizationFunction returns the DRACHOptimizationFunction field if non-nil, zero value otherwise.

### GetDRACHOptimizationFunctionOk

`func (o *GnbDuFunctionSingle) GetDRACHOptimizationFunctionOk() (*DRACHOptimizationFunctionSingle, bool)`

GetDRACHOptimizationFunctionOk returns a tuple with the DRACHOptimizationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDRACHOptimizationFunction

`func (o *GnbDuFunctionSingle) SetDRACHOptimizationFunction(v DRACHOptimizationFunctionSingle)`

SetDRACHOptimizationFunction sets DRACHOptimizationFunction field to given value.

### HasDRACHOptimizationFunction

`func (o *GnbDuFunctionSingle) HasDRACHOptimizationFunction() bool`

HasDRACHOptimizationFunction returns a boolean if a field has been set.

### GetOperatorDU

`func (o *GnbDuFunctionSingle) GetOperatorDU() []OperatorDuSingle`

GetOperatorDU returns the OperatorDU field if non-nil, zero value otherwise.

### GetOperatorDUOk

`func (o *GnbDuFunctionSingle) GetOperatorDUOk() (*[]OperatorDuSingle, bool)`

GetOperatorDUOk returns a tuple with the OperatorDU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorDU

`func (o *GnbDuFunctionSingle) SetOperatorDU(v []OperatorDuSingle)`

SetOperatorDU sets OperatorDU field to given value.

### HasOperatorDU

`func (o *GnbDuFunctionSingle) HasOperatorDU() bool`

HasOperatorDU returns a boolean if a field has been set.

### GetBWPSet

`func (o *GnbDuFunctionSingle) GetBWPSet() []BWPSetSingle`

GetBWPSet returns the BWPSet field if non-nil, zero value otherwise.

### GetBWPSetOk

`func (o *GnbDuFunctionSingle) GetBWPSetOk() (*[]BWPSetSingle, bool)`

GetBWPSetOk returns a tuple with the BWPSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWPSet

`func (o *GnbDuFunctionSingle) SetBWPSet(v []BWPSetSingle)`

SetBWPSet sets BWPSet field to given value.

### HasBWPSet

`func (o *GnbDuFunctionSingle) HasBWPSet() bool`

HasBWPSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


