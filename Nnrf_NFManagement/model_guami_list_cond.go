/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFManagement

import (
	"encoding/json"
)

// GuamiListCond Subscription to a set of AMFs, based on their GUAMIs
type GuamiListCond struct {
	GuamiList []Guami `json:"guamiList"`
}

// NewGuamiListCond instantiates a new GuamiListCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuamiListCond(guamiList []Guami) *GuamiListCond {
	this := GuamiListCond{}
	this.GuamiList = guamiList
	return &this
}

// NewGuamiListCondWithDefaults instantiates a new GuamiListCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuamiListCondWithDefaults() *GuamiListCond {
	this := GuamiListCond{}
	return &this
}

// GetGuamiList returns the GuamiList field value
func (o *GuamiListCond) GetGuamiList() []Guami {
	if o == nil {
		var ret []Guami
		return ret
	}

	return o.GuamiList
}

// GetGuamiListOk returns a tuple with the GuamiList field value
// and a boolean to check if the value has been set.
func (o *GuamiListCond) GetGuamiListOk() ([]Guami, bool) {
	if o == nil {
    return nil, false
	}
	return o.GuamiList, true
}

// SetGuamiList sets field value
func (o *GuamiListCond) SetGuamiList(v []Guami) {
	o.GuamiList = v
}

func (o GuamiListCond) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["guamiList"] = o.GuamiList
	}
	return json.Marshal(toSerialize)
}

type NullableGuamiListCond struct {
	value *GuamiListCond
	isSet bool
}

func (v NullableGuamiListCond) Get() *GuamiListCond {
	return v.value
}

func (v *NullableGuamiListCond) Set(val *GuamiListCond) {
	v.value = val
	v.isSet = true
}

func (v NullableGuamiListCond) IsSet() bool {
	return v.isSet
}

func (v *NullableGuamiListCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuamiListCond(val *GuamiListCond) *NullableGuamiListCond {
	return &NullableGuamiListCond{value: val, isSet: true}
}

func (v NullableGuamiListCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuamiListCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


