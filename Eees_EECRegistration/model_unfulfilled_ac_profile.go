/*
Eees_EECRegistration

API for EEC registration. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_EECRegistration

import (
	"encoding/json"
)

// UnfulfilledAcProfile Desrcibes AC Profile ID and reason sent by EES in EEC Register response.
type UnfulfilledAcProfile struct {
	// The AC ID of a AC profile.
	AcId *string `json:"acId,omitempty"`
	Reason *UnfulfillACProfRsn `json:"reason,omitempty"`
}

// NewUnfulfilledAcProfile instantiates a new UnfulfilledAcProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnfulfilledAcProfile() *UnfulfilledAcProfile {
	this := UnfulfilledAcProfile{}
	return &this
}

// NewUnfulfilledAcProfileWithDefaults instantiates a new UnfulfilledAcProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnfulfilledAcProfileWithDefaults() *UnfulfilledAcProfile {
	this := UnfulfilledAcProfile{}
	return &this
}

// GetAcId returns the AcId field value if set, zero value otherwise.
func (o *UnfulfilledAcProfile) GetAcId() string {
	if o == nil || isNil(o.AcId) {
		var ret string
		return ret
	}
	return *o.AcId
}

// GetAcIdOk returns a tuple with the AcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnfulfilledAcProfile) GetAcIdOk() (*string, bool) {
	if o == nil || isNil(o.AcId) {
    return nil, false
	}
	return o.AcId, true
}

// HasAcId returns a boolean if a field has been set.
func (o *UnfulfilledAcProfile) HasAcId() bool {
	if o != nil && !isNil(o.AcId) {
		return true
	}

	return false
}

// SetAcId gets a reference to the given string and assigns it to the AcId field.
func (o *UnfulfilledAcProfile) SetAcId(v string) {
	o.AcId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *UnfulfilledAcProfile) GetReason() UnfulfillACProfRsn {
	if o == nil || isNil(o.Reason) {
		var ret UnfulfillACProfRsn
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnfulfilledAcProfile) GetReasonOk() (*UnfulfillACProfRsn, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *UnfulfilledAcProfile) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given UnfulfillACProfRsn and assigns it to the Reason field.
func (o *UnfulfilledAcProfile) SetReason(v UnfulfillACProfRsn) {
	o.Reason = &v
}

func (o UnfulfilledAcProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AcId) {
		toSerialize["acId"] = o.AcId
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableUnfulfilledAcProfile struct {
	value *UnfulfilledAcProfile
	isSet bool
}

func (v NullableUnfulfilledAcProfile) Get() *UnfulfilledAcProfile {
	return v.value
}

func (v *NullableUnfulfilledAcProfile) Set(val *UnfulfilledAcProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUnfulfilledAcProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUnfulfilledAcProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnfulfilledAcProfile(val *UnfulfilledAcProfile) *NullableUnfulfilledAcProfile {
	return &NullableUnfulfilledAcProfile{value: val, isSet: true}
}

func (v NullableUnfulfilledAcProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnfulfilledAcProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


