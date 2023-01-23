/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// ProvisionedDataSets Contains the provisioned data sets.
type ProvisionedDataSets struct {
	AmData *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	SmfSelData *SmfSelectionSubscriptionData `json:"smfSelData,omitempty"`
	SmsSubsData *SmsSubscriptionData `json:"smsSubsData,omitempty"`
	SmData *SmSubsData `json:"smData,omitempty"`
	TraceData NullableTraceData `json:"traceData,omitempty"`
	SmsMngData *SmsManagementSubscriptionData `json:"smsMngData,omitempty"`
	LcsPrivacyData *LcsPrivacyData `json:"lcsPrivacyData,omitempty"`
	LcsMoData *LcsMoData `json:"lcsMoData,omitempty"`
	LcsBcaData *LcsBroadcastAssistanceTypesData `json:"lcsBcaData,omitempty"`
	V2xData *V2xSubscriptionData `json:"v2xData,omitempty"`
	ProseData *ProseSubscriptionData `json:"proseData,omitempty"`
	OdbData *OdbData `json:"odbData,omitempty"`
	EeProfileData *EeProfileData `json:"eeProfileData,omitempty"`
	PpProfileData *PpProfileData `json:"ppProfileData,omitempty"`
	NiddAuthData *AuthorizationData `json:"niddAuthData,omitempty"`
	MbsSubscriptionData *MbsSubscriptionData1 `json:"mbsSubscriptionData,omitempty"`
}

// NewProvisionedDataSets instantiates a new ProvisionedDataSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionedDataSets() *ProvisionedDataSets {
	this := ProvisionedDataSets{}
	return &this
}

// NewProvisionedDataSetsWithDefaults instantiates a new ProvisionedDataSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionedDataSetsWithDefaults() *ProvisionedDataSets {
	this := ProvisionedDataSets{}
	return &this
}

// GetAmData returns the AmData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetAmData() AccessAndMobilitySubscriptionData {
	if o == nil || isNil(o.AmData) {
		var ret AccessAndMobilitySubscriptionData
		return ret
	}
	return *o.AmData
}

// GetAmDataOk returns a tuple with the AmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetAmDataOk() (*AccessAndMobilitySubscriptionData, bool) {
	if o == nil || isNil(o.AmData) {
    return nil, false
	}
	return o.AmData, true
}

// HasAmData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasAmData() bool {
	if o != nil && !isNil(o.AmData) {
		return true
	}

	return false
}

// SetAmData gets a reference to the given AccessAndMobilitySubscriptionData and assigns it to the AmData field.
func (o *ProvisionedDataSets) SetAmData(v AccessAndMobilitySubscriptionData) {
	o.AmData = &v
}

// GetSmfSelData returns the SmfSelData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetSmfSelData() SmfSelectionSubscriptionData {
	if o == nil || isNil(o.SmfSelData) {
		var ret SmfSelectionSubscriptionData
		return ret
	}
	return *o.SmfSelData
}

// GetSmfSelDataOk returns a tuple with the SmfSelData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetSmfSelDataOk() (*SmfSelectionSubscriptionData, bool) {
	if o == nil || isNil(o.SmfSelData) {
    return nil, false
	}
	return o.SmfSelData, true
}

// HasSmfSelData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasSmfSelData() bool {
	if o != nil && !isNil(o.SmfSelData) {
		return true
	}

	return false
}

// SetSmfSelData gets a reference to the given SmfSelectionSubscriptionData and assigns it to the SmfSelData field.
func (o *ProvisionedDataSets) SetSmfSelData(v SmfSelectionSubscriptionData) {
	o.SmfSelData = &v
}

// GetSmsSubsData returns the SmsSubsData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetSmsSubsData() SmsSubscriptionData {
	if o == nil || isNil(o.SmsSubsData) {
		var ret SmsSubscriptionData
		return ret
	}
	return *o.SmsSubsData
}

// GetSmsSubsDataOk returns a tuple with the SmsSubsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetSmsSubsDataOk() (*SmsSubscriptionData, bool) {
	if o == nil || isNil(o.SmsSubsData) {
    return nil, false
	}
	return o.SmsSubsData, true
}

// HasSmsSubsData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasSmsSubsData() bool {
	if o != nil && !isNil(o.SmsSubsData) {
		return true
	}

	return false
}

// SetSmsSubsData gets a reference to the given SmsSubscriptionData and assigns it to the SmsSubsData field.
func (o *ProvisionedDataSets) SetSmsSubsData(v SmsSubscriptionData) {
	o.SmsSubsData = &v
}

// GetSmData returns the SmData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetSmData() SmSubsData {
	if o == nil || isNil(o.SmData) {
		var ret SmSubsData
		return ret
	}
	return *o.SmData
}

// GetSmDataOk returns a tuple with the SmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetSmDataOk() (*SmSubsData, bool) {
	if o == nil || isNil(o.SmData) {
    return nil, false
	}
	return o.SmData, true
}

// HasSmData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasSmData() bool {
	if o != nil && !isNil(o.SmData) {
		return true
	}

	return false
}

// SetSmData gets a reference to the given SmSubsData and assigns it to the SmData field.
func (o *ProvisionedDataSets) SetSmData(v SmSubsData) {
	o.SmData = &v
}

// GetTraceData returns the TraceData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProvisionedDataSets) GetTraceData() TraceData {
	if o == nil || isNil(o.TraceData.Get()) {
		var ret TraceData
		return ret
	}
	return *o.TraceData.Get()
}

// GetTraceDataOk returns a tuple with the TraceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProvisionedDataSets) GetTraceDataOk() (*TraceData, bool) {
	if o == nil {
    return nil, false
	}
	return o.TraceData.Get(), o.TraceData.IsSet()
}

// HasTraceData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasTraceData() bool {
	if o != nil && o.TraceData.IsSet() {
		return true
	}

	return false
}

// SetTraceData gets a reference to the given NullableTraceData and assigns it to the TraceData field.
func (o *ProvisionedDataSets) SetTraceData(v TraceData) {
	o.TraceData.Set(&v)
}
// SetTraceDataNil sets the value for TraceData to be an explicit nil
func (o *ProvisionedDataSets) SetTraceDataNil() {
	o.TraceData.Set(nil)
}

// UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
func (o *ProvisionedDataSets) UnsetTraceData() {
	o.TraceData.Unset()
}

// GetSmsMngData returns the SmsMngData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetSmsMngData() SmsManagementSubscriptionData {
	if o == nil || isNil(o.SmsMngData) {
		var ret SmsManagementSubscriptionData
		return ret
	}
	return *o.SmsMngData
}

// GetSmsMngDataOk returns a tuple with the SmsMngData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetSmsMngDataOk() (*SmsManagementSubscriptionData, bool) {
	if o == nil || isNil(o.SmsMngData) {
    return nil, false
	}
	return o.SmsMngData, true
}

// HasSmsMngData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasSmsMngData() bool {
	if o != nil && !isNil(o.SmsMngData) {
		return true
	}

	return false
}

// SetSmsMngData gets a reference to the given SmsManagementSubscriptionData and assigns it to the SmsMngData field.
func (o *ProvisionedDataSets) SetSmsMngData(v SmsManagementSubscriptionData) {
	o.SmsMngData = &v
}

// GetLcsPrivacyData returns the LcsPrivacyData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetLcsPrivacyData() LcsPrivacyData {
	if o == nil || isNil(o.LcsPrivacyData) {
		var ret LcsPrivacyData
		return ret
	}
	return *o.LcsPrivacyData
}

// GetLcsPrivacyDataOk returns a tuple with the LcsPrivacyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetLcsPrivacyDataOk() (*LcsPrivacyData, bool) {
	if o == nil || isNil(o.LcsPrivacyData) {
    return nil, false
	}
	return o.LcsPrivacyData, true
}

// HasLcsPrivacyData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasLcsPrivacyData() bool {
	if o != nil && !isNil(o.LcsPrivacyData) {
		return true
	}

	return false
}

// SetLcsPrivacyData gets a reference to the given LcsPrivacyData and assigns it to the LcsPrivacyData field.
func (o *ProvisionedDataSets) SetLcsPrivacyData(v LcsPrivacyData) {
	o.LcsPrivacyData = &v
}

// GetLcsMoData returns the LcsMoData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetLcsMoData() LcsMoData {
	if o == nil || isNil(o.LcsMoData) {
		var ret LcsMoData
		return ret
	}
	return *o.LcsMoData
}

// GetLcsMoDataOk returns a tuple with the LcsMoData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetLcsMoDataOk() (*LcsMoData, bool) {
	if o == nil || isNil(o.LcsMoData) {
    return nil, false
	}
	return o.LcsMoData, true
}

// HasLcsMoData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasLcsMoData() bool {
	if o != nil && !isNil(o.LcsMoData) {
		return true
	}

	return false
}

// SetLcsMoData gets a reference to the given LcsMoData and assigns it to the LcsMoData field.
func (o *ProvisionedDataSets) SetLcsMoData(v LcsMoData) {
	o.LcsMoData = &v
}

// GetLcsBcaData returns the LcsBcaData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetLcsBcaData() LcsBroadcastAssistanceTypesData {
	if o == nil || isNil(o.LcsBcaData) {
		var ret LcsBroadcastAssistanceTypesData
		return ret
	}
	return *o.LcsBcaData
}

// GetLcsBcaDataOk returns a tuple with the LcsBcaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetLcsBcaDataOk() (*LcsBroadcastAssistanceTypesData, bool) {
	if o == nil || isNil(o.LcsBcaData) {
    return nil, false
	}
	return o.LcsBcaData, true
}

// HasLcsBcaData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasLcsBcaData() bool {
	if o != nil && !isNil(o.LcsBcaData) {
		return true
	}

	return false
}

// SetLcsBcaData gets a reference to the given LcsBroadcastAssistanceTypesData and assigns it to the LcsBcaData field.
func (o *ProvisionedDataSets) SetLcsBcaData(v LcsBroadcastAssistanceTypesData) {
	o.LcsBcaData = &v
}

// GetV2xData returns the V2xData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetV2xData() V2xSubscriptionData {
	if o == nil || isNil(o.V2xData) {
		var ret V2xSubscriptionData
		return ret
	}
	return *o.V2xData
}

// GetV2xDataOk returns a tuple with the V2xData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetV2xDataOk() (*V2xSubscriptionData, bool) {
	if o == nil || isNil(o.V2xData) {
    return nil, false
	}
	return o.V2xData, true
}

// HasV2xData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasV2xData() bool {
	if o != nil && !isNil(o.V2xData) {
		return true
	}

	return false
}

// SetV2xData gets a reference to the given V2xSubscriptionData and assigns it to the V2xData field.
func (o *ProvisionedDataSets) SetV2xData(v V2xSubscriptionData) {
	o.V2xData = &v
}

// GetProseData returns the ProseData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetProseData() ProseSubscriptionData {
	if o == nil || isNil(o.ProseData) {
		var ret ProseSubscriptionData
		return ret
	}
	return *o.ProseData
}

// GetProseDataOk returns a tuple with the ProseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetProseDataOk() (*ProseSubscriptionData, bool) {
	if o == nil || isNil(o.ProseData) {
    return nil, false
	}
	return o.ProseData, true
}

// HasProseData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasProseData() bool {
	if o != nil && !isNil(o.ProseData) {
		return true
	}

	return false
}

// SetProseData gets a reference to the given ProseSubscriptionData and assigns it to the ProseData field.
func (o *ProvisionedDataSets) SetProseData(v ProseSubscriptionData) {
	o.ProseData = &v
}

// GetOdbData returns the OdbData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetOdbData() OdbData {
	if o == nil || isNil(o.OdbData) {
		var ret OdbData
		return ret
	}
	return *o.OdbData
}

// GetOdbDataOk returns a tuple with the OdbData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetOdbDataOk() (*OdbData, bool) {
	if o == nil || isNil(o.OdbData) {
    return nil, false
	}
	return o.OdbData, true
}

// HasOdbData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasOdbData() bool {
	if o != nil && !isNil(o.OdbData) {
		return true
	}

	return false
}

// SetOdbData gets a reference to the given OdbData and assigns it to the OdbData field.
func (o *ProvisionedDataSets) SetOdbData(v OdbData) {
	o.OdbData = &v
}

// GetEeProfileData returns the EeProfileData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetEeProfileData() EeProfileData {
	if o == nil || isNil(o.EeProfileData) {
		var ret EeProfileData
		return ret
	}
	return *o.EeProfileData
}

// GetEeProfileDataOk returns a tuple with the EeProfileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetEeProfileDataOk() (*EeProfileData, bool) {
	if o == nil || isNil(o.EeProfileData) {
    return nil, false
	}
	return o.EeProfileData, true
}

// HasEeProfileData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasEeProfileData() bool {
	if o != nil && !isNil(o.EeProfileData) {
		return true
	}

	return false
}

// SetEeProfileData gets a reference to the given EeProfileData and assigns it to the EeProfileData field.
func (o *ProvisionedDataSets) SetEeProfileData(v EeProfileData) {
	o.EeProfileData = &v
}

// GetPpProfileData returns the PpProfileData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetPpProfileData() PpProfileData {
	if o == nil || isNil(o.PpProfileData) {
		var ret PpProfileData
		return ret
	}
	return *o.PpProfileData
}

// GetPpProfileDataOk returns a tuple with the PpProfileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetPpProfileDataOk() (*PpProfileData, bool) {
	if o == nil || isNil(o.PpProfileData) {
    return nil, false
	}
	return o.PpProfileData, true
}

// HasPpProfileData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasPpProfileData() bool {
	if o != nil && !isNil(o.PpProfileData) {
		return true
	}

	return false
}

// SetPpProfileData gets a reference to the given PpProfileData and assigns it to the PpProfileData field.
func (o *ProvisionedDataSets) SetPpProfileData(v PpProfileData) {
	o.PpProfileData = &v
}

// GetNiddAuthData returns the NiddAuthData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetNiddAuthData() AuthorizationData {
	if o == nil || isNil(o.NiddAuthData) {
		var ret AuthorizationData
		return ret
	}
	return *o.NiddAuthData
}

// GetNiddAuthDataOk returns a tuple with the NiddAuthData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetNiddAuthDataOk() (*AuthorizationData, bool) {
	if o == nil || isNil(o.NiddAuthData) {
    return nil, false
	}
	return o.NiddAuthData, true
}

// HasNiddAuthData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasNiddAuthData() bool {
	if o != nil && !isNil(o.NiddAuthData) {
		return true
	}

	return false
}

// SetNiddAuthData gets a reference to the given AuthorizationData and assigns it to the NiddAuthData field.
func (o *ProvisionedDataSets) SetNiddAuthData(v AuthorizationData) {
	o.NiddAuthData = &v
}

// GetMbsSubscriptionData returns the MbsSubscriptionData field value if set, zero value otherwise.
func (o *ProvisionedDataSets) GetMbsSubscriptionData() MbsSubscriptionData1 {
	if o == nil || isNil(o.MbsSubscriptionData) {
		var ret MbsSubscriptionData1
		return ret
	}
	return *o.MbsSubscriptionData
}

// GetMbsSubscriptionDataOk returns a tuple with the MbsSubscriptionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedDataSets) GetMbsSubscriptionDataOk() (*MbsSubscriptionData1, bool) {
	if o == nil || isNil(o.MbsSubscriptionData) {
    return nil, false
	}
	return o.MbsSubscriptionData, true
}

// HasMbsSubscriptionData returns a boolean if a field has been set.
func (o *ProvisionedDataSets) HasMbsSubscriptionData() bool {
	if o != nil && !isNil(o.MbsSubscriptionData) {
		return true
	}

	return false
}

// SetMbsSubscriptionData gets a reference to the given MbsSubscriptionData1 and assigns it to the MbsSubscriptionData field.
func (o *ProvisionedDataSets) SetMbsSubscriptionData(v MbsSubscriptionData1) {
	o.MbsSubscriptionData = &v
}

func (o ProvisionedDataSets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AmData) {
		toSerialize["amData"] = o.AmData
	}
	if !isNil(o.SmfSelData) {
		toSerialize["smfSelData"] = o.SmfSelData
	}
	if !isNil(o.SmsSubsData) {
		toSerialize["smsSubsData"] = o.SmsSubsData
	}
	if !isNil(o.SmData) {
		toSerialize["smData"] = o.SmData
	}
	if o.TraceData.IsSet() {
		toSerialize["traceData"] = o.TraceData.Get()
	}
	if !isNil(o.SmsMngData) {
		toSerialize["smsMngData"] = o.SmsMngData
	}
	if !isNil(o.LcsPrivacyData) {
		toSerialize["lcsPrivacyData"] = o.LcsPrivacyData
	}
	if !isNil(o.LcsMoData) {
		toSerialize["lcsMoData"] = o.LcsMoData
	}
	if !isNil(o.LcsBcaData) {
		toSerialize["lcsBcaData"] = o.LcsBcaData
	}
	if !isNil(o.V2xData) {
		toSerialize["v2xData"] = o.V2xData
	}
	if !isNil(o.ProseData) {
		toSerialize["proseData"] = o.ProseData
	}
	if !isNil(o.OdbData) {
		toSerialize["odbData"] = o.OdbData
	}
	if !isNil(o.EeProfileData) {
		toSerialize["eeProfileData"] = o.EeProfileData
	}
	if !isNil(o.PpProfileData) {
		toSerialize["ppProfileData"] = o.PpProfileData
	}
	if !isNil(o.NiddAuthData) {
		toSerialize["niddAuthData"] = o.NiddAuthData
	}
	if !isNil(o.MbsSubscriptionData) {
		toSerialize["mbsSubscriptionData"] = o.MbsSubscriptionData
	}
	return json.Marshal(toSerialize)
}

type NullableProvisionedDataSets struct {
	value *ProvisionedDataSets
	isSet bool
}

func (v NullableProvisionedDataSets) Get() *ProvisionedDataSets {
	return v.value
}

func (v *NullableProvisionedDataSets) Set(val *ProvisionedDataSets) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionedDataSets) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionedDataSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionedDataSets(val *ProvisionedDataSets) *NullableProvisionedDataSets {
	return &NullableProvisionedDataSets{value: val, isSet: true}
}

func (v NullableProvisionedDataSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionedDataSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


