/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// TermDensity struct for TermDensity
type TermDensity struct {
	ServAttrCom *ServAttrCom `json:"servAttrCom,omitempty"`
	Density *int32 `json:"density,omitempty"`
}

// NewTermDensity instantiates a new TermDensity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTermDensity() *TermDensity {
	this := TermDensity{}
	return &this
}

// NewTermDensityWithDefaults instantiates a new TermDensity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTermDensityWithDefaults() *TermDensity {
	this := TermDensity{}
	return &this
}

// GetServAttrCom returns the ServAttrCom field value if set, zero value otherwise.
func (o *TermDensity) GetServAttrCom() ServAttrCom {
	if o == nil || isNil(o.ServAttrCom) {
		var ret ServAttrCom
		return ret
	}
	return *o.ServAttrCom
}

// GetServAttrComOk returns a tuple with the ServAttrCom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermDensity) GetServAttrComOk() (*ServAttrCom, bool) {
	if o == nil || isNil(o.ServAttrCom) {
    return nil, false
	}
	return o.ServAttrCom, true
}

// HasServAttrCom returns a boolean if a field has been set.
func (o *TermDensity) HasServAttrCom() bool {
	if o != nil && !isNil(o.ServAttrCom) {
		return true
	}

	return false
}

// SetServAttrCom gets a reference to the given ServAttrCom and assigns it to the ServAttrCom field.
func (o *TermDensity) SetServAttrCom(v ServAttrCom) {
	o.ServAttrCom = &v
}

// GetDensity returns the Density field value if set, zero value otherwise.
func (o *TermDensity) GetDensity() int32 {
	if o == nil || isNil(o.Density) {
		var ret int32
		return ret
	}
	return *o.Density
}

// GetDensityOk returns a tuple with the Density field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermDensity) GetDensityOk() (*int32, bool) {
	if o == nil || isNil(o.Density) {
    return nil, false
	}
	return o.Density, true
}

// HasDensity returns a boolean if a field has been set.
func (o *TermDensity) HasDensity() bool {
	if o != nil && !isNil(o.Density) {
		return true
	}

	return false
}

// SetDensity gets a reference to the given int32 and assigns it to the Density field.
func (o *TermDensity) SetDensity(v int32) {
	o.Density = &v
}

func (o TermDensity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServAttrCom) {
		toSerialize["servAttrCom"] = o.ServAttrCom
	}
	if !isNil(o.Density) {
		toSerialize["density"] = o.Density
	}
	return json.Marshal(toSerialize)
}

type NullableTermDensity struct {
	value *TermDensity
	isSet bool
}

func (v NullableTermDensity) Get() *TermDensity {
	return v.value
}

func (v *NullableTermDensity) Set(val *TermDensity) {
	v.value = val
	v.isSet = true
}

func (v NullableTermDensity) IsSet() bool {
	return v.isSet
}

func (v *NullableTermDensity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermDensity(val *TermDensity) *NullableTermDensity {
	return &NullableTermDensity{value: val, isSet: true}
}

func (v NullableTermDensity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTermDensity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


