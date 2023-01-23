/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"time"
)

// UeCommunication Represents UE communication information.
type UeCommunication struct {
	// indicating a time in seconds.
	CommDur *int32 `json:"commDur,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	CommDurVariance *float32 `json:"commDurVariance,omitempty"`
	// indicating a time in seconds.
	PerioTime *int32 `json:"perioTime,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	PerioTimeVariance *float32 `json:"perioTimeVariance,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Ts *time.Time `json:"ts,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	TsVariance *float32 `json:"tsVariance,omitempty"`
	RecurringTime *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	TrafChar *TrafficCharacterization `json:"trafChar,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	Ratio *int32 `json:"ratio,omitempty"`
	PerioCommInd *bool `json:"perioCommInd,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
	AnaOfAppList *AppListForUeComm `json:"anaOfAppList,omitempty"`
	SessInactTimer *SessInactTimerForUeComm `json:"sessInactTimer,omitempty"`
}

// NewUeCommunication instantiates a new UeCommunication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeCommunication() *UeCommunication {
	this := UeCommunication{}
	return &this
}

// NewUeCommunicationWithDefaults instantiates a new UeCommunication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeCommunicationWithDefaults() *UeCommunication {
	this := UeCommunication{}
	return &this
}

// GetCommDur returns the CommDur field value if set, zero value otherwise.
func (o *UeCommunication) GetCommDur() int32 {
	if o == nil || isNil(o.CommDur) {
		var ret int32
		return ret
	}
	return *o.CommDur
}

// GetCommDurOk returns a tuple with the CommDur field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetCommDurOk() (*int32, bool) {
	if o == nil || isNil(o.CommDur) {
    return nil, false
	}
	return o.CommDur, true
}

// HasCommDur returns a boolean if a field has been set.
func (o *UeCommunication) HasCommDur() bool {
	if o != nil && !isNil(o.CommDur) {
		return true
	}

	return false
}

// SetCommDur gets a reference to the given int32 and assigns it to the CommDur field.
func (o *UeCommunication) SetCommDur(v int32) {
	o.CommDur = &v
}

// GetCommDurVariance returns the CommDurVariance field value if set, zero value otherwise.
func (o *UeCommunication) GetCommDurVariance() float32 {
	if o == nil || isNil(o.CommDurVariance) {
		var ret float32
		return ret
	}
	return *o.CommDurVariance
}

// GetCommDurVarianceOk returns a tuple with the CommDurVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetCommDurVarianceOk() (*float32, bool) {
	if o == nil || isNil(o.CommDurVariance) {
    return nil, false
	}
	return o.CommDurVariance, true
}

// HasCommDurVariance returns a boolean if a field has been set.
func (o *UeCommunication) HasCommDurVariance() bool {
	if o != nil && !isNil(o.CommDurVariance) {
		return true
	}

	return false
}

// SetCommDurVariance gets a reference to the given float32 and assigns it to the CommDurVariance field.
func (o *UeCommunication) SetCommDurVariance(v float32) {
	o.CommDurVariance = &v
}

// GetPerioTime returns the PerioTime field value if set, zero value otherwise.
func (o *UeCommunication) GetPerioTime() int32 {
	if o == nil || isNil(o.PerioTime) {
		var ret int32
		return ret
	}
	return *o.PerioTime
}

// GetPerioTimeOk returns a tuple with the PerioTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetPerioTimeOk() (*int32, bool) {
	if o == nil || isNil(o.PerioTime) {
    return nil, false
	}
	return o.PerioTime, true
}

// HasPerioTime returns a boolean if a field has been set.
func (o *UeCommunication) HasPerioTime() bool {
	if o != nil && !isNil(o.PerioTime) {
		return true
	}

	return false
}

// SetPerioTime gets a reference to the given int32 and assigns it to the PerioTime field.
func (o *UeCommunication) SetPerioTime(v int32) {
	o.PerioTime = &v
}

// GetPerioTimeVariance returns the PerioTimeVariance field value if set, zero value otherwise.
func (o *UeCommunication) GetPerioTimeVariance() float32 {
	if o == nil || isNil(o.PerioTimeVariance) {
		var ret float32
		return ret
	}
	return *o.PerioTimeVariance
}

// GetPerioTimeVarianceOk returns a tuple with the PerioTimeVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetPerioTimeVarianceOk() (*float32, bool) {
	if o == nil || isNil(o.PerioTimeVariance) {
    return nil, false
	}
	return o.PerioTimeVariance, true
}

// HasPerioTimeVariance returns a boolean if a field has been set.
func (o *UeCommunication) HasPerioTimeVariance() bool {
	if o != nil && !isNil(o.PerioTimeVariance) {
		return true
	}

	return false
}

// SetPerioTimeVariance gets a reference to the given float32 and assigns it to the PerioTimeVariance field.
func (o *UeCommunication) SetPerioTimeVariance(v float32) {
	o.PerioTimeVariance = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *UeCommunication) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *UeCommunication) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *UeCommunication) SetTs(v time.Time) {
	o.Ts = &v
}

// GetTsVariance returns the TsVariance field value if set, zero value otherwise.
func (o *UeCommunication) GetTsVariance() float32 {
	if o == nil || isNil(o.TsVariance) {
		var ret float32
		return ret
	}
	return *o.TsVariance
}

// GetTsVarianceOk returns a tuple with the TsVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetTsVarianceOk() (*float32, bool) {
	if o == nil || isNil(o.TsVariance) {
    return nil, false
	}
	return o.TsVariance, true
}

// HasTsVariance returns a boolean if a field has been set.
func (o *UeCommunication) HasTsVariance() bool {
	if o != nil && !isNil(o.TsVariance) {
		return true
	}

	return false
}

// SetTsVariance gets a reference to the given float32 and assigns it to the TsVariance field.
func (o *UeCommunication) SetTsVariance(v float32) {
	o.TsVariance = &v
}

// GetRecurringTime returns the RecurringTime field value if set, zero value otherwise.
func (o *UeCommunication) GetRecurringTime() ScheduledCommunicationTime {
	if o == nil || isNil(o.RecurringTime) {
		var ret ScheduledCommunicationTime
		return ret
	}
	return *o.RecurringTime
}

// GetRecurringTimeOk returns a tuple with the RecurringTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetRecurringTimeOk() (*ScheduledCommunicationTime, bool) {
	if o == nil || isNil(o.RecurringTime) {
    return nil, false
	}
	return o.RecurringTime, true
}

// HasRecurringTime returns a boolean if a field has been set.
func (o *UeCommunication) HasRecurringTime() bool {
	if o != nil && !isNil(o.RecurringTime) {
		return true
	}

	return false
}

// SetRecurringTime gets a reference to the given ScheduledCommunicationTime and assigns it to the RecurringTime field.
func (o *UeCommunication) SetRecurringTime(v ScheduledCommunicationTime) {
	o.RecurringTime = &v
}

// GetTrafChar returns the TrafChar field value if set, zero value otherwise.
func (o *UeCommunication) GetTrafChar() TrafficCharacterization {
	if o == nil || isNil(o.TrafChar) {
		var ret TrafficCharacterization
		return ret
	}
	return *o.TrafChar
}

// GetTrafCharOk returns a tuple with the TrafChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetTrafCharOk() (*TrafficCharacterization, bool) {
	if o == nil || isNil(o.TrafChar) {
    return nil, false
	}
	return o.TrafChar, true
}

// HasTrafChar returns a boolean if a field has been set.
func (o *UeCommunication) HasTrafChar() bool {
	if o != nil && !isNil(o.TrafChar) {
		return true
	}

	return false
}

// SetTrafChar gets a reference to the given TrafficCharacterization and assigns it to the TrafChar field.
func (o *UeCommunication) SetTrafChar(v TrafficCharacterization) {
	o.TrafChar = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *UeCommunication) GetRatio() int32 {
	if o == nil || isNil(o.Ratio) {
		var ret int32
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetRatioOk() (*int32, bool) {
	if o == nil || isNil(o.Ratio) {
    return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *UeCommunication) HasRatio() bool {
	if o != nil && !isNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given int32 and assigns it to the Ratio field.
func (o *UeCommunication) SetRatio(v int32) {
	o.Ratio = &v
}

// GetPerioCommInd returns the PerioCommInd field value if set, zero value otherwise.
func (o *UeCommunication) GetPerioCommInd() bool {
	if o == nil || isNil(o.PerioCommInd) {
		var ret bool
		return ret
	}
	return *o.PerioCommInd
}

// GetPerioCommIndOk returns a tuple with the PerioCommInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetPerioCommIndOk() (*bool, bool) {
	if o == nil || isNil(o.PerioCommInd) {
    return nil, false
	}
	return o.PerioCommInd, true
}

// HasPerioCommInd returns a boolean if a field has been set.
func (o *UeCommunication) HasPerioCommInd() bool {
	if o != nil && !isNil(o.PerioCommInd) {
		return true
	}

	return false
}

// SetPerioCommInd gets a reference to the given bool and assigns it to the PerioCommInd field.
func (o *UeCommunication) SetPerioCommInd(v bool) {
	o.PerioCommInd = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *UeCommunication) GetConfidence() int32 {
	if o == nil || isNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetConfidenceOk() (*int32, bool) {
	if o == nil || isNil(o.Confidence) {
    return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *UeCommunication) HasConfidence() bool {
	if o != nil && !isNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *UeCommunication) SetConfidence(v int32) {
	o.Confidence = &v
}

// GetAnaOfAppList returns the AnaOfAppList field value if set, zero value otherwise.
func (o *UeCommunication) GetAnaOfAppList() AppListForUeComm {
	if o == nil || isNil(o.AnaOfAppList) {
		var ret AppListForUeComm
		return ret
	}
	return *o.AnaOfAppList
}

// GetAnaOfAppListOk returns a tuple with the AnaOfAppList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetAnaOfAppListOk() (*AppListForUeComm, bool) {
	if o == nil || isNil(o.AnaOfAppList) {
    return nil, false
	}
	return o.AnaOfAppList, true
}

// HasAnaOfAppList returns a boolean if a field has been set.
func (o *UeCommunication) HasAnaOfAppList() bool {
	if o != nil && !isNil(o.AnaOfAppList) {
		return true
	}

	return false
}

// SetAnaOfAppList gets a reference to the given AppListForUeComm and assigns it to the AnaOfAppList field.
func (o *UeCommunication) SetAnaOfAppList(v AppListForUeComm) {
	o.AnaOfAppList = &v
}

// GetSessInactTimer returns the SessInactTimer field value if set, zero value otherwise.
func (o *UeCommunication) GetSessInactTimer() SessInactTimerForUeComm {
	if o == nil || isNil(o.SessInactTimer) {
		var ret SessInactTimerForUeComm
		return ret
	}
	return *o.SessInactTimer
}

// GetSessInactTimerOk returns a tuple with the SessInactTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunication) GetSessInactTimerOk() (*SessInactTimerForUeComm, bool) {
	if o == nil || isNil(o.SessInactTimer) {
    return nil, false
	}
	return o.SessInactTimer, true
}

// HasSessInactTimer returns a boolean if a field has been set.
func (o *UeCommunication) HasSessInactTimer() bool {
	if o != nil && !isNil(o.SessInactTimer) {
		return true
	}

	return false
}

// SetSessInactTimer gets a reference to the given SessInactTimerForUeComm and assigns it to the SessInactTimer field.
func (o *UeCommunication) SetSessInactTimer(v SessInactTimerForUeComm) {
	o.SessInactTimer = &v
}

func (o UeCommunication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CommDur) {
		toSerialize["commDur"] = o.CommDur
	}
	if !isNil(o.CommDurVariance) {
		toSerialize["commDurVariance"] = o.CommDurVariance
	}
	if !isNil(o.PerioTime) {
		toSerialize["perioTime"] = o.PerioTime
	}
	if !isNil(o.PerioTimeVariance) {
		toSerialize["perioTimeVariance"] = o.PerioTimeVariance
	}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.TsVariance) {
		toSerialize["tsVariance"] = o.TsVariance
	}
	if !isNil(o.RecurringTime) {
		toSerialize["recurringTime"] = o.RecurringTime
	}
	if !isNil(o.TrafChar) {
		toSerialize["trafChar"] = o.TrafChar
	}
	if !isNil(o.Ratio) {
		toSerialize["ratio"] = o.Ratio
	}
	if !isNil(o.PerioCommInd) {
		toSerialize["perioCommInd"] = o.PerioCommInd
	}
	if !isNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	if !isNil(o.AnaOfAppList) {
		toSerialize["anaOfAppList"] = o.AnaOfAppList
	}
	if !isNil(o.SessInactTimer) {
		toSerialize["sessInactTimer"] = o.SessInactTimer
	}
	return json.Marshal(toSerialize)
}

type NullableUeCommunication struct {
	value *UeCommunication
	isSet bool
}

func (v NullableUeCommunication) Get() *UeCommunication {
	return v.value
}

func (v *NullableUeCommunication) Set(val *UeCommunication) {
	v.value = val
	v.isSet = true
}

func (v NullableUeCommunication) IsSet() bool {
	return v.isSet
}

func (v *NullableUeCommunication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeCommunication(val *UeCommunication) *NullableUeCommunication {
	return &NullableUeCommunication{value: val, isSet: true}
}

func (v NullableUeCommunication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeCommunication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


