/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// MDAOutputEntry struct for MDAOutputEntry
type MDAOutputEntry struct {
	MDAOutputIEName *string `json:"mDAOutputIEName,omitempty"`
	MdaOutputIEValue interface{} `json:"mdaOutputIEValue,omitempty"`
	AnalyticsWindow *TimeWindow `json:"analyticsWindow,omitempty"`
	ConfidenceDegree *float32 `json:"confidenceDegree,omitempty"`
}

// NewMDAOutputEntry instantiates a new MDAOutputEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAOutputEntry() *MDAOutputEntry {
	this := MDAOutputEntry{}
	return &this
}

// NewMDAOutputEntryWithDefaults instantiates a new MDAOutputEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAOutputEntryWithDefaults() *MDAOutputEntry {
	this := MDAOutputEntry{}
	return &this
}

// GetMDAOutputIEName returns the MDAOutputIEName field value if set, zero value otherwise.
func (o *MDAOutputEntry) GetMDAOutputIEName() string {
	if o == nil || isNil(o.MDAOutputIEName) {
		var ret string
		return ret
	}
	return *o.MDAOutputIEName
}

// GetMDAOutputIENameOk returns a tuple with the MDAOutputIEName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputEntry) GetMDAOutputIENameOk() (*string, bool) {
	if o == nil || isNil(o.MDAOutputIEName) {
    return nil, false
	}
	return o.MDAOutputIEName, true
}

// HasMDAOutputIEName returns a boolean if a field has been set.
func (o *MDAOutputEntry) HasMDAOutputIEName() bool {
	if o != nil && !isNil(o.MDAOutputIEName) {
		return true
	}

	return false
}

// SetMDAOutputIEName gets a reference to the given string and assigns it to the MDAOutputIEName field.
func (o *MDAOutputEntry) SetMDAOutputIEName(v string) {
	o.MDAOutputIEName = &v
}

// GetMdaOutputIEValue returns the MdaOutputIEValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MDAOutputEntry) GetMdaOutputIEValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MdaOutputIEValue
}

// GetMdaOutputIEValueOk returns a tuple with the MdaOutputIEValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MDAOutputEntry) GetMdaOutputIEValueOk() (*interface{}, bool) {
	if o == nil || isNil(o.MdaOutputIEValue) {
    return nil, false
	}
	return &o.MdaOutputIEValue, true
}

// HasMdaOutputIEValue returns a boolean if a field has been set.
func (o *MDAOutputEntry) HasMdaOutputIEValue() bool {
	if o != nil && isNil(o.MdaOutputIEValue) {
		return true
	}

	return false
}

// SetMdaOutputIEValue gets a reference to the given interface{} and assigns it to the MdaOutputIEValue field.
func (o *MDAOutputEntry) SetMdaOutputIEValue(v interface{}) {
	o.MdaOutputIEValue = v
}

// GetAnalyticsWindow returns the AnalyticsWindow field value if set, zero value otherwise.
func (o *MDAOutputEntry) GetAnalyticsWindow() TimeWindow {
	if o == nil || isNil(o.AnalyticsWindow) {
		var ret TimeWindow
		return ret
	}
	return *o.AnalyticsWindow
}

// GetAnalyticsWindowOk returns a tuple with the AnalyticsWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputEntry) GetAnalyticsWindowOk() (*TimeWindow, bool) {
	if o == nil || isNil(o.AnalyticsWindow) {
    return nil, false
	}
	return o.AnalyticsWindow, true
}

// HasAnalyticsWindow returns a boolean if a field has been set.
func (o *MDAOutputEntry) HasAnalyticsWindow() bool {
	if o != nil && !isNil(o.AnalyticsWindow) {
		return true
	}

	return false
}

// SetAnalyticsWindow gets a reference to the given TimeWindow and assigns it to the AnalyticsWindow field.
func (o *MDAOutputEntry) SetAnalyticsWindow(v TimeWindow) {
	o.AnalyticsWindow = &v
}

// GetConfidenceDegree returns the ConfidenceDegree field value if set, zero value otherwise.
func (o *MDAOutputEntry) GetConfidenceDegree() float32 {
	if o == nil || isNil(o.ConfidenceDegree) {
		var ret float32
		return ret
	}
	return *o.ConfidenceDegree
}

// GetConfidenceDegreeOk returns a tuple with the ConfidenceDegree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputEntry) GetConfidenceDegreeOk() (*float32, bool) {
	if o == nil || isNil(o.ConfidenceDegree) {
    return nil, false
	}
	return o.ConfidenceDegree, true
}

// HasConfidenceDegree returns a boolean if a field has been set.
func (o *MDAOutputEntry) HasConfidenceDegree() bool {
	if o != nil && !isNil(o.ConfidenceDegree) {
		return true
	}

	return false
}

// SetConfidenceDegree gets a reference to the given float32 and assigns it to the ConfidenceDegree field.
func (o *MDAOutputEntry) SetConfidenceDegree(v float32) {
	o.ConfidenceDegree = &v
}

func (o MDAOutputEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MDAOutputIEName) {
		toSerialize["mDAOutputIEName"] = o.MDAOutputIEName
	}
	if o.MdaOutputIEValue != nil {
		toSerialize["mdaOutputIEValue"] = o.MdaOutputIEValue
	}
	if !isNil(o.AnalyticsWindow) {
		toSerialize["analyticsWindow"] = o.AnalyticsWindow
	}
	if !isNil(o.ConfidenceDegree) {
		toSerialize["confidenceDegree"] = o.ConfidenceDegree
	}
	return json.Marshal(toSerialize)
}

type NullableMDAOutputEntry struct {
	value *MDAOutputEntry
	isSet bool
}

func (v NullableMDAOutputEntry) Get() *MDAOutputEntry {
	return v.value
}

func (v *NullableMDAOutputEntry) Set(val *MDAOutputEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAOutputEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAOutputEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAOutputEntry(val *MDAOutputEntry) *NullableMDAOutputEntry {
	return &NullableMDAOutputEntry{value: val, isSet: true}
}

func (v NullableMDAOutputEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAOutputEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


