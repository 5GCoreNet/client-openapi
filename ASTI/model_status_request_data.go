/*
3gpp-asti

API for ASTI.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ASTI

import (
	"encoding/json"
)

// StatusRequestData Contains the parameters for retrieval of the status of the access stratum time distribution for a list of UEs. 
type StatusRequestData struct {
	Gpsis []string `json:"gpsis"`
}

// NewStatusRequestData instantiates a new StatusRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusRequestData(gpsis []string) *StatusRequestData {
	this := StatusRequestData{}
	this.Gpsis = gpsis
	return &this
}

// NewStatusRequestDataWithDefaults instantiates a new StatusRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusRequestDataWithDefaults() *StatusRequestData {
	this := StatusRequestData{}
	return &this
}

// GetGpsis returns the Gpsis field value
func (o *StatusRequestData) GetGpsis() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Gpsis
}

// GetGpsisOk returns a tuple with the Gpsis field value
// and a boolean to check if the value has been set.
func (o *StatusRequestData) GetGpsisOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Gpsis, true
}

// SetGpsis sets field value
func (o *StatusRequestData) SetGpsis(v []string) {
	o.Gpsis = v
}

func (o StatusRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["gpsis"] = o.Gpsis
	}
	return json.Marshal(toSerialize)
}

type NullableStatusRequestData struct {
	value *StatusRequestData
	isSet bool
}

func (v NullableStatusRequestData) Get() *StatusRequestData {
	return v.value
}

func (v *NullableStatusRequestData) Set(val *StatusRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusRequestData(val *StatusRequestData) *NullableStatusRequestData {
	return &NullableStatusRequestData{value: val, isSet: true}
}

func (v NullableStatusRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


