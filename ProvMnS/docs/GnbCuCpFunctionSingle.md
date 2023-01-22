# GnbCuCpFunctionSingle

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
**NrCellCu** | Pointer to [**[]NrCellCuSingle**](NrCellCuSingle.md) |  | [optional] 
**EPXnC** | Pointer to [**[]EPXnCSingle**](EPXnCSingle.md) |  | [optional] 
**EPE1** | Pointer to [**[]EPE1Single**](EPE1Single.md) |  | [optional] 
**EPF1C** | Pointer to [**[]EPF1CSingle**](EPF1CSingle.md) |  | [optional] 
**EPNgC** | Pointer to [**[]EPNgCSingle**](EPNgCSingle.md) |  | [optional] 
**EPX2C** | Pointer to [**[]EPX2CSingle**](EPX2CSingle.md) |  | [optional] 
**DANRManagementFunction** | Pointer to [**DANRManagementFunctionSingle**](DANRManagementFunctionSingle.md) |  | [optional] 
**DESManagementFunction** | Pointer to [**DESManagementFunctionSingle**](DESManagementFunctionSingle.md) |  | [optional] 
**DMROFunction** | Pointer to [**DMROFunctionSingle**](DMROFunctionSingle.md) |  | [optional] 
**DLBOFunction** | Pointer to [**DLBOFunctionSingle**](DLBOFunctionSingle.md) |  | [optional] 

## Methods

### NewGnbCuCpFunctionSingle

`func NewGnbCuCpFunctionSingle(id NullableString, ) *GnbCuCpFunctionSingle`

NewGnbCuCpFunctionSingle instantiates a new GnbCuCpFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGnbCuCpFunctionSingleWithDefaults

`func NewGnbCuCpFunctionSingleWithDefaults() *GnbCuCpFunctionSingle`

NewGnbCuCpFunctionSingleWithDefaults instantiates a new GnbCuCpFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GnbCuCpFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GnbCuCpFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GnbCuCpFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *GnbCuCpFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GnbCuCpFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *GnbCuCpFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *GnbCuCpFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *GnbCuCpFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *GnbCuCpFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *GnbCuCpFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *GnbCuCpFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *GnbCuCpFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *GnbCuCpFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *GnbCuCpFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *GnbCuCpFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *GnbCuCpFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *GnbCuCpFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *GnbCuCpFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GnbCuCpFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GnbCuCpFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GnbCuCpFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *GnbCuCpFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *GnbCuCpFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *GnbCuCpFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *GnbCuCpFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *GnbCuCpFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *GnbCuCpFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *GnbCuCpFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *GnbCuCpFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *GnbCuCpFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *GnbCuCpFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *GnbCuCpFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *GnbCuCpFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *GnbCuCpFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *GnbCuCpFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *GnbCuCpFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *GnbCuCpFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetRRMPolicyRatio

`func (o *GnbCuCpFunctionSingle) GetRRMPolicyRatio() []RRMPolicyRatioSingle`

GetRRMPolicyRatio returns the RRMPolicyRatio field if non-nil, zero value otherwise.

### GetRRMPolicyRatioOk

`func (o *GnbCuCpFunctionSingle) GetRRMPolicyRatioOk() (*[]RRMPolicyRatioSingle, bool)`

GetRRMPolicyRatioOk returns a tuple with the RRMPolicyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRRMPolicyRatio

`func (o *GnbCuCpFunctionSingle) SetRRMPolicyRatio(v []RRMPolicyRatioSingle)`

SetRRMPolicyRatio sets RRMPolicyRatio field to given value.

### HasRRMPolicyRatio

`func (o *GnbCuCpFunctionSingle) HasRRMPolicyRatio() bool`

HasRRMPolicyRatio returns a boolean if a field has been set.

### GetNrCellCu

`func (o *GnbCuCpFunctionSingle) GetNrCellCu() []NrCellCuSingle`

GetNrCellCu returns the NrCellCu field if non-nil, zero value otherwise.

### GetNrCellCuOk

`func (o *GnbCuCpFunctionSingle) GetNrCellCuOk() (*[]NrCellCuSingle, bool)`

GetNrCellCuOk returns a tuple with the NrCellCu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrCellCu

`func (o *GnbCuCpFunctionSingle) SetNrCellCu(v []NrCellCuSingle)`

SetNrCellCu sets NrCellCu field to given value.

### HasNrCellCu

`func (o *GnbCuCpFunctionSingle) HasNrCellCu() bool`

HasNrCellCu returns a boolean if a field has been set.

### GetEPXnC

`func (o *GnbCuCpFunctionSingle) GetEPXnC() []EPXnCSingle`

GetEPXnC returns the EPXnC field if non-nil, zero value otherwise.

### GetEPXnCOk

`func (o *GnbCuCpFunctionSingle) GetEPXnCOk() (*[]EPXnCSingle, bool)`

GetEPXnCOk returns a tuple with the EPXnC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPXnC

`func (o *GnbCuCpFunctionSingle) SetEPXnC(v []EPXnCSingle)`

SetEPXnC sets EPXnC field to given value.

### HasEPXnC

`func (o *GnbCuCpFunctionSingle) HasEPXnC() bool`

HasEPXnC returns a boolean if a field has been set.

### GetEPE1

`func (o *GnbCuCpFunctionSingle) GetEPE1() []EPE1Single`

GetEPE1 returns the EPE1 field if non-nil, zero value otherwise.

### GetEPE1Ok

`func (o *GnbCuCpFunctionSingle) GetEPE1Ok() (*[]EPE1Single, bool)`

GetEPE1Ok returns a tuple with the EPE1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPE1

`func (o *GnbCuCpFunctionSingle) SetEPE1(v []EPE1Single)`

SetEPE1 sets EPE1 field to given value.

### HasEPE1

`func (o *GnbCuCpFunctionSingle) HasEPE1() bool`

HasEPE1 returns a boolean if a field has been set.

### GetEPF1C

`func (o *GnbCuCpFunctionSingle) GetEPF1C() []EPF1CSingle`

GetEPF1C returns the EPF1C field if non-nil, zero value otherwise.

### GetEPF1COk

`func (o *GnbCuCpFunctionSingle) GetEPF1COk() (*[]EPF1CSingle, bool)`

GetEPF1COk returns a tuple with the EPF1C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPF1C

`func (o *GnbCuCpFunctionSingle) SetEPF1C(v []EPF1CSingle)`

SetEPF1C sets EPF1C field to given value.

### HasEPF1C

`func (o *GnbCuCpFunctionSingle) HasEPF1C() bool`

HasEPF1C returns a boolean if a field has been set.

### GetEPNgC

`func (o *GnbCuCpFunctionSingle) GetEPNgC() []EPNgCSingle`

GetEPNgC returns the EPNgC field if non-nil, zero value otherwise.

### GetEPNgCOk

`func (o *GnbCuCpFunctionSingle) GetEPNgCOk() (*[]EPNgCSingle, bool)`

GetEPNgCOk returns a tuple with the EPNgC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPNgC

`func (o *GnbCuCpFunctionSingle) SetEPNgC(v []EPNgCSingle)`

SetEPNgC sets EPNgC field to given value.

### HasEPNgC

`func (o *GnbCuCpFunctionSingle) HasEPNgC() bool`

HasEPNgC returns a boolean if a field has been set.

### GetEPX2C

`func (o *GnbCuCpFunctionSingle) GetEPX2C() []EPX2CSingle`

GetEPX2C returns the EPX2C field if non-nil, zero value otherwise.

### GetEPX2COk

`func (o *GnbCuCpFunctionSingle) GetEPX2COk() (*[]EPX2CSingle, bool)`

GetEPX2COk returns a tuple with the EPX2C field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPX2C

`func (o *GnbCuCpFunctionSingle) SetEPX2C(v []EPX2CSingle)`

SetEPX2C sets EPX2C field to given value.

### HasEPX2C

`func (o *GnbCuCpFunctionSingle) HasEPX2C() bool`

HasEPX2C returns a boolean if a field has been set.

### GetDANRManagementFunction

`func (o *GnbCuCpFunctionSingle) GetDANRManagementFunction() DANRManagementFunctionSingle`

GetDANRManagementFunction returns the DANRManagementFunction field if non-nil, zero value otherwise.

### GetDANRManagementFunctionOk

`func (o *GnbCuCpFunctionSingle) GetDANRManagementFunctionOk() (*DANRManagementFunctionSingle, bool)`

GetDANRManagementFunctionOk returns a tuple with the DANRManagementFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDANRManagementFunction

`func (o *GnbCuCpFunctionSingle) SetDANRManagementFunction(v DANRManagementFunctionSingle)`

SetDANRManagementFunction sets DANRManagementFunction field to given value.

### HasDANRManagementFunction

`func (o *GnbCuCpFunctionSingle) HasDANRManagementFunction() bool`

HasDANRManagementFunction returns a boolean if a field has been set.

### GetDESManagementFunction

`func (o *GnbCuCpFunctionSingle) GetDESManagementFunction() DESManagementFunctionSingle`

GetDESManagementFunction returns the DESManagementFunction field if non-nil, zero value otherwise.

### GetDESManagementFunctionOk

`func (o *GnbCuCpFunctionSingle) GetDESManagementFunctionOk() (*DESManagementFunctionSingle, bool)`

GetDESManagementFunctionOk returns a tuple with the DESManagementFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDESManagementFunction

`func (o *GnbCuCpFunctionSingle) SetDESManagementFunction(v DESManagementFunctionSingle)`

SetDESManagementFunction sets DESManagementFunction field to given value.

### HasDESManagementFunction

`func (o *GnbCuCpFunctionSingle) HasDESManagementFunction() bool`

HasDESManagementFunction returns a boolean if a field has been set.

### GetDMROFunction

`func (o *GnbCuCpFunctionSingle) GetDMROFunction() DMROFunctionSingle`

GetDMROFunction returns the DMROFunction field if non-nil, zero value otherwise.

### GetDMROFunctionOk

`func (o *GnbCuCpFunctionSingle) GetDMROFunctionOk() (*DMROFunctionSingle, bool)`

GetDMROFunctionOk returns a tuple with the DMROFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDMROFunction

`func (o *GnbCuCpFunctionSingle) SetDMROFunction(v DMROFunctionSingle)`

SetDMROFunction sets DMROFunction field to given value.

### HasDMROFunction

`func (o *GnbCuCpFunctionSingle) HasDMROFunction() bool`

HasDMROFunction returns a boolean if a field has been set.

### GetDLBOFunction

`func (o *GnbCuCpFunctionSingle) GetDLBOFunction() DLBOFunctionSingle`

GetDLBOFunction returns the DLBOFunction field if non-nil, zero value otherwise.

### GetDLBOFunctionOk

`func (o *GnbCuCpFunctionSingle) GetDLBOFunctionOk() (*DLBOFunctionSingle, bool)`

GetDLBOFunctionOk returns a tuple with the DLBOFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLBOFunction

`func (o *GnbCuCpFunctionSingle) SetDLBOFunction(v DLBOFunctionSingle)`

SetDLBOFunction sets DLBOFunction field to given value.

### HasDLBOFunction

`func (o *GnbCuCpFunctionSingle) HasDLBOFunction() bool`

HasDLBOFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


