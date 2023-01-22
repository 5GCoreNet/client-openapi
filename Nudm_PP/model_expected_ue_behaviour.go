/*
Nudm_PP

Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_PP

import (
	"encoding/json"
	"time"
)

// ExpectedUeBehaviour struct for ExpectedUeBehaviour
type ExpectedUeBehaviour struct {
	AfInstanceId string `json:"afInstanceId"`
	ReferenceId int32 `json:"referenceId"`
	StationaryIndication *StationaryIndicationRm `json:"stationaryIndication,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	CommunicationDurationTime NullableInt32 `json:"communicationDurationTime,omitempty"`
	ScheduledCommunicationType *ScheduledCommunicationTypeRm `json:"scheduledCommunicationType,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	PeriodicTime NullableInt32 `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTimeRm `json:"scheduledCommunicationTime,omitempty"`
	// Identifies the UE's expected geographical movement. The attribute is only applicable in 5G.
	ExpectedUmts []LocationArea `json:"expectedUmts,omitempty"`
	TrafficProfile *TrafficProfileRm `json:"trafficProfile,omitempty"`
	BatteryIndication *BatteryIndicationRm `json:"batteryIndication,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation *string `json:"mtcProviderInformation,omitempty"`
}

// NewExpectedUeBehaviour instantiates a new ExpectedUeBehaviour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpectedUeBehaviour(afInstanceId string, referenceId int32) *ExpectedUeBehaviour {
	this := ExpectedUeBehaviour{}
	this.AfInstanceId = afInstanceId
	this.ReferenceId = referenceId
	return &this
}

// NewExpectedUeBehaviourWithDefaults instantiates a new ExpectedUeBehaviour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpectedUeBehaviourWithDefaults() *ExpectedUeBehaviour {
	this := ExpectedUeBehaviour{}
	return &this
}

// GetAfInstanceId returns the AfInstanceId field value
func (o *ExpectedUeBehaviour) GetAfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfInstanceId
}

// GetAfInstanceIdOk returns a tuple with the AfInstanceId field value
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetAfInstanceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AfInstanceId, true
}

// SetAfInstanceId sets field value
func (o *ExpectedUeBehaviour) SetAfInstanceId(v string) {
	o.AfInstanceId = v
}

// GetReferenceId returns the ReferenceId field value
func (o *ExpectedUeBehaviour) GetReferenceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetReferenceIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *ExpectedUeBehaviour) SetReferenceId(v int32) {
	o.ReferenceId = v
}

// GetStationaryIndication returns the StationaryIndication field value if set, zero value otherwise.
func (o *ExpectedUeBehaviour) GetStationaryIndication() StationaryIndicationRm {
	if o == nil || isNil(o.StationaryIndication) {
		var ret StationaryIndicationRm
		return ret
	}
	return *o.StationaryIndication
}

// GetStationaryIndicationOk returns a tuple with the StationaryIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetStationaryIndicationOk() (*StationaryIndicationRm, bool) {
	if o == nil || isNil(o.StationaryIndication) {
    return nil, false
	}
	return o.StationaryIndication, true
}

// HasStationaryIndication returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasStationaryIndication() bool {
	if o != nil && !isNil(o.StationaryIndication) {
		return true
	}

	return false
}

// SetStationaryIndication gets a reference to the given StationaryIndicationRm and assigns it to the StationaryIndication field.
func (o *ExpectedUeBehaviour) SetStationaryIndication(v StationaryIndicationRm) {
	o.StationaryIndication = &v
}

// GetCommunicationDurationTime returns the CommunicationDurationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpectedUeBehaviour) GetCommunicationDurationTime() int32 {
	if o == nil || isNil(o.CommunicationDurationTime.Get()) {
		var ret int32
		return ret
	}
	return *o.CommunicationDurationTime.Get()
}

// GetCommunicationDurationTimeOk returns a tuple with the CommunicationDurationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpectedUeBehaviour) GetCommunicationDurationTimeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.CommunicationDurationTime.Get(), o.CommunicationDurationTime.IsSet()
}

// HasCommunicationDurationTime returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasCommunicationDurationTime() bool {
	if o != nil && o.CommunicationDurationTime.IsSet() {
		return true
	}

	return false
}

// SetCommunicationDurationTime gets a reference to the given NullableInt32 and assigns it to the CommunicationDurationTime field.
func (o *ExpectedUeBehaviour) SetCommunicationDurationTime(v int32) {
	o.CommunicationDurationTime.Set(&v)
}
// SetCommunicationDurationTimeNil sets the value for CommunicationDurationTime to be an explicit nil
func (o *ExpectedUeBehaviour) SetCommunicationDurationTimeNil() {
	o.CommunicationDurationTime.Set(nil)
}

// UnsetCommunicationDurationTime ensures that no value is present for CommunicationDurationTime, not even an explicit nil
func (o *ExpectedUeBehaviour) UnsetCommunicationDurationTime() {
	o.CommunicationDurationTime.Unset()
}

// GetScheduledCommunicationType returns the ScheduledCommunicationType field value if set, zero value otherwise.
func (o *ExpectedUeBehaviour) GetScheduledCommunicationType() ScheduledCommunicationTypeRm {
	if o == nil || isNil(o.ScheduledCommunicationType) {
		var ret ScheduledCommunicationTypeRm
		return ret
	}
	return *o.ScheduledCommunicationType
}

// GetScheduledCommunicationTypeOk returns a tuple with the ScheduledCommunicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetScheduledCommunicationTypeOk() (*ScheduledCommunicationTypeRm, bool) {
	if o == nil || isNil(o.ScheduledCommunicationType) {
    return nil, false
	}
	return o.ScheduledCommunicationType, true
}

// HasScheduledCommunicationType returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasScheduledCommunicationType() bool {
	if o != nil && !isNil(o.ScheduledCommunicationType) {
		return true
	}

	return false
}

// SetScheduledCommunicationType gets a reference to the given ScheduledCommunicationTypeRm and assigns it to the ScheduledCommunicationType field.
func (o *ExpectedUeBehaviour) SetScheduledCommunicationType(v ScheduledCommunicationTypeRm) {
	o.ScheduledCommunicationType = &v
}

// GetPeriodicTime returns the PeriodicTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpectedUeBehaviour) GetPeriodicTime() int32 {
	if o == nil || isNil(o.PeriodicTime.Get()) {
		var ret int32
		return ret
	}
	return *o.PeriodicTime.Get()
}

// GetPeriodicTimeOk returns a tuple with the PeriodicTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpectedUeBehaviour) GetPeriodicTimeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.PeriodicTime.Get(), o.PeriodicTime.IsSet()
}

// HasPeriodicTime returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasPeriodicTime() bool {
	if o != nil && o.PeriodicTime.IsSet() {
		return true
	}

	return false
}

// SetPeriodicTime gets a reference to the given NullableInt32 and assigns it to the PeriodicTime field.
func (o *ExpectedUeBehaviour) SetPeriodicTime(v int32) {
	o.PeriodicTime.Set(&v)
}
// SetPeriodicTimeNil sets the value for PeriodicTime to be an explicit nil
func (o *ExpectedUeBehaviour) SetPeriodicTimeNil() {
	o.PeriodicTime.Set(nil)
}

// UnsetPeriodicTime ensures that no value is present for PeriodicTime, not even an explicit nil
func (o *ExpectedUeBehaviour) UnsetPeriodicTime() {
	o.PeriodicTime.Unset()
}

// GetScheduledCommunicationTime returns the ScheduledCommunicationTime field value if set, zero value otherwise.
func (o *ExpectedUeBehaviour) GetScheduledCommunicationTime() ScheduledCommunicationTimeRm {
	if o == nil || isNil(o.ScheduledCommunicationTime) {
		var ret ScheduledCommunicationTimeRm
		return ret
	}
	return *o.ScheduledCommunicationTime
}

// GetScheduledCommunicationTimeOk returns a tuple with the ScheduledCommunicationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetScheduledCommunicationTimeOk() (*ScheduledCommunicationTimeRm, bool) {
	if o == nil || isNil(o.ScheduledCommunicationTime) {
    return nil, false
	}
	return o.ScheduledCommunicationTime, true
}

// HasScheduledCommunicationTime returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasScheduledCommunicationTime() bool {
	if o != nil && !isNil(o.ScheduledCommunicationTime) {
		return true
	}

	return false
}

// SetScheduledCommunicationTime gets a reference to the given ScheduledCommunicationTimeRm and assigns it to the ScheduledCommunicationTime field.
func (o *ExpectedUeBehaviour) SetScheduledCommunicationTime(v ScheduledCommunicationTimeRm) {
	o.ScheduledCommunicationTime = &v
}

// GetExpectedUmts returns the ExpectedUmts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpectedUeBehaviour) GetExpectedUmts() []LocationArea {
	if o == nil {
		var ret []LocationArea
		return ret
	}
	return o.ExpectedUmts
}

// GetExpectedUmtsOk returns a tuple with the ExpectedUmts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpectedUeBehaviour) GetExpectedUmtsOk() ([]LocationArea, bool) {
	if o == nil || isNil(o.ExpectedUmts) {
    return nil, false
	}
	return o.ExpectedUmts, true
}

// HasExpectedUmts returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasExpectedUmts() bool {
	if o != nil && isNil(o.ExpectedUmts) {
		return true
	}

	return false
}

// SetExpectedUmts gets a reference to the given []LocationArea and assigns it to the ExpectedUmts field.
func (o *ExpectedUeBehaviour) SetExpectedUmts(v []LocationArea) {
	o.ExpectedUmts = v
}

// GetTrafficProfile returns the TrafficProfile field value if set, zero value otherwise.
func (o *ExpectedUeBehaviour) GetTrafficProfile() TrafficProfileRm {
	if o == nil || isNil(o.TrafficProfile) {
		var ret TrafficProfileRm
		return ret
	}
	return *o.TrafficProfile
}

// GetTrafficProfileOk returns a tuple with the TrafficProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetTrafficProfileOk() (*TrafficProfileRm, bool) {
	if o == nil || isNil(o.TrafficProfile) {
    return nil, false
	}
	return o.TrafficProfile, true
}

// HasTrafficProfile returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasTrafficProfile() bool {
	if o != nil && !isNil(o.TrafficProfile) {
		return true
	}

	return false
}

// SetTrafficProfile gets a reference to the given TrafficProfileRm and assigns it to the TrafficProfile field.
func (o *ExpectedUeBehaviour) SetTrafficProfile(v TrafficProfileRm) {
	o.TrafficProfile = &v
}

// GetBatteryIndication returns the BatteryIndication field value if set, zero value otherwise.
func (o *ExpectedUeBehaviour) GetBatteryIndication() BatteryIndicationRm {
	if o == nil || isNil(o.BatteryIndication) {
		var ret BatteryIndicationRm
		return ret
	}
	return *o.BatteryIndication
}

// GetBatteryIndicationOk returns a tuple with the BatteryIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetBatteryIndicationOk() (*BatteryIndicationRm, bool) {
	if o == nil || isNil(o.BatteryIndication) {
    return nil, false
	}
	return o.BatteryIndication, true
}

// HasBatteryIndication returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasBatteryIndication() bool {
	if o != nil && !isNil(o.BatteryIndication) {
		return true
	}

	return false
}

// SetBatteryIndication gets a reference to the given BatteryIndicationRm and assigns it to the BatteryIndication field.
func (o *ExpectedUeBehaviour) SetBatteryIndication(v BatteryIndicationRm) {
	o.BatteryIndication = &v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *ExpectedUeBehaviour) GetValidityTime() time.Time {
	if o == nil || isNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ValidityTime) {
    return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasValidityTime() bool {
	if o != nil && !isNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *ExpectedUeBehaviour) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

// GetMtcProviderInformation returns the MtcProviderInformation field value if set, zero value otherwise.
func (o *ExpectedUeBehaviour) GetMtcProviderInformation() string {
	if o == nil || isNil(o.MtcProviderInformation) {
		var ret string
		return ret
	}
	return *o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil || isNil(o.MtcProviderInformation) {
    return nil, false
	}
	return o.MtcProviderInformation, true
}

// HasMtcProviderInformation returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasMtcProviderInformation() bool {
	if o != nil && !isNil(o.MtcProviderInformation) {
		return true
	}

	return false
}

// SetMtcProviderInformation gets a reference to the given string and assigns it to the MtcProviderInformation field.
func (o *ExpectedUeBehaviour) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = &v
}

func (o ExpectedUeBehaviour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["afInstanceId"] = o.AfInstanceId
	}
	if true {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if !isNil(o.StationaryIndication) {
		toSerialize["stationaryIndication"] = o.StationaryIndication
	}
	if o.CommunicationDurationTime.IsSet() {
		toSerialize["communicationDurationTime"] = o.CommunicationDurationTime.Get()
	}
	if !isNil(o.ScheduledCommunicationType) {
		toSerialize["scheduledCommunicationType"] = o.ScheduledCommunicationType
	}
	if o.PeriodicTime.IsSet() {
		toSerialize["periodicTime"] = o.PeriodicTime.Get()
	}
	if !isNil(o.ScheduledCommunicationTime) {
		toSerialize["scheduledCommunicationTime"] = o.ScheduledCommunicationTime
	}
	if o.ExpectedUmts != nil {
		toSerialize["expectedUmts"] = o.ExpectedUmts
	}
	if !isNil(o.TrafficProfile) {
		toSerialize["trafficProfile"] = o.TrafficProfile
	}
	if !isNil(o.BatteryIndication) {
		toSerialize["batteryIndication"] = o.BatteryIndication
	}
	if !isNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !isNil(o.MtcProviderInformation) {
		toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	}
	return json.Marshal(toSerialize)
}

type NullableExpectedUeBehaviour struct {
	value *ExpectedUeBehaviour
	isSet bool
}

func (v NullableExpectedUeBehaviour) Get() *ExpectedUeBehaviour {
	return v.value
}

func (v *NullableExpectedUeBehaviour) Set(val *ExpectedUeBehaviour) {
	v.value = val
	v.isSet = true
}

func (v NullableExpectedUeBehaviour) IsSet() bool {
	return v.isSet
}

func (v *NullableExpectedUeBehaviour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpectedUeBehaviour(val *ExpectedUeBehaviour) *NullableExpectedUeBehaviour {
	return &NullableExpectedUeBehaviour{value: val, isSet: true}
}

func (v NullableExpectedUeBehaviour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpectedUeBehaviour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


