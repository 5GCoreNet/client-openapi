/*
Npcf_AMPolicyControl

Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_AMPolicyControl

import (
	"encoding/json"
)

// PolicyUpdate Represents updated policies that the PCF provides in a notification or in a reply to an Update Request. 
type PolicyUpdate struct {
	// String providing an URI formatted according to RFC 3986.
	ResourceUri string `json:"resourceUri"`
	// Request Triggers that the PCF subscribes.
	Triggers []RequestTrigger `json:"triggers,omitempty"`
	ServAreaRes *ServiceAreaRestriction `json:"servAreaRes,omitempty"`
	WlServAreaRes *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`
	// Unsigned integer representing the \"Subscriber Profile ID for RAT/Frequency Priority\"  as specified in 3GPP TS 36.413. 
	Rfsp *int32 `json:"rfsp,omitempty"`
	// Unsigned integer representing the \"Subscriber Profile ID for RAT/Frequency Priority\"  as specified in 3GPP TS 36.413. 
	TargetRfsp *int32 `json:"targetRfsp,omitempty"`
	SmfSelInfo NullableSmfSelectionData `json:"smfSelInfo,omitempty"`
	UeAmbr *Ambr `json:"ueAmbr,omitempty"`
	// One or more UE-Slice-MBR(s) for S-NSSAI(s) of serving PLMN the allowed NSSAI as part of the AMF Access and Mobility Policy as determined by the PCF. 
	UeSliceMbrs []UeSliceMbr `json:"ueSliceMbrs,omitempty"`
	// Contains the presence reporting area(s) for which reporting was requested. The praId attribute within the PresenceInfo data type is the key of the map. 
	Pras map[string]PresenceInfoRm `json:"pras,omitempty"`
	PcfUeInfo NullablePcfUeCallbackInfo `json:"pcfUeInfo,omitempty"`
	MatchPdus []PduSessionInfo `json:"matchPdus,omitempty"`
	AsTimeDisParam NullableAsTimeDistributionParam `json:"asTimeDisParam,omitempty"`
}

// NewPolicyUpdate instantiates a new PolicyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyUpdate(resourceUri string) *PolicyUpdate {
	this := PolicyUpdate{}
	this.ResourceUri = resourceUri
	return &this
}

// NewPolicyUpdateWithDefaults instantiates a new PolicyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyUpdateWithDefaults() *PolicyUpdate {
	this := PolicyUpdate{}
	return &this
}

// GetResourceUri returns the ResourceUri field value
func (o *PolicyUpdate) GetResourceUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceUri
}

// GetResourceUriOk returns a tuple with the ResourceUri field value
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetResourceUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ResourceUri, true
}

// SetResourceUri sets field value
func (o *PolicyUpdate) SetResourceUri(v string) {
	o.ResourceUri = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyUpdate) GetTriggers() []RequestTrigger {
	if o == nil {
		var ret []RequestTrigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyUpdate) GetTriggersOk() ([]RequestTrigger, bool) {
	if o == nil || isNil(o.Triggers) {
    return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *PolicyUpdate) HasTriggers() bool {
	if o != nil && isNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []RequestTrigger and assigns it to the Triggers field.
func (o *PolicyUpdate) SetTriggers(v []RequestTrigger) {
	o.Triggers = v
}

// GetServAreaRes returns the ServAreaRes field value if set, zero value otherwise.
func (o *PolicyUpdate) GetServAreaRes() ServiceAreaRestriction {
	if o == nil || isNil(o.ServAreaRes) {
		var ret ServiceAreaRestriction
		return ret
	}
	return *o.ServAreaRes
}

// GetServAreaResOk returns a tuple with the ServAreaRes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetServAreaResOk() (*ServiceAreaRestriction, bool) {
	if o == nil || isNil(o.ServAreaRes) {
    return nil, false
	}
	return o.ServAreaRes, true
}

// HasServAreaRes returns a boolean if a field has been set.
func (o *PolicyUpdate) HasServAreaRes() bool {
	if o != nil && !isNil(o.ServAreaRes) {
		return true
	}

	return false
}

// SetServAreaRes gets a reference to the given ServiceAreaRestriction and assigns it to the ServAreaRes field.
func (o *PolicyUpdate) SetServAreaRes(v ServiceAreaRestriction) {
	o.ServAreaRes = &v
}

// GetWlServAreaRes returns the WlServAreaRes field value if set, zero value otherwise.
func (o *PolicyUpdate) GetWlServAreaRes() WirelineServiceAreaRestriction {
	if o == nil || isNil(o.WlServAreaRes) {
		var ret WirelineServiceAreaRestriction
		return ret
	}
	return *o.WlServAreaRes
}

// GetWlServAreaResOk returns a tuple with the WlServAreaRes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetWlServAreaResOk() (*WirelineServiceAreaRestriction, bool) {
	if o == nil || isNil(o.WlServAreaRes) {
    return nil, false
	}
	return o.WlServAreaRes, true
}

// HasWlServAreaRes returns a boolean if a field has been set.
func (o *PolicyUpdate) HasWlServAreaRes() bool {
	if o != nil && !isNil(o.WlServAreaRes) {
		return true
	}

	return false
}

// SetWlServAreaRes gets a reference to the given WirelineServiceAreaRestriction and assigns it to the WlServAreaRes field.
func (o *PolicyUpdate) SetWlServAreaRes(v WirelineServiceAreaRestriction) {
	o.WlServAreaRes = &v
}

// GetRfsp returns the Rfsp field value if set, zero value otherwise.
func (o *PolicyUpdate) GetRfsp() int32 {
	if o == nil || isNil(o.Rfsp) {
		var ret int32
		return ret
	}
	return *o.Rfsp
}

// GetRfspOk returns a tuple with the Rfsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetRfspOk() (*int32, bool) {
	if o == nil || isNil(o.Rfsp) {
    return nil, false
	}
	return o.Rfsp, true
}

// HasRfsp returns a boolean if a field has been set.
func (o *PolicyUpdate) HasRfsp() bool {
	if o != nil && !isNil(o.Rfsp) {
		return true
	}

	return false
}

// SetRfsp gets a reference to the given int32 and assigns it to the Rfsp field.
func (o *PolicyUpdate) SetRfsp(v int32) {
	o.Rfsp = &v
}

// GetTargetRfsp returns the TargetRfsp field value if set, zero value otherwise.
func (o *PolicyUpdate) GetTargetRfsp() int32 {
	if o == nil || isNil(o.TargetRfsp) {
		var ret int32
		return ret
	}
	return *o.TargetRfsp
}

// GetTargetRfspOk returns a tuple with the TargetRfsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetTargetRfspOk() (*int32, bool) {
	if o == nil || isNil(o.TargetRfsp) {
    return nil, false
	}
	return o.TargetRfsp, true
}

// HasTargetRfsp returns a boolean if a field has been set.
func (o *PolicyUpdate) HasTargetRfsp() bool {
	if o != nil && !isNil(o.TargetRfsp) {
		return true
	}

	return false
}

// SetTargetRfsp gets a reference to the given int32 and assigns it to the TargetRfsp field.
func (o *PolicyUpdate) SetTargetRfsp(v int32) {
	o.TargetRfsp = &v
}

// GetSmfSelInfo returns the SmfSelInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyUpdate) GetSmfSelInfo() SmfSelectionData {
	if o == nil || isNil(o.SmfSelInfo.Get()) {
		var ret SmfSelectionData
		return ret
	}
	return *o.SmfSelInfo.Get()
}

// GetSmfSelInfoOk returns a tuple with the SmfSelInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyUpdate) GetSmfSelInfoOk() (*SmfSelectionData, bool) {
	if o == nil {
    return nil, false
	}
	return o.SmfSelInfo.Get(), o.SmfSelInfo.IsSet()
}

// HasSmfSelInfo returns a boolean if a field has been set.
func (o *PolicyUpdate) HasSmfSelInfo() bool {
	if o != nil && o.SmfSelInfo.IsSet() {
		return true
	}

	return false
}

// SetSmfSelInfo gets a reference to the given NullableSmfSelectionData and assigns it to the SmfSelInfo field.
func (o *PolicyUpdate) SetSmfSelInfo(v SmfSelectionData) {
	o.SmfSelInfo.Set(&v)
}
// SetSmfSelInfoNil sets the value for SmfSelInfo to be an explicit nil
func (o *PolicyUpdate) SetSmfSelInfoNil() {
	o.SmfSelInfo.Set(nil)
}

// UnsetSmfSelInfo ensures that no value is present for SmfSelInfo, not even an explicit nil
func (o *PolicyUpdate) UnsetSmfSelInfo() {
	o.SmfSelInfo.Unset()
}

// GetUeAmbr returns the UeAmbr field value if set, zero value otherwise.
func (o *PolicyUpdate) GetUeAmbr() Ambr {
	if o == nil || isNil(o.UeAmbr) {
		var ret Ambr
		return ret
	}
	return *o.UeAmbr
}

// GetUeAmbrOk returns a tuple with the UeAmbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetUeAmbrOk() (*Ambr, bool) {
	if o == nil || isNil(o.UeAmbr) {
    return nil, false
	}
	return o.UeAmbr, true
}

// HasUeAmbr returns a boolean if a field has been set.
func (o *PolicyUpdate) HasUeAmbr() bool {
	if o != nil && !isNil(o.UeAmbr) {
		return true
	}

	return false
}

// SetUeAmbr gets a reference to the given Ambr and assigns it to the UeAmbr field.
func (o *PolicyUpdate) SetUeAmbr(v Ambr) {
	o.UeAmbr = &v
}

// GetUeSliceMbrs returns the UeSliceMbrs field value if set, zero value otherwise.
func (o *PolicyUpdate) GetUeSliceMbrs() []UeSliceMbr {
	if o == nil || isNil(o.UeSliceMbrs) {
		var ret []UeSliceMbr
		return ret
	}
	return o.UeSliceMbrs
}

// GetUeSliceMbrsOk returns a tuple with the UeSliceMbrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUpdate) GetUeSliceMbrsOk() ([]UeSliceMbr, bool) {
	if o == nil || isNil(o.UeSliceMbrs) {
    return nil, false
	}
	return o.UeSliceMbrs, true
}

// HasUeSliceMbrs returns a boolean if a field has been set.
func (o *PolicyUpdate) HasUeSliceMbrs() bool {
	if o != nil && !isNil(o.UeSliceMbrs) {
		return true
	}

	return false
}

// SetUeSliceMbrs gets a reference to the given []UeSliceMbr and assigns it to the UeSliceMbrs field.
func (o *PolicyUpdate) SetUeSliceMbrs(v []UeSliceMbr) {
	o.UeSliceMbrs = v
}

// GetPras returns the Pras field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyUpdate) GetPras() map[string]PresenceInfoRm {
	if o == nil {
		var ret map[string]PresenceInfoRm
		return ret
	}
	return o.Pras
}

// GetPrasOk returns a tuple with the Pras field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyUpdate) GetPrasOk() (*map[string]PresenceInfoRm, bool) {
	if o == nil || isNil(o.Pras) {
    return nil, false
	}
	return &o.Pras, true
}

// HasPras returns a boolean if a field has been set.
func (o *PolicyUpdate) HasPras() bool {
	if o != nil && isNil(o.Pras) {
		return true
	}

	return false
}

// SetPras gets a reference to the given map[string]PresenceInfoRm and assigns it to the Pras field.
func (o *PolicyUpdate) SetPras(v map[string]PresenceInfoRm) {
	o.Pras = v
}

// GetPcfUeInfo returns the PcfUeInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyUpdate) GetPcfUeInfo() PcfUeCallbackInfo {
	if o == nil || isNil(o.PcfUeInfo.Get()) {
		var ret PcfUeCallbackInfo
		return ret
	}
	return *o.PcfUeInfo.Get()
}

// GetPcfUeInfoOk returns a tuple with the PcfUeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyUpdate) GetPcfUeInfoOk() (*PcfUeCallbackInfo, bool) {
	if o == nil {
    return nil, false
	}
	return o.PcfUeInfo.Get(), o.PcfUeInfo.IsSet()
}

// HasPcfUeInfo returns a boolean if a field has been set.
func (o *PolicyUpdate) HasPcfUeInfo() bool {
	if o != nil && o.PcfUeInfo.IsSet() {
		return true
	}

	return false
}

// SetPcfUeInfo gets a reference to the given NullablePcfUeCallbackInfo and assigns it to the PcfUeInfo field.
func (o *PolicyUpdate) SetPcfUeInfo(v PcfUeCallbackInfo) {
	o.PcfUeInfo.Set(&v)
}
// SetPcfUeInfoNil sets the value for PcfUeInfo to be an explicit nil
func (o *PolicyUpdate) SetPcfUeInfoNil() {
	o.PcfUeInfo.Set(nil)
}

// UnsetPcfUeInfo ensures that no value is present for PcfUeInfo, not even an explicit nil
func (o *PolicyUpdate) UnsetPcfUeInfo() {
	o.PcfUeInfo.Unset()
}

// GetMatchPdus returns the MatchPdus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyUpdate) GetMatchPdus() []PduSessionInfo {
	if o == nil {
		var ret []PduSessionInfo
		return ret
	}
	return o.MatchPdus
}

// GetMatchPdusOk returns a tuple with the MatchPdus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyUpdate) GetMatchPdusOk() ([]PduSessionInfo, bool) {
	if o == nil || isNil(o.MatchPdus) {
    return nil, false
	}
	return o.MatchPdus, true
}

// HasMatchPdus returns a boolean if a field has been set.
func (o *PolicyUpdate) HasMatchPdus() bool {
	if o != nil && isNil(o.MatchPdus) {
		return true
	}

	return false
}

// SetMatchPdus gets a reference to the given []PduSessionInfo and assigns it to the MatchPdus field.
func (o *PolicyUpdate) SetMatchPdus(v []PduSessionInfo) {
	o.MatchPdus = v
}

// GetAsTimeDisParam returns the AsTimeDisParam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyUpdate) GetAsTimeDisParam() AsTimeDistributionParam {
	if o == nil || isNil(o.AsTimeDisParam.Get()) {
		var ret AsTimeDistributionParam
		return ret
	}
	return *o.AsTimeDisParam.Get()
}

// GetAsTimeDisParamOk returns a tuple with the AsTimeDisParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyUpdate) GetAsTimeDisParamOk() (*AsTimeDistributionParam, bool) {
	if o == nil {
    return nil, false
	}
	return o.AsTimeDisParam.Get(), o.AsTimeDisParam.IsSet()
}

// HasAsTimeDisParam returns a boolean if a field has been set.
func (o *PolicyUpdate) HasAsTimeDisParam() bool {
	if o != nil && o.AsTimeDisParam.IsSet() {
		return true
	}

	return false
}

// SetAsTimeDisParam gets a reference to the given NullableAsTimeDistributionParam and assigns it to the AsTimeDisParam field.
func (o *PolicyUpdate) SetAsTimeDisParam(v AsTimeDistributionParam) {
	o.AsTimeDisParam.Set(&v)
}
// SetAsTimeDisParamNil sets the value for AsTimeDisParam to be an explicit nil
func (o *PolicyUpdate) SetAsTimeDisParamNil() {
	o.AsTimeDisParam.Set(nil)
}

// UnsetAsTimeDisParam ensures that no value is present for AsTimeDisParam, not even an explicit nil
func (o *PolicyUpdate) UnsetAsTimeDisParam() {
	o.AsTimeDisParam.Unset()
}

func (o PolicyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceUri"] = o.ResourceUri
	}
	if o.Triggers != nil {
		toSerialize["triggers"] = o.Triggers
	}
	if !isNil(o.ServAreaRes) {
		toSerialize["servAreaRes"] = o.ServAreaRes
	}
	if !isNil(o.WlServAreaRes) {
		toSerialize["wlServAreaRes"] = o.WlServAreaRes
	}
	if !isNil(o.Rfsp) {
		toSerialize["rfsp"] = o.Rfsp
	}
	if !isNil(o.TargetRfsp) {
		toSerialize["targetRfsp"] = o.TargetRfsp
	}
	if o.SmfSelInfo.IsSet() {
		toSerialize["smfSelInfo"] = o.SmfSelInfo.Get()
	}
	if !isNil(o.UeAmbr) {
		toSerialize["ueAmbr"] = o.UeAmbr
	}
	if !isNil(o.UeSliceMbrs) {
		toSerialize["ueSliceMbrs"] = o.UeSliceMbrs
	}
	if o.Pras != nil {
		toSerialize["pras"] = o.Pras
	}
	if o.PcfUeInfo.IsSet() {
		toSerialize["pcfUeInfo"] = o.PcfUeInfo.Get()
	}
	if o.MatchPdus != nil {
		toSerialize["matchPdus"] = o.MatchPdus
	}
	if o.AsTimeDisParam.IsSet() {
		toSerialize["asTimeDisParam"] = o.AsTimeDisParam.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyUpdate struct {
	value *PolicyUpdate
	isSet bool
}

func (v NullablePolicyUpdate) Get() *PolicyUpdate {
	return v.value
}

func (v *NullablePolicyUpdate) Set(val *PolicyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyUpdate(val *PolicyUpdate) *NullablePolicyUpdate {
	return &NullablePolicyUpdate{value: val, isSet: true}
}

func (v NullablePolicyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


