/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// NrOperatorCellDuSingleAllOf struct for NrOperatorCellDuSingleAllOf
type NrOperatorCellDuSingleAllOf struct {
	CellLocalId *int32 `json:"cellLocalId,omitempty"`
	AdministrativeState *AdministrativeState `json:"administrativeState,omitempty"`
	PlmnInfoList []PlmnInfo `json:"plmnInfoList,omitempty"`
	NrTac *int32 `json:"nrTac,omitempty"`
}

// NewNrOperatorCellDuSingleAllOf instantiates a new NrOperatorCellDuSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrOperatorCellDuSingleAllOf() *NrOperatorCellDuSingleAllOf {
	this := NrOperatorCellDuSingleAllOf{}
	return &this
}

// NewNrOperatorCellDuSingleAllOfWithDefaults instantiates a new NrOperatorCellDuSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrOperatorCellDuSingleAllOfWithDefaults() *NrOperatorCellDuSingleAllOf {
	this := NrOperatorCellDuSingleAllOf{}
	return &this
}

// GetCellLocalId returns the CellLocalId field value if set, zero value otherwise.
func (o *NrOperatorCellDuSingleAllOf) GetCellLocalId() int32 {
	if o == nil || isNil(o.CellLocalId) {
		var ret int32
		return ret
	}
	return *o.CellLocalId
}

// GetCellLocalIdOk returns a tuple with the CellLocalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrOperatorCellDuSingleAllOf) GetCellLocalIdOk() (*int32, bool) {
	if o == nil || isNil(o.CellLocalId) {
    return nil, false
	}
	return o.CellLocalId, true
}

// HasCellLocalId returns a boolean if a field has been set.
func (o *NrOperatorCellDuSingleAllOf) HasCellLocalId() bool {
	if o != nil && !isNil(o.CellLocalId) {
		return true
	}

	return false
}

// SetCellLocalId gets a reference to the given int32 and assigns it to the CellLocalId field.
func (o *NrOperatorCellDuSingleAllOf) SetCellLocalId(v int32) {
	o.CellLocalId = &v
}

// GetAdministrativeState returns the AdministrativeState field value if set, zero value otherwise.
func (o *NrOperatorCellDuSingleAllOf) GetAdministrativeState() AdministrativeState {
	if o == nil || isNil(o.AdministrativeState) {
		var ret AdministrativeState
		return ret
	}
	return *o.AdministrativeState
}

// GetAdministrativeStateOk returns a tuple with the AdministrativeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrOperatorCellDuSingleAllOf) GetAdministrativeStateOk() (*AdministrativeState, bool) {
	if o == nil || isNil(o.AdministrativeState) {
    return nil, false
	}
	return o.AdministrativeState, true
}

// HasAdministrativeState returns a boolean if a field has been set.
func (o *NrOperatorCellDuSingleAllOf) HasAdministrativeState() bool {
	if o != nil && !isNil(o.AdministrativeState) {
		return true
	}

	return false
}

// SetAdministrativeState gets a reference to the given AdministrativeState and assigns it to the AdministrativeState field.
func (o *NrOperatorCellDuSingleAllOf) SetAdministrativeState(v AdministrativeState) {
	o.AdministrativeState = &v
}

// GetPlmnInfoList returns the PlmnInfoList field value if set, zero value otherwise.
func (o *NrOperatorCellDuSingleAllOf) GetPlmnInfoList() []PlmnInfo {
	if o == nil || isNil(o.PlmnInfoList) {
		var ret []PlmnInfo
		return ret
	}
	return o.PlmnInfoList
}

// GetPlmnInfoListOk returns a tuple with the PlmnInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrOperatorCellDuSingleAllOf) GetPlmnInfoListOk() ([]PlmnInfo, bool) {
	if o == nil || isNil(o.PlmnInfoList) {
    return nil, false
	}
	return o.PlmnInfoList, true
}

// HasPlmnInfoList returns a boolean if a field has been set.
func (o *NrOperatorCellDuSingleAllOf) HasPlmnInfoList() bool {
	if o != nil && !isNil(o.PlmnInfoList) {
		return true
	}

	return false
}

// SetPlmnInfoList gets a reference to the given []PlmnInfo and assigns it to the PlmnInfoList field.
func (o *NrOperatorCellDuSingleAllOf) SetPlmnInfoList(v []PlmnInfo) {
	o.PlmnInfoList = v
}

// GetNrTac returns the NrTac field value if set, zero value otherwise.
func (o *NrOperatorCellDuSingleAllOf) GetNrTac() int32 {
	if o == nil || isNil(o.NrTac) {
		var ret int32
		return ret
	}
	return *o.NrTac
}

// GetNrTacOk returns a tuple with the NrTac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrOperatorCellDuSingleAllOf) GetNrTacOk() (*int32, bool) {
	if o == nil || isNil(o.NrTac) {
    return nil, false
	}
	return o.NrTac, true
}

// HasNrTac returns a boolean if a field has been set.
func (o *NrOperatorCellDuSingleAllOf) HasNrTac() bool {
	if o != nil && !isNil(o.NrTac) {
		return true
	}

	return false
}

// SetNrTac gets a reference to the given int32 and assigns it to the NrTac field.
func (o *NrOperatorCellDuSingleAllOf) SetNrTac(v int32) {
	o.NrTac = &v
}

func (o NrOperatorCellDuSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CellLocalId) {
		toSerialize["cellLocalId"] = o.CellLocalId
	}
	if !isNil(o.AdministrativeState) {
		toSerialize["administrativeState"] = o.AdministrativeState
	}
	if !isNil(o.PlmnInfoList) {
		toSerialize["plmnInfoList"] = o.PlmnInfoList
	}
	if !isNil(o.NrTac) {
		toSerialize["nrTac"] = o.NrTac
	}
	return json.Marshal(toSerialize)
}

type NullableNrOperatorCellDuSingleAllOf struct {
	value *NrOperatorCellDuSingleAllOf
	isSet bool
}

func (v NullableNrOperatorCellDuSingleAllOf) Get() *NrOperatorCellDuSingleAllOf {
	return v.value
}

func (v *NullableNrOperatorCellDuSingleAllOf) Set(val *NrOperatorCellDuSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNrOperatorCellDuSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNrOperatorCellDuSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrOperatorCellDuSingleAllOf(val *NrOperatorCellDuSingleAllOf) *NullableNrOperatorCellDuSingleAllOf {
	return &NullableNrOperatorCellDuSingleAllOf{value: val, isSet: true}
}

func (v NullableNrOperatorCellDuSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrOperatorCellDuSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


