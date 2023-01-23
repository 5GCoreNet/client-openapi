/*
Nucmf_Provisioning

UCMF_Provisioning Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nucmf_Provisioning

import (
	"encoding/json"
)

// RacsData Represents a UE radio capability data provided by the NF service consumer.
type RacsData struct {
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
	// Identifies the configuration related to manufacturer specific UE radio capability. Each element uniquely identifies an RACS configuration for an RACS ID and is identified in the map via the RACS ID as key. The response shall include successfully provisioned RACS data. 
	RacsConfigs map[string]RacsConfiguration `json:"racsConfigs"`
	// Contains the RACS IDs for which the RACS data are not provisioned successfully. The failure reason is also included. Any string value can be used as a key of the map. 
	RacsReports *map[string]RacsFailureReport `json:"racsReports,omitempty"`
}

// NewRacsData instantiates a new RacsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRacsData(racsConfigs map[string]RacsConfiguration) *RacsData {
	this := RacsData{}
	this.RacsConfigs = racsConfigs
	return &this
}

// NewRacsDataWithDefaults instantiates a new RacsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRacsDataWithDefaults() *RacsData {
	this := RacsData{}
	return &this
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *RacsData) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RacsData) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *RacsData) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *RacsData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetRacsConfigs returns the RacsConfigs field value
func (o *RacsData) GetRacsConfigs() map[string]RacsConfiguration {
	if o == nil {
		var ret map[string]RacsConfiguration
		return ret
	}

	return o.RacsConfigs
}

// GetRacsConfigsOk returns a tuple with the RacsConfigs field value
// and a boolean to check if the value has been set.
func (o *RacsData) GetRacsConfigsOk() (*map[string]RacsConfiguration, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RacsConfigs, true
}

// SetRacsConfigs sets field value
func (o *RacsData) SetRacsConfigs(v map[string]RacsConfiguration) {
	o.RacsConfigs = v
}

// GetRacsReports returns the RacsReports field value if set, zero value otherwise.
func (o *RacsData) GetRacsReports() map[string]RacsFailureReport {
	if o == nil || isNil(o.RacsReports) {
		var ret map[string]RacsFailureReport
		return ret
	}
	return *o.RacsReports
}

// GetRacsReportsOk returns a tuple with the RacsReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RacsData) GetRacsReportsOk() (*map[string]RacsFailureReport, bool) {
	if o == nil || isNil(o.RacsReports) {
    return nil, false
	}
	return o.RacsReports, true
}

// HasRacsReports returns a boolean if a field has been set.
func (o *RacsData) HasRacsReports() bool {
	if o != nil && !isNil(o.RacsReports) {
		return true
	}

	return false
}

// SetRacsReports gets a reference to the given map[string]RacsFailureReport and assigns it to the RacsReports field.
func (o *RacsData) SetRacsReports(v map[string]RacsFailureReport) {
	o.RacsReports = &v
}

func (o RacsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if true {
		toSerialize["racsConfigs"] = o.RacsConfigs
	}
	if !isNil(o.RacsReports) {
		toSerialize["racsReports"] = o.RacsReports
	}
	return json.Marshal(toSerialize)
}

type NullableRacsData struct {
	value *RacsData
	isSet bool
}

func (v NullableRacsData) Get() *RacsData {
	return v.value
}

func (v *NullableRacsData) Set(val *RacsData) {
	v.value = val
	v.isSet = true
}

func (v NullableRacsData) IsSet() bool {
	return v.isSet
}

func (v *NullableRacsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRacsData(val *RacsData) *NullableRacsData {
	return &NullableRacsData{value: val, isSet: true}
}

func (v NullableRacsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRacsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


