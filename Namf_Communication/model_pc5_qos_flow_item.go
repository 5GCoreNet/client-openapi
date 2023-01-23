/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// Pc5QosFlowItem Contains a PC5 QOS flow.
type Pc5QosFlowItem struct {
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255. 
	Pqi int32 `json:"pqi"`
	Pc5FlowBitRates *Pc5FlowBitRates `json:"pc5FlowBitRates,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Range *int32 `json:"range,omitempty"`
}

// NewPc5QosFlowItem instantiates a new Pc5QosFlowItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPc5QosFlowItem(pqi int32) *Pc5QosFlowItem {
	this := Pc5QosFlowItem{}
	this.Pqi = pqi
	return &this
}

// NewPc5QosFlowItemWithDefaults instantiates a new Pc5QosFlowItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPc5QosFlowItemWithDefaults() *Pc5QosFlowItem {
	this := Pc5QosFlowItem{}
	return &this
}

// GetPqi returns the Pqi field value
func (o *Pc5QosFlowItem) GetPqi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pqi
}

// GetPqiOk returns a tuple with the Pqi field value
// and a boolean to check if the value has been set.
func (o *Pc5QosFlowItem) GetPqiOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Pqi, true
}

// SetPqi sets field value
func (o *Pc5QosFlowItem) SetPqi(v int32) {
	o.Pqi = v
}

// GetPc5FlowBitRates returns the Pc5FlowBitRates field value if set, zero value otherwise.
func (o *Pc5QosFlowItem) GetPc5FlowBitRates() Pc5FlowBitRates {
	if o == nil || isNil(o.Pc5FlowBitRates) {
		var ret Pc5FlowBitRates
		return ret
	}
	return *o.Pc5FlowBitRates
}

// GetPc5FlowBitRatesOk returns a tuple with the Pc5FlowBitRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pc5QosFlowItem) GetPc5FlowBitRatesOk() (*Pc5FlowBitRates, bool) {
	if o == nil || isNil(o.Pc5FlowBitRates) {
    return nil, false
	}
	return o.Pc5FlowBitRates, true
}

// HasPc5FlowBitRates returns a boolean if a field has been set.
func (o *Pc5QosFlowItem) HasPc5FlowBitRates() bool {
	if o != nil && !isNil(o.Pc5FlowBitRates) {
		return true
	}

	return false
}

// SetPc5FlowBitRates gets a reference to the given Pc5FlowBitRates and assigns it to the Pc5FlowBitRates field.
func (o *Pc5QosFlowItem) SetPc5FlowBitRates(v Pc5FlowBitRates) {
	o.Pc5FlowBitRates = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *Pc5QosFlowItem) GetRange() int32 {
	if o == nil || isNil(o.Range) {
		var ret int32
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pc5QosFlowItem) GetRangeOk() (*int32, bool) {
	if o == nil || isNil(o.Range) {
    return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *Pc5QosFlowItem) HasRange() bool {
	if o != nil && !isNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given int32 and assigns it to the Range field.
func (o *Pc5QosFlowItem) SetRange(v int32) {
	o.Range = &v
}

func (o Pc5QosFlowItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pqi"] = o.Pqi
	}
	if !isNil(o.Pc5FlowBitRates) {
		toSerialize["pc5FlowBitRates"] = o.Pc5FlowBitRates
	}
	if !isNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	return json.Marshal(toSerialize)
}

type NullablePc5QosFlowItem struct {
	value *Pc5QosFlowItem
	isSet bool
}

func (v NullablePc5QosFlowItem) Get() *Pc5QosFlowItem {
	return v.value
}

func (v *NullablePc5QosFlowItem) Set(val *Pc5QosFlowItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePc5QosFlowItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePc5QosFlowItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePc5QosFlowItem(val *Pc5QosFlowItem) *NullablePc5QosFlowItem {
	return &NullablePc5QosFlowItem{value: val, isSet: true}
}

func (v NullablePc5QosFlowItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePc5QosFlowItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


