/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Location

import (
	"encoding/json"
	"os"
)

// LcsBroadcastAssistanceTypesData struct for LcsBroadcastAssistanceTypesData
type LcsBroadcastAssistanceTypesData struct {
	// string with format 'binary' as defined in OpenAPI.
	LocationAssistanceType *os.File `json:"locationAssistanceType"`
}

// NewLcsBroadcastAssistanceTypesData instantiates a new LcsBroadcastAssistanceTypesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLcsBroadcastAssistanceTypesData(locationAssistanceType *os.File) *LcsBroadcastAssistanceTypesData {
	this := LcsBroadcastAssistanceTypesData{}
	this.LocationAssistanceType = locationAssistanceType
	return &this
}

// NewLcsBroadcastAssistanceTypesDataWithDefaults instantiates a new LcsBroadcastAssistanceTypesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLcsBroadcastAssistanceTypesDataWithDefaults() *LcsBroadcastAssistanceTypesData {
	this := LcsBroadcastAssistanceTypesData{}
	return &this
}

// GetLocationAssistanceType returns the LocationAssistanceType field value
func (o *LcsBroadcastAssistanceTypesData) GetLocationAssistanceType() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.LocationAssistanceType
}

// GetLocationAssistanceTypeOk returns a tuple with the LocationAssistanceType field value
// and a boolean to check if the value has been set.
func (o *LcsBroadcastAssistanceTypesData) GetLocationAssistanceTypeOk() (**os.File, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocationAssistanceType, true
}

// SetLocationAssistanceType sets field value
func (o *LcsBroadcastAssistanceTypesData) SetLocationAssistanceType(v *os.File) {
	o.LocationAssistanceType = v
}

func (o LcsBroadcastAssistanceTypesData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["locationAssistanceType"] = o.LocationAssistanceType
	}
	return json.Marshal(toSerialize)
}

type NullableLcsBroadcastAssistanceTypesData struct {
	value *LcsBroadcastAssistanceTypesData
	isSet bool
}

func (v NullableLcsBroadcastAssistanceTypesData) Get() *LcsBroadcastAssistanceTypesData {
	return v.value
}

func (v *NullableLcsBroadcastAssistanceTypesData) Set(val *LcsBroadcastAssistanceTypesData) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsBroadcastAssistanceTypesData) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsBroadcastAssistanceTypesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsBroadcastAssistanceTypesData(val *LcsBroadcastAssistanceTypesData) *NullableLcsBroadcastAssistanceTypesData {
	return &NullableLcsBroadcastAssistanceTypesData{value: val, isSet: true}
}

func (v NullableLcsBroadcastAssistanceTypesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsBroadcastAssistanceTypesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


