/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// UeAnalyticsContextDescriptor Contains information about available UE related analytics contexts.
type UeAnalyticsContextDescriptor struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi"`
	// List of analytics types for which UE related analytics contexts can be retrieved. 
	AnaTypes []NwdafEvent `json:"anaTypes"`
}

// NewUeAnalyticsContextDescriptor instantiates a new UeAnalyticsContextDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeAnalyticsContextDescriptor(supi string, anaTypes []NwdafEvent) *UeAnalyticsContextDescriptor {
	this := UeAnalyticsContextDescriptor{}
	this.Supi = supi
	this.AnaTypes = anaTypes
	return &this
}

// NewUeAnalyticsContextDescriptorWithDefaults instantiates a new UeAnalyticsContextDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeAnalyticsContextDescriptorWithDefaults() *UeAnalyticsContextDescriptor {
	this := UeAnalyticsContextDescriptor{}
	return &this
}

// GetSupi returns the Supi field value
func (o *UeAnalyticsContextDescriptor) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *UeAnalyticsContextDescriptor) GetSupiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *UeAnalyticsContextDescriptor) SetSupi(v string) {
	o.Supi = v
}

// GetAnaTypes returns the AnaTypes field value
func (o *UeAnalyticsContextDescriptor) GetAnaTypes() []NwdafEvent {
	if o == nil {
		var ret []NwdafEvent
		return ret
	}

	return o.AnaTypes
}

// GetAnaTypesOk returns a tuple with the AnaTypes field value
// and a boolean to check if the value has been set.
func (o *UeAnalyticsContextDescriptor) GetAnaTypesOk() ([]NwdafEvent, bool) {
	if o == nil {
    return nil, false
	}
	return o.AnaTypes, true
}

// SetAnaTypes sets field value
func (o *UeAnalyticsContextDescriptor) SetAnaTypes(v []NwdafEvent) {
	o.AnaTypes = v
}

func (o UeAnalyticsContextDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supi"] = o.Supi
	}
	if true {
		toSerialize["anaTypes"] = o.AnaTypes
	}
	return json.Marshal(toSerialize)
}

type NullableUeAnalyticsContextDescriptor struct {
	value *UeAnalyticsContextDescriptor
	isSet bool
}

func (v NullableUeAnalyticsContextDescriptor) Get() *UeAnalyticsContextDescriptor {
	return v.value
}

func (v *NullableUeAnalyticsContextDescriptor) Set(val *UeAnalyticsContextDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableUeAnalyticsContextDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableUeAnalyticsContextDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeAnalyticsContextDescriptor(val *UeAnalyticsContextDescriptor) *NullableUeAnalyticsContextDescriptor {
	return &NullableUeAnalyticsContextDescriptor{value: val, isSet: true}
}

func (v NullableUeAnalyticsContextDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeAnalyticsContextDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


