/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nchf_ConvergedCharging

import (
	"encoding/json"
)

// ThresholdValue Indicates the threshold value(s) for RTT and/or Packet Loss Rate.
type ThresholdValue struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property. 
	RttThres NullableInt32 `json:"rttThres,omitempty"`
	// This data type is defined in the same way as the 'PacketLossRate' data type, but with the OpenAPI 'nullable: true' property. 
	PlrThres NullableInt32 `json:"plrThres,omitempty"`
}

// NewThresholdValue instantiates a new ThresholdValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdValue() *ThresholdValue {
	this := ThresholdValue{}
	return &this
}

// NewThresholdValueWithDefaults instantiates a new ThresholdValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdValueWithDefaults() *ThresholdValue {
	this := ThresholdValue{}
	return &this
}

// GetRttThres returns the RttThres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThresholdValue) GetRttThres() int32 {
	if o == nil || isNil(o.RttThres.Get()) {
		var ret int32
		return ret
	}
	return *o.RttThres.Get()
}

// GetRttThresOk returns a tuple with the RttThres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThresholdValue) GetRttThresOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.RttThres.Get(), o.RttThres.IsSet()
}

// HasRttThres returns a boolean if a field has been set.
func (o *ThresholdValue) HasRttThres() bool {
	if o != nil && o.RttThres.IsSet() {
		return true
	}

	return false
}

// SetRttThres gets a reference to the given NullableInt32 and assigns it to the RttThres field.
func (o *ThresholdValue) SetRttThres(v int32) {
	o.RttThres.Set(&v)
}
// SetRttThresNil sets the value for RttThres to be an explicit nil
func (o *ThresholdValue) SetRttThresNil() {
	o.RttThres.Set(nil)
}

// UnsetRttThres ensures that no value is present for RttThres, not even an explicit nil
func (o *ThresholdValue) UnsetRttThres() {
	o.RttThres.Unset()
}

// GetPlrThres returns the PlrThres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThresholdValue) GetPlrThres() int32 {
	if o == nil || isNil(o.PlrThres.Get()) {
		var ret int32
		return ret
	}
	return *o.PlrThres.Get()
}

// GetPlrThresOk returns a tuple with the PlrThres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThresholdValue) GetPlrThresOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.PlrThres.Get(), o.PlrThres.IsSet()
}

// HasPlrThres returns a boolean if a field has been set.
func (o *ThresholdValue) HasPlrThres() bool {
	if o != nil && o.PlrThres.IsSet() {
		return true
	}

	return false
}

// SetPlrThres gets a reference to the given NullableInt32 and assigns it to the PlrThres field.
func (o *ThresholdValue) SetPlrThres(v int32) {
	o.PlrThres.Set(&v)
}
// SetPlrThresNil sets the value for PlrThres to be an explicit nil
func (o *ThresholdValue) SetPlrThresNil() {
	o.PlrThres.Set(nil)
}

// UnsetPlrThres ensures that no value is present for PlrThres, not even an explicit nil
func (o *ThresholdValue) UnsetPlrThres() {
	o.PlrThres.Unset()
}

func (o ThresholdValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RttThres.IsSet() {
		toSerialize["rttThres"] = o.RttThres.Get()
	}
	if o.PlrThres.IsSet() {
		toSerialize["plrThres"] = o.PlrThres.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableThresholdValue struct {
	value *ThresholdValue
	isSet bool
}

func (v NullableThresholdValue) Get() *ThresholdValue {
	return v.value
}

func (v *NullableThresholdValue) Set(val *ThresholdValue) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdValue) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdValue(val *ThresholdValue) *NullableThresholdValue {
	return &NullableThresholdValue{value: val, isSet: true}
}

func (v NullableThresholdValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


