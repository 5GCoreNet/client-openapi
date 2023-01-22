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

// InterRatEsDeactivationCandidateCellParameters struct for InterRatEsDeactivationCandidateCellParameters
type InterRatEsDeactivationCandidateCellParameters struct {
	LoadThreshold *int32 `json:"loadThreshold,omitempty"`
	TimeDuration *int32 `json:"timeDuration,omitempty"`
}

// NewInterRatEsDeactivationCandidateCellParameters instantiates a new InterRatEsDeactivationCandidateCellParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterRatEsDeactivationCandidateCellParameters() *InterRatEsDeactivationCandidateCellParameters {
	this := InterRatEsDeactivationCandidateCellParameters{}
	return &this
}

// NewInterRatEsDeactivationCandidateCellParametersWithDefaults instantiates a new InterRatEsDeactivationCandidateCellParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterRatEsDeactivationCandidateCellParametersWithDefaults() *InterRatEsDeactivationCandidateCellParameters {
	this := InterRatEsDeactivationCandidateCellParameters{}
	return &this
}

// GetLoadThreshold returns the LoadThreshold field value if set, zero value otherwise.
func (o *InterRatEsDeactivationCandidateCellParameters) GetLoadThreshold() int32 {
	if o == nil || isNil(o.LoadThreshold) {
		var ret int32
		return ret
	}
	return *o.LoadThreshold
}

// GetLoadThresholdOk returns a tuple with the LoadThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterRatEsDeactivationCandidateCellParameters) GetLoadThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.LoadThreshold) {
    return nil, false
	}
	return o.LoadThreshold, true
}

// HasLoadThreshold returns a boolean if a field has been set.
func (o *InterRatEsDeactivationCandidateCellParameters) HasLoadThreshold() bool {
	if o != nil && !isNil(o.LoadThreshold) {
		return true
	}

	return false
}

// SetLoadThreshold gets a reference to the given int32 and assigns it to the LoadThreshold field.
func (o *InterRatEsDeactivationCandidateCellParameters) SetLoadThreshold(v int32) {
	o.LoadThreshold = &v
}

// GetTimeDuration returns the TimeDuration field value if set, zero value otherwise.
func (o *InterRatEsDeactivationCandidateCellParameters) GetTimeDuration() int32 {
	if o == nil || isNil(o.TimeDuration) {
		var ret int32
		return ret
	}
	return *o.TimeDuration
}

// GetTimeDurationOk returns a tuple with the TimeDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterRatEsDeactivationCandidateCellParameters) GetTimeDurationOk() (*int32, bool) {
	if o == nil || isNil(o.TimeDuration) {
    return nil, false
	}
	return o.TimeDuration, true
}

// HasTimeDuration returns a boolean if a field has been set.
func (o *InterRatEsDeactivationCandidateCellParameters) HasTimeDuration() bool {
	if o != nil && !isNil(o.TimeDuration) {
		return true
	}

	return false
}

// SetTimeDuration gets a reference to the given int32 and assigns it to the TimeDuration field.
func (o *InterRatEsDeactivationCandidateCellParameters) SetTimeDuration(v int32) {
	o.TimeDuration = &v
}

func (o InterRatEsDeactivationCandidateCellParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LoadThreshold) {
		toSerialize["loadThreshold"] = o.LoadThreshold
	}
	if !isNil(o.TimeDuration) {
		toSerialize["timeDuration"] = o.TimeDuration
	}
	return json.Marshal(toSerialize)
}

type NullableInterRatEsDeactivationCandidateCellParameters struct {
	value *InterRatEsDeactivationCandidateCellParameters
	isSet bool
}

func (v NullableInterRatEsDeactivationCandidateCellParameters) Get() *InterRatEsDeactivationCandidateCellParameters {
	return v.value
}

func (v *NullableInterRatEsDeactivationCandidateCellParameters) Set(val *InterRatEsDeactivationCandidateCellParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableInterRatEsDeactivationCandidateCellParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableInterRatEsDeactivationCandidateCellParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterRatEsDeactivationCandidateCellParameters(val *InterRatEsDeactivationCandidateCellParameters) *NullableInterRatEsDeactivationCandidateCellParameters {
	return &NullableInterRatEsDeactivationCandidateCellParameters{value: val, isSet: true}
}

func (v NullableInterRatEsDeactivationCandidateCellParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterRatEsDeactivationCandidateCellParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

