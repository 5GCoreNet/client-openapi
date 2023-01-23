/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_OfflineOnlyCharging

import (
	"encoding/json"
)

// RoamingChargingProfile struct for RoamingChargingProfile
type RoamingChargingProfile struct {
	Triggers []Trigger `json:"triggers,omitempty"`
	PartialRecordMethod *PartialRecordMethod `json:"partialRecordMethod,omitempty"`
}

// NewRoamingChargingProfile instantiates a new RoamingChargingProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoamingChargingProfile() *RoamingChargingProfile {
	this := RoamingChargingProfile{}
	return &this
}

// NewRoamingChargingProfileWithDefaults instantiates a new RoamingChargingProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoamingChargingProfileWithDefaults() *RoamingChargingProfile {
	this := RoamingChargingProfile{}
	return &this
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *RoamingChargingProfile) GetTriggers() []Trigger {
	if o == nil || isNil(o.Triggers) {
		var ret []Trigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoamingChargingProfile) GetTriggersOk() ([]Trigger, bool) {
	if o == nil || isNil(o.Triggers) {
    return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *RoamingChargingProfile) HasTriggers() bool {
	if o != nil && !isNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []Trigger and assigns it to the Triggers field.
func (o *RoamingChargingProfile) SetTriggers(v []Trigger) {
	o.Triggers = v
}

// GetPartialRecordMethod returns the PartialRecordMethod field value if set, zero value otherwise.
func (o *RoamingChargingProfile) GetPartialRecordMethod() PartialRecordMethod {
	if o == nil || isNil(o.PartialRecordMethod) {
		var ret PartialRecordMethod
		return ret
	}
	return *o.PartialRecordMethod
}

// GetPartialRecordMethodOk returns a tuple with the PartialRecordMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoamingChargingProfile) GetPartialRecordMethodOk() (*PartialRecordMethod, bool) {
	if o == nil || isNil(o.PartialRecordMethod) {
    return nil, false
	}
	return o.PartialRecordMethod, true
}

// HasPartialRecordMethod returns a boolean if a field has been set.
func (o *RoamingChargingProfile) HasPartialRecordMethod() bool {
	if o != nil && !isNil(o.PartialRecordMethod) {
		return true
	}

	return false
}

// SetPartialRecordMethod gets a reference to the given PartialRecordMethod and assigns it to the PartialRecordMethod field.
func (o *RoamingChargingProfile) SetPartialRecordMethod(v PartialRecordMethod) {
	o.PartialRecordMethod = &v
}

func (o RoamingChargingProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	if !isNil(o.PartialRecordMethod) {
		toSerialize["partialRecordMethod"] = o.PartialRecordMethod
	}
	return json.Marshal(toSerialize)
}

type NullableRoamingChargingProfile struct {
	value *RoamingChargingProfile
	isSet bool
}

func (v NullableRoamingChargingProfile) Get() *RoamingChargingProfile {
	return v.value
}

func (v *NullableRoamingChargingProfile) Set(val *RoamingChargingProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableRoamingChargingProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableRoamingChargingProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoamingChargingProfile(val *RoamingChargingProfile) *NullableRoamingChargingProfile {
	return &NullableRoamingChargingProfile{value: val, isSet: true}
}

func (v NullableRoamingChargingProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoamingChargingProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


