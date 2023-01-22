/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
	"time"
)

// ConditionData Contains conditions of applicability for a rule.
type ConditionData struct {
	// Uniquely identifies the condition data within a PDU session.
	CondId string `json:"condId"`
	// string with format 'date-time' as defined in OpenAPI with 'nullable:true' property.  
	ActivationTime NullableTime `json:"activationTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI with 'nullable:true' property.  
	DeactivationTime NullableTime `json:"deactivationTime,omitempty"`
	AccessType *AccessType `json:"accessType,omitempty"`
	RatType *RatType `json:"ratType,omitempty"`
}

// NewConditionData instantiates a new ConditionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionData(condId string) *ConditionData {
	this := ConditionData{}
	this.CondId = condId
	return &this
}

// NewConditionDataWithDefaults instantiates a new ConditionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionDataWithDefaults() *ConditionData {
	this := ConditionData{}
	return &this
}

// GetCondId returns the CondId field value
func (o *ConditionData) GetCondId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CondId
}

// GetCondIdOk returns a tuple with the CondId field value
// and a boolean to check if the value has been set.
func (o *ConditionData) GetCondIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CondId, true
}

// SetCondId sets field value
func (o *ConditionData) SetCondId(v string) {
	o.CondId = v
}

// GetActivationTime returns the ActivationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConditionData) GetActivationTime() time.Time {
	if o == nil || isNil(o.ActivationTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ActivationTime.Get()
}

// GetActivationTimeOk returns a tuple with the ActivationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConditionData) GetActivationTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.ActivationTime.Get(), o.ActivationTime.IsSet()
}

// HasActivationTime returns a boolean if a field has been set.
func (o *ConditionData) HasActivationTime() bool {
	if o != nil && o.ActivationTime.IsSet() {
		return true
	}

	return false
}

// SetActivationTime gets a reference to the given NullableTime and assigns it to the ActivationTime field.
func (o *ConditionData) SetActivationTime(v time.Time) {
	o.ActivationTime.Set(&v)
}
// SetActivationTimeNil sets the value for ActivationTime to be an explicit nil
func (o *ConditionData) SetActivationTimeNil() {
	o.ActivationTime.Set(nil)
}

// UnsetActivationTime ensures that no value is present for ActivationTime, not even an explicit nil
func (o *ConditionData) UnsetActivationTime() {
	o.ActivationTime.Unset()
}

// GetDeactivationTime returns the DeactivationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConditionData) GetDeactivationTime() time.Time {
	if o == nil || isNil(o.DeactivationTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DeactivationTime.Get()
}

// GetDeactivationTimeOk returns a tuple with the DeactivationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConditionData) GetDeactivationTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.DeactivationTime.Get(), o.DeactivationTime.IsSet()
}

// HasDeactivationTime returns a boolean if a field has been set.
func (o *ConditionData) HasDeactivationTime() bool {
	if o != nil && o.DeactivationTime.IsSet() {
		return true
	}

	return false
}

// SetDeactivationTime gets a reference to the given NullableTime and assigns it to the DeactivationTime field.
func (o *ConditionData) SetDeactivationTime(v time.Time) {
	o.DeactivationTime.Set(&v)
}
// SetDeactivationTimeNil sets the value for DeactivationTime to be an explicit nil
func (o *ConditionData) SetDeactivationTimeNil() {
	o.DeactivationTime.Set(nil)
}

// UnsetDeactivationTime ensures that no value is present for DeactivationTime, not even an explicit nil
func (o *ConditionData) UnsetDeactivationTime() {
	o.DeactivationTime.Unset()
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *ConditionData) GetAccessType() AccessType {
	if o == nil || isNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionData) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || isNil(o.AccessType) {
    return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *ConditionData) HasAccessType() bool {
	if o != nil && !isNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *ConditionData) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *ConditionData) GetRatType() RatType {
	if o == nil || isNil(o.RatType) {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionData) GetRatTypeOk() (*RatType, bool) {
	if o == nil || isNil(o.RatType) {
    return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *ConditionData) HasRatType() bool {
	if o != nil && !isNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *ConditionData) SetRatType(v RatType) {
	o.RatType = &v
}

func (o ConditionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["condId"] = o.CondId
	}
	if o.ActivationTime.IsSet() {
		toSerialize["activationTime"] = o.ActivationTime.Get()
	}
	if o.DeactivationTime.IsSet() {
		toSerialize["deactivationTime"] = o.DeactivationTime.Get()
	}
	if !isNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !isNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	return json.Marshal(toSerialize)
}

type NullableConditionData struct {
	value *ConditionData
	isSet bool
}

func (v NullableConditionData) Get() *ConditionData {
	return v.value
}

func (v *NullableConditionData) Set(val *ConditionData) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionData) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionData(val *ConditionData) *NullableConditionData {
	return &NullableConditionData{value: val, isSet: true}
}

func (v NullableConditionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


