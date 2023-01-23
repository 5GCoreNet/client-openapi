/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// RimRSReportConf struct for RimRSReportConf
type RimRSReportConf struct {
	ReportIndicator *string `json:"reportIndicator,omitempty"`
	ReportInterval *int32 `json:"reportInterval,omitempty"`
	NrofRIMRSReportInfo *int32 `json:"nrofRIMRSReportInfo,omitempty"`
	MaxPropagationDelay *int32 `json:"maxPropagationDelay,omitempty"`
	RimRSReportInfoList []RimRSReportInfo `json:"rimRSReportInfoList,omitempty"`
}

// NewRimRSReportConf instantiates a new RimRSReportConf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRimRSReportConf() *RimRSReportConf {
	this := RimRSReportConf{}
	return &this
}

// NewRimRSReportConfWithDefaults instantiates a new RimRSReportConf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRimRSReportConfWithDefaults() *RimRSReportConf {
	this := RimRSReportConf{}
	return &this
}

// GetReportIndicator returns the ReportIndicator field value if set, zero value otherwise.
func (o *RimRSReportConf) GetReportIndicator() string {
	if o == nil || isNil(o.ReportIndicator) {
		var ret string
		return ret
	}
	return *o.ReportIndicator
}

// GetReportIndicatorOk returns a tuple with the ReportIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSReportConf) GetReportIndicatorOk() (*string, bool) {
	if o == nil || isNil(o.ReportIndicator) {
    return nil, false
	}
	return o.ReportIndicator, true
}

// HasReportIndicator returns a boolean if a field has been set.
func (o *RimRSReportConf) HasReportIndicator() bool {
	if o != nil && !isNil(o.ReportIndicator) {
		return true
	}

	return false
}

// SetReportIndicator gets a reference to the given string and assigns it to the ReportIndicator field.
func (o *RimRSReportConf) SetReportIndicator(v string) {
	o.ReportIndicator = &v
}

// GetReportInterval returns the ReportInterval field value if set, zero value otherwise.
func (o *RimRSReportConf) GetReportInterval() int32 {
	if o == nil || isNil(o.ReportInterval) {
		var ret int32
		return ret
	}
	return *o.ReportInterval
}

// GetReportIntervalOk returns a tuple with the ReportInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSReportConf) GetReportIntervalOk() (*int32, bool) {
	if o == nil || isNil(o.ReportInterval) {
    return nil, false
	}
	return o.ReportInterval, true
}

// HasReportInterval returns a boolean if a field has been set.
func (o *RimRSReportConf) HasReportInterval() bool {
	if o != nil && !isNil(o.ReportInterval) {
		return true
	}

	return false
}

// SetReportInterval gets a reference to the given int32 and assigns it to the ReportInterval field.
func (o *RimRSReportConf) SetReportInterval(v int32) {
	o.ReportInterval = &v
}

// GetNrofRIMRSReportInfo returns the NrofRIMRSReportInfo field value if set, zero value otherwise.
func (o *RimRSReportConf) GetNrofRIMRSReportInfo() int32 {
	if o == nil || isNil(o.NrofRIMRSReportInfo) {
		var ret int32
		return ret
	}
	return *o.NrofRIMRSReportInfo
}

// GetNrofRIMRSReportInfoOk returns a tuple with the NrofRIMRSReportInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSReportConf) GetNrofRIMRSReportInfoOk() (*int32, bool) {
	if o == nil || isNil(o.NrofRIMRSReportInfo) {
    return nil, false
	}
	return o.NrofRIMRSReportInfo, true
}

// HasNrofRIMRSReportInfo returns a boolean if a field has been set.
func (o *RimRSReportConf) HasNrofRIMRSReportInfo() bool {
	if o != nil && !isNil(o.NrofRIMRSReportInfo) {
		return true
	}

	return false
}

// SetNrofRIMRSReportInfo gets a reference to the given int32 and assigns it to the NrofRIMRSReportInfo field.
func (o *RimRSReportConf) SetNrofRIMRSReportInfo(v int32) {
	o.NrofRIMRSReportInfo = &v
}

// GetMaxPropagationDelay returns the MaxPropagationDelay field value if set, zero value otherwise.
func (o *RimRSReportConf) GetMaxPropagationDelay() int32 {
	if o == nil || isNil(o.MaxPropagationDelay) {
		var ret int32
		return ret
	}
	return *o.MaxPropagationDelay
}

// GetMaxPropagationDelayOk returns a tuple with the MaxPropagationDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSReportConf) GetMaxPropagationDelayOk() (*int32, bool) {
	if o == nil || isNil(o.MaxPropagationDelay) {
    return nil, false
	}
	return o.MaxPropagationDelay, true
}

// HasMaxPropagationDelay returns a boolean if a field has been set.
func (o *RimRSReportConf) HasMaxPropagationDelay() bool {
	if o != nil && !isNil(o.MaxPropagationDelay) {
		return true
	}

	return false
}

// SetMaxPropagationDelay gets a reference to the given int32 and assigns it to the MaxPropagationDelay field.
func (o *RimRSReportConf) SetMaxPropagationDelay(v int32) {
	o.MaxPropagationDelay = &v
}

// GetRimRSReportInfoList returns the RimRSReportInfoList field value if set, zero value otherwise.
func (o *RimRSReportConf) GetRimRSReportInfoList() []RimRSReportInfo {
	if o == nil || isNil(o.RimRSReportInfoList) {
		var ret []RimRSReportInfo
		return ret
	}
	return o.RimRSReportInfoList
}

// GetRimRSReportInfoListOk returns a tuple with the RimRSReportInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RimRSReportConf) GetRimRSReportInfoListOk() ([]RimRSReportInfo, bool) {
	if o == nil || isNil(o.RimRSReportInfoList) {
    return nil, false
	}
	return o.RimRSReportInfoList, true
}

// HasRimRSReportInfoList returns a boolean if a field has been set.
func (o *RimRSReportConf) HasRimRSReportInfoList() bool {
	if o != nil && !isNil(o.RimRSReportInfoList) {
		return true
	}

	return false
}

// SetRimRSReportInfoList gets a reference to the given []RimRSReportInfo and assigns it to the RimRSReportInfoList field.
func (o *RimRSReportConf) SetRimRSReportInfoList(v []RimRSReportInfo) {
	o.RimRSReportInfoList = v
}

func (o RimRSReportConf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReportIndicator) {
		toSerialize["reportIndicator"] = o.ReportIndicator
	}
	if !isNil(o.ReportInterval) {
		toSerialize["reportInterval"] = o.ReportInterval
	}
	if !isNil(o.NrofRIMRSReportInfo) {
		toSerialize["nrofRIMRSReportInfo"] = o.NrofRIMRSReportInfo
	}
	if !isNil(o.MaxPropagationDelay) {
		toSerialize["maxPropagationDelay"] = o.MaxPropagationDelay
	}
	if !isNil(o.RimRSReportInfoList) {
		toSerialize["rimRSReportInfoList"] = o.RimRSReportInfoList
	}
	return json.Marshal(toSerialize)
}

type NullableRimRSReportConf struct {
	value *RimRSReportConf
	isSet bool
}

func (v NullableRimRSReportConf) Get() *RimRSReportConf {
	return v.value
}

func (v *NullableRimRSReportConf) Set(val *RimRSReportConf) {
	v.value = val
	v.isSet = true
}

func (v NullableRimRSReportConf) IsSet() bool {
	return v.isSet
}

func (v *NullableRimRSReportConf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRimRSReportConf(val *RimRSReportConf) *NullableRimRSReportConf {
	return &NullableRimRSReportConf{value: val, isSet: true}
}

func (v NullableRimRSReportConf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRimRSReportConf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


