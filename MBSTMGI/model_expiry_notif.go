/*
3gpp-mbs-tmgi

API for the allocation, deallocation and management of TMGI(s) for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MBSTMGI

import (
	"encoding/json"
)

// ExpiryNotif Represents MBS TMGI(s) timer expiry notification information.
type ExpiryNotif struct {
	Tmgis []Tmgi `json:"tmgis"`
}

// NewExpiryNotif instantiates a new ExpiryNotif object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpiryNotif(tmgis []Tmgi) *ExpiryNotif {
	this := ExpiryNotif{}
	this.Tmgis = tmgis
	return &this
}

// NewExpiryNotifWithDefaults instantiates a new ExpiryNotif object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpiryNotifWithDefaults() *ExpiryNotif {
	this := ExpiryNotif{}
	return &this
}

// GetTmgis returns the Tmgis field value
func (o *ExpiryNotif) GetTmgis() []Tmgi {
	if o == nil {
		var ret []Tmgi
		return ret
	}

	return o.Tmgis
}

// GetTmgisOk returns a tuple with the Tmgis field value
// and a boolean to check if the value has been set.
func (o *ExpiryNotif) GetTmgisOk() ([]Tmgi, bool) {
	if o == nil {
    return nil, false
	}
	return o.Tmgis, true
}

// SetTmgis sets field value
func (o *ExpiryNotif) SetTmgis(v []Tmgi) {
	o.Tmgis = v
}

func (o ExpiryNotif) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tmgis"] = o.Tmgis
	}
	return json.Marshal(toSerialize)
}

type NullableExpiryNotif struct {
	value *ExpiryNotif
	isSet bool
}

func (v NullableExpiryNotif) Get() *ExpiryNotif {
	return v.value
}

func (v *NullableExpiryNotif) Set(val *ExpiryNotif) {
	v.value = val
	v.isSet = true
}

func (v NullableExpiryNotif) IsSet() bool {
	return v.isSet
}

func (v *NullableExpiryNotif) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpiryNotif(val *ExpiryNotif) *NullableExpiryNotif {
	return &NullableExpiryNotif{value: val, isSet: true}
}

func (v NullableExpiryNotif) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpiryNotif) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

