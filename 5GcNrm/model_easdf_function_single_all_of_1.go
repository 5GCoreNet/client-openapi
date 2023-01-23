/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// EASDFFunctionSingleAllOf1 struct for EASDFFunctionSingleAllOf1
type EASDFFunctionSingleAllOf1 struct {
	EPN88 []EPN88Single `json:"EP_N88,omitempty"`
}

// NewEASDFFunctionSingleAllOf1 instantiates a new EASDFFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEASDFFunctionSingleAllOf1() *EASDFFunctionSingleAllOf1 {
	this := EASDFFunctionSingleAllOf1{}
	return &this
}

// NewEASDFFunctionSingleAllOf1WithDefaults instantiates a new EASDFFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEASDFFunctionSingleAllOf1WithDefaults() *EASDFFunctionSingleAllOf1 {
	this := EASDFFunctionSingleAllOf1{}
	return &this
}

// GetEPN88 returns the EPN88 field value if set, zero value otherwise.
func (o *EASDFFunctionSingleAllOf1) GetEPN88() []EPN88Single {
	if o == nil || isNil(o.EPN88) {
		var ret []EPN88Single
		return ret
	}
	return o.EPN88
}

// GetEPN88Ok returns a tuple with the EPN88 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASDFFunctionSingleAllOf1) GetEPN88Ok() ([]EPN88Single, bool) {
	if o == nil || isNil(o.EPN88) {
    return nil, false
	}
	return o.EPN88, true
}

// HasEPN88 returns a boolean if a field has been set.
func (o *EASDFFunctionSingleAllOf1) HasEPN88() bool {
	if o != nil && !isNil(o.EPN88) {
		return true
	}

	return false
}

// SetEPN88 gets a reference to the given []EPN88Single and assigns it to the EPN88 field.
func (o *EASDFFunctionSingleAllOf1) SetEPN88(v []EPN88Single) {
	o.EPN88 = v
}

func (o EASDFFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EPN88) {
		toSerialize["EP_N88"] = o.EPN88
	}
	return json.Marshal(toSerialize)
}

type NullableEASDFFunctionSingleAllOf1 struct {
	value *EASDFFunctionSingleAllOf1
	isSet bool
}

func (v NullableEASDFFunctionSingleAllOf1) Get() *EASDFFunctionSingleAllOf1 {
	return v.value
}

func (v *NullableEASDFFunctionSingleAllOf1) Set(val *EASDFFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableEASDFFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableEASDFFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEASDFFunctionSingleAllOf1(val *EASDFFunctionSingleAllOf1) *NullableEASDFFunctionSingleAllOf1 {
	return &NullableEASDFFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullableEASDFFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEASDFFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


