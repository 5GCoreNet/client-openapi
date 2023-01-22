/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package 5GcNrm

import (
	"encoding/json"
)

// SnssaiInfo struct for SnssaiInfo
type SnssaiInfo struct {
	PlmnInfo *PlmnInfo `json:"plmnInfo,omitempty"`
	AdministrativeState *AdministrativeState `json:"administrativeState,omitempty"`
}

// NewSnssaiInfo instantiates a new SnssaiInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiInfo() *SnssaiInfo {
	this := SnssaiInfo{}
	return &this
}

// NewSnssaiInfoWithDefaults instantiates a new SnssaiInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiInfoWithDefaults() *SnssaiInfo {
	this := SnssaiInfo{}
	return &this
}

// GetPlmnInfo returns the PlmnInfo field value if set, zero value otherwise.
func (o *SnssaiInfo) GetPlmnInfo() PlmnInfo {
	if o == nil || isNil(o.PlmnInfo) {
		var ret PlmnInfo
		return ret
	}
	return *o.PlmnInfo
}

// GetPlmnInfoOk returns a tuple with the PlmnInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnssaiInfo) GetPlmnInfoOk() (*PlmnInfo, bool) {
	if o == nil || isNil(o.PlmnInfo) {
    return nil, false
	}
	return o.PlmnInfo, true
}

// HasPlmnInfo returns a boolean if a field has been set.
func (o *SnssaiInfo) HasPlmnInfo() bool {
	if o != nil && !isNil(o.PlmnInfo) {
		return true
	}

	return false
}

// SetPlmnInfo gets a reference to the given PlmnInfo and assigns it to the PlmnInfo field.
func (o *SnssaiInfo) SetPlmnInfo(v PlmnInfo) {
	o.PlmnInfo = &v
}

// GetAdministrativeState returns the AdministrativeState field value if set, zero value otherwise.
func (o *SnssaiInfo) GetAdministrativeState() AdministrativeState {
	if o == nil || isNil(o.AdministrativeState) {
		var ret AdministrativeState
		return ret
	}
	return *o.AdministrativeState
}

// GetAdministrativeStateOk returns a tuple with the AdministrativeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnssaiInfo) GetAdministrativeStateOk() (*AdministrativeState, bool) {
	if o == nil || isNil(o.AdministrativeState) {
    return nil, false
	}
	return o.AdministrativeState, true
}

// HasAdministrativeState returns a boolean if a field has been set.
func (o *SnssaiInfo) HasAdministrativeState() bool {
	if o != nil && !isNil(o.AdministrativeState) {
		return true
	}

	return false
}

// SetAdministrativeState gets a reference to the given AdministrativeState and assigns it to the AdministrativeState field.
func (o *SnssaiInfo) SetAdministrativeState(v AdministrativeState) {
	o.AdministrativeState = &v
}

func (o SnssaiInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PlmnInfo) {
		toSerialize["plmnInfo"] = o.PlmnInfo
	}
	if !isNil(o.AdministrativeState) {
		toSerialize["administrativeState"] = o.AdministrativeState
	}
	return json.Marshal(toSerialize)
}

type NullableSnssaiInfo struct {
	value *SnssaiInfo
	isSet bool
}

func (v NullableSnssaiInfo) Get() *SnssaiInfo {
	return v.value
}

func (v *NullableSnssaiInfo) Set(val *SnssaiInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiInfo(val *SnssaiInfo) *NullableSnssaiInfo {
	return &NullableSnssaiInfo{value: val, isSet: true}
}

func (v NullableSnssaiInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


