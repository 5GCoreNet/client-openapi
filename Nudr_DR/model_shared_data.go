/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// SharedData struct for SharedData
type SharedData struct {
	SharedDataId string `json:"sharedDataId"`
	SharedAmData *AccessAndMobilitySubscriptionData1 `json:"sharedAmData,omitempty"`
	SharedSmsSubsData *SmsSubscriptionData1 `json:"sharedSmsSubsData,omitempty"`
	SharedSmsMngSubsData *SmsManagementSubscriptionData1 `json:"sharedSmsMngSubsData,omitempty"`
	// A map(list of key-value pairs) where Dnn, or optionally the Wildcard DNN, serves as key of DnnConfiguration
	SharedDnnConfigurations *map[string]DnnConfiguration1 `json:"sharedDnnConfigurations,omitempty"`
	SharedTraceData NullableTraceData `json:"sharedTraceData,omitempty"`
	// A map(list of key-value pairs) where singleNssai serves as key of SnssaiInfo
	SharedSnssaiInfos *map[string]SnssaiInfo `json:"sharedSnssaiInfos,omitempty"`
	// A map(list of key-value pairs) where GroupId serves as key of VnGroupData
	SharedVnGroupDatas *map[string]VnGroupData `json:"sharedVnGroupDatas,omitempty"`
	// A map(list of key-value pairs) where JSON pointer pointing to an attribute within the SharedData serves as key of SharedDataTreatmentInstruction
	TreatmentInstructions *map[string]SharedDataTreatmentInstruction `json:"treatmentInstructions,omitempty"`
	SharedSmSubsData *SessionManagementSubscriptionData1 `json:"sharedSmSubsData,omitempty"`
	SharedEcsAddrConfigInfo NullableEcsAddrConfigInfo1 `json:"sharedEcsAddrConfigInfo,omitempty"`
}

// NewSharedData instantiates a new SharedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedData(sharedDataId string) *SharedData {
	this := SharedData{}
	this.SharedDataId = sharedDataId
	return &this
}

// NewSharedDataWithDefaults instantiates a new SharedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedDataWithDefaults() *SharedData {
	this := SharedData{}
	return &this
}

// GetSharedDataId returns the SharedDataId field value
func (o *SharedData) GetSharedDataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SharedDataId
}

// GetSharedDataIdOk returns a tuple with the SharedDataId field value
// and a boolean to check if the value has been set.
func (o *SharedData) GetSharedDataIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SharedDataId, true
}

// SetSharedDataId sets field value
func (o *SharedData) SetSharedDataId(v string) {
	o.SharedDataId = v
}

// GetSharedAmData returns the SharedAmData field value if set, zero value otherwise.
func (o *SharedData) GetSharedAmData() AccessAndMobilitySubscriptionData1 {
	if o == nil || isNil(o.SharedAmData) {
		var ret AccessAndMobilitySubscriptionData1
		return ret
	}
	return *o.SharedAmData
}

// GetSharedAmDataOk returns a tuple with the SharedAmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedData) GetSharedAmDataOk() (*AccessAndMobilitySubscriptionData1, bool) {
	if o == nil || isNil(o.SharedAmData) {
    return nil, false
	}
	return o.SharedAmData, true
}

// HasSharedAmData returns a boolean if a field has been set.
func (o *SharedData) HasSharedAmData() bool {
	if o != nil && !isNil(o.SharedAmData) {
		return true
	}

	return false
}

// SetSharedAmData gets a reference to the given AccessAndMobilitySubscriptionData1 and assigns it to the SharedAmData field.
func (o *SharedData) SetSharedAmData(v AccessAndMobilitySubscriptionData1) {
	o.SharedAmData = &v
}

// GetSharedSmsSubsData returns the SharedSmsSubsData field value if set, zero value otherwise.
func (o *SharedData) GetSharedSmsSubsData() SmsSubscriptionData1 {
	if o == nil || isNil(o.SharedSmsSubsData) {
		var ret SmsSubscriptionData1
		return ret
	}
	return *o.SharedSmsSubsData
}

// GetSharedSmsSubsDataOk returns a tuple with the SharedSmsSubsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedData) GetSharedSmsSubsDataOk() (*SmsSubscriptionData1, bool) {
	if o == nil || isNil(o.SharedSmsSubsData) {
    return nil, false
	}
	return o.SharedSmsSubsData, true
}

// HasSharedSmsSubsData returns a boolean if a field has been set.
func (o *SharedData) HasSharedSmsSubsData() bool {
	if o != nil && !isNil(o.SharedSmsSubsData) {
		return true
	}

	return false
}

// SetSharedSmsSubsData gets a reference to the given SmsSubscriptionData1 and assigns it to the SharedSmsSubsData field.
func (o *SharedData) SetSharedSmsSubsData(v SmsSubscriptionData1) {
	o.SharedSmsSubsData = &v
}

// GetSharedSmsMngSubsData returns the SharedSmsMngSubsData field value if set, zero value otherwise.
func (o *SharedData) GetSharedSmsMngSubsData() SmsManagementSubscriptionData1 {
	if o == nil || isNil(o.SharedSmsMngSubsData) {
		var ret SmsManagementSubscriptionData1
		return ret
	}
	return *o.SharedSmsMngSubsData
}

// GetSharedSmsMngSubsDataOk returns a tuple with the SharedSmsMngSubsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedData) GetSharedSmsMngSubsDataOk() (*SmsManagementSubscriptionData1, bool) {
	if o == nil || isNil(o.SharedSmsMngSubsData) {
    return nil, false
	}
	return o.SharedSmsMngSubsData, true
}

// HasSharedSmsMngSubsData returns a boolean if a field has been set.
func (o *SharedData) HasSharedSmsMngSubsData() bool {
	if o != nil && !isNil(o.SharedSmsMngSubsData) {
		return true
	}

	return false
}

// SetSharedSmsMngSubsData gets a reference to the given SmsManagementSubscriptionData1 and assigns it to the SharedSmsMngSubsData field.
func (o *SharedData) SetSharedSmsMngSubsData(v SmsManagementSubscriptionData1) {
	o.SharedSmsMngSubsData = &v
}

// GetSharedDnnConfigurations returns the SharedDnnConfigurations field value if set, zero value otherwise.
func (o *SharedData) GetSharedDnnConfigurations() map[string]DnnConfiguration1 {
	if o == nil || isNil(o.SharedDnnConfigurations) {
		var ret map[string]DnnConfiguration1
		return ret
	}
	return *o.SharedDnnConfigurations
}

// GetSharedDnnConfigurationsOk returns a tuple with the SharedDnnConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedData) GetSharedDnnConfigurationsOk() (*map[string]DnnConfiguration1, bool) {
	if o == nil || isNil(o.SharedDnnConfigurations) {
    return nil, false
	}
	return o.SharedDnnConfigurations, true
}

// HasSharedDnnConfigurations returns a boolean if a field has been set.
func (o *SharedData) HasSharedDnnConfigurations() bool {
	if o != nil && !isNil(o.SharedDnnConfigurations) {
		return true
	}

	return false
}

// SetSharedDnnConfigurations gets a reference to the given map[string]DnnConfiguration1 and assigns it to the SharedDnnConfigurations field.
func (o *SharedData) SetSharedDnnConfigurations(v map[string]DnnConfiguration1) {
	o.SharedDnnConfigurations = &v
}

// GetSharedTraceData returns the SharedTraceData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedData) GetSharedTraceData() TraceData {
	if o == nil || isNil(o.SharedTraceData.Get()) {
		var ret TraceData
		return ret
	}
	return *o.SharedTraceData.Get()
}

// GetSharedTraceDataOk returns a tuple with the SharedTraceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharedData) GetSharedTraceDataOk() (*TraceData, bool) {
	if o == nil {
    return nil, false
	}
	return o.SharedTraceData.Get(), o.SharedTraceData.IsSet()
}

// HasSharedTraceData returns a boolean if a field has been set.
func (o *SharedData) HasSharedTraceData() bool {
	if o != nil && o.SharedTraceData.IsSet() {
		return true
	}

	return false
}

// SetSharedTraceData gets a reference to the given NullableTraceData and assigns it to the SharedTraceData field.
func (o *SharedData) SetSharedTraceData(v TraceData) {
	o.SharedTraceData.Set(&v)
}
// SetSharedTraceDataNil sets the value for SharedTraceData to be an explicit nil
func (o *SharedData) SetSharedTraceDataNil() {
	o.SharedTraceData.Set(nil)
}

// UnsetSharedTraceData ensures that no value is present for SharedTraceData, not even an explicit nil
func (o *SharedData) UnsetSharedTraceData() {
	o.SharedTraceData.Unset()
}

// GetSharedSnssaiInfos returns the SharedSnssaiInfos field value if set, zero value otherwise.
func (o *SharedData) GetSharedSnssaiInfos() map[string]SnssaiInfo {
	if o == nil || isNil(o.SharedSnssaiInfos) {
		var ret map[string]SnssaiInfo
		return ret
	}
	return *o.SharedSnssaiInfos
}

// GetSharedSnssaiInfosOk returns a tuple with the SharedSnssaiInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedData) GetSharedSnssaiInfosOk() (*map[string]SnssaiInfo, bool) {
	if o == nil || isNil(o.SharedSnssaiInfos) {
    return nil, false
	}
	return o.SharedSnssaiInfos, true
}

// HasSharedSnssaiInfos returns a boolean if a field has been set.
func (o *SharedData) HasSharedSnssaiInfos() bool {
	if o != nil && !isNil(o.SharedSnssaiInfos) {
		return true
	}

	return false
}

// SetSharedSnssaiInfos gets a reference to the given map[string]SnssaiInfo and assigns it to the SharedSnssaiInfos field.
func (o *SharedData) SetSharedSnssaiInfos(v map[string]SnssaiInfo) {
	o.SharedSnssaiInfos = &v
}

// GetSharedVnGroupDatas returns the SharedVnGroupDatas field value if set, zero value otherwise.
func (o *SharedData) GetSharedVnGroupDatas() map[string]VnGroupData {
	if o == nil || isNil(o.SharedVnGroupDatas) {
		var ret map[string]VnGroupData
		return ret
	}
	return *o.SharedVnGroupDatas
}

// GetSharedVnGroupDatasOk returns a tuple with the SharedVnGroupDatas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedData) GetSharedVnGroupDatasOk() (*map[string]VnGroupData, bool) {
	if o == nil || isNil(o.SharedVnGroupDatas) {
    return nil, false
	}
	return o.SharedVnGroupDatas, true
}

// HasSharedVnGroupDatas returns a boolean if a field has been set.
func (o *SharedData) HasSharedVnGroupDatas() bool {
	if o != nil && !isNil(o.SharedVnGroupDatas) {
		return true
	}

	return false
}

// SetSharedVnGroupDatas gets a reference to the given map[string]VnGroupData and assigns it to the SharedVnGroupDatas field.
func (o *SharedData) SetSharedVnGroupDatas(v map[string]VnGroupData) {
	o.SharedVnGroupDatas = &v
}

// GetTreatmentInstructions returns the TreatmentInstructions field value if set, zero value otherwise.
func (o *SharedData) GetTreatmentInstructions() map[string]SharedDataTreatmentInstruction {
	if o == nil || isNil(o.TreatmentInstructions) {
		var ret map[string]SharedDataTreatmentInstruction
		return ret
	}
	return *o.TreatmentInstructions
}

// GetTreatmentInstructionsOk returns a tuple with the TreatmentInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedData) GetTreatmentInstructionsOk() (*map[string]SharedDataTreatmentInstruction, bool) {
	if o == nil || isNil(o.TreatmentInstructions) {
    return nil, false
	}
	return o.TreatmentInstructions, true
}

// HasTreatmentInstructions returns a boolean if a field has been set.
func (o *SharedData) HasTreatmentInstructions() bool {
	if o != nil && !isNil(o.TreatmentInstructions) {
		return true
	}

	return false
}

// SetTreatmentInstructions gets a reference to the given map[string]SharedDataTreatmentInstruction and assigns it to the TreatmentInstructions field.
func (o *SharedData) SetTreatmentInstructions(v map[string]SharedDataTreatmentInstruction) {
	o.TreatmentInstructions = &v
}

// GetSharedSmSubsData returns the SharedSmSubsData field value if set, zero value otherwise.
func (o *SharedData) GetSharedSmSubsData() SessionManagementSubscriptionData1 {
	if o == nil || isNil(o.SharedSmSubsData) {
		var ret SessionManagementSubscriptionData1
		return ret
	}
	return *o.SharedSmSubsData
}

// GetSharedSmSubsDataOk returns a tuple with the SharedSmSubsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharedData) GetSharedSmSubsDataOk() (*SessionManagementSubscriptionData1, bool) {
	if o == nil || isNil(o.SharedSmSubsData) {
    return nil, false
	}
	return o.SharedSmSubsData, true
}

// HasSharedSmSubsData returns a boolean if a field has been set.
func (o *SharedData) HasSharedSmSubsData() bool {
	if o != nil && !isNil(o.SharedSmSubsData) {
		return true
	}

	return false
}

// SetSharedSmSubsData gets a reference to the given SessionManagementSubscriptionData1 and assigns it to the SharedSmSubsData field.
func (o *SharedData) SetSharedSmSubsData(v SessionManagementSubscriptionData1) {
	o.SharedSmSubsData = &v
}

// GetSharedEcsAddrConfigInfo returns the SharedEcsAddrConfigInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharedData) GetSharedEcsAddrConfigInfo() EcsAddrConfigInfo1 {
	if o == nil || isNil(o.SharedEcsAddrConfigInfo.Get()) {
		var ret EcsAddrConfigInfo1
		return ret
	}
	return *o.SharedEcsAddrConfigInfo.Get()
}

// GetSharedEcsAddrConfigInfoOk returns a tuple with the SharedEcsAddrConfigInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharedData) GetSharedEcsAddrConfigInfoOk() (*EcsAddrConfigInfo1, bool) {
	if o == nil {
    return nil, false
	}
	return o.SharedEcsAddrConfigInfo.Get(), o.SharedEcsAddrConfigInfo.IsSet()
}

// HasSharedEcsAddrConfigInfo returns a boolean if a field has been set.
func (o *SharedData) HasSharedEcsAddrConfigInfo() bool {
	if o != nil && o.SharedEcsAddrConfigInfo.IsSet() {
		return true
	}

	return false
}

// SetSharedEcsAddrConfigInfo gets a reference to the given NullableEcsAddrConfigInfo1 and assigns it to the SharedEcsAddrConfigInfo field.
func (o *SharedData) SetSharedEcsAddrConfigInfo(v EcsAddrConfigInfo1) {
	o.SharedEcsAddrConfigInfo.Set(&v)
}
// SetSharedEcsAddrConfigInfoNil sets the value for SharedEcsAddrConfigInfo to be an explicit nil
func (o *SharedData) SetSharedEcsAddrConfigInfoNil() {
	o.SharedEcsAddrConfigInfo.Set(nil)
}

// UnsetSharedEcsAddrConfigInfo ensures that no value is present for SharedEcsAddrConfigInfo, not even an explicit nil
func (o *SharedData) UnsetSharedEcsAddrConfigInfo() {
	o.SharedEcsAddrConfigInfo.Unset()
}

func (o SharedData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sharedDataId"] = o.SharedDataId
	}
	if !isNil(o.SharedAmData) {
		toSerialize["sharedAmData"] = o.SharedAmData
	}
	if !isNil(o.SharedSmsSubsData) {
		toSerialize["sharedSmsSubsData"] = o.SharedSmsSubsData
	}
	if !isNil(o.SharedSmsMngSubsData) {
		toSerialize["sharedSmsMngSubsData"] = o.SharedSmsMngSubsData
	}
	if !isNil(o.SharedDnnConfigurations) {
		toSerialize["sharedDnnConfigurations"] = o.SharedDnnConfigurations
	}
	if o.SharedTraceData.IsSet() {
		toSerialize["sharedTraceData"] = o.SharedTraceData.Get()
	}
	if !isNil(o.SharedSnssaiInfos) {
		toSerialize["sharedSnssaiInfos"] = o.SharedSnssaiInfos
	}
	if !isNil(o.SharedVnGroupDatas) {
		toSerialize["sharedVnGroupDatas"] = o.SharedVnGroupDatas
	}
	if !isNil(o.TreatmentInstructions) {
		toSerialize["treatmentInstructions"] = o.TreatmentInstructions
	}
	if !isNil(o.SharedSmSubsData) {
		toSerialize["sharedSmSubsData"] = o.SharedSmSubsData
	}
	if o.SharedEcsAddrConfigInfo.IsSet() {
		toSerialize["sharedEcsAddrConfigInfo"] = o.SharedEcsAddrConfigInfo.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSharedData struct {
	value *SharedData
	isSet bool
}

func (v NullableSharedData) Get() *SharedData {
	return v.value
}

func (v *NullableSharedData) Set(val *SharedData) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedData) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedData(val *SharedData) *NullableSharedData {
	return &NullableSharedData{value: val, isSet: true}
}

func (v NullableSharedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


