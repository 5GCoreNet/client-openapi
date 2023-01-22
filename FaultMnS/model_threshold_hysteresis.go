/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FaultMnS

import (
	"encoding/json"
)

// ThresholdHysteresis struct for ThresholdHysteresis
type ThresholdHysteresis struct {
	High ThresholdHysteresisHigh `json:"high"`
	Low *float32 `json:"low,omitempty"`
}

// NewThresholdHysteresis instantiates a new ThresholdHysteresis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThresholdHysteresis(high ThresholdHysteresisHigh) *ThresholdHysteresis {
	this := ThresholdHysteresis{}
	this.High = high
	return &this
}

// NewThresholdHysteresisWithDefaults instantiates a new ThresholdHysteresis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThresholdHysteresisWithDefaults() *ThresholdHysteresis {
	this := ThresholdHysteresis{}
	return &this
}

// GetHigh returns the High field value
func (o *ThresholdHysteresis) GetHigh() ThresholdHysteresisHigh {
	if o == nil {
		var ret ThresholdHysteresisHigh
		return ret
	}

	return o.High
}

// GetHighOk returns a tuple with the High field value
// and a boolean to check if the value has been set.
func (o *ThresholdHysteresis) GetHighOk() (*ThresholdHysteresisHigh, bool) {
	if o == nil {
    return nil, false
	}
	return &o.High, true
}

// SetHigh sets field value
func (o *ThresholdHysteresis) SetHigh(v ThresholdHysteresisHigh) {
	o.High = v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *ThresholdHysteresis) GetLow() float32 {
	if o == nil || isNil(o.Low) {
		var ret float32
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThresholdHysteresis) GetLowOk() (*float32, bool) {
	if o == nil || isNil(o.Low) {
    return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *ThresholdHysteresis) HasLow() bool {
	if o != nil && !isNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given float32 and assigns it to the Low field.
func (o *ThresholdHysteresis) SetLow(v float32) {
	o.Low = &v
}

func (o ThresholdHysteresis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["high"] = o.High
	}
	if !isNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	return json.Marshal(toSerialize)
}

type NullableThresholdHysteresis struct {
	value *ThresholdHysteresis
	isSet bool
}

func (v NullableThresholdHysteresis) Get() *ThresholdHysteresis {
	return v.value
}

func (v *NullableThresholdHysteresis) Set(val *ThresholdHysteresis) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdHysteresis) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdHysteresis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdHysteresis(val *ThresholdHysteresis) *NullableThresholdHysteresis {
	return &NullableThresholdHysteresis{value: val, isSet: true}
}

func (v NullableThresholdHysteresis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdHysteresis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


