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

// EUtranFrequencySingleAllOf struct for EUtranFrequencySingleAllOf
type EUtranFrequencySingleAllOf struct {
	Attributes *EUtranFrequencySingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewEUtranFrequencySingleAllOf instantiates a new EUtranFrequencySingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEUtranFrequencySingleAllOf() *EUtranFrequencySingleAllOf {
	this := EUtranFrequencySingleAllOf{}
	return &this
}

// NewEUtranFrequencySingleAllOfWithDefaults instantiates a new EUtranFrequencySingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEUtranFrequencySingleAllOfWithDefaults() *EUtranFrequencySingleAllOf {
	this := EUtranFrequencySingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EUtranFrequencySingleAllOf) GetAttributes() EUtranFrequencySingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret EUtranFrequencySingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EUtranFrequencySingleAllOf) GetAttributesOk() (*EUtranFrequencySingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EUtranFrequencySingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given EUtranFrequencySingleAllOfAttributes and assigns it to the Attributes field.
func (o *EUtranFrequencySingleAllOf) SetAttributes(v EUtranFrequencySingleAllOfAttributes) {
	o.Attributes = &v
}

func (o EUtranFrequencySingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableEUtranFrequencySingleAllOf struct {
	value *EUtranFrequencySingleAllOf
	isSet bool
}

func (v NullableEUtranFrequencySingleAllOf) Get() *EUtranFrequencySingleAllOf {
	return v.value
}

func (v *NullableEUtranFrequencySingleAllOf) Set(val *EUtranFrequencySingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEUtranFrequencySingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEUtranFrequencySingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEUtranFrequencySingleAllOf(val *EUtranFrequencySingleAllOf) *NullableEUtranFrequencySingleAllOf {
	return &NullableEUtranFrequencySingleAllOf{value: val, isSet: true}
}

func (v NullableEUtranFrequencySingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEUtranFrequencySingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


