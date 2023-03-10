/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// TopSliceSubnetProfile struct for TopSliceSubnetProfile
type TopSliceSubnetProfile struct {
	DLLatency *int32 `json:"dLLatency,omitempty"`
	ULLatency *int32 `json:"uLLatency,omitempty"`
	MaxNumberofUEs *int32 `json:"maxNumberofUEs,omitempty"`
	DLThptPerSliceSubnet *XLThpt `json:"dLThptPerSliceSubnet,omitempty"`
	DLThptPerUE *XLThpt `json:"dLThptPerUE,omitempty"`
	ULThptPerSliceSubnet *XLThpt `json:"uLThptPerSliceSubnet,omitempty"`
	ULThptPerUE *XLThpt `json:"uLThptPerUE,omitempty"`
	DLMaxPktSize *int32 `json:"dLMaxPktSize,omitempty"`
	ULMaxPktSize *int32 `json:"uLMaxPktSize,omitempty"`
	MaxNumberOfPDUSessions *int32 `json:"maxNumberOfPDUSessions,omitempty"`
	NROperatingBands *string `json:"nROperatingBands,omitempty"`
	SliceSimultaneousUse *SliceSimultaneousUse `json:"sliceSimultaneousUse,omitempty"`
	EnergyEfficiency *EnergyEfficiency `json:"energyEfficiency,omitempty"`
	Synchronicity *Synchronicity `json:"synchronicity,omitempty"`
	DelayTolerance *DelayTolerance `json:"delayTolerance,omitempty"`
	Positioning *Positioning `json:"positioning,omitempty"`
	TermDensity *TermDensity `json:"termDensity,omitempty"`
	ActivityFactor *int32 `json:"activityFactor,omitempty"`
	CoverageAreaTAList []int32 `json:"coverageAreaTAList,omitempty"`
	ResourceSharingLevel *SharingLevel `json:"resourceSharingLevel,omitempty"`
	UEMobilityLevel *MobilityLevel `json:"uEMobilityLevel,omitempty"`
	UESpeed *int32 `json:"uESpeed,omitempty"`
	Reliability *float32 `json:"reliability,omitempty"`
	SST *int32 `json:"sST,omitempty"`
	DLDeterministicComm *DeterministicComm `json:"dLDeterministicComm,omitempty"`
	ULDeterministicComm *DeterministicComm `json:"uLDeterministicComm,omitempty"`
	SurvivalTime *float32 `json:"survivalTime,omitempty"`
	NssaaSupport *NSSAASupport `json:"nssaaSupport,omitempty"`
	N6Protection *N6Protection `json:"n6Protection,omitempty"`
}

// NewTopSliceSubnetProfile instantiates a new TopSliceSubnetProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopSliceSubnetProfile() *TopSliceSubnetProfile {
	this := TopSliceSubnetProfile{}
	return &this
}

// NewTopSliceSubnetProfileWithDefaults instantiates a new TopSliceSubnetProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopSliceSubnetProfileWithDefaults() *TopSliceSubnetProfile {
	this := TopSliceSubnetProfile{}
	return &this
}

// GetDLLatency returns the DLLatency field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetDLLatency() int32 {
	if o == nil || isNil(o.DLLatency) {
		var ret int32
		return ret
	}
	return *o.DLLatency
}

// GetDLLatencyOk returns a tuple with the DLLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetDLLatencyOk() (*int32, bool) {
	if o == nil || isNil(o.DLLatency) {
    return nil, false
	}
	return o.DLLatency, true
}

// HasDLLatency returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasDLLatency() bool {
	if o != nil && !isNil(o.DLLatency) {
		return true
	}

	return false
}

// SetDLLatency gets a reference to the given int32 and assigns it to the DLLatency field.
func (o *TopSliceSubnetProfile) SetDLLatency(v int32) {
	o.DLLatency = &v
}

// GetULLatency returns the ULLatency field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetULLatency() int32 {
	if o == nil || isNil(o.ULLatency) {
		var ret int32
		return ret
	}
	return *o.ULLatency
}

// GetULLatencyOk returns a tuple with the ULLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetULLatencyOk() (*int32, bool) {
	if o == nil || isNil(o.ULLatency) {
    return nil, false
	}
	return o.ULLatency, true
}

// HasULLatency returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasULLatency() bool {
	if o != nil && !isNil(o.ULLatency) {
		return true
	}

	return false
}

// SetULLatency gets a reference to the given int32 and assigns it to the ULLatency field.
func (o *TopSliceSubnetProfile) SetULLatency(v int32) {
	o.ULLatency = &v
}

// GetMaxNumberofUEs returns the MaxNumberofUEs field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetMaxNumberofUEs() int32 {
	if o == nil || isNil(o.MaxNumberofUEs) {
		var ret int32
		return ret
	}
	return *o.MaxNumberofUEs
}

// GetMaxNumberofUEsOk returns a tuple with the MaxNumberofUEs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetMaxNumberofUEsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxNumberofUEs) {
    return nil, false
	}
	return o.MaxNumberofUEs, true
}

// HasMaxNumberofUEs returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasMaxNumberofUEs() bool {
	if o != nil && !isNil(o.MaxNumberofUEs) {
		return true
	}

	return false
}

// SetMaxNumberofUEs gets a reference to the given int32 and assigns it to the MaxNumberofUEs field.
func (o *TopSliceSubnetProfile) SetMaxNumberofUEs(v int32) {
	o.MaxNumberofUEs = &v
}

// GetDLThptPerSliceSubnet returns the DLThptPerSliceSubnet field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetDLThptPerSliceSubnet() XLThpt {
	if o == nil || isNil(o.DLThptPerSliceSubnet) {
		var ret XLThpt
		return ret
	}
	return *o.DLThptPerSliceSubnet
}

// GetDLThptPerSliceSubnetOk returns a tuple with the DLThptPerSliceSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetDLThptPerSliceSubnetOk() (*XLThpt, bool) {
	if o == nil || isNil(o.DLThptPerSliceSubnet) {
    return nil, false
	}
	return o.DLThptPerSliceSubnet, true
}

// HasDLThptPerSliceSubnet returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasDLThptPerSliceSubnet() bool {
	if o != nil && !isNil(o.DLThptPerSliceSubnet) {
		return true
	}

	return false
}

// SetDLThptPerSliceSubnet gets a reference to the given XLThpt and assigns it to the DLThptPerSliceSubnet field.
func (o *TopSliceSubnetProfile) SetDLThptPerSliceSubnet(v XLThpt) {
	o.DLThptPerSliceSubnet = &v
}

// GetDLThptPerUE returns the DLThptPerUE field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetDLThptPerUE() XLThpt {
	if o == nil || isNil(o.DLThptPerUE) {
		var ret XLThpt
		return ret
	}
	return *o.DLThptPerUE
}

// GetDLThptPerUEOk returns a tuple with the DLThptPerUE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetDLThptPerUEOk() (*XLThpt, bool) {
	if o == nil || isNil(o.DLThptPerUE) {
    return nil, false
	}
	return o.DLThptPerUE, true
}

// HasDLThptPerUE returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasDLThptPerUE() bool {
	if o != nil && !isNil(o.DLThptPerUE) {
		return true
	}

	return false
}

// SetDLThptPerUE gets a reference to the given XLThpt and assigns it to the DLThptPerUE field.
func (o *TopSliceSubnetProfile) SetDLThptPerUE(v XLThpt) {
	o.DLThptPerUE = &v
}

// GetULThptPerSliceSubnet returns the ULThptPerSliceSubnet field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetULThptPerSliceSubnet() XLThpt {
	if o == nil || isNil(o.ULThptPerSliceSubnet) {
		var ret XLThpt
		return ret
	}
	return *o.ULThptPerSliceSubnet
}

// GetULThptPerSliceSubnetOk returns a tuple with the ULThptPerSliceSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetULThptPerSliceSubnetOk() (*XLThpt, bool) {
	if o == nil || isNil(o.ULThptPerSliceSubnet) {
    return nil, false
	}
	return o.ULThptPerSliceSubnet, true
}

// HasULThptPerSliceSubnet returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasULThptPerSliceSubnet() bool {
	if o != nil && !isNil(o.ULThptPerSliceSubnet) {
		return true
	}

	return false
}

// SetULThptPerSliceSubnet gets a reference to the given XLThpt and assigns it to the ULThptPerSliceSubnet field.
func (o *TopSliceSubnetProfile) SetULThptPerSliceSubnet(v XLThpt) {
	o.ULThptPerSliceSubnet = &v
}

// GetULThptPerUE returns the ULThptPerUE field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetULThptPerUE() XLThpt {
	if o == nil || isNil(o.ULThptPerUE) {
		var ret XLThpt
		return ret
	}
	return *o.ULThptPerUE
}

// GetULThptPerUEOk returns a tuple with the ULThptPerUE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetULThptPerUEOk() (*XLThpt, bool) {
	if o == nil || isNil(o.ULThptPerUE) {
    return nil, false
	}
	return o.ULThptPerUE, true
}

// HasULThptPerUE returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasULThptPerUE() bool {
	if o != nil && !isNil(o.ULThptPerUE) {
		return true
	}

	return false
}

// SetULThptPerUE gets a reference to the given XLThpt and assigns it to the ULThptPerUE field.
func (o *TopSliceSubnetProfile) SetULThptPerUE(v XLThpt) {
	o.ULThptPerUE = &v
}

// GetDLMaxPktSize returns the DLMaxPktSize field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetDLMaxPktSize() int32 {
	if o == nil || isNil(o.DLMaxPktSize) {
		var ret int32
		return ret
	}
	return *o.DLMaxPktSize
}

// GetDLMaxPktSizeOk returns a tuple with the DLMaxPktSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetDLMaxPktSizeOk() (*int32, bool) {
	if o == nil || isNil(o.DLMaxPktSize) {
    return nil, false
	}
	return o.DLMaxPktSize, true
}

// HasDLMaxPktSize returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasDLMaxPktSize() bool {
	if o != nil && !isNil(o.DLMaxPktSize) {
		return true
	}

	return false
}

// SetDLMaxPktSize gets a reference to the given int32 and assigns it to the DLMaxPktSize field.
func (o *TopSliceSubnetProfile) SetDLMaxPktSize(v int32) {
	o.DLMaxPktSize = &v
}

// GetULMaxPktSize returns the ULMaxPktSize field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetULMaxPktSize() int32 {
	if o == nil || isNil(o.ULMaxPktSize) {
		var ret int32
		return ret
	}
	return *o.ULMaxPktSize
}

// GetULMaxPktSizeOk returns a tuple with the ULMaxPktSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetULMaxPktSizeOk() (*int32, bool) {
	if o == nil || isNil(o.ULMaxPktSize) {
    return nil, false
	}
	return o.ULMaxPktSize, true
}

// HasULMaxPktSize returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasULMaxPktSize() bool {
	if o != nil && !isNil(o.ULMaxPktSize) {
		return true
	}

	return false
}

// SetULMaxPktSize gets a reference to the given int32 and assigns it to the ULMaxPktSize field.
func (o *TopSliceSubnetProfile) SetULMaxPktSize(v int32) {
	o.ULMaxPktSize = &v
}

// GetMaxNumberOfPDUSessions returns the MaxNumberOfPDUSessions field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetMaxNumberOfPDUSessions() int32 {
	if o == nil || isNil(o.MaxNumberOfPDUSessions) {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfPDUSessions
}

// GetMaxNumberOfPDUSessionsOk returns a tuple with the MaxNumberOfPDUSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetMaxNumberOfPDUSessionsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxNumberOfPDUSessions) {
    return nil, false
	}
	return o.MaxNumberOfPDUSessions, true
}

// HasMaxNumberOfPDUSessions returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasMaxNumberOfPDUSessions() bool {
	if o != nil && !isNil(o.MaxNumberOfPDUSessions) {
		return true
	}

	return false
}

// SetMaxNumberOfPDUSessions gets a reference to the given int32 and assigns it to the MaxNumberOfPDUSessions field.
func (o *TopSliceSubnetProfile) SetMaxNumberOfPDUSessions(v int32) {
	o.MaxNumberOfPDUSessions = &v
}

// GetNROperatingBands returns the NROperatingBands field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetNROperatingBands() string {
	if o == nil || isNil(o.NROperatingBands) {
		var ret string
		return ret
	}
	return *o.NROperatingBands
}

// GetNROperatingBandsOk returns a tuple with the NROperatingBands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetNROperatingBandsOk() (*string, bool) {
	if o == nil || isNil(o.NROperatingBands) {
    return nil, false
	}
	return o.NROperatingBands, true
}

// HasNROperatingBands returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasNROperatingBands() bool {
	if o != nil && !isNil(o.NROperatingBands) {
		return true
	}

	return false
}

// SetNROperatingBands gets a reference to the given string and assigns it to the NROperatingBands field.
func (o *TopSliceSubnetProfile) SetNROperatingBands(v string) {
	o.NROperatingBands = &v
}

// GetSliceSimultaneousUse returns the SliceSimultaneousUse field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetSliceSimultaneousUse() SliceSimultaneousUse {
	if o == nil || isNil(o.SliceSimultaneousUse) {
		var ret SliceSimultaneousUse
		return ret
	}
	return *o.SliceSimultaneousUse
}

// GetSliceSimultaneousUseOk returns a tuple with the SliceSimultaneousUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetSliceSimultaneousUseOk() (*SliceSimultaneousUse, bool) {
	if o == nil || isNil(o.SliceSimultaneousUse) {
    return nil, false
	}
	return o.SliceSimultaneousUse, true
}

// HasSliceSimultaneousUse returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasSliceSimultaneousUse() bool {
	if o != nil && !isNil(o.SliceSimultaneousUse) {
		return true
	}

	return false
}

// SetSliceSimultaneousUse gets a reference to the given SliceSimultaneousUse and assigns it to the SliceSimultaneousUse field.
func (o *TopSliceSubnetProfile) SetSliceSimultaneousUse(v SliceSimultaneousUse) {
	o.SliceSimultaneousUse = &v
}

// GetEnergyEfficiency returns the EnergyEfficiency field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetEnergyEfficiency() EnergyEfficiency {
	if o == nil || isNil(o.EnergyEfficiency) {
		var ret EnergyEfficiency
		return ret
	}
	return *o.EnergyEfficiency
}

// GetEnergyEfficiencyOk returns a tuple with the EnergyEfficiency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetEnergyEfficiencyOk() (*EnergyEfficiency, bool) {
	if o == nil || isNil(o.EnergyEfficiency) {
    return nil, false
	}
	return o.EnergyEfficiency, true
}

// HasEnergyEfficiency returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasEnergyEfficiency() bool {
	if o != nil && !isNil(o.EnergyEfficiency) {
		return true
	}

	return false
}

// SetEnergyEfficiency gets a reference to the given EnergyEfficiency and assigns it to the EnergyEfficiency field.
func (o *TopSliceSubnetProfile) SetEnergyEfficiency(v EnergyEfficiency) {
	o.EnergyEfficiency = &v
}

// GetSynchronicity returns the Synchronicity field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetSynchronicity() Synchronicity {
	if o == nil || isNil(o.Synchronicity) {
		var ret Synchronicity
		return ret
	}
	return *o.Synchronicity
}

// GetSynchronicityOk returns a tuple with the Synchronicity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetSynchronicityOk() (*Synchronicity, bool) {
	if o == nil || isNil(o.Synchronicity) {
    return nil, false
	}
	return o.Synchronicity, true
}

// HasSynchronicity returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasSynchronicity() bool {
	if o != nil && !isNil(o.Synchronicity) {
		return true
	}

	return false
}

// SetSynchronicity gets a reference to the given Synchronicity and assigns it to the Synchronicity field.
func (o *TopSliceSubnetProfile) SetSynchronicity(v Synchronicity) {
	o.Synchronicity = &v
}

// GetDelayTolerance returns the DelayTolerance field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetDelayTolerance() DelayTolerance {
	if o == nil || isNil(o.DelayTolerance) {
		var ret DelayTolerance
		return ret
	}
	return *o.DelayTolerance
}

// GetDelayToleranceOk returns a tuple with the DelayTolerance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetDelayToleranceOk() (*DelayTolerance, bool) {
	if o == nil || isNil(o.DelayTolerance) {
    return nil, false
	}
	return o.DelayTolerance, true
}

// HasDelayTolerance returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasDelayTolerance() bool {
	if o != nil && !isNil(o.DelayTolerance) {
		return true
	}

	return false
}

// SetDelayTolerance gets a reference to the given DelayTolerance and assigns it to the DelayTolerance field.
func (o *TopSliceSubnetProfile) SetDelayTolerance(v DelayTolerance) {
	o.DelayTolerance = &v
}

// GetPositioning returns the Positioning field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetPositioning() Positioning {
	if o == nil || isNil(o.Positioning) {
		var ret Positioning
		return ret
	}
	return *o.Positioning
}

// GetPositioningOk returns a tuple with the Positioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetPositioningOk() (*Positioning, bool) {
	if o == nil || isNil(o.Positioning) {
    return nil, false
	}
	return o.Positioning, true
}

// HasPositioning returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasPositioning() bool {
	if o != nil && !isNil(o.Positioning) {
		return true
	}

	return false
}

// SetPositioning gets a reference to the given Positioning and assigns it to the Positioning field.
func (o *TopSliceSubnetProfile) SetPositioning(v Positioning) {
	o.Positioning = &v
}

// GetTermDensity returns the TermDensity field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetTermDensity() TermDensity {
	if o == nil || isNil(o.TermDensity) {
		var ret TermDensity
		return ret
	}
	return *o.TermDensity
}

// GetTermDensityOk returns a tuple with the TermDensity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetTermDensityOk() (*TermDensity, bool) {
	if o == nil || isNil(o.TermDensity) {
    return nil, false
	}
	return o.TermDensity, true
}

// HasTermDensity returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasTermDensity() bool {
	if o != nil && !isNil(o.TermDensity) {
		return true
	}

	return false
}

// SetTermDensity gets a reference to the given TermDensity and assigns it to the TermDensity field.
func (o *TopSliceSubnetProfile) SetTermDensity(v TermDensity) {
	o.TermDensity = &v
}

// GetActivityFactor returns the ActivityFactor field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetActivityFactor() int32 {
	if o == nil || isNil(o.ActivityFactor) {
		var ret int32
		return ret
	}
	return *o.ActivityFactor
}

// GetActivityFactorOk returns a tuple with the ActivityFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetActivityFactorOk() (*int32, bool) {
	if o == nil || isNil(o.ActivityFactor) {
    return nil, false
	}
	return o.ActivityFactor, true
}

// HasActivityFactor returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasActivityFactor() bool {
	if o != nil && !isNil(o.ActivityFactor) {
		return true
	}

	return false
}

// SetActivityFactor gets a reference to the given int32 and assigns it to the ActivityFactor field.
func (o *TopSliceSubnetProfile) SetActivityFactor(v int32) {
	o.ActivityFactor = &v
}

// GetCoverageAreaTAList returns the CoverageAreaTAList field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetCoverageAreaTAList() []int32 {
	if o == nil || isNil(o.CoverageAreaTAList) {
		var ret []int32
		return ret
	}
	return o.CoverageAreaTAList
}

// GetCoverageAreaTAListOk returns a tuple with the CoverageAreaTAList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetCoverageAreaTAListOk() ([]int32, bool) {
	if o == nil || isNil(o.CoverageAreaTAList) {
    return nil, false
	}
	return o.CoverageAreaTAList, true
}

// HasCoverageAreaTAList returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasCoverageAreaTAList() bool {
	if o != nil && !isNil(o.CoverageAreaTAList) {
		return true
	}

	return false
}

// SetCoverageAreaTAList gets a reference to the given []int32 and assigns it to the CoverageAreaTAList field.
func (o *TopSliceSubnetProfile) SetCoverageAreaTAList(v []int32) {
	o.CoverageAreaTAList = v
}

// GetResourceSharingLevel returns the ResourceSharingLevel field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetResourceSharingLevel() SharingLevel {
	if o == nil || isNil(o.ResourceSharingLevel) {
		var ret SharingLevel
		return ret
	}
	return *o.ResourceSharingLevel
}

// GetResourceSharingLevelOk returns a tuple with the ResourceSharingLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetResourceSharingLevelOk() (*SharingLevel, bool) {
	if o == nil || isNil(o.ResourceSharingLevel) {
    return nil, false
	}
	return o.ResourceSharingLevel, true
}

// HasResourceSharingLevel returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasResourceSharingLevel() bool {
	if o != nil && !isNil(o.ResourceSharingLevel) {
		return true
	}

	return false
}

// SetResourceSharingLevel gets a reference to the given SharingLevel and assigns it to the ResourceSharingLevel field.
func (o *TopSliceSubnetProfile) SetResourceSharingLevel(v SharingLevel) {
	o.ResourceSharingLevel = &v
}

// GetUEMobilityLevel returns the UEMobilityLevel field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetUEMobilityLevel() MobilityLevel {
	if o == nil || isNil(o.UEMobilityLevel) {
		var ret MobilityLevel
		return ret
	}
	return *o.UEMobilityLevel
}

// GetUEMobilityLevelOk returns a tuple with the UEMobilityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetUEMobilityLevelOk() (*MobilityLevel, bool) {
	if o == nil || isNil(o.UEMobilityLevel) {
    return nil, false
	}
	return o.UEMobilityLevel, true
}

// HasUEMobilityLevel returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasUEMobilityLevel() bool {
	if o != nil && !isNil(o.UEMobilityLevel) {
		return true
	}

	return false
}

// SetUEMobilityLevel gets a reference to the given MobilityLevel and assigns it to the UEMobilityLevel field.
func (o *TopSliceSubnetProfile) SetUEMobilityLevel(v MobilityLevel) {
	o.UEMobilityLevel = &v
}

// GetUESpeed returns the UESpeed field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetUESpeed() int32 {
	if o == nil || isNil(o.UESpeed) {
		var ret int32
		return ret
	}
	return *o.UESpeed
}

// GetUESpeedOk returns a tuple with the UESpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetUESpeedOk() (*int32, bool) {
	if o == nil || isNil(o.UESpeed) {
    return nil, false
	}
	return o.UESpeed, true
}

// HasUESpeed returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasUESpeed() bool {
	if o != nil && !isNil(o.UESpeed) {
		return true
	}

	return false
}

// SetUESpeed gets a reference to the given int32 and assigns it to the UESpeed field.
func (o *TopSliceSubnetProfile) SetUESpeed(v int32) {
	o.UESpeed = &v
}

// GetReliability returns the Reliability field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetReliability() float32 {
	if o == nil || isNil(o.Reliability) {
		var ret float32
		return ret
	}
	return *o.Reliability
}

// GetReliabilityOk returns a tuple with the Reliability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetReliabilityOk() (*float32, bool) {
	if o == nil || isNil(o.Reliability) {
    return nil, false
	}
	return o.Reliability, true
}

// HasReliability returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasReliability() bool {
	if o != nil && !isNil(o.Reliability) {
		return true
	}

	return false
}

// SetReliability gets a reference to the given float32 and assigns it to the Reliability field.
func (o *TopSliceSubnetProfile) SetReliability(v float32) {
	o.Reliability = &v
}

// GetSST returns the SST field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetSST() int32 {
	if o == nil || isNil(o.SST) {
		var ret int32
		return ret
	}
	return *o.SST
}

// GetSSTOk returns a tuple with the SST field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetSSTOk() (*int32, bool) {
	if o == nil || isNil(o.SST) {
    return nil, false
	}
	return o.SST, true
}

// HasSST returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasSST() bool {
	if o != nil && !isNil(o.SST) {
		return true
	}

	return false
}

// SetSST gets a reference to the given int32 and assigns it to the SST field.
func (o *TopSliceSubnetProfile) SetSST(v int32) {
	o.SST = &v
}

// GetDLDeterministicComm returns the DLDeterministicComm field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetDLDeterministicComm() DeterministicComm {
	if o == nil || isNil(o.DLDeterministicComm) {
		var ret DeterministicComm
		return ret
	}
	return *o.DLDeterministicComm
}

// GetDLDeterministicCommOk returns a tuple with the DLDeterministicComm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetDLDeterministicCommOk() (*DeterministicComm, bool) {
	if o == nil || isNil(o.DLDeterministicComm) {
    return nil, false
	}
	return o.DLDeterministicComm, true
}

// HasDLDeterministicComm returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasDLDeterministicComm() bool {
	if o != nil && !isNil(o.DLDeterministicComm) {
		return true
	}

	return false
}

// SetDLDeterministicComm gets a reference to the given DeterministicComm and assigns it to the DLDeterministicComm field.
func (o *TopSliceSubnetProfile) SetDLDeterministicComm(v DeterministicComm) {
	o.DLDeterministicComm = &v
}

// GetULDeterministicComm returns the ULDeterministicComm field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetULDeterministicComm() DeterministicComm {
	if o == nil || isNil(o.ULDeterministicComm) {
		var ret DeterministicComm
		return ret
	}
	return *o.ULDeterministicComm
}

// GetULDeterministicCommOk returns a tuple with the ULDeterministicComm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetULDeterministicCommOk() (*DeterministicComm, bool) {
	if o == nil || isNil(o.ULDeterministicComm) {
    return nil, false
	}
	return o.ULDeterministicComm, true
}

// HasULDeterministicComm returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasULDeterministicComm() bool {
	if o != nil && !isNil(o.ULDeterministicComm) {
		return true
	}

	return false
}

// SetULDeterministicComm gets a reference to the given DeterministicComm and assigns it to the ULDeterministicComm field.
func (o *TopSliceSubnetProfile) SetULDeterministicComm(v DeterministicComm) {
	o.ULDeterministicComm = &v
}

// GetSurvivalTime returns the SurvivalTime field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetSurvivalTime() float32 {
	if o == nil || isNil(o.SurvivalTime) {
		var ret float32
		return ret
	}
	return *o.SurvivalTime
}

// GetSurvivalTimeOk returns a tuple with the SurvivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetSurvivalTimeOk() (*float32, bool) {
	if o == nil || isNil(o.SurvivalTime) {
    return nil, false
	}
	return o.SurvivalTime, true
}

// HasSurvivalTime returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasSurvivalTime() bool {
	if o != nil && !isNil(o.SurvivalTime) {
		return true
	}

	return false
}

// SetSurvivalTime gets a reference to the given float32 and assigns it to the SurvivalTime field.
func (o *TopSliceSubnetProfile) SetSurvivalTime(v float32) {
	o.SurvivalTime = &v
}

// GetNssaaSupport returns the NssaaSupport field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetNssaaSupport() NSSAASupport {
	if o == nil || isNil(o.NssaaSupport) {
		var ret NSSAASupport
		return ret
	}
	return *o.NssaaSupport
}

// GetNssaaSupportOk returns a tuple with the NssaaSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetNssaaSupportOk() (*NSSAASupport, bool) {
	if o == nil || isNil(o.NssaaSupport) {
    return nil, false
	}
	return o.NssaaSupport, true
}

// HasNssaaSupport returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasNssaaSupport() bool {
	if o != nil && !isNil(o.NssaaSupport) {
		return true
	}

	return false
}

// SetNssaaSupport gets a reference to the given NSSAASupport and assigns it to the NssaaSupport field.
func (o *TopSliceSubnetProfile) SetNssaaSupport(v NSSAASupport) {
	o.NssaaSupport = &v
}

// GetN6Protection returns the N6Protection field value if set, zero value otherwise.
func (o *TopSliceSubnetProfile) GetN6Protection() N6Protection {
	if o == nil || isNil(o.N6Protection) {
		var ret N6Protection
		return ret
	}
	return *o.N6Protection
}

// GetN6ProtectionOk returns a tuple with the N6Protection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopSliceSubnetProfile) GetN6ProtectionOk() (*N6Protection, bool) {
	if o == nil || isNil(o.N6Protection) {
    return nil, false
	}
	return o.N6Protection, true
}

// HasN6Protection returns a boolean if a field has been set.
func (o *TopSliceSubnetProfile) HasN6Protection() bool {
	if o != nil && !isNil(o.N6Protection) {
		return true
	}

	return false
}

// SetN6Protection gets a reference to the given N6Protection and assigns it to the N6Protection field.
func (o *TopSliceSubnetProfile) SetN6Protection(v N6Protection) {
	o.N6Protection = &v
}

func (o TopSliceSubnetProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DLLatency) {
		toSerialize["dLLatency"] = o.DLLatency
	}
	if !isNil(o.ULLatency) {
		toSerialize["uLLatency"] = o.ULLatency
	}
	if !isNil(o.MaxNumberofUEs) {
		toSerialize["maxNumberofUEs"] = o.MaxNumberofUEs
	}
	if !isNil(o.DLThptPerSliceSubnet) {
		toSerialize["dLThptPerSliceSubnet"] = o.DLThptPerSliceSubnet
	}
	if !isNil(o.DLThptPerUE) {
		toSerialize["dLThptPerUE"] = o.DLThptPerUE
	}
	if !isNil(o.ULThptPerSliceSubnet) {
		toSerialize["uLThptPerSliceSubnet"] = o.ULThptPerSliceSubnet
	}
	if !isNil(o.ULThptPerUE) {
		toSerialize["uLThptPerUE"] = o.ULThptPerUE
	}
	if !isNil(o.DLMaxPktSize) {
		toSerialize["dLMaxPktSize"] = o.DLMaxPktSize
	}
	if !isNil(o.ULMaxPktSize) {
		toSerialize["uLMaxPktSize"] = o.ULMaxPktSize
	}
	if !isNil(o.MaxNumberOfPDUSessions) {
		toSerialize["maxNumberOfPDUSessions"] = o.MaxNumberOfPDUSessions
	}
	if !isNil(o.NROperatingBands) {
		toSerialize["nROperatingBands"] = o.NROperatingBands
	}
	if !isNil(o.SliceSimultaneousUse) {
		toSerialize["sliceSimultaneousUse"] = o.SliceSimultaneousUse
	}
	if !isNil(o.EnergyEfficiency) {
		toSerialize["energyEfficiency"] = o.EnergyEfficiency
	}
	if !isNil(o.Synchronicity) {
		toSerialize["synchronicity"] = o.Synchronicity
	}
	if !isNil(o.DelayTolerance) {
		toSerialize["delayTolerance"] = o.DelayTolerance
	}
	if !isNil(o.Positioning) {
		toSerialize["positioning"] = o.Positioning
	}
	if !isNil(o.TermDensity) {
		toSerialize["termDensity"] = o.TermDensity
	}
	if !isNil(o.ActivityFactor) {
		toSerialize["activityFactor"] = o.ActivityFactor
	}
	if !isNil(o.CoverageAreaTAList) {
		toSerialize["coverageAreaTAList"] = o.CoverageAreaTAList
	}
	if !isNil(o.ResourceSharingLevel) {
		toSerialize["resourceSharingLevel"] = o.ResourceSharingLevel
	}
	if !isNil(o.UEMobilityLevel) {
		toSerialize["uEMobilityLevel"] = o.UEMobilityLevel
	}
	if !isNil(o.UESpeed) {
		toSerialize["uESpeed"] = o.UESpeed
	}
	if !isNil(o.Reliability) {
		toSerialize["reliability"] = o.Reliability
	}
	if !isNil(o.SST) {
		toSerialize["sST"] = o.SST
	}
	if !isNil(o.DLDeterministicComm) {
		toSerialize["dLDeterministicComm"] = o.DLDeterministicComm
	}
	if !isNil(o.ULDeterministicComm) {
		toSerialize["uLDeterministicComm"] = o.ULDeterministicComm
	}
	if !isNil(o.SurvivalTime) {
		toSerialize["survivalTime"] = o.SurvivalTime
	}
	if !isNil(o.NssaaSupport) {
		toSerialize["nssaaSupport"] = o.NssaaSupport
	}
	if !isNil(o.N6Protection) {
		toSerialize["n6Protection"] = o.N6Protection
	}
	return json.Marshal(toSerialize)
}

type NullableTopSliceSubnetProfile struct {
	value *TopSliceSubnetProfile
	isSet bool
}

func (v NullableTopSliceSubnetProfile) Get() *TopSliceSubnetProfile {
	return v.value
}

func (v *NullableTopSliceSubnetProfile) Set(val *TopSliceSubnetProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableTopSliceSubnetProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableTopSliceSubnetProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopSliceSubnetProfile(val *TopSliceSubnetProfile) *NullableTopSliceSubnetProfile {
	return &NullableTopSliceSubnetProfile{value: val, isSet: true}
}

func (v NullableTopSliceSubnetProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopSliceSubnetProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


