/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// IntraRatEsActivationOriginalCellLoadParameters struct for IntraRatEsActivationOriginalCellLoadParameters
type IntraRatEsActivationOriginalCellLoadParameters struct {
	LoadThreshold *int32 `json:"loadThreshold,omitempty"`
	TimeDuration *int32 `json:"timeDuration,omitempty"`
}

// NewIntraRatEsActivationOriginalCellLoadParameters instantiates a new IntraRatEsActivationOriginalCellLoadParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntraRatEsActivationOriginalCellLoadParameters() *IntraRatEsActivationOriginalCellLoadParameters {
	this := IntraRatEsActivationOriginalCellLoadParameters{}
	return &this
}

// NewIntraRatEsActivationOriginalCellLoadParametersWithDefaults instantiates a new IntraRatEsActivationOriginalCellLoadParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntraRatEsActivationOriginalCellLoadParametersWithDefaults() *IntraRatEsActivationOriginalCellLoadParameters {
	this := IntraRatEsActivationOriginalCellLoadParameters{}
	return &this
}

// GetLoadThreshold returns the LoadThreshold field value if set, zero value otherwise.
func (o *IntraRatEsActivationOriginalCellLoadParameters) GetLoadThreshold() int32 {
	if o == nil || isNil(o.LoadThreshold) {
		var ret int32
		return ret
	}
	return *o.LoadThreshold
}

// GetLoadThresholdOk returns a tuple with the LoadThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntraRatEsActivationOriginalCellLoadParameters) GetLoadThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.LoadThreshold) {
    return nil, false
	}
	return o.LoadThreshold, true
}

// HasLoadThreshold returns a boolean if a field has been set.
func (o *IntraRatEsActivationOriginalCellLoadParameters) HasLoadThreshold() bool {
	if o != nil && !isNil(o.LoadThreshold) {
		return true
	}

	return false
}

// SetLoadThreshold gets a reference to the given int32 and assigns it to the LoadThreshold field.
func (o *IntraRatEsActivationOriginalCellLoadParameters) SetLoadThreshold(v int32) {
	o.LoadThreshold = &v
}

// GetTimeDuration returns the TimeDuration field value if set, zero value otherwise.
func (o *IntraRatEsActivationOriginalCellLoadParameters) GetTimeDuration() int32 {
	if o == nil || isNil(o.TimeDuration) {
		var ret int32
		return ret
	}
	return *o.TimeDuration
}

// GetTimeDurationOk returns a tuple with the TimeDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntraRatEsActivationOriginalCellLoadParameters) GetTimeDurationOk() (*int32, bool) {
	if o == nil || isNil(o.TimeDuration) {
    return nil, false
	}
	return o.TimeDuration, true
}

// HasTimeDuration returns a boolean if a field has been set.
func (o *IntraRatEsActivationOriginalCellLoadParameters) HasTimeDuration() bool {
	if o != nil && !isNil(o.TimeDuration) {
		return true
	}

	return false
}

// SetTimeDuration gets a reference to the given int32 and assigns it to the TimeDuration field.
func (o *IntraRatEsActivationOriginalCellLoadParameters) SetTimeDuration(v int32) {
	o.TimeDuration = &v
}

func (o IntraRatEsActivationOriginalCellLoadParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LoadThreshold) {
		toSerialize["loadThreshold"] = o.LoadThreshold
	}
	if !isNil(o.TimeDuration) {
		toSerialize["timeDuration"] = o.TimeDuration
	}
	return json.Marshal(toSerialize)
}

type NullableIntraRatEsActivationOriginalCellLoadParameters struct {
	value *IntraRatEsActivationOriginalCellLoadParameters
	isSet bool
}

func (v NullableIntraRatEsActivationOriginalCellLoadParameters) Get() *IntraRatEsActivationOriginalCellLoadParameters {
	return v.value
}

func (v *NullableIntraRatEsActivationOriginalCellLoadParameters) Set(val *IntraRatEsActivationOriginalCellLoadParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableIntraRatEsActivationOriginalCellLoadParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableIntraRatEsActivationOriginalCellLoadParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntraRatEsActivationOriginalCellLoadParameters(val *IntraRatEsActivationOriginalCellLoadParameters) *NullableIntraRatEsActivationOriginalCellLoadParameters {
	return &NullableIntraRatEsActivationOriginalCellLoadParameters{value: val, isSet: true}
}

func (v NullableIntraRatEsActivationOriginalCellLoadParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntraRatEsActivationOriginalCellLoadParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


