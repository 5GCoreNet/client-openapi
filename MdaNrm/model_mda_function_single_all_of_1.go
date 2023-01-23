/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
)

// MDAFunctionSingleAllOf1 struct for MDAFunctionSingleAllOf1
type MDAFunctionSingleAllOf1 struct {
	MDARequest []MDARequestSingle `json:"MDARequest,omitempty"`
}

// NewMDAFunctionSingleAllOf1 instantiates a new MDAFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAFunctionSingleAllOf1() *MDAFunctionSingleAllOf1 {
	this := MDAFunctionSingleAllOf1{}
	return &this
}

// NewMDAFunctionSingleAllOf1WithDefaults instantiates a new MDAFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAFunctionSingleAllOf1WithDefaults() *MDAFunctionSingleAllOf1 {
	this := MDAFunctionSingleAllOf1{}
	return &this
}

// GetMDARequest returns the MDARequest field value if set, zero value otherwise.
func (o *MDAFunctionSingleAllOf1) GetMDARequest() []MDARequestSingle {
	if o == nil || isNil(o.MDARequest) {
		var ret []MDARequestSingle
		return ret
	}
	return o.MDARequest
}

// GetMDARequestOk returns a tuple with the MDARequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAFunctionSingleAllOf1) GetMDARequestOk() ([]MDARequestSingle, bool) {
	if o == nil || isNil(o.MDARequest) {
    return nil, false
	}
	return o.MDARequest, true
}

// HasMDARequest returns a boolean if a field has been set.
func (o *MDAFunctionSingleAllOf1) HasMDARequest() bool {
	if o != nil && !isNil(o.MDARequest) {
		return true
	}

	return false
}

// SetMDARequest gets a reference to the given []MDARequestSingle and assigns it to the MDARequest field.
func (o *MDAFunctionSingleAllOf1) SetMDARequest(v []MDARequestSingle) {
	o.MDARequest = v
}

func (o MDAFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MDARequest) {
		toSerialize["MDARequest"] = o.MDARequest
	}
	return json.Marshal(toSerialize)
}

type NullableMDAFunctionSingleAllOf1 struct {
	value *MDAFunctionSingleAllOf1
	isSet bool
}

func (v NullableMDAFunctionSingleAllOf1) Get() *MDAFunctionSingleAllOf1 {
	return v.value
}

func (v *NullableMDAFunctionSingleAllOf1) Set(val *MDAFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAFunctionSingleAllOf1(val *MDAFunctionSingleAllOf1) *NullableMDAFunctionSingleAllOf1 {
	return &NullableMDAFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullableMDAFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


