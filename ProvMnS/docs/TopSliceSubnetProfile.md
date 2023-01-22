# TopSliceSubnetProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DLLatency** | Pointer to **int32** |  | [optional] 
**ULLatency** | Pointer to **int32** |  | [optional] 
**MaxNumberofUEs** | Pointer to **int32** |  | [optional] 
**DLThptPerSliceSubnet** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**DLThptPerUE** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**ULThptPerSliceSubnet** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**ULThptPerUE** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**DLMaxPktSize** | Pointer to **int32** |  | [optional] 
**ULMaxPktSize** | Pointer to **int32** |  | [optional] 
**MaxNumberOfPDUSessions** | Pointer to **int32** |  | [optional] 
**NROperatingBands** | Pointer to **string** |  | [optional] 
**SliceSimultaneousUse** | Pointer to [**SliceSimultaneousUse**](SliceSimultaneousUse.md) |  | [optional] 
**EnergyEfficiency** | Pointer to [**EnergyEfficiency**](EnergyEfficiency.md) |  | [optional] 
**Synchronicity** | Pointer to [**Synchronicity**](Synchronicity.md) |  | [optional] 
**DelayTolerance** | Pointer to [**DelayTolerance**](DelayTolerance.md) |  | [optional] 
**Positioning** | Pointer to [**Positioning**](Positioning.md) |  | [optional] 
**TermDensity** | Pointer to [**TermDensity**](TermDensity.md) |  | [optional] 
**ActivityFactor** | Pointer to **int32** |  | [optional] 
**CoverageAreaTAList** | Pointer to **[]int32** |  | [optional] 
**ResourceSharingLevel** | Pointer to [**SharingLevel**](SharingLevel.md) |  | [optional] 
**UEMobilityLevel** | Pointer to [**MobilityLevel**](MobilityLevel.md) |  | [optional] 
**UESpeed** | Pointer to **int32** |  | [optional] 
**Reliability** | Pointer to **float32** |  | [optional] 
**SST** | Pointer to **int32** |  | [optional] 
**DLDeterministicComm** | Pointer to [**DeterministicComm**](DeterministicComm.md) |  | [optional] 
**ULDeterministicComm** | Pointer to [**DeterministicComm**](DeterministicComm.md) |  | [optional] 
**SurvivalTime** | Pointer to **float32** |  | [optional] 
**NssaaSupport** | Pointer to [**NSSAASupport**](NSSAASupport.md) |  | [optional] 
**N6Protection** | Pointer to [**N6Protection**](N6Protection.md) |  | [optional] 

## Methods

### NewTopSliceSubnetProfile

`func NewTopSliceSubnetProfile() *TopSliceSubnetProfile`

NewTopSliceSubnetProfile instantiates a new TopSliceSubnetProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopSliceSubnetProfileWithDefaults

`func NewTopSliceSubnetProfileWithDefaults() *TopSliceSubnetProfile`

NewTopSliceSubnetProfileWithDefaults instantiates a new TopSliceSubnetProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDLLatency

`func (o *TopSliceSubnetProfile) GetDLLatency() int32`

GetDLLatency returns the DLLatency field if non-nil, zero value otherwise.

### GetDLLatencyOk

`func (o *TopSliceSubnetProfile) GetDLLatencyOk() (*int32, bool)`

GetDLLatencyOk returns a tuple with the DLLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLLatency

`func (o *TopSliceSubnetProfile) SetDLLatency(v int32)`

SetDLLatency sets DLLatency field to given value.

### HasDLLatency

`func (o *TopSliceSubnetProfile) HasDLLatency() bool`

HasDLLatency returns a boolean if a field has been set.

### GetULLatency

`func (o *TopSliceSubnetProfile) GetULLatency() int32`

GetULLatency returns the ULLatency field if non-nil, zero value otherwise.

### GetULLatencyOk

`func (o *TopSliceSubnetProfile) GetULLatencyOk() (*int32, bool)`

GetULLatencyOk returns a tuple with the ULLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULLatency

`func (o *TopSliceSubnetProfile) SetULLatency(v int32)`

SetULLatency sets ULLatency field to given value.

### HasULLatency

`func (o *TopSliceSubnetProfile) HasULLatency() bool`

HasULLatency returns a boolean if a field has been set.

### GetMaxNumberofUEs

`func (o *TopSliceSubnetProfile) GetMaxNumberofUEs() int32`

GetMaxNumberofUEs returns the MaxNumberofUEs field if non-nil, zero value otherwise.

### GetMaxNumberofUEsOk

`func (o *TopSliceSubnetProfile) GetMaxNumberofUEsOk() (*int32, bool)`

GetMaxNumberofUEsOk returns a tuple with the MaxNumberofUEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofUEs

`func (o *TopSliceSubnetProfile) SetMaxNumberofUEs(v int32)`

SetMaxNumberofUEs sets MaxNumberofUEs field to given value.

### HasMaxNumberofUEs

`func (o *TopSliceSubnetProfile) HasMaxNumberofUEs() bool`

HasMaxNumberofUEs returns a boolean if a field has been set.

### GetDLThptPerSliceSubnet

`func (o *TopSliceSubnetProfile) GetDLThptPerSliceSubnet() XLThpt`

GetDLThptPerSliceSubnet returns the DLThptPerSliceSubnet field if non-nil, zero value otherwise.

### GetDLThptPerSliceSubnetOk

`func (o *TopSliceSubnetProfile) GetDLThptPerSliceSubnetOk() (*XLThpt, bool)`

GetDLThptPerSliceSubnetOk returns a tuple with the DLThptPerSliceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLThptPerSliceSubnet

`func (o *TopSliceSubnetProfile) SetDLThptPerSliceSubnet(v XLThpt)`

SetDLThptPerSliceSubnet sets DLThptPerSliceSubnet field to given value.

### HasDLThptPerSliceSubnet

`func (o *TopSliceSubnetProfile) HasDLThptPerSliceSubnet() bool`

HasDLThptPerSliceSubnet returns a boolean if a field has been set.

### GetDLThptPerUE

`func (o *TopSliceSubnetProfile) GetDLThptPerUE() XLThpt`

GetDLThptPerUE returns the DLThptPerUE field if non-nil, zero value otherwise.

### GetDLThptPerUEOk

`func (o *TopSliceSubnetProfile) GetDLThptPerUEOk() (*XLThpt, bool)`

GetDLThptPerUEOk returns a tuple with the DLThptPerUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLThptPerUE

`func (o *TopSliceSubnetProfile) SetDLThptPerUE(v XLThpt)`

SetDLThptPerUE sets DLThptPerUE field to given value.

### HasDLThptPerUE

`func (o *TopSliceSubnetProfile) HasDLThptPerUE() bool`

HasDLThptPerUE returns a boolean if a field has been set.

### GetULThptPerSliceSubnet

`func (o *TopSliceSubnetProfile) GetULThptPerSliceSubnet() XLThpt`

GetULThptPerSliceSubnet returns the ULThptPerSliceSubnet field if non-nil, zero value otherwise.

### GetULThptPerSliceSubnetOk

`func (o *TopSliceSubnetProfile) GetULThptPerSliceSubnetOk() (*XLThpt, bool)`

GetULThptPerSliceSubnetOk returns a tuple with the ULThptPerSliceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULThptPerSliceSubnet

`func (o *TopSliceSubnetProfile) SetULThptPerSliceSubnet(v XLThpt)`

SetULThptPerSliceSubnet sets ULThptPerSliceSubnet field to given value.

### HasULThptPerSliceSubnet

`func (o *TopSliceSubnetProfile) HasULThptPerSliceSubnet() bool`

HasULThptPerSliceSubnet returns a boolean if a field has been set.

### GetULThptPerUE

`func (o *TopSliceSubnetProfile) GetULThptPerUE() XLThpt`

GetULThptPerUE returns the ULThptPerUE field if non-nil, zero value otherwise.

### GetULThptPerUEOk

`func (o *TopSliceSubnetProfile) GetULThptPerUEOk() (*XLThpt, bool)`

GetULThptPerUEOk returns a tuple with the ULThptPerUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULThptPerUE

`func (o *TopSliceSubnetProfile) SetULThptPerUE(v XLThpt)`

SetULThptPerUE sets ULThptPerUE field to given value.

### HasULThptPerUE

`func (o *TopSliceSubnetProfile) HasULThptPerUE() bool`

HasULThptPerUE returns a boolean if a field has been set.

### GetDLMaxPktSize

`func (o *TopSliceSubnetProfile) GetDLMaxPktSize() int32`

GetDLMaxPktSize returns the DLMaxPktSize field if non-nil, zero value otherwise.

### GetDLMaxPktSizeOk

`func (o *TopSliceSubnetProfile) GetDLMaxPktSizeOk() (*int32, bool)`

GetDLMaxPktSizeOk returns a tuple with the DLMaxPktSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLMaxPktSize

`func (o *TopSliceSubnetProfile) SetDLMaxPktSize(v int32)`

SetDLMaxPktSize sets DLMaxPktSize field to given value.

### HasDLMaxPktSize

`func (o *TopSliceSubnetProfile) HasDLMaxPktSize() bool`

HasDLMaxPktSize returns a boolean if a field has been set.

### GetULMaxPktSize

`func (o *TopSliceSubnetProfile) GetULMaxPktSize() int32`

GetULMaxPktSize returns the ULMaxPktSize field if non-nil, zero value otherwise.

### GetULMaxPktSizeOk

`func (o *TopSliceSubnetProfile) GetULMaxPktSizeOk() (*int32, bool)`

GetULMaxPktSizeOk returns a tuple with the ULMaxPktSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULMaxPktSize

`func (o *TopSliceSubnetProfile) SetULMaxPktSize(v int32)`

SetULMaxPktSize sets ULMaxPktSize field to given value.

### HasULMaxPktSize

`func (o *TopSliceSubnetProfile) HasULMaxPktSize() bool`

HasULMaxPktSize returns a boolean if a field has been set.

### GetMaxNumberOfPDUSessions

`func (o *TopSliceSubnetProfile) GetMaxNumberOfPDUSessions() int32`

GetMaxNumberOfPDUSessions returns the MaxNumberOfPDUSessions field if non-nil, zero value otherwise.

### GetMaxNumberOfPDUSessionsOk

`func (o *TopSliceSubnetProfile) GetMaxNumberOfPDUSessionsOk() (*int32, bool)`

GetMaxNumberOfPDUSessionsOk returns a tuple with the MaxNumberOfPDUSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfPDUSessions

`func (o *TopSliceSubnetProfile) SetMaxNumberOfPDUSessions(v int32)`

SetMaxNumberOfPDUSessions sets MaxNumberOfPDUSessions field to given value.

### HasMaxNumberOfPDUSessions

`func (o *TopSliceSubnetProfile) HasMaxNumberOfPDUSessions() bool`

HasMaxNumberOfPDUSessions returns a boolean if a field has been set.

### GetNROperatingBands

`func (o *TopSliceSubnetProfile) GetNROperatingBands() string`

GetNROperatingBands returns the NROperatingBands field if non-nil, zero value otherwise.

### GetNROperatingBandsOk

`func (o *TopSliceSubnetProfile) GetNROperatingBandsOk() (*string, bool)`

GetNROperatingBandsOk returns a tuple with the NROperatingBands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNROperatingBands

`func (o *TopSliceSubnetProfile) SetNROperatingBands(v string)`

SetNROperatingBands sets NROperatingBands field to given value.

### HasNROperatingBands

`func (o *TopSliceSubnetProfile) HasNROperatingBands() bool`

HasNROperatingBands returns a boolean if a field has been set.

### GetSliceSimultaneousUse

`func (o *TopSliceSubnetProfile) GetSliceSimultaneousUse() SliceSimultaneousUse`

GetSliceSimultaneousUse returns the SliceSimultaneousUse field if non-nil, zero value otherwise.

### GetSliceSimultaneousUseOk

`func (o *TopSliceSubnetProfile) GetSliceSimultaneousUseOk() (*SliceSimultaneousUse, bool)`

GetSliceSimultaneousUseOk returns a tuple with the SliceSimultaneousUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceSimultaneousUse

`func (o *TopSliceSubnetProfile) SetSliceSimultaneousUse(v SliceSimultaneousUse)`

SetSliceSimultaneousUse sets SliceSimultaneousUse field to given value.

### HasSliceSimultaneousUse

`func (o *TopSliceSubnetProfile) HasSliceSimultaneousUse() bool`

HasSliceSimultaneousUse returns a boolean if a field has been set.

### GetEnergyEfficiency

`func (o *TopSliceSubnetProfile) GetEnergyEfficiency() EnergyEfficiency`

GetEnergyEfficiency returns the EnergyEfficiency field if non-nil, zero value otherwise.

### GetEnergyEfficiencyOk

`func (o *TopSliceSubnetProfile) GetEnergyEfficiencyOk() (*EnergyEfficiency, bool)`

GetEnergyEfficiencyOk returns a tuple with the EnergyEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyEfficiency

`func (o *TopSliceSubnetProfile) SetEnergyEfficiency(v EnergyEfficiency)`

SetEnergyEfficiency sets EnergyEfficiency field to given value.

### HasEnergyEfficiency

`func (o *TopSliceSubnetProfile) HasEnergyEfficiency() bool`

HasEnergyEfficiency returns a boolean if a field has been set.

### GetSynchronicity

`func (o *TopSliceSubnetProfile) GetSynchronicity() Synchronicity`

GetSynchronicity returns the Synchronicity field if non-nil, zero value otherwise.

### GetSynchronicityOk

`func (o *TopSliceSubnetProfile) GetSynchronicityOk() (*Synchronicity, bool)`

GetSynchronicityOk returns a tuple with the Synchronicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronicity

`func (o *TopSliceSubnetProfile) SetSynchronicity(v Synchronicity)`

SetSynchronicity sets Synchronicity field to given value.

### HasSynchronicity

`func (o *TopSliceSubnetProfile) HasSynchronicity() bool`

HasSynchronicity returns a boolean if a field has been set.

### GetDelayTolerance

`func (o *TopSliceSubnetProfile) GetDelayTolerance() DelayTolerance`

GetDelayTolerance returns the DelayTolerance field if non-nil, zero value otherwise.

### GetDelayToleranceOk

`func (o *TopSliceSubnetProfile) GetDelayToleranceOk() (*DelayTolerance, bool)`

GetDelayToleranceOk returns a tuple with the DelayTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayTolerance

`func (o *TopSliceSubnetProfile) SetDelayTolerance(v DelayTolerance)`

SetDelayTolerance sets DelayTolerance field to given value.

### HasDelayTolerance

`func (o *TopSliceSubnetProfile) HasDelayTolerance() bool`

HasDelayTolerance returns a boolean if a field has been set.

### GetPositioning

`func (o *TopSliceSubnetProfile) GetPositioning() Positioning`

GetPositioning returns the Positioning field if non-nil, zero value otherwise.

### GetPositioningOk

`func (o *TopSliceSubnetProfile) GetPositioningOk() (*Positioning, bool)`

GetPositioningOk returns a tuple with the Positioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioning

`func (o *TopSliceSubnetProfile) SetPositioning(v Positioning)`

SetPositioning sets Positioning field to given value.

### HasPositioning

`func (o *TopSliceSubnetProfile) HasPositioning() bool`

HasPositioning returns a boolean if a field has been set.

### GetTermDensity

`func (o *TopSliceSubnetProfile) GetTermDensity() TermDensity`

GetTermDensity returns the TermDensity field if non-nil, zero value otherwise.

### GetTermDensityOk

`func (o *TopSliceSubnetProfile) GetTermDensityOk() (*TermDensity, bool)`

GetTermDensityOk returns a tuple with the TermDensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermDensity

`func (o *TopSliceSubnetProfile) SetTermDensity(v TermDensity)`

SetTermDensity sets TermDensity field to given value.

### HasTermDensity

`func (o *TopSliceSubnetProfile) HasTermDensity() bool`

HasTermDensity returns a boolean if a field has been set.

### GetActivityFactor

`func (o *TopSliceSubnetProfile) GetActivityFactor() int32`

GetActivityFactor returns the ActivityFactor field if non-nil, zero value otherwise.

### GetActivityFactorOk

`func (o *TopSliceSubnetProfile) GetActivityFactorOk() (*int32, bool)`

GetActivityFactorOk returns a tuple with the ActivityFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityFactor

`func (o *TopSliceSubnetProfile) SetActivityFactor(v int32)`

SetActivityFactor sets ActivityFactor field to given value.

### HasActivityFactor

`func (o *TopSliceSubnetProfile) HasActivityFactor() bool`

HasActivityFactor returns a boolean if a field has been set.

### GetCoverageAreaTAList

`func (o *TopSliceSubnetProfile) GetCoverageAreaTAList() []int32`

GetCoverageAreaTAList returns the CoverageAreaTAList field if non-nil, zero value otherwise.

### GetCoverageAreaTAListOk

`func (o *TopSliceSubnetProfile) GetCoverageAreaTAListOk() (*[]int32, bool)`

GetCoverageAreaTAListOk returns a tuple with the CoverageAreaTAList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageAreaTAList

`func (o *TopSliceSubnetProfile) SetCoverageAreaTAList(v []int32)`

SetCoverageAreaTAList sets CoverageAreaTAList field to given value.

### HasCoverageAreaTAList

`func (o *TopSliceSubnetProfile) HasCoverageAreaTAList() bool`

HasCoverageAreaTAList returns a boolean if a field has been set.

### GetResourceSharingLevel

`func (o *TopSliceSubnetProfile) GetResourceSharingLevel() SharingLevel`

GetResourceSharingLevel returns the ResourceSharingLevel field if non-nil, zero value otherwise.

### GetResourceSharingLevelOk

`func (o *TopSliceSubnetProfile) GetResourceSharingLevelOk() (*SharingLevel, bool)`

GetResourceSharingLevelOk returns a tuple with the ResourceSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSharingLevel

`func (o *TopSliceSubnetProfile) SetResourceSharingLevel(v SharingLevel)`

SetResourceSharingLevel sets ResourceSharingLevel field to given value.

### HasResourceSharingLevel

`func (o *TopSliceSubnetProfile) HasResourceSharingLevel() bool`

HasResourceSharingLevel returns a boolean if a field has been set.

### GetUEMobilityLevel

`func (o *TopSliceSubnetProfile) GetUEMobilityLevel() MobilityLevel`

GetUEMobilityLevel returns the UEMobilityLevel field if non-nil, zero value otherwise.

### GetUEMobilityLevelOk

`func (o *TopSliceSubnetProfile) GetUEMobilityLevelOk() (*MobilityLevel, bool)`

GetUEMobilityLevelOk returns a tuple with the UEMobilityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUEMobilityLevel

`func (o *TopSliceSubnetProfile) SetUEMobilityLevel(v MobilityLevel)`

SetUEMobilityLevel sets UEMobilityLevel field to given value.

### HasUEMobilityLevel

`func (o *TopSliceSubnetProfile) HasUEMobilityLevel() bool`

HasUEMobilityLevel returns a boolean if a field has been set.

### GetUESpeed

`func (o *TopSliceSubnetProfile) GetUESpeed() int32`

GetUESpeed returns the UESpeed field if non-nil, zero value otherwise.

### GetUESpeedOk

`func (o *TopSliceSubnetProfile) GetUESpeedOk() (*int32, bool)`

GetUESpeedOk returns a tuple with the UESpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUESpeed

`func (o *TopSliceSubnetProfile) SetUESpeed(v int32)`

SetUESpeed sets UESpeed field to given value.

### HasUESpeed

`func (o *TopSliceSubnetProfile) HasUESpeed() bool`

HasUESpeed returns a boolean if a field has been set.

### GetReliability

`func (o *TopSliceSubnetProfile) GetReliability() float32`

GetReliability returns the Reliability field if non-nil, zero value otherwise.

### GetReliabilityOk

`func (o *TopSliceSubnetProfile) GetReliabilityOk() (*float32, bool)`

GetReliabilityOk returns a tuple with the Reliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliability

`func (o *TopSliceSubnetProfile) SetReliability(v float32)`

SetReliability sets Reliability field to given value.

### HasReliability

`func (o *TopSliceSubnetProfile) HasReliability() bool`

HasReliability returns a boolean if a field has been set.

### GetSST

`func (o *TopSliceSubnetProfile) GetSST() int32`

GetSST returns the SST field if non-nil, zero value otherwise.

### GetSSTOk

`func (o *TopSliceSubnetProfile) GetSSTOk() (*int32, bool)`

GetSSTOk returns a tuple with the SST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSST

`func (o *TopSliceSubnetProfile) SetSST(v int32)`

SetSST sets SST field to given value.

### HasSST

`func (o *TopSliceSubnetProfile) HasSST() bool`

HasSST returns a boolean if a field has been set.

### GetDLDeterministicComm

`func (o *TopSliceSubnetProfile) GetDLDeterministicComm() DeterministicComm`

GetDLDeterministicComm returns the DLDeterministicComm field if non-nil, zero value otherwise.

### GetDLDeterministicCommOk

`func (o *TopSliceSubnetProfile) GetDLDeterministicCommOk() (*DeterministicComm, bool)`

GetDLDeterministicCommOk returns a tuple with the DLDeterministicComm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLDeterministicComm

`func (o *TopSliceSubnetProfile) SetDLDeterministicComm(v DeterministicComm)`

SetDLDeterministicComm sets DLDeterministicComm field to given value.

### HasDLDeterministicComm

`func (o *TopSliceSubnetProfile) HasDLDeterministicComm() bool`

HasDLDeterministicComm returns a boolean if a field has been set.

### GetULDeterministicComm

`func (o *TopSliceSubnetProfile) GetULDeterministicComm() DeterministicComm`

GetULDeterministicComm returns the ULDeterministicComm field if non-nil, zero value otherwise.

### GetULDeterministicCommOk

`func (o *TopSliceSubnetProfile) GetULDeterministicCommOk() (*DeterministicComm, bool)`

GetULDeterministicCommOk returns a tuple with the ULDeterministicComm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULDeterministicComm

`func (o *TopSliceSubnetProfile) SetULDeterministicComm(v DeterministicComm)`

SetULDeterministicComm sets ULDeterministicComm field to given value.

### HasULDeterministicComm

`func (o *TopSliceSubnetProfile) HasULDeterministicComm() bool`

HasULDeterministicComm returns a boolean if a field has been set.

### GetSurvivalTime

`func (o *TopSliceSubnetProfile) GetSurvivalTime() float32`

GetSurvivalTime returns the SurvivalTime field if non-nil, zero value otherwise.

### GetSurvivalTimeOk

`func (o *TopSliceSubnetProfile) GetSurvivalTimeOk() (*float32, bool)`

GetSurvivalTimeOk returns a tuple with the SurvivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurvivalTime

`func (o *TopSliceSubnetProfile) SetSurvivalTime(v float32)`

SetSurvivalTime sets SurvivalTime field to given value.

### HasSurvivalTime

`func (o *TopSliceSubnetProfile) HasSurvivalTime() bool`

HasSurvivalTime returns a boolean if a field has been set.

### GetNssaaSupport

`func (o *TopSliceSubnetProfile) GetNssaaSupport() NSSAASupport`

GetNssaaSupport returns the NssaaSupport field if non-nil, zero value otherwise.

### GetNssaaSupportOk

`func (o *TopSliceSubnetProfile) GetNssaaSupportOk() (*NSSAASupport, bool)`

GetNssaaSupportOk returns a tuple with the NssaaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssaaSupport

`func (o *TopSliceSubnetProfile) SetNssaaSupport(v NSSAASupport)`

SetNssaaSupport sets NssaaSupport field to given value.

### HasNssaaSupport

`func (o *TopSliceSubnetProfile) HasNssaaSupport() bool`

HasNssaaSupport returns a boolean if a field has been set.

### GetN6Protection

`func (o *TopSliceSubnetProfile) GetN6Protection() N6Protection`

GetN6Protection returns the N6Protection field if non-nil, zero value otherwise.

### GetN6ProtectionOk

`func (o *TopSliceSubnetProfile) GetN6ProtectionOk() (*N6Protection, bool)`

GetN6ProtectionOk returns a tuple with the N6Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN6Protection

`func (o *TopSliceSubnetProfile) SetN6Protection(v N6Protection)`

SetN6Protection sets N6Protection field to given value.

### HasN6Protection

`func (o *TopSliceSubnetProfile) HasN6Protection() bool`

HasN6Protection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


