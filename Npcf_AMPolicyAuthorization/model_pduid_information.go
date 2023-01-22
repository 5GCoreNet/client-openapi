/*
Npcf_AMPolicyAuthorization Service API

PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_AMPolicyAuthorization

import (
	"encoding/json"
	"time"
)

// PduidInformation Contains the ProSe Discovery UE ID and its validity timer.
type PduidInformation struct {
	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry"`
	// Contains the PDUID.
	Pduid string `json:"pduid"`
}

// NewPduidInformation instantiates a new PduidInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduidInformation(expiry time.Time, pduid string) *PduidInformation {
	this := PduidInformation{}
	this.Expiry = expiry
	this.Pduid = pduid
	return &this
}

// NewPduidInformationWithDefaults instantiates a new PduidInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduidInformationWithDefaults() *PduidInformation {
	this := PduidInformation{}
	return &this
}

// GetExpiry returns the Expiry field value
func (o *PduidInformation) GetExpiry() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value
// and a boolean to check if the value has been set.
func (o *PduidInformation) GetExpiryOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Expiry, true
}

// SetExpiry sets field value
func (o *PduidInformation) SetExpiry(v time.Time) {
	o.Expiry = v
}

// GetPduid returns the Pduid field value
func (o *PduidInformation) GetPduid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pduid
}

// GetPduidOk returns a tuple with the Pduid field value
// and a boolean to check if the value has been set.
func (o *PduidInformation) GetPduidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Pduid, true
}

// SetPduid sets field value
func (o *PduidInformation) SetPduid(v string) {
	o.Pduid = v
}

func (o PduidInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["expiry"] = o.Expiry
	}
	if true {
		toSerialize["pduid"] = o.Pduid
	}
	return json.Marshal(toSerialize)
}

type NullablePduidInformation struct {
	value *PduidInformation
	isSet bool
}

func (v NullablePduidInformation) Get() *PduidInformation {
	return v.value
}

func (v *NullablePduidInformation) Set(val *PduidInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePduidInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePduidInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduidInformation(val *PduidInformation) *NullablePduidInformation {
	return &NullablePduidInformation{value: val, isSet: true}
}

func (v NullablePduidInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduidInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


