/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SS_Events

import (
	"encoding/json"
)

// LocationAreaMonReport Event report to notify the VAL UEs moving in or out from a given location.
type LocationAreaMonReport struct {
	// List of identities of all VAL UEs present in the given location area.
	CurPreUEs []ValTargetUe `json:"curPreUEs,omitempty"`
	MoveInOutUEs *MoveInOutUEDetails `json:"moveInOutUEs,omitempty"`
	TrigEvnt *MonLocTriggerEvent `json:"trigEvnt,omitempty"`
}

// NewLocationAreaMonReport instantiates a new LocationAreaMonReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationAreaMonReport() *LocationAreaMonReport {
	this := LocationAreaMonReport{}
	return &this
}

// NewLocationAreaMonReportWithDefaults instantiates a new LocationAreaMonReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationAreaMonReportWithDefaults() *LocationAreaMonReport {
	this := LocationAreaMonReport{}
	return &this
}

// GetCurPreUEs returns the CurPreUEs field value if set, zero value otherwise.
func (o *LocationAreaMonReport) GetCurPreUEs() []ValTargetUe {
	if o == nil || isNil(o.CurPreUEs) {
		var ret []ValTargetUe
		return ret
	}
	return o.CurPreUEs
}

// GetCurPreUEsOk returns a tuple with the CurPreUEs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationAreaMonReport) GetCurPreUEsOk() ([]ValTargetUe, bool) {
	if o == nil || isNil(o.CurPreUEs) {
    return nil, false
	}
	return o.CurPreUEs, true
}

// HasCurPreUEs returns a boolean if a field has been set.
func (o *LocationAreaMonReport) HasCurPreUEs() bool {
	if o != nil && !isNil(o.CurPreUEs) {
		return true
	}

	return false
}

// SetCurPreUEs gets a reference to the given []ValTargetUe and assigns it to the CurPreUEs field.
func (o *LocationAreaMonReport) SetCurPreUEs(v []ValTargetUe) {
	o.CurPreUEs = v
}

// GetMoveInOutUEs returns the MoveInOutUEs field value if set, zero value otherwise.
func (o *LocationAreaMonReport) GetMoveInOutUEs() MoveInOutUEDetails {
	if o == nil || isNil(o.MoveInOutUEs) {
		var ret MoveInOutUEDetails
		return ret
	}
	return *o.MoveInOutUEs
}

// GetMoveInOutUEsOk returns a tuple with the MoveInOutUEs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationAreaMonReport) GetMoveInOutUEsOk() (*MoveInOutUEDetails, bool) {
	if o == nil || isNil(o.MoveInOutUEs) {
    return nil, false
	}
	return o.MoveInOutUEs, true
}

// HasMoveInOutUEs returns a boolean if a field has been set.
func (o *LocationAreaMonReport) HasMoveInOutUEs() bool {
	if o != nil && !isNil(o.MoveInOutUEs) {
		return true
	}

	return false
}

// SetMoveInOutUEs gets a reference to the given MoveInOutUEDetails and assigns it to the MoveInOutUEs field.
func (o *LocationAreaMonReport) SetMoveInOutUEs(v MoveInOutUEDetails) {
	o.MoveInOutUEs = &v
}

// GetTrigEvnt returns the TrigEvnt field value if set, zero value otherwise.
func (o *LocationAreaMonReport) GetTrigEvnt() MonLocTriggerEvent {
	if o == nil || isNil(o.TrigEvnt) {
		var ret MonLocTriggerEvent
		return ret
	}
	return *o.TrigEvnt
}

// GetTrigEvntOk returns a tuple with the TrigEvnt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationAreaMonReport) GetTrigEvntOk() (*MonLocTriggerEvent, bool) {
	if o == nil || isNil(o.TrigEvnt) {
    return nil, false
	}
	return o.TrigEvnt, true
}

// HasTrigEvnt returns a boolean if a field has been set.
func (o *LocationAreaMonReport) HasTrigEvnt() bool {
	if o != nil && !isNil(o.TrigEvnt) {
		return true
	}

	return false
}

// SetTrigEvnt gets a reference to the given MonLocTriggerEvent and assigns it to the TrigEvnt field.
func (o *LocationAreaMonReport) SetTrigEvnt(v MonLocTriggerEvent) {
	o.TrigEvnt = &v
}

func (o LocationAreaMonReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CurPreUEs) {
		toSerialize["curPreUEs"] = o.CurPreUEs
	}
	if !isNil(o.MoveInOutUEs) {
		toSerialize["moveInOutUEs"] = o.MoveInOutUEs
	}
	if !isNil(o.TrigEvnt) {
		toSerialize["trigEvnt"] = o.TrigEvnt
	}
	return json.Marshal(toSerialize)
}

type NullableLocationAreaMonReport struct {
	value *LocationAreaMonReport
	isSet bool
}

func (v NullableLocationAreaMonReport) Get() *LocationAreaMonReport {
	return v.value
}

func (v *NullableLocationAreaMonReport) Set(val *LocationAreaMonReport) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationAreaMonReport) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationAreaMonReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationAreaMonReport(val *LocationAreaMonReport) *NullableLocationAreaMonReport {
	return &NullableLocationAreaMonReport{value: val, isSet: true}
}

func (v NullableLocationAreaMonReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationAreaMonReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


