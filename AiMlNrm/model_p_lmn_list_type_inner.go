/*
AI/ML NRM

OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package AiMlNrm

import (
	"encoding/json"
)

// PLMNListTypeInner struct for PLMNListTypeInner
type PLMNListTypeInner struct {
	Mcc string `json:"mcc"`
	Mnc string `json:"mnc"`
}

// NewPLMNListTypeInner instantiates a new PLMNListTypeInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPLMNListTypeInner(mcc string, mnc string) *PLMNListTypeInner {
	this := PLMNListTypeInner{}
	this.Mcc = mcc
	this.Mnc = mnc
	return &this
}

// NewPLMNListTypeInnerWithDefaults instantiates a new PLMNListTypeInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPLMNListTypeInnerWithDefaults() *PLMNListTypeInner {
	this := PLMNListTypeInner{}
	return &this
}

// GetMcc returns the Mcc field value
func (o *PLMNListTypeInner) GetMcc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value
// and a boolean to check if the value has been set.
func (o *PLMNListTypeInner) GetMccOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mcc, true
}

// SetMcc sets field value
func (o *PLMNListTypeInner) SetMcc(v string) {
	o.Mcc = v
}

// GetMnc returns the Mnc field value
func (o *PLMNListTypeInner) GetMnc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value
// and a boolean to check if the value has been set.
func (o *PLMNListTypeInner) GetMncOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mnc, true
}

// SetMnc sets field value
func (o *PLMNListTypeInner) SetMnc(v string) {
	o.Mnc = v
}

func (o PLMNListTypeInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mcc"] = o.Mcc
	}
	if true {
		toSerialize["mnc"] = o.Mnc
	}
	return json.Marshal(toSerialize)
}

type NullablePLMNListTypeInner struct {
	value *PLMNListTypeInner
	isSet bool
}

func (v NullablePLMNListTypeInner) Get() *PLMNListTypeInner {
	return v.value
}

func (v *NullablePLMNListTypeInner) Set(val *PLMNListTypeInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePLMNListTypeInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePLMNListTypeInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePLMNListTypeInner(val *PLMNListTypeInner) *NullablePLMNListTypeInner {
	return &NullablePLMNListTypeInner{value: val, isSet: true}
}

func (v NullablePLMNListTypeInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePLMNListTypeInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

