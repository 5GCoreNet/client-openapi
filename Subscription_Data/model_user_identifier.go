/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Subscription_Data

import (
	"encoding/json"
	"time"
)

// UserIdentifier Represents the user identifier.
type UserIdentifier struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
}

// NewUserIdentifier instantiates a new UserIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentifier(supi string) *UserIdentifier {
	this := UserIdentifier{}
	this.Supi = supi
	return &this
}

// NewUserIdentifierWithDefaults instantiates a new UserIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentifierWithDefaults() *UserIdentifier {
	this := UserIdentifier{}
	return &this
}

// GetSupi returns the Supi field value
func (o *UserIdentifier) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *UserIdentifier) GetSupiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *UserIdentifier) SetSupi(v string) {
	o.Supi = v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *UserIdentifier) GetGpsi() string {
	if o == nil || isNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentifier) GetGpsiOk() (*string, bool) {
	if o == nil || isNil(o.Gpsi) {
    return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *UserIdentifier) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *UserIdentifier) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *UserIdentifier) GetValidityTime() time.Time {
	if o == nil || isNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentifier) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ValidityTime) {
    return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *UserIdentifier) HasValidityTime() bool {
	if o != nil && !isNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *UserIdentifier) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

func (o UserIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !isNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	return json.Marshal(toSerialize)
}

type NullableUserIdentifier struct {
	value *UserIdentifier
	isSet bool
}

func (v NullableUserIdentifier) Get() *UserIdentifier {
	return v.value
}

func (v *NullableUserIdentifier) Set(val *UserIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentifier(val *UserIdentifier) *NullableUserIdentifier {
	return &NullableUserIdentifier{value: val, isSet: true}
}

func (v NullableUserIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


