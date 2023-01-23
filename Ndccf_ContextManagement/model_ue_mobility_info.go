/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// UeMobilityInfo Contains UE mobility information associated with an application.
type UeMobilityInfo struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi"`
	// String providing an application identifier.
	AppId *string `json:"appId,omitempty"`
	UeTrajs []UeTrajectoryInfo `json:"ueTrajs"`
}

// NewUeMobilityInfo instantiates a new UeMobilityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeMobilityInfo(supi string, ueTrajs []UeTrajectoryInfo) *UeMobilityInfo {
	this := UeMobilityInfo{}
	this.Supi = supi
	this.UeTrajs = ueTrajs
	return &this
}

// NewUeMobilityInfoWithDefaults instantiates a new UeMobilityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeMobilityInfoWithDefaults() *UeMobilityInfo {
	this := UeMobilityInfo{}
	return &this
}

// GetSupi returns the Supi field value
func (o *UeMobilityInfo) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *UeMobilityInfo) GetSupiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *UeMobilityInfo) SetSupi(v string) {
	o.Supi = v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *UeMobilityInfo) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobilityInfo) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *UeMobilityInfo) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *UeMobilityInfo) SetAppId(v string) {
	o.AppId = &v
}

// GetUeTrajs returns the UeTrajs field value
func (o *UeMobilityInfo) GetUeTrajs() []UeTrajectoryInfo {
	if o == nil {
		var ret []UeTrajectoryInfo
		return ret
	}

	return o.UeTrajs
}

// GetUeTrajsOk returns a tuple with the UeTrajs field value
// and a boolean to check if the value has been set.
func (o *UeMobilityInfo) GetUeTrajsOk() ([]UeTrajectoryInfo, bool) {
	if o == nil {
    return nil, false
	}
	return o.UeTrajs, true
}

// SetUeTrajs sets field value
func (o *UeMobilityInfo) SetUeTrajs(v []UeTrajectoryInfo) {
	o.UeTrajs = v
}

func (o UeMobilityInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if true {
		toSerialize["ueTrajs"] = o.UeTrajs
	}
	return json.Marshal(toSerialize)
}

type NullableUeMobilityInfo struct {
	value *UeMobilityInfo
	isSet bool
}

func (v NullableUeMobilityInfo) Get() *UeMobilityInfo {
	return v.value
}

func (v *NullableUeMobilityInfo) Set(val *UeMobilityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUeMobilityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUeMobilityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeMobilityInfo(val *UeMobilityInfo) *NullableUeMobilityInfo {
	return &NullableUeMobilityInfo{value: val, isSet: true}
}

func (v NullableUeMobilityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeMobilityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


