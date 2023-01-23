/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// MonitoringConfiguration struct for MonitoringConfiguration
type MonitoringConfiguration struct {
	EventType EventType `json:"eventType"`
	ImmediateFlag *bool `json:"immediateFlag,omitempty"`
	LocationReportingConfiguration *LocationReportingConfiguration `json:"locationReportingConfiguration,omitempty"`
	AssociationType *AssociationType `json:"associationType,omitempty"`
	DatalinkReportCfg *DatalinkReportingConfiguration `json:"datalinkReportCfg,omitempty"`
	LossConnectivityCfg *LossConnectivityCfg `json:"lossConnectivityCfg,omitempty"`
	// indicating a time in seconds.
	MaximumLatency *int32 `json:"maximumLatency,omitempty"`
	// indicating a time in seconds.
	MaximumResponseTime *int32 `json:"maximumResponseTime,omitempty"`
	SuggestedPacketNumDl *int32 `json:"suggestedPacketNumDl,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	SingleNssai *Snssai `json:"singleNssai,omitempty"`
	PduSessionStatusCfg *PduSessionStatusCfg `json:"pduSessionStatusCfg,omitempty"`
	ReachabilityForSmsCfg *ReachabilityForSmsConfiguration `json:"reachabilityForSmsCfg,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation *string `json:"mtcProviderInformation,omitempty"`
	AfId *string `json:"afId,omitempty"`
	ReachabilityForDataCfg *ReachabilityForDataConfiguration `json:"reachabilityForDataCfg,omitempty"`
	IdleStatusInd *bool `json:"idleStatusInd,omitempty"`
}

// NewMonitoringConfiguration instantiates a new MonitoringConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringConfiguration(eventType EventType) *MonitoringConfiguration {
	this := MonitoringConfiguration{}
	this.EventType = eventType
	var idleStatusInd bool = false
	this.IdleStatusInd = &idleStatusInd
	return &this
}

// NewMonitoringConfigurationWithDefaults instantiates a new MonitoringConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringConfigurationWithDefaults() *MonitoringConfiguration {
	this := MonitoringConfiguration{}
	var idleStatusInd bool = false
	this.IdleStatusInd = &idleStatusInd
	return &this
}

// GetEventType returns the EventType field value
func (o *MonitoringConfiguration) GetEventType() EventType {
	if o == nil {
		var ret EventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetEventTypeOk() (*EventType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *MonitoringConfiguration) SetEventType(v EventType) {
	o.EventType = v
}

// GetImmediateFlag returns the ImmediateFlag field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetImmediateFlag() bool {
	if o == nil || isNil(o.ImmediateFlag) {
		var ret bool
		return ret
	}
	return *o.ImmediateFlag
}

// GetImmediateFlagOk returns a tuple with the ImmediateFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetImmediateFlagOk() (*bool, bool) {
	if o == nil || isNil(o.ImmediateFlag) {
    return nil, false
	}
	return o.ImmediateFlag, true
}

// HasImmediateFlag returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasImmediateFlag() bool {
	if o != nil && !isNil(o.ImmediateFlag) {
		return true
	}

	return false
}

// SetImmediateFlag gets a reference to the given bool and assigns it to the ImmediateFlag field.
func (o *MonitoringConfiguration) SetImmediateFlag(v bool) {
	o.ImmediateFlag = &v
}

// GetLocationReportingConfiguration returns the LocationReportingConfiguration field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetLocationReportingConfiguration() LocationReportingConfiguration {
	if o == nil || isNil(o.LocationReportingConfiguration) {
		var ret LocationReportingConfiguration
		return ret
	}
	return *o.LocationReportingConfiguration
}

// GetLocationReportingConfigurationOk returns a tuple with the LocationReportingConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetLocationReportingConfigurationOk() (*LocationReportingConfiguration, bool) {
	if o == nil || isNil(o.LocationReportingConfiguration) {
    return nil, false
	}
	return o.LocationReportingConfiguration, true
}

// HasLocationReportingConfiguration returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasLocationReportingConfiguration() bool {
	if o != nil && !isNil(o.LocationReportingConfiguration) {
		return true
	}

	return false
}

// SetLocationReportingConfiguration gets a reference to the given LocationReportingConfiguration and assigns it to the LocationReportingConfiguration field.
func (o *MonitoringConfiguration) SetLocationReportingConfiguration(v LocationReportingConfiguration) {
	o.LocationReportingConfiguration = &v
}

// GetAssociationType returns the AssociationType field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetAssociationType() AssociationType {
	if o == nil || isNil(o.AssociationType) {
		var ret AssociationType
		return ret
	}
	return *o.AssociationType
}

// GetAssociationTypeOk returns a tuple with the AssociationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetAssociationTypeOk() (*AssociationType, bool) {
	if o == nil || isNil(o.AssociationType) {
    return nil, false
	}
	return o.AssociationType, true
}

// HasAssociationType returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasAssociationType() bool {
	if o != nil && !isNil(o.AssociationType) {
		return true
	}

	return false
}

// SetAssociationType gets a reference to the given AssociationType and assigns it to the AssociationType field.
func (o *MonitoringConfiguration) SetAssociationType(v AssociationType) {
	o.AssociationType = &v
}

// GetDatalinkReportCfg returns the DatalinkReportCfg field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetDatalinkReportCfg() DatalinkReportingConfiguration {
	if o == nil || isNil(o.DatalinkReportCfg) {
		var ret DatalinkReportingConfiguration
		return ret
	}
	return *o.DatalinkReportCfg
}

// GetDatalinkReportCfgOk returns a tuple with the DatalinkReportCfg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetDatalinkReportCfgOk() (*DatalinkReportingConfiguration, bool) {
	if o == nil || isNil(o.DatalinkReportCfg) {
    return nil, false
	}
	return o.DatalinkReportCfg, true
}

// HasDatalinkReportCfg returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasDatalinkReportCfg() bool {
	if o != nil && !isNil(o.DatalinkReportCfg) {
		return true
	}

	return false
}

// SetDatalinkReportCfg gets a reference to the given DatalinkReportingConfiguration and assigns it to the DatalinkReportCfg field.
func (o *MonitoringConfiguration) SetDatalinkReportCfg(v DatalinkReportingConfiguration) {
	o.DatalinkReportCfg = &v
}

// GetLossConnectivityCfg returns the LossConnectivityCfg field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetLossConnectivityCfg() LossConnectivityCfg {
	if o == nil || isNil(o.LossConnectivityCfg) {
		var ret LossConnectivityCfg
		return ret
	}
	return *o.LossConnectivityCfg
}

// GetLossConnectivityCfgOk returns a tuple with the LossConnectivityCfg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetLossConnectivityCfgOk() (*LossConnectivityCfg, bool) {
	if o == nil || isNil(o.LossConnectivityCfg) {
    return nil, false
	}
	return o.LossConnectivityCfg, true
}

// HasLossConnectivityCfg returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasLossConnectivityCfg() bool {
	if o != nil && !isNil(o.LossConnectivityCfg) {
		return true
	}

	return false
}

// SetLossConnectivityCfg gets a reference to the given LossConnectivityCfg and assigns it to the LossConnectivityCfg field.
func (o *MonitoringConfiguration) SetLossConnectivityCfg(v LossConnectivityCfg) {
	o.LossConnectivityCfg = &v
}

// GetMaximumLatency returns the MaximumLatency field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetMaximumLatency() int32 {
	if o == nil || isNil(o.MaximumLatency) {
		var ret int32
		return ret
	}
	return *o.MaximumLatency
}

// GetMaximumLatencyOk returns a tuple with the MaximumLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetMaximumLatencyOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumLatency) {
    return nil, false
	}
	return o.MaximumLatency, true
}

// HasMaximumLatency returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasMaximumLatency() bool {
	if o != nil && !isNil(o.MaximumLatency) {
		return true
	}

	return false
}

// SetMaximumLatency gets a reference to the given int32 and assigns it to the MaximumLatency field.
func (o *MonitoringConfiguration) SetMaximumLatency(v int32) {
	o.MaximumLatency = &v
}

// GetMaximumResponseTime returns the MaximumResponseTime field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetMaximumResponseTime() int32 {
	if o == nil || isNil(o.MaximumResponseTime) {
		var ret int32
		return ret
	}
	return *o.MaximumResponseTime
}

// GetMaximumResponseTimeOk returns a tuple with the MaximumResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetMaximumResponseTimeOk() (*int32, bool) {
	if o == nil || isNil(o.MaximumResponseTime) {
    return nil, false
	}
	return o.MaximumResponseTime, true
}

// HasMaximumResponseTime returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasMaximumResponseTime() bool {
	if o != nil && !isNil(o.MaximumResponseTime) {
		return true
	}

	return false
}

// SetMaximumResponseTime gets a reference to the given int32 and assigns it to the MaximumResponseTime field.
func (o *MonitoringConfiguration) SetMaximumResponseTime(v int32) {
	o.MaximumResponseTime = &v
}

// GetSuggestedPacketNumDl returns the SuggestedPacketNumDl field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetSuggestedPacketNumDl() int32 {
	if o == nil || isNil(o.SuggestedPacketNumDl) {
		var ret int32
		return ret
	}
	return *o.SuggestedPacketNumDl
}

// GetSuggestedPacketNumDlOk returns a tuple with the SuggestedPacketNumDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetSuggestedPacketNumDlOk() (*int32, bool) {
	if o == nil || isNil(o.SuggestedPacketNumDl) {
    return nil, false
	}
	return o.SuggestedPacketNumDl, true
}

// HasSuggestedPacketNumDl returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasSuggestedPacketNumDl() bool {
	if o != nil && !isNil(o.SuggestedPacketNumDl) {
		return true
	}

	return false
}

// SetSuggestedPacketNumDl gets a reference to the given int32 and assigns it to the SuggestedPacketNumDl field.
func (o *MonitoringConfiguration) SetSuggestedPacketNumDl(v int32) {
	o.SuggestedPacketNumDl = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *MonitoringConfiguration) SetDnn(v string) {
	o.Dnn = &v
}

// GetSingleNssai returns the SingleNssai field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetSingleNssai() Snssai {
	if o == nil || isNil(o.SingleNssai) {
		var ret Snssai
		return ret
	}
	return *o.SingleNssai
}

// GetSingleNssaiOk returns a tuple with the SingleNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetSingleNssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.SingleNssai) {
    return nil, false
	}
	return o.SingleNssai, true
}

// HasSingleNssai returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasSingleNssai() bool {
	if o != nil && !isNil(o.SingleNssai) {
		return true
	}

	return false
}

// SetSingleNssai gets a reference to the given Snssai and assigns it to the SingleNssai field.
func (o *MonitoringConfiguration) SetSingleNssai(v Snssai) {
	o.SingleNssai = &v
}

// GetPduSessionStatusCfg returns the PduSessionStatusCfg field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetPduSessionStatusCfg() PduSessionStatusCfg {
	if o == nil || isNil(o.PduSessionStatusCfg) {
		var ret PduSessionStatusCfg
		return ret
	}
	return *o.PduSessionStatusCfg
}

// GetPduSessionStatusCfgOk returns a tuple with the PduSessionStatusCfg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetPduSessionStatusCfgOk() (*PduSessionStatusCfg, bool) {
	if o == nil || isNil(o.PduSessionStatusCfg) {
    return nil, false
	}
	return o.PduSessionStatusCfg, true
}

// HasPduSessionStatusCfg returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasPduSessionStatusCfg() bool {
	if o != nil && !isNil(o.PduSessionStatusCfg) {
		return true
	}

	return false
}

// SetPduSessionStatusCfg gets a reference to the given PduSessionStatusCfg and assigns it to the PduSessionStatusCfg field.
func (o *MonitoringConfiguration) SetPduSessionStatusCfg(v PduSessionStatusCfg) {
	o.PduSessionStatusCfg = &v
}

// GetReachabilityForSmsCfg returns the ReachabilityForSmsCfg field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetReachabilityForSmsCfg() ReachabilityForSmsConfiguration {
	if o == nil || isNil(o.ReachabilityForSmsCfg) {
		var ret ReachabilityForSmsConfiguration
		return ret
	}
	return *o.ReachabilityForSmsCfg
}

// GetReachabilityForSmsCfgOk returns a tuple with the ReachabilityForSmsCfg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetReachabilityForSmsCfgOk() (*ReachabilityForSmsConfiguration, bool) {
	if o == nil || isNil(o.ReachabilityForSmsCfg) {
    return nil, false
	}
	return o.ReachabilityForSmsCfg, true
}

// HasReachabilityForSmsCfg returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasReachabilityForSmsCfg() bool {
	if o != nil && !isNil(o.ReachabilityForSmsCfg) {
		return true
	}

	return false
}

// SetReachabilityForSmsCfg gets a reference to the given ReachabilityForSmsConfiguration and assigns it to the ReachabilityForSmsCfg field.
func (o *MonitoringConfiguration) SetReachabilityForSmsCfg(v ReachabilityForSmsConfiguration) {
	o.ReachabilityForSmsCfg = &v
}

// GetMtcProviderInformation returns the MtcProviderInformation field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetMtcProviderInformation() string {
	if o == nil || isNil(o.MtcProviderInformation) {
		var ret string
		return ret
	}
	return *o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil || isNil(o.MtcProviderInformation) {
    return nil, false
	}
	return o.MtcProviderInformation, true
}

// HasMtcProviderInformation returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasMtcProviderInformation() bool {
	if o != nil && !isNil(o.MtcProviderInformation) {
		return true
	}

	return false
}

// SetMtcProviderInformation gets a reference to the given string and assigns it to the MtcProviderInformation field.
func (o *MonitoringConfiguration) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = &v
}

// GetAfId returns the AfId field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetAfId() string {
	if o == nil || isNil(o.AfId) {
		var ret string
		return ret
	}
	return *o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetAfIdOk() (*string, bool) {
	if o == nil || isNil(o.AfId) {
    return nil, false
	}
	return o.AfId, true
}

// HasAfId returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasAfId() bool {
	if o != nil && !isNil(o.AfId) {
		return true
	}

	return false
}

// SetAfId gets a reference to the given string and assigns it to the AfId field.
func (o *MonitoringConfiguration) SetAfId(v string) {
	o.AfId = &v
}

// GetReachabilityForDataCfg returns the ReachabilityForDataCfg field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetReachabilityForDataCfg() ReachabilityForDataConfiguration {
	if o == nil || isNil(o.ReachabilityForDataCfg) {
		var ret ReachabilityForDataConfiguration
		return ret
	}
	return *o.ReachabilityForDataCfg
}

// GetReachabilityForDataCfgOk returns a tuple with the ReachabilityForDataCfg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetReachabilityForDataCfgOk() (*ReachabilityForDataConfiguration, bool) {
	if o == nil || isNil(o.ReachabilityForDataCfg) {
    return nil, false
	}
	return o.ReachabilityForDataCfg, true
}

// HasReachabilityForDataCfg returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasReachabilityForDataCfg() bool {
	if o != nil && !isNil(o.ReachabilityForDataCfg) {
		return true
	}

	return false
}

// SetReachabilityForDataCfg gets a reference to the given ReachabilityForDataConfiguration and assigns it to the ReachabilityForDataCfg field.
func (o *MonitoringConfiguration) SetReachabilityForDataCfg(v ReachabilityForDataConfiguration) {
	o.ReachabilityForDataCfg = &v
}

// GetIdleStatusInd returns the IdleStatusInd field value if set, zero value otherwise.
func (o *MonitoringConfiguration) GetIdleStatusInd() bool {
	if o == nil || isNil(o.IdleStatusInd) {
		var ret bool
		return ret
	}
	return *o.IdleStatusInd
}

// GetIdleStatusIndOk returns a tuple with the IdleStatusInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfiguration) GetIdleStatusIndOk() (*bool, bool) {
	if o == nil || isNil(o.IdleStatusInd) {
    return nil, false
	}
	return o.IdleStatusInd, true
}

// HasIdleStatusInd returns a boolean if a field has been set.
func (o *MonitoringConfiguration) HasIdleStatusInd() bool {
	if o != nil && !isNil(o.IdleStatusInd) {
		return true
	}

	return false
}

// SetIdleStatusInd gets a reference to the given bool and assigns it to the IdleStatusInd field.
func (o *MonitoringConfiguration) SetIdleStatusInd(v bool) {
	o.IdleStatusInd = &v
}

func (o MonitoringConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if !isNil(o.ImmediateFlag) {
		toSerialize["immediateFlag"] = o.ImmediateFlag
	}
	if !isNil(o.LocationReportingConfiguration) {
		toSerialize["locationReportingConfiguration"] = o.LocationReportingConfiguration
	}
	if !isNil(o.AssociationType) {
		toSerialize["associationType"] = o.AssociationType
	}
	if !isNil(o.DatalinkReportCfg) {
		toSerialize["datalinkReportCfg"] = o.DatalinkReportCfg
	}
	if !isNil(o.LossConnectivityCfg) {
		toSerialize["lossConnectivityCfg"] = o.LossConnectivityCfg
	}
	if !isNil(o.MaximumLatency) {
		toSerialize["maximumLatency"] = o.MaximumLatency
	}
	if !isNil(o.MaximumResponseTime) {
		toSerialize["maximumResponseTime"] = o.MaximumResponseTime
	}
	if !isNil(o.SuggestedPacketNumDl) {
		toSerialize["suggestedPacketNumDl"] = o.SuggestedPacketNumDl
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.SingleNssai) {
		toSerialize["singleNssai"] = o.SingleNssai
	}
	if !isNil(o.PduSessionStatusCfg) {
		toSerialize["pduSessionStatusCfg"] = o.PduSessionStatusCfg
	}
	if !isNil(o.ReachabilityForSmsCfg) {
		toSerialize["reachabilityForSmsCfg"] = o.ReachabilityForSmsCfg
	}
	if !isNil(o.MtcProviderInformation) {
		toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	}
	if !isNil(o.AfId) {
		toSerialize["afId"] = o.AfId
	}
	if !isNil(o.ReachabilityForDataCfg) {
		toSerialize["reachabilityForDataCfg"] = o.ReachabilityForDataCfg
	}
	if !isNil(o.IdleStatusInd) {
		toSerialize["idleStatusInd"] = o.IdleStatusInd
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringConfiguration struct {
	value *MonitoringConfiguration
	isSet bool
}

func (v NullableMonitoringConfiguration) Get() *MonitoringConfiguration {
	return v.value
}

func (v *NullableMonitoringConfiguration) Set(val *MonitoringConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringConfiguration(val *MonitoringConfiguration) *NullableMonitoringConfiguration {
	return &NullableMonitoringConfiguration{value: val, isSet: true}
}

func (v NullableMonitoringConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


