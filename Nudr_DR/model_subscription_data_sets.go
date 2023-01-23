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

// SubscriptionDataSets struct for SubscriptionDataSets
type SubscriptionDataSets struct {
	AmData *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	SmfSelData *SmfSelectionSubscriptionData `json:"smfSelData,omitempty"`
	UecAmfData *UeContextInAmfData `json:"uecAmfData,omitempty"`
	UecSmfData *UeContextInSmfData `json:"uecSmfData,omitempty"`
	UecSmsfData *UeContextInSmsfData `json:"uecSmsfData,omitempty"`
	SmsSubsData *SmsSubscriptionData `json:"smsSubsData,omitempty"`
	SmData *SmSubsData `json:"smData,omitempty"`
	TraceData NullableTraceData1 `json:"traceData,omitempty"`
	SmsMngData *SmsManagementSubscriptionData `json:"smsMngData,omitempty"`
	LcsPrivacyData *LcsPrivacyData `json:"lcsPrivacyData,omitempty"`
	LcsMoData *LcsMoData `json:"lcsMoData,omitempty"`
	V2xData *V2xSubscriptionData `json:"v2xData,omitempty"`
	LcsBroadcastAssistanceTypesData *LcsBroadcastAssistanceTypesData `json:"lcsBroadcastAssistanceTypesData,omitempty"`
	ProseData *ProseSubscriptionData `json:"proseData,omitempty"`
	MbsData *MbsSubscriptionData `json:"mbsData,omitempty"`
	UcData *UcSubscriptionData `json:"ucData,omitempty"`
}

// NewSubscriptionDataSets instantiates a new SubscriptionDataSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDataSets() *SubscriptionDataSets {
	this := SubscriptionDataSets{}
	return &this
}

// NewSubscriptionDataSetsWithDefaults instantiates a new SubscriptionDataSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDataSetsWithDefaults() *SubscriptionDataSets {
	this := SubscriptionDataSets{}
	return &this
}

// GetAmData returns the AmData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetAmData() AccessAndMobilitySubscriptionData {
	if o == nil || isNil(o.AmData) {
		var ret AccessAndMobilitySubscriptionData
		return ret
	}
	return *o.AmData
}

// GetAmDataOk returns a tuple with the AmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetAmDataOk() (*AccessAndMobilitySubscriptionData, bool) {
	if o == nil || isNil(o.AmData) {
    return nil, false
	}
	return o.AmData, true
}

// HasAmData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasAmData() bool {
	if o != nil && !isNil(o.AmData) {
		return true
	}

	return false
}

// SetAmData gets a reference to the given AccessAndMobilitySubscriptionData and assigns it to the AmData field.
func (o *SubscriptionDataSets) SetAmData(v AccessAndMobilitySubscriptionData) {
	o.AmData = &v
}

// GetSmfSelData returns the SmfSelData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetSmfSelData() SmfSelectionSubscriptionData {
	if o == nil || isNil(o.SmfSelData) {
		var ret SmfSelectionSubscriptionData
		return ret
	}
	return *o.SmfSelData
}

// GetSmfSelDataOk returns a tuple with the SmfSelData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetSmfSelDataOk() (*SmfSelectionSubscriptionData, bool) {
	if o == nil || isNil(o.SmfSelData) {
    return nil, false
	}
	return o.SmfSelData, true
}

// HasSmfSelData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasSmfSelData() bool {
	if o != nil && !isNil(o.SmfSelData) {
		return true
	}

	return false
}

// SetSmfSelData gets a reference to the given SmfSelectionSubscriptionData and assigns it to the SmfSelData field.
func (o *SubscriptionDataSets) SetSmfSelData(v SmfSelectionSubscriptionData) {
	o.SmfSelData = &v
}

// GetUecAmfData returns the UecAmfData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetUecAmfData() UeContextInAmfData {
	if o == nil || isNil(o.UecAmfData) {
		var ret UeContextInAmfData
		return ret
	}
	return *o.UecAmfData
}

// GetUecAmfDataOk returns a tuple with the UecAmfData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetUecAmfDataOk() (*UeContextInAmfData, bool) {
	if o == nil || isNil(o.UecAmfData) {
    return nil, false
	}
	return o.UecAmfData, true
}

// HasUecAmfData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasUecAmfData() bool {
	if o != nil && !isNil(o.UecAmfData) {
		return true
	}

	return false
}

// SetUecAmfData gets a reference to the given UeContextInAmfData and assigns it to the UecAmfData field.
func (o *SubscriptionDataSets) SetUecAmfData(v UeContextInAmfData) {
	o.UecAmfData = &v
}

// GetUecSmfData returns the UecSmfData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetUecSmfData() UeContextInSmfData {
	if o == nil || isNil(o.UecSmfData) {
		var ret UeContextInSmfData
		return ret
	}
	return *o.UecSmfData
}

// GetUecSmfDataOk returns a tuple with the UecSmfData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetUecSmfDataOk() (*UeContextInSmfData, bool) {
	if o == nil || isNil(o.UecSmfData) {
    return nil, false
	}
	return o.UecSmfData, true
}

// HasUecSmfData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasUecSmfData() bool {
	if o != nil && !isNil(o.UecSmfData) {
		return true
	}

	return false
}

// SetUecSmfData gets a reference to the given UeContextInSmfData and assigns it to the UecSmfData field.
func (o *SubscriptionDataSets) SetUecSmfData(v UeContextInSmfData) {
	o.UecSmfData = &v
}

// GetUecSmsfData returns the UecSmsfData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetUecSmsfData() UeContextInSmsfData {
	if o == nil || isNil(o.UecSmsfData) {
		var ret UeContextInSmsfData
		return ret
	}
	return *o.UecSmsfData
}

// GetUecSmsfDataOk returns a tuple with the UecSmsfData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetUecSmsfDataOk() (*UeContextInSmsfData, bool) {
	if o == nil || isNil(o.UecSmsfData) {
    return nil, false
	}
	return o.UecSmsfData, true
}

// HasUecSmsfData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasUecSmsfData() bool {
	if o != nil && !isNil(o.UecSmsfData) {
		return true
	}

	return false
}

// SetUecSmsfData gets a reference to the given UeContextInSmsfData and assigns it to the UecSmsfData field.
func (o *SubscriptionDataSets) SetUecSmsfData(v UeContextInSmsfData) {
	o.UecSmsfData = &v
}

// GetSmsSubsData returns the SmsSubsData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetSmsSubsData() SmsSubscriptionData {
	if o == nil || isNil(o.SmsSubsData) {
		var ret SmsSubscriptionData
		return ret
	}
	return *o.SmsSubsData
}

// GetSmsSubsDataOk returns a tuple with the SmsSubsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetSmsSubsDataOk() (*SmsSubscriptionData, bool) {
	if o == nil || isNil(o.SmsSubsData) {
    return nil, false
	}
	return o.SmsSubsData, true
}

// HasSmsSubsData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasSmsSubsData() bool {
	if o != nil && !isNil(o.SmsSubsData) {
		return true
	}

	return false
}

// SetSmsSubsData gets a reference to the given SmsSubscriptionData and assigns it to the SmsSubsData field.
func (o *SubscriptionDataSets) SetSmsSubsData(v SmsSubscriptionData) {
	o.SmsSubsData = &v
}

// GetSmData returns the SmData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetSmData() SmSubsData {
	if o == nil || isNil(o.SmData) {
		var ret SmSubsData
		return ret
	}
	return *o.SmData
}

// GetSmDataOk returns a tuple with the SmData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetSmDataOk() (*SmSubsData, bool) {
	if o == nil || isNil(o.SmData) {
    return nil, false
	}
	return o.SmData, true
}

// HasSmData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasSmData() bool {
	if o != nil && !isNil(o.SmData) {
		return true
	}

	return false
}

// SetSmData gets a reference to the given SmSubsData and assigns it to the SmData field.
func (o *SubscriptionDataSets) SetSmData(v SmSubsData) {
	o.SmData = &v
}

// GetTraceData returns the TraceData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionDataSets) GetTraceData() TraceData1 {
	if o == nil || isNil(o.TraceData.Get()) {
		var ret TraceData1
		return ret
	}
	return *o.TraceData.Get()
}

// GetTraceDataOk returns a tuple with the TraceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionDataSets) GetTraceDataOk() (*TraceData1, bool) {
	if o == nil {
    return nil, false
	}
	return o.TraceData.Get(), o.TraceData.IsSet()
}

// HasTraceData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasTraceData() bool {
	if o != nil && o.TraceData.IsSet() {
		return true
	}

	return false
}

// SetTraceData gets a reference to the given NullableTraceData1 and assigns it to the TraceData field.
func (o *SubscriptionDataSets) SetTraceData(v TraceData1) {
	o.TraceData.Set(&v)
}
// SetTraceDataNil sets the value for TraceData to be an explicit nil
func (o *SubscriptionDataSets) SetTraceDataNil() {
	o.TraceData.Set(nil)
}

// UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
func (o *SubscriptionDataSets) UnsetTraceData() {
	o.TraceData.Unset()
}

// GetSmsMngData returns the SmsMngData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetSmsMngData() SmsManagementSubscriptionData {
	if o == nil || isNil(o.SmsMngData) {
		var ret SmsManagementSubscriptionData
		return ret
	}
	return *o.SmsMngData
}

// GetSmsMngDataOk returns a tuple with the SmsMngData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetSmsMngDataOk() (*SmsManagementSubscriptionData, bool) {
	if o == nil || isNil(o.SmsMngData) {
    return nil, false
	}
	return o.SmsMngData, true
}

// HasSmsMngData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasSmsMngData() bool {
	if o != nil && !isNil(o.SmsMngData) {
		return true
	}

	return false
}

// SetSmsMngData gets a reference to the given SmsManagementSubscriptionData and assigns it to the SmsMngData field.
func (o *SubscriptionDataSets) SetSmsMngData(v SmsManagementSubscriptionData) {
	o.SmsMngData = &v
}

// GetLcsPrivacyData returns the LcsPrivacyData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetLcsPrivacyData() LcsPrivacyData {
	if o == nil || isNil(o.LcsPrivacyData) {
		var ret LcsPrivacyData
		return ret
	}
	return *o.LcsPrivacyData
}

// GetLcsPrivacyDataOk returns a tuple with the LcsPrivacyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetLcsPrivacyDataOk() (*LcsPrivacyData, bool) {
	if o == nil || isNil(o.LcsPrivacyData) {
    return nil, false
	}
	return o.LcsPrivacyData, true
}

// HasLcsPrivacyData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasLcsPrivacyData() bool {
	if o != nil && !isNil(o.LcsPrivacyData) {
		return true
	}

	return false
}

// SetLcsPrivacyData gets a reference to the given LcsPrivacyData and assigns it to the LcsPrivacyData field.
func (o *SubscriptionDataSets) SetLcsPrivacyData(v LcsPrivacyData) {
	o.LcsPrivacyData = &v
}

// GetLcsMoData returns the LcsMoData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetLcsMoData() LcsMoData {
	if o == nil || isNil(o.LcsMoData) {
		var ret LcsMoData
		return ret
	}
	return *o.LcsMoData
}

// GetLcsMoDataOk returns a tuple with the LcsMoData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetLcsMoDataOk() (*LcsMoData, bool) {
	if o == nil || isNil(o.LcsMoData) {
    return nil, false
	}
	return o.LcsMoData, true
}

// HasLcsMoData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasLcsMoData() bool {
	if o != nil && !isNil(o.LcsMoData) {
		return true
	}

	return false
}

// SetLcsMoData gets a reference to the given LcsMoData and assigns it to the LcsMoData field.
func (o *SubscriptionDataSets) SetLcsMoData(v LcsMoData) {
	o.LcsMoData = &v
}

// GetV2xData returns the V2xData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetV2xData() V2xSubscriptionData {
	if o == nil || isNil(o.V2xData) {
		var ret V2xSubscriptionData
		return ret
	}
	return *o.V2xData
}

// GetV2xDataOk returns a tuple with the V2xData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetV2xDataOk() (*V2xSubscriptionData, bool) {
	if o == nil || isNil(o.V2xData) {
    return nil, false
	}
	return o.V2xData, true
}

// HasV2xData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasV2xData() bool {
	if o != nil && !isNil(o.V2xData) {
		return true
	}

	return false
}

// SetV2xData gets a reference to the given V2xSubscriptionData and assigns it to the V2xData field.
func (o *SubscriptionDataSets) SetV2xData(v V2xSubscriptionData) {
	o.V2xData = &v
}

// GetLcsBroadcastAssistanceTypesData returns the LcsBroadcastAssistanceTypesData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetLcsBroadcastAssistanceTypesData() LcsBroadcastAssistanceTypesData {
	if o == nil || isNil(o.LcsBroadcastAssistanceTypesData) {
		var ret LcsBroadcastAssistanceTypesData
		return ret
	}
	return *o.LcsBroadcastAssistanceTypesData
}

// GetLcsBroadcastAssistanceTypesDataOk returns a tuple with the LcsBroadcastAssistanceTypesData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetLcsBroadcastAssistanceTypesDataOk() (*LcsBroadcastAssistanceTypesData, bool) {
	if o == nil || isNil(o.LcsBroadcastAssistanceTypesData) {
    return nil, false
	}
	return o.LcsBroadcastAssistanceTypesData, true
}

// HasLcsBroadcastAssistanceTypesData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasLcsBroadcastAssistanceTypesData() bool {
	if o != nil && !isNil(o.LcsBroadcastAssistanceTypesData) {
		return true
	}

	return false
}

// SetLcsBroadcastAssistanceTypesData gets a reference to the given LcsBroadcastAssistanceTypesData and assigns it to the LcsBroadcastAssistanceTypesData field.
func (o *SubscriptionDataSets) SetLcsBroadcastAssistanceTypesData(v LcsBroadcastAssistanceTypesData) {
	o.LcsBroadcastAssistanceTypesData = &v
}

// GetProseData returns the ProseData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetProseData() ProseSubscriptionData {
	if o == nil || isNil(o.ProseData) {
		var ret ProseSubscriptionData
		return ret
	}
	return *o.ProseData
}

// GetProseDataOk returns a tuple with the ProseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetProseDataOk() (*ProseSubscriptionData, bool) {
	if o == nil || isNil(o.ProseData) {
    return nil, false
	}
	return o.ProseData, true
}

// HasProseData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasProseData() bool {
	if o != nil && !isNil(o.ProseData) {
		return true
	}

	return false
}

// SetProseData gets a reference to the given ProseSubscriptionData and assigns it to the ProseData field.
func (o *SubscriptionDataSets) SetProseData(v ProseSubscriptionData) {
	o.ProseData = &v
}

// GetMbsData returns the MbsData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetMbsData() MbsSubscriptionData {
	if o == nil || isNil(o.MbsData) {
		var ret MbsSubscriptionData
		return ret
	}
	return *o.MbsData
}

// GetMbsDataOk returns a tuple with the MbsData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetMbsDataOk() (*MbsSubscriptionData, bool) {
	if o == nil || isNil(o.MbsData) {
    return nil, false
	}
	return o.MbsData, true
}

// HasMbsData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasMbsData() bool {
	if o != nil && !isNil(o.MbsData) {
		return true
	}

	return false
}

// SetMbsData gets a reference to the given MbsSubscriptionData and assigns it to the MbsData field.
func (o *SubscriptionDataSets) SetMbsData(v MbsSubscriptionData) {
	o.MbsData = &v
}

// GetUcData returns the UcData field value if set, zero value otherwise.
func (o *SubscriptionDataSets) GetUcData() UcSubscriptionData {
	if o == nil || isNil(o.UcData) {
		var ret UcSubscriptionData
		return ret
	}
	return *o.UcData
}

// GetUcDataOk returns a tuple with the UcData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionDataSets) GetUcDataOk() (*UcSubscriptionData, bool) {
	if o == nil || isNil(o.UcData) {
    return nil, false
	}
	return o.UcData, true
}

// HasUcData returns a boolean if a field has been set.
func (o *SubscriptionDataSets) HasUcData() bool {
	if o != nil && !isNil(o.UcData) {
		return true
	}

	return false
}

// SetUcData gets a reference to the given UcSubscriptionData and assigns it to the UcData field.
func (o *SubscriptionDataSets) SetUcData(v UcSubscriptionData) {
	o.UcData = &v
}

func (o SubscriptionDataSets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AmData) {
		toSerialize["amData"] = o.AmData
	}
	if !isNil(o.SmfSelData) {
		toSerialize["smfSelData"] = o.SmfSelData
	}
	if !isNil(o.UecAmfData) {
		toSerialize["uecAmfData"] = o.UecAmfData
	}
	if !isNil(o.UecSmfData) {
		toSerialize["uecSmfData"] = o.UecSmfData
	}
	if !isNil(o.UecSmsfData) {
		toSerialize["uecSmsfData"] = o.UecSmsfData
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
	if !isNil(o.V2xData) {
		toSerialize["v2xData"] = o.V2xData
	}
	if !isNil(o.LcsBroadcastAssistanceTypesData) {
		toSerialize["lcsBroadcastAssistanceTypesData"] = o.LcsBroadcastAssistanceTypesData
	}
	if !isNil(o.ProseData) {
		toSerialize["proseData"] = o.ProseData
	}
	if !isNil(o.MbsData) {
		toSerialize["mbsData"] = o.MbsData
	}
	if !isNil(o.UcData) {
		toSerialize["ucData"] = o.UcData
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionDataSets struct {
	value *SubscriptionDataSets
	isSet bool
}

func (v NullableSubscriptionDataSets) Get() *SubscriptionDataSets {
	return v.value
}

func (v *NullableSubscriptionDataSets) Set(val *SubscriptionDataSets) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDataSets) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDataSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDataSets(val *SubscriptionDataSets) *NullableSubscriptionDataSets {
	return &NullableSubscriptionDataSets{value: val, isSet: true}
}

func (v NullableSubscriptionDataSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDataSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


