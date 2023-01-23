/*
3gpp-nidd-configuration-trigger

API for NIDD Configuration Trigger.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NIDDConfigurationTrigger

import (
	"encoding/json"
)

// NiddConfigurationTrigger Represents a NIDD configuration trigger.
type NiddConfigurationTrigger struct {
	// Identifies the trigger receiving entity.
	AfId string `json:"afId"`
	// Identifies the trigger sending entity.
	NefId string `json:"nefId"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat"`
}

// NewNiddConfigurationTrigger instantiates a new NiddConfigurationTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiddConfigurationTrigger(afId string, nefId string, gpsi string, suppFeat string) *NiddConfigurationTrigger {
	this := NiddConfigurationTrigger{}
	this.AfId = afId
	this.NefId = nefId
	this.Gpsi = gpsi
	this.SuppFeat = suppFeat
	return &this
}

// NewNiddConfigurationTriggerWithDefaults instantiates a new NiddConfigurationTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiddConfigurationTriggerWithDefaults() *NiddConfigurationTrigger {
	this := NiddConfigurationTrigger{}
	return &this
}

// GetAfId returns the AfId field value
func (o *NiddConfigurationTrigger) GetAfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value
// and a boolean to check if the value has been set.
func (o *NiddConfigurationTrigger) GetAfIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AfId, true
}

// SetAfId sets field value
func (o *NiddConfigurationTrigger) SetAfId(v string) {
	o.AfId = v
}

// GetNefId returns the NefId field value
func (o *NiddConfigurationTrigger) GetNefId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NefId
}

// GetNefIdOk returns a tuple with the NefId field value
// and a boolean to check if the value has been set.
func (o *NiddConfigurationTrigger) GetNefIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NefId, true
}

// SetNefId sets field value
func (o *NiddConfigurationTrigger) SetNefId(v string) {
	o.NefId = v
}

// GetGpsi returns the Gpsi field value
func (o *NiddConfigurationTrigger) GetGpsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value
// and a boolean to check if the value has been set.
func (o *NiddConfigurationTrigger) GetGpsiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Gpsi, true
}

// SetGpsi sets field value
func (o *NiddConfigurationTrigger) SetGpsi(v string) {
	o.Gpsi = v
}

// GetSuppFeat returns the SuppFeat field value
func (o *NiddConfigurationTrigger) GetSuppFeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value
// and a boolean to check if the value has been set.
func (o *NiddConfigurationTrigger) GetSuppFeatOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SuppFeat, true
}

// SetSuppFeat sets field value
func (o *NiddConfigurationTrigger) SetSuppFeat(v string) {
	o.SuppFeat = v
}

func (o NiddConfigurationTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["afId"] = o.AfId
	}
	if true {
		toSerialize["nefId"] = o.NefId
	}
	if true {
		toSerialize["gpsi"] = o.Gpsi
	}
	if true {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableNiddConfigurationTrigger struct {
	value *NiddConfigurationTrigger
	isSet bool
}

func (v NullableNiddConfigurationTrigger) Get() *NiddConfigurationTrigger {
	return v.value
}

func (v *NullableNiddConfigurationTrigger) Set(val *NiddConfigurationTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableNiddConfigurationTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableNiddConfigurationTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiddConfigurationTrigger(val *NiddConfigurationTrigger) *NullableNiddConfigurationTrigger {
	return &NullableNiddConfigurationTrigger{value: val, isSet: true}
}

func (v NullableNiddConfigurationTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiddConfigurationTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


