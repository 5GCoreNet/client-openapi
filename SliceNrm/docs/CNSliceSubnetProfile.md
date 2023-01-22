# CNSliceSubnetProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxNumberofUEs** | Pointer to **int32** |  | [optional] 
**DLLatency** | Pointer to **float32** |  | [optional] 
**ULLatency** | Pointer to **float32** |  | [optional] 
**DLThptPerSliceSubnet** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**DLThptPerUE** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**ULThptPerSliceSubnet** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**ULThptPerUE** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**MaxNumberOfPDUSessions** | Pointer to **int32** |  | [optional] 
**CoverageAreaTAList** | Pointer to **[]int32** |  | [optional] 
**ResourceSharingLevel** | Pointer to [**SharingLevel**](SharingLevel.md) |  | [optional] 
**DLMaxPktSize** | Pointer to **int32** |  | [optional] 
**ULMaxPktSize** | Pointer to **int32** |  | [optional] 
**DelayTolerance** | Pointer to [**DelayTolerance**](DelayTolerance.md) |  | [optional] 
**Synchronicity** | Pointer to [**SynchronicityRANSubnet**](SynchronicityRANSubnet.md) |  | [optional] 
**SliceSimultaneousUse** | Pointer to [**SliceSimultaneousUse**](SliceSimultaneousUse.md) |  | [optional] 
**Reliability** | Pointer to **float32** |  | [optional] 
**EnergyEfficiency** | Pointer to **float32** |  | [optional] 
**DLDeterministicComm** | Pointer to [**DeterministicComm**](DeterministicComm.md) |  | [optional] 
**ULDeterministicComm** | Pointer to [**DeterministicComm**](DeterministicComm.md) |  | [optional] 
**SurvivalTime** | Pointer to **float32** |  | [optional] 
**NssaaSupport** | Pointer to [**NSSAASupport**](NSSAASupport.md) |  | [optional] 
**N6Protection** | Pointer to [**N6Protection**](N6Protection.md) |  | [optional] 

## Methods

### NewCNSliceSubnetProfile

`func NewCNSliceSubnetProfile() *CNSliceSubnetProfile`

NewCNSliceSubnetProfile instantiates a new CNSliceSubnetProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCNSliceSubnetProfileWithDefaults

`func NewCNSliceSubnetProfileWithDefaults() *CNSliceSubnetProfile`

NewCNSliceSubnetProfileWithDefaults instantiates a new CNSliceSubnetProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxNumberofUEs

`func (o *CNSliceSubnetProfile) GetMaxNumberofUEs() int32`

GetMaxNumberofUEs returns the MaxNumberofUEs field if non-nil, zero value otherwise.

### GetMaxNumberofUEsOk

`func (o *CNSliceSubnetProfile) GetMaxNumberofUEsOk() (*int32, bool)`

GetMaxNumberofUEsOk returns a tuple with the MaxNumberofUEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofUEs

`func (o *CNSliceSubnetProfile) SetMaxNumberofUEs(v int32)`

SetMaxNumberofUEs sets MaxNumberofUEs field to given value.

### HasMaxNumberofUEs

`func (o *CNSliceSubnetProfile) HasMaxNumberofUEs() bool`

HasMaxNumberofUEs returns a boolean if a field has been set.

### GetDLLatency

`func (o *CNSliceSubnetProfile) GetDLLatency() float32`

GetDLLatency returns the DLLatency field if non-nil, zero value otherwise.

### GetDLLatencyOk

`func (o *CNSliceSubnetProfile) GetDLLatencyOk() (*float32, bool)`

GetDLLatencyOk returns a tuple with the DLLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLLatency

`func (o *CNSliceSubnetProfile) SetDLLatency(v float32)`

SetDLLatency sets DLLatency field to given value.

### HasDLLatency

`func (o *CNSliceSubnetProfile) HasDLLatency() bool`

HasDLLatency returns a boolean if a field has been set.

### GetULLatency

`func (o *CNSliceSubnetProfile) GetULLatency() float32`

GetULLatency returns the ULLatency field if non-nil, zero value otherwise.

### GetULLatencyOk

`func (o *CNSliceSubnetProfile) GetULLatencyOk() (*float32, bool)`

GetULLatencyOk returns a tuple with the ULLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULLatency

`func (o *CNSliceSubnetProfile) SetULLatency(v float32)`

SetULLatency sets ULLatency field to given value.

### HasULLatency

`func (o *CNSliceSubnetProfile) HasULLatency() bool`

HasULLatency returns a boolean if a field has been set.

### GetDLThptPerSliceSubnet

`func (o *CNSliceSubnetProfile) GetDLThptPerSliceSubnet() XLThpt`

GetDLThptPerSliceSubnet returns the DLThptPerSliceSubnet field if non-nil, zero value otherwise.

### GetDLThptPerSliceSubnetOk

`func (o *CNSliceSubnetProfile) GetDLThptPerSliceSubnetOk() (*XLThpt, bool)`

GetDLThptPerSliceSubnetOk returns a tuple with the DLThptPerSliceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLThptPerSliceSubnet

`func (o *CNSliceSubnetProfile) SetDLThptPerSliceSubnet(v XLThpt)`

SetDLThptPerSliceSubnet sets DLThptPerSliceSubnet field to given value.

### HasDLThptPerSliceSubnet

`func (o *CNSliceSubnetProfile) HasDLThptPerSliceSubnet() bool`

HasDLThptPerSliceSubnet returns a boolean if a field has been set.

### GetDLThptPerUE

`func (o *CNSliceSubnetProfile) GetDLThptPerUE() XLThpt`

GetDLThptPerUE returns the DLThptPerUE field if non-nil, zero value otherwise.

### GetDLThptPerUEOk

`func (o *CNSliceSubnetProfile) GetDLThptPerUEOk() (*XLThpt, bool)`

GetDLThptPerUEOk returns a tuple with the DLThptPerUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLThptPerUE

`func (o *CNSliceSubnetProfile) SetDLThptPerUE(v XLThpt)`

SetDLThptPerUE sets DLThptPerUE field to given value.

### HasDLThptPerUE

`func (o *CNSliceSubnetProfile) HasDLThptPerUE() bool`

HasDLThptPerUE returns a boolean if a field has been set.

### GetULThptPerSliceSubnet

`func (o *CNSliceSubnetProfile) GetULThptPerSliceSubnet() XLThpt`

GetULThptPerSliceSubnet returns the ULThptPerSliceSubnet field if non-nil, zero value otherwise.

### GetULThptPerSliceSubnetOk

`func (o *CNSliceSubnetProfile) GetULThptPerSliceSubnetOk() (*XLThpt, bool)`

GetULThptPerSliceSubnetOk returns a tuple with the ULThptPerSliceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULThptPerSliceSubnet

`func (o *CNSliceSubnetProfile) SetULThptPerSliceSubnet(v XLThpt)`

SetULThptPerSliceSubnet sets ULThptPerSliceSubnet field to given value.

### HasULThptPerSliceSubnet

`func (o *CNSliceSubnetProfile) HasULThptPerSliceSubnet() bool`

HasULThptPerSliceSubnet returns a boolean if a field has been set.

### GetULThptPerUE

`func (o *CNSliceSubnetProfile) GetULThptPerUE() XLThpt`

GetULThptPerUE returns the ULThptPerUE field if non-nil, zero value otherwise.

### GetULThptPerUEOk

`func (o *CNSliceSubnetProfile) GetULThptPerUEOk() (*XLThpt, bool)`

GetULThptPerUEOk returns a tuple with the ULThptPerUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULThptPerUE

`func (o *CNSliceSubnetProfile) SetULThptPerUE(v XLThpt)`

SetULThptPerUE sets ULThptPerUE field to given value.

### HasULThptPerUE

`func (o *CNSliceSubnetProfile) HasULThptPerUE() bool`

HasULThptPerUE returns a boolean if a field has been set.

### GetMaxNumberOfPDUSessions

`func (o *CNSliceSubnetProfile) GetMaxNumberOfPDUSessions() int32`

GetMaxNumberOfPDUSessions returns the MaxNumberOfPDUSessions field if non-nil, zero value otherwise.

### GetMaxNumberOfPDUSessionsOk

`func (o *CNSliceSubnetProfile) GetMaxNumberOfPDUSessionsOk() (*int32, bool)`

GetMaxNumberOfPDUSessionsOk returns a tuple with the MaxNumberOfPDUSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfPDUSessions

`func (o *CNSliceSubnetProfile) SetMaxNumberOfPDUSessions(v int32)`

SetMaxNumberOfPDUSessions sets MaxNumberOfPDUSessions field to given value.

### HasMaxNumberOfPDUSessions

`func (o *CNSliceSubnetProfile) HasMaxNumberOfPDUSessions() bool`

HasMaxNumberOfPDUSessions returns a boolean if a field has been set.

### GetCoverageAreaTAList

`func (o *CNSliceSubnetProfile) GetCoverageAreaTAList() []int32`

GetCoverageAreaTAList returns the CoverageAreaTAList field if non-nil, zero value otherwise.

### GetCoverageAreaTAListOk

`func (o *CNSliceSubnetProfile) GetCoverageAreaTAListOk() (*[]int32, bool)`

GetCoverageAreaTAListOk returns a tuple with the CoverageAreaTAList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageAreaTAList

`func (o *CNSliceSubnetProfile) SetCoverageAreaTAList(v []int32)`

SetCoverageAreaTAList sets CoverageAreaTAList field to given value.

### HasCoverageAreaTAList

`func (o *CNSliceSubnetProfile) HasCoverageAreaTAList() bool`

HasCoverageAreaTAList returns a boolean if a field has been set.

### GetResourceSharingLevel

`func (o *CNSliceSubnetProfile) GetResourceSharingLevel() SharingLevel`

GetResourceSharingLevel returns the ResourceSharingLevel field if non-nil, zero value otherwise.

### GetResourceSharingLevelOk

`func (o *CNSliceSubnetProfile) GetResourceSharingLevelOk() (*SharingLevel, bool)`

GetResourceSharingLevelOk returns a tuple with the ResourceSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSharingLevel

`func (o *CNSliceSubnetProfile) SetResourceSharingLevel(v SharingLevel)`

SetResourceSharingLevel sets ResourceSharingLevel field to given value.

### HasResourceSharingLevel

`func (o *CNSliceSubnetProfile) HasResourceSharingLevel() bool`

HasResourceSharingLevel returns a boolean if a field has been set.

### GetDLMaxPktSize

`func (o *CNSliceSubnetProfile) GetDLMaxPktSize() int32`

GetDLMaxPktSize returns the DLMaxPktSize field if non-nil, zero value otherwise.

### GetDLMaxPktSizeOk

`func (o *CNSliceSubnetProfile) GetDLMaxPktSizeOk() (*int32, bool)`

GetDLMaxPktSizeOk returns a tuple with the DLMaxPktSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLMaxPktSize

`func (o *CNSliceSubnetProfile) SetDLMaxPktSize(v int32)`

SetDLMaxPktSize sets DLMaxPktSize field to given value.

### HasDLMaxPktSize

`func (o *CNSliceSubnetProfile) HasDLMaxPktSize() bool`

HasDLMaxPktSize returns a boolean if a field has been set.

### GetULMaxPktSize

`func (o *CNSliceSubnetProfile) GetULMaxPktSize() int32`

GetULMaxPktSize returns the ULMaxPktSize field if non-nil, zero value otherwise.

### GetULMaxPktSizeOk

`func (o *CNSliceSubnetProfile) GetULMaxPktSizeOk() (*int32, bool)`

GetULMaxPktSizeOk returns a tuple with the ULMaxPktSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULMaxPktSize

`func (o *CNSliceSubnetProfile) SetULMaxPktSize(v int32)`

SetULMaxPktSize sets ULMaxPktSize field to given value.

### HasULMaxPktSize

`func (o *CNSliceSubnetProfile) HasULMaxPktSize() bool`

HasULMaxPktSize returns a boolean if a field has been set.

### GetDelayTolerance

`func (o *CNSliceSubnetProfile) GetDelayTolerance() DelayTolerance`

GetDelayTolerance returns the DelayTolerance field if non-nil, zero value otherwise.

### GetDelayToleranceOk

`func (o *CNSliceSubnetProfile) GetDelayToleranceOk() (*DelayTolerance, bool)`

GetDelayToleranceOk returns a tuple with the DelayTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayTolerance

`func (o *CNSliceSubnetProfile) SetDelayTolerance(v DelayTolerance)`

SetDelayTolerance sets DelayTolerance field to given value.

### HasDelayTolerance

`func (o *CNSliceSubnetProfile) HasDelayTolerance() bool`

HasDelayTolerance returns a boolean if a field has been set.

### GetSynchronicity

`func (o *CNSliceSubnetProfile) GetSynchronicity() SynchronicityRANSubnet`

GetSynchronicity returns the Synchronicity field if non-nil, zero value otherwise.

### GetSynchronicityOk

`func (o *CNSliceSubnetProfile) GetSynchronicityOk() (*SynchronicityRANSubnet, bool)`

GetSynchronicityOk returns a tuple with the Synchronicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronicity

`func (o *CNSliceSubnetProfile) SetSynchronicity(v SynchronicityRANSubnet)`

SetSynchronicity sets Synchronicity field to given value.

### HasSynchronicity

`func (o *CNSliceSubnetProfile) HasSynchronicity() bool`

HasSynchronicity returns a boolean if a field has been set.

### GetSliceSimultaneousUse

`func (o *CNSliceSubnetProfile) GetSliceSimultaneousUse() SliceSimultaneousUse`

GetSliceSimultaneousUse returns the SliceSimultaneousUse field if non-nil, zero value otherwise.

### GetSliceSimultaneousUseOk

`func (o *CNSliceSubnetProfile) GetSliceSimultaneousUseOk() (*SliceSimultaneousUse, bool)`

GetSliceSimultaneousUseOk returns a tuple with the SliceSimultaneousUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceSimultaneousUse

`func (o *CNSliceSubnetProfile) SetSliceSimultaneousUse(v SliceSimultaneousUse)`

SetSliceSimultaneousUse sets SliceSimultaneousUse field to given value.

### HasSliceSimultaneousUse

`func (o *CNSliceSubnetProfile) HasSliceSimultaneousUse() bool`

HasSliceSimultaneousUse returns a boolean if a field has been set.

### GetReliability

`func (o *CNSliceSubnetProfile) GetReliability() float32`

GetReliability returns the Reliability field if non-nil, zero value otherwise.

### GetReliabilityOk

`func (o *CNSliceSubnetProfile) GetReliabilityOk() (*float32, bool)`

GetReliabilityOk returns a tuple with the Reliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliability

`func (o *CNSliceSubnetProfile) SetReliability(v float32)`

SetReliability sets Reliability field to given value.

### HasReliability

`func (o *CNSliceSubnetProfile) HasReliability() bool`

HasReliability returns a boolean if a field has been set.

### GetEnergyEfficiency

`func (o *CNSliceSubnetProfile) GetEnergyEfficiency() float32`

GetEnergyEfficiency returns the EnergyEfficiency field if non-nil, zero value otherwise.

### GetEnergyEfficiencyOk

`func (o *CNSliceSubnetProfile) GetEnergyEfficiencyOk() (*float32, bool)`

GetEnergyEfficiencyOk returns a tuple with the EnergyEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyEfficiency

`func (o *CNSliceSubnetProfile) SetEnergyEfficiency(v float32)`

SetEnergyEfficiency sets EnergyEfficiency field to given value.

### HasEnergyEfficiency

`func (o *CNSliceSubnetProfile) HasEnergyEfficiency() bool`

HasEnergyEfficiency returns a boolean if a field has been set.

### GetDLDeterministicComm

`func (o *CNSliceSubnetProfile) GetDLDeterministicComm() DeterministicComm`

GetDLDeterministicComm returns the DLDeterministicComm field if non-nil, zero value otherwise.

### GetDLDeterministicCommOk

`func (o *CNSliceSubnetProfile) GetDLDeterministicCommOk() (*DeterministicComm, bool)`

GetDLDeterministicCommOk returns a tuple with the DLDeterministicComm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLDeterministicComm

`func (o *CNSliceSubnetProfile) SetDLDeterministicComm(v DeterministicComm)`

SetDLDeterministicComm sets DLDeterministicComm field to given value.

### HasDLDeterministicComm

`func (o *CNSliceSubnetProfile) HasDLDeterministicComm() bool`

HasDLDeterministicComm returns a boolean if a field has been set.

### GetULDeterministicComm

`func (o *CNSliceSubnetProfile) GetULDeterministicComm() DeterministicComm`

GetULDeterministicComm returns the ULDeterministicComm field if non-nil, zero value otherwise.

### GetULDeterministicCommOk

`func (o *CNSliceSubnetProfile) GetULDeterministicCommOk() (*DeterministicComm, bool)`

GetULDeterministicCommOk returns a tuple with the ULDeterministicComm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULDeterministicComm

`func (o *CNSliceSubnetProfile) SetULDeterministicComm(v DeterministicComm)`

SetULDeterministicComm sets ULDeterministicComm field to given value.

### HasULDeterministicComm

`func (o *CNSliceSubnetProfile) HasULDeterministicComm() bool`

HasULDeterministicComm returns a boolean if a field has been set.

### GetSurvivalTime

`func (o *CNSliceSubnetProfile) GetSurvivalTime() float32`

GetSurvivalTime returns the SurvivalTime field if non-nil, zero value otherwise.

### GetSurvivalTimeOk

`func (o *CNSliceSubnetProfile) GetSurvivalTimeOk() (*float32, bool)`

GetSurvivalTimeOk returns a tuple with the SurvivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurvivalTime

`func (o *CNSliceSubnetProfile) SetSurvivalTime(v float32)`

SetSurvivalTime sets SurvivalTime field to given value.

### HasSurvivalTime

`func (o *CNSliceSubnetProfile) HasSurvivalTime() bool`

HasSurvivalTime returns a boolean if a field has been set.

### GetNssaaSupport

`func (o *CNSliceSubnetProfile) GetNssaaSupport() NSSAASupport`

GetNssaaSupport returns the NssaaSupport field if non-nil, zero value otherwise.

### GetNssaaSupportOk

`func (o *CNSliceSubnetProfile) GetNssaaSupportOk() (*NSSAASupport, bool)`

GetNssaaSupportOk returns a tuple with the NssaaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssaaSupport

`func (o *CNSliceSubnetProfile) SetNssaaSupport(v NSSAASupport)`

SetNssaaSupport sets NssaaSupport field to given value.

### HasNssaaSupport

`func (o *CNSliceSubnetProfile) HasNssaaSupport() bool`

HasNssaaSupport returns a boolean if a field has been set.

### GetN6Protection

`func (o *CNSliceSubnetProfile) GetN6Protection() N6Protection`

GetN6Protection returns the N6Protection field if non-nil, zero value otherwise.

### GetN6ProtectionOk

`func (o *CNSliceSubnetProfile) GetN6ProtectionOk() (*N6Protection, bool)`

GetN6ProtectionOk returns a tuple with the N6Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN6Protection

`func (o *CNSliceSubnetProfile) SetN6Protection(v N6Protection)`

SetN6Protection sets N6Protection field to given value.

### HasN6Protection

`func (o *CNSliceSubnetProfile) HasN6Protection() bool`

HasN6Protection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


