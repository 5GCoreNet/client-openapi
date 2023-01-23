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

// ReachabilityForDataConfiguration struct for ReachabilityForDataConfiguration
type ReachabilityForDataConfiguration struct {
	ReportCfg ReachabilityForDataReportConfig `json:"reportCfg"`
	// indicating a time in seconds.
	MinInterval *int32 `json:"minInterval,omitempty"`
}

// NewReachabilityForDataConfiguration instantiates a new ReachabilityForDataConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReachabilityForDataConfiguration(reportCfg ReachabilityForDataReportConfig) *ReachabilityForDataConfiguration {
	this := ReachabilityForDataConfiguration{}
	this.ReportCfg = reportCfg
	return &this
}

// NewReachabilityForDataConfigurationWithDefaults instantiates a new ReachabilityForDataConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReachabilityForDataConfigurationWithDefaults() *ReachabilityForDataConfiguration {
	this := ReachabilityForDataConfiguration{}
	return &this
}

// GetReportCfg returns the ReportCfg field value
func (o *ReachabilityForDataConfiguration) GetReportCfg() ReachabilityForDataReportConfig {
	if o == nil {
		var ret ReachabilityForDataReportConfig
		return ret
	}

	return o.ReportCfg
}

// GetReportCfgOk returns a tuple with the ReportCfg field value
// and a boolean to check if the value has been set.
func (o *ReachabilityForDataConfiguration) GetReportCfgOk() (*ReachabilityForDataReportConfig, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReportCfg, true
}

// SetReportCfg sets field value
func (o *ReachabilityForDataConfiguration) SetReportCfg(v ReachabilityForDataReportConfig) {
	o.ReportCfg = v
}

// GetMinInterval returns the MinInterval field value if set, zero value otherwise.
func (o *ReachabilityForDataConfiguration) GetMinInterval() int32 {
	if o == nil || isNil(o.MinInterval) {
		var ret int32
		return ret
	}
	return *o.MinInterval
}

// GetMinIntervalOk returns a tuple with the MinInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReachabilityForDataConfiguration) GetMinIntervalOk() (*int32, bool) {
	if o == nil || isNil(o.MinInterval) {
    return nil, false
	}
	return o.MinInterval, true
}

// HasMinInterval returns a boolean if a field has been set.
func (o *ReachabilityForDataConfiguration) HasMinInterval() bool {
	if o != nil && !isNil(o.MinInterval) {
		return true
	}

	return false
}

// SetMinInterval gets a reference to the given int32 and assigns it to the MinInterval field.
func (o *ReachabilityForDataConfiguration) SetMinInterval(v int32) {
	o.MinInterval = &v
}

func (o ReachabilityForDataConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reportCfg"] = o.ReportCfg
	}
	if !isNil(o.MinInterval) {
		toSerialize["minInterval"] = o.MinInterval
	}
	return json.Marshal(toSerialize)
}

type NullableReachabilityForDataConfiguration struct {
	value *ReachabilityForDataConfiguration
	isSet bool
}

func (v NullableReachabilityForDataConfiguration) Get() *ReachabilityForDataConfiguration {
	return v.value
}

func (v *NullableReachabilityForDataConfiguration) Set(val *ReachabilityForDataConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityForDataConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityForDataConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityForDataConfiguration(val *ReachabilityForDataConfiguration) *NullableReachabilityForDataConfiguration {
	return &NullableReachabilityForDataConfiguration{value: val, isSet: true}
}

func (v NullableReachabilityForDataConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityForDataConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


