/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MBSUserDataIngestSession

import (
	"encoding/json"
)

// MBSDistributionSessionInfo Represents MBS Distribution Session information.
type MBSDistributionSessionInfo struct {
	MbsDistSessionId *string `json:"mbsDistSessionId,omitempty"`
	MbsSessionId *MbsSessionId `json:"mbsSessionId,omitempty"`
	MbsServInfo *MbsServiceInfo `json:"mbsServInfo,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxContBitRate string `json:"maxContBitRate"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	MaxContDelay *int32 `json:"maxContDelay,omitempty"`
	DistrMethod DistributionMethod `json:"distrMethod"`
	FecConfig *FECConfig `json:"fecConfig,omitempty"`
	ObjDistrInfo *ObjectDistrMethInfo `json:"objDistrInfo,omitempty"`
	PckDistrInfo *PacketDistrMethInfo `json:"pckDistrInfo,omitempty"`
	TrafficMarkingInfo *string `json:"trafficMarkingInfo,omitempty"`
	MbsDistSessState *DistSessionState `json:"mbsDistSessState,omitempty"`
	TgtServAreas *MbsServiceArea `json:"tgtServAreas,omitempty"`
	ExtTgtServAreas *ExternalMbsServiceArea `json:"extTgtServAreas,omitempty"`
	// MBS Frequency Selection Area Identifier
	MbsFSAId *string `json:"mbsFSAId,omitempty"`
	// Represents an indication that this MBS Distribution Session belongs to a location- dependent MBS. This attribute shall be set to \"true\" to indicate that the MBS  Distribution Session belongs to a location-dependent MBS; or set to \"false\" to  indicate that the MBS Distribution Session does not belong to a location-dependent MBS. The default value is \"false\", if omitted. 
	LocationDependent *bool `json:"locationDependent,omitempty"`
	// Represents an indication that this MBS Distribution Session belongs to a multiplex, i.e.  forms part of a set of MBS Distribution Sessions under the same parent MBS User Data  Ingest Session with identical or empty sets of target service areas and multiplexed onto  the same MBS Session at the MB-SMF. 
	MultiplexedServFlag *bool `json:"multiplexedServFlag,omitempty"`
	// Represents an indication that this MBS Distribution Session is not open to any UE, i.e.  restricted to a set of UEs according to their MBS related subscription information. This attribute may be included only if the parent MBS User Service is of Multicast service type. This attribute shall be set to \"true\" to indicate that this MBS Distribution Session is restricted to a set of UE(s); or set to \"false\" to indicate that this MBS Distribution Session is open to any UE. The default value is \"false\", if omitted. 
	RestrictedFlag *bool `json:"restrictedFlag,omitempty"`
}

// NewMBSDistributionSessionInfo instantiates a new MBSDistributionSessionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMBSDistributionSessionInfo(maxContBitRate string, distrMethod DistributionMethod) *MBSDistributionSessionInfo {
	this := MBSDistributionSessionInfo{}
	this.MaxContBitRate = maxContBitRate
	this.DistrMethod = distrMethod
	var locationDependent bool = false
	this.LocationDependent = &locationDependent
	var multiplexedServFlag bool = false
	this.MultiplexedServFlag = &multiplexedServFlag
	var restrictedFlag bool = false
	this.RestrictedFlag = &restrictedFlag
	return &this
}

// NewMBSDistributionSessionInfoWithDefaults instantiates a new MBSDistributionSessionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMBSDistributionSessionInfoWithDefaults() *MBSDistributionSessionInfo {
	this := MBSDistributionSessionInfo{}
	var locationDependent bool = false
	this.LocationDependent = &locationDependent
	var multiplexedServFlag bool = false
	this.MultiplexedServFlag = &multiplexedServFlag
	var restrictedFlag bool = false
	this.RestrictedFlag = &restrictedFlag
	return &this
}

// GetMbsDistSessionId returns the MbsDistSessionId field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetMbsDistSessionId() string {
	if o == nil || isNil(o.MbsDistSessionId) {
		var ret string
		return ret
	}
	return *o.MbsDistSessionId
}

// GetMbsDistSessionIdOk returns a tuple with the MbsDistSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetMbsDistSessionIdOk() (*string, bool) {
	if o == nil || isNil(o.MbsDistSessionId) {
    return nil, false
	}
	return o.MbsDistSessionId, true
}

// HasMbsDistSessionId returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasMbsDistSessionId() bool {
	if o != nil && !isNil(o.MbsDistSessionId) {
		return true
	}

	return false
}

// SetMbsDistSessionId gets a reference to the given string and assigns it to the MbsDistSessionId field.
func (o *MBSDistributionSessionInfo) SetMbsDistSessionId(v string) {
	o.MbsDistSessionId = &v
}

// GetMbsSessionId returns the MbsSessionId field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetMbsSessionId() MbsSessionId {
	if o == nil || isNil(o.MbsSessionId) {
		var ret MbsSessionId
		return ret
	}
	return *o.MbsSessionId
}

// GetMbsSessionIdOk returns a tuple with the MbsSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetMbsSessionIdOk() (*MbsSessionId, bool) {
	if o == nil || isNil(o.MbsSessionId) {
    return nil, false
	}
	return o.MbsSessionId, true
}

// HasMbsSessionId returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasMbsSessionId() bool {
	if o != nil && !isNil(o.MbsSessionId) {
		return true
	}

	return false
}

// SetMbsSessionId gets a reference to the given MbsSessionId and assigns it to the MbsSessionId field.
func (o *MBSDistributionSessionInfo) SetMbsSessionId(v MbsSessionId) {
	o.MbsSessionId = &v
}

// GetMbsServInfo returns the MbsServInfo field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetMbsServInfo() MbsServiceInfo {
	if o == nil || isNil(o.MbsServInfo) {
		var ret MbsServiceInfo
		return ret
	}
	return *o.MbsServInfo
}

// GetMbsServInfoOk returns a tuple with the MbsServInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetMbsServInfoOk() (*MbsServiceInfo, bool) {
	if o == nil || isNil(o.MbsServInfo) {
    return nil, false
	}
	return o.MbsServInfo, true
}

// HasMbsServInfo returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasMbsServInfo() bool {
	if o != nil && !isNil(o.MbsServInfo) {
		return true
	}

	return false
}

// SetMbsServInfo gets a reference to the given MbsServiceInfo and assigns it to the MbsServInfo field.
func (o *MBSDistributionSessionInfo) SetMbsServInfo(v MbsServiceInfo) {
	o.MbsServInfo = &v
}

// GetMaxContBitRate returns the MaxContBitRate field value
func (o *MBSDistributionSessionInfo) GetMaxContBitRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxContBitRate
}

// GetMaxContBitRateOk returns a tuple with the MaxContBitRate field value
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetMaxContBitRateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MaxContBitRate, true
}

// SetMaxContBitRate sets field value
func (o *MBSDistributionSessionInfo) SetMaxContBitRate(v string) {
	o.MaxContBitRate = v
}

// GetMaxContDelay returns the MaxContDelay field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetMaxContDelay() int32 {
	if o == nil || isNil(o.MaxContDelay) {
		var ret int32
		return ret
	}
	return *o.MaxContDelay
}

// GetMaxContDelayOk returns a tuple with the MaxContDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetMaxContDelayOk() (*int32, bool) {
	if o == nil || isNil(o.MaxContDelay) {
    return nil, false
	}
	return o.MaxContDelay, true
}

// HasMaxContDelay returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasMaxContDelay() bool {
	if o != nil && !isNil(o.MaxContDelay) {
		return true
	}

	return false
}

// SetMaxContDelay gets a reference to the given int32 and assigns it to the MaxContDelay field.
func (o *MBSDistributionSessionInfo) SetMaxContDelay(v int32) {
	o.MaxContDelay = &v
}

// GetDistrMethod returns the DistrMethod field value
func (o *MBSDistributionSessionInfo) GetDistrMethod() DistributionMethod {
	if o == nil {
		var ret DistributionMethod
		return ret
	}

	return o.DistrMethod
}

// GetDistrMethodOk returns a tuple with the DistrMethod field value
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetDistrMethodOk() (*DistributionMethod, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DistrMethod, true
}

// SetDistrMethod sets field value
func (o *MBSDistributionSessionInfo) SetDistrMethod(v DistributionMethod) {
	o.DistrMethod = v
}

// GetFecConfig returns the FecConfig field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetFecConfig() FECConfig {
	if o == nil || isNil(o.FecConfig) {
		var ret FECConfig
		return ret
	}
	return *o.FecConfig
}

// GetFecConfigOk returns a tuple with the FecConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetFecConfigOk() (*FECConfig, bool) {
	if o == nil || isNil(o.FecConfig) {
    return nil, false
	}
	return o.FecConfig, true
}

// HasFecConfig returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasFecConfig() bool {
	if o != nil && !isNil(o.FecConfig) {
		return true
	}

	return false
}

// SetFecConfig gets a reference to the given FECConfig and assigns it to the FecConfig field.
func (o *MBSDistributionSessionInfo) SetFecConfig(v FECConfig) {
	o.FecConfig = &v
}

// GetObjDistrInfo returns the ObjDistrInfo field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetObjDistrInfo() ObjectDistrMethInfo {
	if o == nil || isNil(o.ObjDistrInfo) {
		var ret ObjectDistrMethInfo
		return ret
	}
	return *o.ObjDistrInfo
}

// GetObjDistrInfoOk returns a tuple with the ObjDistrInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetObjDistrInfoOk() (*ObjectDistrMethInfo, bool) {
	if o == nil || isNil(o.ObjDistrInfo) {
    return nil, false
	}
	return o.ObjDistrInfo, true
}

// HasObjDistrInfo returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasObjDistrInfo() bool {
	if o != nil && !isNil(o.ObjDistrInfo) {
		return true
	}

	return false
}

// SetObjDistrInfo gets a reference to the given ObjectDistrMethInfo and assigns it to the ObjDistrInfo field.
func (o *MBSDistributionSessionInfo) SetObjDistrInfo(v ObjectDistrMethInfo) {
	o.ObjDistrInfo = &v
}

// GetPckDistrInfo returns the PckDistrInfo field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetPckDistrInfo() PacketDistrMethInfo {
	if o == nil || isNil(o.PckDistrInfo) {
		var ret PacketDistrMethInfo
		return ret
	}
	return *o.PckDistrInfo
}

// GetPckDistrInfoOk returns a tuple with the PckDistrInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetPckDistrInfoOk() (*PacketDistrMethInfo, bool) {
	if o == nil || isNil(o.PckDistrInfo) {
    return nil, false
	}
	return o.PckDistrInfo, true
}

// HasPckDistrInfo returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasPckDistrInfo() bool {
	if o != nil && !isNil(o.PckDistrInfo) {
		return true
	}

	return false
}

// SetPckDistrInfo gets a reference to the given PacketDistrMethInfo and assigns it to the PckDistrInfo field.
func (o *MBSDistributionSessionInfo) SetPckDistrInfo(v PacketDistrMethInfo) {
	o.PckDistrInfo = &v
}

// GetTrafficMarkingInfo returns the TrafficMarkingInfo field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetTrafficMarkingInfo() string {
	if o == nil || isNil(o.TrafficMarkingInfo) {
		var ret string
		return ret
	}
	return *o.TrafficMarkingInfo
}

// GetTrafficMarkingInfoOk returns a tuple with the TrafficMarkingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetTrafficMarkingInfoOk() (*string, bool) {
	if o == nil || isNil(o.TrafficMarkingInfo) {
    return nil, false
	}
	return o.TrafficMarkingInfo, true
}

// HasTrafficMarkingInfo returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasTrafficMarkingInfo() bool {
	if o != nil && !isNil(o.TrafficMarkingInfo) {
		return true
	}

	return false
}

// SetTrafficMarkingInfo gets a reference to the given string and assigns it to the TrafficMarkingInfo field.
func (o *MBSDistributionSessionInfo) SetTrafficMarkingInfo(v string) {
	o.TrafficMarkingInfo = &v
}

// GetMbsDistSessState returns the MbsDistSessState field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetMbsDistSessState() DistSessionState {
	if o == nil || isNil(o.MbsDistSessState) {
		var ret DistSessionState
		return ret
	}
	return *o.MbsDistSessState
}

// GetMbsDistSessStateOk returns a tuple with the MbsDistSessState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetMbsDistSessStateOk() (*DistSessionState, bool) {
	if o == nil || isNil(o.MbsDistSessState) {
    return nil, false
	}
	return o.MbsDistSessState, true
}

// HasMbsDistSessState returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasMbsDistSessState() bool {
	if o != nil && !isNil(o.MbsDistSessState) {
		return true
	}

	return false
}

// SetMbsDistSessState gets a reference to the given DistSessionState and assigns it to the MbsDistSessState field.
func (o *MBSDistributionSessionInfo) SetMbsDistSessState(v DistSessionState) {
	o.MbsDistSessState = &v
}

// GetTgtServAreas returns the TgtServAreas field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetTgtServAreas() MbsServiceArea {
	if o == nil || isNil(o.TgtServAreas) {
		var ret MbsServiceArea
		return ret
	}
	return *o.TgtServAreas
}

// GetTgtServAreasOk returns a tuple with the TgtServAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetTgtServAreasOk() (*MbsServiceArea, bool) {
	if o == nil || isNil(o.TgtServAreas) {
    return nil, false
	}
	return o.TgtServAreas, true
}

// HasTgtServAreas returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasTgtServAreas() bool {
	if o != nil && !isNil(o.TgtServAreas) {
		return true
	}

	return false
}

// SetTgtServAreas gets a reference to the given MbsServiceArea and assigns it to the TgtServAreas field.
func (o *MBSDistributionSessionInfo) SetTgtServAreas(v MbsServiceArea) {
	o.TgtServAreas = &v
}

// GetExtTgtServAreas returns the ExtTgtServAreas field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetExtTgtServAreas() ExternalMbsServiceArea {
	if o == nil || isNil(o.ExtTgtServAreas) {
		var ret ExternalMbsServiceArea
		return ret
	}
	return *o.ExtTgtServAreas
}

// GetExtTgtServAreasOk returns a tuple with the ExtTgtServAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetExtTgtServAreasOk() (*ExternalMbsServiceArea, bool) {
	if o == nil || isNil(o.ExtTgtServAreas) {
    return nil, false
	}
	return o.ExtTgtServAreas, true
}

// HasExtTgtServAreas returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasExtTgtServAreas() bool {
	if o != nil && !isNil(o.ExtTgtServAreas) {
		return true
	}

	return false
}

// SetExtTgtServAreas gets a reference to the given ExternalMbsServiceArea and assigns it to the ExtTgtServAreas field.
func (o *MBSDistributionSessionInfo) SetExtTgtServAreas(v ExternalMbsServiceArea) {
	o.ExtTgtServAreas = &v
}

// GetMbsFSAId returns the MbsFSAId field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetMbsFSAId() string {
	if o == nil || isNil(o.MbsFSAId) {
		var ret string
		return ret
	}
	return *o.MbsFSAId
}

// GetMbsFSAIdOk returns a tuple with the MbsFSAId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetMbsFSAIdOk() (*string, bool) {
	if o == nil || isNil(o.MbsFSAId) {
    return nil, false
	}
	return o.MbsFSAId, true
}

// HasMbsFSAId returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasMbsFSAId() bool {
	if o != nil && !isNil(o.MbsFSAId) {
		return true
	}

	return false
}

// SetMbsFSAId gets a reference to the given string and assigns it to the MbsFSAId field.
func (o *MBSDistributionSessionInfo) SetMbsFSAId(v string) {
	o.MbsFSAId = &v
}

// GetLocationDependent returns the LocationDependent field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetLocationDependent() bool {
	if o == nil || isNil(o.LocationDependent) {
		var ret bool
		return ret
	}
	return *o.LocationDependent
}

// GetLocationDependentOk returns a tuple with the LocationDependent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetLocationDependentOk() (*bool, bool) {
	if o == nil || isNil(o.LocationDependent) {
    return nil, false
	}
	return o.LocationDependent, true
}

// HasLocationDependent returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasLocationDependent() bool {
	if o != nil && !isNil(o.LocationDependent) {
		return true
	}

	return false
}

// SetLocationDependent gets a reference to the given bool and assigns it to the LocationDependent field.
func (o *MBSDistributionSessionInfo) SetLocationDependent(v bool) {
	o.LocationDependent = &v
}

// GetMultiplexedServFlag returns the MultiplexedServFlag field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetMultiplexedServFlag() bool {
	if o == nil || isNil(o.MultiplexedServFlag) {
		var ret bool
		return ret
	}
	return *o.MultiplexedServFlag
}

// GetMultiplexedServFlagOk returns a tuple with the MultiplexedServFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetMultiplexedServFlagOk() (*bool, bool) {
	if o == nil || isNil(o.MultiplexedServFlag) {
    return nil, false
	}
	return o.MultiplexedServFlag, true
}

// HasMultiplexedServFlag returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasMultiplexedServFlag() bool {
	if o != nil && !isNil(o.MultiplexedServFlag) {
		return true
	}

	return false
}

// SetMultiplexedServFlag gets a reference to the given bool and assigns it to the MultiplexedServFlag field.
func (o *MBSDistributionSessionInfo) SetMultiplexedServFlag(v bool) {
	o.MultiplexedServFlag = &v
}

// GetRestrictedFlag returns the RestrictedFlag field value if set, zero value otherwise.
func (o *MBSDistributionSessionInfo) GetRestrictedFlag() bool {
	if o == nil || isNil(o.RestrictedFlag) {
		var ret bool
		return ret
	}
	return *o.RestrictedFlag
}

// GetRestrictedFlagOk returns a tuple with the RestrictedFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSDistributionSessionInfo) GetRestrictedFlagOk() (*bool, bool) {
	if o == nil || isNil(o.RestrictedFlag) {
    return nil, false
	}
	return o.RestrictedFlag, true
}

// HasRestrictedFlag returns a boolean if a field has been set.
func (o *MBSDistributionSessionInfo) HasRestrictedFlag() bool {
	if o != nil && !isNil(o.RestrictedFlag) {
		return true
	}

	return false
}

// SetRestrictedFlag gets a reference to the given bool and assigns it to the RestrictedFlag field.
func (o *MBSDistributionSessionInfo) SetRestrictedFlag(v bool) {
	o.RestrictedFlag = &v
}

func (o MBSDistributionSessionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MbsDistSessionId) {
		toSerialize["mbsDistSessionId"] = o.MbsDistSessionId
	}
	if !isNil(o.MbsSessionId) {
		toSerialize["mbsSessionId"] = o.MbsSessionId
	}
	if !isNil(o.MbsServInfo) {
		toSerialize["mbsServInfo"] = o.MbsServInfo
	}
	if true {
		toSerialize["maxContBitRate"] = o.MaxContBitRate
	}
	if !isNil(o.MaxContDelay) {
		toSerialize["maxContDelay"] = o.MaxContDelay
	}
	if true {
		toSerialize["distrMethod"] = o.DistrMethod
	}
	if !isNil(o.FecConfig) {
		toSerialize["fecConfig"] = o.FecConfig
	}
	if !isNil(o.ObjDistrInfo) {
		toSerialize["objDistrInfo"] = o.ObjDistrInfo
	}
	if !isNil(o.PckDistrInfo) {
		toSerialize["pckDistrInfo"] = o.PckDistrInfo
	}
	if !isNil(o.TrafficMarkingInfo) {
		toSerialize["trafficMarkingInfo"] = o.TrafficMarkingInfo
	}
	if !isNil(o.MbsDistSessState) {
		toSerialize["mbsDistSessState"] = o.MbsDistSessState
	}
	if !isNil(o.TgtServAreas) {
		toSerialize["tgtServAreas"] = o.TgtServAreas
	}
	if !isNil(o.ExtTgtServAreas) {
		toSerialize["extTgtServAreas"] = o.ExtTgtServAreas
	}
	if !isNil(o.MbsFSAId) {
		toSerialize["mbsFSAId"] = o.MbsFSAId
	}
	if !isNil(o.LocationDependent) {
		toSerialize["locationDependent"] = o.LocationDependent
	}
	if !isNil(o.MultiplexedServFlag) {
		toSerialize["multiplexedServFlag"] = o.MultiplexedServFlag
	}
	if !isNil(o.RestrictedFlag) {
		toSerialize["restrictedFlag"] = o.RestrictedFlag
	}
	return json.Marshal(toSerialize)
}

type NullableMBSDistributionSessionInfo struct {
	value *MBSDistributionSessionInfo
	isSet bool
}

func (v NullableMBSDistributionSessionInfo) Get() *MBSDistributionSessionInfo {
	return v.value
}

func (v *NullableMBSDistributionSessionInfo) Set(val *MBSDistributionSessionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMBSDistributionSessionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMBSDistributionSessionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMBSDistributionSessionInfo(val *MBSDistributionSessionInfo) *NullableMBSDistributionSessionInfo {
	return &NullableMBSDistributionSessionInfo{value: val, isSet: true}
}

func (v NullableMBSDistributionSessionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMBSDistributionSessionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


