/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
)

// DESManagementFunctionSingleAllOf struct for DESManagementFunctionSingleAllOf
type DESManagementFunctionSingleAllOf struct {
	Attributes *DESManagementFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewDESManagementFunctionSingleAllOf instantiates a new DESManagementFunctionSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDESManagementFunctionSingleAllOf() *DESManagementFunctionSingleAllOf {
	this := DESManagementFunctionSingleAllOf{}
	return &this
}

// NewDESManagementFunctionSingleAllOfWithDefaults instantiates a new DESManagementFunctionSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDESManagementFunctionSingleAllOfWithDefaults() *DESManagementFunctionSingleAllOf {
	this := DESManagementFunctionSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DESManagementFunctionSingleAllOf) GetAttributes() DESManagementFunctionSingleAllOfAttributes {
	if o == nil || isNil(o.Attributes) {
		var ret DESManagementFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DESManagementFunctionSingleAllOf) GetAttributesOk() (*DESManagementFunctionSingleAllOfAttributes, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DESManagementFunctionSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given DESManagementFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *DESManagementFunctionSingleAllOf) SetAttributes(v DESManagementFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o DESManagementFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableDESManagementFunctionSingleAllOf struct {
	value *DESManagementFunctionSingleAllOf
	isSet bool
}

func (v NullableDESManagementFunctionSingleAllOf) Get() *DESManagementFunctionSingleAllOf {
	return v.value
}

func (v *NullableDESManagementFunctionSingleAllOf) Set(val *DESManagementFunctionSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDESManagementFunctionSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDESManagementFunctionSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDESManagementFunctionSingleAllOf(val *DESManagementFunctionSingleAllOf) *NullableDESManagementFunctionSingleAllOf {
	return &NullableDESManagementFunctionSingleAllOf{value: val, isSet: true}
}

func (v NullableDESManagementFunctionSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDESManagementFunctionSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


