/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
)

// DnaiInformation DNAI Information
type DnaiInformation struct {
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai string `json:"dnai"`
	NoDnaiChangeInd *bool `json:"noDnaiChangeInd,omitempty"`
	NoLocalPsaChangeInd *bool `json:"noLocalPsaChangeInd,omitempty"`
}

// NewDnaiInformation instantiates a new DnaiInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnaiInformation(dnai string) *DnaiInformation {
	this := DnaiInformation{}
	this.Dnai = dnai
	return &this
}

// NewDnaiInformationWithDefaults instantiates a new DnaiInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnaiInformationWithDefaults() *DnaiInformation {
	this := DnaiInformation{}
	return &this
}

// GetDnai returns the Dnai field value
func (o *DnaiInformation) GetDnai() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnai
}

// GetDnaiOk returns a tuple with the Dnai field value
// and a boolean to check if the value has been set.
func (o *DnaiInformation) GetDnaiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Dnai, true
}

// SetDnai sets field value
func (o *DnaiInformation) SetDnai(v string) {
	o.Dnai = v
}

// GetNoDnaiChangeInd returns the NoDnaiChangeInd field value if set, zero value otherwise.
func (o *DnaiInformation) GetNoDnaiChangeInd() bool {
	if o == nil || isNil(o.NoDnaiChangeInd) {
		var ret bool
		return ret
	}
	return *o.NoDnaiChangeInd
}

// GetNoDnaiChangeIndOk returns a tuple with the NoDnaiChangeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaiInformation) GetNoDnaiChangeIndOk() (*bool, bool) {
	if o == nil || isNil(o.NoDnaiChangeInd) {
    return nil, false
	}
	return o.NoDnaiChangeInd, true
}

// HasNoDnaiChangeInd returns a boolean if a field has been set.
func (o *DnaiInformation) HasNoDnaiChangeInd() bool {
	if o != nil && !isNil(o.NoDnaiChangeInd) {
		return true
	}

	return false
}

// SetNoDnaiChangeInd gets a reference to the given bool and assigns it to the NoDnaiChangeInd field.
func (o *DnaiInformation) SetNoDnaiChangeInd(v bool) {
	o.NoDnaiChangeInd = &v
}

// GetNoLocalPsaChangeInd returns the NoLocalPsaChangeInd field value if set, zero value otherwise.
func (o *DnaiInformation) GetNoLocalPsaChangeInd() bool {
	if o == nil || isNil(o.NoLocalPsaChangeInd) {
		var ret bool
		return ret
	}
	return *o.NoLocalPsaChangeInd
}

// GetNoLocalPsaChangeIndOk returns a tuple with the NoLocalPsaChangeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnaiInformation) GetNoLocalPsaChangeIndOk() (*bool, bool) {
	if o == nil || isNil(o.NoLocalPsaChangeInd) {
    return nil, false
	}
	return o.NoLocalPsaChangeInd, true
}

// HasNoLocalPsaChangeInd returns a boolean if a field has been set.
func (o *DnaiInformation) HasNoLocalPsaChangeInd() bool {
	if o != nil && !isNil(o.NoLocalPsaChangeInd) {
		return true
	}

	return false
}

// SetNoLocalPsaChangeInd gets a reference to the given bool and assigns it to the NoLocalPsaChangeInd field.
func (o *DnaiInformation) SetNoLocalPsaChangeInd(v bool) {
	o.NoLocalPsaChangeInd = &v
}

func (o DnaiInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnai"] = o.Dnai
	}
	if !isNil(o.NoDnaiChangeInd) {
		toSerialize["noDnaiChangeInd"] = o.NoDnaiChangeInd
	}
	if !isNil(o.NoLocalPsaChangeInd) {
		toSerialize["noLocalPsaChangeInd"] = o.NoLocalPsaChangeInd
	}
	return json.Marshal(toSerialize)
}

type NullableDnaiInformation struct {
	value *DnaiInformation
	isSet bool
}

func (v NullableDnaiInformation) Get() *DnaiInformation {
	return v.value
}

func (v *NullableDnaiInformation) Set(val *DnaiInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaiInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaiInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaiInformation(val *DnaiInformation) *NullableDnaiInformation {
	return &NullableDnaiInformation{value: val, isSet: true}
}

func (v NullableDnaiInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaiInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


