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

// DLBOFunctionSingleAllOf struct for DLBOFunctionSingleAllOf
type DLBOFunctionSingleAllOf struct {
	Attributes *DLBOFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewDLBOFunctionSingleAllOf instantiates a new DLBOFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDLBOFunctionSingleAllOf() *DLBOFunctionSingleAllOf {
	this := DLBOFunctionSingleAllOf{}
	return &this
}

// NewDLBOFunctionSingleAllOfWithDefaults instantiates a new DLBOFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDLBOFunctionSingleAllOfWithDefaults() *DLBOFunctionSingleAllOf {
	this := DLBOFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DLBOFunctionSingleAllOf) GetAttributes() DLBOFunctionSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret DLBOFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLBOFunctionSingleAllOf) GetAttributesOk() (*DLBOFunctionSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DLBOFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given DLBOFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *DLBOFunctionSingleAllOf) SetAttributes(v DLBOFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o DLBOFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableDLBOFunctionSingleAllOf struct {
	value *DLBOFunctionSingleAllOf
	isSet bool
}

func (v NullableDLBOFunctionSingleAllOf) Get() *DLBOFunctionSingleAllOf {
	return v.value
}

func (v *NullableDLBOFunctionSingleAllOf) Set(val *DLBOFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDLBOFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDLBOFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDLBOFunctionSingleAllOf(val *DLBOFunctionSingleAllOf) *NullableDLBOFunctionSingleAllOf {
	return &NullableDLBOFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableDLBOFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDLBOFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


