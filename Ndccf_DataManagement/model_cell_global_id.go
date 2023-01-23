/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// CellGlobalId Contains a Cell Global Identification as defined in 3GPP TS 23.003, clause 4.3.1.
type CellGlobalId struct {
	PlmnId PlmnId `json:"plmnId"`
	Lac string `json:"lac"`
	CellId string `json:"cellId"`
}

// NewCellGlobalId instantiates a new CellGlobalId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCellGlobalId(plmnId PlmnId, lac string, cellId string) *CellGlobalId {
	this := CellGlobalId{}
	this.PlmnId = plmnId
	this.Lac = lac
	this.CellId = cellId
	return &this
}

// NewCellGlobalIdWithDefaults instantiates a new CellGlobalId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCellGlobalIdWithDefaults() *CellGlobalId {
	this := CellGlobalId{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *CellGlobalId) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *CellGlobalId) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *CellGlobalId) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetLac returns the Lac field value
func (o *CellGlobalId) GetLac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lac
}

// GetLacOk returns a tuple with the Lac field value
// and a boolean to check if the value has been set.
func (o *CellGlobalId) GetLacOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Lac, true
}

// SetLac sets field value
func (o *CellGlobalId) SetLac(v string) {
	o.Lac = v
}

// GetCellId returns the CellId field value
func (o *CellGlobalId) GetCellId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CellId
}

// GetCellIdOk returns a tuple with the CellId field value
// and a boolean to check if the value has been set.
func (o *CellGlobalId) GetCellIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CellId, true
}

// SetCellId sets field value
func (o *CellGlobalId) SetCellId(v string) {
	o.CellId = v
}

func (o CellGlobalId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plmnId"] = o.PlmnId
	}
	if true {
		toSerialize["lac"] = o.Lac
	}
	if true {
		toSerialize["cellId"] = o.CellId
	}
	return json.Marshal(toSerialize)
}

type NullableCellGlobalId struct {
	value *CellGlobalId
	isSet bool
}

func (v NullableCellGlobalId) Get() *CellGlobalId {
	return v.value
}

func (v *NullableCellGlobalId) Set(val *CellGlobalId) {
	v.value = val
	v.isSet = true
}

func (v NullableCellGlobalId) IsSet() bool {
	return v.isSet
}

func (v *NullableCellGlobalId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCellGlobalId(val *CellGlobalId) *NullableCellGlobalId {
	return &NullableCellGlobalId{value: val, isSet: true}
}

func (v NullableCellGlobalId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCellGlobalId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


