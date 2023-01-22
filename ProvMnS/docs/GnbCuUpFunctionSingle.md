# GnbCuUpFunctionSingle

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
**EPE1** | Pointer to [**EPE1Single**](EPE1Single.md) |  | [optional] 
**EPXnU** | Pointer to [**[]EPXnUSingle**](EPXnUSingle.md) |  | [optional] 
**EPF1U** | Pointer to [**[]EPF1USingle**](EPF1USingle.md) |  | [optional] 
**EPNgU** | Pointer to [**[]EPNgUSingle**](EPNgUSingle.md) |  | [optional] 
**EPX2U** | Pointer to [**[]EPX2USingle**](EPX2USingle.md) |  | [optional] 
**EPS1U** | Pointer to [**[]EPS1USingle**](EPS1USingle.md) |  | [optional] 

## Methods

### NewGnbCuUpFunctionSingle

`func NewGnbCuUpFunctionSingle(id NullableString, ) *GnbCuUpFunctionSingle`

NewGnbCuUpFunctionSingle instantiates a new GnbCuUpFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGnbCuUpFunctionSingleWithDefaults

`func NewGnbCuUpFunctionSingleWithDefaults() *GnbCuUpFunctionSingle`

NewGnbCuUpFunctionSingleWithDefaults instantiates a new GnbCuUpFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GnbCuUpFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GnbCuUpFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GnbCuUpFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *GnbCuUpFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GnbCuUpFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *GnbCuUpFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *GnbCuUpFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *GnbCuUpFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *GnbCuUpFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *GnbCuUpFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *GnbCuUpFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *GnbCuUpFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *GnbCuUpFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *GnbCuUpFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *GnbCuUpFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *GnbCuUpFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *GnbCuUpFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *GnbCuUpFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GnbCuUpFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GnbCuUpFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GnbCuUpFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *GnbCuUpFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *GnbCuUpFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *GnbCuUpFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *GnbCuUpFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *GnbCuUpFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *GnbCuUpFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *GnbCuUpFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *GnbCuUpFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *GnbCuUpFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *GnbCuUpFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *GnbCuUpFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *GnbCuUpFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *GnbCuUpFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *GnbCuUpFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *GnbCuUpFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *GnbCuUpFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetRRMPolicyRatio

`func (o *GnbCuUpFunctionSingle) GetRRMPolicyRatio() []RRMPolicyRatioSingle`

GetRRMPolicyRatio returns the RRMPolicyRatio field if non-nil, zero value otherwise.

### GetRRMPolicyRatioOk

`func (o *GnbCuUpFunctionSingle) GetRRMPolicyRatioOk() (*[]RRMPolicyRatioSingle, bool)`

GetRRMPolicyRatioOk returns a tuple with the RRMPolicyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRRMPolicyRatio

`func (o *GnbCuUpFunctionSingle) SetRRMPolicyRatio(v []RRMPolicyRatioSingle)`

SetRRMPolicyRatio sets RRMPolicyRatio field to given value.

### HasRRMPolicyRatio

`func (o *GnbCuUpFunctionSingle) HasRRMPolicyRatio() bool`

HasRRMPolicyRatio returns a boolean if a field has been set.

### GetEPE1

`func (o *GnbCuUpFunctionSingle) GetEPE1() EPE1Single`

GetEPE1 returns the EPE1 field if non-nil, zero value otherwise.

### GetEPE1Ok

`func (o *GnbCuUpFunctionSingle) GetEPE1Ok() (*EPE1Single, bool)`

GetEPE1Ok returns a tuple with the EPE1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPE1

`func (o *GnbCuUpFunctionSingle) SetEPE1(v EPE1Single)`

SetEPE1 sets EPE1 field to given value.

### HasEPE1

`func (o *GnbCuUpFunctionSingle) HasEPE1() bool`

HasEPE1 returns a boolean if a field has been set.

### GetEPXnU

`func (o *GnbCuUpFunctionSingle) GetEPXnU() []EPXnUSingle`

GetEPXnU returns the EPXnU field if non-nil, zero value otherwise.

### GetEPXnUOk

`func (o *GnbCuUpFunctionSingle) GetEPXnUOk() (*[]EPXnUSingle, bool)`

GetEPXnUOk returns a tuple with the EPXnU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPXnU

`func (o *GnbCuUpFunctionSingle) SetEPXnU(v []EPXnUSingle)`

SetEPXnU sets EPXnU field to given value.

### HasEPXnU

`func (o *GnbCuUpFunctionSingle) HasEPXnU() bool`

HasEPXnU returns a boolean if a field has been set.

### GetEPF1U

`func (o *GnbCuUpFunctionSingle) GetEPF1U() []EPF1USingle`

GetEPF1U returns the EPF1U field if non-nil, zero value otherwise.

### GetEPF1UOk

`func (o *GnbCuUpFunctionSingle) GetEPF1UOk() (*[]EPF1USingle, bool)`

GetEPF1UOk returns a tuple with the EPF1U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPF1U

`func (o *GnbCuUpFunctionSingle) SetEPF1U(v []EPF1USingle)`

SetEPF1U sets EPF1U field to given value.

### HasEPF1U

`func (o *GnbCuUpFunctionSingle) HasEPF1U() bool`

HasEPF1U returns a boolean if a field has been set.

### GetEPNgU

`func (o *GnbCuUpFunctionSingle) GetEPNgU() []EPNgUSingle`

GetEPNgU returns the EPNgU field if non-nil, zero value otherwise.

### GetEPNgUOk

`func (o *GnbCuUpFunctionSingle) GetEPNgUOk() (*[]EPNgUSingle, bool)`

GetEPNgUOk returns a tuple with the EPNgU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPNgU

`func (o *GnbCuUpFunctionSingle) SetEPNgU(v []EPNgUSingle)`

SetEPNgU sets EPNgU field to given value.

### HasEPNgU

`func (o *GnbCuUpFunctionSingle) HasEPNgU() bool`

HasEPNgU returns a boolean if a field has been set.

### GetEPX2U

`func (o *GnbCuUpFunctionSingle) GetEPX2U() []EPX2USingle`

GetEPX2U returns the EPX2U field if non-nil, zero value otherwise.

### GetEPX2UOk

`func (o *GnbCuUpFunctionSingle) GetEPX2UOk() (*[]EPX2USingle, bool)`

GetEPX2UOk returns a tuple with the EPX2U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPX2U

`func (o *GnbCuUpFunctionSingle) SetEPX2U(v []EPX2USingle)`

SetEPX2U sets EPX2U field to given value.

### HasEPX2U

`func (o *GnbCuUpFunctionSingle) HasEPX2U() bool`

HasEPX2U returns a boolean if a field has been set.

### GetEPS1U

`func (o *GnbCuUpFunctionSingle) GetEPS1U() []EPS1USingle`

GetEPS1U returns the EPS1U field if non-nil, zero value otherwise.

### GetEPS1UOk

`func (o *GnbCuUpFunctionSingle) GetEPS1UOk() (*[]EPS1USingle, bool)`

GetEPS1UOk returns a tuple with the EPS1U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPS1U

`func (o *GnbCuUpFunctionSingle) SetEPS1U(v []EPS1USingle)`

SetEPS1U sets EPS1U field to given value.

### HasEPS1U

`func (o *GnbCuUpFunctionSingle) HasEPS1U() bool`

HasEPS1U returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


