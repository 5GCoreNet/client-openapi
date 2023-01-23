/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// Throughput struct for Throughput
type Throughput struct {
	// string with format 'float' as defined in OpenAPI.
	GuaranteedThpt *float32 `json:"guaranteedThpt,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	MaximumThpt *float32 `json:"maximumThpt,omitempty"`
}

// NewThroughput instantiates a new Throughput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThroughput() *Throughput {
	this := Throughput{}
	return &this
}

// NewThroughputWithDefaults instantiates a new Throughput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThroughputWithDefaults() *Throughput {
	this := Throughput{}
	return &this
}

// GetGuaranteedThpt returns the GuaranteedThpt field value if set, zero value otherwise.
func (o *Throughput) GetGuaranteedThpt() float32 {
	if o == nil || isNil(o.GuaranteedThpt) {
		var ret float32
		return ret
	}
	return *o.GuaranteedThpt
}

// GetGuaranteedThptOk returns a tuple with the GuaranteedThpt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Throughput) GetGuaranteedThptOk() (*float32, bool) {
	if o == nil || isNil(o.GuaranteedThpt) {
    return nil, false
	}
	return o.GuaranteedThpt, true
}

// HasGuaranteedThpt returns a boolean if a field has been set.
func (o *Throughput) HasGuaranteedThpt() bool {
	if o != nil && !isNil(o.GuaranteedThpt) {
		return true
	}

	return false
}

// SetGuaranteedThpt gets a reference to the given float32 and assigns it to the GuaranteedThpt field.
func (o *Throughput) SetGuaranteedThpt(v float32) {
	o.GuaranteedThpt = &v
}

// GetMaximumThpt returns the MaximumThpt field value if set, zero value otherwise.
func (o *Throughput) GetMaximumThpt() float32 {
	if o == nil || isNil(o.MaximumThpt) {
		var ret float32
		return ret
	}
	return *o.MaximumThpt
}

// GetMaximumThptOk returns a tuple with the MaximumThpt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Throughput) GetMaximumThptOk() (*float32, bool) {
	if o == nil || isNil(o.MaximumThpt) {
    return nil, false
	}
	return o.MaximumThpt, true
}

// HasMaximumThpt returns a boolean if a field has been set.
func (o *Throughput) HasMaximumThpt() bool {
	if o != nil && !isNil(o.MaximumThpt) {
		return true
	}

	return false
}

// SetMaximumThpt gets a reference to the given float32 and assigns it to the MaximumThpt field.
func (o *Throughput) SetMaximumThpt(v float32) {
	o.MaximumThpt = &v
}

func (o Throughput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GuaranteedThpt) {
		toSerialize["guaranteedThpt"] = o.GuaranteedThpt
	}
	if !isNil(o.MaximumThpt) {
		toSerialize["maximumThpt"] = o.MaximumThpt
	}
	return json.Marshal(toSerialize)
}

type NullableThroughput struct {
	value *Throughput
	isSet bool
}

func (v NullableThroughput) Get() *Throughput {
	return v.value
}

func (v *NullableThroughput) Set(val *Throughput) {
	v.value = val
	v.isSet = true
}

func (v NullableThroughput) IsSet() bool {
	return v.isSet
}

func (v *NullableThroughput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThroughput(val *Throughput) *NullableThroughput {
	return &NullableThroughput{value: val, isSet: true}
}

func (v NullableThroughput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThroughput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


