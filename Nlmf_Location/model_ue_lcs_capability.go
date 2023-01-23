/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Location

import (
	"encoding/json"
)

// UeLcsCapability Indicates the LCS capability supported by the UE..
type UeLcsCapability struct {
	LppSupport *bool `json:"lppSupport,omitempty"`
	CiotOptimisation *bool `json:"ciotOptimisation,omitempty"`
}

// NewUeLcsCapability instantiates a new UeLcsCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeLcsCapability() *UeLcsCapability {
	this := UeLcsCapability{}
	var lppSupport bool = true
	this.LppSupport = &lppSupport
	var ciotOptimisation bool = false
	this.CiotOptimisation = &ciotOptimisation
	return &this
}

// NewUeLcsCapabilityWithDefaults instantiates a new UeLcsCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeLcsCapabilityWithDefaults() *UeLcsCapability {
	this := UeLcsCapability{}
	var lppSupport bool = true
	this.LppSupport = &lppSupport
	var ciotOptimisation bool = false
	this.CiotOptimisation = &ciotOptimisation
	return &this
}

// GetLppSupport returns the LppSupport field value if set, zero value otherwise.
func (o *UeLcsCapability) GetLppSupport() bool {
	if o == nil || isNil(o.LppSupport) {
		var ret bool
		return ret
	}
	return *o.LppSupport
}

// GetLppSupportOk returns a tuple with the LppSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeLcsCapability) GetLppSupportOk() (*bool, bool) {
	if o == nil || isNil(o.LppSupport) {
    return nil, false
	}
	return o.LppSupport, true
}

// HasLppSupport returns a boolean if a field has been set.
func (o *UeLcsCapability) HasLppSupport() bool {
	if o != nil && !isNil(o.LppSupport) {
		return true
	}

	return false
}

// SetLppSupport gets a reference to the given bool and assigns it to the LppSupport field.
func (o *UeLcsCapability) SetLppSupport(v bool) {
	o.LppSupport = &v
}

// GetCiotOptimisation returns the CiotOptimisation field value if set, zero value otherwise.
func (o *UeLcsCapability) GetCiotOptimisation() bool {
	if o == nil || isNil(o.CiotOptimisation) {
		var ret bool
		return ret
	}
	return *o.CiotOptimisation
}

// GetCiotOptimisationOk returns a tuple with the CiotOptimisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeLcsCapability) GetCiotOptimisationOk() (*bool, bool) {
	if o == nil || isNil(o.CiotOptimisation) {
    return nil, false
	}
	return o.CiotOptimisation, true
}

// HasCiotOptimisation returns a boolean if a field has been set.
func (o *UeLcsCapability) HasCiotOptimisation() bool {
	if o != nil && !isNil(o.CiotOptimisation) {
		return true
	}

	return false
}

// SetCiotOptimisation gets a reference to the given bool and assigns it to the CiotOptimisation field.
func (o *UeLcsCapability) SetCiotOptimisation(v bool) {
	o.CiotOptimisation = &v
}

func (o UeLcsCapability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LppSupport) {
		toSerialize["lppSupport"] = o.LppSupport
	}
	if !isNil(o.CiotOptimisation) {
		toSerialize["ciotOptimisation"] = o.CiotOptimisation
	}
	return json.Marshal(toSerialize)
}

type NullableUeLcsCapability struct {
	value *UeLcsCapability
	isSet bool
}

func (v NullableUeLcsCapability) Get() *UeLcsCapability {
	return v.value
}

func (v *NullableUeLcsCapability) Set(val *UeLcsCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableUeLcsCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableUeLcsCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeLcsCapability(val *UeLcsCapability) *NullableUeLcsCapability {
	return &NullableUeLcsCapability{value: val, isSet: true}
}

func (v NullableUeLcsCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeLcsCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


