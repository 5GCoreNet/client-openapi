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

// Pc5FlowBitRates it shall represent the PC5 Flow Bit Rates
type Pc5FlowBitRates struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	GuaFbr *string `json:"guaFbr,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxFbr *string `json:"maxFbr,omitempty"`
}

// NewPc5FlowBitRates instantiates a new Pc5FlowBitRates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPc5FlowBitRates() *Pc5FlowBitRates {
	this := Pc5FlowBitRates{}
	return &this
}

// NewPc5FlowBitRatesWithDefaults instantiates a new Pc5FlowBitRates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPc5FlowBitRatesWithDefaults() *Pc5FlowBitRates {
	this := Pc5FlowBitRates{}
	return &this
}

// GetGuaFbr returns the GuaFbr field value if set, zero value otherwise.
func (o *Pc5FlowBitRates) GetGuaFbr() string {
	if o == nil || isNil(o.GuaFbr) {
		var ret string
		return ret
	}
	return *o.GuaFbr
}

// GetGuaFbrOk returns a tuple with the GuaFbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pc5FlowBitRates) GetGuaFbrOk() (*string, bool) {
	if o == nil || isNil(o.GuaFbr) {
    return nil, false
	}
	return o.GuaFbr, true
}

// HasGuaFbr returns a boolean if a field has been set.
func (o *Pc5FlowBitRates) HasGuaFbr() bool {
	if o != nil && !isNil(o.GuaFbr) {
		return true
	}

	return false
}

// SetGuaFbr gets a reference to the given string and assigns it to the GuaFbr field.
func (o *Pc5FlowBitRates) SetGuaFbr(v string) {
	o.GuaFbr = &v
}

// GetMaxFbr returns the MaxFbr field value if set, zero value otherwise.
func (o *Pc5FlowBitRates) GetMaxFbr() string {
	if o == nil || isNil(o.MaxFbr) {
		var ret string
		return ret
	}
	return *o.MaxFbr
}

// GetMaxFbrOk returns a tuple with the MaxFbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pc5FlowBitRates) GetMaxFbrOk() (*string, bool) {
	if o == nil || isNil(o.MaxFbr) {
    return nil, false
	}
	return o.MaxFbr, true
}

// HasMaxFbr returns a boolean if a field has been set.
func (o *Pc5FlowBitRates) HasMaxFbr() bool {
	if o != nil && !isNil(o.MaxFbr) {
		return true
	}

	return false
}

// SetMaxFbr gets a reference to the given string and assigns it to the MaxFbr field.
func (o *Pc5FlowBitRates) SetMaxFbr(v string) {
	o.MaxFbr = &v
}

func (o Pc5FlowBitRates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GuaFbr) {
		toSerialize["guaFbr"] = o.GuaFbr
	}
	if !isNil(o.MaxFbr) {
		toSerialize["maxFbr"] = o.MaxFbr
	}
	return json.Marshal(toSerialize)
}

type NullablePc5FlowBitRates struct {
	value *Pc5FlowBitRates
	isSet bool
}

func (v NullablePc5FlowBitRates) Get() *Pc5FlowBitRates {
	return v.value
}

func (v *NullablePc5FlowBitRates) Set(val *Pc5FlowBitRates) {
	v.value = val
	v.isSet = true
}

func (v NullablePc5FlowBitRates) IsSet() bool {
	return v.isSet
}

func (v *NullablePc5FlowBitRates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePc5FlowBitRates(val *Pc5FlowBitRates) *NullablePc5FlowBitRates {
	return &NullablePc5FlowBitRates{value: val, isSet: true}
}

func (v NullablePc5FlowBitRates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePc5FlowBitRates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


