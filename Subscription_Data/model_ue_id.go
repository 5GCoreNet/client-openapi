/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
)

// UeId struct for UeId
type UeId struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi"`
	GpsiList []string `json:"gpsiList,omitempty"`
}

// NewUeId instantiates a new UeId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeId(supi string) *UeId {
	this := UeId{}
	this.Supi = supi
	return &this
}

// NewUeIdWithDefaults instantiates a new UeId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeIdWithDefaults() *UeId {
	this := UeId{}
	return &this
}

// GetSupi returns the Supi field value
func (o *UeId) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *UeId) GetSupiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *UeId) SetSupi(v string) {
	o.Supi = v
}

// GetGpsiList returns the GpsiList field value if set, zero value otherwise.
func (o *UeId) GetGpsiList() []string {
	if o == nil || isNil(o.GpsiList) {
		var ret []string
		return ret
	}
	return o.GpsiList
}

// GetGpsiListOk returns a tuple with the GpsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeId) GetGpsiListOk() ([]string, bool) {
	if o == nil || isNil(o.GpsiList) {
    return nil, false
	}
	return o.GpsiList, true
}

// HasGpsiList returns a boolean if a field has been set.
func (o *UeId) HasGpsiList() bool {
	if o != nil && !isNil(o.GpsiList) {
		return true
	}

	return false
}

// SetGpsiList gets a reference to the given []string and assigns it to the GpsiList field.
func (o *UeId) SetGpsiList(v []string) {
	o.GpsiList = v
}

func (o UeId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.GpsiList) {
		toSerialize["gpsiList"] = o.GpsiList
	}
	return json.Marshal(toSerialize)
}

type NullableUeId struct {
	value *UeId
	isSet bool
}

func (v NullableUeId) Get() *UeId {
	return v.value
}

func (v *NullableUeId) Set(val *UeId) {
	v.value = val
	v.isSet = true
}

func (v NullableUeId) IsSet() bool {
	return v.isSet
}

func (v *NullableUeId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeId(val *UeId) *NullableUeId {
	return &NullableUeId{value: val, isSet: true}
}

func (v NullableUeId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


