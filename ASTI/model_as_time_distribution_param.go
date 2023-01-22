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

// AsTimeDistributionParam Contains the 5G access stratum time distribution parameters.
type AsTimeDistributionParam struct {
	// When this attribute is included and set to true, it indicates that the access stratum time distribution via Uu reference point is activated. 
	AsTimeDisEnabled *bool `json:"asTimeDisEnabled,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	TimeSyncErrBdgt *int32 `json:"timeSyncErrBdgt,omitempty"`
	TempValidity *TemporalValidity `json:"tempValidity,omitempty"`
}

// NewAsTimeDistributionParam instantiates a new AsTimeDistributionParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsTimeDistributionParam() *AsTimeDistributionParam {
	this := AsTimeDistributionParam{}
	return &this
}

// NewAsTimeDistributionParamWithDefaults instantiates a new AsTimeDistributionParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsTimeDistributionParamWithDefaults() *AsTimeDistributionParam {
	this := AsTimeDistributionParam{}
	return &this
}

// GetAsTimeDisEnabled returns the AsTimeDisEnabled field value if set, zero value otherwise.
func (o *AsTimeDistributionParam) GetAsTimeDisEnabled() bool {
	if o == nil || isNil(o.AsTimeDisEnabled) {
		var ret bool
		return ret
	}
	return *o.AsTimeDisEnabled
}

// GetAsTimeDisEnabledOk returns a tuple with the AsTimeDisEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsTimeDistributionParam) GetAsTimeDisEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.AsTimeDisEnabled) {
    return nil, false
	}
	return o.AsTimeDisEnabled, true
}

// HasAsTimeDisEnabled returns a boolean if a field has been set.
func (o *AsTimeDistributionParam) HasAsTimeDisEnabled() bool {
	if o != nil && !isNil(o.AsTimeDisEnabled) {
		return true
	}

	return false
}

// SetAsTimeDisEnabled gets a reference to the given bool and assigns it to the AsTimeDisEnabled field.
func (o *AsTimeDistributionParam) SetAsTimeDisEnabled(v bool) {
	o.AsTimeDisEnabled = &v
}

// GetTimeSyncErrBdgt returns the TimeSyncErrBdgt field value if set, zero value otherwise.
func (o *AsTimeDistributionParam) GetTimeSyncErrBdgt() int32 {
	if o == nil || isNil(o.TimeSyncErrBdgt) {
		var ret int32
		return ret
	}
	return *o.TimeSyncErrBdgt
}

// GetTimeSyncErrBdgtOk returns a tuple with the TimeSyncErrBdgt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsTimeDistributionParam) GetTimeSyncErrBdgtOk() (*int32, bool) {
	if o == nil || isNil(o.TimeSyncErrBdgt) {
    return nil, false
	}
	return o.TimeSyncErrBdgt, true
}

// HasTimeSyncErrBdgt returns a boolean if a field has been set.
func (o *AsTimeDistributionParam) HasTimeSyncErrBdgt() bool {
	if o != nil && !isNil(o.TimeSyncErrBdgt) {
		return true
	}

	return false
}

// SetTimeSyncErrBdgt gets a reference to the given int32 and assigns it to the TimeSyncErrBdgt field.
func (o *AsTimeDistributionParam) SetTimeSyncErrBdgt(v int32) {
	o.TimeSyncErrBdgt = &v
}

// GetTempValidity returns the TempValidity field value if set, zero value otherwise.
func (o *AsTimeDistributionParam) GetTempValidity() TemporalValidity {
	if o == nil || isNil(o.TempValidity) {
		var ret TemporalValidity
		return ret
	}
	return *o.TempValidity
}

// GetTempValidityOk returns a tuple with the TempValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsTimeDistributionParam) GetTempValidityOk() (*TemporalValidity, bool) {
	if o == nil || isNil(o.TempValidity) {
    return nil, false
	}
	return o.TempValidity, true
}

// HasTempValidity returns a boolean if a field has been set.
func (o *AsTimeDistributionParam) HasTempValidity() bool {
	if o != nil && !isNil(o.TempValidity) {
		return true
	}

	return false
}

// SetTempValidity gets a reference to the given TemporalValidity and assigns it to the TempValidity field.
func (o *AsTimeDistributionParam) SetTempValidity(v TemporalValidity) {
	o.TempValidity = &v
}

func (o AsTimeDistributionParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AsTimeDisEnabled) {
		toSerialize["asTimeDisEnabled"] = o.AsTimeDisEnabled
	}
	if !isNil(o.TimeSyncErrBdgt) {
		toSerialize["timeSyncErrBdgt"] = o.TimeSyncErrBdgt
	}
	if !isNil(o.TempValidity) {
		toSerialize["tempValidity"] = o.TempValidity
	}
	return json.Marshal(toSerialize)
}

type NullableAsTimeDistributionParam struct {
	value *AsTimeDistributionParam
	isSet bool
}

func (v NullableAsTimeDistributionParam) Get() *AsTimeDistributionParam {
	return v.value
}

func (v *NullableAsTimeDistributionParam) Set(val *AsTimeDistributionParam) {
	v.value = val
	v.isSet = true
}

func (v NullableAsTimeDistributionParam) IsSet() bool {
	return v.isSet
}

func (v *NullableAsTimeDistributionParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsTimeDistributionParam(val *AsTimeDistributionParam) *NullableAsTimeDistributionParam {
	return &NullableAsTimeDistributionParam{value: val, isSet: true}
}

func (v NullableAsTimeDistributionParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsTimeDistributionParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


