# RANSliceSubnetProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoverageAreaTAList** | Pointer to **[]int32** |  | [optional] 
**DLLatency** | Pointer to **float32** |  | [optional] 
**ULLatency** | Pointer to **float32** |  | [optional] 
**UEMobilityLevel** | Pointer to [**MobilityLevel**](MobilityLevel.md) |  | [optional] 
**ResourceSharingLevel** | Pointer to [**SharingLevel**](SharingLevel.md) |  | [optional] 
**MaxNumberofUEs** | Pointer to **int32** |  | [optional] 
**ActivityFactor** | Pointer to **int32** |  | [optional] 
**DLThptPerSliceSubnet** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**DLThptPerUE** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**ULThptPerSliceSubnet** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**ULThptPerUE** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**UESpeed** | Pointer to **int32** |  | [optional] 
**Reliability** | Pointer to **float32** |  | [optional] 
**SST** | Pointer to **int32** |  | [optional] 
**DLMaxPktSize** | Pointer to **int32** |  | [optional] 
**ULMaxPktSize** | Pointer to **int32** |  | [optional] 
**NROperatingBands** | Pointer to **string** |  | [optional] 
**DelayTolerance** | Pointer to [**DelayTolerance**](DelayTolerance.md) |  | [optional] 
**Positioning** | Pointer to [**PositioningRANSubnet**](PositioningRANSubnet.md) |  | [optional] 
**SliceSimultaneousUse** | Pointer to [**SliceSimultaneousUse**](SliceSimultaneousUse.md) |  | [optional] 
**EnergyEfficiency** | Pointer to **float32** |  | [optional] 
**TermDensity** | Pointer to [**TermDensity**](TermDensity.md) |  | [optional] 
**SurvivalTime** | Pointer to **float32** |  | [optional] 
**Synchronicity** | Pointer to [**SynchronicityRANSubnet**](SynchronicityRANSubnet.md) |  | [optional] 
**DLDeterministicComm** | Pointer to [**DeterministicComm**](DeterministicComm.md) |  | [optional] 
**ULDeterministicComm** | Pointer to [**DeterministicComm**](DeterministicComm.md) |  | [optional] 

## Methods

### NewRANSliceSubnetProfile

`func NewRANSliceSubnetProfile() *RANSliceSubnetProfile`

NewRANSliceSubnetProfile instantiates a new RANSliceSubnetProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRANSliceSubnetProfileWithDefaults

`func NewRANSliceSubnetProfileWithDefaults() *RANSliceSubnetProfile`

NewRANSliceSubnetProfileWithDefaults instantiates a new RANSliceSubnetProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoverageAreaTAList

`func (o *RANSliceSubnetProfile) GetCoverageAreaTAList() []int32`

GetCoverageAreaTAList returns the CoverageAreaTAList field if non-nil, zero value otherwise.

### GetCoverageAreaTAListOk

`func (o *RANSliceSubnetProfile) GetCoverageAreaTAListOk() (*[]int32, bool)`

GetCoverageAreaTAListOk returns a tuple with the CoverageAreaTAList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageAreaTAList

`func (o *RANSliceSubnetProfile) SetCoverageAreaTAList(v []int32)`

SetCoverageAreaTAList sets CoverageAreaTAList field to given value.

### HasCoverageAreaTAList

`func (o *RANSliceSubnetProfile) HasCoverageAreaTAList() bool`

HasCoverageAreaTAList returns a boolean if a field has been set.

### GetDLLatency

`func (o *RANSliceSubnetProfile) GetDLLatency() float32`

GetDLLatency returns the DLLatency field if non-nil, zero value otherwise.

### GetDLLatencyOk

`func (o *RANSliceSubnetProfile) GetDLLatencyOk() (*float32, bool)`

GetDLLatencyOk returns a tuple with the DLLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLLatency

`func (o *RANSliceSubnetProfile) SetDLLatency(v float32)`

SetDLLatency sets DLLatency field to given value.

### HasDLLatency

`func (o *RANSliceSubnetProfile) HasDLLatency() bool`

HasDLLatency returns a boolean if a field has been set.

### GetULLatency

`func (o *RANSliceSubnetProfile) GetULLatency() float32`

GetULLatency returns the ULLatency field if non-nil, zero value otherwise.

### GetULLatencyOk

`func (o *RANSliceSubnetProfile) GetULLatencyOk() (*float32, bool)`

GetULLatencyOk returns a tuple with the ULLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULLatency

`func (o *RANSliceSubnetProfile) SetULLatency(v float32)`

SetULLatency sets ULLatency field to given value.

### HasULLatency

`func (o *RANSliceSubnetProfile) HasULLatency() bool`

HasULLatency returns a boolean if a field has been set.

### GetUEMobilityLevel

`func (o *RANSliceSubnetProfile) GetUEMobilityLevel() MobilityLevel`

GetUEMobilityLevel returns the UEMobilityLevel field if non-nil, zero value otherwise.

### GetUEMobilityLevelOk

`func (o *RANSliceSubnetProfile) GetUEMobilityLevelOk() (*MobilityLevel, bool)`

GetUEMobilityLevelOk returns a tuple with the UEMobilityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUEMobilityLevel

`func (o *RANSliceSubnetProfile) SetUEMobilityLevel(v MobilityLevel)`

SetUEMobilityLevel sets UEMobilityLevel field to given value.

### HasUEMobilityLevel

`func (o *RANSliceSubnetProfile) HasUEMobilityLevel() bool`

HasUEMobilityLevel returns a boolean if a field has been set.

### GetResourceSharingLevel

`func (o *RANSliceSubnetProfile) GetResourceSharingLevel() SharingLevel`

GetResourceSharingLevel returns the ResourceSharingLevel field if non-nil, zero value otherwise.

### GetResourceSharingLevelOk

`func (o *RANSliceSubnetProfile) GetResourceSharingLevelOk() (*SharingLevel, bool)`

GetResourceSharingLevelOk returns a tuple with the ResourceSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSharingLevel

`func (o *RANSliceSubnetProfile) SetResourceSharingLevel(v SharingLevel)`

SetResourceSharingLevel sets ResourceSharingLevel field to given value.

### HasResourceSharingLevel

`func (o *RANSliceSubnetProfile) HasResourceSharingLevel() bool`

HasResourceSharingLevel returns a boolean if a field has been set.

### GetMaxNumberofUEs

`func (o *RANSliceSubnetProfile) GetMaxNumberofUEs() int32`

GetMaxNumberofUEs returns the MaxNumberofUEs field if non-nil, zero value otherwise.

### GetMaxNumberofUEsOk

`func (o *RANSliceSubnetProfile) GetMaxNumberofUEsOk() (*int32, bool)`

GetMaxNumberofUEsOk returns a tuple with the MaxNumberofUEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofUEs

`func (o *RANSliceSubnetProfile) SetMaxNumberofUEs(v int32)`

SetMaxNumberofUEs sets MaxNumberofUEs field to given value.

### HasMaxNumberofUEs

`func (o *RANSliceSubnetProfile) HasMaxNumberofUEs() bool`

HasMaxNumberofUEs returns a boolean if a field has been set.

### GetActivityFactor

`func (o *RANSliceSubnetProfile) GetActivityFactor() int32`

GetActivityFactor returns the ActivityFactor field if non-nil, zero value otherwise.

### GetActivityFactorOk

`func (o *RANSliceSubnetProfile) GetActivityFactorOk() (*int32, bool)`

GetActivityFactorOk returns a tuple with the ActivityFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityFactor

`func (o *RANSliceSubnetProfile) SetActivityFactor(v int32)`

SetActivityFactor sets ActivityFactor field to given value.

### HasActivityFactor

`func (o *RANSliceSubnetProfile) HasActivityFactor() bool`

HasActivityFactor returns a boolean if a field has been set.

### GetDLThptPerSliceSubnet

`func (o *RANSliceSubnetProfile) GetDLThptPerSliceSubnet() XLThpt`

GetDLThptPerSliceSubnet returns the DLThptPerSliceSubnet field if non-nil, zero value otherwise.

### GetDLThptPerSliceSubnetOk

`func (o *RANSliceSubnetProfile) GetDLThptPerSliceSubnetOk() (*XLThpt, bool)`

GetDLThptPerSliceSubnetOk returns a tuple with the DLThptPerSliceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLThptPerSliceSubnet

`func (o *RANSliceSubnetProfile) SetDLThptPerSliceSubnet(v XLThpt)`

SetDLThptPerSliceSubnet sets DLThptPerSliceSubnet field to given value.

### HasDLThptPerSliceSubnet

`func (o *RANSliceSubnetProfile) HasDLThptPerSliceSubnet() bool`

HasDLThptPerSliceSubnet returns a boolean if a field has been set.

### GetDLThptPerUE

`func (o *RANSliceSubnetProfile) GetDLThptPerUE() XLThpt`

GetDLThptPerUE returns the DLThptPerUE field if non-nil, zero value otherwise.

### GetDLThptPerUEOk

`func (o *RANSliceSubnetProfile) GetDLThptPerUEOk() (*XLThpt, bool)`

GetDLThptPerUEOk returns a tuple with the DLThptPerUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLThptPerUE

`func (o *RANSliceSubnetProfile) SetDLThptPerUE(v XLThpt)`

SetDLThptPerUE sets DLThptPerUE field to given value.

### HasDLThptPerUE

`func (o *RANSliceSubnetProfile) HasDLThptPerUE() bool`

HasDLThptPerUE returns a boolean if a field has been set.

### GetULThptPerSliceSubnet

`func (o *RANSliceSubnetProfile) GetULThptPerSliceSubnet() XLThpt`

GetULThptPerSliceSubnet returns the ULThptPerSliceSubnet field if non-nil, zero value otherwise.

### GetULThptPerSliceSubnetOk

`func (o *RANSliceSubnetProfile) GetULThptPerSliceSubnetOk() (*XLThpt, bool)`

GetULThptPerSliceSubnetOk returns a tuple with the ULThptPerSliceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULThptPerSliceSubnet

`func (o *RANSliceSubnetProfile) SetULThptPerSliceSubnet(v XLThpt)`

SetULThptPerSliceSubnet sets ULThptPerSliceSubnet field to given value.

### HasULThptPerSliceSubnet

`func (o *RANSliceSubnetProfile) HasULThptPerSliceSubnet() bool`

HasULThptPerSliceSubnet returns a boolean if a field has been set.

### GetULThptPerUE

`func (o *RANSliceSubnetProfile) GetULThptPerUE() XLThpt`

GetULThptPerUE returns the ULThptPerUE field if non-nil, zero value otherwise.

### GetULThptPerUEOk

`func (o *RANSliceSubnetProfile) GetULThptPerUEOk() (*XLThpt, bool)`

GetULThptPerUEOk returns a tuple with the ULThptPerUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULThptPerUE

`func (o *RANSliceSubnetProfile) SetULThptPerUE(v XLThpt)`

SetULThptPerUE sets ULThptPerUE field to given value.

### HasULThptPerUE

`func (o *RANSliceSubnetProfile) HasULThptPerUE() bool`

HasULThptPerUE returns a boolean if a field has been set.

### GetUESpeed

`func (o *RANSliceSubnetProfile) GetUESpeed() int32`

GetUESpeed returns the UESpeed field if non-nil, zero value otherwise.

### GetUESpeedOk

`func (o *RANSliceSubnetProfile) GetUESpeedOk() (*int32, bool)`

GetUESpeedOk returns a tuple with the UESpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUESpeed

`func (o *RANSliceSubnetProfile) SetUESpeed(v int32)`

SetUESpeed sets UESpeed field to given value.

### HasUESpeed

`func (o *RANSliceSubnetProfile) HasUESpeed() bool`

HasUESpeed returns a boolean if a field has been set.

### GetReliability

`func (o *RANSliceSubnetProfile) GetReliability() float32`

GetReliability returns the Reliability field if non-nil, zero value otherwise.

### GetReliabilityOk

`func (o *RANSliceSubnetProfile) GetReliabilityOk() (*float32, bool)`

GetReliabilityOk returns a tuple with the Reliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliability

`func (o *RANSliceSubnetProfile) SetReliability(v float32)`

SetReliability sets Reliability field to given value.

### HasReliability

`func (o *RANSliceSubnetProfile) HasReliability() bool`

HasReliability returns a boolean if a field has been set.

### GetSST

`func (o *RANSliceSubnetProfile) GetSST() int32`

GetSST returns the SST field if non-nil, zero value otherwise.

### GetSSTOk

`func (o *RANSliceSubnetProfile) GetSSTOk() (*int32, bool)`

GetSSTOk returns a tuple with the SST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSST

`func (o *RANSliceSubnetProfile) SetSST(v int32)`

SetSST sets SST field to given value.

### HasSST

`func (o *RANSliceSubnetProfile) HasSST() bool`

HasSST returns a boolean if a field has been set.

### GetDLMaxPktSize

`func (o *RANSliceSubnetProfile) GetDLMaxPktSize() int32`

GetDLMaxPktSize returns the DLMaxPktSize field if non-nil, zero value otherwise.

### GetDLMaxPktSizeOk

`func (o *RANSliceSubnetProfile) GetDLMaxPktSizeOk() (*int32, bool)`

GetDLMaxPktSizeOk returns a tuple with the DLMaxPktSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLMaxPktSize

`func (o *RANSliceSubnetProfile) SetDLMaxPktSize(v int32)`

SetDLMaxPktSize sets DLMaxPktSize field to given value.

### HasDLMaxPktSize

`func (o *RANSliceSubnetProfile) HasDLMaxPktSize() bool`

HasDLMaxPktSize returns a boolean if a field has been set.

### GetULMaxPktSize

`func (o *RANSliceSubnetProfile) GetULMaxPktSize() int32`

GetULMaxPktSize returns the ULMaxPktSize field if non-nil, zero value otherwise.

### GetULMaxPktSizeOk

`func (o *RANSliceSubnetProfile) GetULMaxPktSizeOk() (*int32, bool)`

GetULMaxPktSizeOk returns a tuple with the ULMaxPktSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULMaxPktSize

`func (o *RANSliceSubnetProfile) SetULMaxPktSize(v int32)`

SetULMaxPktSize sets ULMaxPktSize field to given value.

### HasULMaxPktSize

`func (o *RANSliceSubnetProfile) HasULMaxPktSize() bool`

HasULMaxPktSize returns a boolean if a field has been set.

### GetNROperatingBands

`func (o *RANSliceSubnetProfile) GetNROperatingBands() string`

GetNROperatingBands returns the NROperatingBands field if non-nil, zero value otherwise.

### GetNROperatingBandsOk

`func (o *RANSliceSubnetProfile) GetNROperatingBandsOk() (*string, bool)`

GetNROperatingBandsOk returns a tuple with the NROperatingBands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNROperatingBands

`func (o *RANSliceSubnetProfile) SetNROperatingBands(v string)`

SetNROperatingBands sets NROperatingBands field to given value.

### HasNROperatingBands

`func (o *RANSliceSubnetProfile) HasNROperatingBands() bool`

HasNROperatingBands returns a boolean if a field has been set.

### GetDelayTolerance

`func (o *RANSliceSubnetProfile) GetDelayTolerance() DelayTolerance`

GetDelayTolerance returns the DelayTolerance field if non-nil, zero value otherwise.

### GetDelayToleranceOk

`func (o *RANSliceSubnetProfile) GetDelayToleranceOk() (*DelayTolerance, bool)`

GetDelayToleranceOk returns a tuple with the DelayTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayTolerance

`func (o *RANSliceSubnetProfile) SetDelayTolerance(v DelayTolerance)`

SetDelayTolerance sets DelayTolerance field to given value.

### HasDelayTolerance

`func (o *RANSliceSubnetProfile) HasDelayTolerance() bool`

HasDelayTolerance returns a boolean if a field has been set.

### GetPositioning

`func (o *RANSliceSubnetProfile) GetPositioning() PositioningRANSubnet`

GetPositioning returns the Positioning field if non-nil, zero value otherwise.

### GetPositioningOk

`func (o *RANSliceSubnetProfile) GetPositioningOk() (*PositioningRANSubnet, bool)`

GetPositioningOk returns a tuple with the Positioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioning

`func (o *RANSliceSubnetProfile) SetPositioning(v PositioningRANSubnet)`

SetPositioning sets Positioning field to given value.

### HasPositioning

`func (o *RANSliceSubnetProfile) HasPositioning() bool`

HasPositioning returns a boolean if a field has been set.

### GetSliceSimultaneousUse

`func (o *RANSliceSubnetProfile) GetSliceSimultaneousUse() SliceSimultaneousUse`

GetSliceSimultaneousUse returns the SliceSimultaneousUse field if non-nil, zero value otherwise.

### GetSliceSimultaneousUseOk

`func (o *RANSliceSubnetProfile) GetSliceSimultaneousUseOk() (*SliceSimultaneousUse, bool)`

GetSliceSimultaneousUseOk returns a tuple with the SliceSimultaneousUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceSimultaneousUse

`func (o *RANSliceSubnetProfile) SetSliceSimultaneousUse(v SliceSimultaneousUse)`

SetSliceSimultaneousUse sets SliceSimultaneousUse field to given value.

### HasSliceSimultaneousUse

`func (o *RANSliceSubnetProfile) HasSliceSimultaneousUse() bool`

HasSliceSimultaneousUse returns a boolean if a field has been set.

### GetEnergyEfficiency

`func (o *RANSliceSubnetProfile) GetEnergyEfficiency() float32`

GetEnergyEfficiency returns the EnergyEfficiency field if non-nil, zero value otherwise.

### GetEnergyEfficiencyOk

`func (o *RANSliceSubnetProfile) GetEnergyEfficiencyOk() (*float32, bool)`

GetEnergyEfficiencyOk returns a tuple with the EnergyEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyEfficiency

`func (o *RANSliceSubnetProfile) SetEnergyEfficiency(v float32)`

SetEnergyEfficiency sets EnergyEfficiency field to given value.

### HasEnergyEfficiency

`func (o *RANSliceSubnetProfile) HasEnergyEfficiency() bool`

HasEnergyEfficiency returns a boolean if a field has been set.

### GetTermDensity

`func (o *RANSliceSubnetProfile) GetTermDensity() TermDensity`

GetTermDensity returns the TermDensity field if non-nil, zero value otherwise.

### GetTermDensityOk

`func (o *RANSliceSubnetProfile) GetTermDensityOk() (*TermDensity, bool)`

GetTermDensityOk returns a tuple with the TermDensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermDensity

`func (o *RANSliceSubnetProfile) SetTermDensity(v TermDensity)`

SetTermDensity sets TermDensity field to given value.

### HasTermDensity

`func (o *RANSliceSubnetProfile) HasTermDensity() bool`

HasTermDensity returns a boolean if a field has been set.

### GetSurvivalTime

`func (o *RANSliceSubnetProfile) GetSurvivalTime() float32`

GetSurvivalTime returns the SurvivalTime field if non-nil, zero value otherwise.

### GetSurvivalTimeOk

`func (o *RANSliceSubnetProfile) GetSurvivalTimeOk() (*float32, bool)`

GetSurvivalTimeOk returns a tuple with the SurvivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurvivalTime

`func (o *RANSliceSubnetProfile) SetSurvivalTime(v float32)`

SetSurvivalTime sets SurvivalTime field to given value.

### HasSurvivalTime

`func (o *RANSliceSubnetProfile) HasSurvivalTime() bool`

HasSurvivalTime returns a boolean if a field has been set.

### GetSynchronicity

`func (o *RANSliceSubnetProfile) GetSynchronicity() SynchronicityRANSubnet`

GetSynchronicity returns the Synchronicity field if non-nil, zero value otherwise.

### GetSynchronicityOk

`func (o *RANSliceSubnetProfile) GetSynchronicityOk() (*SynchronicityRANSubnet, bool)`

GetSynchronicityOk returns a tuple with the Synchronicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronicity

`func (o *RANSliceSubnetProfile) SetSynchronicity(v SynchronicityRANSubnet)`

SetSynchronicity sets Synchronicity field to given value.

### HasSynchronicity

`func (o *RANSliceSubnetProfile) HasSynchronicity() bool`

HasSynchronicity returns a boolean if a field has been set.

### GetDLDeterministicComm

`func (o *RANSliceSubnetProfile) GetDLDeterministicComm() DeterministicComm`

GetDLDeterministicComm returns the DLDeterministicComm field if non-nil, zero value otherwise.

### GetDLDeterministicCommOk

`func (o *RANSliceSubnetProfile) GetDLDeterministicCommOk() (*DeterministicComm, bool)`

GetDLDeterministicCommOk returns a tuple with the DLDeterministicComm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLDeterministicComm

`func (o *RANSliceSubnetProfile) SetDLDeterministicComm(v DeterministicComm)`

SetDLDeterministicComm sets DLDeterministicComm field to given value.

### HasDLDeterministicComm

`func (o *RANSliceSubnetProfile) HasDLDeterministicComm() bool`

HasDLDeterministicComm returns a boolean if a field has been set.

### GetULDeterministicComm

`func (o *RANSliceSubnetProfile) GetULDeterministicComm() DeterministicComm`

GetULDeterministicComm returns the ULDeterministicComm field if non-nil, zero value otherwise.

### GetULDeterministicCommOk

`func (o *RANSliceSubnetProfile) GetULDeterministicCommOk() (*DeterministicComm, bool)`

GetULDeterministicCommOk returns a tuple with the ULDeterministicComm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULDeterministicComm

`func (o *RANSliceSubnetProfile) SetULDeterministicComm(v DeterministicComm)`

SetULDeterministicComm sets ULDeterministicComm field to given value.

### HasULDeterministicComm

`func (o *RANSliceSubnetProfile) HasULDeterministicComm() bool`

HasULDeterministicComm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


