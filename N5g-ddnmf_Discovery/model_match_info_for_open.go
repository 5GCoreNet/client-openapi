/*
N5g-ddnmf_Discovery API

N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package N5g-ddnmf_Discovery

import (
	"encoding/json"
)

// MatchInfoForOpen Represents a report including a matching result, and the information that can be used for charging purpose for the open discovery type. 
type MatchInfoForOpen struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi"`
	AppId []string `json:"appId"`
}

// NewMatchInfoForOpen instantiates a new MatchInfoForOpen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchInfoForOpen(supi string, appId []string) *MatchInfoForOpen {
	this := MatchInfoForOpen{}
	this.Supi = supi
	this.AppId = appId
	return &this
}

// NewMatchInfoForOpenWithDefaults instantiates a new MatchInfoForOpen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchInfoForOpenWithDefaults() *MatchInfoForOpen {
	this := MatchInfoForOpen{}
	return &this
}

// GetSupi returns the Supi field value
func (o *MatchInfoForOpen) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *MatchInfoForOpen) GetSupiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *MatchInfoForOpen) SetSupi(v string) {
	o.Supi = v
}

// GetAppId returns the AppId field value
func (o *MatchInfoForOpen) GetAppId() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *MatchInfoForOpen) GetAppIdOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AppId, true
}

// SetAppId sets field value
func (o *MatchInfoForOpen) SetAppId(v []string) {
	o.AppId = v
}

func (o MatchInfoForOpen) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supi"] = o.Supi
	}
	if true {
		toSerialize["appId"] = o.AppId
	}
	return json.Marshal(toSerialize)
}

type NullableMatchInfoForOpen struct {
	value *MatchInfoForOpen
	isSet bool
}

func (v NullableMatchInfoForOpen) Get() *MatchInfoForOpen {
	return v.value
}

func (v *NullableMatchInfoForOpen) Set(val *MatchInfoForOpen) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchInfoForOpen) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchInfoForOpen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchInfoForOpen(val *MatchInfoForOpen) *NullableMatchInfoForOpen {
	return &NullableMatchInfoForOpen{value: val, isSet: true}
}

func (v NullableMatchInfoForOpen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchInfoForOpen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


