/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_Events

import (
	"encoding/json"
)

// MoveInOutUEDetails List of UEs moved in and out.
type MoveInOutUEDetails struct {
	// List of identities of VAL UEs who moved in to given location area since previous notification. 
	MoveInUEs []ValTargetUe `json:"moveInUEs,omitempty"`
	// List of identities of VAL UEs who moved out of the given location area since previous notification. 
	MoveOutUEs []ValTargetUe `json:"moveOutUEs,omitempty"`
}

// NewMoveInOutUEDetails instantiates a new MoveInOutUEDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveInOutUEDetails() *MoveInOutUEDetails {
	this := MoveInOutUEDetails{}
	return &this
}

// NewMoveInOutUEDetailsWithDefaults instantiates a new MoveInOutUEDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveInOutUEDetailsWithDefaults() *MoveInOutUEDetails {
	this := MoveInOutUEDetails{}
	return &this
}

// GetMoveInUEs returns the MoveInUEs field value if set, zero value otherwise.
func (o *MoveInOutUEDetails) GetMoveInUEs() []ValTargetUe {
	if o == nil || isNil(o.MoveInUEs) {
		var ret []ValTargetUe
		return ret
	}
	return o.MoveInUEs
}

// GetMoveInUEsOk returns a tuple with the MoveInUEs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveInOutUEDetails) GetMoveInUEsOk() ([]ValTargetUe, bool) {
	if o == nil || isNil(o.MoveInUEs) {
    return nil, false
	}
	return o.MoveInUEs, true
}

// HasMoveInUEs returns a boolean if a field has been set.
func (o *MoveInOutUEDetails) HasMoveInUEs() bool {
	if o != nil && !isNil(o.MoveInUEs) {
		return true
	}

	return false
}

// SetMoveInUEs gets a reference to the given []ValTargetUe and assigns it to the MoveInUEs field.
func (o *MoveInOutUEDetails) SetMoveInUEs(v []ValTargetUe) {
	o.MoveInUEs = v
}

// GetMoveOutUEs returns the MoveOutUEs field value if set, zero value otherwise.
func (o *MoveInOutUEDetails) GetMoveOutUEs() []ValTargetUe {
	if o == nil || isNil(o.MoveOutUEs) {
		var ret []ValTargetUe
		return ret
	}
	return o.MoveOutUEs
}

// GetMoveOutUEsOk returns a tuple with the MoveOutUEs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveInOutUEDetails) GetMoveOutUEsOk() ([]ValTargetUe, bool) {
	if o == nil || isNil(o.MoveOutUEs) {
    return nil, false
	}
	return o.MoveOutUEs, true
}

// HasMoveOutUEs returns a boolean if a field has been set.
func (o *MoveInOutUEDetails) HasMoveOutUEs() bool {
	if o != nil && !isNil(o.MoveOutUEs) {
		return true
	}

	return false
}

// SetMoveOutUEs gets a reference to the given []ValTargetUe and assigns it to the MoveOutUEs field.
func (o *MoveInOutUEDetails) SetMoveOutUEs(v []ValTargetUe) {
	o.MoveOutUEs = v
}

func (o MoveInOutUEDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MoveInUEs) {
		toSerialize["moveInUEs"] = o.MoveInUEs
	}
	if !isNil(o.MoveOutUEs) {
		toSerialize["moveOutUEs"] = o.MoveOutUEs
	}
	return json.Marshal(toSerialize)
}

type NullableMoveInOutUEDetails struct {
	value *MoveInOutUEDetails
	isSet bool
}

func (v NullableMoveInOutUEDetails) Get() *MoveInOutUEDetails {
	return v.value
}

func (v *NullableMoveInOutUEDetails) Set(val *MoveInOutUEDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveInOutUEDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveInOutUEDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveInOutUEDetails(val *MoveInOutUEDetails) *NullableMoveInOutUEDetails {
	return &NullableMoveInOutUEDetails{value: val, isSet: true}
}

func (v NullableMoveInOutUEDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveInOutUEDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


