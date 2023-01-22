/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFDiscovery

import (
	"encoding/json"
)

// QueryParameter Contains the name and value of a query parameter
type QueryParameter struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

// NewQueryParameter instantiates a new QueryParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryParameter(name string, value string) *QueryParameter {
	this := QueryParameter{}
	this.Name = name
	this.Value = value
	return &this
}

// NewQueryParameterWithDefaults instantiates a new QueryParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryParameterWithDefaults() *QueryParameter {
	this := QueryParameter{}
	return &this
}

// GetName returns the Name field value
func (o *QueryParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *QueryParameter) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *QueryParameter) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *QueryParameter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *QueryParameter) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *QueryParameter) SetValue(v string) {
	o.Value = v
}

func (o QueryParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableQueryParameter struct {
	value *QueryParameter
	isSet bool
}

func (v NullableQueryParameter) Get() *QueryParameter {
	return v.value
}

func (v *NullableQueryParameter) Set(val *QueryParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryParameter(val *QueryParameter) *NullableQueryParameter {
	return &NullableQueryParameter{value: val, isSet: true}
}

func (v NullableQueryParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


