/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MBSUserDataIngestSession

import (
	"encoding/json"
)

// NcgiTai List of NR cell ids, with their pertaining TAIs
type NcgiTai struct {
	Tai Tai `json:"tai"`
	// List of List of NR cell ids
	CellList []Ncgi `json:"cellList"`
}

// NewNcgiTai instantiates a new NcgiTai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNcgiTai(tai Tai, cellList []Ncgi) *NcgiTai {
	this := NcgiTai{}
	this.Tai = tai
	this.CellList = cellList
	return &this
}

// NewNcgiTaiWithDefaults instantiates a new NcgiTai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNcgiTaiWithDefaults() *NcgiTai {
	this := NcgiTai{}
	return &this
}

// GetTai returns the Tai field value
func (o *NcgiTai) GetTai() Tai {
	if o == nil {
		var ret Tai
		return ret
	}

	return o.Tai
}

// GetTaiOk returns a tuple with the Tai field value
// and a boolean to check if the value has been set.
func (o *NcgiTai) GetTaiOk() (*Tai, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tai, true
}

// SetTai sets field value
func (o *NcgiTai) SetTai(v Tai) {
	o.Tai = v
}

// GetCellList returns the CellList field value
func (o *NcgiTai) GetCellList() []Ncgi {
	if o == nil {
		var ret []Ncgi
		return ret
	}

	return o.CellList
}

// GetCellListOk returns a tuple with the CellList field value
// and a boolean to check if the value has been set.
func (o *NcgiTai) GetCellListOk() ([]Ncgi, bool) {
	if o == nil {
    return nil, false
	}
	return o.CellList, true
}

// SetCellList sets field value
func (o *NcgiTai) SetCellList(v []Ncgi) {
	o.CellList = v
}

func (o NcgiTai) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tai"] = o.Tai
	}
	if true {
		toSerialize["cellList"] = o.CellList
	}
	return json.Marshal(toSerialize)
}

type NullableNcgiTai struct {
	value *NcgiTai
	isSet bool
}

func (v NullableNcgiTai) Get() *NcgiTai {
	return v.value
}

func (v *NullableNcgiTai) Set(val *NcgiTai) {
	v.value = val
	v.isSet = true
}

func (v NullableNcgiTai) IsSet() bool {
	return v.isSet
}

func (v *NullableNcgiTai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNcgiTai(val *NcgiTai) *NullableNcgiTai {
	return &NullableNcgiTai{value: val, isSet: true}
}

func (v NullableNcgiTai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNcgiTai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


