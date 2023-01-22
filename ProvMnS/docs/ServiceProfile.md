# ServiceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceProfileId** | Pointer to **string** |  | [optional] 
**PlmnInfoList** | Pointer to [**[]PlmnInfo**](PlmnInfo.md) |  | [optional] 
**MaxNumberofUEs** | Pointer to **float32** |  | [optional] 
**DLLatency** | Pointer to **float32** |  | [optional] 
**ULLatency** | Pointer to **float32** |  | [optional] 
**UEMobilityLevel** | Pointer to [**MobilityLevel**](MobilityLevel.md) |  | [optional] 
**Sst** | Pointer to **int32** |  | [optional] 
**NetworkSliceSharingIndicator** | Pointer to [**NetworkSliceSharingIndicator**](NetworkSliceSharingIndicator.md) |  | [optional] 
**Availability** | Pointer to **float32** |  | [optional] 
**DelayTolerance** | Pointer to [**DelayTolerance**](DelayTolerance.md) |  | [optional] 
**DLDeterministicComm** | Pointer to [**DeterministicComm**](DeterministicComm.md) |  | [optional] 
**ULDeterministicComm** | Pointer to [**DeterministicComm**](DeterministicComm.md) |  | [optional] 
**DLThptPerSlice** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**DLThptPerUE** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**ULThptPerSlice** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**ULThptPerUE** | Pointer to [**XLThpt**](XLThpt.md) |  | [optional] 
**DLMaxPktSize** | Pointer to [**MaxPktSize**](MaxPktSize.md) |  | [optional] 
**ULMaxPktSize** | Pointer to [**MaxPktSize**](MaxPktSize.md) |  | [optional] 
**MaxNumberofPDUSessions** | Pointer to [**MaxNumberofPDUSessions**](MaxNumberofPDUSessions.md) |  | [optional] 
**KPIMonitoring** | Pointer to [**KPIMonitoring**](KPIMonitoring.md) |  | [optional] 
**NBIoT** | Pointer to [**NBIoT**](NBIoT.md) |  | [optional] 
**RadioSpectrum** | Pointer to [**RadioSpectrum**](RadioSpectrum.md) |  | [optional] 
**Synchronicity** | Pointer to [**Synchronicity**](Synchronicity.md) |  | [optional] 
**Positioning** | Pointer to [**Positioning**](Positioning.md) |  | [optional] 
**UserMgmtOpen** | Pointer to [**UserMgmtOpen**](UserMgmtOpen.md) |  | [optional] 
**V2XModels** | Pointer to [**V2XCommModels**](V2XCommModels.md) |  | [optional] 
**CoverageArea** | Pointer to **string** |  | [optional] 
**TermDensity** | Pointer to [**TermDensity**](TermDensity.md) |  | [optional] 
**ActivityFactor** | Pointer to **float32** |  | [optional] 
**UESpeed** | Pointer to **int32** |  | [optional] 
**Jitter** | Pointer to **int32** |  | [optional] 
**SurvivalTime** | Pointer to **float32** |  | [optional] 
**Reliability** | Pointer to **float32** |  | [optional] 
**MaxDLDataVolume** | Pointer to **float32** |  | [optional] 
**MaxULDataVolume** | Pointer to **float32** |  | [optional] 
**SliceSimultaneousUse** | Pointer to [**SliceSimultaneousUse**](SliceSimultaneousUse.md) |  | [optional] 
**EnergyEfficiency** | Pointer to [**EnergyEfficiency**](EnergyEfficiency.md) |  | [optional] 
**NssaaSupport** | Pointer to [**NSSAASupport**](NSSAASupport.md) |  | [optional] 
**N6Protection** | Pointer to [**N6Protection**](N6Protection.md) |  | [optional] 

## Methods

### NewServiceProfile

`func NewServiceProfile() *ServiceProfile`

NewServiceProfile instantiates a new ServiceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileWithDefaults

`func NewServiceProfileWithDefaults() *ServiceProfile`

NewServiceProfileWithDefaults instantiates a new ServiceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceProfileId

`func (o *ServiceProfile) GetServiceProfileId() string`

GetServiceProfileId returns the ServiceProfileId field if non-nil, zero value otherwise.

### GetServiceProfileIdOk

`func (o *ServiceProfile) GetServiceProfileIdOk() (*string, bool)`

GetServiceProfileIdOk returns a tuple with the ServiceProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProfileId

`func (o *ServiceProfile) SetServiceProfileId(v string)`

SetServiceProfileId sets ServiceProfileId field to given value.

### HasServiceProfileId

`func (o *ServiceProfile) HasServiceProfileId() bool`

HasServiceProfileId returns a boolean if a field has been set.

### GetPlmnInfoList

`func (o *ServiceProfile) GetPlmnInfoList() []PlmnInfo`

GetPlmnInfoList returns the PlmnInfoList field if non-nil, zero value otherwise.

### GetPlmnInfoListOk

`func (o *ServiceProfile) GetPlmnInfoListOk() (*[]PlmnInfo, bool)`

GetPlmnInfoListOk returns a tuple with the PlmnInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnInfoList

`func (o *ServiceProfile) SetPlmnInfoList(v []PlmnInfo)`

SetPlmnInfoList sets PlmnInfoList field to given value.

### HasPlmnInfoList

`func (o *ServiceProfile) HasPlmnInfoList() bool`

HasPlmnInfoList returns a boolean if a field has been set.

### GetMaxNumberofUEs

`func (o *ServiceProfile) GetMaxNumberofUEs() float32`

GetMaxNumberofUEs returns the MaxNumberofUEs field if non-nil, zero value otherwise.

### GetMaxNumberofUEsOk

`func (o *ServiceProfile) GetMaxNumberofUEsOk() (*float32, bool)`

GetMaxNumberofUEsOk returns a tuple with the MaxNumberofUEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofUEs

`func (o *ServiceProfile) SetMaxNumberofUEs(v float32)`

SetMaxNumberofUEs sets MaxNumberofUEs field to given value.

### HasMaxNumberofUEs

`func (o *ServiceProfile) HasMaxNumberofUEs() bool`

HasMaxNumberofUEs returns a boolean if a field has been set.

### GetDLLatency

`func (o *ServiceProfile) GetDLLatency() float32`

GetDLLatency returns the DLLatency field if non-nil, zero value otherwise.

### GetDLLatencyOk

`func (o *ServiceProfile) GetDLLatencyOk() (*float32, bool)`

GetDLLatencyOk returns a tuple with the DLLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLLatency

`func (o *ServiceProfile) SetDLLatency(v float32)`

SetDLLatency sets DLLatency field to given value.

### HasDLLatency

`func (o *ServiceProfile) HasDLLatency() bool`

HasDLLatency returns a boolean if a field has been set.

### GetULLatency

`func (o *ServiceProfile) GetULLatency() float32`

GetULLatency returns the ULLatency field if non-nil, zero value otherwise.

### GetULLatencyOk

`func (o *ServiceProfile) GetULLatencyOk() (*float32, bool)`

GetULLatencyOk returns a tuple with the ULLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULLatency

`func (o *ServiceProfile) SetULLatency(v float32)`

SetULLatency sets ULLatency field to given value.

### HasULLatency

`func (o *ServiceProfile) HasULLatency() bool`

HasULLatency returns a boolean if a field has been set.

### GetUEMobilityLevel

`func (o *ServiceProfile) GetUEMobilityLevel() MobilityLevel`

GetUEMobilityLevel returns the UEMobilityLevel field if non-nil, zero value otherwise.

### GetUEMobilityLevelOk

`func (o *ServiceProfile) GetUEMobilityLevelOk() (*MobilityLevel, bool)`

GetUEMobilityLevelOk returns a tuple with the UEMobilityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUEMobilityLevel

`func (o *ServiceProfile) SetUEMobilityLevel(v MobilityLevel)`

SetUEMobilityLevel sets UEMobilityLevel field to given value.

### HasUEMobilityLevel

`func (o *ServiceProfile) HasUEMobilityLevel() bool`

HasUEMobilityLevel returns a boolean if a field has been set.

### GetSst

`func (o *ServiceProfile) GetSst() int32`

GetSst returns the Sst field if non-nil, zero value otherwise.

### GetSstOk

`func (o *ServiceProfile) GetSstOk() (*int32, bool)`

GetSstOk returns a tuple with the Sst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSst

`func (o *ServiceProfile) SetSst(v int32)`

SetSst sets Sst field to given value.

### HasSst

`func (o *ServiceProfile) HasSst() bool`

HasSst returns a boolean if a field has been set.

### GetNetworkSliceSharingIndicator

`func (o *ServiceProfile) GetNetworkSliceSharingIndicator() NetworkSliceSharingIndicator`

GetNetworkSliceSharingIndicator returns the NetworkSliceSharingIndicator field if non-nil, zero value otherwise.

### GetNetworkSliceSharingIndicatorOk

`func (o *ServiceProfile) GetNetworkSliceSharingIndicatorOk() (*NetworkSliceSharingIndicator, bool)`

GetNetworkSliceSharingIndicatorOk returns a tuple with the NetworkSliceSharingIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSliceSharingIndicator

`func (o *ServiceProfile) SetNetworkSliceSharingIndicator(v NetworkSliceSharingIndicator)`

SetNetworkSliceSharingIndicator sets NetworkSliceSharingIndicator field to given value.

### HasNetworkSliceSharingIndicator

`func (o *ServiceProfile) HasNetworkSliceSharingIndicator() bool`

HasNetworkSliceSharingIndicator returns a boolean if a field has been set.

### GetAvailability

`func (o *ServiceProfile) GetAvailability() float32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ServiceProfile) GetAvailabilityOk() (*float32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ServiceProfile) SetAvailability(v float32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ServiceProfile) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetDelayTolerance

`func (o *ServiceProfile) GetDelayTolerance() DelayTolerance`

GetDelayTolerance returns the DelayTolerance field if non-nil, zero value otherwise.

### GetDelayToleranceOk

`func (o *ServiceProfile) GetDelayToleranceOk() (*DelayTolerance, bool)`

GetDelayToleranceOk returns a tuple with the DelayTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayTolerance

`func (o *ServiceProfile) SetDelayTolerance(v DelayTolerance)`

SetDelayTolerance sets DelayTolerance field to given value.

### HasDelayTolerance

`func (o *ServiceProfile) HasDelayTolerance() bool`

HasDelayTolerance returns a boolean if a field has been set.

### GetDLDeterministicComm

`func (o *ServiceProfile) GetDLDeterministicComm() DeterministicComm`

GetDLDeterministicComm returns the DLDeterministicComm field if non-nil, zero value otherwise.

### GetDLDeterministicCommOk

`func (o *ServiceProfile) GetDLDeterministicCommOk() (*DeterministicComm, bool)`

GetDLDeterministicCommOk returns a tuple with the DLDeterministicComm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLDeterministicComm

`func (o *ServiceProfile) SetDLDeterministicComm(v DeterministicComm)`

SetDLDeterministicComm sets DLDeterministicComm field to given value.

### HasDLDeterministicComm

`func (o *ServiceProfile) HasDLDeterministicComm() bool`

HasDLDeterministicComm returns a boolean if a field has been set.

### GetULDeterministicComm

`func (o *ServiceProfile) GetULDeterministicComm() DeterministicComm`

GetULDeterministicComm returns the ULDeterministicComm field if non-nil, zero value otherwise.

### GetULDeterministicCommOk

`func (o *ServiceProfile) GetULDeterministicCommOk() (*DeterministicComm, bool)`

GetULDeterministicCommOk returns a tuple with the ULDeterministicComm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULDeterministicComm

`func (o *ServiceProfile) SetULDeterministicComm(v DeterministicComm)`

SetULDeterministicComm sets ULDeterministicComm field to given value.

### HasULDeterministicComm

`func (o *ServiceProfile) HasULDeterministicComm() bool`

HasULDeterministicComm returns a boolean if a field has been set.

### GetDLThptPerSlice

`func (o *ServiceProfile) GetDLThptPerSlice() XLThpt`

GetDLThptPerSlice returns the DLThptPerSlice field if non-nil, zero value otherwise.

### GetDLThptPerSliceOk

`func (o *ServiceProfile) GetDLThptPerSliceOk() (*XLThpt, bool)`

GetDLThptPerSliceOk returns a tuple with the DLThptPerSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLThptPerSlice

`func (o *ServiceProfile) SetDLThptPerSlice(v XLThpt)`

SetDLThptPerSlice sets DLThptPerSlice field to given value.

### HasDLThptPerSlice

`func (o *ServiceProfile) HasDLThptPerSlice() bool`

HasDLThptPerSlice returns a boolean if a field has been set.

### GetDLThptPerUE

`func (o *ServiceProfile) GetDLThptPerUE() XLThpt`

GetDLThptPerUE returns the DLThptPerUE field if non-nil, zero value otherwise.

### GetDLThptPerUEOk

`func (o *ServiceProfile) GetDLThptPerUEOk() (*XLThpt, bool)`

GetDLThptPerUEOk returns a tuple with the DLThptPerUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLThptPerUE

`func (o *ServiceProfile) SetDLThptPerUE(v XLThpt)`

SetDLThptPerUE sets DLThptPerUE field to given value.

### HasDLThptPerUE

`func (o *ServiceProfile) HasDLThptPerUE() bool`

HasDLThptPerUE returns a boolean if a field has been set.

### GetULThptPerSlice

`func (o *ServiceProfile) GetULThptPerSlice() XLThpt`

GetULThptPerSlice returns the ULThptPerSlice field if non-nil, zero value otherwise.

### GetULThptPerSliceOk

`func (o *ServiceProfile) GetULThptPerSliceOk() (*XLThpt, bool)`

GetULThptPerSliceOk returns a tuple with the ULThptPerSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULThptPerSlice

`func (o *ServiceProfile) SetULThptPerSlice(v XLThpt)`

SetULThptPerSlice sets ULThptPerSlice field to given value.

### HasULThptPerSlice

`func (o *ServiceProfile) HasULThptPerSlice() bool`

HasULThptPerSlice returns a boolean if a field has been set.

### GetULThptPerUE

`func (o *ServiceProfile) GetULThptPerUE() XLThpt`

GetULThptPerUE returns the ULThptPerUE field if non-nil, zero value otherwise.

### GetULThptPerUEOk

`func (o *ServiceProfile) GetULThptPerUEOk() (*XLThpt, bool)`

GetULThptPerUEOk returns a tuple with the ULThptPerUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULThptPerUE

`func (o *ServiceProfile) SetULThptPerUE(v XLThpt)`

SetULThptPerUE sets ULThptPerUE field to given value.

### HasULThptPerUE

`func (o *ServiceProfile) HasULThptPerUE() bool`

HasULThptPerUE returns a boolean if a field has been set.

### GetDLMaxPktSize

`func (o *ServiceProfile) GetDLMaxPktSize() MaxPktSize`

GetDLMaxPktSize returns the DLMaxPktSize field if non-nil, zero value otherwise.

### GetDLMaxPktSizeOk

`func (o *ServiceProfile) GetDLMaxPktSizeOk() (*MaxPktSize, bool)`

GetDLMaxPktSizeOk returns a tuple with the DLMaxPktSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLMaxPktSize

`func (o *ServiceProfile) SetDLMaxPktSize(v MaxPktSize)`

SetDLMaxPktSize sets DLMaxPktSize field to given value.

### HasDLMaxPktSize

`func (o *ServiceProfile) HasDLMaxPktSize() bool`

HasDLMaxPktSize returns a boolean if a field has been set.

### GetULMaxPktSize

`func (o *ServiceProfile) GetULMaxPktSize() MaxPktSize`

GetULMaxPktSize returns the ULMaxPktSize field if non-nil, zero value otherwise.

### GetULMaxPktSizeOk

`func (o *ServiceProfile) GetULMaxPktSizeOk() (*MaxPktSize, bool)`

GetULMaxPktSizeOk returns a tuple with the ULMaxPktSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetULMaxPktSize

`func (o *ServiceProfile) SetULMaxPktSize(v MaxPktSize)`

SetULMaxPktSize sets ULMaxPktSize field to given value.

### HasULMaxPktSize

`func (o *ServiceProfile) HasULMaxPktSize() bool`

HasULMaxPktSize returns a boolean if a field has been set.

### GetMaxNumberofPDUSessions

`func (o *ServiceProfile) GetMaxNumberofPDUSessions() MaxNumberofPDUSessions`

GetMaxNumberofPDUSessions returns the MaxNumberofPDUSessions field if non-nil, zero value otherwise.

### GetMaxNumberofPDUSessionsOk

`func (o *ServiceProfile) GetMaxNumberofPDUSessionsOk() (*MaxNumberofPDUSessions, bool)`

GetMaxNumberofPDUSessionsOk returns a tuple with the MaxNumberofPDUSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberofPDUSessions

`func (o *ServiceProfile) SetMaxNumberofPDUSessions(v MaxNumberofPDUSessions)`

SetMaxNumberofPDUSessions sets MaxNumberofPDUSessions field to given value.

### HasMaxNumberofPDUSessions

`func (o *ServiceProfile) HasMaxNumberofPDUSessions() bool`

HasMaxNumberofPDUSessions returns a boolean if a field has been set.

### GetKPIMonitoring

`func (o *ServiceProfile) GetKPIMonitoring() KPIMonitoring`

GetKPIMonitoring returns the KPIMonitoring field if non-nil, zero value otherwise.

### GetKPIMonitoringOk

`func (o *ServiceProfile) GetKPIMonitoringOk() (*KPIMonitoring, bool)`

GetKPIMonitoringOk returns a tuple with the KPIMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKPIMonitoring

`func (o *ServiceProfile) SetKPIMonitoring(v KPIMonitoring)`

SetKPIMonitoring sets KPIMonitoring field to given value.

### HasKPIMonitoring

`func (o *ServiceProfile) HasKPIMonitoring() bool`

HasKPIMonitoring returns a boolean if a field has been set.

### GetNBIoT

`func (o *ServiceProfile) GetNBIoT() NBIoT`

GetNBIoT returns the NBIoT field if non-nil, zero value otherwise.

### GetNBIoTOk

`func (o *ServiceProfile) GetNBIoTOk() (*NBIoT, bool)`

GetNBIoTOk returns a tuple with the NBIoT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNBIoT

`func (o *ServiceProfile) SetNBIoT(v NBIoT)`

SetNBIoT sets NBIoT field to given value.

### HasNBIoT

`func (o *ServiceProfile) HasNBIoT() bool`

HasNBIoT returns a boolean if a field has been set.

### GetRadioSpectrum

`func (o *ServiceProfile) GetRadioSpectrum() RadioSpectrum`

GetRadioSpectrum returns the RadioSpectrum field if non-nil, zero value otherwise.

### GetRadioSpectrumOk

`func (o *ServiceProfile) GetRadioSpectrumOk() (*RadioSpectrum, bool)`

GetRadioSpectrumOk returns a tuple with the RadioSpectrum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioSpectrum

`func (o *ServiceProfile) SetRadioSpectrum(v RadioSpectrum)`

SetRadioSpectrum sets RadioSpectrum field to given value.

### HasRadioSpectrum

`func (o *ServiceProfile) HasRadioSpectrum() bool`

HasRadioSpectrum returns a boolean if a field has been set.

### GetSynchronicity

`func (o *ServiceProfile) GetSynchronicity() Synchronicity`

GetSynchronicity returns the Synchronicity field if non-nil, zero value otherwise.

### GetSynchronicityOk

`func (o *ServiceProfile) GetSynchronicityOk() (*Synchronicity, bool)`

GetSynchronicityOk returns a tuple with the Synchronicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronicity

`func (o *ServiceProfile) SetSynchronicity(v Synchronicity)`

SetSynchronicity sets Synchronicity field to given value.

### HasSynchronicity

`func (o *ServiceProfile) HasSynchronicity() bool`

HasSynchronicity returns a boolean if a field has been set.

### GetPositioning

`func (o *ServiceProfile) GetPositioning() Positioning`

GetPositioning returns the Positioning field if non-nil, zero value otherwise.

### GetPositioningOk

`func (o *ServiceProfile) GetPositioningOk() (*Positioning, bool)`

GetPositioningOk returns a tuple with the Positioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositioning

`func (o *ServiceProfile) SetPositioning(v Positioning)`

SetPositioning sets Positioning field to given value.

### HasPositioning

`func (o *ServiceProfile) HasPositioning() bool`

HasPositioning returns a boolean if a field has been set.

### GetUserMgmtOpen

`func (o *ServiceProfile) GetUserMgmtOpen() UserMgmtOpen`

GetUserMgmtOpen returns the UserMgmtOpen field if non-nil, zero value otherwise.

### GetUserMgmtOpenOk

`func (o *ServiceProfile) GetUserMgmtOpenOk() (*UserMgmtOpen, bool)`

GetUserMgmtOpenOk returns a tuple with the UserMgmtOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMgmtOpen

`func (o *ServiceProfile) SetUserMgmtOpen(v UserMgmtOpen)`

SetUserMgmtOpen sets UserMgmtOpen field to given value.

### HasUserMgmtOpen

`func (o *ServiceProfile) HasUserMgmtOpen() bool`

HasUserMgmtOpen returns a boolean if a field has been set.

### GetV2XModels

`func (o *ServiceProfile) GetV2XModels() V2XCommModels`

GetV2XModels returns the V2XModels field if non-nil, zero value otherwise.

### GetV2XModelsOk

`func (o *ServiceProfile) GetV2XModelsOk() (*V2XCommModels, bool)`

GetV2XModelsOk returns a tuple with the V2XModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2XModels

`func (o *ServiceProfile) SetV2XModels(v V2XCommModels)`

SetV2XModels sets V2XModels field to given value.

### HasV2XModels

`func (o *ServiceProfile) HasV2XModels() bool`

HasV2XModels returns a boolean if a field has been set.

### GetCoverageArea

`func (o *ServiceProfile) GetCoverageArea() string`

GetCoverageArea returns the CoverageArea field if non-nil, zero value otherwise.

### GetCoverageAreaOk

`func (o *ServiceProfile) GetCoverageAreaOk() (*string, bool)`

GetCoverageAreaOk returns a tuple with the CoverageArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageArea

`func (o *ServiceProfile) SetCoverageArea(v string)`

SetCoverageArea sets CoverageArea field to given value.

### HasCoverageArea

`func (o *ServiceProfile) HasCoverageArea() bool`

HasCoverageArea returns a boolean if a field has been set.

### GetTermDensity

`func (o *ServiceProfile) GetTermDensity() TermDensity`

GetTermDensity returns the TermDensity field if non-nil, zero value otherwise.

### GetTermDensityOk

`func (o *ServiceProfile) GetTermDensityOk() (*TermDensity, bool)`

GetTermDensityOk returns a tuple with the TermDensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermDensity

`func (o *ServiceProfile) SetTermDensity(v TermDensity)`

SetTermDensity sets TermDensity field to given value.

### HasTermDensity

`func (o *ServiceProfile) HasTermDensity() bool`

HasTermDensity returns a boolean if a field has been set.

### GetActivityFactor

`func (o *ServiceProfile) GetActivityFactor() float32`

GetActivityFactor returns the ActivityFactor field if non-nil, zero value otherwise.

### GetActivityFactorOk

`func (o *ServiceProfile) GetActivityFactorOk() (*float32, bool)`

GetActivityFactorOk returns a tuple with the ActivityFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityFactor

`func (o *ServiceProfile) SetActivityFactor(v float32)`

SetActivityFactor sets ActivityFactor field to given value.

### HasActivityFactor

`func (o *ServiceProfile) HasActivityFactor() bool`

HasActivityFactor returns a boolean if a field has been set.

### GetUESpeed

`func (o *ServiceProfile) GetUESpeed() int32`

GetUESpeed returns the UESpeed field if non-nil, zero value otherwise.

### GetUESpeedOk

`func (o *ServiceProfile) GetUESpeedOk() (*int32, bool)`

GetUESpeedOk returns a tuple with the UESpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUESpeed

`func (o *ServiceProfile) SetUESpeed(v int32)`

SetUESpeed sets UESpeed field to given value.

### HasUESpeed

`func (o *ServiceProfile) HasUESpeed() bool`

HasUESpeed returns a boolean if a field has been set.

### GetJitter

`func (o *ServiceProfile) GetJitter() int32`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *ServiceProfile) GetJitterOk() (*int32, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *ServiceProfile) SetJitter(v int32)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *ServiceProfile) HasJitter() bool`

HasJitter returns a boolean if a field has been set.

### GetSurvivalTime

`func (o *ServiceProfile) GetSurvivalTime() float32`

GetSurvivalTime returns the SurvivalTime field if non-nil, zero value otherwise.

### GetSurvivalTimeOk

`func (o *ServiceProfile) GetSurvivalTimeOk() (*float32, bool)`

GetSurvivalTimeOk returns a tuple with the SurvivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurvivalTime

`func (o *ServiceProfile) SetSurvivalTime(v float32)`

SetSurvivalTime sets SurvivalTime field to given value.

### HasSurvivalTime

`func (o *ServiceProfile) HasSurvivalTime() bool`

HasSurvivalTime returns a boolean if a field has been set.

### GetReliability

`func (o *ServiceProfile) GetReliability() float32`

GetReliability returns the Reliability field if non-nil, zero value otherwise.

### GetReliabilityOk

`func (o *ServiceProfile) GetReliabilityOk() (*float32, bool)`

GetReliabilityOk returns a tuple with the Reliability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliability

`func (o *ServiceProfile) SetReliability(v float32)`

SetReliability sets Reliability field to given value.

### HasReliability

`func (o *ServiceProfile) HasReliability() bool`

HasReliability returns a boolean if a field has been set.

### GetMaxDLDataVolume

`func (o *ServiceProfile) GetMaxDLDataVolume() float32`

GetMaxDLDataVolume returns the MaxDLDataVolume field if non-nil, zero value otherwise.

### GetMaxDLDataVolumeOk

`func (o *ServiceProfile) GetMaxDLDataVolumeOk() (*float32, bool)`

GetMaxDLDataVolumeOk returns a tuple with the MaxDLDataVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDLDataVolume

`func (o *ServiceProfile) SetMaxDLDataVolume(v float32)`

SetMaxDLDataVolume sets MaxDLDataVolume field to given value.

### HasMaxDLDataVolume

`func (o *ServiceProfile) HasMaxDLDataVolume() bool`

HasMaxDLDataVolume returns a boolean if a field has been set.

### GetMaxULDataVolume

`func (o *ServiceProfile) GetMaxULDataVolume() float32`

GetMaxULDataVolume returns the MaxULDataVolume field if non-nil, zero value otherwise.

### GetMaxULDataVolumeOk

`func (o *ServiceProfile) GetMaxULDataVolumeOk() (*float32, bool)`

GetMaxULDataVolumeOk returns a tuple with the MaxULDataVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxULDataVolume

`func (o *ServiceProfile) SetMaxULDataVolume(v float32)`

SetMaxULDataVolume sets MaxULDataVolume field to given value.

### HasMaxULDataVolume

`func (o *ServiceProfile) HasMaxULDataVolume() bool`

HasMaxULDataVolume returns a boolean if a field has been set.

### GetSliceSimultaneousUse

`func (o *ServiceProfile) GetSliceSimultaneousUse() SliceSimultaneousUse`

GetSliceSimultaneousUse returns the SliceSimultaneousUse field if non-nil, zero value otherwise.

### GetSliceSimultaneousUseOk

`func (o *ServiceProfile) GetSliceSimultaneousUseOk() (*SliceSimultaneousUse, bool)`

GetSliceSimultaneousUseOk returns a tuple with the SliceSimultaneousUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceSimultaneousUse

`func (o *ServiceProfile) SetSliceSimultaneousUse(v SliceSimultaneousUse)`

SetSliceSimultaneousUse sets SliceSimultaneousUse field to given value.

### HasSliceSimultaneousUse

`func (o *ServiceProfile) HasSliceSimultaneousUse() bool`

HasSliceSimultaneousUse returns a boolean if a field has been set.

### GetEnergyEfficiency

`func (o *ServiceProfile) GetEnergyEfficiency() EnergyEfficiency`

GetEnergyEfficiency returns the EnergyEfficiency field if non-nil, zero value otherwise.

### GetEnergyEfficiencyOk

`func (o *ServiceProfile) GetEnergyEfficiencyOk() (*EnergyEfficiency, bool)`

GetEnergyEfficiencyOk returns a tuple with the EnergyEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyEfficiency

`func (o *ServiceProfile) SetEnergyEfficiency(v EnergyEfficiency)`

SetEnergyEfficiency sets EnergyEfficiency field to given value.

### HasEnergyEfficiency

`func (o *ServiceProfile) HasEnergyEfficiency() bool`

HasEnergyEfficiency returns a boolean if a field has been set.

### GetNssaaSupport

`func (o *ServiceProfile) GetNssaaSupport() NSSAASupport`

GetNssaaSupport returns the NssaaSupport field if non-nil, zero value otherwise.

### GetNssaaSupportOk

`func (o *ServiceProfile) GetNssaaSupportOk() (*NSSAASupport, bool)`

GetNssaaSupportOk returns a tuple with the NssaaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssaaSupport

`func (o *ServiceProfile) SetNssaaSupport(v NSSAASupport)`

SetNssaaSupport sets NssaaSupport field to given value.

### HasNssaaSupport

`func (o *ServiceProfile) HasNssaaSupport() bool`

HasNssaaSupport returns a boolean if a field has been set.

### GetN6Protection

`func (o *ServiceProfile) GetN6Protection() N6Protection`

GetN6Protection returns the N6Protection field if non-nil, zero value otherwise.

### GetN6ProtectionOk

`func (o *ServiceProfile) GetN6ProtectionOk() (*N6Protection, bool)`

GetN6ProtectionOk returns a tuple with the N6Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN6Protection

`func (o *ServiceProfile) SetN6Protection(v N6Protection)`

SetN6Protection sets N6Protection field to given value.

### HasN6Protection

`func (o *ServiceProfile) HasN6Protection() bool`

HasN6Protection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


