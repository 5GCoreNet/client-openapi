/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package 5GcNrm

import (
	"encoding/json"
)

// ManagedElementSingleAllOf struct for ManagedElementSingleAllOf
type ManagedElementSingleAllOf struct {
	Attributes *ManagedElementAttr `json:"attributes,omitempty"`
}

// NewManagedElementSingleAllOf instantiates a new ManagedElementSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedElementSingleAllOf() *ManagedElementSingleAllOf {
	this := ManagedElementSingleAllOf{}
	return &this
}

// NewManagedElementSingleAllOfWithDefaults instantiates a new ManagedElementSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedElementSingleAllOfWithDefaults() *ManagedElementSingleAllOf {
	this := ManagedElementSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf) GetAttributes() ManagedElementAttr {
	if o == nil || isNil(o.Attributes) {
		var ret ManagedElementAttr
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf) GetAttributesOk() (*ManagedElementAttr, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagedElementAttr and assigns it to the Attributes field.
func (o *ManagedElementSingleAllOf) SetAttributes(v ManagedElementAttr) {
	o.Attributes = &v
}

func (o ManagedElementSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableManagedElementSingleAllOf struct {
	value *ManagedElementSingleAllOf
	isSet bool
}

func (v NullableManagedElementSingleAllOf) Get() *ManagedElementSingleAllOf {
	return v.value
}

func (v *NullableManagedElementSingleAllOf) Set(val *ManagedElementSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedElementSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedElementSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedElementSingleAllOf(val *ManagedElementSingleAllOf) *NullableManagedElementSingleAllOf {
	return &NullableManagedElementSingleAllOf{value: val, isSet: true}
}

func (v NullableManagedElementSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedElementSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


