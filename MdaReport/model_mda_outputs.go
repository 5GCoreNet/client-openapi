/*
MDA Report

OAS 3.0.1 specification of the MDA Report © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaReport

import (
	"encoding/json"
)

// MDAOutputs struct for MDAOutputs
type MDAOutputs struct {
	MDAType *string `json:"mDAType,omitempty"`
	MdaOutputList []MDAOutputEntry `json:"mdaOutputList,omitempty"`
	MDARequestRef *string `json:"mDARequestRef,omitempty"`
}

// NewMDAOutputs instantiates a new MDAOutputs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDAOutputs() *MDAOutputs {
	this := MDAOutputs{}
	return &this
}

// NewMDAOutputsWithDefaults instantiates a new MDAOutputs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDAOutputsWithDefaults() *MDAOutputs {
	this := MDAOutputs{}
	return &this
}

// GetMDAType returns the MDAType field value if set, zero value otherwise.
func (o *MDAOutputs) GetMDAType() string {
	if o == nil || isNil(o.MDAType) {
		var ret string
		return ret
	}
	return *o.MDAType
}

// GetMDATypeOk returns a tuple with the MDAType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputs) GetMDATypeOk() (*string, bool) {
	if o == nil || isNil(o.MDAType) {
    return nil, false
	}
	return o.MDAType, true
}

// HasMDAType returns a boolean if a field has been set.
func (o *MDAOutputs) HasMDAType() bool {
	if o != nil && !isNil(o.MDAType) {
		return true
	}

	return false
}

// SetMDAType gets a reference to the given string and assigns it to the MDAType field.
func (o *MDAOutputs) SetMDAType(v string) {
	o.MDAType = &v
}

// GetMdaOutputList returns the MdaOutputList field value if set, zero value otherwise.
func (o *MDAOutputs) GetMdaOutputList() []MDAOutputEntry {
	if o == nil || isNil(o.MdaOutputList) {
		var ret []MDAOutputEntry
		return ret
	}
	return o.MdaOutputList
}

// GetMdaOutputListOk returns a tuple with the MdaOutputList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputs) GetMdaOutputListOk() ([]MDAOutputEntry, bool) {
	if o == nil || isNil(o.MdaOutputList) {
    return nil, false
	}
	return o.MdaOutputList, true
}

// HasMdaOutputList returns a boolean if a field has been set.
func (o *MDAOutputs) HasMdaOutputList() bool {
	if o != nil && !isNil(o.MdaOutputList) {
		return true
	}

	return false
}

// SetMdaOutputList gets a reference to the given []MDAOutputEntry and assigns it to the MdaOutputList field.
func (o *MDAOutputs) SetMdaOutputList(v []MDAOutputEntry) {
	o.MdaOutputList = v
}

// GetMDARequestRef returns the MDARequestRef field value if set, zero value otherwise.
func (o *MDAOutputs) GetMDARequestRef() string {
	if o == nil || isNil(o.MDARequestRef) {
		var ret string
		return ret
	}
	return *o.MDARequestRef
}

// GetMDARequestRefOk returns a tuple with the MDARequestRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDAOutputs) GetMDARequestRefOk() (*string, bool) {
	if o == nil || isNil(o.MDARequestRef) {
    return nil, false
	}
	return o.MDARequestRef, true
}

// HasMDARequestRef returns a boolean if a field has been set.
func (o *MDAOutputs) HasMDARequestRef() bool {
	if o != nil && !isNil(o.MDARequestRef) {
		return true
	}

	return false
}

// SetMDARequestRef gets a reference to the given string and assigns it to the MDARequestRef field.
func (o *MDAOutputs) SetMDARequestRef(v string) {
	o.MDARequestRef = &v
}

func (o MDAOutputs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MDAType) {
		toSerialize["mDAType"] = o.MDAType
	}
	if !isNil(o.MdaOutputList) {
		toSerialize["mdaOutputList"] = o.MdaOutputList
	}
	if !isNil(o.MDARequestRef) {
		toSerialize["mDARequestRef"] = o.MDARequestRef
	}
	return json.Marshal(toSerialize)
}

type NullableMDAOutputs struct {
	value *MDAOutputs
	isSet bool
}

func (v NullableMDAOutputs) Get() *MDAOutputs {
	return v.value
}

func (v *NullableMDAOutputs) Set(val *MDAOutputs) {
	v.value = val
	v.isSet = true
}

func (v NullableMDAOutputs) IsSet() bool {
	return v.isSet
}

func (v *NullableMDAOutputs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDAOutputs(val *MDAOutputs) *NullableMDAOutputs {
	return &NullableMDAOutputs{value: val, isSet: true}
}

func (v NullableMDAOutputs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDAOutputs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


