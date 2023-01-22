/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// EdrxParameters1 struct for EdrxParameters1
type EdrxParameters1 struct {
	RatType RatType `json:"ratType"`
	EdrxValue string `json:"edrxValue"`
}

// NewEdrxParameters1 instantiates a new EdrxParameters1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdrxParameters1(ratType RatType, edrxValue string) *EdrxParameters1 {
	this := EdrxParameters1{}
	this.RatType = ratType
	this.EdrxValue = edrxValue
	return &this
}

// NewEdrxParameters1WithDefaults instantiates a new EdrxParameters1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdrxParameters1WithDefaults() *EdrxParameters1 {
	this := EdrxParameters1{}
	return &this
}

// GetRatType returns the RatType field value
func (o *EdrxParameters1) GetRatType() RatType {
	if o == nil {
		var ret RatType
		return ret
	}

	return o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value
// and a boolean to check if the value has been set.
func (o *EdrxParameters1) GetRatTypeOk() (*RatType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RatType, true
}

// SetRatType sets field value
func (o *EdrxParameters1) SetRatType(v RatType) {
	o.RatType = v
}

// GetEdrxValue returns the EdrxValue field value
func (o *EdrxParameters1) GetEdrxValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EdrxValue
}

// GetEdrxValueOk returns a tuple with the EdrxValue field value
// and a boolean to check if the value has been set.
func (o *EdrxParameters1) GetEdrxValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EdrxValue, true
}

// SetEdrxValue sets field value
func (o *EdrxParameters1) SetEdrxValue(v string) {
	o.EdrxValue = v
}

func (o EdrxParameters1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ratType"] = o.RatType
	}
	if true {
		toSerialize["edrxValue"] = o.EdrxValue
	}
	return json.Marshal(toSerialize)
}

type NullableEdrxParameters1 struct {
	value *EdrxParameters1
	isSet bool
}

func (v NullableEdrxParameters1) Get() *EdrxParameters1 {
	return v.value
}

func (v *NullableEdrxParameters1) Set(val *EdrxParameters1) {
	v.value = val
	v.isSet = true
}

func (v NullableEdrxParameters1) IsSet() bool {
	return v.isSet
}

func (v *NullableEdrxParameters1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdrxParameters1(val *EdrxParameters1) *NullableEdrxParameters1 {
	return &NullableEdrxParameters1{value: val, isSet: true}
}

func (v NullableEdrxParameters1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdrxParameters1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


