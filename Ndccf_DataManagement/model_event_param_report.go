/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ndccf_DataManagement

import (
	"encoding/json"
)

// EventParamReport Represents a summarized report for one event parameter.
type EventParamReport struct {
	// The name of the reported parameter.
	Name string `json:"name"`
	// The list of values of the reported parameter.
	Values []interface{} `json:"values"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	Area *NetworkAreaInfo `json:"area,omitempty"`
	Spacing *NumberAverage `json:"spacing,omitempty"`
	Duration *NumberAverage `json:"duration,omitempty"`
	AvgAndVar *NumberAverage `json:"avgAndVar,omitempty"`
	MostFreqVal interface{} `json:"mostFreqVal,omitempty"`
	LeastFreqVal interface{} `json:"leastFreqVal,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Count *int32 `json:"count,omitempty"`
	// The minimum value of the parameter.
	MinValue *string `json:"minValue,omitempty"`
	// The maximum value of the parameter.
	MaxValue *string `json:"maxValue,omitempty"`
}

// NewEventParamReport instantiates a new EventParamReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventParamReport(name string, values []interface{}) *EventParamReport {
	this := EventParamReport{}
	this.Name = name
	this.Values = values
	return &this
}

// NewEventParamReportWithDefaults instantiates a new EventParamReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventParamReportWithDefaults() *EventParamReport {
	this := EventParamReport{}
	return &this
}

// GetName returns the Name field value
func (o *EventParamReport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventParamReport) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventParamReport) SetName(v string) {
	o.Name = v
}

// GetValues returns the Values field value
func (o *EventParamReport) GetValues() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *EventParamReport) GetValuesOk() ([]interface{}, bool) {
	if o == nil {
    return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *EventParamReport) SetValues(v []interface{}) {
	o.Values = v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *EventParamReport) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventParamReport) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
    return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *EventParamReport) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *EventParamReport) SetSupi(v string) {
	o.Supi = &v
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *EventParamReport) GetArea() NetworkAreaInfo {
	if o == nil || isNil(o.Area) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventParamReport) GetAreaOk() (*NetworkAreaInfo, bool) {
	if o == nil || isNil(o.Area) {
    return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *EventParamReport) HasArea() bool {
	if o != nil && !isNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given NetworkAreaInfo and assigns it to the Area field.
func (o *EventParamReport) SetArea(v NetworkAreaInfo) {
	o.Area = &v
}

// GetSpacing returns the Spacing field value if set, zero value otherwise.
func (o *EventParamReport) GetSpacing() NumberAverage {
	if o == nil || isNil(o.Spacing) {
		var ret NumberAverage
		return ret
	}
	return *o.Spacing
}

// GetSpacingOk returns a tuple with the Spacing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventParamReport) GetSpacingOk() (*NumberAverage, bool) {
	if o == nil || isNil(o.Spacing) {
    return nil, false
	}
	return o.Spacing, true
}

// HasSpacing returns a boolean if a field has been set.
func (o *EventParamReport) HasSpacing() bool {
	if o != nil && !isNil(o.Spacing) {
		return true
	}

	return false
}

// SetSpacing gets a reference to the given NumberAverage and assigns it to the Spacing field.
func (o *EventParamReport) SetSpacing(v NumberAverage) {
	o.Spacing = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *EventParamReport) GetDuration() NumberAverage {
	if o == nil || isNil(o.Duration) {
		var ret NumberAverage
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventParamReport) GetDurationOk() (*NumberAverage, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *EventParamReport) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NumberAverage and assigns it to the Duration field.
func (o *EventParamReport) SetDuration(v NumberAverage) {
	o.Duration = &v
}

// GetAvgAndVar returns the AvgAndVar field value if set, zero value otherwise.
func (o *EventParamReport) GetAvgAndVar() NumberAverage {
	if o == nil || isNil(o.AvgAndVar) {
		var ret NumberAverage
		return ret
	}
	return *o.AvgAndVar
}

// GetAvgAndVarOk returns a tuple with the AvgAndVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventParamReport) GetAvgAndVarOk() (*NumberAverage, bool) {
	if o == nil || isNil(o.AvgAndVar) {
    return nil, false
	}
	return o.AvgAndVar, true
}

// HasAvgAndVar returns a boolean if a field has been set.
func (o *EventParamReport) HasAvgAndVar() bool {
	if o != nil && !isNil(o.AvgAndVar) {
		return true
	}

	return false
}

// SetAvgAndVar gets a reference to the given NumberAverage and assigns it to the AvgAndVar field.
func (o *EventParamReport) SetAvgAndVar(v NumberAverage) {
	o.AvgAndVar = &v
}

// GetMostFreqVal returns the MostFreqVal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventParamReport) GetMostFreqVal() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MostFreqVal
}

// GetMostFreqValOk returns a tuple with the MostFreqVal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventParamReport) GetMostFreqValOk() (*interface{}, bool) {
	if o == nil || isNil(o.MostFreqVal) {
    return nil, false
	}
	return &o.MostFreqVal, true
}

// HasMostFreqVal returns a boolean if a field has been set.
func (o *EventParamReport) HasMostFreqVal() bool {
	if o != nil && isNil(o.MostFreqVal) {
		return true
	}

	return false
}

// SetMostFreqVal gets a reference to the given interface{} and assigns it to the MostFreqVal field.
func (o *EventParamReport) SetMostFreqVal(v interface{}) {
	o.MostFreqVal = v
}

// GetLeastFreqVal returns the LeastFreqVal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventParamReport) GetLeastFreqVal() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LeastFreqVal
}

// GetLeastFreqValOk returns a tuple with the LeastFreqVal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventParamReport) GetLeastFreqValOk() (*interface{}, bool) {
	if o == nil || isNil(o.LeastFreqVal) {
    return nil, false
	}
	return &o.LeastFreqVal, true
}

// HasLeastFreqVal returns a boolean if a field has been set.
func (o *EventParamReport) HasLeastFreqVal() bool {
	if o != nil && isNil(o.LeastFreqVal) {
		return true
	}

	return false
}

// SetLeastFreqVal gets a reference to the given interface{} and assigns it to the LeastFreqVal field.
func (o *EventParamReport) SetLeastFreqVal(v interface{}) {
	o.LeastFreqVal = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *EventParamReport) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventParamReport) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *EventParamReport) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *EventParamReport) SetCount(v int32) {
	o.Count = &v
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *EventParamReport) GetMinValue() string {
	if o == nil || isNil(o.MinValue) {
		var ret string
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventParamReport) GetMinValueOk() (*string, bool) {
	if o == nil || isNil(o.MinValue) {
    return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *EventParamReport) HasMinValue() bool {
	if o != nil && !isNil(o.MinValue) {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given string and assigns it to the MinValue field.
func (o *EventParamReport) SetMinValue(v string) {
	o.MinValue = &v
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *EventParamReport) GetMaxValue() string {
	if o == nil || isNil(o.MaxValue) {
		var ret string
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventParamReport) GetMaxValueOk() (*string, bool) {
	if o == nil || isNil(o.MaxValue) {
    return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *EventParamReport) HasMaxValue() bool {
	if o != nil && !isNil(o.MaxValue) {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given string and assigns it to the MaxValue field.
func (o *EventParamReport) SetMaxValue(v string) {
	o.MaxValue = &v
}

func (o EventParamReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["values"] = o.Values
	}
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if !isNil(o.Spacing) {
		toSerialize["spacing"] = o.Spacing
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.AvgAndVar) {
		toSerialize["avgAndVar"] = o.AvgAndVar
	}
	if o.MostFreqVal != nil {
		toSerialize["mostFreqVal"] = o.MostFreqVal
	}
	if o.LeastFreqVal != nil {
		toSerialize["leastFreqVal"] = o.LeastFreqVal
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.MinValue) {
		toSerialize["minValue"] = o.MinValue
	}
	if !isNil(o.MaxValue) {
		toSerialize["maxValue"] = o.MaxValue
	}
	return json.Marshal(toSerialize)
}

type NullableEventParamReport struct {
	value *EventParamReport
	isSet bool
}

func (v NullableEventParamReport) Get() *EventParamReport {
	return v.value
}

func (v *NullableEventParamReport) Set(val *EventParamReport) {
	v.value = val
	v.isSet = true
}

func (v NullableEventParamReport) IsSet() bool {
	return v.isSet
}

func (v *NullableEventParamReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventParamReport(val *EventParamReport) *NullableEventParamReport {
	return &NullableEventParamReport{value: val, isSet: true}
}

func (v NullableEventParamReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventParamReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


