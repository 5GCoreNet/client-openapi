/*
3gpp-time-sync-exposure

API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package TimeSyncExposure

import (
	"encoding/json"
)

// PtpCapabilitiesPerUe Contains the supported PTP capabilities per UE.
type PtpCapabilitiesPerUe struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi"`
	PtpCaps []EventFilter `json:"ptpCaps"`
}

// NewPtpCapabilitiesPerUe instantiates a new PtpCapabilitiesPerUe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPtpCapabilitiesPerUe(gpsi string, ptpCaps []EventFilter) *PtpCapabilitiesPerUe {
	this := PtpCapabilitiesPerUe{}
	this.Gpsi = gpsi
	this.PtpCaps = ptpCaps
	return &this
}

// NewPtpCapabilitiesPerUeWithDefaults instantiates a new PtpCapabilitiesPerUe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPtpCapabilitiesPerUeWithDefaults() *PtpCapabilitiesPerUe {
	this := PtpCapabilitiesPerUe{}
	return &this
}

// GetGpsi returns the Gpsi field value
func (o *PtpCapabilitiesPerUe) GetGpsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value
// and a boolean to check if the value has been set.
func (o *PtpCapabilitiesPerUe) GetGpsiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Gpsi, true
}

// SetGpsi sets field value
func (o *PtpCapabilitiesPerUe) SetGpsi(v string) {
	o.Gpsi = v
}

// GetPtpCaps returns the PtpCaps field value
func (o *PtpCapabilitiesPerUe) GetPtpCaps() []EventFilter {
	if o == nil {
		var ret []EventFilter
		return ret
	}

	return o.PtpCaps
}

// GetPtpCapsOk returns a tuple with the PtpCaps field value
// and a boolean to check if the value has been set.
func (o *PtpCapabilitiesPerUe) GetPtpCapsOk() ([]EventFilter, bool) {
	if o == nil {
    return nil, false
	}
	return o.PtpCaps, true
}

// SetPtpCaps sets field value
func (o *PtpCapabilitiesPerUe) SetPtpCaps(v []EventFilter) {
	o.PtpCaps = v
}

func (o PtpCapabilitiesPerUe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["gpsi"] = o.Gpsi
	}
	if true {
		toSerialize["ptpCaps"] = o.PtpCaps
	}
	return json.Marshal(toSerialize)
}

type NullablePtpCapabilitiesPerUe struct {
	value *PtpCapabilitiesPerUe
	isSet bool
}

func (v NullablePtpCapabilitiesPerUe) Get() *PtpCapabilitiesPerUe {
	return v.value
}

func (v *NullablePtpCapabilitiesPerUe) Set(val *PtpCapabilitiesPerUe) {
	v.value = val
	v.isSet = true
}

func (v NullablePtpCapabilitiesPerUe) IsSet() bool {
	return v.isSet
}

func (v *NullablePtpCapabilitiesPerUe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePtpCapabilitiesPerUe(val *PtpCapabilitiesPerUe) *NullablePtpCapabilitiesPerUe {
	return &NullablePtpCapabilitiesPerUe{value: val, isSet: true}
}

func (v NullablePtpCapabilitiesPerUe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePtpCapabilitiesPerUe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


