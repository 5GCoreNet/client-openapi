/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// PtwParameters1 struct for PtwParameters1
type PtwParameters1 struct {
	OperationMode OperationMode `json:"operationMode"`
	PtwValue string `json:"ptwValue"`
}

// NewPtwParameters1 instantiates a new PtwParameters1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPtwParameters1(operationMode OperationMode, ptwValue string) *PtwParameters1 {
	this := PtwParameters1{}
	this.OperationMode = operationMode
	this.PtwValue = ptwValue
	return &this
}

// NewPtwParameters1WithDefaults instantiates a new PtwParameters1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPtwParameters1WithDefaults() *PtwParameters1 {
	this := PtwParameters1{}
	return &this
}

// GetOperationMode returns the OperationMode field value
func (o *PtwParameters1) GetOperationMode() OperationMode {
	if o == nil {
		var ret OperationMode
		return ret
	}

	return o.OperationMode
}

// GetOperationModeOk returns a tuple with the OperationMode field value
// and a boolean to check if the value has been set.
func (o *PtwParameters1) GetOperationModeOk() (*OperationMode, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OperationMode, true
}

// SetOperationMode sets field value
func (o *PtwParameters1) SetOperationMode(v OperationMode) {
	o.OperationMode = v
}

// GetPtwValue returns the PtwValue field value
func (o *PtwParameters1) GetPtwValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PtwValue
}

// GetPtwValueOk returns a tuple with the PtwValue field value
// and a boolean to check if the value has been set.
func (o *PtwParameters1) GetPtwValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PtwValue, true
}

// SetPtwValue sets field value
func (o *PtwParameters1) SetPtwValue(v string) {
	o.PtwValue = v
}

func (o PtwParameters1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["operationMode"] = o.OperationMode
	}
	if true {
		toSerialize["ptwValue"] = o.PtwValue
	}
	return json.Marshal(toSerialize)
}

type NullablePtwParameters1 struct {
	value *PtwParameters1
	isSet bool
}

func (v NullablePtwParameters1) Get() *PtwParameters1 {
	return v.value
}

func (v *NullablePtwParameters1) Set(val *PtwParameters1) {
	v.value = val
	v.isSet = true
}

func (v NullablePtwParameters1) IsSet() bool {
	return v.isSet
}

func (v *NullablePtwParameters1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePtwParameters1(val *PtwParameters1) *NullablePtwParameters1 {
	return &NullablePtwParameters1{value: val, isSet: true}
}

func (v NullablePtwParameters1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePtwParameters1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


