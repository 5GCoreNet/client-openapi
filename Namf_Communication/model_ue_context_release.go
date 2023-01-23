/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// UEContextRelease Data within a Release UE Context request
type UEContextRelease struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	UnauthenticatedSupi *bool `json:"unauthenticatedSupi,omitempty"`
	NgapCause NgApCause `json:"ngapCause"`
}

// NewUEContextRelease instantiates a new UEContextRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUEContextRelease(ngapCause NgApCause) *UEContextRelease {
	this := UEContextRelease{}
	var unauthenticatedSupi bool = false
	this.UnauthenticatedSupi = &unauthenticatedSupi
	this.NgapCause = ngapCause
	return &this
}

// NewUEContextReleaseWithDefaults instantiates a new UEContextRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUEContextReleaseWithDefaults() *UEContextRelease {
	this := UEContextRelease{}
	var unauthenticatedSupi bool = false
	this.UnauthenticatedSupi = &unauthenticatedSupi
	return &this
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *UEContextRelease) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UEContextRelease) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
    return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *UEContextRelease) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *UEContextRelease) SetSupi(v string) {
	o.Supi = &v
}

// GetUnauthenticatedSupi returns the UnauthenticatedSupi field value if set, zero value otherwise.
func (o *UEContextRelease) GetUnauthenticatedSupi() bool {
	if o == nil || isNil(o.UnauthenticatedSupi) {
		var ret bool
		return ret
	}
	return *o.UnauthenticatedSupi
}

// GetUnauthenticatedSupiOk returns a tuple with the UnauthenticatedSupi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UEContextRelease) GetUnauthenticatedSupiOk() (*bool, bool) {
	if o == nil || isNil(o.UnauthenticatedSupi) {
    return nil, false
	}
	return o.UnauthenticatedSupi, true
}

// HasUnauthenticatedSupi returns a boolean if a field has been set.
func (o *UEContextRelease) HasUnauthenticatedSupi() bool {
	if o != nil && !isNil(o.UnauthenticatedSupi) {
		return true
	}

	return false
}

// SetUnauthenticatedSupi gets a reference to the given bool and assigns it to the UnauthenticatedSupi field.
func (o *UEContextRelease) SetUnauthenticatedSupi(v bool) {
	o.UnauthenticatedSupi = &v
}

// GetNgapCause returns the NgapCause field value
func (o *UEContextRelease) GetNgapCause() NgApCause {
	if o == nil {
		var ret NgApCause
		return ret
	}

	return o.NgapCause
}

// GetNgapCauseOk returns a tuple with the NgapCause field value
// and a boolean to check if the value has been set.
func (o *UEContextRelease) GetNgapCauseOk() (*NgApCause, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NgapCause, true
}

// SetNgapCause sets field value
func (o *UEContextRelease) SetNgapCause(v NgApCause) {
	o.NgapCause = v
}

func (o UEContextRelease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.UnauthenticatedSupi) {
		toSerialize["unauthenticatedSupi"] = o.UnauthenticatedSupi
	}
	if true {
		toSerialize["ngapCause"] = o.NgapCause
	}
	return json.Marshal(toSerialize)
}

type NullableUEContextRelease struct {
	value *UEContextRelease
	isSet bool
}

func (v NullableUEContextRelease) Get() *UEContextRelease {
	return v.value
}

func (v *NullableUEContextRelease) Set(val *UEContextRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableUEContextRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableUEContextRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUEContextRelease(val *UEContextRelease) *NullableUEContextRelease {
	return &NullableUEContextRelease{value: val, isSet: true}
}

func (v NullableUEContextRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUEContextRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


