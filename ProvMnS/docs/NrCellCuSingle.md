# NrCellCuSingle

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
**NRCellRelation** | Pointer to [**[]NRCellRelationSingle**](NRCellRelationSingle.md) |  | [optional] 
**EUtranCellRelation** | Pointer to [**[]EUtranCellRelationSingle**](EUtranCellRelationSingle.md) |  | [optional] 
**NRFreqRelation** | Pointer to [**[]NRFreqRelationSingle**](NRFreqRelationSingle.md) |  | [optional] 
**EUtranFreqRelation** | Pointer to [**[]EUtranFreqRelationSingle**](EUtranFreqRelationSingle.md) |  | [optional] 
**DESManagementFunction** | Pointer to [**DESManagementFunctionSingle**](DESManagementFunctionSingle.md) |  | [optional] 
**DMROFunction** | Pointer to [**DMROFunctionSingle**](DMROFunctionSingle.md) |  | [optional] 
**DLBOFunction** | Pointer to [**DLBOFunctionSingle**](DLBOFunctionSingle.md) |  | [optional] 
**CESManagementFunction** | Pointer to [**CESManagementFunctionSingle**](CESManagementFunctionSingle.md) |  | [optional] 
**DPCIConfigurationFunction** | Pointer to [**DPCIConfigurationFunctionSingle**](DPCIConfigurationFunctionSingle.md) |  | [optional] 

## Methods

### NewNrCellCuSingle

`func NewNrCellCuSingle(id NullableString, ) *NrCellCuSingle`

NewNrCellCuSingle instantiates a new NrCellCuSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrCellCuSingleWithDefaults

`func NewNrCellCuSingleWithDefaults() *NrCellCuSingle`

NewNrCellCuSingleWithDefaults instantiates a new NrCellCuSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NrCellCuSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NrCellCuSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NrCellCuSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *NrCellCuSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NrCellCuSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *NrCellCuSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *NrCellCuSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *NrCellCuSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *NrCellCuSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *NrCellCuSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *NrCellCuSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *NrCellCuSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *NrCellCuSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *NrCellCuSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *NrCellCuSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *NrCellCuSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *NrCellCuSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *NrCellCuSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NrCellCuSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NrCellCuSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NrCellCuSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *NrCellCuSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *NrCellCuSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *NrCellCuSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *NrCellCuSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *NrCellCuSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *NrCellCuSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *NrCellCuSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *NrCellCuSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *NrCellCuSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *NrCellCuSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *NrCellCuSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *NrCellCuSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *NrCellCuSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *NrCellCuSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *NrCellCuSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *NrCellCuSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetRRMPolicyRatio

`func (o *NrCellCuSingle) GetRRMPolicyRatio() []RRMPolicyRatioSingle`

GetRRMPolicyRatio returns the RRMPolicyRatio field if non-nil, zero value otherwise.

### GetRRMPolicyRatioOk

`func (o *NrCellCuSingle) GetRRMPolicyRatioOk() (*[]RRMPolicyRatioSingle, bool)`

GetRRMPolicyRatioOk returns a tuple with the RRMPolicyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRRMPolicyRatio

`func (o *NrCellCuSingle) SetRRMPolicyRatio(v []RRMPolicyRatioSingle)`

SetRRMPolicyRatio sets RRMPolicyRatio field to given value.

### HasRRMPolicyRatio

`func (o *NrCellCuSingle) HasRRMPolicyRatio() bool`

HasRRMPolicyRatio returns a boolean if a field has been set.

### GetNRCellRelation

`func (o *NrCellCuSingle) GetNRCellRelation() []NRCellRelationSingle`

GetNRCellRelation returns the NRCellRelation field if non-nil, zero value otherwise.

### GetNRCellRelationOk

`func (o *NrCellCuSingle) GetNRCellRelationOk() (*[]NRCellRelationSingle, bool)`

GetNRCellRelationOk returns a tuple with the NRCellRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRCellRelation

`func (o *NrCellCuSingle) SetNRCellRelation(v []NRCellRelationSingle)`

SetNRCellRelation sets NRCellRelation field to given value.

### HasNRCellRelation

`func (o *NrCellCuSingle) HasNRCellRelation() bool`

HasNRCellRelation returns a boolean if a field has been set.

### GetEUtranCellRelation

`func (o *NrCellCuSingle) GetEUtranCellRelation() []EUtranCellRelationSingle`

GetEUtranCellRelation returns the EUtranCellRelation field if non-nil, zero value otherwise.

### GetEUtranCellRelationOk

`func (o *NrCellCuSingle) GetEUtranCellRelationOk() (*[]EUtranCellRelationSingle, bool)`

GetEUtranCellRelationOk returns a tuple with the EUtranCellRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUtranCellRelation

`func (o *NrCellCuSingle) SetEUtranCellRelation(v []EUtranCellRelationSingle)`

SetEUtranCellRelation sets EUtranCellRelation field to given value.

### HasEUtranCellRelation

`func (o *NrCellCuSingle) HasEUtranCellRelation() bool`

HasEUtranCellRelation returns a boolean if a field has been set.

### GetNRFreqRelation

`func (o *NrCellCuSingle) GetNRFreqRelation() []NRFreqRelationSingle`

GetNRFreqRelation returns the NRFreqRelation field if non-nil, zero value otherwise.

### GetNRFreqRelationOk

`func (o *NrCellCuSingle) GetNRFreqRelationOk() (*[]NRFreqRelationSingle, bool)`

GetNRFreqRelationOk returns a tuple with the NRFreqRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRFreqRelation

`func (o *NrCellCuSingle) SetNRFreqRelation(v []NRFreqRelationSingle)`

SetNRFreqRelation sets NRFreqRelation field to given value.

### HasNRFreqRelation

`func (o *NrCellCuSingle) HasNRFreqRelation() bool`

HasNRFreqRelation returns a boolean if a field has been set.

### GetEUtranFreqRelation

`func (o *NrCellCuSingle) GetEUtranFreqRelation() []EUtranFreqRelationSingle`

GetEUtranFreqRelation returns the EUtranFreqRelation field if non-nil, zero value otherwise.

### GetEUtranFreqRelationOk

`func (o *NrCellCuSingle) GetEUtranFreqRelationOk() (*[]EUtranFreqRelationSingle, bool)`

GetEUtranFreqRelationOk returns a tuple with the EUtranFreqRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUtranFreqRelation

`func (o *NrCellCuSingle) SetEUtranFreqRelation(v []EUtranFreqRelationSingle)`

SetEUtranFreqRelation sets EUtranFreqRelation field to given value.

### HasEUtranFreqRelation

`func (o *NrCellCuSingle) HasEUtranFreqRelation() bool`

HasEUtranFreqRelation returns a boolean if a field has been set.

### GetDESManagementFunction

`func (o *NrCellCuSingle) GetDESManagementFunction() DESManagementFunctionSingle`

GetDESManagementFunction returns the DESManagementFunction field if non-nil, zero value otherwise.

### GetDESManagementFunctionOk

`func (o *NrCellCuSingle) GetDESManagementFunctionOk() (*DESManagementFunctionSingle, bool)`

GetDESManagementFunctionOk returns a tuple with the DESManagementFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDESManagementFunction

`func (o *NrCellCuSingle) SetDESManagementFunction(v DESManagementFunctionSingle)`

SetDESManagementFunction sets DESManagementFunction field to given value.

### HasDESManagementFunction

`func (o *NrCellCuSingle) HasDESManagementFunction() bool`

HasDESManagementFunction returns a boolean if a field has been set.

### GetDMROFunction

`func (o *NrCellCuSingle) GetDMROFunction() DMROFunctionSingle`

GetDMROFunction returns the DMROFunction field if non-nil, zero value otherwise.

### GetDMROFunctionOk

`func (o *NrCellCuSingle) GetDMROFunctionOk() (*DMROFunctionSingle, bool)`

GetDMROFunctionOk returns a tuple with the DMROFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDMROFunction

`func (o *NrCellCuSingle) SetDMROFunction(v DMROFunctionSingle)`

SetDMROFunction sets DMROFunction field to given value.

### HasDMROFunction

`func (o *NrCellCuSingle) HasDMROFunction() bool`

HasDMROFunction returns a boolean if a field has been set.

### GetDLBOFunction

`func (o *NrCellCuSingle) GetDLBOFunction() DLBOFunctionSingle`

GetDLBOFunction returns the DLBOFunction field if non-nil, zero value otherwise.

### GetDLBOFunctionOk

`func (o *NrCellCuSingle) GetDLBOFunctionOk() (*DLBOFunctionSingle, bool)`

GetDLBOFunctionOk returns a tuple with the DLBOFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLBOFunction

`func (o *NrCellCuSingle) SetDLBOFunction(v DLBOFunctionSingle)`

SetDLBOFunction sets DLBOFunction field to given value.

### HasDLBOFunction

`func (o *NrCellCuSingle) HasDLBOFunction() bool`

HasDLBOFunction returns a boolean if a field has been set.

### GetCESManagementFunction

`func (o *NrCellCuSingle) GetCESManagementFunction() CESManagementFunctionSingle`

GetCESManagementFunction returns the CESManagementFunction field if non-nil, zero value otherwise.

### GetCESManagementFunctionOk

`func (o *NrCellCuSingle) GetCESManagementFunctionOk() (*CESManagementFunctionSingle, bool)`

GetCESManagementFunctionOk returns a tuple with the CESManagementFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCESManagementFunction

`func (o *NrCellCuSingle) SetCESManagementFunction(v CESManagementFunctionSingle)`

SetCESManagementFunction sets CESManagementFunction field to given value.

### HasCESManagementFunction

`func (o *NrCellCuSingle) HasCESManagementFunction() bool`

HasCESManagementFunction returns a boolean if a field has been set.

### GetDPCIConfigurationFunction

`func (o *NrCellCuSingle) GetDPCIConfigurationFunction() DPCIConfigurationFunctionSingle`

GetDPCIConfigurationFunction returns the DPCIConfigurationFunction field if non-nil, zero value otherwise.

### GetDPCIConfigurationFunctionOk

`func (o *NrCellCuSingle) GetDPCIConfigurationFunctionOk() (*DPCIConfigurationFunctionSingle, bool)`

GetDPCIConfigurationFunctionOk returns a tuple with the DPCIConfigurationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDPCIConfigurationFunction

`func (o *NrCellCuSingle) SetDPCIConfigurationFunction(v DPCIConfigurationFunctionSingle)`

SetDPCIConfigurationFunction sets DPCIConfigurationFunction field to given value.

### HasDPCIConfigurationFunction

`func (o *NrCellCuSingle) HasDPCIConfigurationFunction() bool`

HasDPCIConfigurationFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


