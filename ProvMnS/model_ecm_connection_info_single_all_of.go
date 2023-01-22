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

// EcmConnectionInfoSingleAllOf struct for EcmConnectionInfoSingleAllOf
type EcmConnectionInfoSingleAllOf struct {
	Attributes *interface{} `json:"attributes,omitempty"`
}

// NewEcmConnectionInfoSingleAllOf instantiates a new EcmConnectionInfoSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcmConnectionInfoSingleAllOf() *EcmConnectionInfoSingleAllOf {
	this := EcmConnectionInfoSingleAllOf{}
	return &this
}

// NewEcmConnectionInfoSingleAllOfWithDefaults instantiates a new EcmConnectionInfoSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcmConnectionInfoSingleAllOfWithDefaults() *EcmConnectionInfoSingleAllOf {
	this := EcmConnectionInfoSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EcmConnectionInfoSingleAllOf) GetAttributes() interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcmConnectionInfoSingleAllOf) GetAttributesOk() (*interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EcmConnectionInfoSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *EcmConnectionInfoSingleAllOf) SetAttributes(v interface{}) {
	o.Attributes = &v
}

func (o EcmConnectionInfoSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableEcmConnectionInfoSingleAllOf struct {
	value *EcmConnectionInfoSingleAllOf
	isSet bool
}

func (v NullableEcmConnectionInfoSingleAllOf) Get() *EcmConnectionInfoSingleAllOf {
	return v.value
}

func (v *NullableEcmConnectionInfoSingleAllOf) Set(val *EcmConnectionInfoSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEcmConnectionInfoSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEcmConnectionInfoSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcmConnectionInfoSingleAllOf(val *EcmConnectionInfoSingleAllOf) *NullableEcmConnectionInfoSingleAllOf {
	return &NullableEcmConnectionInfoSingleAllOf{value: val, isSet: true}
}

func (v NullableEcmConnectionInfoSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcmConnectionInfoSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


