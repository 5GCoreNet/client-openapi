/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// UeCommunicationInfo Contains UE communication information associated with an application.
type UeCommunicationInfo struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	InterGroupId *string `json:"interGroupId,omitempty"`
	// String providing an application identifier.
	AppId *string `json:"appId,omitempty"`
	Comms []CommunicationCollection `json:"comms"`
}

// NewUeCommunicationInfo instantiates a new UeCommunicationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeCommunicationInfo(comms []CommunicationCollection) *UeCommunicationInfo {
	this := UeCommunicationInfo{}
	this.Comms = comms
	return &this
}

// NewUeCommunicationInfoWithDefaults instantiates a new UeCommunicationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeCommunicationInfoWithDefaults() *UeCommunicationInfo {
	this := UeCommunicationInfo{}
	return &this
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *UeCommunicationInfo) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunicationInfo) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
    return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *UeCommunicationInfo) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *UeCommunicationInfo) SetSupi(v string) {
	o.Supi = &v
}

// GetInterGroupId returns the InterGroupId field value if set, zero value otherwise.
func (o *UeCommunicationInfo) GetInterGroupId() string {
	if o == nil || isNil(o.InterGroupId) {
		var ret string
		return ret
	}
	return *o.InterGroupId
}

// GetInterGroupIdOk returns a tuple with the InterGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunicationInfo) GetInterGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.InterGroupId) {
    return nil, false
	}
	return o.InterGroupId, true
}

// HasInterGroupId returns a boolean if a field has been set.
func (o *UeCommunicationInfo) HasInterGroupId() bool {
	if o != nil && !isNil(o.InterGroupId) {
		return true
	}

	return false
}

// SetInterGroupId gets a reference to the given string and assigns it to the InterGroupId field.
func (o *UeCommunicationInfo) SetInterGroupId(v string) {
	o.InterGroupId = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *UeCommunicationInfo) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeCommunicationInfo) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *UeCommunicationInfo) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *UeCommunicationInfo) SetAppId(v string) {
	o.AppId = &v
}

// GetComms returns the Comms field value
func (o *UeCommunicationInfo) GetComms() []CommunicationCollection {
	if o == nil {
		var ret []CommunicationCollection
		return ret
	}

	return o.Comms
}

// GetCommsOk returns a tuple with the Comms field value
// and a boolean to check if the value has been set.
func (o *UeCommunicationInfo) GetCommsOk() ([]CommunicationCollection, bool) {
	if o == nil {
    return nil, false
	}
	return o.Comms, true
}

// SetComms sets field value
func (o *UeCommunicationInfo) SetComms(v []CommunicationCollection) {
	o.Comms = v
}

func (o UeCommunicationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.InterGroupId) {
		toSerialize["interGroupId"] = o.InterGroupId
	}
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if true {
		toSerialize["comms"] = o.Comms
	}
	return json.Marshal(toSerialize)
}

type NullableUeCommunicationInfo struct {
	value *UeCommunicationInfo
	isSet bool
}

func (v NullableUeCommunicationInfo) Get() *UeCommunicationInfo {
	return v.value
}

func (v *NullableUeCommunicationInfo) Set(val *UeCommunicationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUeCommunicationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUeCommunicationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeCommunicationInfo(val *UeCommunicationInfo) *NullableUeCommunicationInfo {
	return &NullableUeCommunicationInfo{value: val, isSet: true}
}

func (v NullableUeCommunicationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeCommunicationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


